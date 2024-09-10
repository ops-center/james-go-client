package inbox

import (
	"fmt"
	"testing"
)

func TestGetGroupAndAssociatedMemberIdentifier(t *testing.T) {
	jamesIdentifier := ObjectIdentifier{
		ObjectName:     "workload",
		ObjectUniqueID: "1",
		ObjectType:     DbType,
		IsGroupType:    true,
		ParentObject: &ObjectIdentifier{
			ObjectName:     "namespace",
			ObjectUniqueID: "2",
			ObjectType:     DbType,
			IsGroupType:    true,
			ParentObject: &ObjectIdentifier{
				ObjectName:     "cluster",
				ObjectUniqueID: "3",
				ObjectType:     DbType,
				IsGroupType:    true,
			},
		},
	}

	results, err := getGroupAndAssociatedMemberIdentifier(jamesIdentifier, nil)
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 2 {
		t.Error(fmt.Errorf("length of result doesn't match: expected: %v, got: %v", 3, len(results)))
		return
	}

	for _, result := range results {
		fmt.Printf("group: %s, member: %s\n", result.Group.ObjectName, result.Member.ObjectName)
		if !result.Member.HasParentObject() {
			t.Errorf("member %s does not have parent object", result.Member.ObjectName)
			return
		}
		if result.Member.ParentObject.GetUniqueID() != result.Group.GetUniqueID() {
			t.Errorf("member parent reference does not point to group object")
			return
		}
	}
}
