# \CassandraExtraAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PerformActionOnCassandraMappings**](CassandraExtraAPI.md#PerformActionOnCassandraMappings) | **Post** /cassandra/mappings | Perform an action on mappings_sources table



## PerformActionOnCassandraMappings

> PerformActionOnCassandraMappings201Response PerformActionOnCassandraMappings(ctx).Action(action).Execute()

Perform an action on mappings_sources table

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
    action := "action_example" // string | The action to perform on mappings_sources table

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CassandraExtraAPI.PerformActionOnCassandraMappings(context.Background()).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CassandraExtraAPI.PerformActionOnCassandraMappings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformActionOnCassandraMappings`: PerformActionOnCassandraMappings201Response
    fmt.Fprintf(os.Stdout, "Response from `CassandraExtraAPI.PerformActionOnCassandraMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionOnCassandraMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | The action to perform on mappings_sources table | 

### Return type

[**PerformActionOnCassandraMappings201Response**](PerformActionOnCassandraMappings201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

