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

// checks if the SearchTimeflowsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchTimeflowsResponse{}

// SearchTimeflowsResponse struct for SearchTimeflowsResponse
type SearchTimeflowsResponse struct {
	Items []Timeflow `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchTimeflowsResponse instantiates a new SearchTimeflowsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchTimeflowsResponse() *SearchTimeflowsResponse {
	this := SearchTimeflowsResponse{}
	return &this
}

// NewSearchTimeflowsResponseWithDefaults instantiates a new SearchTimeflowsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchTimeflowsResponseWithDefaults() *SearchTimeflowsResponse {
	this := SearchTimeflowsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchTimeflowsResponse) GetItems() []Timeflow {
	if o == nil || IsNil(o.Items) {
		var ret []Timeflow
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchTimeflowsResponse) GetItemsOk() ([]Timeflow, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchTimeflowsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Timeflow and assigns it to the Items field.
func (o *SearchTimeflowsResponse) SetItems(v []Timeflow) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchTimeflowsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchTimeflowsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchTimeflowsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchTimeflowsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchTimeflowsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchTimeflowsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchTimeflowsResponse struct {
	value *SearchTimeflowsResponse
	isSet bool
}

func (v NullableSearchTimeflowsResponse) Get() *SearchTimeflowsResponse {
	return v.value
}

func (v *NullableSearchTimeflowsResponse) Set(val *SearchTimeflowsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchTimeflowsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchTimeflowsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchTimeflowsResponse(val *SearchTimeflowsResponse) *NullableSearchTimeflowsResponse {
	return &NullableSearchTimeflowsResponse{value: val, isSet: true}
}

func (v NullableSearchTimeflowsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchTimeflowsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


