# IpamAddress6DataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeEndAddress6Addr** | Pointer to **string** | An IP address in hexadecimal format. For addresses &lt;b&gt;In use&lt;/b&gt; (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;ip6&lt;/b&gt;), it returns the IP address itself.For free addresses (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;), it returns the last IP address of a range of IPv6 addresses that are not assigned yet. The first address in that range is returned in &lt;b&gt;free_start_ip6_addr&lt;/b&gt;. | [optional] 
**FreeScopeSize** | Pointer to **string** | The number of IP addresses that are not assigned yet (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;) between &lt;b&gt;free_start_ip6_addr&lt;/b&gt; and &lt;b&gt;free_end_ip6_addr&lt;/b&gt;. | [optional] 
**FreeStartAddress6Addr** | Pointer to **string** | An IP address in hexadecimal format. For addresses &lt;b&gt;In use&lt;/b&gt; (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;ip6&lt;/b&gt;), it returns the IP address itself.For free addresses (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;), it returns the first IP address of a range of IPv6 addresses that are not assigned yet. The last address in that range is returned in &lt;b&gt;free_end_ip6_addr&lt;/b&gt;. | [optional] 
**Address6Hostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;ip6_addr&lt;/b&gt;. | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device associated with the IP address. | [optional] 
**DeviceName** | Pointer to **string** | The name of the Device Manager device associated with the IP address. | [optional] 
**InterfaceId** | Pointer to **string** | The database identifier (ID) of the Device Manager interface associated with the IP address. | [optional] 
**InterfaceName** | Pointer to **string** | The name of the Device Manager interface associated with the IP address. | [optional] 
**Address6Addr** | Pointer to **string** | The IPv6 address itself. | [optional] 
**Address6Alias** | Pointer to **string** | The name of the IPv6 alias(es) associated with the IPv6 address. | [optional] 
**Address6ClassName** | Pointer to **string** | The name of the class applied to the IPv6 address, it can be preceded by the class directory. | [optional] 
**Address6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv6 address and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Address6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 address. | [optional] 
**Address6MacAddr** | Pointer to **string** | The MAC address associated with the IPv6 address. | [optional] 
**Address6Name** | Pointer to **string** | The name of the IPv6 address. | [optional] 
**Network6LockNetworkBroadcast** | Pointer to **string** | A way to prevent &lt;b&gt;(1)&lt;/b&gt; users from assigning the broadcast IP address and network IP address of the network the IP address belongs to. | [optional] 
**Address6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentSpaceName** | Pointer to **string** | The name of the space where is located the parent of the network the IPv6 address belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the network the IPv6 address belongs to has no parent network. | [optional] 
**ParentNetwork6ClassName** | Pointer to **string** | The name of the class applied to the parent of the IPv6 network the object belongs to, it can be preceded by the class directory. | [optional] 
**ParentNetwork6EndHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;parent_subnet6_end_ip6_addr&lt;/b&gt;. | [optional] 
**ParentNetwork6EndAddress6Addr** | Pointer to **string** | The last IP address of the parent of the IPv6 network the IP address belongs to. | [optional] 
**ParentNetwork6Id** | Pointer to **string** | The database identifier (ID) of the parent IPv6 network. It identifies the parent of the IPv6 network the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the network the object belongs to has no parent network. | [optional] 
**ParentNetwork6Name** | Pointer to **string** | The name of the parent IPv6 network. &lt;b&gt;#&lt;/b&gt; indicates that the network the object belongs to has no parent network. | [optional] 
**ParentNetwork6Prefix** | Pointer to **string** | The prefix of the parent of the IPv6 network the object belongs to. | [optional] 
**ParentNetwork6Size** | Pointer to **string** | The number of IP addresses of the network parent, in hexadecimal format. | [optional] 
**ParentNetwork6StartHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;parent_subnet6_start_ip6_addr&lt;/b&gt;. | [optional] 
**ParentNetwork6StartAddress6Addr** | Pointer to **string** | The first IP address of the parent of the IPv6 network the IP address belongs to. | [optional] 
**ParentVlsmNetwork6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 subnet-type network, located in the VLSM parent space, from which the parent of the network the IP address belongs to was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the parent of the network the IP address belongs to is not a VLSM block-type network duplicated from a parent space. | [optional] 
**Pool6ClassName** | Pointer to **string** | The name of the class applied to the IPv6 pool the object belongs to, it can be preceded by the class directory. | [optional] 
**Pool6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv6 pool the object belongs to and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Pool6EndAddress6Addr** | Pointer to **string** | The last IP address of the IPv6 pool the IP address belongs to. | [optional] 
**Pool6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 pool the object belongs to. | [optional] 
**Pool6Name** | Pointer to **string** | The name of the IPv6 pool the object belongs to. | [optional] 
**Pool6ReadOnly** | Pointer to **string** | The reservation status of the pool the IPv6 address belongs to. If set &lt;b&gt;1&lt;/b&gt;, the pool is reserved and you cannot assign the IP address. | [optional] 
**Pool6Size** | Pointer to **string** | The number of IP addresses that contains the pool the IPv6 address belongs to. | [optional] 
**Pool6StartAddress6Addr** | Pointer to **string** | The first IP address of the IPv6 pool the IP address belongs to. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space the object belongs to. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**Network6ClassName** | Pointer to **string** | The name of the class applied to the IPv6 network the object belongs to, it can be preceded by the class directory. | [optional] 
**Network6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv6 network the object belongs to and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Network6EndAddress6Addr** | Pointer to **string** | The last IP address of the IPv6 network the object belongs to. | [optional] 
**Network6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 network the object belongs to. | [optional] 
**Network6IsTerminal** | Pointer to **string** | A way to determine if the network the IP address belongs to is terminal (&lt;b&gt;1&lt;/b&gt;) or non-terminal (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**Network6Name** | Pointer to **string** | The name of the IPv6 network the object belongs to. | [optional] 
**Network6Prefix** | Pointer to **string** | The prefix of the IPv6 network the object belongs to. | [optional] 
**Network6Size** | Pointer to **string** | The number of IP addresses the network the object belongs to contains. | [optional] 
**Network6StartAddress6Addr** | Pointer to **string** | The first IP address of the IPv6 network the object belongs to. | [optional] 
**NetworkSize** | Pointer to **string** | The number of IP addresses the network the object belongs to contains. | [optional] 
**Address6Type** | Pointer to **string** | A way to determine if you can assign the IP address (&lt;b&gt;free&lt;/b&gt;) or if it is In use (&lt;b&gt;ip&lt;/b&gt;6). | [optional] 
**VlsmNetwork6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 subnet-type network, located in the VLSM parent space, from which the network the IP address belongs to was duplicated. &lt;b&gt;0&lt;/b&gt; indicates the network the IP address belongs to is not a VLSM block-type network duplicated from a parent space. | [optional] 

## Methods

### NewIpamAddress6DataData

`func NewIpamAddress6DataData() *IpamAddress6DataData`

NewIpamAddress6DataData instantiates a new IpamAddress6DataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAddress6DataDataWithDefaults

`func NewIpamAddress6DataDataWithDefaults() *IpamAddress6DataData`

NewIpamAddress6DataDataWithDefaults instantiates a new IpamAddress6DataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeEndAddress6Addr

`func (o *IpamAddress6DataData) GetFreeEndAddress6Addr() string`

GetFreeEndAddress6Addr returns the FreeEndAddress6Addr field if non-nil, zero value otherwise.

### GetFreeEndAddress6AddrOk

`func (o *IpamAddress6DataData) GetFreeEndAddress6AddrOk() (*string, bool)`

GetFreeEndAddress6AddrOk returns a tuple with the FreeEndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeEndAddress6Addr

`func (o *IpamAddress6DataData) SetFreeEndAddress6Addr(v string)`

SetFreeEndAddress6Addr sets FreeEndAddress6Addr field to given value.

### HasFreeEndAddress6Addr

`func (o *IpamAddress6DataData) HasFreeEndAddress6Addr() bool`

HasFreeEndAddress6Addr returns a boolean if a field has been set.

### GetFreeScopeSize

`func (o *IpamAddress6DataData) GetFreeScopeSize() string`

GetFreeScopeSize returns the FreeScopeSize field if non-nil, zero value otherwise.

### GetFreeScopeSizeOk

`func (o *IpamAddress6DataData) GetFreeScopeSizeOk() (*string, bool)`

GetFreeScopeSizeOk returns a tuple with the FreeScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeScopeSize

`func (o *IpamAddress6DataData) SetFreeScopeSize(v string)`

SetFreeScopeSize sets FreeScopeSize field to given value.

### HasFreeScopeSize

`func (o *IpamAddress6DataData) HasFreeScopeSize() bool`

HasFreeScopeSize returns a boolean if a field has been set.

### GetFreeStartAddress6Addr

`func (o *IpamAddress6DataData) GetFreeStartAddress6Addr() string`

GetFreeStartAddress6Addr returns the FreeStartAddress6Addr field if non-nil, zero value otherwise.

### GetFreeStartAddress6AddrOk

`func (o *IpamAddress6DataData) GetFreeStartAddress6AddrOk() (*string, bool)`

GetFreeStartAddress6AddrOk returns a tuple with the FreeStartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStartAddress6Addr

`func (o *IpamAddress6DataData) SetFreeStartAddress6Addr(v string)`

SetFreeStartAddress6Addr sets FreeStartAddress6Addr field to given value.

### HasFreeStartAddress6Addr

`func (o *IpamAddress6DataData) HasFreeStartAddress6Addr() bool`

HasFreeStartAddress6Addr returns a boolean if a field has been set.

### GetAddress6Hostaddr

`func (o *IpamAddress6DataData) GetAddress6Hostaddr() string`

GetAddress6Hostaddr returns the Address6Hostaddr field if non-nil, zero value otherwise.

### GetAddress6HostaddrOk

`func (o *IpamAddress6DataData) GetAddress6HostaddrOk() (*string, bool)`

GetAddress6HostaddrOk returns a tuple with the Address6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Hostaddr

`func (o *IpamAddress6DataData) SetAddress6Hostaddr(v string)`

SetAddress6Hostaddr sets Address6Hostaddr field to given value.

### HasAddress6Hostaddr

`func (o *IpamAddress6DataData) HasAddress6Hostaddr() bool`

HasAddress6Hostaddr returns a boolean if a field has been set.

### GetDeviceId

`func (o *IpamAddress6DataData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *IpamAddress6DataData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *IpamAddress6DataData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *IpamAddress6DataData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *IpamAddress6DataData) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *IpamAddress6DataData) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *IpamAddress6DataData) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *IpamAddress6DataData) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetInterfaceId

`func (o *IpamAddress6DataData) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *IpamAddress6DataData) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *IpamAddress6DataData) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *IpamAddress6DataData) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceName

`func (o *IpamAddress6DataData) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *IpamAddress6DataData) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *IpamAddress6DataData) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *IpamAddress6DataData) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetAddress6Addr

`func (o *IpamAddress6DataData) GetAddress6Addr() string`

GetAddress6Addr returns the Address6Addr field if non-nil, zero value otherwise.

### GetAddress6AddrOk

`func (o *IpamAddress6DataData) GetAddress6AddrOk() (*string, bool)`

GetAddress6AddrOk returns a tuple with the Address6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Addr

`func (o *IpamAddress6DataData) SetAddress6Addr(v string)`

SetAddress6Addr sets Address6Addr field to given value.

### HasAddress6Addr

`func (o *IpamAddress6DataData) HasAddress6Addr() bool`

HasAddress6Addr returns a boolean if a field has been set.

### GetAddress6Alias

`func (o *IpamAddress6DataData) GetAddress6Alias() string`

GetAddress6Alias returns the Address6Alias field if non-nil, zero value otherwise.

### GetAddress6AliasOk

`func (o *IpamAddress6DataData) GetAddress6AliasOk() (*string, bool)`

GetAddress6AliasOk returns a tuple with the Address6Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Alias

`func (o *IpamAddress6DataData) SetAddress6Alias(v string)`

SetAddress6Alias sets Address6Alias field to given value.

### HasAddress6Alias

`func (o *IpamAddress6DataData) HasAddress6Alias() bool`

HasAddress6Alias returns a boolean if a field has been set.

### GetAddress6ClassName

`func (o *IpamAddress6DataData) GetAddress6ClassName() string`

GetAddress6ClassName returns the Address6ClassName field if non-nil, zero value otherwise.

### GetAddress6ClassNameOk

`func (o *IpamAddress6DataData) GetAddress6ClassNameOk() (*string, bool)`

GetAddress6ClassNameOk returns a tuple with the Address6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6ClassName

`func (o *IpamAddress6DataData) SetAddress6ClassName(v string)`

SetAddress6ClassName sets Address6ClassName field to given value.

### HasAddress6ClassName

`func (o *IpamAddress6DataData) HasAddress6ClassName() bool`

HasAddress6ClassName returns a boolean if a field has been set.

### GetAddress6ClassParameters

`func (o *IpamAddress6DataData) GetAddress6ClassParameters() []ApiClassParameterOutputEntry`

GetAddress6ClassParameters returns the Address6ClassParameters field if non-nil, zero value otherwise.

### GetAddress6ClassParametersOk

`func (o *IpamAddress6DataData) GetAddress6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetAddress6ClassParametersOk returns a tuple with the Address6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6ClassParameters

`func (o *IpamAddress6DataData) SetAddress6ClassParameters(v []ApiClassParameterOutputEntry)`

SetAddress6ClassParameters sets Address6ClassParameters field to given value.

### HasAddress6ClassParameters

`func (o *IpamAddress6DataData) HasAddress6ClassParameters() bool`

HasAddress6ClassParameters returns a boolean if a field has been set.

### GetAddress6Id

`func (o *IpamAddress6DataData) GetAddress6Id() string`

GetAddress6Id returns the Address6Id field if non-nil, zero value otherwise.

### GetAddress6IdOk

`func (o *IpamAddress6DataData) GetAddress6IdOk() (*string, bool)`

GetAddress6IdOk returns a tuple with the Address6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Id

`func (o *IpamAddress6DataData) SetAddress6Id(v string)`

SetAddress6Id sets Address6Id field to given value.

### HasAddress6Id

`func (o *IpamAddress6DataData) HasAddress6Id() bool`

HasAddress6Id returns a boolean if a field has been set.

### GetAddress6MacAddr

`func (o *IpamAddress6DataData) GetAddress6MacAddr() string`

GetAddress6MacAddr returns the Address6MacAddr field if non-nil, zero value otherwise.

### GetAddress6MacAddrOk

`func (o *IpamAddress6DataData) GetAddress6MacAddrOk() (*string, bool)`

GetAddress6MacAddrOk returns a tuple with the Address6MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6MacAddr

`func (o *IpamAddress6DataData) SetAddress6MacAddr(v string)`

SetAddress6MacAddr sets Address6MacAddr field to given value.

### HasAddress6MacAddr

`func (o *IpamAddress6DataData) HasAddress6MacAddr() bool`

HasAddress6MacAddr returns a boolean if a field has been set.

### GetAddress6Name

`func (o *IpamAddress6DataData) GetAddress6Name() string`

GetAddress6Name returns the Address6Name field if non-nil, zero value otherwise.

### GetAddress6NameOk

`func (o *IpamAddress6DataData) GetAddress6NameOk() (*string, bool)`

GetAddress6NameOk returns a tuple with the Address6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Name

`func (o *IpamAddress6DataData) SetAddress6Name(v string)`

SetAddress6Name sets Address6Name field to given value.

### HasAddress6Name

`func (o *IpamAddress6DataData) HasAddress6Name() bool`

HasAddress6Name returns a boolean if a field has been set.

### GetNetwork6LockNetworkBroadcast

`func (o *IpamAddress6DataData) GetNetwork6LockNetworkBroadcast() string`

GetNetwork6LockNetworkBroadcast returns the Network6LockNetworkBroadcast field if non-nil, zero value otherwise.

### GetNetwork6LockNetworkBroadcastOk

`func (o *IpamAddress6DataData) GetNetwork6LockNetworkBroadcastOk() (*string, bool)`

GetNetwork6LockNetworkBroadcastOk returns a tuple with the Network6LockNetworkBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6LockNetworkBroadcast

`func (o *IpamAddress6DataData) SetNetwork6LockNetworkBroadcast(v string)`

SetNetwork6LockNetworkBroadcast sets Network6LockNetworkBroadcast field to given value.

### HasNetwork6LockNetworkBroadcast

`func (o *IpamAddress6DataData) HasNetwork6LockNetworkBroadcast() bool`

HasNetwork6LockNetworkBroadcast returns a boolean if a field has been set.

### GetAddress6Multistatus

`func (o *IpamAddress6DataData) GetAddress6Multistatus() string`

GetAddress6Multistatus returns the Address6Multistatus field if non-nil, zero value otherwise.

### GetAddress6MultistatusOk

`func (o *IpamAddress6DataData) GetAddress6MultistatusOk() (*string, bool)`

GetAddress6MultistatusOk returns a tuple with the Address6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Multistatus

`func (o *IpamAddress6DataData) SetAddress6Multistatus(v string)`

SetAddress6Multistatus sets Address6Multistatus field to given value.

### HasAddress6Multistatus

`func (o *IpamAddress6DataData) HasAddress6Multistatus() bool`

HasAddress6Multistatus returns a boolean if a field has been set.

### GetParentSpaceName

`func (o *IpamAddress6DataData) GetParentSpaceName() string`

GetParentSpaceName returns the ParentSpaceName field if non-nil, zero value otherwise.

### GetParentSpaceNameOk

`func (o *IpamAddress6DataData) GetParentSpaceNameOk() (*string, bool)`

GetParentSpaceNameOk returns a tuple with the ParentSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceName

`func (o *IpamAddress6DataData) SetParentSpaceName(v string)`

SetParentSpaceName sets ParentSpaceName field to given value.

### HasParentSpaceName

`func (o *IpamAddress6DataData) HasParentSpaceName() bool`

HasParentSpaceName returns a boolean if a field has been set.

### GetParentNetwork6ClassName

`func (o *IpamAddress6DataData) GetParentNetwork6ClassName() string`

GetParentNetwork6ClassName returns the ParentNetwork6ClassName field if non-nil, zero value otherwise.

### GetParentNetwork6ClassNameOk

`func (o *IpamAddress6DataData) GetParentNetwork6ClassNameOk() (*string, bool)`

GetParentNetwork6ClassNameOk returns a tuple with the ParentNetwork6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6ClassName

`func (o *IpamAddress6DataData) SetParentNetwork6ClassName(v string)`

SetParentNetwork6ClassName sets ParentNetwork6ClassName field to given value.

### HasParentNetwork6ClassName

`func (o *IpamAddress6DataData) HasParentNetwork6ClassName() bool`

HasParentNetwork6ClassName returns a boolean if a field has been set.

### GetParentNetwork6EndHostaddr

`func (o *IpamAddress6DataData) GetParentNetwork6EndHostaddr() string`

GetParentNetwork6EndHostaddr returns the ParentNetwork6EndHostaddr field if non-nil, zero value otherwise.

### GetParentNetwork6EndHostaddrOk

`func (o *IpamAddress6DataData) GetParentNetwork6EndHostaddrOk() (*string, bool)`

GetParentNetwork6EndHostaddrOk returns a tuple with the ParentNetwork6EndHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6EndHostaddr

`func (o *IpamAddress6DataData) SetParentNetwork6EndHostaddr(v string)`

SetParentNetwork6EndHostaddr sets ParentNetwork6EndHostaddr field to given value.

### HasParentNetwork6EndHostaddr

`func (o *IpamAddress6DataData) HasParentNetwork6EndHostaddr() bool`

HasParentNetwork6EndHostaddr returns a boolean if a field has been set.

### GetParentNetwork6EndAddress6Addr

`func (o *IpamAddress6DataData) GetParentNetwork6EndAddress6Addr() string`

GetParentNetwork6EndAddress6Addr returns the ParentNetwork6EndAddress6Addr field if non-nil, zero value otherwise.

### GetParentNetwork6EndAddress6AddrOk

`func (o *IpamAddress6DataData) GetParentNetwork6EndAddress6AddrOk() (*string, bool)`

GetParentNetwork6EndAddress6AddrOk returns a tuple with the ParentNetwork6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6EndAddress6Addr

`func (o *IpamAddress6DataData) SetParentNetwork6EndAddress6Addr(v string)`

SetParentNetwork6EndAddress6Addr sets ParentNetwork6EndAddress6Addr field to given value.

### HasParentNetwork6EndAddress6Addr

`func (o *IpamAddress6DataData) HasParentNetwork6EndAddress6Addr() bool`

HasParentNetwork6EndAddress6Addr returns a boolean if a field has been set.

### GetParentNetwork6Id

`func (o *IpamAddress6DataData) GetParentNetwork6Id() string`

GetParentNetwork6Id returns the ParentNetwork6Id field if non-nil, zero value otherwise.

### GetParentNetwork6IdOk

`func (o *IpamAddress6DataData) GetParentNetwork6IdOk() (*string, bool)`

GetParentNetwork6IdOk returns a tuple with the ParentNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Id

`func (o *IpamAddress6DataData) SetParentNetwork6Id(v string)`

SetParentNetwork6Id sets ParentNetwork6Id field to given value.

### HasParentNetwork6Id

`func (o *IpamAddress6DataData) HasParentNetwork6Id() bool`

HasParentNetwork6Id returns a boolean if a field has been set.

### GetParentNetwork6Name

`func (o *IpamAddress6DataData) GetParentNetwork6Name() string`

GetParentNetwork6Name returns the ParentNetwork6Name field if non-nil, zero value otherwise.

### GetParentNetwork6NameOk

`func (o *IpamAddress6DataData) GetParentNetwork6NameOk() (*string, bool)`

GetParentNetwork6NameOk returns a tuple with the ParentNetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Name

`func (o *IpamAddress6DataData) SetParentNetwork6Name(v string)`

SetParentNetwork6Name sets ParentNetwork6Name field to given value.

### HasParentNetwork6Name

`func (o *IpamAddress6DataData) HasParentNetwork6Name() bool`

HasParentNetwork6Name returns a boolean if a field has been set.

### GetParentNetwork6Prefix

`func (o *IpamAddress6DataData) GetParentNetwork6Prefix() string`

GetParentNetwork6Prefix returns the ParentNetwork6Prefix field if non-nil, zero value otherwise.

### GetParentNetwork6PrefixOk

`func (o *IpamAddress6DataData) GetParentNetwork6PrefixOk() (*string, bool)`

GetParentNetwork6PrefixOk returns a tuple with the ParentNetwork6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Prefix

`func (o *IpamAddress6DataData) SetParentNetwork6Prefix(v string)`

SetParentNetwork6Prefix sets ParentNetwork6Prefix field to given value.

### HasParentNetwork6Prefix

`func (o *IpamAddress6DataData) HasParentNetwork6Prefix() bool`

HasParentNetwork6Prefix returns a boolean if a field has been set.

### GetParentNetwork6Size

`func (o *IpamAddress6DataData) GetParentNetwork6Size() string`

GetParentNetwork6Size returns the ParentNetwork6Size field if non-nil, zero value otherwise.

### GetParentNetwork6SizeOk

`func (o *IpamAddress6DataData) GetParentNetwork6SizeOk() (*string, bool)`

GetParentNetwork6SizeOk returns a tuple with the ParentNetwork6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Size

`func (o *IpamAddress6DataData) SetParentNetwork6Size(v string)`

SetParentNetwork6Size sets ParentNetwork6Size field to given value.

### HasParentNetwork6Size

`func (o *IpamAddress6DataData) HasParentNetwork6Size() bool`

HasParentNetwork6Size returns a boolean if a field has been set.

### GetParentNetwork6StartHostaddr

`func (o *IpamAddress6DataData) GetParentNetwork6StartHostaddr() string`

GetParentNetwork6StartHostaddr returns the ParentNetwork6StartHostaddr field if non-nil, zero value otherwise.

### GetParentNetwork6StartHostaddrOk

`func (o *IpamAddress6DataData) GetParentNetwork6StartHostaddrOk() (*string, bool)`

GetParentNetwork6StartHostaddrOk returns a tuple with the ParentNetwork6StartHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6StartHostaddr

`func (o *IpamAddress6DataData) SetParentNetwork6StartHostaddr(v string)`

SetParentNetwork6StartHostaddr sets ParentNetwork6StartHostaddr field to given value.

### HasParentNetwork6StartHostaddr

`func (o *IpamAddress6DataData) HasParentNetwork6StartHostaddr() bool`

HasParentNetwork6StartHostaddr returns a boolean if a field has been set.

### GetParentNetwork6StartAddress6Addr

`func (o *IpamAddress6DataData) GetParentNetwork6StartAddress6Addr() string`

GetParentNetwork6StartAddress6Addr returns the ParentNetwork6StartAddress6Addr field if non-nil, zero value otherwise.

### GetParentNetwork6StartAddress6AddrOk

`func (o *IpamAddress6DataData) GetParentNetwork6StartAddress6AddrOk() (*string, bool)`

GetParentNetwork6StartAddress6AddrOk returns a tuple with the ParentNetwork6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6StartAddress6Addr

`func (o *IpamAddress6DataData) SetParentNetwork6StartAddress6Addr(v string)`

SetParentNetwork6StartAddress6Addr sets ParentNetwork6StartAddress6Addr field to given value.

### HasParentNetwork6StartAddress6Addr

`func (o *IpamAddress6DataData) HasParentNetwork6StartAddress6Addr() bool`

HasParentNetwork6StartAddress6Addr returns a boolean if a field has been set.

### GetParentVlsmNetwork6Id

`func (o *IpamAddress6DataData) GetParentVlsmNetwork6Id() string`

GetParentVlsmNetwork6Id returns the ParentVlsmNetwork6Id field if non-nil, zero value otherwise.

### GetParentVlsmNetwork6IdOk

`func (o *IpamAddress6DataData) GetParentVlsmNetwork6IdOk() (*string, bool)`

GetParentVlsmNetwork6IdOk returns a tuple with the ParentVlsmNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVlsmNetwork6Id

`func (o *IpamAddress6DataData) SetParentVlsmNetwork6Id(v string)`

SetParentVlsmNetwork6Id sets ParentVlsmNetwork6Id field to given value.

### HasParentVlsmNetwork6Id

`func (o *IpamAddress6DataData) HasParentVlsmNetwork6Id() bool`

HasParentVlsmNetwork6Id returns a boolean if a field has been set.

### GetPool6ClassName

`func (o *IpamAddress6DataData) GetPool6ClassName() string`

GetPool6ClassName returns the Pool6ClassName field if non-nil, zero value otherwise.

### GetPool6ClassNameOk

`func (o *IpamAddress6DataData) GetPool6ClassNameOk() (*string, bool)`

GetPool6ClassNameOk returns a tuple with the Pool6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ClassName

`func (o *IpamAddress6DataData) SetPool6ClassName(v string)`

SetPool6ClassName sets Pool6ClassName field to given value.

### HasPool6ClassName

`func (o *IpamAddress6DataData) HasPool6ClassName() bool`

HasPool6ClassName returns a boolean if a field has been set.

### GetPool6ClassParameters

`func (o *IpamAddress6DataData) GetPool6ClassParameters() []ApiClassParameterOutputEntry`

GetPool6ClassParameters returns the Pool6ClassParameters field if non-nil, zero value otherwise.

### GetPool6ClassParametersOk

`func (o *IpamAddress6DataData) GetPool6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetPool6ClassParametersOk returns a tuple with the Pool6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ClassParameters

`func (o *IpamAddress6DataData) SetPool6ClassParameters(v []ApiClassParameterOutputEntry)`

SetPool6ClassParameters sets Pool6ClassParameters field to given value.

### HasPool6ClassParameters

`func (o *IpamAddress6DataData) HasPool6ClassParameters() bool`

HasPool6ClassParameters returns a boolean if a field has been set.

### GetPool6EndAddress6Addr

`func (o *IpamAddress6DataData) GetPool6EndAddress6Addr() string`

GetPool6EndAddress6Addr returns the Pool6EndAddress6Addr field if non-nil, zero value otherwise.

### GetPool6EndAddress6AddrOk

`func (o *IpamAddress6DataData) GetPool6EndAddress6AddrOk() (*string, bool)`

GetPool6EndAddress6AddrOk returns a tuple with the Pool6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6EndAddress6Addr

`func (o *IpamAddress6DataData) SetPool6EndAddress6Addr(v string)`

SetPool6EndAddress6Addr sets Pool6EndAddress6Addr field to given value.

### HasPool6EndAddress6Addr

`func (o *IpamAddress6DataData) HasPool6EndAddress6Addr() bool`

HasPool6EndAddress6Addr returns a boolean if a field has been set.

### GetPool6Id

`func (o *IpamAddress6DataData) GetPool6Id() string`

GetPool6Id returns the Pool6Id field if non-nil, zero value otherwise.

### GetPool6IdOk

`func (o *IpamAddress6DataData) GetPool6IdOk() (*string, bool)`

GetPool6IdOk returns a tuple with the Pool6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Id

`func (o *IpamAddress6DataData) SetPool6Id(v string)`

SetPool6Id sets Pool6Id field to given value.

### HasPool6Id

`func (o *IpamAddress6DataData) HasPool6Id() bool`

HasPool6Id returns a boolean if a field has been set.

### GetPool6Name

`func (o *IpamAddress6DataData) GetPool6Name() string`

GetPool6Name returns the Pool6Name field if non-nil, zero value otherwise.

### GetPool6NameOk

`func (o *IpamAddress6DataData) GetPool6NameOk() (*string, bool)`

GetPool6NameOk returns a tuple with the Pool6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Name

`func (o *IpamAddress6DataData) SetPool6Name(v string)`

SetPool6Name sets Pool6Name field to given value.

### HasPool6Name

`func (o *IpamAddress6DataData) HasPool6Name() bool`

HasPool6Name returns a boolean if a field has been set.

### GetPool6ReadOnly

`func (o *IpamAddress6DataData) GetPool6ReadOnly() string`

GetPool6ReadOnly returns the Pool6ReadOnly field if non-nil, zero value otherwise.

### GetPool6ReadOnlyOk

`func (o *IpamAddress6DataData) GetPool6ReadOnlyOk() (*string, bool)`

GetPool6ReadOnlyOk returns a tuple with the Pool6ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ReadOnly

`func (o *IpamAddress6DataData) SetPool6ReadOnly(v string)`

SetPool6ReadOnly sets Pool6ReadOnly field to given value.

### HasPool6ReadOnly

`func (o *IpamAddress6DataData) HasPool6ReadOnly() bool`

HasPool6ReadOnly returns a boolean if a field has been set.

### GetPool6Size

`func (o *IpamAddress6DataData) GetPool6Size() string`

GetPool6Size returns the Pool6Size field if non-nil, zero value otherwise.

### GetPool6SizeOk

`func (o *IpamAddress6DataData) GetPool6SizeOk() (*string, bool)`

GetPool6SizeOk returns a tuple with the Pool6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Size

`func (o *IpamAddress6DataData) SetPool6Size(v string)`

SetPool6Size sets Pool6Size field to given value.

### HasPool6Size

`func (o *IpamAddress6DataData) HasPool6Size() bool`

HasPool6Size returns a boolean if a field has been set.

### GetPool6StartAddress6Addr

`func (o *IpamAddress6DataData) GetPool6StartAddress6Addr() string`

GetPool6StartAddress6Addr returns the Pool6StartAddress6Addr field if non-nil, zero value otherwise.

### GetPool6StartAddress6AddrOk

`func (o *IpamAddress6DataData) GetPool6StartAddress6AddrOk() (*string, bool)`

GetPool6StartAddress6AddrOk returns a tuple with the Pool6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6StartAddress6Addr

`func (o *IpamAddress6DataData) SetPool6StartAddress6Addr(v string)`

SetPool6StartAddress6Addr sets Pool6StartAddress6Addr field to given value.

### HasPool6StartAddress6Addr

`func (o *IpamAddress6DataData) HasPool6StartAddress6Addr() bool`

HasPool6StartAddress6Addr returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *IpamAddress6DataData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *IpamAddress6DataData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *IpamAddress6DataData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *IpamAddress6DataData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *IpamAddress6DataData) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *IpamAddress6DataData) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *IpamAddress6DataData) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *IpamAddress6DataData) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *IpamAddress6DataData) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *IpamAddress6DataData) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *IpamAddress6DataData) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *IpamAddress6DataData) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAddress6DataData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAddress6DataData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAddress6DataData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAddress6DataData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAddress6DataData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAddress6DataData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAddress6DataData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAddress6DataData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetNetwork6ClassName

`func (o *IpamAddress6DataData) GetNetwork6ClassName() string`

GetNetwork6ClassName returns the Network6ClassName field if non-nil, zero value otherwise.

### GetNetwork6ClassNameOk

`func (o *IpamAddress6DataData) GetNetwork6ClassNameOk() (*string, bool)`

GetNetwork6ClassNameOk returns a tuple with the Network6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6ClassName

`func (o *IpamAddress6DataData) SetNetwork6ClassName(v string)`

SetNetwork6ClassName sets Network6ClassName field to given value.

### HasNetwork6ClassName

`func (o *IpamAddress6DataData) HasNetwork6ClassName() bool`

HasNetwork6ClassName returns a boolean if a field has been set.

### GetNetwork6ClassParameters

`func (o *IpamAddress6DataData) GetNetwork6ClassParameters() []ApiClassParameterOutputEntry`

GetNetwork6ClassParameters returns the Network6ClassParameters field if non-nil, zero value otherwise.

### GetNetwork6ClassParametersOk

`func (o *IpamAddress6DataData) GetNetwork6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetNetwork6ClassParametersOk returns a tuple with the Network6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6ClassParameters

`func (o *IpamAddress6DataData) SetNetwork6ClassParameters(v []ApiClassParameterOutputEntry)`

SetNetwork6ClassParameters sets Network6ClassParameters field to given value.

### HasNetwork6ClassParameters

`func (o *IpamAddress6DataData) HasNetwork6ClassParameters() bool`

HasNetwork6ClassParameters returns a boolean if a field has been set.

### GetNetwork6EndAddress6Addr

`func (o *IpamAddress6DataData) GetNetwork6EndAddress6Addr() string`

GetNetwork6EndAddress6Addr returns the Network6EndAddress6Addr field if non-nil, zero value otherwise.

### GetNetwork6EndAddress6AddrOk

`func (o *IpamAddress6DataData) GetNetwork6EndAddress6AddrOk() (*string, bool)`

GetNetwork6EndAddress6AddrOk returns a tuple with the Network6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6EndAddress6Addr

`func (o *IpamAddress6DataData) SetNetwork6EndAddress6Addr(v string)`

SetNetwork6EndAddress6Addr sets Network6EndAddress6Addr field to given value.

### HasNetwork6EndAddress6Addr

`func (o *IpamAddress6DataData) HasNetwork6EndAddress6Addr() bool`

HasNetwork6EndAddress6Addr returns a boolean if a field has been set.

### GetNetwork6Id

`func (o *IpamAddress6DataData) GetNetwork6Id() string`

GetNetwork6Id returns the Network6Id field if non-nil, zero value otherwise.

### GetNetwork6IdOk

`func (o *IpamAddress6DataData) GetNetwork6IdOk() (*string, bool)`

GetNetwork6IdOk returns a tuple with the Network6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Id

`func (o *IpamAddress6DataData) SetNetwork6Id(v string)`

SetNetwork6Id sets Network6Id field to given value.

### HasNetwork6Id

`func (o *IpamAddress6DataData) HasNetwork6Id() bool`

HasNetwork6Id returns a boolean if a field has been set.

### GetNetwork6IsTerminal

`func (o *IpamAddress6DataData) GetNetwork6IsTerminal() string`

GetNetwork6IsTerminal returns the Network6IsTerminal field if non-nil, zero value otherwise.

### GetNetwork6IsTerminalOk

`func (o *IpamAddress6DataData) GetNetwork6IsTerminalOk() (*string, bool)`

GetNetwork6IsTerminalOk returns a tuple with the Network6IsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6IsTerminal

`func (o *IpamAddress6DataData) SetNetwork6IsTerminal(v string)`

SetNetwork6IsTerminal sets Network6IsTerminal field to given value.

### HasNetwork6IsTerminal

`func (o *IpamAddress6DataData) HasNetwork6IsTerminal() bool`

HasNetwork6IsTerminal returns a boolean if a field has been set.

### GetNetwork6Name

`func (o *IpamAddress6DataData) GetNetwork6Name() string`

GetNetwork6Name returns the Network6Name field if non-nil, zero value otherwise.

### GetNetwork6NameOk

`func (o *IpamAddress6DataData) GetNetwork6NameOk() (*string, bool)`

GetNetwork6NameOk returns a tuple with the Network6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Name

`func (o *IpamAddress6DataData) SetNetwork6Name(v string)`

SetNetwork6Name sets Network6Name field to given value.

### HasNetwork6Name

`func (o *IpamAddress6DataData) HasNetwork6Name() bool`

HasNetwork6Name returns a boolean if a field has been set.

### GetNetwork6Prefix

`func (o *IpamAddress6DataData) GetNetwork6Prefix() string`

GetNetwork6Prefix returns the Network6Prefix field if non-nil, zero value otherwise.

### GetNetwork6PrefixOk

`func (o *IpamAddress6DataData) GetNetwork6PrefixOk() (*string, bool)`

GetNetwork6PrefixOk returns a tuple with the Network6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Prefix

`func (o *IpamAddress6DataData) SetNetwork6Prefix(v string)`

SetNetwork6Prefix sets Network6Prefix field to given value.

### HasNetwork6Prefix

`func (o *IpamAddress6DataData) HasNetwork6Prefix() bool`

HasNetwork6Prefix returns a boolean if a field has been set.

### GetNetwork6Size

`func (o *IpamAddress6DataData) GetNetwork6Size() string`

GetNetwork6Size returns the Network6Size field if non-nil, zero value otherwise.

### GetNetwork6SizeOk

`func (o *IpamAddress6DataData) GetNetwork6SizeOk() (*string, bool)`

GetNetwork6SizeOk returns a tuple with the Network6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Size

`func (o *IpamAddress6DataData) SetNetwork6Size(v string)`

SetNetwork6Size sets Network6Size field to given value.

### HasNetwork6Size

`func (o *IpamAddress6DataData) HasNetwork6Size() bool`

HasNetwork6Size returns a boolean if a field has been set.

### GetNetwork6StartAddress6Addr

`func (o *IpamAddress6DataData) GetNetwork6StartAddress6Addr() string`

GetNetwork6StartAddress6Addr returns the Network6StartAddress6Addr field if non-nil, zero value otherwise.

### GetNetwork6StartAddress6AddrOk

`func (o *IpamAddress6DataData) GetNetwork6StartAddress6AddrOk() (*string, bool)`

GetNetwork6StartAddress6AddrOk returns a tuple with the Network6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6StartAddress6Addr

`func (o *IpamAddress6DataData) SetNetwork6StartAddress6Addr(v string)`

SetNetwork6StartAddress6Addr sets Network6StartAddress6Addr field to given value.

### HasNetwork6StartAddress6Addr

`func (o *IpamAddress6DataData) HasNetwork6StartAddress6Addr() bool`

HasNetwork6StartAddress6Addr returns a boolean if a field has been set.

### GetNetworkSize

`func (o *IpamAddress6DataData) GetNetworkSize() string`

GetNetworkSize returns the NetworkSize field if non-nil, zero value otherwise.

### GetNetworkSizeOk

`func (o *IpamAddress6DataData) GetNetworkSizeOk() (*string, bool)`

GetNetworkSizeOk returns a tuple with the NetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSize

`func (o *IpamAddress6DataData) SetNetworkSize(v string)`

SetNetworkSize sets NetworkSize field to given value.

### HasNetworkSize

`func (o *IpamAddress6DataData) HasNetworkSize() bool`

HasNetworkSize returns a boolean if a field has been set.

### GetAddress6Type

`func (o *IpamAddress6DataData) GetAddress6Type() string`

GetAddress6Type returns the Address6Type field if non-nil, zero value otherwise.

### GetAddress6TypeOk

`func (o *IpamAddress6DataData) GetAddress6TypeOk() (*string, bool)`

GetAddress6TypeOk returns a tuple with the Address6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Type

`func (o *IpamAddress6DataData) SetAddress6Type(v string)`

SetAddress6Type sets Address6Type field to given value.

### HasAddress6Type

`func (o *IpamAddress6DataData) HasAddress6Type() bool`

HasAddress6Type returns a boolean if a field has been set.

### GetVlsmNetwork6Id

`func (o *IpamAddress6DataData) GetVlsmNetwork6Id() string`

GetVlsmNetwork6Id returns the VlsmNetwork6Id field if non-nil, zero value otherwise.

### GetVlsmNetwork6IdOk

`func (o *IpamAddress6DataData) GetVlsmNetwork6IdOk() (*string, bool)`

GetVlsmNetwork6IdOk returns a tuple with the VlsmNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmNetwork6Id

`func (o *IpamAddress6DataData) SetVlsmNetwork6Id(v string)`

SetVlsmNetwork6Id sets VlsmNetwork6Id field to given value.

### HasVlsmNetwork6Id

`func (o *IpamAddress6DataData) HasVlsmNetwork6Id() bool`

HasVlsmNetwork6Id returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


