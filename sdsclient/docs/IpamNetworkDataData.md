# IpamNetworkDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkEndHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;end_ip_addr&lt;/b&gt;. | [optional] 
**NetworkEndAddressAddr** | Pointer to **string** | todo(here) :ipam.network.list.output.network_end_address_addr. :  | [optional] 
**NetworkIsInOrphan** | Pointer to **string** | A way to determine if the network has a parent (&lt;b&gt;0&lt;/b&gt;) or if it belongs to a container &lt;b&gt;Orphan networks&lt;/b&gt; (&lt;b&gt;1&lt;/b&gt;). | [optional] 
**NetworkIsTerminal** | Pointer to **string** | A way to determine if a network can contain other networks. If set to &lt;b&gt;1&lt;/b&gt;, the network is terminal and cannot contain other subnet-type networks. Block-type networks are always set to &lt;b&gt;0&lt;/b&gt;. | [optional] 
**NetworkLockNetworkBroadcast** | Pointer to **string** | A way to prevent &lt;b&gt;(1)&lt;/b&gt; users from assigning the broadcast IP address and network IP address of the network. | [optional] 
**NetworkMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentNetworkEndAddressAddr** | Pointer to **string** | todo(here) :ipam.network.list.output.parent_network_end_address_addr. :  | [optional] 
**ParentNetworkIsTerminal** | Pointer to **string** | A way to determine if the parent network is terminal (&lt;b&gt;1&lt;/b&gt;) or non-terminal (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**ParentSpaceId** | Pointer to **string** | The database identifier (ID) of the space where is located the parent network. &lt;b&gt;0&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentSpaceName** | Pointer to **string** | The name of the space where is located the parent network. &lt;b&gt;#&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentNetworkStartAddressAddr** | Pointer to **string** | todo(here) :ipam.network.list.output.parent_network_start_address_addr. :  | [optional] 
**ParentNetworkClassName** | Pointer to **string** | The name of the class applied to the parent IPv4 network, it can be preceded by the class directory. | [optional] 
**ParentNetworkClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the parent IPv4 network and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;...&lt;/b&gt; . | [optional] 
**ParentNetworkId** | Pointer to **string** | The database identifier (ID) of the parent IPv4 network. &lt;b&gt;0&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentNetworkLevel** | Pointer to **string** | The level of the parent network within the space. It returns values between &lt;b&gt;0&lt;/b&gt; (block-type network) and &lt;b&gt;n&lt;/b&gt; (subnet-type network). A value higher than &lt;b&gt;1&lt;/b&gt; indicates a VLSM organization where a block-type network can belong to another subnet-type network. | [optional] 
**ParentNetworkName** | Pointer to **string** | The name of the parent IPv4 network:&lt;ul class&#x3D;dashed &gt;&lt;li&gt;&lt;b&gt;#&lt;/b&gt; indicates that the network has no parent network.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;&lt;b&gt;Default&lt;/b&gt; indicates that the network belongs to an orphan network.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt; | [optional] 
**ParentNetworkPath** | Pointer to **string** | The path toward the parent network in the database. &lt;b&gt;#&lt;/b&gt; indicates the network has no parent network. | [optional] 
**ParentNetworkSize** | Pointer to **string** | The number of IP addresses of the network parent. | [optional] 
**ParentVlsmNetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 subnet-type network, located in the VLSM parent space, from which the parent network was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the parent network is not a VLSM block-type network duplicated from a parent space. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space the object belongs to. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceIsTemplate** | Pointer to **string** | The template status of the space the object belongs to. If the space is used as template (1), all the IPv4 networks, pools and IP addresses it contains are also used as template. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**SpaceParentSpaceId** | Pointer to **string** | The database identifier (ID) of the VLSM parent of the space where is located the network. &lt;b&gt;0&lt;/b&gt; indicates that the space where is located the network has no parent space. | [optional] 
**NetworkStartHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;start_ip_addr&lt;/b&gt;. | [optional] 
**NetworkStartAddressAddr** | Pointer to **string** | todo(here) :ipam.network.list.output.network_start_address_addr. :  | [optional] 
**NetworkAllocatedPercent** | Pointer to **string** | The percentage of subnet-type networks the non-terminal network contains. | [optional] 
**NetworkAllocatedSize** | Pointer to **string** | The sum of the size of all the subnet-type networks that belong to the block-type network. | [optional] 
**NetworkClassName** | Pointer to **string** | The name of the class applied to the IPv4 network, it can be preceded by the class directory. | [optional] 
**NetworkClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**NetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 network. | [optional] 
**NetworkAddressFreeSize** | Pointer to **string** | The total number of free addresses, for terminal networks only. It excludes the network and broadcast IP address. | [optional] 
**NetworkAddressUsedPercent** | Pointer to **string** | The percentage of IP addresses &lt;b&gt;In use&lt;/b&gt; in terminal networks. | [optional] 
**NetworkAddressUsedSize** | Pointer to **string** | The number of IP addresses &lt;b&gt;In use&lt;/b&gt; in terminal networks. | [optional] 
**NetworkIsValid** | Pointer to **string** | The network validity. A valid network (&lt;b&gt;1&lt;/b&gt;) has a size, prefix and/or netmask that match. | [optional] 
**NetworkLevel** | Pointer to **string** | The level of the network within the space. It returns values between &lt;b&gt;0&lt;/b&gt; (block-type network) and &lt;b&gt;n&lt;/b&gt; (subnet-type network). A value higher than &lt;b&gt;1&lt;/b&gt; indicates a VLSM organization where a block-type network can belong to another subnet-type network. | [optional] 
**NetworkName** | Pointer to **string** | The name of the IPv4 network. &lt;b&gt;Default&lt;/b&gt; indicates that the network is an orphan network. | [optional] 
**NetworkSize** | Pointer to **string** | The number of IP addresses the IPv4 network contains. | [optional] 
**NetworkUsedPercent** | Pointer to **string** | The percentage of terminal networks the non-terminal network contains. | [optional] 
**NetworkUsedSize** | Pointer to **string** | The sum of the size of all the terminal networks within the block-type network. This sum includes the terminal networks that might belong to non-terminal subnet-type networks. | [optional] 
**DomainId** | Pointer to **string** | The database identifier (ID) of the VLAN domain associated with the network. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain associated with the network. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the VLAN range associated with the network. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range associated with the network. | [optional] 
**VlanId** | Pointer to **string** | The database identifier (ID) of the VLAN associated with the network. | [optional] 
**VlanName** | Pointer to **string** | The name of the VLAN associated with the network. | [optional] 
**VlanVlanId** | Pointer to **string** | The VLAN identifier (ID) of the VLAN associated with the network. | [optional] 
**VlsmBlockId** | Pointer to **string** | The database identifier (ID) of the IPv4 VLSM block-type network duplicated, in a VLSM child space, from the network. &lt;b&gt;0&lt;/b&gt; indicates that the network is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmSpaceId** | Pointer to **string** | The database identifier (ID) of the VLSM child space where the network is duplicated as a VLSM block-type network. &lt;b&gt;0&lt;/b&gt; indicates that the network is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmSpaceName** | Pointer to **string** | The name of the VLSM child space where the network is duplicated as a VLSM block-type network. &lt;b&gt;0&lt;/b&gt; indicates that the network is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmNetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 subnet-type network, located in the VLSM parent space, from which the network was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the network is not a VLSM block-type network duplicated from a parent space. | [optional] 
**NetworkRipeWaitingState** | Pointer to **string** | The state of the exchange between SOLIDserver and the RIPE for the assigned network: &lt;table&gt;&lt;caption&gt;waiting_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;must_send_mail_add&lt;/td&gt;&lt;td &gt;An email must be sent to the RIPE to notify them of a subnet-type network creation.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_mail_add&lt;/td&gt;&lt;td &gt;A network creation email was sent to the RIPE, no reply has been received yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;must_send_mail_del&lt;/td&gt;&lt;td &gt;An email must be sent to the RIPE to notify them of a subnet-type network deletion.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_mail_del&lt;/td&gt;&lt;td &gt;A network deletion email was sent to the RIPE, no reply has been received yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_aw_confirm&lt;/td&gt;&lt;td &gt;The number of IP addresses of the assigned network exceeds the Assignment Window declared during your RIPE configuration.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**NetworkRipeWaitingStatus** | Pointer to **string** | The status of a RIPE assigned network within SOLIDserver until it is confirmed that you can create or delete it. If set to &lt;b&gt;1&lt;/b&gt;, it is about to be created. If set to &lt;b&gt;2&lt;/b&gt;, it is about to be deleted. | [optional] 

## Methods

### NewIpamNetworkDataData

`func NewIpamNetworkDataData() *IpamNetworkDataData`

NewIpamNetworkDataData instantiates a new IpamNetworkDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamNetworkDataDataWithDefaults

`func NewIpamNetworkDataDataWithDefaults() *IpamNetworkDataData`

NewIpamNetworkDataDataWithDefaults instantiates a new IpamNetworkDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkEndHostaddr

`func (o *IpamNetworkDataData) GetNetworkEndHostaddr() string`

GetNetworkEndHostaddr returns the NetworkEndHostaddr field if non-nil, zero value otherwise.

### GetNetworkEndHostaddrOk

`func (o *IpamNetworkDataData) GetNetworkEndHostaddrOk() (*string, bool)`

GetNetworkEndHostaddrOk returns a tuple with the NetworkEndHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndHostaddr

`func (o *IpamNetworkDataData) SetNetworkEndHostaddr(v string)`

SetNetworkEndHostaddr sets NetworkEndHostaddr field to given value.

### HasNetworkEndHostaddr

`func (o *IpamNetworkDataData) HasNetworkEndHostaddr() bool`

HasNetworkEndHostaddr returns a boolean if a field has been set.

### GetNetworkEndAddressAddr

`func (o *IpamNetworkDataData) GetNetworkEndAddressAddr() string`

GetNetworkEndAddressAddr returns the NetworkEndAddressAddr field if non-nil, zero value otherwise.

### GetNetworkEndAddressAddrOk

`func (o *IpamNetworkDataData) GetNetworkEndAddressAddrOk() (*string, bool)`

GetNetworkEndAddressAddrOk returns a tuple with the NetworkEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndAddressAddr

`func (o *IpamNetworkDataData) SetNetworkEndAddressAddr(v string)`

SetNetworkEndAddressAddr sets NetworkEndAddressAddr field to given value.

### HasNetworkEndAddressAddr

`func (o *IpamNetworkDataData) HasNetworkEndAddressAddr() bool`

HasNetworkEndAddressAddr returns a boolean if a field has been set.

### GetNetworkIsInOrphan

`func (o *IpamNetworkDataData) GetNetworkIsInOrphan() string`

GetNetworkIsInOrphan returns the NetworkIsInOrphan field if non-nil, zero value otherwise.

### GetNetworkIsInOrphanOk

`func (o *IpamNetworkDataData) GetNetworkIsInOrphanOk() (*string, bool)`

GetNetworkIsInOrphanOk returns a tuple with the NetworkIsInOrphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsInOrphan

`func (o *IpamNetworkDataData) SetNetworkIsInOrphan(v string)`

SetNetworkIsInOrphan sets NetworkIsInOrphan field to given value.

### HasNetworkIsInOrphan

`func (o *IpamNetworkDataData) HasNetworkIsInOrphan() bool`

HasNetworkIsInOrphan returns a boolean if a field has been set.

### GetNetworkIsTerminal

`func (o *IpamNetworkDataData) GetNetworkIsTerminal() string`

GetNetworkIsTerminal returns the NetworkIsTerminal field if non-nil, zero value otherwise.

### GetNetworkIsTerminalOk

`func (o *IpamNetworkDataData) GetNetworkIsTerminalOk() (*string, bool)`

GetNetworkIsTerminalOk returns a tuple with the NetworkIsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsTerminal

`func (o *IpamNetworkDataData) SetNetworkIsTerminal(v string)`

SetNetworkIsTerminal sets NetworkIsTerminal field to given value.

### HasNetworkIsTerminal

`func (o *IpamNetworkDataData) HasNetworkIsTerminal() bool`

HasNetworkIsTerminal returns a boolean if a field has been set.

### GetNetworkLockNetworkBroadcast

`func (o *IpamNetworkDataData) GetNetworkLockNetworkBroadcast() string`

GetNetworkLockNetworkBroadcast returns the NetworkLockNetworkBroadcast field if non-nil, zero value otherwise.

### GetNetworkLockNetworkBroadcastOk

`func (o *IpamNetworkDataData) GetNetworkLockNetworkBroadcastOk() (*string, bool)`

GetNetworkLockNetworkBroadcastOk returns a tuple with the NetworkLockNetworkBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLockNetworkBroadcast

`func (o *IpamNetworkDataData) SetNetworkLockNetworkBroadcast(v string)`

SetNetworkLockNetworkBroadcast sets NetworkLockNetworkBroadcast field to given value.

### HasNetworkLockNetworkBroadcast

`func (o *IpamNetworkDataData) HasNetworkLockNetworkBroadcast() bool`

HasNetworkLockNetworkBroadcast returns a boolean if a field has been set.

### GetNetworkMultistatus

`func (o *IpamNetworkDataData) GetNetworkMultistatus() string`

GetNetworkMultistatus returns the NetworkMultistatus field if non-nil, zero value otherwise.

### GetNetworkMultistatusOk

`func (o *IpamNetworkDataData) GetNetworkMultistatusOk() (*string, bool)`

GetNetworkMultistatusOk returns a tuple with the NetworkMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMultistatus

`func (o *IpamNetworkDataData) SetNetworkMultistatus(v string)`

SetNetworkMultistatus sets NetworkMultistatus field to given value.

### HasNetworkMultistatus

`func (o *IpamNetworkDataData) HasNetworkMultistatus() bool`

HasNetworkMultistatus returns a boolean if a field has been set.

### GetParentNetworkEndAddressAddr

`func (o *IpamNetworkDataData) GetParentNetworkEndAddressAddr() string`

GetParentNetworkEndAddressAddr returns the ParentNetworkEndAddressAddr field if non-nil, zero value otherwise.

### GetParentNetworkEndAddressAddrOk

`func (o *IpamNetworkDataData) GetParentNetworkEndAddressAddrOk() (*string, bool)`

GetParentNetworkEndAddressAddrOk returns a tuple with the ParentNetworkEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkEndAddressAddr

`func (o *IpamNetworkDataData) SetParentNetworkEndAddressAddr(v string)`

SetParentNetworkEndAddressAddr sets ParentNetworkEndAddressAddr field to given value.

### HasParentNetworkEndAddressAddr

`func (o *IpamNetworkDataData) HasParentNetworkEndAddressAddr() bool`

HasParentNetworkEndAddressAddr returns a boolean if a field has been set.

### GetParentNetworkIsTerminal

`func (o *IpamNetworkDataData) GetParentNetworkIsTerminal() string`

GetParentNetworkIsTerminal returns the ParentNetworkIsTerminal field if non-nil, zero value otherwise.

### GetParentNetworkIsTerminalOk

`func (o *IpamNetworkDataData) GetParentNetworkIsTerminalOk() (*string, bool)`

GetParentNetworkIsTerminalOk returns a tuple with the ParentNetworkIsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkIsTerminal

`func (o *IpamNetworkDataData) SetParentNetworkIsTerminal(v string)`

SetParentNetworkIsTerminal sets ParentNetworkIsTerminal field to given value.

### HasParentNetworkIsTerminal

`func (o *IpamNetworkDataData) HasParentNetworkIsTerminal() bool`

HasParentNetworkIsTerminal returns a boolean if a field has been set.

### GetParentSpaceId

`func (o *IpamNetworkDataData) GetParentSpaceId() string`

GetParentSpaceId returns the ParentSpaceId field if non-nil, zero value otherwise.

### GetParentSpaceIdOk

`func (o *IpamNetworkDataData) GetParentSpaceIdOk() (*string, bool)`

GetParentSpaceIdOk returns a tuple with the ParentSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceId

`func (o *IpamNetworkDataData) SetParentSpaceId(v string)`

SetParentSpaceId sets ParentSpaceId field to given value.

### HasParentSpaceId

`func (o *IpamNetworkDataData) HasParentSpaceId() bool`

HasParentSpaceId returns a boolean if a field has been set.

### GetParentSpaceName

`func (o *IpamNetworkDataData) GetParentSpaceName() string`

GetParentSpaceName returns the ParentSpaceName field if non-nil, zero value otherwise.

### GetParentSpaceNameOk

`func (o *IpamNetworkDataData) GetParentSpaceNameOk() (*string, bool)`

GetParentSpaceNameOk returns a tuple with the ParentSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceName

`func (o *IpamNetworkDataData) SetParentSpaceName(v string)`

SetParentSpaceName sets ParentSpaceName field to given value.

### HasParentSpaceName

`func (o *IpamNetworkDataData) HasParentSpaceName() bool`

HasParentSpaceName returns a boolean if a field has been set.

### GetParentNetworkStartAddressAddr

`func (o *IpamNetworkDataData) GetParentNetworkStartAddressAddr() string`

GetParentNetworkStartAddressAddr returns the ParentNetworkStartAddressAddr field if non-nil, zero value otherwise.

### GetParentNetworkStartAddressAddrOk

`func (o *IpamNetworkDataData) GetParentNetworkStartAddressAddrOk() (*string, bool)`

GetParentNetworkStartAddressAddrOk returns a tuple with the ParentNetworkStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkStartAddressAddr

`func (o *IpamNetworkDataData) SetParentNetworkStartAddressAddr(v string)`

SetParentNetworkStartAddressAddr sets ParentNetworkStartAddressAddr field to given value.

### HasParentNetworkStartAddressAddr

`func (o *IpamNetworkDataData) HasParentNetworkStartAddressAddr() bool`

HasParentNetworkStartAddressAddr returns a boolean if a field has been set.

### GetParentNetworkClassName

`func (o *IpamNetworkDataData) GetParentNetworkClassName() string`

GetParentNetworkClassName returns the ParentNetworkClassName field if non-nil, zero value otherwise.

### GetParentNetworkClassNameOk

`func (o *IpamNetworkDataData) GetParentNetworkClassNameOk() (*string, bool)`

GetParentNetworkClassNameOk returns a tuple with the ParentNetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkClassName

`func (o *IpamNetworkDataData) SetParentNetworkClassName(v string)`

SetParentNetworkClassName sets ParentNetworkClassName field to given value.

### HasParentNetworkClassName

`func (o *IpamNetworkDataData) HasParentNetworkClassName() bool`

HasParentNetworkClassName returns a boolean if a field has been set.

### GetParentNetworkClassParameters

`func (o *IpamNetworkDataData) GetParentNetworkClassParameters() []ApiClassParameterOutputEntry`

GetParentNetworkClassParameters returns the ParentNetworkClassParameters field if non-nil, zero value otherwise.

### GetParentNetworkClassParametersOk

`func (o *IpamNetworkDataData) GetParentNetworkClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetParentNetworkClassParametersOk returns a tuple with the ParentNetworkClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkClassParameters

`func (o *IpamNetworkDataData) SetParentNetworkClassParameters(v []ApiClassParameterOutputEntry)`

SetParentNetworkClassParameters sets ParentNetworkClassParameters field to given value.

### HasParentNetworkClassParameters

`func (o *IpamNetworkDataData) HasParentNetworkClassParameters() bool`

HasParentNetworkClassParameters returns a boolean if a field has been set.

### GetParentNetworkId

`func (o *IpamNetworkDataData) GetParentNetworkId() string`

GetParentNetworkId returns the ParentNetworkId field if non-nil, zero value otherwise.

### GetParentNetworkIdOk

`func (o *IpamNetworkDataData) GetParentNetworkIdOk() (*string, bool)`

GetParentNetworkIdOk returns a tuple with the ParentNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkId

`func (o *IpamNetworkDataData) SetParentNetworkId(v string)`

SetParentNetworkId sets ParentNetworkId field to given value.

### HasParentNetworkId

`func (o *IpamNetworkDataData) HasParentNetworkId() bool`

HasParentNetworkId returns a boolean if a field has been set.

### GetParentNetworkLevel

`func (o *IpamNetworkDataData) GetParentNetworkLevel() string`

GetParentNetworkLevel returns the ParentNetworkLevel field if non-nil, zero value otherwise.

### GetParentNetworkLevelOk

`func (o *IpamNetworkDataData) GetParentNetworkLevelOk() (*string, bool)`

GetParentNetworkLevelOk returns a tuple with the ParentNetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkLevel

`func (o *IpamNetworkDataData) SetParentNetworkLevel(v string)`

SetParentNetworkLevel sets ParentNetworkLevel field to given value.

### HasParentNetworkLevel

`func (o *IpamNetworkDataData) HasParentNetworkLevel() bool`

HasParentNetworkLevel returns a boolean if a field has been set.

### GetParentNetworkName

`func (o *IpamNetworkDataData) GetParentNetworkName() string`

GetParentNetworkName returns the ParentNetworkName field if non-nil, zero value otherwise.

### GetParentNetworkNameOk

`func (o *IpamNetworkDataData) GetParentNetworkNameOk() (*string, bool)`

GetParentNetworkNameOk returns a tuple with the ParentNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkName

`func (o *IpamNetworkDataData) SetParentNetworkName(v string)`

SetParentNetworkName sets ParentNetworkName field to given value.

### HasParentNetworkName

`func (o *IpamNetworkDataData) HasParentNetworkName() bool`

HasParentNetworkName returns a boolean if a field has been set.

### GetParentNetworkPath

`func (o *IpamNetworkDataData) GetParentNetworkPath() string`

GetParentNetworkPath returns the ParentNetworkPath field if non-nil, zero value otherwise.

### GetParentNetworkPathOk

`func (o *IpamNetworkDataData) GetParentNetworkPathOk() (*string, bool)`

GetParentNetworkPathOk returns a tuple with the ParentNetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkPath

`func (o *IpamNetworkDataData) SetParentNetworkPath(v string)`

SetParentNetworkPath sets ParentNetworkPath field to given value.

### HasParentNetworkPath

`func (o *IpamNetworkDataData) HasParentNetworkPath() bool`

HasParentNetworkPath returns a boolean if a field has been set.

### GetParentNetworkSize

`func (o *IpamNetworkDataData) GetParentNetworkSize() string`

GetParentNetworkSize returns the ParentNetworkSize field if non-nil, zero value otherwise.

### GetParentNetworkSizeOk

`func (o *IpamNetworkDataData) GetParentNetworkSizeOk() (*string, bool)`

GetParentNetworkSizeOk returns a tuple with the ParentNetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkSize

`func (o *IpamNetworkDataData) SetParentNetworkSize(v string)`

SetParentNetworkSize sets ParentNetworkSize field to given value.

### HasParentNetworkSize

`func (o *IpamNetworkDataData) HasParentNetworkSize() bool`

HasParentNetworkSize returns a boolean if a field has been set.

### GetParentVlsmNetworkId

`func (o *IpamNetworkDataData) GetParentVlsmNetworkId() string`

GetParentVlsmNetworkId returns the ParentVlsmNetworkId field if non-nil, zero value otherwise.

### GetParentVlsmNetworkIdOk

`func (o *IpamNetworkDataData) GetParentVlsmNetworkIdOk() (*string, bool)`

GetParentVlsmNetworkIdOk returns a tuple with the ParentVlsmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVlsmNetworkId

`func (o *IpamNetworkDataData) SetParentVlsmNetworkId(v string)`

SetParentVlsmNetworkId sets ParentVlsmNetworkId field to given value.

### HasParentVlsmNetworkId

`func (o *IpamNetworkDataData) HasParentVlsmNetworkId() bool`

HasParentVlsmNetworkId returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *IpamNetworkDataData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *IpamNetworkDataData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *IpamNetworkDataData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *IpamNetworkDataData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *IpamNetworkDataData) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *IpamNetworkDataData) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *IpamNetworkDataData) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *IpamNetworkDataData) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *IpamNetworkDataData) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *IpamNetworkDataData) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *IpamNetworkDataData) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *IpamNetworkDataData) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamNetworkDataData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamNetworkDataData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamNetworkDataData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamNetworkDataData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceIsTemplate

`func (o *IpamNetworkDataData) GetSpaceIsTemplate() string`

GetSpaceIsTemplate returns the SpaceIsTemplate field if non-nil, zero value otherwise.

### GetSpaceIsTemplateOk

`func (o *IpamNetworkDataData) GetSpaceIsTemplateOk() (*string, bool)`

GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceIsTemplate

`func (o *IpamNetworkDataData) SetSpaceIsTemplate(v string)`

SetSpaceIsTemplate sets SpaceIsTemplate field to given value.

### HasSpaceIsTemplate

`func (o *IpamNetworkDataData) HasSpaceIsTemplate() bool`

HasSpaceIsTemplate returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamNetworkDataData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamNetworkDataData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamNetworkDataData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamNetworkDataData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetSpaceParentSpaceId

`func (o *IpamNetworkDataData) GetSpaceParentSpaceId() string`

GetSpaceParentSpaceId returns the SpaceParentSpaceId field if non-nil, zero value otherwise.

### GetSpaceParentSpaceIdOk

`func (o *IpamNetworkDataData) GetSpaceParentSpaceIdOk() (*string, bool)`

GetSpaceParentSpaceIdOk returns a tuple with the SpaceParentSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceParentSpaceId

`func (o *IpamNetworkDataData) SetSpaceParentSpaceId(v string)`

SetSpaceParentSpaceId sets SpaceParentSpaceId field to given value.

### HasSpaceParentSpaceId

`func (o *IpamNetworkDataData) HasSpaceParentSpaceId() bool`

HasSpaceParentSpaceId returns a boolean if a field has been set.

### GetNetworkStartHostaddr

`func (o *IpamNetworkDataData) GetNetworkStartHostaddr() string`

GetNetworkStartHostaddr returns the NetworkStartHostaddr field if non-nil, zero value otherwise.

### GetNetworkStartHostaddrOk

`func (o *IpamNetworkDataData) GetNetworkStartHostaddrOk() (*string, bool)`

GetNetworkStartHostaddrOk returns a tuple with the NetworkStartHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStartHostaddr

`func (o *IpamNetworkDataData) SetNetworkStartHostaddr(v string)`

SetNetworkStartHostaddr sets NetworkStartHostaddr field to given value.

### HasNetworkStartHostaddr

`func (o *IpamNetworkDataData) HasNetworkStartHostaddr() bool`

HasNetworkStartHostaddr returns a boolean if a field has been set.

### GetNetworkStartAddressAddr

`func (o *IpamNetworkDataData) GetNetworkStartAddressAddr() string`

GetNetworkStartAddressAddr returns the NetworkStartAddressAddr field if non-nil, zero value otherwise.

### GetNetworkStartAddressAddrOk

`func (o *IpamNetworkDataData) GetNetworkStartAddressAddrOk() (*string, bool)`

GetNetworkStartAddressAddrOk returns a tuple with the NetworkStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStartAddressAddr

`func (o *IpamNetworkDataData) SetNetworkStartAddressAddr(v string)`

SetNetworkStartAddressAddr sets NetworkStartAddressAddr field to given value.

### HasNetworkStartAddressAddr

`func (o *IpamNetworkDataData) HasNetworkStartAddressAddr() bool`

HasNetworkStartAddressAddr returns a boolean if a field has been set.

### GetNetworkAllocatedPercent

`func (o *IpamNetworkDataData) GetNetworkAllocatedPercent() string`

GetNetworkAllocatedPercent returns the NetworkAllocatedPercent field if non-nil, zero value otherwise.

### GetNetworkAllocatedPercentOk

`func (o *IpamNetworkDataData) GetNetworkAllocatedPercentOk() (*string, bool)`

GetNetworkAllocatedPercentOk returns a tuple with the NetworkAllocatedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAllocatedPercent

`func (o *IpamNetworkDataData) SetNetworkAllocatedPercent(v string)`

SetNetworkAllocatedPercent sets NetworkAllocatedPercent field to given value.

### HasNetworkAllocatedPercent

`func (o *IpamNetworkDataData) HasNetworkAllocatedPercent() bool`

HasNetworkAllocatedPercent returns a boolean if a field has been set.

### GetNetworkAllocatedSize

`func (o *IpamNetworkDataData) GetNetworkAllocatedSize() string`

GetNetworkAllocatedSize returns the NetworkAllocatedSize field if non-nil, zero value otherwise.

### GetNetworkAllocatedSizeOk

`func (o *IpamNetworkDataData) GetNetworkAllocatedSizeOk() (*string, bool)`

GetNetworkAllocatedSizeOk returns a tuple with the NetworkAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAllocatedSize

`func (o *IpamNetworkDataData) SetNetworkAllocatedSize(v string)`

SetNetworkAllocatedSize sets NetworkAllocatedSize field to given value.

### HasNetworkAllocatedSize

`func (o *IpamNetworkDataData) HasNetworkAllocatedSize() bool`

HasNetworkAllocatedSize returns a boolean if a field has been set.

### GetNetworkClassName

`func (o *IpamNetworkDataData) GetNetworkClassName() string`

GetNetworkClassName returns the NetworkClassName field if non-nil, zero value otherwise.

### GetNetworkClassNameOk

`func (o *IpamNetworkDataData) GetNetworkClassNameOk() (*string, bool)`

GetNetworkClassNameOk returns a tuple with the NetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassName

`func (o *IpamNetworkDataData) SetNetworkClassName(v string)`

SetNetworkClassName sets NetworkClassName field to given value.

### HasNetworkClassName

`func (o *IpamNetworkDataData) HasNetworkClassName() bool`

HasNetworkClassName returns a boolean if a field has been set.

### GetNetworkClassParameters

`func (o *IpamNetworkDataData) GetNetworkClassParameters() []ApiClassParameterOutputEntry`

GetNetworkClassParameters returns the NetworkClassParameters field if non-nil, zero value otherwise.

### GetNetworkClassParametersOk

`func (o *IpamNetworkDataData) GetNetworkClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetNetworkClassParametersOk returns a tuple with the NetworkClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassParameters

`func (o *IpamNetworkDataData) SetNetworkClassParameters(v []ApiClassParameterOutputEntry)`

SetNetworkClassParameters sets NetworkClassParameters field to given value.

### HasNetworkClassParameters

`func (o *IpamNetworkDataData) HasNetworkClassParameters() bool`

HasNetworkClassParameters returns a boolean if a field has been set.

### GetNetworkId

`func (o *IpamNetworkDataData) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *IpamNetworkDataData) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *IpamNetworkDataData) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *IpamNetworkDataData) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkAddressFreeSize

`func (o *IpamNetworkDataData) GetNetworkAddressFreeSize() string`

GetNetworkAddressFreeSize returns the NetworkAddressFreeSize field if non-nil, zero value otherwise.

### GetNetworkAddressFreeSizeOk

`func (o *IpamNetworkDataData) GetNetworkAddressFreeSizeOk() (*string, bool)`

GetNetworkAddressFreeSizeOk returns a tuple with the NetworkAddressFreeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressFreeSize

`func (o *IpamNetworkDataData) SetNetworkAddressFreeSize(v string)`

SetNetworkAddressFreeSize sets NetworkAddressFreeSize field to given value.

### HasNetworkAddressFreeSize

`func (o *IpamNetworkDataData) HasNetworkAddressFreeSize() bool`

HasNetworkAddressFreeSize returns a boolean if a field has been set.

### GetNetworkAddressUsedPercent

`func (o *IpamNetworkDataData) GetNetworkAddressUsedPercent() string`

GetNetworkAddressUsedPercent returns the NetworkAddressUsedPercent field if non-nil, zero value otherwise.

### GetNetworkAddressUsedPercentOk

`func (o *IpamNetworkDataData) GetNetworkAddressUsedPercentOk() (*string, bool)`

GetNetworkAddressUsedPercentOk returns a tuple with the NetworkAddressUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressUsedPercent

`func (o *IpamNetworkDataData) SetNetworkAddressUsedPercent(v string)`

SetNetworkAddressUsedPercent sets NetworkAddressUsedPercent field to given value.

### HasNetworkAddressUsedPercent

`func (o *IpamNetworkDataData) HasNetworkAddressUsedPercent() bool`

HasNetworkAddressUsedPercent returns a boolean if a field has been set.

### GetNetworkAddressUsedSize

`func (o *IpamNetworkDataData) GetNetworkAddressUsedSize() string`

GetNetworkAddressUsedSize returns the NetworkAddressUsedSize field if non-nil, zero value otherwise.

### GetNetworkAddressUsedSizeOk

`func (o *IpamNetworkDataData) GetNetworkAddressUsedSizeOk() (*string, bool)`

GetNetworkAddressUsedSizeOk returns a tuple with the NetworkAddressUsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressUsedSize

`func (o *IpamNetworkDataData) SetNetworkAddressUsedSize(v string)`

SetNetworkAddressUsedSize sets NetworkAddressUsedSize field to given value.

### HasNetworkAddressUsedSize

`func (o *IpamNetworkDataData) HasNetworkAddressUsedSize() bool`

HasNetworkAddressUsedSize returns a boolean if a field has been set.

### GetNetworkIsValid

`func (o *IpamNetworkDataData) GetNetworkIsValid() string`

GetNetworkIsValid returns the NetworkIsValid field if non-nil, zero value otherwise.

### GetNetworkIsValidOk

`func (o *IpamNetworkDataData) GetNetworkIsValidOk() (*string, bool)`

GetNetworkIsValidOk returns a tuple with the NetworkIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsValid

`func (o *IpamNetworkDataData) SetNetworkIsValid(v string)`

SetNetworkIsValid sets NetworkIsValid field to given value.

### HasNetworkIsValid

`func (o *IpamNetworkDataData) HasNetworkIsValid() bool`

HasNetworkIsValid returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *IpamNetworkDataData) GetNetworkLevel() string`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *IpamNetworkDataData) GetNetworkLevelOk() (*string, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *IpamNetworkDataData) SetNetworkLevel(v string)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *IpamNetworkDataData) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetNetworkName

`func (o *IpamNetworkDataData) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *IpamNetworkDataData) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *IpamNetworkDataData) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *IpamNetworkDataData) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkSize

`func (o *IpamNetworkDataData) GetNetworkSize() string`

GetNetworkSize returns the NetworkSize field if non-nil, zero value otherwise.

### GetNetworkSizeOk

`func (o *IpamNetworkDataData) GetNetworkSizeOk() (*string, bool)`

GetNetworkSizeOk returns a tuple with the NetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSize

`func (o *IpamNetworkDataData) SetNetworkSize(v string)`

SetNetworkSize sets NetworkSize field to given value.

### HasNetworkSize

`func (o *IpamNetworkDataData) HasNetworkSize() bool`

HasNetworkSize returns a boolean if a field has been set.

### GetNetworkUsedPercent

`func (o *IpamNetworkDataData) GetNetworkUsedPercent() string`

GetNetworkUsedPercent returns the NetworkUsedPercent field if non-nil, zero value otherwise.

### GetNetworkUsedPercentOk

`func (o *IpamNetworkDataData) GetNetworkUsedPercentOk() (*string, bool)`

GetNetworkUsedPercentOk returns a tuple with the NetworkUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUsedPercent

`func (o *IpamNetworkDataData) SetNetworkUsedPercent(v string)`

SetNetworkUsedPercent sets NetworkUsedPercent field to given value.

### HasNetworkUsedPercent

`func (o *IpamNetworkDataData) HasNetworkUsedPercent() bool`

HasNetworkUsedPercent returns a boolean if a field has been set.

### GetNetworkUsedSize

`func (o *IpamNetworkDataData) GetNetworkUsedSize() string`

GetNetworkUsedSize returns the NetworkUsedSize field if non-nil, zero value otherwise.

### GetNetworkUsedSizeOk

`func (o *IpamNetworkDataData) GetNetworkUsedSizeOk() (*string, bool)`

GetNetworkUsedSizeOk returns a tuple with the NetworkUsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUsedSize

`func (o *IpamNetworkDataData) SetNetworkUsedSize(v string)`

SetNetworkUsedSize sets NetworkUsedSize field to given value.

### HasNetworkUsedSize

`func (o *IpamNetworkDataData) HasNetworkUsedSize() bool`

HasNetworkUsedSize returns a boolean if a field has been set.

### GetDomainId

`func (o *IpamNetworkDataData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *IpamNetworkDataData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *IpamNetworkDataData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *IpamNetworkDataData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *IpamNetworkDataData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *IpamNetworkDataData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *IpamNetworkDataData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *IpamNetworkDataData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRangeId

`func (o *IpamNetworkDataData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *IpamNetworkDataData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *IpamNetworkDataData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *IpamNetworkDataData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *IpamNetworkDataData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *IpamNetworkDataData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *IpamNetworkDataData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *IpamNetworkDataData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetVlanId

`func (o *IpamNetworkDataData) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *IpamNetworkDataData) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *IpamNetworkDataData) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *IpamNetworkDataData) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanName

`func (o *IpamNetworkDataData) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *IpamNetworkDataData) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *IpamNetworkDataData) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *IpamNetworkDataData) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetVlanVlanId

`func (o *IpamNetworkDataData) GetVlanVlanId() string`

GetVlanVlanId returns the VlanVlanId field if non-nil, zero value otherwise.

### GetVlanVlanIdOk

`func (o *IpamNetworkDataData) GetVlanVlanIdOk() (*string, bool)`

GetVlanVlanIdOk returns a tuple with the VlanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVlanId

`func (o *IpamNetworkDataData) SetVlanVlanId(v string)`

SetVlanVlanId sets VlanVlanId field to given value.

### HasVlanVlanId

`func (o *IpamNetworkDataData) HasVlanVlanId() bool`

HasVlanVlanId returns a boolean if a field has been set.

### GetVlsmBlockId

`func (o *IpamNetworkDataData) GetVlsmBlockId() string`

GetVlsmBlockId returns the VlsmBlockId field if non-nil, zero value otherwise.

### GetVlsmBlockIdOk

`func (o *IpamNetworkDataData) GetVlsmBlockIdOk() (*string, bool)`

GetVlsmBlockIdOk returns a tuple with the VlsmBlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmBlockId

`func (o *IpamNetworkDataData) SetVlsmBlockId(v string)`

SetVlsmBlockId sets VlsmBlockId field to given value.

### HasVlsmBlockId

`func (o *IpamNetworkDataData) HasVlsmBlockId() bool`

HasVlsmBlockId returns a boolean if a field has been set.

### GetVlsmSpaceId

`func (o *IpamNetworkDataData) GetVlsmSpaceId() string`

GetVlsmSpaceId returns the VlsmSpaceId field if non-nil, zero value otherwise.

### GetVlsmSpaceIdOk

`func (o *IpamNetworkDataData) GetVlsmSpaceIdOk() (*string, bool)`

GetVlsmSpaceIdOk returns a tuple with the VlsmSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceId

`func (o *IpamNetworkDataData) SetVlsmSpaceId(v string)`

SetVlsmSpaceId sets VlsmSpaceId field to given value.

### HasVlsmSpaceId

`func (o *IpamNetworkDataData) HasVlsmSpaceId() bool`

HasVlsmSpaceId returns a boolean if a field has been set.

### GetVlsmSpaceName

`func (o *IpamNetworkDataData) GetVlsmSpaceName() string`

GetVlsmSpaceName returns the VlsmSpaceName field if non-nil, zero value otherwise.

### GetVlsmSpaceNameOk

`func (o *IpamNetworkDataData) GetVlsmSpaceNameOk() (*string, bool)`

GetVlsmSpaceNameOk returns a tuple with the VlsmSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceName

`func (o *IpamNetworkDataData) SetVlsmSpaceName(v string)`

SetVlsmSpaceName sets VlsmSpaceName field to given value.

### HasVlsmSpaceName

`func (o *IpamNetworkDataData) HasVlsmSpaceName() bool`

HasVlsmSpaceName returns a boolean if a field has been set.

### GetVlsmNetworkId

`func (o *IpamNetworkDataData) GetVlsmNetworkId() string`

GetVlsmNetworkId returns the VlsmNetworkId field if non-nil, zero value otherwise.

### GetVlsmNetworkIdOk

`func (o *IpamNetworkDataData) GetVlsmNetworkIdOk() (*string, bool)`

GetVlsmNetworkIdOk returns a tuple with the VlsmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmNetworkId

`func (o *IpamNetworkDataData) SetVlsmNetworkId(v string)`

SetVlsmNetworkId sets VlsmNetworkId field to given value.

### HasVlsmNetworkId

`func (o *IpamNetworkDataData) HasVlsmNetworkId() bool`

HasVlsmNetworkId returns a boolean if a field has been set.

### GetNetworkRipeWaitingState

`func (o *IpamNetworkDataData) GetNetworkRipeWaitingState() string`

GetNetworkRipeWaitingState returns the NetworkRipeWaitingState field if non-nil, zero value otherwise.

### GetNetworkRipeWaitingStateOk

`func (o *IpamNetworkDataData) GetNetworkRipeWaitingStateOk() (*string, bool)`

GetNetworkRipeWaitingStateOk returns a tuple with the NetworkRipeWaitingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRipeWaitingState

`func (o *IpamNetworkDataData) SetNetworkRipeWaitingState(v string)`

SetNetworkRipeWaitingState sets NetworkRipeWaitingState field to given value.

### HasNetworkRipeWaitingState

`func (o *IpamNetworkDataData) HasNetworkRipeWaitingState() bool`

HasNetworkRipeWaitingState returns a boolean if a field has been set.

### GetNetworkRipeWaitingStatus

`func (o *IpamNetworkDataData) GetNetworkRipeWaitingStatus() string`

GetNetworkRipeWaitingStatus returns the NetworkRipeWaitingStatus field if non-nil, zero value otherwise.

### GetNetworkRipeWaitingStatusOk

`func (o *IpamNetworkDataData) GetNetworkRipeWaitingStatusOk() (*string, bool)`

GetNetworkRipeWaitingStatusOk returns a tuple with the NetworkRipeWaitingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRipeWaitingStatus

`func (o *IpamNetworkDataData) SetNetworkRipeWaitingStatus(v string)`

SetNetworkRipeWaitingStatus sets NetworkRipeWaitingStatus field to given value.

### HasNetworkRipeWaitingStatus

`func (o *IpamNetworkDataData) HasNetworkRipeWaitingStatus() bool`

HasNetworkRipeWaitingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


