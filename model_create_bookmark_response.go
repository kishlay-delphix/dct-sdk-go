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

// checks if the CreateBookmarkResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateBookmarkResponse{}

// CreateBookmarkResponse struct for CreateBookmarkResponse
type CreateBookmarkResponse struct {
	Bookmark *Bookmark `json:"bookmark,omitempty"`
	Job *Job `json:"job,omitempty"`
}

// NewCreateBookmarkResponse instantiates a new CreateBookmarkResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateBookmarkResponse() *CreateBookmarkResponse {
	this := CreateBookmarkResponse{}
	return &this
}

// NewCreateBookmarkResponseWithDefaults instantiates a new CreateBookmarkResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateBookmarkResponseWithDefaults() *CreateBookmarkResponse {
	this := CreateBookmarkResponse{}
	return &this
}

// GetBookmark returns the Bookmark field value if set, zero value otherwise.
func (o *CreateBookmarkResponse) GetBookmark() Bookmark {
	if o == nil || IsNil(o.Bookmark) {
		var ret Bookmark
		return ret
	}
	return *o.Bookmark
}

// GetBookmarkOk returns a tuple with the Bookmark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBookmarkResponse) GetBookmarkOk() (*Bookmark, bool) {
	if o == nil || IsNil(o.Bookmark) {
		return nil, false
	}
	return o.Bookmark, true
}

// HasBookmark returns a boolean if a field has been set.
func (o *CreateBookmarkResponse) HasBookmark() bool {
	if o != nil && !IsNil(o.Bookmark) {
		return true
	}

	return false
}

// SetBookmark gets a reference to the given Bookmark and assigns it to the Bookmark field.
func (o *CreateBookmarkResponse) SetBookmark(v Bookmark) {
	o.Bookmark = &v
}

// GetJob returns the Job field value if set, zero value otherwise.
func (o *CreateBookmarkResponse) GetJob() Job {
	if o == nil || IsNil(o.Job) {
		var ret Job
		return ret
	}
	return *o.Job
}

// GetJobOk returns a tuple with the Job field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateBookmarkResponse) GetJobOk() (*Job, bool) {
	if o == nil || IsNil(o.Job) {
		return nil, false
	}
	return o.Job, true
}

// HasJob returns a boolean if a field has been set.
func (o *CreateBookmarkResponse) HasJob() bool {
	if o != nil && !IsNil(o.Job) {
		return true
	}

	return false
}

// SetJob gets a reference to the given Job and assigns it to the Job field.
func (o *CreateBookmarkResponse) SetJob(v Job) {
	o.Job = &v
}

func (o CreateBookmarkResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateBookmarkResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Bookmark) {
		toSerialize["bookmark"] = o.Bookmark
	}
	if !IsNil(o.Job) {
		toSerialize["job"] = o.Job
	}
	return toSerialize, nil
}

type NullableCreateBookmarkResponse struct {
	value *CreateBookmarkResponse
	isSet bool
}

func (v NullableCreateBookmarkResponse) Get() *CreateBookmarkResponse {
	return v.value
}

func (v *NullableCreateBookmarkResponse) Set(val *CreateBookmarkResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateBookmarkResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateBookmarkResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateBookmarkResponse(val *CreateBookmarkResponse) *NullableCreateBookmarkResponse {
	return &NullableCreateBookmarkResponse{value: val, isSet: true}
}

func (v NullableCreateBookmarkResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateBookmarkResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


