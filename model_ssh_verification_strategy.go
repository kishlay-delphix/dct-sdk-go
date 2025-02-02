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

// checks if the SSHVerificationStrategy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SSHVerificationStrategy{}

// SSHVerificationStrategy Mechanism to use for ssh host verification.
type SSHVerificationStrategy struct {
	// The name of the verification strategy.
	Name string `json:"name"`
	// The type of SSH key.
	KeyType *string `json:"key_type,omitempty"`
	// Base64-encoded ssh key of the host for RAW_KEY verification.
	RawKey *string `json:"raw_key,omitempty"`
	// Hash function for the fingerprint for FINGERPRINT verification.
	FingerprintType *string `json:"fingerprint_type,omitempty"`
	// Base-64 encoded fingerprint of the ssh key of the host for FINGERPRINT verification.
	Fingerprint *string `json:"fingerprint,omitempty"`
}

// NewSSHVerificationStrategy instantiates a new SSHVerificationStrategy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSSHVerificationStrategy(name string) *SSHVerificationStrategy {
	this := SSHVerificationStrategy{}
	this.Name = name
	return &this
}

// NewSSHVerificationStrategyWithDefaults instantiates a new SSHVerificationStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSSHVerificationStrategyWithDefaults() *SSHVerificationStrategy {
	this := SSHVerificationStrategy{}
	return &this
}

// GetName returns the Name field value
func (o *SSHVerificationStrategy) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SSHVerificationStrategy) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SSHVerificationStrategy) SetName(v string) {
	o.Name = v
}

// GetKeyType returns the KeyType field value if set, zero value otherwise.
func (o *SSHVerificationStrategy) GetKeyType() string {
	if o == nil || IsNil(o.KeyType) {
		var ret string
		return ret
	}
	return *o.KeyType
}

// GetKeyTypeOk returns a tuple with the KeyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHVerificationStrategy) GetKeyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.KeyType) {
		return nil, false
	}
	return o.KeyType, true
}

// HasKeyType returns a boolean if a field has been set.
func (o *SSHVerificationStrategy) HasKeyType() bool {
	if o != nil && !IsNil(o.KeyType) {
		return true
	}

	return false
}

// SetKeyType gets a reference to the given string and assigns it to the KeyType field.
func (o *SSHVerificationStrategy) SetKeyType(v string) {
	o.KeyType = &v
}

// GetRawKey returns the RawKey field value if set, zero value otherwise.
func (o *SSHVerificationStrategy) GetRawKey() string {
	if o == nil || IsNil(o.RawKey) {
		var ret string
		return ret
	}
	return *o.RawKey
}

// GetRawKeyOk returns a tuple with the RawKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHVerificationStrategy) GetRawKeyOk() (*string, bool) {
	if o == nil || IsNil(o.RawKey) {
		return nil, false
	}
	return o.RawKey, true
}

// HasRawKey returns a boolean if a field has been set.
func (o *SSHVerificationStrategy) HasRawKey() bool {
	if o != nil && !IsNil(o.RawKey) {
		return true
	}

	return false
}

// SetRawKey gets a reference to the given string and assigns it to the RawKey field.
func (o *SSHVerificationStrategy) SetRawKey(v string) {
	o.RawKey = &v
}

// GetFingerprintType returns the FingerprintType field value if set, zero value otherwise.
func (o *SSHVerificationStrategy) GetFingerprintType() string {
	if o == nil || IsNil(o.FingerprintType) {
		var ret string
		return ret
	}
	return *o.FingerprintType
}

// GetFingerprintTypeOk returns a tuple with the FingerprintType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHVerificationStrategy) GetFingerprintTypeOk() (*string, bool) {
	if o == nil || IsNil(o.FingerprintType) {
		return nil, false
	}
	return o.FingerprintType, true
}

// HasFingerprintType returns a boolean if a field has been set.
func (o *SSHVerificationStrategy) HasFingerprintType() bool {
	if o != nil && !IsNil(o.FingerprintType) {
		return true
	}

	return false
}

// SetFingerprintType gets a reference to the given string and assigns it to the FingerprintType field.
func (o *SSHVerificationStrategy) SetFingerprintType(v string) {
	o.FingerprintType = &v
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *SSHVerificationStrategy) GetFingerprint() string {
	if o == nil || IsNil(o.Fingerprint) {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SSHVerificationStrategy) GetFingerprintOk() (*string, bool) {
	if o == nil || IsNil(o.Fingerprint) {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *SSHVerificationStrategy) HasFingerprint() bool {
	if o != nil && !IsNil(o.Fingerprint) {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *SSHVerificationStrategy) SetFingerprint(v string) {
	o.Fingerprint = &v
}

func (o SSHVerificationStrategy) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SSHVerificationStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.KeyType) {
		toSerialize["key_type"] = o.KeyType
	}
	if !IsNil(o.RawKey) {
		toSerialize["raw_key"] = o.RawKey
	}
	if !IsNil(o.FingerprintType) {
		toSerialize["fingerprint_type"] = o.FingerprintType
	}
	if !IsNil(o.Fingerprint) {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	return toSerialize, nil
}

type NullableSSHVerificationStrategy struct {
	value *SSHVerificationStrategy
	isSet bool
}

func (v NullableSSHVerificationStrategy) Get() *SSHVerificationStrategy {
	return v.value
}

func (v *NullableSSHVerificationStrategy) Set(val *SSHVerificationStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableSSHVerificationStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableSSHVerificationStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSSHVerificationStrategy(val *SSHVerificationStrategy) *NullableSSHVerificationStrategy {
	return &NullableSSHVerificationStrategy{value: val, isSet: true}
}

func (v NullableSSHVerificationStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSSHVerificationStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


