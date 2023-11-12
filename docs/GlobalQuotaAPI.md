# \GlobalQuotaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteGlobalQuotaCount**](GlobalQuotaAPI.md#DeleteGlobalQuotaCount) | **Delete** /quota/count | Delete the global quota count
[**DeleteGlobalQuotaSize**](GlobalQuotaAPI.md#DeleteGlobalQuotaSize) | **Delete** /quota/size | Delete the global quota size
[**GetGlobalQuota**](GlobalQuotaAPI.md#GetGlobalQuota) | **Get** /quota | Get the global quota
[**GetGlobalQuotaCount**](GlobalQuotaAPI.md#GetGlobalQuotaCount) | **Get** /quota/count | Get the global quota count
[**GetGlobalQuotaSize**](GlobalQuotaAPI.md#GetGlobalQuotaSize) | **Get** /quota/size | Get the global quota size
[**UpdateGlobalQuota**](GlobalQuotaAPI.md#UpdateGlobalQuota) | **Put** /quota | Update the global quota
[**UpdateGlobalQuotaCount**](GlobalQuotaAPI.md#UpdateGlobalQuotaCount) | **Put** /quota/count | Update the global quota count
[**UpdateGlobalQuotaSize**](GlobalQuotaAPI.md#UpdateGlobalQuotaSize) | **Put** /quota/size | Update the global quota size



## DeleteGlobalQuotaCount

> DeleteGlobalQuotaCount(ctx).Execute()

Delete the global quota count

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
    r, err := apiClient.GlobalQuotaAPI.DeleteGlobalQuotaCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalQuotaAPI.DeleteGlobalQuotaCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalQuotaCountRequest struct via the builder pattern


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


## DeleteGlobalQuotaSize

> DeleteGlobalQuotaSize(ctx).Execute()

Delete the global quota size

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
    r, err := apiClient.GlobalQuotaAPI.DeleteGlobalQuotaSize(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalQuotaAPI.DeleteGlobalQuotaSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGlobalQuotaSizeRequest struct via the builder pattern


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


## GetGlobalQuota

> GetGlobalQuota200Response GetGlobalQuota(ctx).Execute()

Get the global quota

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
    resp, r, err := apiClient.GlobalQuotaAPI.GetGlobalQuota(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalQuotaAPI.GetGlobalQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalQuota`: GetGlobalQuota200Response
    fmt.Fprintf(os.Stdout, "Response from `GlobalQuotaAPI.GetGlobalQuota`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalQuotaRequest struct via the builder pattern


### Return type

[**GetGlobalQuota200Response**](GetGlobalQuota200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalQuotaCount

> int32 GetGlobalQuotaCount(ctx).Execute()

Get the global quota count

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
    resp, r, err := apiClient.GlobalQuotaAPI.GetGlobalQuotaCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalQuotaAPI.GetGlobalQuotaCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalQuotaCount`: int32
    fmt.Fprintf(os.Stdout, "Response from `GlobalQuotaAPI.GetGlobalQuotaCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalQuotaCountRequest struct via the builder pattern


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGlobalQuotaSize

> int32 GetGlobalQuotaSize(ctx).Execute()

Get the global quota size

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
    resp, r, err := apiClient.GlobalQuotaAPI.GetGlobalQuotaSize(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalQuotaAPI.GetGlobalQuotaSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGlobalQuotaSize`: int32
    fmt.Fprintf(os.Stdout, "Response from `GlobalQuotaAPI.GetGlobalQuotaSize`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGlobalQuotaSizeRequest struct via the builder pattern


### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGlobalQuota

> UpdateGlobalQuota(ctx).UpdateGlobalQuotaRequest(updateGlobalQuotaRequest).Execute()

Update the global quota

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
    updateGlobalQuotaRequest := *openapiclient.NewUpdateGlobalQuotaRequest() // UpdateGlobalQuotaRequest | The new global quota

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GlobalQuotaAPI.UpdateGlobalQuota(context.Background()).UpdateGlobalQuotaRequest(updateGlobalQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalQuotaAPI.UpdateGlobalQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateGlobalQuotaRequest** | [**UpdateGlobalQuotaRequest**](UpdateGlobalQuotaRequest.md) | The new global quota | 

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


## UpdateGlobalQuotaCount

> UpdateGlobalQuotaCount(ctx).Body(body).Execute()

Update the global quota count

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
    body := int32(56) // int32 | The new global quota count

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GlobalQuotaAPI.UpdateGlobalQuotaCount(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalQuotaAPI.UpdateGlobalQuotaCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalQuotaCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **int32** | The new global quota count | 

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


## UpdateGlobalQuotaSize

> UpdateGlobalQuotaSize(ctx).Body(body).Execute()

Update the global quota size

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
    body := int32(56) // int32 | The new global quota size

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.GlobalQuotaAPI.UpdateGlobalQuotaSize(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GlobalQuotaAPI.UpdateGlobalQuotaSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGlobalQuotaSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **int32** | The new global quota size | 

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

