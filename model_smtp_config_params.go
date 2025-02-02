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

// checks if the SMTPConfigParams type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SMTPConfigParams{}

// SMTPConfigParams Parameters to read or update SMTP Config
type SMTPConfigParams struct {
	// True if outbound email is enabled.
	Enabled *bool `json:"enabled,omitempty"`
	// IP address or hostname of SMTP relay server.
	Server NullableString `json:"server,omitempty"`
	// Port number to use. A value of -1 indicates the default (25 or 587 for TLS).
	Port NullableInt32 `json:"port,omitempty"`
	// True if username/password authentication should be used.
	AuthenticationEnabled *bool `json:"authentication_enabled,omitempty"`
	// True if TLS (transport layer security) should be used.
	TlsEnabled *bool `json:"tls_enabled,omitempty"`
	// If authentication is enabled, username to use when authenticating to the server.
	Username NullableString `json:"username,omitempty"`
	// If authentication is enabled, password to use when authenticating to the server.
	Password NullableString `json:"password,omitempty"`
	// From address to use when sending mail. If unspecified, 'noreply@delphix.com' is used.
	FromAddress NullableString `json:"from_address,omitempty"`
	// Maximum timeout to wait, in seconds, when sending mail.
	SendTimeout NullableInt32 `json:"send_timeout,omitempty"`
}

// NewSMTPConfigParams instantiates a new SMTPConfigParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMTPConfigParams() *SMTPConfigParams {
	this := SMTPConfigParams{}
	return &this
}

// NewSMTPConfigParamsWithDefaults instantiates a new SMTPConfigParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMTPConfigParamsWithDefaults() *SMTPConfigParams {
	this := SMTPConfigParams{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *SMTPConfigParams) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMTPConfigParams) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *SMTPConfigParams) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetServer returns the Server field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMTPConfigParams) GetServer() string {
	if o == nil || IsNil(o.Server.Get()) {
		var ret string
		return ret
	}
	return *o.Server.Get()
}

// GetServerOk returns a tuple with the Server field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMTPConfigParams) GetServerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Server.Get(), o.Server.IsSet()
}

// HasServer returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasServer() bool {
	if o != nil && o.Server.IsSet() {
		return true
	}

	return false
}

// SetServer gets a reference to the given NullableString and assigns it to the Server field.
func (o *SMTPConfigParams) SetServer(v string) {
	o.Server.Set(&v)
}
// SetServerNil sets the value for Server to be an explicit nil
func (o *SMTPConfigParams) SetServerNil() {
	o.Server.Set(nil)
}

// UnsetServer ensures that no value is present for Server, not even an explicit nil
func (o *SMTPConfigParams) UnsetServer() {
	o.Server.Unset()
}

// GetPort returns the Port field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMTPConfigParams) GetPort() int32 {
	if o == nil || IsNil(o.Port.Get()) {
		var ret int32
		return ret
	}
	return *o.Port.Get()
}

// GetPortOk returns a tuple with the Port field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMTPConfigParams) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port.Get(), o.Port.IsSet()
}

// HasPort returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasPort() bool {
	if o != nil && o.Port.IsSet() {
		return true
	}

	return false
}

// SetPort gets a reference to the given NullableInt32 and assigns it to the Port field.
func (o *SMTPConfigParams) SetPort(v int32) {
	o.Port.Set(&v)
}
// SetPortNil sets the value for Port to be an explicit nil
func (o *SMTPConfigParams) SetPortNil() {
	o.Port.Set(nil)
}

// UnsetPort ensures that no value is present for Port, not even an explicit nil
func (o *SMTPConfigParams) UnsetPort() {
	o.Port.Unset()
}

// GetAuthenticationEnabled returns the AuthenticationEnabled field value if set, zero value otherwise.
func (o *SMTPConfigParams) GetAuthenticationEnabled() bool {
	if o == nil || IsNil(o.AuthenticationEnabled) {
		var ret bool
		return ret
	}
	return *o.AuthenticationEnabled
}

// GetAuthenticationEnabledOk returns a tuple with the AuthenticationEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMTPConfigParams) GetAuthenticationEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.AuthenticationEnabled) {
		return nil, false
	}
	return o.AuthenticationEnabled, true
}

// HasAuthenticationEnabled returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasAuthenticationEnabled() bool {
	if o != nil && !IsNil(o.AuthenticationEnabled) {
		return true
	}

	return false
}

// SetAuthenticationEnabled gets a reference to the given bool and assigns it to the AuthenticationEnabled field.
func (o *SMTPConfigParams) SetAuthenticationEnabled(v bool) {
	o.AuthenticationEnabled = &v
}

// GetTlsEnabled returns the TlsEnabled field value if set, zero value otherwise.
func (o *SMTPConfigParams) GetTlsEnabled() bool {
	if o == nil || IsNil(o.TlsEnabled) {
		var ret bool
		return ret
	}
	return *o.TlsEnabled
}

// GetTlsEnabledOk returns a tuple with the TlsEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SMTPConfigParams) GetTlsEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.TlsEnabled) {
		return nil, false
	}
	return o.TlsEnabled, true
}

// HasTlsEnabled returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasTlsEnabled() bool {
	if o != nil && !IsNil(o.TlsEnabled) {
		return true
	}

	return false
}

// SetTlsEnabled gets a reference to the given bool and assigns it to the TlsEnabled field.
func (o *SMTPConfigParams) SetTlsEnabled(v bool) {
	o.TlsEnabled = &v
}

// GetUsername returns the Username field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMTPConfigParams) GetUsername() string {
	if o == nil || IsNil(o.Username.Get()) {
		var ret string
		return ret
	}
	return *o.Username.Get()
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMTPConfigParams) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Username.Get(), o.Username.IsSet()
}

// HasUsername returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasUsername() bool {
	if o != nil && o.Username.IsSet() {
		return true
	}

	return false
}

// SetUsername gets a reference to the given NullableString and assigns it to the Username field.
func (o *SMTPConfigParams) SetUsername(v string) {
	o.Username.Set(&v)
}
// SetUsernameNil sets the value for Username to be an explicit nil
func (o *SMTPConfigParams) SetUsernameNil() {
	o.Username.Set(nil)
}

// UnsetUsername ensures that no value is present for Username, not even an explicit nil
func (o *SMTPConfigParams) UnsetUsername() {
	o.Username.Unset()
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMTPConfigParams) GetPassword() string {
	if o == nil || IsNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMTPConfigParams) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *SMTPConfigParams) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *SMTPConfigParams) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *SMTPConfigParams) UnsetPassword() {
	o.Password.Unset()
}

// GetFromAddress returns the FromAddress field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMTPConfigParams) GetFromAddress() string {
	if o == nil || IsNil(o.FromAddress.Get()) {
		var ret string
		return ret
	}
	return *o.FromAddress.Get()
}

// GetFromAddressOk returns a tuple with the FromAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMTPConfigParams) GetFromAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FromAddress.Get(), o.FromAddress.IsSet()
}

// HasFromAddress returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasFromAddress() bool {
	if o != nil && o.FromAddress.IsSet() {
		return true
	}

	return false
}

// SetFromAddress gets a reference to the given NullableString and assigns it to the FromAddress field.
func (o *SMTPConfigParams) SetFromAddress(v string) {
	o.FromAddress.Set(&v)
}
// SetFromAddressNil sets the value for FromAddress to be an explicit nil
func (o *SMTPConfigParams) SetFromAddressNil() {
	o.FromAddress.Set(nil)
}

// UnsetFromAddress ensures that no value is present for FromAddress, not even an explicit nil
func (o *SMTPConfigParams) UnsetFromAddress() {
	o.FromAddress.Unset()
}

// GetSendTimeout returns the SendTimeout field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SMTPConfigParams) GetSendTimeout() int32 {
	if o == nil || IsNil(o.SendTimeout.Get()) {
		var ret int32
		return ret
	}
	return *o.SendTimeout.Get()
}

// GetSendTimeoutOk returns a tuple with the SendTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SMTPConfigParams) GetSendTimeoutOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.SendTimeout.Get(), o.SendTimeout.IsSet()
}

// HasSendTimeout returns a boolean if a field has been set.
func (o *SMTPConfigParams) HasSendTimeout() bool {
	if o != nil && o.SendTimeout.IsSet() {
		return true
	}

	return false
}

// SetSendTimeout gets a reference to the given NullableInt32 and assigns it to the SendTimeout field.
func (o *SMTPConfigParams) SetSendTimeout(v int32) {
	o.SendTimeout.Set(&v)
}
// SetSendTimeoutNil sets the value for SendTimeout to be an explicit nil
func (o *SMTPConfigParams) SetSendTimeoutNil() {
	o.SendTimeout.Set(nil)
}

// UnsetSendTimeout ensures that no value is present for SendTimeout, not even an explicit nil
func (o *SMTPConfigParams) UnsetSendTimeout() {
	o.SendTimeout.Unset()
}

func (o SMTPConfigParams) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SMTPConfigParams) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Server.IsSet() {
		toSerialize["server"] = o.Server.Get()
	}
	if o.Port.IsSet() {
		toSerialize["port"] = o.Port.Get()
	}
	if !IsNil(o.AuthenticationEnabled) {
		toSerialize["authentication_enabled"] = o.AuthenticationEnabled
	}
	if !IsNil(o.TlsEnabled) {
		toSerialize["tls_enabled"] = o.TlsEnabled
	}
	if o.Username.IsSet() {
		toSerialize["username"] = o.Username.Get()
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}
	if o.FromAddress.IsSet() {
		toSerialize["from_address"] = o.FromAddress.Get()
	}
	if o.SendTimeout.IsSet() {
		toSerialize["send_timeout"] = o.SendTimeout.Get()
	}
	return toSerialize, nil
}

type NullableSMTPConfigParams struct {
	value *SMTPConfigParams
	isSet bool
}

func (v NullableSMTPConfigParams) Get() *SMTPConfigParams {
	return v.value
}

func (v *NullableSMTPConfigParams) Set(val *SMTPConfigParams) {
	v.value = val
	v.isSet = true
}

func (v NullableSMTPConfigParams) IsSet() bool {
	return v.isSet
}

func (v *NullableSMTPConfigParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMTPConfigParams(val *SMTPConfigParams) *NullableSMTPConfigParams {
	return &NullableSMTPConfigParams{value: val, isSet: true}
}

func (v NullableSMTPConfigParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMTPConfigParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


