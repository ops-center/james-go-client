package james

import (
	"context"
	"github.com/pkg/errors"
	"net/http"

	openapi "github.com/ops-center/james-go-client"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

func (jc *WebAdminClient) createAccount(object Object) error {
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

	r, err := jc.UsersAPI.UpsertUser(context.TODO(), objectAddr).
		UpsertUserRequest(openapi.UpsertUserRequest{Password: password}).Execute()

	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (jc *WebAdminClient) deleteAccount(object Object) error {
	if object.IsGroup() {
		return nil
	}

	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	r, err := jc.UsersAPI.DeleteUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (jc *WebAdminClient) CreateObject(object Object) error {
	if object.IsGroup() {
		return jc.createGroup(object)
	}

	return jc.createAccount(object)
}

func (jc *WebAdminClient) DeleteObject(object Object) error {
	if object.IsGroup() {
		return jc.deleteGroup(object)
	}

	return jc.deleteAccount(object)
}

func (jc *WebAdminClient) AddGroups(groups []GroupAndAssociatedMember) error {
	var groupList []*openapi.Group

	for _, group := range groups {
		groupAddr, err := generateObjectAddr(group.GetGroup())
		if err != nil {
			return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
		}

		member := group.GetMember()
		if member == nil {
			continue
		}

		memberAddr, err := generateObjectAddr(member)
		if err != nil {
			return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))
		}

		groupList = append(groupList, openapi.NewGroup(groupAddr, []string{memberAddr}))
	}

	return jc.addGroups(groupList)
}

func (jc *WebAdminClient) addGroups(groups []*openapi.Group) error {
	r, err := jc.AddressGroupAPI.AddGroups(context.TODO(), groups).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusOK {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (jc *WebAdminClient) AddGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := jc.GetObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
	}
	memberAddr, err := jc.GetObjectAddr(memberObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))

	}

	return jc.addGroupMember(grpAddr, memberAddr)
}

func (jc *WebAdminClient) addGroupMember(grpAddr, memberAddr string) error {
	r, err := jc.AddressGroupAPI.AddMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (jc *WebAdminClient) RemoveGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))

	}
	memberAddr, err := generateObjectAddr(memberObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))

	}

	return jc.removeGroupMember(grpAddr, memberAddr)
}

func (jc *WebAdminClient) removeGroupMember(grpAddr, memberAddr string) error {
	r, err := jc.AddressGroupAPI.RemoveMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (jc *WebAdminClient) createGroup(object Object) error {
	grpAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	r, err := jc.AddressGroupAPI.CreateGroup(context.Background(), grpAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (jc *WebAdminClient) deleteGroup(grpObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
	}

	r, err := jc.AddressGroupAPI.DeleteGroup(context.TODO(), grpAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (jc *WebAdminClient) DeleteGroups(grpObjects []Object) error {
	grpAddresses := make([]string, len(grpObjects))
	for i := 0; i < len(grpObjects); i++ {
		grpAddr, err := jc.GetObjectAddr(grpObjects[i])
		if err != nil {
			return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
		}
		grpAddresses[i] = grpAddr
	}
	r, err := jc.AddressGroupAPI.DeleteGroups(context.TODO(), grpAddresses).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (jc *WebAdminClient) CheckObjectExistence(object Object) error {
	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	return jc.checkAddrExistence(objectAddr)
}

func (jc *WebAdminClient) checkAddrExistence(objectAddr string) error {
	r, err := jc.UsersAPI.ExistsUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		if r != nil && r.StatusCode == http.StatusNotFound {
			return ErrUserNotExists
		}
		return newServerError(r, err)
	}

	return nil
}

func (jc *WebAdminClient) UpdateObjectAddr(oldObject, newObject Object) error {
	oldObjectAddr, err := generateObjectAddr(oldObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	newObjectAddr, err := generateObjectAddr(newObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	if err = jc.checkAddrExistence(oldObjectAddr); err != nil {
		return errors.Errorf("address doesn't exits: %v", err)
	}

	err = jc.createAccount(newObject)
	if err != nil && errors.As(err, &ErrUserAlreadyExists) {
		return err
	}

	r, err := jc.UsersAPI.ChangeUsername(context.TODO(), oldObjectAddr, newObjectAddr).Execute()
	if err != nil {
		return newServerError(r, errors.Errorf("failed to change username: %v", err))
	}
	if r.StatusCode != http.StatusCreated {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (jc *WebAdminClient) GetObjectAddr(object Object) (string, error) {
	addrStr, err := generateObjectAddr(object)
	if err != nil {
		return "", err
	}

	return addrStr, nil
}

func (jc *WebAdminClient) IsObjectExists(object Object) (bool, error) {
	err := jc.CheckObjectExistence(object)
	if err != nil {
		if !errors.Is(err, ErrUserNotExists) {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func (jc *WebAdminClient) CreateObjectIfNotExists(object Object) error {
	exists, err := jc.IsObjectExists(object)
	if err != nil {
		return err
	}

	if !exists {
		return jc.CreateObject(object)
	}

	return nil
}
