# CheckAllComponents200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checks** | Pointer to [**[]HealthCheckResult**](HealthCheckResult.md) |  | [optional] 
**Status** | Pointer to **string** | Aggregated status of all checks | [optional] 

## Methods

### NewCheckAllComponents200Response

`func NewCheckAllComponents200Response() *CheckAllComponents200Response`

NewCheckAllComponents200Response instantiates a new CheckAllComponents200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCheckAllComponents200ResponseWithDefaults

`func NewCheckAllComponents200ResponseWithDefaults() *CheckAllComponents200Response`

NewCheckAllComponents200ResponseWithDefaults instantiates a new CheckAllComponents200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecks

`func (o *CheckAllComponents200Response) GetChecks() []HealthCheckResult`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *CheckAllComponents200Response) GetChecksOk() (*[]HealthCheckResult, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *CheckAllComponents200Response) SetChecks(v []HealthCheckResult)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *CheckAllComponents200Response) HasChecks() bool`

HasChecks returns a boolean if a field has been set.

### GetStatus

`func (o *CheckAllComponents200Response) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CheckAllComponents200Response) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CheckAllComponents200Response) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CheckAllComponents200Response) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


