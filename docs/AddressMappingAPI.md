# \AddressMappingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAddressMapping**](AddressMappingAPI.md#AddAddressMapping) | **Post** /mappings | Add an address mapping
[**ListAddressMappings**](AddressMappingAPI.md#ListAddressMappings) | **Get** /mappings | List all address mappings
[**RemoveAddressMapping**](AddressMappingAPI.md#RemoveAddressMapping) | **Delete** /mappings/address/{mappingSource}/targets/{destinationAddress} | Remove an address mapping



## AddAddressMapping

> AddAddressMapping(ctx, mappingSource, destinationAddress).Execute()

Add an address mapping

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
    mappingSource := "alias@domain.tld" // string | Mapping source (recipient address)
    destinationAddress := "user@domain.tld" // string | Mapping destination (address)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressMappingAPI.AddAddressMapping(context.Background(), mappingSource, destinationAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressMappingAPI.AddAddressMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingSource** | **string** | Mapping source (recipient address) | 
**destinationAddress** | **string** | Mapping destination (address) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAddressMappingRequest struct via the builder pattern


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


## ListAddressMappings

> ListAddressMappings200Response ListAddressMappings(ctx).Execute()

List all address mappings

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
    resp, r, err := apiClient.AddressMappingAPI.ListAddressMappings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressMappingAPI.ListAddressMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAddressMappings`: ListAddressMappings200Response
    fmt.Fprintf(os.Stdout, "Response from `AddressMappingAPI.ListAddressMappings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAddressMappingsRequest struct via the builder pattern


### Return type

[**ListAddressMappings200Response**](ListAddressMappings200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAddressMapping

> RemoveAddressMapping(ctx, mappingSource, destinationAddress).Execute()

Remove an address mapping

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
    mappingSource := "alias@domain.tld" // string | Mapping source (recipient address)
    destinationAddress := "user@domain.tld" // string | Mapping destination (address)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.AddressMappingAPI.RemoveAddressMapping(context.Background(), mappingSource, destinationAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AddressMappingAPI.RemoveAddressMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingSource** | **string** | Mapping source (recipient address) | 
**destinationAddress** | **string** | Mapping destination (address) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAddressMappingRequest struct via the builder pattern


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

