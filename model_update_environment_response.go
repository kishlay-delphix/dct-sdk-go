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

// UpdateEnvironmentResponse struct for UpdateEnvironmentResponse
type UpdateEnvironmentResponse struct {
	// The initiated job id.
	JobId *string `json:"job_id,omitempty"`
}

// NewUpdateEnvironmentResponse instantiates a new UpdateEnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateEnvironmentResponse() *UpdateEnvironmentResponse {
	this := UpdateEnvironmentResponse{}
	return &this
}

// NewUpdateEnvironmentResponseWithDefaults instantiates a new UpdateEnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateEnvironmentResponseWithDefaults() *UpdateEnvironmentResponse {
	this := UpdateEnvironmentResponse{}
	return &this
}

// GetJobId returns the JobId field value if set, zero value otherwise.
func (o *UpdateEnvironmentResponse) GetJobId() string {
	if o == nil || o.JobId == nil {
		var ret string
		return ret
	}
	return *o.JobId
}

// GetJobIdOk returns a tuple with the JobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateEnvironmentResponse) GetJobIdOk() (*string, bool) {
	if o == nil || o.JobId == nil {
		return nil, false
	}
	return o.JobId, true
}

// HasJobId returns a boolean if a field has been set.
func (o *UpdateEnvironmentResponse) HasJobId() bool {
	if o != nil && o.JobId != nil {
		return true
	}

	return false
}

// SetJobId gets a reference to the given string and assigns it to the JobId field.
func (o *UpdateEnvironmentResponse) SetJobId(v string) {
	o.JobId = &v
}

func (o UpdateEnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.JobId != nil {
		toSerialize["job_id"] = o.JobId
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateEnvironmentResponse struct {
	value *UpdateEnvironmentResponse
	isSet bool
}

func (v NullableUpdateEnvironmentResponse) Get() *UpdateEnvironmentResponse {
	return v.value
}

func (v *NullableUpdateEnvironmentResponse) Set(val *UpdateEnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateEnvironmentResponse(val *UpdateEnvironmentResponse) *NullableUpdateEnvironmentResponse {
	return &NullableUpdateEnvironmentResponse{value: val, isSet: true}
}

func (v NullableUpdateEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


