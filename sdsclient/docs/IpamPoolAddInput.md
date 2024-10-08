# IpamPoolAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolEndIpAddr** | Pointer to **string** | The last IP address of the pool. | [optional] 
**PoolSize** | Pointer to **int32** | The size of the pool, the number of IP addresses it contains. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**PoolStartIpAddr** | Pointer to **string** | The first IP address of the pool. | [optional] 
**NetworkId** | Pointer to **int32** | The database identifier (ID) of the IPv4 network, a unique numeric key value automatically incremented when you add an IPv4 network. Use the ID to specify the IPv4 network of your choice. | [optional] 
**PoolName** | Pointer to **string** | The name of the IPv4 pool, each IPv4 pool must have a unique name. | [optional] 
**PoolReadOnly** | Pointer to **int32** | The reservation status of the IPv4 pool. If set 1, the IP addresses it contains cannot be assigned. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**PoolClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**PoolClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamPoolAddInput

`func NewIpamPoolAddInput() *IpamPoolAddInput`

NewIpamPoolAddInput instantiates a new IpamPoolAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamPoolAddInputWithDefaults

`func NewIpamPoolAddInputWithDefaults() *IpamPoolAddInput`

NewIpamPoolAddInputWithDefaults instantiates a new IpamPoolAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolEndIpAddr

`func (o *IpamPoolAddInput) GetPoolEndIpAddr() string`

GetPoolEndIpAddr returns the PoolEndIpAddr field if non-nil, zero value otherwise.

### GetPoolEndIpAddrOk

`func (o *IpamPoolAddInput) GetPoolEndIpAddrOk() (*string, bool)`

GetPoolEndIpAddrOk returns a tuple with the PoolEndIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEndIpAddr

`func (o *IpamPoolAddInput) SetPoolEndIpAddr(v string)`

SetPoolEndIpAddr sets PoolEndIpAddr field to given value.

### HasPoolEndIpAddr

`func (o *IpamPoolAddInput) HasPoolEndIpAddr() bool`

HasPoolEndIpAddr returns a boolean if a field has been set.

### GetPoolSize

`func (o *IpamPoolAddInput) GetPoolSize() int32`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *IpamPoolAddInput) GetPoolSizeOk() (*int32, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *IpamPoolAddInput) SetPoolSize(v int32)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *IpamPoolAddInput) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamPoolAddInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamPoolAddInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamPoolAddInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamPoolAddInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamPoolAddInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamPoolAddInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamPoolAddInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamPoolAddInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetPoolStartIpAddr

`func (o *IpamPoolAddInput) GetPoolStartIpAddr() string`

GetPoolStartIpAddr returns the PoolStartIpAddr field if non-nil, zero value otherwise.

### GetPoolStartIpAddrOk

`func (o *IpamPoolAddInput) GetPoolStartIpAddrOk() (*string, bool)`

GetPoolStartIpAddrOk returns a tuple with the PoolStartIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStartIpAddr

`func (o *IpamPoolAddInput) SetPoolStartIpAddr(v string)`

SetPoolStartIpAddr sets PoolStartIpAddr field to given value.

### HasPoolStartIpAddr

`func (o *IpamPoolAddInput) HasPoolStartIpAddr() bool`

HasPoolStartIpAddr returns a boolean if a field has been set.

### GetNetworkId

`func (o *IpamPoolAddInput) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *IpamPoolAddInput) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *IpamPoolAddInput) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *IpamPoolAddInput) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPoolName

`func (o *IpamPoolAddInput) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *IpamPoolAddInput) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *IpamPoolAddInput) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *IpamPoolAddInput) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolReadOnly

`func (o *IpamPoolAddInput) GetPoolReadOnly() int32`

GetPoolReadOnly returns the PoolReadOnly field if non-nil, zero value otherwise.

### GetPoolReadOnlyOk

`func (o *IpamPoolAddInput) GetPoolReadOnlyOk() (*int32, bool)`

GetPoolReadOnlyOk returns a tuple with the PoolReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolReadOnly

`func (o *IpamPoolAddInput) SetPoolReadOnly(v int32)`

SetPoolReadOnly sets PoolReadOnly field to given value.

### HasPoolReadOnly

`func (o *IpamPoolAddInput) HasPoolReadOnly() bool`

HasPoolReadOnly returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *IpamPoolAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *IpamPoolAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *IpamPoolAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *IpamPoolAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetPoolClassName

`func (o *IpamPoolAddInput) GetPoolClassName() string`

GetPoolClassName returns the PoolClassName field if non-nil, zero value otherwise.

### GetPoolClassNameOk

`func (o *IpamPoolAddInput) GetPoolClassNameOk() (*string, bool)`

GetPoolClassNameOk returns a tuple with the PoolClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassName

`func (o *IpamPoolAddInput) SetPoolClassName(v string)`

SetPoolClassName sets PoolClassName field to given value.

### HasPoolClassName

`func (o *IpamPoolAddInput) HasPoolClassName() bool`

HasPoolClassName returns a boolean if a field has been set.

### GetPoolClassParameters

`func (o *IpamPoolAddInput) GetPoolClassParameters() []ApiClassParameterInputEntry`

GetPoolClassParameters returns the PoolClassParameters field if non-nil, zero value otherwise.

### GetPoolClassParametersOk

`func (o *IpamPoolAddInput) GetPoolClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetPoolClassParametersOk returns a tuple with the PoolClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassParameters

`func (o *IpamPoolAddInput) SetPoolClassParameters(v []ApiClassParameterInputEntry)`

SetPoolClassParameters sets PoolClassParameters field to given value.

### HasPoolClassParameters

`func (o *IpamPoolAddInput) HasPoolClassParameters() bool`

HasPoolClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamPoolAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamPoolAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamPoolAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamPoolAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


