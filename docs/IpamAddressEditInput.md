# IpamAddressEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressHostaddr** | Pointer to **string** | The IP address. | [optional] 
**AddressId** | Pointer to **int32** | The database identifier (ID) of the IPv4 address, a unique numeric key value automatically incremented when you add an IPv4 address. Use the ID to specify which IPv4 address to edit. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**CheckIsServerIp** | Pointer to **int32** | A way to force a validity check, if you configured the IPAM to DHCP replication. If the check is enabled (&lt;b&gt;1&lt;/b&gt;), the configuration of the IP address you are adding must be valid as well for the DHCP. | [optional] 
**StaticId** | Pointer to **int32** | The database identifier (ID) of the DHCP static you want to associate with the IP address. | [optional] 
**LeaseId** | Pointer to **int32** | The database identifier (ID) of the DHCP lease you want to associate with the IP address. | [optional] 
**DeviceId** | Pointer to **int32** | The database identifier (ID) of the Device Manager device you want to associate with the IP address. | [optional] 
**InterfaceId** | Pointer to **int32** | The database identifier (ID) of the Device Manager interface you want to associate with the IP address. | [optional] 
**PortId** | Pointer to **int32** | The database identifier (ID) of the NetChange port you want to associate with the IP address. | [optional] 
**AddressMacAddr** | Pointer to **string** | The MAC address you want to associate with the IPv4 address. | [optional] 
**AddressName** | Pointer to **string** | The name of the IPv4 address, each IPv4 address must have a unique name. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**AddressClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**AddressClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamAddressEditInput

`func NewIpamAddressEditInput() *IpamAddressEditInput`

NewIpamAddressEditInput instantiates a new IpamAddressEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAddressEditInputWithDefaults

`func NewIpamAddressEditInputWithDefaults() *IpamAddressEditInput`

NewIpamAddressEditInputWithDefaults instantiates a new IpamAddressEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressHostaddr

`func (o *IpamAddressEditInput) GetAddressHostaddr() string`

GetAddressHostaddr returns the AddressHostaddr field if non-nil, zero value otherwise.

### GetAddressHostaddrOk

`func (o *IpamAddressEditInput) GetAddressHostaddrOk() (*string, bool)`

GetAddressHostaddrOk returns a tuple with the AddressHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressHostaddr

`func (o *IpamAddressEditInput) SetAddressHostaddr(v string)`

SetAddressHostaddr sets AddressHostaddr field to given value.

### HasAddressHostaddr

`func (o *IpamAddressEditInput) HasAddressHostaddr() bool`

HasAddressHostaddr returns a boolean if a field has been set.

### GetAddressId

`func (o *IpamAddressEditInput) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *IpamAddressEditInput) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *IpamAddressEditInput) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *IpamAddressEditInput) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAddressEditInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAddressEditInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAddressEditInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAddressEditInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAddressEditInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAddressEditInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAddressEditInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAddressEditInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetCheckIsServerIp

`func (o *IpamAddressEditInput) GetCheckIsServerIp() int32`

GetCheckIsServerIp returns the CheckIsServerIp field if non-nil, zero value otherwise.

### GetCheckIsServerIpOk

`func (o *IpamAddressEditInput) GetCheckIsServerIpOk() (*int32, bool)`

GetCheckIsServerIpOk returns a tuple with the CheckIsServerIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckIsServerIp

`func (o *IpamAddressEditInput) SetCheckIsServerIp(v int32)`

SetCheckIsServerIp sets CheckIsServerIp field to given value.

### HasCheckIsServerIp

`func (o *IpamAddressEditInput) HasCheckIsServerIp() bool`

HasCheckIsServerIp returns a boolean if a field has been set.

### GetStaticId

`func (o *IpamAddressEditInput) GetStaticId() int32`

GetStaticId returns the StaticId field if non-nil, zero value otherwise.

### GetStaticIdOk

`func (o *IpamAddressEditInput) GetStaticIdOk() (*int32, bool)`

GetStaticIdOk returns a tuple with the StaticId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticId

`func (o *IpamAddressEditInput) SetStaticId(v int32)`

SetStaticId sets StaticId field to given value.

### HasStaticId

`func (o *IpamAddressEditInput) HasStaticId() bool`

HasStaticId returns a boolean if a field has been set.

### GetLeaseId

`func (o *IpamAddressEditInput) GetLeaseId() int32`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *IpamAddressEditInput) GetLeaseIdOk() (*int32, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *IpamAddressEditInput) SetLeaseId(v int32)`

SetLeaseId sets LeaseId field to given value.

### HasLeaseId

`func (o *IpamAddressEditInput) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.

### GetDeviceId

`func (o *IpamAddressEditInput) GetDeviceId() int32`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IpamAddressEditInput) GetDeviceIdOk() (*int32, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IpamAddressEditInput) SetDeviceId(v int32)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IpamAddressEditInput) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *IpamAddressEditInput) GetInterfaceId() int32`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *IpamAddressEditInput) GetInterfaceIdOk() (*int32, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *IpamAddressEditInput) SetInterfaceId(v int32)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *IpamAddressEditInput) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetPortId

`func (o *IpamAddressEditInput) GetPortId() int32`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *IpamAddressEditInput) GetPortIdOk() (*int32, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *IpamAddressEditInput) SetPortId(v int32)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *IpamAddressEditInput) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetAddressMacAddr

`func (o *IpamAddressEditInput) GetAddressMacAddr() string`

GetAddressMacAddr returns the AddressMacAddr field if non-nil, zero value otherwise.

### GetAddressMacAddrOk

`func (o *IpamAddressEditInput) GetAddressMacAddrOk() (*string, bool)`

GetAddressMacAddrOk returns a tuple with the AddressMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMacAddr

`func (o *IpamAddressEditInput) SetAddressMacAddr(v string)`

SetAddressMacAddr sets AddressMacAddr field to given value.

### HasAddressMacAddr

`func (o *IpamAddressEditInput) HasAddressMacAddr() bool`

HasAddressMacAddr returns a boolean if a field has been set.

### GetAddressName

`func (o *IpamAddressEditInput) GetAddressName() string`

GetAddressName returns the AddressName field if non-nil, zero value otherwise.

### GetAddressNameOk

`func (o *IpamAddressEditInput) GetAddressNameOk() (*string, bool)`

GetAddressNameOk returns a tuple with the AddressName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressName

`func (o *IpamAddressEditInput) SetAddressName(v string)`

SetAddressName sets AddressName field to given value.

### HasAddressName

`func (o *IpamAddressEditInput) HasAddressName() bool`

HasAddressName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *IpamAddressEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *IpamAddressEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *IpamAddressEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *IpamAddressEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetAddressClassName

`func (o *IpamAddressEditInput) GetAddressClassName() string`

GetAddressClassName returns the AddressClassName field if non-nil, zero value otherwise.

### GetAddressClassNameOk

`func (o *IpamAddressEditInput) GetAddressClassNameOk() (*string, bool)`

GetAddressClassNameOk returns a tuple with the AddressClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressClassName

`func (o *IpamAddressEditInput) SetAddressClassName(v string)`

SetAddressClassName sets AddressClassName field to given value.

### HasAddressClassName

`func (o *IpamAddressEditInput) HasAddressClassName() bool`

HasAddressClassName returns a boolean if a field has been set.

### GetAddressClassParameters

`func (o *IpamAddressEditInput) GetAddressClassParameters() []ApiClassParameterInputEntry`

GetAddressClassParameters returns the AddressClassParameters field if non-nil, zero value otherwise.

### GetAddressClassParametersOk

`func (o *IpamAddressEditInput) GetAddressClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetAddressClassParametersOk returns a tuple with the AddressClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressClassParameters

`func (o *IpamAddressEditInput) SetAddressClassParameters(v []ApiClassParameterInputEntry)`

SetAddressClassParameters sets AddressClassParameters field to given value.

### HasAddressClassParameters

`func (o *IpamAddressEditInput) HasAddressClassParameters() bool`

HasAddressClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamAddressEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamAddressEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamAddressEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamAddressEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


