# DnsRrDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RrDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**RrDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**RrDelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DNS server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerCloud** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerComment** | Pointer to **string** | The description of the DNS server the object belongs to. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DNS server the object belongs to.&lt;table&gt;&lt;caption&gt;dns_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DNS server or EfficientIP DNS Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msdaemon&lt;/td&gt;&lt;td &gt;Agentless Microsoft DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ans&lt;/td&gt;&lt;td &gt;Nominum DNS server (ANS)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;aws&lt;/td&gt;&lt;td &gt;Amazon Route 53 server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;other&lt;/td&gt;&lt;td &gt;Generic DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdns&lt;/td&gt;&lt;td &gt;EfficientIP DNS smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DNS server the object belongs to. | [optional] 
**ViewClassName** | Pointer to **string** | The name of the class applied to the DNS view the object belongs to, it can be preceded by the class directory. | [optional] 
**ViewClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ViewId** | Pointer to **string** | The database identifier (ID) of the DNS view the object belongs to. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view the object belongs to. | [optional] 
**ZoneClassName** | Pointer to **string** | The name of the class applied to the DNS zone the object belongs to, it can be preceded by the class directory. | [optional] 
**ZoneForwarders** | Pointer to **string** | The IP address(es) of the forwarder(s) associated with the DNS zone the resource record belongs to. It lists the DNS servers to which any unknown query on this zone should be sent, as follows: &lt;b&gt;&amp;lt;ip_address1&amp;gt;;&amp;lt;ip_address2&amp;gt;;...&lt;/b&gt; . | [optional] 
**ZoneId** | Pointer to **string** | The database identifier (ID) of the DNS zone the object belongs to. | [optional] 
**ZoneIsReverse** | Pointer to **string** | A way to determine if the DNS zone the resource record belongs to provides reverse resolution (&lt;b&gt;1&lt;/b&gt;) or direct/name resolution (&lt;b&gt;0&lt;/b&gt;), | [optional] 
**ZoneIsRpz** | Pointer to **string** | The RPZ status of the DNS zone the resource record belongs to. &lt;b&gt;1&lt;/b&gt; indicates that the DNS zone the record belongs to is a Response Policy Zone. | [optional] 
**ZoneMasters** | Pointer to **string** | For resource records in slave DNS zones, the IP address of the DNS server and, if relevant, the name of the DNS view that contain the master DNS zone, as follows: &lt;b&gt;&amp;lt;ip_addr&amp;gt;;&lt;/b&gt; or &lt;b&gt;&amp;lt;ip_addr&amp;gt; key &amp;lt;dnsview_name&amp;gt;;&lt;/b&gt; . | [optional] 
**ZoneName** | Pointer to **string** | The name of the DNS zone the object belongs to. | [optional] 
**ZoneNameUtf** | Pointer to **string** | The name of the DNS zone the resource record belongs to, in UTF-8 format. | [optional] 
**ZoneSpaceName** | Pointer to **string** | The name of the space associated with the DNS zone the RR belongs to. | [optional] 
**ZoneSortZone** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ZoneType** | Pointer to **string** | The type of the DNS zone the object belongs to, either &lt;b&gt;master&lt;/b&gt;, &lt;b&gt;slave&lt;/b&gt;, &lt;b&gt;forward&lt;/b&gt;, &lt;b&gt;stub&lt;/b&gt;, &lt;b&gt;hint&lt;/b&gt; or &lt;b&gt;delegation-only&lt;/b&gt;. | [optional] 
**RrMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**RrAllValue** | Pointer to **string** | The concatenated values of the DNS resource record, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;, &amp;lt;value2&amp;gt;, &amp;lt;value3&amp;gt;, &amp;lt;value4&amp;gt;, &amp;lt;value5&amp;gt;, &amp;lt;value6&amp;gt;, &amp;lt;value7&amp;gt;&lt;/b&gt;. | [optional] 
**RrClassName** | Pointer to **string** | TODO:dns_rr_list.output.rr_class_name | [optional] 
**RrClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DNS rr the object belongs to and their value | [optional] 
**RrFullName** | Pointer to **string** | The full name of the DNS resource record. | [optional] 
**RrFullNameUtf** | Pointer to **string** | The name of the DNS resource record in UTF-8 format. | [optional] 
**RrGlue** | Pointer to **string** | The shortname of the DNS resource record. | [optional] 
**RrId** | Pointer to **string** | The database identifier (ID) of the DNS resource record. | [optional] 
**RrLastUpdateTime** | Pointer to **string** | TODO:dns_rr_list.output.rr_last_update_time | [optional] 
**RrNameIp4Addr** | Pointer to **string** | Internal use. Not documented. | [optional] 
**RrNameAddressAddr** | Pointer to **string** | Internal use. Not documented. | [optional] 
**RrType** | Pointer to **string** | The type of the DNS resource record.&lt;table&gt;&lt;caption&gt;rr_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Value&lt;/th&gt;&lt;th&gt;Record type description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;SOA&lt;/td&gt;&lt;td &gt;Start of Authority. Defines the zone name, an email contact and various time and refresh values applicable to the zone. It is automatically generated upon creation of a zone and cannot be added manually.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NS&lt;/td&gt;&lt;td &gt;Name Server. Defines the authoritative name server(s) for the domain (defined by the SOA record) or the subdomain. The NS record that indicates which server has authority over a zone is automatically generated upon the creation of a zone, once the server has been synchronized.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;A&lt;/td&gt;&lt;td &gt;IPv4 Address. An IPv4 address for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;PTR&lt;/td&gt;&lt;td &gt;Pointer Record. Address Resolution, from an IP address (IPv4 or IPv6) to a host. Used in reverse mapping.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;AAAA&lt;/td&gt;&lt;td &gt;IPv6 Address. An IPv6 address for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;CNAME&lt;/td&gt;&lt;td &gt;Canonical Name. An alias name for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;MX&lt;/td&gt;&lt;td &gt;Mail Exchange. The mail server/exchanger that services this zone.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;SRV&lt;/td&gt;&lt;td &gt;Services record. Defines services available in the zone, for example, LDAP, HTTP, etc...&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DNAME&lt;/td&gt;&lt;td &gt;Delegation of Reverse Names. Delegation of reverse addresses primarily in IPv6. (Deprecated, use the CNAME RR instead)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;TXT&lt;/td&gt;&lt;td &gt;Text. Information associated with a name.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DS&lt;/td&gt;&lt;td &gt;Delegation Signer, a DNSSEC related RR used to verify the validity of the ZSK of a subdomain.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DNSKEY&lt;/td&gt;&lt;td &gt;DNS Key. It contains the public cryptographic key used to sign the zone with DNSSEC.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;65534&lt;/td&gt;&lt;td &gt;A private type record automatically added to the zone once it is signed with DNSSEC.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;HINFO&lt;/td&gt;&lt;td &gt;System Information. Information about a host: hardware type and operating system description.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;MINFO&lt;/td&gt;&lt;td &gt;Mailbox mail list Information. Defines the mail administrator for a mail list and optionally a mailbox to receive error messages relating to the mail list.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;AFSDB&lt;/td&gt;&lt;td &gt;AFS Database. Location of the AFS servers.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;WKS&lt;/td&gt;&lt;td &gt;Well-Known Service. Defines the services and protocols supported by a host. (Deprecated, use the SRV RR instead)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NAPTR&lt;/td&gt;&lt;td &gt;Naming Authority Pointer Record. General purpose definition of rule set to be used by applications e.g. VoIP.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NSAP&lt;/td&gt;&lt;td &gt;Network Service Access Point. Defines record (equivalent of an A record) maps a host name to an endpoint address.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**RrValueIp4Addr** | Pointer to **string** | Internal use. Not documented. | [optional] 
**RrValueAddressAddr** | Pointer to **string** | Internal use. Not documented. | [optional] 
**RrTtl** | Pointer to **string** | The time to live of the DNS resource record, in seconds. | [optional] 
**RrValue1** | Pointer to **string** | The first or only value required for the DNS resource record, as detailed in the service description. | [optional] 
**RrValue2** | Pointer to **string** | The second value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue3** | Pointer to **string** | The third value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue4** | Pointer to **string** | The fourth value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue5** | Pointer to **string** | The fifth value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue6** | Pointer to **string** | The sixth value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue7** | Pointer to **string** | The seventh value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DNS smart architecture managing the DNS server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DNS smart architecture managing the DNS server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDnsRrDataData

`func NewDnsRrDataData() *DnsRrDataData`

NewDnsRrDataData instantiates a new DnsRrDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRrDataDataWithDefaults

`func NewDnsRrDataDataWithDefaults() *DnsRrDataData`

NewDnsRrDataDataWithDefaults instantiates a new DnsRrDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRrDelayedCreateTime

`func (o *DnsRrDataData) GetRrDelayedCreateTime() string`

GetRrDelayedCreateTime returns the RrDelayedCreateTime field if non-nil, zero value otherwise.

### GetRrDelayedCreateTimeOk

`func (o *DnsRrDataData) GetRrDelayedCreateTimeOk() (*string, bool)`

GetRrDelayedCreateTimeOk returns a tuple with the RrDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrDelayedCreateTime

`func (o *DnsRrDataData) SetRrDelayedCreateTime(v string)`

SetRrDelayedCreateTime sets RrDelayedCreateTime field to given value.

### HasRrDelayedCreateTime

`func (o *DnsRrDataData) HasRrDelayedCreateTime() bool`

HasRrDelayedCreateTime returns a boolean if a field has been set.

### GetRrDelayedDeleteTime

`func (o *DnsRrDataData) GetRrDelayedDeleteTime() string`

GetRrDelayedDeleteTime returns the RrDelayedDeleteTime field if non-nil, zero value otherwise.

### GetRrDelayedDeleteTimeOk

`func (o *DnsRrDataData) GetRrDelayedDeleteTimeOk() (*string, bool)`

GetRrDelayedDeleteTimeOk returns a tuple with the RrDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrDelayedDeleteTime

`func (o *DnsRrDataData) SetRrDelayedDeleteTime(v string)`

SetRrDelayedDeleteTime sets RrDelayedDeleteTime field to given value.

### HasRrDelayedDeleteTime

`func (o *DnsRrDataData) HasRrDelayedDeleteTime() bool`

HasRrDelayedDeleteTime returns a boolean if a field has been set.

### GetRrDelayedTime

`func (o *DnsRrDataData) GetRrDelayedTime() string`

GetRrDelayedTime returns the RrDelayedTime field if non-nil, zero value otherwise.

### GetRrDelayedTimeOk

`func (o *DnsRrDataData) GetRrDelayedTimeOk() (*string, bool)`

GetRrDelayedTimeOk returns a tuple with the RrDelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrDelayedTime

`func (o *DnsRrDataData) SetRrDelayedTime(v string)`

SetRrDelayedTime sets RrDelayedTime field to given value.

### HasRrDelayedTime

`func (o *DnsRrDataData) HasRrDelayedTime() bool`

HasRrDelayedTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DnsRrDataData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DnsRrDataData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DnsRrDataData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DnsRrDataData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerCloud

`func (o *DnsRrDataData) GetServerCloud() string`

GetServerCloud returns the ServerCloud field if non-nil, zero value otherwise.

### GetServerCloudOk

`func (o *DnsRrDataData) GetServerCloudOk() (*string, bool)`

GetServerCloudOk returns a tuple with the ServerCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCloud

`func (o *DnsRrDataData) SetServerCloud(v string)`

SetServerCloud sets ServerCloud field to given value.

### HasServerCloud

`func (o *DnsRrDataData) HasServerCloud() bool`

HasServerCloud returns a boolean if a field has been set.

### GetServerComment

`func (o *DnsRrDataData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DnsRrDataData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DnsRrDataData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DnsRrDataData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerId

`func (o *DnsRrDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsRrDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsRrDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsRrDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsRrDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsRrDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsRrDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsRrDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DnsRrDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DnsRrDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DnsRrDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DnsRrDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DnsRrDataData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DnsRrDataData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DnsRrDataData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DnsRrDataData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetViewClassName

`func (o *DnsRrDataData) GetViewClassName() string`

GetViewClassName returns the ViewClassName field if non-nil, zero value otherwise.

### GetViewClassNameOk

`func (o *DnsRrDataData) GetViewClassNameOk() (*string, bool)`

GetViewClassNameOk returns a tuple with the ViewClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassName

`func (o *DnsRrDataData) SetViewClassName(v string)`

SetViewClassName sets ViewClassName field to given value.

### HasViewClassName

`func (o *DnsRrDataData) HasViewClassName() bool`

HasViewClassName returns a boolean if a field has been set.

### GetViewClassParameters

`func (o *DnsRrDataData) GetViewClassParameters() []ApiClassParameterOutputEntry`

GetViewClassParameters returns the ViewClassParameters field if non-nil, zero value otherwise.

### GetViewClassParametersOk

`func (o *DnsRrDataData) GetViewClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetViewClassParametersOk returns a tuple with the ViewClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassParameters

`func (o *DnsRrDataData) SetViewClassParameters(v []ApiClassParameterOutputEntry)`

SetViewClassParameters sets ViewClassParameters field to given value.

### HasViewClassParameters

`func (o *DnsRrDataData) HasViewClassParameters() bool`

HasViewClassParameters returns a boolean if a field has been set.

### GetViewId

`func (o *DnsRrDataData) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DnsRrDataData) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DnsRrDataData) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DnsRrDataData) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewName

`func (o *DnsRrDataData) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DnsRrDataData) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DnsRrDataData) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DnsRrDataData) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetZoneClassName

`func (o *DnsRrDataData) GetZoneClassName() string`

GetZoneClassName returns the ZoneClassName field if non-nil, zero value otherwise.

### GetZoneClassNameOk

`func (o *DnsRrDataData) GetZoneClassNameOk() (*string, bool)`

GetZoneClassNameOk returns a tuple with the ZoneClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneClassName

`func (o *DnsRrDataData) SetZoneClassName(v string)`

SetZoneClassName sets ZoneClassName field to given value.

### HasZoneClassName

`func (o *DnsRrDataData) HasZoneClassName() bool`

HasZoneClassName returns a boolean if a field has been set.

### GetZoneForwarders

`func (o *DnsRrDataData) GetZoneForwarders() string`

GetZoneForwarders returns the ZoneForwarders field if non-nil, zero value otherwise.

### GetZoneForwardersOk

`func (o *DnsRrDataData) GetZoneForwardersOk() (*string, bool)`

GetZoneForwardersOk returns a tuple with the ZoneForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneForwarders

`func (o *DnsRrDataData) SetZoneForwarders(v string)`

SetZoneForwarders sets ZoneForwarders field to given value.

### HasZoneForwarders

`func (o *DnsRrDataData) HasZoneForwarders() bool`

HasZoneForwarders returns a boolean if a field has been set.

### GetZoneId

`func (o *DnsRrDataData) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnsRrDataData) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnsRrDataData) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DnsRrDataData) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneIsReverse

`func (o *DnsRrDataData) GetZoneIsReverse() string`

GetZoneIsReverse returns the ZoneIsReverse field if non-nil, zero value otherwise.

### GetZoneIsReverseOk

`func (o *DnsRrDataData) GetZoneIsReverseOk() (*string, bool)`

GetZoneIsReverseOk returns a tuple with the ZoneIsReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIsReverse

`func (o *DnsRrDataData) SetZoneIsReverse(v string)`

SetZoneIsReverse sets ZoneIsReverse field to given value.

### HasZoneIsReverse

`func (o *DnsRrDataData) HasZoneIsReverse() bool`

HasZoneIsReverse returns a boolean if a field has been set.

### GetZoneIsRpz

`func (o *DnsRrDataData) GetZoneIsRpz() string`

GetZoneIsRpz returns the ZoneIsRpz field if non-nil, zero value otherwise.

### GetZoneIsRpzOk

`func (o *DnsRrDataData) GetZoneIsRpzOk() (*string, bool)`

GetZoneIsRpzOk returns a tuple with the ZoneIsRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIsRpz

`func (o *DnsRrDataData) SetZoneIsRpz(v string)`

SetZoneIsRpz sets ZoneIsRpz field to given value.

### HasZoneIsRpz

`func (o *DnsRrDataData) HasZoneIsRpz() bool`

HasZoneIsRpz returns a boolean if a field has been set.

### GetZoneMasters

`func (o *DnsRrDataData) GetZoneMasters() string`

GetZoneMasters returns the ZoneMasters field if non-nil, zero value otherwise.

### GetZoneMastersOk

`func (o *DnsRrDataData) GetZoneMastersOk() (*string, bool)`

GetZoneMastersOk returns a tuple with the ZoneMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMasters

`func (o *DnsRrDataData) SetZoneMasters(v string)`

SetZoneMasters sets ZoneMasters field to given value.

### HasZoneMasters

`func (o *DnsRrDataData) HasZoneMasters() bool`

HasZoneMasters returns a boolean if a field has been set.

### GetZoneName

`func (o *DnsRrDataData) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DnsRrDataData) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DnsRrDataData) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *DnsRrDataData) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneNameUtf

`func (o *DnsRrDataData) GetZoneNameUtf() string`

GetZoneNameUtf returns the ZoneNameUtf field if non-nil, zero value otherwise.

### GetZoneNameUtfOk

`func (o *DnsRrDataData) GetZoneNameUtfOk() (*string, bool)`

GetZoneNameUtfOk returns a tuple with the ZoneNameUtf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNameUtf

`func (o *DnsRrDataData) SetZoneNameUtf(v string)`

SetZoneNameUtf sets ZoneNameUtf field to given value.

### HasZoneNameUtf

`func (o *DnsRrDataData) HasZoneNameUtf() bool`

HasZoneNameUtf returns a boolean if a field has been set.

### GetZoneSpaceName

`func (o *DnsRrDataData) GetZoneSpaceName() string`

GetZoneSpaceName returns the ZoneSpaceName field if non-nil, zero value otherwise.

### GetZoneSpaceNameOk

`func (o *DnsRrDataData) GetZoneSpaceNameOk() (*string, bool)`

GetZoneSpaceNameOk returns a tuple with the ZoneSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSpaceName

`func (o *DnsRrDataData) SetZoneSpaceName(v string)`

SetZoneSpaceName sets ZoneSpaceName field to given value.

### HasZoneSpaceName

`func (o *DnsRrDataData) HasZoneSpaceName() bool`

HasZoneSpaceName returns a boolean if a field has been set.

### GetZoneSortZone

`func (o *DnsRrDataData) GetZoneSortZone() string`

GetZoneSortZone returns the ZoneSortZone field if non-nil, zero value otherwise.

### GetZoneSortZoneOk

`func (o *DnsRrDataData) GetZoneSortZoneOk() (*string, bool)`

GetZoneSortZoneOk returns a tuple with the ZoneSortZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSortZone

`func (o *DnsRrDataData) SetZoneSortZone(v string)`

SetZoneSortZone sets ZoneSortZone field to given value.

### HasZoneSortZone

`func (o *DnsRrDataData) HasZoneSortZone() bool`

HasZoneSortZone returns a boolean if a field has been set.

### GetZoneType

`func (o *DnsRrDataData) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *DnsRrDataData) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *DnsRrDataData) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *DnsRrDataData) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetRrMultistatus

`func (o *DnsRrDataData) GetRrMultistatus() string`

GetRrMultistatus returns the RrMultistatus field if non-nil, zero value otherwise.

### GetRrMultistatusOk

`func (o *DnsRrDataData) GetRrMultistatusOk() (*string, bool)`

GetRrMultistatusOk returns a tuple with the RrMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrMultistatus

`func (o *DnsRrDataData) SetRrMultistatus(v string)`

SetRrMultistatus sets RrMultistatus field to given value.

### HasRrMultistatus

`func (o *DnsRrDataData) HasRrMultistatus() bool`

HasRrMultistatus returns a boolean if a field has been set.

### GetRrAllValue

`func (o *DnsRrDataData) GetRrAllValue() string`

GetRrAllValue returns the RrAllValue field if non-nil, zero value otherwise.

### GetRrAllValueOk

`func (o *DnsRrDataData) GetRrAllValueOk() (*string, bool)`

GetRrAllValueOk returns a tuple with the RrAllValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrAllValue

`func (o *DnsRrDataData) SetRrAllValue(v string)`

SetRrAllValue sets RrAllValue field to given value.

### HasRrAllValue

`func (o *DnsRrDataData) HasRrAllValue() bool`

HasRrAllValue returns a boolean if a field has been set.

### GetRrClassName

`func (o *DnsRrDataData) GetRrClassName() string`

GetRrClassName returns the RrClassName field if non-nil, zero value otherwise.

### GetRrClassNameOk

`func (o *DnsRrDataData) GetRrClassNameOk() (*string, bool)`

GetRrClassNameOk returns a tuple with the RrClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrClassName

`func (o *DnsRrDataData) SetRrClassName(v string)`

SetRrClassName sets RrClassName field to given value.

### HasRrClassName

`func (o *DnsRrDataData) HasRrClassName() bool`

HasRrClassName returns a boolean if a field has been set.

### GetRrClassParameters

`func (o *DnsRrDataData) GetRrClassParameters() []ApiClassParameterOutputEntry`

GetRrClassParameters returns the RrClassParameters field if non-nil, zero value otherwise.

### GetRrClassParametersOk

`func (o *DnsRrDataData) GetRrClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRrClassParametersOk returns a tuple with the RrClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrClassParameters

`func (o *DnsRrDataData) SetRrClassParameters(v []ApiClassParameterOutputEntry)`

SetRrClassParameters sets RrClassParameters field to given value.

### HasRrClassParameters

`func (o *DnsRrDataData) HasRrClassParameters() bool`

HasRrClassParameters returns a boolean if a field has been set.

### GetRrFullName

`func (o *DnsRrDataData) GetRrFullName() string`

GetRrFullName returns the RrFullName field if non-nil, zero value otherwise.

### GetRrFullNameOk

`func (o *DnsRrDataData) GetRrFullNameOk() (*string, bool)`

GetRrFullNameOk returns a tuple with the RrFullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrFullName

`func (o *DnsRrDataData) SetRrFullName(v string)`

SetRrFullName sets RrFullName field to given value.

### HasRrFullName

`func (o *DnsRrDataData) HasRrFullName() bool`

HasRrFullName returns a boolean if a field has been set.

### GetRrFullNameUtf

`func (o *DnsRrDataData) GetRrFullNameUtf() string`

GetRrFullNameUtf returns the RrFullNameUtf field if non-nil, zero value otherwise.

### GetRrFullNameUtfOk

`func (o *DnsRrDataData) GetRrFullNameUtfOk() (*string, bool)`

GetRrFullNameUtfOk returns a tuple with the RrFullNameUtf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrFullNameUtf

`func (o *DnsRrDataData) SetRrFullNameUtf(v string)`

SetRrFullNameUtf sets RrFullNameUtf field to given value.

### HasRrFullNameUtf

`func (o *DnsRrDataData) HasRrFullNameUtf() bool`

HasRrFullNameUtf returns a boolean if a field has been set.

### GetRrGlue

`func (o *DnsRrDataData) GetRrGlue() string`

GetRrGlue returns the RrGlue field if non-nil, zero value otherwise.

### GetRrGlueOk

`func (o *DnsRrDataData) GetRrGlueOk() (*string, bool)`

GetRrGlueOk returns a tuple with the RrGlue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrGlue

`func (o *DnsRrDataData) SetRrGlue(v string)`

SetRrGlue sets RrGlue field to given value.

### HasRrGlue

`func (o *DnsRrDataData) HasRrGlue() bool`

HasRrGlue returns a boolean if a field has been set.

### GetRrId

`func (o *DnsRrDataData) GetRrId() string`

GetRrId returns the RrId field if non-nil, zero value otherwise.

### GetRrIdOk

`func (o *DnsRrDataData) GetRrIdOk() (*string, bool)`

GetRrIdOk returns a tuple with the RrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrId

`func (o *DnsRrDataData) SetRrId(v string)`

SetRrId sets RrId field to given value.

### HasRrId

`func (o *DnsRrDataData) HasRrId() bool`

HasRrId returns a boolean if a field has been set.

### GetRrLastUpdateTime

`func (o *DnsRrDataData) GetRrLastUpdateTime() string`

GetRrLastUpdateTime returns the RrLastUpdateTime field if non-nil, zero value otherwise.

### GetRrLastUpdateTimeOk

`func (o *DnsRrDataData) GetRrLastUpdateTimeOk() (*string, bool)`

GetRrLastUpdateTimeOk returns a tuple with the RrLastUpdateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrLastUpdateTime

`func (o *DnsRrDataData) SetRrLastUpdateTime(v string)`

SetRrLastUpdateTime sets RrLastUpdateTime field to given value.

### HasRrLastUpdateTime

`func (o *DnsRrDataData) HasRrLastUpdateTime() bool`

HasRrLastUpdateTime returns a boolean if a field has been set.

### GetRrNameIp4Addr

`func (o *DnsRrDataData) GetRrNameIp4Addr() string`

GetRrNameIp4Addr returns the RrNameIp4Addr field if non-nil, zero value otherwise.

### GetRrNameIp4AddrOk

`func (o *DnsRrDataData) GetRrNameIp4AddrOk() (*string, bool)`

GetRrNameIp4AddrOk returns a tuple with the RrNameIp4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrNameIp4Addr

`func (o *DnsRrDataData) SetRrNameIp4Addr(v string)`

SetRrNameIp4Addr sets RrNameIp4Addr field to given value.

### HasRrNameIp4Addr

`func (o *DnsRrDataData) HasRrNameIp4Addr() bool`

HasRrNameIp4Addr returns a boolean if a field has been set.

### GetRrNameAddressAddr

`func (o *DnsRrDataData) GetRrNameAddressAddr() string`

GetRrNameAddressAddr returns the RrNameAddressAddr field if non-nil, zero value otherwise.

### GetRrNameAddressAddrOk

`func (o *DnsRrDataData) GetRrNameAddressAddrOk() (*string, bool)`

GetRrNameAddressAddrOk returns a tuple with the RrNameAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrNameAddressAddr

`func (o *DnsRrDataData) SetRrNameAddressAddr(v string)`

SetRrNameAddressAddr sets RrNameAddressAddr field to given value.

### HasRrNameAddressAddr

`func (o *DnsRrDataData) HasRrNameAddressAddr() bool`

HasRrNameAddressAddr returns a boolean if a field has been set.

### GetRrType

`func (o *DnsRrDataData) GetRrType() string`

GetRrType returns the RrType field if non-nil, zero value otherwise.

### GetRrTypeOk

`func (o *DnsRrDataData) GetRrTypeOk() (*string, bool)`

GetRrTypeOk returns a tuple with the RrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrType

`func (o *DnsRrDataData) SetRrType(v string)`

SetRrType sets RrType field to given value.

### HasRrType

`func (o *DnsRrDataData) HasRrType() bool`

HasRrType returns a boolean if a field has been set.

### GetRrValueIp4Addr

`func (o *DnsRrDataData) GetRrValueIp4Addr() string`

GetRrValueIp4Addr returns the RrValueIp4Addr field if non-nil, zero value otherwise.

### GetRrValueIp4AddrOk

`func (o *DnsRrDataData) GetRrValueIp4AddrOk() (*string, bool)`

GetRrValueIp4AddrOk returns a tuple with the RrValueIp4Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValueIp4Addr

`func (o *DnsRrDataData) SetRrValueIp4Addr(v string)`

SetRrValueIp4Addr sets RrValueIp4Addr field to given value.

### HasRrValueIp4Addr

`func (o *DnsRrDataData) HasRrValueIp4Addr() bool`

HasRrValueIp4Addr returns a boolean if a field has been set.

### GetRrValueAddressAddr

`func (o *DnsRrDataData) GetRrValueAddressAddr() string`

GetRrValueAddressAddr returns the RrValueAddressAddr field if non-nil, zero value otherwise.

### GetRrValueAddressAddrOk

`func (o *DnsRrDataData) GetRrValueAddressAddrOk() (*string, bool)`

GetRrValueAddressAddrOk returns a tuple with the RrValueAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValueAddressAddr

`func (o *DnsRrDataData) SetRrValueAddressAddr(v string)`

SetRrValueAddressAddr sets RrValueAddressAddr field to given value.

### HasRrValueAddressAddr

`func (o *DnsRrDataData) HasRrValueAddressAddr() bool`

HasRrValueAddressAddr returns a boolean if a field has been set.

### GetRrTtl

`func (o *DnsRrDataData) GetRrTtl() string`

GetRrTtl returns the RrTtl field if non-nil, zero value otherwise.

### GetRrTtlOk

`func (o *DnsRrDataData) GetRrTtlOk() (*string, bool)`

GetRrTtlOk returns a tuple with the RrTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrTtl

`func (o *DnsRrDataData) SetRrTtl(v string)`

SetRrTtl sets RrTtl field to given value.

### HasRrTtl

`func (o *DnsRrDataData) HasRrTtl() bool`

HasRrTtl returns a boolean if a field has been set.

### GetRrValue1

`func (o *DnsRrDataData) GetRrValue1() string`

GetRrValue1 returns the RrValue1 field if non-nil, zero value otherwise.

### GetRrValue1Ok

`func (o *DnsRrDataData) GetRrValue1Ok() (*string, bool)`

GetRrValue1Ok returns a tuple with the RrValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue1

`func (o *DnsRrDataData) SetRrValue1(v string)`

SetRrValue1 sets RrValue1 field to given value.

### HasRrValue1

`func (o *DnsRrDataData) HasRrValue1() bool`

HasRrValue1 returns a boolean if a field has been set.

### GetRrValue2

`func (o *DnsRrDataData) GetRrValue2() string`

GetRrValue2 returns the RrValue2 field if non-nil, zero value otherwise.

### GetRrValue2Ok

`func (o *DnsRrDataData) GetRrValue2Ok() (*string, bool)`

GetRrValue2Ok returns a tuple with the RrValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue2

`func (o *DnsRrDataData) SetRrValue2(v string)`

SetRrValue2 sets RrValue2 field to given value.

### HasRrValue2

`func (o *DnsRrDataData) HasRrValue2() bool`

HasRrValue2 returns a boolean if a field has been set.

### GetRrValue3

`func (o *DnsRrDataData) GetRrValue3() string`

GetRrValue3 returns the RrValue3 field if non-nil, zero value otherwise.

### GetRrValue3Ok

`func (o *DnsRrDataData) GetRrValue3Ok() (*string, bool)`

GetRrValue3Ok returns a tuple with the RrValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue3

`func (o *DnsRrDataData) SetRrValue3(v string)`

SetRrValue3 sets RrValue3 field to given value.

### HasRrValue3

`func (o *DnsRrDataData) HasRrValue3() bool`

HasRrValue3 returns a boolean if a field has been set.

### GetRrValue4

`func (o *DnsRrDataData) GetRrValue4() string`

GetRrValue4 returns the RrValue4 field if non-nil, zero value otherwise.

### GetRrValue4Ok

`func (o *DnsRrDataData) GetRrValue4Ok() (*string, bool)`

GetRrValue4Ok returns a tuple with the RrValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue4

`func (o *DnsRrDataData) SetRrValue4(v string)`

SetRrValue4 sets RrValue4 field to given value.

### HasRrValue4

`func (o *DnsRrDataData) HasRrValue4() bool`

HasRrValue4 returns a boolean if a field has been set.

### GetRrValue5

`func (o *DnsRrDataData) GetRrValue5() string`

GetRrValue5 returns the RrValue5 field if non-nil, zero value otherwise.

### GetRrValue5Ok

`func (o *DnsRrDataData) GetRrValue5Ok() (*string, bool)`

GetRrValue5Ok returns a tuple with the RrValue5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue5

`func (o *DnsRrDataData) SetRrValue5(v string)`

SetRrValue5 sets RrValue5 field to given value.

### HasRrValue5

`func (o *DnsRrDataData) HasRrValue5() bool`

HasRrValue5 returns a boolean if a field has been set.

### GetRrValue6

`func (o *DnsRrDataData) GetRrValue6() string`

GetRrValue6 returns the RrValue6 field if non-nil, zero value otherwise.

### GetRrValue6Ok

`func (o *DnsRrDataData) GetRrValue6Ok() (*string, bool)`

GetRrValue6Ok returns a tuple with the RrValue6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue6

`func (o *DnsRrDataData) SetRrValue6(v string)`

SetRrValue6 sets RrValue6 field to given value.

### HasRrValue6

`func (o *DnsRrDataData) HasRrValue6() bool`

HasRrValue6 returns a boolean if a field has been set.

### GetRrValue7

`func (o *DnsRrDataData) GetRrValue7() string`

GetRrValue7 returns the RrValue7 field if non-nil, zero value otherwise.

### GetRrValue7Ok

`func (o *DnsRrDataData) GetRrValue7Ok() (*string, bool)`

GetRrValue7Ok returns a tuple with the RrValue7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue7

`func (o *DnsRrDataData) SetRrValue7(v string)`

SetRrValue7 sets RrValue7 field to given value.

### HasRrValue7

`func (o *DnsRrDataData) HasRrValue7() bool`

HasRrValue7 returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DnsRrDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DnsRrDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DnsRrDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DnsRrDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DnsRrDataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DnsRrDataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DnsRrDataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DnsRrDataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


