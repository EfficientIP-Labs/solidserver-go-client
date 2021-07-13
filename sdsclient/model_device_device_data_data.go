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

// DeviceDeviceDataData struct for DeviceDeviceDataData
type DeviceDeviceDataData struct {
	// The name of the class applied to the device, it can be preceded by the class directory.
	DeviceClassName *string `json:"device_class_name,omitempty"`
	// #general.output.class_parameters#
	DeviceClassParameters *[]ApiClassParameterOutputEntry `json:"device_class_parameters,omitempty"`
	// The database identifier (ID) of the Device Manager device.
	DeviceId *string `json:"device_id,omitempty"`
	// The IP address associated with the Device Manager device.
	DeviceAddressAddr *string `json:"device_address_addr,omitempty"`
	// The human readable version of the parameter <b>hostdev_ip_addr</b>.
	DeviceAddressFormated *string `json:"device_address_formated,omitempty"`
	// The name of the Device Manager device.
	DeviceName *string `json:"device_name,omitempty"`
	// The database identifier (ID) of the space associated with the Device Manager device.
	DeviceSpaceId *string `json:"device_space_id,omitempty"`
	// The name of the space associated with the Device Manager device.
	DeviceSpaceName *string `json:"device_space_name,omitempty"`
	// The number of interfaces on the Device Manager device that are currently free.
	InterfaceFree *string `json:"interface_free,omitempty"`
	// The total number of interfaces on the Device Manager device.
	InterfaceTotal *string `json:"interface_total,omitempty"`
	// The number of interfaces on the Device Manager device that are currently active.
	InterfaceUsed *string `json:"interface_used,omitempty"`
	// The percentage of interfaces on the Device Manager device that are currently active.
	InterfaceUsedPercent *string `json:"interface_used_percent,omitempty"`
	// The database identifier (ID) of the NetChange network device associated with the Device Manager device.
	DevId *string `json:"dev_id,omitempty"`
	// The name of the NetChange network device associated with the Device Manager device.
	DevName *string `json:"dev_name,omitempty"`
	// The number of ports on the Device Manager device that are currently free.
	DevicePortFree *string `json:"device_port_free,omitempty"`
	// The total number of ports on the Device Manager device.
	DevicePortTotal *string `json:"device_port_total,omitempty"`
	// The number of ports on the Device Manager device that are currently active.
	DevicePortUsed *string `json:"device_port_used,omitempty"`
	// The percentage of ports on the Device Manager device that are currently active.
	DevicePortUsedPercent *string `json:"device_port_used_percent,omitempty"`
	// The object activation status:<ul class=dashed ><li><b>0</b> indicates the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.<br/>                                            </li><li><b>1</b> indicates the object is enabled and managed.<br/>                                            </li><li><b>2</b> indicates the object is unmanaged, disabled or both depending on the context.<br/>                                            </li></ul>By default, <b>row_enabled</b> is set to <b>1</b> when an object is created.
	RowState *string `json:"row_state,omitempty"`
}

// NewDeviceDeviceDataData instantiates a new DeviceDeviceDataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceDeviceDataData() *DeviceDeviceDataData {
	this := DeviceDeviceDataData{}
	return &this
}

// NewDeviceDeviceDataDataWithDefaults instantiates a new DeviceDeviceDataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceDeviceDataDataWithDefaults() *DeviceDeviceDataData {
	this := DeviceDeviceDataData{}
	return &this
}

// GetDeviceClassName returns the DeviceClassName field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDeviceClassName() string {
	if o == nil || o.DeviceClassName == nil {
		var ret string
		return ret
	}
	return *o.DeviceClassName
}

// GetDeviceClassNameOk returns a tuple with the DeviceClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDeviceClassNameOk() (*string, bool) {
	if o == nil || o.DeviceClassName == nil {
		return nil, false
	}
	return o.DeviceClassName, true
}

// HasDeviceClassName returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDeviceClassName() bool {
	if o != nil && o.DeviceClassName != nil {
		return true
	}

	return false
}

// SetDeviceClassName gets a reference to the given string and assigns it to the DeviceClassName field.
func (o *DeviceDeviceDataData) SetDeviceClassName(v string) {
	o.DeviceClassName = &v
}

// GetDeviceClassParameters returns the DeviceClassParameters field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDeviceClassParameters() []ApiClassParameterOutputEntry {
	if o == nil || o.DeviceClassParameters == nil {
		var ret []ApiClassParameterOutputEntry
		return ret
	}
	return *o.DeviceClassParameters
}

// GetDeviceClassParametersOk returns a tuple with the DeviceClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDeviceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool) {
	if o == nil || o.DeviceClassParameters == nil {
		return nil, false
	}
	return o.DeviceClassParameters, true
}

// HasDeviceClassParameters returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDeviceClassParameters() bool {
	if o != nil && o.DeviceClassParameters != nil {
		return true
	}

	return false
}

// SetDeviceClassParameters gets a reference to the given []ApiClassParameterOutputEntry and assigns it to the DeviceClassParameters field.
func (o *DeviceDeviceDataData) SetDeviceClassParameters(v []ApiClassParameterOutputEntry) {
	o.DeviceClassParameters = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DeviceDeviceDataData) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetDeviceAddressAddr returns the DeviceAddressAddr field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDeviceAddressAddr() string {
	if o == nil || o.DeviceAddressAddr == nil {
		var ret string
		return ret
	}
	return *o.DeviceAddressAddr
}

// GetDeviceAddressAddrOk returns a tuple with the DeviceAddressAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDeviceAddressAddrOk() (*string, bool) {
	if o == nil || o.DeviceAddressAddr == nil {
		return nil, false
	}
	return o.DeviceAddressAddr, true
}

// HasDeviceAddressAddr returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDeviceAddressAddr() bool {
	if o != nil && o.DeviceAddressAddr != nil {
		return true
	}

	return false
}

// SetDeviceAddressAddr gets a reference to the given string and assigns it to the DeviceAddressAddr field.
func (o *DeviceDeviceDataData) SetDeviceAddressAddr(v string) {
	o.DeviceAddressAddr = &v
}

// GetDeviceAddressFormated returns the DeviceAddressFormated field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDeviceAddressFormated() string {
	if o == nil || o.DeviceAddressFormated == nil {
		var ret string
		return ret
	}
	return *o.DeviceAddressFormated
}

// GetDeviceAddressFormatedOk returns a tuple with the DeviceAddressFormated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDeviceAddressFormatedOk() (*string, bool) {
	if o == nil || o.DeviceAddressFormated == nil {
		return nil, false
	}
	return o.DeviceAddressFormated, true
}

// HasDeviceAddressFormated returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDeviceAddressFormated() bool {
	if o != nil && o.DeviceAddressFormated != nil {
		return true
	}

	return false
}

// SetDeviceAddressFormated gets a reference to the given string and assigns it to the DeviceAddressFormated field.
func (o *DeviceDeviceDataData) SetDeviceAddressFormated(v string) {
	o.DeviceAddressFormated = &v
}

// GetDeviceName returns the DeviceName field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDeviceName() string {
	if o == nil || o.DeviceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceName
}

// GetDeviceNameOk returns a tuple with the DeviceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDeviceNameOk() (*string, bool) {
	if o == nil || o.DeviceName == nil {
		return nil, false
	}
	return o.DeviceName, true
}

// HasDeviceName returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDeviceName() bool {
	if o != nil && o.DeviceName != nil {
		return true
	}

	return false
}

// SetDeviceName gets a reference to the given string and assigns it to the DeviceName field.
func (o *DeviceDeviceDataData) SetDeviceName(v string) {
	o.DeviceName = &v
}

// GetDeviceSpaceId returns the DeviceSpaceId field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDeviceSpaceId() string {
	if o == nil || o.DeviceSpaceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceSpaceId
}

// GetDeviceSpaceIdOk returns a tuple with the DeviceSpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDeviceSpaceIdOk() (*string, bool) {
	if o == nil || o.DeviceSpaceId == nil {
		return nil, false
	}
	return o.DeviceSpaceId, true
}

// HasDeviceSpaceId returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDeviceSpaceId() bool {
	if o != nil && o.DeviceSpaceId != nil {
		return true
	}

	return false
}

// SetDeviceSpaceId gets a reference to the given string and assigns it to the DeviceSpaceId field.
func (o *DeviceDeviceDataData) SetDeviceSpaceId(v string) {
	o.DeviceSpaceId = &v
}

// GetDeviceSpaceName returns the DeviceSpaceName field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDeviceSpaceName() string {
	if o == nil || o.DeviceSpaceName == nil {
		var ret string
		return ret
	}
	return *o.DeviceSpaceName
}

// GetDeviceSpaceNameOk returns a tuple with the DeviceSpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDeviceSpaceNameOk() (*string, bool) {
	if o == nil || o.DeviceSpaceName == nil {
		return nil, false
	}
	return o.DeviceSpaceName, true
}

// HasDeviceSpaceName returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDeviceSpaceName() bool {
	if o != nil && o.DeviceSpaceName != nil {
		return true
	}

	return false
}

// SetDeviceSpaceName gets a reference to the given string and assigns it to the DeviceSpaceName field.
func (o *DeviceDeviceDataData) SetDeviceSpaceName(v string) {
	o.DeviceSpaceName = &v
}

// GetInterfaceFree returns the InterfaceFree field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetInterfaceFree() string {
	if o == nil || o.InterfaceFree == nil {
		var ret string
		return ret
	}
	return *o.InterfaceFree
}

// GetInterfaceFreeOk returns a tuple with the InterfaceFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetInterfaceFreeOk() (*string, bool) {
	if o == nil || o.InterfaceFree == nil {
		return nil, false
	}
	return o.InterfaceFree, true
}

// HasInterfaceFree returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasInterfaceFree() bool {
	if o != nil && o.InterfaceFree != nil {
		return true
	}

	return false
}

// SetInterfaceFree gets a reference to the given string and assigns it to the InterfaceFree field.
func (o *DeviceDeviceDataData) SetInterfaceFree(v string) {
	o.InterfaceFree = &v
}

// GetInterfaceTotal returns the InterfaceTotal field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetInterfaceTotal() string {
	if o == nil || o.InterfaceTotal == nil {
		var ret string
		return ret
	}
	return *o.InterfaceTotal
}

// GetInterfaceTotalOk returns a tuple with the InterfaceTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetInterfaceTotalOk() (*string, bool) {
	if o == nil || o.InterfaceTotal == nil {
		return nil, false
	}
	return o.InterfaceTotal, true
}

// HasInterfaceTotal returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasInterfaceTotal() bool {
	if o != nil && o.InterfaceTotal != nil {
		return true
	}

	return false
}

// SetInterfaceTotal gets a reference to the given string and assigns it to the InterfaceTotal field.
func (o *DeviceDeviceDataData) SetInterfaceTotal(v string) {
	o.InterfaceTotal = &v
}

// GetInterfaceUsed returns the InterfaceUsed field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetInterfaceUsed() string {
	if o == nil || o.InterfaceUsed == nil {
		var ret string
		return ret
	}
	return *o.InterfaceUsed
}

// GetInterfaceUsedOk returns a tuple with the InterfaceUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetInterfaceUsedOk() (*string, bool) {
	if o == nil || o.InterfaceUsed == nil {
		return nil, false
	}
	return o.InterfaceUsed, true
}

// HasInterfaceUsed returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasInterfaceUsed() bool {
	if o != nil && o.InterfaceUsed != nil {
		return true
	}

	return false
}

// SetInterfaceUsed gets a reference to the given string and assigns it to the InterfaceUsed field.
func (o *DeviceDeviceDataData) SetInterfaceUsed(v string) {
	o.InterfaceUsed = &v
}

// GetInterfaceUsedPercent returns the InterfaceUsedPercent field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetInterfaceUsedPercent() string {
	if o == nil || o.InterfaceUsedPercent == nil {
		var ret string
		return ret
	}
	return *o.InterfaceUsedPercent
}

// GetInterfaceUsedPercentOk returns a tuple with the InterfaceUsedPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetInterfaceUsedPercentOk() (*string, bool) {
	if o == nil || o.InterfaceUsedPercent == nil {
		return nil, false
	}
	return o.InterfaceUsedPercent, true
}

// HasInterfaceUsedPercent returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasInterfaceUsedPercent() bool {
	if o != nil && o.InterfaceUsedPercent != nil {
		return true
	}

	return false
}

// SetInterfaceUsedPercent gets a reference to the given string and assigns it to the InterfaceUsedPercent field.
func (o *DeviceDeviceDataData) SetInterfaceUsedPercent(v string) {
	o.InterfaceUsedPercent = &v
}

// GetDevId returns the DevId field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDevId() string {
	if o == nil || o.DevId == nil {
		var ret string
		return ret
	}
	return *o.DevId
}

// GetDevIdOk returns a tuple with the DevId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDevIdOk() (*string, bool) {
	if o == nil || o.DevId == nil {
		return nil, false
	}
	return o.DevId, true
}

// HasDevId returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDevId() bool {
	if o != nil && o.DevId != nil {
		return true
	}

	return false
}

// SetDevId gets a reference to the given string and assigns it to the DevId field.
func (o *DeviceDeviceDataData) SetDevId(v string) {
	o.DevId = &v
}

// GetDevName returns the DevName field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDevName() string {
	if o == nil || o.DevName == nil {
		var ret string
		return ret
	}
	return *o.DevName
}

// GetDevNameOk returns a tuple with the DevName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDevNameOk() (*string, bool) {
	if o == nil || o.DevName == nil {
		return nil, false
	}
	return o.DevName, true
}

// HasDevName returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDevName() bool {
	if o != nil && o.DevName != nil {
		return true
	}

	return false
}

// SetDevName gets a reference to the given string and assigns it to the DevName field.
func (o *DeviceDeviceDataData) SetDevName(v string) {
	o.DevName = &v
}

// GetDevicePortFree returns the DevicePortFree field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDevicePortFree() string {
	if o == nil || o.DevicePortFree == nil {
		var ret string
		return ret
	}
	return *o.DevicePortFree
}

// GetDevicePortFreeOk returns a tuple with the DevicePortFree field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDevicePortFreeOk() (*string, bool) {
	if o == nil || o.DevicePortFree == nil {
		return nil, false
	}
	return o.DevicePortFree, true
}

// HasDevicePortFree returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDevicePortFree() bool {
	if o != nil && o.DevicePortFree != nil {
		return true
	}

	return false
}

// SetDevicePortFree gets a reference to the given string and assigns it to the DevicePortFree field.
func (o *DeviceDeviceDataData) SetDevicePortFree(v string) {
	o.DevicePortFree = &v
}

// GetDevicePortTotal returns the DevicePortTotal field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDevicePortTotal() string {
	if o == nil || o.DevicePortTotal == nil {
		var ret string
		return ret
	}
	return *o.DevicePortTotal
}

// GetDevicePortTotalOk returns a tuple with the DevicePortTotal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDevicePortTotalOk() (*string, bool) {
	if o == nil || o.DevicePortTotal == nil {
		return nil, false
	}
	return o.DevicePortTotal, true
}

// HasDevicePortTotal returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDevicePortTotal() bool {
	if o != nil && o.DevicePortTotal != nil {
		return true
	}

	return false
}

// SetDevicePortTotal gets a reference to the given string and assigns it to the DevicePortTotal field.
func (o *DeviceDeviceDataData) SetDevicePortTotal(v string) {
	o.DevicePortTotal = &v
}

// GetDevicePortUsed returns the DevicePortUsed field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDevicePortUsed() string {
	if o == nil || o.DevicePortUsed == nil {
		var ret string
		return ret
	}
	return *o.DevicePortUsed
}

// GetDevicePortUsedOk returns a tuple with the DevicePortUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDevicePortUsedOk() (*string, bool) {
	if o == nil || o.DevicePortUsed == nil {
		return nil, false
	}
	return o.DevicePortUsed, true
}

// HasDevicePortUsed returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDevicePortUsed() bool {
	if o != nil && o.DevicePortUsed != nil {
		return true
	}

	return false
}

// SetDevicePortUsed gets a reference to the given string and assigns it to the DevicePortUsed field.
func (o *DeviceDeviceDataData) SetDevicePortUsed(v string) {
	o.DevicePortUsed = &v
}

// GetDevicePortUsedPercent returns the DevicePortUsedPercent field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetDevicePortUsedPercent() string {
	if o == nil || o.DevicePortUsedPercent == nil {
		var ret string
		return ret
	}
	return *o.DevicePortUsedPercent
}

// GetDevicePortUsedPercentOk returns a tuple with the DevicePortUsedPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetDevicePortUsedPercentOk() (*string, bool) {
	if o == nil || o.DevicePortUsedPercent == nil {
		return nil, false
	}
	return o.DevicePortUsedPercent, true
}

// HasDevicePortUsedPercent returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasDevicePortUsedPercent() bool {
	if o != nil && o.DevicePortUsedPercent != nil {
		return true
	}

	return false
}

// SetDevicePortUsedPercent gets a reference to the given string and assigns it to the DevicePortUsedPercent field.
func (o *DeviceDeviceDataData) SetDevicePortUsedPercent(v string) {
	o.DevicePortUsedPercent = &v
}

// GetRowState returns the RowState field value if set, zero value otherwise.
func (o *DeviceDeviceDataData) GetRowState() string {
	if o == nil || o.RowState == nil {
		var ret string
		return ret
	}
	return *o.RowState
}

// GetRowStateOk returns a tuple with the RowState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDataData) GetRowStateOk() (*string, bool) {
	if o == nil || o.RowState == nil {
		return nil, false
	}
	return o.RowState, true
}

// HasRowState returns a boolean if a field has been set.
func (o *DeviceDeviceDataData) HasRowState() bool {
	if o != nil && o.RowState != nil {
		return true
	}

	return false
}

// SetRowState gets a reference to the given string and assigns it to the RowState field.
func (o *DeviceDeviceDataData) SetRowState(v string) {
	o.RowState = &v
}

func (o DeviceDeviceDataData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceClassName != nil {
		toSerialize["device_class_name"] = o.DeviceClassName
	}
	if o.DeviceClassParameters != nil {
		toSerialize["device_class_parameters"] = o.DeviceClassParameters
	}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	if o.DeviceAddressAddr != nil {
		toSerialize["device_address_addr"] = o.DeviceAddressAddr
	}
	if o.DeviceAddressFormated != nil {
		toSerialize["device_address_formated"] = o.DeviceAddressFormated
	}
	if o.DeviceName != nil {
		toSerialize["device_name"] = o.DeviceName
	}
	if o.DeviceSpaceId != nil {
		toSerialize["device_space_id"] = o.DeviceSpaceId
	}
	if o.DeviceSpaceName != nil {
		toSerialize["device_space_name"] = o.DeviceSpaceName
	}
	if o.InterfaceFree != nil {
		toSerialize["interface_free"] = o.InterfaceFree
	}
	if o.InterfaceTotal != nil {
		toSerialize["interface_total"] = o.InterfaceTotal
	}
	if o.InterfaceUsed != nil {
		toSerialize["interface_used"] = o.InterfaceUsed
	}
	if o.InterfaceUsedPercent != nil {
		toSerialize["interface_used_percent"] = o.InterfaceUsedPercent
	}
	if o.DevId != nil {
		toSerialize["dev_id"] = o.DevId
	}
	if o.DevName != nil {
		toSerialize["dev_name"] = o.DevName
	}
	if o.DevicePortFree != nil {
		toSerialize["device_port_free"] = o.DevicePortFree
	}
	if o.DevicePortTotal != nil {
		toSerialize["device_port_total"] = o.DevicePortTotal
	}
	if o.DevicePortUsed != nil {
		toSerialize["device_port_used"] = o.DevicePortUsed
	}
	if o.DevicePortUsedPercent != nil {
		toSerialize["device_port_used_percent"] = o.DevicePortUsedPercent
	}
	if o.RowState != nil {
		toSerialize["row_state"] = o.RowState
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceDeviceDataData struct {
	value *DeviceDeviceDataData
	isSet bool
}

func (v NullableDeviceDeviceDataData) Get() *DeviceDeviceDataData {
	return v.value
}

func (v *NullableDeviceDeviceDataData) Set(val *DeviceDeviceDataData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceDeviceDataData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceDeviceDataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceDeviceDataData(val *DeviceDeviceDataData) *NullableDeviceDeviceDataData {
	return &NullableDeviceDeviceDataData{value: val, isSet: true}
}

func (v NullableDeviceDeviceDataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceDeviceDataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

