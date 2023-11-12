# CreateUserIdentityRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bcc** | Pointer to [**[]ListUserIdentities200ResponseInnerBccInner**](ListUserIdentities200ResponseInnerBccInner.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**HtmlSignature** | Pointer to **string** |  | [optional] 
**MayDelete** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to [**[]ListUserIdentities200ResponseInnerBccInner**](ListUserIdentities200ResponseInnerBccInner.md) |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**TextSignature** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateUserIdentityRequest

`func NewCreateUserIdentityRequest() *CreateUserIdentityRequest`

NewCreateUserIdentityRequest instantiates a new CreateUserIdentityRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateUserIdentityRequestWithDefaults

`func NewCreateUserIdentityRequestWithDefaults() *CreateUserIdentityRequest`

NewCreateUserIdentityRequestWithDefaults instantiates a new CreateUserIdentityRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcc

`func (o *CreateUserIdentityRequest) GetBcc() []ListUserIdentities200ResponseInnerBccInner`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *CreateUserIdentityRequest) GetBccOk() (*[]ListUserIdentities200ResponseInnerBccInner, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *CreateUserIdentityRequest) SetBcc(v []ListUserIdentities200ResponseInnerBccInner)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *CreateUserIdentityRequest) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetEmail

`func (o *CreateUserIdentityRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateUserIdentityRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateUserIdentityRequest) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CreateUserIdentityRequest) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHtmlSignature

`func (o *CreateUserIdentityRequest) GetHtmlSignature() string`

GetHtmlSignature returns the HtmlSignature field if non-nil, zero value otherwise.

### GetHtmlSignatureOk

`func (o *CreateUserIdentityRequest) GetHtmlSignatureOk() (*string, bool)`

GetHtmlSignatureOk returns a tuple with the HtmlSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlSignature

`func (o *CreateUserIdentityRequest) SetHtmlSignature(v string)`

SetHtmlSignature sets HtmlSignature field to given value.

### HasHtmlSignature

`func (o *CreateUserIdentityRequest) HasHtmlSignature() bool`

HasHtmlSignature returns a boolean if a field has been set.

### GetMayDelete

`func (o *CreateUserIdentityRequest) GetMayDelete() bool`

GetMayDelete returns the MayDelete field if non-nil, zero value otherwise.

### GetMayDeleteOk

`func (o *CreateUserIdentityRequest) GetMayDeleteOk() (*bool, bool)`

GetMayDeleteOk returns a tuple with the MayDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayDelete

`func (o *CreateUserIdentityRequest) SetMayDelete(v bool)`

SetMayDelete sets MayDelete field to given value.

### HasMayDelete

`func (o *CreateUserIdentityRequest) HasMayDelete() bool`

HasMayDelete returns a boolean if a field has been set.

### GetName

`func (o *CreateUserIdentityRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateUserIdentityRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateUserIdentityRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateUserIdentityRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplyTo

`func (o *CreateUserIdentityRequest) GetReplyTo() []ListUserIdentities200ResponseInnerBccInner`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *CreateUserIdentityRequest) GetReplyToOk() (*[]ListUserIdentities200ResponseInnerBccInner, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *CreateUserIdentityRequest) SetReplyTo(v []ListUserIdentities200ResponseInnerBccInner)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *CreateUserIdentityRequest) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSortOrder

`func (o *CreateUserIdentityRequest) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *CreateUserIdentityRequest) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *CreateUserIdentityRequest) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *CreateUserIdentityRequest) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetTextSignature

`func (o *CreateUserIdentityRequest) GetTextSignature() string`

GetTextSignature returns the TextSignature field if non-nil, zero value otherwise.

### GetTextSignatureOk

`func (o *CreateUserIdentityRequest) GetTextSignatureOk() (*string, bool)`

GetTextSignatureOk returns a tuple with the TextSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSignature

`func (o *CreateUserIdentityRequest) SetTextSignature(v string)`

SetTextSignature sets TextSignature field to given value.

### HasTextSignature

`func (o *CreateUserIdentityRequest) HasTextSignature() bool`

HasTextSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


