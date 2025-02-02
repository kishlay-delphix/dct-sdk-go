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

// checks if the RefreshVDBByTimestampResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshVDBByTimestampResponse{}

// RefreshVDBByTimestampResponse struct for RefreshVDBByTimestampResponse
type RefreshVDBByTimestampResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewRefreshVDBByTimestampResponse instantiates a new RefreshVDBByTimestampResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBByTimestampResponse() *RefreshVDBByTimestampResponse {
	this := RefreshVDBByTimestampResponse{}
	return &this
}

// NewRefreshVDBByTimestampResponseWithDefaults instantiates a new RefreshVDBByTimestampResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBByTimestampResponseWithDefaults() *RefreshVDBByTimestampResponse {
	this := RefreshVDBByTimestampResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *RefreshVDBByTimestampResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBByTimestampResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *RefreshVDBByTimestampResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *RefreshVDBByTimestampResponse) SetJob(v Job) {
	o.Job = &v
}

func (o RefreshVDBByTimestampResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshVDBByTimestampResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableRefreshVDBByTimestampResponse struct {
	value *RefreshVDBByTimestampResponse
	isSet bool
}

func (v NullableRefreshVDBByTimestampResponse) Get() *RefreshVDBByTimestampResponse {
	return v.value
}

func (v *NullableRefreshVDBByTimestampResponse) Set(val *RefreshVDBByTimestampResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBByTimestampResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBByTimestampResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBByTimestampResponse(val *RefreshVDBByTimestampResponse) *NullableRefreshVDBByTimestampResponse {
	return &NullableRefreshVDBByTimestampResponse{value: val, isSet: true}
}

func (v NullableRefreshVDBByTimestampResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBByTimestampResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


