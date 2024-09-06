# \MailQueueAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteMailsOfMailQueue**](MailQueueAPI.md#DeleteMailsOfMailQueue) | **Delete** /mailQueues/{mailQueueName}/mails | Delete mails from a mail queue
[**FlushMailsOfMailQueue**](MailQueueAPI.md#FlushMailsOfMailQueue) | **Patch** /mailQueues/{mailQueueName} | Flush mails from a mail queue
[**ListMailQueues**](MailQueueAPI.md#ListMailQueues) | **Get** /mailQueues | List mail queues
[**ListMailsOfMailQueue**](MailQueueAPI.md#ListMailsOfMailQueue) | **Get** /mailQueues/{mailQueueName}/mails | List mails of a mail queue
[**RepublishMailQueue**](MailQueueAPI.md#RepublishMailQueue) | **Post** /mailQueues | RabbitMQ republishing a mail queue from Cassandra



## DeleteMailsOfMailQueue

> DeleteMailsOfMailQueue(ctx, mailQueueName).Sender(sender).Name(name).Recipient(recipient).Execute()

Delete mails from a mail queue

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
    mailQueueName := "mailQueueName_example" // string | Name of the mail queue
    sender := "sender_example" // string | Sender mail address (optional)
    name := "name_example" // string | Mail name (optional)
    recipient := "recipient_example" // string | Recipient mail address (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MailQueueAPI.DeleteMailsOfMailQueue(context.Background(), mailQueueName).Sender(sender).Name(name).Recipient(recipient).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailQueueAPI.DeleteMailsOfMailQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailQueueName** | **string** | Name of the mail queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailsOfMailQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sender** | **string** | Sender mail address | 
 **name** | **string** | Mail name | 
 **recipient** | **string** | Recipient mail address | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FlushMailsOfMailQueue

> FlushMailsOfMailQueue(ctx, mailQueueName).FlushMailsOfMailQueueRequest(flushMailsOfMailQueueRequest).Execute()

Flush mails from a mail queue

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
    mailQueueName := "mailQueueName_example" // string | Name of the mail queue
    flushMailsOfMailQueueRequest := *openapiclient.NewFlushMailsOfMailQueueRequest() // FlushMailsOfMailQueueRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MailQueueAPI.FlushMailsOfMailQueue(context.Background(), mailQueueName).FlushMailsOfMailQueueRequest(flushMailsOfMailQueueRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailQueueAPI.FlushMailsOfMailQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailQueueName** | **string** | Name of the mail queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiFlushMailsOfMailQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **flushMailsOfMailQueueRequest** | [**FlushMailsOfMailQueueRequest**](FlushMailsOfMailQueueRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMailQueues

> ListMailQueues(ctx).Execute()

List mail queues

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MailQueueAPI.ListMailQueues(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailQueueAPI.ListMailQueues``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMailQueuesRequest struct via the builder pattern


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMailsOfMailQueue

> []ListMailsOfMailQueue200ResponseInner ListMailsOfMailQueue(ctx, mailQueueName).Limit(limit).Execute()

List mails of a mail queue

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
    mailQueueName := "mailQueueName_example" // string | Name of the mail queue
    limit := int32(56) // int32 | Maximum number of mails returned in a single call (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailQueueAPI.ListMailsOfMailQueue(context.Background(), mailQueueName).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailQueueAPI.ListMailsOfMailQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMailsOfMailQueue`: []ListMailsOfMailQueue200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MailQueueAPI.ListMailsOfMailQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mailQueueName** | **string** | Name of the mail queue | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMailsOfMailQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Maximum number of mails returned in a single call | 

### Return type

[**[]ListMailsOfMailQueue200ResponseInner**](ListMailsOfMailQueue200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RepublishMailQueue

> RepublishMailQueue(ctx).Action(action).OlderThan(olderThan).Execute()

RabbitMQ republishing a mail queue from Cassandra

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
    action := "action_example" // string | Republish action (RepublishNotProcessedMails)
    olderThan := "olderThan_example" // string | Older than duration (e.g., 5h, 7d, 1y)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MailQueueAPI.RepublishMailQueue(context.Background()).Action(action).OlderThan(olderThan).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailQueueAPI.RepublishMailQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRepublishMailQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **action** | **string** | Republish action (RepublishNotProcessedMails) | 
 **olderThan** | **string** | Older than duration (e.g., 5h, 7d, 1y) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

