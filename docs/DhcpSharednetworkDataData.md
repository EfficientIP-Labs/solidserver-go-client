# DhcpSharednetworkDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharednetworkDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**SharednetworkDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SharednetworkId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 shared network. | [optional] 
**SharednetworkName** | Pointer to **string** | The name of the DHCPv4 shared network. | [optional] 
**SharednetworkMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpSharednetworkDataData

`func NewDhcpSharednetworkDataData() *DhcpSharednetworkDataData`

NewDhcpSharednetworkDataData instantiates a new DhcpSharednetworkDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSharednetworkDataDataWithDefaults

`func NewDhcpSharednetworkDataDataWithDefaults() *DhcpSharednetworkDataData`

NewDhcpSharednetworkDataDataWithDefaults instantiates a new DhcpSharednetworkDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharednetworkDelayedCreateTime

`func (o *DhcpSharednetworkDataData) GetSharednetworkDelayedCreateTime() string`

GetSharednetworkDelayedCreateTime returns the SharednetworkDelayedCreateTime field if non-nil, zero value otherwise.

### GetSharednetworkDelayedCreateTimeOk

`func (o *DhcpSharednetworkDataData) GetSharednetworkDelayedCreateTimeOk() (*string, bool)`

GetSharednetworkDelayedCreateTimeOk returns a tuple with the SharednetworkDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkDelayedCreateTime

`func (o *DhcpSharednetworkDataData) SetSharednetworkDelayedCreateTime(v string)`

SetSharednetworkDelayedCreateTime sets SharednetworkDelayedCreateTime field to given value.

### HasSharednetworkDelayedCreateTime

`func (o *DhcpSharednetworkDataData) HasSharednetworkDelayedCreateTime() bool`

HasSharednetworkDelayedCreateTime returns a boolean if a field has been set.

### GetSharednetworkDelayedDeleteTime

`func (o *DhcpSharednetworkDataData) GetSharednetworkDelayedDeleteTime() string`

GetSharednetworkDelayedDeleteTime returns the SharednetworkDelayedDeleteTime field if non-nil, zero value otherwise.

### GetSharednetworkDelayedDeleteTimeOk

`func (o *DhcpSharednetworkDataData) GetSharednetworkDelayedDeleteTimeOk() (*string, bool)`

GetSharednetworkDelayedDeleteTimeOk returns a tuple with the SharednetworkDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkDelayedDeleteTime

`func (o *DhcpSharednetworkDataData) SetSharednetworkDelayedDeleteTime(v string)`

SetSharednetworkDelayedDeleteTime sets SharednetworkDelayedDeleteTime field to given value.

### HasSharednetworkDelayedDeleteTime

`func (o *DhcpSharednetworkDataData) HasSharednetworkDelayedDeleteTime() bool`

HasSharednetworkDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpSharednetworkDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpSharednetworkDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpSharednetworkDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpSharednetworkDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpSharednetworkDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpSharednetworkDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpSharednetworkDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpSharednetworkDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpSharednetworkDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpSharednetworkDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpSharednetworkDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpSharednetworkDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DhcpSharednetworkDataData) GetSharednetworkId() string`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DhcpSharednetworkDataData) GetSharednetworkIdOk() (*string, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DhcpSharednetworkDataData) SetSharednetworkId(v string)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DhcpSharednetworkDataData) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DhcpSharednetworkDataData) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DhcpSharednetworkDataData) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DhcpSharednetworkDataData) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DhcpSharednetworkDataData) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetSharednetworkMultistatus

`func (o *DhcpSharednetworkDataData) GetSharednetworkMultistatus() string`

GetSharednetworkMultistatus returns the SharednetworkMultistatus field if non-nil, zero value otherwise.

### GetSharednetworkMultistatusOk

`func (o *DhcpSharednetworkDataData) GetSharednetworkMultistatusOk() (*string, bool)`

GetSharednetworkMultistatusOk returns a tuple with the SharednetworkMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkMultistatus

`func (o *DhcpSharednetworkDataData) SetSharednetworkMultistatus(v string)`

SetSharednetworkMultistatus sets SharednetworkMultistatus field to given value.

### HasSharednetworkMultistatus

`func (o *DhcpSharednetworkDataData) HasSharednetworkMultistatus() bool`

HasSharednetworkMultistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpSharednetworkDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpSharednetworkDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpSharednetworkDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpSharednetworkDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


