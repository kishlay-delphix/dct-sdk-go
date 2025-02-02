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

// checks if the DSource type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DSource{}

// DSource The Delphix storage-based copy of the source database including its history.
type DSource struct {
	// The dSource object entity ID.
	Id *string `json:"id,omitempty"`
	// The database type of this dSource.
	DatabaseType NullableString `json:"database_type,omitempty"`
	// The container name of this dSource.
	Name NullableString `json:"name,omitempty"`
	// The database version of this dSource.
	DatabaseVersion NullableString `json:"database_version,omitempty"`
	// The content type of the dSource.
	ContentType NullableString `json:"content_type,omitempty"`
	// A universal ID that uniquely identifies the dSource database.
	DataUuid NullableString `json:"data_uuid,omitempty"`
	// The actual space used by this dSource, in bytes.
	StorageSize NullableInt64 `json:"storage_size,omitempty"`
	// The version of the plugin associated with this source database.
	PluginVersion NullableString `json:"plugin_version,omitempty"`
	// The date this dSource was created.
	CreationDate NullableTime `json:"creation_date,omitempty"`
	// The name of the group containing this dSource.
	GroupName NullableString `json:"group_name,omitempty"`
	// A value indicating whether this dSource is enabled.
	Enabled NullableBool `json:"enabled,omitempty"`
	// A reference to the Engine that this dSource belongs to.
	EngineId *string `json:"engine_id,omitempty"`
	// A reference to the Source associated with this dSource.
	SourceId NullableString `json:"source_id,omitempty"`
	// The runtime status of the dSource. 'Unknown' if all attempts to connect to the source failed.
	Status NullableString `json:"status,omitempty"`
	// Name of the Engine where this DSource is hosted
	EngineName NullableString `json:"engine_name,omitempty"`
	// A reference to the CDB associated with this dSource.
	CdbId NullableString `json:"cdb_id,omitempty"`
	// A reference to the currently active timeflow for this dSource.
	CurrentTimeflowId *string `json:"current_timeflow_id,omitempty"`
	// A reference to the previous timeflow for this dSource.
	PreviousTimeflowId *string `json:"previous_timeflow_id,omitempty"`
	Tags []Tag `json:"tags,omitempty"`
}

// NewDSource instantiates a new DSource object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDSource() *DSource {
	this := DSource{}
	return &this
}

// NewDSourceWithDefaults instantiates a new DSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDSourceWithDefaults() *DSource {
	this := DSource{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DSource) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSource) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DSource) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DSource) SetId(v string) {
	o.Id = &v
}

// GetDatabaseType returns the DatabaseType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetDatabaseType() string {
	if o == nil || IsNil(o.DatabaseType.Get()) {
		var ret string
		return ret
	}
	return *o.DatabaseType.Get()
}

// GetDatabaseTypeOk returns a tuple with the DatabaseType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetDatabaseTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatabaseType.Get(), o.DatabaseType.IsSet()
}

// HasDatabaseType returns a boolean if a field has been set.
func (o *DSource) HasDatabaseType() bool {
	if o != nil && o.DatabaseType.IsSet() {
		return true
	}

	return false
}

// SetDatabaseType gets a reference to the given NullableString and assigns it to the DatabaseType field.
func (o *DSource) SetDatabaseType(v string) {
	o.DatabaseType.Set(&v)
}
// SetDatabaseTypeNil sets the value for DatabaseType to be an explicit nil
func (o *DSource) SetDatabaseTypeNil() {
	o.DatabaseType.Set(nil)
}

// UnsetDatabaseType ensures that no value is present for DatabaseType, not even an explicit nil
func (o *DSource) UnsetDatabaseType() {
	o.DatabaseType.Unset()
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *DSource) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *DSource) SetName(v string) {
	o.Name.Set(&v)
}
// SetNameNil sets the value for Name to be an explicit nil
func (o *DSource) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *DSource) UnsetName() {
	o.Name.Unset()
}

// GetDatabaseVersion returns the DatabaseVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetDatabaseVersion() string {
	if o == nil || IsNil(o.DatabaseVersion.Get()) {
		var ret string
		return ret
	}
	return *o.DatabaseVersion.Get()
}

// GetDatabaseVersionOk returns a tuple with the DatabaseVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetDatabaseVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DatabaseVersion.Get(), o.DatabaseVersion.IsSet()
}

// HasDatabaseVersion returns a boolean if a field has been set.
func (o *DSource) HasDatabaseVersion() bool {
	if o != nil && o.DatabaseVersion.IsSet() {
		return true
	}

	return false
}

// SetDatabaseVersion gets a reference to the given NullableString and assigns it to the DatabaseVersion field.
func (o *DSource) SetDatabaseVersion(v string) {
	o.DatabaseVersion.Set(&v)
}
// SetDatabaseVersionNil sets the value for DatabaseVersion to be an explicit nil
func (o *DSource) SetDatabaseVersionNil() {
	o.DatabaseVersion.Set(nil)
}

// UnsetDatabaseVersion ensures that no value is present for DatabaseVersion, not even an explicit nil
func (o *DSource) UnsetDatabaseVersion() {
	o.DatabaseVersion.Unset()
}

// GetContentType returns the ContentType field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetContentType() string {
	if o == nil || IsNil(o.ContentType.Get()) {
		var ret string
		return ret
	}
	return *o.ContentType.Get()
}

// GetContentTypeOk returns a tuple with the ContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetContentTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContentType.Get(), o.ContentType.IsSet()
}

// HasContentType returns a boolean if a field has been set.
func (o *DSource) HasContentType() bool {
	if o != nil && o.ContentType.IsSet() {
		return true
	}

	return false
}

// SetContentType gets a reference to the given NullableString and assigns it to the ContentType field.
func (o *DSource) SetContentType(v string) {
	o.ContentType.Set(&v)
}
// SetContentTypeNil sets the value for ContentType to be an explicit nil
func (o *DSource) SetContentTypeNil() {
	o.ContentType.Set(nil)
}

// UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
func (o *DSource) UnsetContentType() {
	o.ContentType.Unset()
}

// GetDataUuid returns the DataUuid field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetDataUuid() string {
	if o == nil || IsNil(o.DataUuid.Get()) {
		var ret string
		return ret
	}
	return *o.DataUuid.Get()
}

// GetDataUuidOk returns a tuple with the DataUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetDataUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.DataUuid.Get(), o.DataUuid.IsSet()
}

// HasDataUuid returns a boolean if a field has been set.
func (o *DSource) HasDataUuid() bool {
	if o != nil && o.DataUuid.IsSet() {
		return true
	}

	return false
}

// SetDataUuid gets a reference to the given NullableString and assigns it to the DataUuid field.
func (o *DSource) SetDataUuid(v string) {
	o.DataUuid.Set(&v)
}
// SetDataUuidNil sets the value for DataUuid to be an explicit nil
func (o *DSource) SetDataUuidNil() {
	o.DataUuid.Set(nil)
}

// UnsetDataUuid ensures that no value is present for DataUuid, not even an explicit nil
func (o *DSource) UnsetDataUuid() {
	o.DataUuid.Unset()
}

// GetStorageSize returns the StorageSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetStorageSize() int64 {
	if o == nil || IsNil(o.StorageSize.Get()) {
		var ret int64
		return ret
	}
	return *o.StorageSize.Get()
}

// GetStorageSizeOk returns a tuple with the StorageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetStorageSizeOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.StorageSize.Get(), o.StorageSize.IsSet()
}

// HasStorageSize returns a boolean if a field has been set.
func (o *DSource) HasStorageSize() bool {
	if o != nil && o.StorageSize.IsSet() {
		return true
	}

	return false
}

// SetStorageSize gets a reference to the given NullableInt64 and assigns it to the StorageSize field.
func (o *DSource) SetStorageSize(v int64) {
	o.StorageSize.Set(&v)
}
// SetStorageSizeNil sets the value for StorageSize to be an explicit nil
func (o *DSource) SetStorageSizeNil() {
	o.StorageSize.Set(nil)
}

// UnsetStorageSize ensures that no value is present for StorageSize, not even an explicit nil
func (o *DSource) UnsetStorageSize() {
	o.StorageSize.Unset()
}

// GetPluginVersion returns the PluginVersion field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetPluginVersion() string {
	if o == nil || IsNil(o.PluginVersion.Get()) {
		var ret string
		return ret
	}
	return *o.PluginVersion.Get()
}

// GetPluginVersionOk returns a tuple with the PluginVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetPluginVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.PluginVersion.Get(), o.PluginVersion.IsSet()
}

// HasPluginVersion returns a boolean if a field has been set.
func (o *DSource) HasPluginVersion() bool {
	if o != nil && o.PluginVersion.IsSet() {
		return true
	}

	return false
}

// SetPluginVersion gets a reference to the given NullableString and assigns it to the PluginVersion field.
func (o *DSource) SetPluginVersion(v string) {
	o.PluginVersion.Set(&v)
}
// SetPluginVersionNil sets the value for PluginVersion to be an explicit nil
func (o *DSource) SetPluginVersionNil() {
	o.PluginVersion.Set(nil)
}

// UnsetPluginVersion ensures that no value is present for PluginVersion, not even an explicit nil
func (o *DSource) UnsetPluginVersion() {
	o.PluginVersion.Unset()
}

// GetCreationDate returns the CreationDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetCreationDate() time.Time {
	if o == nil || IsNil(o.CreationDate.Get()) {
		var ret time.Time
		return ret
	}
	return *o.CreationDate.Get()
}

// GetCreationDateOk returns a tuple with the CreationDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetCreationDateOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.CreationDate.Get(), o.CreationDate.IsSet()
}

// HasCreationDate returns a boolean if a field has been set.
func (o *DSource) HasCreationDate() bool {
	if o != nil && o.CreationDate.IsSet() {
		return true
	}

	return false
}

// SetCreationDate gets a reference to the given NullableTime and assigns it to the CreationDate field.
func (o *DSource) SetCreationDate(v time.Time) {
	o.CreationDate.Set(&v)
}
// SetCreationDateNil sets the value for CreationDate to be an explicit nil
func (o *DSource) SetCreationDateNil() {
	o.CreationDate.Set(nil)
}

// UnsetCreationDate ensures that no value is present for CreationDate, not even an explicit nil
func (o *DSource) UnsetCreationDate() {
	o.CreationDate.Unset()
}

// GetGroupName returns the GroupName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetGroupName() string {
	if o == nil || IsNil(o.GroupName.Get()) {
		var ret string
		return ret
	}
	return *o.GroupName.Get()
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetGroupNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.GroupName.Get(), o.GroupName.IsSet()
}

// HasGroupName returns a boolean if a field has been set.
func (o *DSource) HasGroupName() bool {
	if o != nil && o.GroupName.IsSet() {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given NullableString and assigns it to the GroupName field.
func (o *DSource) SetGroupName(v string) {
	o.GroupName.Set(&v)
}
// SetGroupNameNil sets the value for GroupName to be an explicit nil
func (o *DSource) SetGroupNameNil() {
	o.GroupName.Set(nil)
}

// UnsetGroupName ensures that no value is present for GroupName, not even an explicit nil
func (o *DSource) UnsetGroupName() {
	o.GroupName.Unset()
}

// GetEnabled returns the Enabled field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled.Get()) {
		var ret bool
		return ret
	}
	return *o.Enabled.Get()
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetEnabledOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Enabled.Get(), o.Enabled.IsSet()
}

// HasEnabled returns a boolean if a field has been set.
func (o *DSource) HasEnabled() bool {
	if o != nil && o.Enabled.IsSet() {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given NullableBool and assigns it to the Enabled field.
func (o *DSource) SetEnabled(v bool) {
	o.Enabled.Set(&v)
}
// SetEnabledNil sets the value for Enabled to be an explicit nil
func (o *DSource) SetEnabledNil() {
	o.Enabled.Set(nil)
}

// UnsetEnabled ensures that no value is present for Enabled, not even an explicit nil
func (o *DSource) UnsetEnabled() {
	o.Enabled.Unset()
}

// GetEngineId returns the EngineId field value if set, zero value otherwise.
func (o *DSource) GetEngineId() string {
	if o == nil || IsNil(o.EngineId) {
		var ret string
		return ret
	}
	return *o.EngineId
}

// GetEngineIdOk returns a tuple with the EngineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSource) GetEngineIdOk() (*string, bool) {
	if o == nil || IsNil(o.EngineId) {
		return nil, false
	}
	return o.EngineId, true
}

// HasEngineId returns a boolean if a field has been set.
func (o *DSource) HasEngineId() bool {
	if o != nil && !IsNil(o.EngineId) {
		return true
	}

	return false
}

// SetEngineId gets a reference to the given string and assigns it to the EngineId field.
func (o *DSource) SetEngineId(v string) {
	o.EngineId = &v
}

// GetSourceId returns the SourceId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetSourceId() string {
	if o == nil || IsNil(o.SourceId.Get()) {
		var ret string
		return ret
	}
	return *o.SourceId.Get()
}

// GetSourceIdOk returns a tuple with the SourceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceId.Get(), o.SourceId.IsSet()
}

// HasSourceId returns a boolean if a field has been set.
func (o *DSource) HasSourceId() bool {
	if o != nil && o.SourceId.IsSet() {
		return true
	}

	return false
}

// SetSourceId gets a reference to the given NullableString and assigns it to the SourceId field.
func (o *DSource) SetSourceId(v string) {
	o.SourceId.Set(&v)
}
// SetSourceIdNil sets the value for SourceId to be an explicit nil
func (o *DSource) SetSourceIdNil() {
	o.SourceId.Set(nil)
}

// UnsetSourceId ensures that no value is present for SourceId, not even an explicit nil
func (o *DSource) UnsetSourceId() {
	o.SourceId.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetStatus() string {
	if o == nil || IsNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *DSource) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *DSource) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *DSource) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *DSource) UnsetStatus() {
	o.Status.Unset()
}

// GetEngineName returns the EngineName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetEngineName() string {
	if o == nil || IsNil(o.EngineName.Get()) {
		var ret string
		return ret
	}
	return *o.EngineName.Get()
}

// GetEngineNameOk returns a tuple with the EngineName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetEngineNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EngineName.Get(), o.EngineName.IsSet()
}

// HasEngineName returns a boolean if a field has been set.
func (o *DSource) HasEngineName() bool {
	if o != nil && o.EngineName.IsSet() {
		return true
	}

	return false
}

// SetEngineName gets a reference to the given NullableString and assigns it to the EngineName field.
func (o *DSource) SetEngineName(v string) {
	o.EngineName.Set(&v)
}
// SetEngineNameNil sets the value for EngineName to be an explicit nil
func (o *DSource) SetEngineNameNil() {
	o.EngineName.Set(nil)
}

// UnsetEngineName ensures that no value is present for EngineName, not even an explicit nil
func (o *DSource) UnsetEngineName() {
	o.EngineName.Unset()
}

// GetCdbId returns the CdbId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *DSource) GetCdbId() string {
	if o == nil || IsNil(o.CdbId.Get()) {
		var ret string
		return ret
	}
	return *o.CdbId.Get()
}

// GetCdbIdOk returns a tuple with the CdbId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *DSource) GetCdbIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CdbId.Get(), o.CdbId.IsSet()
}

// HasCdbId returns a boolean if a field has been set.
func (o *DSource) HasCdbId() bool {
	if o != nil && o.CdbId.IsSet() {
		return true
	}

	return false
}

// SetCdbId gets a reference to the given NullableString and assigns it to the CdbId field.
func (o *DSource) SetCdbId(v string) {
	o.CdbId.Set(&v)
}
// SetCdbIdNil sets the value for CdbId to be an explicit nil
func (o *DSource) SetCdbIdNil() {
	o.CdbId.Set(nil)
}

// UnsetCdbId ensures that no value is present for CdbId, not even an explicit nil
func (o *DSource) UnsetCdbId() {
	o.CdbId.Unset()
}

// GetCurrentTimeflowId returns the CurrentTimeflowId field value if set, zero value otherwise.
func (o *DSource) GetCurrentTimeflowId() string {
	if o == nil || IsNil(o.CurrentTimeflowId) {
		var ret string
		return ret
	}
	return *o.CurrentTimeflowId
}

// GetCurrentTimeflowIdOk returns a tuple with the CurrentTimeflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSource) GetCurrentTimeflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.CurrentTimeflowId) {
		return nil, false
	}
	return o.CurrentTimeflowId, true
}

// HasCurrentTimeflowId returns a boolean if a field has been set.
func (o *DSource) HasCurrentTimeflowId() bool {
	if o != nil && !IsNil(o.CurrentTimeflowId) {
		return true
	}

	return false
}

// SetCurrentTimeflowId gets a reference to the given string and assigns it to the CurrentTimeflowId field.
func (o *DSource) SetCurrentTimeflowId(v string) {
	o.CurrentTimeflowId = &v
}

// GetPreviousTimeflowId returns the PreviousTimeflowId field value if set, zero value otherwise.
func (o *DSource) GetPreviousTimeflowId() string {
	if o == nil || IsNil(o.PreviousTimeflowId) {
		var ret string
		return ret
	}
	return *o.PreviousTimeflowId
}

// GetPreviousTimeflowIdOk returns a tuple with the PreviousTimeflowId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSource) GetPreviousTimeflowIdOk() (*string, bool) {
	if o == nil || IsNil(o.PreviousTimeflowId) {
		return nil, false
	}
	return o.PreviousTimeflowId, true
}

// HasPreviousTimeflowId returns a boolean if a field has been set.
func (o *DSource) HasPreviousTimeflowId() bool {
	if o != nil && !IsNil(o.PreviousTimeflowId) {
		return true
	}

	return false
}

// SetPreviousTimeflowId gets a reference to the given string and assigns it to the PreviousTimeflowId field.
func (o *DSource) SetPreviousTimeflowId(v string) {
	o.PreviousTimeflowId = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DSource) GetTags() []Tag {
	if o == nil || IsNil(o.Tags) {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DSource) GetTagsOk() ([]Tag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DSource) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *DSource) SetTags(v []Tag) {
	o.Tags = v
}

func (o DSource) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.DatabaseType.IsSet() {
		toSerialize["database_type"] = o.DatabaseType.Get()
	}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}
	if o.DatabaseVersion.IsSet() {
		toSerialize["database_version"] = o.DatabaseVersion.Get()
	}
	if o.ContentType.IsSet() {
		toSerialize["content_type"] = o.ContentType.Get()
	}
	if o.DataUuid.IsSet() {
		toSerialize["data_uuid"] = o.DataUuid.Get()
	}
	if o.StorageSize.IsSet() {
		toSerialize["storage_size"] = o.StorageSize.Get()
	}
	if o.PluginVersion.IsSet() {
		toSerialize["plugin_version"] = o.PluginVersion.Get()
	}
	if o.CreationDate.IsSet() {
		toSerialize["creation_date"] = o.CreationDate.Get()
	}
	if o.GroupName.IsSet() {
		toSerialize["group_name"] = o.GroupName.Get()
	}
	if o.Enabled.IsSet() {
		toSerialize["enabled"] = o.Enabled.Get()
	}
	if !IsNil(o.EngineId) {
		toSerialize["engine_id"] = o.EngineId
	}
	if o.SourceId.IsSet() {
		toSerialize["source_id"] = o.SourceId.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.EngineName.IsSet() {
		toSerialize["engine_name"] = o.EngineName.Get()
	}
	if o.CdbId.IsSet() {
		toSerialize["cdb_id"] = o.CdbId.Get()
	}
	if !IsNil(o.CurrentTimeflowId) {
		toSerialize["current_timeflow_id"] = o.CurrentTimeflowId
	}
	if !IsNil(o.PreviousTimeflowId) {
		toSerialize["previous_timeflow_id"] = o.PreviousTimeflowId
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	return toSerialize, nil
}

type NullableDSource struct {
	value *DSource
	isSet bool
}

func (v NullableDSource) Get() *DSource {
	return v.value
}

func (v *NullableDSource) Set(val *DSource) {
	v.value = val
	v.isSet = true
}

func (v NullableDSource) IsSet() bool {
	return v.isSet
}

func (v *NullableDSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDSource(val *DSource) *NullableDSource {
	return &NullableDSource{value: val, isSet: true}
}

func (v NullableDSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


