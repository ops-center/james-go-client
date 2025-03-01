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

// syncOps represents individual synchronization operations.
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

// SyncOpsSet represents a set of synchronization operations combined using bitwise operations.
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
type SyncFilter interface{ SkipSync() bool }

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
	if ok && syncFilter.SkipSync() {
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
		return w.sync(new.GetObjectIdentifier(), new.GetSyncOptions().OnDelete)
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
			err = w.RemoveObjectAlias(obj)
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
