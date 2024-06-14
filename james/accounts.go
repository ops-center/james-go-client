package james

import (
	"context"
	"github.com/pkg/errors"
	"net/http"

	openapi "github.com/ops-center/james-go-client"
)

// todo james: restructure errors
// todo james: manage error handling in a proper way
// todo james: use context in every exported function

const (
	RandomPassword = "H1sJk6ORaS" // todo james: remove that later
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

func (js *Service) createAccount(object Object) error {
	if object.IsGroup() { // don't need to create account for groups
		return nil
	}
	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	r, err := js.Client.UsersAPI.UpsertUser(context.TODO(), objectAddr).
		UpsertUserRequest(openapi.UpsertUserRequest{Password: RandomPassword}).Execute()

	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (js *Service) deleteAccount(object Object) error {
	if object.IsGroup() {
		return nil
	}

	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	r, err := js.Client.UsersAPI.DeleteUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (js *Service) CreateObject(object Object) error {
	if object.IsGroup() {
		return js.createGroup(object)
	}

	return js.createAccount(object)
}

func (js *Service) DeleteObject(object Object) error {
	if object.IsGroup() {
		return js.deleteGroup(object)
	}

	return js.deleteAccount(object)
}

func (js *Service) AddGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := js.GetObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
	}
	memberAddr, err := js.GetObjectAddr(memberObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))

	}

	return js.addGroupMember(grpAddr, memberAddr)
}

func (js *Service) addGroupMember(grpAddr, memberAddr string) error {
	r, err := js.Client.AddressGroupAPI.AddMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (js *Service) RemoveGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))

	}
	memberAddr, err := generateObjectAddr(memberObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate member object address: %v", err))

	}

	return js.removeGroupMember(grpAddr, memberAddr)
}

func (js *Service) removeGroupMember(grpAddr, memberAddr string) error {
	r, err := js.Client.AddressGroupAPI.RemoveMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	if r.StatusCode != http.StatusNoContent {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (js *Service) createGroup(object Object) error {
	grpAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	r, err := js.Client.AddressGroupAPI.CreateGroup(context.Background(), grpAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (js *Service) deleteGroup(grpObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
	}

	r, err := js.Client.AddressGroupAPI.DeleteGroup(context.TODO(), grpAddr).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (js *Service) DeleteGroups(grpObjects []Object) error {
	grpAddresses := make([]string, len(grpObjects))
	for i := 0; i < len(grpObjects); i++ {
		grpAddr, err := js.GetObjectAddr(grpObjects[i])
		if err != nil {
			return newServerError(nil, errors.Errorf("failed to generate group object address: %v", err))
		}
		grpAddresses[i] = grpAddr
	}
	r, err := js.Client.AddressGroupAPI.DeleteGroups(context.TODO(), grpAddresses).Execute()
	if err != nil {
		return newServerError(r, err)
	}

	return nil
}

func (js *Service) CheckObjectExistence(object Object) error {
	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	return js.checkAddrExistence(objectAddr)
}

func (js *Service) checkAddrExistence(objectAddr string) error {
	r, err := js.Client.UsersAPI.ExistsUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		if r != nil && r.StatusCode == http.StatusNotFound {
			return ErrUserNotExists
		}
		return newServerError(r, err)
	}

	return nil
}

func (js *Service) UpdateObjectAddr(oldObject, newObject Object) error {
	oldObjectAddr, err := generateObjectAddr(oldObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}
	newObjectAddr, err := generateObjectAddr(newObject)
	if err != nil {
		return newServerError(nil, errors.Errorf("failed to generate object address: %v", err))
	}

	if err = js.checkAddrExistence(oldObjectAddr); err != nil {
		return errors.Errorf("address doesn't exits: %v", err)
	}

	err = js.createAccount(newObject)
	if err != nil && errors.As(err, &ErrUserAlreadyExists) {
		return err
	}

	r, err := js.Client.UsersAPI.ChangeUsername(context.TODO(), oldObjectAddr, newObjectAddr).Execute()
	if err != nil {
		return newServerError(r, errors.Errorf("failed to change username: %v", err))
	}
	if r.StatusCode != http.StatusCreated {
		return newServerError(r, errors.Errorf("unknown error: status: %v", r.Status))
	}

	return nil
}

func (js *Service) GetObjectAddr(object Object) (string, error) {
	addrStr, err := generateObjectAddr(object)
	if err != nil {
		return "", err
	}

	return addrStr, nil
}

func (js *Service) IsObjectExists(object Object) (bool, error) {
	err := js.CheckObjectExistence(object)
	if err != nil {
		if !errors.Is(err, ErrUserNotExists) {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

func (js *Service) CreateObjectIfNotExists(object Object) error {
	exists, err := js.IsObjectExists(object)
	if err != nil {
		return err
	}

	if !exists {
		return js.CreateObject(object)
	}

	return nil
}
