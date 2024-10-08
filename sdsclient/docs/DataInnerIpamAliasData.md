# DataInnerIpamAliasData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasName** | Pointer to **string** | The name of the alias. | [optional] 
**StaticId** | Pointer to **string** | The database identifier (ID) of the DHCP static associated with the IP address. | [optional] 
**LeaseId** | Pointer to **string** | The database identifier (ID) of the DHCP lease associated with the IP address. | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device associated with the IP address. | [optional] 
**InterfaceId** | Pointer to **string** | The database identifier (ID) of the Device Manager interface associated with the IP address. | [optional] 
**AddressAddr** | Pointer to **string** | The IPv4 alias itself, in hexadecimal format. | [optional] 
**AddressClassName** | Pointer to **string** | The name of the class applied to the object, it can be preceded by the class directory. | [optional] 
**AddressClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv4 alias. | [optional] 
**AddressId** | Pointer to **string** | The database identifier (ID) of the IPv4 address associated with the alias. | [optional] 
**AliasId** | Pointer to **string** | The database identifier (ID) of the IPv4 alias. | [optional] 
**AliasType** | Pointer to **string** | The type of the alias, either &lt;b&gt;CNAME&lt;/b&gt; or &lt;b&gt;A&lt;/b&gt;. | [optional] 
**AddressType** | Pointer to **string** | A way to determine if you can assign the IP alias (&lt;b&gt;free&lt;/b&gt;) or if it is In use (&lt;b&gt;ip&lt;/b&gt;). | [optional] 
**PortId** | Pointer to **string** | The database identifier (ID) of the NetChange port associated with the IP address. | [optional] 
**AddressMacAddr** | Pointer to **string** | The MAC address associated with the IPv4 alias. | [optional] 
**AddressName** | Pointer to **string** | The name of the IPv4 alias. | [optional] 
**PoolId** | Pointer to **string** | The database identifier (ID) of the IPv4 pool the object belongs to. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**NetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 network the object belongs to. | [optional] 

## Methods

### NewDataInnerIpamAliasData

`func NewDataInnerIpamAliasData() *DataInnerIpamAliasData`

NewDataInnerIpamAliasData instantiates a new DataInnerIpamAliasData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerIpamAliasDataWithDefaults

`func NewDataInnerIpamAliasDataWithDefaults() *DataInnerIpamAliasData`

NewDataInnerIpamAliasDataWithDefaults instantiates a new DataInnerIpamAliasData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasName

`func (o *DataInnerIpamAliasData) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *DataInnerIpamAliasData) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *DataInnerIpamAliasData) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *DataInnerIpamAliasData) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetStaticId

`func (o *DataInnerIpamAliasData) GetStaticId() string`

GetStaticId returns the StaticId field if non-nil, zero value otherwise.

### GetStaticIdOk

`func (o *DataInnerIpamAliasData) GetStaticIdOk() (*string, bool)`

GetStaticIdOk returns a tuple with the StaticId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticId

`func (o *DataInnerIpamAliasData) SetStaticId(v string)`

SetStaticId sets StaticId field to given value.

### HasStaticId

`func (o *DataInnerIpamAliasData) HasStaticId() bool`

HasStaticId returns a boolean if a field has been set.

### GetLeaseId

`func (o *DataInnerIpamAliasData) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *DataInnerIpamAliasData) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *DataInnerIpamAliasData) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.

### HasLeaseId

`func (o *DataInnerIpamAliasData) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.

### GetDeviceId

`func (o *DataInnerIpamAliasData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DataInnerIpamAliasData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DataInnerIpamAliasData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DataInnerIpamAliasData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *DataInnerIpamAliasData) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *DataInnerIpamAliasData) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *DataInnerIpamAliasData) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *DataInnerIpamAliasData) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetAddressAddr

`func (o *DataInnerIpamAliasData) GetAddressAddr() string`

GetAddressAddr returns the AddressAddr field if non-nil, zero value otherwise.

### GetAddressAddrOk

`func (o *DataInnerIpamAliasData) GetAddressAddrOk() (*string, bool)`

GetAddressAddrOk returns a tuple with the AddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAddr

`func (o *DataInnerIpamAliasData) SetAddressAddr(v string)`

SetAddressAddr sets AddressAddr field to given value.

### HasAddressAddr

`func (o *DataInnerIpamAliasData) HasAddressAddr() bool`

HasAddressAddr returns a boolean if a field has been set.

### GetAddressClassName

`func (o *DataInnerIpamAliasData) GetAddressClassName() string`

GetAddressClassName returns the AddressClassName field if non-nil, zero value otherwise.

### GetAddressClassNameOk

`func (o *DataInnerIpamAliasData) GetAddressClassNameOk() (*string, bool)`

GetAddressClassNameOk returns a tuple with the AddressClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressClassName

`func (o *DataInnerIpamAliasData) SetAddressClassName(v string)`

SetAddressClassName sets AddressClassName field to given value.

### HasAddressClassName

`func (o *DataInnerIpamAliasData) HasAddressClassName() bool`

HasAddressClassName returns a boolean if a field has been set.

### GetAddressClassParameters

`func (o *DataInnerIpamAliasData) GetAddressClassParameters() []ApiClassParameterOutputEntry`

GetAddressClassParameters returns the AddressClassParameters field if non-nil, zero value otherwise.

### GetAddressClassParametersOk

`func (o *DataInnerIpamAliasData) GetAddressClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetAddressClassParametersOk returns a tuple with the AddressClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressClassParameters

`func (o *DataInnerIpamAliasData) SetAddressClassParameters(v []ApiClassParameterOutputEntry)`

SetAddressClassParameters sets AddressClassParameters field to given value.

### HasAddressClassParameters

`func (o *DataInnerIpamAliasData) HasAddressClassParameters() bool`

HasAddressClassParameters returns a boolean if a field has been set.

### GetAddressId

`func (o *DataInnerIpamAliasData) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *DataInnerIpamAliasData) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *DataInnerIpamAliasData) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *DataInnerIpamAliasData) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetAliasId

`func (o *DataInnerIpamAliasData) GetAliasId() string`

GetAliasId returns the AliasId field if non-nil, zero value otherwise.

### GetAliasIdOk

`func (o *DataInnerIpamAliasData) GetAliasIdOk() (*string, bool)`

GetAliasIdOk returns a tuple with the AliasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasId

`func (o *DataInnerIpamAliasData) SetAliasId(v string)`

SetAliasId sets AliasId field to given value.

### HasAliasId

`func (o *DataInnerIpamAliasData) HasAliasId() bool`

HasAliasId returns a boolean if a field has been set.

### GetAliasType

`func (o *DataInnerIpamAliasData) GetAliasType() string`

GetAliasType returns the AliasType field if non-nil, zero value otherwise.

### GetAliasTypeOk

`func (o *DataInnerIpamAliasData) GetAliasTypeOk() (*string, bool)`

GetAliasTypeOk returns a tuple with the AliasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasType

`func (o *DataInnerIpamAliasData) SetAliasType(v string)`

SetAliasType sets AliasType field to given value.

### HasAliasType

`func (o *DataInnerIpamAliasData) HasAliasType() bool`

HasAliasType returns a boolean if a field has been set.

### GetAddressType

`func (o *DataInnerIpamAliasData) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *DataInnerIpamAliasData) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *DataInnerIpamAliasData) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *DataInnerIpamAliasData) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.

### GetPortId

`func (o *DataInnerIpamAliasData) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *DataInnerIpamAliasData) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *DataInnerIpamAliasData) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *DataInnerIpamAliasData) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetAddressMacAddr

`func (o *DataInnerIpamAliasData) GetAddressMacAddr() string`

GetAddressMacAddr returns the AddressMacAddr field if non-nil, zero value otherwise.

### GetAddressMacAddrOk

`func (o *DataInnerIpamAliasData) GetAddressMacAddrOk() (*string, bool)`

GetAddressMacAddrOk returns a tuple with the AddressMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMacAddr

`func (o *DataInnerIpamAliasData) SetAddressMacAddr(v string)`

SetAddressMacAddr sets AddressMacAddr field to given value.

### HasAddressMacAddr

`func (o *DataInnerIpamAliasData) HasAddressMacAddr() bool`

HasAddressMacAddr returns a boolean if a field has been set.

### GetAddressName

`func (o *DataInnerIpamAliasData) GetAddressName() string`

GetAddressName returns the AddressName field if non-nil, zero value otherwise.

### GetAddressNameOk

`func (o *DataInnerIpamAliasData) GetAddressNameOk() (*string, bool)`

GetAddressNameOk returns a tuple with the AddressName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressName

`func (o *DataInnerIpamAliasData) SetAddressName(v string)`

SetAddressName sets AddressName field to given value.

### HasAddressName

`func (o *DataInnerIpamAliasData) HasAddressName() bool`

HasAddressName returns a boolean if a field has been set.

### GetPoolId

`func (o *DataInnerIpamAliasData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *DataInnerIpamAliasData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *DataInnerIpamAliasData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *DataInnerIpamAliasData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *DataInnerIpamAliasData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *DataInnerIpamAliasData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *DataInnerIpamAliasData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *DataInnerIpamAliasData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceId

`func (o *DataInnerIpamAliasData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DataInnerIpamAliasData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DataInnerIpamAliasData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DataInnerIpamAliasData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *DataInnerIpamAliasData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *DataInnerIpamAliasData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *DataInnerIpamAliasData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *DataInnerIpamAliasData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetNetworkId

`func (o *DataInnerIpamAliasData) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *DataInnerIpamAliasData) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *DataInnerIpamAliasData) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *DataInnerIpamAliasData) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


