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

// checks if the ListUsers200ResponseInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUsers200ResponseInner{}

// ListUsers200ResponseInner struct for ListUsers200ResponseInner
type ListUsers200ResponseInner struct {
	Username *string `json:"username,omitempty"`
}

// NewListUsers200ResponseInner instantiates a new ListUsers200ResponseInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsers200ResponseInner() *ListUsers200ResponseInner {
	this := ListUsers200ResponseInner{}
	return &this
}

// NewListUsers200ResponseInnerWithDefaults instantiates a new ListUsers200ResponseInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsers200ResponseInnerWithDefaults() *ListUsers200ResponseInner {
	this := ListUsers200ResponseInner{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *ListUsers200ResponseInner) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsers200ResponseInner) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *ListUsers200ResponseInner) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *ListUsers200ResponseInner) SetUsername(v string) {
	o.Username = &v
}

func (o ListUsers200ResponseInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUsers200ResponseInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableListUsers200ResponseInner struct {
	value *ListUsers200ResponseInner
	isSet bool
}

func (v NullableListUsers200ResponseInner) Get() *ListUsers200ResponseInner {
	return v.value
}

func (v *NullableListUsers200ResponseInner) Set(val *ListUsers200ResponseInner) {
	v.value = val
	v.isSet = true
}

func (v NullableListUsers200ResponseInner) IsSet() bool {
	return v.isSet
}

func (v *NullableListUsers200ResponseInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListUsers200ResponseInner(val *ListUsers200ResponseInner) *NullableListUsers200ResponseInner {
	return &NullableListUsers200ResponseInner{value: val, isSet: true}
}

func (v NullableListUsers200ResponseInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListUsers200ResponseInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
