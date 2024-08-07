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

	addr = fmt.Sprintf("%s&%s", addr, object.GetBoundedUserIdentity().String())

	if len(addr) > MaxEmailLengthJames {
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

	if len(addr) > MaxEmailLengthJames {
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
