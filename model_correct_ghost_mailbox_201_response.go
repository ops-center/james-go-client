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

// checks if the CorrectGhostMailbox201Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorrectGhostMailbox201Response{}

// CorrectGhostMailbox201Response struct for CorrectGhostMailbox201Response
type CorrectGhostMailbox201Response struct {
	AdditionalInformation *CorrectGhostMailbox201ResponseAdditionalInformation `json:"additionalInformation,omitempty"`
	TaskId                *string                                              `json:"taskId,omitempty"`
	Type                  *string                                              `json:"type,omitempty"`
}

// NewCorrectGhostMailbox201Response instantiates a new CorrectGhostMailbox201Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorrectGhostMailbox201Response() *CorrectGhostMailbox201Response {
	this := CorrectGhostMailbox201Response{}
	return &this
}

// NewCorrectGhostMailbox201ResponseWithDefaults instantiates a new CorrectGhostMailbox201Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorrectGhostMailbox201ResponseWithDefaults() *CorrectGhostMailbox201Response {
	this := CorrectGhostMailbox201Response{}
	return &this
}

// GetAdditionalInformation returns the AdditionalInformation field value if set, zero value otherwise.
func (o *CorrectGhostMailbox201Response) GetAdditionalInformation() CorrectGhostMailbox201ResponseAdditionalInformation {
	if o == nil || IsNil(o.AdditionalInformation) {
		var ret CorrectGhostMailbox201ResponseAdditionalInformation
		return ret
	}
	return *o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectGhostMailbox201Response) GetAdditionalInformationOk() (*CorrectGhostMailbox201ResponseAdditionalInformation, bool) {
	if o == nil || IsNil(o.AdditionalInformation) {
		return nil, false
	}
	return o.AdditionalInformation, true
}

// HasAdditionalInformation returns a boolean if a field has been set.
func (o *CorrectGhostMailbox201Response) HasAdditionalInformation() bool {
	if o != nil && !IsNil(o.AdditionalInformation) {
		return true
	}

	return false
}

// SetAdditionalInformation gets a reference to the given CorrectGhostMailbox201ResponseAdditionalInformation and assigns it to the AdditionalInformation field.
func (o *CorrectGhostMailbox201Response) SetAdditionalInformation(v CorrectGhostMailbox201ResponseAdditionalInformation) {
	o.AdditionalInformation = &v
}

// GetTaskId returns the TaskId field value if set, zero value otherwise.
func (o *CorrectGhostMailbox201Response) GetTaskId() string {
	if o == nil || IsNil(o.TaskId) {
		var ret string
		return ret
	}
	return *o.TaskId
}

// GetTaskIdOk returns a tuple with the TaskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectGhostMailbox201Response) GetTaskIdOk() (*string, bool) {
	if o == nil || IsNil(o.TaskId) {
		return nil, false
	}
	return o.TaskId, true
}

// HasTaskId returns a boolean if a field has been set.
func (o *CorrectGhostMailbox201Response) HasTaskId() bool {
	if o != nil && !IsNil(o.TaskId) {
		return true
	}

	return false
}

// SetTaskId gets a reference to the given string and assigns it to the TaskId field.
func (o *CorrectGhostMailbox201Response) SetTaskId(v string) {
	o.TaskId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CorrectGhostMailbox201Response) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectGhostMailbox201Response) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CorrectGhostMailbox201Response) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CorrectGhostMailbox201Response) SetType(v string) {
	o.Type = &v
}

func (o CorrectGhostMailbox201Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorrectGhostMailbox201Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AdditionalInformation) {
		toSerialize["additionalInformation"] = o.AdditionalInformation
	}
	if !IsNil(o.TaskId) {
		toSerialize["taskId"] = o.TaskId
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCorrectGhostMailbox201Response struct {
	value *CorrectGhostMailbox201Response
	isSet bool
}

func (v NullableCorrectGhostMailbox201Response) Get() *CorrectGhostMailbox201Response {
	return v.value
}

func (v *NullableCorrectGhostMailbox201Response) Set(val *CorrectGhostMailbox201Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCorrectGhostMailbox201Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCorrectGhostMailbox201Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorrectGhostMailbox201Response(val *CorrectGhostMailbox201Response) *NullableCorrectGhostMailbox201Response {
	return &NullableCorrectGhostMailbox201Response{value: val, isSet: true}
}

func (v NullableCorrectGhostMailbox201Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorrectGhostMailbox201Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
