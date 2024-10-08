# DataInnerDhcpRangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RangeDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**RangeDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 server the object belongs to. | [optional] 
**ServerComment** | Pointer to **string** | The description of the DHCPv4 server the object belongs to. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft Windows DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv4 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server the object belongs to. | [optional] 
**RangeAcl** | Pointer to **string** | The list of ACLs associated with the DHCPv4 range, as follows: &lt;b&gt;&amp;lt;ACL_name&amp;gt;;&amp;lt;ACL_name&amp;gt;;&lt;/b&gt;... . | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 range, it can be preceded by the class directory. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 range. | [optional] 
**RangeEndAddr** | Pointer to **string** | The last IP address of the DHCPv4 range. | [optional] 
**RangeEndAddressAddr** | Pointer to **string** | The last IP address of the DHCPv4 range, in hexadecimal format. | [optional] 
**RangeFailoverName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 range. | [optional] 
**RangeLeaseCount** | Pointer to **string** | The total number of leases currently delivered by the DHCPv4 range. | [optional] 
**RangeLeasePercent** | Pointer to **string** | The percentage of leases currently delivered by the DHCPv4 range. | [optional] 
**RangeName** | Pointer to **string** | The start and end IP address of the DHCPv4 range, &lt;b&gt;range_start_addr&lt;/b&gt; and &lt;b&gt;range_end_addr&lt;/b&gt;, as follows: &lt;b&gt;&amp;lt;start-ip&amp;gt;-&amp;lt;end-ip&amp;gt;&lt;/b&gt;. | [optional] 
**RangeSize** | Pointer to **string** | The number of IP addresses the DHCPv4 range contains. | [optional] 
**RangeStartAddr** | Pointer to **string** | The first IP address of the DHCPv4 range. | [optional] 
**RangeStartAddressAddr** | Pointer to **string** | The first IP address of the DHCPv4 range, in hexadecimal format. | [optional] 
**RangeState** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ScopeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**ScopeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 scope the object belongs to. | [optional] 
**ScopeId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeIfAddr** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ScopeIfName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ScopeName** | Pointer to **string** | The name of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeNetAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope the object belongs. | [optional] 
**ScopeNetMask** | Pointer to **string** | The netmask of the DHCPv4 scope the object belongs to. It is expressed in dot-decimal notation and defines the number of addresses the scope contains. | [optional] 
**ScopeSpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the DHCPv4 scope the object belongs to. | [optional] 
**ScopeSize** | Pointer to **string** | The number of IP addresses the DHCPv4 scope the object belongs to contains. | [optional] 
**ScopeStartAddressAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope the object belongs to, in hexadecimal format. | [optional] 
**SharednetworkId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 shared network the object belongs to. | [optional] 
**SharednetworkName** | Pointer to **string** | The name of the DHCPv4 shared network the object belongs to. | [optional] 
**ServerHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server_addr&lt;/b&gt; or &lt;b&gt;server_addr6&lt;/b&gt;. | [optional] 
**ServerAddr6** | Pointer to **string** | The Management IP address of the DHCPv4 server the object belongs to, the IPv6 address configured when adding the server, in hexadecimal format. | [optional] 
**ServerAddr** | Pointer to **string** | The Management IP address of the DHCPv4 server the object belongs to, the IPv4 address configured when adding the server, in hexadecimal format. | [optional] 
**RangeMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpRangeData

`func NewDataInnerDhcpRangeData() *DataInnerDhcpRangeData`

NewDataInnerDhcpRangeData instantiates a new DataInnerDhcpRangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpRangeDataWithDefaults

`func NewDataInnerDhcpRangeDataWithDefaults() *DataInnerDhcpRangeData`

NewDataInnerDhcpRangeDataWithDefaults instantiates a new DataInnerDhcpRangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRangeDelayedCreateTime

`func (o *DataInnerDhcpRangeData) GetRangeDelayedCreateTime() string`

GetRangeDelayedCreateTime returns the RangeDelayedCreateTime field if non-nil, zero value otherwise.

### GetRangeDelayedCreateTimeOk

`func (o *DataInnerDhcpRangeData) GetRangeDelayedCreateTimeOk() (*string, bool)`

GetRangeDelayedCreateTimeOk returns a tuple with the RangeDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDelayedCreateTime

`func (o *DataInnerDhcpRangeData) SetRangeDelayedCreateTime(v string)`

SetRangeDelayedCreateTime sets RangeDelayedCreateTime field to given value.

### HasRangeDelayedCreateTime

`func (o *DataInnerDhcpRangeData) HasRangeDelayedCreateTime() bool`

HasRangeDelayedCreateTime returns a boolean if a field has been set.

### GetRangeDelayedDeleteTime

`func (o *DataInnerDhcpRangeData) GetRangeDelayedDeleteTime() string`

GetRangeDelayedDeleteTime returns the RangeDelayedDeleteTime field if non-nil, zero value otherwise.

### GetRangeDelayedDeleteTimeOk

`func (o *DataInnerDhcpRangeData) GetRangeDelayedDeleteTimeOk() (*string, bool)`

GetRangeDelayedDeleteTimeOk returns a tuple with the RangeDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDelayedDeleteTime

`func (o *DataInnerDhcpRangeData) SetRangeDelayedDeleteTime(v string)`

SetRangeDelayedDeleteTime sets RangeDelayedDeleteTime field to given value.

### HasRangeDelayedDeleteTime

`func (o *DataInnerDhcpRangeData) HasRangeDelayedDeleteTime() bool`

HasRangeDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DataInnerDhcpRangeData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DataInnerDhcpRangeData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DataInnerDhcpRangeData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DataInnerDhcpRangeData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DataInnerDhcpRangeData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DataInnerDhcpRangeData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DataInnerDhcpRangeData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DataInnerDhcpRangeData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerComment

`func (o *DataInnerDhcpRangeData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DataInnerDhcpRangeData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DataInnerDhcpRangeData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DataInnerDhcpRangeData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDhcpRangeData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDhcpRangeData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDhcpRangeData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDhcpRangeData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDhcpRangeData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDhcpRangeData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDhcpRangeData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDhcpRangeData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDhcpRangeData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDhcpRangeData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDhcpRangeData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDhcpRangeData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DataInnerDhcpRangeData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DataInnerDhcpRangeData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DataInnerDhcpRangeData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DataInnerDhcpRangeData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetRangeAcl

`func (o *DataInnerDhcpRangeData) GetRangeAcl() string`

GetRangeAcl returns the RangeAcl field if non-nil, zero value otherwise.

### GetRangeAclOk

`func (o *DataInnerDhcpRangeData) GetRangeAclOk() (*string, bool)`

GetRangeAclOk returns a tuple with the RangeAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeAcl

`func (o *DataInnerDhcpRangeData) SetRangeAcl(v string)`

SetRangeAcl sets RangeAcl field to given value.

### HasRangeAcl

`func (o *DataInnerDhcpRangeData) HasRangeAcl() bool`

HasRangeAcl returns a boolean if a field has been set.

### GetRangeClassName

`func (o *DataInnerDhcpRangeData) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *DataInnerDhcpRangeData) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *DataInnerDhcpRangeData) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *DataInnerDhcpRangeData) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *DataInnerDhcpRangeData) GetRangeClassParameters() []ApiClassParameterOutputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *DataInnerDhcpRangeData) GetRangeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *DataInnerDhcpRangeData) SetRangeClassParameters(v []ApiClassParameterOutputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *DataInnerDhcpRangeData) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetRangeEndAddr

`func (o *DataInnerDhcpRangeData) GetRangeEndAddr() string`

GetRangeEndAddr returns the RangeEndAddr field if non-nil, zero value otherwise.

### GetRangeEndAddrOk

`func (o *DataInnerDhcpRangeData) GetRangeEndAddrOk() (*string, bool)`

GetRangeEndAddrOk returns a tuple with the RangeEndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndAddr

`func (o *DataInnerDhcpRangeData) SetRangeEndAddr(v string)`

SetRangeEndAddr sets RangeEndAddr field to given value.

### HasRangeEndAddr

`func (o *DataInnerDhcpRangeData) HasRangeEndAddr() bool`

HasRangeEndAddr returns a boolean if a field has been set.

### GetRangeEndAddressAddr

`func (o *DataInnerDhcpRangeData) GetRangeEndAddressAddr() string`

GetRangeEndAddressAddr returns the RangeEndAddressAddr field if non-nil, zero value otherwise.

### GetRangeEndAddressAddrOk

`func (o *DataInnerDhcpRangeData) GetRangeEndAddressAddrOk() (*string, bool)`

GetRangeEndAddressAddrOk returns a tuple with the RangeEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndAddressAddr

`func (o *DataInnerDhcpRangeData) SetRangeEndAddressAddr(v string)`

SetRangeEndAddressAddr sets RangeEndAddressAddr field to given value.

### HasRangeEndAddressAddr

`func (o *DataInnerDhcpRangeData) HasRangeEndAddressAddr() bool`

HasRangeEndAddressAddr returns a boolean if a field has been set.

### GetRangeFailoverName

`func (o *DataInnerDhcpRangeData) GetRangeFailoverName() string`

GetRangeFailoverName returns the RangeFailoverName field if non-nil, zero value otherwise.

### GetRangeFailoverNameOk

`func (o *DataInnerDhcpRangeData) GetRangeFailoverNameOk() (*string, bool)`

GetRangeFailoverNameOk returns a tuple with the RangeFailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeFailoverName

`func (o *DataInnerDhcpRangeData) SetRangeFailoverName(v string)`

SetRangeFailoverName sets RangeFailoverName field to given value.

### HasRangeFailoverName

`func (o *DataInnerDhcpRangeData) HasRangeFailoverName() bool`

HasRangeFailoverName returns a boolean if a field has been set.

### GetRangeId

`func (o *DataInnerDhcpRangeData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *DataInnerDhcpRangeData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *DataInnerDhcpRangeData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *DataInnerDhcpRangeData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeLeaseCount

`func (o *DataInnerDhcpRangeData) GetRangeLeaseCount() string`

GetRangeLeaseCount returns the RangeLeaseCount field if non-nil, zero value otherwise.

### GetRangeLeaseCountOk

`func (o *DataInnerDhcpRangeData) GetRangeLeaseCountOk() (*string, bool)`

GetRangeLeaseCountOk returns a tuple with the RangeLeaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeLeaseCount

`func (o *DataInnerDhcpRangeData) SetRangeLeaseCount(v string)`

SetRangeLeaseCount sets RangeLeaseCount field to given value.

### HasRangeLeaseCount

`func (o *DataInnerDhcpRangeData) HasRangeLeaseCount() bool`

HasRangeLeaseCount returns a boolean if a field has been set.

### GetRangeLeasePercent

`func (o *DataInnerDhcpRangeData) GetRangeLeasePercent() string`

GetRangeLeasePercent returns the RangeLeasePercent field if non-nil, zero value otherwise.

### GetRangeLeasePercentOk

`func (o *DataInnerDhcpRangeData) GetRangeLeasePercentOk() (*string, bool)`

GetRangeLeasePercentOk returns a tuple with the RangeLeasePercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeLeasePercent

`func (o *DataInnerDhcpRangeData) SetRangeLeasePercent(v string)`

SetRangeLeasePercent sets RangeLeasePercent field to given value.

### HasRangeLeasePercent

`func (o *DataInnerDhcpRangeData) HasRangeLeasePercent() bool`

HasRangeLeasePercent returns a boolean if a field has been set.

### GetRangeName

`func (o *DataInnerDhcpRangeData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *DataInnerDhcpRangeData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *DataInnerDhcpRangeData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *DataInnerDhcpRangeData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetRangeSize

`func (o *DataInnerDhcpRangeData) GetRangeSize() string`

GetRangeSize returns the RangeSize field if non-nil, zero value otherwise.

### GetRangeSizeOk

`func (o *DataInnerDhcpRangeData) GetRangeSizeOk() (*string, bool)`

GetRangeSizeOk returns a tuple with the RangeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeSize

`func (o *DataInnerDhcpRangeData) SetRangeSize(v string)`

SetRangeSize sets RangeSize field to given value.

### HasRangeSize

`func (o *DataInnerDhcpRangeData) HasRangeSize() bool`

HasRangeSize returns a boolean if a field has been set.

### GetRangeStartAddr

`func (o *DataInnerDhcpRangeData) GetRangeStartAddr() string`

GetRangeStartAddr returns the RangeStartAddr field if non-nil, zero value otherwise.

### GetRangeStartAddrOk

`func (o *DataInnerDhcpRangeData) GetRangeStartAddrOk() (*string, bool)`

GetRangeStartAddrOk returns a tuple with the RangeStartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartAddr

`func (o *DataInnerDhcpRangeData) SetRangeStartAddr(v string)`

SetRangeStartAddr sets RangeStartAddr field to given value.

### HasRangeStartAddr

`func (o *DataInnerDhcpRangeData) HasRangeStartAddr() bool`

HasRangeStartAddr returns a boolean if a field has been set.

### GetRangeStartAddressAddr

`func (o *DataInnerDhcpRangeData) GetRangeStartAddressAddr() string`

GetRangeStartAddressAddr returns the RangeStartAddressAddr field if non-nil, zero value otherwise.

### GetRangeStartAddressAddrOk

`func (o *DataInnerDhcpRangeData) GetRangeStartAddressAddrOk() (*string, bool)`

GetRangeStartAddressAddrOk returns a tuple with the RangeStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartAddressAddr

`func (o *DataInnerDhcpRangeData) SetRangeStartAddressAddr(v string)`

SetRangeStartAddressAddr sets RangeStartAddressAddr field to given value.

### HasRangeStartAddressAddr

`func (o *DataInnerDhcpRangeData) HasRangeStartAddressAddr() bool`

HasRangeStartAddressAddr returns a boolean if a field has been set.

### GetRangeState

`func (o *DataInnerDhcpRangeData) GetRangeState() string`

GetRangeState returns the RangeState field if non-nil, zero value otherwise.

### GetRangeStateOk

`func (o *DataInnerDhcpRangeData) GetRangeStateOk() (*string, bool)`

GetRangeStateOk returns a tuple with the RangeState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeState

`func (o *DataInnerDhcpRangeData) SetRangeState(v string)`

SetRangeState sets RangeState field to given value.

### HasRangeState

`func (o *DataInnerDhcpRangeData) HasRangeState() bool`

HasRangeState returns a boolean if a field has been set.

### GetScopeClassName

`func (o *DataInnerDhcpRangeData) GetScopeClassName() string`

GetScopeClassName returns the ScopeClassName field if non-nil, zero value otherwise.

### GetScopeClassNameOk

`func (o *DataInnerDhcpRangeData) GetScopeClassNameOk() (*string, bool)`

GetScopeClassNameOk returns a tuple with the ScopeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassName

`func (o *DataInnerDhcpRangeData) SetScopeClassName(v string)`

SetScopeClassName sets ScopeClassName field to given value.

### HasScopeClassName

`func (o *DataInnerDhcpRangeData) HasScopeClassName() bool`

HasScopeClassName returns a boolean if a field has been set.

### GetScopeClassParameters

`func (o *DataInnerDhcpRangeData) GetScopeClassParameters() []ApiClassParameterOutputEntry`

GetScopeClassParameters returns the ScopeClassParameters field if non-nil, zero value otherwise.

### GetScopeClassParametersOk

`func (o *DataInnerDhcpRangeData) GetScopeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScopeClassParametersOk returns a tuple with the ScopeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassParameters

`func (o *DataInnerDhcpRangeData) SetScopeClassParameters(v []ApiClassParameterOutputEntry)`

SetScopeClassParameters sets ScopeClassParameters field to given value.

### HasScopeClassParameters

`func (o *DataInnerDhcpRangeData) HasScopeClassParameters() bool`

HasScopeClassParameters returns a boolean if a field has been set.

### GetScopeId

`func (o *DataInnerDhcpRangeData) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DataInnerDhcpRangeData) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DataInnerDhcpRangeData) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *DataInnerDhcpRangeData) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetScopeIfAddr

`func (o *DataInnerDhcpRangeData) GetScopeIfAddr() string`

GetScopeIfAddr returns the ScopeIfAddr field if non-nil, zero value otherwise.

### GetScopeIfAddrOk

`func (o *DataInnerDhcpRangeData) GetScopeIfAddrOk() (*string, bool)`

GetScopeIfAddrOk returns a tuple with the ScopeIfAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeIfAddr

`func (o *DataInnerDhcpRangeData) SetScopeIfAddr(v string)`

SetScopeIfAddr sets ScopeIfAddr field to given value.

### HasScopeIfAddr

`func (o *DataInnerDhcpRangeData) HasScopeIfAddr() bool`

HasScopeIfAddr returns a boolean if a field has been set.

### GetScopeIfName

`func (o *DataInnerDhcpRangeData) GetScopeIfName() string`

GetScopeIfName returns the ScopeIfName field if non-nil, zero value otherwise.

### GetScopeIfNameOk

`func (o *DataInnerDhcpRangeData) GetScopeIfNameOk() (*string, bool)`

GetScopeIfNameOk returns a tuple with the ScopeIfName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeIfName

`func (o *DataInnerDhcpRangeData) SetScopeIfName(v string)`

SetScopeIfName sets ScopeIfName field to given value.

### HasScopeIfName

`func (o *DataInnerDhcpRangeData) HasScopeIfName() bool`

HasScopeIfName returns a boolean if a field has been set.

### GetScopeName

`func (o *DataInnerDhcpRangeData) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DataInnerDhcpRangeData) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DataInnerDhcpRangeData) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DataInnerDhcpRangeData) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeNetAddr

`func (o *DataInnerDhcpRangeData) GetScopeNetAddr() string`

GetScopeNetAddr returns the ScopeNetAddr field if non-nil, zero value otherwise.

### GetScopeNetAddrOk

`func (o *DataInnerDhcpRangeData) GetScopeNetAddrOk() (*string, bool)`

GetScopeNetAddrOk returns a tuple with the ScopeNetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetAddr

`func (o *DataInnerDhcpRangeData) SetScopeNetAddr(v string)`

SetScopeNetAddr sets ScopeNetAddr field to given value.

### HasScopeNetAddr

`func (o *DataInnerDhcpRangeData) HasScopeNetAddr() bool`

HasScopeNetAddr returns a boolean if a field has been set.

### GetScopeNetMask

`func (o *DataInnerDhcpRangeData) GetScopeNetMask() string`

GetScopeNetMask returns the ScopeNetMask field if non-nil, zero value otherwise.

### GetScopeNetMaskOk

`func (o *DataInnerDhcpRangeData) GetScopeNetMaskOk() (*string, bool)`

GetScopeNetMaskOk returns a tuple with the ScopeNetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetMask

`func (o *DataInnerDhcpRangeData) SetScopeNetMask(v string)`

SetScopeNetMask sets ScopeNetMask field to given value.

### HasScopeNetMask

`func (o *DataInnerDhcpRangeData) HasScopeNetMask() bool`

HasScopeNetMask returns a boolean if a field has been set.

### GetScopeSpaceId

`func (o *DataInnerDhcpRangeData) GetScopeSpaceId() string`

GetScopeSpaceId returns the ScopeSpaceId field if non-nil, zero value otherwise.

### GetScopeSpaceIdOk

`func (o *DataInnerDhcpRangeData) GetScopeSpaceIdOk() (*string, bool)`

GetScopeSpaceIdOk returns a tuple with the ScopeSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceId

`func (o *DataInnerDhcpRangeData) SetScopeSpaceId(v string)`

SetScopeSpaceId sets ScopeSpaceId field to given value.

### HasScopeSpaceId

`func (o *DataInnerDhcpRangeData) HasScopeSpaceId() bool`

HasScopeSpaceId returns a boolean if a field has been set.

### GetScopeSize

`func (o *DataInnerDhcpRangeData) GetScopeSize() string`

GetScopeSize returns the ScopeSize field if non-nil, zero value otherwise.

### GetScopeSizeOk

`func (o *DataInnerDhcpRangeData) GetScopeSizeOk() (*string, bool)`

GetScopeSizeOk returns a tuple with the ScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSize

`func (o *DataInnerDhcpRangeData) SetScopeSize(v string)`

SetScopeSize sets ScopeSize field to given value.

### HasScopeSize

`func (o *DataInnerDhcpRangeData) HasScopeSize() bool`

HasScopeSize returns a boolean if a field has been set.

### GetScopeStartAddressAddr

`func (o *DataInnerDhcpRangeData) GetScopeStartAddressAddr() string`

GetScopeStartAddressAddr returns the ScopeStartAddressAddr field if non-nil, zero value otherwise.

### GetScopeStartAddressAddrOk

`func (o *DataInnerDhcpRangeData) GetScopeStartAddressAddrOk() (*string, bool)`

GetScopeStartAddressAddrOk returns a tuple with the ScopeStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeStartAddressAddr

`func (o *DataInnerDhcpRangeData) SetScopeStartAddressAddr(v string)`

SetScopeStartAddressAddr sets ScopeStartAddressAddr field to given value.

### HasScopeStartAddressAddr

`func (o *DataInnerDhcpRangeData) HasScopeStartAddressAddr() bool`

HasScopeStartAddressAddr returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DataInnerDhcpRangeData) GetSharednetworkId() string`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DataInnerDhcpRangeData) GetSharednetworkIdOk() (*string, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DataInnerDhcpRangeData) SetSharednetworkId(v string)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DataInnerDhcpRangeData) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DataInnerDhcpRangeData) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DataInnerDhcpRangeData) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DataInnerDhcpRangeData) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DataInnerDhcpRangeData) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DataInnerDhcpRangeData) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DataInnerDhcpRangeData) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DataInnerDhcpRangeData) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DataInnerDhcpRangeData) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerAddr6

`func (o *DataInnerDhcpRangeData) GetServerAddr6() string`

GetServerAddr6 returns the ServerAddr6 field if non-nil, zero value otherwise.

### GetServerAddr6Ok

`func (o *DataInnerDhcpRangeData) GetServerAddr6Ok() (*string, bool)`

GetServerAddr6Ok returns a tuple with the ServerAddr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr6

`func (o *DataInnerDhcpRangeData) SetServerAddr6(v string)`

SetServerAddr6 sets ServerAddr6 field to given value.

### HasServerAddr6

`func (o *DataInnerDhcpRangeData) HasServerAddr6() bool`

HasServerAddr6 returns a boolean if a field has been set.

### GetServerAddr

`func (o *DataInnerDhcpRangeData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DataInnerDhcpRangeData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DataInnerDhcpRangeData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DataInnerDhcpRangeData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetRangeMultistatus

`func (o *DataInnerDhcpRangeData) GetRangeMultistatus() string`

GetRangeMultistatus returns the RangeMultistatus field if non-nil, zero value otherwise.

### GetRangeMultistatusOk

`func (o *DataInnerDhcpRangeData) GetRangeMultistatusOk() (*string, bool)`

GetRangeMultistatusOk returns a tuple with the RangeMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeMultistatus

`func (o *DataInnerDhcpRangeData) SetRangeMultistatus(v string)`

SetRangeMultistatus sets RangeMultistatus field to given value.

### HasRangeMultistatus

`func (o *DataInnerDhcpRangeData) HasRangeMultistatus() bool`

HasRangeMultistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpRangeData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpRangeData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpRangeData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpRangeData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpRangeData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpRangeData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpRangeData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpRangeData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


