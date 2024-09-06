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

// checks if the UpdateUserSieveQuotaRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserSieveQuotaRequest{}

// UpdateUserSieveQuotaRequest struct for UpdateUserSieveQuotaRequest
type UpdateUserSieveQuotaRequest struct {
	// Bytes count allowed for this user on this server
	Bytes *int32 `json:"bytes,omitempty"`
}

// NewUpdateUserSieveQuotaRequest instantiates a new UpdateUserSieveQuotaRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserSieveQuotaRequest() *UpdateUserSieveQuotaRequest {
	this := UpdateUserSieveQuotaRequest{}
	return &this
}

// NewUpdateUserSieveQuotaRequestWithDefaults instantiates a new UpdateUserSieveQuotaRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserSieveQuotaRequestWithDefaults() *UpdateUserSieveQuotaRequest {
	this := UpdateUserSieveQuotaRequest{}
	return &this
}

// GetBytes returns the Bytes field value if set, zero value otherwise.
func (o *UpdateUserSieveQuotaRequest) GetBytes() int32 {
	if o == nil || IsNil(o.Bytes) {
		var ret int32
		return ret
	}
	return *o.Bytes
}

// GetBytesOk returns a tuple with the Bytes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSieveQuotaRequest) GetBytesOk() (*int32, bool) {
	if o == nil || IsNil(o.Bytes) {
		return nil, false
	}
	return o.Bytes, true
}

// HasBytes returns a boolean if a field has been set.
func (o *UpdateUserSieveQuotaRequest) HasBytes() bool {
	if o != nil && !IsNil(o.Bytes) {
		return true
	}

	return false
}

// SetBytes gets a reference to the given int32 and assigns it to the Bytes field.
func (o *UpdateUserSieveQuotaRequest) SetBytes(v int32) {
	o.Bytes = &v
}

func (o UpdateUserSieveQuotaRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserSieveQuotaRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bytes) {
		toSerialize["bytes"] = o.Bytes
	}
	return toSerialize, nil
}

type NullableUpdateUserSieveQuotaRequest struct {
	value *UpdateUserSieveQuotaRequest
	isSet bool
}

func (v NullableUpdateUserSieveQuotaRequest) Get() *UpdateUserSieveQuotaRequest {
	return v.value
}

func (v *NullableUpdateUserSieveQuotaRequest) Set(val *UpdateUserSieveQuotaRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateUserSieveQuotaRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateUserSieveQuotaRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateUserSieveQuotaRequest(val *UpdateUserSieveQuotaRequest) *NullableUpdateUserSieveQuotaRequest {
	return &NullableUpdateUserSieveQuotaRequest{value: val, isSet: true}
}

func (v NullableUpdateUserSieveQuotaRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateUserSieveQuotaRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
