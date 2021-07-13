# DnsZoneEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server. | [optional] 
**ViewId** | Pointer to **int32** | The database identifier (ID) of the DNS view the object belongs to. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view the object belongs to. | [optional] 
**ZoneId** | Pointer to **int32** | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify which DNS zone to edit. | [optional] 
**ZoneName** | Pointer to **string** | The name of the DNS zone, each DNS zone must have a unique name. For hint zones (&lt;b&gt;dnszone_type&lt;/b&gt;: &lt;b&gt;hint&lt;/b&gt;), you must specify &lt;b&gt;.&lt;/b&gt; (dot) as &lt;b&gt;dnszone_name&lt;/b&gt;. | [optional] 
**ZoneType** | Pointer to **string** | The type of the DNS zone, either &lt;b&gt;master&lt;/b&gt;, &lt;b&gt;slave&lt;/b&gt;, &lt;b&gt;forward&lt;/b&gt;, &lt;b&gt;stub&lt;/b&gt;, &lt;b&gt;hint&lt;/b&gt; or &lt;b&gt;delegation-only&lt;/b&gt;. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DNS server. | [optional] 
**ZoneAdIntegrated** | Pointer to **int32** | The AD integrated status of the DNS zone. Set it to &lt;b&gt;1&lt;/b&gt; to indicate that the DNS zone belongs to an Active Directory integrated Microsoft DNS server. | [optional] 
**ZoneAllowQuery** | Pointer to **string** | The ACL values associated with the allow-query configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAllowTransfer** | Pointer to **string** | The ACL values associated with the allow-transfer configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAllowUpdate** | Pointer to **string** | The ACL values associated with the allow-update configuration of the DNS zone, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ZoneAlsoNotify** | Pointer to **string** | The IP address and port of the DNS server managing the smart architecture the DNS zone belongs to. If the parameter &lt;b&gt;dnszone_notify&lt;/b&gt; is set to &lt;b&gt;yes&lt;/b&gt; or &lt;b&gt;explicit&lt;/b&gt;, the server specified is instantly notified of any slave zones updates. | [optional] 
**ZoneForward** | Pointer to **string** | The forwarding mode of the DNS zone.&lt;table&gt;&lt;caption&gt;dnszone_forward possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;first&lt;/td&gt;&lt;td &gt;The zone sends the queries to the forwarder(s). If no answer is returned, it attempts to answer the queries on its own.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;only&lt;/td&gt;&lt;td &gt;The zone only forwards the queries to the forwarder(s). Required by some reverse forward zones (e.g., in the case of private addresses).&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; If the parameter has no value, it indicates that the forwarding is disabled. | [optional] 
**ZoneForwarders** | Pointer to **string** | The IP address(es) of the forwarder(s) associated with the DNS zone. It lists the DNS servers to which any unknown query on this zone should be sent, as follows: &lt;b&gt;&amp;lt;ip_address1&amp;gt;;&amp;lt;ip_address2&amp;gt;;...&lt;/b&gt; . | [optional] 
**ZoneIsRpz** | Pointer to **int32** | The RPZ status of the DNS zone. Set it to &lt;b&gt;1&lt;/b&gt; to indicate that the DNS zone is a Response Policy Zone. | [optional] 
**ZoneMasters** | Pointer to **string** | For slave DNS zones, the IP address of the DNS server and, if relevant, the name of the DNS view that contain the master DNS zone, as follows: &lt;b&gt;&amp;lt;ip_addr&amp;gt;;&lt;/b&gt; or &lt;b&gt;&amp;lt;ip_addr&amp;gt; key &amp;lt;dnsview_name&amp;gt;;&lt;/b&gt; . | [optional] 
**ZoneNotify** | Pointer to **string** | The notify status of the DNS zone.&lt;table&gt;&lt;caption&gt;dnszone_notify possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;No notify message is sent.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;A notify message is sent to the name servers defined in the NS records of the zone and to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;explicit&lt;/td&gt;&lt;td &gt;A notify message is sent only to the IP address(es) specified in the parameter .&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt;The notify message is not sent to the server itself or to the primary server defined in the SOA record of the zone. | [optional] 
**ZoneOrder** | Pointer to **int32** | The level of the DNS RPZ zone, where 0 represents the highest level in the views hierarchy. The RPZ rules of each zone are reviewed following this order. For non-RPZ zones, that have their parameter zone_is_rpz set to &lt;b&gt;0&lt;/b&gt;, you do not need to set this parameter. | [optional] 
**ZoneResponsePolicy** | Pointer to **string** | The response policy of the DNS zone.&lt;table&gt;&lt;caption&gt;dnszone_response_policy possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Policy&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;given&lt;/td&gt;&lt;td &gt;All the rules specified in the RPZ zone are applied normally.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;disabled&lt;/td&gt;&lt;td &gt;The RPZ zone rules configuration is not applied. All the rules it contains are ignored.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;passthru&lt;/td&gt;&lt;td &gt;The rules specified in the RPZ matching the listed RR names are ignored, no matter the RPZ zone they belong to.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;nxdomain&lt;/td&gt;&lt;td &gt;The rules specified in the RPZ return an NXDOMAIN response.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;nodata&lt;/td&gt;&lt;td &gt;The rules specified in the RPZ return a NODATA response.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;cname &lt;domain-name&gt;&lt;/td&gt;&lt;td &gt;All the rules specified in the RPZ are redirected toward the specified domain name.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt;You can only add RPZ zones on EfficientIP or BIND DNS servers. | [optional] 
**ZoneRpzLog** | Pointer to **int32** | TODO:dns_zone_add.input.zone_rpz_log | [optional] 
**ZoneSpaceId** | Pointer to **int32** | The database identifier (ID) of the space associated with the DNS zone the record belongs to. | [optional] 
**ZoneSpaceName** | Pointer to **string** | The name of the space associated with the DNS zone the record belongs to. | [optional] 
**RowState** | Pointer to **int32** | The object activation status.&lt;ul class&#x3D;dashed &gt;&lt;li&gt;                                                If set to &lt;b&gt;0&lt;/b&gt;, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                If set to &lt;b&gt;1&lt;/b&gt;, the object is enabled and managed.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                If set to &lt;b&gt;2&lt;/b&gt;, the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_enabled&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 
**ZoneUseUpdatePolicy** | Pointer to **int32** | The update policy status of the DNS zone. Set it to &lt;b&gt;1&lt;/b&gt; to indicate that the DNS zone uses a specific GSS-TSIG/update-policy. You can only configure the zone update policy if the parameter &lt;b&gt;gss_enabled&lt;/b&gt; of the server it belongs to is set to &lt;b&gt;1&lt;/b&gt;. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**ZoneClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**ZoneClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDnsZoneEditInput

`func NewDnsZoneEditInput() *DnsZoneEditInput`

NewDnsZoneEditInput instantiates a new DnsZoneEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneEditInputWithDefaults

`func NewDnsZoneEditInputWithDefaults() *DnsZoneEditInput`

NewDnsZoneEditInputWithDefaults instantiates a new DnsZoneEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DnsZoneEditInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsZoneEditInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsZoneEditInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsZoneEditInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsZoneEditInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsZoneEditInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsZoneEditInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsZoneEditInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetViewId

`func (o *DnsZoneEditInput) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DnsZoneEditInput) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DnsZoneEditInput) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DnsZoneEditInput) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewName

`func (o *DnsZoneEditInput) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DnsZoneEditInput) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DnsZoneEditInput) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DnsZoneEditInput) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetZoneId

`func (o *DnsZoneEditInput) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnsZoneEditInput) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnsZoneEditInput) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DnsZoneEditInput) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneName

`func (o *DnsZoneEditInput) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DnsZoneEditInput) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DnsZoneEditInput) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *DnsZoneEditInput) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneType

`func (o *DnsZoneEditInput) GetZoneType() string`

GetZoneType returns the ZoneType field if non-nil, zero value otherwise.

### GetZoneTypeOk

`func (o *DnsZoneEditInput) GetZoneTypeOk() (*string, bool)`

GetZoneTypeOk returns a tuple with the ZoneType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneType

`func (o *DnsZoneEditInput) SetZoneType(v string)`

SetZoneType sets ZoneType field to given value.

### HasZoneType

`func (o *DnsZoneEditInput) HasZoneType() bool`

HasZoneType returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DnsZoneEditInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DnsZoneEditInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DnsZoneEditInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DnsZoneEditInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetZoneAdIntegrated

`func (o *DnsZoneEditInput) GetZoneAdIntegrated() int32`

GetZoneAdIntegrated returns the ZoneAdIntegrated field if non-nil, zero value otherwise.

### GetZoneAdIntegratedOk

`func (o *DnsZoneEditInput) GetZoneAdIntegratedOk() (*int32, bool)`

GetZoneAdIntegratedOk returns a tuple with the ZoneAdIntegrated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAdIntegrated

`func (o *DnsZoneEditInput) SetZoneAdIntegrated(v int32)`

SetZoneAdIntegrated sets ZoneAdIntegrated field to given value.

### HasZoneAdIntegrated

`func (o *DnsZoneEditInput) HasZoneAdIntegrated() bool`

HasZoneAdIntegrated returns a boolean if a field has been set.

### GetZoneAllowQuery

`func (o *DnsZoneEditInput) GetZoneAllowQuery() string`

GetZoneAllowQuery returns the ZoneAllowQuery field if non-nil, zero value otherwise.

### GetZoneAllowQueryOk

`func (o *DnsZoneEditInput) GetZoneAllowQueryOk() (*string, bool)`

GetZoneAllowQueryOk returns a tuple with the ZoneAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowQuery

`func (o *DnsZoneEditInput) SetZoneAllowQuery(v string)`

SetZoneAllowQuery sets ZoneAllowQuery field to given value.

### HasZoneAllowQuery

`func (o *DnsZoneEditInput) HasZoneAllowQuery() bool`

HasZoneAllowQuery returns a boolean if a field has been set.

### GetZoneAllowTransfer

`func (o *DnsZoneEditInput) GetZoneAllowTransfer() string`

GetZoneAllowTransfer returns the ZoneAllowTransfer field if non-nil, zero value otherwise.

### GetZoneAllowTransferOk

`func (o *DnsZoneEditInput) GetZoneAllowTransferOk() (*string, bool)`

GetZoneAllowTransferOk returns a tuple with the ZoneAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowTransfer

`func (o *DnsZoneEditInput) SetZoneAllowTransfer(v string)`

SetZoneAllowTransfer sets ZoneAllowTransfer field to given value.

### HasZoneAllowTransfer

`func (o *DnsZoneEditInput) HasZoneAllowTransfer() bool`

HasZoneAllowTransfer returns a boolean if a field has been set.

### GetZoneAllowUpdate

`func (o *DnsZoneEditInput) GetZoneAllowUpdate() string`

GetZoneAllowUpdate returns the ZoneAllowUpdate field if non-nil, zero value otherwise.

### GetZoneAllowUpdateOk

`func (o *DnsZoneEditInput) GetZoneAllowUpdateOk() (*string, bool)`

GetZoneAllowUpdateOk returns a tuple with the ZoneAllowUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAllowUpdate

`func (o *DnsZoneEditInput) SetZoneAllowUpdate(v string)`

SetZoneAllowUpdate sets ZoneAllowUpdate field to given value.

### HasZoneAllowUpdate

`func (o *DnsZoneEditInput) HasZoneAllowUpdate() bool`

HasZoneAllowUpdate returns a boolean if a field has been set.

### GetZoneAlsoNotify

`func (o *DnsZoneEditInput) GetZoneAlsoNotify() string`

GetZoneAlsoNotify returns the ZoneAlsoNotify field if non-nil, zero value otherwise.

### GetZoneAlsoNotifyOk

`func (o *DnsZoneEditInput) GetZoneAlsoNotifyOk() (*string, bool)`

GetZoneAlsoNotifyOk returns a tuple with the ZoneAlsoNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneAlsoNotify

`func (o *DnsZoneEditInput) SetZoneAlsoNotify(v string)`

SetZoneAlsoNotify sets ZoneAlsoNotify field to given value.

### HasZoneAlsoNotify

`func (o *DnsZoneEditInput) HasZoneAlsoNotify() bool`

HasZoneAlsoNotify returns a boolean if a field has been set.

### GetZoneForward

`func (o *DnsZoneEditInput) GetZoneForward() string`

GetZoneForward returns the ZoneForward field if non-nil, zero value otherwise.

### GetZoneForwardOk

`func (o *DnsZoneEditInput) GetZoneForwardOk() (*string, bool)`

GetZoneForwardOk returns a tuple with the ZoneForward field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneForward

`func (o *DnsZoneEditInput) SetZoneForward(v string)`

SetZoneForward sets ZoneForward field to given value.

### HasZoneForward

`func (o *DnsZoneEditInput) HasZoneForward() bool`

HasZoneForward returns a boolean if a field has been set.

### GetZoneForwarders

`func (o *DnsZoneEditInput) GetZoneForwarders() string`

GetZoneForwarders returns the ZoneForwarders field if non-nil, zero value otherwise.

### GetZoneForwardersOk

`func (o *DnsZoneEditInput) GetZoneForwardersOk() (*string, bool)`

GetZoneForwardersOk returns a tuple with the ZoneForwarders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneForwarders

`func (o *DnsZoneEditInput) SetZoneForwarders(v string)`

SetZoneForwarders sets ZoneForwarders field to given value.

### HasZoneForwarders

`func (o *DnsZoneEditInput) HasZoneForwarders() bool`

HasZoneForwarders returns a boolean if a field has been set.

### GetZoneIsRpz

`func (o *DnsZoneEditInput) GetZoneIsRpz() int32`

GetZoneIsRpz returns the ZoneIsRpz field if non-nil, zero value otherwise.

### GetZoneIsRpzOk

`func (o *DnsZoneEditInput) GetZoneIsRpzOk() (*int32, bool)`

GetZoneIsRpzOk returns a tuple with the ZoneIsRpz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIsRpz

`func (o *DnsZoneEditInput) SetZoneIsRpz(v int32)`

SetZoneIsRpz sets ZoneIsRpz field to given value.

### HasZoneIsRpz

`func (o *DnsZoneEditInput) HasZoneIsRpz() bool`

HasZoneIsRpz returns a boolean if a field has been set.

### GetZoneMasters

`func (o *DnsZoneEditInput) GetZoneMasters() string`

GetZoneMasters returns the ZoneMasters field if non-nil, zero value otherwise.

### GetZoneMastersOk

`func (o *DnsZoneEditInput) GetZoneMastersOk() (*string, bool)`

GetZoneMastersOk returns a tuple with the ZoneMasters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneMasters

`func (o *DnsZoneEditInput) SetZoneMasters(v string)`

SetZoneMasters sets ZoneMasters field to given value.

### HasZoneMasters

`func (o *DnsZoneEditInput) HasZoneMasters() bool`

HasZoneMasters returns a boolean if a field has been set.

### GetZoneNotify

`func (o *DnsZoneEditInput) GetZoneNotify() string`

GetZoneNotify returns the ZoneNotify field if non-nil, zero value otherwise.

### GetZoneNotifyOk

`func (o *DnsZoneEditInput) GetZoneNotifyOk() (*string, bool)`

GetZoneNotifyOk returns a tuple with the ZoneNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneNotify

`func (o *DnsZoneEditInput) SetZoneNotify(v string)`

SetZoneNotify sets ZoneNotify field to given value.

### HasZoneNotify

`func (o *DnsZoneEditInput) HasZoneNotify() bool`

HasZoneNotify returns a boolean if a field has been set.

### GetZoneOrder

`func (o *DnsZoneEditInput) GetZoneOrder() int32`

GetZoneOrder returns the ZoneOrder field if non-nil, zero value otherwise.

### GetZoneOrderOk

`func (o *DnsZoneEditInput) GetZoneOrderOk() (*int32, bool)`

GetZoneOrderOk returns a tuple with the ZoneOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneOrder

`func (o *DnsZoneEditInput) SetZoneOrder(v int32)`

SetZoneOrder sets ZoneOrder field to given value.

### HasZoneOrder

`func (o *DnsZoneEditInput) HasZoneOrder() bool`

HasZoneOrder returns a boolean if a field has been set.

### GetZoneResponsePolicy

`func (o *DnsZoneEditInput) GetZoneResponsePolicy() string`

GetZoneResponsePolicy returns the ZoneResponsePolicy field if non-nil, zero value otherwise.

### GetZoneResponsePolicyOk

`func (o *DnsZoneEditInput) GetZoneResponsePolicyOk() (*string, bool)`

GetZoneResponsePolicyOk returns a tuple with the ZoneResponsePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneResponsePolicy

`func (o *DnsZoneEditInput) SetZoneResponsePolicy(v string)`

SetZoneResponsePolicy sets ZoneResponsePolicy field to given value.

### HasZoneResponsePolicy

`func (o *DnsZoneEditInput) HasZoneResponsePolicy() bool`

HasZoneResponsePolicy returns a boolean if a field has been set.

### GetZoneRpzLog

`func (o *DnsZoneEditInput) GetZoneRpzLog() int32`

GetZoneRpzLog returns the ZoneRpzLog field if non-nil, zero value otherwise.

### GetZoneRpzLogOk

`func (o *DnsZoneEditInput) GetZoneRpzLogOk() (*int32, bool)`

GetZoneRpzLogOk returns a tuple with the ZoneRpzLog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneRpzLog

`func (o *DnsZoneEditInput) SetZoneRpzLog(v int32)`

SetZoneRpzLog sets ZoneRpzLog field to given value.

### HasZoneRpzLog

`func (o *DnsZoneEditInput) HasZoneRpzLog() bool`

HasZoneRpzLog returns a boolean if a field has been set.

### GetZoneSpaceId

`func (o *DnsZoneEditInput) GetZoneSpaceId() int32`

GetZoneSpaceId returns the ZoneSpaceId field if non-nil, zero value otherwise.

### GetZoneSpaceIdOk

`func (o *DnsZoneEditInput) GetZoneSpaceIdOk() (*int32, bool)`

GetZoneSpaceIdOk returns a tuple with the ZoneSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSpaceId

`func (o *DnsZoneEditInput) SetZoneSpaceId(v int32)`

SetZoneSpaceId sets ZoneSpaceId field to given value.

### HasZoneSpaceId

`func (o *DnsZoneEditInput) HasZoneSpaceId() bool`

HasZoneSpaceId returns a boolean if a field has been set.

### GetZoneSpaceName

`func (o *DnsZoneEditInput) GetZoneSpaceName() string`

GetZoneSpaceName returns the ZoneSpaceName field if non-nil, zero value otherwise.

### GetZoneSpaceNameOk

`func (o *DnsZoneEditInput) GetZoneSpaceNameOk() (*string, bool)`

GetZoneSpaceNameOk returns a tuple with the ZoneSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSpaceName

`func (o *DnsZoneEditInput) SetZoneSpaceName(v string)`

SetZoneSpaceName sets ZoneSpaceName field to given value.

### HasZoneSpaceName

`func (o *DnsZoneEditInput) HasZoneSpaceName() bool`

HasZoneSpaceName returns a boolean if a field has been set.

### GetRowState

`func (o *DnsZoneEditInput) GetRowState() int32`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *DnsZoneEditInput) GetRowStateOk() (*int32, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *DnsZoneEditInput) SetRowState(v int32)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *DnsZoneEditInput) HasRowState() bool`

HasRowState returns a boolean if a field has been set.

### GetZoneUseUpdatePolicy

`func (o *DnsZoneEditInput) GetZoneUseUpdatePolicy() int32`

GetZoneUseUpdatePolicy returns the ZoneUseUpdatePolicy field if non-nil, zero value otherwise.

### GetZoneUseUpdatePolicyOk

`func (o *DnsZoneEditInput) GetZoneUseUpdatePolicyOk() (*int32, bool)`

GetZoneUseUpdatePolicyOk returns a tuple with the ZoneUseUpdatePolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneUseUpdatePolicy

`func (o *DnsZoneEditInput) SetZoneUseUpdatePolicy(v int32)`

SetZoneUseUpdatePolicy sets ZoneUseUpdatePolicy field to given value.

### HasZoneUseUpdatePolicy

`func (o *DnsZoneEditInput) HasZoneUseUpdatePolicy() bool`

HasZoneUseUpdatePolicy returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DnsZoneEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DnsZoneEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DnsZoneEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DnsZoneEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetZoneClassName

`func (o *DnsZoneEditInput) GetZoneClassName() string`

GetZoneClassName returns the ZoneClassName field if non-nil, zero value otherwise.

### GetZoneClassNameOk

`func (o *DnsZoneEditInput) GetZoneClassNameOk() (*string, bool)`

GetZoneClassNameOk returns a tuple with the ZoneClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneClassName

`func (o *DnsZoneEditInput) SetZoneClassName(v string)`

SetZoneClassName sets ZoneClassName field to given value.

### HasZoneClassName

`func (o *DnsZoneEditInput) HasZoneClassName() bool`

HasZoneClassName returns a boolean if a field has been set.

### GetZoneClassParameters

`func (o *DnsZoneEditInput) GetZoneClassParameters() []ApiClassParameterInputEntry`

GetZoneClassParameters returns the ZoneClassParameters field if non-nil, zero value otherwise.

### GetZoneClassParametersOk

`func (o *DnsZoneEditInput) GetZoneClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetZoneClassParametersOk returns a tuple with the ZoneClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneClassParameters

`func (o *DnsZoneEditInput) SetZoneClassParameters(v []ApiClassParameterInputEntry)`

SetZoneClassParameters sets ZoneClassParameters field to given value.

### HasZoneClassParameters

`func (o *DnsZoneEditInput) HasZoneClassParameters() bool`

HasZoneClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DnsZoneEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DnsZoneEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DnsZoneEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DnsZoneEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


