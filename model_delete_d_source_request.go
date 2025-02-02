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

// checks if the DeleteDSourceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteDSourceRequest{}

// DeleteDSourceRequest struct for DeleteDSourceRequest
type DeleteDSourceRequest struct {
	// Id of the dSource to delete.
	DsourceId string `json:"dsource_id"`
	// Flag indicating whether to continue the operation upon failures.
	Force *bool `json:"force,omitempty"`
	// The name of the privileged user to run the delete operation as (Oracle only).
	OracleUsername *string `json:"oracle_username,omitempty"`
	// Password for privileged user (Oracle only).
	OraclePassword *string `json:"oracle_password,omitempty"`
}

// NewDeleteDSourceRequest instantiates a new DeleteDSourceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteDSourceRequest(dsourceId string) *DeleteDSourceRequest {
	this := DeleteDSourceRequest{}
	this.DsourceId = dsourceId
	var force bool = false
	this.Force = &force
	return &this
}

// NewDeleteDSourceRequestWithDefaults instantiates a new DeleteDSourceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteDSourceRequestWithDefaults() *DeleteDSourceRequest {
	this := DeleteDSourceRequest{}
	var force bool = false
	this.Force = &force
	return &this
}

// GetDsourceId returns the DsourceId field value
func (o *DeleteDSourceRequest) GetDsourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DsourceId
}

// GetDsourceIdOk returns a tuple with the DsourceId field value
// and a boolean to check if the value has been set.
func (o *DeleteDSourceRequest) GetDsourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DsourceId, true
}

// SetDsourceId sets field value
func (o *DeleteDSourceRequest) SetDsourceId(v string) {
	o.DsourceId = v
}

// GetForce returns the Force field value if set, zero value otherwise.
func (o *DeleteDSourceRequest) GetForce() bool {
	if o == nil || IsNil(o.Force) {
		var ret bool
		return ret
	}
	return *o.Force
}

// GetForceOk returns a tuple with the Force field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDSourceRequest) GetForceOk() (*bool, bool) {
	if o == nil || IsNil(o.Force) {
		return nil, false
	}
	return o.Force, true
}

// HasForce returns a boolean if a field has been set.
func (o *DeleteDSourceRequest) HasForce() bool {
	if o != nil && !IsNil(o.Force) {
		return true
	}

	return false
}

// SetForce gets a reference to the given bool and assigns it to the Force field.
func (o *DeleteDSourceRequest) SetForce(v bool) {
	o.Force = &v
}

// GetOracleUsername returns the OracleUsername field value if set, zero value otherwise.
func (o *DeleteDSourceRequest) GetOracleUsername() string {
	if o == nil || IsNil(o.OracleUsername) {
		var ret string
		return ret
	}
	return *o.OracleUsername
}

// GetOracleUsernameOk returns a tuple with the OracleUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDSourceRequest) GetOracleUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.OracleUsername) {
		return nil, false
	}
	return o.OracleUsername, true
}

// HasOracleUsername returns a boolean if a field has been set.
func (o *DeleteDSourceRequest) HasOracleUsername() bool {
	if o != nil && !IsNil(o.OracleUsername) {
		return true
	}

	return false
}

// SetOracleUsername gets a reference to the given string and assigns it to the OracleUsername field.
func (o *DeleteDSourceRequest) SetOracleUsername(v string) {
	o.OracleUsername = &v
}

// GetOraclePassword returns the OraclePassword field value if set, zero value otherwise.
func (o *DeleteDSourceRequest) GetOraclePassword() string {
	if o == nil || IsNil(o.OraclePassword) {
		var ret string
		return ret
	}
	return *o.OraclePassword
}

// GetOraclePasswordOk returns a tuple with the OraclePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeleteDSourceRequest) GetOraclePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OraclePassword) {
		return nil, false
	}
	return o.OraclePassword, true
}

// HasOraclePassword returns a boolean if a field has been set.
func (o *DeleteDSourceRequest) HasOraclePassword() bool {
	if o != nil && !IsNil(o.OraclePassword) {
		return true
	}

	return false
}

// SetOraclePassword gets a reference to the given string and assigns it to the OraclePassword field.
func (o *DeleteDSourceRequest) SetOraclePassword(v string) {
	o.OraclePassword = &v
}

func (o DeleteDSourceRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteDSourceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dsource_id"] = o.DsourceId
	if !IsNil(o.Force) {
		toSerialize["force"] = o.Force
	}
	if !IsNil(o.OracleUsername) {
		toSerialize["oracle_username"] = o.OracleUsername
	}
	if !IsNil(o.OraclePassword) {
		toSerialize["oracle_password"] = o.OraclePassword
	}
	return toSerialize, nil
}

type NullableDeleteDSourceRequest struct {
	value *DeleteDSourceRequest
	isSet bool
}

func (v NullableDeleteDSourceRequest) Get() *DeleteDSourceRequest {
	return v.value
}

func (v *NullableDeleteDSourceRequest) Set(val *DeleteDSourceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteDSourceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteDSourceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteDSourceRequest(val *DeleteDSourceRequest) *NullableDeleteDSourceRequest {
	return &NullableDeleteDSourceRequest{value: val, isSet: true}
}

func (v NullableDeleteDSourceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteDSourceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


