# IpamAddress6AddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address6Hostaddr** | Pointer to **string** | The IP address. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**DeviceId** | Pointer to **int32** | The database identifier (ID) of the Device Manager device you want to associate with the IP address. | [optional] 
**InterfaceId** | Pointer to **int32** | The database identifier (ID) of the Device Manager interface you want to associate with the IP address. | [optional] 
**Address6MacAddr** | Pointer to **string** | The MAC address you want to associate with the IPv6 address. | [optional] 
**Address6Name** | Pointer to **string** | The name of the IPv6 address, each IPv6 address must have a unique name. | [optional] 
**PortId** | Pointer to **int32** | The database identifier (ID) of the NetChange port you want to associate with the IP address. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**Address6ClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**Address6ClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamAddress6AddInput

`func NewIpamAddress6AddInput() *IpamAddress6AddInput`

NewIpamAddress6AddInput instantiates a new IpamAddress6AddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAddress6AddInputWithDefaults

`func NewIpamAddress6AddInputWithDefaults() *IpamAddress6AddInput`

NewIpamAddress6AddInputWithDefaults instantiates a new IpamAddress6AddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress6Hostaddr

`func (o *IpamAddress6AddInput) GetAddress6Hostaddr() string`

GetAddress6Hostaddr returns the Address6Hostaddr field if non-nil, zero value otherwise.

### GetAddress6HostaddrOk

`func (o *IpamAddress6AddInput) GetAddress6HostaddrOk() (*string, bool)`

GetAddress6HostaddrOk returns a tuple with the Address6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Hostaddr

`func (o *IpamAddress6AddInput) SetAddress6Hostaddr(v string)`

SetAddress6Hostaddr sets Address6Hostaddr field to given value.

### HasAddress6Hostaddr

`func (o *IpamAddress6AddInput) HasAddress6Hostaddr() bool`

HasAddress6Hostaddr returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAddress6AddInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAddress6AddInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAddress6AddInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAddress6AddInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAddress6AddInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAddress6AddInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAddress6AddInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAddress6AddInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetDeviceId

`func (o *IpamAddress6AddInput) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IpamAddress6AddInput) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IpamAddress6AddInput) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IpamAddress6AddInput) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *IpamAddress6AddInput) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *IpamAddress6AddInput) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *IpamAddress6AddInput) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *IpamAddress6AddInput) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetAddress6MacAddr

`func (o *IpamAddress6AddInput) GetAddress6MacAddr() string`

GetAddress6MacAddr returns the Address6MacAddr field if non-nil, zero value otherwise.

### GetAddress6MacAddrOk

`func (o *IpamAddress6AddInput) GetAddress6MacAddrOk() (*string, bool)`

GetAddress6MacAddrOk returns a tuple with the Address6MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6MacAddr

`func (o *IpamAddress6AddInput) SetAddress6MacAddr(v string)`

SetAddress6MacAddr sets Address6MacAddr field to given value.

### HasAddress6MacAddr

`func (o *IpamAddress6AddInput) HasAddress6MacAddr() bool`

HasAddress6MacAddr returns a boolean if a field has been set.

### GetAddress6Name

`func (o *IpamAddress6AddInput) GetAddress6Name() string`

GetAddress6Name returns the Address6Name field if non-nil, zero value otherwise.

### GetAddress6NameOk

`func (o *IpamAddress6AddInput) GetAddress6NameOk() (*string, bool)`

GetAddress6NameOk returns a tuple with the Address6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Name

`func (o *IpamAddress6AddInput) SetAddress6Name(v string)`

SetAddress6Name sets Address6Name field to given value.

### HasAddress6Name

`func (o *IpamAddress6AddInput) HasAddress6Name() bool`

HasAddress6Name returns a boolean if a field has been set.

### GetPortId

`func (o *IpamAddress6AddInput) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *IpamAddress6AddInput) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *IpamAddress6AddInput) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *IpamAddress6AddInput) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *IpamAddress6AddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *IpamAddress6AddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *IpamAddress6AddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *IpamAddress6AddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetAddress6ClassName

`func (o *IpamAddress6AddInput) GetAddress6ClassName() string`

GetAddress6ClassName returns the Address6ClassName field if non-nil, zero value otherwise.

### GetAddress6ClassNameOk

`func (o *IpamAddress6AddInput) GetAddress6ClassNameOk() (*string, bool)`

GetAddress6ClassNameOk returns a tuple with the Address6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6ClassName

`func (o *IpamAddress6AddInput) SetAddress6ClassName(v string)`

SetAddress6ClassName sets Address6ClassName field to given value.

### HasAddress6ClassName

`func (o *IpamAddress6AddInput) HasAddress6ClassName() bool`

HasAddress6ClassName returns a boolean if a field has been set.

### GetAddress6ClassParameters

`func (o *IpamAddress6AddInput) GetAddress6ClassParameters() []ApiClassParameterInputEntry`

GetAddress6ClassParameters returns the Address6ClassParameters field if non-nil, zero value otherwise.

### GetAddress6ClassParametersOk

`func (o *IpamAddress6AddInput) GetAddress6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetAddress6ClassParametersOk returns a tuple with the Address6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6ClassParameters

`func (o *IpamAddress6AddInput) SetAddress6ClassParameters(v []ApiClassParameterInputEntry)`

SetAddress6ClassParameters sets Address6ClassParameters field to given value.

### HasAddress6ClassParameters

`func (o *IpamAddress6AddInput) HasAddress6ClassParameters() bool`

HasAddress6ClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamAddress6AddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamAddress6AddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamAddress6AddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamAddress6AddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


