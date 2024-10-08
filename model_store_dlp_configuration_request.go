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

// checks if the StoreDLPConfigurationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StoreDLPConfigurationRequest{}

// StoreDLPConfigurationRequest struct for StoreDLPConfigurationRequest
type StoreDLPConfigurationRequest struct {
	Rules []StoreDLPConfigurationRequestRulesInner `json:"rules,omitempty"`
}

// NewStoreDLPConfigurationRequest instantiates a new StoreDLPConfigurationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStoreDLPConfigurationRequest() *StoreDLPConfigurationRequest {
	this := StoreDLPConfigurationRequest{}
	return &this
}

// NewStoreDLPConfigurationRequestWithDefaults instantiates a new StoreDLPConfigurationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStoreDLPConfigurationRequestWithDefaults() *StoreDLPConfigurationRequest {
	this := StoreDLPConfigurationRequest{}
	return &this
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *StoreDLPConfigurationRequest) GetRules() []StoreDLPConfigurationRequestRulesInner {
	if o == nil || IsNil(o.Rules) {
		var ret []StoreDLPConfigurationRequestRulesInner
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StoreDLPConfigurationRequest) GetRulesOk() ([]StoreDLPConfigurationRequestRulesInner, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *StoreDLPConfigurationRequest) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []StoreDLPConfigurationRequestRulesInner and assigns it to the Rules field.
func (o *StoreDLPConfigurationRequest) SetRules(v []StoreDLPConfigurationRequestRulesInner) {
	o.Rules = v
}

func (o StoreDLPConfigurationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StoreDLPConfigurationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableStoreDLPConfigurationRequest struct {
	value *StoreDLPConfigurationRequest
	isSet bool
}

func (v NullableStoreDLPConfigurationRequest) Get() *StoreDLPConfigurationRequest {
	return v.value
}

func (v *NullableStoreDLPConfigurationRequest) Set(val *StoreDLPConfigurationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableStoreDLPConfigurationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableStoreDLPConfigurationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStoreDLPConfigurationRequest(val *StoreDLPConfigurationRequest) *NullableStoreDLPConfigurationRequest {
	return &NullableStoreDLPConfigurationRequest{value: val, isSet: true}
}

func (v NullableStoreDLPConfigurationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStoreDLPConfigurationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
