# DnsViewDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**ViewDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DNS server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ServerCloud** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerComment** | Pointer to **string** | The description of the DNS server the object belongs to. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DNS server the object belongs to.&lt;table&gt;&lt;caption&gt;dns_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DNS server or EfficientIP DNS Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msdaemon&lt;/td&gt;&lt;td &gt;Agentless Microsoft DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ans&lt;/td&gt;&lt;td &gt;Nominum DNS server (ANS)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;aws&lt;/td&gt;&lt;td &gt;Amazon Route 53 server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;other&lt;/td&gt;&lt;td &gt;Generic DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdns&lt;/td&gt;&lt;td &gt;EfficientIP DNS smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DNS server the object belongs to. | [optional] 
**ViewAllowQuery** | Pointer to **string** | The ACL values associated with the allow-query configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewAllowRecursion** | Pointer to **string** | The ACL values associated with the allow-recursion configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewAllowTransfer** | Pointer to **string** | The ACL values associated with the allow-transfer configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewClassName** | Pointer to **string** | The name of the class applied to the DNS view, it can be preceded by the class directory. | [optional] 
**ViewClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ViewId** | Pointer to **string** | The database identifier (ID) of the DNS view. | [optional] 
**ViewKeyName** | Pointer to **string** | The name of the DNS TSIG key associated with the DNS view. | [optional] 
**ViewMatchClients** | Pointer to **string** | The ACL values associated with the match clients configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;...&lt;/b&gt; . Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewMatchTo** | Pointer to **string** | The ACL values associated with the match destination configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;...&lt;/b&gt; . Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view. | [optional] 
**ViewOrder** | Pointer to **string** | The level of the DNS view, where &lt;b&gt;0&lt;/b&gt; represents the highest level in the views hierarchy. The parameters &lt;b&gt;dnsview_match_client&lt;/b&gt; and &lt;b&gt;dnsview_match_to&lt;/b&gt; of each view in a server are reviewed following this order. | [optional] 
**ViewRecursion** | Pointer to **string** | The recursion status of the DNS view:&lt;table&gt;&lt;caption&gt;dnsview_recursion possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;The view only provides iterative query behavior - normally resulting in a referral. If the answer to the query already exists in the cache it will be returned whatever the value of this statement.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;The view always provides recursive query behavior if requested by the client.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerGssKeytabId** | Pointer to **string** | The database identifier (ID) of the DNS GSS-TSIG keytab. | [optional] 
**ServerAddr** | Pointer to **string** | The IP address of the DNS server the object belongs to, in hexadecimal format. | [optional] 
**ViewMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DNS smart architecture managing the DNS server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DNS smart architecture managing the DNS server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDnsViewDataData

`func NewDnsViewDataData() *DnsViewDataData`

NewDnsViewDataData instantiates a new DnsViewDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsViewDataDataWithDefaults

`func NewDnsViewDataDataWithDefaults() *DnsViewDataData`

NewDnsViewDataDataWithDefaults instantiates a new DnsViewDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewDelayedCreateTime

`func (o *DnsViewDataData) GetViewDelayedCreateTime() string`

GetViewDelayedCreateTime returns the ViewDelayedCreateTime field if non-nil, zero value otherwise.

### GetViewDelayedCreateTimeOk

`func (o *DnsViewDataData) GetViewDelayedCreateTimeOk() (*string, bool)`

GetViewDelayedCreateTimeOk returns a tuple with the ViewDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDelayedCreateTime

`func (o *DnsViewDataData) SetViewDelayedCreateTime(v string)`

SetViewDelayedCreateTime sets ViewDelayedCreateTime field to given value.

### HasViewDelayedCreateTime

`func (o *DnsViewDataData) HasViewDelayedCreateTime() bool`

HasViewDelayedCreateTime returns a boolean if a field has been set.

### GetViewDelayedDeleteTime

`func (o *DnsViewDataData) GetViewDelayedDeleteTime() string`

GetViewDelayedDeleteTime returns the ViewDelayedDeleteTime field if non-nil, zero value otherwise.

### GetViewDelayedDeleteTimeOk

`func (o *DnsViewDataData) GetViewDelayedDeleteTimeOk() (*string, bool)`

GetViewDelayedDeleteTimeOk returns a tuple with the ViewDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDelayedDeleteTime

`func (o *DnsViewDataData) SetViewDelayedDeleteTime(v string)`

SetViewDelayedDeleteTime sets ViewDelayedDeleteTime field to given value.

### HasViewDelayedDeleteTime

`func (o *DnsViewDataData) HasViewDelayedDeleteTime() bool`

HasViewDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DnsViewDataData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DnsViewDataData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DnsViewDataData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DnsViewDataData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DnsViewDataData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DnsViewDataData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DnsViewDataData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DnsViewDataData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerCloud

`func (o *DnsViewDataData) GetServerCloud() string`

GetServerCloud returns the ServerCloud field if non-nil, zero value otherwise.

### GetServerCloudOk

`func (o *DnsViewDataData) GetServerCloudOk() (*string, bool)`

GetServerCloudOk returns a tuple with the ServerCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCloud

`func (o *DnsViewDataData) SetServerCloud(v string)`

SetServerCloud sets ServerCloud field to given value.

### HasServerCloud

`func (o *DnsViewDataData) HasServerCloud() bool`

HasServerCloud returns a boolean if a field has been set.

### GetServerComment

`func (o *DnsViewDataData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DnsViewDataData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DnsViewDataData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DnsViewDataData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerId

`func (o *DnsViewDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsViewDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsViewDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsViewDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsViewDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsViewDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsViewDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsViewDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DnsViewDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DnsViewDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DnsViewDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DnsViewDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DnsViewDataData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DnsViewDataData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DnsViewDataData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DnsViewDataData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetViewAllowQuery

`func (o *DnsViewDataData) GetViewAllowQuery() string`

GetViewAllowQuery returns the ViewAllowQuery field if non-nil, zero value otherwise.

### GetViewAllowQueryOk

`func (o *DnsViewDataData) GetViewAllowQueryOk() (*string, bool)`

GetViewAllowQueryOk returns a tuple with the ViewAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowQuery

`func (o *DnsViewDataData) SetViewAllowQuery(v string)`

SetViewAllowQuery sets ViewAllowQuery field to given value.

### HasViewAllowQuery

`func (o *DnsViewDataData) HasViewAllowQuery() bool`

HasViewAllowQuery returns a boolean if a field has been set.

### GetViewAllowRecursion

`func (o *DnsViewDataData) GetViewAllowRecursion() string`

GetViewAllowRecursion returns the ViewAllowRecursion field if non-nil, zero value otherwise.

### GetViewAllowRecursionOk

`func (o *DnsViewDataData) GetViewAllowRecursionOk() (*string, bool)`

GetViewAllowRecursionOk returns a tuple with the ViewAllowRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowRecursion

`func (o *DnsViewDataData) SetViewAllowRecursion(v string)`

SetViewAllowRecursion sets ViewAllowRecursion field to given value.

### HasViewAllowRecursion

`func (o *DnsViewDataData) HasViewAllowRecursion() bool`

HasViewAllowRecursion returns a boolean if a field has been set.

### GetViewAllowTransfer

`func (o *DnsViewDataData) GetViewAllowTransfer() string`

GetViewAllowTransfer returns the ViewAllowTransfer field if non-nil, zero value otherwise.

### GetViewAllowTransferOk

`func (o *DnsViewDataData) GetViewAllowTransferOk() (*string, bool)`

GetViewAllowTransferOk returns a tuple with the ViewAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowTransfer

`func (o *DnsViewDataData) SetViewAllowTransfer(v string)`

SetViewAllowTransfer sets ViewAllowTransfer field to given value.

### HasViewAllowTransfer

`func (o *DnsViewDataData) HasViewAllowTransfer() bool`

HasViewAllowTransfer returns a boolean if a field has been set.

### GetViewClassName

`func (o *DnsViewDataData) GetViewClassName() string`

GetViewClassName returns the ViewClassName field if non-nil, zero value otherwise.

### GetViewClassNameOk

`func (o *DnsViewDataData) GetViewClassNameOk() (*string, bool)`

GetViewClassNameOk returns a tuple with the ViewClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassName

`func (o *DnsViewDataData) SetViewClassName(v string)`

SetViewClassName sets ViewClassName field to given value.

### HasViewClassName

`func (o *DnsViewDataData) HasViewClassName() bool`

HasViewClassName returns a boolean if a field has been set.

### GetViewClassParameters

`func (o *DnsViewDataData) GetViewClassParameters() []ApiClassParameterOutputEntry`

GetViewClassParameters returns the ViewClassParameters field if non-nil, zero value otherwise.

### GetViewClassParametersOk

`func (o *DnsViewDataData) GetViewClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetViewClassParametersOk returns a tuple with the ViewClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassParameters

`func (o *DnsViewDataData) SetViewClassParameters(v []ApiClassParameterOutputEntry)`

SetViewClassParameters sets ViewClassParameters field to given value.

### HasViewClassParameters

`func (o *DnsViewDataData) HasViewClassParameters() bool`

HasViewClassParameters returns a boolean if a field has been set.

### GetViewId

`func (o *DnsViewDataData) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DnsViewDataData) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DnsViewDataData) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DnsViewDataData) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewKeyName

`func (o *DnsViewDataData) GetViewKeyName() string`

GetViewKeyName returns the ViewKeyName field if non-nil, zero value otherwise.

### GetViewKeyNameOk

`func (o *DnsViewDataData) GetViewKeyNameOk() (*string, bool)`

GetViewKeyNameOk returns a tuple with the ViewKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewKeyName

`func (o *DnsViewDataData) SetViewKeyName(v string)`

SetViewKeyName sets ViewKeyName field to given value.

### HasViewKeyName

`func (o *DnsViewDataData) HasViewKeyName() bool`

HasViewKeyName returns a boolean if a field has been set.

### GetViewMatchClients

`func (o *DnsViewDataData) GetViewMatchClients() string`

GetViewMatchClients returns the ViewMatchClients field if non-nil, zero value otherwise.

### GetViewMatchClientsOk

`func (o *DnsViewDataData) GetViewMatchClientsOk() (*string, bool)`

GetViewMatchClientsOk returns a tuple with the ViewMatchClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatchClients

`func (o *DnsViewDataData) SetViewMatchClients(v string)`

SetViewMatchClients sets ViewMatchClients field to given value.

### HasViewMatchClients

`func (o *DnsViewDataData) HasViewMatchClients() bool`

HasViewMatchClients returns a boolean if a field has been set.

### GetViewMatchTo

`func (o *DnsViewDataData) GetViewMatchTo() string`

GetViewMatchTo returns the ViewMatchTo field if non-nil, zero value otherwise.

### GetViewMatchToOk

`func (o *DnsViewDataData) GetViewMatchToOk() (*string, bool)`

GetViewMatchToOk returns a tuple with the ViewMatchTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatchTo

`func (o *DnsViewDataData) SetViewMatchTo(v string)`

SetViewMatchTo sets ViewMatchTo field to given value.

### HasViewMatchTo

`func (o *DnsViewDataData) HasViewMatchTo() bool`

HasViewMatchTo returns a boolean if a field has been set.

### GetViewName

`func (o *DnsViewDataData) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DnsViewDataData) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DnsViewDataData) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DnsViewDataData) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetViewOrder

`func (o *DnsViewDataData) GetViewOrder() string`

GetViewOrder returns the ViewOrder field if non-nil, zero value otherwise.

### GetViewOrderOk

`func (o *DnsViewDataData) GetViewOrderOk() (*string, bool)`

GetViewOrderOk returns a tuple with the ViewOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewOrder

`func (o *DnsViewDataData) SetViewOrder(v string)`

SetViewOrder sets ViewOrder field to given value.

### HasViewOrder

`func (o *DnsViewDataData) HasViewOrder() bool`

HasViewOrder returns a boolean if a field has been set.

### GetViewRecursion

`func (o *DnsViewDataData) GetViewRecursion() string`

GetViewRecursion returns the ViewRecursion field if non-nil, zero value otherwise.

### GetViewRecursionOk

`func (o *DnsViewDataData) GetViewRecursionOk() (*string, bool)`

GetViewRecursionOk returns a tuple with the ViewRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRecursion

`func (o *DnsViewDataData) SetViewRecursion(v string)`

SetViewRecursion sets ViewRecursion field to given value.

### HasViewRecursion

`func (o *DnsViewDataData) HasViewRecursion() bool`

HasViewRecursion returns a boolean if a field has been set.

### GetServerGssKeytabId

`func (o *DnsViewDataData) GetServerGssKeytabId() string`

GetServerGssKeytabId returns the ServerGssKeytabId field if non-nil, zero value otherwise.

### GetServerGssKeytabIdOk

`func (o *DnsViewDataData) GetServerGssKeytabIdOk() (*string, bool)`

GetServerGssKeytabIdOk returns a tuple with the ServerGssKeytabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGssKeytabId

`func (o *DnsViewDataData) SetServerGssKeytabId(v string)`

SetServerGssKeytabId sets ServerGssKeytabId field to given value.

### HasServerGssKeytabId

`func (o *DnsViewDataData) HasServerGssKeytabId() bool`

HasServerGssKeytabId returns a boolean if a field has been set.

### GetServerAddr

`func (o *DnsViewDataData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DnsViewDataData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DnsViewDataData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DnsViewDataData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetViewMultistatus

`func (o *DnsViewDataData) GetViewMultistatus() string`

GetViewMultistatus returns the ViewMultistatus field if non-nil, zero value otherwise.

### GetViewMultistatusOk

`func (o *DnsViewDataData) GetViewMultistatusOk() (*string, bool)`

GetViewMultistatusOk returns a tuple with the ViewMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMultistatus

`func (o *DnsViewDataData) SetViewMultistatus(v string)`

SetViewMultistatus sets ViewMultistatus field to given value.

### HasViewMultistatus

`func (o *DnsViewDataData) HasViewMultistatus() bool`

HasViewMultistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DnsViewDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DnsViewDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DnsViewDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DnsViewDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DnsViewDataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DnsViewDataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DnsViewDataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DnsViewDataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


