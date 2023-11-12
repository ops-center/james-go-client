# \UserQuotaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteQuotaCount**](UserQuotaAPI.md#DeleteQuotaCount) | **Delete** /quota/users/{username}/count | Delete the quota count for a user
[**DeleteQuotaSize**](UserQuotaAPI.md#DeleteQuotaSize) | **Delete** /quota/users/{username}/size | Delete the quota size for a user
[**GetQuota**](UserQuotaAPI.md#GetQuota) | **Get** /quota/users/{username} | Get the quota for a user
[**GetQuotaCount**](UserQuotaAPI.md#GetQuotaCount) | **Get** /quota/users/{username}/count | Get the quota count for a user
[**GetQuotaSize**](UserQuotaAPI.md#GetQuotaSize) | **Get** /quota/users/{username}/size | Get the quota size for a user
[**RecomputeCurrentQuotas**](UserQuotaAPI.md#RecomputeCurrentQuotas) | **Post** /quota/users | Recompute current quotas for users
[**SearchByQuota**](UserQuotaAPI.md#SearchByQuota) | **Get** /quota/users | Search users by quota ratio
[**UpdateQuota**](UserQuotaAPI.md#UpdateQuota) | **Put** /quota/users/{username} | Update the quota for a user
[**UpdateQuotaCount**](UserQuotaAPI.md#UpdateQuotaCount) | **Put** /quota/users/{username}/count | Update the quota count for a user
[**UpdateQuotaSize**](UserQuotaAPI.md#UpdateQuotaSize) | **Put** /quota/users/{username}/size | Update the quota size for a user



## DeleteQuotaCount

> DeleteQuotaCount(ctx, username).Execute()

Delete the quota count for a user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.DeleteQuotaCount(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.DeleteQuotaCount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteQuotaCountRequest struct via the builder pattern


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


## DeleteQuotaSize

> DeleteQuotaSize(ctx, username).Execute()

Delete the quota size for a user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.DeleteQuotaSize(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.DeleteQuotaSize``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteQuotaSizeRequest struct via the builder pattern


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


## GetQuota

> GetQuota(ctx, username).Execute()

Get the quota for a user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.GetQuota(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.GetQuota``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetQuotaRequest struct via the builder pattern


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


## GetQuotaCount

> GetQuotaCount(ctx, username).Execute()

Get the quota count for a user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.GetQuotaCount(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.GetQuotaCount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetQuotaCountRequest struct via the builder pattern


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


## GetQuotaSize

> GetQuotaSize(ctx, username).Execute()

Get the quota size for a user

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.GetQuotaSize(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.GetQuotaSize``: %v\n", err)
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

Other parameters are passed through a pointer to a apiGetQuotaSizeRequest struct via the builder pattern


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


## RecomputeCurrentQuotas

> RecomputeCurrentQuotas(ctx).Task(task).RecomputeCurrentQuotasRequest(recomputeCurrentQuotasRequest).Execute()

Recompute current quotas for users

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
    task := "task_example" // string | 
    recomputeCurrentQuotasRequest := *openapiclient.NewRecomputeCurrentQuotasRequest() // RecomputeCurrentQuotasRequest | Additional task options (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.RecomputeCurrentQuotas(context.Background()).Task(task).RecomputeCurrentQuotasRequest(recomputeCurrentQuotasRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.RecomputeCurrentQuotas``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRecomputeCurrentQuotasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **task** | **string** |  | 
 **recomputeCurrentQuotasRequest** | [**RecomputeCurrentQuotasRequest**](RecomputeCurrentQuotasRequest.md) | Additional task options | 

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


## SearchByQuota

> SearchByQuota(ctx).MinOccupationRatio(minOccupationRatio).MaxOccupationRatio(maxOccupationRatio).Limit(limit).Offset(offset).Domain(domain).Execute()

Search users by quota ratio

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
    minOccupationRatio := float32(8.14) // float32 | The minimum occupation ratio of users to be returned (optional)
    maxOccupationRatio := float32(8.14) // float32 | The maximum occupation ratio of users to be returned (optional)
    limit := int32(56) // int32 | The maximum number of users to be returned (optional)
    offset := int32(56) // int32 | The number of users to skip (optional)
    domain := "domain_example" // string | The domain of users to be returned (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.SearchByQuota(context.Background()).MinOccupationRatio(minOccupationRatio).MaxOccupationRatio(maxOccupationRatio).Limit(limit).Offset(offset).Domain(domain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.SearchByQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchByQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **minOccupationRatio** | **float32** | The minimum occupation ratio of users to be returned | 
 **maxOccupationRatio** | **float32** | The maximum occupation ratio of users to be returned | 
 **limit** | **int32** | The maximum number of users to be returned | 
 **offset** | **int32** | The number of users to skip | 
 **domain** | **string** | The domain of users to be returned | 

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


## UpdateQuota

> UpdateQuota(ctx, username).UpdateGlobalQuotaRequest(updateGlobalQuotaRequest).Execute()

Update the quota for a user

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
    updateGlobalQuotaRequest := *openapiclient.NewUpdateGlobalQuotaRequest() // UpdateGlobalQuotaRequest | The new quota for the user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.UpdateQuota(context.Background(), username).UpdateGlobalQuotaRequest(updateGlobalQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.UpdateQuota``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGlobalQuotaRequest** | [**UpdateGlobalQuotaRequest**](UpdateGlobalQuotaRequest.md) | The new quota for the user | 

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


## UpdateQuotaCount

> UpdateQuotaCount(ctx, username).Body(body).Execute()

Update the quota count for a user

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
    body := int32(56) // int32 | The new quota count for the user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.UpdateQuotaCount(context.Background(), username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.UpdateQuotaCount``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateQuotaCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **int32** | The new quota count for the user | 

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


## UpdateQuotaSize

> UpdateQuotaSize(ctx, username).Body(body).Execute()

Update the quota size for a user

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
    body := int32(56) // int32 | The new quota size for the user

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserQuotaAPI.UpdateQuotaSize(context.Background(), username).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserQuotaAPI.UpdateQuotaSize``: %v\n", err)
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

Other parameters are passed through a pointer to a apiUpdateQuotaSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **int32** | The new quota size for the user | 

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

