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

// checks if the CorrectGhostMailbox201ResponseAdditionalInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CorrectGhostMailbox201ResponseAdditionalInformation{}

// CorrectGhostMailbox201ResponseAdditionalInformation struct for CorrectGhostMailbox201ResponseAdditionalInformation
type CorrectGhostMailbox201ResponseAdditionalInformation struct {
	MessageFailedCount *int32  `json:"messageFailedCount,omitempty"`
	MessageMovedCount  *int32  `json:"messageMovedCount,omitempty"`
	NewMailboxId       *string `json:"newMailboxId,omitempty"`
	OldMailboxId       *string `json:"oldMailboxId,omitempty"`
	TotalMessageCount  *int32  `json:"totalMessageCount,omitempty"`
}

// NewCorrectGhostMailbox201ResponseAdditionalInformation instantiates a new CorrectGhostMailbox201ResponseAdditionalInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCorrectGhostMailbox201ResponseAdditionalInformation() *CorrectGhostMailbox201ResponseAdditionalInformation {
	this := CorrectGhostMailbox201ResponseAdditionalInformation{}
	return &this
}

// NewCorrectGhostMailbox201ResponseAdditionalInformationWithDefaults instantiates a new CorrectGhostMailbox201ResponseAdditionalInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCorrectGhostMailbox201ResponseAdditionalInformationWithDefaults() *CorrectGhostMailbox201ResponseAdditionalInformation {
	this := CorrectGhostMailbox201ResponseAdditionalInformation{}
	return &this
}

// GetMessageFailedCount returns the MessageFailedCount field value if set, zero value otherwise.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetMessageFailedCount() int32 {
	if o == nil || IsNil(o.MessageFailedCount) {
		var ret int32
		return ret
	}
	return *o.MessageFailedCount
}

// GetMessageFailedCountOk returns a tuple with the MessageFailedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetMessageFailedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageFailedCount) {
		return nil, false
	}
	return o.MessageFailedCount, true
}

// HasMessageFailedCount returns a boolean if a field has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) HasMessageFailedCount() bool {
	if o != nil && !IsNil(o.MessageFailedCount) {
		return true
	}

	return false
}

// SetMessageFailedCount gets a reference to the given int32 and assigns it to the MessageFailedCount field.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) SetMessageFailedCount(v int32) {
	o.MessageFailedCount = &v
}

// GetMessageMovedCount returns the MessageMovedCount field value if set, zero value otherwise.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetMessageMovedCount() int32 {
	if o == nil || IsNil(o.MessageMovedCount) {
		var ret int32
		return ret
	}
	return *o.MessageMovedCount
}

// GetMessageMovedCountOk returns a tuple with the MessageMovedCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetMessageMovedCountOk() (*int32, bool) {
	if o == nil || IsNil(o.MessageMovedCount) {
		return nil, false
	}
	return o.MessageMovedCount, true
}

// HasMessageMovedCount returns a boolean if a field has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) HasMessageMovedCount() bool {
	if o != nil && !IsNil(o.MessageMovedCount) {
		return true
	}

	return false
}

// SetMessageMovedCount gets a reference to the given int32 and assigns it to the MessageMovedCount field.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) SetMessageMovedCount(v int32) {
	o.MessageMovedCount = &v
}

// GetNewMailboxId returns the NewMailboxId field value if set, zero value otherwise.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetNewMailboxId() string {
	if o == nil || IsNil(o.NewMailboxId) {
		var ret string
		return ret
	}
	return *o.NewMailboxId
}

// GetNewMailboxIdOk returns a tuple with the NewMailboxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetNewMailboxIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewMailboxId) {
		return nil, false
	}
	return o.NewMailboxId, true
}

// HasNewMailboxId returns a boolean if a field has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) HasNewMailboxId() bool {
	if o != nil && !IsNil(o.NewMailboxId) {
		return true
	}

	return false
}

// SetNewMailboxId gets a reference to the given string and assigns it to the NewMailboxId field.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) SetNewMailboxId(v string) {
	o.NewMailboxId = &v
}

// GetOldMailboxId returns the OldMailboxId field value if set, zero value otherwise.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetOldMailboxId() string {
	if o == nil || IsNil(o.OldMailboxId) {
		var ret string
		return ret
	}
	return *o.OldMailboxId
}

// GetOldMailboxIdOk returns a tuple with the OldMailboxId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetOldMailboxIdOk() (*string, bool) {
	if o == nil || IsNil(o.OldMailboxId) {
		return nil, false
	}
	return o.OldMailboxId, true
}

// HasOldMailboxId returns a boolean if a field has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) HasOldMailboxId() bool {
	if o != nil && !IsNil(o.OldMailboxId) {
		return true
	}

	return false
}

// SetOldMailboxId gets a reference to the given string and assigns it to the OldMailboxId field.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) SetOldMailboxId(v string) {
	o.OldMailboxId = &v
}

// GetTotalMessageCount returns the TotalMessageCount field value if set, zero value otherwise.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetTotalMessageCount() int32 {
	if o == nil || IsNil(o.TotalMessageCount) {
		var ret int32
		return ret
	}
	return *o.TotalMessageCount
}

// GetTotalMessageCountOk returns a tuple with the TotalMessageCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) GetTotalMessageCountOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalMessageCount) {
		return nil, false
	}
	return o.TotalMessageCount, true
}

// HasTotalMessageCount returns a boolean if a field has been set.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) HasTotalMessageCount() bool {
	if o != nil && !IsNil(o.TotalMessageCount) {
		return true
	}

	return false
}

// SetTotalMessageCount gets a reference to the given int32 and assigns it to the TotalMessageCount field.
func (o *CorrectGhostMailbox201ResponseAdditionalInformation) SetTotalMessageCount(v int32) {
	o.TotalMessageCount = &v
}

func (o CorrectGhostMailbox201ResponseAdditionalInformation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CorrectGhostMailbox201ResponseAdditionalInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MessageFailedCount) {
		toSerialize["messageFailedCount"] = o.MessageFailedCount
	}
	if !IsNil(o.MessageMovedCount) {
		toSerialize["messageMovedCount"] = o.MessageMovedCount
	}
	if !IsNil(o.NewMailboxId) {
		toSerialize["newMailboxId"] = o.NewMailboxId
	}
	if !IsNil(o.OldMailboxId) {
		toSerialize["oldMailboxId"] = o.OldMailboxId
	}
	if !IsNil(o.TotalMessageCount) {
		toSerialize["totalMessageCount"] = o.TotalMessageCount
	}
	return toSerialize, nil
}

type NullableCorrectGhostMailbox201ResponseAdditionalInformation struct {
	value *CorrectGhostMailbox201ResponseAdditionalInformation
	isSet bool
}

func (v NullableCorrectGhostMailbox201ResponseAdditionalInformation) Get() *CorrectGhostMailbox201ResponseAdditionalInformation {
	return v.value
}

func (v *NullableCorrectGhostMailbox201ResponseAdditionalInformation) Set(val *CorrectGhostMailbox201ResponseAdditionalInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableCorrectGhostMailbox201ResponseAdditionalInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableCorrectGhostMailbox201ResponseAdditionalInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCorrectGhostMailbox201ResponseAdditionalInformation(val *CorrectGhostMailbox201ResponseAdditionalInformation) *NullableCorrectGhostMailbox201ResponseAdditionalInformation {
	return &NullableCorrectGhostMailbox201ResponseAdditionalInformation{value: val, isSet: true}
}

func (v NullableCorrectGhostMailbox201ResponseAdditionalInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCorrectGhostMailbox201ResponseAdditionalInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
