# DhcpRangeEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. Use the ID to specify the DHCPv4 server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server. | [optional] 
**RangeEndAddr** | Pointer to **string** | The last IP address of the DHCPv4 range. | [optional] 
**RangeId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 range, a unique numeric key value automatically incremented when you add a DHCPv4 range. Use the ID to specify which DHCPv4 range to edit. | [optional] 
**RangeStartAddr** | Pointer to **string** | The first IP address of the DHCPv4 range. | [optional] 
**ScopeId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 scope, a unique numeric key value automatically incremented when you add a DHCPv4 scope. Use the ID to specify the DHCPv4 scope of your choice. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**RangeAcl** | Pointer to **string** | The list of ACLs associated with the DHCPv4 range, as follows: &lt;b&gt;&amp;lt;ACL_name&amp;gt;;&amp;lt;ACL_name&amp;gt;;&lt;/b&gt;... . | [optional] 
**RangeName** | Pointer to **string** | The name of the DHCPv4 range, each DHCPv4 range must have a unique name. | [optional] 
**ScopeName** | Pointer to **string** | The name of the DHCPv4 scope. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpRangeEditInput

`func NewDhcpRangeEditInput() *DhcpRangeEditInput`

NewDhcpRangeEditInput instantiates a new DhcpRangeEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpRangeEditInputWithDefaults

`func NewDhcpRangeEditInputWithDefaults() *DhcpRangeEditInput`

NewDhcpRangeEditInputWithDefaults instantiates a new DhcpRangeEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DhcpRangeEditInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpRangeEditInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpRangeEditInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpRangeEditInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpRangeEditInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpRangeEditInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpRangeEditInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpRangeEditInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetRangeEndAddr

`func (o *DhcpRangeEditInput) GetRangeEndAddr() string`

GetRangeEndAddr returns the RangeEndAddr field if non-nil, zero value otherwise.

### GetRangeEndAddrOk

`func (o *DhcpRangeEditInput) GetRangeEndAddrOk() (*string, bool)`

GetRangeEndAddrOk returns a tuple with the RangeEndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndAddr

`func (o *DhcpRangeEditInput) SetRangeEndAddr(v string)`

SetRangeEndAddr sets RangeEndAddr field to given value.

### HasRangeEndAddr

`func (o *DhcpRangeEditInput) HasRangeEndAddr() bool`

HasRangeEndAddr returns a boolean if a field has been set.

### GetRangeId

`func (o *DhcpRangeEditInput) GetRangeId() int32`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *DhcpRangeEditInput) GetRangeIdOk() (*int32, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *DhcpRangeEditInput) SetRangeId(v int32)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *DhcpRangeEditInput) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeStartAddr

`func (o *DhcpRangeEditInput) GetRangeStartAddr() string`

GetRangeStartAddr returns the RangeStartAddr field if non-nil, zero value otherwise.

### GetRangeStartAddrOk

`func (o *DhcpRangeEditInput) GetRangeStartAddrOk() (*string, bool)`

GetRangeStartAddrOk returns a tuple with the RangeStartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartAddr

`func (o *DhcpRangeEditInput) SetRangeStartAddr(v string)`

SetRangeStartAddr sets RangeStartAddr field to given value.

### HasRangeStartAddr

`func (o *DhcpRangeEditInput) HasRangeStartAddr() bool`

HasRangeStartAddr returns a boolean if a field has been set.

### GetScopeId

`func (o *DhcpRangeEditInput) GetScopeId() int32`

GetScopeId returns the ScopeId field if non-nil, zero value otherwise.

### GetScopeIdOk

`func (o *DhcpRangeEditInput) GetScopeIdOk() (*int32, bool)`

GetScopeIdOk returns a tuple with the ScopeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeId

`func (o *DhcpRangeEditInput) SetScopeId(v int32)`

SetScopeId sets ScopeId field to given value.

### HasScopeId

`func (o *DhcpRangeEditInput) HasScopeId() bool`

HasScopeId returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DhcpRangeEditInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DhcpRangeEditInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DhcpRangeEditInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DhcpRangeEditInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetRangeAcl

`func (o *DhcpRangeEditInput) GetRangeAcl() string`

GetRangeAcl returns the RangeAcl field if non-nil, zero value otherwise.

### GetRangeAclOk

`func (o *DhcpRangeEditInput) GetRangeAclOk() (*string, bool)`

GetRangeAclOk returns a tuple with the RangeAcl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeAcl

`func (o *DhcpRangeEditInput) SetRangeAcl(v string)`

SetRangeAcl sets RangeAcl field to given value.

### HasRangeAcl

`func (o *DhcpRangeEditInput) HasRangeAcl() bool`

HasRangeAcl returns a boolean if a field has been set.

### GetRangeName

`func (o *DhcpRangeEditInput) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *DhcpRangeEditInput) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *DhcpRangeEditInput) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *DhcpRangeEditInput) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetScopeName

`func (o *DhcpRangeEditInput) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DhcpRangeEditInput) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DhcpRangeEditInput) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DhcpRangeEditInput) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DhcpRangeEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DhcpRangeEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DhcpRangeEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DhcpRangeEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetRangeClassName

`func (o *DhcpRangeEditInput) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *DhcpRangeEditInput) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *DhcpRangeEditInput) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *DhcpRangeEditInput) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *DhcpRangeEditInput) GetRangeClassParameters() []ApiClassParameterInputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *DhcpRangeEditInput) GetRangeClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *DhcpRangeEditInput) SetRangeClassParameters(v []ApiClassParameterInputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *DhcpRangeEditInput) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpRangeEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpRangeEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpRangeEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpRangeEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


