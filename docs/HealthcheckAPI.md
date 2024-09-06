# \HealthcheckAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckAllComponents**](HealthcheckAPI.md#CheckAllComponents) | **Get** /healthcheck | Check all components
[**CheckComponent**](HealthcheckAPI.md#CheckComponent) | **Get** /healthcheck/checks/{componentName} | Check single component
[**ListAllHealthChecks**](HealthcheckAPI.md#ListAllHealthChecks) | **Get** /healthcheck/checks | List all health checks



## CheckAllComponents

> CheckAllComponents200Response CheckAllComponents(ctx).Execute()

Check all components

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
    resp, r, err := apiClient.HealthcheckAPI.CheckAllComponents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPI.CheckAllComponents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckAllComponents`: CheckAllComponents200Response
    fmt.Fprintf(os.Stdout, "Response from `HealthcheckAPI.CheckAllComponents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCheckAllComponentsRequest struct via the builder pattern


### Return type

[**CheckAllComponents200Response**](CheckAllComponents200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckComponent

> HealthCheckResult CheckComponent(ctx, componentName).Execute()

Check single component

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
    componentName := "componentName_example" // string | URL encoded name of the component

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.HealthcheckAPI.CheckComponent(context.Background(), componentName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPI.CheckComponent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CheckComponent`: HealthCheckResult
    fmt.Fprintf(os.Stdout, "Response from `HealthcheckAPI.CheckComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentName** | **string** | URL encoded name of the component | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**HealthCheckResult**](HealthCheckResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAllHealthChecks

> []HealthCheckInfo ListAllHealthChecks(ctx).Execute()

List all health checks

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
    resp, r, err := apiClient.HealthcheckAPI.ListAllHealthChecks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `HealthcheckAPI.ListAllHealthChecks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAllHealthChecks`: []HealthCheckInfo
    fmt.Fprintf(os.Stdout, "Response from `HealthcheckAPI.ListAllHealthChecks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAllHealthChecksRequest struct via the builder pattern


### Return type

[**[]HealthCheckInfo**](HealthCheckInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

