# DnsViewEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server. | [optional] 
**ViewId** | Pointer to **int32** | The database identifier (ID) of the DNS view, a unique numeric key value automatically incremented when you add a DNS view. Use the ID to specify the DNS view of your choice. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view, each DNS view must have a unique name. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DNS server. | [optional] 
**ViewAllowQuery** | Pointer to **string** | The ACL values associated with the allow-query configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewAllowRecursion** | Pointer to **string** | The ACL values associated with the allow-recursion configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewAllowTransfer** | Pointer to **string** | The ACL values associated with the allow-transfer configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;... &lt;/b&gt;. Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewMatchClients** | Pointer to **string** | The ACL values associated with the match clients configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;...&lt;/b&gt; . Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewMatchTo** | Pointer to **string** | The ACL values associated with the match destination configuration of the DNS view, as follows: &lt;b&gt;&amp;lt;value1&amp;gt;;&amp;lt;value2&amp;gt;;...&lt;/b&gt; . Values may include IP and network addresses, the name of TSIG keys and ACLs, preceded by &lt;b&gt;!&lt;/b&gt; if the access is denied. | [optional] 
**ViewOrder** | Pointer to **int32** | The level of the DNS view, where &lt;b&gt;0&lt;/b&gt; represents the highest level in the views hierarchy. The parameters &lt;b&gt;dnsview_match_client&lt;/b&gt; and &lt;b&gt;view_match_to&lt;/b&gt; of each view in a server are reviewed following this order. | [optional] 
**ViewRecursion** | Pointer to **string** | The recursion status of the DNS view:&lt;table&gt;&lt;caption&gt;view_recursion possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;no&lt;/td&gt;&lt;td &gt;The view only provides iterative query behavior - normally resulting in a referral. If the answer to the query already exists in the cache it will be returned whatever the value of this statement.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;yes&lt;/td&gt;&lt;td &gt;The view always provides recursive query behavior if requested by the client.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**ViewClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**ViewClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDnsViewEditInput

`func NewDnsViewEditInput() *DnsViewEditInput`

NewDnsViewEditInput instantiates a new DnsViewEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsViewEditInputWithDefaults

`func NewDnsViewEditInputWithDefaults() *DnsViewEditInput`

NewDnsViewEditInputWithDefaults instantiates a new DnsViewEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DnsViewEditInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsViewEditInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsViewEditInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsViewEditInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsViewEditInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsViewEditInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsViewEditInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsViewEditInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetViewId

`func (o *DnsViewEditInput) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DnsViewEditInput) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DnsViewEditInput) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DnsViewEditInput) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewName

`func (o *DnsViewEditInput) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DnsViewEditInput) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DnsViewEditInput) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DnsViewEditInput) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DnsViewEditInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DnsViewEditInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DnsViewEditInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DnsViewEditInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetViewAllowQuery

`func (o *DnsViewEditInput) GetViewAllowQuery() string`

GetViewAllowQuery returns the ViewAllowQuery field if non-nil, zero value otherwise.

### GetViewAllowQueryOk

`func (o *DnsViewEditInput) GetViewAllowQueryOk() (*string, bool)`

GetViewAllowQueryOk returns a tuple with the ViewAllowQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowQuery

`func (o *DnsViewEditInput) SetViewAllowQuery(v string)`

SetViewAllowQuery sets ViewAllowQuery field to given value.

### HasViewAllowQuery

`func (o *DnsViewEditInput) HasViewAllowQuery() bool`

HasViewAllowQuery returns a boolean if a field has been set.

### GetViewAllowRecursion

`func (o *DnsViewEditInput) GetViewAllowRecursion() string`

GetViewAllowRecursion returns the ViewAllowRecursion field if non-nil, zero value otherwise.

### GetViewAllowRecursionOk

`func (o *DnsViewEditInput) GetViewAllowRecursionOk() (*string, bool)`

GetViewAllowRecursionOk returns a tuple with the ViewAllowRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowRecursion

`func (o *DnsViewEditInput) SetViewAllowRecursion(v string)`

SetViewAllowRecursion sets ViewAllowRecursion field to given value.

### HasViewAllowRecursion

`func (o *DnsViewEditInput) HasViewAllowRecursion() bool`

HasViewAllowRecursion returns a boolean if a field has been set.

### GetViewAllowTransfer

`func (o *DnsViewEditInput) GetViewAllowTransfer() string`

GetViewAllowTransfer returns the ViewAllowTransfer field if non-nil, zero value otherwise.

### GetViewAllowTransferOk

`func (o *DnsViewEditInput) GetViewAllowTransferOk() (*string, bool)`

GetViewAllowTransferOk returns a tuple with the ViewAllowTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewAllowTransfer

`func (o *DnsViewEditInput) SetViewAllowTransfer(v string)`

SetViewAllowTransfer sets ViewAllowTransfer field to given value.

### HasViewAllowTransfer

`func (o *DnsViewEditInput) HasViewAllowTransfer() bool`

HasViewAllowTransfer returns a boolean if a field has been set.

### GetViewMatchClients

`func (o *DnsViewEditInput) GetViewMatchClients() string`

GetViewMatchClients returns the ViewMatchClients field if non-nil, zero value otherwise.

### GetViewMatchClientsOk

`func (o *DnsViewEditInput) GetViewMatchClientsOk() (*string, bool)`

GetViewMatchClientsOk returns a tuple with the ViewMatchClients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatchClients

`func (o *DnsViewEditInput) SetViewMatchClients(v string)`

SetViewMatchClients sets ViewMatchClients field to given value.

### HasViewMatchClients

`func (o *DnsViewEditInput) HasViewMatchClients() bool`

HasViewMatchClients returns a boolean if a field has been set.

### GetViewMatchTo

`func (o *DnsViewEditInput) GetViewMatchTo() string`

GetViewMatchTo returns the ViewMatchTo field if non-nil, zero value otherwise.

### GetViewMatchToOk

`func (o *DnsViewEditInput) GetViewMatchToOk() (*string, bool)`

GetViewMatchToOk returns a tuple with the ViewMatchTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMatchTo

`func (o *DnsViewEditInput) SetViewMatchTo(v string)`

SetViewMatchTo sets ViewMatchTo field to given value.

### HasViewMatchTo

`func (o *DnsViewEditInput) HasViewMatchTo() bool`

HasViewMatchTo returns a boolean if a field has been set.

### GetViewOrder

`func (o *DnsViewEditInput) GetViewOrder() int32`

GetViewOrder returns the ViewOrder field if non-nil, zero value otherwise.

### GetViewOrderOk

`func (o *DnsViewEditInput) GetViewOrderOk() (*int32, bool)`

GetViewOrderOk returns a tuple with the ViewOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewOrder

`func (o *DnsViewEditInput) SetViewOrder(v int32)`

SetViewOrder sets ViewOrder field to given value.

### HasViewOrder

`func (o *DnsViewEditInput) HasViewOrder() bool`

HasViewOrder returns a boolean if a field has been set.

### GetViewRecursion

`func (o *DnsViewEditInput) GetViewRecursion() string`

GetViewRecursion returns the ViewRecursion field if non-nil, zero value otherwise.

### GetViewRecursionOk

`func (o *DnsViewEditInput) GetViewRecursionOk() (*string, bool)`

GetViewRecursionOk returns a tuple with the ViewRecursion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewRecursion

`func (o *DnsViewEditInput) SetViewRecursion(v string)`

SetViewRecursion sets ViewRecursion field to given value.

### HasViewRecursion

`func (o *DnsViewEditInput) HasViewRecursion() bool`

HasViewRecursion returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DnsViewEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DnsViewEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DnsViewEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DnsViewEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetViewClassName

`func (o *DnsViewEditInput) GetViewClassName() string`

GetViewClassName returns the ViewClassName field if non-nil, zero value otherwise.

### GetViewClassNameOk

`func (o *DnsViewEditInput) GetViewClassNameOk() (*string, bool)`

GetViewClassNameOk returns a tuple with the ViewClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassName

`func (o *DnsViewEditInput) SetViewClassName(v string)`

SetViewClassName sets ViewClassName field to given value.

### HasViewClassName

`func (o *DnsViewEditInput) HasViewClassName() bool`

HasViewClassName returns a boolean if a field has been set.

### GetViewClassParameters

`func (o *DnsViewEditInput) GetViewClassParameters() []ApiClassParameterInputEntry`

GetViewClassParameters returns the ViewClassParameters field if non-nil, zero value otherwise.

### GetViewClassParametersOk

`func (o *DnsViewEditInput) GetViewClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetViewClassParametersOk returns a tuple with the ViewClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewClassParameters

`func (o *DnsViewEditInput) SetViewClassParameters(v []ApiClassParameterInputEntry)`

SetViewClassParameters sets ViewClassParameters field to given value.

### HasViewClassParameters

`func (o *DnsViewEditInput) HasViewClassParameters() bool`

HasViewClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DnsViewEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DnsViewEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DnsViewEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DnsViewEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


