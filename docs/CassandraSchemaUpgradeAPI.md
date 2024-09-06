# \CassandraSchemaUpgradeAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetLatestAvailableSchemaVersion**](CassandraSchemaUpgradeAPI.md#GetLatestAvailableSchemaVersion) | **Get** /cassandra/version/latest | Retrieve latest available Cassandra schema version
[**GetSchemaVersion**](CassandraSchemaUpgradeAPI.md#GetSchemaVersion) | **Get** /cassandra/version | Retrieve current Cassandra schema version
[**UpgradeSchemaVersion**](CassandraSchemaUpgradeAPI.md#UpgradeSchemaVersion) | **Post** /cassandra/version/upgrade | Upgrade to a specific Cassandra schema version
[**UpgradeToLatestSchemaVersion**](CassandraSchemaUpgradeAPI.md#UpgradeToLatestSchemaVersion) | **Post** /cassandra/version/upgrade/latest | Upgrade to the latest Cassandra schema version



## GetLatestAvailableSchemaVersion

> GetLatestAvailableSchemaVersion200Response GetLatestAvailableSchemaVersion(ctx).Execute()

Retrieve latest available Cassandra schema version

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
    resp, r, err := apiClient.CassandraSchemaUpgradeAPI.GetLatestAvailableSchemaVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CassandraSchemaUpgradeAPI.GetLatestAvailableSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLatestAvailableSchemaVersion`: GetLatestAvailableSchemaVersion200Response
    fmt.Fprintf(os.Stdout, "Response from `CassandraSchemaUpgradeAPI.GetLatestAvailableSchemaVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLatestAvailableSchemaVersionRequest struct via the builder pattern


### Return type

[**GetLatestAvailableSchemaVersion200Response**](GetLatestAvailableSchemaVersion200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSchemaVersion

> GetSchemaVersion200Response GetSchemaVersion(ctx).Execute()

Retrieve current Cassandra schema version

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
    resp, r, err := apiClient.CassandraSchemaUpgradeAPI.GetSchemaVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CassandraSchemaUpgradeAPI.GetSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSchemaVersion`: GetSchemaVersion200Response
    fmt.Fprintf(os.Stdout, "Response from `CassandraSchemaUpgradeAPI.GetSchemaVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSchemaVersionRequest struct via the builder pattern


### Return type

[**GetSchemaVersion200Response**](GetSchemaVersion200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeSchemaVersion

> UpgradeSchemaVersion200Response UpgradeSchemaVersion(ctx).GetLatestAvailableSchemaVersion200Response(getLatestAvailableSchemaVersion200Response).Execute()

Upgrade to a specific Cassandra schema version

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
    getLatestAvailableSchemaVersion200Response := *openapiclient.NewGetLatestAvailableSchemaVersion200Response() // GetLatestAvailableSchemaVersion200Response | The target schema version to upgrade to

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CassandraSchemaUpgradeAPI.UpgradeSchemaVersion(context.Background()).GetLatestAvailableSchemaVersion200Response(getLatestAvailableSchemaVersion200Response).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CassandraSchemaUpgradeAPI.UpgradeSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeSchemaVersion`: UpgradeSchemaVersion200Response
    fmt.Fprintf(os.Stdout, "Response from `CassandraSchemaUpgradeAPI.UpgradeSchemaVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeSchemaVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **getLatestAvailableSchemaVersion200Response** | [**GetLatestAvailableSchemaVersion200Response**](GetLatestAvailableSchemaVersion200Response.md) | The target schema version to upgrade to | 

### Return type

[**UpgradeSchemaVersion200Response**](UpgradeSchemaVersion200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpgradeToLatestSchemaVersion

> UpgradeToLatestSchemaVersion200Response UpgradeToLatestSchemaVersion(ctx).Execute()

Upgrade to the latest Cassandra schema version

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
    resp, r, err := apiClient.CassandraSchemaUpgradeAPI.UpgradeToLatestSchemaVersion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CassandraSchemaUpgradeAPI.UpgradeToLatestSchemaVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpgradeToLatestSchemaVersion`: UpgradeToLatestSchemaVersion200Response
    fmt.Fprintf(os.Stdout, "Response from `CassandraSchemaUpgradeAPI.UpgradeToLatestSchemaVersion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpgradeToLatestSchemaVersionRequest struct via the builder pattern


### Return type

[**UpgradeToLatestSchemaVersion200Response**](UpgradeToLatestSchemaVersion200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

