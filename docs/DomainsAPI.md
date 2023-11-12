# \DomainsAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDomain**](DomainsAPI.md#CreateDomain) | **Put** /domains/{domainToBeCreated} | Create a domain
[**CreateDomainAlias**](DomainsAPI.md#CreateDomainAlias) | **Put** /domains/{domainName}/aliases | Create an alias for a domain
[**DeleteDomain**](DomainsAPI.md#DeleteDomain) | **Delete** /domains/{domainToBeDeleted} | Delete a domain
[**DeleteDomainAlias**](DomainsAPI.md#DeleteDomainAlias) | **Delete** /domains/{domainName}/aliases | Delete an alias for a domain
[**DeleteUserDataOfDomain**](DomainsAPI.md#DeleteUserDataOfDomain) | **Post** /domains/{domainToBeUsed} | Delete all users data of a domain
[**ExistsDomain**](DomainsAPI.md#ExistsDomain) | **Get** /domains/{domainName} | Test if a domain exists
[**ListDomainAliases**](DomainsAPI.md#ListDomainAliases) | **Get** /domains/{domainName}/aliases | Get the list of aliases for a domain
[**ListDomains**](DomainsAPI.md#ListDomains) | **Get** /domains | Get the list of domains



## CreateDomain

> CreateDomain(ctx, domainToBeCreated).Execute()

Create a domain

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
    domainToBeCreated := "domainToBeCreated_example" // string | Name of the domain to be created

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.CreateDomain(context.Background(), domainToBeCreated).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.CreateDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeCreated** | **string** | Name of the domain to be created | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainRequest struct via the builder pattern


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


## CreateDomainAlias

> CreateDomainAlias(ctx, domainName).SourceDomainName(sourceDomainName).Execute()

Create an alias for a domain

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
    domainName := "domainName_example" // string | Name of the destination domain for the alias
    sourceDomainName := "sourceDomainName_example" // string | Name of the source domain for the alias

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.CreateDomainAlias(context.Background(), domainName).SourceDomainName(sourceDomainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.CreateDomainAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Name of the destination domain for the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDomainAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceDomainName** | **string** | Name of the source domain for the alias | 

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


## DeleteDomain

> DeleteDomain(ctx, domainToBeDeleted).Execute()

Delete a domain

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
    domainToBeDeleted := "domainToBeDeleted_example" // string | Name of the domain to be deleted

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.DeleteDomain(context.Background(), domainToBeDeleted).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeDeleted** | **string** | Name of the domain to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainRequest struct via the builder pattern


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


## DeleteDomainAlias

> DeleteDomainAlias(ctx, domainName).SourceDomainName(sourceDomainName).Execute()

Delete an alias for a domain

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
    domainName := "domainName_example" // string | Name of the destination domain for the alias
    sourceDomainName := "sourceDomainName_example" // string | Name of the source domain for the alias

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.DeleteDomainAlias(context.Background(), domainName).SourceDomainName(sourceDomainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteDomainAlias``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Name of the destination domain for the alias | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDomainAliasRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sourceDomainName** | **string** | Name of the source domain for the alias | 

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


## DeleteUserDataOfDomain

> DeleteUserDataOfDomain(ctx, domainToBeUsed).Action(action).Execute()

Delete all users data of a domain

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
    domainToBeUsed := "domainToBeUsed_example" // string | Name of the domain to delete user data from
    action := "action_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.DeleteUserDataOfDomain(context.Background(), domainToBeUsed).Action(action).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.DeleteUserDataOfDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainToBeUsed** | **string** | Name of the domain to delete user data from | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserDataOfDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **action** | **string** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExistsDomain

> ExistsDomain(ctx, domainName).Execute()

Test if a domain exists

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
    domainName := "domainName_example" // string | Name of the domain to be tested

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    r, err := apiClient.DomainsAPI.ExistsDomain(context.Background(), domainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.ExistsDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Name of the domain to be tested | 

### Other Parameters

Other parameters are passed through a pointer to a apiExistsDomainRequest struct via the builder pattern


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


## ListDomainAliases

> []DomainAlias ListDomainAliases(ctx, domainName).Execute()

Get the list of aliases for a domain

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
    domainName := "domainName_example" // string | Name of the domain to retrieve aliases for

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DomainsAPI.ListDomainAliases(context.Background(), domainName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.ListDomainAliases``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomainAliases`: []DomainAlias
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.ListDomainAliases`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**domainName** | **string** | Name of the domain to retrieve aliases for | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainAliasesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]DomainAlias**](DomainAlias.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDomains

> []string ListDomains(ctx).Execute()

Get the list of domains

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
    resp, r, err := apiClient.DomainsAPI.ListDomains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DomainsAPI.ListDomains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDomains`: []string
    fmt.Fprintf(os.Stdout, "Response from `DomainsAPI.ListDomains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDomainsRequest struct via the builder pattern


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

