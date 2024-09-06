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

// checks if the CheckAllComponents200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CheckAllComponents200Response{}

// CheckAllComponents200Response struct for CheckAllComponents200Response
type CheckAllComponents200Response struct {
	Checks []HealthCheckResult `json:"checks,omitempty"`
	// Aggregated status of all checks
	Status *string `json:"status,omitempty"`
}

// NewCheckAllComponents200Response instantiates a new CheckAllComponents200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckAllComponents200Response() *CheckAllComponents200Response {
	this := CheckAllComponents200Response{}
	return &this
}

// NewCheckAllComponents200ResponseWithDefaults instantiates a new CheckAllComponents200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckAllComponents200ResponseWithDefaults() *CheckAllComponents200Response {
	this := CheckAllComponents200Response{}
	return &this
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *CheckAllComponents200Response) GetChecks() []HealthCheckResult {
	if o == nil || IsNil(o.Checks) {
		var ret []HealthCheckResult
		return ret
	}
	return o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAllComponents200Response) GetChecksOk() ([]HealthCheckResult, bool) {
	if o == nil || IsNil(o.Checks) {
		return nil, false
	}
	return o.Checks, true
}

// HasChecks returns a boolean if a field has been set.
func (o *CheckAllComponents200Response) HasChecks() bool {
	if o != nil && !IsNil(o.Checks) {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []HealthCheckResult and assigns it to the Checks field.
func (o *CheckAllComponents200Response) SetChecks(v []HealthCheckResult) {
	o.Checks = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CheckAllComponents200Response) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckAllComponents200Response) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CheckAllComponents200Response) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CheckAllComponents200Response) SetStatus(v string) {
	o.Status = &v
}

func (o CheckAllComponents200Response) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CheckAllComponents200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Checks) {
		toSerialize["checks"] = o.Checks
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableCheckAllComponents200Response struct {
	value *CheckAllComponents200Response
	isSet bool
}

func (v NullableCheckAllComponents200Response) Get() *CheckAllComponents200Response {
	return v.value
}

func (v *NullableCheckAllComponents200Response) Set(val *CheckAllComponents200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckAllComponents200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckAllComponents200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckAllComponents200Response(val *CheckAllComponents200Response) *NullableCheckAllComponents200Response {
	return &NullableCheckAllComponents200Response{value: val, isSet: true}
}

func (v NullableCheckAllComponents200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckAllComponents200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
