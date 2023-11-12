# ExecutionReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdditionalInformation** | Pointer to **map[string]interface{}** | Additional information about the task | [optional] 
**CancelledDate** | Pointer to **NullableTime** | The cancellation date of the task (null if not cancelled) | [optional] 
**CompletedDate** | Pointer to **time.Time** | The completion date of the task execution | [optional] 
**FailedDate** | Pointer to **NullableTime** | The failure date of the task (null if not failed) | [optional] 
**StartedDate** | Pointer to **time.Time** | The start date of the task execution | [optional] 
**Status** | Pointer to **string** | The status of the task | [optional] 
**SubmitDate** | Pointer to **time.Time** | The submission date of the task | [optional] 
**TaskId** | Pointer to **string** | The ID of the task | [optional] 
**Type** | Pointer to **string** | The type of the task | [optional] 

## Methods

### NewExecutionReport

`func NewExecutionReport() *ExecutionReport`

NewExecutionReport instantiates a new ExecutionReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionReportWithDefaults

`func NewExecutionReportWithDefaults() *ExecutionReport`

NewExecutionReportWithDefaults instantiates a new ExecutionReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdditionalInformation

`func (o *ExecutionReport) GetAdditionalInformation() map[string]interface{}`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *ExecutionReport) GetAdditionalInformationOk() (*map[string]interface{}, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *ExecutionReport) SetAdditionalInformation(v map[string]interface{})`

SetAdditionalInformation sets AdditionalInformation field to given value.

### HasAdditionalInformation

`func (o *ExecutionReport) HasAdditionalInformation() bool`

HasAdditionalInformation returns a boolean if a field has been set.

### GetCancelledDate

`func (o *ExecutionReport) GetCancelledDate() time.Time`

GetCancelledDate returns the CancelledDate field if non-nil, zero value otherwise.

### GetCancelledDateOk

`func (o *ExecutionReport) GetCancelledDateOk() (*time.Time, bool)`

GetCancelledDateOk returns a tuple with the CancelledDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledDate

`func (o *ExecutionReport) SetCancelledDate(v time.Time)`

SetCancelledDate sets CancelledDate field to given value.

### HasCancelledDate

`func (o *ExecutionReport) HasCancelledDate() bool`

HasCancelledDate returns a boolean if a field has been set.

### SetCancelledDateNil

`func (o *ExecutionReport) SetCancelledDateNil(b bool)`

 SetCancelledDateNil sets the value for CancelledDate to be an explicit nil

### UnsetCancelledDate
`func (o *ExecutionReport) UnsetCancelledDate()`

UnsetCancelledDate ensures that no value is present for CancelledDate, not even an explicit nil
### GetCompletedDate

`func (o *ExecutionReport) GetCompletedDate() time.Time`

GetCompletedDate returns the CompletedDate field if non-nil, zero value otherwise.

### GetCompletedDateOk

`func (o *ExecutionReport) GetCompletedDateOk() (*time.Time, bool)`

GetCompletedDateOk returns a tuple with the CompletedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedDate

`func (o *ExecutionReport) SetCompletedDate(v time.Time)`

SetCompletedDate sets CompletedDate field to given value.

### HasCompletedDate

`func (o *ExecutionReport) HasCompletedDate() bool`

HasCompletedDate returns a boolean if a field has been set.

### GetFailedDate

`func (o *ExecutionReport) GetFailedDate() time.Time`

GetFailedDate returns the FailedDate field if non-nil, zero value otherwise.

### GetFailedDateOk

`func (o *ExecutionReport) GetFailedDateOk() (*time.Time, bool)`

GetFailedDateOk returns a tuple with the FailedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedDate

`func (o *ExecutionReport) SetFailedDate(v time.Time)`

SetFailedDate sets FailedDate field to given value.

### HasFailedDate

`func (o *ExecutionReport) HasFailedDate() bool`

HasFailedDate returns a boolean if a field has been set.

### SetFailedDateNil

`func (o *ExecutionReport) SetFailedDateNil(b bool)`

 SetFailedDateNil sets the value for FailedDate to be an explicit nil

### UnsetFailedDate
`func (o *ExecutionReport) UnsetFailedDate()`

UnsetFailedDate ensures that no value is present for FailedDate, not even an explicit nil
### GetStartedDate

`func (o *ExecutionReport) GetStartedDate() time.Time`

GetStartedDate returns the StartedDate field if non-nil, zero value otherwise.

### GetStartedDateOk

`func (o *ExecutionReport) GetStartedDateOk() (*time.Time, bool)`

GetStartedDateOk returns a tuple with the StartedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedDate

`func (o *ExecutionReport) SetStartedDate(v time.Time)`

SetStartedDate sets StartedDate field to given value.

### HasStartedDate

`func (o *ExecutionReport) HasStartedDate() bool`

HasStartedDate returns a boolean if a field has been set.

### GetStatus

`func (o *ExecutionReport) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExecutionReport) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExecutionReport) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExecutionReport) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubmitDate

`func (o *ExecutionReport) GetSubmitDate() time.Time`

GetSubmitDate returns the SubmitDate field if non-nil, zero value otherwise.

### GetSubmitDateOk

`func (o *ExecutionReport) GetSubmitDateOk() (*time.Time, bool)`

GetSubmitDateOk returns a tuple with the SubmitDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmitDate

`func (o *ExecutionReport) SetSubmitDate(v time.Time)`

SetSubmitDate sets SubmitDate field to given value.

### HasSubmitDate

`func (o *ExecutionReport) HasSubmitDate() bool`

HasSubmitDate returns a boolean if a field has been set.

### GetTaskId

`func (o *ExecutionReport) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *ExecutionReport) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *ExecutionReport) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *ExecutionReport) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### GetType

`func (o *ExecutionReport) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ExecutionReport) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ExecutionReport) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ExecutionReport) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


