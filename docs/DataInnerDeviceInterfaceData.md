# DataInnerDeviceInterfaceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InterfaceAddTime** | Pointer to **string** | The time at which the Device Manager port or interface has been added, in decimal UNIX date format. | [optional] 
**DataId** | Pointer to **string** | The database identifier (ID) of the Custom Database entry associated with the Device Manager port or interface. | [optional] 
**DeviceClassName** | Pointer to **string** | The name of the class applied to the device the object belongs to, it can be preceded by the class directory. | [optional] 
**DeviceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the Device Manager device the object belongs to. | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device the object belongs to. | [optional] 
**DeviceName** | Pointer to **string** | The name of the Device Manager device. | [optional] 
**InterfaceAutoLink** | Pointer to **string** | The DM device and port or interface to which the object is automatically linked, as follows: &lt;b&gt;&amp;lt;hostdev_name&amp;gt; ( &amp;lt;hostiface_name&amp;gt; )&lt;/b&gt; | [optional] 
**InterfaceClassName** | Pointer to **string** | The name of the class applied to the port or interface, it can be preceded by the class directory. | [optional] 
**InterfaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the port or interface. | [optional] 
**InterfaceId** | Pointer to **string** | The database identifier (ID) of the Device Manager port or interface. | [optional] 
**InterfaceAddressAddr** | Pointer to **string** | The IP address associated with the Device Manager port or interface. | [optional] 
**InterfaceAddressFormated** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;interface_address_addr&lt;/b&gt;. | [optional] 
**InterfaceMac** | Pointer to **string** | The MAC address associated with the Device Manager port or interface. | [optional] 
**InterfaceManualLink** | Pointer to **string** | The DM device and port or interface to which the object is manually linked, as follows: &lt;b&gt;&amp;lt;hostdev_name&amp;gt; ( &amp;lt;hostiface_name&amp;gt; )&lt;/b&gt; | [optional] 
**InterfaceName** | Pointer to **string** | The name of the Device Manager port or interface. | [optional] 
**InterfaceSpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the Device Manager port or interface. | [optional] 
**InterfaceSpaceName** | Pointer to **string** | The name of the space associated with the Device Manager port or interface. | [optional] 
**InterfaceType** | Pointer to **string** | A way to indicate if the object is either a &lt;b&gt;port&lt;/b&gt; or an &lt;b&gt;interface&lt;/b&gt;. | [optional] 
**DevEnabled** | Pointer to **string** | The activation status of the NetChange network device associated with the DM port or interface.&lt;ul class&#x3D;dashed &gt;&lt;li&gt; If set to &lt;b&gt;1&lt;/b&gt;, the NetChange network device is enabled and managed.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; If set to &lt;b&gt;2&lt;/b&gt;, the NetChange network device is unmanaged, disabled or both depending on the context.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**DevId** | Pointer to **string** | The database identifier (ID) of the NetChange network device associated with the Device Manager device the object belongs to. | [optional] 
**DevName** | Pointer to **string** | The name of the NetChange network device associated with the Device Manager device the object belongs to. | [optional] 
**PortEnabled** | Pointer to **string** | The activation status of the NetChange port associated with the DM port or interface.&lt;ul class&#x3D;dashed &gt;&lt;li&gt; If set to &lt;b&gt;1&lt;/b&gt;, the NetChange port is enabled and managed.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; If set to &lt;b&gt;2&lt;/b&gt;, the NetChange port is unmanaged, disabled or both depending on the context.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**PortId** | Pointer to **string** | The database identifier (ID) of the NetChange port associated with the Device Manager port or interface. | [optional] 
**PortName** | Pointer to **string** | The name of the NetChange port associated with the DM port or interface. | [optional] 
**InterfaceModifyTime** | Pointer to **string** | The last time the Device Manager port or interface data was reconciled, in decimal UNIX date format. | [optional] 
**InterfaceNbIp** | Pointer to **string** | The number of IP addresses associated with the Device Manager port or interface. | [optional] 
**PearIfaceId** | Pointer to **string** | The database identifier (ID) of the Device Manager port or interface linked with &lt;b&gt;interface_id&lt;/b&gt;. | [optional] 
**PearIplIfaceId** | Pointer to **string** | The database identifier (ID) of the NetChange port linked with &lt;b&gt;interface_id&lt;/b&gt;. | [optional] 
**RowState** | Pointer to **string** | The object activation status:&lt;ul class&#x3D;dashed &gt;&lt;li&gt; &lt;b&gt;0&lt;/b&gt; indicates the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; &lt;b&gt;1&lt;/b&gt; indicates the object is enabled and managed.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; &lt;b&gt;2&lt;/b&gt; indicates the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_state&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 
**InterfaceVendorKey** | Pointer to **string** | Internal use. Not documented. | [optional] 
**InterfaceVendorMac** | Pointer to **string** | The vendor details of the client associated with the Device Manager port or interface. | [optional] 

## Methods

### NewDataInnerDeviceInterfaceData

`func NewDataInnerDeviceInterfaceData() *DataInnerDeviceInterfaceData`

NewDataInnerDeviceInterfaceData instantiates a new DataInnerDeviceInterfaceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDeviceInterfaceDataWithDefaults

`func NewDataInnerDeviceInterfaceDataWithDefaults() *DataInnerDeviceInterfaceData`

NewDataInnerDeviceInterfaceDataWithDefaults instantiates a new DataInnerDeviceInterfaceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInterfaceAddTime

`func (o *DataInnerDeviceInterfaceData) GetInterfaceAddTime() string`

GetInterfaceAddTime returns the InterfaceAddTime field if non-nil, zero value otherwise.

### GetInterfaceAddTimeOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceAddTimeOk() (*string, bool)`

GetInterfaceAddTimeOk returns a tuple with the InterfaceAddTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceAddTime

`func (o *DataInnerDeviceInterfaceData) SetInterfaceAddTime(v string)`

SetInterfaceAddTime sets InterfaceAddTime field to given value.

### HasInterfaceAddTime

`func (o *DataInnerDeviceInterfaceData) HasInterfaceAddTime() bool`

HasInterfaceAddTime returns a boolean if a field has been set.

### GetDataId

`func (o *DataInnerDeviceInterfaceData) GetDataId() string`

GetDataId returns the DataId field if non-nil, zero value otherwise.

### GetDataIdOk

`func (o *DataInnerDeviceInterfaceData) GetDataIdOk() (*string, bool)`

GetDataIdOk returns a tuple with the DataId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataId

`func (o *DataInnerDeviceInterfaceData) SetDataId(v string)`

SetDataId sets DataId field to given value.

### HasDataId

`func (o *DataInnerDeviceInterfaceData) HasDataId() bool`

HasDataId returns a boolean if a field has been set.

### GetDeviceClassName

`func (o *DataInnerDeviceInterfaceData) GetDeviceClassName() string`

GetDeviceClassName returns the DeviceClassName field if non-nil, zero value otherwise.

### GetDeviceClassNameOk

`func (o *DataInnerDeviceInterfaceData) GetDeviceClassNameOk() (*string, bool)`

GetDeviceClassNameOk returns a tuple with the DeviceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassName

`func (o *DataInnerDeviceInterfaceData) SetDeviceClassName(v string)`

SetDeviceClassName sets DeviceClassName field to given value.

### HasDeviceClassName

`func (o *DataInnerDeviceInterfaceData) HasDeviceClassName() bool`

HasDeviceClassName returns a boolean if a field has been set.

### GetDeviceClassParameters

`func (o *DataInnerDeviceInterfaceData) GetDeviceClassParameters() []ApiClassParameterOutputEntry`

GetDeviceClassParameters returns the DeviceClassParameters field if non-nil, zero value otherwise.

### GetDeviceClassParametersOk

`func (o *DataInnerDeviceInterfaceData) GetDeviceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetDeviceClassParametersOk returns a tuple with the DeviceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassParameters

`func (o *DataInnerDeviceInterfaceData) SetDeviceClassParameters(v []ApiClassParameterOutputEntry)`

SetDeviceClassParameters sets DeviceClassParameters field to given value.

### HasDeviceClassParameters

`func (o *DataInnerDeviceInterfaceData) HasDeviceClassParameters() bool`

HasDeviceClassParameters returns a boolean if a field has been set.

### GetDeviceId

`func (o *DataInnerDeviceInterfaceData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DataInnerDeviceInterfaceData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DataInnerDeviceInterfaceData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DataInnerDeviceInterfaceData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *DataInnerDeviceInterfaceData) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DataInnerDeviceInterfaceData) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DataInnerDeviceInterfaceData) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DataInnerDeviceInterfaceData) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetInterfaceAutoLink

`func (o *DataInnerDeviceInterfaceData) GetInterfaceAutoLink() string`

GetInterfaceAutoLink returns the InterfaceAutoLink field if non-nil, zero value otherwise.

### GetInterfaceAutoLinkOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceAutoLinkOk() (*string, bool)`

GetInterfaceAutoLinkOk returns a tuple with the InterfaceAutoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceAutoLink

`func (o *DataInnerDeviceInterfaceData) SetInterfaceAutoLink(v string)`

SetInterfaceAutoLink sets InterfaceAutoLink field to given value.

### HasInterfaceAutoLink

`func (o *DataInnerDeviceInterfaceData) HasInterfaceAutoLink() bool`

HasInterfaceAutoLink returns a boolean if a field has been set.

### GetInterfaceClassName

`func (o *DataInnerDeviceInterfaceData) GetInterfaceClassName() string`

GetInterfaceClassName returns the InterfaceClassName field if non-nil, zero value otherwise.

### GetInterfaceClassNameOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceClassNameOk() (*string, bool)`

GetInterfaceClassNameOk returns a tuple with the InterfaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceClassName

`func (o *DataInnerDeviceInterfaceData) SetInterfaceClassName(v string)`

SetInterfaceClassName sets InterfaceClassName field to given value.

### HasInterfaceClassName

`func (o *DataInnerDeviceInterfaceData) HasInterfaceClassName() bool`

HasInterfaceClassName returns a boolean if a field has been set.

### GetInterfaceClassParameters

`func (o *DataInnerDeviceInterfaceData) GetInterfaceClassParameters() []ApiClassParameterOutputEntry`

GetInterfaceClassParameters returns the InterfaceClassParameters field if non-nil, zero value otherwise.

### GetInterfaceClassParametersOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetInterfaceClassParametersOk returns a tuple with the InterfaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceClassParameters

`func (o *DataInnerDeviceInterfaceData) SetInterfaceClassParameters(v []ApiClassParameterOutputEntry)`

SetInterfaceClassParameters sets InterfaceClassParameters field to given value.

### HasInterfaceClassParameters

`func (o *DataInnerDeviceInterfaceData) HasInterfaceClassParameters() bool`

HasInterfaceClassParameters returns a boolean if a field has been set.

### GetInterfaceId

`func (o *DataInnerDeviceInterfaceData) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *DataInnerDeviceInterfaceData) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *DataInnerDeviceInterfaceData) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceAddressAddr

`func (o *DataInnerDeviceInterfaceData) GetInterfaceAddressAddr() string`

GetInterfaceAddressAddr returns the InterfaceAddressAddr field if non-nil, zero value otherwise.

### GetInterfaceAddressAddrOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceAddressAddrOk() (*string, bool)`

GetInterfaceAddressAddrOk returns a tuple with the InterfaceAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceAddressAddr

`func (o *DataInnerDeviceInterfaceData) SetInterfaceAddressAddr(v string)`

SetInterfaceAddressAddr sets InterfaceAddressAddr field to given value.

### HasInterfaceAddressAddr

`func (o *DataInnerDeviceInterfaceData) HasInterfaceAddressAddr() bool`

HasInterfaceAddressAddr returns a boolean if a field has been set.

### GetInterfaceAddressFormated

`func (o *DataInnerDeviceInterfaceData) GetInterfaceAddressFormated() string`

GetInterfaceAddressFormated returns the InterfaceAddressFormated field if non-nil, zero value otherwise.

### GetInterfaceAddressFormatedOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceAddressFormatedOk() (*string, bool)`

GetInterfaceAddressFormatedOk returns a tuple with the InterfaceAddressFormated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceAddressFormated

`func (o *DataInnerDeviceInterfaceData) SetInterfaceAddressFormated(v string)`

SetInterfaceAddressFormated sets InterfaceAddressFormated field to given value.

### HasInterfaceAddressFormated

`func (o *DataInnerDeviceInterfaceData) HasInterfaceAddressFormated() bool`

HasInterfaceAddressFormated returns a boolean if a field has been set.

### GetInterfaceMac

`func (o *DataInnerDeviceInterfaceData) GetInterfaceMac() string`

GetInterfaceMac returns the InterfaceMac field if non-nil, zero value otherwise.

### GetInterfaceMacOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceMacOk() (*string, bool)`

GetInterfaceMacOk returns a tuple with the InterfaceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceMac

`func (o *DataInnerDeviceInterfaceData) SetInterfaceMac(v string)`

SetInterfaceMac sets InterfaceMac field to given value.

### HasInterfaceMac

`func (o *DataInnerDeviceInterfaceData) HasInterfaceMac() bool`

HasInterfaceMac returns a boolean if a field has been set.

### GetInterfaceManualLink

`func (o *DataInnerDeviceInterfaceData) GetInterfaceManualLink() string`

GetInterfaceManualLink returns the InterfaceManualLink field if non-nil, zero value otherwise.

### GetInterfaceManualLinkOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceManualLinkOk() (*string, bool)`

GetInterfaceManualLinkOk returns a tuple with the InterfaceManualLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceManualLink

`func (o *DataInnerDeviceInterfaceData) SetInterfaceManualLink(v string)`

SetInterfaceManualLink sets InterfaceManualLink field to given value.

### HasInterfaceManualLink

`func (o *DataInnerDeviceInterfaceData) HasInterfaceManualLink() bool`

HasInterfaceManualLink returns a boolean if a field has been set.

### GetInterfaceName

`func (o *DataInnerDeviceInterfaceData) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *DataInnerDeviceInterfaceData) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *DataInnerDeviceInterfaceData) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceSpaceId

`func (o *DataInnerDeviceInterfaceData) GetInterfaceSpaceId() string`

GetInterfaceSpaceId returns the InterfaceSpaceId field if non-nil, zero value otherwise.

### GetInterfaceSpaceIdOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceSpaceIdOk() (*string, bool)`

GetInterfaceSpaceIdOk returns a tuple with the InterfaceSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceSpaceId

`func (o *DataInnerDeviceInterfaceData) SetInterfaceSpaceId(v string)`

SetInterfaceSpaceId sets InterfaceSpaceId field to given value.

### HasInterfaceSpaceId

`func (o *DataInnerDeviceInterfaceData) HasInterfaceSpaceId() bool`

HasInterfaceSpaceId returns a boolean if a field has been set.

### GetInterfaceSpaceName

`func (o *DataInnerDeviceInterfaceData) GetInterfaceSpaceName() string`

GetInterfaceSpaceName returns the InterfaceSpaceName field if non-nil, zero value otherwise.

### GetInterfaceSpaceNameOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceSpaceNameOk() (*string, bool)`

GetInterfaceSpaceNameOk returns a tuple with the InterfaceSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceSpaceName

`func (o *DataInnerDeviceInterfaceData) SetInterfaceSpaceName(v string)`

SetInterfaceSpaceName sets InterfaceSpaceName field to given value.

### HasInterfaceSpaceName

`func (o *DataInnerDeviceInterfaceData) HasInterfaceSpaceName() bool`

HasInterfaceSpaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *DataInnerDeviceInterfaceData) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *DataInnerDeviceInterfaceData) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *DataInnerDeviceInterfaceData) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetDevEnabled

`func (o *DataInnerDeviceInterfaceData) GetDevEnabled() string`

GetDevEnabled returns the DevEnabled field if non-nil, zero value otherwise.

### GetDevEnabledOk

`func (o *DataInnerDeviceInterfaceData) GetDevEnabledOk() (*string, bool)`

GetDevEnabledOk returns a tuple with the DevEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevEnabled

`func (o *DataInnerDeviceInterfaceData) SetDevEnabled(v string)`

SetDevEnabled sets DevEnabled field to given value.

### HasDevEnabled

`func (o *DataInnerDeviceInterfaceData) HasDevEnabled() bool`

HasDevEnabled returns a boolean if a field has been set.

### GetDevId

`func (o *DataInnerDeviceInterfaceData) GetDevId() string`

GetDevId returns the DevId field if non-nil, zero value otherwise.

### GetDevIdOk

`func (o *DataInnerDeviceInterfaceData) GetDevIdOk() (*string, bool)`

GetDevIdOk returns a tuple with the DevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevId

`func (o *DataInnerDeviceInterfaceData) SetDevId(v string)`

SetDevId sets DevId field to given value.

### HasDevId

`func (o *DataInnerDeviceInterfaceData) HasDevId() bool`

HasDevId returns a boolean if a field has been set.

### GetDevName

`func (o *DataInnerDeviceInterfaceData) GetDevName() string`

GetDevName returns the DevName field if non-nil, zero value otherwise.

### GetDevNameOk

`func (o *DataInnerDeviceInterfaceData) GetDevNameOk() (*string, bool)`

GetDevNameOk returns a tuple with the DevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevName

`func (o *DataInnerDeviceInterfaceData) SetDevName(v string)`

SetDevName sets DevName field to given value.

### HasDevName

`func (o *DataInnerDeviceInterfaceData) HasDevName() bool`

HasDevName returns a boolean if a field has been set.

### GetPortEnabled

`func (o *DataInnerDeviceInterfaceData) GetPortEnabled() string`

GetPortEnabled returns the PortEnabled field if non-nil, zero value otherwise.

### GetPortEnabledOk

`func (o *DataInnerDeviceInterfaceData) GetPortEnabledOk() (*string, bool)`

GetPortEnabledOk returns a tuple with the PortEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortEnabled

`func (o *DataInnerDeviceInterfaceData) SetPortEnabled(v string)`

SetPortEnabled sets PortEnabled field to given value.

### HasPortEnabled

`func (o *DataInnerDeviceInterfaceData) HasPortEnabled() bool`

HasPortEnabled returns a boolean if a field has been set.

### GetPortId

`func (o *DataInnerDeviceInterfaceData) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *DataInnerDeviceInterfaceData) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *DataInnerDeviceInterfaceData) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *DataInnerDeviceInterfaceData) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPortName

`func (o *DataInnerDeviceInterfaceData) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *DataInnerDeviceInterfaceData) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *DataInnerDeviceInterfaceData) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *DataInnerDeviceInterfaceData) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetInterfaceModifyTime

`func (o *DataInnerDeviceInterfaceData) GetInterfaceModifyTime() string`

GetInterfaceModifyTime returns the InterfaceModifyTime field if non-nil, zero value otherwise.

### GetInterfaceModifyTimeOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceModifyTimeOk() (*string, bool)`

GetInterfaceModifyTimeOk returns a tuple with the InterfaceModifyTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceModifyTime

`func (o *DataInnerDeviceInterfaceData) SetInterfaceModifyTime(v string)`

SetInterfaceModifyTime sets InterfaceModifyTime field to given value.

### HasInterfaceModifyTime

`func (o *DataInnerDeviceInterfaceData) HasInterfaceModifyTime() bool`

HasInterfaceModifyTime returns a boolean if a field has been set.

### GetInterfaceNbIp

`func (o *DataInnerDeviceInterfaceData) GetInterfaceNbIp() string`

GetInterfaceNbIp returns the InterfaceNbIp field if non-nil, zero value otherwise.

### GetInterfaceNbIpOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceNbIpOk() (*string, bool)`

GetInterfaceNbIpOk returns a tuple with the InterfaceNbIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceNbIp

`func (o *DataInnerDeviceInterfaceData) SetInterfaceNbIp(v string)`

SetInterfaceNbIp sets InterfaceNbIp field to given value.

### HasInterfaceNbIp

`func (o *DataInnerDeviceInterfaceData) HasInterfaceNbIp() bool`

HasInterfaceNbIp returns a boolean if a field has been set.

### GetPearIfaceId

`func (o *DataInnerDeviceInterfaceData) GetPearIfaceId() string`

GetPearIfaceId returns the PearIfaceId field if non-nil, zero value otherwise.

### GetPearIfaceIdOk

`func (o *DataInnerDeviceInterfaceData) GetPearIfaceIdOk() (*string, bool)`

GetPearIfaceIdOk returns a tuple with the PearIfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPearIfaceId

`func (o *DataInnerDeviceInterfaceData) SetPearIfaceId(v string)`

SetPearIfaceId sets PearIfaceId field to given value.

### HasPearIfaceId

`func (o *DataInnerDeviceInterfaceData) HasPearIfaceId() bool`

HasPearIfaceId returns a boolean if a field has been set.

### GetPearIplIfaceId

`func (o *DataInnerDeviceInterfaceData) GetPearIplIfaceId() string`

GetPearIplIfaceId returns the PearIplIfaceId field if non-nil, zero value otherwise.

### GetPearIplIfaceIdOk

`func (o *DataInnerDeviceInterfaceData) GetPearIplIfaceIdOk() (*string, bool)`

GetPearIplIfaceIdOk returns a tuple with the PearIplIfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPearIplIfaceId

`func (o *DataInnerDeviceInterfaceData) SetPearIplIfaceId(v string)`

SetPearIplIfaceId sets PearIplIfaceId field to given value.

### HasPearIplIfaceId

`func (o *DataInnerDeviceInterfaceData) HasPearIplIfaceId() bool`

HasPearIplIfaceId returns a boolean if a field has been set.

### GetRowState

`func (o *DataInnerDeviceInterfaceData) GetRowState() string`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *DataInnerDeviceInterfaceData) GetRowStateOk() (*string, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *DataInnerDeviceInterfaceData) SetRowState(v string)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *DataInnerDeviceInterfaceData) HasRowState() bool`

HasRowState returns a boolean if a field has been set.

### GetInterfaceVendorKey

`func (o *DataInnerDeviceInterfaceData) GetInterfaceVendorKey() string`

GetInterfaceVendorKey returns the InterfaceVendorKey field if non-nil, zero value otherwise.

### GetInterfaceVendorKeyOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceVendorKeyOk() (*string, bool)`

GetInterfaceVendorKeyOk returns a tuple with the InterfaceVendorKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceVendorKey

`func (o *DataInnerDeviceInterfaceData) SetInterfaceVendorKey(v string)`

SetInterfaceVendorKey sets InterfaceVendorKey field to given value.

### HasInterfaceVendorKey

`func (o *DataInnerDeviceInterfaceData) HasInterfaceVendorKey() bool`

HasInterfaceVendorKey returns a boolean if a field has been set.

### GetInterfaceVendorMac

`func (o *DataInnerDeviceInterfaceData) GetInterfaceVendorMac() string`

GetInterfaceVendorMac returns the InterfaceVendorMac field if non-nil, zero value otherwise.

### GetInterfaceVendorMacOk

`func (o *DataInnerDeviceInterfaceData) GetInterfaceVendorMacOk() (*string, bool)`

GetInterfaceVendorMacOk returns a tuple with the InterfaceVendorMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceVendorMac

`func (o *DataInnerDeviceInterfaceData) SetInterfaceVendorMac(v string)`

SetInterfaceVendorMac sets InterfaceVendorMac field to given value.

### HasInterfaceVendorMac

`func (o *DataInnerDeviceInterfaceData) HasInterfaceVendorMac() bool`

HasInterfaceVendorMac returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


