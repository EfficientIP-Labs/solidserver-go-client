# DataInnerDhcpLeaseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerClusterRole** | Pointer to **string** | The role of the server the object belongs to in the cluster, either &lt;b&gt;active (M)&lt;/b&gt;, &lt;b&gt;passive (B)&lt;/b&gt; or &lt;b&gt;N/A (#)&lt;/b&gt;. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft Windows DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv4 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server the object belongs to. | [optional] 
**LeaseAddr** | Pointer to **string** | The IP address associated with the DHCPv4 lease. | [optional] 
**LeaseCircuitId** | Pointer to **string** | The circuit identifier (ID) of the relay agent associated with the DHCPv4 lease. | [optional] 
**LeaseClientIdent** | Pointer to **string** | The client identifier (ID) of the client associated with the DHCPv4 lease. | [optional] 
**LeaseClientname** | Pointer to **string** | The name of the client associated with the DHCPv4 lease. | [optional] 
**LeaseDomain** | Pointer to **string** | The domain name associated with the DHCPv4 lease. | [optional] 
**LeaseEndTime** | Pointer to **string** | The expiration time of the lease, in decimal UNIX date format. | [optional] 
**LeaseFingerbankOs** | Pointer to **string** | The operating system details of the client associated with the DHCPv4 lease. | [optional] 
**LeaseFirstTime** | Pointer to **string** | The first time the DHCPv4 lease has been attributed to the client, in decimal UNIX date format. | [optional] 
**LeaseGiaddr** | Pointer to **string** | The gateway IP address of the relay agent of the DHCPv4 lease. | [optional] 
**LeaseId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 lease. | [optional] 
**LeaseAddressAddr** | Pointer to **string** | The IP address associated with the DHCPv4 lease, in hexadecimal format. | [optional] 
**LeaseMacAddr** | Pointer to **string** | The MAC address associated with the IPv4 lease. | [optional] 
**LeaseName** | Pointer to **string** | The name of the DHCPv4 lease. | [optional] 
**LeasePeriod** | Pointer to **string** | The duration time (time to live) of the DHCPv4 lease, in seconds. | [optional] 
**LeaseRemoteId** | Pointer to **string** | The remote identifier (ID) of the relay agent associated with the DHCPv4 lease. | [optional] 
**LeaseTime** | Pointer to **string** | The last time the DHCPv4 lease has been attributed to the client, in decimal UNIX date format. | [optional] 
**LeaseVendorId** | Pointer to **string** | The vendor class identifier (ID) of the client associated with the DHCPv4 lease. | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 range the object belongs to, it can be preceded by the class directory. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 range the object belongs to. | [optional] 
**RangeEndAddr** | Pointer to **string** | The last IP address of the DHCPv4 range the lease belongs to. | [optional] 
**RangeFailoverName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 range the object belongs to. | [optional] 
**RangeName** | Pointer to **string** | The start and end IP address of the DHCPv4 range the object belongs to, &lt;b&gt;range_start_addr&lt;/b&gt; and &lt;b&gt;range_end_addr&lt;/b&gt;, as follows: &lt;b&gt;&amp;lt;start-ip&amp;gt;-&amp;lt;end-ip&amp;gt;&lt;/b&gt;. | [optional] 
**RangeStartAddr** | Pointer to **string** | The first IP address of the DHCPv4 range the lease belongs to. | [optional] 
**ScopeClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 scope the object belongs to, it can be preceded by the class directory. | [optional] 
**ScopeId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeName** | Pointer to **string** | The name of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeNetAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope the object belongs to. | [optional] 
**ScopeSize** | Pointer to **string** | The number of IP addresses the DHCPv4 scope the object belongs to contains. | [optional] 
**SharednetworkId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 shared network the object belongs to. | [optional] 
**SharednetworkName** | Pointer to **string** | The name of the DHCPv4 shared network the object belongs to. | [optional] 
**ServerHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server_addr&lt;/b&gt; or &lt;b&gt;server_addr6&lt;/b&gt;. | [optional] 
**ServerAddr6** | Pointer to **string** | The Management IP address of the DHCPv4 server the object belongs to, the IPv6 address configured when adding the server, in hexadecimal format. | [optional] 
**ServerAddr** | Pointer to **string** | The Management IP address of the DHCPv4 server the object belongs to, the IPv4 address configured when adding the server, in hexadecimal format. | [optional] 
**LeaseMacVendor** | Pointer to **string** | The vendor details of the client associated with the DHCPv4 lease. | [optional] 
**LeaseMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParameterRequestList** | Pointer to **string** | The list of parameters requested with the DHCPv4 lease returned by the server, integers separated by a comma. | [optional] 
**Percent** | Pointer to **string** | The percentage of time the lease has really been in use. | [optional] 
**TimeToExpire** | Pointer to **string** | The time left to the lease before it expires, in seconds. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpLeaseData

`func NewDataInnerDhcpLeaseData() *DataInnerDhcpLeaseData`

NewDataInnerDhcpLeaseData instantiates a new DataInnerDhcpLeaseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpLeaseDataWithDefaults

`func NewDataInnerDhcpLeaseDataWithDefaults() *DataInnerDhcpLeaseData`

NewDataInnerDhcpLeaseDataWithDefaults instantiates a new DataInnerDhcpLeaseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerClusterRole

`func (o *DataInnerDhcpLeaseData) GetServerClusterRole() string`

GetServerClusterRole returns the ServerClusterRole field if non-nil, zero value otherwise.

### GetServerClusterRoleOk

`func (o *DataInnerDhcpLeaseData) GetServerClusterRoleOk() (*string, bool)`

GetServerClusterRoleOk returns a tuple with the ServerClusterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClusterRole

`func (o *DataInnerDhcpLeaseData) SetServerClusterRole(v string)`

SetServerClusterRole sets ServerClusterRole field to given value.

### HasServerClusterRole

`func (o *DataInnerDhcpLeaseData) HasServerClusterRole() bool`

HasServerClusterRole returns a boolean if a field has been set.

### GetServerClassName

`func (o *DataInnerDhcpLeaseData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DataInnerDhcpLeaseData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DataInnerDhcpLeaseData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DataInnerDhcpLeaseData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDhcpLeaseData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDhcpLeaseData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDhcpLeaseData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDhcpLeaseData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDhcpLeaseData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDhcpLeaseData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDhcpLeaseData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDhcpLeaseData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDhcpLeaseData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDhcpLeaseData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDhcpLeaseData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDhcpLeaseData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DataInnerDhcpLeaseData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DataInnerDhcpLeaseData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DataInnerDhcpLeaseData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DataInnerDhcpLeaseData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetLeaseAddr

`func (o *DataInnerDhcpLeaseData) GetLeaseAddr() string`

GetLeaseAddr returns the LeaseAddr field if non-nil, zero value otherwise.

### GetLeaseAddrOk

`func (o *DataInnerDhcpLeaseData) GetLeaseAddrOk() (*string, bool)`

GetLeaseAddrOk returns a tuple with the LeaseAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseAddr

`func (o *DataInnerDhcpLeaseData) SetLeaseAddr(v string)`

SetLeaseAddr sets LeaseAddr field to given value.

### HasLeaseAddr

`func (o *DataInnerDhcpLeaseData) HasLeaseAddr() bool`

HasLeaseAddr returns a boolean if a field has been set.

### GetLeaseCircuitId

`func (o *DataInnerDhcpLeaseData) GetLeaseCircuitId() string`

GetLeaseCircuitId returns the LeaseCircuitId field if non-nil, zero value otherwise.

### GetLeaseCircuitIdOk

`func (o *DataInnerDhcpLeaseData) GetLeaseCircuitIdOk() (*string, bool)`

GetLeaseCircuitIdOk returns a tuple with the LeaseCircuitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseCircuitId

`func (o *DataInnerDhcpLeaseData) SetLeaseCircuitId(v string)`

SetLeaseCircuitId sets LeaseCircuitId field to given value.

### HasLeaseCircuitId

`func (o *DataInnerDhcpLeaseData) HasLeaseCircuitId() bool`

HasLeaseCircuitId returns a boolean if a field has been set.

### GetLeaseClientIdent

`func (o *DataInnerDhcpLeaseData) GetLeaseClientIdent() string`

GetLeaseClientIdent returns the LeaseClientIdent field if non-nil, zero value otherwise.

### GetLeaseClientIdentOk

`func (o *DataInnerDhcpLeaseData) GetLeaseClientIdentOk() (*string, bool)`

GetLeaseClientIdentOk returns a tuple with the LeaseClientIdent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseClientIdent

`func (o *DataInnerDhcpLeaseData) SetLeaseClientIdent(v string)`

SetLeaseClientIdent sets LeaseClientIdent field to given value.

### HasLeaseClientIdent

`func (o *DataInnerDhcpLeaseData) HasLeaseClientIdent() bool`

HasLeaseClientIdent returns a boolean if a field has been set.

### GetLeaseClientname

`func (o *DataInnerDhcpLeaseData) GetLeaseClientname() string`

GetLeaseClientname returns the LeaseClientname field if non-nil, zero value otherwise.

### GetLeaseClientnameOk

`func (o *DataInnerDhcpLeaseData) GetLeaseClientnameOk() (*string, bool)`

GetLeaseClientnameOk returns a tuple with the LeaseClientname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseClientname

`func (o *DataInnerDhcpLeaseData) SetLeaseClientname(v string)`

SetLeaseClientname sets LeaseClientname field to given value.

### HasLeaseClientname

`func (o *DataInnerDhcpLeaseData) HasLeaseClientname() bool`

HasLeaseClientname returns a boolean if a field has been set.

### GetLeaseDomain

`func (o *DataInnerDhcpLeaseData) GetLeaseDomain() string`

GetLeaseDomain returns the LeaseDomain field if non-nil, zero value otherwise.

### GetLeaseDomainOk

`func (o *DataInnerDhcpLeaseData) GetLeaseDomainOk() (*string, bool)`

GetLeaseDomainOk returns a tuple with the LeaseDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseDomain

`func (o *DataInnerDhcpLeaseData) SetLeaseDomain(v string)`

SetLeaseDomain sets LeaseDomain field to given value.

### HasLeaseDomain

`func (o *DataInnerDhcpLeaseData) HasLeaseDomain() bool`

HasLeaseDomain returns a boolean if a field has been set.

### GetLeaseEndTime

`func (o *DataInnerDhcpLeaseData) GetLeaseEndTime() string`

GetLeaseEndTime returns the LeaseEndTime field if non-nil, zero value otherwise.

### GetLeaseEndTimeOk

`func (o *DataInnerDhcpLeaseData) GetLeaseEndTimeOk() (*string, bool)`

GetLeaseEndTimeOk returns a tuple with the LeaseEndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseEndTime

`func (o *DataInnerDhcpLeaseData) SetLeaseEndTime(v string)`

SetLeaseEndTime sets LeaseEndTime field to given value.

### HasLeaseEndTime

`func (o *DataInnerDhcpLeaseData) HasLeaseEndTime() bool`

HasLeaseEndTime returns a boolean if a field has been set.

### GetLeaseFingerbankOs

`func (o *DataInnerDhcpLeaseData) GetLeaseFingerbankOs() string`

GetLeaseFingerbankOs returns the LeaseFingerbankOs field if non-nil, zero value otherwise.

### GetLeaseFingerbankOsOk

`func (o *DataInnerDhcpLeaseData) GetLeaseFingerbankOsOk() (*string, bool)`

GetLeaseFingerbankOsOk returns a tuple with the LeaseFingerbankOs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseFingerbankOs

`func (o *DataInnerDhcpLeaseData) SetLeaseFingerbankOs(v string)`

SetLeaseFingerbankOs sets LeaseFingerbankOs field to given value.

### HasLeaseFingerbankOs

`func (o *DataInnerDhcpLeaseData) HasLeaseFingerbankOs() bool`

HasLeaseFingerbankOs returns a boolean if a field has been set.

### GetLeaseFirstTime

`func (o *DataInnerDhcpLeaseData) GetLeaseFirstTime() string`

GetLeaseFirstTime returns the LeaseFirstTime field if non-nil, zero value otherwise.

### GetLeaseFirstTimeOk

`func (o *DataInnerDhcpLeaseData) GetLeaseFirstTimeOk() (*string, bool)`

GetLeaseFirstTimeOk returns a tuple with the LeaseFirstTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseFirstTime

`func (o *DataInnerDhcpLeaseData) SetLeaseFirstTime(v string)`

SetLeaseFirstTime sets LeaseFirstTime field to given value.

### HasLeaseFirstTime

`func (o *DataInnerDhcpLeaseData) HasLeaseFirstTime() bool`

HasLeaseFirstTime returns a boolean if a field has been set.

### GetLeaseGiaddr

`func (o *DataInnerDhcpLeaseData) GetLeaseGiaddr() string`

GetLeaseGiaddr returns the LeaseGiaddr field if non-nil, zero value otherwise.

### GetLeaseGiaddrOk

`func (o *DataInnerDhcpLeaseData) GetLeaseGiaddrOk() (*string, bool)`

GetLeaseGiaddrOk returns a tuple with the LeaseGiaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseGiaddr

`func (o *DataInnerDhcpLeaseData) SetLeaseGiaddr(v string)`

SetLeaseGiaddr sets LeaseGiaddr field to given value.

### HasLeaseGiaddr

`func (o *DataInnerDhcpLeaseData) HasLeaseGiaddr() bool`

HasLeaseGiaddr returns a boolean if a field has been set.

### GetLeaseId

`func (o *DataInnerDhcpLeaseData) GetLeaseId() string`

GetLeaseId returns the LeaseId field if non-nil, zero value otherwise.

### GetLeaseIdOk

`func (o *DataInnerDhcpLeaseData) GetLeaseIdOk() (*string, bool)`

GetLeaseIdOk returns a tuple with the LeaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseId

`func (o *DataInnerDhcpLeaseData) SetLeaseId(v string)`

SetLeaseId sets LeaseId field to given value.

### HasLeaseId

`func (o *DataInnerDhcpLeaseData) HasLeaseId() bool`

HasLeaseId returns a boolean if a field has been set.

### GetLeaseAddressAddr

`func (o *DataInnerDhcpLeaseData) GetLeaseAddressAddr() string`

GetLeaseAddressAddr returns the LeaseAddressAddr field if non-nil, zero value otherwise.

### GetLeaseAddressAddrOk

`func (o *DataInnerDhcpLeaseData) GetLeaseAddressAddrOk() (*string, bool)`

GetLeaseAddressAddrOk returns a tuple with the LeaseAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseAddressAddr

`func (o *DataInnerDhcpLeaseData) SetLeaseAddressAddr(v string)`

SetLeaseAddressAddr sets LeaseAddressAddr field to given value.

### HasLeaseAddressAddr

`func (o *DataInnerDhcpLeaseData) HasLeaseAddressAddr() bool`

HasLeaseAddressAddr returns a boolean if a field has been set.

### GetLeaseMacAddr

`func (o *DataInnerDhcpLeaseData) GetLeaseMacAddr() string`

GetLeaseMacAddr returns the LeaseMacAddr field if non-nil, zero value otherwise.

### GetLeaseMacAddrOk

`func (o *DataInnerDhcpLeaseData) GetLeaseMacAddrOk() (*string, bool)`

GetLeaseMacAddrOk returns a tuple with the LeaseMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseMacAddr

`func (o *DataInnerDhcpLeaseData) SetLeaseMacAddr(v string)`

SetLeaseMacAddr sets LeaseMacAddr field to given value.

### HasLeaseMacAddr

`func (o *DataInnerDhcpLeaseData) HasLeaseMacAddr() bool`

HasLeaseMacAddr returns a boolean if a field has been set.

### GetLeaseName

`func (o *DataInnerDhcpLeaseData) GetLeaseName() string`

GetLeaseName returns the LeaseName field if non-nil, zero value otherwise.

### GetLeaseNameOk

`func (o *DataInnerDhcpLeaseData) GetLeaseNameOk() (*string, bool)`

GetLeaseNameOk returns a tuple with the LeaseName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseName

`func (o *DataInnerDhcpLeaseData) SetLeaseName(v string)`

SetLeaseName sets LeaseName field to given value.

### HasLeaseName

`func (o *DataInnerDhcpLeaseData) HasLeaseName() bool`

HasLeaseName returns a boolean if a field has been set.

### GetLeasePeriod

`func (o *DataInnerDhcpLeaseData) GetLeasePeriod() string`

GetLeasePeriod returns the LeasePeriod field if non-nil, zero value otherwise.

### GetLeasePeriodOk

`func (o *DataInnerDhcpLeaseData) GetLeasePeriodOk() (*string, bool)`

GetLeasePeriodOk returns a tuple with the LeasePeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeasePeriod

`func (o *DataInnerDhcpLeaseData) SetLeasePeriod(v string)`

SetLeasePeriod sets LeasePeriod field to given value.

### HasLeasePeriod

`func (o *DataInnerDhcpLeaseData) HasLeasePeriod() bool`

HasLeasePeriod returns a boolean if a field has been set.

### GetLeaseRemoteId

`func (o *DataInnerDhcpLeaseData) GetLeaseRemoteId() string`

GetLeaseRemoteId returns the LeaseRemoteId field if non-nil, zero value otherwise.

### GetLeaseRemoteIdOk

`func (o *DataInnerDhcpLeaseData) GetLeaseRemoteIdOk() (*string, bool)`

GetLeaseRemoteIdOk returns a tuple with the LeaseRemoteId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseRemoteId

`func (o *DataInnerDhcpLeaseData) SetLeaseRemoteId(v string)`

SetLeaseRemoteId sets LeaseRemoteId field to given value.

### HasLeaseRemoteId

`func (o *DataInnerDhcpLeaseData) HasLeaseRemoteId() bool`

HasLeaseRemoteId returns a boolean if a field has been set.

### GetLeaseTime

`func (o *DataInnerDhcpLeaseData) GetLeaseTime() string`

GetLeaseTime returns the LeaseTime field if non-nil, zero value otherwise.

### GetLeaseTimeOk

`func (o *DataInnerDhcpLeaseData) GetLeaseTimeOk() (*string, bool)`

GetLeaseTimeOk returns a tuple with the LeaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseTime

`func (o *DataInnerDhcpLeaseData) SetLeaseTime(v string)`

SetLeaseTime sets LeaseTime field to given value.

### HasLeaseTime

`func (o *DataInnerDhcpLeaseData) HasLeaseTime() bool`

HasLeaseTime returns a boolean if a field has been set.

### GetLeaseVendorId

`func (o *DataInnerDhcpLeaseData) GetLeaseVendorId() string`

GetLeaseVendorId returns the LeaseVendorId field if non-nil, zero value otherwise.

### GetLeaseVendorIdOk

`func (o *DataInnerDhcpLeaseData) GetLeaseVendorIdOk() (*string, bool)`

GetLeaseVendorIdOk returns a tuple with the LeaseVendorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseVendorId

`func (o *DataInnerDhcpLeaseData) SetLeaseVendorId(v string)`

SetLeaseVendorId sets LeaseVendorId field to given value.

### HasLeaseVendorId

`func (o *DataInnerDhcpLeaseData) HasLeaseVendorId() bool`

HasLeaseVendorId returns a boolean if a field has been set.

### GetRangeClassName

`func (o *DataInnerDhcpLeaseData) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *DataInnerDhcpLeaseData) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *DataInnerDhcpLeaseData) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *DataInnerDhcpLeaseData) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *DataInnerDhcpLeaseData) GetRangeClassParameters() []ApiClassParameterOutputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *DataInnerDhcpLeaseData) GetRangeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *DataInnerDhcpLeaseData) SetRangeClassParameters(v []ApiClassParameterOutputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *DataInnerDhcpLeaseData) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetRangeEndAddr

`func (o *DataInnerDhcpLeaseData) GetRangeEndAddr() string`

GetRangeEndAddr returns the RangeEndAddr field if non-nil, zero value otherwise.

### GetRangeEndAddrOk

`func (o *DataInnerDhcpLeaseData) GetRangeEndAddrOk() (*string, bool)`

GetRangeEndAddrOk returns a tuple with the RangeEndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndAddr

`func (o *DataInnerDhcpLeaseData) SetRangeEndAddr(v string)`

SetRangeEndAddr sets RangeEndAddr field to given value.

### HasRangeEndAddr

`func (o *DataInnerDhcpLeaseData) HasRangeEndAddr() bool`

HasRangeEndAddr returns a boolean if a field has been set.

### GetRangeFailoverName

`func (o *DataInnerDhcpLeaseData) GetRangeFailoverName() string`

GetRangeFailoverName returns the RangeFailoverName field if non-nil, zero value otherwise.

### GetRangeFailoverNameOk

`func (o *DataInnerDhcpLeaseData) GetRangeFailoverNameOk() (*string, bool)`

GetRangeFailoverNameOk returns a tuple with the RangeFailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeFailoverName

`func (o *DataInnerDhcpLeaseData) SetRangeFailoverName(v string)`

SetRangeFailoverName sets RangeFailoverName field to given value.

### HasRangeFailoverName

`func (o *DataInnerDhcpLeaseData) HasRangeFailoverName() bool`

HasRangeFailoverName returns a boolean if a field has been set.

### GetRangeId

`func (o *DataInnerDhcpLeaseData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *DataInnerDhcpLeaseData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *DataInnerDhcpLeaseData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *DataInnerDhcpLeaseData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *DataInnerDhcpLeaseData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *DataInnerDhcpLeaseData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *DataInnerDhcpLeaseData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *DataInnerDhcpLeaseData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetRangeStartAddr

`func (o *DataInnerDhcpLeaseData) GetRangeStartAddr() string`

GetRangeStartAddr returns the RangeStartAddr field if non-nil, zero value otherwise.

### GetRangeStartAddrOk

`func (o *DataInnerDhcpLeaseData) GetRangeStartAddrOk() (*string, bool)`

GetRangeStartAddrOk returns a tuple with the RangeStartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartAddr

`func (o *DataInnerDhcpLeaseData) SetRangeStartAddr(v string)`

SetRangeStartAddr sets RangeStartAddr field to given value.

### HasRangeStartAddr

`func (o *DataInnerDhcpLeaseData) HasRangeStartAddr() bool`

HasRangeStartAddr returns a boolean if a field has been set.

### GetScopeClassName

`func (o *DataInnerDhcpLeaseData) GetScopeClassName() string`

GetScopeClassName returns the ScopeClassName field if non-nil, zero value otherwise.

### GetScopeClassNameOk

`func (o *DataInnerDhcpLeaseData) GetScopeClassNameOk() (*string, bool)`

GetScopeClassNameOk returns a tuple with the ScopeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassName

`func (o *DataInnerDhcpLeaseData) SetScopeClassName(v string)`

SetScopeClassName sets ScopeClassName field to given value.

### HasScopeClassName

`func (o *DataInnerDhcpLeaseData) HasScopeClassName() bool`

HasScopeClassName returns a boolean if a field has been set.

### GetScopeId

`func (o *DataInnerDhcpLeaseData) GetScopeId() string`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DataInnerDhcpLeaseData) GetScopeIdOk() (*string, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DataInnerDhcpLeaseData) SetScopeId(v string)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *DataInnerDhcpLeaseData) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetScopeName

`func (o *DataInnerDhcpLeaseData) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DataInnerDhcpLeaseData) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DataInnerDhcpLeaseData) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DataInnerDhcpLeaseData) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeNetAddr

`func (o *DataInnerDhcpLeaseData) GetScopeNetAddr() string`

GetScopeNetAddr returns the ScopeNetAddr field if non-nil, zero value otherwise.

### GetScopeNetAddrOk

`func (o *DataInnerDhcpLeaseData) GetScopeNetAddrOk() (*string, bool)`

GetScopeNetAddrOk returns a tuple with the ScopeNetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetAddr

`func (o *DataInnerDhcpLeaseData) SetScopeNetAddr(v string)`

SetScopeNetAddr sets ScopeNetAddr field to given value.

### HasScopeNetAddr

`func (o *DataInnerDhcpLeaseData) HasScopeNetAddr() bool`

HasScopeNetAddr returns a boolean if a field has been set.

### GetScopeSize

`func (o *DataInnerDhcpLeaseData) GetScopeSize() string`

GetScopeSize returns the ScopeSize field if non-nil, zero value otherwise.

### GetScopeSizeOk

`func (o *DataInnerDhcpLeaseData) GetScopeSizeOk() (*string, bool)`

GetScopeSizeOk returns a tuple with the ScopeSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSize

`func (o *DataInnerDhcpLeaseData) SetScopeSize(v string)`

SetScopeSize sets ScopeSize field to given value.

### HasScopeSize

`func (o *DataInnerDhcpLeaseData) HasScopeSize() bool`

HasScopeSize returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DataInnerDhcpLeaseData) GetSharednetworkId() string`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DataInnerDhcpLeaseData) GetSharednetworkIdOk() (*string, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DataInnerDhcpLeaseData) SetSharednetworkId(v string)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DataInnerDhcpLeaseData) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DataInnerDhcpLeaseData) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DataInnerDhcpLeaseData) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DataInnerDhcpLeaseData) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DataInnerDhcpLeaseData) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DataInnerDhcpLeaseData) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DataInnerDhcpLeaseData) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DataInnerDhcpLeaseData) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DataInnerDhcpLeaseData) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerAddr6

`func (o *DataInnerDhcpLeaseData) GetServerAddr6() string`

GetServerAddr6 returns the ServerAddr6 field if non-nil, zero value otherwise.

### GetServerAddr6Ok

`func (o *DataInnerDhcpLeaseData) GetServerAddr6Ok() (*string, bool)`

GetServerAddr6Ok returns a tuple with the ServerAddr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr6

`func (o *DataInnerDhcpLeaseData) SetServerAddr6(v string)`

SetServerAddr6 sets ServerAddr6 field to given value.

### HasServerAddr6

`func (o *DataInnerDhcpLeaseData) HasServerAddr6() bool`

HasServerAddr6 returns a boolean if a field has been set.

### GetServerAddr

`func (o *DataInnerDhcpLeaseData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DataInnerDhcpLeaseData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DataInnerDhcpLeaseData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DataInnerDhcpLeaseData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetLeaseMacVendor

`func (o *DataInnerDhcpLeaseData) GetLeaseMacVendor() string`

GetLeaseMacVendor returns the LeaseMacVendor field if non-nil, zero value otherwise.

### GetLeaseMacVendorOk

`func (o *DataInnerDhcpLeaseData) GetLeaseMacVendorOk() (*string, bool)`

GetLeaseMacVendorOk returns a tuple with the LeaseMacVendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseMacVendor

`func (o *DataInnerDhcpLeaseData) SetLeaseMacVendor(v string)`

SetLeaseMacVendor sets LeaseMacVendor field to given value.

### HasLeaseMacVendor

`func (o *DataInnerDhcpLeaseData) HasLeaseMacVendor() bool`

HasLeaseMacVendor returns a boolean if a field has been set.

### GetLeaseMultistatus

`func (o *DataInnerDhcpLeaseData) GetLeaseMultistatus() string`

GetLeaseMultistatus returns the LeaseMultistatus field if non-nil, zero value otherwise.

### GetLeaseMultistatusOk

`func (o *DataInnerDhcpLeaseData) GetLeaseMultistatusOk() (*string, bool)`

GetLeaseMultistatusOk returns a tuple with the LeaseMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaseMultistatus

`func (o *DataInnerDhcpLeaseData) SetLeaseMultistatus(v string)`

SetLeaseMultistatus sets LeaseMultistatus field to given value.

### HasLeaseMultistatus

`func (o *DataInnerDhcpLeaseData) HasLeaseMultistatus() bool`

HasLeaseMultistatus returns a boolean if a field has been set.

### GetParameterRequestList

`func (o *DataInnerDhcpLeaseData) GetParameterRequestList() string`

GetParameterRequestList returns the ParameterRequestList field if non-nil, zero value otherwise.

### GetParameterRequestListOk

`func (o *DataInnerDhcpLeaseData) GetParameterRequestListOk() (*string, bool)`

GetParameterRequestListOk returns a tuple with the ParameterRequestList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameterRequestList

`func (o *DataInnerDhcpLeaseData) SetParameterRequestList(v string)`

SetParameterRequestList sets ParameterRequestList field to given value.

### HasParameterRequestList

`func (o *DataInnerDhcpLeaseData) HasParameterRequestList() bool`

HasParameterRequestList returns a boolean if a field has been set.

### GetPercent

`func (o *DataInnerDhcpLeaseData) GetPercent() string`

GetPercent returns the Percent field if non-nil, zero value otherwise.

### GetPercentOk

`func (o *DataInnerDhcpLeaseData) GetPercentOk() (*string, bool)`

GetPercentOk returns a tuple with the Percent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPercent

`func (o *DataInnerDhcpLeaseData) SetPercent(v string)`

SetPercent sets Percent field to given value.

### HasPercent

`func (o *DataInnerDhcpLeaseData) HasPercent() bool`

HasPercent returns a boolean if a field has been set.

### GetTimeToExpire

`func (o *DataInnerDhcpLeaseData) GetTimeToExpire() string`

GetTimeToExpire returns the TimeToExpire field if non-nil, zero value otherwise.

### GetTimeToExpireOk

`func (o *DataInnerDhcpLeaseData) GetTimeToExpireOk() (*string, bool)`

GetTimeToExpireOk returns a tuple with the TimeToExpire field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeToExpire

`func (o *DataInnerDhcpLeaseData) SetTimeToExpire(v string)`

SetTimeToExpire sets TimeToExpire field to given value.

### HasTimeToExpire

`func (o *DataInnerDhcpLeaseData) HasTimeToExpire() bool`

HasTimeToExpire returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpLeaseData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpLeaseData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpLeaseData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpLeaseData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpLeaseData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpLeaseData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpLeaseData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpLeaseData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


