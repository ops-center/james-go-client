# Task

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddedMessageIdEntries** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]TaskErrorsInner**](TaskErrorsInner.md) |  | [optional] 
**FixedInconsistencies** | Pointer to [**[]TaskErrorsInner**](TaskErrorsInner.md) |  | [optional] 
**ProcessedImapUidEntries** | Pointer to **int32** |  | [optional] 
**ProcessedMessageIdEntries** | Pointer to **int32** |  | [optional] 
**RemovedMessageIdEntries** | Pointer to **int32** |  | [optional] 
**RunningOptions** | Pointer to [**TaskRunningOptions**](TaskRunningOptions.md) |  | [optional] 
**Timestamp** | Pointer to **time.Time** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedMessageIdEntries** | Pointer to **int32** |  | [optional] 

## Methods

### NewTask

`func NewTask() *Task`

NewTask instantiates a new Task object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskWithDefaults

`func NewTaskWithDefaults() *Task`

NewTaskWithDefaults instantiates a new Task object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddedMessageIdEntries

`func (o *Task) GetAddedMessageIdEntries() int32`

GetAddedMessageIdEntries returns the AddedMessageIdEntries field if non-nil, zero value otherwise.

### GetAddedMessageIdEntriesOk

`func (o *Task) GetAddedMessageIdEntriesOk() (*int32, bool)`

GetAddedMessageIdEntriesOk returns a tuple with the AddedMessageIdEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedMessageIdEntries

`func (o *Task) SetAddedMessageIdEntries(v int32)`

SetAddedMessageIdEntries sets AddedMessageIdEntries field to given value.

### HasAddedMessageIdEntries

`func (o *Task) HasAddedMessageIdEntries() bool`

HasAddedMessageIdEntries returns a boolean if a field has been set.

### GetErrors

`func (o *Task) GetErrors() []TaskErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Task) GetErrorsOk() (*[]TaskErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Task) SetErrors(v []TaskErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Task) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetFixedInconsistencies

`func (o *Task) GetFixedInconsistencies() []TaskErrorsInner`

GetFixedInconsistencies returns the FixedInconsistencies field if non-nil, zero value otherwise.

### GetFixedInconsistenciesOk

`func (o *Task) GetFixedInconsistenciesOk() (*[]TaskErrorsInner, bool)`

GetFixedInconsistenciesOk returns a tuple with the FixedInconsistencies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFixedInconsistencies

`func (o *Task) SetFixedInconsistencies(v []TaskErrorsInner)`

SetFixedInconsistencies sets FixedInconsistencies field to given value.

### HasFixedInconsistencies

`func (o *Task) HasFixedInconsistencies() bool`

HasFixedInconsistencies returns a boolean if a field has been set.

### GetProcessedImapUidEntries

`func (o *Task) GetProcessedImapUidEntries() int32`

GetProcessedImapUidEntries returns the ProcessedImapUidEntries field if non-nil, zero value otherwise.

### GetProcessedImapUidEntriesOk

`func (o *Task) GetProcessedImapUidEntriesOk() (*int32, bool)`

GetProcessedImapUidEntriesOk returns a tuple with the ProcessedImapUidEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedImapUidEntries

`func (o *Task) SetProcessedImapUidEntries(v int32)`

SetProcessedImapUidEntries sets ProcessedImapUidEntries field to given value.

### HasProcessedImapUidEntries

`func (o *Task) HasProcessedImapUidEntries() bool`

HasProcessedImapUidEntries returns a boolean if a field has been set.

### GetProcessedMessageIdEntries

`func (o *Task) GetProcessedMessageIdEntries() int32`

GetProcessedMessageIdEntries returns the ProcessedMessageIdEntries field if non-nil, zero value otherwise.

### GetProcessedMessageIdEntriesOk

`func (o *Task) GetProcessedMessageIdEntriesOk() (*int32, bool)`

GetProcessedMessageIdEntriesOk returns a tuple with the ProcessedMessageIdEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessedMessageIdEntries

`func (o *Task) SetProcessedMessageIdEntries(v int32)`

SetProcessedMessageIdEntries sets ProcessedMessageIdEntries field to given value.

### HasProcessedMessageIdEntries

`func (o *Task) HasProcessedMessageIdEntries() bool`

HasProcessedMessageIdEntries returns a boolean if a field has been set.

### GetRemovedMessageIdEntries

`func (o *Task) GetRemovedMessageIdEntries() int32`

GetRemovedMessageIdEntries returns the RemovedMessageIdEntries field if non-nil, zero value otherwise.

### GetRemovedMessageIdEntriesOk

`func (o *Task) GetRemovedMessageIdEntriesOk() (*int32, bool)`

GetRemovedMessageIdEntriesOk returns a tuple with the RemovedMessageIdEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemovedMessageIdEntries

`func (o *Task) SetRemovedMessageIdEntries(v int32)`

SetRemovedMessageIdEntries sets RemovedMessageIdEntries field to given value.

### HasRemovedMessageIdEntries

`func (o *Task) HasRemovedMessageIdEntries() bool`

HasRemovedMessageIdEntries returns a boolean if a field has been set.

### GetRunningOptions

`func (o *Task) GetRunningOptions() TaskRunningOptions`

GetRunningOptions returns the RunningOptions field if non-nil, zero value otherwise.

### GetRunningOptionsOk

`func (o *Task) GetRunningOptionsOk() (*TaskRunningOptions, bool)`

GetRunningOptionsOk returns a tuple with the RunningOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunningOptions

`func (o *Task) SetRunningOptions(v TaskRunningOptions)`

SetRunningOptions sets RunningOptions field to given value.

### HasRunningOptions

`func (o *Task) HasRunningOptions() bool`

HasRunningOptions returns a boolean if a field has been set.

### GetTimestamp

`func (o *Task) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Task) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Task) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Task) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *Task) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Task) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Task) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Task) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedMessageIdEntries

`func (o *Task) GetUpdatedMessageIdEntries() int32`

GetUpdatedMessageIdEntries returns the UpdatedMessageIdEntries field if non-nil, zero value otherwise.

### GetUpdatedMessageIdEntriesOk

`func (o *Task) GetUpdatedMessageIdEntriesOk() (*int32, bool)`

GetUpdatedMessageIdEntriesOk returns a tuple with the UpdatedMessageIdEntries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedMessageIdEntries

`func (o *Task) SetUpdatedMessageIdEntries(v int32)`

SetUpdatedMessageIdEntries sets UpdatedMessageIdEntries field to given value.

### HasUpdatedMessageIdEntries

`func (o *Task) HasUpdatedMessageIdEntries() bool`

HasUpdatedMessageIdEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


