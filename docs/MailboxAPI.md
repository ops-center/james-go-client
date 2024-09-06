# \MailboxAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PerformActionsOnMailboxes**](MailboxAPI.md#PerformActionsOnMailboxes) | **Post** /mailboxes | Perform actions on mailboxes
[**ReindexMailbox**](MailboxAPI.md#ReindexMailbox) | **Post** /mailboxes/{mailboxId} | Reindex a mailbox



## PerformActionsOnMailboxes

> PerformActionsOnMailboxes201Response PerformActionsOnMailboxes(ctx).Action(action).Execute()

Perform actions on mailboxes

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
    action := "action_example" // string | Comma-separated list of actions to perform on mailboxes

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxAPI.PerformActionsOnMailboxes(context.Background()).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxAPI.PerformActionsOnMailboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PerformActionsOnMailboxes`: PerformActionsOnMailboxes201Response
    fmt.Fprintf(os.Stdout, "Response from `MailboxAPI.PerformActionsOnMailboxes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPerformActionsOnMailboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Comma-separated list of actions to perform on mailboxes | 

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


## ReindexMailbox

> PerformActionsOnMailboxes201Response ReindexMailbox(ctx, mailboxId).Task(task).MessagesPerSecond(messagesPerSecond).Mode(mode).Execute()

Reindex a mailbox

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
    mailboxId := "mailboxId_example" // string | ID of the mailbox to reindex
    task := "task_example" // string | The reindexing task to perform
    messagesPerSecond := int32(56) // int32 | The rate at which messages should be processed per second (optional)
    mode := "mode_example" // string | The reindexing mode used (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailboxAPI.ReindexMailbox(context.Background(), mailboxId).Task(task).MessagesPerSecond(messagesPerSecond).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailboxAPI.ReindexMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ReindexMailbox`: PerformActionsOnMailboxes201Response
    fmt.Fprintf(os.Stdout, "Response from `MailboxAPI.ReindexMailbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailboxId** | **string** | ID of the mailbox to reindex | 

### Other Parameters

Other parameters are passed through a pointer to a apiReindexMailboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **task** | **string** | The reindexing task to perform | 
 **messagesPerSecond** | **int32** | The rate at which messages should be processed per second | 
 **mode** | **string** | The reindexing mode used | 

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

