/*
Delphix DCT API

Delphix DCT API

API version: 1.0
Contact: support@delphix.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package delphix_dct_api

import (
	"encoding/json"
)

// ListVDBGroupsResponse struct for ListVDBGroupsResponse
type ListVDBGroupsResponse struct {
	Items []VDBGroup `json:"items,omitempty"`
	// Sadly, sometimes requests to the API are not successful. Failures can occur for a wide range of reasons. The Errors object contains information about full or partial failures which might have occurred during the request.
	Errors []Error `json:"errors,omitempty"`
}

// NewListVDBGroupsResponse instantiates a new ListVDBGroupsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListVDBGroupsResponse() *ListVDBGroupsResponse {
	this := ListVDBGroupsResponse{}
	return &this
}

// NewListVDBGroupsResponseWithDefaults instantiates a new ListVDBGroupsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListVDBGroupsResponseWithDefaults() *ListVDBGroupsResponse {
	this := ListVDBGroupsResponse{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *ListVDBGroupsResponse) GetItems() []VDBGroup {
	if o == nil || o.Items == nil {
		var ret []VDBGroup
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDBGroupsResponse) GetItemsOk() ([]VDBGroup, bool) {
	if o == nil || o.Items == nil {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *ListVDBGroupsResponse) HasItems() bool {
	if o != nil && o.Items != nil {
		return true
	}

	return false
}

// SetItems gets a reference to the given []VDBGroup and assigns it to the Items field.
func (o *ListVDBGroupsResponse) SetItems(v []VDBGroup) {
	o.Items = v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ListVDBGroupsResponse) GetErrors() []Error {
	if o == nil || o.Errors == nil {
		var ret []Error
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListVDBGroupsResponse) GetErrorsOk() ([]Error, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ListVDBGroupsResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []Error and assigns it to the Errors field.
func (o *ListVDBGroupsResponse) SetErrors(v []Error) {
	o.Errors = v
}

func (o ListVDBGroupsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Items != nil {
		toSerialize["items"] = o.Items
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableListVDBGroupsResponse struct {
	value *ListVDBGroupsResponse
	isSet bool
}

func (v NullableListVDBGroupsResponse) Get() *ListVDBGroupsResponse {
	return v.value
}

func (v *NullableListVDBGroupsResponse) Set(val *ListVDBGroupsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListVDBGroupsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListVDBGroupsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListVDBGroupsResponse(val *ListVDBGroupsResponse) *NullableListVDBGroupsResponse {
	return &NullableListVDBGroupsResponse{value: val, isSet: true}
}

func (v NullableListVDBGroupsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListVDBGroupsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


