# DhcpScopeDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**ScopeDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server the object belongs to. | [optional] 
**FailoverId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 failover channel associated with the object. | [optional] 
**FailoverName** | Pointer to **string** | The name of the DHCPv4 failover channel associated with the object. | [optional] 
**ScopeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 scope, it can be preceded by the class directory. | [optional] 
**ScopeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ScopeEndAddressAddr** | Pointer to **string** | The last IP address of the DHCPv4 scope, in hexadecimal format. | [optional] 
**ScopeId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 scope. | [optional] 
**ScopeName** | Pointer to **string** | The name of the DHCPv4 scope. | [optional] 
**ScopeNetAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope. | [optional] 
**ScopeNetMask** | Pointer to **string** | The netmask of the DHCPv4 scope. It is expressed in dot-decimal notation and defines the number of addresses the scope contains. | [optional] 
**ScopeSpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the DHCPv4 scope. | [optional] 
**ScopeSpaceName** | Pointer to **string** | The name of the space associated with the DHCPv4 scope. | [optional] 
**ScopeSize** | Pointer to **string** | The number of IP addresses the DHCPv4 scope contains. | [optional] 
**ScopeStartAddressAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope, in hexadecimal format. | [optional] 
**SharednetworkId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 shared network the object belongs to. | [optional] 
**SharednetworkName** | Pointer to **string** | The name of the DHCPv4 shared network the object belongs to. | [optional] 
**ServerAddr** | Pointer to **string** | The IP address of the DHCP server the object belongs to, in hexadecimal format. | [optional] 
**ScopeMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartArch** | Pointer to **string** | The type of the DHCPv4 smart architecture the object belongs to.&lt;table&gt;&lt;caption&gt;vdhcp_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;masterslave&lt;/td&gt;&lt;td &gt;The One-to-One smart architecture sets a pair of DHCP servers in a Master/Backup configuration.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;star&lt;/td&gt;&lt;td &gt;The One-to-Many smart architecture sets a multi-site failover configuration at the cost of n-servers+1.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;splitscope&lt;/td&gt;&lt;td &gt;The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;The Single-Server smart architecture manages a single DHCP server.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpScopeDataData

`func NewDhcpScopeDataData() *DhcpScopeDataData`

NewDhcpScopeDataData instantiates a new DhcpScopeDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpScopeDataDataWithDefaults

`func NewDhcpScopeDataDataWithDefaults() *DhcpScopeDataData`

NewDhcpScopeDataDataWithDefaults instantiates a new DhcpScopeDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeDelayedCreateTime

`func (o *DhcpScopeDataData) GetScopeDelayedCreateTime() string`

GetScopeDelayedCreateTime returns the ScopeDelayedCreateTime field if non-nil, zero value otherwise.

### GetScopeDelayedCreateTimeOk

`func (o *DhcpScopeDataData) GetScopeDelayedCreateTimeOk() (*string, bool)`

GetScopeDelayedCreateTimeOk returns a tuple with the ScopeDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDelayedCreateTime

`func (o *DhcpScopeDataData) SetScopeDelayedCreateTime(v string)`

SetScopeDelayedCreateTime sets ScopeDelayedCreateTime field to given value.

### HasScopeDelayedCreateTime

`func (o *DhcpScopeDataData) HasScopeDelayedCreateTime() bool`

HasScopeDelayedCreateTime returns a boolean if a field has been set.

### GetScopeDelayedDeleteTime

`func (o *DhcpScopeDataData) GetScopeDelayedDeleteTime() string`

GetScopeDelayedDeleteTime returns the ScopeDelayedDeleteTime field if non-nil, zero value otherwise.

### GetScopeDelayedDeleteTimeOk

`func (o *DhcpScopeDataData) GetScopeDelayedDeleteTimeOk() (*string, bool)`

GetScopeDelayedDeleteTimeOk returns a tuple with the ScopeDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDelayedDeleteTime

`func (o *DhcpScopeDataData) SetScopeDelayedDeleteTime(v string)`

SetScopeDelayedDeleteTime sets ScopeDelayedDeleteTime field to given value.

### HasScopeDelayedDeleteTime

`func (o *DhcpScopeDataData) HasScopeDelayedDeleteTime() bool`

HasScopeDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DhcpScopeDataData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DhcpScopeDataData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DhcpScopeDataData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DhcpScopeDataData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DhcpScopeDataData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DhcpScopeDataData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DhcpScopeDataData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DhcpScopeDataData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpScopeDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpScopeDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpScopeDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpScopeDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpScopeDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpScopeDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpScopeDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpScopeDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpScopeDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpScopeDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpScopeDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpScopeDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DhcpScopeDataData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DhcpScopeDataData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DhcpScopeDataData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DhcpScopeDataData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetFailoverId

`func (o *DhcpScopeDataData) GetFailoverId() string`

GetFailoverId returns the FailoverId field if non-nil, zero value otherwise.

### GetFailoverIdOk

`func (o *DhcpScopeDataData) GetFailoverIdOk() (*string, bool)`

GetFailoverIdOk returns a tuple with the FailoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverId

`func (o *DhcpScopeDataData) SetFailoverId(v string)`

SetFailoverId sets FailoverId field to given value.

### HasFailoverId

`func (o *DhcpScopeDataData) HasFailoverId() bool`

HasFailoverId returns a boolean if a field has been set.

### GetFailoverName

`func (o *DhcpScopeDataData) GetFailoverName() string`

GetFailoverName returns the FailoverName field if non-nil, zero value otherwise.

### GetFailoverNameOk

`func (o *DhcpScopeDataData) GetFailoverNameOk() (*string, bool)`

GetFailoverNameOk returns a tuple with the FailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverName

`func (o *DhcpScopeDataData) SetFailoverName(v string)`

SetFailoverName sets FailoverName field to given value.

### HasFailoverName

`func (o *DhcpScopeDataData) HasFailoverName() bool`

HasFailoverName returns a boolean if a field has been set.

### GetScopeClassName

`func (o *DhcpScopeDataData) GetScopeClassName() string`

GetScopeClassName returns the ScopeClassName field if non-nil, zero value otherwise.

### GetScopeClassNameOk

`func (o *DhcpScopeDataData) GetScopeClassNameOk() (*string, bool)`

GetScopeClassNameOk returns a tuple with the ScopeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassName

`func (o *DhcpScopeDataData) SetScopeClassName(v string)`

SetScopeClassName sets ScopeClassName field to given value.

### HasScopeClassName

`func (o *DhcpScopeDataData) HasScopeClassName() bool`

HasScopeClassName returns a boolean if a field has been set.

### GetScopeClassParameters

`func (o *DhcpScopeDataData) GetScopeClassParameters() []ApiClassParameterOutputEntry`

GetScopeClassParameters returns the ScopeClassParameters field if non-nil, zero value otherwise.

### GetScopeClassParametersOk

`func (o *DhcpScopeDataData) GetScopeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScopeClassParametersOk returns a tuple with the ScopeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassParameters

`func (o *DhcpScopeDataData) SetScopeClassParameters(v []ApiClassParameterOutputEntry)`

SetScopeClassParameters sets ScopeClassParameters field to given value.

### HasScopeClassParameters

`func (o *DhcpScopeDataData) HasScopeClassParameters() bool`

HasScopeClassParameters returns a boolean if a field has been set.

### GetScopeEndAddressAddr

`func (o *DhcpScopeDataData) GetScopeEndAddressAddr() string`

GetScopeEndAddressAddr returns the ScopeEndAddressAddr field if non-nil, zero value otherwise.

### GetScopeEndAddressAddrOk

`func (o *DhcpScopeDataData) GetScopeEndAddressAddrOk() (*string, bool)`

GetScopeEndAddressAddrOk returns a tuple with the ScopeEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeEndAddressAddr

`func (o *DhcpScopeDataData) SetScopeEndAddressAddr(v string)`

SetScopeEndAddressAddr sets ScopeEndAddressAddr field to given value.

### HasScopeEndAddressAddr

`func (o *DhcpScopeDataData) HasScopeEndAddressAddr() bool`

HasScopeEndAddressAddr returns a boolean if a field has been set.

### GetScopeId

`func (o *DhcpScopeDataData) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DhcpScopeDataData) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DhcpScopeDataData) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *DhcpScopeDataData) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetScopeName

`func (o *DhcpScopeDataData) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DhcpScopeDataData) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DhcpScopeDataData) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DhcpScopeDataData) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeNetAddr

`func (o *DhcpScopeDataData) GetScopeNetAddr() string`

GetScopeNetAddr returns the ScopeNetAddr field if non-nil, zero value otherwise.

### GetScopeNetAddrOk

`func (o *DhcpScopeDataData) GetScopeNetAddrOk() (*string, bool)`

GetScopeNetAddrOk returns a tuple with the ScopeNetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetAddr

`func (o *DhcpScopeDataData) SetScopeNetAddr(v string)`

SetScopeNetAddr sets ScopeNetAddr field to given value.

### HasScopeNetAddr

`func (o *DhcpScopeDataData) HasScopeNetAddr() bool`

HasScopeNetAddr returns a boolean if a field has been set.

### GetScopeNetMask

`func (o *DhcpScopeDataData) GetScopeNetMask() string`

GetScopeNetMask returns the ScopeNetMask field if non-nil, zero value otherwise.

### GetScopeNetMaskOk

`func (o *DhcpScopeDataData) GetScopeNetMaskOk() (*string, bool)`

GetScopeNetMaskOk returns a tuple with the ScopeNetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetMask

`func (o *DhcpScopeDataData) SetScopeNetMask(v string)`

SetScopeNetMask sets ScopeNetMask field to given value.

### HasScopeNetMask

`func (o *DhcpScopeDataData) HasScopeNetMask() bool`

HasScopeNetMask returns a boolean if a field has been set.

### GetScopeSpaceId

`func (o *DhcpScopeDataData) GetScopeSpaceId() string`

GetScopeSpaceId returns the ScopeSpaceId field if non-nil, zero value otherwise.

### GetScopeSpaceIdOk

`func (o *DhcpScopeDataData) GetScopeSpaceIdOk() (*string, bool)`

GetScopeSpaceIdOk returns a tuple with the ScopeSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceId

`func (o *DhcpScopeDataData) SetScopeSpaceId(v string)`

SetScopeSpaceId sets ScopeSpaceId field to given value.

### HasScopeSpaceId

`func (o *DhcpScopeDataData) HasScopeSpaceId() bool`

HasScopeSpaceId returns a boolean if a field has been set.

### GetScopeSpaceName

`func (o *DhcpScopeDataData) GetScopeSpaceName() string`

GetScopeSpaceName returns the ScopeSpaceName field if non-nil, zero value otherwise.

### GetScopeSpaceNameOk

`func (o *DhcpScopeDataData) GetScopeSpaceNameOk() (*string, bool)`

GetScopeSpaceNameOk returns a tuple with the ScopeSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceName

`func (o *DhcpScopeDataData) SetScopeSpaceName(v string)`

SetScopeSpaceName sets ScopeSpaceName field to given value.

### HasScopeSpaceName

`func (o *DhcpScopeDataData) HasScopeSpaceName() bool`

HasScopeSpaceName returns a boolean if a field has been set.

### GetScopeSize

`func (o *DhcpScopeDataData) GetScopeSize() string`

GetScopeSize returns the ScopeSize field if non-nil, zero value otherwise.

### GetScopeSizeOk

`func (o *DhcpScopeDataData) GetScopeSizeOk() (*string, bool)`

GetScopeSizeOk returns a tuple with the ScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSize

`func (o *DhcpScopeDataData) SetScopeSize(v string)`

SetScopeSize sets ScopeSize field to given value.

### HasScopeSize

`func (o *DhcpScopeDataData) HasScopeSize() bool`

HasScopeSize returns a boolean if a field has been set.

### GetScopeStartAddressAddr

`func (o *DhcpScopeDataData) GetScopeStartAddressAddr() string`

GetScopeStartAddressAddr returns the ScopeStartAddressAddr field if non-nil, zero value otherwise.

### GetScopeStartAddressAddrOk

`func (o *DhcpScopeDataData) GetScopeStartAddressAddrOk() (*string, bool)`

GetScopeStartAddressAddrOk returns a tuple with the ScopeStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStartAddressAddr

`func (o *DhcpScopeDataData) SetScopeStartAddressAddr(v string)`

SetScopeStartAddressAddr sets ScopeStartAddressAddr field to given value.

### HasScopeStartAddressAddr

`func (o *DhcpScopeDataData) HasScopeStartAddressAddr() bool`

HasScopeStartAddressAddr returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DhcpScopeDataData) GetSharednetworkId() string`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DhcpScopeDataData) GetSharednetworkIdOk() (*string, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DhcpScopeDataData) SetSharednetworkId(v string)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DhcpScopeDataData) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DhcpScopeDataData) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DhcpScopeDataData) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DhcpScopeDataData) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DhcpScopeDataData) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetServerAddr

`func (o *DhcpScopeDataData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DhcpScopeDataData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DhcpScopeDataData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DhcpScopeDataData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetScopeMultistatus

`func (o *DhcpScopeDataData) GetScopeMultistatus() string`

GetScopeMultistatus returns the ScopeMultistatus field if non-nil, zero value otherwise.

### GetScopeMultistatusOk

`func (o *DhcpScopeDataData) GetScopeMultistatusOk() (*string, bool)`

GetScopeMultistatusOk returns a tuple with the ScopeMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeMultistatus

`func (o *DhcpScopeDataData) SetScopeMultistatus(v string)`

SetScopeMultistatus sets ScopeMultistatus field to given value.

### HasScopeMultistatus

`func (o *DhcpScopeDataData) HasScopeMultistatus() bool`

HasScopeMultistatus returns a boolean if a field has been set.

### GetSmartArch

`func (o *DhcpScopeDataData) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DhcpScopeDataData) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DhcpScopeDataData) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DhcpScopeDataData) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpScopeDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpScopeDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpScopeDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpScopeDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DhcpScopeDataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DhcpScopeDataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DhcpScopeDataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DhcpScopeDataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


