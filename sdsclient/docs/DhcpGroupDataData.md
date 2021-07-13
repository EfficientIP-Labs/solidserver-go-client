# DhcpGroupDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**GroupDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server the object belongs to. | [optional] 
**GroupClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 group, it can be preceded by the class directory. | [optional] 
**GroupClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**GroupId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 group. | [optional] 
**GroupName** | Pointer to **string** | The name of the DHCPv4 group. | [optional] 
**ServerAddr** | Pointer to **string** | The IP address of the DHCP server the object belongs to, in hexadecimal format. | [optional] 
**GroupMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpGroupDataData

`func NewDhcpGroupDataData() *DhcpGroupDataData`

NewDhcpGroupDataData instantiates a new DhcpGroupDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpGroupDataDataWithDefaults

`func NewDhcpGroupDataDataWithDefaults() *DhcpGroupDataData`

NewDhcpGroupDataDataWithDefaults instantiates a new DhcpGroupDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupDelayedCreateTime

`func (o *DhcpGroupDataData) GetGroupDelayedCreateTime() string`

GetGroupDelayedCreateTime returns the GroupDelayedCreateTime field if non-nil, zero value otherwise.

### GetGroupDelayedCreateTimeOk

`func (o *DhcpGroupDataData) GetGroupDelayedCreateTimeOk() (*string, bool)`

GetGroupDelayedCreateTimeOk returns a tuple with the GroupDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDelayedCreateTime

`func (o *DhcpGroupDataData) SetGroupDelayedCreateTime(v string)`

SetGroupDelayedCreateTime sets GroupDelayedCreateTime field to given value.

### HasGroupDelayedCreateTime

`func (o *DhcpGroupDataData) HasGroupDelayedCreateTime() bool`

HasGroupDelayedCreateTime returns a boolean if a field has been set.

### GetGroupDelayedDeleteTime

`func (o *DhcpGroupDataData) GetGroupDelayedDeleteTime() string`

GetGroupDelayedDeleteTime returns the GroupDelayedDeleteTime field if non-nil, zero value otherwise.

### GetGroupDelayedDeleteTimeOk

`func (o *DhcpGroupDataData) GetGroupDelayedDeleteTimeOk() (*string, bool)`

GetGroupDelayedDeleteTimeOk returns a tuple with the GroupDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupDelayedDeleteTime

`func (o *DhcpGroupDataData) SetGroupDelayedDeleteTime(v string)`

SetGroupDelayedDeleteTime sets GroupDelayedDeleteTime field to given value.

### HasGroupDelayedDeleteTime

`func (o *DhcpGroupDataData) HasGroupDelayedDeleteTime() bool`

HasGroupDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DhcpGroupDataData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DhcpGroupDataData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DhcpGroupDataData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DhcpGroupDataData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DhcpGroupDataData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DhcpGroupDataData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DhcpGroupDataData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DhcpGroupDataData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpGroupDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpGroupDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpGroupDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpGroupDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpGroupDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpGroupDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpGroupDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpGroupDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpGroupDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpGroupDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpGroupDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpGroupDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DhcpGroupDataData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DhcpGroupDataData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DhcpGroupDataData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DhcpGroupDataData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetGroupClassName

`func (o *DhcpGroupDataData) GetGroupClassName() string`

GetGroupClassName returns the GroupClassName field if non-nil, zero value otherwise.

### GetGroupClassNameOk

`func (o *DhcpGroupDataData) GetGroupClassNameOk() (*string, bool)`

GetGroupClassNameOk returns a tuple with the GroupClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassName

`func (o *DhcpGroupDataData) SetGroupClassName(v string)`

SetGroupClassName sets GroupClassName field to given value.

### HasGroupClassName

`func (o *DhcpGroupDataData) HasGroupClassName() bool`

HasGroupClassName returns a boolean if a field has been set.

### GetGroupClassParameters

`func (o *DhcpGroupDataData) GetGroupClassParameters() []ApiClassParameterOutputEntry`

GetGroupClassParameters returns the GroupClassParameters field if non-nil, zero value otherwise.

### GetGroupClassParametersOk

`func (o *DhcpGroupDataData) GetGroupClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetGroupClassParametersOk returns a tuple with the GroupClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassParameters

`func (o *DhcpGroupDataData) SetGroupClassParameters(v []ApiClassParameterOutputEntry)`

SetGroupClassParameters sets GroupClassParameters field to given value.

### HasGroupClassParameters

`func (o *DhcpGroupDataData) HasGroupClassParameters() bool`

HasGroupClassParameters returns a boolean if a field has been set.

### GetGroupId

`func (o *DhcpGroupDataData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DhcpGroupDataData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DhcpGroupDataData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DhcpGroupDataData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *DhcpGroupDataData) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DhcpGroupDataData) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DhcpGroupDataData) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DhcpGroupDataData) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetServerAddr

`func (o *DhcpGroupDataData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DhcpGroupDataData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DhcpGroupDataData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DhcpGroupDataData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetGroupMultistatus

`func (o *DhcpGroupDataData) GetGroupMultistatus() string`

GetGroupMultistatus returns the GroupMultistatus field if non-nil, zero value otherwise.

### GetGroupMultistatusOk

`func (o *DhcpGroupDataData) GetGroupMultistatusOk() (*string, bool)`

GetGroupMultistatusOk returns a tuple with the GroupMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupMultistatus

`func (o *DhcpGroupDataData) SetGroupMultistatus(v string)`

SetGroupMultistatus sets GroupMultistatus field to given value.

### HasGroupMultistatus

`func (o *DhcpGroupDataData) HasGroupMultistatus() bool`

HasGroupMultistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpGroupDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpGroupDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpGroupDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpGroupDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DhcpGroupDataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DhcpGroupDataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DhcpGroupDataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DhcpGroupDataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


