# DataInnerDhcpStaticData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StaticDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**StaticDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 server the object belongs to. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft Windows DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv4 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server the object belongs to. | [optional] 
**GroupClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 group the static belongs to, it can be preceded by the class directory. | [optional] 
**GroupClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 group the static belongs to,. | [optional] 
**GroupId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 group. | [optional] 
**GroupName** | Pointer to **string** | The name of the DHCPv4 group associated with the object. | [optional] 
**StaticAddr** | Pointer to **string** | The IP address associated with the DHCPv4 static. | [optional] 
**StaticClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 static, it can be preceded by the class directory. | [optional] 
**StaticClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 static. | [optional] 
**StaticDomain** | Pointer to **string** | The domain name associated with the DHCPv4 static. | [optional] 
**StaticExpireTime** | Pointer to **string** | The expiration time of the lease associated with the DHCPv4 static, in decimal UNIX date format. | [optional] 
**StaticId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 static. | [optional] 
**StaticIdentifier** | Pointer to **string** | The host identifier of the DHCPv4 static, specified as follows: &lt;b&gt;option &amp;lt;option-name&amp;gt; expected value&lt;/b&gt;. | [optional] 
**StaticAddressAddr** | Pointer to **string** | The IP address associated with the DHCPv4 static, in hexadecimal format. | [optional] 
**StaticLastSeen** | Pointer to **string** | The last time the MAC address associated with the DHCPv4 static was seen on the network, in decimal UNIX date format. | [optional] 
**StaticMacAddr** | Pointer to **string** | The MAC address associated with the DHCPv4 static. It is composed of 7 sections, &lt;b&gt;00:11:22:33:44:55:66&lt;/b&gt;, where &lt;b&gt;00&lt;/b&gt; is the MAC address type. The type &lt;b&gt;01&lt;/b&gt; indicates Ethernet. | [optional] 
**StaticName** | Pointer to **string** | The name of the DHCPv4 static. | [optional] 
**ScopeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**ScopeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 scope the object belongs to. | [optional] 
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

### NewDataInnerDhcpStaticData

`func NewDataInnerDhcpStaticData() *DataInnerDhcpStaticData`

NewDataInnerDhcpStaticData instantiates a new DataInnerDhcpStaticData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpStaticDataWithDefaults

`func NewDataInnerDhcpStaticDataWithDefaults() *DataInnerDhcpStaticData`

NewDataInnerDhcpStaticDataWithDefaults instantiates a new DataInnerDhcpStaticData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStaticDelayedCreateTime

`func (o *DataInnerDhcpStaticData) GetStaticDelayedCreateTime() string`

GetStaticDelayedCreateTime returns the StaticDelayedCreateTime field if non-nil, zero value otherwise.

### GetStaticDelayedCreateTimeOk

`func (o *DataInnerDhcpStaticData) GetStaticDelayedCreateTimeOk() (*string, bool)`

GetStaticDelayedCreateTimeOk returns a tuple with the StaticDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDelayedCreateTime

`func (o *DataInnerDhcpStaticData) SetStaticDelayedCreateTime(v string)`

SetStaticDelayedCreateTime sets StaticDelayedCreateTime field to given value.

### HasStaticDelayedCreateTime

`func (o *DataInnerDhcpStaticData) HasStaticDelayedCreateTime() bool`

HasStaticDelayedCreateTime returns a boolean if a field has been set.

### GetStaticDelayedDeleteTime

`func (o *DataInnerDhcpStaticData) GetStaticDelayedDeleteTime() string`

GetStaticDelayedDeleteTime returns the StaticDelayedDeleteTime field if non-nil, zero value otherwise.

### GetStaticDelayedDeleteTimeOk

`func (o *DataInnerDhcpStaticData) GetStaticDelayedDeleteTimeOk() (*string, bool)`

GetStaticDelayedDeleteTimeOk returns a tuple with the StaticDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDelayedDeleteTime

`func (o *DataInnerDhcpStaticData) SetStaticDelayedDeleteTime(v string)`

SetStaticDelayedDeleteTime sets StaticDelayedDeleteTime field to given value.

### HasStaticDelayedDeleteTime

`func (o *DataInnerDhcpStaticData) HasStaticDelayedDeleteTime() bool`

HasStaticDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DataInnerDhcpStaticData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DataInnerDhcpStaticData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DataInnerDhcpStaticData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DataInnerDhcpStaticData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DataInnerDhcpStaticData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DataInnerDhcpStaticData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DataInnerDhcpStaticData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DataInnerDhcpStaticData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDhcpStaticData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDhcpStaticData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDhcpStaticData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDhcpStaticData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDhcpStaticData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDhcpStaticData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDhcpStaticData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDhcpStaticData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDhcpStaticData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDhcpStaticData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDhcpStaticData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDhcpStaticData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DataInnerDhcpStaticData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DataInnerDhcpStaticData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DataInnerDhcpStaticData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DataInnerDhcpStaticData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetGroupClassName

`func (o *DataInnerDhcpStaticData) GetGroupClassName() string`

GetGroupClassName returns the GroupClassName field if non-nil, zero value otherwise.

### GetGroupClassNameOk

`func (o *DataInnerDhcpStaticData) GetGroupClassNameOk() (*string, bool)`

GetGroupClassNameOk returns a tuple with the GroupClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassName

`func (o *DataInnerDhcpStaticData) SetGroupClassName(v string)`

SetGroupClassName sets GroupClassName field to given value.

### HasGroupClassName

`func (o *DataInnerDhcpStaticData) HasGroupClassName() bool`

HasGroupClassName returns a boolean if a field has been set.

### GetGroupClassParameters

`func (o *DataInnerDhcpStaticData) GetGroupClassParameters() []ApiClassParameterOutputEntry`

GetGroupClassParameters returns the GroupClassParameters field if non-nil, zero value otherwise.

### GetGroupClassParametersOk

`func (o *DataInnerDhcpStaticData) GetGroupClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetGroupClassParametersOk returns a tuple with the GroupClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassParameters

`func (o *DataInnerDhcpStaticData) SetGroupClassParameters(v []ApiClassParameterOutputEntry)`

SetGroupClassParameters sets GroupClassParameters field to given value.

### HasGroupClassParameters

`func (o *DataInnerDhcpStaticData) HasGroupClassParameters() bool`

HasGroupClassParameters returns a boolean if a field has been set.

### GetGroupId

`func (o *DataInnerDhcpStaticData) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DataInnerDhcpStaticData) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DataInnerDhcpStaticData) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DataInnerDhcpStaticData) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *DataInnerDhcpStaticData) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DataInnerDhcpStaticData) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DataInnerDhcpStaticData) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DataInnerDhcpStaticData) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStaticAddr

`func (o *DataInnerDhcpStaticData) GetStaticAddr() string`

GetStaticAddr returns the StaticAddr field if non-nil, zero value otherwise.

### GetStaticAddrOk

`func (o *DataInnerDhcpStaticData) GetStaticAddrOk() (*string, bool)`

GetStaticAddrOk returns a tuple with the StaticAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAddr

`func (o *DataInnerDhcpStaticData) SetStaticAddr(v string)`

SetStaticAddr sets StaticAddr field to given value.

### HasStaticAddr

`func (o *DataInnerDhcpStaticData) HasStaticAddr() bool`

HasStaticAddr returns a boolean if a field has been set.

### GetStaticClassName

`func (o *DataInnerDhcpStaticData) GetStaticClassName() string`

GetStaticClassName returns the StaticClassName field if non-nil, zero value otherwise.

### GetStaticClassNameOk

`func (o *DataInnerDhcpStaticData) GetStaticClassNameOk() (*string, bool)`

GetStaticClassNameOk returns a tuple with the StaticClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticClassName

`func (o *DataInnerDhcpStaticData) SetStaticClassName(v string)`

SetStaticClassName sets StaticClassName field to given value.

### HasStaticClassName

`func (o *DataInnerDhcpStaticData) HasStaticClassName() bool`

HasStaticClassName returns a boolean if a field has been set.

### GetStaticClassParameters

`func (o *DataInnerDhcpStaticData) GetStaticClassParameters() []ApiClassParameterOutputEntry`

GetStaticClassParameters returns the StaticClassParameters field if non-nil, zero value otherwise.

### GetStaticClassParametersOk

`func (o *DataInnerDhcpStaticData) GetStaticClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetStaticClassParametersOk returns a tuple with the StaticClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticClassParameters

`func (o *DataInnerDhcpStaticData) SetStaticClassParameters(v []ApiClassParameterOutputEntry)`

SetStaticClassParameters sets StaticClassParameters field to given value.

### HasStaticClassParameters

`func (o *DataInnerDhcpStaticData) HasStaticClassParameters() bool`

HasStaticClassParameters returns a boolean if a field has been set.

### GetStaticDomain

`func (o *DataInnerDhcpStaticData) GetStaticDomain() string`

GetStaticDomain returns the StaticDomain field if non-nil, zero value otherwise.

### GetStaticDomainOk

`func (o *DataInnerDhcpStaticData) GetStaticDomainOk() (*string, bool)`

GetStaticDomainOk returns a tuple with the StaticDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticDomain

`func (o *DataInnerDhcpStaticData) SetStaticDomain(v string)`

SetStaticDomain sets StaticDomain field to given value.

### HasStaticDomain

`func (o *DataInnerDhcpStaticData) HasStaticDomain() bool`

HasStaticDomain returns a boolean if a field has been set.

### GetStaticExpireTime

`func (o *DataInnerDhcpStaticData) GetStaticExpireTime() string`

GetStaticExpireTime returns the StaticExpireTime field if non-nil, zero value otherwise.

### GetStaticExpireTimeOk

`func (o *DataInnerDhcpStaticData) GetStaticExpireTimeOk() (*string, bool)`

GetStaticExpireTimeOk returns a tuple with the StaticExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticExpireTime

`func (o *DataInnerDhcpStaticData) SetStaticExpireTime(v string)`

SetStaticExpireTime sets StaticExpireTime field to given value.

### HasStaticExpireTime

`func (o *DataInnerDhcpStaticData) HasStaticExpireTime() bool`

HasStaticExpireTime returns a boolean if a field has been set.

### GetStaticId

`func (o *DataInnerDhcpStaticData) GetStaticId() string`

GetStaticId returns the StaticId field if non-nil, zero value otherwise.

### GetStaticIdOk

`func (o *DataInnerDhcpStaticData) GetStaticIdOk() (*string, bool)`

GetStaticIdOk returns a tuple with the StaticId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticId

`func (o *DataInnerDhcpStaticData) SetStaticId(v string)`

SetStaticId sets StaticId field to given value.

### HasStaticId

`func (o *DataInnerDhcpStaticData) HasStaticId() bool`

HasStaticId returns a boolean if a field has been set.

### GetStaticIdentifier

`func (o *DataInnerDhcpStaticData) GetStaticIdentifier() string`

GetStaticIdentifier returns the StaticIdentifier field if non-nil, zero value otherwise.

### GetStaticIdentifierOk

`func (o *DataInnerDhcpStaticData) GetStaticIdentifierOk() (*string, bool)`

GetStaticIdentifierOk returns a tuple with the StaticIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIdentifier

`func (o *DataInnerDhcpStaticData) SetStaticIdentifier(v string)`

SetStaticIdentifier sets StaticIdentifier field to given value.

### HasStaticIdentifier

`func (o *DataInnerDhcpStaticData) HasStaticIdentifier() bool`

HasStaticIdentifier returns a boolean if a field has been set.

### GetStaticAddressAddr

`func (o *DataInnerDhcpStaticData) GetStaticAddressAddr() string`

GetStaticAddressAddr returns the StaticAddressAddr field if non-nil, zero value otherwise.

### GetStaticAddressAddrOk

`func (o *DataInnerDhcpStaticData) GetStaticAddressAddrOk() (*string, bool)`

GetStaticAddressAddrOk returns a tuple with the StaticAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAddressAddr

`func (o *DataInnerDhcpStaticData) SetStaticAddressAddr(v string)`

SetStaticAddressAddr sets StaticAddressAddr field to given value.

### HasStaticAddressAddr

`func (o *DataInnerDhcpStaticData) HasStaticAddressAddr() bool`

HasStaticAddressAddr returns a boolean if a field has been set.

### GetStaticLastSeen

`func (o *DataInnerDhcpStaticData) GetStaticLastSeen() string`

GetStaticLastSeen returns the StaticLastSeen field if non-nil, zero value otherwise.

### GetStaticLastSeenOk

`func (o *DataInnerDhcpStaticData) GetStaticLastSeenOk() (*string, bool)`

GetStaticLastSeenOk returns a tuple with the StaticLastSeen field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticLastSeen

`func (o *DataInnerDhcpStaticData) SetStaticLastSeen(v string)`

SetStaticLastSeen sets StaticLastSeen field to given value.

### HasStaticLastSeen

`func (o *DataInnerDhcpStaticData) HasStaticLastSeen() bool`

HasStaticLastSeen returns a boolean if a field has been set.

### GetStaticMacAddr

`func (o *DataInnerDhcpStaticData) GetStaticMacAddr() string`

GetStaticMacAddr returns the StaticMacAddr field if non-nil, zero value otherwise.

### GetStaticMacAddrOk

`func (o *DataInnerDhcpStaticData) GetStaticMacAddrOk() (*string, bool)`

GetStaticMacAddrOk returns a tuple with the StaticMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacAddr

`func (o *DataInnerDhcpStaticData) SetStaticMacAddr(v string)`

SetStaticMacAddr sets StaticMacAddr field to given value.

### HasStaticMacAddr

`func (o *DataInnerDhcpStaticData) HasStaticMacAddr() bool`

HasStaticMacAddr returns a boolean if a field has been set.

### GetStaticName

`func (o *DataInnerDhcpStaticData) GetStaticName() string`

GetStaticName returns the StaticName field if non-nil, zero value otherwise.

### GetStaticNameOk

`func (o *DataInnerDhcpStaticData) GetStaticNameOk() (*string, bool)`

GetStaticNameOk returns a tuple with the StaticName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticName

`func (o *DataInnerDhcpStaticData) SetStaticName(v string)`

SetStaticName sets StaticName field to given value.

### HasStaticName

`func (o *DataInnerDhcpStaticData) HasStaticName() bool`

HasStaticName returns a boolean if a field has been set.

### GetScopeClassName

`func (o *DataInnerDhcpStaticData) GetScopeClassName() string`

GetScopeClassName returns the ScopeClassName field if non-nil, zero value otherwise.

### GetScopeClassNameOk

`func (o *DataInnerDhcpStaticData) GetScopeClassNameOk() (*string, bool)`

GetScopeClassNameOk returns a tuple with the ScopeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassName

`func (o *DataInnerDhcpStaticData) SetScopeClassName(v string)`

SetScopeClassName sets ScopeClassName field to given value.

### HasScopeClassName

`func (o *DataInnerDhcpStaticData) HasScopeClassName() bool`

HasScopeClassName returns a boolean if a field has been set.

### GetScopeClassParameters

`func (o *DataInnerDhcpStaticData) GetScopeClassParameters() []ApiClassParameterOutputEntry`

GetScopeClassParameters returns the ScopeClassParameters field if non-nil, zero value otherwise.

### GetScopeClassParametersOk

`func (o *DataInnerDhcpStaticData) GetScopeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScopeClassParametersOk returns a tuple with the ScopeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassParameters

`func (o *DataInnerDhcpStaticData) SetScopeClassParameters(v []ApiClassParameterOutputEntry)`

SetScopeClassParameters sets ScopeClassParameters field to given value.

### HasScopeClassParameters

`func (o *DataInnerDhcpStaticData) HasScopeClassParameters() bool`

HasScopeClassParameters returns a boolean if a field has been set.

### GetScopeEndAddressAddr

`func (o *DataInnerDhcpStaticData) GetScopeEndAddressAddr() string`

GetScopeEndAddressAddr returns the ScopeEndAddressAddr field if non-nil, zero value otherwise.

### GetScopeEndAddressAddrOk

`func (o *DataInnerDhcpStaticData) GetScopeEndAddressAddrOk() (*string, bool)`

GetScopeEndAddressAddrOk returns a tuple with the ScopeEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeEndAddressAddr

`func (o *DataInnerDhcpStaticData) SetScopeEndAddressAddr(v string)`

SetScopeEndAddressAddr sets ScopeEndAddressAddr field to given value.

### HasScopeEndAddressAddr

`func (o *DataInnerDhcpStaticData) HasScopeEndAddressAddr() bool`

HasScopeEndAddressAddr returns a boolean if a field has been set.

### GetScopeId

`func (o *DataInnerDhcpStaticData) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DataInnerDhcpStaticData) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DataInnerDhcpStaticData) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *DataInnerDhcpStaticData) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetScopeName

`func (o *DataInnerDhcpStaticData) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DataInnerDhcpStaticData) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DataInnerDhcpStaticData) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DataInnerDhcpStaticData) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeNetAddr

`func (o *DataInnerDhcpStaticData) GetScopeNetAddr() string`

GetScopeNetAddr returns the ScopeNetAddr field if non-nil, zero value otherwise.

### GetScopeNetAddrOk

`func (o *DataInnerDhcpStaticData) GetScopeNetAddrOk() (*string, bool)`

GetScopeNetAddrOk returns a tuple with the ScopeNetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetAddr

`func (o *DataInnerDhcpStaticData) SetScopeNetAddr(v string)`

SetScopeNetAddr sets ScopeNetAddr field to given value.

### HasScopeNetAddr

`func (o *DataInnerDhcpStaticData) HasScopeNetAddr() bool`

HasScopeNetAddr returns a boolean if a field has been set.

### GetScopeNetMask

`func (o *DataInnerDhcpStaticData) GetScopeNetMask() string`

GetScopeNetMask returns the ScopeNetMask field if non-nil, zero value otherwise.

### GetScopeNetMaskOk

`func (o *DataInnerDhcpStaticData) GetScopeNetMaskOk() (*string, bool)`

GetScopeNetMaskOk returns a tuple with the ScopeNetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetMask

`func (o *DataInnerDhcpStaticData) SetScopeNetMask(v string)`

SetScopeNetMask sets ScopeNetMask field to given value.

### HasScopeNetMask

`func (o *DataInnerDhcpStaticData) HasScopeNetMask() bool`

HasScopeNetMask returns a boolean if a field has been set.

### GetScopeSpaceId

`func (o *DataInnerDhcpStaticData) GetScopeSpaceId() string`

GetScopeSpaceId returns the ScopeSpaceId field if non-nil, zero value otherwise.

### GetScopeSpaceIdOk

`func (o *DataInnerDhcpStaticData) GetScopeSpaceIdOk() (*string, bool)`

GetScopeSpaceIdOk returns a tuple with the ScopeSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceId

`func (o *DataInnerDhcpStaticData) SetScopeSpaceId(v string)`

SetScopeSpaceId sets ScopeSpaceId field to given value.

### HasScopeSpaceId

`func (o *DataInnerDhcpStaticData) HasScopeSpaceId() bool`

HasScopeSpaceId returns a boolean if a field has been set.

### GetScopeSize

`func (o *DataInnerDhcpStaticData) GetScopeSize() string`

GetScopeSize returns the ScopeSize field if non-nil, zero value otherwise.

### GetScopeSizeOk

`func (o *DataInnerDhcpStaticData) GetScopeSizeOk() (*string, bool)`

GetScopeSizeOk returns a tuple with the ScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSize

`func (o *DataInnerDhcpStaticData) SetScopeSize(v string)`

SetScopeSize sets ScopeSize field to given value.

### HasScopeSize

`func (o *DataInnerDhcpStaticData) HasScopeSize() bool`

HasScopeSize returns a boolean if a field has been set.

### GetScopeStartAddressAddr

`func (o *DataInnerDhcpStaticData) GetScopeStartAddressAddr() string`

GetScopeStartAddressAddr returns the ScopeStartAddressAddr field if non-nil, zero value otherwise.

### GetScopeStartAddressAddrOk

`func (o *DataInnerDhcpStaticData) GetScopeStartAddressAddrOk() (*string, bool)`

GetScopeStartAddressAddrOk returns a tuple with the ScopeStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStartAddressAddr

`func (o *DataInnerDhcpStaticData) SetScopeStartAddressAddr(v string)`

SetScopeStartAddressAddr sets ScopeStartAddressAddr field to given value.

### HasScopeStartAddressAddr

`func (o *DataInnerDhcpStaticData) HasScopeStartAddressAddr() bool`

HasScopeStartAddressAddr returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DataInnerDhcpStaticData) GetSharednetworkId() string`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DataInnerDhcpStaticData) GetSharednetworkIdOk() (*string, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DataInnerDhcpStaticData) SetSharednetworkId(v string)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DataInnerDhcpStaticData) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DataInnerDhcpStaticData) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DataInnerDhcpStaticData) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DataInnerDhcpStaticData) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DataInnerDhcpStaticData) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetStaticMacVendor

`func (o *DataInnerDhcpStaticData) GetStaticMacVendor() string`

GetStaticMacVendor returns the StaticMacVendor field if non-nil, zero value otherwise.

### GetStaticMacVendorOk

`func (o *DataInnerDhcpStaticData) GetStaticMacVendorOk() (*string, bool)`

GetStaticMacVendorOk returns a tuple with the StaticMacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacVendor

`func (o *DataInnerDhcpStaticData) SetStaticMacVendor(v string)`

SetStaticMacVendor sets StaticMacVendor field to given value.

### HasStaticMacVendor

`func (o *DataInnerDhcpStaticData) HasStaticMacVendor() bool`

HasStaticMacVendor returns a boolean if a field has been set.

### GetStaticMultistatus

`func (o *DataInnerDhcpStaticData) GetStaticMultistatus() string`

GetStaticMultistatus returns the StaticMultistatus field if non-nil, zero value otherwise.

### GetStaticMultistatusOk

`func (o *DataInnerDhcpStaticData) GetStaticMultistatusOk() (*string, bool)`

GetStaticMultistatusOk returns a tuple with the StaticMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMultistatus

`func (o *DataInnerDhcpStaticData) SetStaticMultistatus(v string)`

SetStaticMultistatus sets StaticMultistatus field to given value.

### HasStaticMultistatus

`func (o *DataInnerDhcpStaticData) HasStaticMultistatus() bool`

HasStaticMultistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpStaticData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpStaticData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpStaticData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpStaticData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpStaticData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpStaticData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpStaticData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpStaticData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


