# DhcpRange6DataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range6DelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**Server6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 server the object belongs to, it can be preceded by the class directory. | [optional] 
**Server6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 server the object belongs to and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Server6Comment** | Pointer to **string** | The description of the DHCPv6 server the object belongs to. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Version** | Pointer to **string** | The version details of the DHCPv6 server the object belongs to. | [optional] 
**Range6Acl** | Pointer to **string** | The list of ACLs associated with the DHCPv6 range, as follows: &lt;b&gt;&amp;lt;ACL_name&amp;gt;;&amp;lt;ACL_name&amp;gt;;...&lt;/b&gt; . | [optional] 
**Range6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 range, it can be preceded by the class directory. | [optional] 
**Range6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 range and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Range6EndAddress6Addr** | Pointer to **string** | The last IP address of the DHCPv6 range, in hexadecimal format. | [optional] 
**Range6FailoverName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Range6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 range. | [optional] 
**Range6LeaseCount** | Pointer to **string** | The total number of leases currently delivered by the DHCPv6 range. | [optional] 
**Range6Size** | Pointer to **string** | The number of IP addresses the DHCPv6 range contains. | [optional] 
**Range6StartAddress6Addr** | Pointer to **string** | The first IP address of the DHCPv6 range, in hexadecimal format. | [optional] 
**Range6State** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Scope6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**Scope6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 scope the object belongs to and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Scope6EndAddress6Addr** | Pointer to **string** | The last IP address of the DHCPv6 scope the object belongs to, in hexadecimal format. | [optional] 
**Scope6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Name** | Pointer to **string** | The name of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Prefix** | Pointer to **string** | The prefix of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6SpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Size** | Pointer to **string** | The number of IP addresses the DHCPv6 scope the object belongs to contains. | [optional] 
**Scope6StartAddress6Addr** | Pointer to **string** | The first IP address of the DHCPv6 scope the object belongs to, in hexadecimal format. | [optional] 
**Sharednetwork6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 shared network the object belongs to. | [optional] 
**Sharednetwork6Name** | Pointer to **string** | The name of the DHCPv6 shared network the object belongs to. | [optional] 
**Server6Addr** | Pointer to **string** | The IP address of the DHCP server the object belongs to, in hexadecimal format. | [optional] 
**Range6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpRange6DataData

`func NewDhcpRange6DataData() *DhcpRange6DataData`

NewDhcpRange6DataData instantiates a new DhcpRange6DataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpRange6DataDataWithDefaults

`func NewDhcpRange6DataDataWithDefaults() *DhcpRange6DataData`

NewDhcpRange6DataDataWithDefaults instantiates a new DhcpRange6DataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange6DelayedTime

`func (o *DhcpRange6DataData) GetRange6DelayedTime() string`

GetRange6DelayedTime returns the Range6DelayedTime field if non-nil, zero value otherwise.

### GetRange6DelayedTimeOk

`func (o *DhcpRange6DataData) GetRange6DelayedTimeOk() (*string, bool)`

GetRange6DelayedTimeOk returns a tuple with the Range6DelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6DelayedTime

`func (o *DhcpRange6DataData) SetRange6DelayedTime(v string)`

SetRange6DelayedTime sets Range6DelayedTime field to given value.

### HasRange6DelayedTime

`func (o *DhcpRange6DataData) HasRange6DelayedTime() bool`

HasRange6DelayedTime returns a boolean if a field has been set.

### GetServer6ClassName

`func (o *DhcpRange6DataData) GetServer6ClassName() string`

GetServer6ClassName returns the Server6ClassName field if non-nil, zero value otherwise.

### GetServer6ClassNameOk

`func (o *DhcpRange6DataData) GetServer6ClassNameOk() (*string, bool)`

GetServer6ClassNameOk returns a tuple with the Server6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassName

`func (o *DhcpRange6DataData) SetServer6ClassName(v string)`

SetServer6ClassName sets Server6ClassName field to given value.

### HasServer6ClassName

`func (o *DhcpRange6DataData) HasServer6ClassName() bool`

HasServer6ClassName returns a boolean if a field has been set.

### GetServer6ClassParameters

`func (o *DhcpRange6DataData) GetServer6ClassParameters() []ApiClassParameterOutputEntry`

GetServer6ClassParameters returns the Server6ClassParameters field if non-nil, zero value otherwise.

### GetServer6ClassParametersOk

`func (o *DhcpRange6DataData) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassParameters

`func (o *DhcpRange6DataData) SetServer6ClassParameters(v []ApiClassParameterOutputEntry)`

SetServer6ClassParameters sets Server6ClassParameters field to given value.

### HasServer6ClassParameters

`func (o *DhcpRange6DataData) HasServer6ClassParameters() bool`

HasServer6ClassParameters returns a boolean if a field has been set.

### GetServer6Comment

`func (o *DhcpRange6DataData) GetServer6Comment() string`

GetServer6Comment returns the Server6Comment field if non-nil, zero value otherwise.

### GetServer6CommentOk

`func (o *DhcpRange6DataData) GetServer6CommentOk() (*string, bool)`

GetServer6CommentOk returns a tuple with the Server6Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Comment

`func (o *DhcpRange6DataData) SetServer6Comment(v string)`

SetServer6Comment sets Server6Comment field to given value.

### HasServer6Comment

`func (o *DhcpRange6DataData) HasServer6Comment() bool`

HasServer6Comment returns a boolean if a field has been set.

### GetServer6Id

`func (o *DhcpRange6DataData) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DhcpRange6DataData) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DhcpRange6DataData) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DhcpRange6DataData) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DhcpRange6DataData) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DhcpRange6DataData) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DhcpRange6DataData) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DhcpRange6DataData) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DhcpRange6DataData) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DhcpRange6DataData) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DhcpRange6DataData) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DhcpRange6DataData) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetServer6Version

`func (o *DhcpRange6DataData) GetServer6Version() string`

GetServer6Version returns the Server6Version field if non-nil, zero value otherwise.

### GetServer6VersionOk

`func (o *DhcpRange6DataData) GetServer6VersionOk() (*string, bool)`

GetServer6VersionOk returns a tuple with the Server6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Version

`func (o *DhcpRange6DataData) SetServer6Version(v string)`

SetServer6Version sets Server6Version field to given value.

### HasServer6Version

`func (o *DhcpRange6DataData) HasServer6Version() bool`

HasServer6Version returns a boolean if a field has been set.

### GetRange6Acl

`func (o *DhcpRange6DataData) GetRange6Acl() string`

GetRange6Acl returns the Range6Acl field if non-nil, zero value otherwise.

### GetRange6AclOk

`func (o *DhcpRange6DataData) GetRange6AclOk() (*string, bool)`

GetRange6AclOk returns a tuple with the Range6Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Acl

`func (o *DhcpRange6DataData) SetRange6Acl(v string)`

SetRange6Acl sets Range6Acl field to given value.

### HasRange6Acl

`func (o *DhcpRange6DataData) HasRange6Acl() bool`

HasRange6Acl returns a boolean if a field has been set.

### GetRange6ClassName

`func (o *DhcpRange6DataData) GetRange6ClassName() string`

GetRange6ClassName returns the Range6ClassName field if non-nil, zero value otherwise.

### GetRange6ClassNameOk

`func (o *DhcpRange6DataData) GetRange6ClassNameOk() (*string, bool)`

GetRange6ClassNameOk returns a tuple with the Range6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6ClassName

`func (o *DhcpRange6DataData) SetRange6ClassName(v string)`

SetRange6ClassName sets Range6ClassName field to given value.

### HasRange6ClassName

`func (o *DhcpRange6DataData) HasRange6ClassName() bool`

HasRange6ClassName returns a boolean if a field has been set.

### GetRange6ClassParameters

`func (o *DhcpRange6DataData) GetRange6ClassParameters() []ApiClassParameterOutputEntry`

GetRange6ClassParameters returns the Range6ClassParameters field if non-nil, zero value otherwise.

### GetRange6ClassParametersOk

`func (o *DhcpRange6DataData) GetRange6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRange6ClassParametersOk returns a tuple with the Range6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6ClassParameters

`func (o *DhcpRange6DataData) SetRange6ClassParameters(v []ApiClassParameterOutputEntry)`

SetRange6ClassParameters sets Range6ClassParameters field to given value.

### HasRange6ClassParameters

`func (o *DhcpRange6DataData) HasRange6ClassParameters() bool`

HasRange6ClassParameters returns a boolean if a field has been set.

### GetRange6EndAddress6Addr

`func (o *DhcpRange6DataData) GetRange6EndAddress6Addr() string`

GetRange6EndAddress6Addr returns the Range6EndAddress6Addr field if non-nil, zero value otherwise.

### GetRange6EndAddress6AddrOk

`func (o *DhcpRange6DataData) GetRange6EndAddress6AddrOk() (*string, bool)`

GetRange6EndAddress6AddrOk returns a tuple with the Range6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6EndAddress6Addr

`func (o *DhcpRange6DataData) SetRange6EndAddress6Addr(v string)`

SetRange6EndAddress6Addr sets Range6EndAddress6Addr field to given value.

### HasRange6EndAddress6Addr

`func (o *DhcpRange6DataData) HasRange6EndAddress6Addr() bool`

HasRange6EndAddress6Addr returns a boolean if a field has been set.

### GetRange6FailoverName

`func (o *DhcpRange6DataData) GetRange6FailoverName() string`

GetRange6FailoverName returns the Range6FailoverName field if non-nil, zero value otherwise.

### GetRange6FailoverNameOk

`func (o *DhcpRange6DataData) GetRange6FailoverNameOk() (*string, bool)`

GetRange6FailoverNameOk returns a tuple with the Range6FailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6FailoverName

`func (o *DhcpRange6DataData) SetRange6FailoverName(v string)`

SetRange6FailoverName sets Range6FailoverName field to given value.

### HasRange6FailoverName

`func (o *DhcpRange6DataData) HasRange6FailoverName() bool`

HasRange6FailoverName returns a boolean if a field has been set.

### GetRange6Id

`func (o *DhcpRange6DataData) GetRange6Id() string`

GetRange6Id returns the Range6Id field if non-nil, zero value otherwise.

### GetRange6IdOk

`func (o *DhcpRange6DataData) GetRange6IdOk() (*string, bool)`

GetRange6IdOk returns a tuple with the Range6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Id

`func (o *DhcpRange6DataData) SetRange6Id(v string)`

SetRange6Id sets Range6Id field to given value.

### HasRange6Id

`func (o *DhcpRange6DataData) HasRange6Id() bool`

HasRange6Id returns a boolean if a field has been set.

### GetRange6LeaseCount

`func (o *DhcpRange6DataData) GetRange6LeaseCount() string`

GetRange6LeaseCount returns the Range6LeaseCount field if non-nil, zero value otherwise.

### GetRange6LeaseCountOk

`func (o *DhcpRange6DataData) GetRange6LeaseCountOk() (*string, bool)`

GetRange6LeaseCountOk returns a tuple with the Range6LeaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6LeaseCount

`func (o *DhcpRange6DataData) SetRange6LeaseCount(v string)`

SetRange6LeaseCount sets Range6LeaseCount field to given value.

### HasRange6LeaseCount

`func (o *DhcpRange6DataData) HasRange6LeaseCount() bool`

HasRange6LeaseCount returns a boolean if a field has been set.

### GetRange6Size

`func (o *DhcpRange6DataData) GetRange6Size() string`

GetRange6Size returns the Range6Size field if non-nil, zero value otherwise.

### GetRange6SizeOk

`func (o *DhcpRange6DataData) GetRange6SizeOk() (*string, bool)`

GetRange6SizeOk returns a tuple with the Range6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Size

`func (o *DhcpRange6DataData) SetRange6Size(v string)`

SetRange6Size sets Range6Size field to given value.

### HasRange6Size

`func (o *DhcpRange6DataData) HasRange6Size() bool`

HasRange6Size returns a boolean if a field has been set.

### GetRange6StartAddress6Addr

`func (o *DhcpRange6DataData) GetRange6StartAddress6Addr() string`

GetRange6StartAddress6Addr returns the Range6StartAddress6Addr field if non-nil, zero value otherwise.

### GetRange6StartAddress6AddrOk

`func (o *DhcpRange6DataData) GetRange6StartAddress6AddrOk() (*string, bool)`

GetRange6StartAddress6AddrOk returns a tuple with the Range6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6StartAddress6Addr

`func (o *DhcpRange6DataData) SetRange6StartAddress6Addr(v string)`

SetRange6StartAddress6Addr sets Range6StartAddress6Addr field to given value.

### HasRange6StartAddress6Addr

`func (o *DhcpRange6DataData) HasRange6StartAddress6Addr() bool`

HasRange6StartAddress6Addr returns a boolean if a field has been set.

### GetRange6State

`func (o *DhcpRange6DataData) GetRange6State() string`

GetRange6State returns the Range6State field if non-nil, zero value otherwise.

### GetRange6StateOk

`func (o *DhcpRange6DataData) GetRange6StateOk() (*string, bool)`

GetRange6StateOk returns a tuple with the Range6State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6State

`func (o *DhcpRange6DataData) SetRange6State(v string)`

SetRange6State sets Range6State field to given value.

### HasRange6State

`func (o *DhcpRange6DataData) HasRange6State() bool`

HasRange6State returns a boolean if a field has been set.

### GetScope6ClassName

`func (o *DhcpRange6DataData) GetScope6ClassName() string`

GetScope6ClassName returns the Scope6ClassName field if non-nil, zero value otherwise.

### GetScope6ClassNameOk

`func (o *DhcpRange6DataData) GetScope6ClassNameOk() (*string, bool)`

GetScope6ClassNameOk returns a tuple with the Scope6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassName

`func (o *DhcpRange6DataData) SetScope6ClassName(v string)`

SetScope6ClassName sets Scope6ClassName field to given value.

### HasScope6ClassName

`func (o *DhcpRange6DataData) HasScope6ClassName() bool`

HasScope6ClassName returns a boolean if a field has been set.

### GetScope6ClassParameters

`func (o *DhcpRange6DataData) GetScope6ClassParameters() []ApiClassParameterOutputEntry`

GetScope6ClassParameters returns the Scope6ClassParameters field if non-nil, zero value otherwise.

### GetScope6ClassParametersOk

`func (o *DhcpRange6DataData) GetScope6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScope6ClassParametersOk returns a tuple with the Scope6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassParameters

`func (o *DhcpRange6DataData) SetScope6ClassParameters(v []ApiClassParameterOutputEntry)`

SetScope6ClassParameters sets Scope6ClassParameters field to given value.

### HasScope6ClassParameters

`func (o *DhcpRange6DataData) HasScope6ClassParameters() bool`

HasScope6ClassParameters returns a boolean if a field has been set.

### GetScope6EndAddress6Addr

`func (o *DhcpRange6DataData) GetScope6EndAddress6Addr() string`

GetScope6EndAddress6Addr returns the Scope6EndAddress6Addr field if non-nil, zero value otherwise.

### GetScope6EndAddress6AddrOk

`func (o *DhcpRange6DataData) GetScope6EndAddress6AddrOk() (*string, bool)`

GetScope6EndAddress6AddrOk returns a tuple with the Scope6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6EndAddress6Addr

`func (o *DhcpRange6DataData) SetScope6EndAddress6Addr(v string)`

SetScope6EndAddress6Addr sets Scope6EndAddress6Addr field to given value.

### HasScope6EndAddress6Addr

`func (o *DhcpRange6DataData) HasScope6EndAddress6Addr() bool`

HasScope6EndAddress6Addr returns a boolean if a field has been set.

### GetScope6Id

`func (o *DhcpRange6DataData) GetScope6Id() string`

GetScope6Id returns the Scope6Id field if non-nil, zero value otherwise.

### GetScope6IdOk

`func (o *DhcpRange6DataData) GetScope6IdOk() (*string, bool)`

GetScope6IdOk returns a tuple with the Scope6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Id

`func (o *DhcpRange6DataData) SetScope6Id(v string)`

SetScope6Id sets Scope6Id field to given value.

### HasScope6Id

`func (o *DhcpRange6DataData) HasScope6Id() bool`

HasScope6Id returns a boolean if a field has been set.

### GetScope6Name

`func (o *DhcpRange6DataData) GetScope6Name() string`

GetScope6Name returns the Scope6Name field if non-nil, zero value otherwise.

### GetScope6NameOk

`func (o *DhcpRange6DataData) GetScope6NameOk() (*string, bool)`

GetScope6NameOk returns a tuple with the Scope6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Name

`func (o *DhcpRange6DataData) SetScope6Name(v string)`

SetScope6Name sets Scope6Name field to given value.

### HasScope6Name

`func (o *DhcpRange6DataData) HasScope6Name() bool`

HasScope6Name returns a boolean if a field has been set.

### GetScope6Prefix

`func (o *DhcpRange6DataData) GetScope6Prefix() string`

GetScope6Prefix returns the Scope6Prefix field if non-nil, zero value otherwise.

### GetScope6PrefixOk

`func (o *DhcpRange6DataData) GetScope6PrefixOk() (*string, bool)`

GetScope6PrefixOk returns a tuple with the Scope6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Prefix

`func (o *DhcpRange6DataData) SetScope6Prefix(v string)`

SetScope6Prefix sets Scope6Prefix field to given value.

### HasScope6Prefix

`func (o *DhcpRange6DataData) HasScope6Prefix() bool`

HasScope6Prefix returns a boolean if a field has been set.

### GetScope6SpaceId

`func (o *DhcpRange6DataData) GetScope6SpaceId() string`

GetScope6SpaceId returns the Scope6SpaceId field if non-nil, zero value otherwise.

### GetScope6SpaceIdOk

`func (o *DhcpRange6DataData) GetScope6SpaceIdOk() (*string, bool)`

GetScope6SpaceIdOk returns a tuple with the Scope6SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6SpaceId

`func (o *DhcpRange6DataData) SetScope6SpaceId(v string)`

SetScope6SpaceId sets Scope6SpaceId field to given value.

### HasScope6SpaceId

`func (o *DhcpRange6DataData) HasScope6SpaceId() bool`

HasScope6SpaceId returns a boolean if a field has been set.

### GetScope6Size

`func (o *DhcpRange6DataData) GetScope6Size() string`

GetScope6Size returns the Scope6Size field if non-nil, zero value otherwise.

### GetScope6SizeOk

`func (o *DhcpRange6DataData) GetScope6SizeOk() (*string, bool)`

GetScope6SizeOk returns a tuple with the Scope6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Size

`func (o *DhcpRange6DataData) SetScope6Size(v string)`

SetScope6Size sets Scope6Size field to given value.

### HasScope6Size

`func (o *DhcpRange6DataData) HasScope6Size() bool`

HasScope6Size returns a boolean if a field has been set.

### GetScope6StartAddress6Addr

`func (o *DhcpRange6DataData) GetScope6StartAddress6Addr() string`

GetScope6StartAddress6Addr returns the Scope6StartAddress6Addr field if non-nil, zero value otherwise.

### GetScope6StartAddress6AddrOk

`func (o *DhcpRange6DataData) GetScope6StartAddress6AddrOk() (*string, bool)`

GetScope6StartAddress6AddrOk returns a tuple with the Scope6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6StartAddress6Addr

`func (o *DhcpRange6DataData) SetScope6StartAddress6Addr(v string)`

SetScope6StartAddress6Addr sets Scope6StartAddress6Addr field to given value.

### HasScope6StartAddress6Addr

`func (o *DhcpRange6DataData) HasScope6StartAddress6Addr() bool`

HasScope6StartAddress6Addr returns a boolean if a field has been set.

### GetSharednetwork6Id

`func (o *DhcpRange6DataData) GetSharednetwork6Id() string`

GetSharednetwork6Id returns the Sharednetwork6Id field if non-nil, zero value otherwise.

### GetSharednetwork6IdOk

`func (o *DhcpRange6DataData) GetSharednetwork6IdOk() (*string, bool)`

GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Id

`func (o *DhcpRange6DataData) SetSharednetwork6Id(v string)`

SetSharednetwork6Id sets Sharednetwork6Id field to given value.

### HasSharednetwork6Id

`func (o *DhcpRange6DataData) HasSharednetwork6Id() bool`

HasSharednetwork6Id returns a boolean if a field has been set.

### GetSharednetwork6Name

`func (o *DhcpRange6DataData) GetSharednetwork6Name() string`

GetSharednetwork6Name returns the Sharednetwork6Name field if non-nil, zero value otherwise.

### GetSharednetwork6NameOk

`func (o *DhcpRange6DataData) GetSharednetwork6NameOk() (*string, bool)`

GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Name

`func (o *DhcpRange6DataData) SetSharednetwork6Name(v string)`

SetSharednetwork6Name sets Sharednetwork6Name field to given value.

### HasSharednetwork6Name

`func (o *DhcpRange6DataData) HasSharednetwork6Name() bool`

HasSharednetwork6Name returns a boolean if a field has been set.

### GetServer6Addr

`func (o *DhcpRange6DataData) GetServer6Addr() string`

GetServer6Addr returns the Server6Addr field if non-nil, zero value otherwise.

### GetServer6AddrOk

`func (o *DhcpRange6DataData) GetServer6AddrOk() (*string, bool)`

GetServer6AddrOk returns a tuple with the Server6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr

`func (o *DhcpRange6DataData) SetServer6Addr(v string)`

SetServer6Addr sets Server6Addr field to given value.

### HasServer6Addr

`func (o *DhcpRange6DataData) HasServer6Addr() bool`

HasServer6Addr returns a boolean if a field has been set.

### GetRange6Multistatus

`func (o *DhcpRange6DataData) GetRange6Multistatus() string`

GetRange6Multistatus returns the Range6Multistatus field if non-nil, zero value otherwise.

### GetRange6MultistatusOk

`func (o *DhcpRange6DataData) GetRange6MultistatusOk() (*string, bool)`

GetRange6MultistatusOk returns a tuple with the Range6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Multistatus

`func (o *DhcpRange6DataData) SetRange6Multistatus(v string)`

SetRange6Multistatus sets Range6Multistatus field to given value.

### HasRange6Multistatus

`func (o *DhcpRange6DataData) HasRange6Multistatus() bool`

HasRange6Multistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpRange6DataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpRange6DataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpRange6DataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpRange6DataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DhcpRange6DataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DhcpRange6DataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DhcpRange6DataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DhcpRange6DataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


