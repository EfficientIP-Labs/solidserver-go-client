# IpamAlias6DataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AliasName** | Pointer to **string** | The name of the alias. | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device associated with the IP address. | [optional] 
**InterfaceId** | Pointer to **string** | The database identifier (ID) of the Device Manager interface associated with the IP address. | [optional] 
**Address6Addr** | Pointer to **string** | The IPv6 alias itself. | [optional] 
**Address6ClassName** | Pointer to **string** | The name of the class applied to the object. | [optional] 
**Address6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv6 alias and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Address6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 address associated with the alias. | [optional] 
**Address6MacAddr** | Pointer to **string** | The MAC address associated with the IPv6 alias. | [optional] 
**Address6Name** | Pointer to **string** | The name of the IPv6 address associated with the alias. | [optional] 
**Alias6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 alias. | [optional] 
**Alias6Type** | Pointer to **string** | The type of the alias, either &lt;b&gt;CNAME&lt;/b&gt; or &lt;b&gt;AAAA&lt;/b&gt;. | [optional] 
**PortId** | Pointer to **string** | The database identifier (ID) of the NetChange port associated with the IP address. | [optional] 
**Pool6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 pool the object belongs to. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**Network6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 network the object belongs to. | [optional] 

## Methods

### NewIpamAlias6DataData

`func NewIpamAlias6DataData() *IpamAlias6DataData`

NewIpamAlias6DataData instantiates a new IpamAlias6DataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAlias6DataDataWithDefaults

`func NewIpamAlias6DataDataWithDefaults() *IpamAlias6DataData`

NewIpamAlias6DataDataWithDefaults instantiates a new IpamAlias6DataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAliasName

`func (o *IpamAlias6DataData) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *IpamAlias6DataData) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *IpamAlias6DataData) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *IpamAlias6DataData) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetDeviceId

`func (o *IpamAlias6DataData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IpamAlias6DataData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IpamAlias6DataData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IpamAlias6DataData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetInterfaceId

`func (o *IpamAlias6DataData) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *IpamAlias6DataData) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *IpamAlias6DataData) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *IpamAlias6DataData) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetAddress6Addr

`func (o *IpamAlias6DataData) GetAddress6Addr() string`

GetAddress6Addr returns the Address6Addr field if non-nil, zero value otherwise.

### GetAddress6AddrOk

`func (o *IpamAlias6DataData) GetAddress6AddrOk() (*string, bool)`

GetAddress6AddrOk returns a tuple with the Address6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Addr

`func (o *IpamAlias6DataData) SetAddress6Addr(v string)`

SetAddress6Addr sets Address6Addr field to given value.

### HasAddress6Addr

`func (o *IpamAlias6DataData) HasAddress6Addr() bool`

HasAddress6Addr returns a boolean if a field has been set.

### GetAddress6ClassName

`func (o *IpamAlias6DataData) GetAddress6ClassName() string`

GetAddress6ClassName returns the Address6ClassName field if non-nil, zero value otherwise.

### GetAddress6ClassNameOk

`func (o *IpamAlias6DataData) GetAddress6ClassNameOk() (*string, bool)`

GetAddress6ClassNameOk returns a tuple with the Address6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6ClassName

`func (o *IpamAlias6DataData) SetAddress6ClassName(v string)`

SetAddress6ClassName sets Address6ClassName field to given value.

### HasAddress6ClassName

`func (o *IpamAlias6DataData) HasAddress6ClassName() bool`

HasAddress6ClassName returns a boolean if a field has been set.

### GetAddress6ClassParameters

`func (o *IpamAlias6DataData) GetAddress6ClassParameters() []ApiClassParameterOutputEntry`

GetAddress6ClassParameters returns the Address6ClassParameters field if non-nil, zero value otherwise.

### GetAddress6ClassParametersOk

`func (o *IpamAlias6DataData) GetAddress6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetAddress6ClassParametersOk returns a tuple with the Address6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6ClassParameters

`func (o *IpamAlias6DataData) SetAddress6ClassParameters(v []ApiClassParameterOutputEntry)`

SetAddress6ClassParameters sets Address6ClassParameters field to given value.

### HasAddress6ClassParameters

`func (o *IpamAlias6DataData) HasAddress6ClassParameters() bool`

HasAddress6ClassParameters returns a boolean if a field has been set.

### GetAddress6Id

`func (o *IpamAlias6DataData) GetAddress6Id() string`

GetAddress6Id returns the Address6Id field if non-nil, zero value otherwise.

### GetAddress6IdOk

`func (o *IpamAlias6DataData) GetAddress6IdOk() (*string, bool)`

GetAddress6IdOk returns a tuple with the Address6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Id

`func (o *IpamAlias6DataData) SetAddress6Id(v string)`

SetAddress6Id sets Address6Id field to given value.

### HasAddress6Id

`func (o *IpamAlias6DataData) HasAddress6Id() bool`

HasAddress6Id returns a boolean if a field has been set.

### GetAddress6MacAddr

`func (o *IpamAlias6DataData) GetAddress6MacAddr() string`

GetAddress6MacAddr returns the Address6MacAddr field if non-nil, zero value otherwise.

### GetAddress6MacAddrOk

`func (o *IpamAlias6DataData) GetAddress6MacAddrOk() (*string, bool)`

GetAddress6MacAddrOk returns a tuple with the Address6MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6MacAddr

`func (o *IpamAlias6DataData) SetAddress6MacAddr(v string)`

SetAddress6MacAddr sets Address6MacAddr field to given value.

### HasAddress6MacAddr

`func (o *IpamAlias6DataData) HasAddress6MacAddr() bool`

HasAddress6MacAddr returns a boolean if a field has been set.

### GetAddress6Name

`func (o *IpamAlias6DataData) GetAddress6Name() string`

GetAddress6Name returns the Address6Name field if non-nil, zero value otherwise.

### GetAddress6NameOk

`func (o *IpamAlias6DataData) GetAddress6NameOk() (*string, bool)`

GetAddress6NameOk returns a tuple with the Address6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Name

`func (o *IpamAlias6DataData) SetAddress6Name(v string)`

SetAddress6Name sets Address6Name field to given value.

### HasAddress6Name

`func (o *IpamAlias6DataData) HasAddress6Name() bool`

HasAddress6Name returns a boolean if a field has been set.

### GetAlias6Id

`func (o *IpamAlias6DataData) GetAlias6Id() string`

GetAlias6Id returns the Alias6Id field if non-nil, zero value otherwise.

### GetAlias6IdOk

`func (o *IpamAlias6DataData) GetAlias6IdOk() (*string, bool)`

GetAlias6IdOk returns a tuple with the Alias6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias6Id

`func (o *IpamAlias6DataData) SetAlias6Id(v string)`

SetAlias6Id sets Alias6Id field to given value.

### HasAlias6Id

`func (o *IpamAlias6DataData) HasAlias6Id() bool`

HasAlias6Id returns a boolean if a field has been set.

### GetAlias6Type

`func (o *IpamAlias6DataData) GetAlias6Type() string`

GetAlias6Type returns the Alias6Type field if non-nil, zero value otherwise.

### GetAlias6TypeOk

`func (o *IpamAlias6DataData) GetAlias6TypeOk() (*string, bool)`

GetAlias6TypeOk returns a tuple with the Alias6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias6Type

`func (o *IpamAlias6DataData) SetAlias6Type(v string)`

SetAlias6Type sets Alias6Type field to given value.

### HasAlias6Type

`func (o *IpamAlias6DataData) HasAlias6Type() bool`

HasAlias6Type returns a boolean if a field has been set.

### GetPortId

`func (o *IpamAlias6DataData) GetPortId() string`

GetPortId returns the PortId field if non-nil, zero value otherwise.

### GetPortIdOk

`func (o *IpamAlias6DataData) GetPortIdOk() (*string, bool)`

GetPortIdOk returns a tuple with the PortId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortId

`func (o *IpamAlias6DataData) SetPortId(v string)`

SetPortId sets PortId field to given value.

### HasPortId

`func (o *IpamAlias6DataData) HasPortId() bool`

HasPortId returns a boolean if a field has been set.

### GetPool6Id

`func (o *IpamAlias6DataData) GetPool6Id() string`

GetPool6Id returns the Pool6Id field if non-nil, zero value otherwise.

### GetPool6IdOk

`func (o *IpamAlias6DataData) GetPool6IdOk() (*string, bool)`

GetPool6IdOk returns a tuple with the Pool6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Id

`func (o *IpamAlias6DataData) SetPool6Id(v string)`

SetPool6Id sets Pool6Id field to given value.

### HasPool6Id

`func (o *IpamAlias6DataData) HasPool6Id() bool`

HasPool6Id returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *IpamAlias6DataData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *IpamAlias6DataData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *IpamAlias6DataData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *IpamAlias6DataData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAlias6DataData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAlias6DataData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAlias6DataData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAlias6DataData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAlias6DataData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAlias6DataData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAlias6DataData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAlias6DataData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetNetwork6Id

`func (o *IpamAlias6DataData) GetNetwork6Id() string`

GetNetwork6Id returns the Network6Id field if non-nil, zero value otherwise.

### GetNetwork6IdOk

`func (o *IpamAlias6DataData) GetNetwork6IdOk() (*string, bool)`

GetNetwork6IdOk returns a tuple with the Network6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Id

`func (o *IpamAlias6DataData) SetNetwork6Id(v string)`

SetNetwork6Id sets Network6Id field to given value.

### HasNetwork6Id

`func (o *IpamAlias6DataData) HasNetwork6Id() bool`

HasNetwork6Id returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


