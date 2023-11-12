# UpdateUserIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bcc** | Pointer to [**[]ListUserIdentities200ResponseInnerBccInner**](ListUserIdentities200ResponseInnerBccInner.md) |  | [optional] 
**HtmlSignature** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to [**[]ListUserIdentities200ResponseInnerBccInner**](ListUserIdentities200ResponseInnerBccInner.md) |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**TextSignature** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateUserIdentityRequest

`func NewUpdateUserIdentityRequest() *UpdateUserIdentityRequest`

NewUpdateUserIdentityRequest instantiates a new UpdateUserIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserIdentityRequestWithDefaults

`func NewUpdateUserIdentityRequestWithDefaults() *UpdateUserIdentityRequest`

NewUpdateUserIdentityRequestWithDefaults instantiates a new UpdateUserIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcc

`func (o *UpdateUserIdentityRequest) GetBcc() []ListUserIdentities200ResponseInnerBccInner`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *UpdateUserIdentityRequest) GetBccOk() (*[]ListUserIdentities200ResponseInnerBccInner, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *UpdateUserIdentityRequest) SetBcc(v []ListUserIdentities200ResponseInnerBccInner)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *UpdateUserIdentityRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetHtmlSignature

`func (o *UpdateUserIdentityRequest) GetHtmlSignature() string`

GetHtmlSignature returns the HtmlSignature field if non-nil, zero value otherwise.

### GetHtmlSignatureOk

`func (o *UpdateUserIdentityRequest) GetHtmlSignatureOk() (*string, bool)`

GetHtmlSignatureOk returns a tuple with the HtmlSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlSignature

`func (o *UpdateUserIdentityRequest) SetHtmlSignature(v string)`

SetHtmlSignature sets HtmlSignature field to given value.

### HasHtmlSignature

`func (o *UpdateUserIdentityRequest) HasHtmlSignature() bool`

HasHtmlSignature returns a boolean if a field has been set.

### GetName

`func (o *UpdateUserIdentityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserIdentityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserIdentityRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserIdentityRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplyTo

`func (o *UpdateUserIdentityRequest) GetReplyTo() []ListUserIdentities200ResponseInnerBccInner`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *UpdateUserIdentityRequest) GetReplyToOk() (*[]ListUserIdentities200ResponseInnerBccInner, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *UpdateUserIdentityRequest) SetReplyTo(v []ListUserIdentities200ResponseInnerBccInner)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *UpdateUserIdentityRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSortOrder

`func (o *UpdateUserIdentityRequest) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateUserIdentityRequest) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateUserIdentityRequest) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UpdateUserIdentityRequest) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetTextSignature

`func (o *UpdateUserIdentityRequest) GetTextSignature() string`

GetTextSignature returns the TextSignature field if non-nil, zero value otherwise.

### GetTextSignatureOk

`func (o *UpdateUserIdentityRequest) GetTextSignatureOk() (*string, bool)`

GetTextSignatureOk returns a tuple with the TextSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSignature

`func (o *UpdateUserIdentityRequest) SetTextSignature(v string)`

SetTextSignature sets TextSignature field to given value.

### HasTextSignature

`func (o *UpdateUserIdentityRequest) HasTextSignature() bool`

HasTextSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


