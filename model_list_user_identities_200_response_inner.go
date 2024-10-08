/*
Apache JAMES Web Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ListUserIdentities200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUserIdentities200ResponseInner{}

// ListUserIdentities200ResponseInner struct for ListUserIdentities200ResponseInner
type ListUserIdentities200ResponseInner struct {
	Bcc           []ListUserIdentities200ResponseInnerBccInner `json:"bcc,omitempty"`
	Email         *string                                      `json:"email,omitempty"`
	HtmlSignature *string                                      `json:"htmlSignature,omitempty"`
	Id            *string                                      `json:"id,omitempty"`
	MayDelete     *bool                                        `json:"mayDelete,omitempty"`
	Name          *string                                      `json:"name,omitempty"`
	ReplyTo       []ListUserIdentities200ResponseInnerBccInner `json:"replyTo,omitempty"`
	SortOrder     *int32                                       `json:"sortOrder,omitempty"`
	TextSignature *string                                      `json:"textSignature,omitempty"`
}

// NewListUserIdentities200ResponseInner instantiates a new ListUserIdentities200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserIdentities200ResponseInner() *ListUserIdentities200ResponseInner {
	this := ListUserIdentities200ResponseInner{}
	return &this
}

// NewListUserIdentities200ResponseInnerWithDefaults instantiates a new ListUserIdentities200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserIdentities200ResponseInnerWithDefaults() *ListUserIdentities200ResponseInner {
	this := ListUserIdentities200ResponseInner{}
	return &this
}

// GetBcc returns the Bcc field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetBcc() []ListUserIdentities200ResponseInnerBccInner {
	if o == nil || IsNil(o.Bcc) {
		var ret []ListUserIdentities200ResponseInnerBccInner
		return ret
	}
	return o.Bcc
}

// GetBccOk returns a tuple with the Bcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetBccOk() ([]ListUserIdentities200ResponseInnerBccInner, bool) {
	if o == nil || IsNil(o.Bcc) {
		return nil, false
	}
	return o.Bcc, true
}

// HasBcc returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasBcc() bool {
	if o != nil && !IsNil(o.Bcc) {
		return true
	}

	return false
}

// SetBcc gets a reference to the given []ListUserIdentities200ResponseInnerBccInner and assigns it to the Bcc field.
func (o *ListUserIdentities200ResponseInner) SetBcc(v []ListUserIdentities200ResponseInnerBccInner) {
	o.Bcc = v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ListUserIdentities200ResponseInner) SetEmail(v string) {
	o.Email = &v
}

// GetHtmlSignature returns the HtmlSignature field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetHtmlSignature() string {
	if o == nil || IsNil(o.HtmlSignature) {
		var ret string
		return ret
	}
	return *o.HtmlSignature
}

// GetHtmlSignatureOk returns a tuple with the HtmlSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetHtmlSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.HtmlSignature) {
		return nil, false
	}
	return o.HtmlSignature, true
}

// HasHtmlSignature returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasHtmlSignature() bool {
	if o != nil && !IsNil(o.HtmlSignature) {
		return true
	}

	return false
}

// SetHtmlSignature gets a reference to the given string and assigns it to the HtmlSignature field.
func (o *ListUserIdentities200ResponseInner) SetHtmlSignature(v string) {
	o.HtmlSignature = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListUserIdentities200ResponseInner) SetId(v string) {
	o.Id = &v
}

// GetMayDelete returns the MayDelete field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetMayDelete() bool {
	if o == nil || IsNil(o.MayDelete) {
		var ret bool
		return ret
	}
	return *o.MayDelete
}

// GetMayDeleteOk returns a tuple with the MayDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetMayDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.MayDelete) {
		return nil, false
	}
	return o.MayDelete, true
}

// HasMayDelete returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasMayDelete() bool {
	if o != nil && !IsNil(o.MayDelete) {
		return true
	}

	return false
}

// SetMayDelete gets a reference to the given bool and assigns it to the MayDelete field.
func (o *ListUserIdentities200ResponseInner) SetMayDelete(v bool) {
	o.MayDelete = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListUserIdentities200ResponseInner) SetName(v string) {
	o.Name = &v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetReplyTo() []ListUserIdentities200ResponseInnerBccInner {
	if o == nil || IsNil(o.ReplyTo) {
		var ret []ListUserIdentities200ResponseInnerBccInner
		return ret
	}
	return o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetReplyToOk() ([]ListUserIdentities200ResponseInnerBccInner, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given []ListUserIdentities200ResponseInnerBccInner and assigns it to the ReplyTo field.
func (o *ListUserIdentities200ResponseInner) SetReplyTo(v []ListUserIdentities200ResponseInnerBccInner) {
	o.ReplyTo = v
}

// GetSortOrder returns the SortOrder field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetSortOrder() int32 {
	if o == nil || IsNil(o.SortOrder) {
		var ret int32
		return ret
	}
	return *o.SortOrder
}

// GetSortOrderOk returns a tuple with the SortOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetSortOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.SortOrder) {
		return nil, false
	}
	return o.SortOrder, true
}

// HasSortOrder returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasSortOrder() bool {
	if o != nil && !IsNil(o.SortOrder) {
		return true
	}

	return false
}

// SetSortOrder gets a reference to the given int32 and assigns it to the SortOrder field.
func (o *ListUserIdentities200ResponseInner) SetSortOrder(v int32) {
	o.SortOrder = &v
}

// GetTextSignature returns the TextSignature field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInner) GetTextSignature() string {
	if o == nil || IsNil(o.TextSignature) {
		var ret string
		return ret
	}
	return *o.TextSignature
}

// GetTextSignatureOk returns a tuple with the TextSignature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInner) GetTextSignatureOk() (*string, bool) {
	if o == nil || IsNil(o.TextSignature) {
		return nil, false
	}
	return o.TextSignature, true
}

// HasTextSignature returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInner) HasTextSignature() bool {
	if o != nil && !IsNil(o.TextSignature) {
		return true
	}

	return false
}

// SetTextSignature gets a reference to the given string and assigns it to the TextSignature field.
func (o *ListUserIdentities200ResponseInner) SetTextSignature(v string) {
	o.TextSignature = &v
}

func (o ListUserIdentities200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUserIdentities200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bcc) {
		toSerialize["bcc"] = o.Bcc
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.HtmlSignature) {
		toSerialize["htmlSignature"] = o.HtmlSignature
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MayDelete) {
		toSerialize["mayDelete"] = o.MayDelete
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.SortOrder) {
		toSerialize["sortOrder"] = o.SortOrder
	}
	if !IsNil(o.TextSignature) {
		toSerialize["textSignature"] = o.TextSignature
	}
	return toSerialize, nil
}

type NullableListUserIdentities200ResponseInner struct {
	value *ListUserIdentities200ResponseInner
	isSet bool
}

func (v NullableListUserIdentities200ResponseInner) Get() *ListUserIdentities200ResponseInner {
	return v.value
}

func (v *NullableListUserIdentities200ResponseInner) Set(val *ListUserIdentities200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserIdentities200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserIdentities200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserIdentities200ResponseInner(val *ListUserIdentities200ResponseInner) *NullableListUserIdentities200ResponseInner {
	return &NullableListUserIdentities200ResponseInner{value: val, isSet: true}
}

func (v NullableListUserIdentities200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserIdentities200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
