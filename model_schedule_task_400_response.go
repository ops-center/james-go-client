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

// checks if the ScheduleTask400Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleTask400Response{}

// ScheduleTask400Response struct for ScheduleTask400Response
type ScheduleTask400Response struct {
	Error *string `json:"error,omitempty"`
}

// NewScheduleTask400Response instantiates a new ScheduleTask400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleTask400Response() *ScheduleTask400Response {
	this := ScheduleTask400Response{}
	return &this
}

// NewScheduleTask400ResponseWithDefaults instantiates a new ScheduleTask400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleTask400ResponseWithDefaults() *ScheduleTask400Response {
	this := ScheduleTask400Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *ScheduleTask400Response) GetError() string {
	if o == nil || IsNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleTask400Response) GetErrorOk() (*string, bool) {
	if o == nil || IsNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *ScheduleTask400Response) HasError() bool {
	if o != nil && !IsNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *ScheduleTask400Response) SetError(v string) {
	o.Error = &v
}

func (o ScheduleTask400Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleTask400Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Error) {
		toSerialize["error"] = o.Error
	}
	return toSerialize, nil
}

type NullableScheduleTask400Response struct {
	value *ScheduleTask400Response
	isSet bool
}

func (v NullableScheduleTask400Response) Get() *ScheduleTask400Response {
	return v.value
}

func (v *NullableScheduleTask400Response) Set(val *ScheduleTask400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleTask400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleTask400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleTask400Response(val *ScheduleTask400Response) *NullableScheduleTask400Response {
	return &NullableScheduleTask400Response{value: val, isSet: true}
}

func (v NullableScheduleTask400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleTask400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
