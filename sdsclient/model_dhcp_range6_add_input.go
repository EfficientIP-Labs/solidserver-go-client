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

// DhcpRange6AddInput struct for DhcpRange6AddInput
type DhcpRange6AddInput struct {
	// The database identifier (ID) of the DHCPv6 server, a unique numeric key value automatically incremented when you add a DHCPv6 server. Use the ID to specify the DHCPv6 server of your choice.
	Server6Id *int32 `json:"server6_id,omitempty"`
	// The name of the DHCPv6 server.
	Server6Name *string `json:"server6_name,omitempty"`
	// The last IP address of the DHCPv6 range.
	Range6EndAddr *string `json:"range6_end_addr,omitempty"`
	// The first IP address of the DHCPv6 range.
	Range6StartAddr *string `json:"range6_start_addr,omitempty"`
	// The database identifier (ID) of the DHCPv6 scope, a unique numeric key value automatically incremented when you add a DHCPv6 scope. Use the ID to specify the DHCPv6 scope of your choice.
	Scope6Id *int32 `json:"scope6_id,omitempty"`
	// The IP address of the DHCP server.
	Server6Hostaddr *string `json:"server6_hostaddr,omitempty"`
	// The list of ACLs associated with the DHCPv6 range, as follows: <b>&lt;ACL_name&gt;;&lt;ACL_name&gt;;...</b> .
	Range6Acl *string `json:"range6_acl,omitempty"`
	// The name of the DHCPv6 scope.
	Scope6Name *string `json:"scope6_name,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete *[]string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	Range6ClassName *string `json:"range6_class_name,omitempty"`
	// class parameters in json format
	Range6ClassParameters *[]ApiClassParameterInputEntry `json:"range6_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewDhcpRange6AddInput instantiates a new DhcpRange6AddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpRange6AddInput() *DhcpRange6AddInput {
	this := DhcpRange6AddInput{}
	return &this
}

// NewDhcpRange6AddInputWithDefaults instantiates a new DhcpRange6AddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpRange6AddInputWithDefaults() *DhcpRange6AddInput {
	this := DhcpRange6AddInput{}
	return &this
}

// GetServer6Id returns the Server6Id field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetServer6Id() int32 {
	if o == nil || o.Server6Id == nil {
		var ret int32
		return ret
	}
	return *o.Server6Id
}

// GetServer6IdOk returns a tuple with the Server6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetServer6IdOk() (*int32, bool) {
	if o == nil || o.Server6Id == nil {
		return nil, false
	}
	return o.Server6Id, true
}

// HasServer6Id returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasServer6Id() bool {
	if o != nil && o.Server6Id != nil {
		return true
	}

	return false
}

// SetServer6Id gets a reference to the given int32 and assigns it to the Server6Id field.
func (o *DhcpRange6AddInput) SetServer6Id(v int32) {
	o.Server6Id = &v
}

// GetServer6Name returns the Server6Name field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetServer6Name() string {
	if o == nil || o.Server6Name == nil {
		var ret string
		return ret
	}
	return *o.Server6Name
}

// GetServer6NameOk returns a tuple with the Server6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetServer6NameOk() (*string, bool) {
	if o == nil || o.Server6Name == nil {
		return nil, false
	}
	return o.Server6Name, true
}

// HasServer6Name returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasServer6Name() bool {
	if o != nil && o.Server6Name != nil {
		return true
	}

	return false
}

// SetServer6Name gets a reference to the given string and assigns it to the Server6Name field.
func (o *DhcpRange6AddInput) SetServer6Name(v string) {
	o.Server6Name = &v
}

// GetRange6EndAddr returns the Range6EndAddr field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetRange6EndAddr() string {
	if o == nil || o.Range6EndAddr == nil {
		var ret string
		return ret
	}
	return *o.Range6EndAddr
}

// GetRange6EndAddrOk returns a tuple with the Range6EndAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetRange6EndAddrOk() (*string, bool) {
	if o == nil || o.Range6EndAddr == nil {
		return nil, false
	}
	return o.Range6EndAddr, true
}

// HasRange6EndAddr returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasRange6EndAddr() bool {
	if o != nil && o.Range6EndAddr != nil {
		return true
	}

	return false
}

// SetRange6EndAddr gets a reference to the given string and assigns it to the Range6EndAddr field.
func (o *DhcpRange6AddInput) SetRange6EndAddr(v string) {
	o.Range6EndAddr = &v
}

// GetRange6StartAddr returns the Range6StartAddr field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetRange6StartAddr() string {
	if o == nil || o.Range6StartAddr == nil {
		var ret string
		return ret
	}
	return *o.Range6StartAddr
}

// GetRange6StartAddrOk returns a tuple with the Range6StartAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetRange6StartAddrOk() (*string, bool) {
	if o == nil || o.Range6StartAddr == nil {
		return nil, false
	}
	return o.Range6StartAddr, true
}

// HasRange6StartAddr returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasRange6StartAddr() bool {
	if o != nil && o.Range6StartAddr != nil {
		return true
	}

	return false
}

// SetRange6StartAddr gets a reference to the given string and assigns it to the Range6StartAddr field.
func (o *DhcpRange6AddInput) SetRange6StartAddr(v string) {
	o.Range6StartAddr = &v
}

// GetScope6Id returns the Scope6Id field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetScope6Id() int32 {
	if o == nil || o.Scope6Id == nil {
		var ret int32
		return ret
	}
	return *o.Scope6Id
}

// GetScope6IdOk returns a tuple with the Scope6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetScope6IdOk() (*int32, bool) {
	if o == nil || o.Scope6Id == nil {
		return nil, false
	}
	return o.Scope6Id, true
}

// HasScope6Id returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasScope6Id() bool {
	if o != nil && o.Scope6Id != nil {
		return true
	}

	return false
}

// SetScope6Id gets a reference to the given int32 and assigns it to the Scope6Id field.
func (o *DhcpRange6AddInput) SetScope6Id(v int32) {
	o.Scope6Id = &v
}

// GetServer6Hostaddr returns the Server6Hostaddr field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetServer6Hostaddr() string {
	if o == nil || o.Server6Hostaddr == nil {
		var ret string
		return ret
	}
	return *o.Server6Hostaddr
}

// GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetServer6HostaddrOk() (*string, bool) {
	if o == nil || o.Server6Hostaddr == nil {
		return nil, false
	}
	return o.Server6Hostaddr, true
}

// HasServer6Hostaddr returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasServer6Hostaddr() bool {
	if o != nil && o.Server6Hostaddr != nil {
		return true
	}

	return false
}

// SetServer6Hostaddr gets a reference to the given string and assigns it to the Server6Hostaddr field.
func (o *DhcpRange6AddInput) SetServer6Hostaddr(v string) {
	o.Server6Hostaddr = &v
}

// GetRange6Acl returns the Range6Acl field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetRange6Acl() string {
	if o == nil || o.Range6Acl == nil {
		var ret string
		return ret
	}
	return *o.Range6Acl
}

// GetRange6AclOk returns a tuple with the Range6Acl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetRange6AclOk() (*string, bool) {
	if o == nil || o.Range6Acl == nil {
		return nil, false
	}
	return o.Range6Acl, true
}

// HasRange6Acl returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasRange6Acl() bool {
	if o != nil && o.Range6Acl != nil {
		return true
	}

	return false
}

// SetRange6Acl gets a reference to the given string and assigns it to the Range6Acl field.
func (o *DhcpRange6AddInput) SetRange6Acl(v string) {
	o.Range6Acl = &v
}

// GetScope6Name returns the Scope6Name field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetScope6Name() string {
	if o == nil || o.Scope6Name == nil {
		var ret string
		return ret
	}
	return *o.Scope6Name
}

// GetScope6NameOk returns a tuple with the Scope6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetScope6NameOk() (*string, bool) {
	if o == nil || o.Scope6Name == nil {
		return nil, false
	}
	return o.Scope6Name, true
}

// HasScope6Name returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasScope6Name() bool {
	if o != nil && o.Scope6Name != nil {
		return true
	}

	return false
}

// SetScope6Name gets a reference to the given string and assigns it to the Scope6Name field.
func (o *DhcpRange6AddInput) SetScope6Name(v string) {
	o.Scope6Name = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetClassParametersToDelete() []string {
	if o == nil || o.ClassParametersToDelete == nil {
		var ret []string
		return ret
	}
	return *o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetClassParametersToDeleteOk() (*[]string, bool) {
	if o == nil || o.ClassParametersToDelete == nil {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasClassParametersToDelete() bool {
	if o != nil && o.ClassParametersToDelete != nil {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *DhcpRange6AddInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = &v
}

// GetRange6ClassName returns the Range6ClassName field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetRange6ClassName() string {
	if o == nil || o.Range6ClassName == nil {
		var ret string
		return ret
	}
	return *o.Range6ClassName
}

// GetRange6ClassNameOk returns a tuple with the Range6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetRange6ClassNameOk() (*string, bool) {
	if o == nil || o.Range6ClassName == nil {
		return nil, false
	}
	return o.Range6ClassName, true
}

// HasRange6ClassName returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasRange6ClassName() bool {
	if o != nil && o.Range6ClassName != nil {
		return true
	}

	return false
}

// SetRange6ClassName gets a reference to the given string and assigns it to the Range6ClassName field.
func (o *DhcpRange6AddInput) SetRange6ClassName(v string) {
	o.Range6ClassName = &v
}

// GetRange6ClassParameters returns the Range6ClassParameters field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetRange6ClassParameters() []ApiClassParameterInputEntry {
	if o == nil || o.Range6ClassParameters == nil {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return *o.Range6ClassParameters
}

// GetRange6ClassParametersOk returns a tuple with the Range6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetRange6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool) {
	if o == nil || o.Range6ClassParameters == nil {
		return nil, false
	}
	return o.Range6ClassParameters, true
}

// HasRange6ClassParameters returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasRange6ClassParameters() bool {
	if o != nil && o.Range6ClassParameters != nil {
		return true
	}

	return false
}

// SetRange6ClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the Range6ClassParameters field.
func (o *DhcpRange6AddInput) SetRange6ClassParameters(v []ApiClassParameterInputEntry) {
	o.Range6ClassParameters = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DhcpRange6AddInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRange6AddInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DhcpRange6AddInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *DhcpRange6AddInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o DhcpRange6AddInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Server6Id != nil {
		toSerialize["server6_id"] = o.Server6Id
	}
	if o.Server6Name != nil {
		toSerialize["server6_name"] = o.Server6Name
	}
	if o.Range6EndAddr != nil {
		toSerialize["range6_end_addr"] = o.Range6EndAddr
	}
	if o.Range6StartAddr != nil {
		toSerialize["range6_start_addr"] = o.Range6StartAddr
	}
	if o.Scope6Id != nil {
		toSerialize["scope6_id"] = o.Scope6Id
	}
	if o.Server6Hostaddr != nil {
		toSerialize["server6_hostaddr"] = o.Server6Hostaddr
	}
	if o.Range6Acl != nil {
		toSerialize["range6_acl"] = o.Range6Acl
	}
	if o.Scope6Name != nil {
		toSerialize["scope6_name"] = o.Scope6Name
	}
	if o.ClassParametersToDelete != nil {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if o.Range6ClassName != nil {
		toSerialize["range6_class_name"] = o.Range6ClassName
	}
	if o.Range6ClassParameters != nil {
		toSerialize["range6_class_parameters"] = o.Range6ClassParameters
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpRange6AddInput struct {
	value *DhcpRange6AddInput
	isSet bool
}

func (v NullableDhcpRange6AddInput) Get() *DhcpRange6AddInput {
	return v.value
}

func (v *NullableDhcpRange6AddInput) Set(val *DhcpRange6AddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpRange6AddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpRange6AddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpRange6AddInput(val *DhcpRange6AddInput) *NullableDhcpRange6AddInput {
	return &NullableDhcpRange6AddInput{value: val, isSet: true}
}

func (v NullableDhcpRange6AddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpRange6AddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


