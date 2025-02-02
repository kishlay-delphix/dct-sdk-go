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
	"time"
)

// checks if the VDBInventoryData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VDBInventoryData{}

// VDBInventoryData struct for VDBInventoryData
type VDBInventoryData struct {
	EngineName *string `json:"engine_name,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
	ParentName *string `json:"parent_name,omitempty"`
	ParentId *string `json:"parent_id,omitempty"`
	CreationDate *time.Time `json:"creation_date,omitempty"`
	ParentTimeflowLocation *string `json:"parent_timeflow_location,omitempty"`
	ParentTimeflowTimestamp *time.Time `json:"parent_timeflow_timestamp,omitempty"`
	ParentTimeflowTimezone *string `json:"parent_timeflow_timezone,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	Status *string `json:"status,omitempty"`
}

// NewVDBInventoryData instantiates a new VDBInventoryData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVDBInventoryData() *VDBInventoryData {
	this := VDBInventoryData{}
	return &this
}

// NewVDBInventoryDataWithDefaults instantiates a new VDBInventoryData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVDBInventoryDataWithDefaults() *VDBInventoryData {
	this := VDBInventoryData{}
	return &this
}

// GetEngineName returns the EngineName field value if set, zero value otherwise.
func (o *VDBInventoryData) GetEngineName() string {
	if o == nil || IsNil(o.EngineName) {
		var ret string
		return ret
	}
	return *o.EngineName
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetEngineNameOk() (*string, bool) {
	if o == nil || IsNil(o.EngineName) {
		return nil, false
	}
	return o.EngineName, true
}

// HasEngineName returns a boolean if a field has been set.
func (o *VDBInventoryData) HasEngineName() bool {
	if o != nil && !IsNil(o.EngineName) {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given string and assigns it to the EngineName field.
func (o *VDBInventoryData) SetEngineName(v string) {
	o.EngineName = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *VDBInventoryData) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *VDBInventoryData) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *VDBInventoryData) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *VDBInventoryData) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *VDBInventoryData) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *VDBInventoryData) SetType(v string) {
	o.Type = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VDBInventoryData) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VDBInventoryData) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VDBInventoryData) SetVersion(v string) {
	o.Version = &v
}

// GetParentName returns the ParentName field value if set, zero value otherwise.
func (o *VDBInventoryData) GetParentName() string {
	if o == nil || IsNil(o.ParentName) {
		var ret string
		return ret
	}
	return *o.ParentName
}

// GetParentNameOk returns a tuple with the ParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetParentNameOk() (*string, bool) {
	if o == nil || IsNil(o.ParentName) {
		return nil, false
	}
	return o.ParentName, true
}

// HasParentName returns a boolean if a field has been set.
func (o *VDBInventoryData) HasParentName() bool {
	if o != nil && !IsNil(o.ParentName) {
		return true
	}

	return false
}

// SetParentName gets a reference to the given string and assigns it to the ParentName field.
func (o *VDBInventoryData) SetParentName(v string) {
	o.ParentName = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *VDBInventoryData) GetParentId() string {
	if o == nil || IsNil(o.ParentId) {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.ParentId) {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *VDBInventoryData) HasParentId() bool {
	if o != nil && !IsNil(o.ParentId) {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *VDBInventoryData) SetParentId(v string) {
	o.ParentId = &v
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise.
func (o *VDBInventoryData) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetCreationDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationDate) {
		return nil, false
	}
	return o.CreationDate, true
}

// HasCreationDate returns a boolean if a field has been set.
func (o *VDBInventoryData) HasCreationDate() bool {
	if o != nil && !IsNil(o.CreationDate) {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given time.Time and assigns it to the CreationDate field.
func (o *VDBInventoryData) SetCreationDate(v time.Time) {
	o.CreationDate = &v
}

// GetParentTimeflowLocation returns the ParentTimeflowLocation field value if set, zero value otherwise.
func (o *VDBInventoryData) GetParentTimeflowLocation() string {
	if o == nil || IsNil(o.ParentTimeflowLocation) {
		var ret string
		return ret
	}
	return *o.ParentTimeflowLocation
}

// GetParentTimeflowLocationOk returns a tuple with the ParentTimeflowLocation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetParentTimeflowLocationOk() (*string, bool) {
	if o == nil || IsNil(o.ParentTimeflowLocation) {
		return nil, false
	}
	return o.ParentTimeflowLocation, true
}

// HasParentTimeflowLocation returns a boolean if a field has been set.
func (o *VDBInventoryData) HasParentTimeflowLocation() bool {
	if o != nil && !IsNil(o.ParentTimeflowLocation) {
		return true
	}

	return false
}

// SetParentTimeflowLocation gets a reference to the given string and assigns it to the ParentTimeflowLocation field.
func (o *VDBInventoryData) SetParentTimeflowLocation(v string) {
	o.ParentTimeflowLocation = &v
}

// GetParentTimeflowTimestamp returns the ParentTimeflowTimestamp field value if set, zero value otherwise.
func (o *VDBInventoryData) GetParentTimeflowTimestamp() time.Time {
	if o == nil || IsNil(o.ParentTimeflowTimestamp) {
		var ret time.Time
		return ret
	}
	return *o.ParentTimeflowTimestamp
}

// GetParentTimeflowTimestampOk returns a tuple with the ParentTimeflowTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetParentTimeflowTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ParentTimeflowTimestamp) {
		return nil, false
	}
	return o.ParentTimeflowTimestamp, true
}

// HasParentTimeflowTimestamp returns a boolean if a field has been set.
func (o *VDBInventoryData) HasParentTimeflowTimestamp() bool {
	if o != nil && !IsNil(o.ParentTimeflowTimestamp) {
		return true
	}

	return false
}

// SetParentTimeflowTimestamp gets a reference to the given time.Time and assigns it to the ParentTimeflowTimestamp field.
func (o *VDBInventoryData) SetParentTimeflowTimestamp(v time.Time) {
	o.ParentTimeflowTimestamp = &v
}

// GetParentTimeflowTimezone returns the ParentTimeflowTimezone field value if set, zero value otherwise.
func (o *VDBInventoryData) GetParentTimeflowTimezone() string {
	if o == nil || IsNil(o.ParentTimeflowTimezone) {
		var ret string
		return ret
	}
	return *o.ParentTimeflowTimezone
}

// GetParentTimeflowTimezoneOk returns a tuple with the ParentTimeflowTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetParentTimeflowTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.ParentTimeflowTimezone) {
		return nil, false
	}
	return o.ParentTimeflowTimezone, true
}

// HasParentTimeflowTimezone returns a boolean if a field has been set.
func (o *VDBInventoryData) HasParentTimeflowTimezone() bool {
	if o != nil && !IsNil(o.ParentTimeflowTimezone) {
		return true
	}

	return false
}

// SetParentTimeflowTimezone gets a reference to the given string and assigns it to the ParentTimeflowTimezone field.
func (o *VDBInventoryData) SetParentTimeflowTimezone(v string) {
	o.ParentTimeflowTimezone = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *VDBInventoryData) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *VDBInventoryData) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *VDBInventoryData) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VDBInventoryData) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VDBInventoryData) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VDBInventoryData) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *VDBInventoryData) SetStatus(v string) {
	o.Status = &v
}

func (o VDBInventoryData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VDBInventoryData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EngineName) {
		toSerialize["engine_name"] = o.EngineName
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.ParentName) {
		toSerialize["parent_name"] = o.ParentName
	}
	if !IsNil(o.ParentId) {
		toSerialize["parent_id"] = o.ParentId
	}
	if !IsNil(o.CreationDate) {
		toSerialize["creation_date"] = o.CreationDate
	}
	if !IsNil(o.ParentTimeflowLocation) {
		toSerialize["parent_timeflow_location"] = o.ParentTimeflowLocation
	}
	if !IsNil(o.ParentTimeflowTimestamp) {
		toSerialize["parent_timeflow_timestamp"] = o.ParentTimeflowTimestamp
	}
	if !IsNil(o.ParentTimeflowTimezone) {
		toSerialize["parent_timeflow_timezone"] = o.ParentTimeflowTimezone
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableVDBInventoryData struct {
	value *VDBInventoryData
	isSet bool
}

func (v NullableVDBInventoryData) Get() *VDBInventoryData {
	return v.value
}

func (v *NullableVDBInventoryData) Set(val *VDBInventoryData) {
	v.value = val
	v.isSet = true
}

func (v NullableVDBInventoryData) IsSet() bool {
	return v.isSet
}

func (v *NullableVDBInventoryData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVDBInventoryData(val *VDBInventoryData) *NullableVDBInventoryData {
	return &NullableVDBInventoryData{value: val, isSet: true}
}

func (v NullableVDBInventoryData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVDBInventoryData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


