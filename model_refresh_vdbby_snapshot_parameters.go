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

// RefreshVDBBySnapshotParameters struct for RefreshVDBBySnapshotParameters
type RefreshVDBBySnapshotParameters struct {
	// The ID of the snapshot from which to execute the operation. If the snapshot_id is not, selects the latest snapshot.
	SnapshotId *string `json:"snapshot_id,omitempty"`
}

// NewRefreshVDBBySnapshotParameters instantiates a new RefreshVDBBySnapshotParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRefreshVDBBySnapshotParameters() *RefreshVDBBySnapshotParameters {
	this := RefreshVDBBySnapshotParameters{}
	return &this
}

// NewRefreshVDBBySnapshotParametersWithDefaults instantiates a new RefreshVDBBySnapshotParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRefreshVDBBySnapshotParametersWithDefaults() *RefreshVDBBySnapshotParameters {
	this := RefreshVDBBySnapshotParameters{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *RefreshVDBBySnapshotParameters) GetSnapshotId() string {
	if o == nil || o.SnapshotId == nil {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RefreshVDBBySnapshotParameters) GetSnapshotIdOk() (*string, bool) {
	if o == nil || o.SnapshotId == nil {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *RefreshVDBBySnapshotParameters) HasSnapshotId() bool {
	if o != nil && o.SnapshotId != nil {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *RefreshVDBBySnapshotParameters) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

func (o RefreshVDBBySnapshotParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SnapshotId != nil {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	return json.Marshal(toSerialize)
}

type NullableRefreshVDBBySnapshotParameters struct {
	value *RefreshVDBBySnapshotParameters
	isSet bool
}

func (v NullableRefreshVDBBySnapshotParameters) Get() *RefreshVDBBySnapshotParameters {
	return v.value
}

func (v *NullableRefreshVDBBySnapshotParameters) Set(val *RefreshVDBBySnapshotParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRefreshVDBBySnapshotParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRefreshVDBBySnapshotParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRefreshVDBBySnapshotParameters(val *RefreshVDBBySnapshotParameters) *NullableRefreshVDBBySnapshotParameters {
	return &NullableRefreshVDBBySnapshotParameters{value: val, isSet: true}
}

func (v NullableRefreshVDBBySnapshotParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRefreshVDBBySnapshotParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


