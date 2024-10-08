# DataInnerDnsServerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerAwsDelegationSet** | Pointer to **string** | The reusable delegation set ID configured on the Amazon Route 53 public server. | [optional] 
**ServerAwsKeyid** | Pointer to **string** | The AWS access key identifier (ID) of the DNS server. | [optional] 
**ServerAwsRoleArn** | Pointer to **string** | The Amazon Resource Name (ARN) of the role used for the Amazon Route 53 DNS server. | [optional] 
**ServerAwsRoleExternalId** | Pointer to **string** | The external ID of the role used for the Amazon Route 53 DNS server. | [optional] 
**ServerAwsRoleSessionName** | Pointer to **string** | The session name of the role used for the Amazon Route 53 DNS server. | [optional] 
**ServerAwsUseRole** | Pointer to **string** | The configuration of the option &lt;b&gt;Use a role&lt;/b&gt; of the Amazon Route 53 DNS server, either enabled (&lt;b&gt;1&lt;/b&gt;) or not (&lt;b&gt;0&lt;/b&gt;). If enabled, the role can rely on ARN, an external ID or a session. | [optional] 
**ServerAzGroup** | Pointer to **string** | For Microsoft Azure servers, the resource group of the DNS server. | [optional] 
**ServerAzKeyid** | Pointer to **string** | For Microsoft Azure servers, the Azure Application ID of the DNS server. | [optional] 
**ServerAzSubscriptionid** | Pointer to **string** | For Microsoft Azure servers, the subscription ID of the DNS server. | [optional] 
**ServerAzTenantid** | Pointer to **string** | For Microsoft Azure servers, the tenant ID of the DNS server. | [optional] 
**ConnectionprofileName** | Pointer to **string** | The name of the connection profile used as connection method for the DNS server. | [optional] 
**ServerAllowQuery** | Pointer to **string** | The ACL values associated with the allow-query configuration of the DNS server, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ServerAllowQueryCache** | Pointer to **string** | The ACL values associated with the allow-query-cache configuration of the DNS server, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ServerAllowRecursion** | Pointer to **string** | The ACL values associated with the allow-recursion configuration of the DNS server, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ServerAllowTransfer** | Pointer to **string** | The ACL values associated with the allow-transfer configuration of the DNS server, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ServerAlsoNotify** | Pointer to **string** | The IP address and port of the DNS server managing the smart architecture. If the parameter &lt;b&gt;server_notify&lt;/b&gt; is set to &lt;b&gt;yes&lt;/b&gt; or &lt;b&gt;explicit&lt;/b&gt;, the server specified is instantly notified of any slave zones updates. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DNS server, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DNS server. | [optional] 
**ServerCloud** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerCloudPrivate** | Pointer to **string** | The type of the DNS Amazon Route 53 or Azure server, either Public (&lt;b&gt;0&lt;/b&gt;) or Private (&lt;b&gt;1&lt;/b&gt;). If the server is not a cloud server, &lt;b&gt;0&lt;/b&gt; is returned. | [optional] 
**ServerComment** | Pointer to **string** | The description of the DNS server. | [optional] 
**ServerForceHybrid** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerForward** | Pointer to **string** | The forwarding mode of the DNS server. No value indicates that the forwarding is disabled: &lt;table&gt;&lt;caption&gt;server_forward possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;first&lt;/td&gt;&lt;td &gt;The server sends the queries to the forwarder(s). If no answer is returned, it attempts to answer the queries on its own.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;only&lt;/td&gt;&lt;td &gt;The server only forwards the queries to the forwarder(s). Required by some reverse forward zones (e.g., in the case of private addresses).&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerForwarders** | Pointer to **string** | The IP address(es) of the forwarder(s) associated with the DNS server. It lists the DNS servers to which any unknown zone should be sent, as follows: &lt;b&gt;&amp;lt;ip_address1&amp;gt;;&amp;lt;ip_address2&amp;gt;;...&lt;/b&gt; . | [optional] 
**ServerHybrid** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. | [optional] 
**ServerKeyName** | Pointer to **string** | The name of the DNS TSIG key associated with the DNS server. | [optional] 
**ServerKeyProto** | Pointer to **string** | The encryption protocol of the TSIG key associated with the DNS server. | [optional] 
**ServerKeyValue** | Pointer to **string** | The value of the TSIG key associated with the DNS server. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server. | [optional] 
**ServerNotify** | Pointer to **string** | The notify status of the DNS server:&lt;table&gt;&lt;caption&gt;server_notify possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;No notify message is sent when changes are performed in the master zones.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;The notify messages are sent to the target of the NS records of the master zone. They are also sent to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;explicit&lt;/td&gt;&lt;td &gt;The notify messages are only sent to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerRecursion** | Pointer to **string** | The recursion status of the DNS server:&lt;table&gt;&lt;caption&gt;server_recursion possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;The server only provides iterative query behavior - normally resulting in a referral. If the answer to the query already exists in the cache it will be returned whatever the value of this statement.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;The server always provides recursive query behavior if requested by the client.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerRole** | Pointer to **string** | The role of the DNS server in the smart architecture, either &lt;b&gt;master&lt;/b&gt;, &lt;b&gt;hidden-master&lt;/b&gt;, &lt;b&gt;pseudo-master&lt;/b&gt; or &lt;b&gt;slave&lt;/b&gt;. | [optional] 
**ServerRpzBreakDnssec** | Pointer to **string** | The configuration of the option &lt;b&gt;Enable break-dnssec on the server&lt;/b&gt;. It applies to the RPZ zones of the server, if it is not set at zone level. By default it is set to &lt;b&gt;0 (no)&lt;/b&gt;, if set to &lt;b&gt;1 (yes)&lt;/b&gt; the server processes policies on all DNSSEC queries. | [optional] 
**ServerRpzMaxPolicyTtl** | Pointer to **string** | The configuration of the option &lt;b&gt;Server max policy TTL&lt;/b&gt;, i.e. the number of seconds of the max policy Time To Live of the server. It applies to the RPZ zones of the server, if it is not set at zone level. By default it is empty, i.e. set to &lt;b&gt;5 &lt;/b&gt;seconds. | [optional] 
**ServerRpzMinNsDots** | Pointer to **string** | The configuration of the option &lt;b&gt;Server minimum dot separators&lt;/b&gt;&lt;b&gt;of the server&lt;/b&gt;, i.e. the minimum number of dot separators required in any QNAME to process policies. It applies to the RPZ zones of the server, if it is not set at zone level. By default it is empty, i.e. set to &lt;b&gt;1&lt;/b&gt;. | [optional] 
**ServerRpzQnameWaitRecurse** | Pointer to **string** | The configuration of the option &lt;b&gt;Enable qname-wait-recurse on the server&lt;/b&gt;. It applies to the RPZ zones of the server, if it is not set at zone level. By default it is set to &lt;b&gt;0 (no)&lt;/b&gt;, if set to &lt;b&gt;1 (yes)&lt;/b&gt; the server only process policies when the results of any query are available. | [optional] 
**ServerRpzRecursiveOnly** | Pointer to **string** | The configuration of the option &lt;b&gt;Enable recursive-only on the server&lt;/b&gt;. It applies to the RPZ zones of the server, if it is not set at zone level. By default it is set to &lt;b&gt;1 (yes)&lt;/b&gt;, the server only processes policies on recursive queries. | [optional] 
**ServerState** | Pointer to **string** | The status of the DNS server:&lt;table&gt;&lt;caption&gt;server_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ER&lt;/td&gt;&lt;td &gt;The license used in SOLIDserver is not compliant with the added server: the license is invalid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ES&lt;/td&gt;&lt;td &gt;The server configuration could not be parsed properly.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ET&lt;/td&gt;&lt;td &gt;The server does not answer anymore due to a scheduled configuration of the server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IC&lt;/td&gt;&lt;td &gt;The SSL credentials are invalid&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IP&lt;/td&gt;&lt;td &gt;The account used to add the Microsoft Windows DNS server does not have sufficient privileges to manage it.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IR&lt;/td&gt;&lt;td &gt;SOLIDserver cannot resolve the AWS DNS service. The Amazon services are unreachable and the Amazon Route 53 server cannot be managed. Make sure that the DNS resolvers declared on the page are valid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IS&lt;/td&gt;&lt;td &gt;There was a setting error during the server declaration. For instance, some settings were added to a server that does not support them or a smart architecture is not managing any physical server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IT&lt;/td&gt;&lt;td &gt;The server editing performed from the GUI is not pushed to the server because SOLIDserver time and date are incorrect. You must use the UTC system on the appliance, especially when managing Amazon Route 53 servers.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;LS&lt;/td&gt;&lt;td &gt;The server configuration is not viable.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;N&lt;/td&gt;&lt;td &gt;The server does not have a status as it has not synchronized yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;UE&lt;/td&gt;&lt;td &gt;An error occurred that SOLIDserver could not identify.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;Y&lt;/td&gt;&lt;td &gt;The server is operational.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerSynching** | Pointer to **string** | The synchronization status of the DNS server. &lt;b&gt;1&lt;/b&gt; indicates that the server is currently being synchronized. | [optional] 
**ServerType** | Pointer to **string** | The type of the DNS server: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msdaemon&lt;/td&gt;&lt;td &gt;Microsoft Windows DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;aws&lt;/td&gt;&lt;td &gt;Amazon Route 53 server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;other&lt;/td&gt;&lt;td &gt;Generic DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdns&lt;/td&gt;&lt;td &gt;EfficientIP DNS smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DNS server. | [optional] 
**ServerVpcList** | Pointer to **string** | The list of cloud networks configured on the DNS private server, separated by a comma:&lt;ul class&#x3D;dashed &gt;&lt;li&gt; For an Azure server, it returns the list of virtual networks.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; For an Amazon Route 53 server, it returns the list of Virtual Private Cloud (VPC).&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**ServerBlastEnabled** | Pointer to **string** | The status of the service DNS Guardian, either enabled (&lt;b&gt;1&lt;/b&gt;) or disabled (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**ServerGuardianPushStatus** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerBlastStatus** | Pointer to **string** | The status of the Guardian server, either OK (&lt;b&gt;1&lt;/b&gt;), Stopped (&lt;b&gt;2&lt;/b&gt;), Invalid Credentials (&lt;b&gt;4&lt;/b&gt;) or Timeout (&lt;b&gt;5&lt;/b&gt;). | [optional] 
**ServerGslbSupported** | Pointer to **string** | The license GSLB activation status. &lt;b&gt;1&lt;/b&gt; indicates your license includes GSLB and your appliance supports it. | [optional] 
**ServerGuardianGuiManagementSupported** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerGuardianSupported** | Pointer to **string** | The license Guardian activation status. &lt;b&gt;1&lt;/b&gt; indicates your license includes Guardian and your appliance supports its latest features. | [optional] 
**ServerDnssecValidation** | Pointer to **string** | The DNSSEC resolution status of the DNS server. &lt;b&gt;yes&lt;/b&gt; indicates it is enabled. | [optional] 
**ServerGssEnabled** | Pointer to **string** | The GSS-TSIG status of the DNS server. &lt;b&gt;1&lt;/b&gt; indicates that GSS-TSIG is enabled on the server. | [optional] 
**ServerGssKeytabId** | Pointer to **string** | The database identifier (ID) of the DNS GSS-TSIG keytab. | [optional] 
**ServerGuardianStatsOnlySupported** | Pointer to **string** | A way to determine if the server only retrieves Guardian statistics (&lt;b&gt;1&lt;/b&gt;) or not (&lt;b&gt;0&lt;/b&gt;). On appliances where the service &lt;b&gt;DNS Guardian / GSLB server&lt;/b&gt; is not configured, the server cannot be configured with Guardian options, it can only retrieve statistics. | [optional] 
**ServerHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server_addr&lt;/b&gt; or &lt;b&gt;server_addr6&lt;/b&gt;. | [optional] 
**ServerAddr6** | Pointer to **string** | The IPv6 address of the DNS server, in hexadecimal format. | [optional] 
**ServerAddr** | Pointer to **string** | The IPv4 address of the DNS server, in hexadecimal format. | [optional] 
**ServerIpmLogin** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIpmIsPackage** | Pointer to **string** | The DNS server package information. &lt;b&gt;Y&lt;/b&gt; for an EfficientIP Package server, &lt;b&gt;N&lt;/b&gt; for an appliance or virtual machine, &lt;b&gt;U&lt;/b&gt; the package information is irrelevant. For servers with a &lt;b&gt;server_type&lt;/b&gt; set to &lt;b&gt;ipm&lt;/b&gt;, &lt;b&gt;U&lt;/b&gt; indicates either EfficientIP Packages or appliances/virtual machines. | [optional] 
**ServerIpmProtocol** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIpmType** | Pointer to **string** | The engine type of the DNS server: &lt;b&gt;named&lt;/b&gt; (BIND engine), &lt;b&gt;nsd&lt;/b&gt; (NSD engine) or &lt;b&gt;unbound&lt;/b&gt; (Unbound engine). | [optional] 
**ServerIsolated** | Pointer to **string** | A way to determine if the server can update any other module &lt;b&gt;(1)&lt;/b&gt;. | [optional] 
**ServerLdapDomain** | Pointer to **string** | For Microsoft Windows servers, the domain of the DNS server. | [optional] 
**ServerLdapUser** | Pointer to **string** | For Microsoft Windows servers, the login of the user communicating with the DNS server. | [optional] 
**ServerMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ServerQuerylogState** | Pointer to **string** | The DNS querylog status. &lt;b&gt;1&lt;/b&gt; indicates that the DNS server querylog is enabled. | [optional] 
**ReverseProxyConf** | Pointer to **string** | The URL of the HTTP(S) reverse proxy server that forwards client queries to the DNS server, if you configured one. | [optional] 
**ServerSnmpId** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatEnabled** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatNiceness** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatPeriod** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatTime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**TotalSmartMembers** | Pointer to **string** | The total number of servers managed by the DNS smart architecture. | [optional] 
**SmartArch** | Pointer to **string** | The type of the DNS smart architecture:&lt;table&gt;&lt;caption&gt;smart_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;masterslave&lt;/td&gt;&lt;td &gt;Master/Slave&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;stealth&lt;/td&gt;&lt;td &gt;Stealth&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;multimaster&lt;/td&gt;&lt;td &gt;Multi-Master&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;Single-Server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;farm&lt;/td&gt;&lt;td &gt;Farm&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SmartMembersName** | Pointer to **string** | The list of the servers managed by the DNS smart architecture, as follows: &lt;b&gt;&amp;lt;dns_name&amp;gt;,&amp;lt;dns_name&amp;gt;,...&lt;/b&gt; . | [optional] 
**SmartParentArch** | Pointer to **string** | The type of the DNS smart architecture managing the DNS server. No value indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DNS smart architecture managing the DNS server. &lt;b&gt;0&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DNS smart architecture managing the DNS server. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartPublicNsList** | Pointer to **string** | The list of the published name servers associated with the DNS smart architecture, as follows: &lt;b&gt;&amp;lt;ns1&amp;gt;;&amp;lt;ns2&amp;gt;;...&lt;/b&gt; . | [optional] 
**ServerWindnsPort** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerWindnsProtocol** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerWindnsUseSsl** | Pointer to **string** | Internal use. Not documented. | [optional] 

## Methods

### NewDataInnerDnsServerData

`func NewDataInnerDnsServerData() *DataInnerDnsServerData`

NewDataInnerDnsServerData instantiates a new DataInnerDnsServerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDnsServerDataWithDefaults

`func NewDataInnerDnsServerDataWithDefaults() *DataInnerDnsServerData`

NewDataInnerDnsServerDataWithDefaults instantiates a new DataInnerDnsServerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerAwsDelegationSet

`func (o *DataInnerDnsServerData) GetServerAwsDelegationSet() string`

GetServerAwsDelegationSet returns the ServerAwsDelegationSet field if non-nil, zero value otherwise.

### GetServerAwsDelegationSetOk

`func (o *DataInnerDnsServerData) GetServerAwsDelegationSetOk() (*string, bool)`

GetServerAwsDelegationSetOk returns a tuple with the ServerAwsDelegationSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAwsDelegationSet

`func (o *DataInnerDnsServerData) SetServerAwsDelegationSet(v string)`

SetServerAwsDelegationSet sets ServerAwsDelegationSet field to given value.

### HasServerAwsDelegationSet

`func (o *DataInnerDnsServerData) HasServerAwsDelegationSet() bool`

HasServerAwsDelegationSet returns a boolean if a field has been set.

### GetServerAwsKeyid

`func (o *DataInnerDnsServerData) GetServerAwsKeyid() string`

GetServerAwsKeyid returns the ServerAwsKeyid field if non-nil, zero value otherwise.

### GetServerAwsKeyidOk

`func (o *DataInnerDnsServerData) GetServerAwsKeyidOk() (*string, bool)`

GetServerAwsKeyidOk returns a tuple with the ServerAwsKeyid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAwsKeyid

`func (o *DataInnerDnsServerData) SetServerAwsKeyid(v string)`

SetServerAwsKeyid sets ServerAwsKeyid field to given value.

### HasServerAwsKeyid

`func (o *DataInnerDnsServerData) HasServerAwsKeyid() bool`

HasServerAwsKeyid returns a boolean if a field has been set.

### GetServerAwsRoleArn

`func (o *DataInnerDnsServerData) GetServerAwsRoleArn() string`

GetServerAwsRoleArn returns the ServerAwsRoleArn field if non-nil, zero value otherwise.

### GetServerAwsRoleArnOk

`func (o *DataInnerDnsServerData) GetServerAwsRoleArnOk() (*string, bool)`

GetServerAwsRoleArnOk returns a tuple with the ServerAwsRoleArn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAwsRoleArn

`func (o *DataInnerDnsServerData) SetServerAwsRoleArn(v string)`

SetServerAwsRoleArn sets ServerAwsRoleArn field to given value.

### HasServerAwsRoleArn

`func (o *DataInnerDnsServerData) HasServerAwsRoleArn() bool`

HasServerAwsRoleArn returns a boolean if a field has been set.

### GetServerAwsRoleExternalId

`func (o *DataInnerDnsServerData) GetServerAwsRoleExternalId() string`

GetServerAwsRoleExternalId returns the ServerAwsRoleExternalId field if non-nil, zero value otherwise.

### GetServerAwsRoleExternalIdOk

`func (o *DataInnerDnsServerData) GetServerAwsRoleExternalIdOk() (*string, bool)`

GetServerAwsRoleExternalIdOk returns a tuple with the ServerAwsRoleExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAwsRoleExternalId

`func (o *DataInnerDnsServerData) SetServerAwsRoleExternalId(v string)`

SetServerAwsRoleExternalId sets ServerAwsRoleExternalId field to given value.

### HasServerAwsRoleExternalId

`func (o *DataInnerDnsServerData) HasServerAwsRoleExternalId() bool`

HasServerAwsRoleExternalId returns a boolean if a field has been set.

### GetServerAwsRoleSessionName

`func (o *DataInnerDnsServerData) GetServerAwsRoleSessionName() string`

GetServerAwsRoleSessionName returns the ServerAwsRoleSessionName field if non-nil, zero value otherwise.

### GetServerAwsRoleSessionNameOk

`func (o *DataInnerDnsServerData) GetServerAwsRoleSessionNameOk() (*string, bool)`

GetServerAwsRoleSessionNameOk returns a tuple with the ServerAwsRoleSessionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAwsRoleSessionName

`func (o *DataInnerDnsServerData) SetServerAwsRoleSessionName(v string)`

SetServerAwsRoleSessionName sets ServerAwsRoleSessionName field to given value.

### HasServerAwsRoleSessionName

`func (o *DataInnerDnsServerData) HasServerAwsRoleSessionName() bool`

HasServerAwsRoleSessionName returns a boolean if a field has been set.

### GetServerAwsUseRole

`func (o *DataInnerDnsServerData) GetServerAwsUseRole() string`

GetServerAwsUseRole returns the ServerAwsUseRole field if non-nil, zero value otherwise.

### GetServerAwsUseRoleOk

`func (o *DataInnerDnsServerData) GetServerAwsUseRoleOk() (*string, bool)`

GetServerAwsUseRoleOk returns a tuple with the ServerAwsUseRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAwsUseRole

`func (o *DataInnerDnsServerData) SetServerAwsUseRole(v string)`

SetServerAwsUseRole sets ServerAwsUseRole field to given value.

### HasServerAwsUseRole

`func (o *DataInnerDnsServerData) HasServerAwsUseRole() bool`

HasServerAwsUseRole returns a boolean if a field has been set.

### GetServerAzGroup

`func (o *DataInnerDnsServerData) GetServerAzGroup() string`

GetServerAzGroup returns the ServerAzGroup field if non-nil, zero value otherwise.

### GetServerAzGroupOk

`func (o *DataInnerDnsServerData) GetServerAzGroupOk() (*string, bool)`

GetServerAzGroupOk returns a tuple with the ServerAzGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAzGroup

`func (o *DataInnerDnsServerData) SetServerAzGroup(v string)`

SetServerAzGroup sets ServerAzGroup field to given value.

### HasServerAzGroup

`func (o *DataInnerDnsServerData) HasServerAzGroup() bool`

HasServerAzGroup returns a boolean if a field has been set.

### GetServerAzKeyid

`func (o *DataInnerDnsServerData) GetServerAzKeyid() string`

GetServerAzKeyid returns the ServerAzKeyid field if non-nil, zero value otherwise.

### GetServerAzKeyidOk

`func (o *DataInnerDnsServerData) GetServerAzKeyidOk() (*string, bool)`

GetServerAzKeyidOk returns a tuple with the ServerAzKeyid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAzKeyid

`func (o *DataInnerDnsServerData) SetServerAzKeyid(v string)`

SetServerAzKeyid sets ServerAzKeyid field to given value.

### HasServerAzKeyid

`func (o *DataInnerDnsServerData) HasServerAzKeyid() bool`

HasServerAzKeyid returns a boolean if a field has been set.

### GetServerAzSubscriptionid

`func (o *DataInnerDnsServerData) GetServerAzSubscriptionid() string`

GetServerAzSubscriptionid returns the ServerAzSubscriptionid field if non-nil, zero value otherwise.

### GetServerAzSubscriptionidOk

`func (o *DataInnerDnsServerData) GetServerAzSubscriptionidOk() (*string, bool)`

GetServerAzSubscriptionidOk returns a tuple with the ServerAzSubscriptionid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAzSubscriptionid

`func (o *DataInnerDnsServerData) SetServerAzSubscriptionid(v string)`

SetServerAzSubscriptionid sets ServerAzSubscriptionid field to given value.

### HasServerAzSubscriptionid

`func (o *DataInnerDnsServerData) HasServerAzSubscriptionid() bool`

HasServerAzSubscriptionid returns a boolean if a field has been set.

### GetServerAzTenantid

`func (o *DataInnerDnsServerData) GetServerAzTenantid() string`

GetServerAzTenantid returns the ServerAzTenantid field if non-nil, zero value otherwise.

### GetServerAzTenantidOk

`func (o *DataInnerDnsServerData) GetServerAzTenantidOk() (*string, bool)`

GetServerAzTenantidOk returns a tuple with the ServerAzTenantid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAzTenantid

`func (o *DataInnerDnsServerData) SetServerAzTenantid(v string)`

SetServerAzTenantid sets ServerAzTenantid field to given value.

### HasServerAzTenantid

`func (o *DataInnerDnsServerData) HasServerAzTenantid() bool`

HasServerAzTenantid returns a boolean if a field has been set.

### GetConnectionprofileName

`func (o *DataInnerDnsServerData) GetConnectionprofileName() string`

GetConnectionprofileName returns the ConnectionprofileName field if non-nil, zero value otherwise.

### GetConnectionprofileNameOk

`func (o *DataInnerDnsServerData) GetConnectionprofileNameOk() (*string, bool)`

GetConnectionprofileNameOk returns a tuple with the ConnectionprofileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionprofileName

`func (o *DataInnerDnsServerData) SetConnectionprofileName(v string)`

SetConnectionprofileName sets ConnectionprofileName field to given value.

### HasConnectionprofileName

`func (o *DataInnerDnsServerData) HasConnectionprofileName() bool`

HasConnectionprofileName returns a boolean if a field has been set.

### GetServerAllowQuery

`func (o *DataInnerDnsServerData) GetServerAllowQuery() string`

GetServerAllowQuery returns the ServerAllowQuery field if non-nil, zero value otherwise.

### GetServerAllowQueryOk

`func (o *DataInnerDnsServerData) GetServerAllowQueryOk() (*string, bool)`

GetServerAllowQueryOk returns a tuple with the ServerAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllowQuery

`func (o *DataInnerDnsServerData) SetServerAllowQuery(v string)`

SetServerAllowQuery sets ServerAllowQuery field to given value.

### HasServerAllowQuery

`func (o *DataInnerDnsServerData) HasServerAllowQuery() bool`

HasServerAllowQuery returns a boolean if a field has been set.

### GetServerAllowQueryCache

`func (o *DataInnerDnsServerData) GetServerAllowQueryCache() string`

GetServerAllowQueryCache returns the ServerAllowQueryCache field if non-nil, zero value otherwise.

### GetServerAllowQueryCacheOk

`func (o *DataInnerDnsServerData) GetServerAllowQueryCacheOk() (*string, bool)`

GetServerAllowQueryCacheOk returns a tuple with the ServerAllowQueryCache field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllowQueryCache

`func (o *DataInnerDnsServerData) SetServerAllowQueryCache(v string)`

SetServerAllowQueryCache sets ServerAllowQueryCache field to given value.

### HasServerAllowQueryCache

`func (o *DataInnerDnsServerData) HasServerAllowQueryCache() bool`

HasServerAllowQueryCache returns a boolean if a field has been set.

### GetServerAllowRecursion

`func (o *DataInnerDnsServerData) GetServerAllowRecursion() string`

GetServerAllowRecursion returns the ServerAllowRecursion field if non-nil, zero value otherwise.

### GetServerAllowRecursionOk

`func (o *DataInnerDnsServerData) GetServerAllowRecursionOk() (*string, bool)`

GetServerAllowRecursionOk returns a tuple with the ServerAllowRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllowRecursion

`func (o *DataInnerDnsServerData) SetServerAllowRecursion(v string)`

SetServerAllowRecursion sets ServerAllowRecursion field to given value.

### HasServerAllowRecursion

`func (o *DataInnerDnsServerData) HasServerAllowRecursion() bool`

HasServerAllowRecursion returns a boolean if a field has been set.

### GetServerAllowTransfer

`func (o *DataInnerDnsServerData) GetServerAllowTransfer() string`

GetServerAllowTransfer returns the ServerAllowTransfer field if non-nil, zero value otherwise.

### GetServerAllowTransferOk

`func (o *DataInnerDnsServerData) GetServerAllowTransferOk() (*string, bool)`

GetServerAllowTransferOk returns a tuple with the ServerAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAllowTransfer

`func (o *DataInnerDnsServerData) SetServerAllowTransfer(v string)`

SetServerAllowTransfer sets ServerAllowTransfer field to given value.

### HasServerAllowTransfer

`func (o *DataInnerDnsServerData) HasServerAllowTransfer() bool`

HasServerAllowTransfer returns a boolean if a field has been set.

### GetServerAlsoNotify

`func (o *DataInnerDnsServerData) GetServerAlsoNotify() string`

GetServerAlsoNotify returns the ServerAlsoNotify field if non-nil, zero value otherwise.

### GetServerAlsoNotifyOk

`func (o *DataInnerDnsServerData) GetServerAlsoNotifyOk() (*string, bool)`

GetServerAlsoNotifyOk returns a tuple with the ServerAlsoNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAlsoNotify

`func (o *DataInnerDnsServerData) SetServerAlsoNotify(v string)`

SetServerAlsoNotify sets ServerAlsoNotify field to given value.

### HasServerAlsoNotify

`func (o *DataInnerDnsServerData) HasServerAlsoNotify() bool`

HasServerAlsoNotify returns a boolean if a field has been set.

### GetServerClassName

`func (o *DataInnerDnsServerData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DataInnerDnsServerData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DataInnerDnsServerData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DataInnerDnsServerData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DataInnerDnsServerData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DataInnerDnsServerData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DataInnerDnsServerData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DataInnerDnsServerData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerCloud

`func (o *DataInnerDnsServerData) GetServerCloud() string`

GetServerCloud returns the ServerCloud field if non-nil, zero value otherwise.

### GetServerCloudOk

`func (o *DataInnerDnsServerData) GetServerCloudOk() (*string, bool)`

GetServerCloudOk returns a tuple with the ServerCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCloud

`func (o *DataInnerDnsServerData) SetServerCloud(v string)`

SetServerCloud sets ServerCloud field to given value.

### HasServerCloud

`func (o *DataInnerDnsServerData) HasServerCloud() bool`

HasServerCloud returns a boolean if a field has been set.

### GetServerCloudPrivate

`func (o *DataInnerDnsServerData) GetServerCloudPrivate() string`

GetServerCloudPrivate returns the ServerCloudPrivate field if non-nil, zero value otherwise.

### GetServerCloudPrivateOk

`func (o *DataInnerDnsServerData) GetServerCloudPrivateOk() (*string, bool)`

GetServerCloudPrivateOk returns a tuple with the ServerCloudPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCloudPrivate

`func (o *DataInnerDnsServerData) SetServerCloudPrivate(v string)`

SetServerCloudPrivate sets ServerCloudPrivate field to given value.

### HasServerCloudPrivate

`func (o *DataInnerDnsServerData) HasServerCloudPrivate() bool`

HasServerCloudPrivate returns a boolean if a field has been set.

### GetServerComment

`func (o *DataInnerDnsServerData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DataInnerDnsServerData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DataInnerDnsServerData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DataInnerDnsServerData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerForceHybrid

`func (o *DataInnerDnsServerData) GetServerForceHybrid() string`

GetServerForceHybrid returns the ServerForceHybrid field if non-nil, zero value otherwise.

### GetServerForceHybridOk

`func (o *DataInnerDnsServerData) GetServerForceHybridOk() (*string, bool)`

GetServerForceHybridOk returns a tuple with the ServerForceHybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerForceHybrid

`func (o *DataInnerDnsServerData) SetServerForceHybrid(v string)`

SetServerForceHybrid sets ServerForceHybrid field to given value.

### HasServerForceHybrid

`func (o *DataInnerDnsServerData) HasServerForceHybrid() bool`

HasServerForceHybrid returns a boolean if a field has been set.

### GetServerForward

`func (o *DataInnerDnsServerData) GetServerForward() string`

GetServerForward returns the ServerForward field if non-nil, zero value otherwise.

### GetServerForwardOk

`func (o *DataInnerDnsServerData) GetServerForwardOk() (*string, bool)`

GetServerForwardOk returns a tuple with the ServerForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerForward

`func (o *DataInnerDnsServerData) SetServerForward(v string)`

SetServerForward sets ServerForward field to given value.

### HasServerForward

`func (o *DataInnerDnsServerData) HasServerForward() bool`

HasServerForward returns a boolean if a field has been set.

### GetServerForwarders

`func (o *DataInnerDnsServerData) GetServerForwarders() string`

GetServerForwarders returns the ServerForwarders field if non-nil, zero value otherwise.

### GetServerForwardersOk

`func (o *DataInnerDnsServerData) GetServerForwardersOk() (*string, bool)`

GetServerForwardersOk returns a tuple with the ServerForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerForwarders

`func (o *DataInnerDnsServerData) SetServerForwarders(v string)`

SetServerForwarders sets ServerForwarders field to given value.

### HasServerForwarders

`func (o *DataInnerDnsServerData) HasServerForwarders() bool`

HasServerForwarders returns a boolean if a field has been set.

### GetServerHybrid

`func (o *DataInnerDnsServerData) GetServerHybrid() string`

GetServerHybrid returns the ServerHybrid field if non-nil, zero value otherwise.

### GetServerHybridOk

`func (o *DataInnerDnsServerData) GetServerHybridOk() (*string, bool)`

GetServerHybridOk returns a tuple with the ServerHybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHybrid

`func (o *DataInnerDnsServerData) SetServerHybrid(v string)`

SetServerHybrid sets ServerHybrid field to given value.

### HasServerHybrid

`func (o *DataInnerDnsServerData) HasServerHybrid() bool`

HasServerHybrid returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDnsServerData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDnsServerData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDnsServerData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDnsServerData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerKeyName

`func (o *DataInnerDnsServerData) GetServerKeyName() string`

GetServerKeyName returns the ServerKeyName field if non-nil, zero value otherwise.

### GetServerKeyNameOk

`func (o *DataInnerDnsServerData) GetServerKeyNameOk() (*string, bool)`

GetServerKeyNameOk returns a tuple with the ServerKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeyName

`func (o *DataInnerDnsServerData) SetServerKeyName(v string)`

SetServerKeyName sets ServerKeyName field to given value.

### HasServerKeyName

`func (o *DataInnerDnsServerData) HasServerKeyName() bool`

HasServerKeyName returns a boolean if a field has been set.

### GetServerKeyProto

`func (o *DataInnerDnsServerData) GetServerKeyProto() string`

GetServerKeyProto returns the ServerKeyProto field if non-nil, zero value otherwise.

### GetServerKeyProtoOk

`func (o *DataInnerDnsServerData) GetServerKeyProtoOk() (*string, bool)`

GetServerKeyProtoOk returns a tuple with the ServerKeyProto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeyProto

`func (o *DataInnerDnsServerData) SetServerKeyProto(v string)`

SetServerKeyProto sets ServerKeyProto field to given value.

### HasServerKeyProto

`func (o *DataInnerDnsServerData) HasServerKeyProto() bool`

HasServerKeyProto returns a boolean if a field has been set.

### GetServerKeyValue

`func (o *DataInnerDnsServerData) GetServerKeyValue() string`

GetServerKeyValue returns the ServerKeyValue field if non-nil, zero value otherwise.

### GetServerKeyValueOk

`func (o *DataInnerDnsServerData) GetServerKeyValueOk() (*string, bool)`

GetServerKeyValueOk returns a tuple with the ServerKeyValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerKeyValue

`func (o *DataInnerDnsServerData) SetServerKeyValue(v string)`

SetServerKeyValue sets ServerKeyValue field to given value.

### HasServerKeyValue

`func (o *DataInnerDnsServerData) HasServerKeyValue() bool`

HasServerKeyValue returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDnsServerData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDnsServerData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDnsServerData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDnsServerData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerNotify

`func (o *DataInnerDnsServerData) GetServerNotify() string`

GetServerNotify returns the ServerNotify field if non-nil, zero value otherwise.

### GetServerNotifyOk

`func (o *DataInnerDnsServerData) GetServerNotifyOk() (*string, bool)`

GetServerNotifyOk returns a tuple with the ServerNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerNotify

`func (o *DataInnerDnsServerData) SetServerNotify(v string)`

SetServerNotify sets ServerNotify field to given value.

### HasServerNotify

`func (o *DataInnerDnsServerData) HasServerNotify() bool`

HasServerNotify returns a boolean if a field has been set.

### GetServerRecursion

`func (o *DataInnerDnsServerData) GetServerRecursion() string`

GetServerRecursion returns the ServerRecursion field if non-nil, zero value otherwise.

### GetServerRecursionOk

`func (o *DataInnerDnsServerData) GetServerRecursionOk() (*string, bool)`

GetServerRecursionOk returns a tuple with the ServerRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRecursion

`func (o *DataInnerDnsServerData) SetServerRecursion(v string)`

SetServerRecursion sets ServerRecursion field to given value.

### HasServerRecursion

`func (o *DataInnerDnsServerData) HasServerRecursion() bool`

HasServerRecursion returns a boolean if a field has been set.

### GetServerRole

`func (o *DataInnerDnsServerData) GetServerRole() string`

GetServerRole returns the ServerRole field if non-nil, zero value otherwise.

### GetServerRoleOk

`func (o *DataInnerDnsServerData) GetServerRoleOk() (*string, bool)`

GetServerRoleOk returns a tuple with the ServerRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRole

`func (o *DataInnerDnsServerData) SetServerRole(v string)`

SetServerRole sets ServerRole field to given value.

### HasServerRole

`func (o *DataInnerDnsServerData) HasServerRole() bool`

HasServerRole returns a boolean if a field has been set.

### GetServerRpzBreakDnssec

`func (o *DataInnerDnsServerData) GetServerRpzBreakDnssec() string`

GetServerRpzBreakDnssec returns the ServerRpzBreakDnssec field if non-nil, zero value otherwise.

### GetServerRpzBreakDnssecOk

`func (o *DataInnerDnsServerData) GetServerRpzBreakDnssecOk() (*string, bool)`

GetServerRpzBreakDnssecOk returns a tuple with the ServerRpzBreakDnssec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRpzBreakDnssec

`func (o *DataInnerDnsServerData) SetServerRpzBreakDnssec(v string)`

SetServerRpzBreakDnssec sets ServerRpzBreakDnssec field to given value.

### HasServerRpzBreakDnssec

`func (o *DataInnerDnsServerData) HasServerRpzBreakDnssec() bool`

HasServerRpzBreakDnssec returns a boolean if a field has been set.

### GetServerRpzMaxPolicyTtl

`func (o *DataInnerDnsServerData) GetServerRpzMaxPolicyTtl() string`

GetServerRpzMaxPolicyTtl returns the ServerRpzMaxPolicyTtl field if non-nil, zero value otherwise.

### GetServerRpzMaxPolicyTtlOk

`func (o *DataInnerDnsServerData) GetServerRpzMaxPolicyTtlOk() (*string, bool)`

GetServerRpzMaxPolicyTtlOk returns a tuple with the ServerRpzMaxPolicyTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRpzMaxPolicyTtl

`func (o *DataInnerDnsServerData) SetServerRpzMaxPolicyTtl(v string)`

SetServerRpzMaxPolicyTtl sets ServerRpzMaxPolicyTtl field to given value.

### HasServerRpzMaxPolicyTtl

`func (o *DataInnerDnsServerData) HasServerRpzMaxPolicyTtl() bool`

HasServerRpzMaxPolicyTtl returns a boolean if a field has been set.

### GetServerRpzMinNsDots

`func (o *DataInnerDnsServerData) GetServerRpzMinNsDots() string`

GetServerRpzMinNsDots returns the ServerRpzMinNsDots field if non-nil, zero value otherwise.

### GetServerRpzMinNsDotsOk

`func (o *DataInnerDnsServerData) GetServerRpzMinNsDotsOk() (*string, bool)`

GetServerRpzMinNsDotsOk returns a tuple with the ServerRpzMinNsDots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRpzMinNsDots

`func (o *DataInnerDnsServerData) SetServerRpzMinNsDots(v string)`

SetServerRpzMinNsDots sets ServerRpzMinNsDots field to given value.

### HasServerRpzMinNsDots

`func (o *DataInnerDnsServerData) HasServerRpzMinNsDots() bool`

HasServerRpzMinNsDots returns a boolean if a field has been set.

### GetServerRpzQnameWaitRecurse

`func (o *DataInnerDnsServerData) GetServerRpzQnameWaitRecurse() string`

GetServerRpzQnameWaitRecurse returns the ServerRpzQnameWaitRecurse field if non-nil, zero value otherwise.

### GetServerRpzQnameWaitRecurseOk

`func (o *DataInnerDnsServerData) GetServerRpzQnameWaitRecurseOk() (*string, bool)`

GetServerRpzQnameWaitRecurseOk returns a tuple with the ServerRpzQnameWaitRecurse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRpzQnameWaitRecurse

`func (o *DataInnerDnsServerData) SetServerRpzQnameWaitRecurse(v string)`

SetServerRpzQnameWaitRecurse sets ServerRpzQnameWaitRecurse field to given value.

### HasServerRpzQnameWaitRecurse

`func (o *DataInnerDnsServerData) HasServerRpzQnameWaitRecurse() bool`

HasServerRpzQnameWaitRecurse returns a boolean if a field has been set.

### GetServerRpzRecursiveOnly

`func (o *DataInnerDnsServerData) GetServerRpzRecursiveOnly() string`

GetServerRpzRecursiveOnly returns the ServerRpzRecursiveOnly field if non-nil, zero value otherwise.

### GetServerRpzRecursiveOnlyOk

`func (o *DataInnerDnsServerData) GetServerRpzRecursiveOnlyOk() (*string, bool)`

GetServerRpzRecursiveOnlyOk returns a tuple with the ServerRpzRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerRpzRecursiveOnly

`func (o *DataInnerDnsServerData) SetServerRpzRecursiveOnly(v string)`

SetServerRpzRecursiveOnly sets ServerRpzRecursiveOnly field to given value.

### HasServerRpzRecursiveOnly

`func (o *DataInnerDnsServerData) HasServerRpzRecursiveOnly() bool`

HasServerRpzRecursiveOnly returns a boolean if a field has been set.

### GetServerState

`func (o *DataInnerDnsServerData) GetServerState() string`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *DataInnerDnsServerData) GetServerStateOk() (*string, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *DataInnerDnsServerData) SetServerState(v string)`

SetServerState sets ServerState field to given value.

### HasServerState

`func (o *DataInnerDnsServerData) HasServerState() bool`

HasServerState returns a boolean if a field has been set.

### GetServerSynching

`func (o *DataInnerDnsServerData) GetServerSynching() string`

GetServerSynching returns the ServerSynching field if non-nil, zero value otherwise.

### GetServerSynchingOk

`func (o *DataInnerDnsServerData) GetServerSynchingOk() (*string, bool)`

GetServerSynchingOk returns a tuple with the ServerSynching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSynching

`func (o *DataInnerDnsServerData) SetServerSynching(v string)`

SetServerSynching sets ServerSynching field to given value.

### HasServerSynching

`func (o *DataInnerDnsServerData) HasServerSynching() bool`

HasServerSynching returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDnsServerData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDnsServerData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDnsServerData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDnsServerData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DataInnerDnsServerData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DataInnerDnsServerData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DataInnerDnsServerData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DataInnerDnsServerData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerVpcList

`func (o *DataInnerDnsServerData) GetServerVpcList() string`

GetServerVpcList returns the ServerVpcList field if non-nil, zero value otherwise.

### GetServerVpcListOk

`func (o *DataInnerDnsServerData) GetServerVpcListOk() (*string, bool)`

GetServerVpcListOk returns a tuple with the ServerVpcList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVpcList

`func (o *DataInnerDnsServerData) SetServerVpcList(v string)`

SetServerVpcList sets ServerVpcList field to given value.

### HasServerVpcList

`func (o *DataInnerDnsServerData) HasServerVpcList() bool`

HasServerVpcList returns a boolean if a field has been set.

### GetServerBlastEnabled

`func (o *DataInnerDnsServerData) GetServerBlastEnabled() string`

GetServerBlastEnabled returns the ServerBlastEnabled field if non-nil, zero value otherwise.

### GetServerBlastEnabledOk

`func (o *DataInnerDnsServerData) GetServerBlastEnabledOk() (*string, bool)`

GetServerBlastEnabledOk returns a tuple with the ServerBlastEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBlastEnabled

`func (o *DataInnerDnsServerData) SetServerBlastEnabled(v string)`

SetServerBlastEnabled sets ServerBlastEnabled field to given value.

### HasServerBlastEnabled

`func (o *DataInnerDnsServerData) HasServerBlastEnabled() bool`

HasServerBlastEnabled returns a boolean if a field has been set.

### GetServerGuardianPushStatus

`func (o *DataInnerDnsServerData) GetServerGuardianPushStatus() string`

GetServerGuardianPushStatus returns the ServerGuardianPushStatus field if non-nil, zero value otherwise.

### GetServerGuardianPushStatusOk

`func (o *DataInnerDnsServerData) GetServerGuardianPushStatusOk() (*string, bool)`

GetServerGuardianPushStatusOk returns a tuple with the ServerGuardianPushStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGuardianPushStatus

`func (o *DataInnerDnsServerData) SetServerGuardianPushStatus(v string)`

SetServerGuardianPushStatus sets ServerGuardianPushStatus field to given value.

### HasServerGuardianPushStatus

`func (o *DataInnerDnsServerData) HasServerGuardianPushStatus() bool`

HasServerGuardianPushStatus returns a boolean if a field has been set.

### GetServerBlastStatus

`func (o *DataInnerDnsServerData) GetServerBlastStatus() string`

GetServerBlastStatus returns the ServerBlastStatus field if non-nil, zero value otherwise.

### GetServerBlastStatusOk

`func (o *DataInnerDnsServerData) GetServerBlastStatusOk() (*string, bool)`

GetServerBlastStatusOk returns a tuple with the ServerBlastStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerBlastStatus

`func (o *DataInnerDnsServerData) SetServerBlastStatus(v string)`

SetServerBlastStatus sets ServerBlastStatus field to given value.

### HasServerBlastStatus

`func (o *DataInnerDnsServerData) HasServerBlastStatus() bool`

HasServerBlastStatus returns a boolean if a field has been set.

### GetServerGslbSupported

`func (o *DataInnerDnsServerData) GetServerGslbSupported() string`

GetServerGslbSupported returns the ServerGslbSupported field if non-nil, zero value otherwise.

### GetServerGslbSupportedOk

`func (o *DataInnerDnsServerData) GetServerGslbSupportedOk() (*string, bool)`

GetServerGslbSupportedOk returns a tuple with the ServerGslbSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGslbSupported

`func (o *DataInnerDnsServerData) SetServerGslbSupported(v string)`

SetServerGslbSupported sets ServerGslbSupported field to given value.

### HasServerGslbSupported

`func (o *DataInnerDnsServerData) HasServerGslbSupported() bool`

HasServerGslbSupported returns a boolean if a field has been set.

### GetServerGuardianGuiManagementSupported

`func (o *DataInnerDnsServerData) GetServerGuardianGuiManagementSupported() string`

GetServerGuardianGuiManagementSupported returns the ServerGuardianGuiManagementSupported field if non-nil, zero value otherwise.

### GetServerGuardianGuiManagementSupportedOk

`func (o *DataInnerDnsServerData) GetServerGuardianGuiManagementSupportedOk() (*string, bool)`

GetServerGuardianGuiManagementSupportedOk returns a tuple with the ServerGuardianGuiManagementSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGuardianGuiManagementSupported

`func (o *DataInnerDnsServerData) SetServerGuardianGuiManagementSupported(v string)`

SetServerGuardianGuiManagementSupported sets ServerGuardianGuiManagementSupported field to given value.

### HasServerGuardianGuiManagementSupported

`func (o *DataInnerDnsServerData) HasServerGuardianGuiManagementSupported() bool`

HasServerGuardianGuiManagementSupported returns a boolean if a field has been set.

### GetServerGuardianSupported

`func (o *DataInnerDnsServerData) GetServerGuardianSupported() string`

GetServerGuardianSupported returns the ServerGuardianSupported field if non-nil, zero value otherwise.

### GetServerGuardianSupportedOk

`func (o *DataInnerDnsServerData) GetServerGuardianSupportedOk() (*string, bool)`

GetServerGuardianSupportedOk returns a tuple with the ServerGuardianSupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGuardianSupported

`func (o *DataInnerDnsServerData) SetServerGuardianSupported(v string)`

SetServerGuardianSupported sets ServerGuardianSupported field to given value.

### HasServerGuardianSupported

`func (o *DataInnerDnsServerData) HasServerGuardianSupported() bool`

HasServerGuardianSupported returns a boolean if a field has been set.

### GetServerDnssecValidation

`func (o *DataInnerDnsServerData) GetServerDnssecValidation() string`

GetServerDnssecValidation returns the ServerDnssecValidation field if non-nil, zero value otherwise.

### GetServerDnssecValidationOk

`func (o *DataInnerDnsServerData) GetServerDnssecValidationOk() (*string, bool)`

GetServerDnssecValidationOk returns a tuple with the ServerDnssecValidation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDnssecValidation

`func (o *DataInnerDnsServerData) SetServerDnssecValidation(v string)`

SetServerDnssecValidation sets ServerDnssecValidation field to given value.

### HasServerDnssecValidation

`func (o *DataInnerDnsServerData) HasServerDnssecValidation() bool`

HasServerDnssecValidation returns a boolean if a field has been set.

### GetServerGssEnabled

`func (o *DataInnerDnsServerData) GetServerGssEnabled() string`

GetServerGssEnabled returns the ServerGssEnabled field if non-nil, zero value otherwise.

### GetServerGssEnabledOk

`func (o *DataInnerDnsServerData) GetServerGssEnabledOk() (*string, bool)`

GetServerGssEnabledOk returns a tuple with the ServerGssEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGssEnabled

`func (o *DataInnerDnsServerData) SetServerGssEnabled(v string)`

SetServerGssEnabled sets ServerGssEnabled field to given value.

### HasServerGssEnabled

`func (o *DataInnerDnsServerData) HasServerGssEnabled() bool`

HasServerGssEnabled returns a boolean if a field has been set.

### GetServerGssKeytabId

`func (o *DataInnerDnsServerData) GetServerGssKeytabId() string`

GetServerGssKeytabId returns the ServerGssKeytabId field if non-nil, zero value otherwise.

### GetServerGssKeytabIdOk

`func (o *DataInnerDnsServerData) GetServerGssKeytabIdOk() (*string, bool)`

GetServerGssKeytabIdOk returns a tuple with the ServerGssKeytabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGssKeytabId

`func (o *DataInnerDnsServerData) SetServerGssKeytabId(v string)`

SetServerGssKeytabId sets ServerGssKeytabId field to given value.

### HasServerGssKeytabId

`func (o *DataInnerDnsServerData) HasServerGssKeytabId() bool`

HasServerGssKeytabId returns a boolean if a field has been set.

### GetServerGuardianStatsOnlySupported

`func (o *DataInnerDnsServerData) GetServerGuardianStatsOnlySupported() string`

GetServerGuardianStatsOnlySupported returns the ServerGuardianStatsOnlySupported field if non-nil, zero value otherwise.

### GetServerGuardianStatsOnlySupportedOk

`func (o *DataInnerDnsServerData) GetServerGuardianStatsOnlySupportedOk() (*string, bool)`

GetServerGuardianStatsOnlySupportedOk returns a tuple with the ServerGuardianStatsOnlySupported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGuardianStatsOnlySupported

`func (o *DataInnerDnsServerData) SetServerGuardianStatsOnlySupported(v string)`

SetServerGuardianStatsOnlySupported sets ServerGuardianStatsOnlySupported field to given value.

### HasServerGuardianStatsOnlySupported

`func (o *DataInnerDnsServerData) HasServerGuardianStatsOnlySupported() bool`

HasServerGuardianStatsOnlySupported returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DataInnerDnsServerData) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DataInnerDnsServerData) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DataInnerDnsServerData) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DataInnerDnsServerData) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerAddr6

`func (o *DataInnerDnsServerData) GetServerAddr6() string`

GetServerAddr6 returns the ServerAddr6 field if non-nil, zero value otherwise.

### GetServerAddr6Ok

`func (o *DataInnerDnsServerData) GetServerAddr6Ok() (*string, bool)`

GetServerAddr6Ok returns a tuple with the ServerAddr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr6

`func (o *DataInnerDnsServerData) SetServerAddr6(v string)`

SetServerAddr6 sets ServerAddr6 field to given value.

### HasServerAddr6

`func (o *DataInnerDnsServerData) HasServerAddr6() bool`

HasServerAddr6 returns a boolean if a field has been set.

### GetServerAddr

`func (o *DataInnerDnsServerData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DataInnerDnsServerData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DataInnerDnsServerData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DataInnerDnsServerData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetServerIpmLogin

`func (o *DataInnerDnsServerData) GetServerIpmLogin() string`

GetServerIpmLogin returns the ServerIpmLogin field if non-nil, zero value otherwise.

### GetServerIpmLoginOk

`func (o *DataInnerDnsServerData) GetServerIpmLoginOk() (*string, bool)`

GetServerIpmLoginOk returns a tuple with the ServerIpmLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmLogin

`func (o *DataInnerDnsServerData) SetServerIpmLogin(v string)`

SetServerIpmLogin sets ServerIpmLogin field to given value.

### HasServerIpmLogin

`func (o *DataInnerDnsServerData) HasServerIpmLogin() bool`

HasServerIpmLogin returns a boolean if a field has been set.

### GetServerIpmIsPackage

`func (o *DataInnerDnsServerData) GetServerIpmIsPackage() string`

GetServerIpmIsPackage returns the ServerIpmIsPackage field if non-nil, zero value otherwise.

### GetServerIpmIsPackageOk

`func (o *DataInnerDnsServerData) GetServerIpmIsPackageOk() (*string, bool)`

GetServerIpmIsPackageOk returns a tuple with the ServerIpmIsPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmIsPackage

`func (o *DataInnerDnsServerData) SetServerIpmIsPackage(v string)`

SetServerIpmIsPackage sets ServerIpmIsPackage field to given value.

### HasServerIpmIsPackage

`func (o *DataInnerDnsServerData) HasServerIpmIsPackage() bool`

HasServerIpmIsPackage returns a boolean if a field has been set.

### GetServerIpmProtocol

`func (o *DataInnerDnsServerData) GetServerIpmProtocol() string`

GetServerIpmProtocol returns the ServerIpmProtocol field if non-nil, zero value otherwise.

### GetServerIpmProtocolOk

`func (o *DataInnerDnsServerData) GetServerIpmProtocolOk() (*string, bool)`

GetServerIpmProtocolOk returns a tuple with the ServerIpmProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmProtocol

`func (o *DataInnerDnsServerData) SetServerIpmProtocol(v string)`

SetServerIpmProtocol sets ServerIpmProtocol field to given value.

### HasServerIpmProtocol

`func (o *DataInnerDnsServerData) HasServerIpmProtocol() bool`

HasServerIpmProtocol returns a boolean if a field has been set.

### GetServerIpmType

`func (o *DataInnerDnsServerData) GetServerIpmType() string`

GetServerIpmType returns the ServerIpmType field if non-nil, zero value otherwise.

### GetServerIpmTypeOk

`func (o *DataInnerDnsServerData) GetServerIpmTypeOk() (*string, bool)`

GetServerIpmTypeOk returns a tuple with the ServerIpmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmType

`func (o *DataInnerDnsServerData) SetServerIpmType(v string)`

SetServerIpmType sets ServerIpmType field to given value.

### HasServerIpmType

`func (o *DataInnerDnsServerData) HasServerIpmType() bool`

HasServerIpmType returns a boolean if a field has been set.

### GetServerIsolated

`func (o *DataInnerDnsServerData) GetServerIsolated() string`

GetServerIsolated returns the ServerIsolated field if non-nil, zero value otherwise.

### GetServerIsolatedOk

`func (o *DataInnerDnsServerData) GetServerIsolatedOk() (*string, bool)`

GetServerIsolatedOk returns a tuple with the ServerIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIsolated

`func (o *DataInnerDnsServerData) SetServerIsolated(v string)`

SetServerIsolated sets ServerIsolated field to given value.

### HasServerIsolated

`func (o *DataInnerDnsServerData) HasServerIsolated() bool`

HasServerIsolated returns a boolean if a field has been set.

### GetServerLdapDomain

`func (o *DataInnerDnsServerData) GetServerLdapDomain() string`

GetServerLdapDomain returns the ServerLdapDomain field if non-nil, zero value otherwise.

### GetServerLdapDomainOk

`func (o *DataInnerDnsServerData) GetServerLdapDomainOk() (*string, bool)`

GetServerLdapDomainOk returns a tuple with the ServerLdapDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLdapDomain

`func (o *DataInnerDnsServerData) SetServerLdapDomain(v string)`

SetServerLdapDomain sets ServerLdapDomain field to given value.

### HasServerLdapDomain

`func (o *DataInnerDnsServerData) HasServerLdapDomain() bool`

HasServerLdapDomain returns a boolean if a field has been set.

### GetServerLdapUser

`func (o *DataInnerDnsServerData) GetServerLdapUser() string`

GetServerLdapUser returns the ServerLdapUser field if non-nil, zero value otherwise.

### GetServerLdapUserOk

`func (o *DataInnerDnsServerData) GetServerLdapUserOk() (*string, bool)`

GetServerLdapUserOk returns a tuple with the ServerLdapUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLdapUser

`func (o *DataInnerDnsServerData) SetServerLdapUser(v string)`

SetServerLdapUser sets ServerLdapUser field to given value.

### HasServerLdapUser

`func (o *DataInnerDnsServerData) HasServerLdapUser() bool`

HasServerLdapUser returns a boolean if a field has been set.

### GetServerMultistatus

`func (o *DataInnerDnsServerData) GetServerMultistatus() string`

GetServerMultistatus returns the ServerMultistatus field if non-nil, zero value otherwise.

### GetServerMultistatusOk

`func (o *DataInnerDnsServerData) GetServerMultistatusOk() (*string, bool)`

GetServerMultistatusOk returns a tuple with the ServerMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMultistatus

`func (o *DataInnerDnsServerData) SetServerMultistatus(v string)`

SetServerMultistatus sets ServerMultistatus field to given value.

### HasServerMultistatus

`func (o *DataInnerDnsServerData) HasServerMultistatus() bool`

HasServerMultistatus returns a boolean if a field has been set.

### GetServerQuerylogState

`func (o *DataInnerDnsServerData) GetServerQuerylogState() string`

GetServerQuerylogState returns the ServerQuerylogState field if non-nil, zero value otherwise.

### GetServerQuerylogStateOk

`func (o *DataInnerDnsServerData) GetServerQuerylogStateOk() (*string, bool)`

GetServerQuerylogStateOk returns a tuple with the ServerQuerylogState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerQuerylogState

`func (o *DataInnerDnsServerData) SetServerQuerylogState(v string)`

SetServerQuerylogState sets ServerQuerylogState field to given value.

### HasServerQuerylogState

`func (o *DataInnerDnsServerData) HasServerQuerylogState() bool`

HasServerQuerylogState returns a boolean if a field has been set.

### GetReverseProxyConf

`func (o *DataInnerDnsServerData) GetReverseProxyConf() string`

GetReverseProxyConf returns the ReverseProxyConf field if non-nil, zero value otherwise.

### GetReverseProxyConfOk

`func (o *DataInnerDnsServerData) GetReverseProxyConfOk() (*string, bool)`

GetReverseProxyConfOk returns a tuple with the ReverseProxyConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseProxyConf

`func (o *DataInnerDnsServerData) SetReverseProxyConf(v string)`

SetReverseProxyConf sets ReverseProxyConf field to given value.

### HasReverseProxyConf

`func (o *DataInnerDnsServerData) HasReverseProxyConf() bool`

HasReverseProxyConf returns a boolean if a field has been set.

### GetServerSnmpId

`func (o *DataInnerDnsServerData) GetServerSnmpId() string`

GetServerSnmpId returns the ServerSnmpId field if non-nil, zero value otherwise.

### GetServerSnmpIdOk

`func (o *DataInnerDnsServerData) GetServerSnmpIdOk() (*string, bool)`

GetServerSnmpIdOk returns a tuple with the ServerSnmpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpId

`func (o *DataInnerDnsServerData) SetServerSnmpId(v string)`

SetServerSnmpId sets ServerSnmpId field to given value.

### HasServerSnmpId

`func (o *DataInnerDnsServerData) HasServerSnmpId() bool`

HasServerSnmpId returns a boolean if a field has been set.

### GetServerStatEnabled

`func (o *DataInnerDnsServerData) GetServerStatEnabled() string`

GetServerStatEnabled returns the ServerStatEnabled field if non-nil, zero value otherwise.

### GetServerStatEnabledOk

`func (o *DataInnerDnsServerData) GetServerStatEnabledOk() (*string, bool)`

GetServerStatEnabledOk returns a tuple with the ServerStatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatEnabled

`func (o *DataInnerDnsServerData) SetServerStatEnabled(v string)`

SetServerStatEnabled sets ServerStatEnabled field to given value.

### HasServerStatEnabled

`func (o *DataInnerDnsServerData) HasServerStatEnabled() bool`

HasServerStatEnabled returns a boolean if a field has been set.

### GetServerStatNiceness

`func (o *DataInnerDnsServerData) GetServerStatNiceness() string`

GetServerStatNiceness returns the ServerStatNiceness field if non-nil, zero value otherwise.

### GetServerStatNicenessOk

`func (o *DataInnerDnsServerData) GetServerStatNicenessOk() (*string, bool)`

GetServerStatNicenessOk returns a tuple with the ServerStatNiceness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatNiceness

`func (o *DataInnerDnsServerData) SetServerStatNiceness(v string)`

SetServerStatNiceness sets ServerStatNiceness field to given value.

### HasServerStatNiceness

`func (o *DataInnerDnsServerData) HasServerStatNiceness() bool`

HasServerStatNiceness returns a boolean if a field has been set.

### GetServerStatPeriod

`func (o *DataInnerDnsServerData) GetServerStatPeriod() string`

GetServerStatPeriod returns the ServerStatPeriod field if non-nil, zero value otherwise.

### GetServerStatPeriodOk

`func (o *DataInnerDnsServerData) GetServerStatPeriodOk() (*string, bool)`

GetServerStatPeriodOk returns a tuple with the ServerStatPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatPeriod

`func (o *DataInnerDnsServerData) SetServerStatPeriod(v string)`

SetServerStatPeriod sets ServerStatPeriod field to given value.

### HasServerStatPeriod

`func (o *DataInnerDnsServerData) HasServerStatPeriod() bool`

HasServerStatPeriod returns a boolean if a field has been set.

### GetServerStatTime

`func (o *DataInnerDnsServerData) GetServerStatTime() string`

GetServerStatTime returns the ServerStatTime field if non-nil, zero value otherwise.

### GetServerStatTimeOk

`func (o *DataInnerDnsServerData) GetServerStatTimeOk() (*string, bool)`

GetServerStatTimeOk returns a tuple with the ServerStatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatTime

`func (o *DataInnerDnsServerData) SetServerStatTime(v string)`

SetServerStatTime sets ServerStatTime field to given value.

### HasServerStatTime

`func (o *DataInnerDnsServerData) HasServerStatTime() bool`

HasServerStatTime returns a boolean if a field has been set.

### GetTotalSmartMembers

`func (o *DataInnerDnsServerData) GetTotalSmartMembers() string`

GetTotalSmartMembers returns the TotalSmartMembers field if non-nil, zero value otherwise.

### GetTotalSmartMembersOk

`func (o *DataInnerDnsServerData) GetTotalSmartMembersOk() (*string, bool)`

GetTotalSmartMembersOk returns a tuple with the TotalSmartMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSmartMembers

`func (o *DataInnerDnsServerData) SetTotalSmartMembers(v string)`

SetTotalSmartMembers sets TotalSmartMembers field to given value.

### HasTotalSmartMembers

`func (o *DataInnerDnsServerData) HasTotalSmartMembers() bool`

HasTotalSmartMembers returns a boolean if a field has been set.

### GetSmartArch

`func (o *DataInnerDnsServerData) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DataInnerDnsServerData) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DataInnerDnsServerData) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DataInnerDnsServerData) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartMembersName

`func (o *DataInnerDnsServerData) GetSmartMembersName() string`

GetSmartMembersName returns the SmartMembersName field if non-nil, zero value otherwise.

### GetSmartMembersNameOk

`func (o *DataInnerDnsServerData) GetSmartMembersNameOk() (*string, bool)`

GetSmartMembersNameOk returns a tuple with the SmartMembersName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartMembersName

`func (o *DataInnerDnsServerData) SetSmartMembersName(v string)`

SetSmartMembersName sets SmartMembersName field to given value.

### HasSmartMembersName

`func (o *DataInnerDnsServerData) HasSmartMembersName() bool`

HasSmartMembersName returns a boolean if a field has been set.

### GetSmartParentArch

`func (o *DataInnerDnsServerData) GetSmartParentArch() string`

GetSmartParentArch returns the SmartParentArch field if non-nil, zero value otherwise.

### GetSmartParentArchOk

`func (o *DataInnerDnsServerData) GetSmartParentArchOk() (*string, bool)`

GetSmartParentArchOk returns a tuple with the SmartParentArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentArch

`func (o *DataInnerDnsServerData) SetSmartParentArch(v string)`

SetSmartParentArch sets SmartParentArch field to given value.

### HasSmartParentArch

`func (o *DataInnerDnsServerData) HasSmartParentArch() bool`

HasSmartParentArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDnsServerData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDnsServerData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDnsServerData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDnsServerData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDnsServerData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDnsServerData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDnsServerData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDnsServerData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.

### GetSmartPublicNsList

`func (o *DataInnerDnsServerData) GetSmartPublicNsList() string`

GetSmartPublicNsList returns the SmartPublicNsList field if non-nil, zero value otherwise.

### GetSmartPublicNsListOk

`func (o *DataInnerDnsServerData) GetSmartPublicNsListOk() (*string, bool)`

GetSmartPublicNsListOk returns a tuple with the SmartPublicNsList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartPublicNsList

`func (o *DataInnerDnsServerData) SetSmartPublicNsList(v string)`

SetSmartPublicNsList sets SmartPublicNsList field to given value.

### HasSmartPublicNsList

`func (o *DataInnerDnsServerData) HasSmartPublicNsList() bool`

HasSmartPublicNsList returns a boolean if a field has been set.

### GetServerWindnsPort

`func (o *DataInnerDnsServerData) GetServerWindnsPort() string`

GetServerWindnsPort returns the ServerWindnsPort field if non-nil, zero value otherwise.

### GetServerWindnsPortOk

`func (o *DataInnerDnsServerData) GetServerWindnsPortOk() (*string, bool)`

GetServerWindnsPortOk returns a tuple with the ServerWindnsPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWindnsPort

`func (o *DataInnerDnsServerData) SetServerWindnsPort(v string)`

SetServerWindnsPort sets ServerWindnsPort field to given value.

### HasServerWindnsPort

`func (o *DataInnerDnsServerData) HasServerWindnsPort() bool`

HasServerWindnsPort returns a boolean if a field has been set.

### GetServerWindnsProtocol

`func (o *DataInnerDnsServerData) GetServerWindnsProtocol() string`

GetServerWindnsProtocol returns the ServerWindnsProtocol field if non-nil, zero value otherwise.

### GetServerWindnsProtocolOk

`func (o *DataInnerDnsServerData) GetServerWindnsProtocolOk() (*string, bool)`

GetServerWindnsProtocolOk returns a tuple with the ServerWindnsProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWindnsProtocol

`func (o *DataInnerDnsServerData) SetServerWindnsProtocol(v string)`

SetServerWindnsProtocol sets ServerWindnsProtocol field to given value.

### HasServerWindnsProtocol

`func (o *DataInnerDnsServerData) HasServerWindnsProtocol() bool`

HasServerWindnsProtocol returns a boolean if a field has been set.

### GetServerWindnsUseSsl

`func (o *DataInnerDnsServerData) GetServerWindnsUseSsl() string`

GetServerWindnsUseSsl returns the ServerWindnsUseSsl field if non-nil, zero value otherwise.

### GetServerWindnsUseSslOk

`func (o *DataInnerDnsServerData) GetServerWindnsUseSslOk() (*string, bool)`

GetServerWindnsUseSslOk returns a tuple with the ServerWindnsUseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWindnsUseSsl

`func (o *DataInnerDnsServerData) SetServerWindnsUseSsl(v string)`

SetServerWindnsUseSsl sets ServerWindnsUseSsl field to given value.

### HasServerWindnsUseSsl

`func (o *DataInnerDnsServerData) HasServerWindnsUseSsl() bool`

HasServerWindnsUseSsl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


