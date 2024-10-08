# DataInnerIpamNetwork6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network6EndHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;end_address6_addr&lt;/b&gt;. | [optional] 
**EndAddress6Addr** | Pointer to **string** | The last IP address of the IPv6 network, in hexadecimal format. | [optional] 
**Network6IsInOrphan** | Pointer to **string** | A way to determine if the network has a parent (&lt;b&gt;0&lt;/b&gt;) or if it belongs to a container &lt;b&gt;Orphan networks&lt;/b&gt; (&lt;b&gt;1&lt;/b&gt;). | [optional] 
**Network6IsTerminal** | Pointer to **string** | A way to determine if a network can contain other networks. If set to &lt;b&gt;1&lt;/b&gt;, the network is terminal and cannot contain other subnet-type networks. Block-type networks are always set to &lt;b&gt;0&lt;/b&gt;. | [optional] 
**Network6LockNetworkBroadcast** | Pointer to **string** | A way to prevent &lt;b&gt;(1)&lt;/b&gt; users from assigning the broadcast IP address and network IP address of the network. | [optional] 
**Network6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentEndAddress6Addr** | Pointer to **string** | The last IP address of the parent IPv6 network, in hexadecimal format. | [optional] 
**ParentNetwork6IsTerminal** | Pointer to **string** | A way to determine if the parent network is terminal (&lt;b&gt;1&lt;/b&gt;) or non-terminal (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**ParentNetwork6PercentAllocated** | Pointer to **string** | The percentage of subnet-type networks the parent network contains. | [optional] 
**ParentNetwork6PercentUsed** | Pointer to **string** | The percentage of terminal networks the parent network contains. | [optional] 
**ParentSpaceId** | Pointer to **string** | The database identifier (ID) of the space where is located the parent network. &lt;b&gt;0&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentSpaceName** | Pointer to **string** | The name of the space where is located the parent network. &lt;b&gt;#&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentStartAddress6Addr** | Pointer to **string** | The first IP address of the parent IPv6 network, in hexadecimal format. | [optional] 
**ParentNetwork6ClassName** | Pointer to **string** | The name of the class applied to the parent IPv6 network, it can be preceded by the class directory. | [optional] 
**ParentNetwork6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the parent IPv6 network and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;...&lt;/b&gt; . | [optional] 
**ParentNetwork6Id** | Pointer to **string** | The database identifier (ID) of the parent IPv6 network. &lt;b&gt;0&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentNetwork6Name** | Pointer to **string** | The name of the parent IPv6 network. &lt;b&gt;#&lt;/b&gt; indicates that the network has no parent network. | [optional] 
**ParentNetwork6Prefix** | Pointer to **string** | The prefix of the parent of the IPv6 network the object belongs to. | [optional] 
**ParentNetworkLevel** | Pointer to **string** | The level of the parent network within the space. It returns values between &lt;b&gt;0&lt;/b&gt; (block-type network) and &lt;b&gt;n&lt;/b&gt; (subnet-type network). A value higher than &lt;b&gt;1&lt;/b&gt; indicates a VLSM organization where a block-type network can belong to another subnet-type network. | [optional] 
**ParentNetworkPath** | Pointer to **string** | The path toward the parent network in the database. &lt;b&gt;#&lt;/b&gt; indicates the network has no parent network. | [optional] 
**ParentNetworkSize** | Pointer to **string** | The number of IP addresses of the network parent, in hexadecimal format. | [optional] 
**ParentVlsmNetwork6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 subnet-type network, located in the VLSM parent space, from which the parent network was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the parent network is not a VLSM block-type network duplicated from a parent space. | [optional] 
**PercentAllocated** | Pointer to **string** | The percentage of subnet-type networks the non-terminal network contains. | [optional] 
**PercentUsed** | Pointer to **string** | The percentage of terminal networks the non-terminal network contains. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the space the object belongs to. | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space the object belongs to. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceIsTemplate** | Pointer to **string** | The template status of the space the object belongs to. If the space is used as template (1), all the IPv4 networks, pools and IP addresses it contains are also used as template. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**SpaceParentSpaceId** | Pointer to **string** | The database identifier (ID) of the VLSM parent of the space where is located the network. &lt;b&gt;0&lt;/b&gt; indicates that the space where is located the network has no parent space. | [optional] 
**Network6StartHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;start_address6_addr&lt;/b&gt;. | [optional] 
**StartAddress6Addr** | Pointer to **string** | The first IP address of the IPv6 network, in hexadecimal format. | [optional] 
**Network6ClassName** | Pointer to **string** | The name of the class applied to the IPv6 network, it can be preceded by the class directory. | [optional] 
**Network6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv6 network. | [optional] 
**Network6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 network. | [optional] 
**Network6IsValid** | Pointer to **string** | The network validity. A valid network (&lt;b&gt;1&lt;/b&gt;) has a prefix and last IP address that match. | [optional] 
**Network6Name** | Pointer to **string** | The name of the IPv6 network. | [optional] 
**Network6Prefix** | Pointer to **string** | The prefix of the IPv6 network. | [optional] 
**NetworkLevel** | Pointer to **string** | The level of the network within the space. It returns values between &lt;b&gt;0&lt;/b&gt; (block-type network) and &lt;b&gt;n&lt;/b&gt; (subnet-type network). A value higher than &lt;b&gt;1&lt;/b&gt; indicates a VLSM organization where a block-type network can belong to another subnet-type network. | [optional] 
**NetworkPath** | Pointer to **string** | The path toward the network in the database from the containing block-type network down to the subnet-type network: &lt;b&gt;&amp;lt;block-network-start-IP&amp;gt;#&amp;lt;block-network-ID&amp;gt;#&amp;lt;subnet-network-start-IP&amp;gt;#&amp;lt;subnet-network-ID&amp;gt;&lt;/b&gt;. The IP address is returned in hexadecimal format.&lt;ul class&#x3D;dashed &gt;&lt;li&gt; In network-based VLSM organizations, the path includes all the subnet-type networks there are from the containing block-type network down to the subnet-type network specified in &lt;b&gt;subnet_id&lt;/b&gt;.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; In space-based VLSM organizations, the path includes the block-type network of the top parent space and all the subnet-type networks there are until the network specified in &lt;b&gt;subnet_id&lt;/b&gt;. Only one block-type network is returned.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**NetworkSize** | Pointer to **string** | The number of IP addresses the IPv6 network contains. | [optional] 
**Network6OperationDetailsAddedOn** | Pointer to **string** | The creation date of the IPv6 network, in decimal UNIX date format. | [optional] 
**Network6OperationDetailsCallStack** | Pointer to **string** | The call stack of the IPv6 network operation details, as follows: &lt;b&gt;&amp;lt;service1&amp;gt;&amp;amp;&amp;lt;service2&amp;gt;&amp;amp;&amp;lt;service3&amp;gt;&lt;/b&gt;... . | [optional] 
**Network6OperationDetailsOriginModule** | Pointer to **string** | The name of the module where the IPv6 network addition originated. | [optional] 
**Network6OperationDetailsRequestedBy** | Pointer to **string** | The login of the user who requested the IPv6 network. | [optional] 
**Network6OperationDetailsAddedBy** | Pointer to **string** | The login of the user who added the IPv6 network. | [optional] 
**Network6OperationDetailsLastUpdatedOn** | Pointer to **string** | The last time the IPv6 network was updated, in decimal UNIX date format. | [optional] 
**DomainId** | Pointer to **string** | The database identifier (ID) of the VLAN domain associated with the network. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain associated with the network. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the VLAN range associated with the network. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range associated with the network. | [optional] 
**VlanId** | Pointer to **string** | The database identifier (ID) of the VLAN associated with the network. | [optional] 
**VlanName** | Pointer to **string** | The name of the VLAN associated with the network. | [optional] 
**VlanVlanId** | Pointer to **string** | The VLAN identifier (ID) of the VLAN associated with the network. | [optional] 
**VlsmBlock6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 VLSM block-type network duplicated, in a VLSM child space, from the network. &lt;b&gt;0&lt;/b&gt; indicates that the network is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmSpaceId** | Pointer to **string** | The database identifier (ID) of the VLSM child space where the network is duplicated as a VLSM block-type network. &lt;b&gt;0&lt;/b&gt; indicates that the network is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmSpaceName** | Pointer to **string** | The name of the VLSM child space where the network is duplicated as a VLSM block-type network. &lt;b&gt;0&lt;/b&gt; indicates that the network is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmNetwork6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 subnet-type network, located in the VLSM parent space, from which the network was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the network is not a VLSM block-type network duplicated from a parent space. | [optional] 
**Network6RipeWaitingState** | Pointer to **string** | The state of the exchange between SOLIDserver and the RIPE for the assigned network: &lt;table&gt;&lt;caption&gt;network6_ripe_waiting_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;must_send_mail_add&lt;/td&gt;&lt;td &gt;An email must be sent to the RIPE to notify them of a subnet-type network creation.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_mail_add&lt;/td&gt;&lt;td &gt;A network creation email was sent to the RIPE, no reply has been received yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;must_send_mail_del&lt;/td&gt;&lt;td &gt;An email must be sent to the RIPE to notify them of a subnet-type network deletion.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_mail_del&lt;/td&gt;&lt;td &gt;A network deletion email was sent to the RIPE, no reply has been received yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;wait_aw_confirm&lt;/td&gt;&lt;td &gt;The number of IP addresses of the assigned network exceeds the Assignment Window declared during your RIPE configuration.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Network6RipeWaitingStatus** | Pointer to **string** | The status of a RIPE assigned network within SOLIDserver until it is confirmed that you can create or delete it. If set to &lt;b&gt;1&lt;/b&gt;, it is about to be created. If set to &lt;b&gt;2&lt;/b&gt;, it is about to be deleted. | [optional] 

## Methods

### NewDataInnerIpamNetwork6Data

`func NewDataInnerIpamNetwork6Data() *DataInnerIpamNetwork6Data`

NewDataInnerIpamNetwork6Data instantiates a new DataInnerIpamNetwork6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerIpamNetwork6DataWithDefaults

`func NewDataInnerIpamNetwork6DataWithDefaults() *DataInnerIpamNetwork6Data`

NewDataInnerIpamNetwork6DataWithDefaults instantiates a new DataInnerIpamNetwork6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork6EndHostaddr

`func (o *DataInnerIpamNetwork6Data) GetNetwork6EndHostaddr() string`

GetNetwork6EndHostaddr returns the Network6EndHostaddr field if non-nil, zero value otherwise.

### GetNetwork6EndHostaddrOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6EndHostaddrOk() (*string, bool)`

GetNetwork6EndHostaddrOk returns a tuple with the Network6EndHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6EndHostaddr

`func (o *DataInnerIpamNetwork6Data) SetNetwork6EndHostaddr(v string)`

SetNetwork6EndHostaddr sets Network6EndHostaddr field to given value.

### HasNetwork6EndHostaddr

`func (o *DataInnerIpamNetwork6Data) HasNetwork6EndHostaddr() bool`

HasNetwork6EndHostaddr returns a boolean if a field has been set.

### GetEndAddress6Addr

`func (o *DataInnerIpamNetwork6Data) GetEndAddress6Addr() string`

GetEndAddress6Addr returns the EndAddress6Addr field if non-nil, zero value otherwise.

### GetEndAddress6AddrOk

`func (o *DataInnerIpamNetwork6Data) GetEndAddress6AddrOk() (*string, bool)`

GetEndAddress6AddrOk returns a tuple with the EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress6Addr

`func (o *DataInnerIpamNetwork6Data) SetEndAddress6Addr(v string)`

SetEndAddress6Addr sets EndAddress6Addr field to given value.

### HasEndAddress6Addr

`func (o *DataInnerIpamNetwork6Data) HasEndAddress6Addr() bool`

HasEndAddress6Addr returns a boolean if a field has been set.

### GetNetwork6IsInOrphan

`func (o *DataInnerIpamNetwork6Data) GetNetwork6IsInOrphan() string`

GetNetwork6IsInOrphan returns the Network6IsInOrphan field if non-nil, zero value otherwise.

### GetNetwork6IsInOrphanOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6IsInOrphanOk() (*string, bool)`

GetNetwork6IsInOrphanOk returns a tuple with the Network6IsInOrphan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6IsInOrphan

`func (o *DataInnerIpamNetwork6Data) SetNetwork6IsInOrphan(v string)`

SetNetwork6IsInOrphan sets Network6IsInOrphan field to given value.

### HasNetwork6IsInOrphan

`func (o *DataInnerIpamNetwork6Data) HasNetwork6IsInOrphan() bool`

HasNetwork6IsInOrphan returns a boolean if a field has been set.

### GetNetwork6IsTerminal

`func (o *DataInnerIpamNetwork6Data) GetNetwork6IsTerminal() string`

GetNetwork6IsTerminal returns the Network6IsTerminal field if non-nil, zero value otherwise.

### GetNetwork6IsTerminalOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6IsTerminalOk() (*string, bool)`

GetNetwork6IsTerminalOk returns a tuple with the Network6IsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6IsTerminal

`func (o *DataInnerIpamNetwork6Data) SetNetwork6IsTerminal(v string)`

SetNetwork6IsTerminal sets Network6IsTerminal field to given value.

### HasNetwork6IsTerminal

`func (o *DataInnerIpamNetwork6Data) HasNetwork6IsTerminal() bool`

HasNetwork6IsTerminal returns a boolean if a field has been set.

### GetNetwork6LockNetworkBroadcast

`func (o *DataInnerIpamNetwork6Data) GetNetwork6LockNetworkBroadcast() string`

GetNetwork6LockNetworkBroadcast returns the Network6LockNetworkBroadcast field if non-nil, zero value otherwise.

### GetNetwork6LockNetworkBroadcastOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6LockNetworkBroadcastOk() (*string, bool)`

GetNetwork6LockNetworkBroadcastOk returns a tuple with the Network6LockNetworkBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6LockNetworkBroadcast

`func (o *DataInnerIpamNetwork6Data) SetNetwork6LockNetworkBroadcast(v string)`

SetNetwork6LockNetworkBroadcast sets Network6LockNetworkBroadcast field to given value.

### HasNetwork6LockNetworkBroadcast

`func (o *DataInnerIpamNetwork6Data) HasNetwork6LockNetworkBroadcast() bool`

HasNetwork6LockNetworkBroadcast returns a boolean if a field has been set.

### GetNetwork6Multistatus

`func (o *DataInnerIpamNetwork6Data) GetNetwork6Multistatus() string`

GetNetwork6Multistatus returns the Network6Multistatus field if non-nil, zero value otherwise.

### GetNetwork6MultistatusOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6MultistatusOk() (*string, bool)`

GetNetwork6MultistatusOk returns a tuple with the Network6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Multistatus

`func (o *DataInnerIpamNetwork6Data) SetNetwork6Multistatus(v string)`

SetNetwork6Multistatus sets Network6Multistatus field to given value.

### HasNetwork6Multistatus

`func (o *DataInnerIpamNetwork6Data) HasNetwork6Multistatus() bool`

HasNetwork6Multistatus returns a boolean if a field has been set.

### GetParentEndAddress6Addr

`func (o *DataInnerIpamNetwork6Data) GetParentEndAddress6Addr() string`

GetParentEndAddress6Addr returns the ParentEndAddress6Addr field if non-nil, zero value otherwise.

### GetParentEndAddress6AddrOk

`func (o *DataInnerIpamNetwork6Data) GetParentEndAddress6AddrOk() (*string, bool)`

GetParentEndAddress6AddrOk returns a tuple with the ParentEndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentEndAddress6Addr

`func (o *DataInnerIpamNetwork6Data) SetParentEndAddress6Addr(v string)`

SetParentEndAddress6Addr sets ParentEndAddress6Addr field to given value.

### HasParentEndAddress6Addr

`func (o *DataInnerIpamNetwork6Data) HasParentEndAddress6Addr() bool`

HasParentEndAddress6Addr returns a boolean if a field has been set.

### GetParentNetwork6IsTerminal

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6IsTerminal() string`

GetParentNetwork6IsTerminal returns the ParentNetwork6IsTerminal field if non-nil, zero value otherwise.

### GetParentNetwork6IsTerminalOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6IsTerminalOk() (*string, bool)`

GetParentNetwork6IsTerminalOk returns a tuple with the ParentNetwork6IsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6IsTerminal

`func (o *DataInnerIpamNetwork6Data) SetParentNetwork6IsTerminal(v string)`

SetParentNetwork6IsTerminal sets ParentNetwork6IsTerminal field to given value.

### HasParentNetwork6IsTerminal

`func (o *DataInnerIpamNetwork6Data) HasParentNetwork6IsTerminal() bool`

HasParentNetwork6IsTerminal returns a boolean if a field has been set.

### GetParentNetwork6PercentAllocated

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6PercentAllocated() string`

GetParentNetwork6PercentAllocated returns the ParentNetwork6PercentAllocated field if non-nil, zero value otherwise.

### GetParentNetwork6PercentAllocatedOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6PercentAllocatedOk() (*string, bool)`

GetParentNetwork6PercentAllocatedOk returns a tuple with the ParentNetwork6PercentAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6PercentAllocated

`func (o *DataInnerIpamNetwork6Data) SetParentNetwork6PercentAllocated(v string)`

SetParentNetwork6PercentAllocated sets ParentNetwork6PercentAllocated field to given value.

### HasParentNetwork6PercentAllocated

`func (o *DataInnerIpamNetwork6Data) HasParentNetwork6PercentAllocated() bool`

HasParentNetwork6PercentAllocated returns a boolean if a field has been set.

### GetParentNetwork6PercentUsed

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6PercentUsed() string`

GetParentNetwork6PercentUsed returns the ParentNetwork6PercentUsed field if non-nil, zero value otherwise.

### GetParentNetwork6PercentUsedOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6PercentUsedOk() (*string, bool)`

GetParentNetwork6PercentUsedOk returns a tuple with the ParentNetwork6PercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6PercentUsed

`func (o *DataInnerIpamNetwork6Data) SetParentNetwork6PercentUsed(v string)`

SetParentNetwork6PercentUsed sets ParentNetwork6PercentUsed field to given value.

### HasParentNetwork6PercentUsed

`func (o *DataInnerIpamNetwork6Data) HasParentNetwork6PercentUsed() bool`

HasParentNetwork6PercentUsed returns a boolean if a field has been set.

### GetParentSpaceId

`func (o *DataInnerIpamNetwork6Data) GetParentSpaceId() string`

GetParentSpaceId returns the ParentSpaceId field if non-nil, zero value otherwise.

### GetParentSpaceIdOk

`func (o *DataInnerIpamNetwork6Data) GetParentSpaceIdOk() (*string, bool)`

GetParentSpaceIdOk returns a tuple with the ParentSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceId

`func (o *DataInnerIpamNetwork6Data) SetParentSpaceId(v string)`

SetParentSpaceId sets ParentSpaceId field to given value.

### HasParentSpaceId

`func (o *DataInnerIpamNetwork6Data) HasParentSpaceId() bool`

HasParentSpaceId returns a boolean if a field has been set.

### GetParentSpaceName

`func (o *DataInnerIpamNetwork6Data) GetParentSpaceName() string`

GetParentSpaceName returns the ParentSpaceName field if non-nil, zero value otherwise.

### GetParentSpaceNameOk

`func (o *DataInnerIpamNetwork6Data) GetParentSpaceNameOk() (*string, bool)`

GetParentSpaceNameOk returns a tuple with the ParentSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceName

`func (o *DataInnerIpamNetwork6Data) SetParentSpaceName(v string)`

SetParentSpaceName sets ParentSpaceName field to given value.

### HasParentSpaceName

`func (o *DataInnerIpamNetwork6Data) HasParentSpaceName() bool`

HasParentSpaceName returns a boolean if a field has been set.

### GetParentStartAddress6Addr

`func (o *DataInnerIpamNetwork6Data) GetParentStartAddress6Addr() string`

GetParentStartAddress6Addr returns the ParentStartAddress6Addr field if non-nil, zero value otherwise.

### GetParentStartAddress6AddrOk

`func (o *DataInnerIpamNetwork6Data) GetParentStartAddress6AddrOk() (*string, bool)`

GetParentStartAddress6AddrOk returns a tuple with the ParentStartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentStartAddress6Addr

`func (o *DataInnerIpamNetwork6Data) SetParentStartAddress6Addr(v string)`

SetParentStartAddress6Addr sets ParentStartAddress6Addr field to given value.

### HasParentStartAddress6Addr

`func (o *DataInnerIpamNetwork6Data) HasParentStartAddress6Addr() bool`

HasParentStartAddress6Addr returns a boolean if a field has been set.

### GetParentNetwork6ClassName

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6ClassName() string`

GetParentNetwork6ClassName returns the ParentNetwork6ClassName field if non-nil, zero value otherwise.

### GetParentNetwork6ClassNameOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6ClassNameOk() (*string, bool)`

GetParentNetwork6ClassNameOk returns a tuple with the ParentNetwork6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6ClassName

`func (o *DataInnerIpamNetwork6Data) SetParentNetwork6ClassName(v string)`

SetParentNetwork6ClassName sets ParentNetwork6ClassName field to given value.

### HasParentNetwork6ClassName

`func (o *DataInnerIpamNetwork6Data) HasParentNetwork6ClassName() bool`

HasParentNetwork6ClassName returns a boolean if a field has been set.

### GetParentNetwork6ClassParameters

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6ClassParameters() []ApiClassParameterOutputEntry`

GetParentNetwork6ClassParameters returns the ParentNetwork6ClassParameters field if non-nil, zero value otherwise.

### GetParentNetwork6ClassParametersOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetParentNetwork6ClassParametersOk returns a tuple with the ParentNetwork6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6ClassParameters

`func (o *DataInnerIpamNetwork6Data) SetParentNetwork6ClassParameters(v []ApiClassParameterOutputEntry)`

SetParentNetwork6ClassParameters sets ParentNetwork6ClassParameters field to given value.

### HasParentNetwork6ClassParameters

`func (o *DataInnerIpamNetwork6Data) HasParentNetwork6ClassParameters() bool`

HasParentNetwork6ClassParameters returns a boolean if a field has been set.

### GetParentNetwork6Id

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6Id() string`

GetParentNetwork6Id returns the ParentNetwork6Id field if non-nil, zero value otherwise.

### GetParentNetwork6IdOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6IdOk() (*string, bool)`

GetParentNetwork6IdOk returns a tuple with the ParentNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Id

`func (o *DataInnerIpamNetwork6Data) SetParentNetwork6Id(v string)`

SetParentNetwork6Id sets ParentNetwork6Id field to given value.

### HasParentNetwork6Id

`func (o *DataInnerIpamNetwork6Data) HasParentNetwork6Id() bool`

HasParentNetwork6Id returns a boolean if a field has been set.

### GetParentNetwork6Name

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6Name() string`

GetParentNetwork6Name returns the ParentNetwork6Name field if non-nil, zero value otherwise.

### GetParentNetwork6NameOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6NameOk() (*string, bool)`

GetParentNetwork6NameOk returns a tuple with the ParentNetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Name

`func (o *DataInnerIpamNetwork6Data) SetParentNetwork6Name(v string)`

SetParentNetwork6Name sets ParentNetwork6Name field to given value.

### HasParentNetwork6Name

`func (o *DataInnerIpamNetwork6Data) HasParentNetwork6Name() bool`

HasParentNetwork6Name returns a boolean if a field has been set.

### GetParentNetwork6Prefix

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6Prefix() string`

GetParentNetwork6Prefix returns the ParentNetwork6Prefix field if non-nil, zero value otherwise.

### GetParentNetwork6PrefixOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetwork6PrefixOk() (*string, bool)`

GetParentNetwork6PrefixOk returns a tuple with the ParentNetwork6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Prefix

`func (o *DataInnerIpamNetwork6Data) SetParentNetwork6Prefix(v string)`

SetParentNetwork6Prefix sets ParentNetwork6Prefix field to given value.

### HasParentNetwork6Prefix

`func (o *DataInnerIpamNetwork6Data) HasParentNetwork6Prefix() bool`

HasParentNetwork6Prefix returns a boolean if a field has been set.

### GetParentNetworkLevel

`func (o *DataInnerIpamNetwork6Data) GetParentNetworkLevel() string`

GetParentNetworkLevel returns the ParentNetworkLevel field if non-nil, zero value otherwise.

### GetParentNetworkLevelOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetworkLevelOk() (*string, bool)`

GetParentNetworkLevelOk returns a tuple with the ParentNetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkLevel

`func (o *DataInnerIpamNetwork6Data) SetParentNetworkLevel(v string)`

SetParentNetworkLevel sets ParentNetworkLevel field to given value.

### HasParentNetworkLevel

`func (o *DataInnerIpamNetwork6Data) HasParentNetworkLevel() bool`

HasParentNetworkLevel returns a boolean if a field has been set.

### GetParentNetworkPath

`func (o *DataInnerIpamNetwork6Data) GetParentNetworkPath() string`

GetParentNetworkPath returns the ParentNetworkPath field if non-nil, zero value otherwise.

### GetParentNetworkPathOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetworkPathOk() (*string, bool)`

GetParentNetworkPathOk returns a tuple with the ParentNetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkPath

`func (o *DataInnerIpamNetwork6Data) SetParentNetworkPath(v string)`

SetParentNetworkPath sets ParentNetworkPath field to given value.

### HasParentNetworkPath

`func (o *DataInnerIpamNetwork6Data) HasParentNetworkPath() bool`

HasParentNetworkPath returns a boolean if a field has been set.

### GetParentNetworkSize

`func (o *DataInnerIpamNetwork6Data) GetParentNetworkSize() string`

GetParentNetworkSize returns the ParentNetworkSize field if non-nil, zero value otherwise.

### GetParentNetworkSizeOk

`func (o *DataInnerIpamNetwork6Data) GetParentNetworkSizeOk() (*string, bool)`

GetParentNetworkSizeOk returns a tuple with the ParentNetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkSize

`func (o *DataInnerIpamNetwork6Data) SetParentNetworkSize(v string)`

SetParentNetworkSize sets ParentNetworkSize field to given value.

### HasParentNetworkSize

`func (o *DataInnerIpamNetwork6Data) HasParentNetworkSize() bool`

HasParentNetworkSize returns a boolean if a field has been set.

### GetParentVlsmNetwork6Id

`func (o *DataInnerIpamNetwork6Data) GetParentVlsmNetwork6Id() string`

GetParentVlsmNetwork6Id returns the ParentVlsmNetwork6Id field if non-nil, zero value otherwise.

### GetParentVlsmNetwork6IdOk

`func (o *DataInnerIpamNetwork6Data) GetParentVlsmNetwork6IdOk() (*string, bool)`

GetParentVlsmNetwork6IdOk returns a tuple with the ParentVlsmNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentVlsmNetwork6Id

`func (o *DataInnerIpamNetwork6Data) SetParentVlsmNetwork6Id(v string)`

SetParentVlsmNetwork6Id sets ParentVlsmNetwork6Id field to given value.

### HasParentVlsmNetwork6Id

`func (o *DataInnerIpamNetwork6Data) HasParentVlsmNetwork6Id() bool`

HasParentVlsmNetwork6Id returns a boolean if a field has been set.

### GetPercentAllocated

`func (o *DataInnerIpamNetwork6Data) GetPercentAllocated() string`

GetPercentAllocated returns the PercentAllocated field if non-nil, zero value otherwise.

### GetPercentAllocatedOk

`func (o *DataInnerIpamNetwork6Data) GetPercentAllocatedOk() (*string, bool)`

GetPercentAllocatedOk returns a tuple with the PercentAllocated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentAllocated

`func (o *DataInnerIpamNetwork6Data) SetPercentAllocated(v string)`

SetPercentAllocated sets PercentAllocated field to given value.

### HasPercentAllocated

`func (o *DataInnerIpamNetwork6Data) HasPercentAllocated() bool`

HasPercentAllocated returns a boolean if a field has been set.

### GetPercentUsed

`func (o *DataInnerIpamNetwork6Data) GetPercentUsed() string`

GetPercentUsed returns the PercentUsed field if non-nil, zero value otherwise.

### GetPercentUsedOk

`func (o *DataInnerIpamNetwork6Data) GetPercentUsedOk() (*string, bool)`

GetPercentUsedOk returns a tuple with the PercentUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercentUsed

`func (o *DataInnerIpamNetwork6Data) SetPercentUsed(v string)`

SetPercentUsed sets PercentUsed field to given value.

### HasPercentUsed

`func (o *DataInnerIpamNetwork6Data) HasPercentUsed() bool`

HasPercentUsed returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *DataInnerIpamNetwork6Data) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *DataInnerIpamNetwork6Data) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *DataInnerIpamNetwork6Data) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *DataInnerIpamNetwork6Data) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *DataInnerIpamNetwork6Data) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *DataInnerIpamNetwork6Data) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *DataInnerIpamNetwork6Data) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *DataInnerIpamNetwork6Data) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *DataInnerIpamNetwork6Data) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *DataInnerIpamNetwork6Data) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *DataInnerIpamNetwork6Data) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *DataInnerIpamNetwork6Data) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *DataInnerIpamNetwork6Data) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DataInnerIpamNetwork6Data) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DataInnerIpamNetwork6Data) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DataInnerIpamNetwork6Data) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceIsTemplate

`func (o *DataInnerIpamNetwork6Data) GetSpaceIsTemplate() string`

GetSpaceIsTemplate returns the SpaceIsTemplate field if non-nil, zero value otherwise.

### GetSpaceIsTemplateOk

`func (o *DataInnerIpamNetwork6Data) GetSpaceIsTemplateOk() (*string, bool)`

GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceIsTemplate

`func (o *DataInnerIpamNetwork6Data) SetSpaceIsTemplate(v string)`

SetSpaceIsTemplate sets SpaceIsTemplate field to given value.

### HasSpaceIsTemplate

`func (o *DataInnerIpamNetwork6Data) HasSpaceIsTemplate() bool`

HasSpaceIsTemplate returns a boolean if a field has been set.

### GetSpaceName

`func (o *DataInnerIpamNetwork6Data) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *DataInnerIpamNetwork6Data) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *DataInnerIpamNetwork6Data) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *DataInnerIpamNetwork6Data) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetSpaceParentSpaceId

`func (o *DataInnerIpamNetwork6Data) GetSpaceParentSpaceId() string`

GetSpaceParentSpaceId returns the SpaceParentSpaceId field if non-nil, zero value otherwise.

### GetSpaceParentSpaceIdOk

`func (o *DataInnerIpamNetwork6Data) GetSpaceParentSpaceIdOk() (*string, bool)`

GetSpaceParentSpaceIdOk returns a tuple with the SpaceParentSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceParentSpaceId

`func (o *DataInnerIpamNetwork6Data) SetSpaceParentSpaceId(v string)`

SetSpaceParentSpaceId sets SpaceParentSpaceId field to given value.

### HasSpaceParentSpaceId

`func (o *DataInnerIpamNetwork6Data) HasSpaceParentSpaceId() bool`

HasSpaceParentSpaceId returns a boolean if a field has been set.

### GetNetwork6StartHostaddr

`func (o *DataInnerIpamNetwork6Data) GetNetwork6StartHostaddr() string`

GetNetwork6StartHostaddr returns the Network6StartHostaddr field if non-nil, zero value otherwise.

### GetNetwork6StartHostaddrOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6StartHostaddrOk() (*string, bool)`

GetNetwork6StartHostaddrOk returns a tuple with the Network6StartHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6StartHostaddr

`func (o *DataInnerIpamNetwork6Data) SetNetwork6StartHostaddr(v string)`

SetNetwork6StartHostaddr sets Network6StartHostaddr field to given value.

### HasNetwork6StartHostaddr

`func (o *DataInnerIpamNetwork6Data) HasNetwork6StartHostaddr() bool`

HasNetwork6StartHostaddr returns a boolean if a field has been set.

### GetStartAddress6Addr

`func (o *DataInnerIpamNetwork6Data) GetStartAddress6Addr() string`

GetStartAddress6Addr returns the StartAddress6Addr field if non-nil, zero value otherwise.

### GetStartAddress6AddrOk

`func (o *DataInnerIpamNetwork6Data) GetStartAddress6AddrOk() (*string, bool)`

GetStartAddress6AddrOk returns a tuple with the StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress6Addr

`func (o *DataInnerIpamNetwork6Data) SetStartAddress6Addr(v string)`

SetStartAddress6Addr sets StartAddress6Addr field to given value.

### HasStartAddress6Addr

`func (o *DataInnerIpamNetwork6Data) HasStartAddress6Addr() bool`

HasStartAddress6Addr returns a boolean if a field has been set.

### GetNetwork6ClassName

`func (o *DataInnerIpamNetwork6Data) GetNetwork6ClassName() string`

GetNetwork6ClassName returns the Network6ClassName field if non-nil, zero value otherwise.

### GetNetwork6ClassNameOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6ClassNameOk() (*string, bool)`

GetNetwork6ClassNameOk returns a tuple with the Network6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6ClassName

`func (o *DataInnerIpamNetwork6Data) SetNetwork6ClassName(v string)`

SetNetwork6ClassName sets Network6ClassName field to given value.

### HasNetwork6ClassName

`func (o *DataInnerIpamNetwork6Data) HasNetwork6ClassName() bool`

HasNetwork6ClassName returns a boolean if a field has been set.

### GetNetwork6ClassParameters

`func (o *DataInnerIpamNetwork6Data) GetNetwork6ClassParameters() []ApiClassParameterOutputEntry`

GetNetwork6ClassParameters returns the Network6ClassParameters field if non-nil, zero value otherwise.

### GetNetwork6ClassParametersOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetNetwork6ClassParametersOk returns a tuple with the Network6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6ClassParameters

`func (o *DataInnerIpamNetwork6Data) SetNetwork6ClassParameters(v []ApiClassParameterOutputEntry)`

SetNetwork6ClassParameters sets Network6ClassParameters field to given value.

### HasNetwork6ClassParameters

`func (o *DataInnerIpamNetwork6Data) HasNetwork6ClassParameters() bool`

HasNetwork6ClassParameters returns a boolean if a field has been set.

### GetNetwork6Id

`func (o *DataInnerIpamNetwork6Data) GetNetwork6Id() string`

GetNetwork6Id returns the Network6Id field if non-nil, zero value otherwise.

### GetNetwork6IdOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6IdOk() (*string, bool)`

GetNetwork6IdOk returns a tuple with the Network6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Id

`func (o *DataInnerIpamNetwork6Data) SetNetwork6Id(v string)`

SetNetwork6Id sets Network6Id field to given value.

### HasNetwork6Id

`func (o *DataInnerIpamNetwork6Data) HasNetwork6Id() bool`

HasNetwork6Id returns a boolean if a field has been set.

### GetNetwork6IsValid

`func (o *DataInnerIpamNetwork6Data) GetNetwork6IsValid() string`

GetNetwork6IsValid returns the Network6IsValid field if non-nil, zero value otherwise.

### GetNetwork6IsValidOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6IsValidOk() (*string, bool)`

GetNetwork6IsValidOk returns a tuple with the Network6IsValid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6IsValid

`func (o *DataInnerIpamNetwork6Data) SetNetwork6IsValid(v string)`

SetNetwork6IsValid sets Network6IsValid field to given value.

### HasNetwork6IsValid

`func (o *DataInnerIpamNetwork6Data) HasNetwork6IsValid() bool`

HasNetwork6IsValid returns a boolean if a field has been set.

### GetNetwork6Name

`func (o *DataInnerIpamNetwork6Data) GetNetwork6Name() string`

GetNetwork6Name returns the Network6Name field if non-nil, zero value otherwise.

### GetNetwork6NameOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6NameOk() (*string, bool)`

GetNetwork6NameOk returns a tuple with the Network6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Name

`func (o *DataInnerIpamNetwork6Data) SetNetwork6Name(v string)`

SetNetwork6Name sets Network6Name field to given value.

### HasNetwork6Name

`func (o *DataInnerIpamNetwork6Data) HasNetwork6Name() bool`

HasNetwork6Name returns a boolean if a field has been set.

### GetNetwork6Prefix

`func (o *DataInnerIpamNetwork6Data) GetNetwork6Prefix() string`

GetNetwork6Prefix returns the Network6Prefix field if non-nil, zero value otherwise.

### GetNetwork6PrefixOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6PrefixOk() (*string, bool)`

GetNetwork6PrefixOk returns a tuple with the Network6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Prefix

`func (o *DataInnerIpamNetwork6Data) SetNetwork6Prefix(v string)`

SetNetwork6Prefix sets Network6Prefix field to given value.

### HasNetwork6Prefix

`func (o *DataInnerIpamNetwork6Data) HasNetwork6Prefix() bool`

HasNetwork6Prefix returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *DataInnerIpamNetwork6Data) GetNetworkLevel() string`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *DataInnerIpamNetwork6Data) GetNetworkLevelOk() (*string, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *DataInnerIpamNetwork6Data) SetNetworkLevel(v string)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *DataInnerIpamNetwork6Data) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetNetworkPath

`func (o *DataInnerIpamNetwork6Data) GetNetworkPath() string`

GetNetworkPath returns the NetworkPath field if non-nil, zero value otherwise.

### GetNetworkPathOk

`func (o *DataInnerIpamNetwork6Data) GetNetworkPathOk() (*string, bool)`

GetNetworkPathOk returns a tuple with the NetworkPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPath

`func (o *DataInnerIpamNetwork6Data) SetNetworkPath(v string)`

SetNetworkPath sets NetworkPath field to given value.

### HasNetworkPath

`func (o *DataInnerIpamNetwork6Data) HasNetworkPath() bool`

HasNetworkPath returns a boolean if a field has been set.

### GetNetworkSize

`func (o *DataInnerIpamNetwork6Data) GetNetworkSize() string`

GetNetworkSize returns the NetworkSize field if non-nil, zero value otherwise.

### GetNetworkSizeOk

`func (o *DataInnerIpamNetwork6Data) GetNetworkSizeOk() (*string, bool)`

GetNetworkSizeOk returns a tuple with the NetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSize

`func (o *DataInnerIpamNetwork6Data) SetNetworkSize(v string)`

SetNetworkSize sets NetworkSize field to given value.

### HasNetworkSize

`func (o *DataInnerIpamNetwork6Data) HasNetworkSize() bool`

HasNetworkSize returns a boolean if a field has been set.

### GetNetwork6OperationDetailsAddedOn

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsAddedOn() string`

GetNetwork6OperationDetailsAddedOn returns the Network6OperationDetailsAddedOn field if non-nil, zero value otherwise.

### GetNetwork6OperationDetailsAddedOnOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsAddedOnOk() (*string, bool)`

GetNetwork6OperationDetailsAddedOnOk returns a tuple with the Network6OperationDetailsAddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6OperationDetailsAddedOn

`func (o *DataInnerIpamNetwork6Data) SetNetwork6OperationDetailsAddedOn(v string)`

SetNetwork6OperationDetailsAddedOn sets Network6OperationDetailsAddedOn field to given value.

### HasNetwork6OperationDetailsAddedOn

`func (o *DataInnerIpamNetwork6Data) HasNetwork6OperationDetailsAddedOn() bool`

HasNetwork6OperationDetailsAddedOn returns a boolean if a field has been set.

### GetNetwork6OperationDetailsCallStack

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsCallStack() string`

GetNetwork6OperationDetailsCallStack returns the Network6OperationDetailsCallStack field if non-nil, zero value otherwise.

### GetNetwork6OperationDetailsCallStackOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsCallStackOk() (*string, bool)`

GetNetwork6OperationDetailsCallStackOk returns a tuple with the Network6OperationDetailsCallStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6OperationDetailsCallStack

`func (o *DataInnerIpamNetwork6Data) SetNetwork6OperationDetailsCallStack(v string)`

SetNetwork6OperationDetailsCallStack sets Network6OperationDetailsCallStack field to given value.

### HasNetwork6OperationDetailsCallStack

`func (o *DataInnerIpamNetwork6Data) HasNetwork6OperationDetailsCallStack() bool`

HasNetwork6OperationDetailsCallStack returns a boolean if a field has been set.

### GetNetwork6OperationDetailsOriginModule

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsOriginModule() string`

GetNetwork6OperationDetailsOriginModule returns the Network6OperationDetailsOriginModule field if non-nil, zero value otherwise.

### GetNetwork6OperationDetailsOriginModuleOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsOriginModuleOk() (*string, bool)`

GetNetwork6OperationDetailsOriginModuleOk returns a tuple with the Network6OperationDetailsOriginModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6OperationDetailsOriginModule

`func (o *DataInnerIpamNetwork6Data) SetNetwork6OperationDetailsOriginModule(v string)`

SetNetwork6OperationDetailsOriginModule sets Network6OperationDetailsOriginModule field to given value.

### HasNetwork6OperationDetailsOriginModule

`func (o *DataInnerIpamNetwork6Data) HasNetwork6OperationDetailsOriginModule() bool`

HasNetwork6OperationDetailsOriginModule returns a boolean if a field has been set.

### GetNetwork6OperationDetailsRequestedBy

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsRequestedBy() string`

GetNetwork6OperationDetailsRequestedBy returns the Network6OperationDetailsRequestedBy field if non-nil, zero value otherwise.

### GetNetwork6OperationDetailsRequestedByOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsRequestedByOk() (*string, bool)`

GetNetwork6OperationDetailsRequestedByOk returns a tuple with the Network6OperationDetailsRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6OperationDetailsRequestedBy

`func (o *DataInnerIpamNetwork6Data) SetNetwork6OperationDetailsRequestedBy(v string)`

SetNetwork6OperationDetailsRequestedBy sets Network6OperationDetailsRequestedBy field to given value.

### HasNetwork6OperationDetailsRequestedBy

`func (o *DataInnerIpamNetwork6Data) HasNetwork6OperationDetailsRequestedBy() bool`

HasNetwork6OperationDetailsRequestedBy returns a boolean if a field has been set.

### GetNetwork6OperationDetailsAddedBy

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsAddedBy() string`

GetNetwork6OperationDetailsAddedBy returns the Network6OperationDetailsAddedBy field if non-nil, zero value otherwise.

### GetNetwork6OperationDetailsAddedByOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsAddedByOk() (*string, bool)`

GetNetwork6OperationDetailsAddedByOk returns a tuple with the Network6OperationDetailsAddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6OperationDetailsAddedBy

`func (o *DataInnerIpamNetwork6Data) SetNetwork6OperationDetailsAddedBy(v string)`

SetNetwork6OperationDetailsAddedBy sets Network6OperationDetailsAddedBy field to given value.

### HasNetwork6OperationDetailsAddedBy

`func (o *DataInnerIpamNetwork6Data) HasNetwork6OperationDetailsAddedBy() bool`

HasNetwork6OperationDetailsAddedBy returns a boolean if a field has been set.

### GetNetwork6OperationDetailsLastUpdatedOn

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsLastUpdatedOn() string`

GetNetwork6OperationDetailsLastUpdatedOn returns the Network6OperationDetailsLastUpdatedOn field if non-nil, zero value otherwise.

### GetNetwork6OperationDetailsLastUpdatedOnOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6OperationDetailsLastUpdatedOnOk() (*string, bool)`

GetNetwork6OperationDetailsLastUpdatedOnOk returns a tuple with the Network6OperationDetailsLastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6OperationDetailsLastUpdatedOn

`func (o *DataInnerIpamNetwork6Data) SetNetwork6OperationDetailsLastUpdatedOn(v string)`

SetNetwork6OperationDetailsLastUpdatedOn sets Network6OperationDetailsLastUpdatedOn field to given value.

### HasNetwork6OperationDetailsLastUpdatedOn

`func (o *DataInnerIpamNetwork6Data) HasNetwork6OperationDetailsLastUpdatedOn() bool`

HasNetwork6OperationDetailsLastUpdatedOn returns a boolean if a field has been set.

### GetDomainId

`func (o *DataInnerIpamNetwork6Data) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DataInnerIpamNetwork6Data) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DataInnerIpamNetwork6Data) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DataInnerIpamNetwork6Data) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *DataInnerIpamNetwork6Data) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DataInnerIpamNetwork6Data) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DataInnerIpamNetwork6Data) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DataInnerIpamNetwork6Data) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRangeId

`func (o *DataInnerIpamNetwork6Data) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *DataInnerIpamNetwork6Data) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *DataInnerIpamNetwork6Data) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *DataInnerIpamNetwork6Data) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *DataInnerIpamNetwork6Data) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *DataInnerIpamNetwork6Data) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *DataInnerIpamNetwork6Data) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *DataInnerIpamNetwork6Data) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetVlanId

`func (o *DataInnerIpamNetwork6Data) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *DataInnerIpamNetwork6Data) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *DataInnerIpamNetwork6Data) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *DataInnerIpamNetwork6Data) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanName

`func (o *DataInnerIpamNetwork6Data) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *DataInnerIpamNetwork6Data) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *DataInnerIpamNetwork6Data) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *DataInnerIpamNetwork6Data) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetVlanVlanId

`func (o *DataInnerIpamNetwork6Data) GetVlanVlanId() string`

GetVlanVlanId returns the VlanVlanId field if non-nil, zero value otherwise.

### GetVlanVlanIdOk

`func (o *DataInnerIpamNetwork6Data) GetVlanVlanIdOk() (*string, bool)`

GetVlanVlanIdOk returns a tuple with the VlanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVlanId

`func (o *DataInnerIpamNetwork6Data) SetVlanVlanId(v string)`

SetVlanVlanId sets VlanVlanId field to given value.

### HasVlanVlanId

`func (o *DataInnerIpamNetwork6Data) HasVlanVlanId() bool`

HasVlanVlanId returns a boolean if a field has been set.

### GetVlsmBlock6Id

`func (o *DataInnerIpamNetwork6Data) GetVlsmBlock6Id() string`

GetVlsmBlock6Id returns the VlsmBlock6Id field if non-nil, zero value otherwise.

### GetVlsmBlock6IdOk

`func (o *DataInnerIpamNetwork6Data) GetVlsmBlock6IdOk() (*string, bool)`

GetVlsmBlock6IdOk returns a tuple with the VlsmBlock6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmBlock6Id

`func (o *DataInnerIpamNetwork6Data) SetVlsmBlock6Id(v string)`

SetVlsmBlock6Id sets VlsmBlock6Id field to given value.

### HasVlsmBlock6Id

`func (o *DataInnerIpamNetwork6Data) HasVlsmBlock6Id() bool`

HasVlsmBlock6Id returns a boolean if a field has been set.

### GetVlsmSpaceId

`func (o *DataInnerIpamNetwork6Data) GetVlsmSpaceId() string`

GetVlsmSpaceId returns the VlsmSpaceId field if non-nil, zero value otherwise.

### GetVlsmSpaceIdOk

`func (o *DataInnerIpamNetwork6Data) GetVlsmSpaceIdOk() (*string, bool)`

GetVlsmSpaceIdOk returns a tuple with the VlsmSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceId

`func (o *DataInnerIpamNetwork6Data) SetVlsmSpaceId(v string)`

SetVlsmSpaceId sets VlsmSpaceId field to given value.

### HasVlsmSpaceId

`func (o *DataInnerIpamNetwork6Data) HasVlsmSpaceId() bool`

HasVlsmSpaceId returns a boolean if a field has been set.

### GetVlsmSpaceName

`func (o *DataInnerIpamNetwork6Data) GetVlsmSpaceName() string`

GetVlsmSpaceName returns the VlsmSpaceName field if non-nil, zero value otherwise.

### GetVlsmSpaceNameOk

`func (o *DataInnerIpamNetwork6Data) GetVlsmSpaceNameOk() (*string, bool)`

GetVlsmSpaceNameOk returns a tuple with the VlsmSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceName

`func (o *DataInnerIpamNetwork6Data) SetVlsmSpaceName(v string)`

SetVlsmSpaceName sets VlsmSpaceName field to given value.

### HasVlsmSpaceName

`func (o *DataInnerIpamNetwork6Data) HasVlsmSpaceName() bool`

HasVlsmSpaceName returns a boolean if a field has been set.

### GetVlsmNetwork6Id

`func (o *DataInnerIpamNetwork6Data) GetVlsmNetwork6Id() string`

GetVlsmNetwork6Id returns the VlsmNetwork6Id field if non-nil, zero value otherwise.

### GetVlsmNetwork6IdOk

`func (o *DataInnerIpamNetwork6Data) GetVlsmNetwork6IdOk() (*string, bool)`

GetVlsmNetwork6IdOk returns a tuple with the VlsmNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmNetwork6Id

`func (o *DataInnerIpamNetwork6Data) SetVlsmNetwork6Id(v string)`

SetVlsmNetwork6Id sets VlsmNetwork6Id field to given value.

### HasVlsmNetwork6Id

`func (o *DataInnerIpamNetwork6Data) HasVlsmNetwork6Id() bool`

HasVlsmNetwork6Id returns a boolean if a field has been set.

### GetNetwork6RipeWaitingState

`func (o *DataInnerIpamNetwork6Data) GetNetwork6RipeWaitingState() string`

GetNetwork6RipeWaitingState returns the Network6RipeWaitingState field if non-nil, zero value otherwise.

### GetNetwork6RipeWaitingStateOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6RipeWaitingStateOk() (*string, bool)`

GetNetwork6RipeWaitingStateOk returns a tuple with the Network6RipeWaitingState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6RipeWaitingState

`func (o *DataInnerIpamNetwork6Data) SetNetwork6RipeWaitingState(v string)`

SetNetwork6RipeWaitingState sets Network6RipeWaitingState field to given value.

### HasNetwork6RipeWaitingState

`func (o *DataInnerIpamNetwork6Data) HasNetwork6RipeWaitingState() bool`

HasNetwork6RipeWaitingState returns a boolean if a field has been set.

### GetNetwork6RipeWaitingStatus

`func (o *DataInnerIpamNetwork6Data) GetNetwork6RipeWaitingStatus() string`

GetNetwork6RipeWaitingStatus returns the Network6RipeWaitingStatus field if non-nil, zero value otherwise.

### GetNetwork6RipeWaitingStatusOk

`func (o *DataInnerIpamNetwork6Data) GetNetwork6RipeWaitingStatusOk() (*string, bool)`

GetNetwork6RipeWaitingStatusOk returns a tuple with the Network6RipeWaitingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6RipeWaitingStatus

`func (o *DataInnerIpamNetwork6Data) SetNetwork6RipeWaitingStatus(v string)`

SetNetwork6RipeWaitingStatus sets Network6RipeWaitingStatus field to given value.

### HasNetwork6RipeWaitingStatus

`func (o *DataInnerIpamNetwork6Data) HasNetwork6RipeWaitingStatus() bool`

HasNetwork6RipeWaitingStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


