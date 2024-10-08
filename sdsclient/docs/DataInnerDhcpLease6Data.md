# DataInnerDhcpLease6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 server the object belongs to, it can be preceded by the class directory. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;server6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv6 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Version** | Pointer to **string** | The version details of the DHCPv6 server the object belongs to. | [optional] 
**Lease6ClientDuid** | Pointer to **string** | The client DHCP Unique Identifier (DUID) associated with the DHCPv6 lease. | [optional] 
**Lease6Clientname** | Pointer to **string** | The name of the client associated with the DHCPv6 lease. | [optional] 
**Lease6Domain** | Pointer to **string** | The domain name associated with the DHCPv6 lease. | [optional] 
**Lease6EndTime** | Pointer to **string** | The expiration time of the lease, in decimal UNIX date format. | [optional] 
**Lease6FirstTime** | Pointer to **string** | The first time the DHCPv6 lease has been attributed to the client, in decimal UNIX date format. | [optional] 
**Lease6Giaddr** | Pointer to **string** | The gateway IP address of the relay agent of the DHCPv6 lease. | [optional] 
**Lease6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 lease. | [optional] 
**Lease6Address6Addr** | Pointer to **string** | The IP address associated with the DHCPv6 lease, in hexadecimal format. | [optional] 
**Lease6MacAddr** | Pointer to **string** | The MAC address associated with the DHCPv6 lease. | [optional] 
**Lease6Name** | Pointer to **string** | The name of the DHCPv6 lease. | [optional] 
**Lease6Period** | Pointer to **string** | The duration time (time to live) of the DHCPv6 lease, in seconds. | [optional] 
**Lease6Time** | Pointer to **string** | The last time the DHCPv6 lease has been attributed to the client, in decimal UNIX date format. | [optional] 
**Range6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 range the object belongs to, it can be preceded by the class directory. | [optional] 
**Range6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 range the object belongs to. | [optional] 
**Range6EndAddress6Addr** | Pointer to **string** | The last IP address of the DHCPv6 range the object belongs to, in hexadecimal format. | [optional] 
**Range6FailoverName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Range6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 range the object belongs to. | [optional] 
**Range6StartAddress6Addr** | Pointer to **string** | The first IP address of the DHCPv6 range the object belongs to, in hexadecimal format. | [optional] 
**Scope6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**Scope6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Name** | Pointer to **string** | The name of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Prefix** | Pointer to **string** | The prefix of the DHCPv6 scope the object belongs to. | [optional] 
**Scope6Size** | Pointer to **string** | The number of IP addresses the DHCPv6 scope the object belongs to contains. | [optional] 
**Scope6StartAddress6Addr** | Pointer to **string** | The first IP address of the DHCPv6 scope the object belongs to, in hexadecimal format. | [optional] 
**Sharednetwork6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 shared network the object belongs to. | [optional] 
**Sharednetwork6Name** | Pointer to **string** | The name of the DHCPv6 shared network the object belongs to. | [optional] 
**Server6Hostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server6_addr&lt;/b&gt; or &lt;b&gt;server6_addr6&lt;/b&gt;. | [optional] 
**Server6Addr6** | Pointer to **string** | The Management IP address of the DHCPv6 server the object belongs to, the IPv6 address configured when adding the server, in hexadecimal format. | [optional] 
**Server6Addr** | Pointer to **string** | The Management IP address of the DHCPv6 server the object belongs to, the IPv4 address configured when adding the server, in hexadecimal format. | [optional] 
**Lease6MacVendor** | Pointer to **string** | The vendor details of the client associated with the DHCPv6 lease. | [optional] 
**Lease6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**Lease6Percent** | Pointer to **string** | The percentage of time the lease has really been in use. | [optional] 
**Lease6TimeToExpire** | Pointer to **string** | The time left to the lease before it expires, in seconds. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpLease6Data

`func NewDataInnerDhcpLease6Data() *DataInnerDhcpLease6Data`

NewDataInnerDhcpLease6Data instantiates a new DataInnerDhcpLease6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpLease6DataWithDefaults

`func NewDataInnerDhcpLease6DataWithDefaults() *DataInnerDhcpLease6Data`

NewDataInnerDhcpLease6DataWithDefaults instantiates a new DataInnerDhcpLease6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer6ClassName

`func (o *DataInnerDhcpLease6Data) GetServer6ClassName() string`

GetServer6ClassName returns the Server6ClassName field if non-nil, zero value otherwise.

### GetServer6ClassNameOk

`func (o *DataInnerDhcpLease6Data) GetServer6ClassNameOk() (*string, bool)`

GetServer6ClassNameOk returns a tuple with the Server6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassName

`func (o *DataInnerDhcpLease6Data) SetServer6ClassName(v string)`

SetServer6ClassName sets Server6ClassName field to given value.

### HasServer6ClassName

`func (o *DataInnerDhcpLease6Data) HasServer6ClassName() bool`

HasServer6ClassName returns a boolean if a field has been set.

### GetServer6Id

`func (o *DataInnerDhcpLease6Data) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DataInnerDhcpLease6Data) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DataInnerDhcpLease6Data) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DataInnerDhcpLease6Data) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DataInnerDhcpLease6Data) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DataInnerDhcpLease6Data) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DataInnerDhcpLease6Data) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DataInnerDhcpLease6Data) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DataInnerDhcpLease6Data) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DataInnerDhcpLease6Data) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DataInnerDhcpLease6Data) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DataInnerDhcpLease6Data) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetServer6Version

`func (o *DataInnerDhcpLease6Data) GetServer6Version() string`

GetServer6Version returns the Server6Version field if non-nil, zero value otherwise.

### GetServer6VersionOk

`func (o *DataInnerDhcpLease6Data) GetServer6VersionOk() (*string, bool)`

GetServer6VersionOk returns a tuple with the Server6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Version

`func (o *DataInnerDhcpLease6Data) SetServer6Version(v string)`

SetServer6Version sets Server6Version field to given value.

### HasServer6Version

`func (o *DataInnerDhcpLease6Data) HasServer6Version() bool`

HasServer6Version returns a boolean if a field has been set.

### GetLease6ClientDuid

`func (o *DataInnerDhcpLease6Data) GetLease6ClientDuid() string`

GetLease6ClientDuid returns the Lease6ClientDuid field if non-nil, zero value otherwise.

### GetLease6ClientDuidOk

`func (o *DataInnerDhcpLease6Data) GetLease6ClientDuidOk() (*string, bool)`

GetLease6ClientDuidOk returns a tuple with the Lease6ClientDuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6ClientDuid

`func (o *DataInnerDhcpLease6Data) SetLease6ClientDuid(v string)`

SetLease6ClientDuid sets Lease6ClientDuid field to given value.

### HasLease6ClientDuid

`func (o *DataInnerDhcpLease6Data) HasLease6ClientDuid() bool`

HasLease6ClientDuid returns a boolean if a field has been set.

### GetLease6Clientname

`func (o *DataInnerDhcpLease6Data) GetLease6Clientname() string`

GetLease6Clientname returns the Lease6Clientname field if non-nil, zero value otherwise.

### GetLease6ClientnameOk

`func (o *DataInnerDhcpLease6Data) GetLease6ClientnameOk() (*string, bool)`

GetLease6ClientnameOk returns a tuple with the Lease6Clientname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Clientname

`func (o *DataInnerDhcpLease6Data) SetLease6Clientname(v string)`

SetLease6Clientname sets Lease6Clientname field to given value.

### HasLease6Clientname

`func (o *DataInnerDhcpLease6Data) HasLease6Clientname() bool`

HasLease6Clientname returns a boolean if a field has been set.

### GetLease6Domain

`func (o *DataInnerDhcpLease6Data) GetLease6Domain() string`

GetLease6Domain returns the Lease6Domain field if non-nil, zero value otherwise.

### GetLease6DomainOk

`func (o *DataInnerDhcpLease6Data) GetLease6DomainOk() (*string, bool)`

GetLease6DomainOk returns a tuple with the Lease6Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Domain

`func (o *DataInnerDhcpLease6Data) SetLease6Domain(v string)`

SetLease6Domain sets Lease6Domain field to given value.

### HasLease6Domain

`func (o *DataInnerDhcpLease6Data) HasLease6Domain() bool`

HasLease6Domain returns a boolean if a field has been set.

### GetLease6EndTime

`func (o *DataInnerDhcpLease6Data) GetLease6EndTime() string`

GetLease6EndTime returns the Lease6EndTime field if non-nil, zero value otherwise.

### GetLease6EndTimeOk

`func (o *DataInnerDhcpLease6Data) GetLease6EndTimeOk() (*string, bool)`

GetLease6EndTimeOk returns a tuple with the Lease6EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6EndTime

`func (o *DataInnerDhcpLease6Data) SetLease6EndTime(v string)`

SetLease6EndTime sets Lease6EndTime field to given value.

### HasLease6EndTime

`func (o *DataInnerDhcpLease6Data) HasLease6EndTime() bool`

HasLease6EndTime returns a boolean if a field has been set.

### GetLease6FirstTime

`func (o *DataInnerDhcpLease6Data) GetLease6FirstTime() string`

GetLease6FirstTime returns the Lease6FirstTime field if non-nil, zero value otherwise.

### GetLease6FirstTimeOk

`func (o *DataInnerDhcpLease6Data) GetLease6FirstTimeOk() (*string, bool)`

GetLease6FirstTimeOk returns a tuple with the Lease6FirstTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6FirstTime

`func (o *DataInnerDhcpLease6Data) SetLease6FirstTime(v string)`

SetLease6FirstTime sets Lease6FirstTime field to given value.

### HasLease6FirstTime

`func (o *DataInnerDhcpLease6Data) HasLease6FirstTime() bool`

HasLease6FirstTime returns a boolean if a field has been set.

### GetLease6Giaddr

`func (o *DataInnerDhcpLease6Data) GetLease6Giaddr() string`

GetLease6Giaddr returns the Lease6Giaddr field if non-nil, zero value otherwise.

### GetLease6GiaddrOk

`func (o *DataInnerDhcpLease6Data) GetLease6GiaddrOk() (*string, bool)`

GetLease6GiaddrOk returns a tuple with the Lease6Giaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Giaddr

`func (o *DataInnerDhcpLease6Data) SetLease6Giaddr(v string)`

SetLease6Giaddr sets Lease6Giaddr field to given value.

### HasLease6Giaddr

`func (o *DataInnerDhcpLease6Data) HasLease6Giaddr() bool`

HasLease6Giaddr returns a boolean if a field has been set.

### GetLease6Id

`func (o *DataInnerDhcpLease6Data) GetLease6Id() string`

GetLease6Id returns the Lease6Id field if non-nil, zero value otherwise.

### GetLease6IdOk

`func (o *DataInnerDhcpLease6Data) GetLease6IdOk() (*string, bool)`

GetLease6IdOk returns a tuple with the Lease6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Id

`func (o *DataInnerDhcpLease6Data) SetLease6Id(v string)`

SetLease6Id sets Lease6Id field to given value.

### HasLease6Id

`func (o *DataInnerDhcpLease6Data) HasLease6Id() bool`

HasLease6Id returns a boolean if a field has been set.

### GetLease6Address6Addr

`func (o *DataInnerDhcpLease6Data) GetLease6Address6Addr() string`

GetLease6Address6Addr returns the Lease6Address6Addr field if non-nil, zero value otherwise.

### GetLease6Address6AddrOk

`func (o *DataInnerDhcpLease6Data) GetLease6Address6AddrOk() (*string, bool)`

GetLease6Address6AddrOk returns a tuple with the Lease6Address6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Address6Addr

`func (o *DataInnerDhcpLease6Data) SetLease6Address6Addr(v string)`

SetLease6Address6Addr sets Lease6Address6Addr field to given value.

### HasLease6Address6Addr

`func (o *DataInnerDhcpLease6Data) HasLease6Address6Addr() bool`

HasLease6Address6Addr returns a boolean if a field has been set.

### GetLease6MacAddr

`func (o *DataInnerDhcpLease6Data) GetLease6MacAddr() string`

GetLease6MacAddr returns the Lease6MacAddr field if non-nil, zero value otherwise.

### GetLease6MacAddrOk

`func (o *DataInnerDhcpLease6Data) GetLease6MacAddrOk() (*string, bool)`

GetLease6MacAddrOk returns a tuple with the Lease6MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6MacAddr

`func (o *DataInnerDhcpLease6Data) SetLease6MacAddr(v string)`

SetLease6MacAddr sets Lease6MacAddr field to given value.

### HasLease6MacAddr

`func (o *DataInnerDhcpLease6Data) HasLease6MacAddr() bool`

HasLease6MacAddr returns a boolean if a field has been set.

### GetLease6Name

`func (o *DataInnerDhcpLease6Data) GetLease6Name() string`

GetLease6Name returns the Lease6Name field if non-nil, zero value otherwise.

### GetLease6NameOk

`func (o *DataInnerDhcpLease6Data) GetLease6NameOk() (*string, bool)`

GetLease6NameOk returns a tuple with the Lease6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Name

`func (o *DataInnerDhcpLease6Data) SetLease6Name(v string)`

SetLease6Name sets Lease6Name field to given value.

### HasLease6Name

`func (o *DataInnerDhcpLease6Data) HasLease6Name() bool`

HasLease6Name returns a boolean if a field has been set.

### GetLease6Period

`func (o *DataInnerDhcpLease6Data) GetLease6Period() string`

GetLease6Period returns the Lease6Period field if non-nil, zero value otherwise.

### GetLease6PeriodOk

`func (o *DataInnerDhcpLease6Data) GetLease6PeriodOk() (*string, bool)`

GetLease6PeriodOk returns a tuple with the Lease6Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Period

`func (o *DataInnerDhcpLease6Data) SetLease6Period(v string)`

SetLease6Period sets Lease6Period field to given value.

### HasLease6Period

`func (o *DataInnerDhcpLease6Data) HasLease6Period() bool`

HasLease6Period returns a boolean if a field has been set.

### GetLease6Time

`func (o *DataInnerDhcpLease6Data) GetLease6Time() string`

GetLease6Time returns the Lease6Time field if non-nil, zero value otherwise.

### GetLease6TimeOk

`func (o *DataInnerDhcpLease6Data) GetLease6TimeOk() (*string, bool)`

GetLease6TimeOk returns a tuple with the Lease6Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Time

`func (o *DataInnerDhcpLease6Data) SetLease6Time(v string)`

SetLease6Time sets Lease6Time field to given value.

### HasLease6Time

`func (o *DataInnerDhcpLease6Data) HasLease6Time() bool`

HasLease6Time returns a boolean if a field has been set.

### GetRange6ClassName

`func (o *DataInnerDhcpLease6Data) GetRange6ClassName() string`

GetRange6ClassName returns the Range6ClassName field if non-nil, zero value otherwise.

### GetRange6ClassNameOk

`func (o *DataInnerDhcpLease6Data) GetRange6ClassNameOk() (*string, bool)`

GetRange6ClassNameOk returns a tuple with the Range6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6ClassName

`func (o *DataInnerDhcpLease6Data) SetRange6ClassName(v string)`

SetRange6ClassName sets Range6ClassName field to given value.

### HasRange6ClassName

`func (o *DataInnerDhcpLease6Data) HasRange6ClassName() bool`

HasRange6ClassName returns a boolean if a field has been set.

### GetRange6ClassParameters

`func (o *DataInnerDhcpLease6Data) GetRange6ClassParameters() []ApiClassParameterOutputEntry`

GetRange6ClassParameters returns the Range6ClassParameters field if non-nil, zero value otherwise.

### GetRange6ClassParametersOk

`func (o *DataInnerDhcpLease6Data) GetRange6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRange6ClassParametersOk returns a tuple with the Range6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6ClassParameters

`func (o *DataInnerDhcpLease6Data) SetRange6ClassParameters(v []ApiClassParameterOutputEntry)`

SetRange6ClassParameters sets Range6ClassParameters field to given value.

### HasRange6ClassParameters

`func (o *DataInnerDhcpLease6Data) HasRange6ClassParameters() bool`

HasRange6ClassParameters returns a boolean if a field has been set.

### GetRange6EndAddress6Addr

`func (o *DataInnerDhcpLease6Data) GetRange6EndAddress6Addr() string`

GetRange6EndAddress6Addr returns the Range6EndAddress6Addr field if non-nil, zero value otherwise.

### GetRange6EndAddress6AddrOk

`func (o *DataInnerDhcpLease6Data) GetRange6EndAddress6AddrOk() (*string, bool)`

GetRange6EndAddress6AddrOk returns a tuple with the Range6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6EndAddress6Addr

`func (o *DataInnerDhcpLease6Data) SetRange6EndAddress6Addr(v string)`

SetRange6EndAddress6Addr sets Range6EndAddress6Addr field to given value.

### HasRange6EndAddress6Addr

`func (o *DataInnerDhcpLease6Data) HasRange6EndAddress6Addr() bool`

HasRange6EndAddress6Addr returns a boolean if a field has been set.

### GetRange6FailoverName

`func (o *DataInnerDhcpLease6Data) GetRange6FailoverName() string`

GetRange6FailoverName returns the Range6FailoverName field if non-nil, zero value otherwise.

### GetRange6FailoverNameOk

`func (o *DataInnerDhcpLease6Data) GetRange6FailoverNameOk() (*string, bool)`

GetRange6FailoverNameOk returns a tuple with the Range6FailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6FailoverName

`func (o *DataInnerDhcpLease6Data) SetRange6FailoverName(v string)`

SetRange6FailoverName sets Range6FailoverName field to given value.

### HasRange6FailoverName

`func (o *DataInnerDhcpLease6Data) HasRange6FailoverName() bool`

HasRange6FailoverName returns a boolean if a field has been set.

### GetRange6Id

`func (o *DataInnerDhcpLease6Data) GetRange6Id() string`

GetRange6Id returns the Range6Id field if non-nil, zero value otherwise.

### GetRange6IdOk

`func (o *DataInnerDhcpLease6Data) GetRange6IdOk() (*string, bool)`

GetRange6IdOk returns a tuple with the Range6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Id

`func (o *DataInnerDhcpLease6Data) SetRange6Id(v string)`

SetRange6Id sets Range6Id field to given value.

### HasRange6Id

`func (o *DataInnerDhcpLease6Data) HasRange6Id() bool`

HasRange6Id returns a boolean if a field has been set.

### GetRange6StartAddress6Addr

`func (o *DataInnerDhcpLease6Data) GetRange6StartAddress6Addr() string`

GetRange6StartAddress6Addr returns the Range6StartAddress6Addr field if non-nil, zero value otherwise.

### GetRange6StartAddress6AddrOk

`func (o *DataInnerDhcpLease6Data) GetRange6StartAddress6AddrOk() (*string, bool)`

GetRange6StartAddress6AddrOk returns a tuple with the Range6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6StartAddress6Addr

`func (o *DataInnerDhcpLease6Data) SetRange6StartAddress6Addr(v string)`

SetRange6StartAddress6Addr sets Range6StartAddress6Addr field to given value.

### HasRange6StartAddress6Addr

`func (o *DataInnerDhcpLease6Data) HasRange6StartAddress6Addr() bool`

HasRange6StartAddress6Addr returns a boolean if a field has been set.

### GetScope6ClassName

`func (o *DataInnerDhcpLease6Data) GetScope6ClassName() string`

GetScope6ClassName returns the Scope6ClassName field if non-nil, zero value otherwise.

### GetScope6ClassNameOk

`func (o *DataInnerDhcpLease6Data) GetScope6ClassNameOk() (*string, bool)`

GetScope6ClassNameOk returns a tuple with the Scope6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassName

`func (o *DataInnerDhcpLease6Data) SetScope6ClassName(v string)`

SetScope6ClassName sets Scope6ClassName field to given value.

### HasScope6ClassName

`func (o *DataInnerDhcpLease6Data) HasScope6ClassName() bool`

HasScope6ClassName returns a boolean if a field has been set.

### GetScope6Id

`func (o *DataInnerDhcpLease6Data) GetScope6Id() string`

GetScope6Id returns the Scope6Id field if non-nil, zero value otherwise.

### GetScope6IdOk

`func (o *DataInnerDhcpLease6Data) GetScope6IdOk() (*string, bool)`

GetScope6IdOk returns a tuple with the Scope6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Id

`func (o *DataInnerDhcpLease6Data) SetScope6Id(v string)`

SetScope6Id sets Scope6Id field to given value.

### HasScope6Id

`func (o *DataInnerDhcpLease6Data) HasScope6Id() bool`

HasScope6Id returns a boolean if a field has been set.

### GetScope6Name

`func (o *DataInnerDhcpLease6Data) GetScope6Name() string`

GetScope6Name returns the Scope6Name field if non-nil, zero value otherwise.

### GetScope6NameOk

`func (o *DataInnerDhcpLease6Data) GetScope6NameOk() (*string, bool)`

GetScope6NameOk returns a tuple with the Scope6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Name

`func (o *DataInnerDhcpLease6Data) SetScope6Name(v string)`

SetScope6Name sets Scope6Name field to given value.

### HasScope6Name

`func (o *DataInnerDhcpLease6Data) HasScope6Name() bool`

HasScope6Name returns a boolean if a field has been set.

### GetScope6Prefix

`func (o *DataInnerDhcpLease6Data) GetScope6Prefix() string`

GetScope6Prefix returns the Scope6Prefix field if non-nil, zero value otherwise.

### GetScope6PrefixOk

`func (o *DataInnerDhcpLease6Data) GetScope6PrefixOk() (*string, bool)`

GetScope6PrefixOk returns a tuple with the Scope6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Prefix

`func (o *DataInnerDhcpLease6Data) SetScope6Prefix(v string)`

SetScope6Prefix sets Scope6Prefix field to given value.

### HasScope6Prefix

`func (o *DataInnerDhcpLease6Data) HasScope6Prefix() bool`

HasScope6Prefix returns a boolean if a field has been set.

### GetScope6Size

`func (o *DataInnerDhcpLease6Data) GetScope6Size() string`

GetScope6Size returns the Scope6Size field if non-nil, zero value otherwise.

### GetScope6SizeOk

`func (o *DataInnerDhcpLease6Data) GetScope6SizeOk() (*string, bool)`

GetScope6SizeOk returns a tuple with the Scope6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Size

`func (o *DataInnerDhcpLease6Data) SetScope6Size(v string)`

SetScope6Size sets Scope6Size field to given value.

### HasScope6Size

`func (o *DataInnerDhcpLease6Data) HasScope6Size() bool`

HasScope6Size returns a boolean if a field has been set.

### GetScope6StartAddress6Addr

`func (o *DataInnerDhcpLease6Data) GetScope6StartAddress6Addr() string`

GetScope6StartAddress6Addr returns the Scope6StartAddress6Addr field if non-nil, zero value otherwise.

### GetScope6StartAddress6AddrOk

`func (o *DataInnerDhcpLease6Data) GetScope6StartAddress6AddrOk() (*string, bool)`

GetScope6StartAddress6AddrOk returns a tuple with the Scope6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6StartAddress6Addr

`func (o *DataInnerDhcpLease6Data) SetScope6StartAddress6Addr(v string)`

SetScope6StartAddress6Addr sets Scope6StartAddress6Addr field to given value.

### HasScope6StartAddress6Addr

`func (o *DataInnerDhcpLease6Data) HasScope6StartAddress6Addr() bool`

HasScope6StartAddress6Addr returns a boolean if a field has been set.

### GetSharednetwork6Id

`func (o *DataInnerDhcpLease6Data) GetSharednetwork6Id() string`

GetSharednetwork6Id returns the Sharednetwork6Id field if non-nil, zero value otherwise.

### GetSharednetwork6IdOk

`func (o *DataInnerDhcpLease6Data) GetSharednetwork6IdOk() (*string, bool)`

GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Id

`func (o *DataInnerDhcpLease6Data) SetSharednetwork6Id(v string)`

SetSharednetwork6Id sets Sharednetwork6Id field to given value.

### HasSharednetwork6Id

`func (o *DataInnerDhcpLease6Data) HasSharednetwork6Id() bool`

HasSharednetwork6Id returns a boolean if a field has been set.

### GetSharednetwork6Name

`func (o *DataInnerDhcpLease6Data) GetSharednetwork6Name() string`

GetSharednetwork6Name returns the Sharednetwork6Name field if non-nil, zero value otherwise.

### GetSharednetwork6NameOk

`func (o *DataInnerDhcpLease6Data) GetSharednetwork6NameOk() (*string, bool)`

GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Name

`func (o *DataInnerDhcpLease6Data) SetSharednetwork6Name(v string)`

SetSharednetwork6Name sets Sharednetwork6Name field to given value.

### HasSharednetwork6Name

`func (o *DataInnerDhcpLease6Data) HasSharednetwork6Name() bool`

HasSharednetwork6Name returns a boolean if a field has been set.

### GetServer6Hostaddr

`func (o *DataInnerDhcpLease6Data) GetServer6Hostaddr() string`

GetServer6Hostaddr returns the Server6Hostaddr field if non-nil, zero value otherwise.

### GetServer6HostaddrOk

`func (o *DataInnerDhcpLease6Data) GetServer6HostaddrOk() (*string, bool)`

GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Hostaddr

`func (o *DataInnerDhcpLease6Data) SetServer6Hostaddr(v string)`

SetServer6Hostaddr sets Server6Hostaddr field to given value.

### HasServer6Hostaddr

`func (o *DataInnerDhcpLease6Data) HasServer6Hostaddr() bool`

HasServer6Hostaddr returns a boolean if a field has been set.

### GetServer6Addr6

`func (o *DataInnerDhcpLease6Data) GetServer6Addr6() string`

GetServer6Addr6 returns the Server6Addr6 field if non-nil, zero value otherwise.

### GetServer6Addr6Ok

`func (o *DataInnerDhcpLease6Data) GetServer6Addr6Ok() (*string, bool)`

GetServer6Addr6Ok returns a tuple with the Server6Addr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr6

`func (o *DataInnerDhcpLease6Data) SetServer6Addr6(v string)`

SetServer6Addr6 sets Server6Addr6 field to given value.

### HasServer6Addr6

`func (o *DataInnerDhcpLease6Data) HasServer6Addr6() bool`

HasServer6Addr6 returns a boolean if a field has been set.

### GetServer6Addr

`func (o *DataInnerDhcpLease6Data) GetServer6Addr() string`

GetServer6Addr returns the Server6Addr field if non-nil, zero value otherwise.

### GetServer6AddrOk

`func (o *DataInnerDhcpLease6Data) GetServer6AddrOk() (*string, bool)`

GetServer6AddrOk returns a tuple with the Server6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr

`func (o *DataInnerDhcpLease6Data) SetServer6Addr(v string)`

SetServer6Addr sets Server6Addr field to given value.

### HasServer6Addr

`func (o *DataInnerDhcpLease6Data) HasServer6Addr() bool`

HasServer6Addr returns a boolean if a field has been set.

### GetLease6MacVendor

`func (o *DataInnerDhcpLease6Data) GetLease6MacVendor() string`

GetLease6MacVendor returns the Lease6MacVendor field if non-nil, zero value otherwise.

### GetLease6MacVendorOk

`func (o *DataInnerDhcpLease6Data) GetLease6MacVendorOk() (*string, bool)`

GetLease6MacVendorOk returns a tuple with the Lease6MacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6MacVendor

`func (o *DataInnerDhcpLease6Data) SetLease6MacVendor(v string)`

SetLease6MacVendor sets Lease6MacVendor field to given value.

### HasLease6MacVendor

`func (o *DataInnerDhcpLease6Data) HasLease6MacVendor() bool`

HasLease6MacVendor returns a boolean if a field has been set.

### GetLease6Multistatus

`func (o *DataInnerDhcpLease6Data) GetLease6Multistatus() string`

GetLease6Multistatus returns the Lease6Multistatus field if non-nil, zero value otherwise.

### GetLease6MultistatusOk

`func (o *DataInnerDhcpLease6Data) GetLease6MultistatusOk() (*string, bool)`

GetLease6MultistatusOk returns a tuple with the Lease6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Multistatus

`func (o *DataInnerDhcpLease6Data) SetLease6Multistatus(v string)`

SetLease6Multistatus sets Lease6Multistatus field to given value.

### HasLease6Multistatus

`func (o *DataInnerDhcpLease6Data) HasLease6Multistatus() bool`

HasLease6Multistatus returns a boolean if a field has been set.

### GetLease6Percent

`func (o *DataInnerDhcpLease6Data) GetLease6Percent() string`

GetLease6Percent returns the Lease6Percent field if non-nil, zero value otherwise.

### GetLease6PercentOk

`func (o *DataInnerDhcpLease6Data) GetLease6PercentOk() (*string, bool)`

GetLease6PercentOk returns a tuple with the Lease6Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6Percent

`func (o *DataInnerDhcpLease6Data) SetLease6Percent(v string)`

SetLease6Percent sets Lease6Percent field to given value.

### HasLease6Percent

`func (o *DataInnerDhcpLease6Data) HasLease6Percent() bool`

HasLease6Percent returns a boolean if a field has been set.

### GetLease6TimeToExpire

`func (o *DataInnerDhcpLease6Data) GetLease6TimeToExpire() string`

GetLease6TimeToExpire returns the Lease6TimeToExpire field if non-nil, zero value otherwise.

### GetLease6TimeToExpireOk

`func (o *DataInnerDhcpLease6Data) GetLease6TimeToExpireOk() (*string, bool)`

GetLease6TimeToExpireOk returns a tuple with the Lease6TimeToExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLease6TimeToExpire

`func (o *DataInnerDhcpLease6Data) SetLease6TimeToExpire(v string)`

SetLease6TimeToExpire sets Lease6TimeToExpire field to given value.

### HasLease6TimeToExpire

`func (o *DataInnerDhcpLease6Data) HasLease6TimeToExpire() bool`

HasLease6TimeToExpire returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpLease6Data) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpLease6Data) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpLease6Data) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpLease6Data) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpLease6Data) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpLease6Data) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpLease6Data) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpLease6Data) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


