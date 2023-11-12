# \UserMappingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUserMappings**](UserMappingAPI.md#ListUserMappings) | **Get** /mappings/user/{userAddress} | Listing User Mappings



## ListUserMappings

> []ListUserMappings200ResponseInner ListUserMappings(ctx, userAddress).Execute()

Listing User Mappings

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
    userAddress := "user123@domain.tld" // string | User address

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMappingAPI.ListUserMappings(context.Background(), userAddress).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMappingAPI.ListUserMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUserMappings`: []ListUserMappings200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserMappingAPI.ListUserMappings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userAddress** | **string** | User address | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUserMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListUserMappings200ResponseInner**](ListUserMappings200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

