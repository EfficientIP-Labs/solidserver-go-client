# DhcpRange6EditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 server, a unique numeric key value automatically incremented when you add a DHCPv6 server. Use the ID to specify the DHCPv6 server of your choice. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server. | [optional] 
**Range6EndAddr** | Pointer to **string** | The last IP address of the DHCPv6 range. | [optional] 
**Range6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 range, a unique numeric key value automatically incremented when you add a DHCPv6 range. Use the ID to specify the DHCPv6 range of your choice. | [optional] 
**Range6StartAddr** | Pointer to **string** | The first IP address of the DHCPv6 range. | [optional] 
**Scope6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 scope, a unique numeric key value automatically incremented when you add a DHCPv6 scope. Use the ID to specify the DHCPv6 scope of your choice. | [optional] 
**Server6Hostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**Range6Acl** | Pointer to **string** | The list of ACLs associated with the DHCPv6 range, as follows: &lt;b&gt;&amp;lt;ACL_name&amp;gt;;&amp;lt;ACL_name&amp;gt;;...&lt;/b&gt; . | [optional] 
**Scope6Name** | Pointer to **string** | The name of the DHCPv6 scope. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**Range6ClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**Range6ClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpRange6EditInput

`func NewDhcpRange6EditInput() *DhcpRange6EditInput`

NewDhcpRange6EditInput instantiates a new DhcpRange6EditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpRange6EditInputWithDefaults

`func NewDhcpRange6EditInputWithDefaults() *DhcpRange6EditInput`

NewDhcpRange6EditInputWithDefaults instantiates a new DhcpRange6EditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer6Id

`func (o *DhcpRange6EditInput) GetServer6Id() int32`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DhcpRange6EditInput) GetServer6IdOk() (*int32, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DhcpRange6EditInput) SetServer6Id(v int32)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DhcpRange6EditInput) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DhcpRange6EditInput) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DhcpRange6EditInput) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DhcpRange6EditInput) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DhcpRange6EditInput) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetRange6EndAddr

`func (o *DhcpRange6EditInput) GetRange6EndAddr() string`

GetRange6EndAddr returns the Range6EndAddr field if non-nil, zero value otherwise.

### GetRange6EndAddrOk

`func (o *DhcpRange6EditInput) GetRange6EndAddrOk() (*string, bool)`

GetRange6EndAddrOk returns a tuple with the Range6EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6EndAddr

`func (o *DhcpRange6EditInput) SetRange6EndAddr(v string)`

SetRange6EndAddr sets Range6EndAddr field to given value.

### HasRange6EndAddr

`func (o *DhcpRange6EditInput) HasRange6EndAddr() bool`

HasRange6EndAddr returns a boolean if a field has been set.

### GetRange6Id

`func (o *DhcpRange6EditInput) GetRange6Id() int32`

GetRange6Id returns the Range6Id field if non-nil, zero value otherwise.

### GetRange6IdOk

`func (o *DhcpRange6EditInput) GetRange6IdOk() (*int32, bool)`

GetRange6IdOk returns a tuple with the Range6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Id

`func (o *DhcpRange6EditInput) SetRange6Id(v int32)`

SetRange6Id sets Range6Id field to given value.

### HasRange6Id

`func (o *DhcpRange6EditInput) HasRange6Id() bool`

HasRange6Id returns a boolean if a field has been set.

### GetRange6StartAddr

`func (o *DhcpRange6EditInput) GetRange6StartAddr() string`

GetRange6StartAddr returns the Range6StartAddr field if non-nil, zero value otherwise.

### GetRange6StartAddrOk

`func (o *DhcpRange6EditInput) GetRange6StartAddrOk() (*string, bool)`

GetRange6StartAddrOk returns a tuple with the Range6StartAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6StartAddr

`func (o *DhcpRange6EditInput) SetRange6StartAddr(v string)`

SetRange6StartAddr sets Range6StartAddr field to given value.

### HasRange6StartAddr

`func (o *DhcpRange6EditInput) HasRange6StartAddr() bool`

HasRange6StartAddr returns a boolean if a field has been set.

### GetScope6Id

`func (o *DhcpRange6EditInput) GetScope6Id() int32`

GetScope6Id returns the Scope6Id field if non-nil, zero value otherwise.

### GetScope6IdOk

`func (o *DhcpRange6EditInput) GetScope6IdOk() (*int32, bool)`

GetScope6IdOk returns a tuple with the Scope6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Id

`func (o *DhcpRange6EditInput) SetScope6Id(v int32)`

SetScope6Id sets Scope6Id field to given value.

### HasScope6Id

`func (o *DhcpRange6EditInput) HasScope6Id() bool`

HasScope6Id returns a boolean if a field has been set.

### GetServer6Hostaddr

`func (o *DhcpRange6EditInput) GetServer6Hostaddr() string`

GetServer6Hostaddr returns the Server6Hostaddr field if non-nil, zero value otherwise.

### GetServer6HostaddrOk

`func (o *DhcpRange6EditInput) GetServer6HostaddrOk() (*string, bool)`

GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Hostaddr

`func (o *DhcpRange6EditInput) SetServer6Hostaddr(v string)`

SetServer6Hostaddr sets Server6Hostaddr field to given value.

### HasServer6Hostaddr

`func (o *DhcpRange6EditInput) HasServer6Hostaddr() bool`

HasServer6Hostaddr returns a boolean if a field has been set.

### GetRange6Acl

`func (o *DhcpRange6EditInput) GetRange6Acl() string`

GetRange6Acl returns the Range6Acl field if non-nil, zero value otherwise.

### GetRange6AclOk

`func (o *DhcpRange6EditInput) GetRange6AclOk() (*string, bool)`

GetRange6AclOk returns a tuple with the Range6Acl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6Acl

`func (o *DhcpRange6EditInput) SetRange6Acl(v string)`

SetRange6Acl sets Range6Acl field to given value.

### HasRange6Acl

`func (o *DhcpRange6EditInput) HasRange6Acl() bool`

HasRange6Acl returns a boolean if a field has been set.

### GetScope6Name

`func (o *DhcpRange6EditInput) GetScope6Name() string`

GetScope6Name returns the Scope6Name field if non-nil, zero value otherwise.

### GetScope6NameOk

`func (o *DhcpRange6EditInput) GetScope6NameOk() (*string, bool)`

GetScope6NameOk returns a tuple with the Scope6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope6Name

`func (o *DhcpRange6EditInput) SetScope6Name(v string)`

SetScope6Name sets Scope6Name field to given value.

### HasScope6Name

`func (o *DhcpRange6EditInput) HasScope6Name() bool`

HasScope6Name returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DhcpRange6EditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DhcpRange6EditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DhcpRange6EditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DhcpRange6EditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetRange6ClassName

`func (o *DhcpRange6EditInput) GetRange6ClassName() string`

GetRange6ClassName returns the Range6ClassName field if non-nil, zero value otherwise.

### GetRange6ClassNameOk

`func (o *DhcpRange6EditInput) GetRange6ClassNameOk() (*string, bool)`

GetRange6ClassNameOk returns a tuple with the Range6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6ClassName

`func (o *DhcpRange6EditInput) SetRange6ClassName(v string)`

SetRange6ClassName sets Range6ClassName field to given value.

### HasRange6ClassName

`func (o *DhcpRange6EditInput) HasRange6ClassName() bool`

HasRange6ClassName returns a boolean if a field has been set.

### GetRange6ClassParameters

`func (o *DhcpRange6EditInput) GetRange6ClassParameters() []ApiClassParameterInputEntry`

GetRange6ClassParameters returns the Range6ClassParameters field if non-nil, zero value otherwise.

### GetRange6ClassParametersOk

`func (o *DhcpRange6EditInput) GetRange6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetRange6ClassParametersOk returns a tuple with the Range6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange6ClassParameters

`func (o *DhcpRange6EditInput) SetRange6ClassParameters(v []ApiClassParameterInputEntry)`

SetRange6ClassParameters sets Range6ClassParameters field to given value.

### HasRange6ClassParameters

`func (o *DhcpRange6EditInput) HasRange6ClassParameters() bool`

HasRange6ClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpRange6EditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpRange6EditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpRange6EditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpRange6EditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


