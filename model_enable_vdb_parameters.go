/*
Delphix DCT API

Delphix DCT API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// EnableVDBParameters Parameters to enable a VDB.
type EnableVDBParameters struct {
	// Whether to attempt a startup of the VDB after the enable.
	AttemptStart *bool `json:"attempt_start,omitempty"`
}

// NewEnableVDBParameters instantiates a new EnableVDBParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableVDBParameters() *EnableVDBParameters {
	this := EnableVDBParameters{}
	var attemptStart bool = true
	this.AttemptStart = &attemptStart
	return &this
}

// NewEnableVDBParametersWithDefaults instantiates a new EnableVDBParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableVDBParametersWithDefaults() *EnableVDBParameters {
	this := EnableVDBParameters{}
	var attemptStart bool = true
	this.AttemptStart = &attemptStart
	return &this
}

// GetAttemptStart returns the AttemptStart field value if set, zero value otherwise.
func (o *EnableVDBParameters) GetAttemptStart() bool {
	if o == nil || o.AttemptStart == nil {
		var ret bool
		return ret
	}
	return *o.AttemptStart
}

// GetAttemptStartOk returns a tuple with the AttemptStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableVDBParameters) GetAttemptStartOk() (*bool, bool) {
	if o == nil || o.AttemptStart == nil {
		return nil, false
	}
	return o.AttemptStart, true
}

// HasAttemptStart returns a boolean if a field has been set.
func (o *EnableVDBParameters) HasAttemptStart() bool {
	if o != nil && o.AttemptStart != nil {
		return true
	}

	return false
}

// SetAttemptStart gets a reference to the given bool and assigns it to the AttemptStart field.
func (o *EnableVDBParameters) SetAttemptStart(v bool) {
	o.AttemptStart = &v
}

func (o EnableVDBParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AttemptStart != nil {
		toSerialize["attempt_start"] = o.AttemptStart
	}
	return json.Marshal(toSerialize)
}

type NullableEnableVDBParameters struct {
	value *EnableVDBParameters
	isSet bool
}

func (v NullableEnableVDBParameters) Get() *EnableVDBParameters {
	return v.value
}

func (v *NullableEnableVDBParameters) Set(val *EnableVDBParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableVDBParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableVDBParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableVDBParameters(val *EnableVDBParameters) *NullableEnableVDBParameters {
	return &NullableEnableVDBParameters{value: val, isSet: true}
}

func (v NullableEnableVDBParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableVDBParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


