package inbox

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

func GenerateObjectAddr(object Object, domain string) (string, error) {
	return generateObjectAddr(object, domain)
}

func generateObjectAddr(object Object, domain string) (string, error) {
	addr, err := GenerateBaseAddress(object)
	if err != nil {
		return "", err
	}

	resourceOwner := object.ResourceOwnerIdentity()
	if resourceOwner != nil {
		addr = fmt.Sprintf("%s&%s", addr, resourceOwner.String())
	}

	if len(addr) > MaxEmailLength {
		return "", ErrMaxEmailLengthExceeded
	}
	return fmt.Sprintf("%s@%s", addr, domain), nil
}

func GenerateBaseAddress(object Object) (string, error) {
	if object == nil {
		return "", errors.New("nil object reference")
	}

	if val := reflect.ValueOf(object); val.Kind() == reflect.Ptr && val.IsNil() {
		return "", errors.New("nil object pointer")
	}

	addr := fmt.Sprintf("%s$%s", object.GetType(), object.GetUniqueID())
	for object.HasParentObject() {
		parent, err := object.GetParentObject()
		if err != nil {
			return "", fmt.Errorf("parent resolution failed: %w", err)
		}

		if parent == nil {
			return "", errors.New("nil parent object")
		}

		addr = fmt.Sprintf("%s.%s$%s", addr, parent.GetType(), parent.GetUniqueID())
		object = parent
	}

	return addr, nil
}

func GetObjectIdentifierFromObjectInterface(object Object) (*ObjectIdentifier, error) {
	return getObjectIdentifierFromObjectInterface(object)
}

func getObjectIdentifierFromObjectInterface(object Object) (*ObjectIdentifier, error) {
	if object == nil || (reflect.ValueOf(object).Kind() == reflect.Ptr && reflect.ValueOf(object).IsNil()) {
		return nil, nil
	}

	var parentObjectIdentifier *ObjectIdentifier = nil

	if object.HasParentObject() {
		parentObject, err := object.GetParentObject()
		if err != nil {
			return nil, err
		}

		parentObjectIdentifier, err = getObjectIdentifierFromObjectInterface(parentObject)
		if err != nil {
			return nil, err
		}
	}

	return &ObjectIdentifier{
		ObjectUniqueID:   object.GetUniqueID(),
		ObjectType:       object.GetType(),
		IsGroupType:      object.IsGroup(),
		ParentObject:     parentObjectIdentifier,
		AdditionalClaims: object.AdditionalTokenClaims(),
		ResourceOwner:    object.ResourceOwnerIdentity(),
	}, nil
}

func GetGroupAndAssociatedMemberIdentifier(object Object) ([]GroupAndAssociatedMemberIdentifier, error) {
	if object == nil || (reflect.ValueOf(object).Kind() == reflect.Ptr && reflect.ValueOf(object).IsNil()) {
		return nil, nil
	}

	groups := make([]GroupAndAssociatedMemberIdentifier, 0)
	obj, err := getObjectIdentifierFromObjectInterface(object)
	if err != nil {
		return nil, err
	}

	for ; obj != nil && obj.HasParentObject(); obj = obj.ParentObject {
		groups = append(groups, GroupAndAssociatedMemberIdentifier{
			Group:  *obj,
			Member: *obj.ParentObject,
		})
	}

	return groups, nil
}

func generateRandomPassword() (string, error) {
	pass, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return pass.String(), nil
}

func isEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

func isValidDomainFormat(domain string) bool {
	re := regexp.MustCompile(`^(?i)[a-z0-9-]{1,63}(\.[a-z0-9-]{1,63})+\.?$`)
	return re.MatchString(domain)
}
