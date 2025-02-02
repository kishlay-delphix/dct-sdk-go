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

// checks if the RefreshVDBByTimestampParametersAllOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RefreshVDBByTimestampParametersAllOf{}

// RefreshVDBByTimestampParametersAllOf struct for RefreshVDBByTimestampParametersAllOf
type RefreshVDBByTimestampParametersAllOf struct {
	// ID of the dataset to refresh to, mutually exclusive with timeflow_id.
	DatasetId *string `json:"dataset_id,omitempty"`
	// ID of the timeflow to refresh to, mutually exclusive with dataset_id.
	TimeflowId *string `json:"timeflow_id,omitempty"`
}

// NewRefreshVDBByTimestampParametersAllOf instantiates a new RefreshVDBByTimestampParametersAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBByTimestampParametersAllOf() *RefreshVDBByTimestampParametersAllOf {
	this := RefreshVDBByTimestampParametersAllOf{}
	return &this
}

// NewRefreshVDBByTimestampParametersAllOfWithDefaults instantiates a new RefreshVDBByTimestampParametersAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBByTimestampParametersAllOfWithDefaults() *RefreshVDBByTimestampParametersAllOf {
	this := RefreshVDBByTimestampParametersAllOf{}
	return &this
}

// GetDatasetId returns the DatasetId field value if set, zero value otherwise.
func (o *RefreshVDBByTimestampParametersAllOf) GetDatasetId() string {
	if o == nil || IsNil(o.DatasetId) {
		var ret string
		return ret
	}
	return *o.DatasetId
}

// GetDatasetIdOk returns a tuple with the DatasetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBByTimestampParametersAllOf) GetDatasetIdOk() (*string, bool) {
	if o == nil || IsNil(o.DatasetId) {
		return nil, false
	}
	return o.DatasetId, true
}

// HasDatasetId returns a boolean if a field has been set.
func (o *RefreshVDBByTimestampParametersAllOf) HasDatasetId() bool {
	if o != nil && !IsNil(o.DatasetId) {
		return true
	}

	return false
}

// SetDatasetId gets a reference to the given string and assigns it to the DatasetId field.
func (o *RefreshVDBByTimestampParametersAllOf) SetDatasetId(v string) {
	o.DatasetId = &v
}

// GetTimeflowId returns the TimeflowId field value if set, zero value otherwise.
func (o *RefreshVDBByTimestampParametersAllOf) GetTimeflowId() string {
	if o == nil || IsNil(o.TimeflowId) {
		var ret string
		return ret
	}
	return *o.TimeflowId
}

// GetTimeflowIdOk returns a tuple with the TimeflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBByTimestampParametersAllOf) GetTimeflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.TimeflowId) {
		return nil, false
	}
	return o.TimeflowId, true
}

// HasTimeflowId returns a boolean if a field has been set.
func (o *RefreshVDBByTimestampParametersAllOf) HasTimeflowId() bool {
	if o != nil && !IsNil(o.TimeflowId) {
		return true
	}

	return false
}

// SetTimeflowId gets a reference to the given string and assigns it to the TimeflowId field.
func (o *RefreshVDBByTimestampParametersAllOf) SetTimeflowId(v string) {
	o.TimeflowId = &v
}

func (o RefreshVDBByTimestampParametersAllOf) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RefreshVDBByTimestampParametersAllOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DatasetId) {
		toSerialize["dataset_id"] = o.DatasetId
	}
	if !IsNil(o.TimeflowId) {
		toSerialize["timeflow_id"] = o.TimeflowId
	}
	return toSerialize, nil
}

type NullableRefreshVDBByTimestampParametersAllOf struct {
	value *RefreshVDBByTimestampParametersAllOf
	isSet bool
}

func (v NullableRefreshVDBByTimestampParametersAllOf) Get() *RefreshVDBByTimestampParametersAllOf {
	return v.value
}

func (v *NullableRefreshVDBByTimestampParametersAllOf) Set(val *RefreshVDBByTimestampParametersAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBByTimestampParametersAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBByTimestampParametersAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBByTimestampParametersAllOf(val *RefreshVDBByTimestampParametersAllOf) *NullableRefreshVDBByTimestampParametersAllOf {
	return &NullableRefreshVDBByTimestampParametersAllOf{value: val, isSet: true}
}

func (v NullableRefreshVDBByTimestampParametersAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBByTimestampParametersAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


