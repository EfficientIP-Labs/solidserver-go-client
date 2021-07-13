# DhcpGroupEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. Use the ID to specify the DHCPv4 server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server. | [optional] 
**GroupId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 group, a unique numeric key value automatically incremented when you add a DHCPv4 group. Use the ID to specify the DHCPv4 group of your choice. | [optional] 
**GroupName** | Pointer to **string** | The name of the DHCPv4 group, each DHCPv4 group must have a unique name. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**GroupClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**GroupClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpGroupEditInput

`func NewDhcpGroupEditInput() *DhcpGroupEditInput`

NewDhcpGroupEditInput instantiates a new DhcpGroupEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpGroupEditInputWithDefaults

`func NewDhcpGroupEditInputWithDefaults() *DhcpGroupEditInput`

NewDhcpGroupEditInputWithDefaults instantiates a new DhcpGroupEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DhcpGroupEditInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpGroupEditInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpGroupEditInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpGroupEditInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpGroupEditInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpGroupEditInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpGroupEditInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpGroupEditInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetGroupId

`func (o *DhcpGroupEditInput) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DhcpGroupEditInput) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DhcpGroupEditInput) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DhcpGroupEditInput) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *DhcpGroupEditInput) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DhcpGroupEditInput) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DhcpGroupEditInput) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DhcpGroupEditInput) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DhcpGroupEditInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DhcpGroupEditInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DhcpGroupEditInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DhcpGroupEditInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DhcpGroupEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DhcpGroupEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DhcpGroupEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DhcpGroupEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetGroupClassName

`func (o *DhcpGroupEditInput) GetGroupClassName() string`

GetGroupClassName returns the GroupClassName field if non-nil, zero value otherwise.

### GetGroupClassNameOk

`func (o *DhcpGroupEditInput) GetGroupClassNameOk() (*string, bool)`

GetGroupClassNameOk returns a tuple with the GroupClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassName

`func (o *DhcpGroupEditInput) SetGroupClassName(v string)`

SetGroupClassName sets GroupClassName field to given value.

### HasGroupClassName

`func (o *DhcpGroupEditInput) HasGroupClassName() bool`

HasGroupClassName returns a boolean if a field has been set.

### GetGroupClassParameters

`func (o *DhcpGroupEditInput) GetGroupClassParameters() []ApiClassParameterInputEntry`

GetGroupClassParameters returns the GroupClassParameters field if non-nil, zero value otherwise.

### GetGroupClassParametersOk

`func (o *DhcpGroupEditInput) GetGroupClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetGroupClassParametersOk returns a tuple with the GroupClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassParameters

`func (o *DhcpGroupEditInput) SetGroupClassParameters(v []ApiClassParameterInputEntry)`

SetGroupClassParameters sets GroupClassParameters field to given value.

### HasGroupClassParameters

`func (o *DhcpGroupEditInput) HasGroupClassParameters() bool`

HasGroupClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpGroupEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpGroupEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpGroupEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpGroupEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


