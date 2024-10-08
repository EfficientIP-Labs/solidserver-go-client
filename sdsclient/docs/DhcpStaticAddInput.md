# DhcpStaticAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. Use the ID to specify the DHCPv4 server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server. | [optional] 
**StaticAddr** | Pointer to **string** | The IP address associated with the DHCPv4 static. | [optional] 
**StaticIdentifier** | Pointer to **string** | The host identifier you want to associate with the IPv4 static. An option and value to look for to identify clients and assign them the static, specified as follows: &lt;b&gt;option &amp;lt;option-name&amp;gt; expected value&lt;/b&gt;. | [optional] 
**StaticMacAddr** | Pointer to **string** | The MAC address you want to associate with the IPv4 static, it must include the MAC address type. The address has 7 sections, &lt;b&gt;00:11:22:33:44:55:66&lt;/b&gt; , where &lt;b&gt;00&lt;/b&gt; indicates the type. For Ethernet, type in &lt;b&gt;01&lt;/b&gt;. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**GroupId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 group, a unique numeric key value automatically incremented when you add a DHCPv4 group. Use the ID to specify the DHCPv4 group of your choice. | [optional] 
**GroupName** | Pointer to **string** | The name of the DHCPv4 group. | [optional] 
**StaticName** | Pointer to **string** | The name of the DHCPv4 static, each DHCPv4 static must have a unique name. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**StaticClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**StaticClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpStaticAddInput

`func NewDhcpStaticAddInput() *DhcpStaticAddInput`

NewDhcpStaticAddInput instantiates a new DhcpStaticAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpStaticAddInputWithDefaults

`func NewDhcpStaticAddInputWithDefaults() *DhcpStaticAddInput`

NewDhcpStaticAddInputWithDefaults instantiates a new DhcpStaticAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DhcpStaticAddInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpStaticAddInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpStaticAddInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpStaticAddInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpStaticAddInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpStaticAddInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpStaticAddInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpStaticAddInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetStaticAddr

`func (o *DhcpStaticAddInput) GetStaticAddr() string`

GetStaticAddr returns the StaticAddr field if non-nil, zero value otherwise.

### GetStaticAddrOk

`func (o *DhcpStaticAddInput) GetStaticAddrOk() (*string, bool)`

GetStaticAddrOk returns a tuple with the StaticAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticAddr

`func (o *DhcpStaticAddInput) SetStaticAddr(v string)`

SetStaticAddr sets StaticAddr field to given value.

### HasStaticAddr

`func (o *DhcpStaticAddInput) HasStaticAddr() bool`

HasStaticAddr returns a boolean if a field has been set.

### GetStaticIdentifier

`func (o *DhcpStaticAddInput) GetStaticIdentifier() string`

GetStaticIdentifier returns the StaticIdentifier field if non-nil, zero value otherwise.

### GetStaticIdentifierOk

`func (o *DhcpStaticAddInput) GetStaticIdentifierOk() (*string, bool)`

GetStaticIdentifierOk returns a tuple with the StaticIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticIdentifier

`func (o *DhcpStaticAddInput) SetStaticIdentifier(v string)`

SetStaticIdentifier sets StaticIdentifier field to given value.

### HasStaticIdentifier

`func (o *DhcpStaticAddInput) HasStaticIdentifier() bool`

HasStaticIdentifier returns a boolean if a field has been set.

### GetStaticMacAddr

`func (o *DhcpStaticAddInput) GetStaticMacAddr() string`

GetStaticMacAddr returns the StaticMacAddr field if non-nil, zero value otherwise.

### GetStaticMacAddrOk

`func (o *DhcpStaticAddInput) GetStaticMacAddrOk() (*string, bool)`

GetStaticMacAddrOk returns a tuple with the StaticMacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticMacAddr

`func (o *DhcpStaticAddInput) SetStaticMacAddr(v string)`

SetStaticMacAddr sets StaticMacAddr field to given value.

### HasStaticMacAddr

`func (o *DhcpStaticAddInput) HasStaticMacAddr() bool`

HasStaticMacAddr returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DhcpStaticAddInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DhcpStaticAddInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DhcpStaticAddInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DhcpStaticAddInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetGroupId

`func (o *DhcpStaticAddInput) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DhcpStaticAddInput) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DhcpStaticAddInput) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *DhcpStaticAddInput) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupName

`func (o *DhcpStaticAddInput) GetGroupName() string`

GetGroupName returns the GroupName field if non-nil, zero value otherwise.

### GetGroupNameOk

`func (o *DhcpStaticAddInput) GetGroupNameOk() (*string, bool)`

GetGroupNameOk returns a tuple with the GroupName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupName

`func (o *DhcpStaticAddInput) SetGroupName(v string)`

SetGroupName sets GroupName field to given value.

### HasGroupName

`func (o *DhcpStaticAddInput) HasGroupName() bool`

HasGroupName returns a boolean if a field has been set.

### GetStaticName

`func (o *DhcpStaticAddInput) GetStaticName() string`

GetStaticName returns the StaticName field if non-nil, zero value otherwise.

### GetStaticNameOk

`func (o *DhcpStaticAddInput) GetStaticNameOk() (*string, bool)`

GetStaticNameOk returns a tuple with the StaticName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticName

`func (o *DhcpStaticAddInput) SetStaticName(v string)`

SetStaticName sets StaticName field to given value.

### HasStaticName

`func (o *DhcpStaticAddInput) HasStaticName() bool`

HasStaticName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DhcpStaticAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DhcpStaticAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DhcpStaticAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DhcpStaticAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetStaticClassName

`func (o *DhcpStaticAddInput) GetStaticClassName() string`

GetStaticClassName returns the StaticClassName field if non-nil, zero value otherwise.

### GetStaticClassNameOk

`func (o *DhcpStaticAddInput) GetStaticClassNameOk() (*string, bool)`

GetStaticClassNameOk returns a tuple with the StaticClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticClassName

`func (o *DhcpStaticAddInput) SetStaticClassName(v string)`

SetStaticClassName sets StaticClassName field to given value.

### HasStaticClassName

`func (o *DhcpStaticAddInput) HasStaticClassName() bool`

HasStaticClassName returns a boolean if a field has been set.

### GetStaticClassParameters

`func (o *DhcpStaticAddInput) GetStaticClassParameters() []ApiClassParameterInputEntry`

GetStaticClassParameters returns the StaticClassParameters field if non-nil, zero value otherwise.

### GetStaticClassParametersOk

`func (o *DhcpStaticAddInput) GetStaticClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetStaticClassParametersOk returns a tuple with the StaticClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticClassParameters

`func (o *DhcpStaticAddInput) SetStaticClassParameters(v []ApiClassParameterInputEntry)`

SetStaticClassParameters sets StaticClassParameters field to given value.

### HasStaticClassParameters

`func (o *DhcpStaticAddInput) HasStaticClassParameters() bool`

HasStaticClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpStaticAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpStaticAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpStaticAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpStaticAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


