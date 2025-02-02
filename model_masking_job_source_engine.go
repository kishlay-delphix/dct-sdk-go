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

// checks if the MaskingJobSourceEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskingJobSourceEngine{}

// MaskingJobSourceEngine A masking job's source engine.
type MaskingJobSourceEngine struct {
	// The MaskingJob entity ID.
	MaskingJobId *string `json:"masking_job_id,omitempty"`
	// The ID of the Engine serving as the source for the MaskingJob.
	SourceEngineId *string `json:"source_engine_id,omitempty"`
}

// NewMaskingJobSourceEngine instantiates a new MaskingJobSourceEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskingJobSourceEngine() *MaskingJobSourceEngine {
	this := MaskingJobSourceEngine{}
	return &this
}

// NewMaskingJobSourceEngineWithDefaults instantiates a new MaskingJobSourceEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskingJobSourceEngineWithDefaults() *MaskingJobSourceEngine {
	this := MaskingJobSourceEngine{}
	return &this
}

// GetMaskingJobId returns the MaskingJobId field value if set, zero value otherwise.
func (o *MaskingJobSourceEngine) GetMaskingJobId() string {
	if o == nil || IsNil(o.MaskingJobId) {
		var ret string
		return ret
	}
	return *o.MaskingJobId
}

// GetMaskingJobIdOk returns a tuple with the MaskingJobId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJobSourceEngine) GetMaskingJobIdOk() (*string, bool) {
	if o == nil || IsNil(o.MaskingJobId) {
		return nil, false
	}
	return o.MaskingJobId, true
}

// HasMaskingJobId returns a boolean if a field has been set.
func (o *MaskingJobSourceEngine) HasMaskingJobId() bool {
	if o != nil && !IsNil(o.MaskingJobId) {
		return true
	}

	return false
}

// SetMaskingJobId gets a reference to the given string and assigns it to the MaskingJobId field.
func (o *MaskingJobSourceEngine) SetMaskingJobId(v string) {
	o.MaskingJobId = &v
}

// GetSourceEngineId returns the SourceEngineId field value if set, zero value otherwise.
func (o *MaskingJobSourceEngine) GetSourceEngineId() string {
	if o == nil || IsNil(o.SourceEngineId) {
		var ret string
		return ret
	}
	return *o.SourceEngineId
}

// GetSourceEngineIdOk returns a tuple with the SourceEngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingJobSourceEngine) GetSourceEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceEngineId) {
		return nil, false
	}
	return o.SourceEngineId, true
}

// HasSourceEngineId returns a boolean if a field has been set.
func (o *MaskingJobSourceEngine) HasSourceEngineId() bool {
	if o != nil && !IsNil(o.SourceEngineId) {
		return true
	}

	return false
}

// SetSourceEngineId gets a reference to the given string and assigns it to the SourceEngineId field.
func (o *MaskingJobSourceEngine) SetSourceEngineId(v string) {
	o.SourceEngineId = &v
}

func (o MaskingJobSourceEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskingJobSourceEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaskingJobId) {
		toSerialize["masking_job_id"] = o.MaskingJobId
	}
	if !IsNil(o.SourceEngineId) {
		toSerialize["source_engine_id"] = o.SourceEngineId
	}
	return toSerialize, nil
}

type NullableMaskingJobSourceEngine struct {
	value *MaskingJobSourceEngine
	isSet bool
}

func (v NullableMaskingJobSourceEngine) Get() *MaskingJobSourceEngine {
	return v.value
}

func (v *NullableMaskingJobSourceEngine) Set(val *MaskingJobSourceEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskingJobSourceEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskingJobSourceEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskingJobSourceEngine(val *MaskingJobSourceEngine) *NullableMaskingJobSourceEngine {
	return &NullableMaskingJobSourceEngine{value: val, isSet: true}
}

func (v NullableMaskingJobSourceEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskingJobSourceEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


