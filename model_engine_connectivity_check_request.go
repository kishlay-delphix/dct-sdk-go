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

// EngineConnectivityCheckRequest Parameters to check connectivity between engine and remote host.
type EngineConnectivityCheckRequest struct {
	EngineId string `json:"engine_id"`
	Host string `json:"host"`
	Port NullableInt32 `json:"port"`
}

// NewEngineConnectivityCheckRequest instantiates a new EngineConnectivityCheckRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngineConnectivityCheckRequest(engineId string, host string, port NullableInt32) *EngineConnectivityCheckRequest {
	this := EngineConnectivityCheckRequest{}
	this.EngineId = engineId
	this.Host = host
	this.Port = port
	return &this
}

// NewEngineConnectivityCheckRequestWithDefaults instantiates a new EngineConnectivityCheckRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineConnectivityCheckRequestWithDefaults() *EngineConnectivityCheckRequest {
	this := EngineConnectivityCheckRequest{}
	return &this
}

// GetEngineId returns the EngineId field value
func (o *EngineConnectivityCheckRequest) GetEngineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value
// and a boolean to check if the value has been set.
func (o *EngineConnectivityCheckRequest) GetEngineIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.EngineId, true
}

// SetEngineId sets field value
func (o *EngineConnectivityCheckRequest) SetEngineId(v string) {
	o.EngineId = v
}

// GetHost returns the Host field value
func (o *EngineConnectivityCheckRequest) GetHost() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Host
}

// GetHostOk returns a tuple with the Host field value
// and a boolean to check if the value has been set.
func (o *EngineConnectivityCheckRequest) GetHostOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Host, true
}

// SetHost sets field value
func (o *EngineConnectivityCheckRequest) SetHost(v string) {
	o.Host = v
}

// GetPort returns the Port field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *EngineConnectivityCheckRequest) GetPort() int32 {
	if o == nil || o.Port.Get() == nil {
		var ret int32
		return ret
	}

	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EngineConnectivityCheckRequest) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// SetPort sets field value
func (o *EngineConnectivityCheckRequest) SetPort(v int32) {
	o.Port.Set(&v)
}

func (o EngineConnectivityCheckRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["engine_id"] = o.EngineId
	}
	if true {
		toSerialize["host"] = o.Host
	}
	if true {
		toSerialize["port"] = o.Port.Get()
	}
	return json.Marshal(toSerialize)
}

type NullableEngineConnectivityCheckRequest struct {
	value *EngineConnectivityCheckRequest
	isSet bool
}

func (v NullableEngineConnectivityCheckRequest) Get() *EngineConnectivityCheckRequest {
	return v.value
}

func (v *NullableEngineConnectivityCheckRequest) Set(val *EngineConnectivityCheckRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEngineConnectivityCheckRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEngineConnectivityCheckRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngineConnectivityCheckRequest(val *EngineConnectivityCheckRequest) *NullableEngineConnectivityCheckRequest {
	return &NullableEngineConnectivityCheckRequest{value: val, isSet: true}
}

func (v NullableEngineConnectivityCheckRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngineConnectivityCheckRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


