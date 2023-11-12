# \GhostMailboxAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CorrectGhostMailbox**](GhostMailboxAPI.md#CorrectGhostMailbox) | **Post** /cassandra/mailbox/merging | Correct ghost mailbox by merging



## CorrectGhostMailbox

> CorrectGhostMailbox201Response CorrectGhostMailbox(ctx).CorrectGhostMailboxRequest(correctGhostMailboxRequest).Execute()

Correct ghost mailbox by merging

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
    correctGhostMailboxRequest := *openapiclient.NewCorrectGhostMailboxRequest() // CorrectGhostMailboxRequest | Parameters for merging mailboxes

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GhostMailboxAPI.CorrectGhostMailbox(context.Background()).CorrectGhostMailboxRequest(correctGhostMailboxRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GhostMailboxAPI.CorrectGhostMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CorrectGhostMailbox`: CorrectGhostMailbox201Response
    fmt.Fprintf(os.Stdout, "Response from `GhostMailboxAPI.CorrectGhostMailbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCorrectGhostMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **correctGhostMailboxRequest** | [**CorrectGhostMailboxRequest**](CorrectGhostMailboxRequest.md) | Parameters for merging mailboxes | 

### Return type

[**CorrectGhostMailbox201Response**](CorrectGhostMailbox201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

