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

// DhcpStatic6EditInput struct for DhcpStatic6EditInput
type DhcpStatic6EditInput struct {
	// The database identifier (ID) of the DHCPv6 server, a unique numeric key value automatically incremented when you add a DHCPv6 server. Use the ID to specify the DHCPv6 server of your choice.
	Server6Id *int32 `json:"server6_id,omitempty"`
	// The name of the DHCPv6 server.
	Server6Name *string `json:"server6_name,omitempty"`
	// The client DHCP Unique Identifier (DUID) associated with the DHCPv6 static.
	Static6ClientDuid *string `json:"static6_client_duid,omitempty"`
	// The database identifier (ID) of the DHCPv6 static, a unique numeric key value automatically incremented when you add a DHCPv6 static. Use the ID to specify which DHCPv6 static to edit.
	Static6Id *int32 `json:"static6_id,omitempty"`
	// The MAC address you want to associate with the IPv6 static.
	Static6MacAddr *string `json:"static6_mac_addr,omitempty"`
	// The name of the DHCPv6 static, each DHCPv6 static must have a unique name.
	Static6Name *string `json:"static6_name,omitempty"`
	// The IP address of the DHCP server.
	Server6Hostaddr *string `json:"server6_hostaddr,omitempty"`
	// The database identifier (ID) of the DHCPv6 group, a unique numeric key value automatically incremented when you add a DHCPv6 group. Use the ID to specify the DHCPv6 group of your choice.
	Group6Id *int32 `json:"group6_id,omitempty"`
	// The name of the DHCPv6 group.
	Group6Name *string `json:"group6_name,omitempty"`
	// The IP address associated with the DHCPv6 static.
	Static6Addr *string `json:"static6_addr,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete *[]string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	Static6ClassName *string `json:"static6_class_name,omitempty"`
	// class parameters in json format
	Static6ClassParameters *[]ApiClassParameterInputEntry `json:"static6_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewDhcpStatic6EditInput instantiates a new DhcpStatic6EditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpStatic6EditInput() *DhcpStatic6EditInput {
	this := DhcpStatic6EditInput{}
	return &this
}

// NewDhcpStatic6EditInputWithDefaults instantiates a new DhcpStatic6EditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpStatic6EditInputWithDefaults() *DhcpStatic6EditInput {
	this := DhcpStatic6EditInput{}
	return &this
}

// GetServer6Id returns the Server6Id field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetServer6Id() int32 {
	if o == nil || o.Server6Id == nil {
		var ret int32
		return ret
	}
	return *o.Server6Id
}

// GetServer6IdOk returns a tuple with the Server6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetServer6IdOk() (*int32, bool) {
	if o == nil || o.Server6Id == nil {
		return nil, false
	}
	return o.Server6Id, true
}

// HasServer6Id returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasServer6Id() bool {
	if o != nil && o.Server6Id != nil {
		return true
	}

	return false
}

// SetServer6Id gets a reference to the given int32 and assigns it to the Server6Id field.
func (o *DhcpStatic6EditInput) SetServer6Id(v int32) {
	o.Server6Id = &v
}

// GetServer6Name returns the Server6Name field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetServer6Name() string {
	if o == nil || o.Server6Name == nil {
		var ret string
		return ret
	}
	return *o.Server6Name
}

// GetServer6NameOk returns a tuple with the Server6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetServer6NameOk() (*string, bool) {
	if o == nil || o.Server6Name == nil {
		return nil, false
	}
	return o.Server6Name, true
}

// HasServer6Name returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasServer6Name() bool {
	if o != nil && o.Server6Name != nil {
		return true
	}

	return false
}

// SetServer6Name gets a reference to the given string and assigns it to the Server6Name field.
func (o *DhcpStatic6EditInput) SetServer6Name(v string) {
	o.Server6Name = &v
}

// GetStatic6ClientDuid returns the Static6ClientDuid field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetStatic6ClientDuid() string {
	if o == nil || o.Static6ClientDuid == nil {
		var ret string
		return ret
	}
	return *o.Static6ClientDuid
}

// GetStatic6ClientDuidOk returns a tuple with the Static6ClientDuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetStatic6ClientDuidOk() (*string, bool) {
	if o == nil || o.Static6ClientDuid == nil {
		return nil, false
	}
	return o.Static6ClientDuid, true
}

// HasStatic6ClientDuid returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasStatic6ClientDuid() bool {
	if o != nil && o.Static6ClientDuid != nil {
		return true
	}

	return false
}

// SetStatic6ClientDuid gets a reference to the given string and assigns it to the Static6ClientDuid field.
func (o *DhcpStatic6EditInput) SetStatic6ClientDuid(v string) {
	o.Static6ClientDuid = &v
}

// GetStatic6Id returns the Static6Id field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetStatic6Id() int32 {
	if o == nil || o.Static6Id == nil {
		var ret int32
		return ret
	}
	return *o.Static6Id
}

// GetStatic6IdOk returns a tuple with the Static6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetStatic6IdOk() (*int32, bool) {
	if o == nil || o.Static6Id == nil {
		return nil, false
	}
	return o.Static6Id, true
}

// HasStatic6Id returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasStatic6Id() bool {
	if o != nil && o.Static6Id != nil {
		return true
	}

	return false
}

// SetStatic6Id gets a reference to the given int32 and assigns it to the Static6Id field.
func (o *DhcpStatic6EditInput) SetStatic6Id(v int32) {
	o.Static6Id = &v
}

// GetStatic6MacAddr returns the Static6MacAddr field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetStatic6MacAddr() string {
	if o == nil || o.Static6MacAddr == nil {
		var ret string
		return ret
	}
	return *o.Static6MacAddr
}

// GetStatic6MacAddrOk returns a tuple with the Static6MacAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetStatic6MacAddrOk() (*string, bool) {
	if o == nil || o.Static6MacAddr == nil {
		return nil, false
	}
	return o.Static6MacAddr, true
}

// HasStatic6MacAddr returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasStatic6MacAddr() bool {
	if o != nil && o.Static6MacAddr != nil {
		return true
	}

	return false
}

// SetStatic6MacAddr gets a reference to the given string and assigns it to the Static6MacAddr field.
func (o *DhcpStatic6EditInput) SetStatic6MacAddr(v string) {
	o.Static6MacAddr = &v
}

// GetStatic6Name returns the Static6Name field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetStatic6Name() string {
	if o == nil || o.Static6Name == nil {
		var ret string
		return ret
	}
	return *o.Static6Name
}

// GetStatic6NameOk returns a tuple with the Static6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetStatic6NameOk() (*string, bool) {
	if o == nil || o.Static6Name == nil {
		return nil, false
	}
	return o.Static6Name, true
}

// HasStatic6Name returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasStatic6Name() bool {
	if o != nil && o.Static6Name != nil {
		return true
	}

	return false
}

// SetStatic6Name gets a reference to the given string and assigns it to the Static6Name field.
func (o *DhcpStatic6EditInput) SetStatic6Name(v string) {
	o.Static6Name = &v
}

// GetServer6Hostaddr returns the Server6Hostaddr field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetServer6Hostaddr() string {
	if o == nil || o.Server6Hostaddr == nil {
		var ret string
		return ret
	}
	return *o.Server6Hostaddr
}

// GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetServer6HostaddrOk() (*string, bool) {
	if o == nil || o.Server6Hostaddr == nil {
		return nil, false
	}
	return o.Server6Hostaddr, true
}

// HasServer6Hostaddr returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasServer6Hostaddr() bool {
	if o != nil && o.Server6Hostaddr != nil {
		return true
	}

	return false
}

// SetServer6Hostaddr gets a reference to the given string and assigns it to the Server6Hostaddr field.
func (o *DhcpStatic6EditInput) SetServer6Hostaddr(v string) {
	o.Server6Hostaddr = &v
}

// GetGroup6Id returns the Group6Id field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetGroup6Id() int32 {
	if o == nil || o.Group6Id == nil {
		var ret int32
		return ret
	}
	return *o.Group6Id
}

// GetGroup6IdOk returns a tuple with the Group6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetGroup6IdOk() (*int32, bool) {
	if o == nil || o.Group6Id == nil {
		return nil, false
	}
	return o.Group6Id, true
}

// HasGroup6Id returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasGroup6Id() bool {
	if o != nil && o.Group6Id != nil {
		return true
	}

	return false
}

// SetGroup6Id gets a reference to the given int32 and assigns it to the Group6Id field.
func (o *DhcpStatic6EditInput) SetGroup6Id(v int32) {
	o.Group6Id = &v
}

// GetGroup6Name returns the Group6Name field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetGroup6Name() string {
	if o == nil || o.Group6Name == nil {
		var ret string
		return ret
	}
	return *o.Group6Name
}

// GetGroup6NameOk returns a tuple with the Group6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetGroup6NameOk() (*string, bool) {
	if o == nil || o.Group6Name == nil {
		return nil, false
	}
	return o.Group6Name, true
}

// HasGroup6Name returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasGroup6Name() bool {
	if o != nil && o.Group6Name != nil {
		return true
	}

	return false
}

// SetGroup6Name gets a reference to the given string and assigns it to the Group6Name field.
func (o *DhcpStatic6EditInput) SetGroup6Name(v string) {
	o.Group6Name = &v
}

// GetStatic6Addr returns the Static6Addr field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetStatic6Addr() string {
	if o == nil || o.Static6Addr == nil {
		var ret string
		return ret
	}
	return *o.Static6Addr
}

// GetStatic6AddrOk returns a tuple with the Static6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetStatic6AddrOk() (*string, bool) {
	if o == nil || o.Static6Addr == nil {
		return nil, false
	}
	return o.Static6Addr, true
}

// HasStatic6Addr returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasStatic6Addr() bool {
	if o != nil && o.Static6Addr != nil {
		return true
	}

	return false
}

// SetStatic6Addr gets a reference to the given string and assigns it to the Static6Addr field.
func (o *DhcpStatic6EditInput) SetStatic6Addr(v string) {
	o.Static6Addr = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetClassParametersToDelete() []string {
	if o == nil || o.ClassParametersToDelete == nil {
		var ret []string
		return ret
	}
	return *o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetClassParametersToDeleteOk() (*[]string, bool) {
	if o == nil || o.ClassParametersToDelete == nil {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasClassParametersToDelete() bool {
	if o != nil && o.ClassParametersToDelete != nil {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *DhcpStatic6EditInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = &v
}

// GetStatic6ClassName returns the Static6ClassName field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetStatic6ClassName() string {
	if o == nil || o.Static6ClassName == nil {
		var ret string
		return ret
	}
	return *o.Static6ClassName
}

// GetStatic6ClassNameOk returns a tuple with the Static6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetStatic6ClassNameOk() (*string, bool) {
	if o == nil || o.Static6ClassName == nil {
		return nil, false
	}
	return o.Static6ClassName, true
}

// HasStatic6ClassName returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasStatic6ClassName() bool {
	if o != nil && o.Static6ClassName != nil {
		return true
	}

	return false
}

// SetStatic6ClassName gets a reference to the given string and assigns it to the Static6ClassName field.
func (o *DhcpStatic6EditInput) SetStatic6ClassName(v string) {
	o.Static6ClassName = &v
}

// GetStatic6ClassParameters returns the Static6ClassParameters field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetStatic6ClassParameters() []ApiClassParameterInputEntry {
	if o == nil || o.Static6ClassParameters == nil {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return *o.Static6ClassParameters
}

// GetStatic6ClassParametersOk returns a tuple with the Static6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetStatic6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool) {
	if o == nil || o.Static6ClassParameters == nil {
		return nil, false
	}
	return o.Static6ClassParameters, true
}

// HasStatic6ClassParameters returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasStatic6ClassParameters() bool {
	if o != nil && o.Static6ClassParameters != nil {
		return true
	}

	return false
}

// SetStatic6ClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the Static6ClassParameters field.
func (o *DhcpStatic6EditInput) SetStatic6ClassParameters(v []ApiClassParameterInputEntry) {
	o.Static6ClassParameters = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DhcpStatic6EditInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6EditInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DhcpStatic6EditInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *DhcpStatic6EditInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o DhcpStatic6EditInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Server6Id != nil {
		toSerialize["server6_id"] = o.Server6Id
	}
	if o.Server6Name != nil {
		toSerialize["server6_name"] = o.Server6Name
	}
	if o.Static6ClientDuid != nil {
		toSerialize["static6_client_duid"] = o.Static6ClientDuid
	}
	if o.Static6Id != nil {
		toSerialize["static6_id"] = o.Static6Id
	}
	if o.Static6MacAddr != nil {
		toSerialize["static6_mac_addr"] = o.Static6MacAddr
	}
	if o.Static6Name != nil {
		toSerialize["static6_name"] = o.Static6Name
	}
	if o.Server6Hostaddr != nil {
		toSerialize["server6_hostaddr"] = o.Server6Hostaddr
	}
	if o.Group6Id != nil {
		toSerialize["group6_id"] = o.Group6Id
	}
	if o.Group6Name != nil {
		toSerialize["group6_name"] = o.Group6Name
	}
	if o.Static6Addr != nil {
		toSerialize["static6_addr"] = o.Static6Addr
	}
	if o.ClassParametersToDelete != nil {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if o.Static6ClassName != nil {
		toSerialize["static6_class_name"] = o.Static6ClassName
	}
	if o.Static6ClassParameters != nil {
		toSerialize["static6_class_parameters"] = o.Static6ClassParameters
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpStatic6EditInput struct {
	value *DhcpStatic6EditInput
	isSet bool
}

func (v NullableDhcpStatic6EditInput) Get() *DhcpStatic6EditInput {
	return v.value
}

func (v *NullableDhcpStatic6EditInput) Set(val *DhcpStatic6EditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpStatic6EditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpStatic6EditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpStatic6EditInput(val *DhcpStatic6EditInput) *NullableDhcpStatic6EditInput {
	return &NullableDhcpStatic6EditInput{value: val, isSet: true}
}

func (v NullableDhcpStatic6EditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpStatic6EditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


