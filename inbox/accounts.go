package inbox

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	apis "go.opscenter.dev/james-go-client"
)

var ErrUserAlreadyExists = errors.New("user already exists")

func (w *WebAdminClient) createAccount(object Object) error {
	if object.IsGroup() { // don't need to create account for groups
		return nil
	}
	objectAddr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}

	password, err := generateRandomPassword()
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate random password: %w", err))
	}

	r, err := w.UsersAPI.UpsertUser(context.TODO(), objectAddr).
		UpsertUserRequest(apis.UpsertUserRequest{Password: password}).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) deleteAccount(object Object) error {
	if object.IsGroup() {
		return nil
	}

	objectAddr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}

	if err = w.deleteAccountAddr(objectAddr); err != nil {
		return err
	}

	return nil
}

func (w *WebAdminClient) deleteAccountAddr(objectAddr string) error {
	r, err := w.UsersAPI.DeleteUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) deleteMultipleAccounts(objectAddresses []string) error {
	if len(objectAddresses) == 0 {
		return nil
	}
	r, err := w.UsersAPI.DeleteUsers(context.TODO(), objectAddresses).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) CreateObject(object Object) error {
	if object.IsGroup() {
		return w.createGroup(object)
	}

	return w.createAccount(object)
}

func (w *WebAdminClient) DeleteObject(object Object) error {
	if object.IsGroup() {
		return w.deleteGroup(object)
	}

	return w.deleteAccount(object)
}

func (w *WebAdminClient) AddGroups(groups []GroupAndAssociatedMember) ([]apis.GroupsWithMembersInfo, error) {
	groupMap := make(map[string]*apis.Group)

	for _, group := range groups {

		groupAddr, err := generateObjectAddr(group.GetGroup(), w.emailDomain)
		if err != nil {
			return nil, newServerError(nil, fmt.Errorf("failed to generate group object address: %w", err))
		}

		memberAddr, err := generateObjectAddr(group.GetMember(), w.emailDomain)
		if err != nil {
			return nil, newServerError(nil, fmt.Errorf("failed to generate member object address: %w", err))
		}

		if apisGroup, ok := groupMap[groupAddr]; !ok {
			groupMap[groupAddr] = apis.NewGroup(groupAddr, []string{memberAddr})
		} else {
			apisGroup.MemberAddrs = append(apisGroup.MemberAddrs, memberAddr)
		}
	}

	var groupList []*apis.Group
	for _, group := range groupMap {
		groupList = append(groupList, group)
	}

	return w.addGroups(groupList)
}

func (w *WebAdminClient) AddObjectAlias(object Object) error {
	userAddr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}
	addrAlias, err := object.GetAddressAlias()
	if err != nil {
		return newServerError(nil, fmt.Errorf("couldn't get address alias: %w", err))
	}
	addrAlias = fmt.Sprintf("%s@%s", addrAlias, w.emailDomain)

	return w.addAddressAlias(userAddr, addrAlias)
}

func (w *WebAdminClient) addAddressAlias(userAddr string, alias string) error {
	r, err := w.AddressAliasAPI.CreateAlias(context.TODO(), userAddr).SourceAddress(alias).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) RemoveObjectAlias(object Object) error {
	userAddr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}
	addrAlias, err := object.GetAddressAlias()
	if err != nil {
		return newServerError(nil, fmt.Errorf("couldn't get address alias: %w", err))
	}
	addrAlias = fmt.Sprintf("%s@%s", addrAlias, w.emailDomain)

	return w.removeAddressAlias(userAddr, addrAlias)
}

func (w *WebAdminClient) removeAddressAlias(userAddr string, alias string) error {
	r, err := w.AddressAliasAPI.DeleteAlias(context.TODO(), userAddr).SourceAddress(alias).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) ListObjectAliases(object Object) ([]apis.GetAlias200ResponseInner, error) {
	userAddr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return nil, newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}

	return w.listAddressAliases(userAddr)
}

func (w *WebAdminClient) listAddressAliases(userAddr string) ([]apis.GetAlias200ResponseInner, error) {
	aliases, r, err := w.AddressAliasAPI.GetAlias(context.TODO(), userAddr).Execute()
	if err != nil {
		return nil, newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return nil, newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return aliases, nil
}

func (w *WebAdminClient) RemoveAllObjectAliases(object Object) error {
	userAddr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}

	return w.removeAllAddressAliasesOfAUserAddr(userAddr)
}

func (w *WebAdminClient) removeAllAddressAliasesOfAUserAddr(userAddr string) error {
	aliases, err := w.listAddressAliases(userAddr)
	if err != nil {
		return err
	}

	for _, alias := range aliases {
		r, err := w.AddressAliasAPI.DeleteAlias(context.TODO(), userAddr).SourceAddress(*alias.Source).Execute()
		if err != nil {
			return newServerError(r, err)
		}

		if r.StatusCode != http.StatusNoContent {
			return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
		}
	}

	return nil
}

func (w *WebAdminClient) deleteAllAddressAliasesOfAListOfUsers(userAddresses []string) error {
	r, err := w.AddressAliasAPI.DeleteAliases(context.TODO(), userAddresses).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) addGroups(groups []*apis.Group) ([]apis.GroupsWithMembersInfo, error) {
	r, err := w.AddressGroupAPI.AddGroups(context.TODO(), groups).Execute()
	if err != nil {
		return nil, newServerError(r, fmt.Errorf("failed to add groups: %w", err))
	}

	if r.StatusCode != http.StatusOK {
		return nil, newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, newServerError(r, fmt.Errorf("failed to read body: %w", err))
	}

	var addGroupsStatus []apis.GroupsWithMembersInfo
	err = json.Unmarshal(body, &addGroupsStatus)
	if err != nil {
		return nil, newServerError(r, fmt.Errorf("failed to unmarshal body: %w", err))
	}

	var failedGroups []apis.GroupsWithMembersInfo

	for _, group := range addGroupsStatus {
		if group.Status != apis.GroupAddStatusSuccess {
			failedGroups = append(failedGroups, group)
			continue
		}

		candidate := apis.GroupsWithMembersInfo{
			Address:     group.Address,
			Status:      group.Status,
			Reason:      group.Reason,
			MembersInfo: nil,
		}

		for _, member := range group.MembersInfo {
			if member.Status != apis.GroupAddStatusSuccess {
				candidate.MembersInfo = append(candidate.MembersInfo, member)
			}
		}

		if candidate.MembersInfo != nil {
			failedGroups = append(failedGroups, candidate)
		}
	}

	if failedGroups != nil {
		return failedGroups, newServerError(r, fmt.Errorf("failed to add groups: %s", failedGroups))
	}

	return nil, nil
}

func (w *WebAdminClient) AddGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := w.GetObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate group object address: %w", err))
	}
	memberAddr, err := w.GetObjectAddr(memberObject)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate member object address: %w", err))
	}

	return w.addGroupMember(grpAddr, memberAddr)
}

func (w *WebAdminClient) addGroupMember(grpAddr, memberAddr string) error {
	r, err := w.AddressGroupAPI.AddMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) RemoveGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate group object address: %w", err))
	}
	memberAddr, err := generateObjectAddr(memberObject, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate member object address: %w", err))
	}

	return w.removeGroupMember(grpAddr, memberAddr)
}

func (w *WebAdminClient) removeGroupMember(grpAddr, memberAddr string) error {
	r, err := w.AddressGroupAPI.RemoveMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) CheckGroupMemberExistence(grpObject, memberObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate group object address: %w", err))
	}
	memberAddr, err := generateObjectAddr(memberObject, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate member object address: %w", err))
	}

	return w.checkGroupMemberExistence(grpAddr, memberAddr)
}

func (w *WebAdminClient) checkGroupMemberExistence(grpAddr, memberAddr string) error {
	r, err := w.AddressGroupAPI.CheckGruopMemberExistence(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) CheckMultipleGroupMemberPairExistence(groupMemberPair []apis.GroupMemberPair) error {
	r, err := w.AddressGroupAPI.CheckMultipleGroupMemberPairExistence(context.TODO(), groupMemberPair).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) createGroup(object Object) error {
	grpAddr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}
	r, err := w.AddressGroupAPI.CreateGroup(context.Background(), grpAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (w *WebAdminClient) deleteGroup(grpObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate group object address: %w", err))
	}

	r, err := w.AddressGroupAPI.DeleteGroup(context.TODO(), grpAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (w *WebAdminClient) DeleteGroups(grpObjects []Object) error {
	grpAddresses := make([]string, len(grpObjects))
	for i := 0; i < len(grpObjects); i++ {
		grpAddr, err := w.GetObjectAddr(grpObjects[i])
		if err != nil {
			return newServerError(nil, fmt.Errorf("failed to generate group object address: %w", err))
		}
		grpAddresses[i] = grpAddr
	}

	if err := w.deleteGroupAddresses(grpAddresses); err != nil {
		return err
	}

	return nil
}

func (w *WebAdminClient) deleteGroupAddresses(grpAddresses []string) error {
	if len(grpAddresses) == 0 {
		return nil
	}
	r, err := w.AddressGroupAPI.DeleteGroups(context.TODO(), grpAddresses).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) CheckObjectExistence(object Object) error {
	objectAddr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}

	return w.checkAddrExistence(objectAddr)
}

func (w *WebAdminClient) checkAddrExistence(objectAddr string) error {
	r, err := w.UsersAPI.ExistsUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		if r != nil && r.StatusCode == http.StatusNotFound {
			return ErrUserNotExists
		}
		return newServerError(r, err)
	}

	return nil
}

func (w *WebAdminClient) UpdateObjectAddr(oldObject, newObject Object) error {
	oldObjectAddr, err := generateObjectAddr(oldObject, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}
	newObjectAddr, err := generateObjectAddr(newObject, w.emailDomain)
	if err != nil {
		return newServerError(nil, fmt.Errorf("failed to generate object address: %w", err))
	}

	if err = w.checkAddrExistence(oldObjectAddr); err != nil {
		return fmt.Errorf("address doesn't exits: %w", err)
	}

	err = w.createAccount(newObject)
	if err != nil && errors.Is(err, ErrUserAlreadyExists) {
		return err
	}

	r, err := w.UsersAPI.ChangeUsername(context.TODO(), oldObjectAddr, newObjectAddr).Execute()
	if err != nil {
		return newServerError(r, fmt.Errorf("failed to change username: %w", err))
	}
	if r.StatusCode != http.StatusCreated {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) GetObjectAddr(object Object) (string, error) {
	addrStr, err := generateObjectAddr(object, w.emailDomain)
	if err != nil {
		return "", err
	}

	return addrStr, nil
}

func (w *WebAdminClient) IsObjectExists(object Object) (bool, error) {
	err := w.CheckObjectExistence(object)
	if err != nil {
		if !errors.Is(err, ErrUserNotExists) {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func (w *WebAdminClient) CreateObjectIfNotExists(object Object) error {
	exists, err := w.IsObjectExists(object)
	if err != nil {
		return err
	}

	if !exists {
		return w.CreateObject(object)
	}

	return nil
}

func (w *WebAdminClient) HealthCheck() (*apis.CheckAllComponents200Response, error) {
	checkAllComponentsResponse, r, err := w.HealthcheckAPI.CheckAllComponents(context.TODO()).Execute()
	if err != nil {
		return nil, newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return nil, newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return checkAllComponentsResponse, nil
}

func (w *WebAdminClient) CreateDomains(domains ...string) error {
	for _, domain := range domains {
		r, err := w.DomainsAPI.CreateDomain(context.TODO(), domain).Execute()
		if err != nil {
			return newServerError(r, err)
		}

		if r.StatusCode != http.StatusNoContent {
			return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
		}
	}

	return nil
}

func (w *WebAdminClient) GetAllUsers() ([]string, error) {
	listUsersResponse, r, err := w.UsersAPI.ListUsers(context.TODO()).Execute()
	if err != nil {
		return nil, newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return nil, newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	var users []string
	for _, userResp := range listUsersResponse {
		users = append(users, *userResp.Username)
	}

	return users, nil
}

func (w *WebAdminClient) DeleteUserAddresses(userAddresses []string) error {
	r, err := w.UsersAPI.DeleteUsers(context.TODO(), userAddresses).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent && r.StatusCode != http.StatusOK {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) GetAllGroups() ([]string, error) {
	listGroupsResponse, r, err := w.AddressGroupAPI.ListGroups(context.TODO()).Execute()
	if err != nil {
		return nil, newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return nil, newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return listGroupsResponse, nil
}

func (w *WebAdminClient) ListAliasAddresses() ([]string, error) {
	listAliasesResponse, r, err := w.AddressAliasAPI.ListAliases(context.TODO()).Execute()
	if err != nil {
		return nil, newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return nil, newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return listAliasesResponse, nil
}

func (w *WebAdminClient) DeleteAliasAddresses(aliasAddresses []string) error {
	r, err := w.AddressAliasAPI.DeleteAliases(context.TODO(), aliasAddresses).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent && r.StatusCode != http.StatusOK {
		return newServerError(r, fmt.Errorf("unknown error: status: %s", r.Status))
	}

	return nil
}

func (w *WebAdminClient) CreateCloudAppcodeDomain() error {
	return w.CreateDomains(w.emailDomain)
}
