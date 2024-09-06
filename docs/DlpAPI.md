# \DlpAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FetchDLPConfiguration**](DlpAPI.md#FetchDLPConfiguration) | **Get** /dlp/rules/{senderDomain}/rules/{ruleId} | Fetch a DLP configuration item by sender domain and rule id
[**ListDLPConfiguration**](DlpAPI.md#ListDLPConfiguration) | **Get** /dlp/rules/{senderDomain} | List DLP configuration by sender domain
[**RemoveDLPConfiguration**](DlpAPI.md#RemoveDLPConfiguration) | **Delete** /dlp/rules/{senderDomain} | Remove DLP configuration by sender domain
[**StoreDLPConfiguration**](DlpAPI.md#StoreDLPConfiguration) | **Put** /dlp/rules/{senderDomain} | Store DLP configuration by sender domain



## FetchDLPConfiguration

> FetchDLPConfiguration(ctx, senderDomain, ruleId).Execute()

Fetch a DLP configuration item by sender domain and rule id

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
    senderDomain := "senderDomain_example" // string | 
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DlpAPI.FetchDLPConfiguration(context.Background(), senderDomain, ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DlpAPI.FetchDLPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**senderDomain** | **string** |  | 
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiFetchDLPConfigurationRequest struct via the builder pattern


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


## ListDLPConfiguration

> ListDLPConfiguration(ctx, senderDomain).Execute()

List DLP configuration by sender domain

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
    senderDomain := "senderDomain_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DlpAPI.ListDLPConfiguration(context.Background(), senderDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DlpAPI.ListDLPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**senderDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDLPConfigurationRequest struct via the builder pattern


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


## RemoveDLPConfiguration

> RemoveDLPConfiguration(ctx, senderDomain).Execute()

Remove DLP configuration by sender domain

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
    senderDomain := "senderDomain_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DlpAPI.RemoveDLPConfiguration(context.Background(), senderDomain).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DlpAPI.RemoveDLPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**senderDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDLPConfigurationRequest struct via the builder pattern


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


## StoreDLPConfiguration

> StoreDLPConfiguration(ctx, senderDomain).StoreDLPConfigurationRequest(storeDLPConfigurationRequest).Execute()

Store DLP configuration by sender domain

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
    senderDomain := "senderDomain_example" // string | 
    storeDLPConfigurationRequest := *openapiclient.NewStoreDLPConfigurationRequest() // StoreDLPConfigurationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DlpAPI.StoreDLPConfiguration(context.Background(), senderDomain).StoreDLPConfigurationRequest(storeDLPConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DlpAPI.StoreDLPConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**senderDomain** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreDLPConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **storeDLPConfigurationRequest** | [**StoreDLPConfigurationRequest**](StoreDLPConfigurationRequest.md) |  | 

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

