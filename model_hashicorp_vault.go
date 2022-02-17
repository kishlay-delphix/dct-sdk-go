/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// HashicorpVault Configuration required to connect and authenticate with an Hashicorp Vault which stores engines username and passwords.
type HashicorpVault struct {
	Id *int64 `json:"id,omitempty"`
	// Environment variables to set when invoking the Vault CLI tool. The environment variables will be used both to login to the vault (if this step is required) and to retrieve engine username and passwords. 
	EnvVariables *map[string]string `json:"env_variables,omitempty"`
	// Arguments to the \"vault\" CLI tool to be used to fetch a client token (or \"login\"). If supporting files, such as TLS certificates, must be used to authenticate, they can be mounted to the /etc/config directory. This property must not be set when using the TOKEN authentication method as login is not required. 
	LoginCommandArgs []string `json:"login_command_args,omitempty"`
}

// NewHashicorpVault instantiates a new HashicorpVault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHashicorpVault() *HashicorpVault {
	this := HashicorpVault{}
	return &this
}

// NewHashicorpVaultWithDefaults instantiates a new HashicorpVault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHashicorpVaultWithDefaults() *HashicorpVault {
	this := HashicorpVault{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *HashicorpVault) GetId() int64 {
	if o == nil || o.Id == nil {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashicorpVault) GetIdOk() (*int64, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *HashicorpVault) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *HashicorpVault) SetId(v int64) {
	o.Id = &v
}

// GetEnvVariables returns the EnvVariables field value if set, zero value otherwise.
func (o *HashicorpVault) GetEnvVariables() map[string]string {
	if o == nil || o.EnvVariables == nil {
		var ret map[string]string
		return ret
	}
	return *o.EnvVariables
}

// GetEnvVariablesOk returns a tuple with the EnvVariables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashicorpVault) GetEnvVariablesOk() (*map[string]string, bool) {
	if o == nil || o.EnvVariables == nil {
		return nil, false
	}
	return o.EnvVariables, true
}

// HasEnvVariables returns a boolean if a field has been set.
func (o *HashicorpVault) HasEnvVariables() bool {
	if o != nil && o.EnvVariables != nil {
		return true
	}

	return false
}

// SetEnvVariables gets a reference to the given map[string]string and assigns it to the EnvVariables field.
func (o *HashicorpVault) SetEnvVariables(v map[string]string) {
	o.EnvVariables = &v
}

// GetLoginCommandArgs returns the LoginCommandArgs field value if set, zero value otherwise.
func (o *HashicorpVault) GetLoginCommandArgs() []string {
	if o == nil || o.LoginCommandArgs == nil {
		var ret []string
		return ret
	}
	return o.LoginCommandArgs
}

// GetLoginCommandArgsOk returns a tuple with the LoginCommandArgs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HashicorpVault) GetLoginCommandArgsOk() ([]string, bool) {
	if o == nil || o.LoginCommandArgs == nil {
		return nil, false
	}
	return o.LoginCommandArgs, true
}

// HasLoginCommandArgs returns a boolean if a field has been set.
func (o *HashicorpVault) HasLoginCommandArgs() bool {
	if o != nil && o.LoginCommandArgs != nil {
		return true
	}

	return false
}

// SetLoginCommandArgs gets a reference to the given []string and assigns it to the LoginCommandArgs field.
func (o *HashicorpVault) SetLoginCommandArgs(v []string) {
	o.LoginCommandArgs = v
}

func (o HashicorpVault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.EnvVariables != nil {
		toSerialize["env_variables"] = o.EnvVariables
	}
	if o.LoginCommandArgs != nil {
		toSerialize["login_command_args"] = o.LoginCommandArgs
	}
	return json.Marshal(toSerialize)
}

type NullableHashicorpVault struct {
	value *HashicorpVault
	isSet bool
}

func (v NullableHashicorpVault) Get() *HashicorpVault {
	return v.value
}

func (v *NullableHashicorpVault) Set(val *HashicorpVault) {
	v.value = val
	v.isSet = true
}

func (v NullableHashicorpVault) IsSet() bool {
	return v.isSet
}

func (v *NullableHashicorpVault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHashicorpVault(val *HashicorpVault) *NullableHashicorpVault {
	return &NullableHashicorpVault{value: val, isSet: true}
}

func (v NullableHashicorpVault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHashicorpVault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


