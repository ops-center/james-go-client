# \AddressForwardAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDestination**](AddressForwardAPI.md#AddDestination) | **Put** /address/forwards/{userAddress} | Add a new destination to a forward
[**DeleteDestination**](AddressForwardAPI.md#DeleteDestination) | **Delete** /address/forwards/{userAddress} | Remove a destination from a forward
[**ListDestinations**](AddressForwardAPI.md#ListDestinations) | **Get** /address/forwards/{userAddress} | List destinations in a forward
[**ListForwards**](AddressForwardAPI.md#ListForwards) | **Get** /address/forwards | List address forwards



## AddDestination

> AddDestination(ctx, userAddress).DestinationAddress(destinationAddress).Execute()

Add a new destination to a forward

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
    destinationAddress := "destination@domain.com" // string | Destination mail address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressForwardAPI.AddDestination(context.Background(), userAddress).DestinationAddress(destinationAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressForwardAPI.AddDestination``: %v\n", err)
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

Other parameters are passed through a pointer to a apiAddDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationAddress** | **string** | Destination mail address | 

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


## DeleteDestination

> DeleteDestination(ctx, userAddress).DestinationAddress(destinationAddress).Execute()

Remove a destination from a forward

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
    destinationAddress := "destination@domain.com" // string | Destination mail address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressForwardAPI.DeleteDestination(context.Background(), userAddress).DestinationAddress(destinationAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressForwardAPI.DeleteDestination``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDestinationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **destinationAddress** | **string** | Destination mail address | 

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


## ListDestinations

> []ListDestinations200ResponseInner ListDestinations(ctx, userAddress).Execute()

List destinations in a forward

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
    resp, r, err := apiClient.AddressForwardAPI.ListDestinations(context.Background(), userAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressForwardAPI.ListDestinations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDestinations`: []ListDestinations200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `AddressForwardAPI.ListDestinations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAddress** | **string** | User mail address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDestinationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListDestinations200ResponseInner**](ListDestinations200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListForwards

> []string ListForwards(ctx).Execute()

List address forwards

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
    resp, r, err := apiClient.AddressForwardAPI.ListForwards(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressForwardAPI.ListForwards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListForwards`: []string
    fmt.Fprintf(os.Stdout, "Response from `AddressForwardAPI.ListForwards`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListForwardsRequest struct via the builder pattern


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

