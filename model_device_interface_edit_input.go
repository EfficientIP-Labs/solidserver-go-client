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

// DeviceInterfaceEditInput struct for DeviceInterfaceEditInput
type DeviceInterfaceEditInput struct {
	// The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice.
	DeviceId *int32 `json:"device_id,omitempty"`
	// The name of the Device Manager device.
	DeviceName *string `json:"device_name,omitempty"`
	// The IP addresses you want to associate with the Device Manager port or interface, as follows <b>&lt;ip4_list&gt;,&lt;ip6_list&gt;</b>.
	InterfaceAddr *string `json:"interface_addr,omitempty"`
	// The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify which port or interface to edit.
	InterfaceId *int32 `json:"interface_id,omitempty"`
	// The MAC address you want to associate with the Device Manager port or interface.
	InterfaceMac *string `json:"interface_mac,omitempty"`
	// The name of the Device Manager port or interface, each port or interface must have a unique name.
	InterfaceName *string `json:"interface_name,omitempty"`
	// A way to indicate if the object is either a <b>port</b> or an <b>interface</b>.
	InterfaceType *string `json:"interface_type,omitempty"`
	// The list of IPv4 addresses you want to associate to the Device Manager port or interface, in hexadecimal format, as follows: <b>&lt;ip_address&gt;, &lt;ip_address&gt;,...</b>
	InterfaceIp4List *string `json:"interface_ip4_list,omitempty"`
	// todo(here) :device.interface.edit.input.interface_address6_list. : List of strings
	InterfaceAddress6List *string `json:"interface_address6_list,omitempty"`
	// The database identifier (ID) of a NetChange port you want to associate with the Device Manager port or interface.
	PortId *int32 `json:"port_id,omitempty"`
	// The object activation status.<ul class=dashed ><li>                                                If set to <b>0</b>, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.<br/>                                            </li><li>                                                If set to <b>1</b>, the object is enabled and managed.<br/>                                            </li><li>                                                If set to <b>2</b>, the object is unmanaged, disabled or both depending on the context.<br/>                                            </li></ul>By default, <b>row_enabled</b> is set to <b>1</b> when an object is created.
	RowState *int32 `json:"row_state,omitempty"`
	// The database identifier (ID) of a space you want to associate with the Device Manager port or interface.
	SpaceId *int32 `json:"space_id,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete *[]string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	InterfaceClassName *string `json:"interface_class_name,omitempty"`
	// class parameters in json format
	InterfaceClassParameters *[]ApiClassParameterInputEntry `json:"interface_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewDeviceInterfaceEditInput instantiates a new DeviceInterfaceEditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceInterfaceEditInput() *DeviceInterfaceEditInput {
	this := DeviceInterfaceEditInput{}
	return &this
}

// NewDeviceInterfaceEditInputWithDefaults instantiates a new DeviceInterfaceEditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceInterfaceEditInputWithDefaults() *DeviceInterfaceEditInput {
	this := DeviceInterfaceEditInput{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetDeviceId() int32 {
	if o == nil || o.DeviceId == nil {
		var ret int32
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetDeviceIdOk() (*int32, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given int32 and assigns it to the DeviceId field.
func (o *DeviceInterfaceEditInput) SetDeviceId(v int32) {
	o.DeviceId = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *DeviceInterfaceEditInput) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetInterfaceAddr returns the InterfaceAddr field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceAddr() string {
	if o == nil || o.InterfaceAddr == nil {
		var ret string
		return ret
	}
	return *o.InterfaceAddr
}

// GetInterfaceAddrOk returns a tuple with the InterfaceAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceAddrOk() (*string, bool) {
	if o == nil || o.InterfaceAddr == nil {
		return nil, false
	}
	return o.InterfaceAddr, true
}

// HasInterfaceAddr returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceAddr() bool {
	if o != nil && o.InterfaceAddr != nil {
		return true
	}

	return false
}

// SetInterfaceAddr gets a reference to the given string and assigns it to the InterfaceAddr field.
func (o *DeviceInterfaceEditInput) SetInterfaceAddr(v string) {
	o.InterfaceAddr = &v
}

// GetInterfaceId returns the InterfaceId field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceId() int32 {
	if o == nil || o.InterfaceId == nil {
		var ret int32
		return ret
	}
	return *o.InterfaceId
}

// GetInterfaceIdOk returns a tuple with the InterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceIdOk() (*int32, bool) {
	if o == nil || o.InterfaceId == nil {
		return nil, false
	}
	return o.InterfaceId, true
}

// HasInterfaceId returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceId() bool {
	if o != nil && o.InterfaceId != nil {
		return true
	}

	return false
}

// SetInterfaceId gets a reference to the given int32 and assigns it to the InterfaceId field.
func (o *DeviceInterfaceEditInput) SetInterfaceId(v int32) {
	o.InterfaceId = &v
}

// GetInterfaceMac returns the InterfaceMac field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceMac() string {
	if o == nil || o.InterfaceMac == nil {
		var ret string
		return ret
	}
	return *o.InterfaceMac
}

// GetInterfaceMacOk returns a tuple with the InterfaceMac field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceMacOk() (*string, bool) {
	if o == nil || o.InterfaceMac == nil {
		return nil, false
	}
	return o.InterfaceMac, true
}

// HasInterfaceMac returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceMac() bool {
	if o != nil && o.InterfaceMac != nil {
		return true
	}

	return false
}

// SetInterfaceMac gets a reference to the given string and assigns it to the InterfaceMac field.
func (o *DeviceInterfaceEditInput) SetInterfaceMac(v string) {
	o.InterfaceMac = &v
}

// GetInterfaceName returns the InterfaceName field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceName() string {
	if o == nil || o.InterfaceName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceName
}

// GetInterfaceNameOk returns a tuple with the InterfaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceNameOk() (*string, bool) {
	if o == nil || o.InterfaceName == nil {
		return nil, false
	}
	return o.InterfaceName, true
}

// HasInterfaceName returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceName() bool {
	if o != nil && o.InterfaceName != nil {
		return true
	}

	return false
}

// SetInterfaceName gets a reference to the given string and assigns it to the InterfaceName field.
func (o *DeviceInterfaceEditInput) SetInterfaceName(v string) {
	o.InterfaceName = &v
}

// GetInterfaceType returns the InterfaceType field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceType() string {
	if o == nil || o.InterfaceType == nil {
		var ret string
		return ret
	}
	return *o.InterfaceType
}

// GetInterfaceTypeOk returns a tuple with the InterfaceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceTypeOk() (*string, bool) {
	if o == nil || o.InterfaceType == nil {
		return nil, false
	}
	return o.InterfaceType, true
}

// HasInterfaceType returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceType() bool {
	if o != nil && o.InterfaceType != nil {
		return true
	}

	return false
}

// SetInterfaceType gets a reference to the given string and assigns it to the InterfaceType field.
func (o *DeviceInterfaceEditInput) SetInterfaceType(v string) {
	o.InterfaceType = &v
}

// GetInterfaceIp4List returns the InterfaceIp4List field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceIp4List() string {
	if o == nil || o.InterfaceIp4List == nil {
		var ret string
		return ret
	}
	return *o.InterfaceIp4List
}

// GetInterfaceIp4ListOk returns a tuple with the InterfaceIp4List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceIp4ListOk() (*string, bool) {
	if o == nil || o.InterfaceIp4List == nil {
		return nil, false
	}
	return o.InterfaceIp4List, true
}

// HasInterfaceIp4List returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceIp4List() bool {
	if o != nil && o.InterfaceIp4List != nil {
		return true
	}

	return false
}

// SetInterfaceIp4List gets a reference to the given string and assigns it to the InterfaceIp4List field.
func (o *DeviceInterfaceEditInput) SetInterfaceIp4List(v string) {
	o.InterfaceIp4List = &v
}

// GetInterfaceAddress6List returns the InterfaceAddress6List field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceAddress6List() string {
	if o == nil || o.InterfaceAddress6List == nil {
		var ret string
		return ret
	}
	return *o.InterfaceAddress6List
}

// GetInterfaceAddress6ListOk returns a tuple with the InterfaceAddress6List field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceAddress6ListOk() (*string, bool) {
	if o == nil || o.InterfaceAddress6List == nil {
		return nil, false
	}
	return o.InterfaceAddress6List, true
}

// HasInterfaceAddress6List returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceAddress6List() bool {
	if o != nil && o.InterfaceAddress6List != nil {
		return true
	}

	return false
}

// SetInterfaceAddress6List gets a reference to the given string and assigns it to the InterfaceAddress6List field.
func (o *DeviceInterfaceEditInput) SetInterfaceAddress6List(v string) {
	o.InterfaceAddress6List = &v
}

// GetPortId returns the PortId field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetPortId() int32 {
	if o == nil || o.PortId == nil {
		var ret int32
		return ret
	}
	return *o.PortId
}

// GetPortIdOk returns a tuple with the PortId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetPortIdOk() (*int32, bool) {
	if o == nil || o.PortId == nil {
		return nil, false
	}
	return o.PortId, true
}

// HasPortId returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasPortId() bool {
	if o != nil && o.PortId != nil {
		return true
	}

	return false
}

// SetPortId gets a reference to the given int32 and assigns it to the PortId field.
func (o *DeviceInterfaceEditInput) SetPortId(v int32) {
	o.PortId = &v
}

// GetRowState returns the RowState field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetRowState() int32 {
	if o == nil || o.RowState == nil {
		var ret int32
		return ret
	}
	return *o.RowState
}

// GetRowStateOk returns a tuple with the RowState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetRowStateOk() (*int32, bool) {
	if o == nil || o.RowState == nil {
		return nil, false
	}
	return o.RowState, true
}

// HasRowState returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasRowState() bool {
	if o != nil && o.RowState != nil {
		return true
	}

	return false
}

// SetRowState gets a reference to the given int32 and assigns it to the RowState field.
func (o *DeviceInterfaceEditInput) SetRowState(v int32) {
	o.RowState = &v
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetSpaceId() int32 {
	if o == nil || o.SpaceId == nil {
		var ret int32
		return ret
	}
	return *o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetSpaceIdOk() (*int32, bool) {
	if o == nil || o.SpaceId == nil {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasSpaceId() bool {
	if o != nil && o.SpaceId != nil {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given int32 and assigns it to the SpaceId field.
func (o *DeviceInterfaceEditInput) SetSpaceId(v int32) {
	o.SpaceId = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetClassParametersToDelete() []string {
	if o == nil || o.ClassParametersToDelete == nil {
		var ret []string
		return ret
	}
	return *o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetClassParametersToDeleteOk() (*[]string, bool) {
	if o == nil || o.ClassParametersToDelete == nil {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasClassParametersToDelete() bool {
	if o != nil && o.ClassParametersToDelete != nil {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *DeviceInterfaceEditInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = &v
}

// GetInterfaceClassName returns the InterfaceClassName field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceClassName() string {
	if o == nil || o.InterfaceClassName == nil {
		var ret string
		return ret
	}
	return *o.InterfaceClassName
}

// GetInterfaceClassNameOk returns a tuple with the InterfaceClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceClassNameOk() (*string, bool) {
	if o == nil || o.InterfaceClassName == nil {
		return nil, false
	}
	return o.InterfaceClassName, true
}

// HasInterfaceClassName returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceClassName() bool {
	if o != nil && o.InterfaceClassName != nil {
		return true
	}

	return false
}

// SetInterfaceClassName gets a reference to the given string and assigns it to the InterfaceClassName field.
func (o *DeviceInterfaceEditInput) SetInterfaceClassName(v string) {
	o.InterfaceClassName = &v
}

// GetInterfaceClassParameters returns the InterfaceClassParameters field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetInterfaceClassParameters() []ApiClassParameterInputEntry {
	if o == nil || o.InterfaceClassParameters == nil {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return *o.InterfaceClassParameters
}

// GetInterfaceClassParametersOk returns a tuple with the InterfaceClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetInterfaceClassParametersOk() (*[]ApiClassParameterInputEntry, bool) {
	if o == nil || o.InterfaceClassParameters == nil {
		return nil, false
	}
	return o.InterfaceClassParameters, true
}

// HasInterfaceClassParameters returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasInterfaceClassParameters() bool {
	if o != nil && o.InterfaceClassParameters != nil {
		return true
	}

	return false
}

// SetInterfaceClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the InterfaceClassParameters field.
func (o *DeviceInterfaceEditInput) SetInterfaceClassParameters(v []ApiClassParameterInputEntry) {
	o.InterfaceClassParameters = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DeviceInterfaceEditInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceInterfaceEditInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DeviceInterfaceEditInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *DeviceInterfaceEditInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o DeviceInterfaceEditInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.DeviceName != nil {
		toSerialize["device_name"] = o.DeviceName
	}
	if o.InterfaceAddr != nil {
		toSerialize["interface_addr"] = o.InterfaceAddr
	}
	if o.InterfaceId != nil {
		toSerialize["interface_id"] = o.InterfaceId
	}
	if o.InterfaceMac != nil {
		toSerialize["interface_mac"] = o.InterfaceMac
	}
	if o.InterfaceName != nil {
		toSerialize["interface_name"] = o.InterfaceName
	}
	if o.InterfaceType != nil {
		toSerialize["interface_type"] = o.InterfaceType
	}
	if o.InterfaceIp4List != nil {
		toSerialize["interface_ip4_list"] = o.InterfaceIp4List
	}
	if o.InterfaceAddress6List != nil {
		toSerialize["interface_address6_list"] = o.InterfaceAddress6List
	}
	if o.PortId != nil {
		toSerialize["port_id"] = o.PortId
	}
	if o.RowState != nil {
		toSerialize["row_state"] = o.RowState
	}
	if o.SpaceId != nil {
		toSerialize["space_id"] = o.SpaceId
	}
	if o.ClassParametersToDelete != nil {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if o.InterfaceClassName != nil {
		toSerialize["interface_class_name"] = o.InterfaceClassName
	}
	if o.InterfaceClassParameters != nil {
		toSerialize["interface_class_parameters"] = o.InterfaceClassParameters
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceInterfaceEditInput struct {
	value *DeviceInterfaceEditInput
	isSet bool
}

func (v NullableDeviceInterfaceEditInput) Get() *DeviceInterfaceEditInput {
	return v.value
}

func (v *NullableDeviceInterfaceEditInput) Set(val *DeviceInterfaceEditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceInterfaceEditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceInterfaceEditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceInterfaceEditInput(val *DeviceInterfaceEditInput) *NullableDeviceInterfaceEditInput {
	return &NullableDeviceInterfaceEditInput{value: val, isSet: true}
}

func (v NullableDeviceInterfaceEditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceInterfaceEditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

