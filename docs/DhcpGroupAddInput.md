# DhcpGroupAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. Use the ID to specify the DHCPv4 server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server. | [optional] 
**GroupName** | Pointer to **string** | The name of the DHCPv4 group, each DHCPv4 group must have a unique name. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**GroupClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**GroupClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpGroupAddInput

`func NewDhcpGroupAddInput() *DhcpGroupAddInput`

NewDhcpGroupAddInput instantiates a new DhcpGroupAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpGroupAddInputWithDefaults

`func NewDhcpGroupAddInputWithDefaults() *DhcpGroupAddInput`

NewDhcpGroupAddInputWithDefaults instantiates a new DhcpGroupAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DhcpGroupAddInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpGroupAddInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpGroupAddInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpGroupAddInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpGroupAddInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpGroupAddInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpGroupAddInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpGroupAddInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetGroupName

`func (o *DhcpGroupAddInput) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DhcpGroupAddInput) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DhcpGroupAddInput) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DhcpGroupAddInput) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DhcpGroupAddInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DhcpGroupAddInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DhcpGroupAddInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DhcpGroupAddInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DhcpGroupAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DhcpGroupAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DhcpGroupAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DhcpGroupAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetGroupClassName

`func (o *DhcpGroupAddInput) GetGroupClassName() string`

GetGroupClassName returns the GroupClassName field if non-nil, zero value otherwise.

### GetGroupClassNameOk

`func (o *DhcpGroupAddInput) GetGroupClassNameOk() (*string, bool)`

GetGroupClassNameOk returns a tuple with the GroupClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassName

`func (o *DhcpGroupAddInput) SetGroupClassName(v string)`

SetGroupClassName sets GroupClassName field to given value.

### HasGroupClassName

`func (o *DhcpGroupAddInput) HasGroupClassName() bool`

HasGroupClassName returns a boolean if a field has been set.

### GetGroupClassParameters

`func (o *DhcpGroupAddInput) GetGroupClassParameters() []ApiClassParameterInputEntry`

GetGroupClassParameters returns the GroupClassParameters field if non-nil, zero value otherwise.

### GetGroupClassParametersOk

`func (o *DhcpGroupAddInput) GetGroupClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetGroupClassParametersOk returns a tuple with the GroupClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupClassParameters

`func (o *DhcpGroupAddInput) SetGroupClassParameters(v []ApiClassParameterInputEntry)`

SetGroupClassParameters sets GroupClassParameters field to given value.

### HasGroupClassParameters

`func (o *DhcpGroupAddInput) HasGroupClassParameters() bool`

HasGroupClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpGroupAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpGroupAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpGroupAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpGroupAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


