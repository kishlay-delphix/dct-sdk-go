/*
Delphix DCT API

Delphix DCT API

API version: 3.1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// checks if the StopVDBResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StopVDBResponse{}

// StopVDBResponse struct for StopVDBResponse
type StopVDBResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewStopVDBResponse instantiates a new StopVDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStopVDBResponse() *StopVDBResponse {
	this := StopVDBResponse{}
	return &this
}

// NewStopVDBResponseWithDefaults instantiates a new StopVDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStopVDBResponseWithDefaults() *StopVDBResponse {
	this := StopVDBResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *StopVDBResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StopVDBResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *StopVDBResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *StopVDBResponse) SetJob(v Job) {
	o.Job = &v
}

func (o StopVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StopVDBResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableStopVDBResponse struct {
	value *StopVDBResponse
	isSet bool
}

func (v NullableStopVDBResponse) Get() *StopVDBResponse {
	return v.value
}

func (v *NullableStopVDBResponse) Set(val *StopVDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableStopVDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableStopVDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopVDBResponse(val *StopVDBResponse) *NullableStopVDBResponse {
	return &NullableStopVDBResponse{value: val, isSet: true}
}

func (v NullableStopVDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopVDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


