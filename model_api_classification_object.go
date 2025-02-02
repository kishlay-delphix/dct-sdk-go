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

// checks if the ApiClassificationObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiClassificationObject{}

// ApiClassificationObject An API classification object which classifies APIs as automation or management.
type ApiClassificationObject struct {
	// HTTP method of the API.
	ApiMethod *string `json:"api_method,omitempty"`
	// context path of the API.
	Path *string `json:"path,omitempty"`
	// Either this API is automation or not.
	IsAutomation *bool `json:"is_automation,omitempty"`
	// The start date and time from when this api's is_automation definition has changed.
	StartDate *time.Time `json:"start_date,omitempty"`
	// The end date and time from when this api's is_automation definition has changed.
	EndDate *time.Time `json:"end_date,omitempty"`
}

// NewApiClassificationObject instantiates a new ApiClassificationObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiClassificationObject() *ApiClassificationObject {
	this := ApiClassificationObject{}
	return &this
}

// NewApiClassificationObjectWithDefaults instantiates a new ApiClassificationObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiClassificationObjectWithDefaults() *ApiClassificationObject {
	this := ApiClassificationObject{}
	return &this
}

// GetApiMethod returns the ApiMethod field value if set, zero value otherwise.
func (o *ApiClassificationObject) GetApiMethod() string {
	if o == nil || IsNil(o.ApiMethod) {
		var ret string
		return ret
	}
	return *o.ApiMethod
}

// GetApiMethodOk returns a tuple with the ApiMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClassificationObject) GetApiMethodOk() (*string, bool) {
	if o == nil || IsNil(o.ApiMethod) {
		return nil, false
	}
	return o.ApiMethod, true
}

// HasApiMethod returns a boolean if a field has been set.
func (o *ApiClassificationObject) HasApiMethod() bool {
	if o != nil && !IsNil(o.ApiMethod) {
		return true
	}

	return false
}

// SetApiMethod gets a reference to the given string and assigns it to the ApiMethod field.
func (o *ApiClassificationObject) SetApiMethod(v string) {
	o.ApiMethod = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *ApiClassificationObject) GetPath() string {
	if o == nil || IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClassificationObject) GetPathOk() (*string, bool) {
	if o == nil || IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *ApiClassificationObject) HasPath() bool {
	if o != nil && !IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *ApiClassificationObject) SetPath(v string) {
	o.Path = &v
}

// GetIsAutomation returns the IsAutomation field value if set, zero value otherwise.
func (o *ApiClassificationObject) GetIsAutomation() bool {
	if o == nil || IsNil(o.IsAutomation) {
		var ret bool
		return ret
	}
	return *o.IsAutomation
}

// GetIsAutomationOk returns a tuple with the IsAutomation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClassificationObject) GetIsAutomationOk() (*bool, bool) {
	if o == nil || IsNil(o.IsAutomation) {
		return nil, false
	}
	return o.IsAutomation, true
}

// HasIsAutomation returns a boolean if a field has been set.
func (o *ApiClassificationObject) HasIsAutomation() bool {
	if o != nil && !IsNil(o.IsAutomation) {
		return true
	}

	return false
}

// SetIsAutomation gets a reference to the given bool and assigns it to the IsAutomation field.
func (o *ApiClassificationObject) SetIsAutomation(v bool) {
	o.IsAutomation = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *ApiClassificationObject) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClassificationObject) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *ApiClassificationObject) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *ApiClassificationObject) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *ApiClassificationObject) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiClassificationObject) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *ApiClassificationObject) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *ApiClassificationObject) SetEndDate(v time.Time) {
	o.EndDate = &v
}

func (o ApiClassificationObject) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiClassificationObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiMethod) {
		toSerialize["api_method"] = o.ApiMethod
	}
	if !IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	if !IsNil(o.IsAutomation) {
		toSerialize["is_automation"] = o.IsAutomation
	}
	if !IsNil(o.StartDate) {
		toSerialize["start_date"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["end_date"] = o.EndDate
	}
	return toSerialize, nil
}

type NullableApiClassificationObject struct {
	value *ApiClassificationObject
	isSet bool
}

func (v NullableApiClassificationObject) Get() *ApiClassificationObject {
	return v.value
}

func (v *NullableApiClassificationObject) Set(val *ApiClassificationObject) {
	v.value = val
	v.isSet = true
}

func (v NullableApiClassificationObject) IsSet() bool {
	return v.isSet
}

func (v *NullableApiClassificationObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiClassificationObject(val *ApiClassificationObject) *NullableApiClassificationObject {
	return &NullableApiClassificationObject{value: val, isSet: true}
}

func (v NullableApiClassificationObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiClassificationObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


