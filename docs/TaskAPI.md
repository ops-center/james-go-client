# \TaskAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AwaitTaskCompletion**](TaskAPI.md#AwaitTaskCompletion) | **Get** /tasks/{taskId}/await | Await the completion of a task
[**CancelTask**](TaskAPI.md#CancelTask) | **Delete** /tasks/{taskId} | Cancel a task
[**GetTask**](TaskAPI.md#GetTask) | **Get** /tasks/{taskId} | Get a task&#39;s details
[**ListTasks**](TaskAPI.md#ListTasks) | **Get** /tasks | List tasks



## AwaitTaskCompletion

> ExecutionReport AwaitTaskCompletion(ctx, taskId).Timeout(timeout).Execute()

Await the completion of a task

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
    taskId := "taskId_example" // string | The ID of the task
    timeout := "timeout_example" // string | The timeout duration in the format Nunit (e.g., 30s, 5m, 7d, 1y) (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.AwaitTaskCompletion(context.Background(), taskId).Timeout(timeout).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.AwaitTaskCompletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AwaitTaskCompletion`: ExecutionReport
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.AwaitTaskCompletion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiAwaitTaskCompletionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **timeout** | **string** | The timeout duration in the format Nunit (e.g., 30s, 5m, 7d, 1y) | 

### Return type

[**ExecutionReport**](ExecutionReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelTask

> CancelTask(ctx, taskId).Execute()

Cancel a task

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
    taskId := "taskId_example" // string | The ID of the task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.TaskAPI.CancelTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.CancelTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelTaskRequest struct via the builder pattern


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


## GetTask

> ExecutionReport GetTask(ctx, taskId).Execute()

Get a task's details

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
    taskId := "taskId_example" // string | The ID of the task

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.GetTask(context.Background(), taskId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.GetTask``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTask`: ExecutionReport
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.GetTask`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the task | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaskRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ExecutionReport**](ExecutionReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTasks

> []ExecutionReport ListTasks(ctx).Status(status).Type_(type_).SubmittedBefore(submittedBefore).SubmittedAfter(submittedAfter).StartedBefore(startedBefore).StartedAfter(startedAfter).CompletedBefore(completedBefore).CompletedAfter(completedAfter).FailedBefore(failedBefore).FailedAfter(failedAfter).Offset(offset).Limit(limit).Execute()

List tasks

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "github.com/searchlight/james-go-client"
)

func main() {
    status := "status_example" // string | The status of the tasks (waiting, inProgress, canceledRequested, completed, canceled, failed) (optional)
    type_ := "type__example" // string | The type of the tasks (optional)
    submittedBefore := time.Now() // time.Time | Return tasks submitted before this date (optional)
    submittedAfter := time.Now() // time.Time | Return tasks submitted after this date (optional)
    startedBefore := time.Now() // time.Time | Return tasks started before this date (optional)
    startedAfter := time.Now() // time.Time | Return tasks started after this date (optional)
    completedBefore := time.Now() // time.Time | Return tasks completed before this date (optional)
    completedAfter := time.Now() // time.Time | Return tasks completed after this date (optional)
    failedBefore := time.Now() // time.Time | Return tasks failed before this date (optional)
    failedAfter := time.Now() // time.Time | Return tasks failed after this date (optional)
    offset := int32(56) // int32 | Number of tasks to skip in the response (optional)
    limit := int32(56) // int32 | Maximum number of tasks to return in one call (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaskAPI.ListTasks(context.Background()).Status(status).Type_(type_).SubmittedBefore(submittedBefore).SubmittedAfter(submittedAfter).StartedBefore(startedBefore).StartedAfter(startedAfter).CompletedBefore(completedBefore).CompletedAfter(completedAfter).FailedBefore(failedBefore).FailedAfter(failedAfter).Offset(offset).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaskAPI.ListTasks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTasks`: []ExecutionReport
    fmt.Fprintf(os.Stdout, "Response from `TaskAPI.ListTasks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTasksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | The status of the tasks (waiting, inProgress, canceledRequested, completed, canceled, failed) | 
 **type_** | **string** | The type of the tasks | 
 **submittedBefore** | **time.Time** | Return tasks submitted before this date | 
 **submittedAfter** | **time.Time** | Return tasks submitted after this date | 
 **startedBefore** | **time.Time** | Return tasks started before this date | 
 **startedAfter** | **time.Time** | Return tasks started after this date | 
 **completedBefore** | **time.Time** | Return tasks completed before this date | 
 **completedAfter** | **time.Time** | Return tasks completed after this date | 
 **failedBefore** | **time.Time** | Return tasks failed before this date | 
 **failedAfter** | **time.Time** | Return tasks failed after this date | 
 **offset** | **int32** | Number of tasks to skip in the response | 
 **limit** | **int32** | Maximum number of tasks to return in one call | 

### Return type

[**[]ExecutionReport**](ExecutionReport.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

