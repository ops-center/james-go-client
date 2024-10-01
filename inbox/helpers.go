package inbox

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/uuid"
)

func generateObjectAddr(object Object) (string, error) {
	addr, err := generateAddr(object, "")
	if err != nil {
		return "", err
	}

	boundedUserIdentity := object.GetBoundedUserIdentity()
	if boundedUserIdentity != nil {
		addr = fmt.Sprintf("%s&%s", addr, boundedUserIdentity.String())
	}

	if len(addr) > MaxEmailLength {
		return "", ErrMaxEmailLengthExceeded
	}
	return fmt.Sprintf("%s@%s", addr, GlobalMailDomain), nil
}

func generateAddr(object Object, addr string) (string, error) {
	if object == nil {
		return addr, nil
	}
	if isEmptyString(addr) {
		addr = fmt.Sprintf("%s&%s$%s", object.GetType(), object.GetName(), object.GetUniqueID())
	} else {
		addr = fmt.Sprintf("%s.%s&%s$%s", addr, object.GetType(), object.GetName(), object.GetUniqueID())
	}

	if len(addr) > MaxEmailLength {
		return "", ErrMaxEmailLengthExceeded
	}

	if object.HasParentObject() {
		parentObject, err := object.GetParentObject()
		if err != nil {
			return "", err
		}
		return generateAddr(parentObject, addr)
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
		ObjectName:          object.GetName(),
		ObjectUniqueID:      object.GetUniqueID(),
		ObjectType:          object.GetType(),
		IsGroupType:         object.IsGroup(),
		ParentObject:        parentObjectIdentifier,
		AdditionalClaims:    object.AdditionalTokenClaims(),
		BoundedUserIdentity: object.GetBoundedUserIdentity(),
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
