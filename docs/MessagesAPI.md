# \MessagesAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ReindexEmail**](MessagesAPI.md#ReindexEmail) | **Post** /messages/{messageId} | Reindex a single mail by messageId
[**ScheduleTask**](MessagesAPI.md#ScheduleTask) | **Post** /messages | Schedule a task for fixing message inconsistencies



## ReindexEmail

> PerformActionsOnMailboxes201Response ReindexEmail(ctx, messageId).Task(task).Execute()

Reindex a single mail by messageId

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
    messageId := "messageId_example" // string | Implementation-dependent valid messageId
    task := "task_example" // string | Task type for reindexing

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.ReindexEmail(context.Background(), messageId).Task(task).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.ReindexEmail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReindexEmail`: PerformActionsOnMailboxes201Response
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.ReindexEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**messageId** | **string** | Implementation-dependent valid messageId | 

### Other Parameters

Other parameters are passed through a pointer to a apiReindexEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | **string** | Task type for reindexing | 

### Return type

[**PerformActionsOnMailboxes201Response**](PerformActionsOnMailboxes201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScheduleTask

> PerformActionsOnMailboxes201Response ScheduleTask(ctx).Task(task).MessagesPerSecond(messagesPerSecond).Mode(mode).Execute()

Schedule a task for fixing message inconsistencies

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
    task := "task_example" // string | Task type for fixing message inconsistencies
    messagesPerSecond := int32(56) // int32 | Rate of messages to be processed per second (optional)
    mode := "mode_example" // string | Reindexing mode used (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MessagesAPI.ScheduleTask(context.Background()).Task(task).MessagesPerSecond(messagesPerSecond).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.ScheduleTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ScheduleTask`: PerformActionsOnMailboxes201Response
    fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.ScheduleTask`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiScheduleTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **task** | **string** | Task type for fixing message inconsistencies | 
 **messagesPerSecond** | **int32** | Rate of messages to be processed per second | 
 **mode** | **string** | Reindexing mode used | 

### Return type

[**PerformActionsOnMailboxes201Response**](PerformActionsOnMailboxes201Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

