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

// checks if the RefreshVDBFromBookmarkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshVDBFromBookmarkResponse{}

// RefreshVDBFromBookmarkResponse struct for RefreshVDBFromBookmarkResponse
type RefreshVDBFromBookmarkResponse struct {
	Job *Job `json:"job,omitempty"`
}

// NewRefreshVDBFromBookmarkResponse instantiates a new RefreshVDBFromBookmarkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBFromBookmarkResponse() *RefreshVDBFromBookmarkResponse {
	this := RefreshVDBFromBookmarkResponse{}
	return &this
}

// NewRefreshVDBFromBookmarkResponseWithDefaults instantiates a new RefreshVDBFromBookmarkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBFromBookmarkResponseWithDefaults() *RefreshVDBFromBookmarkResponse {
	this := RefreshVDBFromBookmarkResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *RefreshVDBFromBookmarkResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBFromBookmarkResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *RefreshVDBFromBookmarkResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *RefreshVDBFromBookmarkResponse) SetJob(v Job) {
	o.Job = &v
}

func (o RefreshVDBFromBookmarkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshVDBFromBookmarkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableRefreshVDBFromBookmarkResponse struct {
	value *RefreshVDBFromBookmarkResponse
	isSet bool
}

func (v NullableRefreshVDBFromBookmarkResponse) Get() *RefreshVDBFromBookmarkResponse {
	return v.value
}

func (v *NullableRefreshVDBFromBookmarkResponse) Set(val *RefreshVDBFromBookmarkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBFromBookmarkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBFromBookmarkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBFromBookmarkResponse(val *RefreshVDBFromBookmarkResponse) *NullableRefreshVDBFromBookmarkResponse {
	return &NullableRefreshVDBFromBookmarkResponse{value: val, isSet: true}
}

func (v NullableRefreshVDBFromBookmarkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBFromBookmarkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


