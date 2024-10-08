# DataInnerDhcpFailoverData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerClusterRole** | Pointer to **string** | The role of the server the object belongs to in the cluster, either &lt;b&gt;active (M)&lt;/b&gt;, &lt;b&gt;passive (B)&lt;/b&gt; or &lt;b&gt;N/A (#)&lt;/b&gt;. | [optional] 
**FailoverDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**FailoverDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerState** | Pointer to **string** | The status of the DHCPv4 smart architecture. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft Windows DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv4 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**FailoverAddr** | Pointer to **string** | The IP address of the primary DHCPv4 server. | [optional] 
**FailoverAutoPartnerDown** | Pointer to **string** | The time after which the DHCPv4 failover channel automatically switches to &lt;b&gt;partner-down&lt;/b&gt; after being in &lt;b&gt;communication-interrupted&lt;/b&gt; state, in hours. | [optional] 
**FailoverId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 failover channel, a unique numeric key value automatically incremented when you add a failover channel. | [optional] 
**FailoverMclt** | Pointer to **string** | The maximum client lead time (MCLT) of the failover channel, in seconds. It indicates for how long each DHCP server can extend the lease of a client, beyond the time known by its partner server. | [optional] 
**FailoverName** | Pointer to **string** | The name of the DHCPv4 failover channel. | [optional] 
**FailoverPeerAddr** | Pointer to **string** | The IP address of the secondary DHCPv4 server. | [optional] 
**FailoverPeerPort** | Pointer to **string** | The port number of the secondary DHCPv4 server. | [optional] 
**FailoverPort** | Pointer to **string** | The port number of the primary DHCPv4 server. | [optional] 
**FailoverSplit** | Pointer to **string** | Internal use. Not documented. | [optional] 
**FailoverState** | Pointer to **string** | The status of the DHCPv4 failover channel, either &lt;b&gt;startup&lt;/b&gt;, &lt;b&gt;normal&lt;/b&gt;, &lt;b&gt;communications-interrupted&lt;/b&gt; or &lt;b&gt;recover-wait&lt;/b&gt;. | [optional] 
**FailoverType** | Pointer to **string** | The type of the DHCPv4 failover channel, either &lt;b&gt;primary&lt;/b&gt; or &lt;b&gt;secondary&lt;/b&gt;. | [optional] 
**ServerHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server_addr&lt;/b&gt; or &lt;b&gt;server_addr6&lt;/b&gt;. | [optional] 
**ServerAddr6** | Pointer to **string** | The Management IP address of the DHCPv4 server, the IPv6 address configured when adding the server, on which the failover channel is configured, in hexadecimal format. | [optional] 
**ServerAddr** | Pointer to **string** | The Management IP address of the DHCPv4 server, the IPv4 address configured when adding the server, on which the failover channel is configured, in hexadecimal format. | [optional] 
**FailoverMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**PeerServerId** | Pointer to **string** | The database identifier (ID) of the secondary DHCPv4 server. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpFailoverData

`func NewDataInnerDhcpFailoverData() *DataInnerDhcpFailoverData`

NewDataInnerDhcpFailoverData instantiates a new DataInnerDhcpFailoverData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpFailoverDataWithDefaults

`func NewDataInnerDhcpFailoverDataWithDefaults() *DataInnerDhcpFailoverData`

NewDataInnerDhcpFailoverDataWithDefaults instantiates a new DataInnerDhcpFailoverData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerClusterRole

`func (o *DataInnerDhcpFailoverData) GetServerClusterRole() string`

GetServerClusterRole returns the ServerClusterRole field if non-nil, zero value otherwise.

### GetServerClusterRoleOk

`func (o *DataInnerDhcpFailoverData) GetServerClusterRoleOk() (*string, bool)`

GetServerClusterRoleOk returns a tuple with the ServerClusterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClusterRole

`func (o *DataInnerDhcpFailoverData) SetServerClusterRole(v string)`

SetServerClusterRole sets ServerClusterRole field to given value.

### HasServerClusterRole

`func (o *DataInnerDhcpFailoverData) HasServerClusterRole() bool`

HasServerClusterRole returns a boolean if a field has been set.

### GetFailoverDelayedCreateTime

`func (o *DataInnerDhcpFailoverData) GetFailoverDelayedCreateTime() string`

GetFailoverDelayedCreateTime returns the FailoverDelayedCreateTime field if non-nil, zero value otherwise.

### GetFailoverDelayedCreateTimeOk

`func (o *DataInnerDhcpFailoverData) GetFailoverDelayedCreateTimeOk() (*string, bool)`

GetFailoverDelayedCreateTimeOk returns a tuple with the FailoverDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverDelayedCreateTime

`func (o *DataInnerDhcpFailoverData) SetFailoverDelayedCreateTime(v string)`

SetFailoverDelayedCreateTime sets FailoverDelayedCreateTime field to given value.

### HasFailoverDelayedCreateTime

`func (o *DataInnerDhcpFailoverData) HasFailoverDelayedCreateTime() bool`

HasFailoverDelayedCreateTime returns a boolean if a field has been set.

### GetFailoverDelayedDeleteTime

`func (o *DataInnerDhcpFailoverData) GetFailoverDelayedDeleteTime() string`

GetFailoverDelayedDeleteTime returns the FailoverDelayedDeleteTime field if non-nil, zero value otherwise.

### GetFailoverDelayedDeleteTimeOk

`func (o *DataInnerDhcpFailoverData) GetFailoverDelayedDeleteTimeOk() (*string, bool)`

GetFailoverDelayedDeleteTimeOk returns a tuple with the FailoverDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverDelayedDeleteTime

`func (o *DataInnerDhcpFailoverData) SetFailoverDelayedDeleteTime(v string)`

SetFailoverDelayedDeleteTime sets FailoverDelayedDeleteTime field to given value.

### HasFailoverDelayedDeleteTime

`func (o *DataInnerDhcpFailoverData) HasFailoverDelayedDeleteTime() bool`

HasFailoverDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDhcpFailoverData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDhcpFailoverData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDhcpFailoverData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDhcpFailoverData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDhcpFailoverData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDhcpFailoverData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDhcpFailoverData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDhcpFailoverData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerState

`func (o *DataInnerDhcpFailoverData) GetServerState() string`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *DataInnerDhcpFailoverData) GetServerStateOk() (*string, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *DataInnerDhcpFailoverData) SetServerState(v string)`

SetServerState sets ServerState field to given value.

### HasServerState

`func (o *DataInnerDhcpFailoverData) HasServerState() bool`

HasServerState returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDhcpFailoverData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDhcpFailoverData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDhcpFailoverData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDhcpFailoverData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetFailoverAddr

`func (o *DataInnerDhcpFailoverData) GetFailoverAddr() string`

GetFailoverAddr returns the FailoverAddr field if non-nil, zero value otherwise.

### GetFailoverAddrOk

`func (o *DataInnerDhcpFailoverData) GetFailoverAddrOk() (*string, bool)`

GetFailoverAddrOk returns a tuple with the FailoverAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAddr

`func (o *DataInnerDhcpFailoverData) SetFailoverAddr(v string)`

SetFailoverAddr sets FailoverAddr field to given value.

### HasFailoverAddr

`func (o *DataInnerDhcpFailoverData) HasFailoverAddr() bool`

HasFailoverAddr returns a boolean if a field has been set.

### GetFailoverAutoPartnerDown

`func (o *DataInnerDhcpFailoverData) GetFailoverAutoPartnerDown() string`

GetFailoverAutoPartnerDown returns the FailoverAutoPartnerDown field if non-nil, zero value otherwise.

### GetFailoverAutoPartnerDownOk

`func (o *DataInnerDhcpFailoverData) GetFailoverAutoPartnerDownOk() (*string, bool)`

GetFailoverAutoPartnerDownOk returns a tuple with the FailoverAutoPartnerDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAutoPartnerDown

`func (o *DataInnerDhcpFailoverData) SetFailoverAutoPartnerDown(v string)`

SetFailoverAutoPartnerDown sets FailoverAutoPartnerDown field to given value.

### HasFailoverAutoPartnerDown

`func (o *DataInnerDhcpFailoverData) HasFailoverAutoPartnerDown() bool`

HasFailoverAutoPartnerDown returns a boolean if a field has been set.

### GetFailoverId

`func (o *DataInnerDhcpFailoverData) GetFailoverId() string`

GetFailoverId returns the FailoverId field if non-nil, zero value otherwise.

### GetFailoverIdOk

`func (o *DataInnerDhcpFailoverData) GetFailoverIdOk() (*string, bool)`

GetFailoverIdOk returns a tuple with the FailoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverId

`func (o *DataInnerDhcpFailoverData) SetFailoverId(v string)`

SetFailoverId sets FailoverId field to given value.

### HasFailoverId

`func (o *DataInnerDhcpFailoverData) HasFailoverId() bool`

HasFailoverId returns a boolean if a field has been set.

### GetFailoverMclt

`func (o *DataInnerDhcpFailoverData) GetFailoverMclt() string`

GetFailoverMclt returns the FailoverMclt field if non-nil, zero value otherwise.

### GetFailoverMcltOk

`func (o *DataInnerDhcpFailoverData) GetFailoverMcltOk() (*string, bool)`

GetFailoverMcltOk returns a tuple with the FailoverMclt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMclt

`func (o *DataInnerDhcpFailoverData) SetFailoverMclt(v string)`

SetFailoverMclt sets FailoverMclt field to given value.

### HasFailoverMclt

`func (o *DataInnerDhcpFailoverData) HasFailoverMclt() bool`

HasFailoverMclt returns a boolean if a field has been set.

### GetFailoverName

`func (o *DataInnerDhcpFailoverData) GetFailoverName() string`

GetFailoverName returns the FailoverName field if non-nil, zero value otherwise.

### GetFailoverNameOk

`func (o *DataInnerDhcpFailoverData) GetFailoverNameOk() (*string, bool)`

GetFailoverNameOk returns a tuple with the FailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverName

`func (o *DataInnerDhcpFailoverData) SetFailoverName(v string)`

SetFailoverName sets FailoverName field to given value.

### HasFailoverName

`func (o *DataInnerDhcpFailoverData) HasFailoverName() bool`

HasFailoverName returns a boolean if a field has been set.

### GetFailoverPeerAddr

`func (o *DataInnerDhcpFailoverData) GetFailoverPeerAddr() string`

GetFailoverPeerAddr returns the FailoverPeerAddr field if non-nil, zero value otherwise.

### GetFailoverPeerAddrOk

`func (o *DataInnerDhcpFailoverData) GetFailoverPeerAddrOk() (*string, bool)`

GetFailoverPeerAddrOk returns a tuple with the FailoverPeerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPeerAddr

`func (o *DataInnerDhcpFailoverData) SetFailoverPeerAddr(v string)`

SetFailoverPeerAddr sets FailoverPeerAddr field to given value.

### HasFailoverPeerAddr

`func (o *DataInnerDhcpFailoverData) HasFailoverPeerAddr() bool`

HasFailoverPeerAddr returns a boolean if a field has been set.

### GetFailoverPeerPort

`func (o *DataInnerDhcpFailoverData) GetFailoverPeerPort() string`

GetFailoverPeerPort returns the FailoverPeerPort field if non-nil, zero value otherwise.

### GetFailoverPeerPortOk

`func (o *DataInnerDhcpFailoverData) GetFailoverPeerPortOk() (*string, bool)`

GetFailoverPeerPortOk returns a tuple with the FailoverPeerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPeerPort

`func (o *DataInnerDhcpFailoverData) SetFailoverPeerPort(v string)`

SetFailoverPeerPort sets FailoverPeerPort field to given value.

### HasFailoverPeerPort

`func (o *DataInnerDhcpFailoverData) HasFailoverPeerPort() bool`

HasFailoverPeerPort returns a boolean if a field has been set.

### GetFailoverPort

`func (o *DataInnerDhcpFailoverData) GetFailoverPort() string`

GetFailoverPort returns the FailoverPort field if non-nil, zero value otherwise.

### GetFailoverPortOk

`func (o *DataInnerDhcpFailoverData) GetFailoverPortOk() (*string, bool)`

GetFailoverPortOk returns a tuple with the FailoverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPort

`func (o *DataInnerDhcpFailoverData) SetFailoverPort(v string)`

SetFailoverPort sets FailoverPort field to given value.

### HasFailoverPort

`func (o *DataInnerDhcpFailoverData) HasFailoverPort() bool`

HasFailoverPort returns a boolean if a field has been set.

### GetFailoverSplit

`func (o *DataInnerDhcpFailoverData) GetFailoverSplit() string`

GetFailoverSplit returns the FailoverSplit field if non-nil, zero value otherwise.

### GetFailoverSplitOk

`func (o *DataInnerDhcpFailoverData) GetFailoverSplitOk() (*string, bool)`

GetFailoverSplitOk returns a tuple with the FailoverSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverSplit

`func (o *DataInnerDhcpFailoverData) SetFailoverSplit(v string)`

SetFailoverSplit sets FailoverSplit field to given value.

### HasFailoverSplit

`func (o *DataInnerDhcpFailoverData) HasFailoverSplit() bool`

HasFailoverSplit returns a boolean if a field has been set.

### GetFailoverState

`func (o *DataInnerDhcpFailoverData) GetFailoverState() string`

GetFailoverState returns the FailoverState field if non-nil, zero value otherwise.

### GetFailoverStateOk

`func (o *DataInnerDhcpFailoverData) GetFailoverStateOk() (*string, bool)`

GetFailoverStateOk returns a tuple with the FailoverState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverState

`func (o *DataInnerDhcpFailoverData) SetFailoverState(v string)`

SetFailoverState sets FailoverState field to given value.

### HasFailoverState

`func (o *DataInnerDhcpFailoverData) HasFailoverState() bool`

HasFailoverState returns a boolean if a field has been set.

### GetFailoverType

`func (o *DataInnerDhcpFailoverData) GetFailoverType() string`

GetFailoverType returns the FailoverType field if non-nil, zero value otherwise.

### GetFailoverTypeOk

`func (o *DataInnerDhcpFailoverData) GetFailoverTypeOk() (*string, bool)`

GetFailoverTypeOk returns a tuple with the FailoverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverType

`func (o *DataInnerDhcpFailoverData) SetFailoverType(v string)`

SetFailoverType sets FailoverType field to given value.

### HasFailoverType

`func (o *DataInnerDhcpFailoverData) HasFailoverType() bool`

HasFailoverType returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DataInnerDhcpFailoverData) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DataInnerDhcpFailoverData) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DataInnerDhcpFailoverData) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DataInnerDhcpFailoverData) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerAddr6

`func (o *DataInnerDhcpFailoverData) GetServerAddr6() string`

GetServerAddr6 returns the ServerAddr6 field if non-nil, zero value otherwise.

### GetServerAddr6Ok

`func (o *DataInnerDhcpFailoverData) GetServerAddr6Ok() (*string, bool)`

GetServerAddr6Ok returns a tuple with the ServerAddr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr6

`func (o *DataInnerDhcpFailoverData) SetServerAddr6(v string)`

SetServerAddr6 sets ServerAddr6 field to given value.

### HasServerAddr6

`func (o *DataInnerDhcpFailoverData) HasServerAddr6() bool`

HasServerAddr6 returns a boolean if a field has been set.

### GetServerAddr

`func (o *DataInnerDhcpFailoverData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DataInnerDhcpFailoverData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DataInnerDhcpFailoverData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DataInnerDhcpFailoverData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetFailoverMultistatus

`func (o *DataInnerDhcpFailoverData) GetFailoverMultistatus() string`

GetFailoverMultistatus returns the FailoverMultistatus field if non-nil, zero value otherwise.

### GetFailoverMultistatusOk

`func (o *DataInnerDhcpFailoverData) GetFailoverMultistatusOk() (*string, bool)`

GetFailoverMultistatusOk returns a tuple with the FailoverMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMultistatus

`func (o *DataInnerDhcpFailoverData) SetFailoverMultistatus(v string)`

SetFailoverMultistatus sets FailoverMultistatus field to given value.

### HasFailoverMultistatus

`func (o *DataInnerDhcpFailoverData) HasFailoverMultistatus() bool`

HasFailoverMultistatus returns a boolean if a field has been set.

### GetPeerServerId

`func (o *DataInnerDhcpFailoverData) GetPeerServerId() string`

GetPeerServerId returns the PeerServerId field if non-nil, zero value otherwise.

### GetPeerServerIdOk

`func (o *DataInnerDhcpFailoverData) GetPeerServerIdOk() (*string, bool)`

GetPeerServerIdOk returns a tuple with the PeerServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerId

`func (o *DataInnerDhcpFailoverData) SetPeerServerId(v string)`

SetPeerServerId sets PeerServerId field to given value.

### HasPeerServerId

`func (o *DataInnerDhcpFailoverData) HasPeerServerId() bool`

HasPeerServerId returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpFailoverData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpFailoverData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpFailoverData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpFailoverData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


