# IpamAliasDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasName** | Pointer to **string** | The name of the IPv4 alias. | [optional] 
**StaticId** | Pointer to **string** | The database identifier (ID) of the DHCP static associated with the IP address. | [optional] 
**LeaseId** | Pointer to **string** | The database identifier (ID) of the DHCP lease associated the IP address. | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device associated with the IP address. | [optional] 
**InterfaceId** | Pointer to **string** | The database identifier (ID) of the Device Manager interface associated with the IP address. | [optional] 
**AddressAddr** | Pointer to **string** | The IPv4 alias itself, in hexadecimal format. | [optional] 
**AddressClassName** | Pointer to **string** | The name of the class applied to the object. | [optional] 
**AddressClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**AddressId** | Pointer to **string** | The database identifier (ID) of the IPv4 address associated with the alias. | [optional] 
**AliasId** | Pointer to **string** | The database identifier (ID) of the IPv4 alias. | [optional] 
**AliasType** | Pointer to **string** | The type of the alias, either &lt;b&gt;CNAME&lt;/b&gt; or &lt;b&gt;A&lt;/b&gt;. | [optional] 
**AddressType** | Pointer to **string** | A way to determine if you can assign the IP alias (&lt;b&gt;free&lt;/b&gt;) or if it is In use (&lt;b&gt;ip&lt;/b&gt;). | [optional] 
**PortId** | Pointer to **string** | The database identifier (ID) of the NetChange port associated with the IP address. | [optional] 
**AddressMacAddr** | Pointer to **string** | The MAC address associated with the IPv4 alias. | [optional] 
**AddressName** | Pointer to **string** | todo(here) :ipam.alias.list.output.address_name. :  | [optional] 
**PoolId** | Pointer to **string** | The database identifier (ID) of the IPv4 pool the object belongs to. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**NetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 network the object belongs to. | [optional] 

## Methods

### NewIpamAliasDataData

`func NewIpamAliasDataData() *IpamAliasDataData`

NewIpamAliasDataData instantiates a new IpamAliasDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAliasDataDataWithDefaults

`func NewIpamAliasDataDataWithDefaults() *IpamAliasDataData`

NewIpamAliasDataDataWithDefaults instantiates a new IpamAliasDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasName

`func (o *IpamAliasDataData) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *IpamAliasDataData) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *IpamAliasDataData) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *IpamAliasDataData) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetStaticId

`func (o *IpamAliasDataData) GetStaticId() string`

GetStaticId returns the StaticId field if non-nil, zero value otherwise.

### GetStaticIdOk

`func (o *IpamAliasDataData) GetStaticIdOk() (*string, bool)`

GetStaticIdOk returns a tuple with the StaticId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticId

`func (o *IpamAliasDataData) SetStaticId(v string)`

SetStaticId sets StaticId field to given value.

### HasStaticId

`func (o *IpamAliasDataData) HasStaticId() bool`

HasStaticId returns a boolean if a field has been set.

### GetLeaseId

`func (o *IpamAliasDataData) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *IpamAliasDataData) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *IpamAliasDataData) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.

### HasLeaseId

`func (o *IpamAliasDataData) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.

### GetDeviceId

`func (o *IpamAliasDataData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IpamAliasDataData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IpamAliasDataData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IpamAliasDataData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *IpamAliasDataData) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *IpamAliasDataData) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *IpamAliasDataData) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *IpamAliasDataData) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetAddressAddr

`func (o *IpamAliasDataData) GetAddressAddr() string`

GetAddressAddr returns the AddressAddr field if non-nil, zero value otherwise.

### GetAddressAddrOk

`func (o *IpamAliasDataData) GetAddressAddrOk() (*string, bool)`

GetAddressAddrOk returns a tuple with the AddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAddr

`func (o *IpamAliasDataData) SetAddressAddr(v string)`

SetAddressAddr sets AddressAddr field to given value.

### HasAddressAddr

`func (o *IpamAliasDataData) HasAddressAddr() bool`

HasAddressAddr returns a boolean if a field has been set.

### GetAddressClassName

`func (o *IpamAliasDataData) GetAddressClassName() string`

GetAddressClassName returns the AddressClassName field if non-nil, zero value otherwise.

### GetAddressClassNameOk

`func (o *IpamAliasDataData) GetAddressClassNameOk() (*string, bool)`

GetAddressClassNameOk returns a tuple with the AddressClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressClassName

`func (o *IpamAliasDataData) SetAddressClassName(v string)`

SetAddressClassName sets AddressClassName field to given value.

### HasAddressClassName

`func (o *IpamAliasDataData) HasAddressClassName() bool`

HasAddressClassName returns a boolean if a field has been set.

### GetAddressClassParameters

`func (o *IpamAliasDataData) GetAddressClassParameters() []ApiClassParameterOutputEntry`

GetAddressClassParameters returns the AddressClassParameters field if non-nil, zero value otherwise.

### GetAddressClassParametersOk

`func (o *IpamAliasDataData) GetAddressClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetAddressClassParametersOk returns a tuple with the AddressClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressClassParameters

`func (o *IpamAliasDataData) SetAddressClassParameters(v []ApiClassParameterOutputEntry)`

SetAddressClassParameters sets AddressClassParameters field to given value.

### HasAddressClassParameters

`func (o *IpamAliasDataData) HasAddressClassParameters() bool`

HasAddressClassParameters returns a boolean if a field has been set.

### GetAddressId

`func (o *IpamAliasDataData) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *IpamAliasDataData) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *IpamAliasDataData) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *IpamAliasDataData) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetAliasId

`func (o *IpamAliasDataData) GetAliasId() string`

GetAliasId returns the AliasId field if non-nil, zero value otherwise.

### GetAliasIdOk

`func (o *IpamAliasDataData) GetAliasIdOk() (*string, bool)`

GetAliasIdOk returns a tuple with the AliasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasId

`func (o *IpamAliasDataData) SetAliasId(v string)`

SetAliasId sets AliasId field to given value.

### HasAliasId

`func (o *IpamAliasDataData) HasAliasId() bool`

HasAliasId returns a boolean if a field has been set.

### GetAliasType

`func (o *IpamAliasDataData) GetAliasType() string`

GetAliasType returns the AliasType field if non-nil, zero value otherwise.

### GetAliasTypeOk

`func (o *IpamAliasDataData) GetAliasTypeOk() (*string, bool)`

GetAliasTypeOk returns a tuple with the AliasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasType

`func (o *IpamAliasDataData) SetAliasType(v string)`

SetAliasType sets AliasType field to given value.

### HasAliasType

`func (o *IpamAliasDataData) HasAliasType() bool`

HasAliasType returns a boolean if a field has been set.

### GetAddressType

`func (o *IpamAliasDataData) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *IpamAliasDataData) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *IpamAliasDataData) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *IpamAliasDataData) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetPortId

`func (o *IpamAliasDataData) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *IpamAliasDataData) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *IpamAliasDataData) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *IpamAliasDataData) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetAddressMacAddr

`func (o *IpamAliasDataData) GetAddressMacAddr() string`

GetAddressMacAddr returns the AddressMacAddr field if non-nil, zero value otherwise.

### GetAddressMacAddrOk

`func (o *IpamAliasDataData) GetAddressMacAddrOk() (*string, bool)`

GetAddressMacAddrOk returns a tuple with the AddressMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMacAddr

`func (o *IpamAliasDataData) SetAddressMacAddr(v string)`

SetAddressMacAddr sets AddressMacAddr field to given value.

### HasAddressMacAddr

`func (o *IpamAliasDataData) HasAddressMacAddr() bool`

HasAddressMacAddr returns a boolean if a field has been set.

### GetAddressName

`func (o *IpamAliasDataData) GetAddressName() string`

GetAddressName returns the AddressName field if non-nil, zero value otherwise.

### GetAddressNameOk

`func (o *IpamAliasDataData) GetAddressNameOk() (*string, bool)`

GetAddressNameOk returns a tuple with the AddressName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressName

`func (o *IpamAliasDataData) SetAddressName(v string)`

SetAddressName sets AddressName field to given value.

### HasAddressName

`func (o *IpamAliasDataData) HasAddressName() bool`

HasAddressName returns a boolean if a field has been set.

### GetPoolId

`func (o *IpamAliasDataData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *IpamAliasDataData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *IpamAliasDataData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *IpamAliasDataData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *IpamAliasDataData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *IpamAliasDataData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *IpamAliasDataData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *IpamAliasDataData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAliasDataData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAliasDataData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAliasDataData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAliasDataData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAliasDataData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAliasDataData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAliasDataData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAliasDataData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetNetworkId

`func (o *IpamAliasDataData) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *IpamAliasDataData) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *IpamAliasDataData) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *IpamAliasDataData) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


