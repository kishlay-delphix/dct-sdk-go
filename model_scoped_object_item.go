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

// checks if the ScopedObjectItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScopedObjectItem{}

// ScopedObjectItem struct for ScopedObjectItem
type ScopedObjectItem struct {
	// ID of the object
	ObjectId string `json:"object_id"`
	ObjectType ObjectTypeEnum `json:"object_type"`
	Permission *PermissionEnum `json:"permission,omitempty"`
}

// NewScopedObjectItem instantiates a new ScopedObjectItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScopedObjectItem(objectId string, objectType ObjectTypeEnum) *ScopedObjectItem {
	this := ScopedObjectItem{}
	this.ObjectId = objectId
	this.ObjectType = objectType
	return &this
}

// NewScopedObjectItemWithDefaults instantiates a new ScopedObjectItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScopedObjectItemWithDefaults() *ScopedObjectItem {
	this := ScopedObjectItem{}
	return &this
}

// GetObjectId returns the ObjectId field value
func (o *ScopedObjectItem) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *ScopedObjectItem) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *ScopedObjectItem) SetObjectId(v string) {
	o.ObjectId = v
}

// GetObjectType returns the ObjectType field value
func (o *ScopedObjectItem) GetObjectType() ObjectTypeEnum {
	if o == nil {
		var ret ObjectTypeEnum
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ScopedObjectItem) GetObjectTypeOk() (*ObjectTypeEnum, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ScopedObjectItem) SetObjectType(v ObjectTypeEnum) {
	o.ObjectType = v
}

// GetPermission returns the Permission field value if set, zero value otherwise.
func (o *ScopedObjectItem) GetPermission() PermissionEnum {
	if o == nil || IsNil(o.Permission) {
		var ret PermissionEnum
		return ret
	}
	return *o.Permission
}

// GetPermissionOk returns a tuple with the Permission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScopedObjectItem) GetPermissionOk() (*PermissionEnum, bool) {
	if o == nil || IsNil(o.Permission) {
		return nil, false
	}
	return o.Permission, true
}

// HasPermission returns a boolean if a field has been set.
func (o *ScopedObjectItem) HasPermission() bool {
	if o != nil && !IsNil(o.Permission) {
		return true
	}

	return false
}

// SetPermission gets a reference to the given PermissionEnum and assigns it to the Permission field.
func (o *ScopedObjectItem) SetPermission(v PermissionEnum) {
	o.Permission = &v
}

func (o ScopedObjectItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScopedObjectItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["object_id"] = o.ObjectId
	toSerialize["object_type"] = o.ObjectType
	if !IsNil(o.Permission) {
		toSerialize["permission"] = o.Permission
	}
	return toSerialize, nil
}

type NullableScopedObjectItem struct {
	value *ScopedObjectItem
	isSet bool
}

func (v NullableScopedObjectItem) Get() *ScopedObjectItem {
	return v.value
}

func (v *NullableScopedObjectItem) Set(val *ScopedObjectItem) {
	v.value = val
	v.isSet = true
}

func (v NullableScopedObjectItem) IsSet() bool {
	return v.isSet
}

func (v *NullableScopedObjectItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScopedObjectItem(val *ScopedObjectItem) *NullableScopedObjectItem {
	return &NullableScopedObjectItem{value: val, isSet: true}
}

func (v NullableScopedObjectItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScopedObjectItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


