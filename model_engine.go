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

// checks if the Engine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Engine{}

// Engine A Delphix Virtualization or Masking Engine.
type Engine struct {
	// The Engine object entity ID.
	Id *string `json:"id,omitempty"`
	// The unique identifier generated by this engine.
	Uuid NullableString `json:"uuid,omitempty"`
	// The type of this engine.
	Type NullableString `json:"type,omitempty"`
	// The engine version.
	Version NullableString `json:"version,omitempty"`
	// The name of this engine.
	Name *string `json:"name,omitempty"`
	// The hostname of this engine.
	Hostname *string `json:"hostname,omitempty"`
	// The registration status of this engine.
	RegistrationStatus NullableString `json:"registration_status,omitempty"`
	// The connection status of this engine.
	ConnectionStatus NullableString `json:"connection_status,omitempty"`
	// The last time a connection was established with this engine.
	LastConnectionTime NullableTime `json:"last_connection_time,omitempty"`
	// The total amount of storage allocated to the engine's boot partition, in bytes.
	BootStorageCapacity NullableInt64 `json:"boot_storage_capacity,omitempty"`
	// The total number of CPU cores on this engine.
	CpuCoreCount NullableInt32 `json:"cpu_core_count,omitempty"`
	// The model of the processors on this engine.
	CpuType NullableString `json:"cpu_type,omitempty"`
	// The total amount of memory on this engine, in bytes.
	MemorySize NullableInt64 `json:"memory_size,omitempty"`
	// The total amount of storage allocated for engine objects and system metadata, in bytes.
	DataStorageCapacity NullableInt64 `json:"data_storage_capacity,omitempty"`
	// The amount of storage used by engine objects and system metadata, in bytes.
	DataStorageUsed NullableInt64 `json:"data_storage_used,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
}

// NewEngine instantiates a new Engine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEngine() *Engine {
	this := Engine{}
	return &this
}

// NewEngineWithDefaults instantiates a new Engine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEngineWithDefaults() *Engine {
	this := Engine{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Engine) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Engine) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Engine) SetId(v string) {
	o.Id = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetUuid() string {
	if o == nil || IsNil(o.Uuid.Get()) {
		var ret string
		return ret
	}
	return *o.Uuid.Get()
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Uuid.Get(), o.Uuid.IsSet()
}

// HasUuid returns a boolean if a field has been set.
func (o *Engine) HasUuid() bool {
	if o != nil && o.Uuid.IsSet() {
		return true
	}

	return false
}

// SetUuid gets a reference to the given NullableString and assigns it to the Uuid field.
func (o *Engine) SetUuid(v string) {
	o.Uuid.Set(&v)
}
// SetUuidNil sets the value for Uuid to be an explicit nil
func (o *Engine) SetUuidNil() {
	o.Uuid.Set(nil)
}

// UnsetUuid ensures that no value is present for Uuid, not even an explicit nil
func (o *Engine) UnsetUuid() {
	o.Uuid.Unset()
}

// GetType returns the Type field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetType() string {
	if o == nil || IsNil(o.Type.Get()) {
		var ret string
		return ret
	}
	return *o.Type.Get()
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type.Get(), o.Type.IsSet()
}

// HasType returns a boolean if a field has been set.
func (o *Engine) HasType() bool {
	if o != nil && o.Type.IsSet() {
		return true
	}

	return false
}

// SetType gets a reference to the given NullableString and assigns it to the Type field.
func (o *Engine) SetType(v string) {
	o.Type.Set(&v)
}
// SetTypeNil sets the value for Type to be an explicit nil
func (o *Engine) SetTypeNil() {
	o.Type.Set(nil)
}

// UnsetType ensures that no value is present for Type, not even an explicit nil
func (o *Engine) UnsetType() {
	o.Type.Unset()
}

// GetVersion returns the Version field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetVersion() string {
	if o == nil || IsNil(o.Version.Get()) {
		var ret string
		return ret
	}
	return *o.Version.Get()
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Version.Get(), o.Version.IsSet()
}

// HasVersion returns a boolean if a field has been set.
func (o *Engine) HasVersion() bool {
	if o != nil && o.Version.IsSet() {
		return true
	}

	return false
}

// SetVersion gets a reference to the given NullableString and assigns it to the Version field.
func (o *Engine) SetVersion(v string) {
	o.Version.Set(&v)
}
// SetVersionNil sets the value for Version to be an explicit nil
func (o *Engine) SetVersionNil() {
	o.Version.Set(nil)
}

// UnsetVersion ensures that no value is present for Version, not even an explicit nil
func (o *Engine) UnsetVersion() {
	o.Version.Unset()
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Engine) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Engine) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Engine) SetName(v string) {
	o.Name = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *Engine) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *Engine) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *Engine) SetHostname(v string) {
	o.Hostname = &v
}

// GetRegistrationStatus returns the RegistrationStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetRegistrationStatus() string {
	if o == nil || IsNil(o.RegistrationStatus.Get()) {
		var ret string
		return ret
	}
	return *o.RegistrationStatus.Get()
}

// GetRegistrationStatusOk returns a tuple with the RegistrationStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetRegistrationStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RegistrationStatus.Get(), o.RegistrationStatus.IsSet()
}

// HasRegistrationStatus returns a boolean if a field has been set.
func (o *Engine) HasRegistrationStatus() bool {
	if o != nil && o.RegistrationStatus.IsSet() {
		return true
	}

	return false
}

// SetRegistrationStatus gets a reference to the given NullableString and assigns it to the RegistrationStatus field.
func (o *Engine) SetRegistrationStatus(v string) {
	o.RegistrationStatus.Set(&v)
}
// SetRegistrationStatusNil sets the value for RegistrationStatus to be an explicit nil
func (o *Engine) SetRegistrationStatusNil() {
	o.RegistrationStatus.Set(nil)
}

// UnsetRegistrationStatus ensures that no value is present for RegistrationStatus, not even an explicit nil
func (o *Engine) UnsetRegistrationStatus() {
	o.RegistrationStatus.Unset()
}

// GetConnectionStatus returns the ConnectionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetConnectionStatus() string {
	if o == nil || IsNil(o.ConnectionStatus.Get()) {
		var ret string
		return ret
	}
	return *o.ConnectionStatus.Get()
}

// GetConnectionStatusOk returns a tuple with the ConnectionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetConnectionStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConnectionStatus.Get(), o.ConnectionStatus.IsSet()
}

// HasConnectionStatus returns a boolean if a field has been set.
func (o *Engine) HasConnectionStatus() bool {
	if o != nil && o.ConnectionStatus.IsSet() {
		return true
	}

	return false
}

// SetConnectionStatus gets a reference to the given NullableString and assigns it to the ConnectionStatus field.
func (o *Engine) SetConnectionStatus(v string) {
	o.ConnectionStatus.Set(&v)
}
// SetConnectionStatusNil sets the value for ConnectionStatus to be an explicit nil
func (o *Engine) SetConnectionStatusNil() {
	o.ConnectionStatus.Set(nil)
}

// UnsetConnectionStatus ensures that no value is present for ConnectionStatus, not even an explicit nil
func (o *Engine) UnsetConnectionStatus() {
	o.ConnectionStatus.Unset()
}

// GetLastConnectionTime returns the LastConnectionTime field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetLastConnectionTime() time.Time {
	if o == nil || IsNil(o.LastConnectionTime.Get()) {
		var ret time.Time
		return ret
	}
	return *o.LastConnectionTime.Get()
}

// GetLastConnectionTimeOk returns a tuple with the LastConnectionTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetLastConnectionTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastConnectionTime.Get(), o.LastConnectionTime.IsSet()
}

// HasLastConnectionTime returns a boolean if a field has been set.
func (o *Engine) HasLastConnectionTime() bool {
	if o != nil && o.LastConnectionTime.IsSet() {
		return true
	}

	return false
}

// SetLastConnectionTime gets a reference to the given NullableTime and assigns it to the LastConnectionTime field.
func (o *Engine) SetLastConnectionTime(v time.Time) {
	o.LastConnectionTime.Set(&v)
}
// SetLastConnectionTimeNil sets the value for LastConnectionTime to be an explicit nil
func (o *Engine) SetLastConnectionTimeNil() {
	o.LastConnectionTime.Set(nil)
}

// UnsetLastConnectionTime ensures that no value is present for LastConnectionTime, not even an explicit nil
func (o *Engine) UnsetLastConnectionTime() {
	o.LastConnectionTime.Unset()
}

// GetBootStorageCapacity returns the BootStorageCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetBootStorageCapacity() int64 {
	if o == nil || IsNil(o.BootStorageCapacity.Get()) {
		var ret int64
		return ret
	}
	return *o.BootStorageCapacity.Get()
}

// GetBootStorageCapacityOk returns a tuple with the BootStorageCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetBootStorageCapacityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.BootStorageCapacity.Get(), o.BootStorageCapacity.IsSet()
}

// HasBootStorageCapacity returns a boolean if a field has been set.
func (o *Engine) HasBootStorageCapacity() bool {
	if o != nil && o.BootStorageCapacity.IsSet() {
		return true
	}

	return false
}

// SetBootStorageCapacity gets a reference to the given NullableInt64 and assigns it to the BootStorageCapacity field.
func (o *Engine) SetBootStorageCapacity(v int64) {
	o.BootStorageCapacity.Set(&v)
}
// SetBootStorageCapacityNil sets the value for BootStorageCapacity to be an explicit nil
func (o *Engine) SetBootStorageCapacityNil() {
	o.BootStorageCapacity.Set(nil)
}

// UnsetBootStorageCapacity ensures that no value is present for BootStorageCapacity, not even an explicit nil
func (o *Engine) UnsetBootStorageCapacity() {
	o.BootStorageCapacity.Unset()
}

// GetCpuCoreCount returns the CpuCoreCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetCpuCoreCount() int32 {
	if o == nil || IsNil(o.CpuCoreCount.Get()) {
		var ret int32
		return ret
	}
	return *o.CpuCoreCount.Get()
}

// GetCpuCoreCountOk returns a tuple with the CpuCoreCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetCpuCoreCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuCoreCount.Get(), o.CpuCoreCount.IsSet()
}

// HasCpuCoreCount returns a boolean if a field has been set.
func (o *Engine) HasCpuCoreCount() bool {
	if o != nil && o.CpuCoreCount.IsSet() {
		return true
	}

	return false
}

// SetCpuCoreCount gets a reference to the given NullableInt32 and assigns it to the CpuCoreCount field.
func (o *Engine) SetCpuCoreCount(v int32) {
	o.CpuCoreCount.Set(&v)
}
// SetCpuCoreCountNil sets the value for CpuCoreCount to be an explicit nil
func (o *Engine) SetCpuCoreCountNil() {
	o.CpuCoreCount.Set(nil)
}

// UnsetCpuCoreCount ensures that no value is present for CpuCoreCount, not even an explicit nil
func (o *Engine) UnsetCpuCoreCount() {
	o.CpuCoreCount.Unset()
}

// GetCpuType returns the CpuType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetCpuType() string {
	if o == nil || IsNil(o.CpuType.Get()) {
		var ret string
		return ret
	}
	return *o.CpuType.Get()
}

// GetCpuTypeOk returns a tuple with the CpuType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetCpuTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CpuType.Get(), o.CpuType.IsSet()
}

// HasCpuType returns a boolean if a field has been set.
func (o *Engine) HasCpuType() bool {
	if o != nil && o.CpuType.IsSet() {
		return true
	}

	return false
}

// SetCpuType gets a reference to the given NullableString and assigns it to the CpuType field.
func (o *Engine) SetCpuType(v string) {
	o.CpuType.Set(&v)
}
// SetCpuTypeNil sets the value for CpuType to be an explicit nil
func (o *Engine) SetCpuTypeNil() {
	o.CpuType.Set(nil)
}

// UnsetCpuType ensures that no value is present for CpuType, not even an explicit nil
func (o *Engine) UnsetCpuType() {
	o.CpuType.Unset()
}

// GetMemorySize returns the MemorySize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetMemorySize() int64 {
	if o == nil || IsNil(o.MemorySize.Get()) {
		var ret int64
		return ret
	}
	return *o.MemorySize.Get()
}

// GetMemorySizeOk returns a tuple with the MemorySize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetMemorySizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.MemorySize.Get(), o.MemorySize.IsSet()
}

// HasMemorySize returns a boolean if a field has been set.
func (o *Engine) HasMemorySize() bool {
	if o != nil && o.MemorySize.IsSet() {
		return true
	}

	return false
}

// SetMemorySize gets a reference to the given NullableInt64 and assigns it to the MemorySize field.
func (o *Engine) SetMemorySize(v int64) {
	o.MemorySize.Set(&v)
}
// SetMemorySizeNil sets the value for MemorySize to be an explicit nil
func (o *Engine) SetMemorySizeNil() {
	o.MemorySize.Set(nil)
}

// UnsetMemorySize ensures that no value is present for MemorySize, not even an explicit nil
func (o *Engine) UnsetMemorySize() {
	o.MemorySize.Unset()
}

// GetDataStorageCapacity returns the DataStorageCapacity field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetDataStorageCapacity() int64 {
	if o == nil || IsNil(o.DataStorageCapacity.Get()) {
		var ret int64
		return ret
	}
	return *o.DataStorageCapacity.Get()
}

// GetDataStorageCapacityOk returns a tuple with the DataStorageCapacity field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetDataStorageCapacityOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataStorageCapacity.Get(), o.DataStorageCapacity.IsSet()
}

// HasDataStorageCapacity returns a boolean if a field has been set.
func (o *Engine) HasDataStorageCapacity() bool {
	if o != nil && o.DataStorageCapacity.IsSet() {
		return true
	}

	return false
}

// SetDataStorageCapacity gets a reference to the given NullableInt64 and assigns it to the DataStorageCapacity field.
func (o *Engine) SetDataStorageCapacity(v int64) {
	o.DataStorageCapacity.Set(&v)
}
// SetDataStorageCapacityNil sets the value for DataStorageCapacity to be an explicit nil
func (o *Engine) SetDataStorageCapacityNil() {
	o.DataStorageCapacity.Set(nil)
}

// UnsetDataStorageCapacity ensures that no value is present for DataStorageCapacity, not even an explicit nil
func (o *Engine) UnsetDataStorageCapacity() {
	o.DataStorageCapacity.Unset()
}

// GetDataStorageUsed returns the DataStorageUsed field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Engine) GetDataStorageUsed() int64 {
	if o == nil || IsNil(o.DataStorageUsed.Get()) {
		var ret int64
		return ret
	}
	return *o.DataStorageUsed.Get()
}

// GetDataStorageUsedOk returns a tuple with the DataStorageUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Engine) GetDataStorageUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataStorageUsed.Get(), o.DataStorageUsed.IsSet()
}

// HasDataStorageUsed returns a boolean if a field has been set.
func (o *Engine) HasDataStorageUsed() bool {
	if o != nil && o.DataStorageUsed.IsSet() {
		return true
	}

	return false
}

// SetDataStorageUsed gets a reference to the given NullableInt64 and assigns it to the DataStorageUsed field.
func (o *Engine) SetDataStorageUsed(v int64) {
	o.DataStorageUsed.Set(&v)
}
// SetDataStorageUsedNil sets the value for DataStorageUsed to be an explicit nil
func (o *Engine) SetDataStorageUsedNil() {
	o.DataStorageUsed.Set(nil)
}

// UnsetDataStorageUsed ensures that no value is present for DataStorageUsed, not even an explicit nil
func (o *Engine) UnsetDataStorageUsed() {
	o.DataStorageUsed.Unset()
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *Engine) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Engine) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *Engine) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *Engine) SetTags(v []Tag) {
	o.Tags = v
}

func (o Engine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Engine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.Uuid.IsSet() {
		toSerialize["uuid"] = o.Uuid.Get()
	}
	if o.Type.IsSet() {
		toSerialize["type"] = o.Type.Get()
	}
	if o.Version.IsSet() {
		toSerialize["version"] = o.Version.Get()
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if o.RegistrationStatus.IsSet() {
		toSerialize["registration_status"] = o.RegistrationStatus.Get()
	}
	if o.ConnectionStatus.IsSet() {
		toSerialize["connection_status"] = o.ConnectionStatus.Get()
	}
	if o.LastConnectionTime.IsSet() {
		toSerialize["last_connection_time"] = o.LastConnectionTime.Get()
	}
	if o.BootStorageCapacity.IsSet() {
		toSerialize["boot_storage_capacity"] = o.BootStorageCapacity.Get()
	}
	if o.CpuCoreCount.IsSet() {
		toSerialize["cpu_core_count"] = o.CpuCoreCount.Get()
	}
	if o.CpuType.IsSet() {
		toSerialize["cpu_type"] = o.CpuType.Get()
	}
	if o.MemorySize.IsSet() {
		toSerialize["memory_size"] = o.MemorySize.Get()
	}
	if o.DataStorageCapacity.IsSet() {
		toSerialize["data_storage_capacity"] = o.DataStorageCapacity.Get()
	}
	if o.DataStorageUsed.IsSet() {
		toSerialize["data_storage_used"] = o.DataStorageUsed.Get()
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableEngine struct {
	value *Engine
	isSet bool
}

func (v NullableEngine) Get() *Engine {
	return v.value
}

func (v *NullableEngine) Set(val *Engine) {
	v.value = val
	v.isSet = true
}

func (v NullableEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEngine(val *Engine) *NullableEngine {
	return &NullableEngine{value: val, isSet: true}
}

func (v NullableEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


