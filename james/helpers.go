package james

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

func generateObjectAddr(object Object) (string, error) {
	addr, err := genObjectTreeAddr(object, "")
	if err != nil {
		return "", err
	}

	if object.IsGroup() {
		addr = fmt.Sprintf("%s.x&type-grp", addr)
	} else {
		addr = fmt.Sprintf("%s.x&type-acc", addr)
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

func genObjectTreeAddr(object Object, addr string) (string, error) {
	if object == nil {
		return addr, nil
	}
	if isEmptyString(addr) {
		addr = fmt.Sprintf("%s$%s&%s", object.GetName(), object.GetUniqueID(), object.GetType())
	} else {
		addr = fmt.Sprintf("%s.%s$%s&%s", addr, object.GetName(), object.GetUniqueID(), object.GetType())
	}

	if len(addr) > MaxEmailLength {
		return "", ErrMaxEmailLengthExceeded
	}

	if object.HasParentObject() {
		parentObject, err := object.GetParentObject()
		if err != nil {
			return "", err
		}
		return genObjectTreeAddr(parentObject, addr)
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

func GetGroupAndAssociatedMembersIdentifier(object Object) ([]GroupAndAssociatedMembersIdentifier, error) {
	return getGroupAndAssociatedMembersIdentifier(object, nil)
}

func getGroupAndAssociatedMembersIdentifier(object Object, childObject Object) ([]GroupAndAssociatedMembersIdentifier, error) {
	if object == nil || (reflect.ValueOf(object).Kind() == reflect.Ptr && reflect.ValueOf(object).IsNil()) {
		return nil, nil
	}

	objectIdentifier, err := getObjectIdentifierFromObjectInterface(object)
	if err != nil {
		return nil, err
	}
	member, err := getObjectIdentifierFromObjectInterface(childObject)
	if err != nil {
		return nil, err
	}
	result := GroupAndAssociatedMembersIdentifier{
		Group: objectIdentifier,
		Members: func() []*ObjectIdentifier {
			if member != nil {
				return []*ObjectIdentifier{member}
			}
			return nil
		}(),
	}

	resultOfParentObject, err := getGroupAndAssociatedMembersIdentifier(objectIdentifier.ParentObject, object)
	if err != nil {
		return nil, err
	}

	return append(resultOfParentObject, result), nil
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
