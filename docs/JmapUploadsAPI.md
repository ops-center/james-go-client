# \JmapUploadsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CleanUploadRepository**](JmapUploadsAPI.md#CleanUploadRepository) | **Delete** /jmap/uploads | Clean upload repository



## CleanUploadRepository

> CleanUploadRepository201Response CleanUploadRepository(ctx).Scope(scope).CleanUploadRepositoryRequest(cleanUploadRepositoryRequest).Execute()

Clean upload repository

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
    scope := "scope_example" // string | Specifies the scope of cleanup (e.g., \"expired\" uploads)
    cleanUploadRepositoryRequest := *openapiclient.NewCleanUploadRepositoryRequest() // CleanUploadRepositoryRequest | Task details (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.JmapUploadsAPI.CleanUploadRepository(context.Background()).Scope(scope).CleanUploadRepositoryRequest(cleanUploadRepositoryRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `JmapUploadsAPI.CleanUploadRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CleanUploadRepository`: CleanUploadRepository201Response
    fmt.Fprintf(os.Stdout, "Response from `JmapUploadsAPI.CleanUploadRepository`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCleanUploadRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | Specifies the scope of cleanup (e.g., \&quot;expired\&quot; uploads) | 
 **cleanUploadRepositoryRequest** | [**CleanUploadRepositoryRequest**](CleanUploadRepositoryRequest.md) | Task details | 

### Return type

[**CleanUploadRepository201Response**](CleanUploadRepository201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

