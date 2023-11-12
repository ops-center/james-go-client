# \SieveQuotaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSieveQuota**](SieveQuotaAPI.md#GetUserSieveQuota) | **Get** /sieve/quota/users/{userEmail} | Retrieve user sieve quota
[**RemoveUserSieveQuota**](SieveQuotaAPI.md#RemoveUserSieveQuota) | **Delete** /sieve/quota/users/{userEmail} | Remove user sieve quota
[**SieveQuotaDefaultDelete**](SieveQuotaAPI.md#SieveQuotaDefaultDelete) | **Delete** /sieve/quota/default | Remove global sieve quota
[**SieveQuotaDefaultGet**](SieveQuotaAPI.md#SieveQuotaDefaultGet) | **Get** /sieve/quota/default | Retrieve global sieve quota
[**UpdateGlobalSieveQuota**](SieveQuotaAPI.md#UpdateGlobalSieveQuota) | **Put** /sieve/quota/default | Update global sieve quota
[**UpdateUserSieveQuota**](SieveQuotaAPI.md#UpdateUserSieveQuota) | **Put** /sieve/quota/users/{userEmail} | Update user sieve quota



## GetUserSieveQuota

> GetUserSieveQuota(ctx, userEmail).Execute()

Retrieve user sieve quota

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
    userEmail := "userEmail_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SieveQuotaAPI.GetUserSieveQuota(context.Background(), userEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SieveQuotaAPI.GetUserSieveQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEmail** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSieveQuotaRequest struct via the builder pattern


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


## RemoveUserSieveQuota

> RemoveUserSieveQuota(ctx, userEmail).Execute()

Remove user sieve quota

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
    userEmail := "userEmail_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SieveQuotaAPI.RemoveUserSieveQuota(context.Background(), userEmail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SieveQuotaAPI.RemoveUserSieveQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEmail** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserSieveQuotaRequest struct via the builder pattern


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


## SieveQuotaDefaultDelete

> SieveQuotaDefaultDelete(ctx).Execute()

Remove global sieve quota

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
    r, err := apiClient.SieveQuotaAPI.SieveQuotaDefaultDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SieveQuotaAPI.SieveQuotaDefaultDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSieveQuotaDefaultDeleteRequest struct via the builder pattern


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


## SieveQuotaDefaultGet

> SieveQuotaDefaultGet(ctx).Execute()

Retrieve global sieve quota

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
    r, err := apiClient.SieveQuotaAPI.SieveQuotaDefaultGet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SieveQuotaAPI.SieveQuotaDefaultGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiSieveQuotaDefaultGetRequest struct via the builder pattern


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


## UpdateGlobalSieveQuota

> UpdateGlobalSieveQuota(ctx).UpdateGlobalSieveQuotaRequest(updateGlobalSieveQuotaRequest).Execute()

Update global sieve quota

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
    updateGlobalSieveQuotaRequest := *openapiclient.NewUpdateGlobalSieveQuotaRequest() // UpdateGlobalSieveQuotaRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SieveQuotaAPI.UpdateGlobalSieveQuota(context.Background()).UpdateGlobalSieveQuotaRequest(updateGlobalSieveQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SieveQuotaAPI.UpdateGlobalSieveQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalSieveQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGlobalSieveQuotaRequest** | [**UpdateGlobalSieveQuotaRequest**](UpdateGlobalSieveQuotaRequest.md) |  | 

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


## UpdateUserSieveQuota

> UpdateUserSieveQuota(ctx, userEmail).UpdateUserSieveQuotaRequest(updateUserSieveQuotaRequest).Execute()

Update user sieve quota

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
    userEmail := "userEmail_example" // string | 
    updateUserSieveQuotaRequest := *openapiclient.NewUpdateUserSieveQuotaRequest() // UpdateUserSieveQuotaRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.SieveQuotaAPI.UpdateUserSieveQuota(context.Background(), userEmail).UpdateUserSieveQuotaRequest(updateUserSieveQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SieveQuotaAPI.UpdateUserSieveQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userEmail** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserSieveQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUserSieveQuotaRequest** | [**UpdateUserSieveQuotaRequest**](UpdateUserSieveQuotaRequest.md) |  | 

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

