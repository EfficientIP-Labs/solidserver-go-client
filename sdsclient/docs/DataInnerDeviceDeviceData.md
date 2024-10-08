# DataInnerDeviceDeviceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceClassName** | Pointer to **string** | The name of the class applied to the device, it can be preceded by the class directory. | [optional] 
**DeviceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the device. | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device. | [optional] 
**DeviceAddressAddr** | Pointer to **string** | The IP address associated with the Device Manager device. | [optional] 
**DeviceAddressFormated** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;device_address_addr&lt;/b&gt;. | [optional] 
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
**RowState** | Pointer to **string** | The object activation status:&lt;ul class&#x3D;dashed &gt;&lt;li&gt; &lt;b&gt;0&lt;/b&gt; indicates the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; &lt;b&gt;1&lt;/b&gt; indicates the object is enabled and managed.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; &lt;b&gt;2&lt;/b&gt; indicates the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_state&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 

## Methods

### NewDataInnerDeviceDeviceData

`func NewDataInnerDeviceDeviceData() *DataInnerDeviceDeviceData`

NewDataInnerDeviceDeviceData instantiates a new DataInnerDeviceDeviceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDeviceDeviceDataWithDefaults

`func NewDataInnerDeviceDeviceDataWithDefaults() *DataInnerDeviceDeviceData`

NewDataInnerDeviceDeviceDataWithDefaults instantiates a new DataInnerDeviceDeviceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceClassName

`func (o *DataInnerDeviceDeviceData) GetDeviceClassName() string`

GetDeviceClassName returns the DeviceClassName field if non-nil, zero value otherwise.

### GetDeviceClassNameOk

`func (o *DataInnerDeviceDeviceData) GetDeviceClassNameOk() (*string, bool)`

GetDeviceClassNameOk returns a tuple with the DeviceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassName

`func (o *DataInnerDeviceDeviceData) SetDeviceClassName(v string)`

SetDeviceClassName sets DeviceClassName field to given value.

### HasDeviceClassName

`func (o *DataInnerDeviceDeviceData) HasDeviceClassName() bool`

HasDeviceClassName returns a boolean if a field has been set.

### GetDeviceClassParameters

`func (o *DataInnerDeviceDeviceData) GetDeviceClassParameters() []ApiClassParameterOutputEntry`

GetDeviceClassParameters returns the DeviceClassParameters field if non-nil, zero value otherwise.

### GetDeviceClassParametersOk

`func (o *DataInnerDeviceDeviceData) GetDeviceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetDeviceClassParametersOk returns a tuple with the DeviceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceClassParameters

`func (o *DataInnerDeviceDeviceData) SetDeviceClassParameters(v []ApiClassParameterOutputEntry)`

SetDeviceClassParameters sets DeviceClassParameters field to given value.

### HasDeviceClassParameters

`func (o *DataInnerDeviceDeviceData) HasDeviceClassParameters() bool`

HasDeviceClassParameters returns a boolean if a field has been set.

### GetDeviceId

`func (o *DataInnerDeviceDeviceData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DataInnerDeviceDeviceData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DataInnerDeviceDeviceData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DataInnerDeviceDeviceData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceAddressAddr

`func (o *DataInnerDeviceDeviceData) GetDeviceAddressAddr() string`

GetDeviceAddressAddr returns the DeviceAddressAddr field if non-nil, zero value otherwise.

### GetDeviceAddressAddrOk

`func (o *DataInnerDeviceDeviceData) GetDeviceAddressAddrOk() (*string, bool)`

GetDeviceAddressAddrOk returns a tuple with the DeviceAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAddressAddr

`func (o *DataInnerDeviceDeviceData) SetDeviceAddressAddr(v string)`

SetDeviceAddressAddr sets DeviceAddressAddr field to given value.

### HasDeviceAddressAddr

`func (o *DataInnerDeviceDeviceData) HasDeviceAddressAddr() bool`

HasDeviceAddressAddr returns a boolean if a field has been set.

### GetDeviceAddressFormated

`func (o *DataInnerDeviceDeviceData) GetDeviceAddressFormated() string`

GetDeviceAddressFormated returns the DeviceAddressFormated field if non-nil, zero value otherwise.

### GetDeviceAddressFormatedOk

`func (o *DataInnerDeviceDeviceData) GetDeviceAddressFormatedOk() (*string, bool)`

GetDeviceAddressFormatedOk returns a tuple with the DeviceAddressFormated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceAddressFormated

`func (o *DataInnerDeviceDeviceData) SetDeviceAddressFormated(v string)`

SetDeviceAddressFormated sets DeviceAddressFormated field to given value.

### HasDeviceAddressFormated

`func (o *DataInnerDeviceDeviceData) HasDeviceAddressFormated() bool`

HasDeviceAddressFormated returns a boolean if a field has been set.

### GetDeviceName

`func (o *DataInnerDeviceDeviceData) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DataInnerDeviceDeviceData) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DataInnerDeviceDeviceData) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DataInnerDeviceDeviceData) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetDeviceSpaceId

`func (o *DataInnerDeviceDeviceData) GetDeviceSpaceId() string`

GetDeviceSpaceId returns the DeviceSpaceId field if non-nil, zero value otherwise.

### GetDeviceSpaceIdOk

`func (o *DataInnerDeviceDeviceData) GetDeviceSpaceIdOk() (*string, bool)`

GetDeviceSpaceIdOk returns a tuple with the DeviceSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSpaceId

`func (o *DataInnerDeviceDeviceData) SetDeviceSpaceId(v string)`

SetDeviceSpaceId sets DeviceSpaceId field to given value.

### HasDeviceSpaceId

`func (o *DataInnerDeviceDeviceData) HasDeviceSpaceId() bool`

HasDeviceSpaceId returns a boolean if a field has been set.

### GetDeviceSpaceName

`func (o *DataInnerDeviceDeviceData) GetDeviceSpaceName() string`

GetDeviceSpaceName returns the DeviceSpaceName field if non-nil, zero value otherwise.

### GetDeviceSpaceNameOk

`func (o *DataInnerDeviceDeviceData) GetDeviceSpaceNameOk() (*string, bool)`

GetDeviceSpaceNameOk returns a tuple with the DeviceSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSpaceName

`func (o *DataInnerDeviceDeviceData) SetDeviceSpaceName(v string)`

SetDeviceSpaceName sets DeviceSpaceName field to given value.

### HasDeviceSpaceName

`func (o *DataInnerDeviceDeviceData) HasDeviceSpaceName() bool`

HasDeviceSpaceName returns a boolean if a field has been set.

### GetInterfaceFree

`func (o *DataInnerDeviceDeviceData) GetInterfaceFree() string`

GetInterfaceFree returns the InterfaceFree field if non-nil, zero value otherwise.

### GetInterfaceFreeOk

`func (o *DataInnerDeviceDeviceData) GetInterfaceFreeOk() (*string, bool)`

GetInterfaceFreeOk returns a tuple with the InterfaceFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceFree

`func (o *DataInnerDeviceDeviceData) SetInterfaceFree(v string)`

SetInterfaceFree sets InterfaceFree field to given value.

### HasInterfaceFree

`func (o *DataInnerDeviceDeviceData) HasInterfaceFree() bool`

HasInterfaceFree returns a boolean if a field has been set.

### GetInterfaceTotal

`func (o *DataInnerDeviceDeviceData) GetInterfaceTotal() string`

GetInterfaceTotal returns the InterfaceTotal field if non-nil, zero value otherwise.

### GetInterfaceTotalOk

`func (o *DataInnerDeviceDeviceData) GetInterfaceTotalOk() (*string, bool)`

GetInterfaceTotalOk returns a tuple with the InterfaceTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceTotal

`func (o *DataInnerDeviceDeviceData) SetInterfaceTotal(v string)`

SetInterfaceTotal sets InterfaceTotal field to given value.

### HasInterfaceTotal

`func (o *DataInnerDeviceDeviceData) HasInterfaceTotal() bool`

HasInterfaceTotal returns a boolean if a field has been set.

### GetInterfaceUsed

`func (o *DataInnerDeviceDeviceData) GetInterfaceUsed() string`

GetInterfaceUsed returns the InterfaceUsed field if non-nil, zero value otherwise.

### GetInterfaceUsedOk

`func (o *DataInnerDeviceDeviceData) GetInterfaceUsedOk() (*string, bool)`

GetInterfaceUsedOk returns a tuple with the InterfaceUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUsed

`func (o *DataInnerDeviceDeviceData) SetInterfaceUsed(v string)`

SetInterfaceUsed sets InterfaceUsed field to given value.

### HasInterfaceUsed

`func (o *DataInnerDeviceDeviceData) HasInterfaceUsed() bool`

HasInterfaceUsed returns a boolean if a field has been set.

### GetInterfaceUsedPercent

`func (o *DataInnerDeviceDeviceData) GetInterfaceUsedPercent() string`

GetInterfaceUsedPercent returns the InterfaceUsedPercent field if non-nil, zero value otherwise.

### GetInterfaceUsedPercentOk

`func (o *DataInnerDeviceDeviceData) GetInterfaceUsedPercentOk() (*string, bool)`

GetInterfaceUsedPercentOk returns a tuple with the InterfaceUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceUsedPercent

`func (o *DataInnerDeviceDeviceData) SetInterfaceUsedPercent(v string)`

SetInterfaceUsedPercent sets InterfaceUsedPercent field to given value.

### HasInterfaceUsedPercent

`func (o *DataInnerDeviceDeviceData) HasInterfaceUsedPercent() bool`

HasInterfaceUsedPercent returns a boolean if a field has been set.

### GetDevId

`func (o *DataInnerDeviceDeviceData) GetDevId() string`

GetDevId returns the DevId field if non-nil, zero value otherwise.

### GetDevIdOk

`func (o *DataInnerDeviceDeviceData) GetDevIdOk() (*string, bool)`

GetDevIdOk returns a tuple with the DevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevId

`func (o *DataInnerDeviceDeviceData) SetDevId(v string)`

SetDevId sets DevId field to given value.

### HasDevId

`func (o *DataInnerDeviceDeviceData) HasDevId() bool`

HasDevId returns a boolean if a field has been set.

### GetDevName

`func (o *DataInnerDeviceDeviceData) GetDevName() string`

GetDevName returns the DevName field if non-nil, zero value otherwise.

### GetDevNameOk

`func (o *DataInnerDeviceDeviceData) GetDevNameOk() (*string, bool)`

GetDevNameOk returns a tuple with the DevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevName

`func (o *DataInnerDeviceDeviceData) SetDevName(v string)`

SetDevName sets DevName field to given value.

### HasDevName

`func (o *DataInnerDeviceDeviceData) HasDevName() bool`

HasDevName returns a boolean if a field has been set.

### GetDevicePortFree

`func (o *DataInnerDeviceDeviceData) GetDevicePortFree() string`

GetDevicePortFree returns the DevicePortFree field if non-nil, zero value otherwise.

### GetDevicePortFreeOk

`func (o *DataInnerDeviceDeviceData) GetDevicePortFreeOk() (*string, bool)`

GetDevicePortFreeOk returns a tuple with the DevicePortFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortFree

`func (o *DataInnerDeviceDeviceData) SetDevicePortFree(v string)`

SetDevicePortFree sets DevicePortFree field to given value.

### HasDevicePortFree

`func (o *DataInnerDeviceDeviceData) HasDevicePortFree() bool`

HasDevicePortFree returns a boolean if a field has been set.

### GetDevicePortTotal

`func (o *DataInnerDeviceDeviceData) GetDevicePortTotal() string`

GetDevicePortTotal returns the DevicePortTotal field if non-nil, zero value otherwise.

### GetDevicePortTotalOk

`func (o *DataInnerDeviceDeviceData) GetDevicePortTotalOk() (*string, bool)`

GetDevicePortTotalOk returns a tuple with the DevicePortTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortTotal

`func (o *DataInnerDeviceDeviceData) SetDevicePortTotal(v string)`

SetDevicePortTotal sets DevicePortTotal field to given value.

### HasDevicePortTotal

`func (o *DataInnerDeviceDeviceData) HasDevicePortTotal() bool`

HasDevicePortTotal returns a boolean if a field has been set.

### GetDevicePortUsed

`func (o *DataInnerDeviceDeviceData) GetDevicePortUsed() string`

GetDevicePortUsed returns the DevicePortUsed field if non-nil, zero value otherwise.

### GetDevicePortUsedOk

`func (o *DataInnerDeviceDeviceData) GetDevicePortUsedOk() (*string, bool)`

GetDevicePortUsedOk returns a tuple with the DevicePortUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortUsed

`func (o *DataInnerDeviceDeviceData) SetDevicePortUsed(v string)`

SetDevicePortUsed sets DevicePortUsed field to given value.

### HasDevicePortUsed

`func (o *DataInnerDeviceDeviceData) HasDevicePortUsed() bool`

HasDevicePortUsed returns a boolean if a field has been set.

### GetDevicePortUsedPercent

`func (o *DataInnerDeviceDeviceData) GetDevicePortUsedPercent() string`

GetDevicePortUsedPercent returns the DevicePortUsedPercent field if non-nil, zero value otherwise.

### GetDevicePortUsedPercentOk

`func (o *DataInnerDeviceDeviceData) GetDevicePortUsedPercentOk() (*string, bool)`

GetDevicePortUsedPercentOk returns a tuple with the DevicePortUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevicePortUsedPercent

`func (o *DataInnerDeviceDeviceData) SetDevicePortUsedPercent(v string)`

SetDevicePortUsedPercent sets DevicePortUsedPercent field to given value.

### HasDevicePortUsedPercent

`func (o *DataInnerDeviceDeviceData) HasDevicePortUsedPercent() bool`

HasDevicePortUsedPercent returns a boolean if a field has been set.

### GetRowState

`func (o *DataInnerDeviceDeviceData) GetRowState() string`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *DataInnerDeviceDeviceData) GetRowStateOk() (*string, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *DataInnerDeviceDeviceData) SetRowState(v string)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *DataInnerDeviceDeviceData) HasRowState() bool`

HasRowState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


