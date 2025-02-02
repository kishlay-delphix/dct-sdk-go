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

// checks if the DataPointByTimestampParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataPointByTimestampParameters{}

// DataPointByTimestampParameters struct for DataPointByTimestampParameters
type DataPointByTimestampParameters struct {
	// The point in time from which to execute the operation. Mutually exclusive with timestamp_in_database_timezone. If the timestamp is not set, selects the latest point.
	Timestamp *time.Time `json:"timestamp,omitempty"`
	// The point in time from which to execute the operation, expressed as a date-time in the timezone of the source database. Mutually exclusive with timestamp.
	TimestampInDatabaseTimezone *string `json:"timestamp_in_database_timezone,omitempty"`
}

// NewDataPointByTimestampParameters instantiates a new DataPointByTimestampParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataPointByTimestampParameters() *DataPointByTimestampParameters {
	this := DataPointByTimestampParameters{}
	return &this
}

// NewDataPointByTimestampParametersWithDefaults instantiates a new DataPointByTimestampParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataPointByTimestampParametersWithDefaults() *DataPointByTimestampParameters {
	this := DataPointByTimestampParameters{}
	return &this
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DataPointByTimestampParameters) GetTimestamp() time.Time {
	if o == nil || IsNil(o.Timestamp) {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPointByTimestampParameters) GetTimestampOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DataPointByTimestampParameters) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *DataPointByTimestampParameters) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetTimestampInDatabaseTimezone returns the TimestampInDatabaseTimezone field value if set, zero value otherwise.
func (o *DataPointByTimestampParameters) GetTimestampInDatabaseTimezone() string {
	if o == nil || IsNil(o.TimestampInDatabaseTimezone) {
		var ret string
		return ret
	}
	return *o.TimestampInDatabaseTimezone
}

// GetTimestampInDatabaseTimezoneOk returns a tuple with the TimestampInDatabaseTimezone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataPointByTimestampParameters) GetTimestampInDatabaseTimezoneOk() (*string, bool) {
	if o == nil || IsNil(o.TimestampInDatabaseTimezone) {
		return nil, false
	}
	return o.TimestampInDatabaseTimezone, true
}

// HasTimestampInDatabaseTimezone returns a boolean if a field has been set.
func (o *DataPointByTimestampParameters) HasTimestampInDatabaseTimezone() bool {
	if o != nil && !IsNil(o.TimestampInDatabaseTimezone) {
		return true
	}

	return false
}

// SetTimestampInDatabaseTimezone gets a reference to the given string and assigns it to the TimestampInDatabaseTimezone field.
func (o *DataPointByTimestampParameters) SetTimestampInDatabaseTimezone(v string) {
	o.TimestampInDatabaseTimezone = &v
}

func (o DataPointByTimestampParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataPointByTimestampParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.TimestampInDatabaseTimezone) {
		toSerialize["timestamp_in_database_timezone"] = o.TimestampInDatabaseTimezone
	}
	return toSerialize, nil
}

type NullableDataPointByTimestampParameters struct {
	value *DataPointByTimestampParameters
	isSet bool
}

func (v NullableDataPointByTimestampParameters) Get() *DataPointByTimestampParameters {
	return v.value
}

func (v *NullableDataPointByTimestampParameters) Set(val *DataPointByTimestampParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableDataPointByTimestampParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableDataPointByTimestampParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataPointByTimestampParameters(val *DataPointByTimestampParameters) *NullableDataPointByTimestampParameters {
	return &NullableDataPointByTimestampParameters{value: val, isSet: true}
}

func (v NullableDataPointByTimestampParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataPointByTimestampParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


