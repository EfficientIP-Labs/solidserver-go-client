# IpamAddressDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticId** | Pointer to **string** | The database identifier (ID) of the DHCP static associated with the IP address. | [optional] 
**LeaseEndTime** | Pointer to **string** | The expiration time of the lease, if the IP address was imported from the DHCP, in decimal UNIX date format. | [optional] 
**LeaseId** | Pointer to **string** | The database identifier (ID) of the DHCP lease associated the IP address. | [optional] 
**FreeEndAddressAddr** | Pointer to **string** | An IP address in hexadecimal format. For addresses &lt;b&gt;In use&lt;/b&gt; (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;ip&lt;/b&gt;), it returns the IP address itself.For free addresses (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;), it returns the last IP address of a range of IPv4 addresses that are not assigned yet. The first address in that range is returned in &lt;b&gt;free_start_ip_addr&lt;/b&gt;. | [optional] 
**FreeScopeSize** | Pointer to **string** | The number of IP addresses that are not assigned yet (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;) between &lt;b&gt;free_start_ip_addr&lt;/b&gt; and &lt;b&gt;free_end_ip_addr&lt;/b&gt;. | [optional] 
**FreeStartAddressAddr** | Pointer to **string** | An IP address in hexadecimal format. For addresses &lt;b&gt;In use&lt;/b&gt; (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;ip&lt;/b&gt;), it returns the IP address itself.For free addresses (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;), it returns the first IP address of a range of IPv4 addresses that are not assigned yet. The last address in that range is returned in &lt;b&gt;free_end_ip_addr&lt;/b&gt;. | [optional] 
**AddressHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;ip_addr&lt;/b&gt;. | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device associated with the IP address. | [optional] 
**DeviceName** | Pointer to **string** | The name of the Device Manager device associated with the IP address. | [optional] 
**InterfaceId** | Pointer to **string** | The database identifier (ID) of the Device Manager interface associated with the IP address. | [optional] 
**InterfaceName** | Pointer to **string** | The name of the Device Manager interface associated with the IP address. | [optional] 
**AddressAddr** | Pointer to **string** | The IPv4 address itself, in hexadecimal format. | [optional] 
**AddressAlias** | Pointer to **string** | The name of the IPv4 alias(es) associated with the IPv4 address. | [optional] 
**AddressClassName** | Pointer to **string** | The name of the class applied to the IPv4 address, it can be preceded by the class directory. | [optional] 
**AddressClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**AddressId** | Pointer to **string** | The database identifier (ID) of the IPv4 address. | [optional] 
**DevId** | Pointer to **string** | The database identifier (ID) of the NetChange network device associated with the IP address. | [optional] 
**DevName** | Pointer to **string** | The name of the NetChange network device associated with the IP address. | [optional] 
**PortIfvlan** | Pointer to **string** | The VLAN identifier (ID) of the NetChange port, for IP addresses which MAC addresses is imported from NetChange. | [optional] 
**PortName** | Pointer to **string** | The name of the NetChange port associated with the IP address. | [optional] 
**PortPortnumber** | Pointer to **string** | The number of the port, for IP addresses which MAC addresses is imported from NetChange. | [optional] 
**PortSlotnumber** | Pointer to **string** | The slot number of the port, for IP addresses which MAC addresses is imported from NetChange. | [optional] 
**AddressMacLastSeen** | Pointer to **string** | The last time the MAC address associated with the IP address was seen on the network, in decimal UNIX date format. | [optional] 
**NetworkLockNetworkBroadcast** | Pointer to **string** | A way to prevent &lt;b&gt;(1)&lt;/b&gt; users from assigning the broadcast IP address and network IP address of the network the IP address belongs to. | [optional] 
**AddressMacAddr** | Pointer to **string** | The MAC address associated with the IPv4 address. | [optional] 
**AddressMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**AddressName** | Pointer to **string** | The name of the IPv4 address. | [optional] 
**ParentNetworkClassName** | Pointer to **string** | The name of the class applied to the parent of the IPv4 network the object belongs to, it can be preceded by the class directory. | [optional] 
**ParentNetworkEndHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;parent_subnet_end_ip_addr&lt;/b&gt;. | [optional] 
**ParentNetworkEndAddressAddr** | Pointer to **string** | The last IP address of the parent of the IPv4 network the IP address belongs to. | [optional] 
**ParentNetworkId** | Pointer to **string** | The database identifier (ID) of the parent IPv4 network. It identifies the parent of the IPv4 network the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the network the object belongs to has no parent network. | [optional] 
**ParentNetworkName** | Pointer to **string** | The name of the parent IPv4 network:&lt;ul class&#x3D;dashed &gt;&lt;li&gt;&lt;b&gt;#&lt;/b&gt; indicates that the network the object belongs to has no parent network.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;&lt;b&gt;Default&lt;/b&gt; indicates that the network the object belongs to is in an orphan network.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt; | [optional] 
**ParentNetworkSize** | Pointer to **string** | The number of IP addresses of the parent of the network the object belongs to. | [optional] 
**ParentNetworkStartHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;parent_subnet_start_ip_addr&lt;/b&gt;. | [optional] 
**ParentNetworkStartAddressAddr** | Pointer to **string** | The first IP address of the parent of the IPv4 network the IP address belongs to. | [optional] 
**ParentVlsmNetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 subnet-type network, located in the VLSM parent space, from which the parent of the network the IP address belongs to was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the parent of the network the IP address belongs to is not a VLSM block-type network duplicated from a parent space. | [optional] 
**PoolClassName** | Pointer to **string** | The name of the class applied to the IPv4 pool the object belongs to, it can be preceded by the class directory. | [optional] 
**PoolClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**PoolEndAddressAddr** | Pointer to **string** | The last IP address of the IPv4 pool the IP address belongs to. | [optional] 
**PoolId** | Pointer to **string** | The database identifier (ID) of the IPv4 pool the object belongs to. | [optional] 
**PoolName** | Pointer to **string** | The name of the IPv4 pool the object belongs to. | [optional] 
**PoolReadOnly** | Pointer to **string** | The reservation status of the pool the IPv4 address belongs to. If set &lt;b&gt;1&lt;/b&gt;, the pool is reserved and you cannot assign the IP address. | [optional] 
**PoolSize** | Pointer to **string** | The number of IP addresses that contains the pool the IPv4 address belongs to. | [optional] 
**PoolStartAddressAddr** | Pointer to **string** | The first IP address of the IPv4 pool the IP address belongs to. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space the object belongs to. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceIsTemplate** | Pointer to **string** | The template status of the space the object belongs to. If the space is used as template (1), all the IPv4 networks, pools and IP addresses it contains are also used as template. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**NetworkClassName** | Pointer to **string** | The name of the class applied to the IPv4 network the object belongs to, it can be preceded by the class directory. | [optional] 
**NetworkClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**NetworkEndAddressAddr** | Pointer to **string** | The last IP address of the IPv4 network the object belongs to. | [optional] 
**NetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 network the object belongs to. | [optional] 
**NetworkIsTerminal** | Pointer to **string** | A way to determine if the network the IP address belongs to is terminal (&lt;b&gt;1&lt;/b&gt;) or non-terminal (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**NetworkName** | Pointer to **string** | The name of the IPv4 network the object belongs to. &lt;b&gt;Default&lt;/b&gt; indicates that the network the object belongs to is an orphan network. | [optional] 
**NetworkSize** | Pointer to **string** | The number of IP addresses the network the object belongs to contains. | [optional] 
**NetworkStartAddressAddr** | Pointer to **string** | The first IP address of the IPv4 network the object belongs to. | [optional] 
**TagContainerDhcpstatic** | Pointer to **string** | A way to determine if the terminal network or pool the IP address belongs to is configured to Create DHCP static (&lt;b&gt;1&lt;/b&gt;) or not (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**TagPoolDhcprange** | Pointer to **string** | A way to determine if the pool the IP address belongs to is configured to Create DHCP range (&lt;b&gt;1&lt;/b&gt;) or not (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**AddressType** | Pointer to **string** | A way to determine if you can assign the IP address (&lt;b&gt;free&lt;/b&gt;) or if it is In use (&lt;b&gt;ip&lt;/b&gt;). | [optional] 

## Methods

### NewIpamAddressDataData

`func NewIpamAddressDataData() *IpamAddressDataData`

NewIpamAddressDataData instantiates a new IpamAddressDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAddressDataDataWithDefaults

`func NewIpamAddressDataDataWithDefaults() *IpamAddressDataData`

NewIpamAddressDataDataWithDefaults instantiates a new IpamAddressDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticId

`func (o *IpamAddressDataData) GetStaticId() string`

GetStaticId returns the StaticId field if non-nil, zero value otherwise.

### GetStaticIdOk

`func (o *IpamAddressDataData) GetStaticIdOk() (*string, bool)`

GetStaticIdOk returns a tuple with the StaticId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticId

`func (o *IpamAddressDataData) SetStaticId(v string)`

SetStaticId sets StaticId field to given value.

### HasStaticId

`func (o *IpamAddressDataData) HasStaticId() bool`

HasStaticId returns a boolean if a field has been set.

### GetLeaseEndTime

`func (o *IpamAddressDataData) GetLeaseEndTime() string`

GetLeaseEndTime returns the LeaseEndTime field if non-nil, zero value otherwise.

### GetLeaseEndTimeOk

`func (o *IpamAddressDataData) GetLeaseEndTimeOk() (*string, bool)`

GetLeaseEndTimeOk returns a tuple with the LeaseEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseEndTime

`func (o *IpamAddressDataData) SetLeaseEndTime(v string)`

SetLeaseEndTime sets LeaseEndTime field to given value.

### HasLeaseEndTime

`func (o *IpamAddressDataData) HasLeaseEndTime() bool`

HasLeaseEndTime returns a boolean if a field has been set.

### GetLeaseId

`func (o *IpamAddressDataData) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *IpamAddressDataData) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *IpamAddressDataData) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.

### HasLeaseId

`func (o *IpamAddressDataData) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.

### GetFreeEndAddressAddr

`func (o *IpamAddressDataData) GetFreeEndAddressAddr() string`

GetFreeEndAddressAddr returns the FreeEndAddressAddr field if non-nil, zero value otherwise.

### GetFreeEndAddressAddrOk

`func (o *IpamAddressDataData) GetFreeEndAddressAddrOk() (*string, bool)`

GetFreeEndAddressAddrOk returns a tuple with the FreeEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeEndAddressAddr

`func (o *IpamAddressDataData) SetFreeEndAddressAddr(v string)`

SetFreeEndAddressAddr sets FreeEndAddressAddr field to given value.

### HasFreeEndAddressAddr

`func (o *IpamAddressDataData) HasFreeEndAddressAddr() bool`

HasFreeEndAddressAddr returns a boolean if a field has been set.

### GetFreeScopeSize

`func (o *IpamAddressDataData) GetFreeScopeSize() string`

GetFreeScopeSize returns the FreeScopeSize field if non-nil, zero value otherwise.

### GetFreeScopeSizeOk

`func (o *IpamAddressDataData) GetFreeScopeSizeOk() (*string, bool)`

GetFreeScopeSizeOk returns a tuple with the FreeScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeScopeSize

`func (o *IpamAddressDataData) SetFreeScopeSize(v string)`

SetFreeScopeSize sets FreeScopeSize field to given value.

### HasFreeScopeSize

`func (o *IpamAddressDataData) HasFreeScopeSize() bool`

HasFreeScopeSize returns a boolean if a field has been set.

### GetFreeStartAddressAddr

`func (o *IpamAddressDataData) GetFreeStartAddressAddr() string`

GetFreeStartAddressAddr returns the FreeStartAddressAddr field if non-nil, zero value otherwise.

### GetFreeStartAddressAddrOk

`func (o *IpamAddressDataData) GetFreeStartAddressAddrOk() (*string, bool)`

GetFreeStartAddressAddrOk returns a tuple with the FreeStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStartAddressAddr

`func (o *IpamAddressDataData) SetFreeStartAddressAddr(v string)`

SetFreeStartAddressAddr sets FreeStartAddressAddr field to given value.

### HasFreeStartAddressAddr

`func (o *IpamAddressDataData) HasFreeStartAddressAddr() bool`

HasFreeStartAddressAddr returns a boolean if a field has been set.

### GetAddressHostaddr

`func (o *IpamAddressDataData) GetAddressHostaddr() string`

GetAddressHostaddr returns the AddressHostaddr field if non-nil, zero value otherwise.

### GetAddressHostaddrOk

`func (o *IpamAddressDataData) GetAddressHostaddrOk() (*string, bool)`

GetAddressHostaddrOk returns a tuple with the AddressHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressHostaddr

`func (o *IpamAddressDataData) SetAddressHostaddr(v string)`

SetAddressHostaddr sets AddressHostaddr field to given value.

### HasAddressHostaddr

`func (o *IpamAddressDataData) HasAddressHostaddr() bool`

HasAddressHostaddr returns a boolean if a field has been set.

### GetDeviceId

`func (o *IpamAddressDataData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IpamAddressDataData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IpamAddressDataData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IpamAddressDataData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *IpamAddressDataData) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *IpamAddressDataData) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *IpamAddressDataData) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *IpamAddressDataData) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetInterfaceId

`func (o *IpamAddressDataData) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *IpamAddressDataData) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *IpamAddressDataData) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *IpamAddressDataData) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *IpamAddressDataData) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *IpamAddressDataData) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *IpamAddressDataData) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *IpamAddressDataData) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetAddressAddr

`func (o *IpamAddressDataData) GetAddressAddr() string`

GetAddressAddr returns the AddressAddr field if non-nil, zero value otherwise.

### GetAddressAddrOk

`func (o *IpamAddressDataData) GetAddressAddrOk() (*string, bool)`

GetAddressAddrOk returns a tuple with the AddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAddr

`func (o *IpamAddressDataData) SetAddressAddr(v string)`

SetAddressAddr sets AddressAddr field to given value.

### HasAddressAddr

`func (o *IpamAddressDataData) HasAddressAddr() bool`

HasAddressAddr returns a boolean if a field has been set.

### GetAddressAlias

`func (o *IpamAddressDataData) GetAddressAlias() string`

GetAddressAlias returns the AddressAlias field if non-nil, zero value otherwise.

### GetAddressAliasOk

`func (o *IpamAddressDataData) GetAddressAliasOk() (*string, bool)`

GetAddressAliasOk returns a tuple with the AddressAlias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressAlias

`func (o *IpamAddressDataData) SetAddressAlias(v string)`

SetAddressAlias sets AddressAlias field to given value.

### HasAddressAlias

`func (o *IpamAddressDataData) HasAddressAlias() bool`

HasAddressAlias returns a boolean if a field has been set.

### GetAddressClassName

`func (o *IpamAddressDataData) GetAddressClassName() string`

GetAddressClassName returns the AddressClassName field if non-nil, zero value otherwise.

### GetAddressClassNameOk

`func (o *IpamAddressDataData) GetAddressClassNameOk() (*string, bool)`

GetAddressClassNameOk returns a tuple with the AddressClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressClassName

`func (o *IpamAddressDataData) SetAddressClassName(v string)`

SetAddressClassName sets AddressClassName field to given value.

### HasAddressClassName

`func (o *IpamAddressDataData) HasAddressClassName() bool`

HasAddressClassName returns a boolean if a field has been set.

### GetAddressClassParameters

`func (o *IpamAddressDataData) GetAddressClassParameters() []ApiClassParameterOutputEntry`

GetAddressClassParameters returns the AddressClassParameters field if non-nil, zero value otherwise.

### GetAddressClassParametersOk

`func (o *IpamAddressDataData) GetAddressClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetAddressClassParametersOk returns a tuple with the AddressClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressClassParameters

`func (o *IpamAddressDataData) SetAddressClassParameters(v []ApiClassParameterOutputEntry)`

SetAddressClassParameters sets AddressClassParameters field to given value.

### HasAddressClassParameters

`func (o *IpamAddressDataData) HasAddressClassParameters() bool`

HasAddressClassParameters returns a boolean if a field has been set.

### GetAddressId

`func (o *IpamAddressDataData) GetAddressId() string`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *IpamAddressDataData) GetAddressIdOk() (*string, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *IpamAddressDataData) SetAddressId(v string)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *IpamAddressDataData) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetDevId

`func (o *IpamAddressDataData) GetDevId() string`

GetDevId returns the DevId field if non-nil, zero value otherwise.

### GetDevIdOk

`func (o *IpamAddressDataData) GetDevIdOk() (*string, bool)`

GetDevIdOk returns a tuple with the DevId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevId

`func (o *IpamAddressDataData) SetDevId(v string)`

SetDevId sets DevId field to given value.

### HasDevId

`func (o *IpamAddressDataData) HasDevId() bool`

HasDevId returns a boolean if a field has been set.

### GetDevName

`func (o *IpamAddressDataData) GetDevName() string`

GetDevName returns the DevName field if non-nil, zero value otherwise.

### GetDevNameOk

`func (o *IpamAddressDataData) GetDevNameOk() (*string, bool)`

GetDevNameOk returns a tuple with the DevName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevName

`func (o *IpamAddressDataData) SetDevName(v string)`

SetDevName sets DevName field to given value.

### HasDevName

`func (o *IpamAddressDataData) HasDevName() bool`

HasDevName returns a boolean if a field has been set.

### GetPortIfvlan

`func (o *IpamAddressDataData) GetPortIfvlan() string`

GetPortIfvlan returns the PortIfvlan field if non-nil, zero value otherwise.

### GetPortIfvlanOk

`func (o *IpamAddressDataData) GetPortIfvlanOk() (*string, bool)`

GetPortIfvlanOk returns a tuple with the PortIfvlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortIfvlan

`func (o *IpamAddressDataData) SetPortIfvlan(v string)`

SetPortIfvlan sets PortIfvlan field to given value.

### HasPortIfvlan

`func (o *IpamAddressDataData) HasPortIfvlan() bool`

HasPortIfvlan returns a boolean if a field has been set.

### GetPortName

`func (o *IpamAddressDataData) GetPortName() string`

GetPortName returns the PortName field if non-nil, zero value otherwise.

### GetPortNameOk

`func (o *IpamAddressDataData) GetPortNameOk() (*string, bool)`

GetPortNameOk returns a tuple with the PortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortName

`func (o *IpamAddressDataData) SetPortName(v string)`

SetPortName sets PortName field to given value.

### HasPortName

`func (o *IpamAddressDataData) HasPortName() bool`

HasPortName returns a boolean if a field has been set.

### GetPortPortnumber

`func (o *IpamAddressDataData) GetPortPortnumber() string`

GetPortPortnumber returns the PortPortnumber field if non-nil, zero value otherwise.

### GetPortPortnumberOk

`func (o *IpamAddressDataData) GetPortPortnumberOk() (*string, bool)`

GetPortPortnumberOk returns a tuple with the PortPortnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortPortnumber

`func (o *IpamAddressDataData) SetPortPortnumber(v string)`

SetPortPortnumber sets PortPortnumber field to given value.

### HasPortPortnumber

`func (o *IpamAddressDataData) HasPortPortnumber() bool`

HasPortPortnumber returns a boolean if a field has been set.

### GetPortSlotnumber

`func (o *IpamAddressDataData) GetPortSlotnumber() string`

GetPortSlotnumber returns the PortSlotnumber field if non-nil, zero value otherwise.

### GetPortSlotnumberOk

`func (o *IpamAddressDataData) GetPortSlotnumberOk() (*string, bool)`

GetPortSlotnumberOk returns a tuple with the PortSlotnumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortSlotnumber

`func (o *IpamAddressDataData) SetPortSlotnumber(v string)`

SetPortSlotnumber sets PortSlotnumber field to given value.

### HasPortSlotnumber

`func (o *IpamAddressDataData) HasPortSlotnumber() bool`

HasPortSlotnumber returns a boolean if a field has been set.

### GetAddressMacLastSeen

`func (o *IpamAddressDataData) GetAddressMacLastSeen() string`

GetAddressMacLastSeen returns the AddressMacLastSeen field if non-nil, zero value otherwise.

### GetAddressMacLastSeenOk

`func (o *IpamAddressDataData) GetAddressMacLastSeenOk() (*string, bool)`

GetAddressMacLastSeenOk returns a tuple with the AddressMacLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMacLastSeen

`func (o *IpamAddressDataData) SetAddressMacLastSeen(v string)`

SetAddressMacLastSeen sets AddressMacLastSeen field to given value.

### HasAddressMacLastSeen

`func (o *IpamAddressDataData) HasAddressMacLastSeen() bool`

HasAddressMacLastSeen returns a boolean if a field has been set.

### GetNetworkLockNetworkBroadcast

`func (o *IpamAddressDataData) GetNetworkLockNetworkBroadcast() string`

GetNetworkLockNetworkBroadcast returns the NetworkLockNetworkBroadcast field if non-nil, zero value otherwise.

### GetNetworkLockNetworkBroadcastOk

`func (o *IpamAddressDataData) GetNetworkLockNetworkBroadcastOk() (*string, bool)`

GetNetworkLockNetworkBroadcastOk returns a tuple with the NetworkLockNetworkBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLockNetworkBroadcast

`func (o *IpamAddressDataData) SetNetworkLockNetworkBroadcast(v string)`

SetNetworkLockNetworkBroadcast sets NetworkLockNetworkBroadcast field to given value.

### HasNetworkLockNetworkBroadcast

`func (o *IpamAddressDataData) HasNetworkLockNetworkBroadcast() bool`

HasNetworkLockNetworkBroadcast returns a boolean if a field has been set.

### GetAddressMacAddr

`func (o *IpamAddressDataData) GetAddressMacAddr() string`

GetAddressMacAddr returns the AddressMacAddr field if non-nil, zero value otherwise.

### GetAddressMacAddrOk

`func (o *IpamAddressDataData) GetAddressMacAddrOk() (*string, bool)`

GetAddressMacAddrOk returns a tuple with the AddressMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMacAddr

`func (o *IpamAddressDataData) SetAddressMacAddr(v string)`

SetAddressMacAddr sets AddressMacAddr field to given value.

### HasAddressMacAddr

`func (o *IpamAddressDataData) HasAddressMacAddr() bool`

HasAddressMacAddr returns a boolean if a field has been set.

### GetAddressMultistatus

`func (o *IpamAddressDataData) GetAddressMultistatus() string`

GetAddressMultistatus returns the AddressMultistatus field if non-nil, zero value otherwise.

### GetAddressMultistatusOk

`func (o *IpamAddressDataData) GetAddressMultistatusOk() (*string, bool)`

GetAddressMultistatusOk returns a tuple with the AddressMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressMultistatus

`func (o *IpamAddressDataData) SetAddressMultistatus(v string)`

SetAddressMultistatus sets AddressMultistatus field to given value.

### HasAddressMultistatus

`func (o *IpamAddressDataData) HasAddressMultistatus() bool`

HasAddressMultistatus returns a boolean if a field has been set.

### GetAddressName

`func (o *IpamAddressDataData) GetAddressName() string`

GetAddressName returns the AddressName field if non-nil, zero value otherwise.

### GetAddressNameOk

`func (o *IpamAddressDataData) GetAddressNameOk() (*string, bool)`

GetAddressNameOk returns a tuple with the AddressName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressName

`func (o *IpamAddressDataData) SetAddressName(v string)`

SetAddressName sets AddressName field to given value.

### HasAddressName

`func (o *IpamAddressDataData) HasAddressName() bool`

HasAddressName returns a boolean if a field has been set.

### GetParentNetworkClassName

`func (o *IpamAddressDataData) GetParentNetworkClassName() string`

GetParentNetworkClassName returns the ParentNetworkClassName field if non-nil, zero value otherwise.

### GetParentNetworkClassNameOk

`func (o *IpamAddressDataData) GetParentNetworkClassNameOk() (*string, bool)`

GetParentNetworkClassNameOk returns a tuple with the ParentNetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkClassName

`func (o *IpamAddressDataData) SetParentNetworkClassName(v string)`

SetParentNetworkClassName sets ParentNetworkClassName field to given value.

### HasParentNetworkClassName

`func (o *IpamAddressDataData) HasParentNetworkClassName() bool`

HasParentNetworkClassName returns a boolean if a field has been set.

### GetParentNetworkEndHostaddr

`func (o *IpamAddressDataData) GetParentNetworkEndHostaddr() string`

GetParentNetworkEndHostaddr returns the ParentNetworkEndHostaddr field if non-nil, zero value otherwise.

### GetParentNetworkEndHostaddrOk

`func (o *IpamAddressDataData) GetParentNetworkEndHostaddrOk() (*string, bool)`

GetParentNetworkEndHostaddrOk returns a tuple with the ParentNetworkEndHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkEndHostaddr

`func (o *IpamAddressDataData) SetParentNetworkEndHostaddr(v string)`

SetParentNetworkEndHostaddr sets ParentNetworkEndHostaddr field to given value.

### HasParentNetworkEndHostaddr

`func (o *IpamAddressDataData) HasParentNetworkEndHostaddr() bool`

HasParentNetworkEndHostaddr returns a boolean if a field has been set.

### GetParentNetworkEndAddressAddr

`func (o *IpamAddressDataData) GetParentNetworkEndAddressAddr() string`

GetParentNetworkEndAddressAddr returns the ParentNetworkEndAddressAddr field if non-nil, zero value otherwise.

### GetParentNetworkEndAddressAddrOk

`func (o *IpamAddressDataData) GetParentNetworkEndAddressAddrOk() (*string, bool)`

GetParentNetworkEndAddressAddrOk returns a tuple with the ParentNetworkEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkEndAddressAddr

`func (o *IpamAddressDataData) SetParentNetworkEndAddressAddr(v string)`

SetParentNetworkEndAddressAddr sets ParentNetworkEndAddressAddr field to given value.

### HasParentNetworkEndAddressAddr

`func (o *IpamAddressDataData) HasParentNetworkEndAddressAddr() bool`

HasParentNetworkEndAddressAddr returns a boolean if a field has been set.

### GetParentNetworkId

`func (o *IpamAddressDataData) GetParentNetworkId() string`

GetParentNetworkId returns the ParentNetworkId field if non-nil, zero value otherwise.

### GetParentNetworkIdOk

`func (o *IpamAddressDataData) GetParentNetworkIdOk() (*string, bool)`

GetParentNetworkIdOk returns a tuple with the ParentNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkId

`func (o *IpamAddressDataData) SetParentNetworkId(v string)`

SetParentNetworkId sets ParentNetworkId field to given value.

### HasParentNetworkId

`func (o *IpamAddressDataData) HasParentNetworkId() bool`

HasParentNetworkId returns a boolean if a field has been set.

### GetParentNetworkName

`func (o *IpamAddressDataData) GetParentNetworkName() string`

GetParentNetworkName returns the ParentNetworkName field if non-nil, zero value otherwise.

### GetParentNetworkNameOk

`func (o *IpamAddressDataData) GetParentNetworkNameOk() (*string, bool)`

GetParentNetworkNameOk returns a tuple with the ParentNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkName

`func (o *IpamAddressDataData) SetParentNetworkName(v string)`

SetParentNetworkName sets ParentNetworkName field to given value.

### HasParentNetworkName

`func (o *IpamAddressDataData) HasParentNetworkName() bool`

HasParentNetworkName returns a boolean if a field has been set.

### GetParentNetworkSize

`func (o *IpamAddressDataData) GetParentNetworkSize() string`

GetParentNetworkSize returns the ParentNetworkSize field if non-nil, zero value otherwise.

### GetParentNetworkSizeOk

`func (o *IpamAddressDataData) GetParentNetworkSizeOk() (*string, bool)`

GetParentNetworkSizeOk returns a tuple with the ParentNetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkSize

`func (o *IpamAddressDataData) SetParentNetworkSize(v string)`

SetParentNetworkSize sets ParentNetworkSize field to given value.

### HasParentNetworkSize

`func (o *IpamAddressDataData) HasParentNetworkSize() bool`

HasParentNetworkSize returns a boolean if a field has been set.

### GetParentNetworkStartHostaddr

`func (o *IpamAddressDataData) GetParentNetworkStartHostaddr() string`

GetParentNetworkStartHostaddr returns the ParentNetworkStartHostaddr field if non-nil, zero value otherwise.

### GetParentNetworkStartHostaddrOk

`func (o *IpamAddressDataData) GetParentNetworkStartHostaddrOk() (*string, bool)`

GetParentNetworkStartHostaddrOk returns a tuple with the ParentNetworkStartHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkStartHostaddr

`func (o *IpamAddressDataData) SetParentNetworkStartHostaddr(v string)`

SetParentNetworkStartHostaddr sets ParentNetworkStartHostaddr field to given value.

### HasParentNetworkStartHostaddr

`func (o *IpamAddressDataData) HasParentNetworkStartHostaddr() bool`

HasParentNetworkStartHostaddr returns a boolean if a field has been set.

### GetParentNetworkStartAddressAddr

`func (o *IpamAddressDataData) GetParentNetworkStartAddressAddr() string`

GetParentNetworkStartAddressAddr returns the ParentNetworkStartAddressAddr field if non-nil, zero value otherwise.

### GetParentNetworkStartAddressAddrOk

`func (o *IpamAddressDataData) GetParentNetworkStartAddressAddrOk() (*string, bool)`

GetParentNetworkStartAddressAddrOk returns a tuple with the ParentNetworkStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkStartAddressAddr

`func (o *IpamAddressDataData) SetParentNetworkStartAddressAddr(v string)`

SetParentNetworkStartAddressAddr sets ParentNetworkStartAddressAddr field to given value.

### HasParentNetworkStartAddressAddr

`func (o *IpamAddressDataData) HasParentNetworkStartAddressAddr() bool`

HasParentNetworkStartAddressAddr returns a boolean if a field has been set.

### GetParentVlsmNetworkId

`func (o *IpamAddressDataData) GetParentVlsmNetworkId() string`

GetParentVlsmNetworkId returns the ParentVlsmNetworkId field if non-nil, zero value otherwise.

### GetParentVlsmNetworkIdOk

`func (o *IpamAddressDataData) GetParentVlsmNetworkIdOk() (*string, bool)`

GetParentVlsmNetworkIdOk returns a tuple with the ParentVlsmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVlsmNetworkId

`func (o *IpamAddressDataData) SetParentVlsmNetworkId(v string)`

SetParentVlsmNetworkId sets ParentVlsmNetworkId field to given value.

### HasParentVlsmNetworkId

`func (o *IpamAddressDataData) HasParentVlsmNetworkId() bool`

HasParentVlsmNetworkId returns a boolean if a field has been set.

### GetPoolClassName

`func (o *IpamAddressDataData) GetPoolClassName() string`

GetPoolClassName returns the PoolClassName field if non-nil, zero value otherwise.

### GetPoolClassNameOk

`func (o *IpamAddressDataData) GetPoolClassNameOk() (*string, bool)`

GetPoolClassNameOk returns a tuple with the PoolClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassName

`func (o *IpamAddressDataData) SetPoolClassName(v string)`

SetPoolClassName sets PoolClassName field to given value.

### HasPoolClassName

`func (o *IpamAddressDataData) HasPoolClassName() bool`

HasPoolClassName returns a boolean if a field has been set.

### GetPoolClassParameters

`func (o *IpamAddressDataData) GetPoolClassParameters() []ApiClassParameterOutputEntry`

GetPoolClassParameters returns the PoolClassParameters field if non-nil, zero value otherwise.

### GetPoolClassParametersOk

`func (o *IpamAddressDataData) GetPoolClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetPoolClassParametersOk returns a tuple with the PoolClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassParameters

`func (o *IpamAddressDataData) SetPoolClassParameters(v []ApiClassParameterOutputEntry)`

SetPoolClassParameters sets PoolClassParameters field to given value.

### HasPoolClassParameters

`func (o *IpamAddressDataData) HasPoolClassParameters() bool`

HasPoolClassParameters returns a boolean if a field has been set.

### GetPoolEndAddressAddr

`func (o *IpamAddressDataData) GetPoolEndAddressAddr() string`

GetPoolEndAddressAddr returns the PoolEndAddressAddr field if non-nil, zero value otherwise.

### GetPoolEndAddressAddrOk

`func (o *IpamAddressDataData) GetPoolEndAddressAddrOk() (*string, bool)`

GetPoolEndAddressAddrOk returns a tuple with the PoolEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEndAddressAddr

`func (o *IpamAddressDataData) SetPoolEndAddressAddr(v string)`

SetPoolEndAddressAddr sets PoolEndAddressAddr field to given value.

### HasPoolEndAddressAddr

`func (o *IpamAddressDataData) HasPoolEndAddressAddr() bool`

HasPoolEndAddressAddr returns a boolean if a field has been set.

### GetPoolId

`func (o *IpamAddressDataData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *IpamAddressDataData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *IpamAddressDataData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *IpamAddressDataData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *IpamAddressDataData) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *IpamAddressDataData) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *IpamAddressDataData) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *IpamAddressDataData) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolReadOnly

`func (o *IpamAddressDataData) GetPoolReadOnly() string`

GetPoolReadOnly returns the PoolReadOnly field if non-nil, zero value otherwise.

### GetPoolReadOnlyOk

`func (o *IpamAddressDataData) GetPoolReadOnlyOk() (*string, bool)`

GetPoolReadOnlyOk returns a tuple with the PoolReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolReadOnly

`func (o *IpamAddressDataData) SetPoolReadOnly(v string)`

SetPoolReadOnly sets PoolReadOnly field to given value.

### HasPoolReadOnly

`func (o *IpamAddressDataData) HasPoolReadOnly() bool`

HasPoolReadOnly returns a boolean if a field has been set.

### GetPoolSize

`func (o *IpamAddressDataData) GetPoolSize() string`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *IpamAddressDataData) GetPoolSizeOk() (*string, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *IpamAddressDataData) SetPoolSize(v string)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *IpamAddressDataData) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.

### GetPoolStartAddressAddr

`func (o *IpamAddressDataData) GetPoolStartAddressAddr() string`

GetPoolStartAddressAddr returns the PoolStartAddressAddr field if non-nil, zero value otherwise.

### GetPoolStartAddressAddrOk

`func (o *IpamAddressDataData) GetPoolStartAddressAddrOk() (*string, bool)`

GetPoolStartAddressAddrOk returns a tuple with the PoolStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStartAddressAddr

`func (o *IpamAddressDataData) SetPoolStartAddressAddr(v string)`

SetPoolStartAddressAddr sets PoolStartAddressAddr field to given value.

### HasPoolStartAddressAddr

`func (o *IpamAddressDataData) HasPoolStartAddressAddr() bool`

HasPoolStartAddressAddr returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *IpamAddressDataData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *IpamAddressDataData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *IpamAddressDataData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *IpamAddressDataData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *IpamAddressDataData) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *IpamAddressDataData) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *IpamAddressDataData) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *IpamAddressDataData) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *IpamAddressDataData) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *IpamAddressDataData) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *IpamAddressDataData) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *IpamAddressDataData) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAddressDataData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAddressDataData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAddressDataData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAddressDataData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceIsTemplate

`func (o *IpamAddressDataData) GetSpaceIsTemplate() string`

GetSpaceIsTemplate returns the SpaceIsTemplate field if non-nil, zero value otherwise.

### GetSpaceIsTemplateOk

`func (o *IpamAddressDataData) GetSpaceIsTemplateOk() (*string, bool)`

GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceIsTemplate

`func (o *IpamAddressDataData) SetSpaceIsTemplate(v string)`

SetSpaceIsTemplate sets SpaceIsTemplate field to given value.

### HasSpaceIsTemplate

`func (o *IpamAddressDataData) HasSpaceIsTemplate() bool`

HasSpaceIsTemplate returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAddressDataData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAddressDataData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAddressDataData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAddressDataData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetNetworkClassName

`func (o *IpamAddressDataData) GetNetworkClassName() string`

GetNetworkClassName returns the NetworkClassName field if non-nil, zero value otherwise.

### GetNetworkClassNameOk

`func (o *IpamAddressDataData) GetNetworkClassNameOk() (*string, bool)`

GetNetworkClassNameOk returns a tuple with the NetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassName

`func (o *IpamAddressDataData) SetNetworkClassName(v string)`

SetNetworkClassName sets NetworkClassName field to given value.

### HasNetworkClassName

`func (o *IpamAddressDataData) HasNetworkClassName() bool`

HasNetworkClassName returns a boolean if a field has been set.

### GetNetworkClassParameters

`func (o *IpamAddressDataData) GetNetworkClassParameters() []ApiClassParameterOutputEntry`

GetNetworkClassParameters returns the NetworkClassParameters field if non-nil, zero value otherwise.

### GetNetworkClassParametersOk

`func (o *IpamAddressDataData) GetNetworkClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetNetworkClassParametersOk returns a tuple with the NetworkClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassParameters

`func (o *IpamAddressDataData) SetNetworkClassParameters(v []ApiClassParameterOutputEntry)`

SetNetworkClassParameters sets NetworkClassParameters field to given value.

### HasNetworkClassParameters

`func (o *IpamAddressDataData) HasNetworkClassParameters() bool`

HasNetworkClassParameters returns a boolean if a field has been set.

### GetNetworkEndAddressAddr

`func (o *IpamAddressDataData) GetNetworkEndAddressAddr() string`

GetNetworkEndAddressAddr returns the NetworkEndAddressAddr field if non-nil, zero value otherwise.

### GetNetworkEndAddressAddrOk

`func (o *IpamAddressDataData) GetNetworkEndAddressAddrOk() (*string, bool)`

GetNetworkEndAddressAddrOk returns a tuple with the NetworkEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndAddressAddr

`func (o *IpamAddressDataData) SetNetworkEndAddressAddr(v string)`

SetNetworkEndAddressAddr sets NetworkEndAddressAddr field to given value.

### HasNetworkEndAddressAddr

`func (o *IpamAddressDataData) HasNetworkEndAddressAddr() bool`

HasNetworkEndAddressAddr returns a boolean if a field has been set.

### GetNetworkId

`func (o *IpamAddressDataData) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *IpamAddressDataData) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *IpamAddressDataData) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *IpamAddressDataData) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkIsTerminal

`func (o *IpamAddressDataData) GetNetworkIsTerminal() string`

GetNetworkIsTerminal returns the NetworkIsTerminal field if non-nil, zero value otherwise.

### GetNetworkIsTerminalOk

`func (o *IpamAddressDataData) GetNetworkIsTerminalOk() (*string, bool)`

GetNetworkIsTerminalOk returns a tuple with the NetworkIsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsTerminal

`func (o *IpamAddressDataData) SetNetworkIsTerminal(v string)`

SetNetworkIsTerminal sets NetworkIsTerminal field to given value.

### HasNetworkIsTerminal

`func (o *IpamAddressDataData) HasNetworkIsTerminal() bool`

HasNetworkIsTerminal returns a boolean if a field has been set.

### GetNetworkName

`func (o *IpamAddressDataData) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *IpamAddressDataData) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *IpamAddressDataData) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *IpamAddressDataData) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkSize

`func (o *IpamAddressDataData) GetNetworkSize() string`

GetNetworkSize returns the NetworkSize field if non-nil, zero value otherwise.

### GetNetworkSizeOk

`func (o *IpamAddressDataData) GetNetworkSizeOk() (*string, bool)`

GetNetworkSizeOk returns a tuple with the NetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSize

`func (o *IpamAddressDataData) SetNetworkSize(v string)`

SetNetworkSize sets NetworkSize field to given value.

### HasNetworkSize

`func (o *IpamAddressDataData) HasNetworkSize() bool`

HasNetworkSize returns a boolean if a field has been set.

### GetNetworkStartAddressAddr

`func (o *IpamAddressDataData) GetNetworkStartAddressAddr() string`

GetNetworkStartAddressAddr returns the NetworkStartAddressAddr field if non-nil, zero value otherwise.

### GetNetworkStartAddressAddrOk

`func (o *IpamAddressDataData) GetNetworkStartAddressAddrOk() (*string, bool)`

GetNetworkStartAddressAddrOk returns a tuple with the NetworkStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStartAddressAddr

`func (o *IpamAddressDataData) SetNetworkStartAddressAddr(v string)`

SetNetworkStartAddressAddr sets NetworkStartAddressAddr field to given value.

### HasNetworkStartAddressAddr

`func (o *IpamAddressDataData) HasNetworkStartAddressAddr() bool`

HasNetworkStartAddressAddr returns a boolean if a field has been set.

### GetTagContainerDhcpstatic

`func (o *IpamAddressDataData) GetTagContainerDhcpstatic() string`

GetTagContainerDhcpstatic returns the TagContainerDhcpstatic field if non-nil, zero value otherwise.

### GetTagContainerDhcpstaticOk

`func (o *IpamAddressDataData) GetTagContainerDhcpstaticOk() (*string, bool)`

GetTagContainerDhcpstaticOk returns a tuple with the TagContainerDhcpstatic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagContainerDhcpstatic

`func (o *IpamAddressDataData) SetTagContainerDhcpstatic(v string)`

SetTagContainerDhcpstatic sets TagContainerDhcpstatic field to given value.

### HasTagContainerDhcpstatic

`func (o *IpamAddressDataData) HasTagContainerDhcpstatic() bool`

HasTagContainerDhcpstatic returns a boolean if a field has been set.

### GetTagPoolDhcprange

`func (o *IpamAddressDataData) GetTagPoolDhcprange() string`

GetTagPoolDhcprange returns the TagPoolDhcprange field if non-nil, zero value otherwise.

### GetTagPoolDhcprangeOk

`func (o *IpamAddressDataData) GetTagPoolDhcprangeOk() (*string, bool)`

GetTagPoolDhcprangeOk returns a tuple with the TagPoolDhcprange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagPoolDhcprange

`func (o *IpamAddressDataData) SetTagPoolDhcprange(v string)`

SetTagPoolDhcprange sets TagPoolDhcprange field to given value.

### HasTagPoolDhcprange

`func (o *IpamAddressDataData) HasTagPoolDhcprange() bool`

HasTagPoolDhcprange returns a boolean if a field has been set.

### GetAddressType

`func (o *IpamAddressDataData) GetAddressType() string`

GetAddressType returns the AddressType field if non-nil, zero value otherwise.

### GetAddressTypeOk

`func (o *IpamAddressDataData) GetAddressTypeOk() (*string, bool)`

GetAddressTypeOk returns a tuple with the AddressType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressType

`func (o *IpamAddressDataData) SetAddressType(v string)`

SetAddressType sets AddressType field to given value.

### HasAddressType

`func (o *IpamAddressDataData) HasAddressType() bool`

HasAddressType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


