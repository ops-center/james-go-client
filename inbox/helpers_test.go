package inbox

import (
	"fmt"
	"github.com/google/go-cmp/cmp"
	"sort"
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

	results, err := GetGroupAndAssociatedMemberIdentifier(jamesIdentifier)
	if err != nil {
		t.Error(err)
		return
	}

	expected := []GroupAndAssociatedMemberIdentifier{
		{
			Group: ObjectIdentifier{
				ObjectName:     "database",
				ObjectUniqueID: "1",
				ObjectType:     7,
				IsGroupType:    true,
				ParentObject: &ObjectIdentifier{
					ObjectName:     "namespace",
					ObjectUniqueID: "2",
					ObjectType:     7,
					IsGroupType:    true,
					ParentObject: &ObjectIdentifier{
						ObjectName:     "cluster",
						ObjectUniqueID: "3",
						ObjectType:     7,
						IsGroupType:    true,
					},
				},
			},
			Member: ObjectIdentifier{
				ObjectName:     "namespace",
				ObjectUniqueID: "2",
				ObjectType:     7,
				IsGroupType:    true,
				ParentObject: &ObjectIdentifier{
					ObjectName:     "cluster",
					ObjectUniqueID: "3",
					ObjectType:     7,
					IsGroupType:    true,
				},
			},
		},
		{
			Group: ObjectIdentifier{
				ObjectName:     "namespace",
				ObjectUniqueID: "2",
				ObjectType:     7,
				IsGroupType:    true,
				ParentObject: &ObjectIdentifier{
					ObjectName:     "cluster",
					ObjectUniqueID: "3",
					ObjectType:     7,
					IsGroupType:    true,
				},
			},
			Member: ObjectIdentifier{
				ObjectName:     "cluster",
				ObjectUniqueID: "3",
				ObjectType:     7,
				IsGroupType:    true,
			},
		},
	}

	if err = isEqual(results, expected); err != nil {
		t.Error(err)
	}
}

func isEqual(results, expected []GroupAndAssociatedMemberIdentifier) error {
	if len(results) != len(expected) {
		return fmt.Errorf("slice length mismatch. expected %d, got %d", len(expected), len(results))
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].Group.ObjectName == results[j].Group.ObjectName {
			return results[i].Member.ObjectName < results[j].Member.ObjectName
		}
		return results[i].Group.ObjectName < results[j].Group.ObjectName
	})

	sort.Slice(expected, func(i, j int) bool {
		if expected[i].Group.ObjectName == expected[j].Group.ObjectName {
			return expected[i].Member.ObjectName < expected[j].Member.ObjectName
		}
		return expected[i].Group.ObjectName < expected[j].Group.ObjectName
	})

	for i := 0; i < len(results); i++ {
		if diff := cmp.Diff(results[i], expected[i]); len(diff) != 0 {
			return fmt.Errorf("difference between expected output and result:\n%s", diff)
		}
	}

	return nil
}
