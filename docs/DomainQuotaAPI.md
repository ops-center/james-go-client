# \DomainQuotaAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDomainQuotaCount**](DomainQuotaAPI.md#DeleteDomainQuotaCount) | **Delete** /quota/domains/{domainToBeUsed}/count | Delete the quota count for a domain
[**DeleteDomainQuotaSize**](DomainQuotaAPI.md#DeleteDomainQuotaSize) | **Delete** /quota/domains/{domainToBeUsed}/size | Delete the quota size for a domain
[**GetDomainQuota**](DomainQuotaAPI.md#GetDomainQuota) | **Get** /quota/domains/{domainToBeUsed} | Get the quota for a domain
[**GetDomainQuotaCount**](DomainQuotaAPI.md#GetDomainQuotaCount) | **Get** /quota/domains/{domainToBeUsed}/count | Get the quota count for a domain
[**GetDomainQuotaSize**](DomainQuotaAPI.md#GetDomainQuotaSize) | **Get** /quota/domains/{domainToBeUsed}/size | Get the quota size for a domain
[**UpdateDomainQuota**](DomainQuotaAPI.md#UpdateDomainQuota) | **Put** /quota/domains/{domainToBeUsed} | Update the quota for a domain
[**UpdateDomainQuotaCount**](DomainQuotaAPI.md#UpdateDomainQuotaCount) | **Put** /quota/domains/{domainToBeUsed}/count | Update the quota count for a domain
[**UpdateDomainQuotaSize**](DomainQuotaAPI.md#UpdateDomainQuotaSize) | **Put** /quota/domains/{domainToBeUsed}/size | Update the quota size for a domain



## DeleteDomainQuotaCount

> DeleteDomainQuotaCount(ctx, domainToBeUsed).Execute()

Delete the quota count for a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainQuotaAPI.DeleteDomainQuotaCount(context.Background(), domainToBeUsed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainQuotaAPI.DeleteDomainQuotaCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainQuotaCountRequest struct via the builder pattern


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


## DeleteDomainQuotaSize

> DeleteDomainQuotaSize(ctx, domainToBeUsed).Execute()

Delete the quota size for a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainQuotaAPI.DeleteDomainQuotaSize(context.Background(), domainToBeUsed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainQuotaAPI.DeleteDomainQuotaSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainQuotaSizeRequest struct via the builder pattern


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


## GetDomainQuota

> GetDomainQuota(ctx, domainToBeUsed).Execute()

Get the quota for a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainQuotaAPI.GetDomainQuota(context.Background(), domainToBeUsed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainQuotaAPI.GetDomainQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainQuotaRequest struct via the builder pattern


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


## GetDomainQuotaCount

> GetDomainQuotaCount(ctx, domainToBeUsed).Execute()

Get the quota count for a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainQuotaAPI.GetDomainQuotaCount(context.Background(), domainToBeUsed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainQuotaAPI.GetDomainQuotaCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainQuotaCountRequest struct via the builder pattern


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


## GetDomainQuotaSize

> GetDomainQuotaSize(ctx, domainToBeUsed).Execute()

Get the quota size for a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainQuotaAPI.GetDomainQuotaSize(context.Background(), domainToBeUsed).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainQuotaAPI.GetDomainQuotaSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDomainQuotaSizeRequest struct via the builder pattern


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


## UpdateDomainQuota

> UpdateDomainQuota(ctx, domainToBeUsed).UpdateGlobalQuotaRequest(updateGlobalQuotaRequest).Execute()

Update the quota for a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | 
    updateGlobalQuotaRequest := *openapiclient.NewUpdateGlobalQuotaRequest() // UpdateGlobalQuotaRequest | The new quota for the domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainQuotaAPI.UpdateDomainQuota(context.Background(), domainToBeUsed).UpdateGlobalQuotaRequest(updateGlobalQuotaRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainQuotaAPI.UpdateDomainQuota``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainQuotaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGlobalQuotaRequest** | [**UpdateGlobalQuotaRequest**](UpdateGlobalQuotaRequest.md) | The new quota for the domain | 

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


## UpdateDomainQuotaCount

> UpdateDomainQuotaCount(ctx, domainToBeUsed).Body(body).Execute()

Update the quota count for a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | 
    body := int32(56) // int32 | The new quota count for the domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainQuotaAPI.UpdateDomainQuotaCount(context.Background(), domainToBeUsed).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainQuotaAPI.UpdateDomainQuotaCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainQuotaCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **int32** | The new quota count for the domain | 

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


## UpdateDomainQuotaSize

> UpdateDomainQuotaSize(ctx, domainToBeUsed).Body(body).Execute()

Update the quota size for a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | 
    body := int32(56) // int32 | The new quota size for the domain

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainQuotaAPI.UpdateDomainQuotaSize(context.Background(), domainToBeUsed).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainQuotaAPI.UpdateDomainQuotaSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDomainQuotaSizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **int32** | The new quota size for the domain | 

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

