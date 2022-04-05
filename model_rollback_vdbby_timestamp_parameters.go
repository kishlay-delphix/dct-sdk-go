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
	"time"
)

// RollbackVDBByTimestampParameters struct for RollbackVDBByTimestampParameters
type RollbackVDBByTimestampParameters struct {
	// The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp.
	TimestampInDatabaseTimezone *string `json:"timestamp_in_database_timezone,omitempty"`
}

// NewRollbackVDBByTimestampParameters instantiates a new RollbackVDBByTimestampParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRollbackVDBByTimestampParameters() *RollbackVDBByTimestampParameters {
	this := RollbackVDBByTimestampParameters{}
	return &this
}

// NewRollbackVDBByTimestampParametersWithDefaults instantiates a new RollbackVDBByTimestampParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollbackVDBByTimestampParametersWithDefaults() *RollbackVDBByTimestampParameters {
	this := RollbackVDBByTimestampParameters{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *RollbackVDBByTimestampParameters) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackVDBByTimestampParameters) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *RollbackVDBByTimestampParameters) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *RollbackVDBByTimestampParameters) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTimestampInDatabaseTimezone returns the TimestampInDatabaseTimezone field value if set, zero value otherwise.
func (o *RollbackVDBByTimestampParameters) GetTimestampInDatabaseTimezone() string {
	if o == nil || o.TimestampInDatabaseTimezone == nil {
		var ret string
		return ret
	}
	return *o.TimestampInDatabaseTimezone
}

// GetTimestampInDatabaseTimezoneOk returns a tuple with the TimestampInDatabaseTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RollbackVDBByTimestampParameters) GetTimestampInDatabaseTimezoneOk() (*string, bool) {
	if o == nil || o.TimestampInDatabaseTimezone == nil {
		return nil, false
	}
	return o.TimestampInDatabaseTimezone, true
}

// HasTimestampInDatabaseTimezone returns a boolean if a field has been set.
func (o *RollbackVDBByTimestampParameters) HasTimestampInDatabaseTimezone() bool {
	if o != nil && o.TimestampInDatabaseTimezone != nil {
		return true
	}

	return false
}

// SetTimestampInDatabaseTimezone gets a reference to the given string and assigns it to the TimestampInDatabaseTimezone field.
func (o *RollbackVDBByTimestampParameters) SetTimestampInDatabaseTimezone(v string) {
	o.TimestampInDatabaseTimezone = &v
}

func (o RollbackVDBByTimestampParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.TimestampInDatabaseTimezone != nil {
		toSerialize["timestamp_in_database_timezone"] = o.TimestampInDatabaseTimezone
	}
	return json.Marshal(toSerialize)
}

type NullableRollbackVDBByTimestampParameters struct {
	value *RollbackVDBByTimestampParameters
	isSet bool
}

func (v NullableRollbackVDBByTimestampParameters) Get() *RollbackVDBByTimestampParameters {
	return v.value
}

func (v *NullableRollbackVDBByTimestampParameters) Set(val *RollbackVDBByTimestampParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableRollbackVDBByTimestampParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableRollbackVDBByTimestampParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollbackVDBByTimestampParameters(val *RollbackVDBByTimestampParameters) *NullableRollbackVDBByTimestampParameters {
	return &NullableRollbackVDBByTimestampParameters{value: val, isSet: true}
}

func (v NullableRollbackVDBByTimestampParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollbackVDBByTimestampParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


