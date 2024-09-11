package inbox

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"net/http"

	openapi "go.opscenter.dev/james-go-client"
)

var ErrUserAlreadyExists = errors.New("user already exists")

func (w *WebAdminClient) createAccount(object Object) error {
	if object.IsGroup() { // don't need to create account for groups
		return nil
	}
	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	password, err := generateRandomPassword()
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate random password: %v", err))
	}

	r, err := w.UsersAPI.UpsertUser(context.TODO(), objectAddr).
		UpsertUserRequest(openapi.UpsertUserRequest{Password: password}).Execute()

	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (w *WebAdminClient) deleteAccount(object Object) error {
	if object.IsGroup() {
		return nil
	}

	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	r, err := w.UsersAPI.DeleteUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
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

func (w *WebAdminClient) AddGroups(groups []GroupAndAssociatedMember) error {
	var groupList []*openapi.Group

	for _, group := range groups {

		groupAddr, err := generateObjectAddr(group.GetGroup())
		if err != nil {
			return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
		}

		memberAddr, err := generateObjectAddr(group.GetMember())
		if err != nil {
			return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))
		}

		groupList = append(groupList, openapi.NewGroup(groupAddr, []string{memberAddr}))
	}

	return w.addGroups(groupList)
}

func (w *WebAdminClient) AddObjectAlias(object Object) error {
	userAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	addrAlias, err := object.GetAddressAlias()
	if err != nil {
		return newServerError(nil, errors.Errorf("couldn't get address alias: %v", err))
	}
	addrAlias = fmt.Sprintf("%s@%s", addrAlias, GlobalMailDomain)

	return w.addAddressAlias(userAddr, addrAlias)
}

func (w *WebAdminClient) addAddressAlias(userAddr string, alias string) error {
	r, err := w.AddressAliasAPI.CreateAlias(context.TODO(), userAddr).SourceAddress(alias).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (w *WebAdminClient) RemoveObjectAlias(object Object) error {
	userAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	addrAlias, err := object.GetAddressAlias()
	if err != nil {
		return newServerError(nil, errors.Errorf("couldn't get address alias: %v", err))
	}

	return w.removeAddressAlias(userAddr, addrAlias)
}

func (w *WebAdminClient) removeAddressAlias(userAddr string, alias string) error {
	r, err := w.AddressAliasAPI.DeleteAlias(context.TODO(), userAddr).SourceAddress(alias).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (w *WebAdminClient) ListObjectAliases(object Object) ([]openapi.GetAlias200ResponseInner, error) {
	userAddr, err := generateObjectAddr(object)
	if err != nil {
		return nil, newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	return w.listAddressAliases(userAddr)
}

func (w *WebAdminClient) listAddressAliases(userAddr string) ([]openapi.GetAlias200ResponseInner, error) {
	aliases, r, err := w.AddressAliasAPI.GetAlias(context.TODO(), userAddr).Execute()
	if err != nil {
		return nil, newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return nil, newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return aliases, nil
}

func (w *WebAdminClient) RemoveAllObjectAliases(object Object) error {
	userAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	return w.removeAllAddressAliases(userAddr)
}

func (w *WebAdminClient) removeAllAddressAliases(userAddr string) error {
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
			return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
		}
	}

	return nil
}

func (w *WebAdminClient) addGroups(groups []*openapi.Group) error {
	r, err := w.AddressGroupAPI.AddGroups(context.TODO(), groups).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (w *WebAdminClient) AddGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := w.GetObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
	}
	memberAddr, err := w.GetObjectAddr(memberObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))

	}

	return w.addGroupMember(grpAddr, memberAddr)
}

func (w *WebAdminClient) addGroupMember(grpAddr, memberAddr string) error {
	r, err := w.AddressGroupAPI.AddMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (w *WebAdminClient) RemoveGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))

	}
	memberAddr, err := generateObjectAddr(memberObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))

	}

	return w.removeGroupMember(grpAddr, memberAddr)
}

func (w *WebAdminClient) removeGroupMember(grpAddr, memberAddr string) error {
	r, err := w.AddressGroupAPI.RemoveMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (w *WebAdminClient) CheckGroupMemberExistence(grpObject, memberObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
	}
	memberAddr, err := generateObjectAddr(memberObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))
	}

	return w.checkGroupMemberExistence(grpAddr, memberAddr)
}

func (w *WebAdminClient) checkGroupMemberExistence(grpAddr, memberAddr string) error {
	r, err := w.AddressGroupAPI.CheckGruopMemberExistence(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (w *WebAdminClient) createGroup(object Object) error {
	grpAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	r, err := w.AddressGroupAPI.CreateGroup(context.Background(), grpAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (w *WebAdminClient) deleteGroup(grpObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
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
			return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
		}
		grpAddresses[i] = grpAddr
	}
	r, err := w.AddressGroupAPI.DeleteGroups(context.TODO(), grpAddresses).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (w *WebAdminClient) CheckObjectExistence(object Object) error {
	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
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
	oldObjectAddr, err := generateObjectAddr(oldObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	newObjectAddr, err := generateObjectAddr(newObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	if err = w.checkAddrExistence(oldObjectAddr); err != nil {
		return errors.Errorf("address doesn't exits: %v", err)
	}

	err = w.createAccount(newObject)
	if err != nil && errors.As(err, &ErrUserAlreadyExists) {
		return err
	}

	r, err := w.UsersAPI.ChangeUsername(context.TODO(), oldObjectAddr, newObjectAddr).Execute()
	if err != nil {
		return newServerError(r, errors.Errorf("failed to change username: %v", err))
	}
	if r.StatusCode != http.StatusCreated {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (w *WebAdminClient) GetObjectAddr(object Object) (string, error) {
	addrStr, err := generateObjectAddr(object)
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

func (w *WebAdminClient) HealthCheck() (*openapi.CheckAllComponents200Response, error) {
	checkAllComponentsResponse, r, err := w.HealthcheckAPI.CheckAllComponents(context.TODO()).Execute()
	if err != nil {
		return nil, newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return nil, newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
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
			return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
		}
	}

	return nil
}

func (w *WebAdminClient) CreateCloudAppcodeDomain() error {
	return w.CreateDomains(GlobalMailDomain)
}
