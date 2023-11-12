# \DeletedMessageVaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ExportDeletedMessages**](DeletedMessageVaultAPI.md#ExportDeletedMessages) | **Post** /deletedMessages/users/{user}/actions/export | Export deleted messages for a specific user
[**PurgeMessage**](DeletedMessageVaultAPI.md#PurgeMessage) | **Delete** /deletedMessages/users/{user}/messages/{messageId} | Permanently remove a deleted message
[**PurgeMessages**](DeletedMessageVaultAPI.md#PurgeMessages) | **Delete** /deletedMessages | Purge all expired deleted messages
[**RestoreDeletedMessages**](DeletedMessageVaultAPI.md#RestoreDeletedMessages) | **Post** /deletedMessages/users/{user}/actions/restore | Restore deleted messages for a specific user



## ExportDeletedMessages

> ExportDeletedMessages(ctx, user).ExportTo(exportTo).ExportDeletedMessagesRequest(exportDeletedMessagesRequest).Execute()

Export deleted messages for a specific user

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
    user := "user_example" // string | The user to export deleted messages from
    exportTo := "exportTo_example" // string | The email address to export the messages to
    exportDeletedMessagesRequest := *openapiclient.NewExportDeletedMessagesRequest() // ExportDeletedMessagesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeletedMessageVaultAPI.ExportDeletedMessages(context.Background(), user).ExportTo(exportTo).ExportDeletedMessagesRequest(exportDeletedMessagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletedMessageVaultAPI.ExportDeletedMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | The user to export deleted messages from | 

### Other Parameters

Other parameters are passed through a pointer to a apiExportDeletedMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportTo** | **string** | The email address to export the messages to | 
 **exportDeletedMessagesRequest** | [**ExportDeletedMessagesRequest**](ExportDeletedMessagesRequest.md) |  | 

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


## PurgeMessage

> PurgeMessage(ctx, user, messageId).Execute()

Permanently remove a deleted message

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
    user := "user_example" // string | The user of the deleted message
    messageId := "messageId_example" // string | The ID of the deleted message

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeletedMessageVaultAPI.PurgeMessage(context.Background(), user, messageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletedMessageVaultAPI.PurgeMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | The user of the deleted message | 
**messageId** | **string** | The ID of the deleted message | 

### Other Parameters

Other parameters are passed through a pointer to a apiPurgeMessageRequest struct via the builder pattern


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


## PurgeMessages

> PurgeMessages(ctx).Scope(scope).Execute()

Purge all expired deleted messages

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
    scope := "scope_example" // string | The scope of messages to purge

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeletedMessageVaultAPI.PurgeMessages(context.Background()).Scope(scope).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletedMessageVaultAPI.PurgeMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPurgeMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **scope** | **string** | The scope of messages to purge | 

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


## RestoreDeletedMessages

> RestoreDeletedMessages(ctx, user).ExportDeletedMessagesRequest(exportDeletedMessagesRequest).Execute()

Restore deleted messages for a specific user

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
    user := "user_example" // string | The user to restore deleted messages for
    exportDeletedMessagesRequest := *openapiclient.NewExportDeletedMessagesRequest() // ExportDeletedMessagesRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DeletedMessageVaultAPI.RestoreDeletedMessages(context.Background(), user).ExportDeletedMessagesRequest(exportDeletedMessagesRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeletedMessageVaultAPI.RestoreDeletedMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**user** | **string** | The user to restore deleted messages for | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreDeletedMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **exportDeletedMessagesRequest** | [**ExportDeletedMessagesRequest**](ExportDeletedMessagesRequest.md) |  | 

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

