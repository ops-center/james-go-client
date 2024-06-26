package james

import (
	"fmt"
	"sort"
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

	type keyValPair struct {
		key, val string
	}
	bui := []keyValPair{}

	for key, value := range object.GetBoundedUserIdentity() {
		bui = append(bui, keyValPair{key, value})
	}

	sort.Slice(bui, func(i, j int) bool {
		return bui[i].key < bui[j].key
	})

	for _, kv := range bui {
		addr = fmt.Sprintf("%s&%s-%s", addr, kv.key, kv.val)
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
