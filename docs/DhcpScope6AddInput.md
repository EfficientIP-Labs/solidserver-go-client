# DhcpScope6AddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 server, a unique numeric key value automatically incremented when you add a DHCPv6 server. Use the ID to specify the DHCPv6 server of your choice. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server. | [optional] 
**Scope6EndAddr** | Pointer to **string** | The last IP address of the DHCPv6 scope. | [optional] 
**Scope6Prefix** | Pointer to **string** | The prefix of the DHCPv6 scope, an integer that defines the number of address the scope contains. | [optional] 
**Scope6StartAddr** | Pointer to **string** | The first IP address of the DHCPv6 scope. | [optional] 
**Server6Hostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**Failover6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 failover channel, a unique numeric key value automatically incremented when you add a DHCPv6 failover channel. Use the ID to specify the DHCPv6 failover channel of your choice. | [optional] 
**Failover6Name** | Pointer to **string** | The name of the DHCPv6 failover channel. | [optional] 
**Scope6Name** | Pointer to **string** | The name of the DHCPv6 scope, each DHCPv6 scope must have a unique name. | [optional] 
**Scope6SpaceId** | Pointer to **int32** | The database identifier (ID) of an existing space you want to associate with the DHCPv6 scope. | [optional] 
**Scope6SpaceName** | Pointer to **string** | The name of an existing space you want to associate with the DHCPv6 scope. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**Scope6ClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**Scope6ClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpScope6AddInput

`func NewDhcpScope6AddInput() *DhcpScope6AddInput`

NewDhcpScope6AddInput instantiates a new DhcpScope6AddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpScope6AddInputWithDefaults

`func NewDhcpScope6AddInputWithDefaults() *DhcpScope6AddInput`

NewDhcpScope6AddInputWithDefaults instantiates a new DhcpScope6AddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer6Id

`func (o *DhcpScope6AddInput) GetServer6Id() int32`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DhcpScope6AddInput) GetServer6IdOk() (*int32, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DhcpScope6AddInput) SetServer6Id(v int32)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DhcpScope6AddInput) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DhcpScope6AddInput) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DhcpScope6AddInput) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DhcpScope6AddInput) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DhcpScope6AddInput) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetScope6EndAddr

`func (o *DhcpScope6AddInput) GetScope6EndAddr() string`

GetScope6EndAddr returns the Scope6EndAddr field if non-nil, zero value otherwise.

### GetScope6EndAddrOk

`func (o *DhcpScope6AddInput) GetScope6EndAddrOk() (*string, bool)`

GetScope6EndAddrOk returns a tuple with the Scope6EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6EndAddr

`func (o *DhcpScope6AddInput) SetScope6EndAddr(v string)`

SetScope6EndAddr sets Scope6EndAddr field to given value.

### HasScope6EndAddr

`func (o *DhcpScope6AddInput) HasScope6EndAddr() bool`

HasScope6EndAddr returns a boolean if a field has been set.

### GetScope6Prefix

`func (o *DhcpScope6AddInput) GetScope6Prefix() string`

GetScope6Prefix returns the Scope6Prefix field if non-nil, zero value otherwise.

### GetScope6PrefixOk

`func (o *DhcpScope6AddInput) GetScope6PrefixOk() (*string, bool)`

GetScope6PrefixOk returns a tuple with the Scope6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Prefix

`func (o *DhcpScope6AddInput) SetScope6Prefix(v string)`

SetScope6Prefix sets Scope6Prefix field to given value.

### HasScope6Prefix

`func (o *DhcpScope6AddInput) HasScope6Prefix() bool`

HasScope6Prefix returns a boolean if a field has been set.

### GetScope6StartAddr

`func (o *DhcpScope6AddInput) GetScope6StartAddr() string`

GetScope6StartAddr returns the Scope6StartAddr field if non-nil, zero value otherwise.

### GetScope6StartAddrOk

`func (o *DhcpScope6AddInput) GetScope6StartAddrOk() (*string, bool)`

GetScope6StartAddrOk returns a tuple with the Scope6StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6StartAddr

`func (o *DhcpScope6AddInput) SetScope6StartAddr(v string)`

SetScope6StartAddr sets Scope6StartAddr field to given value.

### HasScope6StartAddr

`func (o *DhcpScope6AddInput) HasScope6StartAddr() bool`

HasScope6StartAddr returns a boolean if a field has been set.

### GetServer6Hostaddr

`func (o *DhcpScope6AddInput) GetServer6Hostaddr() string`

GetServer6Hostaddr returns the Server6Hostaddr field if non-nil, zero value otherwise.

### GetServer6HostaddrOk

`func (o *DhcpScope6AddInput) GetServer6HostaddrOk() (*string, bool)`

GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Hostaddr

`func (o *DhcpScope6AddInput) SetServer6Hostaddr(v string)`

SetServer6Hostaddr sets Server6Hostaddr field to given value.

### HasServer6Hostaddr

`func (o *DhcpScope6AddInput) HasServer6Hostaddr() bool`

HasServer6Hostaddr returns a boolean if a field has been set.

### GetFailover6Id

`func (o *DhcpScope6AddInput) GetFailover6Id() int32`

GetFailover6Id returns the Failover6Id field if non-nil, zero value otherwise.

### GetFailover6IdOk

`func (o *DhcpScope6AddInput) GetFailover6IdOk() (*int32, bool)`

GetFailover6IdOk returns a tuple with the Failover6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover6Id

`func (o *DhcpScope6AddInput) SetFailover6Id(v int32)`

SetFailover6Id sets Failover6Id field to given value.

### HasFailover6Id

`func (o *DhcpScope6AddInput) HasFailover6Id() bool`

HasFailover6Id returns a boolean if a field has been set.

### GetFailover6Name

`func (o *DhcpScope6AddInput) GetFailover6Name() string`

GetFailover6Name returns the Failover6Name field if non-nil, zero value otherwise.

### GetFailover6NameOk

`func (o *DhcpScope6AddInput) GetFailover6NameOk() (*string, bool)`

GetFailover6NameOk returns a tuple with the Failover6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailover6Name

`func (o *DhcpScope6AddInput) SetFailover6Name(v string)`

SetFailover6Name sets Failover6Name field to given value.

### HasFailover6Name

`func (o *DhcpScope6AddInput) HasFailover6Name() bool`

HasFailover6Name returns a boolean if a field has been set.

### GetScope6Name

`func (o *DhcpScope6AddInput) GetScope6Name() string`

GetScope6Name returns the Scope6Name field if non-nil, zero value otherwise.

### GetScope6NameOk

`func (o *DhcpScope6AddInput) GetScope6NameOk() (*string, bool)`

GetScope6NameOk returns a tuple with the Scope6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Name

`func (o *DhcpScope6AddInput) SetScope6Name(v string)`

SetScope6Name sets Scope6Name field to given value.

### HasScope6Name

`func (o *DhcpScope6AddInput) HasScope6Name() bool`

HasScope6Name returns a boolean if a field has been set.

### GetScope6SpaceId

`func (o *DhcpScope6AddInput) GetScope6SpaceId() int32`

GetScope6SpaceId returns the Scope6SpaceId field if non-nil, zero value otherwise.

### GetScope6SpaceIdOk

`func (o *DhcpScope6AddInput) GetScope6SpaceIdOk() (*int32, bool)`

GetScope6SpaceIdOk returns a tuple with the Scope6SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6SpaceId

`func (o *DhcpScope6AddInput) SetScope6SpaceId(v int32)`

SetScope6SpaceId sets Scope6SpaceId field to given value.

### HasScope6SpaceId

`func (o *DhcpScope6AddInput) HasScope6SpaceId() bool`

HasScope6SpaceId returns a boolean if a field has been set.

### GetScope6SpaceName

`func (o *DhcpScope6AddInput) GetScope6SpaceName() string`

GetScope6SpaceName returns the Scope6SpaceName field if non-nil, zero value otherwise.

### GetScope6SpaceNameOk

`func (o *DhcpScope6AddInput) GetScope6SpaceNameOk() (*string, bool)`

GetScope6SpaceNameOk returns a tuple with the Scope6SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6SpaceName

`func (o *DhcpScope6AddInput) SetScope6SpaceName(v string)`

SetScope6SpaceName sets Scope6SpaceName field to given value.

### HasScope6SpaceName

`func (o *DhcpScope6AddInput) HasScope6SpaceName() bool`

HasScope6SpaceName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DhcpScope6AddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DhcpScope6AddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DhcpScope6AddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DhcpScope6AddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetScope6ClassName

`func (o *DhcpScope6AddInput) GetScope6ClassName() string`

GetScope6ClassName returns the Scope6ClassName field if non-nil, zero value otherwise.

### GetScope6ClassNameOk

`func (o *DhcpScope6AddInput) GetScope6ClassNameOk() (*string, bool)`

GetScope6ClassNameOk returns a tuple with the Scope6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassName

`func (o *DhcpScope6AddInput) SetScope6ClassName(v string)`

SetScope6ClassName sets Scope6ClassName field to given value.

### HasScope6ClassName

`func (o *DhcpScope6AddInput) HasScope6ClassName() bool`

HasScope6ClassName returns a boolean if a field has been set.

### GetScope6ClassParameters

`func (o *DhcpScope6AddInput) GetScope6ClassParameters() []ApiClassParameterInputEntry`

GetScope6ClassParameters returns the Scope6ClassParameters field if non-nil, zero value otherwise.

### GetScope6ClassParametersOk

`func (o *DhcpScope6AddInput) GetScope6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetScope6ClassParametersOk returns a tuple with the Scope6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6ClassParameters

`func (o *DhcpScope6AddInput) SetScope6ClassParameters(v []ApiClassParameterInputEntry)`

SetScope6ClassParameters sets Scope6ClassParameters field to given value.

### HasScope6ClassParameters

`func (o *DhcpScope6AddInput) HasScope6ClassParameters() bool`

HasScope6ClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpScope6AddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpScope6AddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpScope6AddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpScope6AddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


