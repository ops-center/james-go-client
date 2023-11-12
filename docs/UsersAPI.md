# \UsersAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDelegatedUser**](UsersAPI.md#AddDelegatedUser) | **Put** /users/{baseUser}/authorizedUsers/delegatedUser | Add a delegated user to a base user
[**ChangeUsername**](UsersAPI.md#ChangeUsername) | **Post** /users/{oldUser}/rename/{newUser} | Change a username
[**CreateUserIdentity**](UsersAPI.md#CreateUserIdentity) | **Post** /users/{username}/identities | Create a JMAP user identity
[**DeleteUser**](UsersAPI.md#DeleteUser) | **Delete** /users/{username} | Delete a user
[**ExistsUser**](UsersAPI.md#ExistsUser) | **Head** /users/{username} | Test user existence
[**ListAllowedFromHeaders**](UsersAPI.md#ListAllowedFromHeaders) | **Get** /users/{givenUser}/allowedFromHeaders | Retrieve the list of allowed From headers for a given user
[**ListDelegatedUsers**](UsersAPI.md#ListDelegatedUsers) | **Get** /users/{baseUser}/authorizedUsers | Retrieve the list of delegated users of a base user
[**ListUserIdentities**](UsersAPI.md#ListUserIdentities) | **Get** /users/{username}/identities | Retrieve the user identities
[**ListUsers**](UsersAPI.md#ListUsers) | **Get** /users | Retrieve the user list
[**RemoveAllDelegatedUsers**](UsersAPI.md#RemoveAllDelegatedUsers) | **Delete** /users/{baseUser}/authorizedUsers | Remove all delegated users of a base user
[**RemoveDelegatedUser**](UsersAPI.md#RemoveDelegatedUser) | **Delete** /users/{baseUser}/authorizedUsers/delegatedUser | Remove a delegated user from a base user
[**UpdateUserIdentity**](UsersAPI.md#UpdateUserIdentity) | **Put** /users/{username}/identities/{identityId} | Update a JMAP user identity
[**UpsertUser**](UsersAPI.md#UpsertUser) | **Put** /users/{username} | Create or Update User



## AddDelegatedUser

> AddDelegatedUser(ctx, baseUser, delegatedUser).Execute()

Add a delegated user to a base user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    baseUser := "baseUser_example" // string | 
    delegatedUser := "delegatedUser_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.AddDelegatedUser(context.Background(), baseUser, delegatedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.AddDelegatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**baseUser** | **string** |  | 
**delegatedUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDelegatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## ChangeUsername

> ChangeUsername(ctx, oldUser, newUser).Action(action).Execute()

Change a username

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    oldUser := "oldUser_example" // string | 
    newUser := "newUser_example" // string | 
    action := "action_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.ChangeUsername(context.Background(), oldUser, newUser).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ChangeUsername``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**oldUser** | **string** |  | 
**newUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiChangeUsernameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **action** | **string** |  | 

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


## CreateUserIdentity

> CreateUserIdentity(ctx, username).CreateUserIdentityRequest(createUserIdentityRequest).Execute()

Create a JMAP user identity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    username := "username_example" // string | 
    createUserIdentityRequest := *openapiclient.NewCreateUserIdentityRequest() // CreateUserIdentityRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.CreateUserIdentity(context.Background(), username).CreateUserIdentityRequest(createUserIdentityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.CreateUserIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUserIdentityRequest** | [**CreateUserIdentityRequest**](CreateUserIdentityRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUser

> DeleteUser(ctx, username).Execute()

Delete a user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    username := "username_example" // string | The username to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.DeleteUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.DeleteUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ExistsUser

> ExistsUser(ctx, username).Execute()

Test user existence

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    username := "username_example" // string | The username to be tested

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.ExistsUser(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ExistsUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username to be tested | 

### Other Parameters

Other parameters are passed through a pointer to a apiExistsUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListAllowedFromHeaders

> []string ListAllowedFromHeaders(ctx, givenUser).Execute()

Retrieve the list of allowed From headers for a given user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    givenUser := "givenUser_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.ListAllowedFromHeaders(context.Background(), givenUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListAllowedFromHeaders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllowedFromHeaders`: []string
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListAllowedFromHeaders`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**givenUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListAllowedFromHeadersRequest struct via the builder pattern


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


## ListDelegatedUsers

> []string ListDelegatedUsers(ctx, baseUser).Execute()

Retrieve the list of delegated users of a base user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    baseUser := "baseUser_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.ListDelegatedUsers(context.Background(), baseUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListDelegatedUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDelegatedUsers`: []string
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListDelegatedUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**baseUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDelegatedUsersRequest struct via the builder pattern


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


## ListUserIdentities

> []ListUserIdentities200ResponseInner ListUserIdentities(ctx, username).Default_(default_).Execute()

Retrieve the user identities

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    username := "username_example" // string | 
    default_ := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.ListUserIdentities(context.Background(), username).Default_(default_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUserIdentities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserIdentities`: []ListUserIdentities200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUserIdentities`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserIdentitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **default_** | **bool** |  | 

### Return type

[**[]ListUserIdentities200ResponseInner**](ListUserIdentities200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsers

> []ListUsers200ResponseInner ListUsers(ctx).Execute()

Retrieve the user list

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UsersAPI.ListUsers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.ListUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsers`: []ListUsers200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UsersAPI.ListUsers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersRequest struct via the builder pattern


### Return type

[**[]ListUsers200ResponseInner**](ListUsers200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAllDelegatedUsers

> RemoveAllDelegatedUsers(ctx, baseUser).Execute()

Remove all delegated users of a base user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    baseUser := "baseUser_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.RemoveAllDelegatedUsers(context.Background(), baseUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveAllDelegatedUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**baseUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAllDelegatedUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## RemoveDelegatedUser

> RemoveDelegatedUser(ctx, baseUser, delegatedUser).Execute()

Remove a delegated user from a base user

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    baseUser := "baseUser_example" // string | 
    delegatedUser := "delegatedUser_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.RemoveDelegatedUser(context.Background(), baseUser, delegatedUser).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.RemoveDelegatedUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**baseUser** | **string** |  | 
**delegatedUser** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDelegatedUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## UpdateUserIdentity

> UpdateUserIdentity(ctx, username, identityId).UpdateUserIdentityRequest(updateUserIdentityRequest).Execute()

Update a JMAP user identity

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    username := "username_example" // string | 
    identityId := "identityId_example" // string | 
    updateUserIdentityRequest := *openapiclient.NewUpdateUserIdentityRequest() // UpdateUserIdentityRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.UpdateUserIdentity(context.Background(), username, identityId).UpdateUserIdentityRequest(updateUserIdentityRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpdateUserIdentity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**identityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserIdentityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateUserIdentityRequest** | [**UpdateUserIdentityRequest**](UpdateUserIdentityRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertUser

> UpsertUser(ctx, username).UpsertUserRequest(upsertUserRequest).Force(force).Execute()

Create or Update User

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    username := "username_example" // string | The username of the user to be created or updated
    upsertUserRequest := *openapiclient.NewUpsertUserRequest("Password_example") // UpsertUserRequest | 
    force := true // bool | Indicates whether to force update the password (for updating password only) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UsersAPI.UpsertUser(context.Background(), username).UpsertUserRequest(upsertUserRequest).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UpsertUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** | The username of the user to be created or updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **upsertUserRequest** | [**UpsertUserRequest**](UpsertUserRequest.md) |  | 
 **force** | **bool** | Indicates whether to force update the password (for updating password only) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

