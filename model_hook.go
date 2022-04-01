/*
Delphix API Gateway

Delphix API Gateway API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// Hook struct for Hook
type Hook struct {
	Name *string `json:"name,omitempty"`
	Command string `json:"command"`
	Shell *string `json:"shell,omitempty"`
}

// NewHook instantiates a new Hook object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHook(command string) *Hook {
	this := Hook{}
	this.Command = command
	var shell string = "bash"
	this.Shell = &shell
	return &this
}

// NewHookWithDefaults instantiates a new Hook object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHookWithDefaults() *Hook {
	this := Hook{}
	var shell string = "bash"
	this.Shell = &shell
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Hook) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hook) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Hook) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Hook) SetName(v string) {
	o.Name = &v
}

// GetCommand returns the Command field value
func (o *Hook) GetCommand() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *Hook) GetCommandOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Command, true
}

// SetCommand sets field value
func (o *Hook) SetCommand(v string) {
	o.Command = v
}

// GetShell returns the Shell field value if set, zero value otherwise.
func (o *Hook) GetShell() string {
	if o == nil || o.Shell == nil {
		var ret string
		return ret
	}
	return *o.Shell
}

// GetShellOk returns a tuple with the Shell field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Hook) GetShellOk() (*string, bool) {
	if o == nil || o.Shell == nil {
		return nil, false
	}
	return o.Shell, true
}

// HasShell returns a boolean if a field has been set.
func (o *Hook) HasShell() bool {
	if o != nil && o.Shell != nil {
		return true
	}

	return false
}

// SetShell gets a reference to the given string and assigns it to the Shell field.
func (o *Hook) SetShell(v string) {
	o.Shell = &v
}

func (o Hook) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["command"] = o.Command
	}
	if o.Shell != nil {
		toSerialize["shell"] = o.Shell
	}
	return json.Marshal(toSerialize)
}

type NullableHook struct {
	value *Hook
	isSet bool
}

func (v NullableHook) Get() *Hook {
	return v.value
}

func (v *NullableHook) Set(val *Hook) {
	v.value = val
	v.isSet = true
}

func (v NullableHook) IsSet() bool {
	return v.isSet
}

func (v *NullableHook) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHook(val *Hook) *NullableHook {
	return &NullableHook{value: val, isSet: true}
}

func (v NullableHook) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHook) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


