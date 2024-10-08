# DataInnerIpamNetworkData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkEndHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;network_end_ip_addr&lt;/b&gt;. | [optional] 
**NetworkEndIpAddr** | Pointer to **string** | The last IP address of the IPv4 network, in hexadecimal format. | [optional] 
**NetworkIsInOrphan** | Pointer to **string** | A way to determine if the network has a parent (&lt;b&gt;0&lt;/b&gt;) or if it belongs to a container &lt;b&gt;Orphan networks&lt;/b&gt; (&lt;b&gt;1&lt;/b&gt;). | [optional] 
**NetworkIsTerminal** | Pointer to **string** | A way to determine if a network can contain other networks. If set to &lt;b&gt;1&lt;/b&gt;, the network is terminal and cannot contain other subnet-type networks. Block-type networks are always set to &lt;b&gt;0&lt;/b&gt;. | [optional] 
**NetworkLockNetworkBroadcast** | Pointer to **string** | A way to prevent &lt;b&gt;(1)&lt;/b&gt; users from assigning the broadcast IP address and network IP address of the network. | [optional] 
**NetworkMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentNetworkEndIpAddr** | Pointer to **string** | The last IP address of the parent IPv4 network, in hexadecimal format. | [optional] 
**ParentNetworkIsTerminal** | Pointer to **string** | A way to determine if the parent network is terminal (&lt;b&gt;1&lt;/b&gt;) or non-terminal (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**ParentSpaceId** | Pointer to **string** | The database identifier (ID) of the space where is located the parent network. &lt;b&gt;0&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentSpaceName** | Pointer to **string** | The name of the space where is located the parent network. &lt;b&gt;#&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentNetworkStartIpAddr** | Pointer to **string** | The first IP address of the parent IPv4 network, in hexadecimal format. | [optional] 
**ParentNetworkClassName** | Pointer to **string** | The name of the class applied to the parent IPv4 network, it can be preceded by the class directory. | [optional] 
**ParentNetworkClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the parent IPv4 network and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;...&lt;/b&gt; . | [optional] 
**ParentNetworkId** | Pointer to **string** | The database identifier (ID) of the parent IPv4 network. &lt;b&gt;0&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentNetworkLevel** | Pointer to **string** | The level of the parent network within the space. It returns values between &lt;b&gt;0&lt;/b&gt; (block-type network) and &lt;b&gt;n&lt;/b&gt; (subnet-type network). A value higher than &lt;b&gt;1&lt;/b&gt; indicates a VLSM organization where a block-type network can belong to another subnet-type network. | [optional] 
**ParentNetworkName** | Pointer to **string** | The name of the parent IPv4 network:&lt;ul class&#x3D;dashed &gt;&lt;li&gt; &lt;b&gt;#&lt;/b&gt; indicates that the network has no parent network.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; &lt;b&gt;Default&lt;/b&gt; indicates that the network belongs to an orphan network.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**ParentNetworkPath** | Pointer to **string** | The path toward the parent network in the database. &lt;b&gt;#&lt;/b&gt; indicates the network has no parent network. | [optional] 
**ParentNetworkSize** | Pointer to **string** | The number of IP addresses of the network parent. | [optional] 
**ParentVlsmNetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 subnet-type network, located in the VLSM parent space, from which the parent network was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the parent network is not a VLSM block-type network duplicated from a parent space. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the space the object belongs to. | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space the object belongs to. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceIsTemplate** | Pointer to **string** | The template status of the space the object belongs to. If the space is used as template (1), all the IPv4 networks, pools and IP addresses it contains are also used as template. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**SpaceParentSpaceId** | Pointer to **string** | The database identifier (ID) of the VLSM parent of the space where is located the network. &lt;b&gt;0&lt;/b&gt; indicates that the space where is located the network has no parent space. | [optional] 
**NetworkStartHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;network_start_ip_addr&lt;/b&gt;. | [optional] 
**NetworkStartIpAddr** | Pointer to **string** | The first IP address of the IPv4 network, in hexadecimal format. | [optional] 
**NetworkAllocatedPercent** | Pointer to **string** | The percentage of subnet-type networks the non-terminal network contains. | [optional] 
**NetworkAllocatedSize** | Pointer to **string** | The sum of the size of all the subnet-type networks that belong to the block-type network. | [optional] 
**NetworkClassName** | Pointer to **string** | The name of the class applied to the IPv4 network, it can be preceded by the class directory. | [optional] 
**NetworkClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv4 network. | [optional] 
**NetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 network. | [optional] 
**NetworkAddressFreeSize** | Pointer to **string** | The total number of free addresses, for terminal networks only. It excludes the network and broadcast IP address. | [optional] 
**NetworkAddressUsedPercent** | Pointer to **string** | The percentage of IP addresses &lt;b&gt;In use&lt;/b&gt; in terminal networks. | [optional] 
**NetworkAddressUsedSize** | Pointer to **string** | The number of IP addresses &lt;b&gt;In use&lt;/b&gt; in terminal networks. | [optional] 
**NetworkIsValid** | Pointer to **string** | The network validity. A valid network (&lt;b&gt;1&lt;/b&gt;) has a size, prefix and/or netmask that match. | [optional] 
**NetworkLevel** | Pointer to **string** | The level of the network within the space. It returns values between &lt;b&gt;0&lt;/b&gt; (block-type network) and &lt;b&gt;n&lt;/b&gt; (subnet-type network). A value higher than &lt;b&gt;1&lt;/b&gt; indicates a VLSM organization where a block-type network can belong to another subnet-type network. | [optional] 
**NetworkName** | Pointer to **string** | The name of the IPv4 network. &lt;b&gt;Default&lt;/b&gt; indicates that the network is an orphan network. | [optional] 
**NetworkPath** | Pointer to **string** | The path toward the network in the database from the containing block-type network down to the subnet-type network: &lt;b&gt;&amp;lt;block-network-start-IP&amp;gt;#&amp;lt;block-network-ID&amp;gt;#&amp;lt;subnet-network-start-IP&amp;gt;#&amp;lt;subnet-network-ID&amp;gt;&lt;/b&gt;. The IP address is returned in hexadecimal format.&lt;ul class&#x3D;dashed &gt;&lt;li&gt; In network-based VLSM organizations, the path includes all the subnet-type networks there are from the containing block-type network down to the subnet-type network specified in &lt;b&gt;network_id&lt;/b&gt;.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; In space-based VLSM organizations, the path includes the block-type network of the top parent space and all the subnet-type networks there are until the network specified in &lt;b&gt;network_id&lt;/b&gt;. Only one block-type network is returned.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**NetworkSize** | Pointer to **string** | The number of IP addresses the IPv4 network contains. | [optional] 
**NetworkUsedPercent** | Pointer to **string** | The percentage of terminal networks the non-terminal network contains. | [optional] 
**NetworkUsedSize** | Pointer to **string** | The sum of the size of all the terminal networks within the block-type network. This sum includes the terminal networks that might belong to non-terminal subnet-type networks. | [optional] 
**NetworkOperationDetailsAddedOn** | Pointer to **string** | The creation date of the IPv4 network, in decimal UNIX date format. | [optional] 
**NetworkOperationDetailsCallStack** | Pointer to **string** | The call stack of the IPv4 network operation details, as follows: &lt;b&gt;&amp;lt;service1&amp;gt;&amp;amp;&amp;lt;service2&amp;gt;&amp;amp;&amp;lt;service3&amp;gt;&lt;/b&gt;... . | [optional] 
**NetworkOperationDetailsOriginModule** | Pointer to **string** | The name of the module where the IPv4 network addition originated. | [optional] 
**NetworkOperationDetailsRequestedBy** | Pointer to **string** | The login of the user who requested the IPv4 network. | [optional] 
**NetworkOperationDetailsAddedBy** | Pointer to **string** | The login of the user who added the IPv4 network. | [optional] 
**NetworkOperationDetailsLastUpdatedOn** | Pointer to **string** | The last time the IPv4 network was updated, in decimal UNIX date format. | [optional] 
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
**NetworkRipeWaitingState** | Pointer to **string** | The state of the exchange between SOLIDserver and the RIPE for the assigned network: &lt;table&gt;&lt;caption&gt;network_ripe_waiting_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;must_send_mail_add&lt;/td&gt;&lt;td &gt;An email must be sent to the RIPE to notify them of a subnet-type network creation.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_mail_add&lt;/td&gt;&lt;td &gt;A network creation email was sent to the RIPE, no reply has been received yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;must_send_mail_del&lt;/td&gt;&lt;td &gt;An email must be sent to the RIPE to notify them of a subnet-type network deletion.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_mail_del&lt;/td&gt;&lt;td &gt;A network deletion email was sent to the RIPE, no reply has been received yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_aw_confirm&lt;/td&gt;&lt;td &gt;The number of IP addresses of the assigned network exceeds the Assignment Window declared during your RIPE configuration.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**NetworkRipeWaitingStatus** | Pointer to **string** | The status of a RIPE assigned network within SOLIDserver until it is confirmed that you can create or delete it. If set to &lt;b&gt;1&lt;/b&gt;, it is about to be created. If set to &lt;b&gt;2&lt;/b&gt;, it is about to be deleted. | [optional] 

## Methods

### NewDataInnerIpamNetworkData

`func NewDataInnerIpamNetworkData() *DataInnerIpamNetworkData`

NewDataInnerIpamNetworkData instantiates a new DataInnerIpamNetworkData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerIpamNetworkDataWithDefaults

`func NewDataInnerIpamNetworkDataWithDefaults() *DataInnerIpamNetworkData`

NewDataInnerIpamNetworkDataWithDefaults instantiates a new DataInnerIpamNetworkData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkEndHostaddr

`func (o *DataInnerIpamNetworkData) GetNetworkEndHostaddr() string`

GetNetworkEndHostaddr returns the NetworkEndHostaddr field if non-nil, zero value otherwise.

### GetNetworkEndHostaddrOk

`func (o *DataInnerIpamNetworkData) GetNetworkEndHostaddrOk() (*string, bool)`

GetNetworkEndHostaddrOk returns a tuple with the NetworkEndHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndHostaddr

`func (o *DataInnerIpamNetworkData) SetNetworkEndHostaddr(v string)`

SetNetworkEndHostaddr sets NetworkEndHostaddr field to given value.

### HasNetworkEndHostaddr

`func (o *DataInnerIpamNetworkData) HasNetworkEndHostaddr() bool`

HasNetworkEndHostaddr returns a boolean if a field has been set.

### GetNetworkEndIpAddr

`func (o *DataInnerIpamNetworkData) GetNetworkEndIpAddr() string`

GetNetworkEndIpAddr returns the NetworkEndIpAddr field if non-nil, zero value otherwise.

### GetNetworkEndIpAddrOk

`func (o *DataInnerIpamNetworkData) GetNetworkEndIpAddrOk() (*string, bool)`

GetNetworkEndIpAddrOk returns a tuple with the NetworkEndIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndIpAddr

`func (o *DataInnerIpamNetworkData) SetNetworkEndIpAddr(v string)`

SetNetworkEndIpAddr sets NetworkEndIpAddr field to given value.

### HasNetworkEndIpAddr

`func (o *DataInnerIpamNetworkData) HasNetworkEndIpAddr() bool`

HasNetworkEndIpAddr returns a boolean if a field has been set.

### GetNetworkIsInOrphan

`func (o *DataInnerIpamNetworkData) GetNetworkIsInOrphan() string`

GetNetworkIsInOrphan returns the NetworkIsInOrphan field if non-nil, zero value otherwise.

### GetNetworkIsInOrphanOk

`func (o *DataInnerIpamNetworkData) GetNetworkIsInOrphanOk() (*string, bool)`

GetNetworkIsInOrphanOk returns a tuple with the NetworkIsInOrphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsInOrphan

`func (o *DataInnerIpamNetworkData) SetNetworkIsInOrphan(v string)`

SetNetworkIsInOrphan sets NetworkIsInOrphan field to given value.

### HasNetworkIsInOrphan

`func (o *DataInnerIpamNetworkData) HasNetworkIsInOrphan() bool`

HasNetworkIsInOrphan returns a boolean if a field has been set.

### GetNetworkIsTerminal

`func (o *DataInnerIpamNetworkData) GetNetworkIsTerminal() string`

GetNetworkIsTerminal returns the NetworkIsTerminal field if non-nil, zero value otherwise.

### GetNetworkIsTerminalOk

`func (o *DataInnerIpamNetworkData) GetNetworkIsTerminalOk() (*string, bool)`

GetNetworkIsTerminalOk returns a tuple with the NetworkIsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsTerminal

`func (o *DataInnerIpamNetworkData) SetNetworkIsTerminal(v string)`

SetNetworkIsTerminal sets NetworkIsTerminal field to given value.

### HasNetworkIsTerminal

`func (o *DataInnerIpamNetworkData) HasNetworkIsTerminal() bool`

HasNetworkIsTerminal returns a boolean if a field has been set.

### GetNetworkLockNetworkBroadcast

`func (o *DataInnerIpamNetworkData) GetNetworkLockNetworkBroadcast() string`

GetNetworkLockNetworkBroadcast returns the NetworkLockNetworkBroadcast field if non-nil, zero value otherwise.

### GetNetworkLockNetworkBroadcastOk

`func (o *DataInnerIpamNetworkData) GetNetworkLockNetworkBroadcastOk() (*string, bool)`

GetNetworkLockNetworkBroadcastOk returns a tuple with the NetworkLockNetworkBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLockNetworkBroadcast

`func (o *DataInnerIpamNetworkData) SetNetworkLockNetworkBroadcast(v string)`

SetNetworkLockNetworkBroadcast sets NetworkLockNetworkBroadcast field to given value.

### HasNetworkLockNetworkBroadcast

`func (o *DataInnerIpamNetworkData) HasNetworkLockNetworkBroadcast() bool`

HasNetworkLockNetworkBroadcast returns a boolean if a field has been set.

### GetNetworkMultistatus

`func (o *DataInnerIpamNetworkData) GetNetworkMultistatus() string`

GetNetworkMultistatus returns the NetworkMultistatus field if non-nil, zero value otherwise.

### GetNetworkMultistatusOk

`func (o *DataInnerIpamNetworkData) GetNetworkMultistatusOk() (*string, bool)`

GetNetworkMultistatusOk returns a tuple with the NetworkMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMultistatus

`func (o *DataInnerIpamNetworkData) SetNetworkMultistatus(v string)`

SetNetworkMultistatus sets NetworkMultistatus field to given value.

### HasNetworkMultistatus

`func (o *DataInnerIpamNetworkData) HasNetworkMultistatus() bool`

HasNetworkMultistatus returns a boolean if a field has been set.

### GetParentNetworkEndIpAddr

`func (o *DataInnerIpamNetworkData) GetParentNetworkEndIpAddr() string`

GetParentNetworkEndIpAddr returns the ParentNetworkEndIpAddr field if non-nil, zero value otherwise.

### GetParentNetworkEndIpAddrOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkEndIpAddrOk() (*string, bool)`

GetParentNetworkEndIpAddrOk returns a tuple with the ParentNetworkEndIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkEndIpAddr

`func (o *DataInnerIpamNetworkData) SetParentNetworkEndIpAddr(v string)`

SetParentNetworkEndIpAddr sets ParentNetworkEndIpAddr field to given value.

### HasParentNetworkEndIpAddr

`func (o *DataInnerIpamNetworkData) HasParentNetworkEndIpAddr() bool`

HasParentNetworkEndIpAddr returns a boolean if a field has been set.

### GetParentNetworkIsTerminal

`func (o *DataInnerIpamNetworkData) GetParentNetworkIsTerminal() string`

GetParentNetworkIsTerminal returns the ParentNetworkIsTerminal field if non-nil, zero value otherwise.

### GetParentNetworkIsTerminalOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkIsTerminalOk() (*string, bool)`

GetParentNetworkIsTerminalOk returns a tuple with the ParentNetworkIsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkIsTerminal

`func (o *DataInnerIpamNetworkData) SetParentNetworkIsTerminal(v string)`

SetParentNetworkIsTerminal sets ParentNetworkIsTerminal field to given value.

### HasParentNetworkIsTerminal

`func (o *DataInnerIpamNetworkData) HasParentNetworkIsTerminal() bool`

HasParentNetworkIsTerminal returns a boolean if a field has been set.

### GetParentSpaceId

`func (o *DataInnerIpamNetworkData) GetParentSpaceId() string`

GetParentSpaceId returns the ParentSpaceId field if non-nil, zero value otherwise.

### GetParentSpaceIdOk

`func (o *DataInnerIpamNetworkData) GetParentSpaceIdOk() (*string, bool)`

GetParentSpaceIdOk returns a tuple with the ParentSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceId

`func (o *DataInnerIpamNetworkData) SetParentSpaceId(v string)`

SetParentSpaceId sets ParentSpaceId field to given value.

### HasParentSpaceId

`func (o *DataInnerIpamNetworkData) HasParentSpaceId() bool`

HasParentSpaceId returns a boolean if a field has been set.

### GetParentSpaceName

`func (o *DataInnerIpamNetworkData) GetParentSpaceName() string`

GetParentSpaceName returns the ParentSpaceName field if non-nil, zero value otherwise.

### GetParentSpaceNameOk

`func (o *DataInnerIpamNetworkData) GetParentSpaceNameOk() (*string, bool)`

GetParentSpaceNameOk returns a tuple with the ParentSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceName

`func (o *DataInnerIpamNetworkData) SetParentSpaceName(v string)`

SetParentSpaceName sets ParentSpaceName field to given value.

### HasParentSpaceName

`func (o *DataInnerIpamNetworkData) HasParentSpaceName() bool`

HasParentSpaceName returns a boolean if a field has been set.

### GetParentNetworkStartIpAddr

`func (o *DataInnerIpamNetworkData) GetParentNetworkStartIpAddr() string`

GetParentNetworkStartIpAddr returns the ParentNetworkStartIpAddr field if non-nil, zero value otherwise.

### GetParentNetworkStartIpAddrOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkStartIpAddrOk() (*string, bool)`

GetParentNetworkStartIpAddrOk returns a tuple with the ParentNetworkStartIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkStartIpAddr

`func (o *DataInnerIpamNetworkData) SetParentNetworkStartIpAddr(v string)`

SetParentNetworkStartIpAddr sets ParentNetworkStartIpAddr field to given value.

### HasParentNetworkStartIpAddr

`func (o *DataInnerIpamNetworkData) HasParentNetworkStartIpAddr() bool`

HasParentNetworkStartIpAddr returns a boolean if a field has been set.

### GetParentNetworkClassName

`func (o *DataInnerIpamNetworkData) GetParentNetworkClassName() string`

GetParentNetworkClassName returns the ParentNetworkClassName field if non-nil, zero value otherwise.

### GetParentNetworkClassNameOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkClassNameOk() (*string, bool)`

GetParentNetworkClassNameOk returns a tuple with the ParentNetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkClassName

`func (o *DataInnerIpamNetworkData) SetParentNetworkClassName(v string)`

SetParentNetworkClassName sets ParentNetworkClassName field to given value.

### HasParentNetworkClassName

`func (o *DataInnerIpamNetworkData) HasParentNetworkClassName() bool`

HasParentNetworkClassName returns a boolean if a field has been set.

### GetParentNetworkClassParameters

`func (o *DataInnerIpamNetworkData) GetParentNetworkClassParameters() []ApiClassParameterOutputEntry`

GetParentNetworkClassParameters returns the ParentNetworkClassParameters field if non-nil, zero value otherwise.

### GetParentNetworkClassParametersOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetParentNetworkClassParametersOk returns a tuple with the ParentNetworkClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkClassParameters

`func (o *DataInnerIpamNetworkData) SetParentNetworkClassParameters(v []ApiClassParameterOutputEntry)`

SetParentNetworkClassParameters sets ParentNetworkClassParameters field to given value.

### HasParentNetworkClassParameters

`func (o *DataInnerIpamNetworkData) HasParentNetworkClassParameters() bool`

HasParentNetworkClassParameters returns a boolean if a field has been set.

### GetParentNetworkId

`func (o *DataInnerIpamNetworkData) GetParentNetworkId() string`

GetParentNetworkId returns the ParentNetworkId field if non-nil, zero value otherwise.

### GetParentNetworkIdOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkIdOk() (*string, bool)`

GetParentNetworkIdOk returns a tuple with the ParentNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkId

`func (o *DataInnerIpamNetworkData) SetParentNetworkId(v string)`

SetParentNetworkId sets ParentNetworkId field to given value.

### HasParentNetworkId

`func (o *DataInnerIpamNetworkData) HasParentNetworkId() bool`

HasParentNetworkId returns a boolean if a field has been set.

### GetParentNetworkLevel

`func (o *DataInnerIpamNetworkData) GetParentNetworkLevel() string`

GetParentNetworkLevel returns the ParentNetworkLevel field if non-nil, zero value otherwise.

### GetParentNetworkLevelOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkLevelOk() (*string, bool)`

GetParentNetworkLevelOk returns a tuple with the ParentNetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkLevel

`func (o *DataInnerIpamNetworkData) SetParentNetworkLevel(v string)`

SetParentNetworkLevel sets ParentNetworkLevel field to given value.

### HasParentNetworkLevel

`func (o *DataInnerIpamNetworkData) HasParentNetworkLevel() bool`

HasParentNetworkLevel returns a boolean if a field has been set.

### GetParentNetworkName

`func (o *DataInnerIpamNetworkData) GetParentNetworkName() string`

GetParentNetworkName returns the ParentNetworkName field if non-nil, zero value otherwise.

### GetParentNetworkNameOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkNameOk() (*string, bool)`

GetParentNetworkNameOk returns a tuple with the ParentNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkName

`func (o *DataInnerIpamNetworkData) SetParentNetworkName(v string)`

SetParentNetworkName sets ParentNetworkName field to given value.

### HasParentNetworkName

`func (o *DataInnerIpamNetworkData) HasParentNetworkName() bool`

HasParentNetworkName returns a boolean if a field has been set.

### GetParentNetworkPath

`func (o *DataInnerIpamNetworkData) GetParentNetworkPath() string`

GetParentNetworkPath returns the ParentNetworkPath field if non-nil, zero value otherwise.

### GetParentNetworkPathOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkPathOk() (*string, bool)`

GetParentNetworkPathOk returns a tuple with the ParentNetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkPath

`func (o *DataInnerIpamNetworkData) SetParentNetworkPath(v string)`

SetParentNetworkPath sets ParentNetworkPath field to given value.

### HasParentNetworkPath

`func (o *DataInnerIpamNetworkData) HasParentNetworkPath() bool`

HasParentNetworkPath returns a boolean if a field has been set.

### GetParentNetworkSize

`func (o *DataInnerIpamNetworkData) GetParentNetworkSize() string`

GetParentNetworkSize returns the ParentNetworkSize field if non-nil, zero value otherwise.

### GetParentNetworkSizeOk

`func (o *DataInnerIpamNetworkData) GetParentNetworkSizeOk() (*string, bool)`

GetParentNetworkSizeOk returns a tuple with the ParentNetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkSize

`func (o *DataInnerIpamNetworkData) SetParentNetworkSize(v string)`

SetParentNetworkSize sets ParentNetworkSize field to given value.

### HasParentNetworkSize

`func (o *DataInnerIpamNetworkData) HasParentNetworkSize() bool`

HasParentNetworkSize returns a boolean if a field has been set.

### GetParentVlsmNetworkId

`func (o *DataInnerIpamNetworkData) GetParentVlsmNetworkId() string`

GetParentVlsmNetworkId returns the ParentVlsmNetworkId field if non-nil, zero value otherwise.

### GetParentVlsmNetworkIdOk

`func (o *DataInnerIpamNetworkData) GetParentVlsmNetworkIdOk() (*string, bool)`

GetParentVlsmNetworkIdOk returns a tuple with the ParentVlsmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVlsmNetworkId

`func (o *DataInnerIpamNetworkData) SetParentVlsmNetworkId(v string)`

SetParentVlsmNetworkId sets ParentVlsmNetworkId field to given value.

### HasParentVlsmNetworkId

`func (o *DataInnerIpamNetworkData) HasParentVlsmNetworkId() bool`

HasParentVlsmNetworkId returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *DataInnerIpamNetworkData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *DataInnerIpamNetworkData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *DataInnerIpamNetworkData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *DataInnerIpamNetworkData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *DataInnerIpamNetworkData) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *DataInnerIpamNetworkData) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *DataInnerIpamNetworkData) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *DataInnerIpamNetworkData) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *DataInnerIpamNetworkData) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *DataInnerIpamNetworkData) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *DataInnerIpamNetworkData) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *DataInnerIpamNetworkData) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *DataInnerIpamNetworkData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DataInnerIpamNetworkData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DataInnerIpamNetworkData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DataInnerIpamNetworkData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceIsTemplate

`func (o *DataInnerIpamNetworkData) GetSpaceIsTemplate() string`

GetSpaceIsTemplate returns the SpaceIsTemplate field if non-nil, zero value otherwise.

### GetSpaceIsTemplateOk

`func (o *DataInnerIpamNetworkData) GetSpaceIsTemplateOk() (*string, bool)`

GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceIsTemplate

`func (o *DataInnerIpamNetworkData) SetSpaceIsTemplate(v string)`

SetSpaceIsTemplate sets SpaceIsTemplate field to given value.

### HasSpaceIsTemplate

`func (o *DataInnerIpamNetworkData) HasSpaceIsTemplate() bool`

HasSpaceIsTemplate returns a boolean if a field has been set.

### GetSpaceName

`func (o *DataInnerIpamNetworkData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *DataInnerIpamNetworkData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *DataInnerIpamNetworkData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *DataInnerIpamNetworkData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetSpaceParentSpaceId

`func (o *DataInnerIpamNetworkData) GetSpaceParentSpaceId() string`

GetSpaceParentSpaceId returns the SpaceParentSpaceId field if non-nil, zero value otherwise.

### GetSpaceParentSpaceIdOk

`func (o *DataInnerIpamNetworkData) GetSpaceParentSpaceIdOk() (*string, bool)`

GetSpaceParentSpaceIdOk returns a tuple with the SpaceParentSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceParentSpaceId

`func (o *DataInnerIpamNetworkData) SetSpaceParentSpaceId(v string)`

SetSpaceParentSpaceId sets SpaceParentSpaceId field to given value.

### HasSpaceParentSpaceId

`func (o *DataInnerIpamNetworkData) HasSpaceParentSpaceId() bool`

HasSpaceParentSpaceId returns a boolean if a field has been set.

### GetNetworkStartHostaddr

`func (o *DataInnerIpamNetworkData) GetNetworkStartHostaddr() string`

GetNetworkStartHostaddr returns the NetworkStartHostaddr field if non-nil, zero value otherwise.

### GetNetworkStartHostaddrOk

`func (o *DataInnerIpamNetworkData) GetNetworkStartHostaddrOk() (*string, bool)`

GetNetworkStartHostaddrOk returns a tuple with the NetworkStartHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStartHostaddr

`func (o *DataInnerIpamNetworkData) SetNetworkStartHostaddr(v string)`

SetNetworkStartHostaddr sets NetworkStartHostaddr field to given value.

### HasNetworkStartHostaddr

`func (o *DataInnerIpamNetworkData) HasNetworkStartHostaddr() bool`

HasNetworkStartHostaddr returns a boolean if a field has been set.

### GetNetworkStartIpAddr

`func (o *DataInnerIpamNetworkData) GetNetworkStartIpAddr() string`

GetNetworkStartIpAddr returns the NetworkStartIpAddr field if non-nil, zero value otherwise.

### GetNetworkStartIpAddrOk

`func (o *DataInnerIpamNetworkData) GetNetworkStartIpAddrOk() (*string, bool)`

GetNetworkStartIpAddrOk returns a tuple with the NetworkStartIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStartIpAddr

`func (o *DataInnerIpamNetworkData) SetNetworkStartIpAddr(v string)`

SetNetworkStartIpAddr sets NetworkStartIpAddr field to given value.

### HasNetworkStartIpAddr

`func (o *DataInnerIpamNetworkData) HasNetworkStartIpAddr() bool`

HasNetworkStartIpAddr returns a boolean if a field has been set.

### GetNetworkAllocatedPercent

`func (o *DataInnerIpamNetworkData) GetNetworkAllocatedPercent() string`

GetNetworkAllocatedPercent returns the NetworkAllocatedPercent field if non-nil, zero value otherwise.

### GetNetworkAllocatedPercentOk

`func (o *DataInnerIpamNetworkData) GetNetworkAllocatedPercentOk() (*string, bool)`

GetNetworkAllocatedPercentOk returns a tuple with the NetworkAllocatedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAllocatedPercent

`func (o *DataInnerIpamNetworkData) SetNetworkAllocatedPercent(v string)`

SetNetworkAllocatedPercent sets NetworkAllocatedPercent field to given value.

### HasNetworkAllocatedPercent

`func (o *DataInnerIpamNetworkData) HasNetworkAllocatedPercent() bool`

HasNetworkAllocatedPercent returns a boolean if a field has been set.

### GetNetworkAllocatedSize

`func (o *DataInnerIpamNetworkData) GetNetworkAllocatedSize() string`

GetNetworkAllocatedSize returns the NetworkAllocatedSize field if non-nil, zero value otherwise.

### GetNetworkAllocatedSizeOk

`func (o *DataInnerIpamNetworkData) GetNetworkAllocatedSizeOk() (*string, bool)`

GetNetworkAllocatedSizeOk returns a tuple with the NetworkAllocatedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAllocatedSize

`func (o *DataInnerIpamNetworkData) SetNetworkAllocatedSize(v string)`

SetNetworkAllocatedSize sets NetworkAllocatedSize field to given value.

### HasNetworkAllocatedSize

`func (o *DataInnerIpamNetworkData) HasNetworkAllocatedSize() bool`

HasNetworkAllocatedSize returns a boolean if a field has been set.

### GetNetworkClassName

`func (o *DataInnerIpamNetworkData) GetNetworkClassName() string`

GetNetworkClassName returns the NetworkClassName field if non-nil, zero value otherwise.

### GetNetworkClassNameOk

`func (o *DataInnerIpamNetworkData) GetNetworkClassNameOk() (*string, bool)`

GetNetworkClassNameOk returns a tuple with the NetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassName

`func (o *DataInnerIpamNetworkData) SetNetworkClassName(v string)`

SetNetworkClassName sets NetworkClassName field to given value.

### HasNetworkClassName

`func (o *DataInnerIpamNetworkData) HasNetworkClassName() bool`

HasNetworkClassName returns a boolean if a field has been set.

### GetNetworkClassParameters

`func (o *DataInnerIpamNetworkData) GetNetworkClassParameters() []ApiClassParameterOutputEntry`

GetNetworkClassParameters returns the NetworkClassParameters field if non-nil, zero value otherwise.

### GetNetworkClassParametersOk

`func (o *DataInnerIpamNetworkData) GetNetworkClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetNetworkClassParametersOk returns a tuple with the NetworkClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassParameters

`func (o *DataInnerIpamNetworkData) SetNetworkClassParameters(v []ApiClassParameterOutputEntry)`

SetNetworkClassParameters sets NetworkClassParameters field to given value.

### HasNetworkClassParameters

`func (o *DataInnerIpamNetworkData) HasNetworkClassParameters() bool`

HasNetworkClassParameters returns a boolean if a field has been set.

### GetNetworkId

`func (o *DataInnerIpamNetworkData) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *DataInnerIpamNetworkData) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *DataInnerIpamNetworkData) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *DataInnerIpamNetworkData) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkAddressFreeSize

`func (o *DataInnerIpamNetworkData) GetNetworkAddressFreeSize() string`

GetNetworkAddressFreeSize returns the NetworkAddressFreeSize field if non-nil, zero value otherwise.

### GetNetworkAddressFreeSizeOk

`func (o *DataInnerIpamNetworkData) GetNetworkAddressFreeSizeOk() (*string, bool)`

GetNetworkAddressFreeSizeOk returns a tuple with the NetworkAddressFreeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressFreeSize

`func (o *DataInnerIpamNetworkData) SetNetworkAddressFreeSize(v string)`

SetNetworkAddressFreeSize sets NetworkAddressFreeSize field to given value.

### HasNetworkAddressFreeSize

`func (o *DataInnerIpamNetworkData) HasNetworkAddressFreeSize() bool`

HasNetworkAddressFreeSize returns a boolean if a field has been set.

### GetNetworkAddressUsedPercent

`func (o *DataInnerIpamNetworkData) GetNetworkAddressUsedPercent() string`

GetNetworkAddressUsedPercent returns the NetworkAddressUsedPercent field if non-nil, zero value otherwise.

### GetNetworkAddressUsedPercentOk

`func (o *DataInnerIpamNetworkData) GetNetworkAddressUsedPercentOk() (*string, bool)`

GetNetworkAddressUsedPercentOk returns a tuple with the NetworkAddressUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressUsedPercent

`func (o *DataInnerIpamNetworkData) SetNetworkAddressUsedPercent(v string)`

SetNetworkAddressUsedPercent sets NetworkAddressUsedPercent field to given value.

### HasNetworkAddressUsedPercent

`func (o *DataInnerIpamNetworkData) HasNetworkAddressUsedPercent() bool`

HasNetworkAddressUsedPercent returns a boolean if a field has been set.

### GetNetworkAddressUsedSize

`func (o *DataInnerIpamNetworkData) GetNetworkAddressUsedSize() string`

GetNetworkAddressUsedSize returns the NetworkAddressUsedSize field if non-nil, zero value otherwise.

### GetNetworkAddressUsedSizeOk

`func (o *DataInnerIpamNetworkData) GetNetworkAddressUsedSizeOk() (*string, bool)`

GetNetworkAddressUsedSizeOk returns a tuple with the NetworkAddressUsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddressUsedSize

`func (o *DataInnerIpamNetworkData) SetNetworkAddressUsedSize(v string)`

SetNetworkAddressUsedSize sets NetworkAddressUsedSize field to given value.

### HasNetworkAddressUsedSize

`func (o *DataInnerIpamNetworkData) HasNetworkAddressUsedSize() bool`

HasNetworkAddressUsedSize returns a boolean if a field has been set.

### GetNetworkIsValid

`func (o *DataInnerIpamNetworkData) GetNetworkIsValid() string`

GetNetworkIsValid returns the NetworkIsValid field if non-nil, zero value otherwise.

### GetNetworkIsValidOk

`func (o *DataInnerIpamNetworkData) GetNetworkIsValidOk() (*string, bool)`

GetNetworkIsValidOk returns a tuple with the NetworkIsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsValid

`func (o *DataInnerIpamNetworkData) SetNetworkIsValid(v string)`

SetNetworkIsValid sets NetworkIsValid field to given value.

### HasNetworkIsValid

`func (o *DataInnerIpamNetworkData) HasNetworkIsValid() bool`

HasNetworkIsValid returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *DataInnerIpamNetworkData) GetNetworkLevel() string`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *DataInnerIpamNetworkData) GetNetworkLevelOk() (*string, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *DataInnerIpamNetworkData) SetNetworkLevel(v string)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *DataInnerIpamNetworkData) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetNetworkName

`func (o *DataInnerIpamNetworkData) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *DataInnerIpamNetworkData) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *DataInnerIpamNetworkData) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *DataInnerIpamNetworkData) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkPath

`func (o *DataInnerIpamNetworkData) GetNetworkPath() string`

GetNetworkPath returns the NetworkPath field if non-nil, zero value otherwise.

### GetNetworkPathOk

`func (o *DataInnerIpamNetworkData) GetNetworkPathOk() (*string, bool)`

GetNetworkPathOk returns a tuple with the NetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPath

`func (o *DataInnerIpamNetworkData) SetNetworkPath(v string)`

SetNetworkPath sets NetworkPath field to given value.

### HasNetworkPath

`func (o *DataInnerIpamNetworkData) HasNetworkPath() bool`

HasNetworkPath returns a boolean if a field has been set.

### GetNetworkSize

`func (o *DataInnerIpamNetworkData) GetNetworkSize() string`

GetNetworkSize returns the NetworkSize field if non-nil, zero value otherwise.

### GetNetworkSizeOk

`func (o *DataInnerIpamNetworkData) GetNetworkSizeOk() (*string, bool)`

GetNetworkSizeOk returns a tuple with the NetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSize

`func (o *DataInnerIpamNetworkData) SetNetworkSize(v string)`

SetNetworkSize sets NetworkSize field to given value.

### HasNetworkSize

`func (o *DataInnerIpamNetworkData) HasNetworkSize() bool`

HasNetworkSize returns a boolean if a field has been set.

### GetNetworkUsedPercent

`func (o *DataInnerIpamNetworkData) GetNetworkUsedPercent() string`

GetNetworkUsedPercent returns the NetworkUsedPercent field if non-nil, zero value otherwise.

### GetNetworkUsedPercentOk

`func (o *DataInnerIpamNetworkData) GetNetworkUsedPercentOk() (*string, bool)`

GetNetworkUsedPercentOk returns a tuple with the NetworkUsedPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUsedPercent

`func (o *DataInnerIpamNetworkData) SetNetworkUsedPercent(v string)`

SetNetworkUsedPercent sets NetworkUsedPercent field to given value.

### HasNetworkUsedPercent

`func (o *DataInnerIpamNetworkData) HasNetworkUsedPercent() bool`

HasNetworkUsedPercent returns a boolean if a field has been set.

### GetNetworkUsedSize

`func (o *DataInnerIpamNetworkData) GetNetworkUsedSize() string`

GetNetworkUsedSize returns the NetworkUsedSize field if non-nil, zero value otherwise.

### GetNetworkUsedSizeOk

`func (o *DataInnerIpamNetworkData) GetNetworkUsedSizeOk() (*string, bool)`

GetNetworkUsedSizeOk returns a tuple with the NetworkUsedSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkUsedSize

`func (o *DataInnerIpamNetworkData) SetNetworkUsedSize(v string)`

SetNetworkUsedSize sets NetworkUsedSize field to given value.

### HasNetworkUsedSize

`func (o *DataInnerIpamNetworkData) HasNetworkUsedSize() bool`

HasNetworkUsedSize returns a boolean if a field has been set.

### GetNetworkOperationDetailsAddedOn

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsAddedOn() string`

GetNetworkOperationDetailsAddedOn returns the NetworkOperationDetailsAddedOn field if non-nil, zero value otherwise.

### GetNetworkOperationDetailsAddedOnOk

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsAddedOnOk() (*string, bool)`

GetNetworkOperationDetailsAddedOnOk returns a tuple with the NetworkOperationDetailsAddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOperationDetailsAddedOn

`func (o *DataInnerIpamNetworkData) SetNetworkOperationDetailsAddedOn(v string)`

SetNetworkOperationDetailsAddedOn sets NetworkOperationDetailsAddedOn field to given value.

### HasNetworkOperationDetailsAddedOn

`func (o *DataInnerIpamNetworkData) HasNetworkOperationDetailsAddedOn() bool`

HasNetworkOperationDetailsAddedOn returns a boolean if a field has been set.

### GetNetworkOperationDetailsCallStack

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsCallStack() string`

GetNetworkOperationDetailsCallStack returns the NetworkOperationDetailsCallStack field if non-nil, zero value otherwise.

### GetNetworkOperationDetailsCallStackOk

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsCallStackOk() (*string, bool)`

GetNetworkOperationDetailsCallStackOk returns a tuple with the NetworkOperationDetailsCallStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOperationDetailsCallStack

`func (o *DataInnerIpamNetworkData) SetNetworkOperationDetailsCallStack(v string)`

SetNetworkOperationDetailsCallStack sets NetworkOperationDetailsCallStack field to given value.

### HasNetworkOperationDetailsCallStack

`func (o *DataInnerIpamNetworkData) HasNetworkOperationDetailsCallStack() bool`

HasNetworkOperationDetailsCallStack returns a boolean if a field has been set.

### GetNetworkOperationDetailsOriginModule

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsOriginModule() string`

GetNetworkOperationDetailsOriginModule returns the NetworkOperationDetailsOriginModule field if non-nil, zero value otherwise.

### GetNetworkOperationDetailsOriginModuleOk

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsOriginModuleOk() (*string, bool)`

GetNetworkOperationDetailsOriginModuleOk returns a tuple with the NetworkOperationDetailsOriginModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOperationDetailsOriginModule

`func (o *DataInnerIpamNetworkData) SetNetworkOperationDetailsOriginModule(v string)`

SetNetworkOperationDetailsOriginModule sets NetworkOperationDetailsOriginModule field to given value.

### HasNetworkOperationDetailsOriginModule

`func (o *DataInnerIpamNetworkData) HasNetworkOperationDetailsOriginModule() bool`

HasNetworkOperationDetailsOriginModule returns a boolean if a field has been set.

### GetNetworkOperationDetailsRequestedBy

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsRequestedBy() string`

GetNetworkOperationDetailsRequestedBy returns the NetworkOperationDetailsRequestedBy field if non-nil, zero value otherwise.

### GetNetworkOperationDetailsRequestedByOk

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsRequestedByOk() (*string, bool)`

GetNetworkOperationDetailsRequestedByOk returns a tuple with the NetworkOperationDetailsRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOperationDetailsRequestedBy

`func (o *DataInnerIpamNetworkData) SetNetworkOperationDetailsRequestedBy(v string)`

SetNetworkOperationDetailsRequestedBy sets NetworkOperationDetailsRequestedBy field to given value.

### HasNetworkOperationDetailsRequestedBy

`func (o *DataInnerIpamNetworkData) HasNetworkOperationDetailsRequestedBy() bool`

HasNetworkOperationDetailsRequestedBy returns a boolean if a field has been set.

### GetNetworkOperationDetailsAddedBy

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsAddedBy() string`

GetNetworkOperationDetailsAddedBy returns the NetworkOperationDetailsAddedBy field if non-nil, zero value otherwise.

### GetNetworkOperationDetailsAddedByOk

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsAddedByOk() (*string, bool)`

GetNetworkOperationDetailsAddedByOk returns a tuple with the NetworkOperationDetailsAddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOperationDetailsAddedBy

`func (o *DataInnerIpamNetworkData) SetNetworkOperationDetailsAddedBy(v string)`

SetNetworkOperationDetailsAddedBy sets NetworkOperationDetailsAddedBy field to given value.

### HasNetworkOperationDetailsAddedBy

`func (o *DataInnerIpamNetworkData) HasNetworkOperationDetailsAddedBy() bool`

HasNetworkOperationDetailsAddedBy returns a boolean if a field has been set.

### GetNetworkOperationDetailsLastUpdatedOn

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsLastUpdatedOn() string`

GetNetworkOperationDetailsLastUpdatedOn returns the NetworkOperationDetailsLastUpdatedOn field if non-nil, zero value otherwise.

### GetNetworkOperationDetailsLastUpdatedOnOk

`func (o *DataInnerIpamNetworkData) GetNetworkOperationDetailsLastUpdatedOnOk() (*string, bool)`

GetNetworkOperationDetailsLastUpdatedOnOk returns a tuple with the NetworkOperationDetailsLastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkOperationDetailsLastUpdatedOn

`func (o *DataInnerIpamNetworkData) SetNetworkOperationDetailsLastUpdatedOn(v string)`

SetNetworkOperationDetailsLastUpdatedOn sets NetworkOperationDetailsLastUpdatedOn field to given value.

### HasNetworkOperationDetailsLastUpdatedOn

`func (o *DataInnerIpamNetworkData) HasNetworkOperationDetailsLastUpdatedOn() bool`

HasNetworkOperationDetailsLastUpdatedOn returns a boolean if a field has been set.

### GetDomainId

`func (o *DataInnerIpamNetworkData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DataInnerIpamNetworkData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DataInnerIpamNetworkData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DataInnerIpamNetworkData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *DataInnerIpamNetworkData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DataInnerIpamNetworkData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DataInnerIpamNetworkData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DataInnerIpamNetworkData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRangeId

`func (o *DataInnerIpamNetworkData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *DataInnerIpamNetworkData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *DataInnerIpamNetworkData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *DataInnerIpamNetworkData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *DataInnerIpamNetworkData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *DataInnerIpamNetworkData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *DataInnerIpamNetworkData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *DataInnerIpamNetworkData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetVlanId

`func (o *DataInnerIpamNetworkData) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *DataInnerIpamNetworkData) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *DataInnerIpamNetworkData) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *DataInnerIpamNetworkData) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanName

`func (o *DataInnerIpamNetworkData) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *DataInnerIpamNetworkData) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *DataInnerIpamNetworkData) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *DataInnerIpamNetworkData) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetVlanVlanId

`func (o *DataInnerIpamNetworkData) GetVlanVlanId() string`

GetVlanVlanId returns the VlanVlanId field if non-nil, zero value otherwise.

### GetVlanVlanIdOk

`func (o *DataInnerIpamNetworkData) GetVlanVlanIdOk() (*string, bool)`

GetVlanVlanIdOk returns a tuple with the VlanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVlanId

`func (o *DataInnerIpamNetworkData) SetVlanVlanId(v string)`

SetVlanVlanId sets VlanVlanId field to given value.

### HasVlanVlanId

`func (o *DataInnerIpamNetworkData) HasVlanVlanId() bool`

HasVlanVlanId returns a boolean if a field has been set.

### GetVlsmBlockId

`func (o *DataInnerIpamNetworkData) GetVlsmBlockId() string`

GetVlsmBlockId returns the VlsmBlockId field if non-nil, zero value otherwise.

### GetVlsmBlockIdOk

`func (o *DataInnerIpamNetworkData) GetVlsmBlockIdOk() (*string, bool)`

GetVlsmBlockIdOk returns a tuple with the VlsmBlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmBlockId

`func (o *DataInnerIpamNetworkData) SetVlsmBlockId(v string)`

SetVlsmBlockId sets VlsmBlockId field to given value.

### HasVlsmBlockId

`func (o *DataInnerIpamNetworkData) HasVlsmBlockId() bool`

HasVlsmBlockId returns a boolean if a field has been set.

### GetVlsmSpaceId

`func (o *DataInnerIpamNetworkData) GetVlsmSpaceId() string`

GetVlsmSpaceId returns the VlsmSpaceId field if non-nil, zero value otherwise.

### GetVlsmSpaceIdOk

`func (o *DataInnerIpamNetworkData) GetVlsmSpaceIdOk() (*string, bool)`

GetVlsmSpaceIdOk returns a tuple with the VlsmSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceId

`func (o *DataInnerIpamNetworkData) SetVlsmSpaceId(v string)`

SetVlsmSpaceId sets VlsmSpaceId field to given value.

### HasVlsmSpaceId

`func (o *DataInnerIpamNetworkData) HasVlsmSpaceId() bool`

HasVlsmSpaceId returns a boolean if a field has been set.

### GetVlsmSpaceName

`func (o *DataInnerIpamNetworkData) GetVlsmSpaceName() string`

GetVlsmSpaceName returns the VlsmSpaceName field if non-nil, zero value otherwise.

### GetVlsmSpaceNameOk

`func (o *DataInnerIpamNetworkData) GetVlsmSpaceNameOk() (*string, bool)`

GetVlsmSpaceNameOk returns a tuple with the VlsmSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceName

`func (o *DataInnerIpamNetworkData) SetVlsmSpaceName(v string)`

SetVlsmSpaceName sets VlsmSpaceName field to given value.

### HasVlsmSpaceName

`func (o *DataInnerIpamNetworkData) HasVlsmSpaceName() bool`

HasVlsmSpaceName returns a boolean if a field has been set.

### GetVlsmNetworkId

`func (o *DataInnerIpamNetworkData) GetVlsmNetworkId() string`

GetVlsmNetworkId returns the VlsmNetworkId field if non-nil, zero value otherwise.

### GetVlsmNetworkIdOk

`func (o *DataInnerIpamNetworkData) GetVlsmNetworkIdOk() (*string, bool)`

GetVlsmNetworkIdOk returns a tuple with the VlsmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmNetworkId

`func (o *DataInnerIpamNetworkData) SetVlsmNetworkId(v string)`

SetVlsmNetworkId sets VlsmNetworkId field to given value.

### HasVlsmNetworkId

`func (o *DataInnerIpamNetworkData) HasVlsmNetworkId() bool`

HasVlsmNetworkId returns a boolean if a field has been set.

### GetNetworkRipeWaitingState

`func (o *DataInnerIpamNetworkData) GetNetworkRipeWaitingState() string`

GetNetworkRipeWaitingState returns the NetworkRipeWaitingState field if non-nil, zero value otherwise.

### GetNetworkRipeWaitingStateOk

`func (o *DataInnerIpamNetworkData) GetNetworkRipeWaitingStateOk() (*string, bool)`

GetNetworkRipeWaitingStateOk returns a tuple with the NetworkRipeWaitingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRipeWaitingState

`func (o *DataInnerIpamNetworkData) SetNetworkRipeWaitingState(v string)`

SetNetworkRipeWaitingState sets NetworkRipeWaitingState field to given value.

### HasNetworkRipeWaitingState

`func (o *DataInnerIpamNetworkData) HasNetworkRipeWaitingState() bool`

HasNetworkRipeWaitingState returns a boolean if a field has been set.

### GetNetworkRipeWaitingStatus

`func (o *DataInnerIpamNetworkData) GetNetworkRipeWaitingStatus() string`

GetNetworkRipeWaitingStatus returns the NetworkRipeWaitingStatus field if non-nil, zero value otherwise.

### GetNetworkRipeWaitingStatusOk

`func (o *DataInnerIpamNetworkData) GetNetworkRipeWaitingStatusOk() (*string, bool)`

GetNetworkRipeWaitingStatusOk returns a tuple with the NetworkRipeWaitingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkRipeWaitingStatus

`func (o *DataInnerIpamNetworkData) SetNetworkRipeWaitingStatus(v string)`

SetNetworkRipeWaitingStatus sets NetworkRipeWaitingStatus field to given value.

### HasNetworkRipeWaitingStatus

`func (o *DataInnerIpamNetworkData) HasNetworkRipeWaitingStatus() bool`

HasNetworkRipeWaitingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


