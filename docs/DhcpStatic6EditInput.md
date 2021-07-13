# DhcpStatic6EditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 server, a unique numeric key value automatically incremented when you add a DHCPv6 server. Use the ID to specify the DHCPv6 server of your choice. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server. | [optional] 
**Static6ClientDuid** | Pointer to **string** | The client DHCP Unique Identifier (DUID) associated with the DHCPv6 static. | [optional] 
**Static6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 static, a unique numeric key value automatically incremented when you add a DHCPv6 static. Use the ID to specify which DHCPv6 static to edit. | [optional] 
**Static6MacAddr** | Pointer to **string** | The MAC address you want to associate with the IPv6 static. | [optional] 
**Static6Name** | Pointer to **string** | The name of the DHCPv6 static, each DHCPv6 static must have a unique name. | [optional] 
**Server6Hostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**Group6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 group, a unique numeric key value automatically incremented when you add a DHCPv6 group. Use the ID to specify the DHCPv6 group of your choice. | [optional] 
**Group6Name** | Pointer to **string** | The name of the DHCPv6 group. | [optional] 
**Static6Addr** | Pointer to **string** | The IP address associated with the DHCPv6 static. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**Static6ClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**Static6ClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpStatic6EditInput

`func NewDhcpStatic6EditInput() *DhcpStatic6EditInput`

NewDhcpStatic6EditInput instantiates a new DhcpStatic6EditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpStatic6EditInputWithDefaults

`func NewDhcpStatic6EditInputWithDefaults() *DhcpStatic6EditInput`

NewDhcpStatic6EditInputWithDefaults instantiates a new DhcpStatic6EditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer6Id

`func (o *DhcpStatic6EditInput) GetServer6Id() int32`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DhcpStatic6EditInput) GetServer6IdOk() (*int32, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DhcpStatic6EditInput) SetServer6Id(v int32)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DhcpStatic6EditInput) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DhcpStatic6EditInput) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DhcpStatic6EditInput) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DhcpStatic6EditInput) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DhcpStatic6EditInput) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetStatic6ClientDuid

`func (o *DhcpStatic6EditInput) GetStatic6ClientDuid() string`

GetStatic6ClientDuid returns the Static6ClientDuid field if non-nil, zero value otherwise.

### GetStatic6ClientDuidOk

`func (o *DhcpStatic6EditInput) GetStatic6ClientDuidOk() (*string, bool)`

GetStatic6ClientDuidOk returns a tuple with the Static6ClientDuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6ClientDuid

`func (o *DhcpStatic6EditInput) SetStatic6ClientDuid(v string)`

SetStatic6ClientDuid sets Static6ClientDuid field to given value.

### HasStatic6ClientDuid

`func (o *DhcpStatic6EditInput) HasStatic6ClientDuid() bool`

HasStatic6ClientDuid returns a boolean if a field has been set.

### GetStatic6Id

`func (o *DhcpStatic6EditInput) GetStatic6Id() int32`

GetStatic6Id returns the Static6Id field if non-nil, zero value otherwise.

### GetStatic6IdOk

`func (o *DhcpStatic6EditInput) GetStatic6IdOk() (*int32, bool)`

GetStatic6IdOk returns a tuple with the Static6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Id

`func (o *DhcpStatic6EditInput) SetStatic6Id(v int32)`

SetStatic6Id sets Static6Id field to given value.

### HasStatic6Id

`func (o *DhcpStatic6EditInput) HasStatic6Id() bool`

HasStatic6Id returns a boolean if a field has been set.

### GetStatic6MacAddr

`func (o *DhcpStatic6EditInput) GetStatic6MacAddr() string`

GetStatic6MacAddr returns the Static6MacAddr field if non-nil, zero value otherwise.

### GetStatic6MacAddrOk

`func (o *DhcpStatic6EditInput) GetStatic6MacAddrOk() (*string, bool)`

GetStatic6MacAddrOk returns a tuple with the Static6MacAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6MacAddr

`func (o *DhcpStatic6EditInput) SetStatic6MacAddr(v string)`

SetStatic6MacAddr sets Static6MacAddr field to given value.

### HasStatic6MacAddr

`func (o *DhcpStatic6EditInput) HasStatic6MacAddr() bool`

HasStatic6MacAddr returns a boolean if a field has been set.

### GetStatic6Name

`func (o *DhcpStatic6EditInput) GetStatic6Name() string`

GetStatic6Name returns the Static6Name field if non-nil, zero value otherwise.

### GetStatic6NameOk

`func (o *DhcpStatic6EditInput) GetStatic6NameOk() (*string, bool)`

GetStatic6NameOk returns a tuple with the Static6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Name

`func (o *DhcpStatic6EditInput) SetStatic6Name(v string)`

SetStatic6Name sets Static6Name field to given value.

### HasStatic6Name

`func (o *DhcpStatic6EditInput) HasStatic6Name() bool`

HasStatic6Name returns a boolean if a field has been set.

### GetServer6Hostaddr

`func (o *DhcpStatic6EditInput) GetServer6Hostaddr() string`

GetServer6Hostaddr returns the Server6Hostaddr field if non-nil, zero value otherwise.

### GetServer6HostaddrOk

`func (o *DhcpStatic6EditInput) GetServer6HostaddrOk() (*string, bool)`

GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Hostaddr

`func (o *DhcpStatic6EditInput) SetServer6Hostaddr(v string)`

SetServer6Hostaddr sets Server6Hostaddr field to given value.

### HasServer6Hostaddr

`func (o *DhcpStatic6EditInput) HasServer6Hostaddr() bool`

HasServer6Hostaddr returns a boolean if a field has been set.

### GetGroup6Id

`func (o *DhcpStatic6EditInput) GetGroup6Id() int32`

GetGroup6Id returns the Group6Id field if non-nil, zero value otherwise.

### GetGroup6IdOk

`func (o *DhcpStatic6EditInput) GetGroup6IdOk() (*int32, bool)`

GetGroup6IdOk returns a tuple with the Group6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Id

`func (o *DhcpStatic6EditInput) SetGroup6Id(v int32)`

SetGroup6Id sets Group6Id field to given value.

### HasGroup6Id

`func (o *DhcpStatic6EditInput) HasGroup6Id() bool`

HasGroup6Id returns a boolean if a field has been set.

### GetGroup6Name

`func (o *DhcpStatic6EditInput) GetGroup6Name() string`

GetGroup6Name returns the Group6Name field if non-nil, zero value otherwise.

### GetGroup6NameOk

`func (o *DhcpStatic6EditInput) GetGroup6NameOk() (*string, bool)`

GetGroup6NameOk returns a tuple with the Group6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Name

`func (o *DhcpStatic6EditInput) SetGroup6Name(v string)`

SetGroup6Name sets Group6Name field to given value.

### HasGroup6Name

`func (o *DhcpStatic6EditInput) HasGroup6Name() bool`

HasGroup6Name returns a boolean if a field has been set.

### GetStatic6Addr

`func (o *DhcpStatic6EditInput) GetStatic6Addr() string`

GetStatic6Addr returns the Static6Addr field if non-nil, zero value otherwise.

### GetStatic6AddrOk

`func (o *DhcpStatic6EditInput) GetStatic6AddrOk() (*string, bool)`

GetStatic6AddrOk returns a tuple with the Static6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6Addr

`func (o *DhcpStatic6EditInput) SetStatic6Addr(v string)`

SetStatic6Addr sets Static6Addr field to given value.

### HasStatic6Addr

`func (o *DhcpStatic6EditInput) HasStatic6Addr() bool`

HasStatic6Addr returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DhcpStatic6EditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DhcpStatic6EditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DhcpStatic6EditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DhcpStatic6EditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetStatic6ClassName

`func (o *DhcpStatic6EditInput) GetStatic6ClassName() string`

GetStatic6ClassName returns the Static6ClassName field if non-nil, zero value otherwise.

### GetStatic6ClassNameOk

`func (o *DhcpStatic6EditInput) GetStatic6ClassNameOk() (*string, bool)`

GetStatic6ClassNameOk returns a tuple with the Static6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6ClassName

`func (o *DhcpStatic6EditInput) SetStatic6ClassName(v string)`

SetStatic6ClassName sets Static6ClassName field to given value.

### HasStatic6ClassName

`func (o *DhcpStatic6EditInput) HasStatic6ClassName() bool`

HasStatic6ClassName returns a boolean if a field has been set.

### GetStatic6ClassParameters

`func (o *DhcpStatic6EditInput) GetStatic6ClassParameters() []ApiClassParameterInputEntry`

GetStatic6ClassParameters returns the Static6ClassParameters field if non-nil, zero value otherwise.

### GetStatic6ClassParametersOk

`func (o *DhcpStatic6EditInput) GetStatic6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetStatic6ClassParametersOk returns a tuple with the Static6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatic6ClassParameters

`func (o *DhcpStatic6EditInput) SetStatic6ClassParameters(v []ApiClassParameterInputEntry)`

SetStatic6ClassParameters sets Static6ClassParameters field to given value.

### HasStatic6ClassParameters

`func (o *DhcpStatic6EditInput) HasStatic6ClassParameters() bool`

HasStatic6ClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpStatic6EditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpStatic6EditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpStatic6EditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpStatic6EditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


