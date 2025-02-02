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
	"time"
)

// checks if the ProductHistory type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductHistory{}

// ProductHistory struct for ProductHistory
type ProductHistory struct {
	// Product Version.
	Version *string `json:"version,omitempty"`
	// This version installed on date.
	InstalledOn *time.Time `json:"installed_on,omitempty"`
}

// NewProductHistory instantiates a new ProductHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductHistory() *ProductHistory {
	this := ProductHistory{}
	return &this
}

// NewProductHistoryWithDefaults instantiates a new ProductHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductHistoryWithDefaults() *ProductHistory {
	this := ProductHistory{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *ProductHistory) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductHistory) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *ProductHistory) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *ProductHistory) SetVersion(v string) {
	o.Version = &v
}

// GetInstalledOn returns the InstalledOn field value if set, zero value otherwise.
func (o *ProductHistory) GetInstalledOn() time.Time {
	if o == nil || IsNil(o.InstalledOn) {
		var ret time.Time
		return ret
	}
	return *o.InstalledOn
}

// GetInstalledOnOk returns a tuple with the InstalledOn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductHistory) GetInstalledOnOk() (*time.Time, bool) {
	if o == nil || IsNil(o.InstalledOn) {
		return nil, false
	}
	return o.InstalledOn, true
}

// HasInstalledOn returns a boolean if a field has been set.
func (o *ProductHistory) HasInstalledOn() bool {
	if o != nil && !IsNil(o.InstalledOn) {
		return true
	}

	return false
}

// SetInstalledOn gets a reference to the given time.Time and assigns it to the InstalledOn field.
func (o *ProductHistory) SetInstalledOn(v time.Time) {
	o.InstalledOn = &v
}

func (o ProductHistory) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductHistory) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.InstalledOn) {
		toSerialize["installed_on"] = o.InstalledOn
	}
	return toSerialize, nil
}

type NullableProductHistory struct {
	value *ProductHistory
	isSet bool
}

func (v NullableProductHistory) Get() *ProductHistory {
	return v.value
}

func (v *NullableProductHistory) Set(val *ProductHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableProductHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableProductHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductHistory(val *ProductHistory) *NullableProductHistory {
	return &NullableProductHistory{value: val, isSet: true}
}

func (v NullableProductHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


