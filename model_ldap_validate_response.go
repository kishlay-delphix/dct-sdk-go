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

// checks if the LdapValidateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LdapValidateResponse{}

// LdapValidateResponse struct for LdapValidateResponse
type LdapValidateResponse struct {
	// Validation message for LDAP config.
	Message *string `json:"message,omitempty"`
}

// NewLdapValidateResponse instantiates a new LdapValidateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLdapValidateResponse() *LdapValidateResponse {
	this := LdapValidateResponse{}
	return &this
}

// NewLdapValidateResponseWithDefaults instantiates a new LdapValidateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLdapValidateResponseWithDefaults() *LdapValidateResponse {
	this := LdapValidateResponse{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *LdapValidateResponse) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LdapValidateResponse) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *LdapValidateResponse) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *LdapValidateResponse) SetMessage(v string) {
	o.Message = &v
}

func (o LdapValidateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LdapValidateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableLdapValidateResponse struct {
	value *LdapValidateResponse
	isSet bool
}

func (v NullableLdapValidateResponse) Get() *LdapValidateResponse {
	return v.value
}

func (v *NullableLdapValidateResponse) Set(val *LdapValidateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableLdapValidateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableLdapValidateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdapValidateResponse(val *LdapValidateResponse) *NullableLdapValidateResponse {
	return &NullableLdapValidateResponse{value: val, isSet: true}
}

func (v NullableLdapValidateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdapValidateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


