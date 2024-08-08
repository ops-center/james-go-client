# \DefaultAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddressGroupsAddGroupsPost**](DefaultAPI.md#AddressGroupsAddGroupsPost) | **Post** /address/groups/add-groups | Add multiple groups with members



## AddressGroupsAddGroupsPost

> AddGroups200Response AddressGroupsAddGroupsPost(ctx).Group(group).Execute()

Add multiple groups with members



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	group := []openapiclient.Group{*openapiclient.NewGroup("Group1", []string{"Member1"})} // []Group | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.AddressGroupsAddGroupsPost(context.Background()).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.AddressGroupsAddGroupsPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddressGroupsAddGroupsPost`: AddGroups200Response
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.AddressGroupsAddGroupsPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddressGroupsAddGroupsPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **group** | [**[]Group**](Group.md) |  | 

### Return type

[**AddGroups200Response**](AddGroups200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

