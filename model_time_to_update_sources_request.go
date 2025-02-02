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

// checks if the TimeToUpdateSourcesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TimeToUpdateSourcesRequest{}

// TimeToUpdateSourcesRequest struct for TimeToUpdateSourcesRequest
type TimeToUpdateSourcesRequest struct {
	// no. of times same engine needs to be registered
	EnginesCount int32 `json:"engines_count"`
	// list of engine hostnames to be registered engines_count times
	EnginesList []string `json:"engines_list"`
	// no. of actual sources on 1 engine - this no. needs to be same for all engines
	ExistingNoOfSources int32 `json:"existing_no_of_sources"`
	// no. of times every source needs to be saved in data library
	SourcesCount int32 `json:"sources_count"`
}

// NewTimeToUpdateSourcesRequest instantiates a new TimeToUpdateSourcesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTimeToUpdateSourcesRequest(enginesCount int32, enginesList []string, existingNoOfSources int32, sourcesCount int32) *TimeToUpdateSourcesRequest {
	this := TimeToUpdateSourcesRequest{}
	this.EnginesCount = enginesCount
	this.EnginesList = enginesList
	this.ExistingNoOfSources = existingNoOfSources
	this.SourcesCount = sourcesCount
	return &this
}

// NewTimeToUpdateSourcesRequestWithDefaults instantiates a new TimeToUpdateSourcesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTimeToUpdateSourcesRequestWithDefaults() *TimeToUpdateSourcesRequest {
	this := TimeToUpdateSourcesRequest{}
	return &this
}

// GetEnginesCount returns the EnginesCount field value
func (o *TimeToUpdateSourcesRequest) GetEnginesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.EnginesCount
}

// GetEnginesCountOk returns a tuple with the EnginesCount field value
// and a boolean to check if the value has been set.
func (o *TimeToUpdateSourcesRequest) GetEnginesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EnginesCount, true
}

// SetEnginesCount sets field value
func (o *TimeToUpdateSourcesRequest) SetEnginesCount(v int32) {
	o.EnginesCount = v
}

// GetEnginesList returns the EnginesList field value
func (o *TimeToUpdateSourcesRequest) GetEnginesList() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.EnginesList
}

// GetEnginesListOk returns a tuple with the EnginesList field value
// and a boolean to check if the value has been set.
func (o *TimeToUpdateSourcesRequest) GetEnginesListOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnginesList, true
}

// SetEnginesList sets field value
func (o *TimeToUpdateSourcesRequest) SetEnginesList(v []string) {
	o.EnginesList = v
}

// GetExistingNoOfSources returns the ExistingNoOfSources field value
func (o *TimeToUpdateSourcesRequest) GetExistingNoOfSources() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExistingNoOfSources
}

// GetExistingNoOfSourcesOk returns a tuple with the ExistingNoOfSources field value
// and a boolean to check if the value has been set.
func (o *TimeToUpdateSourcesRequest) GetExistingNoOfSourcesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExistingNoOfSources, true
}

// SetExistingNoOfSources sets field value
func (o *TimeToUpdateSourcesRequest) SetExistingNoOfSources(v int32) {
	o.ExistingNoOfSources = v
}

// GetSourcesCount returns the SourcesCount field value
func (o *TimeToUpdateSourcesRequest) GetSourcesCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SourcesCount
}

// GetSourcesCountOk returns a tuple with the SourcesCount field value
// and a boolean to check if the value has been set.
func (o *TimeToUpdateSourcesRequest) GetSourcesCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcesCount, true
}

// SetSourcesCount sets field value
func (o *TimeToUpdateSourcesRequest) SetSourcesCount(v int32) {
	o.SourcesCount = v
}

func (o TimeToUpdateSourcesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TimeToUpdateSourcesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["engines_count"] = o.EnginesCount
	toSerialize["engines_list"] = o.EnginesList
	toSerialize["existing_no_of_sources"] = o.ExistingNoOfSources
	toSerialize["sources_count"] = o.SourcesCount
	return toSerialize, nil
}

type NullableTimeToUpdateSourcesRequest struct {
	value *TimeToUpdateSourcesRequest
	isSet bool
}

func (v NullableTimeToUpdateSourcesRequest) Get() *TimeToUpdateSourcesRequest {
	return v.value
}

func (v *NullableTimeToUpdateSourcesRequest) Set(val *TimeToUpdateSourcesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeToUpdateSourcesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeToUpdateSourcesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeToUpdateSourcesRequest(val *TimeToUpdateSourcesRequest) *NullableTimeToUpdateSourcesRequest {
	return &NullableTimeToUpdateSourcesRequest{value: val, isSet: true}
}

func (v NullableTimeToUpdateSourcesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeToUpdateSourcesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


