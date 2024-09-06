# \EventDeadLetterAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAllEvents**](EventDeadLetterAPI.md#DeleteAllEvents) | **Delete** /events/deadLetter/groups/{groupName} | Delete all events of a group
[**DeleteEvent**](EventDeadLetterAPI.md#DeleteEvent) | **Delete** /events/deadLetter/groups/{groupName}/{insertionId} | Delete an event
[**GetEvent**](EventDeadLetterAPI.md#GetEvent) | **Get** /events/deadLetter/groups/{groupName}/{insertionId} | Get event details
[**ListFailedEvents**](EventDeadLetterAPI.md#ListFailedEvents) | **Get** /events/deadLetter/groups/{groupName} | List failed events for a given group
[**ListMailboxListenerGroups**](EventDeadLetterAPI.md#ListMailboxListenerGroups) | **Get** /events/deadLetter/groups | List Mailbox Listener Groups
[**RedeliverAllEvents**](EventDeadLetterAPI.md#RedeliverAllEvents) | **Post** /events/deadLetter/groups | Redeliver all events
[**RedeliverEvent**](EventDeadLetterAPI.md#RedeliverEvent) | **Post** /events/deadLetter/groups/{groupName}/{insertionId}/reDeliver | Redeliver a single event
[**RedeliverGroupEvents**](EventDeadLetterAPI.md#RedeliverGroupEvents) | **Post** /events/deadLetter/groups/{groupName}/reDeliver | Redeliver group events



## DeleteAllEvents

> DeleteAllEvents(ctx, groupName).Execute()

Delete all events of a group

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
    groupName := "groupName_example" // string | The name of the mailbox listener group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventDeadLetterAPI.DeleteAllEvents(context.Background(), groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventDeadLetterAPI.DeleteAllEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The name of the mailbox listener group | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAllEventsRequest struct via the builder pattern


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


## DeleteEvent

> DeleteEvent(ctx, groupName, insertionId).Execute()

Delete an event

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
    groupName := "groupName_example" // string | The name of the mailbox listener group
    insertionId := "insertionId_example" // string | The insertion ID of the event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventDeadLetterAPI.DeleteEvent(context.Background(), groupName, insertionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventDeadLetterAPI.DeleteEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The name of the mailbox listener group | 
**insertionId** | **string** | The insertion ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventRequest struct via the builder pattern


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


## GetEvent

> GetEvent(ctx, groupName, insertionId).Execute()

Get event details

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
    groupName := "groupName_example" // string | The name of the mailbox listener group
    insertionId := "insertionId_example" // string | The insertion ID of the event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventDeadLetterAPI.GetEvent(context.Background(), groupName, insertionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventDeadLetterAPI.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The name of the mailbox listener group | 
**insertionId** | **string** | The insertion ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


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


## ListFailedEvents

> []string ListFailedEvents(ctx, groupName).Execute()

List failed events for a given group

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
    groupName := "groupName_example" // string | The name of the mailbox listener group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventDeadLetterAPI.ListFailedEvents(context.Background(), groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventDeadLetterAPI.ListFailedEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFailedEvents`: []string
    fmt.Fprintf(os.Stdout, "Response from `EventDeadLetterAPI.ListFailedEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The name of the mailbox listener group | 

### Other Parameters

Other parameters are passed through a pointer to a apiListFailedEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMailboxListenerGroups

> []string ListMailboxListenerGroups(ctx).Execute()

List Mailbox Listener Groups

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
    resp, r, err := apiClient.EventDeadLetterAPI.ListMailboxListenerGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventDeadLetterAPI.ListMailboxListenerGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMailboxListenerGroups`: []string
    fmt.Fprintf(os.Stdout, "Response from `EventDeadLetterAPI.ListMailboxListenerGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMailboxListenerGroupsRequest struct via the builder pattern


### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RedeliverAllEvents

> RedeliverAllEvents(ctx).Execute()

Redeliver all events

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
    r, err := apiClient.EventDeadLetterAPI.RedeliverAllEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventDeadLetterAPI.RedeliverAllEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRedeliverAllEventsRequest struct via the builder pattern


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


## RedeliverEvent

> RedeliverEvent(ctx, groupName, insertionId).Execute()

Redeliver a single event

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
    groupName := "groupName_example" // string | The name of the mailbox listener group
    insertionId := "insertionId_example" // string | The insertion ID of the event

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventDeadLetterAPI.RedeliverEvent(context.Background(), groupName, insertionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventDeadLetterAPI.RedeliverEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The name of the mailbox listener group | 
**insertionId** | **string** | The insertion ID of the event | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeliverEventRequest struct via the builder pattern


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


## RedeliverGroupEvents

> RedeliverGroupEvents(ctx, groupName).Execute()

Redeliver group events

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
    groupName := "groupName_example" // string | The name of the mailbox listener group

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.EventDeadLetterAPI.RedeliverGroupEvents(context.Background(), groupName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventDeadLetterAPI.RedeliverGroupEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupName** | **string** | The name of the mailbox listener group | 

### Other Parameters

Other parameters are passed through a pointer to a apiRedeliverGroupEventsRequest struct via the builder pattern


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

