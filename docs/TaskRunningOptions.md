# TaskRunningOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagesPerSecond** | Pointer to **int32** |  | [optional] 
**Mode** | Pointer to **string** |  | [optional] 

## Methods

### NewTaskRunningOptions

`func NewTaskRunningOptions() *TaskRunningOptions`

NewTaskRunningOptions instantiates a new TaskRunningOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaskRunningOptionsWithDefaults

`func NewTaskRunningOptionsWithDefaults() *TaskRunningOptions`

NewTaskRunningOptionsWithDefaults instantiates a new TaskRunningOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagesPerSecond

`func (o *TaskRunningOptions) GetMessagesPerSecond() int32`

GetMessagesPerSecond returns the MessagesPerSecond field if non-nil, zero value otherwise.

### GetMessagesPerSecondOk

`func (o *TaskRunningOptions) GetMessagesPerSecondOk() (*int32, bool)`

GetMessagesPerSecondOk returns a tuple with the MessagesPerSecond field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagesPerSecond

`func (o *TaskRunningOptions) SetMessagesPerSecond(v int32)`

SetMessagesPerSecond sets MessagesPerSecond field to given value.

### HasMessagesPerSecond

`func (o *TaskRunningOptions) HasMessagesPerSecond() bool`

HasMessagesPerSecond returns a boolean if a field has been set.

### GetMode

`func (o *TaskRunningOptions) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TaskRunningOptions) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TaskRunningOptions) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *TaskRunningOptions) HasMode() bool`

HasMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


