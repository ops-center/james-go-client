# \GarbageCollectionAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RunBlobGarbageCollector**](GarbageCollectionAPI.md#RunBlobGarbageCollector) | **Delete** /blobs | Run blob garbage collection



## RunBlobGarbageCollector

> RunBlobGarbageCollector200Response RunBlobGarbageCollector(ctx).Scope(scope).AssociatedProbability(associatedProbability).ExpectedBlobCount(expectedBlobCount).Execute()

Run blob garbage collection

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
    scope := "scope_example" // string | Specifies the scope of garbage collection (e.g., \"unreferenced\" blobs)
    associatedProbability := float64(1.2) // float64 |  (optional) (default to 0.01)
    expectedBlobCount := int32(56) // int32 |  (optional) (default to 1000000)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GarbageCollectionAPI.RunBlobGarbageCollector(context.Background()).Scope(scope).AssociatedProbability(associatedProbability).ExpectedBlobCount(expectedBlobCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GarbageCollectionAPI.RunBlobGarbageCollector``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RunBlobGarbageCollector`: RunBlobGarbageCollector200Response
    fmt.Fprintf(os.Stdout, "Response from `GarbageCollectionAPI.RunBlobGarbageCollector`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRunBlobGarbageCollectorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Specifies the scope of garbage collection (e.g., \&quot;unreferenced\&quot; blobs) | 
 **associatedProbability** | **float64** |  | [default to 0.01]
 **expectedBlobCount** | **int32** |  | [default to 1000000]

### Return type

[**RunBlobGarbageCollector200Response**](RunBlobGarbageCollector200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

