# DataInnerDhcpScope6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope6DelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**Server6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 server the object belongs to, it can be preceded by the class directory. | [optional] 
**Server6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 server the object belongs to. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;server6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv6 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Version** | Pointer to **string** | The version details of the DHCPv6 server the object belongs to. | [optional] 
**Failover6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 failover channel associated with the object. | [optional] 
**Failover6Name** | Pointer to **string** | The name of the DHCPv6 failover channel associated with the object. | [optional] 
**Scope6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 scope, it can be preceded by the class directory. | [optional] 
**Scope6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 scope. | [optional] 
**Scope6EndAddress6Addr** | Pointer to **string** | The last IP address of the DHCPv6 scope, in hexadecimal format. | [optional] 
**Scope6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 scope. | [optional] 
**Scope6Name** | Pointer to **string** | The name of the DHCPv6 scope. | [optional] 
**Scope6Prefix** | Pointer to **string** | The prefix of the DHCPv6 scope. | [optional] 
**Scope6SpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the DHCPv6 scope. | [optional] 
**Scope6SpaceName** | Pointer to **string** | The name of the space associated with the DHCPv6 scope. | [optional] 
**Scope6Size** | Pointer to **string** | The number of IP addresses the DHCPv6 scope contains. | [optional] 
**Scope6SortName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Scope6StartAddress6Addr** | Pointer to **string** | The first IP address of the DHCPv6 scope, in hexadecimal format. | [optional] 
**Sharednetwork6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 shared network the object belongs to. | [optional] 
**Sharednetwork6Name** | Pointer to **string** | The name of the DHCPv6 shared network the object belongs to. | [optional] 
**Server6Hostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server6_addr&lt;/b&gt; or &lt;b&gt;server6_addr6&lt;/b&gt;. | [optional] 
**Server6Addr6** | Pointer to **string** | The Management IP address of the DHCPv6 server the object belongs to, the IPv6 address configured when adding the server, in hexadecimal format. | [optional] 
**Server6Addr** | Pointer to **string** | The Management IP address of the DHCPv6 server the object belongs to, the IPv4 address configured when adding the server, in hexadecimal format. | [optional] 
**Scope6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartArch** | Pointer to **string** | The type of the DHCPv6 smart architecture the object belongs to.&lt;table&gt;&lt;caption&gt;smart_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;The Single-Server smart architecture manages a single DHCPv6 server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;splitscope&lt;/td&gt;&lt;td &gt;The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;stateless&lt;/td&gt;&lt;td &gt;The Stateless smart architecture offers a limited number of options to the DHCP clients. The IP address is delivered thanks to the subnet gateway and it is impossible to create any ranges or statics or to retrieve any leases.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpScope6Data

`func NewDataInnerDhcpScope6Data() *DataInnerDhcpScope6Data`

NewDataInnerDhcpScope6Data instantiates a new DataInnerDhcpScope6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpScope6DataWithDefaults

`func NewDataInnerDhcpScope6DataWithDefaults() *DataInnerDhcpScope6Data`

NewDataInnerDhcpScope6DataWithDefaults instantiates a new DataInnerDhcpScope6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope6DelayedTime

`func (o *DataInnerDhcpScope6Data) GetScope6DelayedTime() string`

GetScope6DelayedTime returns the Scope6DelayedTime field if non-nil, zero value otherwise.

### GetScope6DelayedTimeOk

`func (o *DataInnerDhcpScope6Data) GetScope6DelayedTimeOk() (*string, bool)`

GetScope6DelayedTimeOk returns a tuple with the Scope6DelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6DelayedTime

`func (o *DataInnerDhcpScope6Data) SetScope6DelayedTime(v string)`

SetScope6DelayedTime sets Scope6DelayedTime field to given value.

### HasScope6DelayedTime

`func (o *DataInnerDhcpScope6Data) HasScope6DelayedTime() bool`

HasScope6DelayedTime returns a boolean if a field has been set.

### GetServer6ClassName

`func (o *DataInnerDhcpScope6Data) GetServer6ClassName() string`

GetServer6ClassName returns the Server6ClassName field if non-nil, zero value otherwise.

### GetServer6ClassNameOk

`func (o *DataInnerDhcpScope6Data) GetServer6ClassNameOk() (*string, bool)`

GetServer6ClassNameOk returns a tuple with the Server6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassName

`func (o *DataInnerDhcpScope6Data) SetServer6ClassName(v string)`

SetServer6ClassName sets Server6ClassName field to given value.

### HasServer6ClassName

`func (o *DataInnerDhcpScope6Data) HasServer6ClassName() bool`

HasServer6ClassName returns a boolean if a field has been set.

### GetServer6ClassParameters

`func (o *DataInnerDhcpScope6Data) GetServer6ClassParameters() []ApiClassParameterOutputEntry`

GetServer6ClassParameters returns the Server6ClassParameters field if non-nil, zero value otherwise.

### GetServer6ClassParametersOk

`func (o *DataInnerDhcpScope6Data) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassParameters

`func (o *DataInnerDhcpScope6Data) SetServer6ClassParameters(v []ApiClassParameterOutputEntry)`

SetServer6ClassParameters sets Server6ClassParameters field to given value.

### HasServer6ClassParameters

`func (o *DataInnerDhcpScope6Data) HasServer6ClassParameters() bool`

HasServer6ClassParameters returns a boolean if a field has been set.

### GetServer6Id

`func (o *DataInnerDhcpScope6Data) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DataInnerDhcpScope6Data) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DataInnerDhcpScope6Data) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DataInnerDhcpScope6Data) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DataInnerDhcpScope6Data) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DataInnerDhcpScope6Data) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DataInnerDhcpScope6Data) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DataInnerDhcpScope6Data) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DataInnerDhcpScope6Data) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DataInnerDhcpScope6Data) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DataInnerDhcpScope6Data) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DataInnerDhcpScope6Data) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetServer6Version

`func (o *DataInnerDhcpScope6Data) GetServer6Version() string`

GetServer6Version returns the Server6Version field if non-nil, zero value otherwise.

### GetServer6VersionOk

`func (o *DataInnerDhcpScope6Data) GetServer6VersionOk() (*string, bool)`

GetServer6VersionOk returns a tuple with the Server6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Version

`func (o *DataInnerDhcpScope6Data) SetServer6Version(v string)`

SetServer6Version sets Server6Version field to given value.

### HasServer6Version

`func (o *DataInnerDhcpScope6Data) HasServer6Version() bool`

HasServer6Version returns a boolean if a field has been set.

### GetFailover6Id

`func (o *DataInnerDhcpScope6Data) GetFailover6Id() string`

GetFailover6Id returns the Failover6Id field if non-nil, zero value otherwise.

### GetFailover6IdOk

`func (o *DataInnerDhcpScope6Data) GetFailover6IdOk() (*string, bool)`

GetFailover6IdOk returns a tuple with the Failover6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover6Id

`func (o *DataInnerDhcpScope6Data) SetFailover6Id(v string)`

SetFailover6Id sets Failover6Id field to given value.

### HasFailover6Id

`func (o *DataInnerDhcpScope6Data) HasFailover6Id() bool`

HasFailover6Id returns a boolean if a field has been set.

### GetFailover6Name

`func (o *DataInnerDhcpScope6Data) GetFailover6Name() string`

GetFailover6Name returns the Failover6Name field if non-nil, zero value otherwise.

### GetFailover6NameOk

`func (o *DataInnerDhcpScope6Data) GetFailover6NameOk() (*string, bool)`

GetFailover6NameOk returns a tuple with the Failover6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover6Name

`func (o *DataInnerDhcpScope6Data) SetFailover6Name(v string)`

SetFailover6Name sets Failover6Name field to given value.

### HasFailover6Name

`func (o *DataInnerDhcpScope6Data) HasFailover6Name() bool`

HasFailover6Name returns a boolean if a field has been set.

### GetScope6ClassName

`func (o *DataInnerDhcpScope6Data) GetScope6ClassName() string`

GetScope6ClassName returns the Scope6ClassName field if non-nil, zero value otherwise.

### GetScope6ClassNameOk

`func (o *DataInnerDhcpScope6Data) GetScope6ClassNameOk() (*string, bool)`

GetScope6ClassNameOk returns a tuple with the Scope6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassName

`func (o *DataInnerDhcpScope6Data) SetScope6ClassName(v string)`

SetScope6ClassName sets Scope6ClassName field to given value.

### HasScope6ClassName

`func (o *DataInnerDhcpScope6Data) HasScope6ClassName() bool`

HasScope6ClassName returns a boolean if a field has been set.

### GetScope6ClassParameters

`func (o *DataInnerDhcpScope6Data) GetScope6ClassParameters() []ApiClassParameterOutputEntry`

GetScope6ClassParameters returns the Scope6ClassParameters field if non-nil, zero value otherwise.

### GetScope6ClassParametersOk

`func (o *DataInnerDhcpScope6Data) GetScope6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScope6ClassParametersOk returns a tuple with the Scope6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassParameters

`func (o *DataInnerDhcpScope6Data) SetScope6ClassParameters(v []ApiClassParameterOutputEntry)`

SetScope6ClassParameters sets Scope6ClassParameters field to given value.

### HasScope6ClassParameters

`func (o *DataInnerDhcpScope6Data) HasScope6ClassParameters() bool`

HasScope6ClassParameters returns a boolean if a field has been set.

### GetScope6EndAddress6Addr

`func (o *DataInnerDhcpScope6Data) GetScope6EndAddress6Addr() string`

GetScope6EndAddress6Addr returns the Scope6EndAddress6Addr field if non-nil, zero value otherwise.

### GetScope6EndAddress6AddrOk

`func (o *DataInnerDhcpScope6Data) GetScope6EndAddress6AddrOk() (*string, bool)`

GetScope6EndAddress6AddrOk returns a tuple with the Scope6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6EndAddress6Addr

`func (o *DataInnerDhcpScope6Data) SetScope6EndAddress6Addr(v string)`

SetScope6EndAddress6Addr sets Scope6EndAddress6Addr field to given value.

### HasScope6EndAddress6Addr

`func (o *DataInnerDhcpScope6Data) HasScope6EndAddress6Addr() bool`

HasScope6EndAddress6Addr returns a boolean if a field has been set.

### GetScope6Id

`func (o *DataInnerDhcpScope6Data) GetScope6Id() string`

GetScope6Id returns the Scope6Id field if non-nil, zero value otherwise.

### GetScope6IdOk

`func (o *DataInnerDhcpScope6Data) GetScope6IdOk() (*string, bool)`

GetScope6IdOk returns a tuple with the Scope6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Id

`func (o *DataInnerDhcpScope6Data) SetScope6Id(v string)`

SetScope6Id sets Scope6Id field to given value.

### HasScope6Id

`func (o *DataInnerDhcpScope6Data) HasScope6Id() bool`

HasScope6Id returns a boolean if a field has been set.

### GetScope6Name

`func (o *DataInnerDhcpScope6Data) GetScope6Name() string`

GetScope6Name returns the Scope6Name field if non-nil, zero value otherwise.

### GetScope6NameOk

`func (o *DataInnerDhcpScope6Data) GetScope6NameOk() (*string, bool)`

GetScope6NameOk returns a tuple with the Scope6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Name

`func (o *DataInnerDhcpScope6Data) SetScope6Name(v string)`

SetScope6Name sets Scope6Name field to given value.

### HasScope6Name

`func (o *DataInnerDhcpScope6Data) HasScope6Name() bool`

HasScope6Name returns a boolean if a field has been set.

### GetScope6Prefix

`func (o *DataInnerDhcpScope6Data) GetScope6Prefix() string`

GetScope6Prefix returns the Scope6Prefix field if non-nil, zero value otherwise.

### GetScope6PrefixOk

`func (o *DataInnerDhcpScope6Data) GetScope6PrefixOk() (*string, bool)`

GetScope6PrefixOk returns a tuple with the Scope6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Prefix

`func (o *DataInnerDhcpScope6Data) SetScope6Prefix(v string)`

SetScope6Prefix sets Scope6Prefix field to given value.

### HasScope6Prefix

`func (o *DataInnerDhcpScope6Data) HasScope6Prefix() bool`

HasScope6Prefix returns a boolean if a field has been set.

### GetScope6SpaceId

`func (o *DataInnerDhcpScope6Data) GetScope6SpaceId() string`

GetScope6SpaceId returns the Scope6SpaceId field if non-nil, zero value otherwise.

### GetScope6SpaceIdOk

`func (o *DataInnerDhcpScope6Data) GetScope6SpaceIdOk() (*string, bool)`

GetScope6SpaceIdOk returns a tuple with the Scope6SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6SpaceId

`func (o *DataInnerDhcpScope6Data) SetScope6SpaceId(v string)`

SetScope6SpaceId sets Scope6SpaceId field to given value.

### HasScope6SpaceId

`func (o *DataInnerDhcpScope6Data) HasScope6SpaceId() bool`

HasScope6SpaceId returns a boolean if a field has been set.

### GetScope6SpaceName

`func (o *DataInnerDhcpScope6Data) GetScope6SpaceName() string`

GetScope6SpaceName returns the Scope6SpaceName field if non-nil, zero value otherwise.

### GetScope6SpaceNameOk

`func (o *DataInnerDhcpScope6Data) GetScope6SpaceNameOk() (*string, bool)`

GetScope6SpaceNameOk returns a tuple with the Scope6SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6SpaceName

`func (o *DataInnerDhcpScope6Data) SetScope6SpaceName(v string)`

SetScope6SpaceName sets Scope6SpaceName field to given value.

### HasScope6SpaceName

`func (o *DataInnerDhcpScope6Data) HasScope6SpaceName() bool`

HasScope6SpaceName returns a boolean if a field has been set.

### GetScope6Size

`func (o *DataInnerDhcpScope6Data) GetScope6Size() string`

GetScope6Size returns the Scope6Size field if non-nil, zero value otherwise.

### GetScope6SizeOk

`func (o *DataInnerDhcpScope6Data) GetScope6SizeOk() (*string, bool)`

GetScope6SizeOk returns a tuple with the Scope6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Size

`func (o *DataInnerDhcpScope6Data) SetScope6Size(v string)`

SetScope6Size sets Scope6Size field to given value.

### HasScope6Size

`func (o *DataInnerDhcpScope6Data) HasScope6Size() bool`

HasScope6Size returns a boolean if a field has been set.

### GetScope6SortName

`func (o *DataInnerDhcpScope6Data) GetScope6SortName() string`

GetScope6SortName returns the Scope6SortName field if non-nil, zero value otherwise.

### GetScope6SortNameOk

`func (o *DataInnerDhcpScope6Data) GetScope6SortNameOk() (*string, bool)`

GetScope6SortNameOk returns a tuple with the Scope6SortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6SortName

`func (o *DataInnerDhcpScope6Data) SetScope6SortName(v string)`

SetScope6SortName sets Scope6SortName field to given value.

### HasScope6SortName

`func (o *DataInnerDhcpScope6Data) HasScope6SortName() bool`

HasScope6SortName returns a boolean if a field has been set.

### GetScope6StartAddress6Addr

`func (o *DataInnerDhcpScope6Data) GetScope6StartAddress6Addr() string`

GetScope6StartAddress6Addr returns the Scope6StartAddress6Addr field if non-nil, zero value otherwise.

### GetScope6StartAddress6AddrOk

`func (o *DataInnerDhcpScope6Data) GetScope6StartAddress6AddrOk() (*string, bool)`

GetScope6StartAddress6AddrOk returns a tuple with the Scope6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6StartAddress6Addr

`func (o *DataInnerDhcpScope6Data) SetScope6StartAddress6Addr(v string)`

SetScope6StartAddress6Addr sets Scope6StartAddress6Addr field to given value.

### HasScope6StartAddress6Addr

`func (o *DataInnerDhcpScope6Data) HasScope6StartAddress6Addr() bool`

HasScope6StartAddress6Addr returns a boolean if a field has been set.

### GetSharednetwork6Id

`func (o *DataInnerDhcpScope6Data) GetSharednetwork6Id() string`

GetSharednetwork6Id returns the Sharednetwork6Id field if non-nil, zero value otherwise.

### GetSharednetwork6IdOk

`func (o *DataInnerDhcpScope6Data) GetSharednetwork6IdOk() (*string, bool)`

GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Id

`func (o *DataInnerDhcpScope6Data) SetSharednetwork6Id(v string)`

SetSharednetwork6Id sets Sharednetwork6Id field to given value.

### HasSharednetwork6Id

`func (o *DataInnerDhcpScope6Data) HasSharednetwork6Id() bool`

HasSharednetwork6Id returns a boolean if a field has been set.

### GetSharednetwork6Name

`func (o *DataInnerDhcpScope6Data) GetSharednetwork6Name() string`

GetSharednetwork6Name returns the Sharednetwork6Name field if non-nil, zero value otherwise.

### GetSharednetwork6NameOk

`func (o *DataInnerDhcpScope6Data) GetSharednetwork6NameOk() (*string, bool)`

GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Name

`func (o *DataInnerDhcpScope6Data) SetSharednetwork6Name(v string)`

SetSharednetwork6Name sets Sharednetwork6Name field to given value.

### HasSharednetwork6Name

`func (o *DataInnerDhcpScope6Data) HasSharednetwork6Name() bool`

HasSharednetwork6Name returns a boolean if a field has been set.

### GetServer6Hostaddr

`func (o *DataInnerDhcpScope6Data) GetServer6Hostaddr() string`

GetServer6Hostaddr returns the Server6Hostaddr field if non-nil, zero value otherwise.

### GetServer6HostaddrOk

`func (o *DataInnerDhcpScope6Data) GetServer6HostaddrOk() (*string, bool)`

GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Hostaddr

`func (o *DataInnerDhcpScope6Data) SetServer6Hostaddr(v string)`

SetServer6Hostaddr sets Server6Hostaddr field to given value.

### HasServer6Hostaddr

`func (o *DataInnerDhcpScope6Data) HasServer6Hostaddr() bool`

HasServer6Hostaddr returns a boolean if a field has been set.

### GetServer6Addr6

`func (o *DataInnerDhcpScope6Data) GetServer6Addr6() string`

GetServer6Addr6 returns the Server6Addr6 field if non-nil, zero value otherwise.

### GetServer6Addr6Ok

`func (o *DataInnerDhcpScope6Data) GetServer6Addr6Ok() (*string, bool)`

GetServer6Addr6Ok returns a tuple with the Server6Addr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr6

`func (o *DataInnerDhcpScope6Data) SetServer6Addr6(v string)`

SetServer6Addr6 sets Server6Addr6 field to given value.

### HasServer6Addr6

`func (o *DataInnerDhcpScope6Data) HasServer6Addr6() bool`

HasServer6Addr6 returns a boolean if a field has been set.

### GetServer6Addr

`func (o *DataInnerDhcpScope6Data) GetServer6Addr() string`

GetServer6Addr returns the Server6Addr field if non-nil, zero value otherwise.

### GetServer6AddrOk

`func (o *DataInnerDhcpScope6Data) GetServer6AddrOk() (*string, bool)`

GetServer6AddrOk returns a tuple with the Server6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr

`func (o *DataInnerDhcpScope6Data) SetServer6Addr(v string)`

SetServer6Addr sets Server6Addr field to given value.

### HasServer6Addr

`func (o *DataInnerDhcpScope6Data) HasServer6Addr() bool`

HasServer6Addr returns a boolean if a field has been set.

### GetScope6Multistatus

`func (o *DataInnerDhcpScope6Data) GetScope6Multistatus() string`

GetScope6Multistatus returns the Scope6Multistatus field if non-nil, zero value otherwise.

### GetScope6MultistatusOk

`func (o *DataInnerDhcpScope6Data) GetScope6MultistatusOk() (*string, bool)`

GetScope6MultistatusOk returns a tuple with the Scope6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Multistatus

`func (o *DataInnerDhcpScope6Data) SetScope6Multistatus(v string)`

SetScope6Multistatus sets Scope6Multistatus field to given value.

### HasScope6Multistatus

`func (o *DataInnerDhcpScope6Data) HasScope6Multistatus() bool`

HasScope6Multistatus returns a boolean if a field has been set.

### GetSmartArch

`func (o *DataInnerDhcpScope6Data) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DataInnerDhcpScope6Data) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DataInnerDhcpScope6Data) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DataInnerDhcpScope6Data) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpScope6Data) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpScope6Data) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpScope6Data) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpScope6Data) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpScope6Data) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpScope6Data) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpScope6Data) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpScope6Data) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


