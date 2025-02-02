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

// checks if the ProvisionVDBGroupFromBookmarkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionVDBGroupFromBookmarkResponse{}

// ProvisionVDBGroupFromBookmarkResponse struct for ProvisionVDBGroupFromBookmarkResponse
type ProvisionVDBGroupFromBookmarkResponse struct {
	VdbGroup *VDBGroup `json:"vdb_group,omitempty"`
	Job *Job `json:"job,omitempty"`
}

// NewProvisionVDBGroupFromBookmarkResponse instantiates a new ProvisionVDBGroupFromBookmarkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionVDBGroupFromBookmarkResponse() *ProvisionVDBGroupFromBookmarkResponse {
	this := ProvisionVDBGroupFromBookmarkResponse{}
	return &this
}

// NewProvisionVDBGroupFromBookmarkResponseWithDefaults instantiates a new ProvisionVDBGroupFromBookmarkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionVDBGroupFromBookmarkResponseWithDefaults() *ProvisionVDBGroupFromBookmarkResponse {
	this := ProvisionVDBGroupFromBookmarkResponse{}
	return &this
}

// GetVdbGroup returns the VdbGroup field value if set, zero value otherwise.
func (o *ProvisionVDBGroupFromBookmarkResponse) GetVdbGroup() VDBGroup {
	if o == nil || IsNil(o.VdbGroup) {
		var ret VDBGroup
		return ret
	}
	return *o.VdbGroup
}

// GetVdbGroupOk returns a tuple with the VdbGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBGroupFromBookmarkResponse) GetVdbGroupOk() (*VDBGroup, bool) {
	if o == nil || IsNil(o.VdbGroup) {
		return nil, false
	}
	return o.VdbGroup, true
}

// HasVdbGroup returns a boolean if a field has been set.
func (o *ProvisionVDBGroupFromBookmarkResponse) HasVdbGroup() bool {
	if o != nil && !IsNil(o.VdbGroup) {
		return true
	}

	return false
}

// SetVdbGroup gets a reference to the given VDBGroup and assigns it to the VdbGroup field.
func (o *ProvisionVDBGroupFromBookmarkResponse) SetVdbGroup(v VDBGroup) {
	o.VdbGroup = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *ProvisionVDBGroupFromBookmarkResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBGroupFromBookmarkResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *ProvisionVDBGroupFromBookmarkResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *ProvisionVDBGroupFromBookmarkResponse) SetJob(v Job) {
	o.Job = &v
}

func (o ProvisionVDBGroupFromBookmarkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionVDBGroupFromBookmarkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.VdbGroup) {
		toSerialize["vdb_group"] = o.VdbGroup
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableProvisionVDBGroupFromBookmarkResponse struct {
	value *ProvisionVDBGroupFromBookmarkResponse
	isSet bool
}

func (v NullableProvisionVDBGroupFromBookmarkResponse) Get() *ProvisionVDBGroupFromBookmarkResponse {
	return v.value
}

func (v *NullableProvisionVDBGroupFromBookmarkResponse) Set(val *ProvisionVDBGroupFromBookmarkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionVDBGroupFromBookmarkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionVDBGroupFromBookmarkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionVDBGroupFromBookmarkResponse(val *ProvisionVDBGroupFromBookmarkResponse) *NullableProvisionVDBGroupFromBookmarkResponse {
	return &NullableProvisionVDBGroupFromBookmarkResponse{value: val, isSet: true}
}

func (v NullableProvisionVDBGroupFromBookmarkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionVDBGroupFromBookmarkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


