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

// RollbackVDBGroupParameters Parameters to rollback a VDB Group.
type RollbackVDBGroupParameters struct {
	// ID of a bookmark to rollback this VDB Group to.
	BookmarkId *string `json:"bookmark_id,omitempty"`
}

// NewRollbackVDBGroupParameters instantiates a new RollbackVDBGroupParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackVDBGroupParameters() *RollbackVDBGroupParameters {
	this := RollbackVDBGroupParameters{}
	return &this
}

// NewRollbackVDBGroupParametersWithDefaults instantiates a new RollbackVDBGroupParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackVDBGroupParametersWithDefaults() *RollbackVDBGroupParameters {
	this := RollbackVDBGroupParameters{}
	return &this
}

// GetBookmarkId returns the BookmarkId field value if set, zero value otherwise.
func (o *RollbackVDBGroupParameters) GetBookmarkId() string {
	if o == nil || o.BookmarkId == nil {
		var ret string
		return ret
	}
	return *o.BookmarkId
}

// GetBookmarkIdOk returns a tuple with the BookmarkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackVDBGroupParameters) GetBookmarkIdOk() (*string, bool) {
	if o == nil || o.BookmarkId == nil {
		return nil, false
	}
	return o.BookmarkId, true
}

// HasBookmarkId returns a boolean if a field has been set.
func (o *RollbackVDBGroupParameters) HasBookmarkId() bool {
	if o != nil && o.BookmarkId != nil {
		return true
	}

	return false
}

// SetBookmarkId gets a reference to the given string and assigns it to the BookmarkId field.
func (o *RollbackVDBGroupParameters) SetBookmarkId(v string) {
	o.BookmarkId = &v
}

func (o RollbackVDBGroupParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BookmarkId != nil {
		toSerialize["bookmark_id"] = o.BookmarkId
	}
	return json.Marshal(toSerialize)
}

type NullableRollbackVDBGroupParameters struct {
	value *RollbackVDBGroupParameters
	isSet bool
}

func (v NullableRollbackVDBGroupParameters) Get() *RollbackVDBGroupParameters {
	return v.value
}

func (v *NullableRollbackVDBGroupParameters) Set(val *RollbackVDBGroupParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackVDBGroupParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackVDBGroupParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackVDBGroupParameters(val *RollbackVDBGroupParameters) *NullableRollbackVDBGroupParameters {
	return &NullableRollbackVDBGroupParameters{value: val, isSet: true}
}

func (v NullableRollbackVDBGroupParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackVDBGroupParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


