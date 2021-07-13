/*
 * SOLIDserver API
 *
 * OpenAPI 3.0.2 API definition for SOLIDserver service from EfficientIP.<p>Copyright © 2000-2021 EfficientIP</p><p><em>All specifications and information regarding the products in  this document are subject to change without notice and should not be  construed as a commitment by EfficientIP. EfficientIP assumes no  responsibility or liability for any mistakes or inaccuracies that may appear  in this document. All statements and recommendations in this document are  believed to be accurate but are presented without warranty. Users must take  full responsibility for their application of any product.</em></p><p>Generated (Monday 14th of June 2021 12:30:34 PM)</p>
 *
 * API version: 2.0
 * Contact: support-api@efficientip.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdsclient

import (
	"encoding/json"
)

// DhcpSharednetwork6EditInput struct for DhcpSharednetwork6EditInput
type DhcpSharednetwork6EditInput struct {
	// The database identifier (ID) of the DHCPv6 server, a unique numeric key value automatically incremented when you add a DHCPv6 server. Use the ID to specify the DHCPv6 server of your choice.
	Server6Id *int32 `json:"server6_id,omitempty"`
	// The name of the DHCPv6 server.
	Server6Name *string `json:"server6_name,omitempty"`
	// The database identifier (ID) of the DHCPv6 shared network, a unique numeric key value automatically incremented when you add a DHCPv6 shared network. Use the ID to specify the DHCPv6 shared network of your choice.
	Sharednetwork6Id *int32 `json:"sharednetwork6_id,omitempty"`
	// The name of the DHCPv6 shared network.
	Sharednetwork6Name *string `json:"sharednetwork6_name,omitempty"`
	// The IP address of the DHCP server.
	Server6Hostaddr *string `json:"server6_hostaddr,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewDhcpSharednetwork6EditInput instantiates a new DhcpSharednetwork6EditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSharednetwork6EditInput() *DhcpSharednetwork6EditInput {
	this := DhcpSharednetwork6EditInput{}
	return &this
}

// NewDhcpSharednetwork6EditInputWithDefaults instantiates a new DhcpSharednetwork6EditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSharednetwork6EditInputWithDefaults() *DhcpSharednetwork6EditInput {
	this := DhcpSharednetwork6EditInput{}
	return &this
}

// GetServer6Id returns the Server6Id field value if set, zero value otherwise.
func (o *DhcpSharednetwork6EditInput) GetServer6Id() int32 {
	if o == nil || o.Server6Id == nil {
		var ret int32
		return ret
	}
	return *o.Server6Id
}

// GetServer6IdOk returns a tuple with the Server6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6EditInput) GetServer6IdOk() (*int32, bool) {
	if o == nil || o.Server6Id == nil {
		return nil, false
	}
	return o.Server6Id, true
}

// HasServer6Id returns a boolean if a field has been set.
func (o *DhcpSharednetwork6EditInput) HasServer6Id() bool {
	if o != nil && o.Server6Id != nil {
		return true
	}

	return false
}

// SetServer6Id gets a reference to the given int32 and assigns it to the Server6Id field.
func (o *DhcpSharednetwork6EditInput) SetServer6Id(v int32) {
	o.Server6Id = &v
}

// GetServer6Name returns the Server6Name field value if set, zero value otherwise.
func (o *DhcpSharednetwork6EditInput) GetServer6Name() string {
	if o == nil || o.Server6Name == nil {
		var ret string
		return ret
	}
	return *o.Server6Name
}

// GetServer6NameOk returns a tuple with the Server6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6EditInput) GetServer6NameOk() (*string, bool) {
	if o == nil || o.Server6Name == nil {
		return nil, false
	}
	return o.Server6Name, true
}

// HasServer6Name returns a boolean if a field has been set.
func (o *DhcpSharednetwork6EditInput) HasServer6Name() bool {
	if o != nil && o.Server6Name != nil {
		return true
	}

	return false
}

// SetServer6Name gets a reference to the given string and assigns it to the Server6Name field.
func (o *DhcpSharednetwork6EditInput) SetServer6Name(v string) {
	o.Server6Name = &v
}

// GetSharednetwork6Id returns the Sharednetwork6Id field value if set, zero value otherwise.
func (o *DhcpSharednetwork6EditInput) GetSharednetwork6Id() int32 {
	if o == nil || o.Sharednetwork6Id == nil {
		var ret int32
		return ret
	}
	return *o.Sharednetwork6Id
}

// GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6EditInput) GetSharednetwork6IdOk() (*int32, bool) {
	if o == nil || o.Sharednetwork6Id == nil {
		return nil, false
	}
	return o.Sharednetwork6Id, true
}

// HasSharednetwork6Id returns a boolean if a field has been set.
func (o *DhcpSharednetwork6EditInput) HasSharednetwork6Id() bool {
	if o != nil && o.Sharednetwork6Id != nil {
		return true
	}

	return false
}

// SetSharednetwork6Id gets a reference to the given int32 and assigns it to the Sharednetwork6Id field.
func (o *DhcpSharednetwork6EditInput) SetSharednetwork6Id(v int32) {
	o.Sharednetwork6Id = &v
}

// GetSharednetwork6Name returns the Sharednetwork6Name field value if set, zero value otherwise.
func (o *DhcpSharednetwork6EditInput) GetSharednetwork6Name() string {
	if o == nil || o.Sharednetwork6Name == nil {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Name
}

// GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6EditInput) GetSharednetwork6NameOk() (*string, bool) {
	if o == nil || o.Sharednetwork6Name == nil {
		return nil, false
	}
	return o.Sharednetwork6Name, true
}

// HasSharednetwork6Name returns a boolean if a field has been set.
func (o *DhcpSharednetwork6EditInput) HasSharednetwork6Name() bool {
	if o != nil && o.Sharednetwork6Name != nil {
		return true
	}

	return false
}

// SetSharednetwork6Name gets a reference to the given string and assigns it to the Sharednetwork6Name field.
func (o *DhcpSharednetwork6EditInput) SetSharednetwork6Name(v string) {
	o.Sharednetwork6Name = &v
}

// GetServer6Hostaddr returns the Server6Hostaddr field value if set, zero value otherwise.
func (o *DhcpSharednetwork6EditInput) GetServer6Hostaddr() string {
	if o == nil || o.Server6Hostaddr == nil {
		var ret string
		return ret
	}
	return *o.Server6Hostaddr
}

// GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6EditInput) GetServer6HostaddrOk() (*string, bool) {
	if o == nil || o.Server6Hostaddr == nil {
		return nil, false
	}
	return o.Server6Hostaddr, true
}

// HasServer6Hostaddr returns a boolean if a field has been set.
func (o *DhcpSharednetwork6EditInput) HasServer6Hostaddr() bool {
	if o != nil && o.Server6Hostaddr != nil {
		return true
	}

	return false
}

// SetServer6Hostaddr gets a reference to the given string and assigns it to the Server6Hostaddr field.
func (o *DhcpSharednetwork6EditInput) SetServer6Hostaddr(v string) {
	o.Server6Hostaddr = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DhcpSharednetwork6EditInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6EditInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DhcpSharednetwork6EditInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *DhcpSharednetwork6EditInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o DhcpSharednetwork6EditInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Server6Id != nil {
		toSerialize["server6_id"] = o.Server6Id
	}
	if o.Server6Name != nil {
		toSerialize["server6_name"] = o.Server6Name
	}
	if o.Sharednetwork6Id != nil {
		toSerialize["sharednetwork6_id"] = o.Sharednetwork6Id
	}
	if o.Sharednetwork6Name != nil {
		toSerialize["sharednetwork6_name"] = o.Sharednetwork6Name
	}
	if o.Server6Hostaddr != nil {
		toSerialize["server6_hostaddr"] = o.Server6Hostaddr
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpSharednetwork6EditInput struct {
	value *DhcpSharednetwork6EditInput
	isSet bool
}

func (v NullableDhcpSharednetwork6EditInput) Get() *DhcpSharednetwork6EditInput {
	return v.value
}

func (v *NullableDhcpSharednetwork6EditInput) Set(val *DhcpSharednetwork6EditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSharednetwork6EditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSharednetwork6EditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSharednetwork6EditInput(val *DhcpSharednetwork6EditInput) *NullableDhcpSharednetwork6EditInput {
	return &NullableDhcpSharednetwork6EditInput{value: val, isSet: true}
}

func (v NullableDhcpSharednetwork6EditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSharednetwork6EditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

