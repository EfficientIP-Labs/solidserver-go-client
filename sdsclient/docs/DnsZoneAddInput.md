# DnsZoneAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server. | [optional] 
**ViewId** | Pointer to **int32** | The database identifier (ID) of the DNS view the object belongs to. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view the object belongs to. | [optional] 
**ZoneName** | Pointer to **string** | The name of the DNS zone, each DNS zone must have a unique name. For hint zones (&lt;b&gt;zone_type&lt;/b&gt;: &lt;b&gt;hint&lt;/b&gt;), you must type in &lt;b&gt;.&lt;/b&gt; (dot) as &lt;b&gt;zone_name&lt;/b&gt;. | [optional] 
**ZoneType** | Pointer to **string** | The type of the DNS zone, either &lt;b&gt;master&lt;/b&gt;, &lt;b&gt;slave&lt;/b&gt;, &lt;b&gt;forward&lt;/b&gt;, &lt;b&gt;stub&lt;/b&gt;, &lt;b&gt;hint&lt;/b&gt; or &lt;b&gt;delegation-only&lt;/b&gt;. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DNS server. | [optional] 
**ServerDdnsScavenging** | Pointer to **int32** | The DDNS scavenging status of the zone. Set it to &lt;b&gt;1&lt;/b&gt; to enable the scavenging and automatically delete the stale resource records dynamically added via GSS-TSIG. DDNS scavenging is only effective if the parameters &lt;b&gt;zone_use_update_policy&lt;/b&gt; and &lt;b&gt;gss_enabled&lt;/b&gt; are set to &lt;b&gt;1&lt;/b&gt;, and the rule 416 is enabled in the GUI. | [optional] 
**ZoneAdIntegrated** | Pointer to **int32** | The AD integrated status of the DNS zone. Set it to &lt;b&gt;1&lt;/b&gt; to indicate that the DNS zone belongs to an Active Directory integrated Microsoft Windows DNS server. | [optional] 
**ZoneAllowQuery** | Pointer to **string** | The ACL values associated with the allow-query configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAllowTransfer** | Pointer to **string** | The ACL values associated with the allow-transfer configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAllowUpdate** | Pointer to **string** | The ACL values associated with the allow-update configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAlsoNotify** | Pointer to **string** | The IP address and port of the DNS server managing the smart architecture the DNS zone belongs to. If the parameter &lt;b&gt;zone_notify&lt;/b&gt; is set to &lt;b&gt;yes&lt;/b&gt; or &lt;b&gt;explicit&lt;/b&gt;, the server specified is instantly notified of any slave zones updates. | [optional] 
**ZoneForward** | Pointer to **string** | The forwarding mode of the DNS zone.&lt;table&gt;&lt;caption&gt;zone_forward possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;first&lt;/td&gt;&lt;td &gt;The zone sends the queries to the forwarder(s). If no answer is returned, it attempts to answer the queries on its own.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;only&lt;/td&gt;&lt;td &gt;The zone only forwards the queries to the forwarder(s). Required by some reverse forward zones (e.g., in the case of private addresses).&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; If the parameter has no value, it indicates that the forwarding is disabled. | [optional] 
**ZoneForwarders** | Pointer to **string** | The IP address(es) of the forwarder(s) associated with the DNS zone. It lists the DNS servers to which any unknown query on this zone should be sent, as follows: &lt;b&gt;&amp;lt;ip_address1&amp;gt;;&amp;lt;ip_address2&amp;gt;;...&lt;/b&gt; . | [optional] 
**ZoneIsRpz** | Pointer to **int32** | The RPZ status of the DNS zone. Set it to &lt;b&gt;1&lt;/b&gt; to indicate that the DNS zone is a Response Policy Zone. | [optional] 
**ZoneMasters** | Pointer to **string** | For slave DNS zones, the IP address of the DNS server and, if relevant, the name of the DNS view that contain the master DNS zone, as follows: &lt;b&gt;&amp;lt;ip_addr&amp;gt;;&lt;/b&gt; or &lt;b&gt;&amp;lt;ip_addr&amp;gt; key &amp;lt;dnsview_name&amp;gt;;&lt;/b&gt; . | [optional] 
**ZoneNotify** | Pointer to **string** | The notify status of the DNS zone.&lt;table&gt;&lt;caption&gt;zone_notify possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;No notify message is sent.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;A notify message is sent to the name servers defined in the NS records of the zone and to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;explicit&lt;/td&gt;&lt;td &gt;A notify message is sent only to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt;The notify message is not sent to the server itself or to the primary server defined in the SOA record of the zone. | [optional] 
**ZoneOrder** | Pointer to **int32** | The level of the RPZ zone, where 0 represents the highest level in the views hierarchy. The RPZ rules of each zone are reviewed following this order. For non-RPZ zones, that have their parameter zone_is_rpz set to &lt;b&gt;0&lt;/b&gt;, you do not need to set this parameter. | [optional] 
**ZoneResponsePolicy** | Pointer to **string** | The response policy of the DNS zone. All the zones can be set with the policy &lt;b&gt;given&lt;/b&gt;, only RPZ zones can be set with other policies.&lt;table&gt;&lt;caption&gt;zone_response_policy possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Policy&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;given&lt;/td&gt;&lt;td &gt;All the rules specified in the RPZ zone are applied normally.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;disabled&lt;/td&gt;&lt;td &gt;The RPZ zone rules configuration is not applied. All the rules it contains are ignored.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;passthru&lt;/td&gt;&lt;td &gt;The rules specified in the RPZ matching the listed RR names are ignored, no matter the RPZ zone they belong to.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;nxdomain&lt;/td&gt;&lt;td &gt;The rules specified in the RPZ return an NXDOMAIN response.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;nodata&lt;/td&gt;&lt;td &gt;The rules specified in the RPZ return a NODATA response.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;cname &lt;domain-name&gt;&lt;/td&gt;&lt;td &gt;All the rules specified in the RPZ are redirected toward the specified domain name.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt;You can only add RPZ zones on EfficientIP or BIND DNS servers. | [optional] 
**ZoneRpzLog** | Pointer to **int32** | The RPZ logging status, if &lt;b&gt;zone_is_rpz&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; or &lt;b&gt;yes&lt;/b&gt;. It allows you to log all the operations triggered by the RPZ rules of the zone. | [optional] 
**ZoneRpzMaxPolicyTtl** | Pointer to **int32** | A way to set the number of seconds of the max policy Time To Live of the zone. By default, the parameter is empty, and the &lt;b&gt;Server max policy TTL&lt;/b&gt; of the zone is of 5 seconds. It overrides the same option set at server level. | [optional] 
**ZoneRpzRecursiveOnly** | Pointer to **int32** | A way to &lt;b&gt;Enable recursive-only on the server&lt;/b&gt; for the zone. By default it is enabled (&lt;b&gt;1&lt;/b&gt;) and, for this zone, the server only processes policies on recursive queries. It overrides the same option set at server level. | [optional] 
**ZoneSpaceId** | Pointer to **int32** | The database identifier (ID) of the space associated with the DNS zone the record belongs to. | [optional] 
**ZoneSpaceName** | Pointer to **string** | The name of the space associated with the DNS zone the record belongs to. | [optional] 
**RowState** | Pointer to **int32** | The object activation status.&lt;ul&gt;&lt;li&gt; If set to &lt;b&gt;0&lt;/b&gt;, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; If set to &lt;b&gt;1&lt;/b&gt;, the object is enabled and managed.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; If set to &lt;b&gt;2&lt;/b&gt;, the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_state&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 
**ZoneUseUpdatePolicy** | Pointer to **int32** | The update policy status of the DNS zone. Set it to &lt;b&gt;1&lt;/b&gt; to indicate that the DNS zone uses a specific GSS-TSIG/update-policy. You can only configure the zone update policy if the parameter &lt;b&gt;gss_enabled&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt;. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**ZoneClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**ZoneClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDnsZoneAddInput

`func NewDnsZoneAddInput() *DnsZoneAddInput`

NewDnsZoneAddInput instantiates a new DnsZoneAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneAddInputWithDefaults

`func NewDnsZoneAddInputWithDefaults() *DnsZoneAddInput`

NewDnsZoneAddInputWithDefaults instantiates a new DnsZoneAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DnsZoneAddInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsZoneAddInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsZoneAddInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsZoneAddInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsZoneAddInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsZoneAddInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsZoneAddInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsZoneAddInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetViewId

`func (o *DnsZoneAddInput) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DnsZoneAddInput) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DnsZoneAddInput) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DnsZoneAddInput) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewName

`func (o *DnsZoneAddInput) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DnsZoneAddInput) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DnsZoneAddInput) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DnsZoneAddInput) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetZoneName

`func (o *DnsZoneAddInput) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DnsZoneAddInput) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DnsZoneAddInput) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *DnsZoneAddInput) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneType

`func (o *DnsZoneAddInput) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *DnsZoneAddInput) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *DnsZoneAddInput) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *DnsZoneAddInput) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DnsZoneAddInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DnsZoneAddInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DnsZoneAddInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DnsZoneAddInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerDdnsScavenging

`func (o *DnsZoneAddInput) GetServerDdnsScavenging() int32`

GetServerDdnsScavenging returns the ServerDdnsScavenging field if non-nil, zero value otherwise.

### GetServerDdnsScavengingOk

`func (o *DnsZoneAddInput) GetServerDdnsScavengingOk() (*int32, bool)`

GetServerDdnsScavengingOk returns a tuple with the ServerDdnsScavenging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerDdnsScavenging

`func (o *DnsZoneAddInput) SetServerDdnsScavenging(v int32)`

SetServerDdnsScavenging sets ServerDdnsScavenging field to given value.

### HasServerDdnsScavenging

`func (o *DnsZoneAddInput) HasServerDdnsScavenging() bool`

HasServerDdnsScavenging returns a boolean if a field has been set.

### GetZoneAdIntegrated

`func (o *DnsZoneAddInput) GetZoneAdIntegrated() int32`

GetZoneAdIntegrated returns the ZoneAdIntegrated field if non-nil, zero value otherwise.

### GetZoneAdIntegratedOk

`func (o *DnsZoneAddInput) GetZoneAdIntegratedOk() (*int32, bool)`

GetZoneAdIntegratedOk returns a tuple with the ZoneAdIntegrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAdIntegrated

`func (o *DnsZoneAddInput) SetZoneAdIntegrated(v int32)`

SetZoneAdIntegrated sets ZoneAdIntegrated field to given value.

### HasZoneAdIntegrated

`func (o *DnsZoneAddInput) HasZoneAdIntegrated() bool`

HasZoneAdIntegrated returns a boolean if a field has been set.

### GetZoneAllowQuery

`func (o *DnsZoneAddInput) GetZoneAllowQuery() string`

GetZoneAllowQuery returns the ZoneAllowQuery field if non-nil, zero value otherwise.

### GetZoneAllowQueryOk

`func (o *DnsZoneAddInput) GetZoneAllowQueryOk() (*string, bool)`

GetZoneAllowQueryOk returns a tuple with the ZoneAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowQuery

`func (o *DnsZoneAddInput) SetZoneAllowQuery(v string)`

SetZoneAllowQuery sets ZoneAllowQuery field to given value.

### HasZoneAllowQuery

`func (o *DnsZoneAddInput) HasZoneAllowQuery() bool`

HasZoneAllowQuery returns a boolean if a field has been set.

### GetZoneAllowTransfer

`func (o *DnsZoneAddInput) GetZoneAllowTransfer() string`

GetZoneAllowTransfer returns the ZoneAllowTransfer field if non-nil, zero value otherwise.

### GetZoneAllowTransferOk

`func (o *DnsZoneAddInput) GetZoneAllowTransferOk() (*string, bool)`

GetZoneAllowTransferOk returns a tuple with the ZoneAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowTransfer

`func (o *DnsZoneAddInput) SetZoneAllowTransfer(v string)`

SetZoneAllowTransfer sets ZoneAllowTransfer field to given value.

### HasZoneAllowTransfer

`func (o *DnsZoneAddInput) HasZoneAllowTransfer() bool`

HasZoneAllowTransfer returns a boolean if a field has been set.

### GetZoneAllowUpdate

`func (o *DnsZoneAddInput) GetZoneAllowUpdate() string`

GetZoneAllowUpdate returns the ZoneAllowUpdate field if non-nil, zero value otherwise.

### GetZoneAllowUpdateOk

`func (o *DnsZoneAddInput) GetZoneAllowUpdateOk() (*string, bool)`

GetZoneAllowUpdateOk returns a tuple with the ZoneAllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowUpdate

`func (o *DnsZoneAddInput) SetZoneAllowUpdate(v string)`

SetZoneAllowUpdate sets ZoneAllowUpdate field to given value.

### HasZoneAllowUpdate

`func (o *DnsZoneAddInput) HasZoneAllowUpdate() bool`

HasZoneAllowUpdate returns a boolean if a field has been set.

### GetZoneAlsoNotify

`func (o *DnsZoneAddInput) GetZoneAlsoNotify() string`

GetZoneAlsoNotify returns the ZoneAlsoNotify field if non-nil, zero value otherwise.

### GetZoneAlsoNotifyOk

`func (o *DnsZoneAddInput) GetZoneAlsoNotifyOk() (*string, bool)`

GetZoneAlsoNotifyOk returns a tuple with the ZoneAlsoNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAlsoNotify

`func (o *DnsZoneAddInput) SetZoneAlsoNotify(v string)`

SetZoneAlsoNotify sets ZoneAlsoNotify field to given value.

### HasZoneAlsoNotify

`func (o *DnsZoneAddInput) HasZoneAlsoNotify() bool`

HasZoneAlsoNotify returns a boolean if a field has been set.

### GetZoneForward

`func (o *DnsZoneAddInput) GetZoneForward() string`

GetZoneForward returns the ZoneForward field if non-nil, zero value otherwise.

### GetZoneForwardOk

`func (o *DnsZoneAddInput) GetZoneForwardOk() (*string, bool)`

GetZoneForwardOk returns a tuple with the ZoneForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneForward

`func (o *DnsZoneAddInput) SetZoneForward(v string)`

SetZoneForward sets ZoneForward field to given value.

### HasZoneForward

`func (o *DnsZoneAddInput) HasZoneForward() bool`

HasZoneForward returns a boolean if a field has been set.

### GetZoneForwarders

`func (o *DnsZoneAddInput) GetZoneForwarders() string`

GetZoneForwarders returns the ZoneForwarders field if non-nil, zero value otherwise.

### GetZoneForwardersOk

`func (o *DnsZoneAddInput) GetZoneForwardersOk() (*string, bool)`

GetZoneForwardersOk returns a tuple with the ZoneForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneForwarders

`func (o *DnsZoneAddInput) SetZoneForwarders(v string)`

SetZoneForwarders sets ZoneForwarders field to given value.

### HasZoneForwarders

`func (o *DnsZoneAddInput) HasZoneForwarders() bool`

HasZoneForwarders returns a boolean if a field has been set.

### GetZoneIsRpz

`func (o *DnsZoneAddInput) GetZoneIsRpz() int32`

GetZoneIsRpz returns the ZoneIsRpz field if non-nil, zero value otherwise.

### GetZoneIsRpzOk

`func (o *DnsZoneAddInput) GetZoneIsRpzOk() (*int32, bool)`

GetZoneIsRpzOk returns a tuple with the ZoneIsRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIsRpz

`func (o *DnsZoneAddInput) SetZoneIsRpz(v int32)`

SetZoneIsRpz sets ZoneIsRpz field to given value.

### HasZoneIsRpz

`func (o *DnsZoneAddInput) HasZoneIsRpz() bool`

HasZoneIsRpz returns a boolean if a field has been set.

### GetZoneMasters

`func (o *DnsZoneAddInput) GetZoneMasters() string`

GetZoneMasters returns the ZoneMasters field if non-nil, zero value otherwise.

### GetZoneMastersOk

`func (o *DnsZoneAddInput) GetZoneMastersOk() (*string, bool)`

GetZoneMastersOk returns a tuple with the ZoneMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMasters

`func (o *DnsZoneAddInput) SetZoneMasters(v string)`

SetZoneMasters sets ZoneMasters field to given value.

### HasZoneMasters

`func (o *DnsZoneAddInput) HasZoneMasters() bool`

HasZoneMasters returns a boolean if a field has been set.

### GetZoneNotify

`func (o *DnsZoneAddInput) GetZoneNotify() string`

GetZoneNotify returns the ZoneNotify field if non-nil, zero value otherwise.

### GetZoneNotifyOk

`func (o *DnsZoneAddInput) GetZoneNotifyOk() (*string, bool)`

GetZoneNotifyOk returns a tuple with the ZoneNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNotify

`func (o *DnsZoneAddInput) SetZoneNotify(v string)`

SetZoneNotify sets ZoneNotify field to given value.

### HasZoneNotify

`func (o *DnsZoneAddInput) HasZoneNotify() bool`

HasZoneNotify returns a boolean if a field has been set.

### GetZoneOrder

`func (o *DnsZoneAddInput) GetZoneOrder() int32`

GetZoneOrder returns the ZoneOrder field if non-nil, zero value otherwise.

### GetZoneOrderOk

`func (o *DnsZoneAddInput) GetZoneOrderOk() (*int32, bool)`

GetZoneOrderOk returns a tuple with the ZoneOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneOrder

`func (o *DnsZoneAddInput) SetZoneOrder(v int32)`

SetZoneOrder sets ZoneOrder field to given value.

### HasZoneOrder

`func (o *DnsZoneAddInput) HasZoneOrder() bool`

HasZoneOrder returns a boolean if a field has been set.

### GetZoneResponsePolicy

`func (o *DnsZoneAddInput) GetZoneResponsePolicy() string`

GetZoneResponsePolicy returns the ZoneResponsePolicy field if non-nil, zero value otherwise.

### GetZoneResponsePolicyOk

`func (o *DnsZoneAddInput) GetZoneResponsePolicyOk() (*string, bool)`

GetZoneResponsePolicyOk returns a tuple with the ZoneResponsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneResponsePolicy

`func (o *DnsZoneAddInput) SetZoneResponsePolicy(v string)`

SetZoneResponsePolicy sets ZoneResponsePolicy field to given value.

### HasZoneResponsePolicy

`func (o *DnsZoneAddInput) HasZoneResponsePolicy() bool`

HasZoneResponsePolicy returns a boolean if a field has been set.

### GetZoneRpzLog

`func (o *DnsZoneAddInput) GetZoneRpzLog() int32`

GetZoneRpzLog returns the ZoneRpzLog field if non-nil, zero value otherwise.

### GetZoneRpzLogOk

`func (o *DnsZoneAddInput) GetZoneRpzLogOk() (*int32, bool)`

GetZoneRpzLogOk returns a tuple with the ZoneRpzLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRpzLog

`func (o *DnsZoneAddInput) SetZoneRpzLog(v int32)`

SetZoneRpzLog sets ZoneRpzLog field to given value.

### HasZoneRpzLog

`func (o *DnsZoneAddInput) HasZoneRpzLog() bool`

HasZoneRpzLog returns a boolean if a field has been set.

### GetZoneRpzMaxPolicyTtl

`func (o *DnsZoneAddInput) GetZoneRpzMaxPolicyTtl() int32`

GetZoneRpzMaxPolicyTtl returns the ZoneRpzMaxPolicyTtl field if non-nil, zero value otherwise.

### GetZoneRpzMaxPolicyTtlOk

`func (o *DnsZoneAddInput) GetZoneRpzMaxPolicyTtlOk() (*int32, bool)`

GetZoneRpzMaxPolicyTtlOk returns a tuple with the ZoneRpzMaxPolicyTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRpzMaxPolicyTtl

`func (o *DnsZoneAddInput) SetZoneRpzMaxPolicyTtl(v int32)`

SetZoneRpzMaxPolicyTtl sets ZoneRpzMaxPolicyTtl field to given value.

### HasZoneRpzMaxPolicyTtl

`func (o *DnsZoneAddInput) HasZoneRpzMaxPolicyTtl() bool`

HasZoneRpzMaxPolicyTtl returns a boolean if a field has been set.

### GetZoneRpzRecursiveOnly

`func (o *DnsZoneAddInput) GetZoneRpzRecursiveOnly() int32`

GetZoneRpzRecursiveOnly returns the ZoneRpzRecursiveOnly field if non-nil, zero value otherwise.

### GetZoneRpzRecursiveOnlyOk

`func (o *DnsZoneAddInput) GetZoneRpzRecursiveOnlyOk() (*int32, bool)`

GetZoneRpzRecursiveOnlyOk returns a tuple with the ZoneRpzRecursiveOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRpzRecursiveOnly

`func (o *DnsZoneAddInput) SetZoneRpzRecursiveOnly(v int32)`

SetZoneRpzRecursiveOnly sets ZoneRpzRecursiveOnly field to given value.

### HasZoneRpzRecursiveOnly

`func (o *DnsZoneAddInput) HasZoneRpzRecursiveOnly() bool`

HasZoneRpzRecursiveOnly returns a boolean if a field has been set.

### GetZoneSpaceId

`func (o *DnsZoneAddInput) GetZoneSpaceId() int32`

GetZoneSpaceId returns the ZoneSpaceId field if non-nil, zero value otherwise.

### GetZoneSpaceIdOk

`func (o *DnsZoneAddInput) GetZoneSpaceIdOk() (*int32, bool)`

GetZoneSpaceIdOk returns a tuple with the ZoneSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSpaceId

`func (o *DnsZoneAddInput) SetZoneSpaceId(v int32)`

SetZoneSpaceId sets ZoneSpaceId field to given value.

### HasZoneSpaceId

`func (o *DnsZoneAddInput) HasZoneSpaceId() bool`

HasZoneSpaceId returns a boolean if a field has been set.

### GetZoneSpaceName

`func (o *DnsZoneAddInput) GetZoneSpaceName() string`

GetZoneSpaceName returns the ZoneSpaceName field if non-nil, zero value otherwise.

### GetZoneSpaceNameOk

`func (o *DnsZoneAddInput) GetZoneSpaceNameOk() (*string, bool)`

GetZoneSpaceNameOk returns a tuple with the ZoneSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSpaceName

`func (o *DnsZoneAddInput) SetZoneSpaceName(v string)`

SetZoneSpaceName sets ZoneSpaceName field to given value.

### HasZoneSpaceName

`func (o *DnsZoneAddInput) HasZoneSpaceName() bool`

HasZoneSpaceName returns a boolean if a field has been set.

### GetRowState

`func (o *DnsZoneAddInput) GetRowState() int32`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *DnsZoneAddInput) GetRowStateOk() (*int32, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *DnsZoneAddInput) SetRowState(v int32)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *DnsZoneAddInput) HasRowState() bool`

HasRowState returns a boolean if a field has been set.

### GetZoneUseUpdatePolicy

`func (o *DnsZoneAddInput) GetZoneUseUpdatePolicy() int32`

GetZoneUseUpdatePolicy returns the ZoneUseUpdatePolicy field if non-nil, zero value otherwise.

### GetZoneUseUpdatePolicyOk

`func (o *DnsZoneAddInput) GetZoneUseUpdatePolicyOk() (*int32, bool)`

GetZoneUseUpdatePolicyOk returns a tuple with the ZoneUseUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUseUpdatePolicy

`func (o *DnsZoneAddInput) SetZoneUseUpdatePolicy(v int32)`

SetZoneUseUpdatePolicy sets ZoneUseUpdatePolicy field to given value.

### HasZoneUseUpdatePolicy

`func (o *DnsZoneAddInput) HasZoneUseUpdatePolicy() bool`

HasZoneUseUpdatePolicy returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DnsZoneAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DnsZoneAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DnsZoneAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DnsZoneAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetZoneClassName

`func (o *DnsZoneAddInput) GetZoneClassName() string`

GetZoneClassName returns the ZoneClassName field if non-nil, zero value otherwise.

### GetZoneClassNameOk

`func (o *DnsZoneAddInput) GetZoneClassNameOk() (*string, bool)`

GetZoneClassNameOk returns a tuple with the ZoneClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneClassName

`func (o *DnsZoneAddInput) SetZoneClassName(v string)`

SetZoneClassName sets ZoneClassName field to given value.

### HasZoneClassName

`func (o *DnsZoneAddInput) HasZoneClassName() bool`

HasZoneClassName returns a boolean if a field has been set.

### GetZoneClassParameters

`func (o *DnsZoneAddInput) GetZoneClassParameters() []ApiClassParameterInputEntry`

GetZoneClassParameters returns the ZoneClassParameters field if non-nil, zero value otherwise.

### GetZoneClassParametersOk

`func (o *DnsZoneAddInput) GetZoneClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetZoneClassParametersOk returns a tuple with the ZoneClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneClassParameters

`func (o *DnsZoneAddInput) SetZoneClassParameters(v []ApiClassParameterInputEntry)`

SetZoneClassParameters sets ZoneClassParameters field to given value.

### HasZoneClassParameters

`func (o *DnsZoneAddInput) HasZoneClassParameters() bool`

HasZoneClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DnsZoneAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DnsZoneAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DnsZoneAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DnsZoneAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


