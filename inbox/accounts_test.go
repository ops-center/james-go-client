package inbox

import (
	"github.com/pkg/errors"
	"log"
	"math/rand"
	"os"
	"strings"
	"testing"
)

const testRandomSeed = 42

var testRand *rand.Rand

func init() {
	testRand = rand.New(rand.NewSource(int64(testRandomSeed)))
}

func TestAddGroups(t *testing.T) {
	const (
		numberOfGroups  = 100
		membersPerGroup = 1
	)
	client, err := getTestWebAdminClient()
	if err != nil {
		t.Error("failed to get test webAdmin client")
		return
	}

	if failed, err := client.AddGroups(generateGroupAndMember(numberOfGroups, membersPerGroup)); err != nil {
		for _, g := range failed {
			log.Println("failed group name: " + g.Address)
			log.Println("reason " + g.Reason)
			log.Println("status: " + g.Reason)
			for _, m := range g.MembersInfo {
				log.Println("member name: " + m.Address)
				log.Println("reason: " + m.Reason)
				log.Println("status: " + m.Status)
			}
			log.Println("-------------------")
		}
		t.Error(errors.Wrap(err, "failed to add groups"))
		return
	}
}

func generateGroupAndMember(groups, members int) []GroupAndAssociatedMember {
	BounderUserIdentity := UserIdentity{
		OwnerID:   genRandString(5),
		OwnerType: genRandString(5),
		OwnerName: genRandString(5),
	}

	var groupsAndAssociatedMembers []GroupAndAssociatedMember
	for i := 0; i < groups; i++ {
		group := getObjectIdentifier(ClusterType, nil, BounderUserIdentity)

		for j := 0; j < members; j++ {
			member := getObjectIdentifier(GroupType, &group, BounderUserIdentity)
			groupsAndAssociatedMembers = append(
				groupsAndAssociatedMembers,
				&GroupAndAssociatedMemberIdentifier{
					Group:  group,
					Member: member,
				},
			)
		}
	}

	return groupsAndAssociatedMembers
}

func getObjectIdentifier(
	Type ObjectTypeIdentifier,
	parent *ObjectIdentifier,
	BounderUserIdentity UserIdentity,
) ObjectIdentifier {
	return ObjectIdentifier{
		ObjectName:          genRandString(10 + testRand.Intn(10)),
		ObjectType:          Type,
		IsGroupType:         Type == GroupType,
		ParentObject:        parent,
		AdditionalClaims:    nil,
		BoundedUserIdentity: BounderUserIdentity.DeepCopy(),
	}
}

func genRandString(len int) string {
	ret := strings.Builder{}
	for i := 0; i < len; i++ {
		ret.WriteRune(rune(testRand.Intn(26) + 'a'))
	}
	return ret.String()
}

func getTestWebAdminClient() (*WebAdminClient, error) {
	const (
		webAdminSessionEndpoint = "http://<host>:<port>"
	)
	_ = os.Setenv("INBOX_SERVER_TOKEN", "<jwt_token>")
	return NewWebAdminClient(
		&WebAdminConf{
			WebAdminServiceEndpoint: webAdminSessionEndpoint,
		},
	)
}
