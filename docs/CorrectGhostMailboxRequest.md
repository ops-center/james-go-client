# CorrectGhostMailboxRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MergeDestination** | Pointer to **string** | The ID of the destination mailbox (new mailbox) | [optional] 
**MergeOrigin** | Pointer to **string** | The ID of the mailbox to be merged (ghosted mailbox) | [optional] 

## Methods

### NewCorrectGhostMailboxRequest

`func NewCorrectGhostMailboxRequest() *CorrectGhostMailboxRequest`

NewCorrectGhostMailboxRequest instantiates a new CorrectGhostMailboxRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCorrectGhostMailboxRequestWithDefaults

`func NewCorrectGhostMailboxRequestWithDefaults() *CorrectGhostMailboxRequest`

NewCorrectGhostMailboxRequestWithDefaults instantiates a new CorrectGhostMailboxRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMergeDestination

`func (o *CorrectGhostMailboxRequest) GetMergeDestination() string`

GetMergeDestination returns the MergeDestination field if non-nil, zero value otherwise.

### GetMergeDestinationOk

`func (o *CorrectGhostMailboxRequest) GetMergeDestinationOk() (*string, bool)`

GetMergeDestinationOk returns a tuple with the MergeDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeDestination

`func (o *CorrectGhostMailboxRequest) SetMergeDestination(v string)`

SetMergeDestination sets MergeDestination field to given value.

### HasMergeDestination

`func (o *CorrectGhostMailboxRequest) HasMergeDestination() bool`

HasMergeDestination returns a boolean if a field has been set.

### GetMergeOrigin

`func (o *CorrectGhostMailboxRequest) GetMergeOrigin() string`

GetMergeOrigin returns the MergeOrigin field if non-nil, zero value otherwise.

### GetMergeOriginOk

`func (o *CorrectGhostMailboxRequest) GetMergeOriginOk() (*string, bool)`

GetMergeOriginOk returns a tuple with the MergeOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeOrigin

`func (o *CorrectGhostMailboxRequest) SetMergeOrigin(v string)`

SetMergeOrigin sets MergeOrigin field to given value.

### HasMergeOrigin

`func (o *CorrectGhostMailboxRequest) HasMergeOrigin() bool`

HasMergeOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


