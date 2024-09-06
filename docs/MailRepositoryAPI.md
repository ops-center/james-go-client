# \MailRepositoryAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMailRepository**](MailRepositoryAPI.md#CreateMailRepository) | **Put** /mailRepositories/{encodedPathOfTheRepository} | Create a mail repository
[**GetMailRepository**](MailRepositoryAPI.md#GetMailRepository) | **Get** /mailRepositories/{encodedPathOfTheRepository} | Getting additional information for a mail repository
[**ListMailRepositories**](MailRepositoryAPI.md#ListMailRepositories) | **Get** /mailRepositories | Listing mail repositories
[**ListMailsInMailRepository**](MailRepositoryAPI.md#ListMailsInMailRepository) | **Get** /mailRepositories/{encodedPathOfTheRepository}/mails | Listing mails contained in a mail repository



## CreateMailRepository

> CreateMailRepository(ctx, encodedPathOfTheRepository).Protocol(protocol).Execute()

Create a mail repository

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
    encodedPathOfTheRepository := "mailRepo" // string | Encoded resource path of the created mail repository
    protocol := "file" // string | Protocol of the mail repository

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.MailRepositoryAPI.CreateMailRepository(context.Background(), encodedPathOfTheRepository).Protocol(protocol).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailRepositoryAPI.CreateMailRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedPathOfTheRepository** | **string** | Encoded resource path of the created mail repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMailRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **protocol** | **string** | Protocol of the mail repository | 

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


## GetMailRepository

> GetMailRepository200Response GetMailRepository(ctx, encodedPathOfTheRepository).Execute()

Getting additional information for a mail repository

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
    encodedPathOfTheRepository := "var%2Fmail%2Ferror%2F" // string | Encoded resource path of an existing mail repository

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailRepositoryAPI.GetMailRepository(context.Background(), encodedPathOfTheRepository).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailRepositoryAPI.GetMailRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMailRepository`: GetMailRepository200Response
    fmt.Fprintf(os.Stdout, "Response from `MailRepositoryAPI.GetMailRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedPathOfTheRepository** | **string** | Encoded resource path of an existing mail repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMailRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMailRepository200Response**](GetMailRepository200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMailRepositories

> []ListMailRepositories200ResponseInner ListMailRepositories(ctx).Execute()

Listing mail repositories

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
    resp, r, err := apiClient.MailRepositoryAPI.ListMailRepositories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailRepositoryAPI.ListMailRepositories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMailRepositories`: []ListMailRepositories200ResponseInner
    fmt.Fprintf(os.Stdout, "Response from `MailRepositoryAPI.ListMailRepositories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMailRepositoriesRequest struct via the builder pattern


### Return type

[**[]ListMailRepositories200ResponseInner**](ListMailRepositories200ResponseInner.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMailsInMailRepository

> []string ListMailsInMailRepository(ctx, encodedPathOfTheRepository).Limit(limit).Offset(offset).Execute()

Listing mails contained in a mail repository

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
    encodedPathOfTheRepository := "var%2Fmail%2Ferror%2F" // string | Encoded resource path of an existing mail repository
    limit := int32(56) // int32 | Limit the number of returned elements (optional)
    offset := int32(56) // int32 | Skip the specified number of elements (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MailRepositoryAPI.ListMailsInMailRepository(context.Background(), encodedPathOfTheRepository).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MailRepositoryAPI.ListMailsInMailRepository``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMailsInMailRepository`: []string
    fmt.Fprintf(os.Stdout, "Response from `MailRepositoryAPI.ListMailsInMailRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**encodedPathOfTheRepository** | **string** | Encoded resource path of an existing mail repository | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMailsInMailRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** | Limit the number of returned elements | 
 **offset** | **int32** | Skip the specified number of elements | 

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

