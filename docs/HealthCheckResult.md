# HealthCheckResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cause** | Pointer to **NullableString** |  | [optional] 
**ComponentName** | Pointer to **string** |  | [optional] 
**EscapedComponentName** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHealthCheckResult

`func NewHealthCheckResult() *HealthCheckResult`

NewHealthCheckResult instantiates a new HealthCheckResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthCheckResultWithDefaults

`func NewHealthCheckResultWithDefaults() *HealthCheckResult`

NewHealthCheckResultWithDefaults instantiates a new HealthCheckResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCause

`func (o *HealthCheckResult) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *HealthCheckResult) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *HealthCheckResult) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *HealthCheckResult) HasCause() bool`

HasCause returns a boolean if a field has been set.

### SetCauseNil

`func (o *HealthCheckResult) SetCauseNil(b bool)`

 SetCauseNil sets the value for Cause to be an explicit nil

### UnsetCause
`func (o *HealthCheckResult) UnsetCause()`

UnsetCause ensures that no value is present for Cause, not even an explicit nil
### GetComponentName

`func (o *HealthCheckResult) GetComponentName() string`

GetComponentName returns the ComponentName field if non-nil, zero value otherwise.

### GetComponentNameOk

`func (o *HealthCheckResult) GetComponentNameOk() (*string, bool)`

GetComponentNameOk returns a tuple with the ComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentName

`func (o *HealthCheckResult) SetComponentName(v string)`

SetComponentName sets ComponentName field to given value.

### HasComponentName

`func (o *HealthCheckResult) HasComponentName() bool`

HasComponentName returns a boolean if a field has been set.

### GetEscapedComponentName

`func (o *HealthCheckResult) GetEscapedComponentName() string`

GetEscapedComponentName returns the EscapedComponentName field if non-nil, zero value otherwise.

### GetEscapedComponentNameOk

`func (o *HealthCheckResult) GetEscapedComponentNameOk() (*string, bool)`

GetEscapedComponentNameOk returns a tuple with the EscapedComponentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEscapedComponentName

`func (o *HealthCheckResult) SetEscapedComponentName(v string)`

SetEscapedComponentName sets EscapedComponentName field to given value.

### HasEscapedComponentName

`func (o *HealthCheckResult) HasEscapedComponentName() bool`

HasEscapedComponentName returns a boolean if a field has been set.

### GetStatus

`func (o *HealthCheckResult) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthCheckResult) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthCheckResult) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthCheckResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


