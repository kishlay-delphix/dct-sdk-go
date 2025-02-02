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

// checks if the ListJobsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListJobsResponse{}

// ListJobsResponse struct for ListJobsResponse
type ListJobsResponse struct {
	Items []Job `json:"items,omitempty"`
	// Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request.
	Errors []Error `json:"errors,omitempty"`
	ResponseMetadata *PaginatedResponseMetadata `json:"response_metadata,omitempty"`
}

// NewListJobsResponse instantiates a new ListJobsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListJobsResponse() *ListJobsResponse {
	this := ListJobsResponse{}
	return &this
}

// NewListJobsResponseWithDefaults instantiates a new ListJobsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListJobsResponseWithDefaults() *ListJobsResponse {
	this := ListJobsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListJobsResponse) GetItems() []Job {
	if o == nil || IsNil(o.Items) {
		var ret []Job
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobsResponse) GetItemsOk() ([]Job, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListJobsResponse) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Job and assigns it to the Items field.
func (o *ListJobsResponse) SetItems(v []Job) {
	o.Items = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListJobsResponse) GetErrors() []Error {
	if o == nil || IsNil(o.Errors) {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListJobsResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListJobsResponse) SetErrors(v []Error) {
	o.Errors = v
}

// GetResponseMetadata returns the ResponseMetadata field value if set, zero value otherwise.
func (o *ListJobsResponse) GetResponseMetadata() PaginatedResponseMetadata {
	if o == nil || IsNil(o.ResponseMetadata) {
		var ret PaginatedResponseMetadata
		return ret
	}
	return *o.ResponseMetadata
}

// GetResponseMetadataOk returns a tuple with the ResponseMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListJobsResponse) GetResponseMetadataOk() (*PaginatedResponseMetadata, bool) {
	if o == nil || IsNil(o.ResponseMetadata) {
		return nil, false
	}
	return o.ResponseMetadata, true
}

// HasResponseMetadata returns a boolean if a field has been set.
func (o *ListJobsResponse) HasResponseMetadata() bool {
	if o != nil && !IsNil(o.ResponseMetadata) {
		return true
	}

	return false
}

// SetResponseMetadata gets a reference to the given PaginatedResponseMetadata and assigns it to the ResponseMetadata field.
func (o *ListJobsResponse) SetResponseMetadata(v PaginatedResponseMetadata) {
	o.ResponseMetadata = &v
}

func (o ListJobsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListJobsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !IsNil(o.ResponseMetadata) {
		toSerialize["response_metadata"] = o.ResponseMetadata
	}
	return toSerialize, nil
}

type NullableListJobsResponse struct {
	value *ListJobsResponse
	isSet bool
}

func (v NullableListJobsResponse) Get() *ListJobsResponse {
	return v.value
}

func (v *NullableListJobsResponse) Set(val *ListJobsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListJobsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListJobsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListJobsResponse(val *ListJobsResponse) *NullableListJobsResponse {
	return &NullableListJobsResponse{value: val, isSet: true}
}

func (v NullableListJobsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListJobsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


