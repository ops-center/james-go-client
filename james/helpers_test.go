package james

import (
	"fmt"
	"testing"
)

func TestGetGroupAndAssociatedMembersIdentifier(t *testing.T) {
	jamesIdentifier := ObjectIdentifier{
		ObjectName:     "test-obj",
		ObjectUniqueID: "1",
		ObjectType:     DbType,
		IsGroupType:    true,
		ParentObject: &ObjectIdentifier{
			ObjectName:     "test-obj-2",
			ObjectUniqueID: "2",
			ObjectType:     DbType,
			IsGroupType:    true,
			ParentObject: &ObjectIdentifier{
				ObjectName:     "test-obj-3",
				ObjectUniqueID: "3",
				ObjectType:     DbType,
				IsGroupType:    true,
			},
		},
	}

	results, err := getGroupAndAssociatedMembersIdentifier(jamesIdentifier, nil)
	if err != nil {
		t.Error(err)
	}

	if len(results) != 3 {
		t.Error(fmt.Errorf("length of result doesn't match: expected: %v, got: %v", 3, len(results)))
	}

	for idx, result := range results {
		if idx == 2 {
			if result.Members != nil {
				t.Error(fmt.Errorf("length of members of group `%v` doesn't match, expected: %v, got: %v", result.Group.ObjectName, 0, len(result.Members)))
			}
			continue
		}
		if len(result.Members) != 1 {
			t.Error(fmt.Errorf("length of members of group `%v` doesn't match, expected: %v, got: %v", result.Group.ObjectName, 1, len(result.Members)))
		}
	}
}
