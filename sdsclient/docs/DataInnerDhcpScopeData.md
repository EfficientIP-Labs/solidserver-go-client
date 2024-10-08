# DataInnerDhcpScopeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScopeDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**ScopeDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 server the object belongs to. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft Windows DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv4 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server the object belongs to. | [optional] 
**FailoverId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 failover channel associated with the object. | [optional] 
**FailoverName** | Pointer to **string** | The name of the DHCPv4 failover channel associated with the object. | [optional] 
**ScopeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 scope, it can be preceded by the class directory. | [optional] 
**ScopeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 scope. | [optional] 
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
**ServerHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server_addr&lt;/b&gt; or &lt;b&gt;server_addr6&lt;/b&gt;. | [optional] 
**ServerAddr6** | Pointer to **string** | The Management IP address of the DHCPv4 server the object belongs to, the IPv6 address configured when adding the server, in hexadecimal format. | [optional] 
**ServerAddr** | Pointer to **string** | The Management IP address of the DHCPv4 server the object belongs to, the IPv4 address configured when adding the server, in hexadecimal format. | [optional] 
**ScopeMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartArch** | Pointer to **string** | The type of the DHCPv4 smart architecture the object belongs to.&lt;table&gt;&lt;caption&gt;smart_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;masterslave&lt;/td&gt;&lt;td &gt;The One-to-One smart architecture sets a pair of DHCP servers in a Master/Backup configuration.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;star&lt;/td&gt;&lt;td &gt;The One-to-Many smart architecture sets a multi-site failover configuration at the cost of n-servers+1.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;splitscope&lt;/td&gt;&lt;td &gt;The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;The Single-Server smart architecture manages a single DHCP server.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpScopeData

`func NewDataInnerDhcpScopeData() *DataInnerDhcpScopeData`

NewDataInnerDhcpScopeData instantiates a new DataInnerDhcpScopeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpScopeDataWithDefaults

`func NewDataInnerDhcpScopeDataWithDefaults() *DataInnerDhcpScopeData`

NewDataInnerDhcpScopeDataWithDefaults instantiates a new DataInnerDhcpScopeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScopeDelayedCreateTime

`func (o *DataInnerDhcpScopeData) GetScopeDelayedCreateTime() string`

GetScopeDelayedCreateTime returns the ScopeDelayedCreateTime field if non-nil, zero value otherwise.

### GetScopeDelayedCreateTimeOk

`func (o *DataInnerDhcpScopeData) GetScopeDelayedCreateTimeOk() (*string, bool)`

GetScopeDelayedCreateTimeOk returns a tuple with the ScopeDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDelayedCreateTime

`func (o *DataInnerDhcpScopeData) SetScopeDelayedCreateTime(v string)`

SetScopeDelayedCreateTime sets ScopeDelayedCreateTime field to given value.

### HasScopeDelayedCreateTime

`func (o *DataInnerDhcpScopeData) HasScopeDelayedCreateTime() bool`

HasScopeDelayedCreateTime returns a boolean if a field has been set.

### GetScopeDelayedDeleteTime

`func (o *DataInnerDhcpScopeData) GetScopeDelayedDeleteTime() string`

GetScopeDelayedDeleteTime returns the ScopeDelayedDeleteTime field if non-nil, zero value otherwise.

### GetScopeDelayedDeleteTimeOk

`func (o *DataInnerDhcpScopeData) GetScopeDelayedDeleteTimeOk() (*string, bool)`

GetScopeDelayedDeleteTimeOk returns a tuple with the ScopeDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeDelayedDeleteTime

`func (o *DataInnerDhcpScopeData) SetScopeDelayedDeleteTime(v string)`

SetScopeDelayedDeleteTime sets ScopeDelayedDeleteTime field to given value.

### HasScopeDelayedDeleteTime

`func (o *DataInnerDhcpScopeData) HasScopeDelayedDeleteTime() bool`

HasScopeDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DataInnerDhcpScopeData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DataInnerDhcpScopeData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DataInnerDhcpScopeData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DataInnerDhcpScopeData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DataInnerDhcpScopeData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DataInnerDhcpScopeData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DataInnerDhcpScopeData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DataInnerDhcpScopeData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDhcpScopeData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDhcpScopeData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDhcpScopeData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDhcpScopeData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDhcpScopeData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDhcpScopeData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDhcpScopeData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDhcpScopeData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDhcpScopeData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDhcpScopeData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDhcpScopeData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDhcpScopeData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DataInnerDhcpScopeData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DataInnerDhcpScopeData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DataInnerDhcpScopeData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DataInnerDhcpScopeData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetFailoverId

`func (o *DataInnerDhcpScopeData) GetFailoverId() string`

GetFailoverId returns the FailoverId field if non-nil, zero value otherwise.

### GetFailoverIdOk

`func (o *DataInnerDhcpScopeData) GetFailoverIdOk() (*string, bool)`

GetFailoverIdOk returns a tuple with the FailoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverId

`func (o *DataInnerDhcpScopeData) SetFailoverId(v string)`

SetFailoverId sets FailoverId field to given value.

### HasFailoverId

`func (o *DataInnerDhcpScopeData) HasFailoverId() bool`

HasFailoverId returns a boolean if a field has been set.

### GetFailoverName

`func (o *DataInnerDhcpScopeData) GetFailoverName() string`

GetFailoverName returns the FailoverName field if non-nil, zero value otherwise.

### GetFailoverNameOk

`func (o *DataInnerDhcpScopeData) GetFailoverNameOk() (*string, bool)`

GetFailoverNameOk returns a tuple with the FailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverName

`func (o *DataInnerDhcpScopeData) SetFailoverName(v string)`

SetFailoverName sets FailoverName field to given value.

### HasFailoverName

`func (o *DataInnerDhcpScopeData) HasFailoverName() bool`

HasFailoverName returns a boolean if a field has been set.

### GetScopeClassName

`func (o *DataInnerDhcpScopeData) GetScopeClassName() string`

GetScopeClassName returns the ScopeClassName field if non-nil, zero value otherwise.

### GetScopeClassNameOk

`func (o *DataInnerDhcpScopeData) GetScopeClassNameOk() (*string, bool)`

GetScopeClassNameOk returns a tuple with the ScopeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassName

`func (o *DataInnerDhcpScopeData) SetScopeClassName(v string)`

SetScopeClassName sets ScopeClassName field to given value.

### HasScopeClassName

`func (o *DataInnerDhcpScopeData) HasScopeClassName() bool`

HasScopeClassName returns a boolean if a field has been set.

### GetScopeClassParameters

`func (o *DataInnerDhcpScopeData) GetScopeClassParameters() []ApiClassParameterOutputEntry`

GetScopeClassParameters returns the ScopeClassParameters field if non-nil, zero value otherwise.

### GetScopeClassParametersOk

`func (o *DataInnerDhcpScopeData) GetScopeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScopeClassParametersOk returns a tuple with the ScopeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassParameters

`func (o *DataInnerDhcpScopeData) SetScopeClassParameters(v []ApiClassParameterOutputEntry)`

SetScopeClassParameters sets ScopeClassParameters field to given value.

### HasScopeClassParameters

`func (o *DataInnerDhcpScopeData) HasScopeClassParameters() bool`

HasScopeClassParameters returns a boolean if a field has been set.

### GetScopeEndAddressAddr

`func (o *DataInnerDhcpScopeData) GetScopeEndAddressAddr() string`

GetScopeEndAddressAddr returns the ScopeEndAddressAddr field if non-nil, zero value otherwise.

### GetScopeEndAddressAddrOk

`func (o *DataInnerDhcpScopeData) GetScopeEndAddressAddrOk() (*string, bool)`

GetScopeEndAddressAddrOk returns a tuple with the ScopeEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeEndAddressAddr

`func (o *DataInnerDhcpScopeData) SetScopeEndAddressAddr(v string)`

SetScopeEndAddressAddr sets ScopeEndAddressAddr field to given value.

### HasScopeEndAddressAddr

`func (o *DataInnerDhcpScopeData) HasScopeEndAddressAddr() bool`

HasScopeEndAddressAddr returns a boolean if a field has been set.

### GetScopeId

`func (o *DataInnerDhcpScopeData) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DataInnerDhcpScopeData) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DataInnerDhcpScopeData) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *DataInnerDhcpScopeData) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetScopeName

`func (o *DataInnerDhcpScopeData) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DataInnerDhcpScopeData) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DataInnerDhcpScopeData) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DataInnerDhcpScopeData) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeNetAddr

`func (o *DataInnerDhcpScopeData) GetScopeNetAddr() string`

GetScopeNetAddr returns the ScopeNetAddr field if non-nil, zero value otherwise.

### GetScopeNetAddrOk

`func (o *DataInnerDhcpScopeData) GetScopeNetAddrOk() (*string, bool)`

GetScopeNetAddrOk returns a tuple with the ScopeNetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetAddr

`func (o *DataInnerDhcpScopeData) SetScopeNetAddr(v string)`

SetScopeNetAddr sets ScopeNetAddr field to given value.

### HasScopeNetAddr

`func (o *DataInnerDhcpScopeData) HasScopeNetAddr() bool`

HasScopeNetAddr returns a boolean if a field has been set.

### GetScopeNetMask

`func (o *DataInnerDhcpScopeData) GetScopeNetMask() string`

GetScopeNetMask returns the ScopeNetMask field if non-nil, zero value otherwise.

### GetScopeNetMaskOk

`func (o *DataInnerDhcpScopeData) GetScopeNetMaskOk() (*string, bool)`

GetScopeNetMaskOk returns a tuple with the ScopeNetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetMask

`func (o *DataInnerDhcpScopeData) SetScopeNetMask(v string)`

SetScopeNetMask sets ScopeNetMask field to given value.

### HasScopeNetMask

`func (o *DataInnerDhcpScopeData) HasScopeNetMask() bool`

HasScopeNetMask returns a boolean if a field has been set.

### GetScopeSpaceId

`func (o *DataInnerDhcpScopeData) GetScopeSpaceId() string`

GetScopeSpaceId returns the ScopeSpaceId field if non-nil, zero value otherwise.

### GetScopeSpaceIdOk

`func (o *DataInnerDhcpScopeData) GetScopeSpaceIdOk() (*string, bool)`

GetScopeSpaceIdOk returns a tuple with the ScopeSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceId

`func (o *DataInnerDhcpScopeData) SetScopeSpaceId(v string)`

SetScopeSpaceId sets ScopeSpaceId field to given value.

### HasScopeSpaceId

`func (o *DataInnerDhcpScopeData) HasScopeSpaceId() bool`

HasScopeSpaceId returns a boolean if a field has been set.

### GetScopeSpaceName

`func (o *DataInnerDhcpScopeData) GetScopeSpaceName() string`

GetScopeSpaceName returns the ScopeSpaceName field if non-nil, zero value otherwise.

### GetScopeSpaceNameOk

`func (o *DataInnerDhcpScopeData) GetScopeSpaceNameOk() (*string, bool)`

GetScopeSpaceNameOk returns a tuple with the ScopeSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceName

`func (o *DataInnerDhcpScopeData) SetScopeSpaceName(v string)`

SetScopeSpaceName sets ScopeSpaceName field to given value.

### HasScopeSpaceName

`func (o *DataInnerDhcpScopeData) HasScopeSpaceName() bool`

HasScopeSpaceName returns a boolean if a field has been set.

### GetScopeSize

`func (o *DataInnerDhcpScopeData) GetScopeSize() string`

GetScopeSize returns the ScopeSize field if non-nil, zero value otherwise.

### GetScopeSizeOk

`func (o *DataInnerDhcpScopeData) GetScopeSizeOk() (*string, bool)`

GetScopeSizeOk returns a tuple with the ScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSize

`func (o *DataInnerDhcpScopeData) SetScopeSize(v string)`

SetScopeSize sets ScopeSize field to given value.

### HasScopeSize

`func (o *DataInnerDhcpScopeData) HasScopeSize() bool`

HasScopeSize returns a boolean if a field has been set.

### GetScopeStartAddressAddr

`func (o *DataInnerDhcpScopeData) GetScopeStartAddressAddr() string`

GetScopeStartAddressAddr returns the ScopeStartAddressAddr field if non-nil, zero value otherwise.

### GetScopeStartAddressAddrOk

`func (o *DataInnerDhcpScopeData) GetScopeStartAddressAddrOk() (*string, bool)`

GetScopeStartAddressAddrOk returns a tuple with the ScopeStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStartAddressAddr

`func (o *DataInnerDhcpScopeData) SetScopeStartAddressAddr(v string)`

SetScopeStartAddressAddr sets ScopeStartAddressAddr field to given value.

### HasScopeStartAddressAddr

`func (o *DataInnerDhcpScopeData) HasScopeStartAddressAddr() bool`

HasScopeStartAddressAddr returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DataInnerDhcpScopeData) GetSharednetworkId() string`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DataInnerDhcpScopeData) GetSharednetworkIdOk() (*string, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DataInnerDhcpScopeData) SetSharednetworkId(v string)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DataInnerDhcpScopeData) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DataInnerDhcpScopeData) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DataInnerDhcpScopeData) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DataInnerDhcpScopeData) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DataInnerDhcpScopeData) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DataInnerDhcpScopeData) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DataInnerDhcpScopeData) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DataInnerDhcpScopeData) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DataInnerDhcpScopeData) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerAddr6

`func (o *DataInnerDhcpScopeData) GetServerAddr6() string`

GetServerAddr6 returns the ServerAddr6 field if non-nil, zero value otherwise.

### GetServerAddr6Ok

`func (o *DataInnerDhcpScopeData) GetServerAddr6Ok() (*string, bool)`

GetServerAddr6Ok returns a tuple with the ServerAddr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr6

`func (o *DataInnerDhcpScopeData) SetServerAddr6(v string)`

SetServerAddr6 sets ServerAddr6 field to given value.

### HasServerAddr6

`func (o *DataInnerDhcpScopeData) HasServerAddr6() bool`

HasServerAddr6 returns a boolean if a field has been set.

### GetServerAddr

`func (o *DataInnerDhcpScopeData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DataInnerDhcpScopeData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DataInnerDhcpScopeData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DataInnerDhcpScopeData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetScopeMultistatus

`func (o *DataInnerDhcpScopeData) GetScopeMultistatus() string`

GetScopeMultistatus returns the ScopeMultistatus field if non-nil, zero value otherwise.

### GetScopeMultistatusOk

`func (o *DataInnerDhcpScopeData) GetScopeMultistatusOk() (*string, bool)`

GetScopeMultistatusOk returns a tuple with the ScopeMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeMultistatus

`func (o *DataInnerDhcpScopeData) SetScopeMultistatus(v string)`

SetScopeMultistatus sets ScopeMultistatus field to given value.

### HasScopeMultistatus

`func (o *DataInnerDhcpScopeData) HasScopeMultistatus() bool`

HasScopeMultistatus returns a boolean if a field has been set.

### GetSmartArch

`func (o *DataInnerDhcpScopeData) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DataInnerDhcpScopeData) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DataInnerDhcpScopeData) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DataInnerDhcpScopeData) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpScopeData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpScopeData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpScopeData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpScopeData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpScopeData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpScopeData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpScopeData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpScopeData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


