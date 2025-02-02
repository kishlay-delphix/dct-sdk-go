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

// checks if the MaskingRuleset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MaskingRuleset{}

// MaskingRuleset A masking ruleset.
type MaskingRuleset struct {
	// The Ruleset type.
	Type NullableString `json:"type,omitempty"`
	// The name of this Ruleset.
	Name *string `json:"name,omitempty"`
	// Whether refresh drops tables. Only applicable to database ruleset type.
	RefreshDropsTables NullableBool `json:"refresh_drops_tables,omitempty"`
	// The list of algorithms for this Ruleset.
	Algorithms []string `json:"algorithms,omitempty"`
}

// NewMaskingRuleset instantiates a new MaskingRuleset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMaskingRuleset() *MaskingRuleset {
	this := MaskingRuleset{}
	return &this
}

// NewMaskingRulesetWithDefaults instantiates a new MaskingRuleset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMaskingRulesetWithDefaults() *MaskingRuleset {
	this := MaskingRuleset{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaskingRuleset) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaskingRuleset) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *MaskingRuleset) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *MaskingRuleset) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *MaskingRuleset) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *MaskingRuleset) UnsetType() {
	o.Type.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *MaskingRuleset) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingRuleset) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *MaskingRuleset) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *MaskingRuleset) SetName(v string) {
	o.Name = &v
}

// GetRefreshDropsTables returns the RefreshDropsTables field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MaskingRuleset) GetRefreshDropsTables() bool {
	if o == nil || IsNil(o.RefreshDropsTables.Get()) {
		var ret bool
		return ret
	}
	return *o.RefreshDropsTables.Get()
}

// GetRefreshDropsTablesOk returns a tuple with the RefreshDropsTables field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MaskingRuleset) GetRefreshDropsTablesOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.RefreshDropsTables.Get(), o.RefreshDropsTables.IsSet()
}

// HasRefreshDropsTables returns a boolean if a field has been set.
func (o *MaskingRuleset) HasRefreshDropsTables() bool {
	if o != nil && o.RefreshDropsTables.IsSet() {
		return true
	}

	return false
}

// SetRefreshDropsTables gets a reference to the given NullableBool and assigns it to the RefreshDropsTables field.
func (o *MaskingRuleset) SetRefreshDropsTables(v bool) {
	o.RefreshDropsTables.Set(&v)
}
// SetRefreshDropsTablesNil sets the value for RefreshDropsTables to be an explicit nil
func (o *MaskingRuleset) SetRefreshDropsTablesNil() {
	o.RefreshDropsTables.Set(nil)
}

// UnsetRefreshDropsTables ensures that no value is present for RefreshDropsTables, not even an explicit nil
func (o *MaskingRuleset) UnsetRefreshDropsTables() {
	o.RefreshDropsTables.Unset()
}

// GetAlgorithms returns the Algorithms field value if set, zero value otherwise.
func (o *MaskingRuleset) GetAlgorithms() []string {
	if o == nil || IsNil(o.Algorithms) {
		var ret []string
		return ret
	}
	return o.Algorithms
}

// GetAlgorithmsOk returns a tuple with the Algorithms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MaskingRuleset) GetAlgorithmsOk() ([]string, bool) {
	if o == nil || IsNil(o.Algorithms) {
		return nil, false
	}
	return o.Algorithms, true
}

// HasAlgorithms returns a boolean if a field has been set.
func (o *MaskingRuleset) HasAlgorithms() bool {
	if o != nil && !IsNil(o.Algorithms) {
		return true
	}

	return false
}

// SetAlgorithms gets a reference to the given []string and assigns it to the Algorithms field.
func (o *MaskingRuleset) SetAlgorithms(v []string) {
	o.Algorithms = v
}

func (o MaskingRuleset) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MaskingRuleset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if o.RefreshDropsTables.IsSet() {
		toSerialize["refresh_drops_tables"] = o.RefreshDropsTables.Get()
	}
	if !IsNil(o.Algorithms) {
		toSerialize["algorithms"] = o.Algorithms
	}
	return toSerialize, nil
}

type NullableMaskingRuleset struct {
	value *MaskingRuleset
	isSet bool
}

func (v NullableMaskingRuleset) Get() *MaskingRuleset {
	return v.value
}

func (v *NullableMaskingRuleset) Set(val *MaskingRuleset) {
	v.value = val
	v.isSet = true
}

func (v NullableMaskingRuleset) IsSet() bool {
	return v.isSet
}

func (v *NullableMaskingRuleset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMaskingRuleset(val *MaskingRuleset) *NullableMaskingRuleset {
	return &NullableMaskingRuleset{value: val, isSet: true}
}

func (v NullableMaskingRuleset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMaskingRuleset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


