# DataInnerDhcpStatic6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Static6DelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**Server6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 server the object belongs to, it can be preceded by the class directory. | [optional] 
**Server6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 server the object belongs to. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;server6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv6 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Version** | Pointer to **string** | The version details of the DHCPv6 server the object belongs to. | [optional] 
**Group6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 group the static belongs to, it can be preceded by the class directory. | [optional] 
**Group6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 group the static belongs to,. | [optional] 
**Group6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 group. | [optional] 
**Group6Name** | Pointer to **string** | The name of the DHCPv6 group associated with the object. | [optional] 
**Static6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6static, it can be preceded by the class directory. | [optional] 
**Static6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 static. | [optional] 
**Static6ClientDuid** | Pointer to **string** | The client DHCP Unique Identifier (DUID) associated with the DHCPv6 static. | [optional] 
**Static6Domain** | Pointer to **string** | The domain name associated with the DHCPv6 static. | [optional] 
**Static6EndTime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Static6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 static. | [optional] 
**Static6Address6Addr** | Pointer to **string** | The IP address associated with the DHCPv6 static, in hexadecimal format. | [optional] 
**Static6MacAddr** | Pointer to **string** | The MAC address associated with the DHCPv6 static. | [optional] 
**Static6Name** | Pointer to **string** | The name of the DHCPv6 static. | [optional] 
**Static6Prefix6** | Pointer to **string** | The IP address and prefix of the delegated prefix of the DHCPv6 static, in hexadecimal format as follows: &lt;b&gt;&amp;lt;IPv6-address&amp;gt;/&amp;lt;prefix&amp;gt;&lt;/b&gt;. | [optional] 
**Static6Prefix6Addr** | Pointer to **string** | The IP address of the delegated prefix of the DHCPv6 static, in hexadecimal format. | [optional] 
**Static6Prefix6Prefix** | Pointer to **string** | The prefix of the delegated prefix of the DHCPv6 static. | [optional] 
**Static6State** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Static6Time** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Scope6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**Scope6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 scope the object belongs to. | [optional] 
**Scope6EndAddress6Addr** | Pointer to **string** | The last IP address of the DHCPv6 scope the object belongs to, in hexadecimal format. | [optional] 
**Scope6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Name** | Pointer to **string** | The name of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6SpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Size** | Pointer to **string** | The number of IP addresses the DHCPv6 scope the object belongs to contains. | [optional] 
**Scope6StartAddress6Addr** | Pointer to **string** | The first IP address of the DHCPv6 scope the object belongs to, in hexadecimal format. | [optional] 
**Sharednetwork6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 shared network the object belongs to. | [optional] 
**Sharednetwork6Name** | Pointer to **string** | The name of the DHCPv6 shared network the object belongs to. | [optional] 
**Static6MacVendor** | Pointer to **string** | The vendor details of the client associated with the DHCPv6 static. | [optional] 
**Static6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpStatic6Data

`func NewDataInnerDhcpStatic6Data() *DataInnerDhcpStatic6Data`

NewDataInnerDhcpStatic6Data instantiates a new DataInnerDhcpStatic6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpStatic6DataWithDefaults

`func NewDataInnerDhcpStatic6DataWithDefaults() *DataInnerDhcpStatic6Data`

NewDataInnerDhcpStatic6DataWithDefaults instantiates a new DataInnerDhcpStatic6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatic6DelayedTime

`func (o *DataInnerDhcpStatic6Data) GetStatic6DelayedTime() string`

GetStatic6DelayedTime returns the Static6DelayedTime field if non-nil, zero value otherwise.

### GetStatic6DelayedTimeOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6DelayedTimeOk() (*string, bool)`

GetStatic6DelayedTimeOk returns a tuple with the Static6DelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6DelayedTime

`func (o *DataInnerDhcpStatic6Data) SetStatic6DelayedTime(v string)`

SetStatic6DelayedTime sets Static6DelayedTime field to given value.

### HasStatic6DelayedTime

`func (o *DataInnerDhcpStatic6Data) HasStatic6DelayedTime() bool`

HasStatic6DelayedTime returns a boolean if a field has been set.

### GetServer6ClassName

`func (o *DataInnerDhcpStatic6Data) GetServer6ClassName() string`

GetServer6ClassName returns the Server6ClassName field if non-nil, zero value otherwise.

### GetServer6ClassNameOk

`func (o *DataInnerDhcpStatic6Data) GetServer6ClassNameOk() (*string, bool)`

GetServer6ClassNameOk returns a tuple with the Server6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassName

`func (o *DataInnerDhcpStatic6Data) SetServer6ClassName(v string)`

SetServer6ClassName sets Server6ClassName field to given value.

### HasServer6ClassName

`func (o *DataInnerDhcpStatic6Data) HasServer6ClassName() bool`

HasServer6ClassName returns a boolean if a field has been set.

### GetServer6ClassParameters

`func (o *DataInnerDhcpStatic6Data) GetServer6ClassParameters() []ApiClassParameterOutputEntry`

GetServer6ClassParameters returns the Server6ClassParameters field if non-nil, zero value otherwise.

### GetServer6ClassParametersOk

`func (o *DataInnerDhcpStatic6Data) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassParameters

`func (o *DataInnerDhcpStatic6Data) SetServer6ClassParameters(v []ApiClassParameterOutputEntry)`

SetServer6ClassParameters sets Server6ClassParameters field to given value.

### HasServer6ClassParameters

`func (o *DataInnerDhcpStatic6Data) HasServer6ClassParameters() bool`

HasServer6ClassParameters returns a boolean if a field has been set.

### GetServer6Id

`func (o *DataInnerDhcpStatic6Data) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DataInnerDhcpStatic6Data) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DataInnerDhcpStatic6Data) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DataInnerDhcpStatic6Data) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DataInnerDhcpStatic6Data) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DataInnerDhcpStatic6Data) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DataInnerDhcpStatic6Data) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DataInnerDhcpStatic6Data) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DataInnerDhcpStatic6Data) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DataInnerDhcpStatic6Data) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DataInnerDhcpStatic6Data) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DataInnerDhcpStatic6Data) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetServer6Version

`func (o *DataInnerDhcpStatic6Data) GetServer6Version() string`

GetServer6Version returns the Server6Version field if non-nil, zero value otherwise.

### GetServer6VersionOk

`func (o *DataInnerDhcpStatic6Data) GetServer6VersionOk() (*string, bool)`

GetServer6VersionOk returns a tuple with the Server6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Version

`func (o *DataInnerDhcpStatic6Data) SetServer6Version(v string)`

SetServer6Version sets Server6Version field to given value.

### HasServer6Version

`func (o *DataInnerDhcpStatic6Data) HasServer6Version() bool`

HasServer6Version returns a boolean if a field has been set.

### GetGroup6ClassName

`func (o *DataInnerDhcpStatic6Data) GetGroup6ClassName() string`

GetGroup6ClassName returns the Group6ClassName field if non-nil, zero value otherwise.

### GetGroup6ClassNameOk

`func (o *DataInnerDhcpStatic6Data) GetGroup6ClassNameOk() (*string, bool)`

GetGroup6ClassNameOk returns a tuple with the Group6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6ClassName

`func (o *DataInnerDhcpStatic6Data) SetGroup6ClassName(v string)`

SetGroup6ClassName sets Group6ClassName field to given value.

### HasGroup6ClassName

`func (o *DataInnerDhcpStatic6Data) HasGroup6ClassName() bool`

HasGroup6ClassName returns a boolean if a field has been set.

### GetGroup6ClassParameters

`func (o *DataInnerDhcpStatic6Data) GetGroup6ClassParameters() []ApiClassParameterOutputEntry`

GetGroup6ClassParameters returns the Group6ClassParameters field if non-nil, zero value otherwise.

### GetGroup6ClassParametersOk

`func (o *DataInnerDhcpStatic6Data) GetGroup6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetGroup6ClassParametersOk returns a tuple with the Group6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6ClassParameters

`func (o *DataInnerDhcpStatic6Data) SetGroup6ClassParameters(v []ApiClassParameterOutputEntry)`

SetGroup6ClassParameters sets Group6ClassParameters field to given value.

### HasGroup6ClassParameters

`func (o *DataInnerDhcpStatic6Data) HasGroup6ClassParameters() bool`

HasGroup6ClassParameters returns a boolean if a field has been set.

### GetGroup6Id

`func (o *DataInnerDhcpStatic6Data) GetGroup6Id() string`

GetGroup6Id returns the Group6Id field if non-nil, zero value otherwise.

### GetGroup6IdOk

`func (o *DataInnerDhcpStatic6Data) GetGroup6IdOk() (*string, bool)`

GetGroup6IdOk returns a tuple with the Group6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Id

`func (o *DataInnerDhcpStatic6Data) SetGroup6Id(v string)`

SetGroup6Id sets Group6Id field to given value.

### HasGroup6Id

`func (o *DataInnerDhcpStatic6Data) HasGroup6Id() bool`

HasGroup6Id returns a boolean if a field has been set.

### GetGroup6Name

`func (o *DataInnerDhcpStatic6Data) GetGroup6Name() string`

GetGroup6Name returns the Group6Name field if non-nil, zero value otherwise.

### GetGroup6NameOk

`func (o *DataInnerDhcpStatic6Data) GetGroup6NameOk() (*string, bool)`

GetGroup6NameOk returns a tuple with the Group6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Name

`func (o *DataInnerDhcpStatic6Data) SetGroup6Name(v string)`

SetGroup6Name sets Group6Name field to given value.

### HasGroup6Name

`func (o *DataInnerDhcpStatic6Data) HasGroup6Name() bool`

HasGroup6Name returns a boolean if a field has been set.

### GetStatic6ClassName

`func (o *DataInnerDhcpStatic6Data) GetStatic6ClassName() string`

GetStatic6ClassName returns the Static6ClassName field if non-nil, zero value otherwise.

### GetStatic6ClassNameOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6ClassNameOk() (*string, bool)`

GetStatic6ClassNameOk returns a tuple with the Static6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6ClassName

`func (o *DataInnerDhcpStatic6Data) SetStatic6ClassName(v string)`

SetStatic6ClassName sets Static6ClassName field to given value.

### HasStatic6ClassName

`func (o *DataInnerDhcpStatic6Data) HasStatic6ClassName() bool`

HasStatic6ClassName returns a boolean if a field has been set.

### GetStatic6ClassParameters

`func (o *DataInnerDhcpStatic6Data) GetStatic6ClassParameters() []ApiClassParameterOutputEntry`

GetStatic6ClassParameters returns the Static6ClassParameters field if non-nil, zero value otherwise.

### GetStatic6ClassParametersOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetStatic6ClassParametersOk returns a tuple with the Static6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6ClassParameters

`func (o *DataInnerDhcpStatic6Data) SetStatic6ClassParameters(v []ApiClassParameterOutputEntry)`

SetStatic6ClassParameters sets Static6ClassParameters field to given value.

### HasStatic6ClassParameters

`func (o *DataInnerDhcpStatic6Data) HasStatic6ClassParameters() bool`

HasStatic6ClassParameters returns a boolean if a field has been set.

### GetStatic6ClientDuid

`func (o *DataInnerDhcpStatic6Data) GetStatic6ClientDuid() string`

GetStatic6ClientDuid returns the Static6ClientDuid field if non-nil, zero value otherwise.

### GetStatic6ClientDuidOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6ClientDuidOk() (*string, bool)`

GetStatic6ClientDuidOk returns a tuple with the Static6ClientDuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6ClientDuid

`func (o *DataInnerDhcpStatic6Data) SetStatic6ClientDuid(v string)`

SetStatic6ClientDuid sets Static6ClientDuid field to given value.

### HasStatic6ClientDuid

`func (o *DataInnerDhcpStatic6Data) HasStatic6ClientDuid() bool`

HasStatic6ClientDuid returns a boolean if a field has been set.

### GetStatic6Domain

`func (o *DataInnerDhcpStatic6Data) GetStatic6Domain() string`

GetStatic6Domain returns the Static6Domain field if non-nil, zero value otherwise.

### GetStatic6DomainOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6DomainOk() (*string, bool)`

GetStatic6DomainOk returns a tuple with the Static6Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Domain

`func (o *DataInnerDhcpStatic6Data) SetStatic6Domain(v string)`

SetStatic6Domain sets Static6Domain field to given value.

### HasStatic6Domain

`func (o *DataInnerDhcpStatic6Data) HasStatic6Domain() bool`

HasStatic6Domain returns a boolean if a field has been set.

### GetStatic6EndTime

`func (o *DataInnerDhcpStatic6Data) GetStatic6EndTime() string`

GetStatic6EndTime returns the Static6EndTime field if non-nil, zero value otherwise.

### GetStatic6EndTimeOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6EndTimeOk() (*string, bool)`

GetStatic6EndTimeOk returns a tuple with the Static6EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6EndTime

`func (o *DataInnerDhcpStatic6Data) SetStatic6EndTime(v string)`

SetStatic6EndTime sets Static6EndTime field to given value.

### HasStatic6EndTime

`func (o *DataInnerDhcpStatic6Data) HasStatic6EndTime() bool`

HasStatic6EndTime returns a boolean if a field has been set.

### GetStatic6Id

`func (o *DataInnerDhcpStatic6Data) GetStatic6Id() string`

GetStatic6Id returns the Static6Id field if non-nil, zero value otherwise.

### GetStatic6IdOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6IdOk() (*string, bool)`

GetStatic6IdOk returns a tuple with the Static6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Id

`func (o *DataInnerDhcpStatic6Data) SetStatic6Id(v string)`

SetStatic6Id sets Static6Id field to given value.

### HasStatic6Id

`func (o *DataInnerDhcpStatic6Data) HasStatic6Id() bool`

HasStatic6Id returns a boolean if a field has been set.

### GetStatic6Address6Addr

`func (o *DataInnerDhcpStatic6Data) GetStatic6Address6Addr() string`

GetStatic6Address6Addr returns the Static6Address6Addr field if non-nil, zero value otherwise.

### GetStatic6Address6AddrOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6Address6AddrOk() (*string, bool)`

GetStatic6Address6AddrOk returns a tuple with the Static6Address6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Address6Addr

`func (o *DataInnerDhcpStatic6Data) SetStatic6Address6Addr(v string)`

SetStatic6Address6Addr sets Static6Address6Addr field to given value.

### HasStatic6Address6Addr

`func (o *DataInnerDhcpStatic6Data) HasStatic6Address6Addr() bool`

HasStatic6Address6Addr returns a boolean if a field has been set.

### GetStatic6MacAddr

`func (o *DataInnerDhcpStatic6Data) GetStatic6MacAddr() string`

GetStatic6MacAddr returns the Static6MacAddr field if non-nil, zero value otherwise.

### GetStatic6MacAddrOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6MacAddrOk() (*string, bool)`

GetStatic6MacAddrOk returns a tuple with the Static6MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6MacAddr

`func (o *DataInnerDhcpStatic6Data) SetStatic6MacAddr(v string)`

SetStatic6MacAddr sets Static6MacAddr field to given value.

### HasStatic6MacAddr

`func (o *DataInnerDhcpStatic6Data) HasStatic6MacAddr() bool`

HasStatic6MacAddr returns a boolean if a field has been set.

### GetStatic6Name

`func (o *DataInnerDhcpStatic6Data) GetStatic6Name() string`

GetStatic6Name returns the Static6Name field if non-nil, zero value otherwise.

### GetStatic6NameOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6NameOk() (*string, bool)`

GetStatic6NameOk returns a tuple with the Static6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Name

`func (o *DataInnerDhcpStatic6Data) SetStatic6Name(v string)`

SetStatic6Name sets Static6Name field to given value.

### HasStatic6Name

`func (o *DataInnerDhcpStatic6Data) HasStatic6Name() bool`

HasStatic6Name returns a boolean if a field has been set.

### GetStatic6Prefix6

`func (o *DataInnerDhcpStatic6Data) GetStatic6Prefix6() string`

GetStatic6Prefix6 returns the Static6Prefix6 field if non-nil, zero value otherwise.

### GetStatic6Prefix6Ok

`func (o *DataInnerDhcpStatic6Data) GetStatic6Prefix6Ok() (*string, bool)`

GetStatic6Prefix6Ok returns a tuple with the Static6Prefix6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Prefix6

`func (o *DataInnerDhcpStatic6Data) SetStatic6Prefix6(v string)`

SetStatic6Prefix6 sets Static6Prefix6 field to given value.

### HasStatic6Prefix6

`func (o *DataInnerDhcpStatic6Data) HasStatic6Prefix6() bool`

HasStatic6Prefix6 returns a boolean if a field has been set.

### GetStatic6Prefix6Addr

`func (o *DataInnerDhcpStatic6Data) GetStatic6Prefix6Addr() string`

GetStatic6Prefix6Addr returns the Static6Prefix6Addr field if non-nil, zero value otherwise.

### GetStatic6Prefix6AddrOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6Prefix6AddrOk() (*string, bool)`

GetStatic6Prefix6AddrOk returns a tuple with the Static6Prefix6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Prefix6Addr

`func (o *DataInnerDhcpStatic6Data) SetStatic6Prefix6Addr(v string)`

SetStatic6Prefix6Addr sets Static6Prefix6Addr field to given value.

### HasStatic6Prefix6Addr

`func (o *DataInnerDhcpStatic6Data) HasStatic6Prefix6Addr() bool`

HasStatic6Prefix6Addr returns a boolean if a field has been set.

### GetStatic6Prefix6Prefix

`func (o *DataInnerDhcpStatic6Data) GetStatic6Prefix6Prefix() string`

GetStatic6Prefix6Prefix returns the Static6Prefix6Prefix field if non-nil, zero value otherwise.

### GetStatic6Prefix6PrefixOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6Prefix6PrefixOk() (*string, bool)`

GetStatic6Prefix6PrefixOk returns a tuple with the Static6Prefix6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Prefix6Prefix

`func (o *DataInnerDhcpStatic6Data) SetStatic6Prefix6Prefix(v string)`

SetStatic6Prefix6Prefix sets Static6Prefix6Prefix field to given value.

### HasStatic6Prefix6Prefix

`func (o *DataInnerDhcpStatic6Data) HasStatic6Prefix6Prefix() bool`

HasStatic6Prefix6Prefix returns a boolean if a field has been set.

### GetStatic6State

`func (o *DataInnerDhcpStatic6Data) GetStatic6State() string`

GetStatic6State returns the Static6State field if non-nil, zero value otherwise.

### GetStatic6StateOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6StateOk() (*string, bool)`

GetStatic6StateOk returns a tuple with the Static6State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6State

`func (o *DataInnerDhcpStatic6Data) SetStatic6State(v string)`

SetStatic6State sets Static6State field to given value.

### HasStatic6State

`func (o *DataInnerDhcpStatic6Data) HasStatic6State() bool`

HasStatic6State returns a boolean if a field has been set.

### GetStatic6Time

`func (o *DataInnerDhcpStatic6Data) GetStatic6Time() string`

GetStatic6Time returns the Static6Time field if non-nil, zero value otherwise.

### GetStatic6TimeOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6TimeOk() (*string, bool)`

GetStatic6TimeOk returns a tuple with the Static6Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Time

`func (o *DataInnerDhcpStatic6Data) SetStatic6Time(v string)`

SetStatic6Time sets Static6Time field to given value.

### HasStatic6Time

`func (o *DataInnerDhcpStatic6Data) HasStatic6Time() bool`

HasStatic6Time returns a boolean if a field has been set.

### GetScope6ClassName

`func (o *DataInnerDhcpStatic6Data) GetScope6ClassName() string`

GetScope6ClassName returns the Scope6ClassName field if non-nil, zero value otherwise.

### GetScope6ClassNameOk

`func (o *DataInnerDhcpStatic6Data) GetScope6ClassNameOk() (*string, bool)`

GetScope6ClassNameOk returns a tuple with the Scope6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassName

`func (o *DataInnerDhcpStatic6Data) SetScope6ClassName(v string)`

SetScope6ClassName sets Scope6ClassName field to given value.

### HasScope6ClassName

`func (o *DataInnerDhcpStatic6Data) HasScope6ClassName() bool`

HasScope6ClassName returns a boolean if a field has been set.

### GetScope6ClassParameters

`func (o *DataInnerDhcpStatic6Data) GetScope6ClassParameters() []ApiClassParameterOutputEntry`

GetScope6ClassParameters returns the Scope6ClassParameters field if non-nil, zero value otherwise.

### GetScope6ClassParametersOk

`func (o *DataInnerDhcpStatic6Data) GetScope6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScope6ClassParametersOk returns a tuple with the Scope6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassParameters

`func (o *DataInnerDhcpStatic6Data) SetScope6ClassParameters(v []ApiClassParameterOutputEntry)`

SetScope6ClassParameters sets Scope6ClassParameters field to given value.

### HasScope6ClassParameters

`func (o *DataInnerDhcpStatic6Data) HasScope6ClassParameters() bool`

HasScope6ClassParameters returns a boolean if a field has been set.

### GetScope6EndAddress6Addr

`func (o *DataInnerDhcpStatic6Data) GetScope6EndAddress6Addr() string`

GetScope6EndAddress6Addr returns the Scope6EndAddress6Addr field if non-nil, zero value otherwise.

### GetScope6EndAddress6AddrOk

`func (o *DataInnerDhcpStatic6Data) GetScope6EndAddress6AddrOk() (*string, bool)`

GetScope6EndAddress6AddrOk returns a tuple with the Scope6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6EndAddress6Addr

`func (o *DataInnerDhcpStatic6Data) SetScope6EndAddress6Addr(v string)`

SetScope6EndAddress6Addr sets Scope6EndAddress6Addr field to given value.

### HasScope6EndAddress6Addr

`func (o *DataInnerDhcpStatic6Data) HasScope6EndAddress6Addr() bool`

HasScope6EndAddress6Addr returns a boolean if a field has been set.

### GetScope6Id

`func (o *DataInnerDhcpStatic6Data) GetScope6Id() string`

GetScope6Id returns the Scope6Id field if non-nil, zero value otherwise.

### GetScope6IdOk

`func (o *DataInnerDhcpStatic6Data) GetScope6IdOk() (*string, bool)`

GetScope6IdOk returns a tuple with the Scope6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Id

`func (o *DataInnerDhcpStatic6Data) SetScope6Id(v string)`

SetScope6Id sets Scope6Id field to given value.

### HasScope6Id

`func (o *DataInnerDhcpStatic6Data) HasScope6Id() bool`

HasScope6Id returns a boolean if a field has been set.

### GetScope6Name

`func (o *DataInnerDhcpStatic6Data) GetScope6Name() string`

GetScope6Name returns the Scope6Name field if non-nil, zero value otherwise.

### GetScope6NameOk

`func (o *DataInnerDhcpStatic6Data) GetScope6NameOk() (*string, bool)`

GetScope6NameOk returns a tuple with the Scope6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Name

`func (o *DataInnerDhcpStatic6Data) SetScope6Name(v string)`

SetScope6Name sets Scope6Name field to given value.

### HasScope6Name

`func (o *DataInnerDhcpStatic6Data) HasScope6Name() bool`

HasScope6Name returns a boolean if a field has been set.

### GetScope6SpaceId

`func (o *DataInnerDhcpStatic6Data) GetScope6SpaceId() string`

GetScope6SpaceId returns the Scope6SpaceId field if non-nil, zero value otherwise.

### GetScope6SpaceIdOk

`func (o *DataInnerDhcpStatic6Data) GetScope6SpaceIdOk() (*string, bool)`

GetScope6SpaceIdOk returns a tuple with the Scope6SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6SpaceId

`func (o *DataInnerDhcpStatic6Data) SetScope6SpaceId(v string)`

SetScope6SpaceId sets Scope6SpaceId field to given value.

### HasScope6SpaceId

`func (o *DataInnerDhcpStatic6Data) HasScope6SpaceId() bool`

HasScope6SpaceId returns a boolean if a field has been set.

### GetScope6Size

`func (o *DataInnerDhcpStatic6Data) GetScope6Size() string`

GetScope6Size returns the Scope6Size field if non-nil, zero value otherwise.

### GetScope6SizeOk

`func (o *DataInnerDhcpStatic6Data) GetScope6SizeOk() (*string, bool)`

GetScope6SizeOk returns a tuple with the Scope6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Size

`func (o *DataInnerDhcpStatic6Data) SetScope6Size(v string)`

SetScope6Size sets Scope6Size field to given value.

### HasScope6Size

`func (o *DataInnerDhcpStatic6Data) HasScope6Size() bool`

HasScope6Size returns a boolean if a field has been set.

### GetScope6StartAddress6Addr

`func (o *DataInnerDhcpStatic6Data) GetScope6StartAddress6Addr() string`

GetScope6StartAddress6Addr returns the Scope6StartAddress6Addr field if non-nil, zero value otherwise.

### GetScope6StartAddress6AddrOk

`func (o *DataInnerDhcpStatic6Data) GetScope6StartAddress6AddrOk() (*string, bool)`

GetScope6StartAddress6AddrOk returns a tuple with the Scope6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6StartAddress6Addr

`func (o *DataInnerDhcpStatic6Data) SetScope6StartAddress6Addr(v string)`

SetScope6StartAddress6Addr sets Scope6StartAddress6Addr field to given value.

### HasScope6StartAddress6Addr

`func (o *DataInnerDhcpStatic6Data) HasScope6StartAddress6Addr() bool`

HasScope6StartAddress6Addr returns a boolean if a field has been set.

### GetSharednetwork6Id

`func (o *DataInnerDhcpStatic6Data) GetSharednetwork6Id() string`

GetSharednetwork6Id returns the Sharednetwork6Id field if non-nil, zero value otherwise.

### GetSharednetwork6IdOk

`func (o *DataInnerDhcpStatic6Data) GetSharednetwork6IdOk() (*string, bool)`

GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Id

`func (o *DataInnerDhcpStatic6Data) SetSharednetwork6Id(v string)`

SetSharednetwork6Id sets Sharednetwork6Id field to given value.

### HasSharednetwork6Id

`func (o *DataInnerDhcpStatic6Data) HasSharednetwork6Id() bool`

HasSharednetwork6Id returns a boolean if a field has been set.

### GetSharednetwork6Name

`func (o *DataInnerDhcpStatic6Data) GetSharednetwork6Name() string`

GetSharednetwork6Name returns the Sharednetwork6Name field if non-nil, zero value otherwise.

### GetSharednetwork6NameOk

`func (o *DataInnerDhcpStatic6Data) GetSharednetwork6NameOk() (*string, bool)`

GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Name

`func (o *DataInnerDhcpStatic6Data) SetSharednetwork6Name(v string)`

SetSharednetwork6Name sets Sharednetwork6Name field to given value.

### HasSharednetwork6Name

`func (o *DataInnerDhcpStatic6Data) HasSharednetwork6Name() bool`

HasSharednetwork6Name returns a boolean if a field has been set.

### GetStatic6MacVendor

`func (o *DataInnerDhcpStatic6Data) GetStatic6MacVendor() string`

GetStatic6MacVendor returns the Static6MacVendor field if non-nil, zero value otherwise.

### GetStatic6MacVendorOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6MacVendorOk() (*string, bool)`

GetStatic6MacVendorOk returns a tuple with the Static6MacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6MacVendor

`func (o *DataInnerDhcpStatic6Data) SetStatic6MacVendor(v string)`

SetStatic6MacVendor sets Static6MacVendor field to given value.

### HasStatic6MacVendor

`func (o *DataInnerDhcpStatic6Data) HasStatic6MacVendor() bool`

HasStatic6MacVendor returns a boolean if a field has been set.

### GetStatic6Multistatus

`func (o *DataInnerDhcpStatic6Data) GetStatic6Multistatus() string`

GetStatic6Multistatus returns the Static6Multistatus field if non-nil, zero value otherwise.

### GetStatic6MultistatusOk

`func (o *DataInnerDhcpStatic6Data) GetStatic6MultistatusOk() (*string, bool)`

GetStatic6MultistatusOk returns a tuple with the Static6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Multistatus

`func (o *DataInnerDhcpStatic6Data) SetStatic6Multistatus(v string)`

SetStatic6Multistatus sets Static6Multistatus field to given value.

### HasStatic6Multistatus

`func (o *DataInnerDhcpStatic6Data) HasStatic6Multistatus() bool`

HasStatic6Multistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpStatic6Data) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpStatic6Data) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpStatic6Data) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpStatic6Data) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpStatic6Data) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpStatic6Data) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpStatic6Data) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpStatic6Data) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


