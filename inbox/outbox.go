package inbox

import (
	"fmt"
	"reflect"
	"strings"

	"xorm.io/xorm/names"
)

// SyncableEntity defines the interface for entities that can be synchronized.
type SyncableEntity interface {
	GetObjectIdentifier() ObjectIdentifier
	GetSyncOptions() *SyncOptions
	TableName() string
}

// SyncOptions defines the synchronization operations to be performed for insert and delete actions.
type SyncOptions struct {
	OnInsert SyncOpsSet
	OnDelete SyncOpsSet
}

// syncOps represents individual synchronization operation flags using bitmask semantics.
// Each constant represents a distinct synchronization action that can be performed.
//
// Example synchronization operations:
//   - Creating/updating/deleting user accounts
//   - Managing account aliases
//   - Modifying group memberships
type syncOps uint16

const (
	SyncOpsCreateAccount syncOps = 1 << iota
	SyncOpsDeleteAccount
	SyncOpsUpdateAccount
	SyncOpsCreateAccountAlias
	SyncOpsDeleteAccountAlias
	SyncOpsAddAsGroupMember
	SyncOpsRemoveAsGroupMember
)

var syncOpsRegistry = []syncOps{
	SyncOpsCreateAccount,
	SyncOpsDeleteAccount,
	SyncOpsUpdateAccount,
	SyncOpsCreateAccountAlias,
	SyncOpsDeleteAccountAlias,
	SyncOpsAddAsGroupMember,
	SyncOpsRemoveAsGroupMember,
}

// SyncOpsSet represents a combination of synchronization operations using bitwise flags.
type SyncOpsSet uint16

func NewSyncOpsSet(ops ...syncOps) SyncOpsSet {
	var result SyncOpsSet
	for _, op := range ops {
		result |= SyncOpsSet(op)
	}

	return result
}

func (s SyncOpsSet) Has(op syncOps) bool {
	return s&SyncOpsSet(op) != 0
}

func (s SyncOpsSet) RetrieveSyncOps() []syncOps {
	var result []syncOps
	for _, op := range syncOpsRegistry {
		if s.Has(op) {
			result = append(result, op)
		}
	}
	return result
}

func (s SyncOpsSet) Add(ops ...syncOps) SyncOpsSet {
	return NewSyncOpsSet(append(s.RetrieveSyncOps(), ops...)...)
}

// DeriveCleanupOps calculates required cleanup operations based on registered sync operations.
// This is the primary method for determining which data domains need truncation before re-sync.
func (s SyncOpsSet) DeriveCleanupOps() CleanupOpsSet {
	var result CleanupOpsSet

	for syncOp, cleanupOp := range syncCleanupOpsMappingRegistry {
		if s.Has(syncOp) {
			result |= CleanupOpsSet(cleanupOp)
		}
	}

	return result
}

// cleanupOps represents data domain cleanup requirements before full resynchronization.
// These flags ensure stale data is purged from downstream systems when performing
// full consistency re-syncs. Each flag corresponds to a distinct data domain that
// needs truncation based on previously executed synchronization operations.
//
// Cleanup occurs BEFORE re-sync to:
//   - Remove potentially inconsistent existing state
//   - Handle schema/downstream storage changes
//   - Reset synchronization baselines
//
// Example cleanup triggers:
//
//	CleanupOpsDeleteAccounts:    Required when any account-related sync operations
//	                              (create/update/delete) were previously executed,
//	                              ensures old account data is purged before full re-sync
//
//	CleanupOpsDeleteGroups:      Needed after group membership changes to reset
//	                              relationship mappings before rebuilding
//
//	CleanupOpsDeleteAddressAliases: Mandatory when alias operations were performed,
//	                              prevents duplicate/conflicting aliases
type cleanupOps uint16

const (
	CleanupOpsDeleteAccounts cleanupOps = 1 << iota
	CleanupOpsDeleteAddressAliases
)

var cleanupOpsRegistry = []cleanupOps{
	CleanupOpsDeleteAccounts,
	CleanupOpsDeleteAddressAliases,
}

// syncCleanupOpsMappingRegistry defines the cleanup requirements for each synchronization operation.
// This mapping determines which cleanup operations must be performed after specific sync ops.
//
// Note: When adding new sync operations, update this registry to specify their cleanup requirements.
var syncCleanupOpsMappingRegistry = map[syncOps]cleanupOps{
	// Account operations require full account table cleanup
	SyncOpsCreateAccount: CleanupOpsDeleteAccounts,
	SyncOpsDeleteAccount: CleanupOpsDeleteAccounts,
	SyncOpsUpdateAccount: CleanupOpsDeleteAccounts,

	// Alias operations require alias table reset
	SyncOpsCreateAccountAlias: CleanupOpsDeleteAddressAliases,
	SyncOpsDeleteAccountAlias: CleanupOpsDeleteAddressAliases,

	// Group membership changes require group table refresh
	SyncOpsAddAsGroupMember:    CleanupOpsDeleteAccounts,
	SyncOpsRemoveAsGroupMember: CleanupOpsDeleteAccounts,
}

// CleanupOpsSet represents the aggregated cleanup requirements for a resynchronization operation.
// The set is constructed by combining cleanup obligations from all synchronization operations
// that were previously executed, as defined in the syncCleanupOpsRegistry.
//
// Usage Pattern:
//
//	// After identifying required sync operations
//	syncSet := NewSyncOpsSet(SyncOpsCreateAccount, SyncOpsAddAsGroupMember)
//
//	// Convert to clean up requirements
//	cleanupSet := syncSet.DeriveCleanupOps()
//
//	// Execute cleanup
//	if cleanupSet.Has(CleanupOpsDeleteAccounts) {
//	    truncateAccountTables()
//	}
type CleanupOpsSet uint16

func NewCleanupOpsSet(ops ...cleanupOps) CleanupOpsSet {
	var result CleanupOpsSet
	for _, op := range ops {
		result |= CleanupOpsSet(op)
	}

	return result
}

func (s CleanupOpsSet) Has(op cleanupOps) bool {
	return s&CleanupOpsSet(op) != 0
}

// ErrModelNotRegistered represents an error when a table name is not registered in the ModelSyncRegistry.
// - Useful for identifying cases where a SyncableEntity has not been properly registered.
type ErrModelNotRegistered struct {
	TableName string // Name of the unregistered table
}

// Error returns a descriptive error message for ErrModelNotRegistered.
func (e *ErrModelNotRegistered) Error() string {
	return fmt.Sprintf("no model registered for table `%v`", e.TableName)
}

// SyncFilter defines entities that can conditionally skip synchronization.
type SyncFilter interface{ SkipInboxSvcSync(old SyncableEntity) bool }

type EntityRegistry struct {
	mapper names.Mapper
	reg    map[string]SyncableEntity
}

func NewSyncableEntityRegistry(mapper names.Mapper, entities ...SyncableEntity) *EntityRegistry {
	reg := &EntityRegistry{
		mapper: mapper,
		reg:    make(map[string]SyncableEntity),
	}
	reg.Register(entities...)

	return reg
}

func (r *EntityRegistry) Register(entities ...SyncableEntity) {
	for _, entity := range entities {
		if entity == nil {
			continue
		}
		r.reg[entity.TableName()] = entity
	}
}

func (r *EntityRegistry) RegisteredEntities() []SyncableEntity {
	var result []SyncableEntity
	for _, entity := range r.reg {
		result = append(result, entity)
	}

	return result
}

func (r *EntityRegistry) RegisteredEntityNames() []string {
	var result []string
	for _, entity := range r.reg {
		result = append(result, entity.TableName())
	}

	return result
}

func (r *EntityRegistry) IsRegistered(tableName string) bool {
	return r.reg[tableName] != nil
}

func (r *EntityRegistry) GetSyncableEntity(tableName string, data map[string]any) (SyncableEntity, error) {
	if len(data) == 0 {
		return nil, nil
	}
	modelProto, exists := r.reg[tableName]
	if !exists {
		return nil, &ErrModelNotRegistered{TableName: tableName}
	}
	modelType := reflect.TypeOf(modelProto).Elem()
	modelInstance := reflect.New(modelType).Interface()

	// Map column names to struct fields using the mapper
	for column, value := range data {
		fieldName := r.mapper.Table2Obj(column) // Convert column name to struct field name
		field := reflect.ValueOf(modelInstance).Elem().FieldByName(fieldName)

		if !field.IsValid() || !field.CanSet() {
			continue
		}

		// Ensure type compatibility before setting the value
		valueRef := reflect.ValueOf(value)
		if value != nil && valueRef.Type().ConvertibleTo(field.Type()) {
			field.Set(valueRef.Convert(field.Type()))
		}
	}

	se, ok := modelInstance.(SyncableEntity)
	if !ok {
		return nil, fmt.Errorf("model `%v` is not a SyncableEntity", tableName)
	}

	return se, nil
}

func (w *WebAdminClient) SyncToInboxService(new, old SyncableEntity, ops string) error {
	syncFilter, ok := new.(SyncFilter)
	if ok && syncFilter.SkipInboxSvcSync(old) {
		return nil
	}

	switch strings.ToLower(ops) {
	case "insert":
		return w.sync(new.GetObjectIdentifier(), new.GetSyncOptions().OnInsert)
	case "update":
		if err := w.sync(old.GetObjectIdentifier(), old.GetSyncOptions().OnDelete); err != nil {
			return err
		}
		return w.sync(new.GetObjectIdentifier(), new.GetSyncOptions().OnInsert)
	case "delete":
		return w.sync(old.GetObjectIdentifier(), old.GetSyncOptions().OnDelete)
	}

	return nil
}

func (w *WebAdminClient) sync(obj ObjectIdentifier, syncOpsSet SyncOpsSet) error {
	for _, op := range syncOpsRegistry {
		if !syncOpsSet.Has(op) {
			continue
		}
		var err error

		switch op {
		case SyncOpsCreateAccount:
			err = w.CreateObjectIfNotExists(obj)
		case SyncOpsDeleteAccount:
			err = w.DeleteObject(obj)
		case SyncOpsUpdateAccount:
			// TODO: (inbox) here will be needing two object
			// return w.UpdateObjectAddr()
		case SyncOpsCreateAccountAlias:
			err = w.AddObjectAlias(obj)
		case SyncOpsDeleteAccountAlias:
			err = w.RemoveAllObjectAliases(obj)
		case SyncOpsAddAsGroupMember:
			err = w.AddGroupMember(obj.ParentObject, obj)
		case SyncOpsRemoveAsGroupMember:
			err = w.RemoveGroupMember(obj.ParentObject, obj)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func (w *WebAdminClient) CleanupPreAppliedData(preAppliedSyncOps SyncOpsSet) error {
	requiredCleanupOpsSet := preAppliedSyncOps.DeriveCleanupOps()

	for _, ops := range cleanupOpsRegistry {
		if !requiredCleanupOpsSet.Has(ops) {
			continue
		}

		switch ops {
		case CleanupOpsDeleteAccounts:
			users, err := w.GetAllUsers()
			if err != nil {
				return fmt.Errorf("failed to get all users: %w", err)
			}
			if err = w.deleteMultipleAccounts(users); err != nil {
				return fmt.Errorf("failed to delete all accounts: %w", err)
			}

			groups, err := w.GetAllGroups()
			if err != nil {
				return fmt.Errorf("failed to get all groups: %w", err)
			}
			if err = w.deleteGroupAddresses(groups); err != nil {
				return fmt.Errorf("failed to delete group addresses: %w", err)
			}
		case CleanupOpsDeleteAddressAliases:
			addresses, err := w.ListAliasAddresses()
			if err != nil {
				return fmt.Errorf("failed to get all aliases: %w", err)
			}

			if err = w.deleteAllAddressAliasesOfAListOfUsers(addresses); err != nil {
				return fmt.Errorf("failed to remove all aliases: %w", err)
			}
		}
	}

	return nil
}
