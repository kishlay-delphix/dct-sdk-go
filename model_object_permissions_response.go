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

// checks if the ObjectPermissionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ObjectPermissionsResponse{}

// ObjectPermissionsResponse The object permissions for a given object in DCT based on object type and object id.
type ObjectPermissionsResponse struct {
	// The Accounts permitted for this object.
	Accounts []ObjectPermissionAccount `json:"accounts,omitempty"`
}

// NewObjectPermissionsResponse instantiates a new ObjectPermissionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectPermissionsResponse() *ObjectPermissionsResponse {
	this := ObjectPermissionsResponse{}
	return &this
}

// NewObjectPermissionsResponseWithDefaults instantiates a new ObjectPermissionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectPermissionsResponseWithDefaults() *ObjectPermissionsResponse {
	this := ObjectPermissionsResponse{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *ObjectPermissionsResponse) GetAccounts() []ObjectPermissionAccount {
	if o == nil || IsNil(o.Accounts) {
		var ret []ObjectPermissionAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectPermissionsResponse) GetAccountsOk() ([]ObjectPermissionAccount, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *ObjectPermissionsResponse) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []ObjectPermissionAccount and assigns it to the Accounts field.
func (o *ObjectPermissionsResponse) SetAccounts(v []ObjectPermissionAccount) {
	o.Accounts = v
}

func (o ObjectPermissionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ObjectPermissionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	return toSerialize, nil
}

type NullableObjectPermissionsResponse struct {
	value *ObjectPermissionsResponse
	isSet bool
}

func (v NullableObjectPermissionsResponse) Get() *ObjectPermissionsResponse {
	return v.value
}

func (v *NullableObjectPermissionsResponse) Set(val *ObjectPermissionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectPermissionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectPermissionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectPermissionsResponse(val *ObjectPermissionsResponse) *NullableObjectPermissionsResponse {
	return &NullableObjectPermissionsResponse{value: val, isSet: true}
}

func (v NullableObjectPermissionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectPermissionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


