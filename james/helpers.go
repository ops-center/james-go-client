package james

import (
	"fmt"
	"github.com/google/uuid"
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

func getObjectIdentifierFromObjectInterface(object Object) *ObjectIdentifier {
	if object == nil {
		return nil
	}

	return &ObjectIdentifier{
		ObjectName:          object.GetName(),
		ObjectUniqueID:      object.GetUniqueID(),
		ObjectType:          object.GetType(),
		IsGroupType:         object.IsGroup(),
		ParentObject:        getObjectIdentifierFromObjectInterface(object),
		AdditionalClaims:    object.AdditionalTokenClaims(),
		BoundedUserIdentity: object.GetBoundedUserIdentity(),
	}
}

func GetObjectIdentifierFromObjectInterface(object Object) *ObjectIdentifier {
	return getObjectIdentifierFromObjectInterface(object)
}

func getGroupAndAssociatedMembersIdentifier(object Object, childObjects []Object) ([]GroupAndAssociatedMembersIdentifier, error) {
	if object == nil {
		return nil, nil
	}

	result := GroupAndAssociatedMembersIdentifier{
		Group: getObjectIdentifierFromObjectInterface(object),
	}
	for _, childObj := range childObjects {
		result.Members = append(result.Members, getObjectIdentifierFromObjectInterface(childObj))
	}

	parentObject, err := object.GetParentObject()
	if err != nil {
		return nil, err
	}
	resultOfParentObject, err := getGroupAndAssociatedMembersIdentifier(parentObject, append(childObjects, object))
	if err != nil {
		return nil, err
	}

	return append(resultOfParentObject, result), nil
}

func GetGroupAndAssociatedMembersIdentifier(object Object) ([]GroupAndAssociatedMembersIdentifier, error) {
	return getGroupAndAssociatedMembersIdentifier(object, nil)
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
