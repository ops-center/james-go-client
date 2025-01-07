/*
Apache JAMES Web Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

// AddressGroupAPIService AddressGroupAPI service
type AddressGroupAPIService service

type ApiAddMemberRequest struct {
	ctx           context.Context
	ApiService    *AddressGroupAPIService
	groupAddress  string
	memberAddress *string
	groups        []*Group
}

// Member mail address
func (r ApiAddMemberRequest) MemberAddress(memberAddress string) ApiAddMemberRequest {
	r.memberAddress = &memberAddress
	return r
}

func (r ApiAddMemberRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddMemberExecute(r)
}

/*
AddMember Add a group member

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupAddress Group mail address
	@return ApiAddMemberRequest
*/
func (a *AddressGroupAPIService) AddMember(ctx context.Context, groupAddress string) ApiAddMemberRequest {
	return ApiAddMemberRequest{
		ApiService:   a,
		ctx:          ctx,
		groupAddress: groupAddress,
	}
}

// Execute executes the request
func (a *AddressGroupAPIService) AddMemberExecute(r ApiAddMemberRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPut
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.AddMember")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	if r.memberAddress == nil {
		return nil, reportError("memberAddress is required and must be specified")
	}
	localVarPath := fmt.Sprintf("%s/%s/%s", localBasePath+"/address/groups", r.groupAddress, *r.memberAddress)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	parameterAddToHeaderOrQuery(localVarQueryParams, "memberAddress", r.memberAddress, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiAddGroupsRequest struct {
	ctx        context.Context
	ApiService *AddressGroupAPIService
	groups     []*Group
}

const (
	GroupAddStatusFail    = "failed"
	GroupAddStatusSuccess = "success"
)

type GroupMembersInfo struct {
	Address string `json:"address"`
	Status  string `json:"status"`
	Reason  string `json:"reason"`
}

type GroupsWithMembersInfo struct {
	Address     string             `json:"address"`
	Status      string             `json:"status"`
	Reason      string             `json:"reason"`
	MembersInfo []GroupMembersInfo `json:"membersInfo,omitempty"`
}

func (r ApiAddGroupsRequest) Execute() (*http.Response, error) {
	return r.ApiService.AddGroupsExecute(r)
}

/*
AddGroups Add multiple groups with members

Adds multiple groups along with their members.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return AddGroupsRequest
*/
func (a *AddressGroupAPIService) AddGroups(ctx context.Context, groups []*Group) ApiAddGroupsRequest {
	return ApiAddGroupsRequest{
		ApiService: a,
		ctx:        ctx,
		groups:     groups,
	}
}

// Execute executes the request
//
//	@return AddGroups200Response
func (a *AddressGroupAPIService) AddGroupsExecute(r ApiAddGroupsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.AddressGroupsAddGroupsPost")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/address/groups/add-groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groups == nil {
		return nil, reportError("group is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groups
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}

		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiListGroupsRequest struct {
	ctx        context.Context
	ApiService *AddressGroupAPIService
}

func (r ApiListGroupsRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ListGroupsExecute(r)
}

/*
ListGroups List address groups

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiListGroupsRequest
*/
func (a *AddressGroupAPIService) ListGroups(ctx context.Context) ApiListGroupsRequest {
	return ApiListGroupsRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []string
func (a *AddressGroupAPIService) ListGroupsExecute(r ApiListGroupsRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.ListGroups")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/address/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListMembersRequest struct {
	ctx          context.Context
	ApiService   *AddressGroupAPIService
	groupAddress string
}

func (r ApiListMembersRequest) Execute() ([]string, *http.Response, error) {
	return r.ApiService.ListMembersExecute(r)
}

/*
ListMembers List members of a group

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupAddress Group mail address
	@return ApiListMembersRequest
*/
func (a *AddressGroupAPIService) ListMembers(ctx context.Context, groupAddress string) ApiListMembersRequest {
	return ApiListMembersRequest{
		ApiService:   a,
		ctx:          ctx,
		groupAddress: groupAddress,
	}
}

// Execute executes the request
//
//	@return []string
func (a *AddressGroupAPIService) ListMembersExecute(r ApiListMembersRequest) ([]string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.ListMembers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/address/groups/{groupAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupAddress"+"}", url.PathEscape(parameterValueToString(r.groupAddress, "groupAddress")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiRemoveMemberRequest struct {
	ctx           context.Context
	ApiService    *AddressGroupAPIService
	groupAddress  string
	memberAddress *string
}

// Member mail address
func (r ApiRemoveMemberRequest) MemberAddress(memberAddress string) ApiRemoveMemberRequest {
	r.memberAddress = &memberAddress
	return r
}

func (r ApiRemoveMemberRequest) Execute() (*http.Response, error) {
	return r.ApiService.RemoveMemberExecute(r)
}

/*
RemoveMember Remove a group member

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupAddress Group mail address
	@return ApiRemoveMemberRequest
*/
func (a *AddressGroupAPIService) RemoveMember(ctx context.Context, groupAddress string) ApiRemoveMemberRequest {
	return ApiRemoveMemberRequest{
		ApiService:   a,
		ctx:          ctx,
		groupAddress: groupAddress,
	}
}

// Execute executes the request
func (a *AddressGroupAPIService) RemoveMemberExecute(r ApiRemoveMemberRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.RemoveMember")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	if r.memberAddress == nil {
		return nil, reportError("memberAddress is required and must be specified")
	}

	localVarPath := localBasePath + "/address/groups/{groupAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupAddress"+"}", url.PathEscape(parameterValueToString(r.groupAddress, "groupAddress")), -1) + "/" + *r.memberAddress

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// parameterAddToHeaderOrQuery(localVarQueryParams, "memberAddress", r.memberAddress, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCheckGruopMemberExistenceRequest struct {
	ctx           context.Context
	ApiService    *AddressGroupAPIService
	groupAddress  string
	memberAddress *string
}

// Member mail address
func (r ApiCheckGruopMemberExistenceRequest) MemberAddress(memberAddress string) ApiCheckGruopMemberExistenceRequest {
	r.memberAddress = &memberAddress
	return r
}

func (r ApiCheckGruopMemberExistenceRequest) Execute() (*http.Response, error) {
	return r.ApiService.CheckGroupMemberExecute(r)
}

/*
CheckGruopMemberExistence checks if a member exists in the group

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param groupAddress Group mail address
@return ApiRemoveMemberRequest
*/
func (a *AddressGroupAPIService) CheckGruopMemberExistence(ctx context.Context, groupAddress string) ApiCheckGruopMemberExistenceRequest {
	return ApiCheckGruopMemberExistenceRequest{
		ApiService:   a,
		ctx:          ctx,
		groupAddress: groupAddress,
	}
}

// Execute executes the request
func (a *AddressGroupAPIService) CheckGroupMemberExecute(r ApiCheckGruopMemberExistenceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.RemoveMember")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	if r.memberAddress == nil {
		return nil, reportError("memberAddress is required and must be specified")
	}

	localVarPath := localBasePath + "/address/groups/{groupAddress}"
	localVarPath = strings.Replace(localVarPath, "{"+"groupAddress"+"}", url.PathEscape(parameterValueToString(r.groupAddress, "groupAddress")), -1) + "/" + *r.memberAddress

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// parameterAddToHeaderOrQuery(localVarQueryParams, "memberAddress", r.memberAddress, "")
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteGroupRequest struct {
	ctx          context.Context
	ApiService   *AddressGroupAPIService
	groupAddress *string
}

func (r ApiDeleteGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupExecute(r)
}

/*
DeleteGroup deletes a group

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupAddress Group mail address
	@return ApiDeleteGroupRequest
*/
func (a *AddressGroupAPIService) DeleteGroup(ctx context.Context, groupAddress string) ApiDeleteGroupRequest {
	return ApiDeleteGroupRequest{
		ApiService:   a,
		ctx:          ctx,
		groupAddress: &groupAddress,
	}
}

// Execute executes the request
func (a *AddressGroupAPIService) DeleteGroupExecute(r ApiDeleteGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.DeleteGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	if r.groupAddress == nil {
		return nil, reportError("groupAddress is requiered and must be specified")
	}
	localVarPath := localBasePath + "/address/groups/" + *r.groupAddress

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiDeleteGroupsRequest struct {
	ctx            context.Context
	ApiService     *AddressGroupAPIService
	groupAddresses []string
}

func (r ApiDeleteGroupsRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteGroupsExecute(r)
}

/*
DeleteGroups Deletes a list of grups

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupAddresses is a list of Group mail address
	@return ApiDeleteGroupsRequest
*/
func (a *AddressGroupAPIService) DeleteGroups(ctx context.Context, groupAddress []string) ApiDeleteGroupsRequest {
	return ApiDeleteGroupsRequest{
		ApiService:     a,
		ctx:            ctx,
		groupAddresses: groupAddress,
	}
}

// Execute executes the request
func (a *AddressGroupAPIService) DeleteGroupsExecute(r ApiDeleteGroupsRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.DeleteGroups")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	if r.groupAddresses == nil {
		return nil, reportError("groupAddresses is requiered and must be specified")
	}
	localVarPath := localBasePath + "/address/groups"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}

	// body params
	localVarPostBody = r.groupAddresses
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCreateGroupRequest struct {
	ctx          context.Context
	ApiService   *AddressGroupAPIService
	groupAddress *string
}

func (r ApiCreateGroupRequest) Execute() (*http.Response, error) {
	return r.ApiService.CreateGroupExecute(r)
}

/*
CreateGroup creates a group with dummy user

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupAddress Group mail address
	@return ApiCreateGroupRequest
*/
func (a *AddressGroupAPIService) CreateGroup(ctx context.Context, groupAddress string) ApiCreateGroupRequest {
	return ApiCreateGroupRequest{
		ApiService:   a,
		ctx:          ctx,
		groupAddress: &groupAddress,
	}
}

// Execute executes the request
func (a *AddressGroupAPIService) CreateGroupExecute(r ApiCreateGroupRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.CreateGroup")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	if r.groupAddress == nil {
		return nil, reportError("groupAddress is requiered and must be specified")
	}
	localVarPath := localBasePath + "/address/groups/" + *r.groupAddress

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiCheckMultipleGroupMemberPairExistenceRequest struct {
	ctx              context.Context
	ApiService       *AddressGroupAPIService
	groupMemberPairs []GroupMemberPair
}

func (r ApiCheckMultipleGroupMemberPairExistenceRequest) Execute() (*http.Response, error) {
	return r.ApiService.CheckMultipleGruopMemberPairExecute(r)
}

/*
ApiCheckMultipleGruopMemberPairExistenceRequest checks if a member exists in the group for multiple pair of (group,member)

@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
@param groupMemberPair
@return ApiRemoveMemberRequest
*/
func (a *AddressGroupAPIService) CheckMultipleGroupMemberPairExistence(ctx context.Context, groupMemberPair []GroupMemberPair) ApiCheckMultipleGroupMemberPairExistenceRequest {
	return ApiCheckMultipleGroupMemberPairExistenceRequest{
		ApiService:       a,
		ctx:              ctx,
		groupMemberPairs: groupMemberPair,
	}
}

// Execute executes the request
func (a *AddressGroupAPIService) CheckMultipleGruopMemberPairExecute(r ApiCheckMultipleGroupMemberPairExistenceRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodGet
		localVarPostBody   interface{}
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AddressGroupAPIService.CheckMultipleGruopMemberPairExecute")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/address/groups/associations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.groupMemberPairs == nil {
		return nil, reportError("group_members_pairs is required and can't be empty")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.groupMemberPairs

	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}

		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
