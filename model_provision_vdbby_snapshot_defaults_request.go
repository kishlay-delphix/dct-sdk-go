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

// checks if the ProvisionVDBBySnapshotDefaultsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisionVDBBySnapshotDefaultsRequest{}

// ProvisionVDBBySnapshotDefaultsRequest struct for ProvisionVDBBySnapshotDefaultsRequest
type ProvisionVDBBySnapshotDefaultsRequest struct {
	// The ID of the snapshot from which to execute the operation.
	SnapshotId *string `json:"snapshot_id,omitempty"`
	// The ID of the Engine onto which to provision. If the source ID unambiguously identifies a source object, this parameter is unnecessary and ignored.
	EngineId *string `json:"engine_id,omitempty"`
	// The ID of the source object (dSource or VDB) to provision from. If this property is not set, the data_source of the snapshot_id will be used.
	SourceDataId *string `json:"source_data_id,omitempty"`
}

// NewProvisionVDBBySnapshotDefaultsRequest instantiates a new ProvisionVDBBySnapshotDefaultsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisionVDBBySnapshotDefaultsRequest() *ProvisionVDBBySnapshotDefaultsRequest {
	this := ProvisionVDBBySnapshotDefaultsRequest{}
	return &this
}

// NewProvisionVDBBySnapshotDefaultsRequestWithDefaults instantiates a new ProvisionVDBBySnapshotDefaultsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisionVDBBySnapshotDefaultsRequestWithDefaults() *ProvisionVDBBySnapshotDefaultsRequest {
	this := ProvisionVDBBySnapshotDefaultsRequest{}
	return &this
}

// GetSnapshotId returns the SnapshotId field value if set, zero value otherwise.
func (o *ProvisionVDBBySnapshotDefaultsRequest) GetSnapshotId() string {
	if o == nil || IsNil(o.SnapshotId) {
		var ret string
		return ret
	}
	return *o.SnapshotId
}

// GetSnapshotIdOk returns a tuple with the SnapshotId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBBySnapshotDefaultsRequest) GetSnapshotIdOk() (*string, bool) {
	if o == nil || IsNil(o.SnapshotId) {
		return nil, false
	}
	return o.SnapshotId, true
}

// HasSnapshotId returns a boolean if a field has been set.
func (o *ProvisionVDBBySnapshotDefaultsRequest) HasSnapshotId() bool {
	if o != nil && !IsNil(o.SnapshotId) {
		return true
	}

	return false
}

// SetSnapshotId gets a reference to the given string and assigns it to the SnapshotId field.
func (o *ProvisionVDBBySnapshotDefaultsRequest) SetSnapshotId(v string) {
	o.SnapshotId = &v
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *ProvisionVDBBySnapshotDefaultsRequest) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBBySnapshotDefaultsRequest) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *ProvisionVDBBySnapshotDefaultsRequest) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *ProvisionVDBBySnapshotDefaultsRequest) SetEngineId(v string) {
	o.EngineId = &v
}

// GetSourceDataId returns the SourceDataId field value if set, zero value otherwise.
func (o *ProvisionVDBBySnapshotDefaultsRequest) GetSourceDataId() string {
	if o == nil || IsNil(o.SourceDataId) {
		var ret string
		return ret
	}
	return *o.SourceDataId
}

// GetSourceDataIdOk returns a tuple with the SourceDataId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisionVDBBySnapshotDefaultsRequest) GetSourceDataIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceDataId) {
		return nil, false
	}
	return o.SourceDataId, true
}

// HasSourceDataId returns a boolean if a field has been set.
func (o *ProvisionVDBBySnapshotDefaultsRequest) HasSourceDataId() bool {
	if o != nil && !IsNil(o.SourceDataId) {
		return true
	}

	return false
}

// SetSourceDataId gets a reference to the given string and assigns it to the SourceDataId field.
func (o *ProvisionVDBBySnapshotDefaultsRequest) SetSourceDataId(v string) {
	o.SourceDataId = &v
}

func (o ProvisionVDBBySnapshotDefaultsRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisionVDBBySnapshotDefaultsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SnapshotId) {
		toSerialize["snapshot_id"] = o.SnapshotId
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if !IsNil(o.SourceDataId) {
		toSerialize["source_data_id"] = o.SourceDataId
	}
	return toSerialize, nil
}

type NullableProvisionVDBBySnapshotDefaultsRequest struct {
	value *ProvisionVDBBySnapshotDefaultsRequest
	isSet bool
}

func (v NullableProvisionVDBBySnapshotDefaultsRequest) Get() *ProvisionVDBBySnapshotDefaultsRequest {
	return v.value
}

func (v *NullableProvisionVDBBySnapshotDefaultsRequest) Set(val *ProvisionVDBBySnapshotDefaultsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisionVDBBySnapshotDefaultsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisionVDBBySnapshotDefaultsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisionVDBBySnapshotDefaultsRequest(val *ProvisionVDBBySnapshotDefaultsRequest) *NullableProvisionVDBBySnapshotDefaultsRequest {
	return &NullableProvisionVDBBySnapshotDefaultsRequest{value: val, isSet: true}
}

func (v NullableProvisionVDBBySnapshotDefaultsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisionVDBBySnapshotDefaultsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


