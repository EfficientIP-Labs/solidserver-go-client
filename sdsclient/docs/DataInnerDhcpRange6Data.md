# DataInnerDhcpRange6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Range6DelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**Server6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 server the object belongs to, it can be preceded by the class directory. | [optional] 
**Server6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 server the object belongs to. | [optional] 
**Server6Comment** | Pointer to **string** | The description of the DHCPv6 server the object belongs to. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;server6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv6 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Version** | Pointer to **string** | The version details of the DHCPv6 server the object belongs to. | [optional] 
**Range6Acl** | Pointer to **string** | The list of ACLs associated with the DHCPv6 range, as follows: &lt;b&gt;&amp;lt;ACL_name&amp;gt;;&amp;lt;ACL_name&amp;gt;;...&lt;/b&gt; . | [optional] 
**Range6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 range, it can be preceded by the class directory. | [optional] 
**Range6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 range. | [optional] 
**Range6EndAddress6Addr** | Pointer to **string** | The last IP address of the DHCPv6 range, in hexadecimal format. | [optional] 
**Range6FailoverName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Range6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 range. | [optional] 
**Range6LeaseCount** | Pointer to **string** | The total number of leases currently delivered by the DHCPv6 range. | [optional] 
**Range6Name** | Pointer to **string** | The start and end IP address of the DHCPv6 range, &lt;b&gt;range6_start_address6_addr&lt;/b&gt; and &lt;b&gt;range6_end_address6_addr&lt;/b&gt;, in compressed format as follows: &lt;b&gt;&amp;lt;start-ip&amp;gt; - &amp;lt;end-ip&amp;gt;&lt;/b&gt;. | [optional] 
**Range6Size** | Pointer to **string** | The number of IP addresses the DHCPv6 range contains. | [optional] 
**Range6StartAddress6Addr** | Pointer to **string** | The first IP address of the DHCPv6 range, in hexadecimal format. | [optional] 
**Range6State** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Scope6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**Scope6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 scope the object belongs to. | [optional] 
**Scope6EndAddress6Addr** | Pointer to **string** | The last IP address of the DHCPv6 scope the object belongs to, in hexadecimal format. | [optional] 
**Scope6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Name** | Pointer to **string** | The name of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Prefix** | Pointer to **string** | The prefix of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6SpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Size** | Pointer to **string** | The number of IP addresses the DHCPv6 scope the object belongs to contains. | [optional] 
**Scope6StartAddress6Addr** | Pointer to **string** | The first IP address of the DHCPv6 scope the object belongs to, in hexadecimal format. | [optional] 
**Sharednetwork6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 shared network the object belongs to. | [optional] 
**Sharednetwork6Name** | Pointer to **string** | The name of the DHCPv6 shared network the object belongs to. | [optional] 
**Server6Hostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server6_addr&lt;/b&gt; or &lt;b&gt;server6_addr6&lt;/b&gt;. | [optional] 
**Server6Addr6** | Pointer to **string** | The Management IP address of the DHCPv6 server the object belongs to, the IPv6 address configured when adding the server, in hexadecimal format. | [optional] 
**Server6Addr** | Pointer to **string** | The Management IP address of the DHCPv6 server the object belongs to, the IPv4 address configured when adding the server, in hexadecimal format. | [optional] 
**Range6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpRange6Data

`func NewDataInnerDhcpRange6Data() *DataInnerDhcpRange6Data`

NewDataInnerDhcpRange6Data instantiates a new DataInnerDhcpRange6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpRange6DataWithDefaults

`func NewDataInnerDhcpRange6DataWithDefaults() *DataInnerDhcpRange6Data`

NewDataInnerDhcpRange6DataWithDefaults instantiates a new DataInnerDhcpRange6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRange6DelayedTime

`func (o *DataInnerDhcpRange6Data) GetRange6DelayedTime() string`

GetRange6DelayedTime returns the Range6DelayedTime field if non-nil, zero value otherwise.

### GetRange6DelayedTimeOk

`func (o *DataInnerDhcpRange6Data) GetRange6DelayedTimeOk() (*string, bool)`

GetRange6DelayedTimeOk returns a tuple with the Range6DelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6DelayedTime

`func (o *DataInnerDhcpRange6Data) SetRange6DelayedTime(v string)`

SetRange6DelayedTime sets Range6DelayedTime field to given value.

### HasRange6DelayedTime

`func (o *DataInnerDhcpRange6Data) HasRange6DelayedTime() bool`

HasRange6DelayedTime returns a boolean if a field has been set.

### GetServer6ClassName

`func (o *DataInnerDhcpRange6Data) GetServer6ClassName() string`

GetServer6ClassName returns the Server6ClassName field if non-nil, zero value otherwise.

### GetServer6ClassNameOk

`func (o *DataInnerDhcpRange6Data) GetServer6ClassNameOk() (*string, bool)`

GetServer6ClassNameOk returns a tuple with the Server6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassName

`func (o *DataInnerDhcpRange6Data) SetServer6ClassName(v string)`

SetServer6ClassName sets Server6ClassName field to given value.

### HasServer6ClassName

`func (o *DataInnerDhcpRange6Data) HasServer6ClassName() bool`

HasServer6ClassName returns a boolean if a field has been set.

### GetServer6ClassParameters

`func (o *DataInnerDhcpRange6Data) GetServer6ClassParameters() []ApiClassParameterOutputEntry`

GetServer6ClassParameters returns the Server6ClassParameters field if non-nil, zero value otherwise.

### GetServer6ClassParametersOk

`func (o *DataInnerDhcpRange6Data) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassParameters

`func (o *DataInnerDhcpRange6Data) SetServer6ClassParameters(v []ApiClassParameterOutputEntry)`

SetServer6ClassParameters sets Server6ClassParameters field to given value.

### HasServer6ClassParameters

`func (o *DataInnerDhcpRange6Data) HasServer6ClassParameters() bool`

HasServer6ClassParameters returns a boolean if a field has been set.

### GetServer6Comment

`func (o *DataInnerDhcpRange6Data) GetServer6Comment() string`

GetServer6Comment returns the Server6Comment field if non-nil, zero value otherwise.

### GetServer6CommentOk

`func (o *DataInnerDhcpRange6Data) GetServer6CommentOk() (*string, bool)`

GetServer6CommentOk returns a tuple with the Server6Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Comment

`func (o *DataInnerDhcpRange6Data) SetServer6Comment(v string)`

SetServer6Comment sets Server6Comment field to given value.

### HasServer6Comment

`func (o *DataInnerDhcpRange6Data) HasServer6Comment() bool`

HasServer6Comment returns a boolean if a field has been set.

### GetServer6Id

`func (o *DataInnerDhcpRange6Data) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DataInnerDhcpRange6Data) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DataInnerDhcpRange6Data) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DataInnerDhcpRange6Data) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DataInnerDhcpRange6Data) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DataInnerDhcpRange6Data) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DataInnerDhcpRange6Data) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DataInnerDhcpRange6Data) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DataInnerDhcpRange6Data) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DataInnerDhcpRange6Data) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DataInnerDhcpRange6Data) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DataInnerDhcpRange6Data) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetServer6Version

`func (o *DataInnerDhcpRange6Data) GetServer6Version() string`

GetServer6Version returns the Server6Version field if non-nil, zero value otherwise.

### GetServer6VersionOk

`func (o *DataInnerDhcpRange6Data) GetServer6VersionOk() (*string, bool)`

GetServer6VersionOk returns a tuple with the Server6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Version

`func (o *DataInnerDhcpRange6Data) SetServer6Version(v string)`

SetServer6Version sets Server6Version field to given value.

### HasServer6Version

`func (o *DataInnerDhcpRange6Data) HasServer6Version() bool`

HasServer6Version returns a boolean if a field has been set.

### GetRange6Acl

`func (o *DataInnerDhcpRange6Data) GetRange6Acl() string`

GetRange6Acl returns the Range6Acl field if non-nil, zero value otherwise.

### GetRange6AclOk

`func (o *DataInnerDhcpRange6Data) GetRange6AclOk() (*string, bool)`

GetRange6AclOk returns a tuple with the Range6Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Acl

`func (o *DataInnerDhcpRange6Data) SetRange6Acl(v string)`

SetRange6Acl sets Range6Acl field to given value.

### HasRange6Acl

`func (o *DataInnerDhcpRange6Data) HasRange6Acl() bool`

HasRange6Acl returns a boolean if a field has been set.

### GetRange6ClassName

`func (o *DataInnerDhcpRange6Data) GetRange6ClassName() string`

GetRange6ClassName returns the Range6ClassName field if non-nil, zero value otherwise.

### GetRange6ClassNameOk

`func (o *DataInnerDhcpRange6Data) GetRange6ClassNameOk() (*string, bool)`

GetRange6ClassNameOk returns a tuple with the Range6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6ClassName

`func (o *DataInnerDhcpRange6Data) SetRange6ClassName(v string)`

SetRange6ClassName sets Range6ClassName field to given value.

### HasRange6ClassName

`func (o *DataInnerDhcpRange6Data) HasRange6ClassName() bool`

HasRange6ClassName returns a boolean if a field has been set.

### GetRange6ClassParameters

`func (o *DataInnerDhcpRange6Data) GetRange6ClassParameters() []ApiClassParameterOutputEntry`

GetRange6ClassParameters returns the Range6ClassParameters field if non-nil, zero value otherwise.

### GetRange6ClassParametersOk

`func (o *DataInnerDhcpRange6Data) GetRange6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRange6ClassParametersOk returns a tuple with the Range6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6ClassParameters

`func (o *DataInnerDhcpRange6Data) SetRange6ClassParameters(v []ApiClassParameterOutputEntry)`

SetRange6ClassParameters sets Range6ClassParameters field to given value.

### HasRange6ClassParameters

`func (o *DataInnerDhcpRange6Data) HasRange6ClassParameters() bool`

HasRange6ClassParameters returns a boolean if a field has been set.

### GetRange6EndAddress6Addr

`func (o *DataInnerDhcpRange6Data) GetRange6EndAddress6Addr() string`

GetRange6EndAddress6Addr returns the Range6EndAddress6Addr field if non-nil, zero value otherwise.

### GetRange6EndAddress6AddrOk

`func (o *DataInnerDhcpRange6Data) GetRange6EndAddress6AddrOk() (*string, bool)`

GetRange6EndAddress6AddrOk returns a tuple with the Range6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6EndAddress6Addr

`func (o *DataInnerDhcpRange6Data) SetRange6EndAddress6Addr(v string)`

SetRange6EndAddress6Addr sets Range6EndAddress6Addr field to given value.

### HasRange6EndAddress6Addr

`func (o *DataInnerDhcpRange6Data) HasRange6EndAddress6Addr() bool`

HasRange6EndAddress6Addr returns a boolean if a field has been set.

### GetRange6FailoverName

`func (o *DataInnerDhcpRange6Data) GetRange6FailoverName() string`

GetRange6FailoverName returns the Range6FailoverName field if non-nil, zero value otherwise.

### GetRange6FailoverNameOk

`func (o *DataInnerDhcpRange6Data) GetRange6FailoverNameOk() (*string, bool)`

GetRange6FailoverNameOk returns a tuple with the Range6FailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6FailoverName

`func (o *DataInnerDhcpRange6Data) SetRange6FailoverName(v string)`

SetRange6FailoverName sets Range6FailoverName field to given value.

### HasRange6FailoverName

`func (o *DataInnerDhcpRange6Data) HasRange6FailoverName() bool`

HasRange6FailoverName returns a boolean if a field has been set.

### GetRange6Id

`func (o *DataInnerDhcpRange6Data) GetRange6Id() string`

GetRange6Id returns the Range6Id field if non-nil, zero value otherwise.

### GetRange6IdOk

`func (o *DataInnerDhcpRange6Data) GetRange6IdOk() (*string, bool)`

GetRange6IdOk returns a tuple with the Range6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Id

`func (o *DataInnerDhcpRange6Data) SetRange6Id(v string)`

SetRange6Id sets Range6Id field to given value.

### HasRange6Id

`func (o *DataInnerDhcpRange6Data) HasRange6Id() bool`

HasRange6Id returns a boolean if a field has been set.

### GetRange6LeaseCount

`func (o *DataInnerDhcpRange6Data) GetRange6LeaseCount() string`

GetRange6LeaseCount returns the Range6LeaseCount field if non-nil, zero value otherwise.

### GetRange6LeaseCountOk

`func (o *DataInnerDhcpRange6Data) GetRange6LeaseCountOk() (*string, bool)`

GetRange6LeaseCountOk returns a tuple with the Range6LeaseCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6LeaseCount

`func (o *DataInnerDhcpRange6Data) SetRange6LeaseCount(v string)`

SetRange6LeaseCount sets Range6LeaseCount field to given value.

### HasRange6LeaseCount

`func (o *DataInnerDhcpRange6Data) HasRange6LeaseCount() bool`

HasRange6LeaseCount returns a boolean if a field has been set.

### GetRange6Name

`func (o *DataInnerDhcpRange6Data) GetRange6Name() string`

GetRange6Name returns the Range6Name field if non-nil, zero value otherwise.

### GetRange6NameOk

`func (o *DataInnerDhcpRange6Data) GetRange6NameOk() (*string, bool)`

GetRange6NameOk returns a tuple with the Range6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Name

`func (o *DataInnerDhcpRange6Data) SetRange6Name(v string)`

SetRange6Name sets Range6Name field to given value.

### HasRange6Name

`func (o *DataInnerDhcpRange6Data) HasRange6Name() bool`

HasRange6Name returns a boolean if a field has been set.

### GetRange6Size

`func (o *DataInnerDhcpRange6Data) GetRange6Size() string`

GetRange6Size returns the Range6Size field if non-nil, zero value otherwise.

### GetRange6SizeOk

`func (o *DataInnerDhcpRange6Data) GetRange6SizeOk() (*string, bool)`

GetRange6SizeOk returns a tuple with the Range6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Size

`func (o *DataInnerDhcpRange6Data) SetRange6Size(v string)`

SetRange6Size sets Range6Size field to given value.

### HasRange6Size

`func (o *DataInnerDhcpRange6Data) HasRange6Size() bool`

HasRange6Size returns a boolean if a field has been set.

### GetRange6StartAddress6Addr

`func (o *DataInnerDhcpRange6Data) GetRange6StartAddress6Addr() string`

GetRange6StartAddress6Addr returns the Range6StartAddress6Addr field if non-nil, zero value otherwise.

### GetRange6StartAddress6AddrOk

`func (o *DataInnerDhcpRange6Data) GetRange6StartAddress6AddrOk() (*string, bool)`

GetRange6StartAddress6AddrOk returns a tuple with the Range6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6StartAddress6Addr

`func (o *DataInnerDhcpRange6Data) SetRange6StartAddress6Addr(v string)`

SetRange6StartAddress6Addr sets Range6StartAddress6Addr field to given value.

### HasRange6StartAddress6Addr

`func (o *DataInnerDhcpRange6Data) HasRange6StartAddress6Addr() bool`

HasRange6StartAddress6Addr returns a boolean if a field has been set.

### GetRange6State

`func (o *DataInnerDhcpRange6Data) GetRange6State() string`

GetRange6State returns the Range6State field if non-nil, zero value otherwise.

### GetRange6StateOk

`func (o *DataInnerDhcpRange6Data) GetRange6StateOk() (*string, bool)`

GetRange6StateOk returns a tuple with the Range6State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6State

`func (o *DataInnerDhcpRange6Data) SetRange6State(v string)`

SetRange6State sets Range6State field to given value.

### HasRange6State

`func (o *DataInnerDhcpRange6Data) HasRange6State() bool`

HasRange6State returns a boolean if a field has been set.

### GetScope6ClassName

`func (o *DataInnerDhcpRange6Data) GetScope6ClassName() string`

GetScope6ClassName returns the Scope6ClassName field if non-nil, zero value otherwise.

### GetScope6ClassNameOk

`func (o *DataInnerDhcpRange6Data) GetScope6ClassNameOk() (*string, bool)`

GetScope6ClassNameOk returns a tuple with the Scope6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassName

`func (o *DataInnerDhcpRange6Data) SetScope6ClassName(v string)`

SetScope6ClassName sets Scope6ClassName field to given value.

### HasScope6ClassName

`func (o *DataInnerDhcpRange6Data) HasScope6ClassName() bool`

HasScope6ClassName returns a boolean if a field has been set.

### GetScope6ClassParameters

`func (o *DataInnerDhcpRange6Data) GetScope6ClassParameters() []ApiClassParameterOutputEntry`

GetScope6ClassParameters returns the Scope6ClassParameters field if non-nil, zero value otherwise.

### GetScope6ClassParametersOk

`func (o *DataInnerDhcpRange6Data) GetScope6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetScope6ClassParametersOk returns a tuple with the Scope6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassParameters

`func (o *DataInnerDhcpRange6Data) SetScope6ClassParameters(v []ApiClassParameterOutputEntry)`

SetScope6ClassParameters sets Scope6ClassParameters field to given value.

### HasScope6ClassParameters

`func (o *DataInnerDhcpRange6Data) HasScope6ClassParameters() bool`

HasScope6ClassParameters returns a boolean if a field has been set.

### GetScope6EndAddress6Addr

`func (o *DataInnerDhcpRange6Data) GetScope6EndAddress6Addr() string`

GetScope6EndAddress6Addr returns the Scope6EndAddress6Addr field if non-nil, zero value otherwise.

### GetScope6EndAddress6AddrOk

`func (o *DataInnerDhcpRange6Data) GetScope6EndAddress6AddrOk() (*string, bool)`

GetScope6EndAddress6AddrOk returns a tuple with the Scope6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6EndAddress6Addr

`func (o *DataInnerDhcpRange6Data) SetScope6EndAddress6Addr(v string)`

SetScope6EndAddress6Addr sets Scope6EndAddress6Addr field to given value.

### HasScope6EndAddress6Addr

`func (o *DataInnerDhcpRange6Data) HasScope6EndAddress6Addr() bool`

HasScope6EndAddress6Addr returns a boolean if a field has been set.

### GetScope6Id

`func (o *DataInnerDhcpRange6Data) GetScope6Id() string`

GetScope6Id returns the Scope6Id field if non-nil, zero value otherwise.

### GetScope6IdOk

`func (o *DataInnerDhcpRange6Data) GetScope6IdOk() (*string, bool)`

GetScope6IdOk returns a tuple with the Scope6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Id

`func (o *DataInnerDhcpRange6Data) SetScope6Id(v string)`

SetScope6Id sets Scope6Id field to given value.

### HasScope6Id

`func (o *DataInnerDhcpRange6Data) HasScope6Id() bool`

HasScope6Id returns a boolean if a field has been set.

### GetScope6Name

`func (o *DataInnerDhcpRange6Data) GetScope6Name() string`

GetScope6Name returns the Scope6Name field if non-nil, zero value otherwise.

### GetScope6NameOk

`func (o *DataInnerDhcpRange6Data) GetScope6NameOk() (*string, bool)`

GetScope6NameOk returns a tuple with the Scope6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Name

`func (o *DataInnerDhcpRange6Data) SetScope6Name(v string)`

SetScope6Name sets Scope6Name field to given value.

### HasScope6Name

`func (o *DataInnerDhcpRange6Data) HasScope6Name() bool`

HasScope6Name returns a boolean if a field has been set.

### GetScope6Prefix

`func (o *DataInnerDhcpRange6Data) GetScope6Prefix() string`

GetScope6Prefix returns the Scope6Prefix field if non-nil, zero value otherwise.

### GetScope6PrefixOk

`func (o *DataInnerDhcpRange6Data) GetScope6PrefixOk() (*string, bool)`

GetScope6PrefixOk returns a tuple with the Scope6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Prefix

`func (o *DataInnerDhcpRange6Data) SetScope6Prefix(v string)`

SetScope6Prefix sets Scope6Prefix field to given value.

### HasScope6Prefix

`func (o *DataInnerDhcpRange6Data) HasScope6Prefix() bool`

HasScope6Prefix returns a boolean if a field has been set.

### GetScope6SpaceId

`func (o *DataInnerDhcpRange6Data) GetScope6SpaceId() string`

GetScope6SpaceId returns the Scope6SpaceId field if non-nil, zero value otherwise.

### GetScope6SpaceIdOk

`func (o *DataInnerDhcpRange6Data) GetScope6SpaceIdOk() (*string, bool)`

GetScope6SpaceIdOk returns a tuple with the Scope6SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6SpaceId

`func (o *DataInnerDhcpRange6Data) SetScope6SpaceId(v string)`

SetScope6SpaceId sets Scope6SpaceId field to given value.

### HasScope6SpaceId

`func (o *DataInnerDhcpRange6Data) HasScope6SpaceId() bool`

HasScope6SpaceId returns a boolean if a field has been set.

### GetScope6Size

`func (o *DataInnerDhcpRange6Data) GetScope6Size() string`

GetScope6Size returns the Scope6Size field if non-nil, zero value otherwise.

### GetScope6SizeOk

`func (o *DataInnerDhcpRange6Data) GetScope6SizeOk() (*string, bool)`

GetScope6SizeOk returns a tuple with the Scope6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Size

`func (o *DataInnerDhcpRange6Data) SetScope6Size(v string)`

SetScope6Size sets Scope6Size field to given value.

### HasScope6Size

`func (o *DataInnerDhcpRange6Data) HasScope6Size() bool`

HasScope6Size returns a boolean if a field has been set.

### GetScope6StartAddress6Addr

`func (o *DataInnerDhcpRange6Data) GetScope6StartAddress6Addr() string`

GetScope6StartAddress6Addr returns the Scope6StartAddress6Addr field if non-nil, zero value otherwise.

### GetScope6StartAddress6AddrOk

`func (o *DataInnerDhcpRange6Data) GetScope6StartAddress6AddrOk() (*string, bool)`

GetScope6StartAddress6AddrOk returns a tuple with the Scope6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6StartAddress6Addr

`func (o *DataInnerDhcpRange6Data) SetScope6StartAddress6Addr(v string)`

SetScope6StartAddress6Addr sets Scope6StartAddress6Addr field to given value.

### HasScope6StartAddress6Addr

`func (o *DataInnerDhcpRange6Data) HasScope6StartAddress6Addr() bool`

HasScope6StartAddress6Addr returns a boolean if a field has been set.

### GetSharednetwork6Id

`func (o *DataInnerDhcpRange6Data) GetSharednetwork6Id() string`

GetSharednetwork6Id returns the Sharednetwork6Id field if non-nil, zero value otherwise.

### GetSharednetwork6IdOk

`func (o *DataInnerDhcpRange6Data) GetSharednetwork6IdOk() (*string, bool)`

GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Id

`func (o *DataInnerDhcpRange6Data) SetSharednetwork6Id(v string)`

SetSharednetwork6Id sets Sharednetwork6Id field to given value.

### HasSharednetwork6Id

`func (o *DataInnerDhcpRange6Data) HasSharednetwork6Id() bool`

HasSharednetwork6Id returns a boolean if a field has been set.

### GetSharednetwork6Name

`func (o *DataInnerDhcpRange6Data) GetSharednetwork6Name() string`

GetSharednetwork6Name returns the Sharednetwork6Name field if non-nil, zero value otherwise.

### GetSharednetwork6NameOk

`func (o *DataInnerDhcpRange6Data) GetSharednetwork6NameOk() (*string, bool)`

GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Name

`func (o *DataInnerDhcpRange6Data) SetSharednetwork6Name(v string)`

SetSharednetwork6Name sets Sharednetwork6Name field to given value.

### HasSharednetwork6Name

`func (o *DataInnerDhcpRange6Data) HasSharednetwork6Name() bool`

HasSharednetwork6Name returns a boolean if a field has been set.

### GetServer6Hostaddr

`func (o *DataInnerDhcpRange6Data) GetServer6Hostaddr() string`

GetServer6Hostaddr returns the Server6Hostaddr field if non-nil, zero value otherwise.

### GetServer6HostaddrOk

`func (o *DataInnerDhcpRange6Data) GetServer6HostaddrOk() (*string, bool)`

GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Hostaddr

`func (o *DataInnerDhcpRange6Data) SetServer6Hostaddr(v string)`

SetServer6Hostaddr sets Server6Hostaddr field to given value.

### HasServer6Hostaddr

`func (o *DataInnerDhcpRange6Data) HasServer6Hostaddr() bool`

HasServer6Hostaddr returns a boolean if a field has been set.

### GetServer6Addr6

`func (o *DataInnerDhcpRange6Data) GetServer6Addr6() string`

GetServer6Addr6 returns the Server6Addr6 field if non-nil, zero value otherwise.

### GetServer6Addr6Ok

`func (o *DataInnerDhcpRange6Data) GetServer6Addr6Ok() (*string, bool)`

GetServer6Addr6Ok returns a tuple with the Server6Addr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr6

`func (o *DataInnerDhcpRange6Data) SetServer6Addr6(v string)`

SetServer6Addr6 sets Server6Addr6 field to given value.

### HasServer6Addr6

`func (o *DataInnerDhcpRange6Data) HasServer6Addr6() bool`

HasServer6Addr6 returns a boolean if a field has been set.

### GetServer6Addr

`func (o *DataInnerDhcpRange6Data) GetServer6Addr() string`

GetServer6Addr returns the Server6Addr field if non-nil, zero value otherwise.

### GetServer6AddrOk

`func (o *DataInnerDhcpRange6Data) GetServer6AddrOk() (*string, bool)`

GetServer6AddrOk returns a tuple with the Server6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr

`func (o *DataInnerDhcpRange6Data) SetServer6Addr(v string)`

SetServer6Addr sets Server6Addr field to given value.

### HasServer6Addr

`func (o *DataInnerDhcpRange6Data) HasServer6Addr() bool`

HasServer6Addr returns a boolean if a field has been set.

### GetRange6Multistatus

`func (o *DataInnerDhcpRange6Data) GetRange6Multistatus() string`

GetRange6Multistatus returns the Range6Multistatus field if non-nil, zero value otherwise.

### GetRange6MultistatusOk

`func (o *DataInnerDhcpRange6Data) GetRange6MultistatusOk() (*string, bool)`

GetRange6MultistatusOk returns a tuple with the Range6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Multistatus

`func (o *DataInnerDhcpRange6Data) SetRange6Multistatus(v string)`

SetRange6Multistatus sets Range6Multistatus field to given value.

### HasRange6Multistatus

`func (o *DataInnerDhcpRange6Data) HasRange6Multistatus() bool`

HasRange6Multistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpRange6Data) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpRange6Data) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpRange6Data) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpRange6Data) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpRange6Data) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpRange6Data) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpRange6Data) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpRange6Data) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


