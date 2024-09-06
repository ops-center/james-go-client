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

// checks if the ListUserIdentities200ResponseInnerBccInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUserIdentities200ResponseInnerBccInner{}

// ListUserIdentities200ResponseInnerBccInner struct for ListUserIdentities200ResponseInnerBccInner
type ListUserIdentities200ResponseInnerBccInner struct {
	Email *string `json:"email,omitempty"`
	Name  *string `json:"name,omitempty"`
}

// NewListUserIdentities200ResponseInnerBccInner instantiates a new ListUserIdentities200ResponseInnerBccInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUserIdentities200ResponseInnerBccInner() *ListUserIdentities200ResponseInnerBccInner {
	this := ListUserIdentities200ResponseInnerBccInner{}
	return &this
}

// NewListUserIdentities200ResponseInnerBccInnerWithDefaults instantiates a new ListUserIdentities200ResponseInnerBccInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUserIdentities200ResponseInnerBccInnerWithDefaults() *ListUserIdentities200ResponseInnerBccInner {
	this := ListUserIdentities200ResponseInnerBccInner{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInnerBccInner) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInnerBccInner) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInnerBccInner) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *ListUserIdentities200ResponseInnerBccInner) SetEmail(v string) {
	o.Email = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListUserIdentities200ResponseInnerBccInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUserIdentities200ResponseInnerBccInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ListUserIdentities200ResponseInnerBccInner) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListUserIdentities200ResponseInnerBccInner) SetName(v string) {
	o.Name = &v
}

func (o ListUserIdentities200ResponseInnerBccInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUserIdentities200ResponseInnerBccInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableListUserIdentities200ResponseInnerBccInner struct {
	value *ListUserIdentities200ResponseInnerBccInner
	isSet bool
}

func (v NullableListUserIdentities200ResponseInnerBccInner) Get() *ListUserIdentities200ResponseInnerBccInner {
	return v.value
}

func (v *NullableListUserIdentities200ResponseInnerBccInner) Set(val *ListUserIdentities200ResponseInnerBccInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUserIdentities200ResponseInnerBccInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUserIdentities200ResponseInnerBccInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUserIdentities200ResponseInnerBccInner(val *ListUserIdentities200ResponseInnerBccInner) *NullableListUserIdentities200ResponseInnerBccInner {
	return &NullableListUserIdentities200ResponseInnerBccInner{value: val, isSet: true}
}

func (v NullableListUserIdentities200ResponseInnerBccInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUserIdentities200ResponseInnerBccInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
