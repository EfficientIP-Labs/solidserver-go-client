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

// IpamAddress6AddInput struct for IpamAddress6AddInput
type IpamAddress6AddInput struct {
	// The IP address.
	Address6Hostaddr *string `json:"address6_hostaddr,omitempty"`
	// The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice.
	SpaceId *int32 `json:"space_id,omitempty"`
	// The name of the space.
	SpaceName *string `json:"space_name,omitempty"`
	// The database identifier (ID) of the Device Manager device you want to associate with the IP address.
	DeviceId *int32 `json:"device_id,omitempty"`
	// The database identifier (ID) of the Device Manager interface you want to associate with the IP address.
	InterfaceId *int32 `json:"interface_id,omitempty"`
	// The MAC address you want to associate with the IPv6 address.
	Address6MacAddr *string `json:"address6_mac_addr,omitempty"`
	// The name of the IPv6 address, each IPv6 address must have a unique name.
	Address6Name *string `json:"address6_name,omitempty"`
	// The database identifier (ID) of the NetChange port you want to associate with the IP address.
	PortId *int32 `json:"port_id,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete *[]string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	Address6ClassName *string `json:"address6_class_name,omitempty"`
	// class parameters in json format
	Address6ClassParameters *[]ApiClassParameterInputEntry `json:"address6_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewIpamAddress6AddInput instantiates a new IpamAddress6AddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamAddress6AddInput() *IpamAddress6AddInput {
	this := IpamAddress6AddInput{}
	return &this
}

// NewIpamAddress6AddInputWithDefaults instantiates a new IpamAddress6AddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamAddress6AddInputWithDefaults() *IpamAddress6AddInput {
	this := IpamAddress6AddInput{}
	return &this
}

// GetAddress6Hostaddr returns the Address6Hostaddr field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetAddress6Hostaddr() string {
	if o == nil || o.Address6Hostaddr == nil {
		var ret string
		return ret
	}
	return *o.Address6Hostaddr
}

// GetAddress6HostaddrOk returns a tuple with the Address6Hostaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetAddress6HostaddrOk() (*string, bool) {
	if o == nil || o.Address6Hostaddr == nil {
		return nil, false
	}
	return o.Address6Hostaddr, true
}

// HasAddress6Hostaddr returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasAddress6Hostaddr() bool {
	if o != nil && o.Address6Hostaddr != nil {
		return true
	}

	return false
}

// SetAddress6Hostaddr gets a reference to the given string and assigns it to the Address6Hostaddr field.
func (o *IpamAddress6AddInput) SetAddress6Hostaddr(v string) {
	o.Address6Hostaddr = &v
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetSpaceId() int32 {
	if o == nil || o.SpaceId == nil {
		var ret int32
		return ret
	}
	return *o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetSpaceIdOk() (*int32, bool) {
	if o == nil || o.SpaceId == nil {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasSpaceId() bool {
	if o != nil && o.SpaceId != nil {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given int32 and assigns it to the SpaceId field.
func (o *IpamAddress6AddInput) SetSpaceId(v int32) {
	o.SpaceId = &v
}

// GetSpaceName returns the SpaceName field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetSpaceName() string {
	if o == nil || o.SpaceName == nil {
		var ret string
		return ret
	}
	return *o.SpaceName
}

// GetSpaceNameOk returns a tuple with the SpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetSpaceNameOk() (*string, bool) {
	if o == nil || o.SpaceName == nil {
		return nil, false
	}
	return o.SpaceName, true
}

// HasSpaceName returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasSpaceName() bool {
	if o != nil && o.SpaceName != nil {
		return true
	}

	return false
}

// SetSpaceName gets a reference to the given string and assigns it to the SpaceName field.
func (o *IpamAddress6AddInput) SetSpaceName(v string) {
	o.SpaceName = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetDeviceId() int32 {
	if o == nil || o.DeviceId == nil {
		var ret int32
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetDeviceIdOk() (*int32, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int32 and assigns it to the DeviceId field.
func (o *IpamAddress6AddInput) SetDeviceId(v int32) {
	o.DeviceId = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetInterfaceId() int32 {
	if o == nil || o.InterfaceId == nil {
		var ret int32
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetInterfaceIdOk() (*int32, bool) {
	if o == nil || o.InterfaceId == nil {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasInterfaceId() bool {
	if o != nil && o.InterfaceId != nil {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given int32 and assigns it to the InterfaceId field.
func (o *IpamAddress6AddInput) SetInterfaceId(v int32) {
	o.InterfaceId = &v
}

// GetAddress6MacAddr returns the Address6MacAddr field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetAddress6MacAddr() string {
	if o == nil || o.Address6MacAddr == nil {
		var ret string
		return ret
	}
	return *o.Address6MacAddr
}

// GetAddress6MacAddrOk returns a tuple with the Address6MacAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetAddress6MacAddrOk() (*string, bool) {
	if o == nil || o.Address6MacAddr == nil {
		return nil, false
	}
	return o.Address6MacAddr, true
}

// HasAddress6MacAddr returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasAddress6MacAddr() bool {
	if o != nil && o.Address6MacAddr != nil {
		return true
	}

	return false
}

// SetAddress6MacAddr gets a reference to the given string and assigns it to the Address6MacAddr field.
func (o *IpamAddress6AddInput) SetAddress6MacAddr(v string) {
	o.Address6MacAddr = &v
}

// GetAddress6Name returns the Address6Name field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetAddress6Name() string {
	if o == nil || o.Address6Name == nil {
		var ret string
		return ret
	}
	return *o.Address6Name
}

// GetAddress6NameOk returns a tuple with the Address6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetAddress6NameOk() (*string, bool) {
	if o == nil || o.Address6Name == nil {
		return nil, false
	}
	return o.Address6Name, true
}

// HasAddress6Name returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasAddress6Name() bool {
	if o != nil && o.Address6Name != nil {
		return true
	}

	return false
}

// SetAddress6Name gets a reference to the given string and assigns it to the Address6Name field.
func (o *IpamAddress6AddInput) SetAddress6Name(v string) {
	o.Address6Name = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetPortId() int32 {
	if o == nil || o.PortId == nil {
		var ret int32
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetPortIdOk() (*int32, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int32 and assigns it to the PortId field.
func (o *IpamAddress6AddInput) SetPortId(v int32) {
	o.PortId = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetClassParametersToDelete() []string {
	if o == nil || o.ClassParametersToDelete == nil {
		var ret []string
		return ret
	}
	return *o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetClassParametersToDeleteOk() (*[]string, bool) {
	if o == nil || o.ClassParametersToDelete == nil {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasClassParametersToDelete() bool {
	if o != nil && o.ClassParametersToDelete != nil {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *IpamAddress6AddInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = &v
}

// GetAddress6ClassName returns the Address6ClassName field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetAddress6ClassName() string {
	if o == nil || o.Address6ClassName == nil {
		var ret string
		return ret
	}
	return *o.Address6ClassName
}

// GetAddress6ClassNameOk returns a tuple with the Address6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetAddress6ClassNameOk() (*string, bool) {
	if o == nil || o.Address6ClassName == nil {
		return nil, false
	}
	return o.Address6ClassName, true
}

// HasAddress6ClassName returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasAddress6ClassName() bool {
	if o != nil && o.Address6ClassName != nil {
		return true
	}

	return false
}

// SetAddress6ClassName gets a reference to the given string and assigns it to the Address6ClassName field.
func (o *IpamAddress6AddInput) SetAddress6ClassName(v string) {
	o.Address6ClassName = &v
}

// GetAddress6ClassParameters returns the Address6ClassParameters field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetAddress6ClassParameters() []ApiClassParameterInputEntry {
	if o == nil || o.Address6ClassParameters == nil {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return *o.Address6ClassParameters
}

// GetAddress6ClassParametersOk returns a tuple with the Address6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetAddress6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool) {
	if o == nil || o.Address6ClassParameters == nil {
		return nil, false
	}
	return o.Address6ClassParameters, true
}

// HasAddress6ClassParameters returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasAddress6ClassParameters() bool {
	if o != nil && o.Address6ClassParameters != nil {
		return true
	}

	return false
}

// SetAddress6ClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the Address6ClassParameters field.
func (o *IpamAddress6AddInput) SetAddress6ClassParameters(v []ApiClassParameterInputEntry) {
	o.Address6ClassParameters = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *IpamAddress6AddInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddress6AddInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *IpamAddress6AddInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *IpamAddress6AddInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o IpamAddress6AddInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Address6Hostaddr != nil {
		toSerialize["address6_hostaddr"] = o.Address6Hostaddr
	}
	if o.SpaceId != nil {
		toSerialize["space_id"] = o.SpaceId
	}
	if o.SpaceName != nil {
		toSerialize["space_name"] = o.SpaceName
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.InterfaceId != nil {
		toSerialize["interface_id"] = o.InterfaceId
	}
	if o.Address6MacAddr != nil {
		toSerialize["address6_mac_addr"] = o.Address6MacAddr
	}
	if o.Address6Name != nil {
		toSerialize["address6_name"] = o.Address6Name
	}
	if o.PortId != nil {
		toSerialize["port_id"] = o.PortId
	}
	if o.ClassParametersToDelete != nil {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if o.Address6ClassName != nil {
		toSerialize["address6_class_name"] = o.Address6ClassName
	}
	if o.Address6ClassParameters != nil {
		toSerialize["address6_class_parameters"] = o.Address6ClassParameters
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableIpamAddress6AddInput struct {
	value *IpamAddress6AddInput
	isSet bool
}

func (v NullableIpamAddress6AddInput) Get() *IpamAddress6AddInput {
	return v.value
}

func (v *NullableIpamAddress6AddInput) Set(val *IpamAddress6AddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamAddress6AddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamAddress6AddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamAddress6AddInput(val *IpamAddress6AddInput) *NullableIpamAddress6AddInput {
	return &NullableIpamAddress6AddInput{value: val, isSet: true}
}

func (v NullableIpamAddress6AddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamAddress6AddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


