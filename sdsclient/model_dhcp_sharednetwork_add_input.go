/*
SOLIDserver API

OpenAPI 3.0.2 API definition for SOLIDserver service from EfficientIP.<p>Copyright © 2000-2024 EfficientIP</p><p><em>All specifications and information regarding the products in this document are subject to change without notice and should not be construed as a commitment by EfficientIP. EfficientIP assumes no responsibility or liability for any mistakes or inaccuracies that may appear in this document. All statements and recommendations in this document are believed to be accurate but are presented without warranty. Users must take full responsibility for their application of any product.</em></p><p><em>This document aims at detailing EfficientIP proprietary solutions. As our solutions rely on several third-party products, created by other companies or organizations, it may redirect readers to third-party websites and documentation for further information. In such a case, EfficientIP cannot be liable or expected to provide said information on products they do maintain or created.</em></p><p>Generated (Friday 4th of October 2024 03:41:11 PM)</p>

API version: 2.0
Contact: support-api@efficientip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdsclient

import (
	"encoding/json"
)

// checks if the DhcpSharednetworkAddInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpSharednetworkAddInput{}

// DhcpSharednetworkAddInput struct for DhcpSharednetworkAddInput
type DhcpSharednetworkAddInput struct {
	// The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. Use the ID to specify the DHCPv4 server of your choice.
	ServerId *int32 `json:"server_id,omitempty"`
	// The name of the DHCPv4 server.
	ServerName *string `json:"server_name,omitempty"`
	// The name of the DHCPv4 shared network, each DHCPv4 shared network must have a unique name.
	SharednetworkName *string `json:"sharednetwork_name,omitempty"`
	// The IP address of the DHCP server.
	ServerHostaddr *string `json:"server_hostaddr,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewDhcpSharednetworkAddInput instantiates a new DhcpSharednetworkAddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSharednetworkAddInput() *DhcpSharednetworkAddInput {
	this := DhcpSharednetworkAddInput{}
	return &this
}

// NewDhcpSharednetworkAddInputWithDefaults instantiates a new DhcpSharednetworkAddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSharednetworkAddInputWithDefaults() *DhcpSharednetworkAddInput {
	this := DhcpSharednetworkAddInput{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *DhcpSharednetworkAddInput) GetServerId() int32 {
	if o == nil || IsNil(o.ServerId) {
		var ret int32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkAddInput) GetServerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *DhcpSharednetworkAddInput) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int32 and assigns it to the ServerId field.
func (o *DhcpSharednetworkAddInput) SetServerId(v int32) {
	o.ServerId = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *DhcpSharednetworkAddInput) GetServerName() string {
	if o == nil || IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkAddInput) GetServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *DhcpSharednetworkAddInput) HasServerName() bool {
	if o != nil && !IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *DhcpSharednetworkAddInput) SetServerName(v string) {
	o.ServerName = &v
}

// GetSharednetworkName returns the SharednetworkName field value if set, zero value otherwise.
func (o *DhcpSharednetworkAddInput) GetSharednetworkName() string {
	if o == nil || IsNil(o.SharednetworkName) {
		var ret string
		return ret
	}
	return *o.SharednetworkName
}

// GetSharednetworkNameOk returns a tuple with the SharednetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkAddInput) GetSharednetworkNameOk() (*string, bool) {
	if o == nil || IsNil(o.SharednetworkName) {
		return nil, false
	}
	return o.SharednetworkName, true
}

// HasSharednetworkName returns a boolean if a field has been set.
func (o *DhcpSharednetworkAddInput) HasSharednetworkName() bool {
	if o != nil && !IsNil(o.SharednetworkName) {
		return true
	}

	return false
}

// SetSharednetworkName gets a reference to the given string and assigns it to the SharednetworkName field.
func (o *DhcpSharednetworkAddInput) SetSharednetworkName(v string) {
	o.SharednetworkName = &v
}

// GetServerHostaddr returns the ServerHostaddr field value if set, zero value otherwise.
func (o *DhcpSharednetworkAddInput) GetServerHostaddr() string {
	if o == nil || IsNil(o.ServerHostaddr) {
		var ret string
		return ret
	}
	return *o.ServerHostaddr
}

// GetServerHostaddrOk returns a tuple with the ServerHostaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkAddInput) GetServerHostaddrOk() (*string, bool) {
	if o == nil || IsNil(o.ServerHostaddr) {
		return nil, false
	}
	return o.ServerHostaddr, true
}

// HasServerHostaddr returns a boolean if a field has been set.
func (o *DhcpSharednetworkAddInput) HasServerHostaddr() bool {
	if o != nil && !IsNil(o.ServerHostaddr) {
		return true
	}

	return false
}

// SetServerHostaddr gets a reference to the given string and assigns it to the ServerHostaddr field.
func (o *DhcpSharednetworkAddInput) SetServerHostaddr(v string) {
	o.ServerHostaddr = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DhcpSharednetworkAddInput) GetWarnings() string {
	if o == nil || IsNil(o.Warnings) {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkAddInput) GetWarningsOk() (*string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DhcpSharednetworkAddInput) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *DhcpSharednetworkAddInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o DhcpSharednetworkAddInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpSharednetworkAddInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerId) {
		toSerialize["server_id"] = o.ServerId
	}
	if !IsNil(o.ServerName) {
		toSerialize["server_name"] = o.ServerName
	}
	if !IsNil(o.SharednetworkName) {
		toSerialize["sharednetwork_name"] = o.SharednetworkName
	}
	if !IsNil(o.ServerHostaddr) {
		toSerialize["server_hostaddr"] = o.ServerHostaddr
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableDhcpSharednetworkAddInput struct {
	value *DhcpSharednetworkAddInput
	isSet bool
}

func (v NullableDhcpSharednetworkAddInput) Get() *DhcpSharednetworkAddInput {
	return v.value
}

func (v *NullableDhcpSharednetworkAddInput) Set(val *DhcpSharednetworkAddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSharednetworkAddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSharednetworkAddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSharednetworkAddInput(val *DhcpSharednetworkAddInput) *NullableDhcpSharednetworkAddInput {
	return &NullableDhcpSharednetworkAddInput{value: val, isSet: true}
}

func (v NullableDhcpSharednetworkAddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSharednetworkAddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
