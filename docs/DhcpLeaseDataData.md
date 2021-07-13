# DhcpLeaseDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server the object belongs to. | [optional] 
**LeaseAddr** | Pointer to **string** | The IP address associated with the DHCPv4 lease. | [optional] 
**LeaseCircuitId** | Pointer to **string** | The circuit identifier (ID) of the relay agent associated with the DHCPv4 lease. | [optional] 
**LeaseClientIdent** | Pointer to **string** | The client identifier (ID) of the client associated with the DHCPv4 lease. | [optional] 
**LeaseClientname** | Pointer to **string** | The name of the client associated with the DHCPv4 lease. | [optional] 
**LeaseDomain** | Pointer to **string** | The domain name associated with the DHCPv4 lease. | [optional] 
**LeaseEndTime** | Pointer to **string** | The expiration time of the lease, in decimal UNIX date format. | [optional] 
**LeaseFingerbankOs** | Pointer to **string** | The operating system details of the client associated with the DHCPv4 lease. | [optional] 
**LeaseFirstTime** | Pointer to **string** | The first time the DHCPv4 lease has been attributed to the client, in decimal UNIX date format. | [optional] 
**LeaseId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 lease. | [optional] 
**LeaseAddressAddr** | Pointer to **string** | The IP address associated with the DHCPv4 lease, in hexadecimal format. | [optional] 
**LeaseMacAddr** | Pointer to **string** | The MAC address associated with the IPv4 lease. | [optional] 
**LeaseName** | Pointer to **string** | The name of the DHCPv4 lease. | [optional] 
**LeasePeriod** | Pointer to **string** | The duration time (time to live) of the DHCPv4 lease, in seconds. | [optional] 
**LeaseRemoteId** | Pointer to **string** | The remote identifier (ID) of the relay agent associated with the DHCPv4 lease. | [optional] 
**LeaseTime** | Pointer to **string** | The last time the DHCPv4 lease has been attributed to the client, in decimal UNIX date format. | [optional] 
**LeaseVendorId** | Pointer to **string** | The vendor class identifier (ID) of the client associated with the DHCPv4 lease. | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 range the object belongs to, it can be preceded by the class directory. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**RangeEndAddr** | Pointer to **string** | The last IP address of the DHCPv4 range the lease belongs to. | [optional] 
**RangeFailoverName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 range the object belongs to. | [optional] 
**RangeName** | Pointer to **string** | The name of the DHCPv4 range the object belongs to. | [optional] 
**RangeStartAddr** | Pointer to **string** | The first IP address of the DHCPv4 range the lease belongs to. | [optional] 
**ScopeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**ScopeId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeName** | Pointer to **string** | The name of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeNetAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeSize** | Pointer to **string** | The number of IP addresses the DHCPv4 scope the object belongs to contains. | [optional] 
**SharednetworkId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 shared network the object belongs to. | [optional] 
**SharednetworkName** | Pointer to **string** | The name of the DHCPv4 shared network the object belongs to. | [optional] 
**ServerAddr** | Pointer to **string** | The IP address of the DHCP server the object belongs to, in hexadecimal format. | [optional] 
**LeaseMacVendor** | Pointer to **string** | The vendor details of the client associated with the DHCPv4 lease. | [optional] 
**LeaseMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParameterRequestList** | Pointer to **string** | The list of parameters requested with the DHCPv4 lease returned by the server, integers separated by a comma. | [optional] 
**Percent** | Pointer to **string** | The percentage of time the lease has really been in use. | [optional] 
**TimeToExpire** | Pointer to **string** | The time left to the lease before it expires, in seconds. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpLeaseDataData

`func NewDhcpLeaseDataData() *DhcpLeaseDataData`

NewDhcpLeaseDataData instantiates a new DhcpLeaseDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpLeaseDataDataWithDefaults

`func NewDhcpLeaseDataDataWithDefaults() *DhcpLeaseDataData`

NewDhcpLeaseDataDataWithDefaults instantiates a new DhcpLeaseDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerClassName

`func (o *DhcpLeaseDataData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DhcpLeaseDataData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DhcpLeaseDataData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DhcpLeaseDataData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpLeaseDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpLeaseDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpLeaseDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpLeaseDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpLeaseDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpLeaseDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpLeaseDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpLeaseDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpLeaseDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpLeaseDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpLeaseDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpLeaseDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DhcpLeaseDataData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DhcpLeaseDataData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DhcpLeaseDataData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DhcpLeaseDataData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetLeaseAddr

`func (o *DhcpLeaseDataData) GetLeaseAddr() string`

GetLeaseAddr returns the LeaseAddr field if non-nil, zero value otherwise.

### GetLeaseAddrOk

`func (o *DhcpLeaseDataData) GetLeaseAddrOk() (*string, bool)`

GetLeaseAddrOk returns a tuple with the LeaseAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseAddr

`func (o *DhcpLeaseDataData) SetLeaseAddr(v string)`

SetLeaseAddr sets LeaseAddr field to given value.

### HasLeaseAddr

`func (o *DhcpLeaseDataData) HasLeaseAddr() bool`

HasLeaseAddr returns a boolean if a field has been set.

### GetLeaseCircuitId

`func (o *DhcpLeaseDataData) GetLeaseCircuitId() string`

GetLeaseCircuitId returns the LeaseCircuitId field if non-nil, zero value otherwise.

### GetLeaseCircuitIdOk

`func (o *DhcpLeaseDataData) GetLeaseCircuitIdOk() (*string, bool)`

GetLeaseCircuitIdOk returns a tuple with the LeaseCircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseCircuitId

`func (o *DhcpLeaseDataData) SetLeaseCircuitId(v string)`

SetLeaseCircuitId sets LeaseCircuitId field to given value.

### HasLeaseCircuitId

`func (o *DhcpLeaseDataData) HasLeaseCircuitId() bool`

HasLeaseCircuitId returns a boolean if a field has been set.

### GetLeaseClientIdent

`func (o *DhcpLeaseDataData) GetLeaseClientIdent() string`

GetLeaseClientIdent returns the LeaseClientIdent field if non-nil, zero value otherwise.

### GetLeaseClientIdentOk

`func (o *DhcpLeaseDataData) GetLeaseClientIdentOk() (*string, bool)`

GetLeaseClientIdentOk returns a tuple with the LeaseClientIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseClientIdent

`func (o *DhcpLeaseDataData) SetLeaseClientIdent(v string)`

SetLeaseClientIdent sets LeaseClientIdent field to given value.

### HasLeaseClientIdent

`func (o *DhcpLeaseDataData) HasLeaseClientIdent() bool`

HasLeaseClientIdent returns a boolean if a field has been set.

### GetLeaseClientname

`func (o *DhcpLeaseDataData) GetLeaseClientname() string`

GetLeaseClientname returns the LeaseClientname field if non-nil, zero value otherwise.

### GetLeaseClientnameOk

`func (o *DhcpLeaseDataData) GetLeaseClientnameOk() (*string, bool)`

GetLeaseClientnameOk returns a tuple with the LeaseClientname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseClientname

`func (o *DhcpLeaseDataData) SetLeaseClientname(v string)`

SetLeaseClientname sets LeaseClientname field to given value.

### HasLeaseClientname

`func (o *DhcpLeaseDataData) HasLeaseClientname() bool`

HasLeaseClientname returns a boolean if a field has been set.

### GetLeaseDomain

`func (o *DhcpLeaseDataData) GetLeaseDomain() string`

GetLeaseDomain returns the LeaseDomain field if non-nil, zero value otherwise.

### GetLeaseDomainOk

`func (o *DhcpLeaseDataData) GetLeaseDomainOk() (*string, bool)`

GetLeaseDomainOk returns a tuple with the LeaseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDomain

`func (o *DhcpLeaseDataData) SetLeaseDomain(v string)`

SetLeaseDomain sets LeaseDomain field to given value.

### HasLeaseDomain

`func (o *DhcpLeaseDataData) HasLeaseDomain() bool`

HasLeaseDomain returns a boolean if a field has been set.

### GetLeaseEndTime

`func (o *DhcpLeaseDataData) GetLeaseEndTime() string`

GetLeaseEndTime returns the LeaseEndTime field if non-nil, zero value otherwise.

### GetLeaseEndTimeOk

`func (o *DhcpLeaseDataData) GetLeaseEndTimeOk() (*string, bool)`

GetLeaseEndTimeOk returns a tuple with the LeaseEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseEndTime

`func (o *DhcpLeaseDataData) SetLeaseEndTime(v string)`

SetLeaseEndTime sets LeaseEndTime field to given value.

### HasLeaseEndTime

`func (o *DhcpLeaseDataData) HasLeaseEndTime() bool`

HasLeaseEndTime returns a boolean if a field has been set.

### GetLeaseFingerbankOs

`func (o *DhcpLeaseDataData) GetLeaseFingerbankOs() string`

GetLeaseFingerbankOs returns the LeaseFingerbankOs field if non-nil, zero value otherwise.

### GetLeaseFingerbankOsOk

`func (o *DhcpLeaseDataData) GetLeaseFingerbankOsOk() (*string, bool)`

GetLeaseFingerbankOsOk returns a tuple with the LeaseFingerbankOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseFingerbankOs

`func (o *DhcpLeaseDataData) SetLeaseFingerbankOs(v string)`

SetLeaseFingerbankOs sets LeaseFingerbankOs field to given value.

### HasLeaseFingerbankOs

`func (o *DhcpLeaseDataData) HasLeaseFingerbankOs() bool`

HasLeaseFingerbankOs returns a boolean if a field has been set.

### GetLeaseFirstTime

`func (o *DhcpLeaseDataData) GetLeaseFirstTime() string`

GetLeaseFirstTime returns the LeaseFirstTime field if non-nil, zero value otherwise.

### GetLeaseFirstTimeOk

`func (o *DhcpLeaseDataData) GetLeaseFirstTimeOk() (*string, bool)`

GetLeaseFirstTimeOk returns a tuple with the LeaseFirstTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseFirstTime

`func (o *DhcpLeaseDataData) SetLeaseFirstTime(v string)`

SetLeaseFirstTime sets LeaseFirstTime field to given value.

### HasLeaseFirstTime

`func (o *DhcpLeaseDataData) HasLeaseFirstTime() bool`

HasLeaseFirstTime returns a boolean if a field has been set.

### GetLeaseId

`func (o *DhcpLeaseDataData) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *DhcpLeaseDataData) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *DhcpLeaseDataData) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.

### HasLeaseId

`func (o *DhcpLeaseDataData) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.

### GetLeaseAddressAddr

`func (o *DhcpLeaseDataData) GetLeaseAddressAddr() string`

GetLeaseAddressAddr returns the LeaseAddressAddr field if non-nil, zero value otherwise.

### GetLeaseAddressAddrOk

`func (o *DhcpLeaseDataData) GetLeaseAddressAddrOk() (*string, bool)`

GetLeaseAddressAddrOk returns a tuple with the LeaseAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseAddressAddr

`func (o *DhcpLeaseDataData) SetLeaseAddressAddr(v string)`

SetLeaseAddressAddr sets LeaseAddressAddr field to given value.

### HasLeaseAddressAddr

`func (o *DhcpLeaseDataData) HasLeaseAddressAddr() bool`

HasLeaseAddressAddr returns a boolean if a field has been set.

### GetLeaseMacAddr

`func (o *DhcpLeaseDataData) GetLeaseMacAddr() string`

GetLeaseMacAddr returns the LeaseMacAddr field if non-nil, zero value otherwise.

### GetLeaseMacAddrOk

`func (o *DhcpLeaseDataData) GetLeaseMacAddrOk() (*string, bool)`

GetLeaseMacAddrOk returns a tuple with the LeaseMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseMacAddr

`func (o *DhcpLeaseDataData) SetLeaseMacAddr(v string)`

SetLeaseMacAddr sets LeaseMacAddr field to given value.

### HasLeaseMacAddr

`func (o *DhcpLeaseDataData) HasLeaseMacAddr() bool`

HasLeaseMacAddr returns a boolean if a field has been set.

### GetLeaseName

`func (o *DhcpLeaseDataData) GetLeaseName() string`

GetLeaseName returns the LeaseName field if non-nil, zero value otherwise.

### GetLeaseNameOk

`func (o *DhcpLeaseDataData) GetLeaseNameOk() (*string, bool)`

GetLeaseNameOk returns a tuple with the LeaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseName

`func (o *DhcpLeaseDataData) SetLeaseName(v string)`

SetLeaseName sets LeaseName field to given value.

### HasLeaseName

`func (o *DhcpLeaseDataData) HasLeaseName() bool`

HasLeaseName returns a boolean if a field has been set.

### GetLeasePeriod

`func (o *DhcpLeaseDataData) GetLeasePeriod() string`

GetLeasePeriod returns the LeasePeriod field if non-nil, zero value otherwise.

### GetLeasePeriodOk

`func (o *DhcpLeaseDataData) GetLeasePeriodOk() (*string, bool)`

GetLeasePeriodOk returns a tuple with the LeasePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasePeriod

`func (o *DhcpLeaseDataData) SetLeasePeriod(v string)`

SetLeasePeriod sets LeasePeriod field to given value.

### HasLeasePeriod

`func (o *DhcpLeaseDataData) HasLeasePeriod() bool`

HasLeasePeriod returns a boolean if a field has been set.

### GetLeaseRemoteId

`func (o *DhcpLeaseDataData) GetLeaseRemoteId() string`

GetLeaseRemoteId returns the LeaseRemoteId field if non-nil, zero value otherwise.

### GetLeaseRemoteIdOk

`func (o *DhcpLeaseDataData) GetLeaseRemoteIdOk() (*string, bool)`

GetLeaseRemoteIdOk returns a tuple with the LeaseRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseRemoteId

`func (o *DhcpLeaseDataData) SetLeaseRemoteId(v string)`

SetLeaseRemoteId sets LeaseRemoteId field to given value.

### HasLeaseRemoteId

`func (o *DhcpLeaseDataData) HasLeaseRemoteId() bool`

HasLeaseRemoteId returns a boolean if a field has been set.

### GetLeaseTime

`func (o *DhcpLeaseDataData) GetLeaseTime() string`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *DhcpLeaseDataData) GetLeaseTimeOk() (*string, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *DhcpLeaseDataData) SetLeaseTime(v string)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *DhcpLeaseDataData) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetLeaseVendorId

`func (o *DhcpLeaseDataData) GetLeaseVendorId() string`

GetLeaseVendorId returns the LeaseVendorId field if non-nil, zero value otherwise.

### GetLeaseVendorIdOk

`func (o *DhcpLeaseDataData) GetLeaseVendorIdOk() (*string, bool)`

GetLeaseVendorIdOk returns a tuple with the LeaseVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseVendorId

`func (o *DhcpLeaseDataData) SetLeaseVendorId(v string)`

SetLeaseVendorId sets LeaseVendorId field to given value.

### HasLeaseVendorId

`func (o *DhcpLeaseDataData) HasLeaseVendorId() bool`

HasLeaseVendorId returns a boolean if a field has been set.

### GetRangeClassName

`func (o *DhcpLeaseDataData) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *DhcpLeaseDataData) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *DhcpLeaseDataData) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *DhcpLeaseDataData) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *DhcpLeaseDataData) GetRangeClassParameters() []ApiClassParameterOutputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *DhcpLeaseDataData) GetRangeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *DhcpLeaseDataData) SetRangeClassParameters(v []ApiClassParameterOutputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *DhcpLeaseDataData) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetRangeEndAddr

`func (o *DhcpLeaseDataData) GetRangeEndAddr() string`

GetRangeEndAddr returns the RangeEndAddr field if non-nil, zero value otherwise.

### GetRangeEndAddrOk

`func (o *DhcpLeaseDataData) GetRangeEndAddrOk() (*string, bool)`

GetRangeEndAddrOk returns a tuple with the RangeEndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndAddr

`func (o *DhcpLeaseDataData) SetRangeEndAddr(v string)`

SetRangeEndAddr sets RangeEndAddr field to given value.

### HasRangeEndAddr

`func (o *DhcpLeaseDataData) HasRangeEndAddr() bool`

HasRangeEndAddr returns a boolean if a field has been set.

### GetRangeFailoverName

`func (o *DhcpLeaseDataData) GetRangeFailoverName() string`

GetRangeFailoverName returns the RangeFailoverName field if non-nil, zero value otherwise.

### GetRangeFailoverNameOk

`func (o *DhcpLeaseDataData) GetRangeFailoverNameOk() (*string, bool)`

GetRangeFailoverNameOk returns a tuple with the RangeFailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeFailoverName

`func (o *DhcpLeaseDataData) SetRangeFailoverName(v string)`

SetRangeFailoverName sets RangeFailoverName field to given value.

### HasRangeFailoverName

`func (o *DhcpLeaseDataData) HasRangeFailoverName() bool`

HasRangeFailoverName returns a boolean if a field has been set.

### GetRangeId

`func (o *DhcpLeaseDataData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *DhcpLeaseDataData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *DhcpLeaseDataData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *DhcpLeaseDataData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *DhcpLeaseDataData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *DhcpLeaseDataData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *DhcpLeaseDataData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *DhcpLeaseDataData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetRangeStartAddr

`func (o *DhcpLeaseDataData) GetRangeStartAddr() string`

GetRangeStartAddr returns the RangeStartAddr field if non-nil, zero value otherwise.

### GetRangeStartAddrOk

`func (o *DhcpLeaseDataData) GetRangeStartAddrOk() (*string, bool)`

GetRangeStartAddrOk returns a tuple with the RangeStartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartAddr

`func (o *DhcpLeaseDataData) SetRangeStartAddr(v string)`

SetRangeStartAddr sets RangeStartAddr field to given value.

### HasRangeStartAddr

`func (o *DhcpLeaseDataData) HasRangeStartAddr() bool`

HasRangeStartAddr returns a boolean if a field has been set.

### GetScopeClassName

`func (o *DhcpLeaseDataData) GetScopeClassName() string`

GetScopeClassName returns the ScopeClassName field if non-nil, zero value otherwise.

### GetScopeClassNameOk

`func (o *DhcpLeaseDataData) GetScopeClassNameOk() (*string, bool)`

GetScopeClassNameOk returns a tuple with the ScopeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassName

`func (o *DhcpLeaseDataData) SetScopeClassName(v string)`

SetScopeClassName sets ScopeClassName field to given value.

### HasScopeClassName

`func (o *DhcpLeaseDataData) HasScopeClassName() bool`

HasScopeClassName returns a boolean if a field has been set.

### GetScopeId

`func (o *DhcpLeaseDataData) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DhcpLeaseDataData) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DhcpLeaseDataData) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *DhcpLeaseDataData) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetScopeName

`func (o *DhcpLeaseDataData) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DhcpLeaseDataData) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DhcpLeaseDataData) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DhcpLeaseDataData) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeNetAddr

`func (o *DhcpLeaseDataData) GetScopeNetAddr() string`

GetScopeNetAddr returns the ScopeNetAddr field if non-nil, zero value otherwise.

### GetScopeNetAddrOk

`func (o *DhcpLeaseDataData) GetScopeNetAddrOk() (*string, bool)`

GetScopeNetAddrOk returns a tuple with the ScopeNetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetAddr

`func (o *DhcpLeaseDataData) SetScopeNetAddr(v string)`

SetScopeNetAddr sets ScopeNetAddr field to given value.

### HasScopeNetAddr

`func (o *DhcpLeaseDataData) HasScopeNetAddr() bool`

HasScopeNetAddr returns a boolean if a field has been set.

### GetScopeSize

`func (o *DhcpLeaseDataData) GetScopeSize() string`

GetScopeSize returns the ScopeSize field if non-nil, zero value otherwise.

### GetScopeSizeOk

`func (o *DhcpLeaseDataData) GetScopeSizeOk() (*string, bool)`

GetScopeSizeOk returns a tuple with the ScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSize

`func (o *DhcpLeaseDataData) SetScopeSize(v string)`

SetScopeSize sets ScopeSize field to given value.

### HasScopeSize

`func (o *DhcpLeaseDataData) HasScopeSize() bool`

HasScopeSize returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DhcpLeaseDataData) GetSharednetworkId() string`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DhcpLeaseDataData) GetSharednetworkIdOk() (*string, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DhcpLeaseDataData) SetSharednetworkId(v string)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DhcpLeaseDataData) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DhcpLeaseDataData) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DhcpLeaseDataData) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DhcpLeaseDataData) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DhcpLeaseDataData) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetServerAddr

`func (o *DhcpLeaseDataData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DhcpLeaseDataData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DhcpLeaseDataData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DhcpLeaseDataData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetLeaseMacVendor

`func (o *DhcpLeaseDataData) GetLeaseMacVendor() string`

GetLeaseMacVendor returns the LeaseMacVendor field if non-nil, zero value otherwise.

### GetLeaseMacVendorOk

`func (o *DhcpLeaseDataData) GetLeaseMacVendorOk() (*string, bool)`

GetLeaseMacVendorOk returns a tuple with the LeaseMacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseMacVendor

`func (o *DhcpLeaseDataData) SetLeaseMacVendor(v string)`

SetLeaseMacVendor sets LeaseMacVendor field to given value.

### HasLeaseMacVendor

`func (o *DhcpLeaseDataData) HasLeaseMacVendor() bool`

HasLeaseMacVendor returns a boolean if a field has been set.

### GetLeaseMultistatus

`func (o *DhcpLeaseDataData) GetLeaseMultistatus() string`

GetLeaseMultistatus returns the LeaseMultistatus field if non-nil, zero value otherwise.

### GetLeaseMultistatusOk

`func (o *DhcpLeaseDataData) GetLeaseMultistatusOk() (*string, bool)`

GetLeaseMultistatusOk returns a tuple with the LeaseMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseMultistatus

`func (o *DhcpLeaseDataData) SetLeaseMultistatus(v string)`

SetLeaseMultistatus sets LeaseMultistatus field to given value.

### HasLeaseMultistatus

`func (o *DhcpLeaseDataData) HasLeaseMultistatus() bool`

HasLeaseMultistatus returns a boolean if a field has been set.

### GetParameterRequestList

`func (o *DhcpLeaseDataData) GetParameterRequestList() string`

GetParameterRequestList returns the ParameterRequestList field if non-nil, zero value otherwise.

### GetParameterRequestListOk

`func (o *DhcpLeaseDataData) GetParameterRequestListOk() (*string, bool)`

GetParameterRequestListOk returns a tuple with the ParameterRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterRequestList

`func (o *DhcpLeaseDataData) SetParameterRequestList(v string)`

SetParameterRequestList sets ParameterRequestList field to given value.

### HasParameterRequestList

`func (o *DhcpLeaseDataData) HasParameterRequestList() bool`

HasParameterRequestList returns a boolean if a field has been set.

### GetPercent

`func (o *DhcpLeaseDataData) GetPercent() string`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *DhcpLeaseDataData) GetPercentOk() (*string, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *DhcpLeaseDataData) SetPercent(v string)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *DhcpLeaseDataData) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetTimeToExpire

`func (o *DhcpLeaseDataData) GetTimeToExpire() string`

GetTimeToExpire returns the TimeToExpire field if non-nil, zero value otherwise.

### GetTimeToExpireOk

`func (o *DhcpLeaseDataData) GetTimeToExpireOk() (*string, bool)`

GetTimeToExpireOk returns a tuple with the TimeToExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToExpire

`func (o *DhcpLeaseDataData) SetTimeToExpire(v string)`

SetTimeToExpire sets TimeToExpire field to given value.

### HasTimeToExpire

`func (o *DhcpLeaseDataData) HasTimeToExpire() bool`

HasTimeToExpire returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpLeaseDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpLeaseDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpLeaseDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpLeaseDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DhcpLeaseDataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DhcpLeaseDataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DhcpLeaseDataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DhcpLeaseDataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


