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

// checks if the SearchSnapshotsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SearchSnapshotsResponse{}

// SearchSnapshotsResponse struct for SearchSnapshotsResponse
type SearchSnapshotsResponse struct {
	Items []Snapshot `json:"items,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewSearchSnapshotsResponse instantiates a new SearchSnapshotsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSearchSnapshotsResponse() *SearchSnapshotsResponse {
	this := SearchSnapshotsResponse{}
	return &this
}

// NewSearchSnapshotsResponseWithDefaults instantiates a new SearchSnapshotsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSearchSnapshotsResponseWithDefaults() *SearchSnapshotsResponse {
	this := SearchSnapshotsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SearchSnapshotsResponse) GetItems() []Snapshot {
	if o == nil || IsNil(o.Items) {
		var ret []Snapshot
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSnapshotsResponse) GetItemsOk() ([]Snapshot, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SearchSnapshotsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Snapshot and assigns it to the Items field.
func (o *SearchSnapshotsResponse) SetItems(v []Snapshot) {
	o.Items = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *SearchSnapshotsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SearchSnapshotsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *SearchSnapshotsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *SearchSnapshotsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o SearchSnapshotsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SearchSnapshotsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableSearchSnapshotsResponse struct {
	value *SearchSnapshotsResponse
	isSet bool
}

func (v NullableSearchSnapshotsResponse) Get() *SearchSnapshotsResponse {
	return v.value
}

func (v *NullableSearchSnapshotsResponse) Set(val *SearchSnapshotsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSearchSnapshotsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSearchSnapshotsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSearchSnapshotsResponse(val *SearchSnapshotsResponse) *NullableSearchSnapshotsResponse {
	return &NullableSearchSnapshotsResponse{value: val, isSet: true}
}

func (v NullableSearchSnapshotsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSearchSnapshotsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


