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

// TagsResponse struct for TagsResponse
type TagsResponse struct {
	// Array of tags with key value pairs
	Tags []Tag `json:"tags,omitempty"`
}

// NewTagsResponse instantiates a new TagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTagsResponse() *TagsResponse {
	this := TagsResponse{}
	return &this
}

// NewTagsResponseWithDefaults instantiates a new TagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagsResponseWithDefaults() *TagsResponse {
	this := TagsResponse{}
	return &this
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TagsResponse) GetTags() []Tag {
	if o == nil || o.Tags == nil {
		var ret []Tag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TagsResponse) GetTagsOk() ([]Tag, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TagsResponse) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []Tag and assigns it to the Tags field.
func (o *TagsResponse) SetTags(v []Tag) {
	o.Tags = v
}

func (o TagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableTagsResponse struct {
	value *TagsResponse
	isSet bool
}

func (v NullableTagsResponse) Get() *TagsResponse {
	return v.value
}

func (v *NullableTagsResponse) Set(val *TagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTagsResponse(val *TagsResponse) *NullableTagsResponse {
	return &NullableTagsResponse{value: val, isSet: true}
}

func (v NullableTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


