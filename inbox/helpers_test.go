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

func TestIsValidDomainFormat(t *testing.T) {
	tests := []struct {
		emailDomain string
		expected    bool
	}{
		{"example.com", true},
		{"sub.example.com", true},
		{"ace.internal", true},
		{"example.co.bd", true},
		{"example", false},
		{"@example.com", false},
		{"example@com", false},
		{"", false},
	}

	for _, test := range tests {
		result := isValidDomainFormat(test.emailDomain)
		if result != test.expected {
			t.Errorf("IsValidEmailDomain(%q) = %v; expected %v", test.emailDomain, result, test.expected)
		}
	}
}
