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

// checks if the HostCreateParameters type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HostCreateParameters{}

// HostCreateParameters struct for HostCreateParameters
type HostCreateParameters struct {
	// The name to associate with the host.
	Name *string `json:"name,omitempty"`
	// The hostname or IP address of this host.
	Hostname *string `json:"hostname,omitempty"`
	// The list of host/IP addresses to use for NFS export.
	NfsAddresses []string `json:"nfs_addresses,omitempty"`
	// The port number used to connect to the host via SSH.
	SshPort *int32 `json:"ssh_port,omitempty"`
	// Reference to a profile for escalating user privileges.
	PrivilegeElevationProfileReference *string `json:"privilege_elevation_profile_reference,omitempty"`
	// The lowercase alias to use inside the user managed DSP keystore.
	DspKeystoreAlias *string `json:"dsp_keystore_alias,omitempty"`
	// The password for the user managed DSP keystore.
	DspKeystorePassword *string `json:"dsp_keystore_password,omitempty"`
	// The path to the user managed DSP keystore.
	DspKeystorePath *string `json:"dsp_keystore_path,omitempty"`
	// The password for the user managed DSP truststore.
	DspTruststorePassword *string `json:"dsp_truststore_password,omitempty"`
	// The path to the user managed DSP truststore.
	DspTruststorePath *string `json:"dsp_truststore_path,omitempty"`
	// The path to the user managed Java Development Kit (JDK). If not specified, then the OpenJDK will be used.
	JavaHome *string `json:"java_home,omitempty"`
	// The path for the toolkit that resides on the host.
	ToolkitPath *string `json:"toolkit_path,omitempty"`
	// The password for the user managed Oracle JDBC keystore.
	OracleJdbcKeystorePassword *string `json:"oracle_jdbc_keystore_password,omitempty"`
	// The path to the root of the Oracle TDE keystores artifact directories.
	OracleTdeKeystoresRootPath *string `json:"oracle_tde_keystores_root_path,omitempty"`
	SshVerificationStrategy *SSHVerificationStrategy `json:"ssh_verification_strategy,omitempty"`
	// The Virtual IP addresses associated with the OracleClusterNode.
	OracleClusterNodeVirtualIps []OracleVirtualIP `json:"oracle_cluster_node_virtual_ips,omitempty"`
}

// NewHostCreateParameters instantiates a new HostCreateParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHostCreateParameters() *HostCreateParameters {
	this := HostCreateParameters{}
	var sshPort int32 = 22
	this.SshPort = &sshPort
	return &this
}

// NewHostCreateParametersWithDefaults instantiates a new HostCreateParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHostCreateParametersWithDefaults() *HostCreateParameters {
	this := HostCreateParameters{}
	var sshPort int32 = 22
	this.SshPort = &sshPort
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *HostCreateParameters) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *HostCreateParameters) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *HostCreateParameters) SetName(v string) {
	o.Name = &v
}

// GetHostname returns the Hostname field value if set, zero value otherwise.
func (o *HostCreateParameters) GetHostname() string {
	if o == nil || IsNil(o.Hostname) {
		var ret string
		return ret
	}
	return *o.Hostname
}

// GetHostnameOk returns a tuple with the Hostname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetHostnameOk() (*string, bool) {
	if o == nil || IsNil(o.Hostname) {
		return nil, false
	}
	return o.Hostname, true
}

// HasHostname returns a boolean if a field has been set.
func (o *HostCreateParameters) HasHostname() bool {
	if o != nil && !IsNil(o.Hostname) {
		return true
	}

	return false
}

// SetHostname gets a reference to the given string and assigns it to the Hostname field.
func (o *HostCreateParameters) SetHostname(v string) {
	o.Hostname = &v
}

// GetNfsAddresses returns the NfsAddresses field value if set, zero value otherwise.
func (o *HostCreateParameters) GetNfsAddresses() []string {
	if o == nil || IsNil(o.NfsAddresses) {
		var ret []string
		return ret
	}
	return o.NfsAddresses
}

// GetNfsAddressesOk returns a tuple with the NfsAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetNfsAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.NfsAddresses) {
		return nil, false
	}
	return o.NfsAddresses, true
}

// HasNfsAddresses returns a boolean if a field has been set.
func (o *HostCreateParameters) HasNfsAddresses() bool {
	if o != nil && !IsNil(o.NfsAddresses) {
		return true
	}

	return false
}

// SetNfsAddresses gets a reference to the given []string and assigns it to the NfsAddresses field.
func (o *HostCreateParameters) SetNfsAddresses(v []string) {
	o.NfsAddresses = v
}

// GetSshPort returns the SshPort field value if set, zero value otherwise.
func (o *HostCreateParameters) GetSshPort() int32 {
	if o == nil || IsNil(o.SshPort) {
		var ret int32
		return ret
	}
	return *o.SshPort
}

// GetSshPortOk returns a tuple with the SshPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetSshPortOk() (*int32, bool) {
	if o == nil || IsNil(o.SshPort) {
		return nil, false
	}
	return o.SshPort, true
}

// HasSshPort returns a boolean if a field has been set.
func (o *HostCreateParameters) HasSshPort() bool {
	if o != nil && !IsNil(o.SshPort) {
		return true
	}

	return false
}

// SetSshPort gets a reference to the given int32 and assigns it to the SshPort field.
func (o *HostCreateParameters) SetSshPort(v int32) {
	o.SshPort = &v
}

// GetPrivilegeElevationProfileReference returns the PrivilegeElevationProfileReference field value if set, zero value otherwise.
func (o *HostCreateParameters) GetPrivilegeElevationProfileReference() string {
	if o == nil || IsNil(o.PrivilegeElevationProfileReference) {
		var ret string
		return ret
	}
	return *o.PrivilegeElevationProfileReference
}

// GetPrivilegeElevationProfileReferenceOk returns a tuple with the PrivilegeElevationProfileReference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetPrivilegeElevationProfileReferenceOk() (*string, bool) {
	if o == nil || IsNil(o.PrivilegeElevationProfileReference) {
		return nil, false
	}
	return o.PrivilegeElevationProfileReference, true
}

// HasPrivilegeElevationProfileReference returns a boolean if a field has been set.
func (o *HostCreateParameters) HasPrivilegeElevationProfileReference() bool {
	if o != nil && !IsNil(o.PrivilegeElevationProfileReference) {
		return true
	}

	return false
}

// SetPrivilegeElevationProfileReference gets a reference to the given string and assigns it to the PrivilegeElevationProfileReference field.
func (o *HostCreateParameters) SetPrivilegeElevationProfileReference(v string) {
	o.PrivilegeElevationProfileReference = &v
}

// GetDspKeystoreAlias returns the DspKeystoreAlias field value if set, zero value otherwise.
func (o *HostCreateParameters) GetDspKeystoreAlias() string {
	if o == nil || IsNil(o.DspKeystoreAlias) {
		var ret string
		return ret
	}
	return *o.DspKeystoreAlias
}

// GetDspKeystoreAliasOk returns a tuple with the DspKeystoreAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetDspKeystoreAliasOk() (*string, bool) {
	if o == nil || IsNil(o.DspKeystoreAlias) {
		return nil, false
	}
	return o.DspKeystoreAlias, true
}

// HasDspKeystoreAlias returns a boolean if a field has been set.
func (o *HostCreateParameters) HasDspKeystoreAlias() bool {
	if o != nil && !IsNil(o.DspKeystoreAlias) {
		return true
	}

	return false
}

// SetDspKeystoreAlias gets a reference to the given string and assigns it to the DspKeystoreAlias field.
func (o *HostCreateParameters) SetDspKeystoreAlias(v string) {
	o.DspKeystoreAlias = &v
}

// GetDspKeystorePassword returns the DspKeystorePassword field value if set, zero value otherwise.
func (o *HostCreateParameters) GetDspKeystorePassword() string {
	if o == nil || IsNil(o.DspKeystorePassword) {
		var ret string
		return ret
	}
	return *o.DspKeystorePassword
}

// GetDspKeystorePasswordOk returns a tuple with the DspKeystorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetDspKeystorePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DspKeystorePassword) {
		return nil, false
	}
	return o.DspKeystorePassword, true
}

// HasDspKeystorePassword returns a boolean if a field has been set.
func (o *HostCreateParameters) HasDspKeystorePassword() bool {
	if o != nil && !IsNil(o.DspKeystorePassword) {
		return true
	}

	return false
}

// SetDspKeystorePassword gets a reference to the given string and assigns it to the DspKeystorePassword field.
func (o *HostCreateParameters) SetDspKeystorePassword(v string) {
	o.DspKeystorePassword = &v
}

// GetDspKeystorePath returns the DspKeystorePath field value if set, zero value otherwise.
func (o *HostCreateParameters) GetDspKeystorePath() string {
	if o == nil || IsNil(o.DspKeystorePath) {
		var ret string
		return ret
	}
	return *o.DspKeystorePath
}

// GetDspKeystorePathOk returns a tuple with the DspKeystorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetDspKeystorePathOk() (*string, bool) {
	if o == nil || IsNil(o.DspKeystorePath) {
		return nil, false
	}
	return o.DspKeystorePath, true
}

// HasDspKeystorePath returns a boolean if a field has been set.
func (o *HostCreateParameters) HasDspKeystorePath() bool {
	if o != nil && !IsNil(o.DspKeystorePath) {
		return true
	}

	return false
}

// SetDspKeystorePath gets a reference to the given string and assigns it to the DspKeystorePath field.
func (o *HostCreateParameters) SetDspKeystorePath(v string) {
	o.DspKeystorePath = &v
}

// GetDspTruststorePassword returns the DspTruststorePassword field value if set, zero value otherwise.
func (o *HostCreateParameters) GetDspTruststorePassword() string {
	if o == nil || IsNil(o.DspTruststorePassword) {
		var ret string
		return ret
	}
	return *o.DspTruststorePassword
}

// GetDspTruststorePasswordOk returns a tuple with the DspTruststorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetDspTruststorePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DspTruststorePassword) {
		return nil, false
	}
	return o.DspTruststorePassword, true
}

// HasDspTruststorePassword returns a boolean if a field has been set.
func (o *HostCreateParameters) HasDspTruststorePassword() bool {
	if o != nil && !IsNil(o.DspTruststorePassword) {
		return true
	}

	return false
}

// SetDspTruststorePassword gets a reference to the given string and assigns it to the DspTruststorePassword field.
func (o *HostCreateParameters) SetDspTruststorePassword(v string) {
	o.DspTruststorePassword = &v
}

// GetDspTruststorePath returns the DspTruststorePath field value if set, zero value otherwise.
func (o *HostCreateParameters) GetDspTruststorePath() string {
	if o == nil || IsNil(o.DspTruststorePath) {
		var ret string
		return ret
	}
	return *o.DspTruststorePath
}

// GetDspTruststorePathOk returns a tuple with the DspTruststorePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetDspTruststorePathOk() (*string, bool) {
	if o == nil || IsNil(o.DspTruststorePath) {
		return nil, false
	}
	return o.DspTruststorePath, true
}

// HasDspTruststorePath returns a boolean if a field has been set.
func (o *HostCreateParameters) HasDspTruststorePath() bool {
	if o != nil && !IsNil(o.DspTruststorePath) {
		return true
	}

	return false
}

// SetDspTruststorePath gets a reference to the given string and assigns it to the DspTruststorePath field.
func (o *HostCreateParameters) SetDspTruststorePath(v string) {
	o.DspTruststorePath = &v
}

// GetJavaHome returns the JavaHome field value if set, zero value otherwise.
func (o *HostCreateParameters) GetJavaHome() string {
	if o == nil || IsNil(o.JavaHome) {
		var ret string
		return ret
	}
	return *o.JavaHome
}

// GetJavaHomeOk returns a tuple with the JavaHome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetJavaHomeOk() (*string, bool) {
	if o == nil || IsNil(o.JavaHome) {
		return nil, false
	}
	return o.JavaHome, true
}

// HasJavaHome returns a boolean if a field has been set.
func (o *HostCreateParameters) HasJavaHome() bool {
	if o != nil && !IsNil(o.JavaHome) {
		return true
	}

	return false
}

// SetJavaHome gets a reference to the given string and assigns it to the JavaHome field.
func (o *HostCreateParameters) SetJavaHome(v string) {
	o.JavaHome = &v
}

// GetToolkitPath returns the ToolkitPath field value if set, zero value otherwise.
func (o *HostCreateParameters) GetToolkitPath() string {
	if o == nil || IsNil(o.ToolkitPath) {
		var ret string
		return ret
	}
	return *o.ToolkitPath
}

// GetToolkitPathOk returns a tuple with the ToolkitPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetToolkitPathOk() (*string, bool) {
	if o == nil || IsNil(o.ToolkitPath) {
		return nil, false
	}
	return o.ToolkitPath, true
}

// HasToolkitPath returns a boolean if a field has been set.
func (o *HostCreateParameters) HasToolkitPath() bool {
	if o != nil && !IsNil(o.ToolkitPath) {
		return true
	}

	return false
}

// SetToolkitPath gets a reference to the given string and assigns it to the ToolkitPath field.
func (o *HostCreateParameters) SetToolkitPath(v string) {
	o.ToolkitPath = &v
}

// GetOracleJdbcKeystorePassword returns the OracleJdbcKeystorePassword field value if set, zero value otherwise.
func (o *HostCreateParameters) GetOracleJdbcKeystorePassword() string {
	if o == nil || IsNil(o.OracleJdbcKeystorePassword) {
		var ret string
		return ret
	}
	return *o.OracleJdbcKeystorePassword
}

// GetOracleJdbcKeystorePasswordOk returns a tuple with the OracleJdbcKeystorePassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetOracleJdbcKeystorePasswordOk() (*string, bool) {
	if o == nil || IsNil(o.OracleJdbcKeystorePassword) {
		return nil, false
	}
	return o.OracleJdbcKeystorePassword, true
}

// HasOracleJdbcKeystorePassword returns a boolean if a field has been set.
func (o *HostCreateParameters) HasOracleJdbcKeystorePassword() bool {
	if o != nil && !IsNil(o.OracleJdbcKeystorePassword) {
		return true
	}

	return false
}

// SetOracleJdbcKeystorePassword gets a reference to the given string and assigns it to the OracleJdbcKeystorePassword field.
func (o *HostCreateParameters) SetOracleJdbcKeystorePassword(v string) {
	o.OracleJdbcKeystorePassword = &v
}

// GetOracleTdeKeystoresRootPath returns the OracleTdeKeystoresRootPath field value if set, zero value otherwise.
func (o *HostCreateParameters) GetOracleTdeKeystoresRootPath() string {
	if o == nil || IsNil(o.OracleTdeKeystoresRootPath) {
		var ret string
		return ret
	}
	return *o.OracleTdeKeystoresRootPath
}

// GetOracleTdeKeystoresRootPathOk returns a tuple with the OracleTdeKeystoresRootPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetOracleTdeKeystoresRootPathOk() (*string, bool) {
	if o == nil || IsNil(o.OracleTdeKeystoresRootPath) {
		return nil, false
	}
	return o.OracleTdeKeystoresRootPath, true
}

// HasOracleTdeKeystoresRootPath returns a boolean if a field has been set.
func (o *HostCreateParameters) HasOracleTdeKeystoresRootPath() bool {
	if o != nil && !IsNil(o.OracleTdeKeystoresRootPath) {
		return true
	}

	return false
}

// SetOracleTdeKeystoresRootPath gets a reference to the given string and assigns it to the OracleTdeKeystoresRootPath field.
func (o *HostCreateParameters) SetOracleTdeKeystoresRootPath(v string) {
	o.OracleTdeKeystoresRootPath = &v
}

// GetSshVerificationStrategy returns the SshVerificationStrategy field value if set, zero value otherwise.
func (o *HostCreateParameters) GetSshVerificationStrategy() SSHVerificationStrategy {
	if o == nil || IsNil(o.SshVerificationStrategy) {
		var ret SSHVerificationStrategy
		return ret
	}
	return *o.SshVerificationStrategy
}

// GetSshVerificationStrategyOk returns a tuple with the SshVerificationStrategy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetSshVerificationStrategyOk() (*SSHVerificationStrategy, bool) {
	if o == nil || IsNil(o.SshVerificationStrategy) {
		return nil, false
	}
	return o.SshVerificationStrategy, true
}

// HasSshVerificationStrategy returns a boolean if a field has been set.
func (o *HostCreateParameters) HasSshVerificationStrategy() bool {
	if o != nil && !IsNil(o.SshVerificationStrategy) {
		return true
	}

	return false
}

// SetSshVerificationStrategy gets a reference to the given SSHVerificationStrategy and assigns it to the SshVerificationStrategy field.
func (o *HostCreateParameters) SetSshVerificationStrategy(v SSHVerificationStrategy) {
	o.SshVerificationStrategy = &v
}

// GetOracleClusterNodeVirtualIps returns the OracleClusterNodeVirtualIps field value if set, zero value otherwise.
func (o *HostCreateParameters) GetOracleClusterNodeVirtualIps() []OracleVirtualIP {
	if o == nil || IsNil(o.OracleClusterNodeVirtualIps) {
		var ret []OracleVirtualIP
		return ret
	}
	return o.OracleClusterNodeVirtualIps
}

// GetOracleClusterNodeVirtualIpsOk returns a tuple with the OracleClusterNodeVirtualIps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HostCreateParameters) GetOracleClusterNodeVirtualIpsOk() ([]OracleVirtualIP, bool) {
	if o == nil || IsNil(o.OracleClusterNodeVirtualIps) {
		return nil, false
	}
	return o.OracleClusterNodeVirtualIps, true
}

// HasOracleClusterNodeVirtualIps returns a boolean if a field has been set.
func (o *HostCreateParameters) HasOracleClusterNodeVirtualIps() bool {
	if o != nil && !IsNil(o.OracleClusterNodeVirtualIps) {
		return true
	}

	return false
}

// SetOracleClusterNodeVirtualIps gets a reference to the given []OracleVirtualIP and assigns it to the OracleClusterNodeVirtualIps field.
func (o *HostCreateParameters) SetOracleClusterNodeVirtualIps(v []OracleVirtualIP) {
	o.OracleClusterNodeVirtualIps = v
}

func (o HostCreateParameters) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HostCreateParameters) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Hostname) {
		toSerialize["hostname"] = o.Hostname
	}
	if !IsNil(o.NfsAddresses) {
		toSerialize["nfs_addresses"] = o.NfsAddresses
	}
	if !IsNil(o.SshPort) {
		toSerialize["ssh_port"] = o.SshPort
	}
	if !IsNil(o.PrivilegeElevationProfileReference) {
		toSerialize["privilege_elevation_profile_reference"] = o.PrivilegeElevationProfileReference
	}
	if !IsNil(o.DspKeystoreAlias) {
		toSerialize["dsp_keystore_alias"] = o.DspKeystoreAlias
	}
	if !IsNil(o.DspKeystorePassword) {
		toSerialize["dsp_keystore_password"] = o.DspKeystorePassword
	}
	if !IsNil(o.DspKeystorePath) {
		toSerialize["dsp_keystore_path"] = o.DspKeystorePath
	}
	if !IsNil(o.DspTruststorePassword) {
		toSerialize["dsp_truststore_password"] = o.DspTruststorePassword
	}
	if !IsNil(o.DspTruststorePath) {
		toSerialize["dsp_truststore_path"] = o.DspTruststorePath
	}
	if !IsNil(o.JavaHome) {
		toSerialize["java_home"] = o.JavaHome
	}
	if !IsNil(o.ToolkitPath) {
		toSerialize["toolkit_path"] = o.ToolkitPath
	}
	if !IsNil(o.OracleJdbcKeystorePassword) {
		toSerialize["oracle_jdbc_keystore_password"] = o.OracleJdbcKeystorePassword
	}
	if !IsNil(o.OracleTdeKeystoresRootPath) {
		toSerialize["oracle_tde_keystores_root_path"] = o.OracleTdeKeystoresRootPath
	}
	if !IsNil(o.SshVerificationStrategy) {
		toSerialize["ssh_verification_strategy"] = o.SshVerificationStrategy
	}
	if !IsNil(o.OracleClusterNodeVirtualIps) {
		toSerialize["oracle_cluster_node_virtual_ips"] = o.OracleClusterNodeVirtualIps
	}
	return toSerialize, nil
}

type NullableHostCreateParameters struct {
	value *HostCreateParameters
	isSet bool
}

func (v NullableHostCreateParameters) Get() *HostCreateParameters {
	return v.value
}

func (v *NullableHostCreateParameters) Set(val *HostCreateParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableHostCreateParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableHostCreateParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHostCreateParameters(val *HostCreateParameters) *NullableHostCreateParameters {
	return &NullableHostCreateParameters{value: val, isSet: true}
}

func (v NullableHostCreateParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHostCreateParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


