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

// checks if the AccessGroupScopesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccessGroupScopesRequest{}

// AccessGroupScopesRequest struct for AccessGroupScopesRequest
type AccessGroupScopesRequest struct {
	Scopes []AccessGroupScope `json:"scopes"`
}

// NewAccessGroupScopesRequest instantiates a new AccessGroupScopesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessGroupScopesRequest(scopes []AccessGroupScope) *AccessGroupScopesRequest {
	this := AccessGroupScopesRequest{}
	this.Scopes = scopes
	return &this
}

// NewAccessGroupScopesRequestWithDefaults instantiates a new AccessGroupScopesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessGroupScopesRequestWithDefaults() *AccessGroupScopesRequest {
	this := AccessGroupScopesRequest{}
	return &this
}

// GetScopes returns the Scopes field value
func (o *AccessGroupScopesRequest) GetScopes() []AccessGroupScope {
	if o == nil {
		var ret []AccessGroupScope
		return ret
	}

	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value
// and a boolean to check if the value has been set.
func (o *AccessGroupScopesRequest) GetScopesOk() ([]AccessGroupScope, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scopes, true
}

// SetScopes sets field value
func (o *AccessGroupScopesRequest) SetScopes(v []AccessGroupScope) {
	o.Scopes = v
}

func (o AccessGroupScopesRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccessGroupScopesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["scopes"] = o.Scopes
	return toSerialize, nil
}

type NullableAccessGroupScopesRequest struct {
	value *AccessGroupScopesRequest
	isSet bool
}

func (v NullableAccessGroupScopesRequest) Get() *AccessGroupScopesRequest {
	return v.value
}

func (v *NullableAccessGroupScopesRequest) Set(val *AccessGroupScopesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessGroupScopesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessGroupScopesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessGroupScopesRequest(val *AccessGroupScopesRequest) *NullableAccessGroupScopesRequest {
	return &NullableAccessGroupScopesRequest{value: val, isSet: true}
}

func (v NullableAccessGroupScopesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessGroupScopesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


