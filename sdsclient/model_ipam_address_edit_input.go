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

// IpamAddressEditInput struct for IpamAddressEditInput
type IpamAddressEditInput struct {
	// The IP address.
	AddressHostaddr *string `json:"address_hostaddr,omitempty"`
	// The database identifier (ID) of the IPv4 address, a unique numeric key value automatically incremented when you add an IPv4 address. Use the ID to specify which IPv4 address to edit.
	AddressId *int32 `json:"address_id,omitempty"`
	// The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice.
	SpaceId *int32 `json:"space_id,omitempty"`
	// The name of the space.
	SpaceName *string `json:"space_name,omitempty"`
	// A way to force a validity check, if you configured the IPAM to DHCP replication. If the check is enabled (<b>1</b>), the configuration of the IP address you are adding must be valid as well for the DHCP.
	CheckIsServerIp *int32 `json:"check_is_server_ip,omitempty"`
	// The database identifier (ID) of the DHCP static you want to associate with the IP address.
	StaticId *int32 `json:"static_id,omitempty"`
	// The database identifier (ID) of the DHCP lease you want to associate with the IP address.
	LeaseId *int32 `json:"lease_id,omitempty"`
	// The database identifier (ID) of the Device Manager device you want to associate with the IP address.
	DeviceId *int32 `json:"device_id,omitempty"`
	// The database identifier (ID) of the Device Manager interface you want to associate with the IP address.
	InterfaceId *int32 `json:"interface_id,omitempty"`
	// The database identifier (ID) of the NetChange port you want to associate with the IP address.
	PortId *int32 `json:"port_id,omitempty"`
	// The MAC address you want to associate with the IPv4 address.
	AddressMacAddr *string `json:"address_mac_addr,omitempty"`
	// The name of the IPv4 address, each IPv4 address must have a unique name.
	AddressName *string `json:"address_name,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete *[]string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	AddressClassName *string `json:"address_class_name,omitempty"`
	// class parameters in json format
	AddressClassParameters *[]ApiClassParameterInputEntry `json:"address_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewIpamAddressEditInput instantiates a new IpamAddressEditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamAddressEditInput() *IpamAddressEditInput {
	this := IpamAddressEditInput{}
	return &this
}

// NewIpamAddressEditInputWithDefaults instantiates a new IpamAddressEditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamAddressEditInputWithDefaults() *IpamAddressEditInput {
	this := IpamAddressEditInput{}
	return &this
}

// GetAddressHostaddr returns the AddressHostaddr field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetAddressHostaddr() string {
	if o == nil || o.AddressHostaddr == nil {
		var ret string
		return ret
	}
	return *o.AddressHostaddr
}

// GetAddressHostaddrOk returns a tuple with the AddressHostaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetAddressHostaddrOk() (*string, bool) {
	if o == nil || o.AddressHostaddr == nil {
		return nil, false
	}
	return o.AddressHostaddr, true
}

// HasAddressHostaddr returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasAddressHostaddr() bool {
	if o != nil && o.AddressHostaddr != nil {
		return true
	}

	return false
}

// SetAddressHostaddr gets a reference to the given string and assigns it to the AddressHostaddr field.
func (o *IpamAddressEditInput) SetAddressHostaddr(v string) {
	o.AddressHostaddr = &v
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetAddressId() int32 {
	if o == nil || o.AddressId == nil {
		var ret int32
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetAddressIdOk() (*int32, bool) {
	if o == nil || o.AddressId == nil {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasAddressId() bool {
	if o != nil && o.AddressId != nil {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given int32 and assigns it to the AddressId field.
func (o *IpamAddressEditInput) SetAddressId(v int32) {
	o.AddressId = &v
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetSpaceId() int32 {
	if o == nil || o.SpaceId == nil {
		var ret int32
		return ret
	}
	return *o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetSpaceIdOk() (*int32, bool) {
	if o == nil || o.SpaceId == nil {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasSpaceId() bool {
	if o != nil && o.SpaceId != nil {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given int32 and assigns it to the SpaceId field.
func (o *IpamAddressEditInput) SetSpaceId(v int32) {
	o.SpaceId = &v
}

// GetSpaceName returns the SpaceName field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetSpaceName() string {
	if o == nil || o.SpaceName == nil {
		var ret string
		return ret
	}
	return *o.SpaceName
}

// GetSpaceNameOk returns a tuple with the SpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetSpaceNameOk() (*string, bool) {
	if o == nil || o.SpaceName == nil {
		return nil, false
	}
	return o.SpaceName, true
}

// HasSpaceName returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasSpaceName() bool {
	if o != nil && o.SpaceName != nil {
		return true
	}

	return false
}

// SetSpaceName gets a reference to the given string and assigns it to the SpaceName field.
func (o *IpamAddressEditInput) SetSpaceName(v string) {
	o.SpaceName = &v
}

// GetCheckIsServerIp returns the CheckIsServerIp field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetCheckIsServerIp() int32 {
	if o == nil || o.CheckIsServerIp == nil {
		var ret int32
		return ret
	}
	return *o.CheckIsServerIp
}

// GetCheckIsServerIpOk returns a tuple with the CheckIsServerIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetCheckIsServerIpOk() (*int32, bool) {
	if o == nil || o.CheckIsServerIp == nil {
		return nil, false
	}
	return o.CheckIsServerIp, true
}

// HasCheckIsServerIp returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasCheckIsServerIp() bool {
	if o != nil && o.CheckIsServerIp != nil {
		return true
	}

	return false
}

// SetCheckIsServerIp gets a reference to the given int32 and assigns it to the CheckIsServerIp field.
func (o *IpamAddressEditInput) SetCheckIsServerIp(v int32) {
	o.CheckIsServerIp = &v
}

// GetStaticId returns the StaticId field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetStaticId() int32 {
	if o == nil || o.StaticId == nil {
		var ret int32
		return ret
	}
	return *o.StaticId
}

// GetStaticIdOk returns a tuple with the StaticId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetStaticIdOk() (*int32, bool) {
	if o == nil || o.StaticId == nil {
		return nil, false
	}
	return o.StaticId, true
}

// HasStaticId returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasStaticId() bool {
	if o != nil && o.StaticId != nil {
		return true
	}

	return false
}

// SetStaticId gets a reference to the given int32 and assigns it to the StaticId field.
func (o *IpamAddressEditInput) SetStaticId(v int32) {
	o.StaticId = &v
}

// GetLeaseId returns the LeaseId field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetLeaseId() int32 {
	if o == nil || o.LeaseId == nil {
		var ret int32
		return ret
	}
	return *o.LeaseId
}

// GetLeaseIdOk returns a tuple with the LeaseId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetLeaseIdOk() (*int32, bool) {
	if o == nil || o.LeaseId == nil {
		return nil, false
	}
	return o.LeaseId, true
}

// HasLeaseId returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasLeaseId() bool {
	if o != nil && o.LeaseId != nil {
		return true
	}

	return false
}

// SetLeaseId gets a reference to the given int32 and assigns it to the LeaseId field.
func (o *IpamAddressEditInput) SetLeaseId(v int32) {
	o.LeaseId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetDeviceId() int32 {
	if o == nil || o.DeviceId == nil {
		var ret int32
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetDeviceIdOk() (*int32, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int32 and assigns it to the DeviceId field.
func (o *IpamAddressEditInput) SetDeviceId(v int32) {
	o.DeviceId = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetInterfaceId() int32 {
	if o == nil || o.InterfaceId == nil {
		var ret int32
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetInterfaceIdOk() (*int32, bool) {
	if o == nil || o.InterfaceId == nil {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasInterfaceId() bool {
	if o != nil && o.InterfaceId != nil {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given int32 and assigns it to the InterfaceId field.
func (o *IpamAddressEditInput) SetInterfaceId(v int32) {
	o.InterfaceId = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetPortId() int32 {
	if o == nil || o.PortId == nil {
		var ret int32
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetPortIdOk() (*int32, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int32 and assigns it to the PortId field.
func (o *IpamAddressEditInput) SetPortId(v int32) {
	o.PortId = &v
}

// GetAddressMacAddr returns the AddressMacAddr field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetAddressMacAddr() string {
	if o == nil || o.AddressMacAddr == nil {
		var ret string
		return ret
	}
	return *o.AddressMacAddr
}

// GetAddressMacAddrOk returns a tuple with the AddressMacAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetAddressMacAddrOk() (*string, bool) {
	if o == nil || o.AddressMacAddr == nil {
		return nil, false
	}
	return o.AddressMacAddr, true
}

// HasAddressMacAddr returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasAddressMacAddr() bool {
	if o != nil && o.AddressMacAddr != nil {
		return true
	}

	return false
}

// SetAddressMacAddr gets a reference to the given string and assigns it to the AddressMacAddr field.
func (o *IpamAddressEditInput) SetAddressMacAddr(v string) {
	o.AddressMacAddr = &v
}

// GetAddressName returns the AddressName field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetAddressName() string {
	if o == nil || o.AddressName == nil {
		var ret string
		return ret
	}
	return *o.AddressName
}

// GetAddressNameOk returns a tuple with the AddressName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetAddressNameOk() (*string, bool) {
	if o == nil || o.AddressName == nil {
		return nil, false
	}
	return o.AddressName, true
}

// HasAddressName returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasAddressName() bool {
	if o != nil && o.AddressName != nil {
		return true
	}

	return false
}

// SetAddressName gets a reference to the given string and assigns it to the AddressName field.
func (o *IpamAddressEditInput) SetAddressName(v string) {
	o.AddressName = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetClassParametersToDelete() []string {
	if o == nil || o.ClassParametersToDelete == nil {
		var ret []string
		return ret
	}
	return *o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetClassParametersToDeleteOk() (*[]string, bool) {
	if o == nil || o.ClassParametersToDelete == nil {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasClassParametersToDelete() bool {
	if o != nil && o.ClassParametersToDelete != nil {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *IpamAddressEditInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = &v
}

// GetAddressClassName returns the AddressClassName field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetAddressClassName() string {
	if o == nil || o.AddressClassName == nil {
		var ret string
		return ret
	}
	return *o.AddressClassName
}

// GetAddressClassNameOk returns a tuple with the AddressClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetAddressClassNameOk() (*string, bool) {
	if o == nil || o.AddressClassName == nil {
		return nil, false
	}
	return o.AddressClassName, true
}

// HasAddressClassName returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasAddressClassName() bool {
	if o != nil && o.AddressClassName != nil {
		return true
	}

	return false
}

// SetAddressClassName gets a reference to the given string and assigns it to the AddressClassName field.
func (o *IpamAddressEditInput) SetAddressClassName(v string) {
	o.AddressClassName = &v
}

// GetAddressClassParameters returns the AddressClassParameters field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetAddressClassParameters() []ApiClassParameterInputEntry {
	if o == nil || o.AddressClassParameters == nil {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return *o.AddressClassParameters
}

// GetAddressClassParametersOk returns a tuple with the AddressClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetAddressClassParametersOk() (*[]ApiClassParameterInputEntry, bool) {
	if o == nil || o.AddressClassParameters == nil {
		return nil, false
	}
	return o.AddressClassParameters, true
}

// HasAddressClassParameters returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasAddressClassParameters() bool {
	if o != nil && o.AddressClassParameters != nil {
		return true
	}

	return false
}

// SetAddressClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the AddressClassParameters field.
func (o *IpamAddressEditInput) SetAddressClassParameters(v []ApiClassParameterInputEntry) {
	o.AddressClassParameters = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *IpamAddressEditInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAddressEditInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *IpamAddressEditInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *IpamAddressEditInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o IpamAddressEditInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressHostaddr != nil {
		toSerialize["address_hostaddr"] = o.AddressHostaddr
	}
	if o.AddressId != nil {
		toSerialize["address_id"] = o.AddressId
	}
	if o.SpaceId != nil {
		toSerialize["space_id"] = o.SpaceId
	}
	if o.SpaceName != nil {
		toSerialize["space_name"] = o.SpaceName
	}
	if o.CheckIsServerIp != nil {
		toSerialize["check_is_server_ip"] = o.CheckIsServerIp
	}
	if o.StaticId != nil {
		toSerialize["static_id"] = o.StaticId
	}
	if o.LeaseId != nil {
		toSerialize["lease_id"] = o.LeaseId
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.InterfaceId != nil {
		toSerialize["interface_id"] = o.InterfaceId
	}
	if o.PortId != nil {
		toSerialize["port_id"] = o.PortId
	}
	if o.AddressMacAddr != nil {
		toSerialize["address_mac_addr"] = o.AddressMacAddr
	}
	if o.AddressName != nil {
		toSerialize["address_name"] = o.AddressName
	}
	if o.ClassParametersToDelete != nil {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if o.AddressClassName != nil {
		toSerialize["address_class_name"] = o.AddressClassName
	}
	if o.AddressClassParameters != nil {
		toSerialize["address_class_parameters"] = o.AddressClassParameters
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableIpamAddressEditInput struct {
	value *IpamAddressEditInput
	isSet bool
}

func (v NullableIpamAddressEditInput) Get() *IpamAddressEditInput {
	return v.value
}

func (v *NullableIpamAddressEditInput) Set(val *IpamAddressEditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamAddressEditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamAddressEditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamAddressEditInput(val *IpamAddressEditInput) *NullableIpamAddressEditInput {
	return &NullableIpamAddressEditInput{value: val, isSet: true}
}

func (v NullableIpamAddressEditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamAddressEditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


