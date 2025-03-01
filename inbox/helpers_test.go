package inbox

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestGetGroupAndAssociatedMemberIdentifier(t *testing.T) {
	jamesIdentifier := ObjectIdentifier{
		ObjectUniqueID: "1",
		ObjectType:     DbType,
		IsGroupType:    true,
		ParentObject: &ObjectIdentifier{
			ObjectUniqueID: "2",
			ObjectType:     NamespaceType,
			IsGroupType:    true,
			ParentObject: &ObjectIdentifier{
				ObjectUniqueID: "3",
				ObjectType:     ClusterType,
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
				ObjectUniqueID: "1",
				ObjectType:     DbType,
				IsGroupType:    true,
				ParentObject: &ObjectIdentifier{
					ObjectUniqueID: "2",
					ObjectType:     NamespaceType,
					IsGroupType:    true,
					ParentObject: &ObjectIdentifier{
						ObjectUniqueID: "3",
						ObjectType:     ClusterType,
						IsGroupType:    true,
					},
				},
			},
			Member: ObjectIdentifier{
				ObjectUniqueID: "2",
				ObjectType:     NamespaceType,
				IsGroupType:    true,
				ParentObject: &ObjectIdentifier{
					ObjectUniqueID: "3",
					ObjectType:     ClusterType,
					IsGroupType:    true,
				},
			},
		},
		{
			Group: ObjectIdentifier{
				ObjectUniqueID: "2",
				ObjectType:     NamespaceType,
				IsGroupType:    true,
				ParentObject: &ObjectIdentifier{
					ObjectUniqueID: "3",
					ObjectType:     ClusterType,
					IsGroupType:    true,
				},
			},
			Member: ObjectIdentifier{
				ObjectUniqueID: "3",
				ObjectType:     ClusterType,
				IsGroupType:    true,
			},
		},
	}

	if diff := cmp.Diff(results, expected); len(diff) != 0 {
		t.Errorf("difference between expected output and result:\n%s", diff)
	}
}
