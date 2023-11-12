# ListMailsOfMailQueue200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**NextDelivery** | Pointer to **time.Time** |  | [optional] 
**Recipients** | Pointer to **[]string** |  | [optional] 
**Sender** | Pointer to **string** |  | [optional] 

## Methods

### NewListMailsOfMailQueue200ResponseInner

`func NewListMailsOfMailQueue200ResponseInner() *ListMailsOfMailQueue200ResponseInner`

NewListMailsOfMailQueue200ResponseInner instantiates a new ListMailsOfMailQueue200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMailsOfMailQueue200ResponseInnerWithDefaults

`func NewListMailsOfMailQueue200ResponseInnerWithDefaults() *ListMailsOfMailQueue200ResponseInner`

NewListMailsOfMailQueue200ResponseInnerWithDefaults instantiates a new ListMailsOfMailQueue200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListMailsOfMailQueue200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMailsOfMailQueue200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMailsOfMailQueue200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListMailsOfMailQueue200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNextDelivery

`func (o *ListMailsOfMailQueue200ResponseInner) GetNextDelivery() time.Time`

GetNextDelivery returns the NextDelivery field if non-nil, zero value otherwise.

### GetNextDeliveryOk

`func (o *ListMailsOfMailQueue200ResponseInner) GetNextDeliveryOk() (*time.Time, bool)`

GetNextDeliveryOk returns a tuple with the NextDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextDelivery

`func (o *ListMailsOfMailQueue200ResponseInner) SetNextDelivery(v time.Time)`

SetNextDelivery sets NextDelivery field to given value.

### HasNextDelivery

`func (o *ListMailsOfMailQueue200ResponseInner) HasNextDelivery() bool`

HasNextDelivery returns a boolean if a field has been set.

### GetRecipients

`func (o *ListMailsOfMailQueue200ResponseInner) GetRecipients() []string`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *ListMailsOfMailQueue200ResponseInner) GetRecipientsOk() (*[]string, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *ListMailsOfMailQueue200ResponseInner) SetRecipients(v []string)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *ListMailsOfMailQueue200ResponseInner) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSender

`func (o *ListMailsOfMailQueue200ResponseInner) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *ListMailsOfMailQueue200ResponseInner) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *ListMailsOfMailQueue200ResponseInner) SetSender(v string)`

SetSender sets Sender field to given value.

### HasSender

`func (o *ListMailsOfMailQueue200ResponseInner) HasSender() bool`

HasSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


