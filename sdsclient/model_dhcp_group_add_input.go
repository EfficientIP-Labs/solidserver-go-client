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

// DhcpGroupAddInput struct for DhcpGroupAddInput
type DhcpGroupAddInput struct {
	// The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. Use the ID to specify the DHCPv4 server of your choice.
	ServerId *int32 `json:"server_id,omitempty"`
	// The name of the DHCPv4 server.
	ServerName *string `json:"server_name,omitempty"`
	// The name of the DHCPv4 group, each DHCPv4 group must have a unique name.
	GroupName *string `json:"group_name,omitempty"`
	// The IP address of the DHCP server.
	ServerHostaddr *string `json:"server_hostaddr,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete *[]string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	GroupClassName *string `json:"group_class_name,omitempty"`
	// class parameters in json format
	GroupClassParameters *[]ApiClassParameterInputEntry `json:"group_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewDhcpGroupAddInput instantiates a new DhcpGroupAddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpGroupAddInput() *DhcpGroupAddInput {
	this := DhcpGroupAddInput{}
	return &this
}

// NewDhcpGroupAddInputWithDefaults instantiates a new DhcpGroupAddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpGroupAddInputWithDefaults() *DhcpGroupAddInput {
	this := DhcpGroupAddInput{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *DhcpGroupAddInput) GetServerId() int32 {
	if o == nil || o.ServerId == nil {
		var ret int32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroupAddInput) GetServerIdOk() (*int32, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *DhcpGroupAddInput) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int32 and assigns it to the ServerId field.
func (o *DhcpGroupAddInput) SetServerId(v int32) {
	o.ServerId = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *DhcpGroupAddInput) GetServerName() string {
	if o == nil || o.ServerName == nil {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroupAddInput) GetServerNameOk() (*string, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *DhcpGroupAddInput) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *DhcpGroupAddInput) SetServerName(v string) {
	o.ServerName = &v
}

// GetGroupName returns the GroupName field value if set, zero value otherwise.
func (o *DhcpGroupAddInput) GetGroupName() string {
	if o == nil || o.GroupName == nil {
		var ret string
		return ret
	}
	return *o.GroupName
}

// GetGroupNameOk returns a tuple with the GroupName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroupAddInput) GetGroupNameOk() (*string, bool) {
	if o == nil || o.GroupName == nil {
		return nil, false
	}
	return o.GroupName, true
}

// HasGroupName returns a boolean if a field has been set.
func (o *DhcpGroupAddInput) HasGroupName() bool {
	if o != nil && o.GroupName != nil {
		return true
	}

	return false
}

// SetGroupName gets a reference to the given string and assigns it to the GroupName field.
func (o *DhcpGroupAddInput) SetGroupName(v string) {
	o.GroupName = &v
}

// GetServerHostaddr returns the ServerHostaddr field value if set, zero value otherwise.
func (o *DhcpGroupAddInput) GetServerHostaddr() string {
	if o == nil || o.ServerHostaddr == nil {
		var ret string
		return ret
	}
	return *o.ServerHostaddr
}

// GetServerHostaddrOk returns a tuple with the ServerHostaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroupAddInput) GetServerHostaddrOk() (*string, bool) {
	if o == nil || o.ServerHostaddr == nil {
		return nil, false
	}
	return o.ServerHostaddr, true
}

// HasServerHostaddr returns a boolean if a field has been set.
func (o *DhcpGroupAddInput) HasServerHostaddr() bool {
	if o != nil && o.ServerHostaddr != nil {
		return true
	}

	return false
}

// SetServerHostaddr gets a reference to the given string and assigns it to the ServerHostaddr field.
func (o *DhcpGroupAddInput) SetServerHostaddr(v string) {
	o.ServerHostaddr = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *DhcpGroupAddInput) GetClassParametersToDelete() []string {
	if o == nil || o.ClassParametersToDelete == nil {
		var ret []string
		return ret
	}
	return *o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroupAddInput) GetClassParametersToDeleteOk() (*[]string, bool) {
	if o == nil || o.ClassParametersToDelete == nil {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *DhcpGroupAddInput) HasClassParametersToDelete() bool {
	if o != nil && o.ClassParametersToDelete != nil {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *DhcpGroupAddInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = &v
}

// GetGroupClassName returns the GroupClassName field value if set, zero value otherwise.
func (o *DhcpGroupAddInput) GetGroupClassName() string {
	if o == nil || o.GroupClassName == nil {
		var ret string
		return ret
	}
	return *o.GroupClassName
}

// GetGroupClassNameOk returns a tuple with the GroupClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroupAddInput) GetGroupClassNameOk() (*string, bool) {
	if o == nil || o.GroupClassName == nil {
		return nil, false
	}
	return o.GroupClassName, true
}

// HasGroupClassName returns a boolean if a field has been set.
func (o *DhcpGroupAddInput) HasGroupClassName() bool {
	if o != nil && o.GroupClassName != nil {
		return true
	}

	return false
}

// SetGroupClassName gets a reference to the given string and assigns it to the GroupClassName field.
func (o *DhcpGroupAddInput) SetGroupClassName(v string) {
	o.GroupClassName = &v
}

// GetGroupClassParameters returns the GroupClassParameters field value if set, zero value otherwise.
func (o *DhcpGroupAddInput) GetGroupClassParameters() []ApiClassParameterInputEntry {
	if o == nil || o.GroupClassParameters == nil {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return *o.GroupClassParameters
}

// GetGroupClassParametersOk returns a tuple with the GroupClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroupAddInput) GetGroupClassParametersOk() (*[]ApiClassParameterInputEntry, bool) {
	if o == nil || o.GroupClassParameters == nil {
		return nil, false
	}
	return o.GroupClassParameters, true
}

// HasGroupClassParameters returns a boolean if a field has been set.
func (o *DhcpGroupAddInput) HasGroupClassParameters() bool {
	if o != nil && o.GroupClassParameters != nil {
		return true
	}

	return false
}

// SetGroupClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the GroupClassParameters field.
func (o *DhcpGroupAddInput) SetGroupClassParameters(v []ApiClassParameterInputEntry) {
	o.GroupClassParameters = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DhcpGroupAddInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroupAddInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DhcpGroupAddInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *DhcpGroupAddInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o DhcpGroupAddInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServerId != nil {
		toSerialize["server_id"] = o.ServerId
	}
	if o.ServerName != nil {
		toSerialize["server_name"] = o.ServerName
	}
	if o.GroupName != nil {
		toSerialize["group_name"] = o.GroupName
	}
	if o.ServerHostaddr != nil {
		toSerialize["server_hostaddr"] = o.ServerHostaddr
	}
	if o.ClassParametersToDelete != nil {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if o.GroupClassName != nil {
		toSerialize["group_class_name"] = o.GroupClassName
	}
	if o.GroupClassParameters != nil {
		toSerialize["group_class_parameters"] = o.GroupClassParameters
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpGroupAddInput struct {
	value *DhcpGroupAddInput
	isSet bool
}

func (v NullableDhcpGroupAddInput) Get() *DhcpGroupAddInput {
	return v.value
}

func (v *NullableDhcpGroupAddInput) Set(val *DhcpGroupAddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpGroupAddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpGroupAddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpGroupAddInput(val *DhcpGroupAddInput) *NullableDhcpGroupAddInput {
	return &NullableDhcpGroupAddInput{value: val, isSet: true}
}

func (v NullableDhcpGroupAddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpGroupAddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


