# DhcpStaticDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**StaticDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server the object belongs to. | [optional] 
**GroupClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 group the static belongs to, it can be preceded by the class directory. | [optional] 
**GroupClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**GroupId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 group. | [optional] 
**GroupName** | Pointer to **string** | The name of the DHCPv4 group associated with the object. | [optional] 
**StaticAddr** | Pointer to **string** | The IP address associated with the DHCPv4 static. | [optional] 
**StaticClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 static, it can be preceded by the class directory. | [optional] 
**StaticClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**StaticDomain** | Pointer to **string** | The domain name associated with the DHCPv4 static. | [optional] 
**StaticExpireTime** | Pointer to **string** | The expiration time of the lease associated with the DHCPv4 static, in decimal UNIX date format. | [optional] 
**StaticId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 static. | [optional] 
**StaticIdentifier** | Pointer to **string** | TODO:dhcp_static_list.output.static_identifier | [optional] 
**StaticAddressAddr** | Pointer to **string** | The IP address associated with the DHCPv4 static, in hexadecimal format. | [optional] 
**StaticLastSeen** | Pointer to **string** | The last time the MAC address associated with the DHCPv4 static was seen on the network, in decimal UNIX date format. | [optional] 
**StaticMacAddr** | Pointer to **string** | The MAC address associated with the DHCPv4 static. It is composed of 7 sections, &lt;b&gt;00:11:22:33:44:55:66&lt;/b&gt;, where &lt;b&gt;00&lt;/b&gt; is the MAC address type. The type &lt;b&gt;01&lt;/b&gt; indicates Ethernet. | [optional] 
**StaticName** | Pointer to **string** | The name of the DHCPv4 static. | [optional] 
**ScopeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**ScopeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ScopeEndAddressAddr** | Pointer to **string** | The last IP address of the DHCPv4 scope the object belongs to, in hexadecimal format. | [optional] 
**ScopeId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeName** | Pointer to **string** | The name of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeNetAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope the object belongs. | [optional] 
**ScopeNetMask** | Pointer to **string** | The netmask of the DHCPv4 scope the object belongs to. It is expressed in dot-decimal notation and defines the number of addresses the scope contains. | [optional] 
**ScopeSpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the DHCPv4 scope the object belongs to. | [optional] 
**ScopeSize** | Pointer to **string** | The number of IP addresses the DHCPv4 scope the object belongs to contains. | [optional] 
**ScopeStartAddressAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope the object belongs to, in hexadecimal format. | [optional] 
**SharednetworkId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 shared network the object belongs to. | [optional] 
**SharednetworkName** | Pointer to **string** | The name of the DHCPv4 shared network the object belongs to. | [optional] 
**StaticMacVendor** | Pointer to **string** | The vendor details of the client associated with the DHCPv4 static. | [optional] 
**StaticMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpStaticDataData

`func NewDhcpStaticDataData() *DhcpStaticDataData`

NewDhcpStaticDataData instantiates a new DhcpStaticDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpStaticDataDataWithDefaults

`func NewDhcpStaticDataDataWithDefaults() *DhcpStaticDataData`

NewDhcpStaticDataDataWithDefaults instantiates a new DhcpStaticDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticDelayedCreateTime

`func (o *DhcpStaticDataData) GetStaticDelayedCreateTime() string`

GetStaticDelayedCreateTime returns the StaticDelayedCreateTime field if non-nil, zero value otherwise.

### GetStaticDelayedCreateTimeOk

`func (o *DhcpStaticDataData) GetStaticDelayedCreateTimeOk() (*string, bool)`

GetStaticDelayedCreateTimeOk returns a tuple with the StaticDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDelayedCreateTime

`func (o *DhcpStaticDataData) SetStaticDelayedCreateTime(v string)`

SetStaticDelayedCreateTime sets StaticDelayedCreateTime field to given value.

### HasStaticDelayedCreateTime

`func (o *DhcpStaticDataData) HasStaticDelayedCreateTime() bool`

HasStaticDelayedCreateTime returns a boolean if a field has been set.

### GetStaticDelayedDeleteTime

`func (o *DhcpStaticDataData) GetStaticDelayedDeleteTime() string`

GetStaticDelayedDeleteTime returns the StaticDelayedDeleteTime field if non-nil, zero value otherwise.

### GetStaticDelayedDeleteTimeOk

`func (o *DhcpStaticDataData) GetStaticDelayedDeleteTimeOk() (*string, bool)`

GetStaticDelayedDeleteTimeOk returns a tuple with the StaticDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDelayedDeleteTime

`func (o *DhcpStaticDataData) SetStaticDelayedDeleteTime(v string)`

SetStaticDelayedDeleteTime sets StaticDelayedDeleteTime field to given value.

### HasStaticDelayedDeleteTime

`func (o *DhcpStaticDataData) HasStaticDelayedDeleteTime() bool`

HasStaticDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DhcpStaticDataData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DhcpStaticDataData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DhcpStaticDataData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DhcpStaticDataData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DhcpStaticDataData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DhcpStaticDataData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DhcpStaticDataData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DhcpStaticDataData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpStaticDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpStaticDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpStaticDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpStaticDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpStaticDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpStaticDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpStaticDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpStaticDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpStaticDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpStaticDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpStaticDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpStaticDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DhcpStaticDataData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DhcpStaticDataData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DhcpStaticDataData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DhcpStaticDataData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetGroupClassName

`func (o *DhcpStaticDataData) GetGroupClassName() string`

GetGroupClassName returns the GroupClassName field if non-nil, zero value otherwise.

### GetGroupClassNameOk

`func (o *DhcpStaticDataData) GetGroupClassNameOk() (*string, bool)`

GetGroupClassNameOk returns a tuple with the GroupClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassName

`func (o *DhcpStaticDataData) SetGroupClassName(v string)`

SetGroupClassName sets GroupClassName field to given value.

### HasGroupClassName

`func (o *DhcpStaticDataData) HasGroupClassName() bool`

HasGroupClassName returns a boolean if a field has been set.

### GetGroupClassParameters

`func (o *DhcpStaticDataData) GetGroupClassParameters() []ApiClassParameterOutputEntry`

GetGroupClassParameters returns the GroupClassParameters field if non-nil, zero value otherwise.

### GetGroupClassParametersOk

`func (o *DhcpStaticDataData) GetGroupClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetGroupClassParametersOk returns a tuple with the GroupClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassParameters

`func (o *DhcpStaticDataData) SetGroupClassParameters(v []ApiClassParameterOutputEntry)`

SetGroupClassParameters sets GroupClassParameters field to given value.

### HasGroupClassParameters

`func (o *DhcpStaticDataData) HasGroupClassParameters() bool`

HasGroupClassParameters returns a boolean if a field has been set.

### GetGroupId

`func (o *DhcpStaticDataData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DhcpStaticDataData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DhcpStaticDataData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DhcpStaticDataData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *DhcpStaticDataData) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DhcpStaticDataData) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DhcpStaticDataData) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DhcpStaticDataData) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStaticAddr

`func (o *DhcpStaticDataData) GetStaticAddr() string`

GetStaticAddr returns the StaticAddr field if non-nil, zero value otherwise.

### GetStaticAddrOk

`func (o *DhcpStaticDataData) GetStaticAddrOk() (*string, bool)`

GetStaticAddrOk returns a tuple with the StaticAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAddr

`func (o *DhcpStaticDataData) SetStaticAddr(v string)`

SetStaticAddr sets StaticAddr field to given value.

### HasStaticAddr

`func (o *DhcpStaticDataData) HasStaticAddr() bool`

HasStaticAddr returns a boolean if a field has been set.

### GetStaticClassName

`func (o *DhcpStaticDataData) GetStaticClassName() string`

GetStaticClassName returns the StaticClassName field if non-nil, zero value otherwise.

### GetStaticClassNameOk

`func (o *DhcpStaticDataData) GetStaticClassNameOk() (*string, bool)`

GetStaticClassNameOk returns a tuple with the StaticClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticClassName

`func (o *DhcpStaticDataData) SetStaticClassName(v string)`

SetStaticClassName sets StaticClassName field to given value.

### HasStaticClassName

`func (o *DhcpStaticDataData) HasStaticClassName() bool`

HasStaticClassName returns a boolean if a field has been set.

### GetStaticClassParameters

`func (o *DhcpStaticDataData) GetStaticClassParameters() []ApiClassParameterOutputEntry`

GetStaticClassParameters returns the StaticClassParameters field if non-nil, zero value otherwise.

### GetStaticClassParametersOk

`func (o *DhcpStaticDataData) GetStaticClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetStaticClassParametersOk returns a tuple with the StaticClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticClassParameters

`func (o *DhcpStaticDataData) SetStaticClassParameters(v []ApiClassParameterOutputEntry)`

SetStaticClassParameters sets StaticClassParameters field to given value.

### HasStaticClassParameters

`func (o *DhcpStaticDataData) HasStaticClassParameters() bool`

HasStaticClassParameters returns a boolean if a field has been set.

### GetStaticDomain

`func (o *DhcpStaticDataData) GetStaticDomain() string`

GetStaticDomain returns the StaticDomain field if non-nil, zero value otherwise.

### GetStaticDomainOk

`func (o *DhcpStaticDataData) GetStaticDomainOk() (*string, bool)`

GetStaticDomainOk returns a tuple with the StaticDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDomain

`func (o *DhcpStaticDataData) SetStaticDomain(v string)`

SetStaticDomain sets StaticDomain field to given value.

### HasStaticDomain

`func (o *DhcpStaticDataData) HasStaticDomain() bool`

HasStaticDomain returns a boolean if a field has been set.

### GetStaticExpireTime

`func (o *DhcpStaticDataData) GetStaticExpireTime() string`

GetStaticExpireTime returns the StaticExpireTime field if non-nil, zero value otherwise.

### GetStaticExpireTimeOk

`func (o *DhcpStaticDataData) GetStaticExpireTimeOk() (*string, bool)`

GetStaticExpireTimeOk returns a tuple with the StaticExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticExpireTime

`func (o *DhcpStaticDataData) SetStaticExpireTime(v string)`

SetStaticExpireTime sets StaticExpireTime field to given value.

### HasStaticExpireTime

`func (o *DhcpStaticDataData) HasStaticExpireTime() bool`

HasStaticExpireTime returns a boolean if a field has been set.

### GetStaticId

`func (o *DhcpStaticDataData) GetStaticId() string`

GetStaticId returns the StaticId field if non-nil, zero value otherwise.

### GetStaticIdOk

`func (o *DhcpStaticDataData) GetStaticIdOk() (*string, bool)`

GetStaticIdOk returns a tuple with the StaticId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticId

`func (o *DhcpStaticDataData) SetStaticId(v string)`

SetStaticId sets StaticId field to given value.

### HasStaticId

`func (o *DhcpStaticDataData) HasStaticId() bool`

HasStaticId returns a boolean if a field has been set.

### GetStaticIdentifier

`func (o *DhcpStaticDataData) GetStaticIdentifier() string`

GetStaticIdentifier returns the StaticIdentifier field if non-nil, zero value otherwise.

### GetStaticIdentifierOk

`func (o *DhcpStaticDataData) GetStaticIdentifierOk() (*string, bool)`

GetStaticIdentifierOk returns a tuple with the StaticIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIdentifier

`func (o *DhcpStaticDataData) SetStaticIdentifier(v string)`

SetStaticIdentifier sets StaticIdentifier field to given value.

### HasStaticIdentifier

`func (o *DhcpStaticDataData) HasStaticIdentifier() bool`

HasStaticIdentifier returns a boolean if a field has been set.

### GetStaticAddressAddr

`func (o *DhcpStaticDataData) GetStaticAddressAddr() string`

GetStaticAddressAddr returns the StaticAddressAddr field if non-nil, zero value otherwise.

### GetStaticAddressAddrOk

`func (o *DhcpStaticDataData) GetStaticAddressAddrOk() (*string, bool)`

GetStaticAddressAddrOk returns a tuple with the StaticAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAddressAddr

`func (o *DhcpStaticDataData) SetStaticAddressAddr(v string)`

SetStaticAddressAddr sets StaticAddressAddr field to given value.

### HasStaticAddressAddr

`func (o *DhcpStaticDataData) HasStaticAddressAddr() bool`

HasStaticAddressAddr returns a boolean if a field has been set.

### GetStaticLastSeen

`func (o *DhcpStaticDataData) GetStaticLastSeen() string`

GetStaticLastSeen returns the StaticLastSeen field if non-nil, zero value otherwise.

### GetStaticLastSeenOk

`func (o *DhcpStaticDataData) GetStaticLastSeenOk() (*string, bool)`

GetStaticLastSeenOk returns a tuple with the StaticLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticLastSeen

`func (o *DhcpStaticDataData) SetStaticLastSeen(v string)`

SetStaticLastSeen sets StaticLastSeen field to given value.

### HasStaticLastSeen

`func (o *DhcpStaticDataData) HasStaticLastSeen() bool`

HasStaticLastSeen returns a boolean if a field has been set.

### GetStaticMacAddr

`func (o *DhcpStaticDataData) GetStaticMacAddr() string`

GetStaticMacAddr returns the StaticMacAddr field if non-nil, zero value otherwise.

### GetStaticMacAddrOk

`func (o *DhcpStaticDataData) GetStaticMacAddrOk() (*string, bool)`

GetStaticMacAddrOk returns a tuple with the StaticMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacAddr

`func (o *DhcpStaticDataData) SetStaticMacAddr(v string)`

SetStaticMacAddr sets StaticMacAddr field to given value.

### HasStaticMacAddr

`func (o *DhcpStaticDataData) HasStaticMacAddr() bool`

HasStaticMacAddr returns a boolean if a field has been set.

### GetStaticName

`func (o *DhcpStaticDataData) GetStaticName() string`

GetStaticName returns the StaticName field if non-nil, zero value otherwise.

### GetStaticNameOk

`func (o *DhcpStaticDataData) GetStaticNameOk() (*string, bool)`

GetStaticNameOk returns a tuple with the StaticName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticName

`func (o *DhcpStaticDataData) SetStaticName(v string)`

SetStaticName sets StaticName field to given value.

### HasStaticName

`func (o *DhcpStaticDataData) HasStaticName() bool`

HasStaticName returns a boolean if a field has been set.

### GetScopeClassName

`func (o *DhcpStaticDataData) GetScopeClassName() string`

GetScopeClassName returns the ScopeClassName field if non-nil, zero value otherwise.

### GetScopeClassNameOk

`func (o *DhcpStaticDataData) GetScopeClassNameOk() (*string, bool)`

GetScopeClassNameOk returns a tuple with the ScopeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassName

`func (o *DhcpStaticDataData) SetScopeClassName(v string)`

SetScopeClassName sets ScopeClassName field to given value.

### HasScopeClassName

`func (o *DhcpStaticDataData) HasScopeClassName() bool`

HasScopeClassName returns a boolean if a field has been set.

### GetScopeClassParameters

`func (o *DhcpStaticDataData) GetScopeClassParameters() []ApiClassParameterOutputEntry`

GetScopeClassParameters returns the ScopeClassParameters field if non-nil, zero value otherwise.

### GetScopeClassParametersOk

`func (o *DhcpStaticDataData) GetScopeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScopeClassParametersOk returns a tuple with the ScopeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassParameters

`func (o *DhcpStaticDataData) SetScopeClassParameters(v []ApiClassParameterOutputEntry)`

SetScopeClassParameters sets ScopeClassParameters field to given value.

### HasScopeClassParameters

`func (o *DhcpStaticDataData) HasScopeClassParameters() bool`

HasScopeClassParameters returns a boolean if a field has been set.

### GetScopeEndAddressAddr

`func (o *DhcpStaticDataData) GetScopeEndAddressAddr() string`

GetScopeEndAddressAddr returns the ScopeEndAddressAddr field if non-nil, zero value otherwise.

### GetScopeEndAddressAddrOk

`func (o *DhcpStaticDataData) GetScopeEndAddressAddrOk() (*string, bool)`

GetScopeEndAddressAddrOk returns a tuple with the ScopeEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeEndAddressAddr

`func (o *DhcpStaticDataData) SetScopeEndAddressAddr(v string)`

SetScopeEndAddressAddr sets ScopeEndAddressAddr field to given value.

### HasScopeEndAddressAddr

`func (o *DhcpStaticDataData) HasScopeEndAddressAddr() bool`

HasScopeEndAddressAddr returns a boolean if a field has been set.

### GetScopeId

`func (o *DhcpStaticDataData) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DhcpStaticDataData) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DhcpStaticDataData) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *DhcpStaticDataData) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetScopeName

`func (o *DhcpStaticDataData) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DhcpStaticDataData) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DhcpStaticDataData) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DhcpStaticDataData) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeNetAddr

`func (o *DhcpStaticDataData) GetScopeNetAddr() string`

GetScopeNetAddr returns the ScopeNetAddr field if non-nil, zero value otherwise.

### GetScopeNetAddrOk

`func (o *DhcpStaticDataData) GetScopeNetAddrOk() (*string, bool)`

GetScopeNetAddrOk returns a tuple with the ScopeNetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetAddr

`func (o *DhcpStaticDataData) SetScopeNetAddr(v string)`

SetScopeNetAddr sets ScopeNetAddr field to given value.

### HasScopeNetAddr

`func (o *DhcpStaticDataData) HasScopeNetAddr() bool`

HasScopeNetAddr returns a boolean if a field has been set.

### GetScopeNetMask

`func (o *DhcpStaticDataData) GetScopeNetMask() string`

GetScopeNetMask returns the ScopeNetMask field if non-nil, zero value otherwise.

### GetScopeNetMaskOk

`func (o *DhcpStaticDataData) GetScopeNetMaskOk() (*string, bool)`

GetScopeNetMaskOk returns a tuple with the ScopeNetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetMask

`func (o *DhcpStaticDataData) SetScopeNetMask(v string)`

SetScopeNetMask sets ScopeNetMask field to given value.

### HasScopeNetMask

`func (o *DhcpStaticDataData) HasScopeNetMask() bool`

HasScopeNetMask returns a boolean if a field has been set.

### GetScopeSpaceId

`func (o *DhcpStaticDataData) GetScopeSpaceId() string`

GetScopeSpaceId returns the ScopeSpaceId field if non-nil, zero value otherwise.

### GetScopeSpaceIdOk

`func (o *DhcpStaticDataData) GetScopeSpaceIdOk() (*string, bool)`

GetScopeSpaceIdOk returns a tuple with the ScopeSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceId

`func (o *DhcpStaticDataData) SetScopeSpaceId(v string)`

SetScopeSpaceId sets ScopeSpaceId field to given value.

### HasScopeSpaceId

`func (o *DhcpStaticDataData) HasScopeSpaceId() bool`

HasScopeSpaceId returns a boolean if a field has been set.

### GetScopeSize

`func (o *DhcpStaticDataData) GetScopeSize() string`

GetScopeSize returns the ScopeSize field if non-nil, zero value otherwise.

### GetScopeSizeOk

`func (o *DhcpStaticDataData) GetScopeSizeOk() (*string, bool)`

GetScopeSizeOk returns a tuple with the ScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSize

`func (o *DhcpStaticDataData) SetScopeSize(v string)`

SetScopeSize sets ScopeSize field to given value.

### HasScopeSize

`func (o *DhcpStaticDataData) HasScopeSize() bool`

HasScopeSize returns a boolean if a field has been set.

### GetScopeStartAddressAddr

`func (o *DhcpStaticDataData) GetScopeStartAddressAddr() string`

GetScopeStartAddressAddr returns the ScopeStartAddressAddr field if non-nil, zero value otherwise.

### GetScopeStartAddressAddrOk

`func (o *DhcpStaticDataData) GetScopeStartAddressAddrOk() (*string, bool)`

GetScopeStartAddressAddrOk returns a tuple with the ScopeStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStartAddressAddr

`func (o *DhcpStaticDataData) SetScopeStartAddressAddr(v string)`

SetScopeStartAddressAddr sets ScopeStartAddressAddr field to given value.

### HasScopeStartAddressAddr

`func (o *DhcpStaticDataData) HasScopeStartAddressAddr() bool`

HasScopeStartAddressAddr returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DhcpStaticDataData) GetSharednetworkId() string`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DhcpStaticDataData) GetSharednetworkIdOk() (*string, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DhcpStaticDataData) SetSharednetworkId(v string)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DhcpStaticDataData) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DhcpStaticDataData) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DhcpStaticDataData) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DhcpStaticDataData) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DhcpStaticDataData) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetStaticMacVendor

`func (o *DhcpStaticDataData) GetStaticMacVendor() string`

GetStaticMacVendor returns the StaticMacVendor field if non-nil, zero value otherwise.

### GetStaticMacVendorOk

`func (o *DhcpStaticDataData) GetStaticMacVendorOk() (*string, bool)`

GetStaticMacVendorOk returns a tuple with the StaticMacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacVendor

`func (o *DhcpStaticDataData) SetStaticMacVendor(v string)`

SetStaticMacVendor sets StaticMacVendor field to given value.

### HasStaticMacVendor

`func (o *DhcpStaticDataData) HasStaticMacVendor() bool`

HasStaticMacVendor returns a boolean if a field has been set.

### GetStaticMultistatus

`func (o *DhcpStaticDataData) GetStaticMultistatus() string`

GetStaticMultistatus returns the StaticMultistatus field if non-nil, zero value otherwise.

### GetStaticMultistatusOk

`func (o *DhcpStaticDataData) GetStaticMultistatusOk() (*string, bool)`

GetStaticMultistatusOk returns a tuple with the StaticMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMultistatus

`func (o *DhcpStaticDataData) SetStaticMultistatus(v string)`

SetStaticMultistatus sets StaticMultistatus field to given value.

### HasStaticMultistatus

`func (o *DhcpStaticDataData) HasStaticMultistatus() bool`

HasStaticMultistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpStaticDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpStaticDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpStaticDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpStaticDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DhcpStaticDataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DhcpStaticDataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DhcpStaticDataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DhcpStaticDataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


