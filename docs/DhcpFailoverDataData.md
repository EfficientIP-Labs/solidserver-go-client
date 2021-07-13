# DhcpFailoverDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FailoverDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**FailoverDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerState** | Pointer to **string** | The status of the DHCPv4 smart architecture. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**FailoverAddr** | Pointer to **string** | The IP address of the primary DHCPv4 server. | [optional] 
**FailoverAutoPartnerDown** | Pointer to **string** | The time after which the DHCPv4 failover channel automatically switches to &lt;b&gt;partner-down&lt;/b&gt; after being in &lt;b&gt;communication-interrupted&lt;/b&gt; state, in hours. | [optional] 
**FailoverId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 failover channel, a unique numeric key value automatically incremented when you add a failover channel. | [optional] 
**FailoverName** | Pointer to **string** | The name of the DHCPv4 failover channel. | [optional] 
**FailoverPeerAddr** | Pointer to **string** | The IP address of the secondary DHCPv4 server. | [optional] 
**FailoverPeerPort** | Pointer to **string** | The port number of the secondary DHCPv4 server. | [optional] 
**FailoverPort** | Pointer to **string** | The port number of the primary DHCPv4 server. | [optional] 
**FailoverSplit** | Pointer to **string** | Internal use. Not documented. | [optional] 
**FailoverState** | Pointer to **string** | The status of the DHCPv4 failover channel, either &lt;b&gt;startup&lt;/b&gt;, &lt;b&gt;normal&lt;/b&gt;, &lt;b&gt;communications-interrupted&lt;/b&gt; or &lt;b&gt;recover-wait&lt;/b&gt;. | [optional] 
**FailoverType** | Pointer to **string** | The type of the DHCPv4 failover channel, either &lt;b&gt;primary&lt;/b&gt; or &lt;b&gt;secondary&lt;/b&gt;. | [optional] 
**ServerAddr** | Pointer to **string** | The IP address of the DHCP server on which the failover channel is configured, in hexadecimal format. | [optional] 
**FailoverMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**PeerServerId** | Pointer to **string** | The database identifier (ID) of the secondary DHCPv4 server. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpFailoverDataData

`func NewDhcpFailoverDataData() *DhcpFailoverDataData`

NewDhcpFailoverDataData instantiates a new DhcpFailoverDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpFailoverDataDataWithDefaults

`func NewDhcpFailoverDataDataWithDefaults() *DhcpFailoverDataData`

NewDhcpFailoverDataDataWithDefaults instantiates a new DhcpFailoverDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFailoverDelayedCreateTime

`func (o *DhcpFailoverDataData) GetFailoverDelayedCreateTime() string`

GetFailoverDelayedCreateTime returns the FailoverDelayedCreateTime field if non-nil, zero value otherwise.

### GetFailoverDelayedCreateTimeOk

`func (o *DhcpFailoverDataData) GetFailoverDelayedCreateTimeOk() (*string, bool)`

GetFailoverDelayedCreateTimeOk returns a tuple with the FailoverDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverDelayedCreateTime

`func (o *DhcpFailoverDataData) SetFailoverDelayedCreateTime(v string)`

SetFailoverDelayedCreateTime sets FailoverDelayedCreateTime field to given value.

### HasFailoverDelayedCreateTime

`func (o *DhcpFailoverDataData) HasFailoverDelayedCreateTime() bool`

HasFailoverDelayedCreateTime returns a boolean if a field has been set.

### GetFailoverDelayedDeleteTime

`func (o *DhcpFailoverDataData) GetFailoverDelayedDeleteTime() string`

GetFailoverDelayedDeleteTime returns the FailoverDelayedDeleteTime field if non-nil, zero value otherwise.

### GetFailoverDelayedDeleteTimeOk

`func (o *DhcpFailoverDataData) GetFailoverDelayedDeleteTimeOk() (*string, bool)`

GetFailoverDelayedDeleteTimeOk returns a tuple with the FailoverDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverDelayedDeleteTime

`func (o *DhcpFailoverDataData) SetFailoverDelayedDeleteTime(v string)`

SetFailoverDelayedDeleteTime sets FailoverDelayedDeleteTime field to given value.

### HasFailoverDelayedDeleteTime

`func (o *DhcpFailoverDataData) HasFailoverDelayedDeleteTime() bool`

HasFailoverDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpFailoverDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpFailoverDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpFailoverDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpFailoverDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpFailoverDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpFailoverDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpFailoverDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpFailoverDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerState

`func (o *DhcpFailoverDataData) GetServerState() string`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *DhcpFailoverDataData) GetServerStateOk() (*string, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *DhcpFailoverDataData) SetServerState(v string)`

SetServerState sets ServerState field to given value.

### HasServerState

`func (o *DhcpFailoverDataData) HasServerState() bool`

HasServerState returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpFailoverDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpFailoverDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpFailoverDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpFailoverDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetFailoverAddr

`func (o *DhcpFailoverDataData) GetFailoverAddr() string`

GetFailoverAddr returns the FailoverAddr field if non-nil, zero value otherwise.

### GetFailoverAddrOk

`func (o *DhcpFailoverDataData) GetFailoverAddrOk() (*string, bool)`

GetFailoverAddrOk returns a tuple with the FailoverAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAddr

`func (o *DhcpFailoverDataData) SetFailoverAddr(v string)`

SetFailoverAddr sets FailoverAddr field to given value.

### HasFailoverAddr

`func (o *DhcpFailoverDataData) HasFailoverAddr() bool`

HasFailoverAddr returns a boolean if a field has been set.

### GetFailoverAutoPartnerDown

`func (o *DhcpFailoverDataData) GetFailoverAutoPartnerDown() string`

GetFailoverAutoPartnerDown returns the FailoverAutoPartnerDown field if non-nil, zero value otherwise.

### GetFailoverAutoPartnerDownOk

`func (o *DhcpFailoverDataData) GetFailoverAutoPartnerDownOk() (*string, bool)`

GetFailoverAutoPartnerDownOk returns a tuple with the FailoverAutoPartnerDown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverAutoPartnerDown

`func (o *DhcpFailoverDataData) SetFailoverAutoPartnerDown(v string)`

SetFailoverAutoPartnerDown sets FailoverAutoPartnerDown field to given value.

### HasFailoverAutoPartnerDown

`func (o *DhcpFailoverDataData) HasFailoverAutoPartnerDown() bool`

HasFailoverAutoPartnerDown returns a boolean if a field has been set.

### GetFailoverId

`func (o *DhcpFailoverDataData) GetFailoverId() string`

GetFailoverId returns the FailoverId field if non-nil, zero value otherwise.

### GetFailoverIdOk

`func (o *DhcpFailoverDataData) GetFailoverIdOk() (*string, bool)`

GetFailoverIdOk returns a tuple with the FailoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverId

`func (o *DhcpFailoverDataData) SetFailoverId(v string)`

SetFailoverId sets FailoverId field to given value.

### HasFailoverId

`func (o *DhcpFailoverDataData) HasFailoverId() bool`

HasFailoverId returns a boolean if a field has been set.

### GetFailoverName

`func (o *DhcpFailoverDataData) GetFailoverName() string`

GetFailoverName returns the FailoverName field if non-nil, zero value otherwise.

### GetFailoverNameOk

`func (o *DhcpFailoverDataData) GetFailoverNameOk() (*string, bool)`

GetFailoverNameOk returns a tuple with the FailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverName

`func (o *DhcpFailoverDataData) SetFailoverName(v string)`

SetFailoverName sets FailoverName field to given value.

### HasFailoverName

`func (o *DhcpFailoverDataData) HasFailoverName() bool`

HasFailoverName returns a boolean if a field has been set.

### GetFailoverPeerAddr

`func (o *DhcpFailoverDataData) GetFailoverPeerAddr() string`

GetFailoverPeerAddr returns the FailoverPeerAddr field if non-nil, zero value otherwise.

### GetFailoverPeerAddrOk

`func (o *DhcpFailoverDataData) GetFailoverPeerAddrOk() (*string, bool)`

GetFailoverPeerAddrOk returns a tuple with the FailoverPeerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPeerAddr

`func (o *DhcpFailoverDataData) SetFailoverPeerAddr(v string)`

SetFailoverPeerAddr sets FailoverPeerAddr field to given value.

### HasFailoverPeerAddr

`func (o *DhcpFailoverDataData) HasFailoverPeerAddr() bool`

HasFailoverPeerAddr returns a boolean if a field has been set.

### GetFailoverPeerPort

`func (o *DhcpFailoverDataData) GetFailoverPeerPort() string`

GetFailoverPeerPort returns the FailoverPeerPort field if non-nil, zero value otherwise.

### GetFailoverPeerPortOk

`func (o *DhcpFailoverDataData) GetFailoverPeerPortOk() (*string, bool)`

GetFailoverPeerPortOk returns a tuple with the FailoverPeerPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPeerPort

`func (o *DhcpFailoverDataData) SetFailoverPeerPort(v string)`

SetFailoverPeerPort sets FailoverPeerPort field to given value.

### HasFailoverPeerPort

`func (o *DhcpFailoverDataData) HasFailoverPeerPort() bool`

HasFailoverPeerPort returns a boolean if a field has been set.

### GetFailoverPort

`func (o *DhcpFailoverDataData) GetFailoverPort() string`

GetFailoverPort returns the FailoverPort field if non-nil, zero value otherwise.

### GetFailoverPortOk

`func (o *DhcpFailoverDataData) GetFailoverPortOk() (*string, bool)`

GetFailoverPortOk returns a tuple with the FailoverPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverPort

`func (o *DhcpFailoverDataData) SetFailoverPort(v string)`

SetFailoverPort sets FailoverPort field to given value.

### HasFailoverPort

`func (o *DhcpFailoverDataData) HasFailoverPort() bool`

HasFailoverPort returns a boolean if a field has been set.

### GetFailoverSplit

`func (o *DhcpFailoverDataData) GetFailoverSplit() string`

GetFailoverSplit returns the FailoverSplit field if non-nil, zero value otherwise.

### GetFailoverSplitOk

`func (o *DhcpFailoverDataData) GetFailoverSplitOk() (*string, bool)`

GetFailoverSplitOk returns a tuple with the FailoverSplit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverSplit

`func (o *DhcpFailoverDataData) SetFailoverSplit(v string)`

SetFailoverSplit sets FailoverSplit field to given value.

### HasFailoverSplit

`func (o *DhcpFailoverDataData) HasFailoverSplit() bool`

HasFailoverSplit returns a boolean if a field has been set.

### GetFailoverState

`func (o *DhcpFailoverDataData) GetFailoverState() string`

GetFailoverState returns the FailoverState field if non-nil, zero value otherwise.

### GetFailoverStateOk

`func (o *DhcpFailoverDataData) GetFailoverStateOk() (*string, bool)`

GetFailoverStateOk returns a tuple with the FailoverState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverState

`func (o *DhcpFailoverDataData) SetFailoverState(v string)`

SetFailoverState sets FailoverState field to given value.

### HasFailoverState

`func (o *DhcpFailoverDataData) HasFailoverState() bool`

HasFailoverState returns a boolean if a field has been set.

### GetFailoverType

`func (o *DhcpFailoverDataData) GetFailoverType() string`

GetFailoverType returns the FailoverType field if non-nil, zero value otherwise.

### GetFailoverTypeOk

`func (o *DhcpFailoverDataData) GetFailoverTypeOk() (*string, bool)`

GetFailoverTypeOk returns a tuple with the FailoverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverType

`func (o *DhcpFailoverDataData) SetFailoverType(v string)`

SetFailoverType sets FailoverType field to given value.

### HasFailoverType

`func (o *DhcpFailoverDataData) HasFailoverType() bool`

HasFailoverType returns a boolean if a field has been set.

### GetServerAddr

`func (o *DhcpFailoverDataData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DhcpFailoverDataData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DhcpFailoverDataData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DhcpFailoverDataData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetFailoverMultistatus

`func (o *DhcpFailoverDataData) GetFailoverMultistatus() string`

GetFailoverMultistatus returns the FailoverMultistatus field if non-nil, zero value otherwise.

### GetFailoverMultistatusOk

`func (o *DhcpFailoverDataData) GetFailoverMultistatusOk() (*string, bool)`

GetFailoverMultistatusOk returns a tuple with the FailoverMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverMultistatus

`func (o *DhcpFailoverDataData) SetFailoverMultistatus(v string)`

SetFailoverMultistatus sets FailoverMultistatus field to given value.

### HasFailoverMultistatus

`func (o *DhcpFailoverDataData) HasFailoverMultistatus() bool`

HasFailoverMultistatus returns a boolean if a field has been set.

### GetPeerServerId

`func (o *DhcpFailoverDataData) GetPeerServerId() string`

GetPeerServerId returns the PeerServerId field if non-nil, zero value otherwise.

### GetPeerServerIdOk

`func (o *DhcpFailoverDataData) GetPeerServerIdOk() (*string, bool)`

GetPeerServerIdOk returns a tuple with the PeerServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerServerId

`func (o *DhcpFailoverDataData) SetPeerServerId(v string)`

SetPeerServerId sets PeerServerId field to given value.

### HasPeerServerId

`func (o *DhcpFailoverDataData) HasPeerServerId() bool`

HasPeerServerId returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpFailoverDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpFailoverDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpFailoverDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpFailoverDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


