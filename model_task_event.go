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

// checks if the TaskEvent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TaskEvent{}

// TaskEvent A step or event performed by a masking execution.
type TaskEvent struct {
	// The steps or events a task will perform.
	Event *string `json:"event,omitempty"`
	// The state of result of the task event.
	Status *string `json:"status,omitempty"`
}

// NewTaskEvent instantiates a new TaskEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTaskEvent() *TaskEvent {
	this := TaskEvent{}
	return &this
}

// NewTaskEventWithDefaults instantiates a new TaskEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTaskEventWithDefaults() *TaskEvent {
	this := TaskEvent{}
	return &this
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *TaskEvent) GetEvent() string {
	if o == nil || IsNil(o.Event) {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetEventOk() (*string, bool) {
	if o == nil || IsNil(o.Event) {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *TaskEvent) HasEvent() bool {
	if o != nil && !IsNil(o.Event) {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *TaskEvent) SetEvent(v string) {
	o.Event = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *TaskEvent) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TaskEvent) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *TaskEvent) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *TaskEvent) SetStatus(v string) {
	o.Status = &v
}

func (o TaskEvent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TaskEvent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Event) {
		toSerialize["event"] = o.Event
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableTaskEvent struct {
	value *TaskEvent
	isSet bool
}

func (v NullableTaskEvent) Get() *TaskEvent {
	return v.value
}

func (v *NullableTaskEvent) Set(val *TaskEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTaskEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTaskEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaskEvent(val *TaskEvent) *NullableTaskEvent {
	return &NullableTaskEvent{value: val, isSet: true}
}

func (v NullableTaskEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaskEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


