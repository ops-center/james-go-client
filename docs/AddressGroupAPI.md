# \AddressGroupAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddMember**](AddressGroupAPI.md#AddMember) | **Put** /address/groups/{groupAddress} | Add a group member
[**ListGroups**](AddressGroupAPI.md#ListGroups) | **Get** /address/groups | List address groups
[**ListMembers**](AddressGroupAPI.md#ListMembers) | **Get** /address/groups/{groupAddress} | List members of a group
[**RemoveMember**](AddressGroupAPI.md#RemoveMember) | **Delete** /address/groups/{groupAddress} | Remove a group member



## AddMember

> AddMember(ctx, groupAddress).MemberAddress(memberAddress).Execute()

Add a group member

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "go.opscenter.dev/james-go-client"
)

func main() {
    groupAddress := "group@domain.com" // string | Group mail address
    memberAddress := "member@domain.com" // string | Member mail address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressGroupAPI.AddMember(context.Background(), groupAddress).MemberAddress(memberAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressGroupAPI.AddMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupAddress** | **string** | Group mail address | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberAddress** | **string** | Member mail address | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroups

> []string ListGroups(ctx).Execute()

List address groups

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "go.opscenter.dev/james-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressGroupAPI.ListGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressGroupAPI.ListGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `AddressGroupAPI.ListGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMembers

> []string ListMembers(ctx, groupAddress).Execute()

List members of a group

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "go.opscenter.dev/james-go-client"
)

func main() {
    groupAddress := "group@domain.com" // string | Group mail address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressGroupAPI.ListMembers(context.Background(), groupAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressGroupAPI.ListMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMembers`: []string
    fmt.Fprintf(os.Stdout, "Response from `AddressGroupAPI.ListMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupAddress** | **string** | Group mail address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMember

> RemoveMember(ctx, groupAddress).MemberAddress(memberAddress).Execute()

Remove a group member

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "go.opscenter.dev/james-go-client"
)

func main() {
    groupAddress := "group@domain.com" // string | Group mail address
    memberAddress := "member@domain.com" // string | Member mail address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressGroupAPI.RemoveMember(context.Background(), groupAddress).MemberAddress(memberAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressGroupAPI.RemoveMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupAddress** | **string** | Group mail address | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **memberAddress** | **string** | Member mail address | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

