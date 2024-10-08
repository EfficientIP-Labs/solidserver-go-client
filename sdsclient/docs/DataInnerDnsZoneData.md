# DataInnerDnsZoneData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerAwsDelegationSet** | Pointer to **string** | The reusable delegation set ID configured on the Amazon Route 53 public server the zone belongs to. | [optional] 
**ServerDdnsScavenging** | Pointer to **string** | The DDNS scavenging status of the zone, either enabled (&lt;b&gt;1&lt;/b&gt;) or disabled (&lt;b&gt;0&lt;/b&gt;). DDNS scavenging is only effective if the parameters &lt;b&gt;zone_use_update_policy&lt;/b&gt; and &lt;b&gt;server_gss_enabled&lt;/b&gt; are set to &lt;b&gt;1&lt;/b&gt;, and the rule 416 is enabled in the GUI. | [optional] 
**ZoneDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**ZoneDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DNS server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DNS server the object belongs to. | [optional] 
**ServerCloud** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerComment** | Pointer to **string** | The description of the DNS server the object belongs to. | [optional] 
**ServerForceHybrid** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server the object belongs to. | [optional] 
**ServerState** | Pointer to **string** | The status of the DNS server the object belongs to. &lt;table&gt;&lt;caption&gt;server_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ER&lt;/td&gt;&lt;td &gt;The license used in SOLIDserver is not compliant with the added server: the license is invalid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ES&lt;/td&gt;&lt;td &gt;The server configuration could not be parsed properly.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ET&lt;/td&gt;&lt;td &gt;The server does not answer anymore due to a scheduled configuration of the server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IC&lt;/td&gt;&lt;td &gt;The SSL credentials are invalid&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IP&lt;/td&gt;&lt;td &gt;The provided account does not have sufficient privileges to remotely manage the MS server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IR&lt;/td&gt;&lt;td &gt;SOLIDserver cannot resolve the AWS DNS service. The Amazon services are unreachable and the Amazon Route 53 server cannot be managed. Make sure that the DNS resolvers declared on the page are valid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IS&lt;/td&gt;&lt;td &gt;There was a setting error during the server declaration. For instance, some settings were added to a server that does not support them or a smart architecture is not managing any physical server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IT&lt;/td&gt;&lt;td &gt;The server editing performed from the GUI is not pushed to the server because SOLIDserver time and date are incorrect. You must use the UTC system on the appliance, especially when managing Amazon Route 53 servers.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;LS&lt;/td&gt;&lt;td &gt;The server configuration is not viable.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;N&lt;/td&gt;&lt;td &gt;The server does not have a status as it has not synchronized yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;UE&lt;/td&gt;&lt;td &gt;An error occurred that SOLIDserver could not identify.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;Y&lt;/td&gt;&lt;td &gt;The server is operational.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerType** | Pointer to **string** | The type of the DNS server the object belongs to.&lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msdaemon&lt;/td&gt;&lt;td &gt;Microsoft Windows DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;aws&lt;/td&gt;&lt;td &gt;Amazon Route 53 server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;other&lt;/td&gt;&lt;td &gt;Generic DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdns&lt;/td&gt;&lt;td &gt;EfficientIP DNS smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DNS server the object belongs to. | [optional] 
**ServerVpcList** | Pointer to **string** | The list of cloud networks configured on the DNS private server the zone belongs to, separated by a comma:&lt;ul class&#x3D;dashed &gt;&lt;li&gt; For an Azure server, it returns the list of virtual networks.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; For an Amazon Route 53 server, it returns the list of Virtual Private Cloud (VPC).&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**ViewClassName** | Pointer to **string** | The name of the class applied to the DNS view the object belongs to, it can be preceded by the class directory. | [optional] 
**ViewClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DNS view the object belongs to. | [optional] 
**ViewId** | Pointer to **string** | The database identifier (ID) of the DNS view the object belongs to. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view the object belongs to. | [optional] 
**ZoneAdIntegrated** | Pointer to **string** | The AD integrated status of the DNS zone. &lt;b&gt;1&lt;/b&gt; indicates that the DNS zone belongs to an Active Directory integrated Microsoft Windows DNS server. | [optional] 
**ZoneAllowQuery** | Pointer to **string** | The ACL values associated with the allow-query configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAllowTransfer** | Pointer to **string** | The ACL values associated with the allow-transfer configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAllowUpdate** | Pointer to **string** | The ACL values associated with the allow-update configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAlsoNotify** | Pointer to **string** | The IP address and port of the DNS server managing the smart architecture the DNS zone belongs to. If the parameter &lt;b&gt;zone_notify&lt;/b&gt; is set to &lt;b&gt;yes&lt;/b&gt; or &lt;b&gt;explicit&lt;/b&gt;, the server specified is instantly notified of any slave zones updates. | [optional] 
**ZoneClassName** | Pointer to **string** | The name of the class applied to the DNS zone, it can be preceded by the class directory. | [optional] 
**ZoneClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DNS zone. | [optional] 
**ZoneForward** | Pointer to **string** | The forwarding mode of the DNS zone.&lt;table&gt;&lt;caption&gt;zone_forward possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;first&lt;/td&gt;&lt;td &gt;The zone sends the queries to the forwarder(s). If no answer is returned, it attempts to answer the queries on its own.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;only&lt;/td&gt;&lt;td &gt;The zone only forwards the queries to the forwarder(s). Required by some reverse forward zones (e.g., in the case of private addresses).&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; If the parameter has no value, it indicates that the forwarding is disabled. | [optional] 
**ZoneForwarders** | Pointer to **string** | The IP address(es) of the forwarder(s) associated with the DNS zone. It lists the DNS servers to which any unknown query on this zone should be sent, as follows: &lt;b&gt;&amp;lt;ip_address1&amp;gt;;&amp;lt;ip_address2&amp;gt;;...&lt;/b&gt; . | [optional] 
**ZoneId** | Pointer to **string** | The database identifier (ID) of the DNS zone. | [optional] 
**ZoneIsReverse** | Pointer to **string** | A way to determine if the DNS zone provides reverse resolution (1) or direct/name resolution (0), | [optional] 
**ZoneIsRpz** | Pointer to **string** | The RPZ status of the DNS zone. &lt;b&gt;1&lt;/b&gt; indicates that the DNS zone is a Response Policy Zone. | [optional] 
**ZoneMasters** | Pointer to **string** | For slave DNS zones, the IP address of the DNS server and, if relevant, the name of the DNS view that contain the master DNS zone, as follows: &lt;b&gt;&amp;lt;ip_addr&amp;gt;;&lt;/b&gt; or &lt;b&gt;&amp;lt;ip_addr&amp;gt; key &amp;lt;dnsview_name&amp;gt;;&lt;/b&gt; . | [optional] 
**ZoneName** | Pointer to **string** | The name of the DNS zone. | [optional] 
**ZoneNameUtf** | Pointer to **string** | The name of the DNS zone in UTF-8 format. | [optional] 
**ZoneNotify** | Pointer to **string** | The notify status of the DNS zone.&lt;table&gt;&lt;caption&gt;zone_notify possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;No notify message is sent.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;A notify message is sent to the name servers defined in the NS records of the zone and to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;explicit&lt;/td&gt;&lt;td &gt;A notify message is sent only to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt;The notify message is not sent to the server itself or to the primary server defined in the SOA record of the zone. | [optional] 
**ZoneOrder** | Pointer to **string** | The level of the DNS zone, where 0 represents the highest level in the zones hierarchy. The RPZ rules parameters of each zone are reviewed following this order. The zones with the parameter zone_is_rpz set to 0 will always return 0 for the parameter zone_order. | [optional] 
**ZoneResponsePolicy** | Pointer to **string** | The Overriding rule of the RPZ zone. DNS zones are set with the policy &lt;b&gt;given&lt;/b&gt;. | [optional] 
**ZoneRevSortZone** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ZoneRpzLog** | Pointer to **string** | The logging status of an RPZ zone. | [optional] 
**ZoneRpzMaxPolicyTtl** | Pointer to **string** | The configuration of the option &lt;b&gt;Server max policy TTL&lt;/b&gt;, i.e. the number of seconds of the max policy Time To Live of the zone. It overrides the value set at server level. By default it is empty, i.e. set to &lt;b&gt;5&lt;/b&gt; seconds. | [optional] 
**ZoneRpzRecursiveOnly** | Pointer to **string** | The configuration of the &lt;b&gt;option Enable recursive-only on the server&lt;/b&gt; of the zone. It overrides the value set at server level. By default it is set to &lt;b&gt;1 (yes)&lt;/b&gt;, the server only processes policies on recursive queries. | [optional] 
**ZoneSpaceId** | Pointer to **string** | The database identifier (ID) of the space associated with the DNS zone. | [optional] 
**ZoneSpaceName** | Pointer to **string** | The name of the space associated with the DNS zone. | [optional] 
**ZoneSortZone** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ZoneSynching** | Pointer to **string** | The synchronization status of the DNS zone. &lt;b&gt;1&lt;/b&gt; indicates that the zone is currently being synchronized. | [optional] 
**ZoneType** | Pointer to **string** | The type of the DNS zone, either &lt;b&gt;master&lt;/b&gt;, &lt;b&gt;slave&lt;/b&gt;, &lt;b&gt;forward&lt;/b&gt;, &lt;b&gt;stub&lt;/b&gt;, &lt;b&gt;hint&lt;/b&gt; or &lt;b&gt;delegation-only&lt;/b&gt;. | [optional] 
**ZoneXferDone** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ZoneDs** | Pointer to **string** | The DNSSEC delegation signer (DS) fingerprint key associated with the DNS zone, if it is signed. | [optional] 
**ServerGssEnabled** | Pointer to **string** | The GSS-TSIG status of the DNS server the zone belongs to. &lt;b&gt;1&lt;/b&gt; indicates that GSS-TSIG is enabled on the server. | [optional] 
**ServerGssKeytabId** | Pointer to **string** | The database identifier (ID) of the DNS GSS-TSIG keytab. | [optional] 
**ServerHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server_addr&lt;/b&gt; or &lt;b&gt;server_addr6&lt;/b&gt;. | [optional] 
**ServerAddr6** | Pointer to **string** | The IPv6 address of the DNS server the object belongs to, in hexadecimal format. | [optional] 
**ServerAddr** | Pointer to **string** | The IPv4 address of the DNS server the object belongs to, in hexadecimal format. | [optional] 
**ServerIpmProtocol** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIpmType** | Pointer to **string** | The engine type of the DNS server the DNS zone belongs to: &lt;b&gt;named&lt;/b&gt; (BIND engine), &lt;b&gt;nsd&lt;/b&gt; (NSD engine) or &lt;b&gt;unbound&lt;/b&gt; (Unbound engine). | [optional] 
**ZoneMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ZoneNumKeys** | Pointer to **string** | The number of keys associated with the zone. This number of keys includes all ZSK and KSK. | [optional] 
**ZoneUseUpdatePolicy** | Pointer to **string** | The update policy status of the DNS zone. &lt;b&gt;1&lt;/b&gt; indicates that the DNS zone uses a specific GSS-TSIG/update-policy. The parameter &lt;b&gt;server_gss_enabled&lt;/b&gt; must be set to &lt;b&gt;1&lt;/b&gt;. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DNS smart architecture managing the DNS server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DNS smart architecture managing the DNS server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDnsZoneData

`func NewDataInnerDnsZoneData() *DataInnerDnsZoneData`

NewDataInnerDnsZoneData instantiates a new DataInnerDnsZoneData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDnsZoneDataWithDefaults

`func NewDataInnerDnsZoneDataWithDefaults() *DataInnerDnsZoneData`

NewDataInnerDnsZoneDataWithDefaults instantiates a new DataInnerDnsZoneData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerAwsDelegationSet

`func (o *DataInnerDnsZoneData) GetServerAwsDelegationSet() string`

GetServerAwsDelegationSet returns the ServerAwsDelegationSet field if non-nil, zero value otherwise.

### GetServerAwsDelegationSetOk

`func (o *DataInnerDnsZoneData) GetServerAwsDelegationSetOk() (*string, bool)`

GetServerAwsDelegationSetOk returns a tuple with the ServerAwsDelegationSet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAwsDelegationSet

`func (o *DataInnerDnsZoneData) SetServerAwsDelegationSet(v string)`

SetServerAwsDelegationSet sets ServerAwsDelegationSet field to given value.

### HasServerAwsDelegationSet

`func (o *DataInnerDnsZoneData) HasServerAwsDelegationSet() bool`

HasServerAwsDelegationSet returns a boolean if a field has been set.

### GetServerDdnsScavenging

`func (o *DataInnerDnsZoneData) GetServerDdnsScavenging() string`

GetServerDdnsScavenging returns the ServerDdnsScavenging field if non-nil, zero value otherwise.

### GetServerDdnsScavengingOk

`func (o *DataInnerDnsZoneData) GetServerDdnsScavengingOk() (*string, bool)`

GetServerDdnsScavengingOk returns a tuple with the ServerDdnsScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDdnsScavenging

`func (o *DataInnerDnsZoneData) SetServerDdnsScavenging(v string)`

SetServerDdnsScavenging sets ServerDdnsScavenging field to given value.

### HasServerDdnsScavenging

`func (o *DataInnerDnsZoneData) HasServerDdnsScavenging() bool`

HasServerDdnsScavenging returns a boolean if a field has been set.

### GetZoneDelayedCreateTime

`func (o *DataInnerDnsZoneData) GetZoneDelayedCreateTime() string`

GetZoneDelayedCreateTime returns the ZoneDelayedCreateTime field if non-nil, zero value otherwise.

### GetZoneDelayedCreateTimeOk

`func (o *DataInnerDnsZoneData) GetZoneDelayedCreateTimeOk() (*string, bool)`

GetZoneDelayedCreateTimeOk returns a tuple with the ZoneDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDelayedCreateTime

`func (o *DataInnerDnsZoneData) SetZoneDelayedCreateTime(v string)`

SetZoneDelayedCreateTime sets ZoneDelayedCreateTime field to given value.

### HasZoneDelayedCreateTime

`func (o *DataInnerDnsZoneData) HasZoneDelayedCreateTime() bool`

HasZoneDelayedCreateTime returns a boolean if a field has been set.

### GetZoneDelayedDeleteTime

`func (o *DataInnerDnsZoneData) GetZoneDelayedDeleteTime() string`

GetZoneDelayedDeleteTime returns the ZoneDelayedDeleteTime field if non-nil, zero value otherwise.

### GetZoneDelayedDeleteTimeOk

`func (o *DataInnerDnsZoneData) GetZoneDelayedDeleteTimeOk() (*string, bool)`

GetZoneDelayedDeleteTimeOk returns a tuple with the ZoneDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDelayedDeleteTime

`func (o *DataInnerDnsZoneData) SetZoneDelayedDeleteTime(v string)`

SetZoneDelayedDeleteTime sets ZoneDelayedDeleteTime field to given value.

### HasZoneDelayedDeleteTime

`func (o *DataInnerDnsZoneData) HasZoneDelayedDeleteTime() bool`

HasZoneDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DataInnerDnsZoneData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DataInnerDnsZoneData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DataInnerDnsZoneData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DataInnerDnsZoneData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DataInnerDnsZoneData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DataInnerDnsZoneData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DataInnerDnsZoneData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DataInnerDnsZoneData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerCloud

`func (o *DataInnerDnsZoneData) GetServerCloud() string`

GetServerCloud returns the ServerCloud field if non-nil, zero value otherwise.

### GetServerCloudOk

`func (o *DataInnerDnsZoneData) GetServerCloudOk() (*string, bool)`

GetServerCloudOk returns a tuple with the ServerCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCloud

`func (o *DataInnerDnsZoneData) SetServerCloud(v string)`

SetServerCloud sets ServerCloud field to given value.

### HasServerCloud

`func (o *DataInnerDnsZoneData) HasServerCloud() bool`

HasServerCloud returns a boolean if a field has been set.

### GetServerComment

`func (o *DataInnerDnsZoneData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DataInnerDnsZoneData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DataInnerDnsZoneData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DataInnerDnsZoneData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerForceHybrid

`func (o *DataInnerDnsZoneData) GetServerForceHybrid() string`

GetServerForceHybrid returns the ServerForceHybrid field if non-nil, zero value otherwise.

### GetServerForceHybridOk

`func (o *DataInnerDnsZoneData) GetServerForceHybridOk() (*string, bool)`

GetServerForceHybridOk returns a tuple with the ServerForceHybrid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerForceHybrid

`func (o *DataInnerDnsZoneData) SetServerForceHybrid(v string)`

SetServerForceHybrid sets ServerForceHybrid field to given value.

### HasServerForceHybrid

`func (o *DataInnerDnsZoneData) HasServerForceHybrid() bool`

HasServerForceHybrid returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDnsZoneData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDnsZoneData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDnsZoneData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDnsZoneData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDnsZoneData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDnsZoneData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDnsZoneData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDnsZoneData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerState

`func (o *DataInnerDnsZoneData) GetServerState() string`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *DataInnerDnsZoneData) GetServerStateOk() (*string, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *DataInnerDnsZoneData) SetServerState(v string)`

SetServerState sets ServerState field to given value.

### HasServerState

`func (o *DataInnerDnsZoneData) HasServerState() bool`

HasServerState returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDnsZoneData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDnsZoneData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDnsZoneData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDnsZoneData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DataInnerDnsZoneData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DataInnerDnsZoneData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DataInnerDnsZoneData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DataInnerDnsZoneData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerVpcList

`func (o *DataInnerDnsZoneData) GetServerVpcList() string`

GetServerVpcList returns the ServerVpcList field if non-nil, zero value otherwise.

### GetServerVpcListOk

`func (o *DataInnerDnsZoneData) GetServerVpcListOk() (*string, bool)`

GetServerVpcListOk returns a tuple with the ServerVpcList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVpcList

`func (o *DataInnerDnsZoneData) SetServerVpcList(v string)`

SetServerVpcList sets ServerVpcList field to given value.

### HasServerVpcList

`func (o *DataInnerDnsZoneData) HasServerVpcList() bool`

HasServerVpcList returns a boolean if a field has been set.

### GetViewClassName

`func (o *DataInnerDnsZoneData) GetViewClassName() string`

GetViewClassName returns the ViewClassName field if non-nil, zero value otherwise.

### GetViewClassNameOk

`func (o *DataInnerDnsZoneData) GetViewClassNameOk() (*string, bool)`

GetViewClassNameOk returns a tuple with the ViewClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassName

`func (o *DataInnerDnsZoneData) SetViewClassName(v string)`

SetViewClassName sets ViewClassName field to given value.

### HasViewClassName

`func (o *DataInnerDnsZoneData) HasViewClassName() bool`

HasViewClassName returns a boolean if a field has been set.

### GetViewClassParameters

`func (o *DataInnerDnsZoneData) GetViewClassParameters() []ApiClassParameterOutputEntry`

GetViewClassParameters returns the ViewClassParameters field if non-nil, zero value otherwise.

### GetViewClassParametersOk

`func (o *DataInnerDnsZoneData) GetViewClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetViewClassParametersOk returns a tuple with the ViewClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassParameters

`func (o *DataInnerDnsZoneData) SetViewClassParameters(v []ApiClassParameterOutputEntry)`

SetViewClassParameters sets ViewClassParameters field to given value.

### HasViewClassParameters

`func (o *DataInnerDnsZoneData) HasViewClassParameters() bool`

HasViewClassParameters returns a boolean if a field has been set.

### GetViewId

`func (o *DataInnerDnsZoneData) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DataInnerDnsZoneData) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DataInnerDnsZoneData) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DataInnerDnsZoneData) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewName

`func (o *DataInnerDnsZoneData) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DataInnerDnsZoneData) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DataInnerDnsZoneData) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DataInnerDnsZoneData) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetZoneAdIntegrated

`func (o *DataInnerDnsZoneData) GetZoneAdIntegrated() string`

GetZoneAdIntegrated returns the ZoneAdIntegrated field if non-nil, zero value otherwise.

### GetZoneAdIntegratedOk

`func (o *DataInnerDnsZoneData) GetZoneAdIntegratedOk() (*string, bool)`

GetZoneAdIntegratedOk returns a tuple with the ZoneAdIntegrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAdIntegrated

`func (o *DataInnerDnsZoneData) SetZoneAdIntegrated(v string)`

SetZoneAdIntegrated sets ZoneAdIntegrated field to given value.

### HasZoneAdIntegrated

`func (o *DataInnerDnsZoneData) HasZoneAdIntegrated() bool`

HasZoneAdIntegrated returns a boolean if a field has been set.

### GetZoneAllowQuery

`func (o *DataInnerDnsZoneData) GetZoneAllowQuery() string`

GetZoneAllowQuery returns the ZoneAllowQuery field if non-nil, zero value otherwise.

### GetZoneAllowQueryOk

`func (o *DataInnerDnsZoneData) GetZoneAllowQueryOk() (*string, bool)`

GetZoneAllowQueryOk returns a tuple with the ZoneAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowQuery

`func (o *DataInnerDnsZoneData) SetZoneAllowQuery(v string)`

SetZoneAllowQuery sets ZoneAllowQuery field to given value.

### HasZoneAllowQuery

`func (o *DataInnerDnsZoneData) HasZoneAllowQuery() bool`

HasZoneAllowQuery returns a boolean if a field has been set.

### GetZoneAllowTransfer

`func (o *DataInnerDnsZoneData) GetZoneAllowTransfer() string`

GetZoneAllowTransfer returns the ZoneAllowTransfer field if non-nil, zero value otherwise.

### GetZoneAllowTransferOk

`func (o *DataInnerDnsZoneData) GetZoneAllowTransferOk() (*string, bool)`

GetZoneAllowTransferOk returns a tuple with the ZoneAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowTransfer

`func (o *DataInnerDnsZoneData) SetZoneAllowTransfer(v string)`

SetZoneAllowTransfer sets ZoneAllowTransfer field to given value.

### HasZoneAllowTransfer

`func (o *DataInnerDnsZoneData) HasZoneAllowTransfer() bool`

HasZoneAllowTransfer returns a boolean if a field has been set.

### GetZoneAllowUpdate

`func (o *DataInnerDnsZoneData) GetZoneAllowUpdate() string`

GetZoneAllowUpdate returns the ZoneAllowUpdate field if non-nil, zero value otherwise.

### GetZoneAllowUpdateOk

`func (o *DataInnerDnsZoneData) GetZoneAllowUpdateOk() (*string, bool)`

GetZoneAllowUpdateOk returns a tuple with the ZoneAllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowUpdate

`func (o *DataInnerDnsZoneData) SetZoneAllowUpdate(v string)`

SetZoneAllowUpdate sets ZoneAllowUpdate field to given value.

### HasZoneAllowUpdate

`func (o *DataInnerDnsZoneData) HasZoneAllowUpdate() bool`

HasZoneAllowUpdate returns a boolean if a field has been set.

### GetZoneAlsoNotify

`func (o *DataInnerDnsZoneData) GetZoneAlsoNotify() string`

GetZoneAlsoNotify returns the ZoneAlsoNotify field if non-nil, zero value otherwise.

### GetZoneAlsoNotifyOk

`func (o *DataInnerDnsZoneData) GetZoneAlsoNotifyOk() (*string, bool)`

GetZoneAlsoNotifyOk returns a tuple with the ZoneAlsoNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAlsoNotify

`func (o *DataInnerDnsZoneData) SetZoneAlsoNotify(v string)`

SetZoneAlsoNotify sets ZoneAlsoNotify field to given value.

### HasZoneAlsoNotify

`func (o *DataInnerDnsZoneData) HasZoneAlsoNotify() bool`

HasZoneAlsoNotify returns a boolean if a field has been set.

### GetZoneClassName

`func (o *DataInnerDnsZoneData) GetZoneClassName() string`

GetZoneClassName returns the ZoneClassName field if non-nil, zero value otherwise.

### GetZoneClassNameOk

`func (o *DataInnerDnsZoneData) GetZoneClassNameOk() (*string, bool)`

GetZoneClassNameOk returns a tuple with the ZoneClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneClassName

`func (o *DataInnerDnsZoneData) SetZoneClassName(v string)`

SetZoneClassName sets ZoneClassName field to given value.

### HasZoneClassName

`func (o *DataInnerDnsZoneData) HasZoneClassName() bool`

HasZoneClassName returns a boolean if a field has been set.

### GetZoneClassParameters

`func (o *DataInnerDnsZoneData) GetZoneClassParameters() []ApiClassParameterOutputEntry`

GetZoneClassParameters returns the ZoneClassParameters field if non-nil, zero value otherwise.

### GetZoneClassParametersOk

`func (o *DataInnerDnsZoneData) GetZoneClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetZoneClassParametersOk returns a tuple with the ZoneClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneClassParameters

`func (o *DataInnerDnsZoneData) SetZoneClassParameters(v []ApiClassParameterOutputEntry)`

SetZoneClassParameters sets ZoneClassParameters field to given value.

### HasZoneClassParameters

`func (o *DataInnerDnsZoneData) HasZoneClassParameters() bool`

HasZoneClassParameters returns a boolean if a field has been set.

### GetZoneForward

`func (o *DataInnerDnsZoneData) GetZoneForward() string`

GetZoneForward returns the ZoneForward field if non-nil, zero value otherwise.

### GetZoneForwardOk

`func (o *DataInnerDnsZoneData) GetZoneForwardOk() (*string, bool)`

GetZoneForwardOk returns a tuple with the ZoneForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneForward

`func (o *DataInnerDnsZoneData) SetZoneForward(v string)`

SetZoneForward sets ZoneForward field to given value.

### HasZoneForward

`func (o *DataInnerDnsZoneData) HasZoneForward() bool`

HasZoneForward returns a boolean if a field has been set.

### GetZoneForwarders

`func (o *DataInnerDnsZoneData) GetZoneForwarders() string`

GetZoneForwarders returns the ZoneForwarders field if non-nil, zero value otherwise.

### GetZoneForwardersOk

`func (o *DataInnerDnsZoneData) GetZoneForwardersOk() (*string, bool)`

GetZoneForwardersOk returns a tuple with the ZoneForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneForwarders

`func (o *DataInnerDnsZoneData) SetZoneForwarders(v string)`

SetZoneForwarders sets ZoneForwarders field to given value.

### HasZoneForwarders

`func (o *DataInnerDnsZoneData) HasZoneForwarders() bool`

HasZoneForwarders returns a boolean if a field has been set.

### GetZoneId

`func (o *DataInnerDnsZoneData) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DataInnerDnsZoneData) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DataInnerDnsZoneData) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DataInnerDnsZoneData) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneIsReverse

`func (o *DataInnerDnsZoneData) GetZoneIsReverse() string`

GetZoneIsReverse returns the ZoneIsReverse field if non-nil, zero value otherwise.

### GetZoneIsReverseOk

`func (o *DataInnerDnsZoneData) GetZoneIsReverseOk() (*string, bool)`

GetZoneIsReverseOk returns a tuple with the ZoneIsReverse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIsReverse

`func (o *DataInnerDnsZoneData) SetZoneIsReverse(v string)`

SetZoneIsReverse sets ZoneIsReverse field to given value.

### HasZoneIsReverse

`func (o *DataInnerDnsZoneData) HasZoneIsReverse() bool`

HasZoneIsReverse returns a boolean if a field has been set.

### GetZoneIsRpz

`func (o *DataInnerDnsZoneData) GetZoneIsRpz() string`

GetZoneIsRpz returns the ZoneIsRpz field if non-nil, zero value otherwise.

### GetZoneIsRpzOk

`func (o *DataInnerDnsZoneData) GetZoneIsRpzOk() (*string, bool)`

GetZoneIsRpzOk returns a tuple with the ZoneIsRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIsRpz

`func (o *DataInnerDnsZoneData) SetZoneIsRpz(v string)`

SetZoneIsRpz sets ZoneIsRpz field to given value.

### HasZoneIsRpz

`func (o *DataInnerDnsZoneData) HasZoneIsRpz() bool`

HasZoneIsRpz returns a boolean if a field has been set.

### GetZoneMasters

`func (o *DataInnerDnsZoneData) GetZoneMasters() string`

GetZoneMasters returns the ZoneMasters field if non-nil, zero value otherwise.

### GetZoneMastersOk

`func (o *DataInnerDnsZoneData) GetZoneMastersOk() (*string, bool)`

GetZoneMastersOk returns a tuple with the ZoneMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMasters

`func (o *DataInnerDnsZoneData) SetZoneMasters(v string)`

SetZoneMasters sets ZoneMasters field to given value.

### HasZoneMasters

`func (o *DataInnerDnsZoneData) HasZoneMasters() bool`

HasZoneMasters returns a boolean if a field has been set.

### GetZoneName

`func (o *DataInnerDnsZoneData) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DataInnerDnsZoneData) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DataInnerDnsZoneData) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *DataInnerDnsZoneData) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneNameUtf

`func (o *DataInnerDnsZoneData) GetZoneNameUtf() string`

GetZoneNameUtf returns the ZoneNameUtf field if non-nil, zero value otherwise.

### GetZoneNameUtfOk

`func (o *DataInnerDnsZoneData) GetZoneNameUtfOk() (*string, bool)`

GetZoneNameUtfOk returns a tuple with the ZoneNameUtf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNameUtf

`func (o *DataInnerDnsZoneData) SetZoneNameUtf(v string)`

SetZoneNameUtf sets ZoneNameUtf field to given value.

### HasZoneNameUtf

`func (o *DataInnerDnsZoneData) HasZoneNameUtf() bool`

HasZoneNameUtf returns a boolean if a field has been set.

### GetZoneNotify

`func (o *DataInnerDnsZoneData) GetZoneNotify() string`

GetZoneNotify returns the ZoneNotify field if non-nil, zero value otherwise.

### GetZoneNotifyOk

`func (o *DataInnerDnsZoneData) GetZoneNotifyOk() (*string, bool)`

GetZoneNotifyOk returns a tuple with the ZoneNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNotify

`func (o *DataInnerDnsZoneData) SetZoneNotify(v string)`

SetZoneNotify sets ZoneNotify field to given value.

### HasZoneNotify

`func (o *DataInnerDnsZoneData) HasZoneNotify() bool`

HasZoneNotify returns a boolean if a field has been set.

### GetZoneOrder

`func (o *DataInnerDnsZoneData) GetZoneOrder() string`

GetZoneOrder returns the ZoneOrder field if non-nil, zero value otherwise.

### GetZoneOrderOk

`func (o *DataInnerDnsZoneData) GetZoneOrderOk() (*string, bool)`

GetZoneOrderOk returns a tuple with the ZoneOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneOrder

`func (o *DataInnerDnsZoneData) SetZoneOrder(v string)`

SetZoneOrder sets ZoneOrder field to given value.

### HasZoneOrder

`func (o *DataInnerDnsZoneData) HasZoneOrder() bool`

HasZoneOrder returns a boolean if a field has been set.

### GetZoneResponsePolicy

`func (o *DataInnerDnsZoneData) GetZoneResponsePolicy() string`

GetZoneResponsePolicy returns the ZoneResponsePolicy field if non-nil, zero value otherwise.

### GetZoneResponsePolicyOk

`func (o *DataInnerDnsZoneData) GetZoneResponsePolicyOk() (*string, bool)`

GetZoneResponsePolicyOk returns a tuple with the ZoneResponsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneResponsePolicy

`func (o *DataInnerDnsZoneData) SetZoneResponsePolicy(v string)`

SetZoneResponsePolicy sets ZoneResponsePolicy field to given value.

### HasZoneResponsePolicy

`func (o *DataInnerDnsZoneData) HasZoneResponsePolicy() bool`

HasZoneResponsePolicy returns a boolean if a field has been set.

### GetZoneRevSortZone

`func (o *DataInnerDnsZoneData) GetZoneRevSortZone() string`

GetZoneRevSortZone returns the ZoneRevSortZone field if non-nil, zero value otherwise.

### GetZoneRevSortZoneOk

`func (o *DataInnerDnsZoneData) GetZoneRevSortZoneOk() (*string, bool)`

GetZoneRevSortZoneOk returns a tuple with the ZoneRevSortZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRevSortZone

`func (o *DataInnerDnsZoneData) SetZoneRevSortZone(v string)`

SetZoneRevSortZone sets ZoneRevSortZone field to given value.

### HasZoneRevSortZone

`func (o *DataInnerDnsZoneData) HasZoneRevSortZone() bool`

HasZoneRevSortZone returns a boolean if a field has been set.

### GetZoneRpzLog

`func (o *DataInnerDnsZoneData) GetZoneRpzLog() string`

GetZoneRpzLog returns the ZoneRpzLog field if non-nil, zero value otherwise.

### GetZoneRpzLogOk

`func (o *DataInnerDnsZoneData) GetZoneRpzLogOk() (*string, bool)`

GetZoneRpzLogOk returns a tuple with the ZoneRpzLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRpzLog

`func (o *DataInnerDnsZoneData) SetZoneRpzLog(v string)`

SetZoneRpzLog sets ZoneRpzLog field to given value.

### HasZoneRpzLog

`func (o *DataInnerDnsZoneData) HasZoneRpzLog() bool`

HasZoneRpzLog returns a boolean if a field has been set.

### GetZoneRpzMaxPolicyTtl

`func (o *DataInnerDnsZoneData) GetZoneRpzMaxPolicyTtl() string`

GetZoneRpzMaxPolicyTtl returns the ZoneRpzMaxPolicyTtl field if non-nil, zero value otherwise.

### GetZoneRpzMaxPolicyTtlOk

`func (o *DataInnerDnsZoneData) GetZoneRpzMaxPolicyTtlOk() (*string, bool)`

GetZoneRpzMaxPolicyTtlOk returns a tuple with the ZoneRpzMaxPolicyTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRpzMaxPolicyTtl

`func (o *DataInnerDnsZoneData) SetZoneRpzMaxPolicyTtl(v string)`

SetZoneRpzMaxPolicyTtl sets ZoneRpzMaxPolicyTtl field to given value.

### HasZoneRpzMaxPolicyTtl

`func (o *DataInnerDnsZoneData) HasZoneRpzMaxPolicyTtl() bool`

HasZoneRpzMaxPolicyTtl returns a boolean if a field has been set.

### GetZoneRpzRecursiveOnly

`func (o *DataInnerDnsZoneData) GetZoneRpzRecursiveOnly() string`

GetZoneRpzRecursiveOnly returns the ZoneRpzRecursiveOnly field if non-nil, zero value otherwise.

### GetZoneRpzRecursiveOnlyOk

`func (o *DataInnerDnsZoneData) GetZoneRpzRecursiveOnlyOk() (*string, bool)`

GetZoneRpzRecursiveOnlyOk returns a tuple with the ZoneRpzRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRpzRecursiveOnly

`func (o *DataInnerDnsZoneData) SetZoneRpzRecursiveOnly(v string)`

SetZoneRpzRecursiveOnly sets ZoneRpzRecursiveOnly field to given value.

### HasZoneRpzRecursiveOnly

`func (o *DataInnerDnsZoneData) HasZoneRpzRecursiveOnly() bool`

HasZoneRpzRecursiveOnly returns a boolean if a field has been set.

### GetZoneSpaceId

`func (o *DataInnerDnsZoneData) GetZoneSpaceId() string`

GetZoneSpaceId returns the ZoneSpaceId field if non-nil, zero value otherwise.

### GetZoneSpaceIdOk

`func (o *DataInnerDnsZoneData) GetZoneSpaceIdOk() (*string, bool)`

GetZoneSpaceIdOk returns a tuple with the ZoneSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSpaceId

`func (o *DataInnerDnsZoneData) SetZoneSpaceId(v string)`

SetZoneSpaceId sets ZoneSpaceId field to given value.

### HasZoneSpaceId

`func (o *DataInnerDnsZoneData) HasZoneSpaceId() bool`

HasZoneSpaceId returns a boolean if a field has been set.

### GetZoneSpaceName

`func (o *DataInnerDnsZoneData) GetZoneSpaceName() string`

GetZoneSpaceName returns the ZoneSpaceName field if non-nil, zero value otherwise.

### GetZoneSpaceNameOk

`func (o *DataInnerDnsZoneData) GetZoneSpaceNameOk() (*string, bool)`

GetZoneSpaceNameOk returns a tuple with the ZoneSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSpaceName

`func (o *DataInnerDnsZoneData) SetZoneSpaceName(v string)`

SetZoneSpaceName sets ZoneSpaceName field to given value.

### HasZoneSpaceName

`func (o *DataInnerDnsZoneData) HasZoneSpaceName() bool`

HasZoneSpaceName returns a boolean if a field has been set.

### GetZoneSortZone

`func (o *DataInnerDnsZoneData) GetZoneSortZone() string`

GetZoneSortZone returns the ZoneSortZone field if non-nil, zero value otherwise.

### GetZoneSortZoneOk

`func (o *DataInnerDnsZoneData) GetZoneSortZoneOk() (*string, bool)`

GetZoneSortZoneOk returns a tuple with the ZoneSortZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSortZone

`func (o *DataInnerDnsZoneData) SetZoneSortZone(v string)`

SetZoneSortZone sets ZoneSortZone field to given value.

### HasZoneSortZone

`func (o *DataInnerDnsZoneData) HasZoneSortZone() bool`

HasZoneSortZone returns a boolean if a field has been set.

### GetZoneSynching

`func (o *DataInnerDnsZoneData) GetZoneSynching() string`

GetZoneSynching returns the ZoneSynching field if non-nil, zero value otherwise.

### GetZoneSynchingOk

`func (o *DataInnerDnsZoneData) GetZoneSynchingOk() (*string, bool)`

GetZoneSynchingOk returns a tuple with the ZoneSynching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSynching

`func (o *DataInnerDnsZoneData) SetZoneSynching(v string)`

SetZoneSynching sets ZoneSynching field to given value.

### HasZoneSynching

`func (o *DataInnerDnsZoneData) HasZoneSynching() bool`

HasZoneSynching returns a boolean if a field has been set.

### GetZoneType

`func (o *DataInnerDnsZoneData) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *DataInnerDnsZoneData) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *DataInnerDnsZoneData) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *DataInnerDnsZoneData) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetZoneXferDone

`func (o *DataInnerDnsZoneData) GetZoneXferDone() string`

GetZoneXferDone returns the ZoneXferDone field if non-nil, zero value otherwise.

### GetZoneXferDoneOk

`func (o *DataInnerDnsZoneData) GetZoneXferDoneOk() (*string, bool)`

GetZoneXferDoneOk returns a tuple with the ZoneXferDone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneXferDone

`func (o *DataInnerDnsZoneData) SetZoneXferDone(v string)`

SetZoneXferDone sets ZoneXferDone field to given value.

### HasZoneXferDone

`func (o *DataInnerDnsZoneData) HasZoneXferDone() bool`

HasZoneXferDone returns a boolean if a field has been set.

### GetZoneDs

`func (o *DataInnerDnsZoneData) GetZoneDs() string`

GetZoneDs returns the ZoneDs field if non-nil, zero value otherwise.

### GetZoneDsOk

`func (o *DataInnerDnsZoneData) GetZoneDsOk() (*string, bool)`

GetZoneDsOk returns a tuple with the ZoneDs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneDs

`func (o *DataInnerDnsZoneData) SetZoneDs(v string)`

SetZoneDs sets ZoneDs field to given value.

### HasZoneDs

`func (o *DataInnerDnsZoneData) HasZoneDs() bool`

HasZoneDs returns a boolean if a field has been set.

### GetServerGssEnabled

`func (o *DataInnerDnsZoneData) GetServerGssEnabled() string`

GetServerGssEnabled returns the ServerGssEnabled field if non-nil, zero value otherwise.

### GetServerGssEnabledOk

`func (o *DataInnerDnsZoneData) GetServerGssEnabledOk() (*string, bool)`

GetServerGssEnabledOk returns a tuple with the ServerGssEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGssEnabled

`func (o *DataInnerDnsZoneData) SetServerGssEnabled(v string)`

SetServerGssEnabled sets ServerGssEnabled field to given value.

### HasServerGssEnabled

`func (o *DataInnerDnsZoneData) HasServerGssEnabled() bool`

HasServerGssEnabled returns a boolean if a field has been set.

### GetServerGssKeytabId

`func (o *DataInnerDnsZoneData) GetServerGssKeytabId() string`

GetServerGssKeytabId returns the ServerGssKeytabId field if non-nil, zero value otherwise.

### GetServerGssKeytabIdOk

`func (o *DataInnerDnsZoneData) GetServerGssKeytabIdOk() (*string, bool)`

GetServerGssKeytabIdOk returns a tuple with the ServerGssKeytabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGssKeytabId

`func (o *DataInnerDnsZoneData) SetServerGssKeytabId(v string)`

SetServerGssKeytabId sets ServerGssKeytabId field to given value.

### HasServerGssKeytabId

`func (o *DataInnerDnsZoneData) HasServerGssKeytabId() bool`

HasServerGssKeytabId returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DataInnerDnsZoneData) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DataInnerDnsZoneData) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DataInnerDnsZoneData) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DataInnerDnsZoneData) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerAddr6

`func (o *DataInnerDnsZoneData) GetServerAddr6() string`

GetServerAddr6 returns the ServerAddr6 field if non-nil, zero value otherwise.

### GetServerAddr6Ok

`func (o *DataInnerDnsZoneData) GetServerAddr6Ok() (*string, bool)`

GetServerAddr6Ok returns a tuple with the ServerAddr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr6

`func (o *DataInnerDnsZoneData) SetServerAddr6(v string)`

SetServerAddr6 sets ServerAddr6 field to given value.

### HasServerAddr6

`func (o *DataInnerDnsZoneData) HasServerAddr6() bool`

HasServerAddr6 returns a boolean if a field has been set.

### GetServerAddr

`func (o *DataInnerDnsZoneData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DataInnerDnsZoneData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DataInnerDnsZoneData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DataInnerDnsZoneData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetServerIpmProtocol

`func (o *DataInnerDnsZoneData) GetServerIpmProtocol() string`

GetServerIpmProtocol returns the ServerIpmProtocol field if non-nil, zero value otherwise.

### GetServerIpmProtocolOk

`func (o *DataInnerDnsZoneData) GetServerIpmProtocolOk() (*string, bool)`

GetServerIpmProtocolOk returns a tuple with the ServerIpmProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmProtocol

`func (o *DataInnerDnsZoneData) SetServerIpmProtocol(v string)`

SetServerIpmProtocol sets ServerIpmProtocol field to given value.

### HasServerIpmProtocol

`func (o *DataInnerDnsZoneData) HasServerIpmProtocol() bool`

HasServerIpmProtocol returns a boolean if a field has been set.

### GetServerIpmType

`func (o *DataInnerDnsZoneData) GetServerIpmType() string`

GetServerIpmType returns the ServerIpmType field if non-nil, zero value otherwise.

### GetServerIpmTypeOk

`func (o *DataInnerDnsZoneData) GetServerIpmTypeOk() (*string, bool)`

GetServerIpmTypeOk returns a tuple with the ServerIpmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmType

`func (o *DataInnerDnsZoneData) SetServerIpmType(v string)`

SetServerIpmType sets ServerIpmType field to given value.

### HasServerIpmType

`func (o *DataInnerDnsZoneData) HasServerIpmType() bool`

HasServerIpmType returns a boolean if a field has been set.

### GetZoneMultistatus

`func (o *DataInnerDnsZoneData) GetZoneMultistatus() string`

GetZoneMultistatus returns the ZoneMultistatus field if non-nil, zero value otherwise.

### GetZoneMultistatusOk

`func (o *DataInnerDnsZoneData) GetZoneMultistatusOk() (*string, bool)`

GetZoneMultistatusOk returns a tuple with the ZoneMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMultistatus

`func (o *DataInnerDnsZoneData) SetZoneMultistatus(v string)`

SetZoneMultistatus sets ZoneMultistatus field to given value.

### HasZoneMultistatus

`func (o *DataInnerDnsZoneData) HasZoneMultistatus() bool`

HasZoneMultistatus returns a boolean if a field has been set.

### GetZoneNumKeys

`func (o *DataInnerDnsZoneData) GetZoneNumKeys() string`

GetZoneNumKeys returns the ZoneNumKeys field if non-nil, zero value otherwise.

### GetZoneNumKeysOk

`func (o *DataInnerDnsZoneData) GetZoneNumKeysOk() (*string, bool)`

GetZoneNumKeysOk returns a tuple with the ZoneNumKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNumKeys

`func (o *DataInnerDnsZoneData) SetZoneNumKeys(v string)`

SetZoneNumKeys sets ZoneNumKeys field to given value.

### HasZoneNumKeys

`func (o *DataInnerDnsZoneData) HasZoneNumKeys() bool`

HasZoneNumKeys returns a boolean if a field has been set.

### GetZoneUseUpdatePolicy

`func (o *DataInnerDnsZoneData) GetZoneUseUpdatePolicy() string`

GetZoneUseUpdatePolicy returns the ZoneUseUpdatePolicy field if non-nil, zero value otherwise.

### GetZoneUseUpdatePolicyOk

`func (o *DataInnerDnsZoneData) GetZoneUseUpdatePolicyOk() (*string, bool)`

GetZoneUseUpdatePolicyOk returns a tuple with the ZoneUseUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUseUpdatePolicy

`func (o *DataInnerDnsZoneData) SetZoneUseUpdatePolicy(v string)`

SetZoneUseUpdatePolicy sets ZoneUseUpdatePolicy field to given value.

### HasZoneUseUpdatePolicy

`func (o *DataInnerDnsZoneData) HasZoneUseUpdatePolicy() bool`

HasZoneUseUpdatePolicy returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDnsZoneData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDnsZoneData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDnsZoneData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDnsZoneData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDnsZoneData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDnsZoneData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDnsZoneData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDnsZoneData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


