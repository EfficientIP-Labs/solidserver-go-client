# DeviceDeviceDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceClassName** | Pointer to **string** | The name of the class applied to the device, it can be preceded by the class directory. | [optional] 
**DeviceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device. | [optional] 
**DeviceAddressAddr** | Pointer to **string** | The IP address associated with the Device Manager device. | [optional] 
**DeviceAddressFormated** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;hostdev_ip_addr&lt;/b&gt;. | [optional] 
**DeviceName** | Pointer to **string** | The name of the Device Manager device. | [optional] 
**DeviceSpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the Device Manager device. | [optional] 
**DeviceSpaceName** | Pointer to **string** | The name of the space associated with the Device Manager device. | [optional] 
**InterfaceFree** | Pointer to **string** | The number of interfaces on the Device Manager device that are currently free. | [optional] 
**InterfaceTotal** | Pointer to **string** | The total number of interfaces on the Device Manager device. | [optional] 
**InterfaceUsed** | Pointer to **string** | The number of interfaces on the Device Manager device that are currently active. | [optional] 
**InterfaceUsedPercent** | Pointer to **string** | The percentage of interfaces on the Device Manager device that are currently active. | [optional] 
**DevId** | Pointer to **string** | The database identifier (ID) of the NetChange network device associated with the Device Manager device. | [optional] 
**DevName** | Pointer to **string** | The name of the NetChange network device associated with the Device Manager device. | [optional] 
**DevicePortFree** | Pointer to **string** | The number of ports on the Device Manager device that are currently free. | [optional] 
**DevicePortTotal** | Pointer to **string** | The total number of ports on the Device Manager device. | [optional] 
**DevicePortUsed** | Pointer to **string** | The number of ports on the Device Manager device that are currently active. | [optional] 
**DevicePortUsedPercent** | Pointer to **string** | The percentage of ports on the Device Manager device that are currently active. | [optional] 
**RowState** | Pointer to **string** | The object activation status:&lt;ul class&#x3D;dashed &gt;&lt;li&gt;&lt;b&gt;0&lt;/b&gt; indicates the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;&lt;b&gt;1&lt;/b&gt; indicates the object is enabled and managed.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;&lt;b&gt;2&lt;/b&gt; indicates the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_enabled&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 

## Methods

### NewDeviceDeviceDataData

`func NewDeviceDeviceDataData() *DeviceDeviceDataData`

NewDeviceDeviceDataData instantiates a new DeviceDeviceDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceDeviceDataDataWithDefaults

`func NewDeviceDeviceDataDataWithDefaults() *DeviceDeviceDataData`

NewDeviceDeviceDataDataWithDefaults instantiates a new DeviceDeviceDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceClassName

`func (o *DeviceDeviceDataData) GetDeviceClassName() string`

GetDeviceClassName returns the DeviceClassName field if non-nil, zero value otherwise.

### GetDeviceClassNameOk

`func (o *DeviceDeviceDataData) GetDeviceClassNameOk() (*string, bool)`

GetDeviceClassNameOk returns a tuple with the DeviceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassName

`func (o *DeviceDeviceDataData) SetDeviceClassName(v string)`

SetDeviceClassName sets DeviceClassName field to given value.

### HasDeviceClassName

`func (o *DeviceDeviceDataData) HasDeviceClassName() bool`

HasDeviceClassName returns a boolean if a field has been set.

### GetDeviceClassParameters

`func (o *DeviceDeviceDataData) GetDeviceClassParameters() []ApiClassParameterOutputEntry`

GetDeviceClassParameters returns the DeviceClassParameters field if non-nil, zero value otherwise.

### GetDeviceClassParametersOk

`func (o *DeviceDeviceDataData) GetDeviceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetDeviceClassParametersOk returns a tuple with the DeviceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassParameters

`func (o *DeviceDeviceDataData) SetDeviceClassParameters(v []ApiClassParameterOutputEntry)`

SetDeviceClassParameters sets DeviceClassParameters field to given value.

### HasDeviceClassParameters

`func (o *DeviceDeviceDataData) HasDeviceClassParameters() bool`

HasDeviceClassParameters returns a boolean if a field has been set.

### GetDeviceId

`func (o *DeviceDeviceDataData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceDeviceDataData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceDeviceDataData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceDeviceDataData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceAddressAddr

`func (o *DeviceDeviceDataData) GetDeviceAddressAddr() string`

GetDeviceAddressAddr returns the DeviceAddressAddr field if non-nil, zero value otherwise.

### GetDeviceAddressAddrOk

`func (o *DeviceDeviceDataData) GetDeviceAddressAddrOk() (*string, bool)`

GetDeviceAddressAddrOk returns a tuple with the DeviceAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAddressAddr

`func (o *DeviceDeviceDataData) SetDeviceAddressAddr(v string)`

SetDeviceAddressAddr sets DeviceAddressAddr field to given value.

### HasDeviceAddressAddr

`func (o *DeviceDeviceDataData) HasDeviceAddressAddr() bool`

HasDeviceAddressAddr returns a boolean if a field has been set.

### GetDeviceAddressFormated

`func (o *DeviceDeviceDataData) GetDeviceAddressFormated() string`

GetDeviceAddressFormated returns the DeviceAddressFormated field if non-nil, zero value otherwise.

### GetDeviceAddressFormatedOk

`func (o *DeviceDeviceDataData) GetDeviceAddressFormatedOk() (*string, bool)`

GetDeviceAddressFormatedOk returns a tuple with the DeviceAddressFormated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAddressFormated

`func (o *DeviceDeviceDataData) SetDeviceAddressFormated(v string)`

SetDeviceAddressFormated sets DeviceAddressFormated field to given value.

### HasDeviceAddressFormated

`func (o *DeviceDeviceDataData) HasDeviceAddressFormated() bool`

HasDeviceAddressFormated returns a boolean if a field has been set.

### GetDeviceName

`func (o *DeviceDeviceDataData) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceDeviceDataData) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceDeviceDataData) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceDeviceDataData) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceSpaceId

`func (o *DeviceDeviceDataData) GetDeviceSpaceId() string`

GetDeviceSpaceId returns the DeviceSpaceId field if non-nil, zero value otherwise.

### GetDeviceSpaceIdOk

`func (o *DeviceDeviceDataData) GetDeviceSpaceIdOk() (*string, bool)`

GetDeviceSpaceIdOk returns a tuple with the DeviceSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSpaceId

`func (o *DeviceDeviceDataData) SetDeviceSpaceId(v string)`

SetDeviceSpaceId sets DeviceSpaceId field to given value.

### HasDeviceSpaceId

`func (o *DeviceDeviceDataData) HasDeviceSpaceId() bool`

HasDeviceSpaceId returns a boolean if a field has been set.

### GetDeviceSpaceName

`func (o *DeviceDeviceDataData) GetDeviceSpaceName() string`

GetDeviceSpaceName returns the DeviceSpaceName field if non-nil, zero value otherwise.

### GetDeviceSpaceNameOk

`func (o *DeviceDeviceDataData) GetDeviceSpaceNameOk() (*string, bool)`

GetDeviceSpaceNameOk returns a tuple with the DeviceSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSpaceName

`func (o *DeviceDeviceDataData) SetDeviceSpaceName(v string)`

SetDeviceSpaceName sets DeviceSpaceName field to given value.

### HasDeviceSpaceName

`func (o *DeviceDeviceDataData) HasDeviceSpaceName() bool`

HasDeviceSpaceName returns a boolean if a field has been set.

### GetInterfaceFree

`func (o *DeviceDeviceDataData) GetInterfaceFree() string`

GetInterfaceFree returns the InterfaceFree field if non-nil, zero value otherwise.

### GetInterfaceFreeOk

`func (o *DeviceDeviceDataData) GetInterfaceFreeOk() (*string, bool)`

GetInterfaceFreeOk returns a tuple with the InterfaceFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceFree

`func (o *DeviceDeviceDataData) SetInterfaceFree(v string)`

SetInterfaceFree sets InterfaceFree field to given value.

### HasInterfaceFree

`func (o *DeviceDeviceDataData) HasInterfaceFree() bool`

HasInterfaceFree returns a boolean if a field has been set.

### GetInterfaceTotal

`func (o *DeviceDeviceDataData) GetInterfaceTotal() string`

GetInterfaceTotal returns the InterfaceTotal field if non-nil, zero value otherwise.

### GetInterfaceTotalOk

`func (o *DeviceDeviceDataData) GetInterfaceTotalOk() (*string, bool)`

GetInterfaceTotalOk returns a tuple with the InterfaceTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceTotal

`func (o *DeviceDeviceDataData) SetInterfaceTotal(v string)`

SetInterfaceTotal sets InterfaceTotal field to given value.

### HasInterfaceTotal

`func (o *DeviceDeviceDataData) HasInterfaceTotal() bool`

HasInterfaceTotal returns a boolean if a field has been set.

### GetInterfaceUsed

`func (o *DeviceDeviceDataData) GetInterfaceUsed() string`

GetInterfaceUsed returns the InterfaceUsed field if non-nil, zero value otherwise.

### GetInterfaceUsedOk

`func (o *DeviceDeviceDataData) GetInterfaceUsedOk() (*string, bool)`

GetInterfaceUsedOk returns a tuple with the InterfaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUsed

`func (o *DeviceDeviceDataData) SetInterfaceUsed(v string)`

SetInterfaceUsed sets InterfaceUsed field to given value.

### HasInterfaceUsed

`func (o *DeviceDeviceDataData) HasInterfaceUsed() bool`

HasInterfaceUsed returns a boolean if a field has been set.

### GetInterfaceUsedPercent

`func (o *DeviceDeviceDataData) GetInterfaceUsedPercent() string`

GetInterfaceUsedPercent returns the InterfaceUsedPercent field if non-nil, zero value otherwise.

### GetInterfaceUsedPercentOk

`func (o *DeviceDeviceDataData) GetInterfaceUsedPercentOk() (*string, bool)`

GetInterfaceUsedPercentOk returns a tuple with the InterfaceUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUsedPercent

`func (o *DeviceDeviceDataData) SetInterfaceUsedPercent(v string)`

SetInterfaceUsedPercent sets InterfaceUsedPercent field to given value.

### HasInterfaceUsedPercent

`func (o *DeviceDeviceDataData) HasInterfaceUsedPercent() bool`

HasInterfaceUsedPercent returns a boolean if a field has been set.

### GetDevId

`func (o *DeviceDeviceDataData) GetDevId() string`

GetDevId returns the DevId field if non-nil, zero value otherwise.

### GetDevIdOk

`func (o *DeviceDeviceDataData) GetDevIdOk() (*string, bool)`

GetDevIdOk returns a tuple with the DevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevId

`func (o *DeviceDeviceDataData) SetDevId(v string)`

SetDevId sets DevId field to given value.

### HasDevId

`func (o *DeviceDeviceDataData) HasDevId() bool`

HasDevId returns a boolean if a field has been set.

### GetDevName

`func (o *DeviceDeviceDataData) GetDevName() string`

GetDevName returns the DevName field if non-nil, zero value otherwise.

### GetDevNameOk

`func (o *DeviceDeviceDataData) GetDevNameOk() (*string, bool)`

GetDevNameOk returns a tuple with the DevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevName

`func (o *DeviceDeviceDataData) SetDevName(v string)`

SetDevName sets DevName field to given value.

### HasDevName

`func (o *DeviceDeviceDataData) HasDevName() bool`

HasDevName returns a boolean if a field has been set.

### GetDevicePortFree

`func (o *DeviceDeviceDataData) GetDevicePortFree() string`

GetDevicePortFree returns the DevicePortFree field if non-nil, zero value otherwise.

### GetDevicePortFreeOk

`func (o *DeviceDeviceDataData) GetDevicePortFreeOk() (*string, bool)`

GetDevicePortFreeOk returns a tuple with the DevicePortFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortFree

`func (o *DeviceDeviceDataData) SetDevicePortFree(v string)`

SetDevicePortFree sets DevicePortFree field to given value.

### HasDevicePortFree

`func (o *DeviceDeviceDataData) HasDevicePortFree() bool`

HasDevicePortFree returns a boolean if a field has been set.

### GetDevicePortTotal

`func (o *DeviceDeviceDataData) GetDevicePortTotal() string`

GetDevicePortTotal returns the DevicePortTotal field if non-nil, zero value otherwise.

### GetDevicePortTotalOk

`func (o *DeviceDeviceDataData) GetDevicePortTotalOk() (*string, bool)`

GetDevicePortTotalOk returns a tuple with the DevicePortTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortTotal

`func (o *DeviceDeviceDataData) SetDevicePortTotal(v string)`

SetDevicePortTotal sets DevicePortTotal field to given value.

### HasDevicePortTotal

`func (o *DeviceDeviceDataData) HasDevicePortTotal() bool`

HasDevicePortTotal returns a boolean if a field has been set.

### GetDevicePortUsed

`func (o *DeviceDeviceDataData) GetDevicePortUsed() string`

GetDevicePortUsed returns the DevicePortUsed field if non-nil, zero value otherwise.

### GetDevicePortUsedOk

`func (o *DeviceDeviceDataData) GetDevicePortUsedOk() (*string, bool)`

GetDevicePortUsedOk returns a tuple with the DevicePortUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortUsed

`func (o *DeviceDeviceDataData) SetDevicePortUsed(v string)`

SetDevicePortUsed sets DevicePortUsed field to given value.

### HasDevicePortUsed

`func (o *DeviceDeviceDataData) HasDevicePortUsed() bool`

HasDevicePortUsed returns a boolean if a field has been set.

### GetDevicePortUsedPercent

`func (o *DeviceDeviceDataData) GetDevicePortUsedPercent() string`

GetDevicePortUsedPercent returns the DevicePortUsedPercent field if non-nil, zero value otherwise.

### GetDevicePortUsedPercentOk

`func (o *DeviceDeviceDataData) GetDevicePortUsedPercentOk() (*string, bool)`

GetDevicePortUsedPercentOk returns a tuple with the DevicePortUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortUsedPercent

`func (o *DeviceDeviceDataData) SetDevicePortUsedPercent(v string)`

SetDevicePortUsedPercent sets DevicePortUsedPercent field to given value.

### HasDevicePortUsedPercent

`func (o *DeviceDeviceDataData) HasDevicePortUsedPercent() bool`

HasDevicePortUsedPercent returns a boolean if a field has been set.

### GetRowState

`func (o *DeviceDeviceDataData) GetRowState() string`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *DeviceDeviceDataData) GetRowStateOk() (*string, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *DeviceDeviceDataData) SetRowState(v string)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *DeviceDeviceDataData) HasRowState() bool`

HasRowState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


