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

// checks if the CreateEnvironmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateEnvironmentResponse{}

// CreateEnvironmentResponse struct for CreateEnvironmentResponse
type CreateEnvironmentResponse struct {
	Job *Job `json:"job,omitempty"`
	// The id of environment created.
	EnvironmentId *string `json:"environment_id,omitempty"`
}

// NewCreateEnvironmentResponse instantiates a new CreateEnvironmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateEnvironmentResponse() *CreateEnvironmentResponse {
	this := CreateEnvironmentResponse{}
	return &this
}

// NewCreateEnvironmentResponseWithDefaults instantiates a new CreateEnvironmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateEnvironmentResponseWithDefaults() *CreateEnvironmentResponse {
	this := CreateEnvironmentResponse{}
	return &this
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CreateEnvironmentResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CreateEnvironmentResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *CreateEnvironmentResponse) SetJob(v Job) {
	o.Job = &v
}

// GetEnvironmentId returns the EnvironmentId field value if set, zero value otherwise.
func (o *CreateEnvironmentResponse) GetEnvironmentId() string {
	if o == nil || IsNil(o.EnvironmentId) {
		var ret string
		return ret
	}
	return *o.EnvironmentId
}

// GetEnvironmentIdOk returns a tuple with the EnvironmentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateEnvironmentResponse) GetEnvironmentIdOk() (*string, bool) {
	if o == nil || IsNil(o.EnvironmentId) {
		return nil, false
	}
	return o.EnvironmentId, true
}

// HasEnvironmentId returns a boolean if a field has been set.
func (o *CreateEnvironmentResponse) HasEnvironmentId() bool {
	if o != nil && !IsNil(o.EnvironmentId) {
		return true
	}

	return false
}

// SetEnvironmentId gets a reference to the given string and assigns it to the EnvironmentId field.
func (o *CreateEnvironmentResponse) SetEnvironmentId(v string) {
	o.EnvironmentId = &v
}

func (o CreateEnvironmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateEnvironmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	if !IsNil(o.EnvironmentId) {
		toSerialize["environment_id"] = o.EnvironmentId
	}
	return toSerialize, nil
}

type NullableCreateEnvironmentResponse struct {
	value *CreateEnvironmentResponse
	isSet bool
}

func (v NullableCreateEnvironmentResponse) Get() *CreateEnvironmentResponse {
	return v.value
}

func (v *NullableCreateEnvironmentResponse) Set(val *CreateEnvironmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateEnvironmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateEnvironmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateEnvironmentResponse(val *CreateEnvironmentResponse) *NullableCreateEnvironmentResponse {
	return &NullableCreateEnvironmentResponse{value: val, isSet: true}
}

func (v NullableCreateEnvironmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateEnvironmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


