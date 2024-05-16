package james

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	openapi "github.com/searchlight/james-go-client"
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
		return fmt.Errorf("generate object address: %v", err)
	}
	r, err := js.Client.UsersAPI.UpsertUser(context.TODO(), objectAddr).
		UpsertUserRequest(openapi.UpsertUserRequest{Password: RandomPassword}).Execute()

	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}
	if r.StatusCode == http.StatusConflict {
		return ErrUserAlreadyExists
	}
	if r.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to create user account, status: %v", r.Status)
	}

	return nil
}

func (js *Service) deleteAccount(object Object) error {
	if object.IsGroup() {
		return nil
	}

	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return fmt.Errorf("generate object address: %v", err)
	}

	r, err := js.Client.UsersAPI.DeleteUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	if r.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete user account, status: %v", r.Status)
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
		return fmt.Errorf("generate object address: %v", err)
	}
	memberAddr, err := js.GetObjectAddr(memberObject)
	if err != nil {
		return fmt.Errorf("generate object address: %v", err)
	}

	return js.addGroupMember(grpAddr, memberAddr)
}

func (js *Service) addGroupMember(grpAddr, memberAddr string) error {
	r, err := js.Client.AddressGroupAPI.AddMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	if r.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to add group member, status: %v", r.Status)
	}

	return nil
}

func (js *Service) RemoveGroupMember(grpObject, memberObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return fmt.Errorf("generate object address: %v", err)
	}
	memberAddr, err := generateObjectAddr(memberObject)
	if err != nil {
		return fmt.Errorf("generate object address: %v", err)
	}

	return js.removeGroupMember(grpAddr, memberAddr)
}

func (js *Service) removeGroupMember(grpAddr, memberAddr string) error {
	r, err := js.Client.AddressGroupAPI.RemoveMember(context.TODO(), grpAddr).MemberAddress(memberAddr).Execute()
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	if r.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to remove group member, status: %v", r.Status)
	}

	return nil
}

func (js *Service) createGroup(object Object) error {
	grpAddr, err := generateObjectAddr(object)
	if err != nil {
		return err
	}
	r, err := js.Client.AddressGroupAPI.CreateGroup(context.Background(), grpAddr).Execute()
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}
	if r.StatusCode >= 300 {
		return fmt.Errorf("failed to create group, status: %v", r.Status)
	}

	return nil
}

func (js *Service) deleteGroup(grpObject Object) error {
	grpAddr, err := generateObjectAddr(grpObject)
	if err != nil {
		return err
	}

	r, err := js.Client.AddressGroupAPI.DeleteGroup(context.TODO(), grpAddr).Execute()
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	if r.StatusCode >= 300 {
		return fmt.Errorf("failed get delete group, status code: %v", r.StatusCode)
	}

	return nil
}

func (js *Service) DeleteGroups(grpObjects []Object) error {
	grpAddresses := make([]string, len(grpObjects))
	for i := 0; i < len(grpObjects); i++ {
		grpAddr, err := js.GetObjectAddr(grpObjects[i])
		if err != nil {
			return err
		}
		grpAddresses[i] = grpAddr
	}
	r, err := js.Client.AddressGroupAPI.DeleteGroups(context.TODO(), grpAddresses).Execute()
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	if r.StatusCode >= 300 {
		return fmt.Errorf("failed get delete group list, status code: %v", r.StatusCode)
	}

	return nil
}

func (js *Service) CheckObjectExistence(object Object) error {
	objectAddr, err := generateObjectAddr(object)
	if err != nil {
		return fmt.Errorf("generate object address: %v", err)
	}

	return js.checkAddrExistence(objectAddr)
}

func (js *Service) checkAddrExistence(objectAddr string) error {
	r, err := js.Client.UsersAPI.ExistsUser(context.TODO(), objectAddr).Execute()
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf("user doesn't exits, status: %v", r.Status)
	}

	return nil
}

func (js *Service) UpdateObjectAddr(oldObject, newObject Object) error {
	oldObjectAddr, err := generateObjectAddr(oldObject)
	if err != nil {
		return fmt.Errorf("generate object address: %v", err)
	}
	newObjectAddr, err := generateObjectAddr(newObject)
	if err != nil {
		return fmt.Errorf("generate object address: %v", err)
	}

	if err = js.checkAddrExistence(oldObjectAddr); err != nil {
		return fmt.Errorf("update username: %v", err)
	}

	err = js.createAccount(newObject)
	if err != nil && errors.As(err, &ErrUserAlreadyExists) {
		return err
	}

	r, err := js.Client.UsersAPI.ChangeUsername(context.TODO(), oldObjectAddr, newObjectAddr).Execute()
	if err != nil {
		return fmt.Errorf("server error: %v", err)
	}
	if r.StatusCode != http.StatusCreated {
		// todo: get the server error, unmarshal it and generated new error
		return fmt.Errorf("failed to update usrename")
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
