package james

import (
	"fmt"
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

	bui := object.GetBoundedUserIdentity()
	if bui.OwnerName != "" {
		addr = fmt.Sprintf("%s&%s-%s", addr, "ownername", bui.OwnerName)
	}
	if bui.OwnerID != "" {
		addr = fmt.Sprintf("%s&%s-%s", addr, "ownerid", bui.OwnerID)
	}
	if bui.OwnerType != "" {
		addr = fmt.Sprintf("%s&%s-%s", addr, "ownertype", bui.OwnerType)
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

	if object.HasParentObject() {
		parentObject, err := object.GetParentObject()
		if err != nil {
			return "", err
		}
		return genObjectTreeAddr(parentObject, addr)
	}

	return addr, nil
}

func isEmptyString(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}
