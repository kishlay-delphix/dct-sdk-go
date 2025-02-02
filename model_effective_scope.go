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

// checks if the EffectiveScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EffectiveScope{}

// EffectiveScope struct for EffectiveScope
type EffectiveScope struct {
	// Id of the access group scope.
	Id *string `json:"id,omitempty"`
	// Name of the access group scope.
	Name *string `json:"name,omitempty"`
}

// NewEffectiveScope instantiates a new EffectiveScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEffectiveScope() *EffectiveScope {
	this := EffectiveScope{}
	return &this
}

// NewEffectiveScopeWithDefaults instantiates a new EffectiveScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEffectiveScopeWithDefaults() *EffectiveScope {
	this := EffectiveScope{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *EffectiveScope) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectiveScope) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *EffectiveScope) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *EffectiveScope) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EffectiveScope) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EffectiveScope) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EffectiveScope) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EffectiveScope) SetName(v string) {
	o.Name = &v
}

func (o EffectiveScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EffectiveScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	return toSerialize, nil
}

type NullableEffectiveScope struct {
	value *EffectiveScope
	isSet bool
}

func (v NullableEffectiveScope) Get() *EffectiveScope {
	return v.value
}

func (v *NullableEffectiveScope) Set(val *EffectiveScope) {
	v.value = val
	v.isSet = true
}

func (v NullableEffectiveScope) IsSet() bool {
	return v.isSet
}

func (v *NullableEffectiveScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEffectiveScope(val *EffectiveScope) *NullableEffectiveScope {
	return &NullableEffectiveScope{value: val, isSet: true}
}

func (v NullableEffectiveScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEffectiveScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


