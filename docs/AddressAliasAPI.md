# \AddressAliasAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAlias**](AddressAliasAPI.md#CreateAlias) | **Put** /address/aliases/{userAddress} | Add a new alias to a user
[**DeleteAlias**](AddressAliasAPI.md#DeleteAlias) | **Delete** /address/aliases/{userAddress} | Remove an alias from a user
[**GetAlias**](AddressAliasAPI.md#GetAlias) | **Get** /address/aliases/{userAddress} | List alias sources of a user
[**ListAliases**](AddressAliasAPI.md#ListAliases) | **Get** /address/aliases | List users with aliases



## CreateAlias

> CreateAlias(ctx, userAddress).SourceAddress(sourceAddress).Execute()

Add a new alias to a user

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
    userAddress := "user@domain.com" // string | User mail address
    sourceAddress := "alias@domain.com" // string | Alias source mail address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressAliasAPI.CreateAlias(context.Background(), userAddress).SourceAddress(sourceAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAliasAPI.CreateAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAddress** | **string** | User mail address | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceAddress** | **string** | Alias source mail address | 

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


## DeleteAlias

> DeleteAlias(ctx, userAddress).SourceAddress(sourceAddress).Execute()

Remove an alias from a user

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
    userAddress := "user@domain.com" // string | User mail address
    sourceAddress := "alias@domain.com" // string | Alias source mail address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressAliasAPI.DeleteAlias(context.Background(), userAddress).SourceAddress(sourceAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAliasAPI.DeleteAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAddress** | **string** | User mail address | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceAddress** | **string** | Alias source mail address | 

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


## GetAlias

> []GetAlias200ResponseInner GetAlias(ctx, userAddress).Execute()

List alias sources of a user

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
    userAddress := "user@domain.com" // string | User mail address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AddressAliasAPI.GetAlias(context.Background(), userAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAliasAPI.GetAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAlias`: []GetAlias200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AddressAliasAPI.GetAlias`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAddress** | **string** | User mail address | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetAlias200ResponseInner**](GetAlias200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAliases

> []string ListAliases(ctx).Execute()

List users with aliases

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
    resp, r, err := apiClient.AddressAliasAPI.ListAliases(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressAliasAPI.ListAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAliases`: []string
    fmt.Fprintf(os.Stdout, "Response from `AddressAliasAPI.ListAliases`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAliasesRequest struct via the builder pattern


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

