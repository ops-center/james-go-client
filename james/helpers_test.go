package james

import (
	"fmt"
	"testing"
)

func TestGetGroupAndAssociatedMembersIdentifier(t *testing.T) {
	jamesIdentifier := &ObjectIdentifier{
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

	results, err := GetGroupAndAssociatedMembersIdentifier(jamesIdentifier)
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 2 {
		t.Error(fmt.Errorf("length of result doesn't match: expected: %v, got: %v", 2, len(results)))
		return
	}

	for _, result := range results {
		/*	if idx == 2 {
			if result.Members != nil {
				t.Error(fmt.Errorf("length of members of group `%v` doesn't match, expected: %v, got: %v", result.Group.ObjectName, 0, len(result.Members)))
			}
			continue
		}*/
		if len(result.Members) != 1 {
			t.Error(fmt.Errorf("length of members of group `%v` doesn't match, expected: %v, got: %v", result.Group.ObjectName, 1, len(result.Members)))
			return
		}
		fmt.Printf("Group: %s\nMember: %s\n\n", result.Group.ObjectName, result.Members[0].ObjectName)
	}

}
