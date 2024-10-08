# DataInnerDnsViewData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**ViewDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DNS server the object belongs to, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DNS server the object belongs to. | [optional] 
**ServerCloud** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerComment** | Pointer to **string** | The description of the DNS server the object belongs to. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DNS server the object belongs to.&lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msdaemon&lt;/td&gt;&lt;td &gt;Microsoft Windows DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;aws&lt;/td&gt;&lt;td &gt;Amazon Route 53 server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;other&lt;/td&gt;&lt;td &gt;Generic DNS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdns&lt;/td&gt;&lt;td &gt;EfficientIP DNS smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DNS server the object belongs to. | [optional] 
**ViewAllowQuery** | Pointer to **string** | The ACL values associated with the allow-query configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewAllowRecursion** | Pointer to **string** | The ACL values associated with the allow-recursion configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewAllowTransfer** | Pointer to **string** | The ACL values associated with the allow-transfer configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewClassName** | Pointer to **string** | The name of the class applied to the DNS view, it can be preceded by the class directory. | [optional] 
**ViewClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DNS view. | [optional] 
**ViewId** | Pointer to **string** | The database identifier (ID) of the DNS view. | [optional] 
**ViewKeyName** | Pointer to **string** | The name of the DNS TSIG key associated with the DNS view. | [optional] 
**ViewMatchClients** | Pointer to **string** | The ACL values associated with the match clients configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;...&lt;/b&gt; . Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewMatchTo** | Pointer to **string** | The ACL values associated with the match destination configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;...&lt;/b&gt; . Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view. | [optional] 
**ViewOrder** | Pointer to **string** | The level of the DNS view, where &lt;b&gt;0&lt;/b&gt; represents the highest level in the views hierarchy. The parameters &lt;b&gt;dnsview_match_client&lt;/b&gt; and &lt;b&gt;view_match_to&lt;/b&gt; of each view in a server are reviewed following this order. | [optional] 
**ViewRecursion** | Pointer to **string** | The recursion status of the DNS view:&lt;table&gt;&lt;caption&gt;view_recursion possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;The view only provides iterative query behavior - normally resulting in a referral. If the answer to the query already exists in the cache it will be returned whatever the value of this statement.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;The view always provides recursive query behavior if requested by the client.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerGssKeytabId** | Pointer to **string** | The database identifier (ID) of the DNS GSS-TSIG keytab. | [optional] 
**ServerHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server_addr&lt;/b&gt; or &lt;b&gt;server_addr6&lt;/b&gt;. | [optional] 
**ServerAddr6** | Pointer to **string** | The IPv6 address of the DNS server the object belongs to, in hexadecimal format. | [optional] 
**ServerAddr** | Pointer to **string** | The IPv4 address of the DNS server the object belongs to, in hexadecimal format. | [optional] 
**ViewMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DNS smart architecture managing the DNS server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DNS smart architecture managing the DNS server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDnsViewData

`func NewDataInnerDnsViewData() *DataInnerDnsViewData`

NewDataInnerDnsViewData instantiates a new DataInnerDnsViewData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDnsViewDataWithDefaults

`func NewDataInnerDnsViewDataWithDefaults() *DataInnerDnsViewData`

NewDataInnerDnsViewDataWithDefaults instantiates a new DataInnerDnsViewData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewDelayedCreateTime

`func (o *DataInnerDnsViewData) GetViewDelayedCreateTime() string`

GetViewDelayedCreateTime returns the ViewDelayedCreateTime field if non-nil, zero value otherwise.

### GetViewDelayedCreateTimeOk

`func (o *DataInnerDnsViewData) GetViewDelayedCreateTimeOk() (*string, bool)`

GetViewDelayedCreateTimeOk returns a tuple with the ViewDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDelayedCreateTime

`func (o *DataInnerDnsViewData) SetViewDelayedCreateTime(v string)`

SetViewDelayedCreateTime sets ViewDelayedCreateTime field to given value.

### HasViewDelayedCreateTime

`func (o *DataInnerDnsViewData) HasViewDelayedCreateTime() bool`

HasViewDelayedCreateTime returns a boolean if a field has been set.

### GetViewDelayedDeleteTime

`func (o *DataInnerDnsViewData) GetViewDelayedDeleteTime() string`

GetViewDelayedDeleteTime returns the ViewDelayedDeleteTime field if non-nil, zero value otherwise.

### GetViewDelayedDeleteTimeOk

`func (o *DataInnerDnsViewData) GetViewDelayedDeleteTimeOk() (*string, bool)`

GetViewDelayedDeleteTimeOk returns a tuple with the ViewDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewDelayedDeleteTime

`func (o *DataInnerDnsViewData) SetViewDelayedDeleteTime(v string)`

SetViewDelayedDeleteTime sets ViewDelayedDeleteTime field to given value.

### HasViewDelayedDeleteTime

`func (o *DataInnerDnsViewData) HasViewDelayedDeleteTime() bool`

HasViewDelayedDeleteTime returns a boolean if a field has been set.

### GetServerClassName

`func (o *DataInnerDnsViewData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DataInnerDnsViewData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DataInnerDnsViewData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DataInnerDnsViewData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DataInnerDnsViewData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DataInnerDnsViewData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DataInnerDnsViewData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DataInnerDnsViewData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerCloud

`func (o *DataInnerDnsViewData) GetServerCloud() string`

GetServerCloud returns the ServerCloud field if non-nil, zero value otherwise.

### GetServerCloudOk

`func (o *DataInnerDnsViewData) GetServerCloudOk() (*string, bool)`

GetServerCloudOk returns a tuple with the ServerCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCloud

`func (o *DataInnerDnsViewData) SetServerCloud(v string)`

SetServerCloud sets ServerCloud field to given value.

### HasServerCloud

`func (o *DataInnerDnsViewData) HasServerCloud() bool`

HasServerCloud returns a boolean if a field has been set.

### GetServerComment

`func (o *DataInnerDnsViewData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DataInnerDnsViewData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DataInnerDnsViewData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DataInnerDnsViewData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDnsViewData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDnsViewData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDnsViewData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDnsViewData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDnsViewData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDnsViewData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDnsViewData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDnsViewData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDnsViewData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDnsViewData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDnsViewData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDnsViewData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerVersion

`func (o *DataInnerDnsViewData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DataInnerDnsViewData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DataInnerDnsViewData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DataInnerDnsViewData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetViewAllowQuery

`func (o *DataInnerDnsViewData) GetViewAllowQuery() string`

GetViewAllowQuery returns the ViewAllowQuery field if non-nil, zero value otherwise.

### GetViewAllowQueryOk

`func (o *DataInnerDnsViewData) GetViewAllowQueryOk() (*string, bool)`

GetViewAllowQueryOk returns a tuple with the ViewAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowQuery

`func (o *DataInnerDnsViewData) SetViewAllowQuery(v string)`

SetViewAllowQuery sets ViewAllowQuery field to given value.

### HasViewAllowQuery

`func (o *DataInnerDnsViewData) HasViewAllowQuery() bool`

HasViewAllowQuery returns a boolean if a field has been set.

### GetViewAllowRecursion

`func (o *DataInnerDnsViewData) GetViewAllowRecursion() string`

GetViewAllowRecursion returns the ViewAllowRecursion field if non-nil, zero value otherwise.

### GetViewAllowRecursionOk

`func (o *DataInnerDnsViewData) GetViewAllowRecursionOk() (*string, bool)`

GetViewAllowRecursionOk returns a tuple with the ViewAllowRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowRecursion

`func (o *DataInnerDnsViewData) SetViewAllowRecursion(v string)`

SetViewAllowRecursion sets ViewAllowRecursion field to given value.

### HasViewAllowRecursion

`func (o *DataInnerDnsViewData) HasViewAllowRecursion() bool`

HasViewAllowRecursion returns a boolean if a field has been set.

### GetViewAllowTransfer

`func (o *DataInnerDnsViewData) GetViewAllowTransfer() string`

GetViewAllowTransfer returns the ViewAllowTransfer field if non-nil, zero value otherwise.

### GetViewAllowTransferOk

`func (o *DataInnerDnsViewData) GetViewAllowTransferOk() (*string, bool)`

GetViewAllowTransferOk returns a tuple with the ViewAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowTransfer

`func (o *DataInnerDnsViewData) SetViewAllowTransfer(v string)`

SetViewAllowTransfer sets ViewAllowTransfer field to given value.

### HasViewAllowTransfer

`func (o *DataInnerDnsViewData) HasViewAllowTransfer() bool`

HasViewAllowTransfer returns a boolean if a field has been set.

### GetViewClassName

`func (o *DataInnerDnsViewData) GetViewClassName() string`

GetViewClassName returns the ViewClassName field if non-nil, zero value otherwise.

### GetViewClassNameOk

`func (o *DataInnerDnsViewData) GetViewClassNameOk() (*string, bool)`

GetViewClassNameOk returns a tuple with the ViewClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassName

`func (o *DataInnerDnsViewData) SetViewClassName(v string)`

SetViewClassName sets ViewClassName field to given value.

### HasViewClassName

`func (o *DataInnerDnsViewData) HasViewClassName() bool`

HasViewClassName returns a boolean if a field has been set.

### GetViewClassParameters

`func (o *DataInnerDnsViewData) GetViewClassParameters() []ApiClassParameterOutputEntry`

GetViewClassParameters returns the ViewClassParameters field if non-nil, zero value otherwise.

### GetViewClassParametersOk

`func (o *DataInnerDnsViewData) GetViewClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetViewClassParametersOk returns a tuple with the ViewClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassParameters

`func (o *DataInnerDnsViewData) SetViewClassParameters(v []ApiClassParameterOutputEntry)`

SetViewClassParameters sets ViewClassParameters field to given value.

### HasViewClassParameters

`func (o *DataInnerDnsViewData) HasViewClassParameters() bool`

HasViewClassParameters returns a boolean if a field has been set.

### GetViewId

`func (o *DataInnerDnsViewData) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DataInnerDnsViewData) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DataInnerDnsViewData) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DataInnerDnsViewData) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewKeyName

`func (o *DataInnerDnsViewData) GetViewKeyName() string`

GetViewKeyName returns the ViewKeyName field if non-nil, zero value otherwise.

### GetViewKeyNameOk

`func (o *DataInnerDnsViewData) GetViewKeyNameOk() (*string, bool)`

GetViewKeyNameOk returns a tuple with the ViewKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewKeyName

`func (o *DataInnerDnsViewData) SetViewKeyName(v string)`

SetViewKeyName sets ViewKeyName field to given value.

### HasViewKeyName

`func (o *DataInnerDnsViewData) HasViewKeyName() bool`

HasViewKeyName returns a boolean if a field has been set.

### GetViewMatchClients

`func (o *DataInnerDnsViewData) GetViewMatchClients() string`

GetViewMatchClients returns the ViewMatchClients field if non-nil, zero value otherwise.

### GetViewMatchClientsOk

`func (o *DataInnerDnsViewData) GetViewMatchClientsOk() (*string, bool)`

GetViewMatchClientsOk returns a tuple with the ViewMatchClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatchClients

`func (o *DataInnerDnsViewData) SetViewMatchClients(v string)`

SetViewMatchClients sets ViewMatchClients field to given value.

### HasViewMatchClients

`func (o *DataInnerDnsViewData) HasViewMatchClients() bool`

HasViewMatchClients returns a boolean if a field has been set.

### GetViewMatchTo

`func (o *DataInnerDnsViewData) GetViewMatchTo() string`

GetViewMatchTo returns the ViewMatchTo field if non-nil, zero value otherwise.

### GetViewMatchToOk

`func (o *DataInnerDnsViewData) GetViewMatchToOk() (*string, bool)`

GetViewMatchToOk returns a tuple with the ViewMatchTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatchTo

`func (o *DataInnerDnsViewData) SetViewMatchTo(v string)`

SetViewMatchTo sets ViewMatchTo field to given value.

### HasViewMatchTo

`func (o *DataInnerDnsViewData) HasViewMatchTo() bool`

HasViewMatchTo returns a boolean if a field has been set.

### GetViewName

`func (o *DataInnerDnsViewData) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DataInnerDnsViewData) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DataInnerDnsViewData) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DataInnerDnsViewData) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetViewOrder

`func (o *DataInnerDnsViewData) GetViewOrder() string`

GetViewOrder returns the ViewOrder field if non-nil, zero value otherwise.

### GetViewOrderOk

`func (o *DataInnerDnsViewData) GetViewOrderOk() (*string, bool)`

GetViewOrderOk returns a tuple with the ViewOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewOrder

`func (o *DataInnerDnsViewData) SetViewOrder(v string)`

SetViewOrder sets ViewOrder field to given value.

### HasViewOrder

`func (o *DataInnerDnsViewData) HasViewOrder() bool`

HasViewOrder returns a boolean if a field has been set.

### GetViewRecursion

`func (o *DataInnerDnsViewData) GetViewRecursion() string`

GetViewRecursion returns the ViewRecursion field if non-nil, zero value otherwise.

### GetViewRecursionOk

`func (o *DataInnerDnsViewData) GetViewRecursionOk() (*string, bool)`

GetViewRecursionOk returns a tuple with the ViewRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRecursion

`func (o *DataInnerDnsViewData) SetViewRecursion(v string)`

SetViewRecursion sets ViewRecursion field to given value.

### HasViewRecursion

`func (o *DataInnerDnsViewData) HasViewRecursion() bool`

HasViewRecursion returns a boolean if a field has been set.

### GetServerGssKeytabId

`func (o *DataInnerDnsViewData) GetServerGssKeytabId() string`

GetServerGssKeytabId returns the ServerGssKeytabId field if non-nil, zero value otherwise.

### GetServerGssKeytabIdOk

`func (o *DataInnerDnsViewData) GetServerGssKeytabIdOk() (*string, bool)`

GetServerGssKeytabIdOk returns a tuple with the ServerGssKeytabId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerGssKeytabId

`func (o *DataInnerDnsViewData) SetServerGssKeytabId(v string)`

SetServerGssKeytabId sets ServerGssKeytabId field to given value.

### HasServerGssKeytabId

`func (o *DataInnerDnsViewData) HasServerGssKeytabId() bool`

HasServerGssKeytabId returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DataInnerDnsViewData) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DataInnerDnsViewData) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DataInnerDnsViewData) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DataInnerDnsViewData) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerAddr6

`func (o *DataInnerDnsViewData) GetServerAddr6() string`

GetServerAddr6 returns the ServerAddr6 field if non-nil, zero value otherwise.

### GetServerAddr6Ok

`func (o *DataInnerDnsViewData) GetServerAddr6Ok() (*string, bool)`

GetServerAddr6Ok returns a tuple with the ServerAddr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr6

`func (o *DataInnerDnsViewData) SetServerAddr6(v string)`

SetServerAddr6 sets ServerAddr6 field to given value.

### HasServerAddr6

`func (o *DataInnerDnsViewData) HasServerAddr6() bool`

HasServerAddr6 returns a boolean if a field has been set.

### GetServerAddr

`func (o *DataInnerDnsViewData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DataInnerDnsViewData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DataInnerDnsViewData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DataInnerDnsViewData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetViewMultistatus

`func (o *DataInnerDnsViewData) GetViewMultistatus() string`

GetViewMultistatus returns the ViewMultistatus field if non-nil, zero value otherwise.

### GetViewMultistatusOk

`func (o *DataInnerDnsViewData) GetViewMultistatusOk() (*string, bool)`

GetViewMultistatusOk returns a tuple with the ViewMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMultistatus

`func (o *DataInnerDnsViewData) SetViewMultistatus(v string)`

SetViewMultistatus sets ViewMultistatus field to given value.

### HasViewMultistatus

`func (o *DataInnerDnsViewData) HasViewMultistatus() bool`

HasViewMultistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDnsViewData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDnsViewData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDnsViewData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDnsViewData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDnsViewData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDnsViewData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDnsViewData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDnsViewData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


