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

// checks if the RefreshVDBBySnapshotResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshVDBBySnapshotResponse{}

// RefreshVDBBySnapshotResponse struct for RefreshVDBBySnapshotResponse
type RefreshVDBBySnapshotResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewRefreshVDBBySnapshotResponse instantiates a new RefreshVDBBySnapshotResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBBySnapshotResponse() *RefreshVDBBySnapshotResponse {
	this := RefreshVDBBySnapshotResponse{}
	return &this
}

// NewRefreshVDBBySnapshotResponseWithDefaults instantiates a new RefreshVDBBySnapshotResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBBySnapshotResponseWithDefaults() *RefreshVDBBySnapshotResponse {
	this := RefreshVDBBySnapshotResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *RefreshVDBBySnapshotResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBBySnapshotResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *RefreshVDBBySnapshotResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *RefreshVDBBySnapshotResponse) SetJob(v Job) {
	o.Job = &v
}

func (o RefreshVDBBySnapshotResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshVDBBySnapshotResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableRefreshVDBBySnapshotResponse struct {
	value *RefreshVDBBySnapshotResponse
	isSet bool
}

func (v NullableRefreshVDBBySnapshotResponse) Get() *RefreshVDBBySnapshotResponse {
	return v.value
}

func (v *NullableRefreshVDBBySnapshotResponse) Set(val *RefreshVDBBySnapshotResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBBySnapshotResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBBySnapshotResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBBySnapshotResponse(val *RefreshVDBBySnapshotResponse) *NullableRefreshVDBBySnapshotResponse {
	return &NullableRefreshVDBBySnapshotResponse{value: val, isSet: true}
}

func (v NullableRefreshVDBBySnapshotResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBBySnapshotResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


