# DnsServerDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdDomain** | Pointer to **string** | TODO:dns_server_list.output.ad_domain | [optional] 
**AdUser** | Pointer to **string** | TODO:dns_server_list.output.ad_user | [optional] 
**ServerAwsKeyid** | Pointer to **string** | The AWS access key identifier (ID) of the DNS server. | [optional] 
**AzGroup** | Pointer to **string** | TODO:dns_server_list.output.az_group | [optional] 
**AzKeyid** | Pointer to **string** | TODO:dns_server_list.output.az_keyid | [optional] 
**AzSubscriptionid** | Pointer to **string** | TODO:dns_server_list.output.az_subscriptionid | [optional] 
**AzTenantid** | Pointer to **string** | TODO:dns_server_list.output.az_tenantid | [optional] 
**ServerAllowQuery** | Pointer to **string** | The ACL values associated with the allow-query configuration of the DNS server, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ServerAllowQueryCache** | Pointer to **string** | The ACL values associated with the allow-query-cache configuration of the DNS server, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ServerAllowRecursion** | Pointer to **string** | The ACL values associated with the allow-recursion configuration of the DNS server, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ServerAllowTransfer** | Pointer to **string** | The ACL values associated with the allow-transfer configuration of the DNS server, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ServerAlsoNotify** | Pointer to **string** | The IP address and port of the DNS server managing the smart architecture. If the parameter &lt;b&gt;dns_notify&lt;/b&gt; is set to &lt;b&gt;yes&lt;/b&gt; or &lt;b&gt;explicit&lt;/b&gt;, the server specified is instantly notified of any slave zones updates. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DNS server, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ServerCloud** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerComment** | Pointer to **string** | The description of the DNS server. | [optional] 
**ServerForceHybrid** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerForward** | Pointer to **string** | The forwarding mode of the DNS server. No value indicates that the forwarding is disabled: &lt;table&gt;&lt;caption&gt;dns_forward possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;first&lt;/td&gt;&lt;td &gt;The server sends the queries to the forwarder(s). If no answer is returned, it attempts to answer the queries on its own.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;only&lt;/td&gt;&lt;td &gt;The server only forwards the queries to the forwarder(s). Required by some reverse forward zones (e.g., in the case of private addresses).&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerForwarders** | Pointer to **string** | The IP address(es) of the forwarder(s) associated with the DNS server. It lists the DNS servers to which any unknown zone should be sent, as follows: &lt;b&gt;&amp;lt;ip_address1&amp;gt;;&amp;lt;ip_address2&amp;gt;;...&lt;/b&gt; . | [optional] 
**ServerHybrid** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. | [optional] 
**ServerKeyName** | Pointer to **string** | The name of the DNS TSIG key associated with the DNS server. | [optional] 
**ServerKeyProto** | Pointer to **string** | The encryption protocol of the TSIG key associated with the DNS server. | [optional] 
**ServerKeyValue** | Pointer to **string** | The value of the TSIG key associated with the DNS server. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server. | [optional] 
**ServerNotify** | Pointer to **string** | The notify status of the DNS server:&lt;table&gt;&lt;caption&gt;dns_notify possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;No notify message is sent when changes are performed in the master zones.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;The notify messages are sent to the target of the NS records of the master zone. They are also sent to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;explicit&lt;/td&gt;&lt;td &gt;The notify messages are only sent to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerRecursion** | Pointer to **string** | The recursion status of the DNS server:&lt;table&gt;&lt;caption&gt;dns_recursion possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;The server only provides iterative query behavior - normally resulting in a referral. If the answer to the query already exists in the cache it will be returned whatever the value of this statement.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;The server always provides recursive query behavior if requested by the client.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerRole** | Pointer to **string** | The role of the DNS server in the smart architecture, either &lt;b&gt;master&lt;/b&gt;, &lt;b&gt;hidden-master&lt;/b&gt;, &lt;b&gt;pseudo-master&lt;/b&gt; or &lt;b&gt;slave&lt;/b&gt;. | [optional] 
**ServerState** | Pointer to **string** | The status of the DNS server:&lt;table&gt;&lt;caption&gt;dns_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ER&lt;/td&gt;&lt;td &gt;The license used in SOLIDserver is not compliant with the added server: the license is invalid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ES&lt;/td&gt;&lt;td &gt;The server configuration could not be parsed properly.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ET&lt;/td&gt;&lt;td &gt;The server does not answer anymore due to a scheduled configuration of the server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IC&lt;/td&gt;&lt;td &gt;The SSL credentials are invalid&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IP&lt;/td&gt;&lt;td &gt;The provided account does not have sufficient privileges to remotely manage the MS server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IR&lt;/td&gt;&lt;td &gt;SOLIDserver cannot resolve the AWS DNS service. The Amazon services are unreachable and the Amazon Route 53 server cannot be managed. Make sure that the DNS resolvers declared on the page  are valid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IS&lt;/td&gt;&lt;td &gt;There was a setting error during the server declaration. For instance, some settings were added to a server that does not support them or a smart architecture is not managing any physical server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IT&lt;/td&gt;&lt;td &gt;The server editions performed from the GUI are not pushed to the server because SOLIDserver time and date are incorrect. You must use the UTC system on the appliance, especially when managing Amazon Route 53 servers.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;LS&lt;/td&gt;&lt;td &gt;The server configuration is not viable.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;N&lt;/td&gt;&lt;td &gt;The server does not have a status as it has not synchronized yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;UE&lt;/td&gt;&lt;td &gt;An error occurred that SOLIDserver could not identify.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;Y&lt;/td&gt;&lt;td &gt;The server is operational.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerSynching** | Pointer to **string** | The synchronization status of the DNS server. &lt;b&gt;1&lt;/b&gt; indicates that the server is currently being synchronized. | [optional] 
**ServerType** | Pointer to **string** | The type of the DNS server: &lt;table&gt;&lt;caption&gt;dns_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DNS server or EfficientIP DNS Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msdaemon&lt;/td&gt;&lt;td &gt;Agentless Microsoft DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ans&lt;/td&gt;&lt;td &gt;Nominum DNS server (ANS)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;aws&lt;/td&gt;&lt;td &gt;Amazon Route 53 server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;other&lt;/td&gt;&lt;td &gt;Generic DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdns&lt;/td&gt;&lt;td &gt;EfficientIP DNS smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DNS server. | [optional] 
**ServerBlastEnabled** | Pointer to **string** | The status of the service DNS Guardian, either enabled (&lt;b&gt;1&lt;/b&gt;) or disabled (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**ServerBlastStatus** | Pointer to **string** | The status of the Guardian server, either OK (&lt;b&gt;1&lt;/b&gt;), Stopped (&lt;b&gt;2&lt;/b&gt;), Invalid Credentials (&lt;b&gt;4&lt;/b&gt;) or Timeout (&lt;b&gt;5&lt;/b&gt;). | [optional] 
**ServerGslbSupported** | Pointer to **string** | The license GSLB activation status. &lt;b&gt;1&lt;/b&gt; indicates your license includes GSLB and your appliance supports it. | [optional] 
**ServerGuardianGuiManagementSupported** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerGuardianSupported** | Pointer to **string** | The license Guardian activation status. &lt;b&gt;1&lt;/b&gt; indicates your license includes Guardian and your appliance supports its latest features. | [optional] 
**ServerDnssecValidation** | Pointer to **string** | The DNSSEC resolution status of the DNS server. &lt;b&gt;yes&lt;/b&gt; indicates it is enabled. | [optional] 
**ServerGssEnabled** | Pointer to **string** | The GSS-TSIG status of the DNS server. &lt;b&gt;1&lt;/b&gt; indicates that GSS-TSIG is enabled on the server. | [optional] 
**ServerGssKeytabId** | Pointer to **string** | The database identifier (ID) of the DNS GSS-TSIG keytab. | [optional] 
**ServerAddr** | Pointer to **string** | The IP address of the DNS server, in hexadecimal format. | [optional] 
**ServerIpmLogin** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIpmIsPackage** | Pointer to **string** | The DNS server package information. &lt;b&gt;Y&lt;/b&gt; for an EfficientIP DNS Package server, &lt;b&gt;N&lt;/b&gt; for an appliance or virtual machine, &lt;b&gt;U&lt;/b&gt; the package information is irrelevant. For servers with a &lt;b&gt;dns_type&lt;/b&gt; set to &lt;b&gt;ipm&lt;/b&gt;, &lt;b&gt;U&lt;/b&gt; indicates either EfficientIP DNS Packages or appliances/virtual machines. | [optional] 
**ServerIpmProtocol** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIpmType** | Pointer to **string** | The engine type of the DNS server: &lt;b&gt;named&lt;/b&gt; (BIND engine), &lt;b&gt;nsd&lt;/b&gt; (NSD engine) or &lt;b&gt;unbound&lt;/b&gt; (Unbound engine). | [optional] 
**ServerIsolated** | Pointer to **string** | A way to determine if the server can update any other module &lt;b&gt;(1)&lt;/b&gt;. | [optional] 
**ServerLdapDomain** | Pointer to **string** | The LDAP domain associated with the DNS server. | [optional] 
**ServerLdapUser** | Pointer to **string** | The LDAP login associated with the DNS server. | [optional] 
**ServerMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ServerQuerylogState** | Pointer to **string** | The DNS querylog status. &lt;b&gt;1&lt;/b&gt; indicates that the DNS server querylog is enabled. | [optional] 
**ReverseProxyConf** | Pointer to **string** | TODO:dns_server_list.output.reverse_proxy_conf | [optional] 
**ServerSnmpId** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatEnabled** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatNiceness** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatPeriod** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatTime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**TotalSmartMembers** | Pointer to **string** | The total number of servers managed by the DNS smart architecture. | [optional] 
**SmartArch** | Pointer to **string** | The type of the DNS smart architecture:&lt;table&gt;&lt;caption&gt;vdns_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;masterslave&lt;/td&gt;&lt;td &gt;Master/Slave&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;stealth&lt;/td&gt;&lt;td &gt;Stealth&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;multimaster&lt;/td&gt;&lt;td &gt;Multi-Master&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;Single-Server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;farm&lt;/td&gt;&lt;td &gt;Farm&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SmartMembersName** | Pointer to **string** | The list of the servers managed by the DNS smart architecture, as follows: &lt;b&gt;&amp;lt;dns_name&amp;gt;,&amp;lt;dns_name&amp;gt;,...&lt;/b&gt; . | [optional] 
**SmartParentArch** | Pointer to **string** | The type of the DNS smart architecture managing the DNS server. No value indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DNS smart architecture managing the DNS server. &lt;b&gt;0&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DNS smart architecture managing the DNS server. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartPublicNsList** | Pointer to **string** | The list of the published name servers associated with the DNS smart architecture, as follows: &lt;b&gt;&amp;lt;ns1&amp;gt;;&amp;lt;ns2&amp;gt;;...&lt;/b&gt; . | [optional] 
**ServerWindnsPort** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerWindnsProtocol** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerWindnsUseSsl** | Pointer to **string** | Internal use. Not documented. | [optional] 

## Methods

### NewDnsServerDataData

`func NewDnsServerDataData() *DnsServerDataData`

NewDnsServerDataData instantiates a new DnsServerDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsServerDataDataWithDefaults

`func NewDnsServerDataDataWithDefaults() *DnsServerDataData`

NewDnsServerDataDataWithDefaults instantiates a new DnsServerDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdDomain

`func (o *DnsServerDataData) GetAdDomain() string`

GetAdDomain returns the AdDomain field if non-nil, zero value otherwise.

### GetAdDomainOk

`func (o *DnsServerDataData) GetAdDomainOk() (*string, bool)`

GetAdDomainOk returns a tuple with the AdDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdDomain

`func (o *DnsServerDataData) SetAdDomain(v string)`

SetAdDomain sets AdDomain field to given value.

### HasAdDomain

`func (o *DnsServerDataData) HasAdDomain() bool`

HasAdDomain returns a boolean if a field has been set.

### GetAdUser

`func (o *DnsServerDataData) GetAdUser() string`

GetAdUser returns the AdUser field if non-nil, zero value otherwise.

### GetAdUserOk

`func (o *DnsServerDataData) GetAdUserOk() (*string, bool)`

GetAdUserOk returns a tuple with the AdUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdUser

`func (o *DnsServerDataData) SetAdUser(v string)`

SetAdUser sets AdUser field to given value.

### HasAdUser

`func (o *DnsServerDataData) HasAdUser() bool`

HasAdUser returns a boolean if a field has been set.

### GetServerAwsKeyid

`func (o *DnsServerDataData) GetServerAwsKeyid() string`

GetServerAwsKeyid returns the ServerAwsKeyid field if non-nil, zero value otherwise.

### GetServerAwsKeyidOk

`func (o *DnsServerDataData) GetServerAwsKeyidOk() (*string, bool)`

GetServerAwsKeyidOk returns a tuple with the ServerAwsKeyid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAwsKeyid

`func (o *DnsServerDataData) SetServerAwsKeyid(v string)`

SetServerAwsKeyid sets ServerAwsKeyid field to given value.

### HasServerAwsKeyid

`func (o *DnsServerDataData) HasServerAwsKeyid() bool`

HasServerAwsKeyid returns a boolean if a field has been set.

### GetAzGroup

`func (o *DnsServerDataData) GetAzGroup() string`

GetAzGroup returns the AzGroup field if non-nil, zero value otherwise.

### GetAzGroupOk

`func (o *DnsServerDataData) GetAzGroupOk() (*string, bool)`

GetAzGroupOk returns a tuple with the AzGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzGroup

`func (o *DnsServerDataData) SetAzGroup(v string)`

SetAzGroup sets AzGroup field to given value.

### HasAzGroup

`func (o *DnsServerDataData) HasAzGroup() bool`

HasAzGroup returns a boolean if a field has been set.

### GetAzKeyid

`func (o *DnsServerDataData) GetAzKeyid() string`

GetAzKeyid returns the AzKeyid field if non-nil, zero value otherwise.

### GetAzKeyidOk

`func (o *DnsServerDataData) GetAzKeyidOk() (*string, bool)`

GetAzKeyidOk returns a tuple with the AzKeyid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzKeyid

`func (o *DnsServerDataData) SetAzKeyid(v string)`

SetAzKeyid sets AzKeyid field to given value.

### HasAzKeyid

`func (o *DnsServerDataData) HasAzKeyid() bool`

HasAzKeyid returns a boolean if a field has been set.

### GetAzSubscriptionid

`func (o *DnsServerDataData) GetAzSubscriptionid() string`

GetAzSubscriptionid returns the AzSubscriptionid field if non-nil, zero value otherwise.

### GetAzSubscriptionidOk

`func (o *DnsServerDataData) GetAzSubscriptionidOk() (*string, bool)`

GetAzSubscriptionidOk returns a tuple with the AzSubscriptionid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzSubscriptionid

`func (o *DnsServerDataData) SetAzSubscriptionid(v string)`

SetAzSubscriptionid sets AzSubscriptionid field to given value.

### HasAzSubscriptionid

`func (o *DnsServerDataData) HasAzSubscriptionid() bool`

HasAzSubscriptionid returns a boolean if a field has been set.

### GetAzTenantid

`func (o *DnsServerDataData) GetAzTenantid() string`

GetAzTenantid returns the AzTenantid field if non-nil, zero value otherwise.

### GetAzTenantidOk

`func (o *DnsServerDataData) GetAzTenantidOk() (*string, bool)`

GetAzTenantidOk returns a tuple with the AzTenantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzTenantid

`func (o *DnsServerDataData) SetAzTenantid(v string)`

SetAzTenantid sets AzTenantid field to given value.

### HasAzTenantid

`func (o *DnsServerDataData) HasAzTenantid() bool`

HasAzTenantid returns a boolean if a field has been set.

### GetServerAllowQuery

`func (o *DnsServerDataData) GetServerAllowQuery() string`

GetServerAllowQuery returns the ServerAllowQuery field if non-nil, zero value otherwise.

### GetServerAllowQueryOk

`func (o *DnsServerDataData) GetServerAllowQueryOk() (*string, bool)`

GetServerAllowQueryOk returns a tuple with the ServerAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllowQuery

`func (o *DnsServerDataData) SetServerAllowQuery(v string)`

SetServerAllowQuery sets ServerAllowQuery field to given value.

### HasServerAllowQuery

`func (o *DnsServerDataData) HasServerAllowQuery() bool`

HasServerAllowQuery returns a boolean if a field has been set.

### GetServerAllowQueryCache

`func (o *DnsServerDataData) GetServerAllowQueryCache() string`

GetServerAllowQueryCache returns the ServerAllowQueryCache field if non-nil, zero value otherwise.

### GetServerAllowQueryCacheOk

`func (o *DnsServerDataData) GetServerAllowQueryCacheOk() (*string, bool)`

GetServerAllowQueryCacheOk returns a tuple with the ServerAllowQueryCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllowQueryCache

`func (o *DnsServerDataData) SetServerAllowQueryCache(v string)`

SetServerAllowQueryCache sets ServerAllowQueryCache field to given value.

### HasServerAllowQueryCache

`func (o *DnsServerDataData) HasServerAllowQueryCache() bool`

HasServerAllowQueryCache returns a boolean if a field has been set.

### GetServerAllowRecursion

`func (o *DnsServerDataData) GetServerAllowRecursion() string`

GetServerAllowRecursion returns the ServerAllowRecursion field if non-nil, zero value otherwise.

### GetServerAllowRecursionOk

`func (o *DnsServerDataData) GetServerAllowRecursionOk() (*string, bool)`

GetServerAllowRecursionOk returns a tuple with the ServerAllowRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllowRecursion

`func (o *DnsServerDataData) SetServerAllowRecursion(v string)`

SetServerAllowRecursion sets ServerAllowRecursion field to given value.

### HasServerAllowRecursion

`func (o *DnsServerDataData) HasServerAllowRecursion() bool`

HasServerAllowRecursion returns a boolean if a field has been set.

### GetServerAllowTransfer

`func (o *DnsServerDataData) GetServerAllowTransfer() string`

GetServerAllowTransfer returns the ServerAllowTransfer field if non-nil, zero value otherwise.

### GetServerAllowTransferOk

`func (o *DnsServerDataData) GetServerAllowTransferOk() (*string, bool)`

GetServerAllowTransferOk returns a tuple with the ServerAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllowTransfer

`func (o *DnsServerDataData) SetServerAllowTransfer(v string)`

SetServerAllowTransfer sets ServerAllowTransfer field to given value.

### HasServerAllowTransfer

`func (o *DnsServerDataData) HasServerAllowTransfer() bool`

HasServerAllowTransfer returns a boolean if a field has been set.

### GetServerAlsoNotify

`func (o *DnsServerDataData) GetServerAlsoNotify() string`

GetServerAlsoNotify returns the ServerAlsoNotify field if non-nil, zero value otherwise.

### GetServerAlsoNotifyOk

`func (o *DnsServerDataData) GetServerAlsoNotifyOk() (*string, bool)`

GetServerAlsoNotifyOk returns a tuple with the ServerAlsoNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAlsoNotify

`func (o *DnsServerDataData) SetServerAlsoNotify(v string)`

SetServerAlsoNotify sets ServerAlsoNotify field to given value.

### HasServerAlsoNotify

`func (o *DnsServerDataData) HasServerAlsoNotify() bool`

HasServerAlsoNotify returns a boolean if a field has been set.

### GetServerClassName

`func (o *DnsServerDataData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DnsServerDataData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DnsServerDataData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DnsServerDataData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DnsServerDataData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DnsServerDataData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DnsServerDataData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DnsServerDataData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerCloud

`func (o *DnsServerDataData) GetServerCloud() string`

GetServerCloud returns the ServerCloud field if non-nil, zero value otherwise.

### GetServerCloudOk

`func (o *DnsServerDataData) GetServerCloudOk() (*string, bool)`

GetServerCloudOk returns a tuple with the ServerCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCloud

`func (o *DnsServerDataData) SetServerCloud(v string)`

SetServerCloud sets ServerCloud field to given value.

### HasServerCloud

`func (o *DnsServerDataData) HasServerCloud() bool`

HasServerCloud returns a boolean if a field has been set.

### GetServerComment

`func (o *DnsServerDataData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DnsServerDataData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DnsServerDataData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DnsServerDataData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerForceHybrid

`func (o *DnsServerDataData) GetServerForceHybrid() string`

GetServerForceHybrid returns the ServerForceHybrid field if non-nil, zero value otherwise.

### GetServerForceHybridOk

`func (o *DnsServerDataData) GetServerForceHybridOk() (*string, bool)`

GetServerForceHybridOk returns a tuple with the ServerForceHybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerForceHybrid

`func (o *DnsServerDataData) SetServerForceHybrid(v string)`

SetServerForceHybrid sets ServerForceHybrid field to given value.

### HasServerForceHybrid

`func (o *DnsServerDataData) HasServerForceHybrid() bool`

HasServerForceHybrid returns a boolean if a field has been set.

### GetServerForward

`func (o *DnsServerDataData) GetServerForward() string`

GetServerForward returns the ServerForward field if non-nil, zero value otherwise.

### GetServerForwardOk

`func (o *DnsServerDataData) GetServerForwardOk() (*string, bool)`

GetServerForwardOk returns a tuple with the ServerForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerForward

`func (o *DnsServerDataData) SetServerForward(v string)`

SetServerForward sets ServerForward field to given value.

### HasServerForward

`func (o *DnsServerDataData) HasServerForward() bool`

HasServerForward returns a boolean if a field has been set.

### GetServerForwarders

`func (o *DnsServerDataData) GetServerForwarders() string`

GetServerForwarders returns the ServerForwarders field if non-nil, zero value otherwise.

### GetServerForwardersOk

`func (o *DnsServerDataData) GetServerForwardersOk() (*string, bool)`

GetServerForwardersOk returns a tuple with the ServerForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerForwarders

`func (o *DnsServerDataData) SetServerForwarders(v string)`

SetServerForwarders sets ServerForwarders field to given value.

### HasServerForwarders

`func (o *DnsServerDataData) HasServerForwarders() bool`

HasServerForwarders returns a boolean if a field has been set.

### GetServerHybrid

`func (o *DnsServerDataData) GetServerHybrid() string`

GetServerHybrid returns the ServerHybrid field if non-nil, zero value otherwise.

### GetServerHybridOk

`func (o *DnsServerDataData) GetServerHybridOk() (*string, bool)`

GetServerHybridOk returns a tuple with the ServerHybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHybrid

`func (o *DnsServerDataData) SetServerHybrid(v string)`

SetServerHybrid sets ServerHybrid field to given value.

### HasServerHybrid

`func (o *DnsServerDataData) HasServerHybrid() bool`

HasServerHybrid returns a boolean if a field has been set.

### GetServerId

`func (o *DnsServerDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsServerDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsServerDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsServerDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerKeyName

`func (o *DnsServerDataData) GetServerKeyName() string`

GetServerKeyName returns the ServerKeyName field if non-nil, zero value otherwise.

### GetServerKeyNameOk

`func (o *DnsServerDataData) GetServerKeyNameOk() (*string, bool)`

GetServerKeyNameOk returns a tuple with the ServerKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeyName

`func (o *DnsServerDataData) SetServerKeyName(v string)`

SetServerKeyName sets ServerKeyName field to given value.

### HasServerKeyName

`func (o *DnsServerDataData) HasServerKeyName() bool`

HasServerKeyName returns a boolean if a field has been set.

### GetServerKeyProto

`func (o *DnsServerDataData) GetServerKeyProto() string`

GetServerKeyProto returns the ServerKeyProto field if non-nil, zero value otherwise.

### GetServerKeyProtoOk

`func (o *DnsServerDataData) GetServerKeyProtoOk() (*string, bool)`

GetServerKeyProtoOk returns a tuple with the ServerKeyProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeyProto

`func (o *DnsServerDataData) SetServerKeyProto(v string)`

SetServerKeyProto sets ServerKeyProto field to given value.

### HasServerKeyProto

`func (o *DnsServerDataData) HasServerKeyProto() bool`

HasServerKeyProto returns a boolean if a field has been set.

### GetServerKeyValue

`func (o *DnsServerDataData) GetServerKeyValue() string`

GetServerKeyValue returns the ServerKeyValue field if non-nil, zero value otherwise.

### GetServerKeyValueOk

`func (o *DnsServerDataData) GetServerKeyValueOk() (*string, bool)`

GetServerKeyValueOk returns a tuple with the ServerKeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeyValue

`func (o *DnsServerDataData) SetServerKeyValue(v string)`

SetServerKeyValue sets ServerKeyValue field to given value.

### HasServerKeyValue

`func (o *DnsServerDataData) HasServerKeyValue() bool`

HasServerKeyValue returns a boolean if a field has been set.

### GetServerName

`func (o *DnsServerDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsServerDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsServerDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsServerDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerNotify

`func (o *DnsServerDataData) GetServerNotify() string`

GetServerNotify returns the ServerNotify field if non-nil, zero value otherwise.

### GetServerNotifyOk

`func (o *DnsServerDataData) GetServerNotifyOk() (*string, bool)`

GetServerNotifyOk returns a tuple with the ServerNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNotify

`func (o *DnsServerDataData) SetServerNotify(v string)`

SetServerNotify sets ServerNotify field to given value.

### HasServerNotify

`func (o *DnsServerDataData) HasServerNotify() bool`

HasServerNotify returns a boolean if a field has been set.

### GetServerRecursion

`func (o *DnsServerDataData) GetServerRecursion() string`

GetServerRecursion returns the ServerRecursion field if non-nil, zero value otherwise.

### GetServerRecursionOk

`func (o *DnsServerDataData) GetServerRecursionOk() (*string, bool)`

GetServerRecursionOk returns a tuple with the ServerRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRecursion

`func (o *DnsServerDataData) SetServerRecursion(v string)`

SetServerRecursion sets ServerRecursion field to given value.

### HasServerRecursion

`func (o *DnsServerDataData) HasServerRecursion() bool`

HasServerRecursion returns a boolean if a field has been set.

### GetServerRole

`func (o *DnsServerDataData) GetServerRole() string`

GetServerRole returns the ServerRole field if non-nil, zero value otherwise.

### GetServerRoleOk

`func (o *DnsServerDataData) GetServerRoleOk() (*string, bool)`

GetServerRoleOk returns a tuple with the ServerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRole

`func (o *DnsServerDataData) SetServerRole(v string)`

SetServerRole sets ServerRole field to given value.

### HasServerRole

`func (o *DnsServerDataData) HasServerRole() bool`

HasServerRole returns a boolean if a field has been set.

### GetServerState

`func (o *DnsServerDataData) GetServerState() string`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *DnsServerDataData) GetServerStateOk() (*string, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *DnsServerDataData) SetServerState(v string)`

SetServerState sets ServerState field to given value.

### HasServerState

`func (o *DnsServerDataData) HasServerState() bool`

HasServerState returns a boolean if a field has been set.

### GetServerSynching

`func (o *DnsServerDataData) GetServerSynching() string`

GetServerSynching returns the ServerSynching field if non-nil, zero value otherwise.

### GetServerSynchingOk

`func (o *DnsServerDataData) GetServerSynchingOk() (*string, bool)`

GetServerSynchingOk returns a tuple with the ServerSynching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSynching

`func (o *DnsServerDataData) SetServerSynching(v string)`

SetServerSynching sets ServerSynching field to given value.

### HasServerSynching

`func (o *DnsServerDataData) HasServerSynching() bool`

HasServerSynching returns a boolean if a field has been set.

### GetServerType

`func (o *DnsServerDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DnsServerDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DnsServerDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DnsServerDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DnsServerDataData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DnsServerDataData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DnsServerDataData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DnsServerDataData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerBlastEnabled

`func (o *DnsServerDataData) GetServerBlastEnabled() string`

GetServerBlastEnabled returns the ServerBlastEnabled field if non-nil, zero value otherwise.

### GetServerBlastEnabledOk

`func (o *DnsServerDataData) GetServerBlastEnabledOk() (*string, bool)`

GetServerBlastEnabledOk returns a tuple with the ServerBlastEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBlastEnabled

`func (o *DnsServerDataData) SetServerBlastEnabled(v string)`

SetServerBlastEnabled sets ServerBlastEnabled field to given value.

### HasServerBlastEnabled

`func (o *DnsServerDataData) HasServerBlastEnabled() bool`

HasServerBlastEnabled returns a boolean if a field has been set.

### GetServerBlastStatus

`func (o *DnsServerDataData) GetServerBlastStatus() string`

GetServerBlastStatus returns the ServerBlastStatus field if non-nil, zero value otherwise.

### GetServerBlastStatusOk

`func (o *DnsServerDataData) GetServerBlastStatusOk() (*string, bool)`

GetServerBlastStatusOk returns a tuple with the ServerBlastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBlastStatus

`func (o *DnsServerDataData) SetServerBlastStatus(v string)`

SetServerBlastStatus sets ServerBlastStatus field to given value.

### HasServerBlastStatus

`func (o *DnsServerDataData) HasServerBlastStatus() bool`

HasServerBlastStatus returns a boolean if a field has been set.

### GetServerGslbSupported

`func (o *DnsServerDataData) GetServerGslbSupported() string`

GetServerGslbSupported returns the ServerGslbSupported field if non-nil, zero value otherwise.

### GetServerGslbSupportedOk

`func (o *DnsServerDataData) GetServerGslbSupportedOk() (*string, bool)`

GetServerGslbSupportedOk returns a tuple with the ServerGslbSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGslbSupported

`func (o *DnsServerDataData) SetServerGslbSupported(v string)`

SetServerGslbSupported sets ServerGslbSupported field to given value.

### HasServerGslbSupported

`func (o *DnsServerDataData) HasServerGslbSupported() bool`

HasServerGslbSupported returns a boolean if a field has been set.

### GetServerGuardianGuiManagementSupported

`func (o *DnsServerDataData) GetServerGuardianGuiManagementSupported() string`

GetServerGuardianGuiManagementSupported returns the ServerGuardianGuiManagementSupported field if non-nil, zero value otherwise.

### GetServerGuardianGuiManagementSupportedOk

`func (o *DnsServerDataData) GetServerGuardianGuiManagementSupportedOk() (*string, bool)`

GetServerGuardianGuiManagementSupportedOk returns a tuple with the ServerGuardianGuiManagementSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGuardianGuiManagementSupported

`func (o *DnsServerDataData) SetServerGuardianGuiManagementSupported(v string)`

SetServerGuardianGuiManagementSupported sets ServerGuardianGuiManagementSupported field to given value.

### HasServerGuardianGuiManagementSupported

`func (o *DnsServerDataData) HasServerGuardianGuiManagementSupported() bool`

HasServerGuardianGuiManagementSupported returns a boolean if a field has been set.

### GetServerGuardianSupported

`func (o *DnsServerDataData) GetServerGuardianSupported() string`

GetServerGuardianSupported returns the ServerGuardianSupported field if non-nil, zero value otherwise.

### GetServerGuardianSupportedOk

`func (o *DnsServerDataData) GetServerGuardianSupportedOk() (*string, bool)`

GetServerGuardianSupportedOk returns a tuple with the ServerGuardianSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGuardianSupported

`func (o *DnsServerDataData) SetServerGuardianSupported(v string)`

SetServerGuardianSupported sets ServerGuardianSupported field to given value.

### HasServerGuardianSupported

`func (o *DnsServerDataData) HasServerGuardianSupported() bool`

HasServerGuardianSupported returns a boolean if a field has been set.

### GetServerDnssecValidation

`func (o *DnsServerDataData) GetServerDnssecValidation() string`

GetServerDnssecValidation returns the ServerDnssecValidation field if non-nil, zero value otherwise.

### GetServerDnssecValidationOk

`func (o *DnsServerDataData) GetServerDnssecValidationOk() (*string, bool)`

GetServerDnssecValidationOk returns a tuple with the ServerDnssecValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDnssecValidation

`func (o *DnsServerDataData) SetServerDnssecValidation(v string)`

SetServerDnssecValidation sets ServerDnssecValidation field to given value.

### HasServerDnssecValidation

`func (o *DnsServerDataData) HasServerDnssecValidation() bool`

HasServerDnssecValidation returns a boolean if a field has been set.

### GetServerGssEnabled

`func (o *DnsServerDataData) GetServerGssEnabled() string`

GetServerGssEnabled returns the ServerGssEnabled field if non-nil, zero value otherwise.

### GetServerGssEnabledOk

`func (o *DnsServerDataData) GetServerGssEnabledOk() (*string, bool)`

GetServerGssEnabledOk returns a tuple with the ServerGssEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGssEnabled

`func (o *DnsServerDataData) SetServerGssEnabled(v string)`

SetServerGssEnabled sets ServerGssEnabled field to given value.

### HasServerGssEnabled

`func (o *DnsServerDataData) HasServerGssEnabled() bool`

HasServerGssEnabled returns a boolean if a field has been set.

### GetServerGssKeytabId

`func (o *DnsServerDataData) GetServerGssKeytabId() string`

GetServerGssKeytabId returns the ServerGssKeytabId field if non-nil, zero value otherwise.

### GetServerGssKeytabIdOk

`func (o *DnsServerDataData) GetServerGssKeytabIdOk() (*string, bool)`

GetServerGssKeytabIdOk returns a tuple with the ServerGssKeytabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGssKeytabId

`func (o *DnsServerDataData) SetServerGssKeytabId(v string)`

SetServerGssKeytabId sets ServerGssKeytabId field to given value.

### HasServerGssKeytabId

`func (o *DnsServerDataData) HasServerGssKeytabId() bool`

HasServerGssKeytabId returns a boolean if a field has been set.

### GetServerAddr

`func (o *DnsServerDataData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DnsServerDataData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DnsServerDataData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DnsServerDataData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetServerIpmLogin

`func (o *DnsServerDataData) GetServerIpmLogin() string`

GetServerIpmLogin returns the ServerIpmLogin field if non-nil, zero value otherwise.

### GetServerIpmLoginOk

`func (o *DnsServerDataData) GetServerIpmLoginOk() (*string, bool)`

GetServerIpmLoginOk returns a tuple with the ServerIpmLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmLogin

`func (o *DnsServerDataData) SetServerIpmLogin(v string)`

SetServerIpmLogin sets ServerIpmLogin field to given value.

### HasServerIpmLogin

`func (o *DnsServerDataData) HasServerIpmLogin() bool`

HasServerIpmLogin returns a boolean if a field has been set.

### GetServerIpmIsPackage

`func (o *DnsServerDataData) GetServerIpmIsPackage() string`

GetServerIpmIsPackage returns the ServerIpmIsPackage field if non-nil, zero value otherwise.

### GetServerIpmIsPackageOk

`func (o *DnsServerDataData) GetServerIpmIsPackageOk() (*string, bool)`

GetServerIpmIsPackageOk returns a tuple with the ServerIpmIsPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmIsPackage

`func (o *DnsServerDataData) SetServerIpmIsPackage(v string)`

SetServerIpmIsPackage sets ServerIpmIsPackage field to given value.

### HasServerIpmIsPackage

`func (o *DnsServerDataData) HasServerIpmIsPackage() bool`

HasServerIpmIsPackage returns a boolean if a field has been set.

### GetServerIpmProtocol

`func (o *DnsServerDataData) GetServerIpmProtocol() string`

GetServerIpmProtocol returns the ServerIpmProtocol field if non-nil, zero value otherwise.

### GetServerIpmProtocolOk

`func (o *DnsServerDataData) GetServerIpmProtocolOk() (*string, bool)`

GetServerIpmProtocolOk returns a tuple with the ServerIpmProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmProtocol

`func (o *DnsServerDataData) SetServerIpmProtocol(v string)`

SetServerIpmProtocol sets ServerIpmProtocol field to given value.

### HasServerIpmProtocol

`func (o *DnsServerDataData) HasServerIpmProtocol() bool`

HasServerIpmProtocol returns a boolean if a field has been set.

### GetServerIpmType

`func (o *DnsServerDataData) GetServerIpmType() string`

GetServerIpmType returns the ServerIpmType field if non-nil, zero value otherwise.

### GetServerIpmTypeOk

`func (o *DnsServerDataData) GetServerIpmTypeOk() (*string, bool)`

GetServerIpmTypeOk returns a tuple with the ServerIpmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmType

`func (o *DnsServerDataData) SetServerIpmType(v string)`

SetServerIpmType sets ServerIpmType field to given value.

### HasServerIpmType

`func (o *DnsServerDataData) HasServerIpmType() bool`

HasServerIpmType returns a boolean if a field has been set.

### GetServerIsolated

`func (o *DnsServerDataData) GetServerIsolated() string`

GetServerIsolated returns the ServerIsolated field if non-nil, zero value otherwise.

### GetServerIsolatedOk

`func (o *DnsServerDataData) GetServerIsolatedOk() (*string, bool)`

GetServerIsolatedOk returns a tuple with the ServerIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIsolated

`func (o *DnsServerDataData) SetServerIsolated(v string)`

SetServerIsolated sets ServerIsolated field to given value.

### HasServerIsolated

`func (o *DnsServerDataData) HasServerIsolated() bool`

HasServerIsolated returns a boolean if a field has been set.

### GetServerLdapDomain

`func (o *DnsServerDataData) GetServerLdapDomain() string`

GetServerLdapDomain returns the ServerLdapDomain field if non-nil, zero value otherwise.

### GetServerLdapDomainOk

`func (o *DnsServerDataData) GetServerLdapDomainOk() (*string, bool)`

GetServerLdapDomainOk returns a tuple with the ServerLdapDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLdapDomain

`func (o *DnsServerDataData) SetServerLdapDomain(v string)`

SetServerLdapDomain sets ServerLdapDomain field to given value.

### HasServerLdapDomain

`func (o *DnsServerDataData) HasServerLdapDomain() bool`

HasServerLdapDomain returns a boolean if a field has been set.

### GetServerLdapUser

`func (o *DnsServerDataData) GetServerLdapUser() string`

GetServerLdapUser returns the ServerLdapUser field if non-nil, zero value otherwise.

### GetServerLdapUserOk

`func (o *DnsServerDataData) GetServerLdapUserOk() (*string, bool)`

GetServerLdapUserOk returns a tuple with the ServerLdapUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLdapUser

`func (o *DnsServerDataData) SetServerLdapUser(v string)`

SetServerLdapUser sets ServerLdapUser field to given value.

### HasServerLdapUser

`func (o *DnsServerDataData) HasServerLdapUser() bool`

HasServerLdapUser returns a boolean if a field has been set.

### GetServerMultistatus

`func (o *DnsServerDataData) GetServerMultistatus() string`

GetServerMultistatus returns the ServerMultistatus field if non-nil, zero value otherwise.

### GetServerMultistatusOk

`func (o *DnsServerDataData) GetServerMultistatusOk() (*string, bool)`

GetServerMultistatusOk returns a tuple with the ServerMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMultistatus

`func (o *DnsServerDataData) SetServerMultistatus(v string)`

SetServerMultistatus sets ServerMultistatus field to given value.

### HasServerMultistatus

`func (o *DnsServerDataData) HasServerMultistatus() bool`

HasServerMultistatus returns a boolean if a field has been set.

### GetServerQuerylogState

`func (o *DnsServerDataData) GetServerQuerylogState() string`

GetServerQuerylogState returns the ServerQuerylogState field if non-nil, zero value otherwise.

### GetServerQuerylogStateOk

`func (o *DnsServerDataData) GetServerQuerylogStateOk() (*string, bool)`

GetServerQuerylogStateOk returns a tuple with the ServerQuerylogState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerQuerylogState

`func (o *DnsServerDataData) SetServerQuerylogState(v string)`

SetServerQuerylogState sets ServerQuerylogState field to given value.

### HasServerQuerylogState

`func (o *DnsServerDataData) HasServerQuerylogState() bool`

HasServerQuerylogState returns a boolean if a field has been set.

### GetReverseProxyConf

`func (o *DnsServerDataData) GetReverseProxyConf() string`

GetReverseProxyConf returns the ReverseProxyConf field if non-nil, zero value otherwise.

### GetReverseProxyConfOk

`func (o *DnsServerDataData) GetReverseProxyConfOk() (*string, bool)`

GetReverseProxyConfOk returns a tuple with the ReverseProxyConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseProxyConf

`func (o *DnsServerDataData) SetReverseProxyConf(v string)`

SetReverseProxyConf sets ReverseProxyConf field to given value.

### HasReverseProxyConf

`func (o *DnsServerDataData) HasReverseProxyConf() bool`

HasReverseProxyConf returns a boolean if a field has been set.

### GetServerSnmpId

`func (o *DnsServerDataData) GetServerSnmpId() string`

GetServerSnmpId returns the ServerSnmpId field if non-nil, zero value otherwise.

### GetServerSnmpIdOk

`func (o *DnsServerDataData) GetServerSnmpIdOk() (*string, bool)`

GetServerSnmpIdOk returns a tuple with the ServerSnmpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpId

`func (o *DnsServerDataData) SetServerSnmpId(v string)`

SetServerSnmpId sets ServerSnmpId field to given value.

### HasServerSnmpId

`func (o *DnsServerDataData) HasServerSnmpId() bool`

HasServerSnmpId returns a boolean if a field has been set.

### GetServerStatEnabled

`func (o *DnsServerDataData) GetServerStatEnabled() string`

GetServerStatEnabled returns the ServerStatEnabled field if non-nil, zero value otherwise.

### GetServerStatEnabledOk

`func (o *DnsServerDataData) GetServerStatEnabledOk() (*string, bool)`

GetServerStatEnabledOk returns a tuple with the ServerStatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatEnabled

`func (o *DnsServerDataData) SetServerStatEnabled(v string)`

SetServerStatEnabled sets ServerStatEnabled field to given value.

### HasServerStatEnabled

`func (o *DnsServerDataData) HasServerStatEnabled() bool`

HasServerStatEnabled returns a boolean if a field has been set.

### GetServerStatNiceness

`func (o *DnsServerDataData) GetServerStatNiceness() string`

GetServerStatNiceness returns the ServerStatNiceness field if non-nil, zero value otherwise.

### GetServerStatNicenessOk

`func (o *DnsServerDataData) GetServerStatNicenessOk() (*string, bool)`

GetServerStatNicenessOk returns a tuple with the ServerStatNiceness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatNiceness

`func (o *DnsServerDataData) SetServerStatNiceness(v string)`

SetServerStatNiceness sets ServerStatNiceness field to given value.

### HasServerStatNiceness

`func (o *DnsServerDataData) HasServerStatNiceness() bool`

HasServerStatNiceness returns a boolean if a field has been set.

### GetServerStatPeriod

`func (o *DnsServerDataData) GetServerStatPeriod() string`

GetServerStatPeriod returns the ServerStatPeriod field if non-nil, zero value otherwise.

### GetServerStatPeriodOk

`func (o *DnsServerDataData) GetServerStatPeriodOk() (*string, bool)`

GetServerStatPeriodOk returns a tuple with the ServerStatPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatPeriod

`func (o *DnsServerDataData) SetServerStatPeriod(v string)`

SetServerStatPeriod sets ServerStatPeriod field to given value.

### HasServerStatPeriod

`func (o *DnsServerDataData) HasServerStatPeriod() bool`

HasServerStatPeriod returns a boolean if a field has been set.

### GetServerStatTime

`func (o *DnsServerDataData) GetServerStatTime() string`

GetServerStatTime returns the ServerStatTime field if non-nil, zero value otherwise.

### GetServerStatTimeOk

`func (o *DnsServerDataData) GetServerStatTimeOk() (*string, bool)`

GetServerStatTimeOk returns a tuple with the ServerStatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatTime

`func (o *DnsServerDataData) SetServerStatTime(v string)`

SetServerStatTime sets ServerStatTime field to given value.

### HasServerStatTime

`func (o *DnsServerDataData) HasServerStatTime() bool`

HasServerStatTime returns a boolean if a field has been set.

### GetTotalSmartMembers

`func (o *DnsServerDataData) GetTotalSmartMembers() string`

GetTotalSmartMembers returns the TotalSmartMembers field if non-nil, zero value otherwise.

### GetTotalSmartMembersOk

`func (o *DnsServerDataData) GetTotalSmartMembersOk() (*string, bool)`

GetTotalSmartMembersOk returns a tuple with the TotalSmartMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSmartMembers

`func (o *DnsServerDataData) SetTotalSmartMembers(v string)`

SetTotalSmartMembers sets TotalSmartMembers field to given value.

### HasTotalSmartMembers

`func (o *DnsServerDataData) HasTotalSmartMembers() bool`

HasTotalSmartMembers returns a boolean if a field has been set.

### GetSmartArch

`func (o *DnsServerDataData) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DnsServerDataData) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DnsServerDataData) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DnsServerDataData) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartMembersName

`func (o *DnsServerDataData) GetSmartMembersName() string`

GetSmartMembersName returns the SmartMembersName field if non-nil, zero value otherwise.

### GetSmartMembersNameOk

`func (o *DnsServerDataData) GetSmartMembersNameOk() (*string, bool)`

GetSmartMembersNameOk returns a tuple with the SmartMembersName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartMembersName

`func (o *DnsServerDataData) SetSmartMembersName(v string)`

SetSmartMembersName sets SmartMembersName field to given value.

### HasSmartMembersName

`func (o *DnsServerDataData) HasSmartMembersName() bool`

HasSmartMembersName returns a boolean if a field has been set.

### GetSmartParentArch

`func (o *DnsServerDataData) GetSmartParentArch() string`

GetSmartParentArch returns the SmartParentArch field if non-nil, zero value otherwise.

### GetSmartParentArchOk

`func (o *DnsServerDataData) GetSmartParentArchOk() (*string, bool)`

GetSmartParentArchOk returns a tuple with the SmartParentArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentArch

`func (o *DnsServerDataData) SetSmartParentArch(v string)`

SetSmartParentArch sets SmartParentArch field to given value.

### HasSmartParentArch

`func (o *DnsServerDataData) HasSmartParentArch() bool`

HasSmartParentArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DnsServerDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DnsServerDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DnsServerDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DnsServerDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DnsServerDataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DnsServerDataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DnsServerDataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DnsServerDataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.

### GetSmartPublicNsList

`func (o *DnsServerDataData) GetSmartPublicNsList() string`

GetSmartPublicNsList returns the SmartPublicNsList field if non-nil, zero value otherwise.

### GetSmartPublicNsListOk

`func (o *DnsServerDataData) GetSmartPublicNsListOk() (*string, bool)`

GetSmartPublicNsListOk returns a tuple with the SmartPublicNsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartPublicNsList

`func (o *DnsServerDataData) SetSmartPublicNsList(v string)`

SetSmartPublicNsList sets SmartPublicNsList field to given value.

### HasSmartPublicNsList

`func (o *DnsServerDataData) HasSmartPublicNsList() bool`

HasSmartPublicNsList returns a boolean if a field has been set.

### GetServerWindnsPort

`func (o *DnsServerDataData) GetServerWindnsPort() string`

GetServerWindnsPort returns the ServerWindnsPort field if non-nil, zero value otherwise.

### GetServerWindnsPortOk

`func (o *DnsServerDataData) GetServerWindnsPortOk() (*string, bool)`

GetServerWindnsPortOk returns a tuple with the ServerWindnsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWindnsPort

`func (o *DnsServerDataData) SetServerWindnsPort(v string)`

SetServerWindnsPort sets ServerWindnsPort field to given value.

### HasServerWindnsPort

`func (o *DnsServerDataData) HasServerWindnsPort() bool`

HasServerWindnsPort returns a boolean if a field has been set.

### GetServerWindnsProtocol

`func (o *DnsServerDataData) GetServerWindnsProtocol() string`

GetServerWindnsProtocol returns the ServerWindnsProtocol field if non-nil, zero value otherwise.

### GetServerWindnsProtocolOk

`func (o *DnsServerDataData) GetServerWindnsProtocolOk() (*string, bool)`

GetServerWindnsProtocolOk returns a tuple with the ServerWindnsProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWindnsProtocol

`func (o *DnsServerDataData) SetServerWindnsProtocol(v string)`

SetServerWindnsProtocol sets ServerWindnsProtocol field to given value.

### HasServerWindnsProtocol

`func (o *DnsServerDataData) HasServerWindnsProtocol() bool`

HasServerWindnsProtocol returns a boolean if a field has been set.

### GetServerWindnsUseSsl

`func (o *DnsServerDataData) GetServerWindnsUseSsl() string`

GetServerWindnsUseSsl returns the ServerWindnsUseSsl field if non-nil, zero value otherwise.

### GetServerWindnsUseSslOk

`func (o *DnsServerDataData) GetServerWindnsUseSslOk() (*string, bool)`

GetServerWindnsUseSslOk returns a tuple with the ServerWindnsUseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWindnsUseSsl

`func (o *DnsServerDataData) SetServerWindnsUseSsl(v string)`

SetServerWindnsUseSsl sets ServerWindnsUseSsl field to given value.

### HasServerWindnsUseSsl

`func (o *DnsServerDataData) HasServerWindnsUseSsl() bool`

HasServerWindnsUseSsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


