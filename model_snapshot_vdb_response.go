/*
Delphix DCT API

Delphix DCT API

API version: 2.0.0

Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// SnapshotVDBResponse struct for SnapshotVDBResponse
type SnapshotVDBResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewSnapshotVDBResponse instantiates a new SnapshotVDBResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnapshotVDBResponse() *SnapshotVDBResponse {
	this := SnapshotVDBResponse{}
	return &this
}

// NewSnapshotVDBResponseWithDefaults instantiates a new SnapshotVDBResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnapshotVDBResponseWithDefaults() *SnapshotVDBResponse {
	this := SnapshotVDBResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *SnapshotVDBResponse) GetJob() Job {
	if o == nil || o.Job == nil {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnapshotVDBResponse) GetJobOk() (*Job, bool) {
	if o == nil || o.Job == nil {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *SnapshotVDBResponse) HasJob() bool {
	if o != nil && o.Job != nil {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *SnapshotVDBResponse) SetJob(v Job) {
	o.Job = &v
}

func (o SnapshotVDBResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Job != nil {
		toSerialize["job"] = o.Job
	}
	return json.Marshal(toSerialize)
}

type NullableSnapshotVDBResponse struct {
	value *SnapshotVDBResponse
	isSet bool
}

func (v NullableSnapshotVDBResponse) Get() *SnapshotVDBResponse {
	return v.value
}

func (v *NullableSnapshotVDBResponse) Set(val *SnapshotVDBResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSnapshotVDBResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSnapshotVDBResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnapshotVDBResponse(val *SnapshotVDBResponse) *NullableSnapshotVDBResponse {
	return &NullableSnapshotVDBResponse{value: val, isSet: true}
}

func (v NullableSnapshotVDBResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnapshotVDBResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


