# ListUserIdentities200ResponseInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bcc** | Pointer to [**[]ListUserIdentities200ResponseInnerBccInner**](ListUserIdentities200ResponseInnerBccInner.md) |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**HtmlSignature** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MayDelete** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to [**[]ListUserIdentities200ResponseInnerBccInner**](ListUserIdentities200ResponseInnerBccInner.md) |  | [optional] 
**SortOrder** | Pointer to **int32** |  | [optional] 
**TextSignature** | Pointer to **string** |  | [optional] 

## Methods

### NewListUserIdentities200ResponseInner

`func NewListUserIdentities200ResponseInner() *ListUserIdentities200ResponseInner`

NewListUserIdentities200ResponseInner instantiates a new ListUserIdentities200ResponseInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserIdentities200ResponseInnerWithDefaults

`func NewListUserIdentities200ResponseInnerWithDefaults() *ListUserIdentities200ResponseInner`

NewListUserIdentities200ResponseInnerWithDefaults instantiates a new ListUserIdentities200ResponseInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBcc

`func (o *ListUserIdentities200ResponseInner) GetBcc() []ListUserIdentities200ResponseInnerBccInner`

GetBcc returns the Bcc field if non-nil, zero value otherwise.

### GetBccOk

`func (o *ListUserIdentities200ResponseInner) GetBccOk() (*[]ListUserIdentities200ResponseInnerBccInner, bool)`

GetBccOk returns a tuple with the Bcc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBcc

`func (o *ListUserIdentities200ResponseInner) SetBcc(v []ListUserIdentities200ResponseInnerBccInner)`

SetBcc sets Bcc field to given value.

### HasBcc

`func (o *ListUserIdentities200ResponseInner) HasBcc() bool`

HasBcc returns a boolean if a field has been set.

### GetEmail

`func (o *ListUserIdentities200ResponseInner) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ListUserIdentities200ResponseInner) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ListUserIdentities200ResponseInner) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ListUserIdentities200ResponseInner) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetHtmlSignature

`func (o *ListUserIdentities200ResponseInner) GetHtmlSignature() string`

GetHtmlSignature returns the HtmlSignature field if non-nil, zero value otherwise.

### GetHtmlSignatureOk

`func (o *ListUserIdentities200ResponseInner) GetHtmlSignatureOk() (*string, bool)`

GetHtmlSignatureOk returns a tuple with the HtmlSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlSignature

`func (o *ListUserIdentities200ResponseInner) SetHtmlSignature(v string)`

SetHtmlSignature sets HtmlSignature field to given value.

### HasHtmlSignature

`func (o *ListUserIdentities200ResponseInner) HasHtmlSignature() bool`

HasHtmlSignature returns a boolean if a field has been set.

### GetId

`func (o *ListUserIdentities200ResponseInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListUserIdentities200ResponseInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListUserIdentities200ResponseInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListUserIdentities200ResponseInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMayDelete

`func (o *ListUserIdentities200ResponseInner) GetMayDelete() bool`

GetMayDelete returns the MayDelete field if non-nil, zero value otherwise.

### GetMayDeleteOk

`func (o *ListUserIdentities200ResponseInner) GetMayDeleteOk() (*bool, bool)`

GetMayDeleteOk returns a tuple with the MayDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMayDelete

`func (o *ListUserIdentities200ResponseInner) SetMayDelete(v bool)`

SetMayDelete sets MayDelete field to given value.

### HasMayDelete

`func (o *ListUserIdentities200ResponseInner) HasMayDelete() bool`

HasMayDelete returns a boolean if a field has been set.

### GetName

`func (o *ListUserIdentities200ResponseInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListUserIdentities200ResponseInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListUserIdentities200ResponseInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListUserIdentities200ResponseInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReplyTo

`func (o *ListUserIdentities200ResponseInner) GetReplyTo() []ListUserIdentities200ResponseInnerBccInner`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *ListUserIdentities200ResponseInner) GetReplyToOk() (*[]ListUserIdentities200ResponseInnerBccInner, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *ListUserIdentities200ResponseInner) SetReplyTo(v []ListUserIdentities200ResponseInnerBccInner)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *ListUserIdentities200ResponseInner) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSortOrder

`func (o *ListUserIdentities200ResponseInner) GetSortOrder() int32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListUserIdentities200ResponseInner) GetSortOrderOk() (*int32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListUserIdentities200ResponseInner) SetSortOrder(v int32)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ListUserIdentities200ResponseInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetTextSignature

`func (o *ListUserIdentities200ResponseInner) GetTextSignature() string`

GetTextSignature returns the TextSignature field if non-nil, zero value otherwise.

### GetTextSignatureOk

`func (o *ListUserIdentities200ResponseInner) GetTextSignatureOk() (*string, bool)`

GetTextSignatureOk returns a tuple with the TextSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextSignature

`func (o *ListUserIdentities200ResponseInner) SetTextSignature(v string)`

SetTextSignature sets TextSignature field to given value.

### HasTextSignature

`func (o *ListUserIdentities200ResponseInner) HasTextSignature() bool`

HasTextSignature returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


