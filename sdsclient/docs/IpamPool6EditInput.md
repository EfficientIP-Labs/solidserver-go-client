# IpamPool6EditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool6EndIpAddr** | Pointer to **string** | The last IP address of the pool. | [optional] 
**Pool6Id** | Pointer to **int32** | The database identifier (ID) of the IPv6 pool, a unique numeric key value automatically incremented when you add an IPv6 pool. Use the ID to specify the IPv6 pool of your choice. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**Pool6StartIpAddr** | Pointer to **string** | The first IP address of the pool. | [optional] 
**Network6Id** | Pointer to **int32** | The database identifier (ID) of the IPv6 network, a unique numeric key value automatically incremented when you add an IPv6 network. Use the ID to specify the IPv6 network of your choice. | [optional] 
**Pool6Name** | Pointer to **string** | The name of the IPv6 pool, each IPv6 pool must have a unique name. | [optional] 
**Pool6ReadOnly** | Pointer to **int32** | The reservation status of the IPv6 pool. If set 1, the IP addresses it contains cannot be assigned. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**Pool6ClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**Pool6ClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamPool6EditInput

`func NewIpamPool6EditInput() *IpamPool6EditInput`

NewIpamPool6EditInput instantiates a new IpamPool6EditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamPool6EditInputWithDefaults

`func NewIpamPool6EditInputWithDefaults() *IpamPool6EditInput`

NewIpamPool6EditInputWithDefaults instantiates a new IpamPool6EditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool6EndIpAddr

`func (o *IpamPool6EditInput) GetPool6EndIpAddr() string`

GetPool6EndIpAddr returns the Pool6EndIpAddr field if non-nil, zero value otherwise.

### GetPool6EndIpAddrOk

`func (o *IpamPool6EditInput) GetPool6EndIpAddrOk() (*string, bool)`

GetPool6EndIpAddrOk returns a tuple with the Pool6EndIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6EndIpAddr

`func (o *IpamPool6EditInput) SetPool6EndIpAddr(v string)`

SetPool6EndIpAddr sets Pool6EndIpAddr field to given value.

### HasPool6EndIpAddr

`func (o *IpamPool6EditInput) HasPool6EndIpAddr() bool`

HasPool6EndIpAddr returns a boolean if a field has been set.

### GetPool6Id

`func (o *IpamPool6EditInput) GetPool6Id() int32`

GetPool6Id returns the Pool6Id field if non-nil, zero value otherwise.

### GetPool6IdOk

`func (o *IpamPool6EditInput) GetPool6IdOk() (*int32, bool)`

GetPool6IdOk returns a tuple with the Pool6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Id

`func (o *IpamPool6EditInput) SetPool6Id(v int32)`

SetPool6Id sets Pool6Id field to given value.

### HasPool6Id

`func (o *IpamPool6EditInput) HasPool6Id() bool`

HasPool6Id returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamPool6EditInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamPool6EditInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamPool6EditInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamPool6EditInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamPool6EditInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamPool6EditInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamPool6EditInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamPool6EditInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetPool6StartIpAddr

`func (o *IpamPool6EditInput) GetPool6StartIpAddr() string`

GetPool6StartIpAddr returns the Pool6StartIpAddr field if non-nil, zero value otherwise.

### GetPool6StartIpAddrOk

`func (o *IpamPool6EditInput) GetPool6StartIpAddrOk() (*string, bool)`

GetPool6StartIpAddrOk returns a tuple with the Pool6StartIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6StartIpAddr

`func (o *IpamPool6EditInput) SetPool6StartIpAddr(v string)`

SetPool6StartIpAddr sets Pool6StartIpAddr field to given value.

### HasPool6StartIpAddr

`func (o *IpamPool6EditInput) HasPool6StartIpAddr() bool`

HasPool6StartIpAddr returns a boolean if a field has been set.

### GetNetwork6Id

`func (o *IpamPool6EditInput) GetNetwork6Id() int32`

GetNetwork6Id returns the Network6Id field if non-nil, zero value otherwise.

### GetNetwork6IdOk

`func (o *IpamPool6EditInput) GetNetwork6IdOk() (*int32, bool)`

GetNetwork6IdOk returns a tuple with the Network6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Id

`func (o *IpamPool6EditInput) SetNetwork6Id(v int32)`

SetNetwork6Id sets Network6Id field to given value.

### HasNetwork6Id

`func (o *IpamPool6EditInput) HasNetwork6Id() bool`

HasNetwork6Id returns a boolean if a field has been set.

### GetPool6Name

`func (o *IpamPool6EditInput) GetPool6Name() string`

GetPool6Name returns the Pool6Name field if non-nil, zero value otherwise.

### GetPool6NameOk

`func (o *IpamPool6EditInput) GetPool6NameOk() (*string, bool)`

GetPool6NameOk returns a tuple with the Pool6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Name

`func (o *IpamPool6EditInput) SetPool6Name(v string)`

SetPool6Name sets Pool6Name field to given value.

### HasPool6Name

`func (o *IpamPool6EditInput) HasPool6Name() bool`

HasPool6Name returns a boolean if a field has been set.

### GetPool6ReadOnly

`func (o *IpamPool6EditInput) GetPool6ReadOnly() int32`

GetPool6ReadOnly returns the Pool6ReadOnly field if non-nil, zero value otherwise.

### GetPool6ReadOnlyOk

`func (o *IpamPool6EditInput) GetPool6ReadOnlyOk() (*int32, bool)`

GetPool6ReadOnlyOk returns a tuple with the Pool6ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ReadOnly

`func (o *IpamPool6EditInput) SetPool6ReadOnly(v int32)`

SetPool6ReadOnly sets Pool6ReadOnly field to given value.

### HasPool6ReadOnly

`func (o *IpamPool6EditInput) HasPool6ReadOnly() bool`

HasPool6ReadOnly returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *IpamPool6EditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *IpamPool6EditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *IpamPool6EditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *IpamPool6EditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetPool6ClassName

`func (o *IpamPool6EditInput) GetPool6ClassName() string`

GetPool6ClassName returns the Pool6ClassName field if non-nil, zero value otherwise.

### GetPool6ClassNameOk

`func (o *IpamPool6EditInput) GetPool6ClassNameOk() (*string, bool)`

GetPool6ClassNameOk returns a tuple with the Pool6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ClassName

`func (o *IpamPool6EditInput) SetPool6ClassName(v string)`

SetPool6ClassName sets Pool6ClassName field to given value.

### HasPool6ClassName

`func (o *IpamPool6EditInput) HasPool6ClassName() bool`

HasPool6ClassName returns a boolean if a field has been set.

### GetPool6ClassParameters

`func (o *IpamPool6EditInput) GetPool6ClassParameters() []ApiClassParameterInputEntry`

GetPool6ClassParameters returns the Pool6ClassParameters field if non-nil, zero value otherwise.

### GetPool6ClassParametersOk

`func (o *IpamPool6EditInput) GetPool6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetPool6ClassParametersOk returns a tuple with the Pool6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ClassParameters

`func (o *IpamPool6EditInput) SetPool6ClassParameters(v []ApiClassParameterInputEntry)`

SetPool6ClassParameters sets Pool6ClassParameters field to given value.

### HasPool6ClassParameters

`func (o *IpamPool6EditInput) HasPool6ClassParameters() bool`

HasPool6ClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamPool6EditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamPool6EditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamPool6EditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamPool6EditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


