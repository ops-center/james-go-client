package inbox

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetGroupAndAssociatedMemberIdentifier(t *testing.T) {
	jamesIdentifier := ObjectIdentifier{
		ObjectName:     "database",
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

	expectedParent := make(map[string]string)
	expectedParent["database"] = "namespace"
	expectedParent["namespace"] = "cluster"

	results, err := GetGroupAndAssociatedMemberIdentifier(jamesIdentifier)
	if err != nil {
		t.Error(err)
		return
	}

	if len(results) != 2 {
		t.Error(fmt.Errorf("length of result doesn't match: expected: %v, got: %v", 3, len(results)))
		return
	}

	prettyJSON, err := json.MarshalIndent(results, "  ", "   ")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("%v\n", string(prettyJSON))

	for _, g := range results {
		if g.Member.HasParentObject() && g.Member.ParentObject.ObjectUniqueID == g.Group.ObjectUniqueID {
			t.Errorf("member's parent's unique id cannot be the same as its group's")
		}

		verifyStructure := func(identifier ObjectIdentifier) {
			for ; identifier.ParentObject != nil; identifier = *identifier.ParentObject {
				if identifier.ParentObject.ObjectName != expectedParent[identifier.ObjectName] {
					t.Errorf("for object %s, expected parent %s, got %s\n", identifier.ObjectName, expectedParent[identifier.ObjectName], identifier.ParentObject.ObjectName)
				}
			}
			return
		}
		verifyStructure(g.Group)
		verifyStructure(g.Member)
	}
}
