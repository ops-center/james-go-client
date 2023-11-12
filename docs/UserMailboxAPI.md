# \UserMailboxAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearMailbox**](UserMailboxAPI.md#ClearMailbox) | **Delete** /users/{username}/mailboxes/{mailboxName}/messages | Clear mailbox content
[**CountEmails**](UserMailboxAPI.md#CountEmails) | **Get** /users/{username}/mailboxes/{mailboxName}/messageCount | Count emails in a mailbox
[**CountUnseenEmails**](UserMailboxAPI.md#CountUnseenEmails) | **Get** /users/{username}/mailboxes/{mailboxName}/unseenMessageCount | Count unseen emails in a mailbox
[**CreateMailbox**](UserMailboxAPI.md#CreateMailbox) | **Put** /users/{username}/mailboxes/{mailboxNameToBeCreated} | Create a mailbox
[**DeleteMailbox**](UserMailboxAPI.md#DeleteMailbox) | **Delete** /users/{username}/mailboxes/INBOX.work | Delete a mailbox and its children
[**DeleteMailboxes**](UserMailboxAPI.md#DeleteMailboxes) | **Delete** /users/{username}/mailboxes | Delete user mailboxes
[**ExistsMailbox**](UserMailboxAPI.md#ExistsMailbox) | **Get** /users/{username}/mailboxes/{mailboxNameToBeTested} | Test existence of a mailbox
[**ExportMailboxes**](UserMailboxAPI.md#ExportMailboxes) | **Post** /users/{username}/mailboxes?action&#x3D;export | Export user mailboxes
[**ListMailboxes**](UserMailboxAPI.md#ListMailboxes) | **Get** /users/{username}/mailboxes | List user mailboxes
[**RecomputeCassandraFilteringProjection**](UserMailboxAPI.md#RecomputeCassandraFilteringProjection) | **Post** /mailboxes?task&#x3D;populateFilteringProjection | Recompute Cassandra filtering projection
[**RecomputeMessageViewProjection**](UserMailboxAPI.md#RecomputeMessageViewProjection) | **Post** /users/{username}/mailboxes?task&#x3D;recomputeFastViewProjectionItems | Recompute User JMAP fast message view projection
[**ReindexEmails**](UserMailboxAPI.md#ReindexEmails) | **Post** /users/{username}/mailboxes?task&#x3D;reIndex | Reindex a user&#39;s mails
[**SubscribeAllMailboxes**](UserMailboxAPI.md#SubscribeAllMailboxes) | **Post** /users/{username}/mailboxes?task&#x3D;subscribeAll | Subscribe a user to all of their mailboxes



## ClearMailbox

> ClearMailbox(ctx, username, mailboxName).Execute()

Clear mailbox content

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
    username := "username_example" // string | 
    mailboxName := "mailboxName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.ClearMailbox(context.Background(), username, mailboxName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.ClearMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**mailboxName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiClearMailboxRequest struct via the builder pattern


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


## CountEmails

> CountEmails(ctx, username, mailboxName).Execute()

Count emails in a mailbox

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
    username := "username_example" // string | 
    mailboxName := "mailboxName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.CountEmails(context.Background(), username, mailboxName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.CountEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**mailboxName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountEmailsRequest struct via the builder pattern


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


## CountUnseenEmails

> CountUnseenEmails(ctx, username, mailboxName).Execute()

Count unseen emails in a mailbox

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
    username := "username_example" // string | 
    mailboxName := "mailboxName_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.CountUnseenEmails(context.Background(), username, mailboxName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.CountUnseenEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**mailboxName** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountUnseenEmailsRequest struct via the builder pattern


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


## CreateMailbox

> CreateMailbox(ctx, username, mailboxNameToBeCreated).Execute()

Create a mailbox

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
    username := "username_example" // string | 
    mailboxNameToBeCreated := "mailboxNameToBeCreated_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.CreateMailbox(context.Background(), username, mailboxNameToBeCreated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.CreateMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**mailboxNameToBeCreated** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMailboxRequest struct via the builder pattern


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


## DeleteMailbox

> DeleteMailbox(ctx, username, mailboxNameToBeDeleted).Execute()

Delete a mailbox and its children

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
    username := "username_example" // string | 
    mailboxNameToBeDeleted := "mailboxNameToBeDeleted_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.DeleteMailbox(context.Background(), username, mailboxNameToBeDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.DeleteMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**mailboxNameToBeDeleted** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailboxRequest struct via the builder pattern


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


## DeleteMailboxes

> DeleteMailboxes(ctx, username).Execute()

Delete user mailboxes

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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.DeleteMailboxes(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.DeleteMailboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMailboxesRequest struct via the builder pattern


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


## ExistsMailbox

> ExistsMailbox(ctx, username, mailboxNameToBeTested).Execute()

Test existence of a mailbox

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
    username := "username_example" // string | 
    mailboxNameToBeTested := "mailboxNameToBeTested_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.ExistsMailbox(context.Background(), username, mailboxNameToBeTested).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.ExistsMailbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 
**mailboxNameToBeTested** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExistsMailboxRequest struct via the builder pattern


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


## ExportMailboxes

> ExportMailboxes(ctx, username).Execute()

Export user mailboxes

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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.ExportMailboxes(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.ExportMailboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportMailboxesRequest struct via the builder pattern


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


## ListMailboxes

> []ListMailboxes200ResponseInner ListMailboxes(ctx, username).Execute()

List user mailboxes

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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserMailboxAPI.ListMailboxes(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.ListMailboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMailboxes`: []ListMailboxes200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `UserMailboxAPI.ListMailboxes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMailboxesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ListMailboxes200ResponseInner**](ListMailboxes200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RecomputeCassandraFilteringProjection

> RecomputeCassandraFilteringProjection(ctx).Execute()

Recompute Cassandra filtering projection

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.RecomputeCassandraFilteringProjection(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.RecomputeCassandraFilteringProjection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRecomputeCassandraFilteringProjectionRequest struct via the builder pattern


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


## RecomputeMessageViewProjection

> RecomputeMessageViewProjection(ctx, username).MessagesPerSecond(messagesPerSecond).Execute()

Recompute User JMAP fast message view projection

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
    username := "username_example" // string | 
    messagesPerSecond := int32(56) // int32 |  (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.RecomputeMessageViewProjection(context.Background(), username).MessagesPerSecond(messagesPerSecond).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.RecomputeMessageViewProjection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRecomputeMessageViewProjectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messagesPerSecond** | **int32** |  | [default to 10]

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


## ReindexEmails

> ReindexEmails(ctx, username).MessagesPerSecond(messagesPerSecond).Mode(mode).Execute()

Reindex a user's mails

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
    username := "username_example" // string | 
    messagesPerSecond := int32(56) // int32 |  (optional) (default to 50)
    mode := "mode_example" // string |  (optional) (default to "rebuildAll")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.ReindexEmails(context.Background(), username).MessagesPerSecond(messagesPerSecond).Mode(mode).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.ReindexEmails``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReindexEmailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **messagesPerSecond** | **int32** |  | [default to 50]
 **mode** | **string** |  | [default to &quot;rebuildAll&quot;]

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


## SubscribeAllMailboxes

> SubscribeAllMailboxes(ctx, username).Execute()

Subscribe a user to all of their mailboxes

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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.UserMailboxAPI.SubscribeAllMailboxes(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserMailboxAPI.SubscribeAllMailboxes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubscribeAllMailboxesRequest struct via the builder pattern


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

