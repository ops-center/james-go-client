# \RegexMappingAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddRegexMapping**](RegexMappingAPI.md#AddRegexMapping) | **Post** /mappings/regex/{mappingSource}/targets/{regex} | Add a regex mapping
[**RemoveRegexMapping**](RegexMappingAPI.md#RemoveRegexMapping) | **Delete** /mappings/regex/{mappingSource}/targets/{regex} | Remove a regex mapping



## AddRegexMapping

> AddRegexMapping(ctx, mappingSource, regex).Execute()

Add a regex mapping

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
    mappingSource := "james@domain.tld" // string | Mapping source (recipient address)
    regex := "james@.*:james-intern@james.org" // string | Java Regular Expression (regex)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegexMappingAPI.AddRegexMapping(context.Background(), mappingSource, regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegexMappingAPI.AddRegexMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingSource** | **string** | Mapping source (recipient address) | 
**regex** | **string** | Java Regular Expression (regex) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddRegexMappingRequest struct via the builder pattern


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


## RemoveRegexMapping

> RemoveRegexMapping(ctx, mappingSource, regex).Execute()

Remove a regex mapping

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
    mappingSource := "james@domain.tld" // string | Mapping source (recipient address)
    regex := "[O_O]:james-intern@james.org" // string | Java Regular Expression (regex)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.RegexMappingAPI.RemoveRegexMapping(context.Background(), mappingSource, regex).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RegexMappingAPI.RemoveRegexMapping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mappingSource** | **string** | Mapping source (recipient address) | 
**regex** | **string** | Java Regular Expression (regex) | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveRegexMappingRequest struct via the builder pattern


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

