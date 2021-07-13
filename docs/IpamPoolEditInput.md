# IpamPoolEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolEndAddressAddr** | Pointer to **string** | todo(here) :ipam.pool.edit.input.pool_end_address_addr. : IPv4 address | [optional] 
**PoolId** | Pointer to **int32** | The database identifier (ID) of the IPv4 pool, a unique numeric key value automatically incremented when you add an IPv4 pool. Use the ID to specify which IPv4 pool to edit. | [optional] 
**PoolSize** | Pointer to **int32** | The size of the pool, the number of IP addresses it contains. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**PoolStartAddressAddr** | Pointer to **string** | todo(here) :ipam.pool.edit.input.pool_start_address_addr. : IPv4 address | [optional] 
**NetworkId** | Pointer to **int32** | The database identifier (ID) of the IPv4 network, a unique numeric key value automatically incremented when you add an IPv4 network. Use the ID to specify the IPv4 network of your choice. | [optional] 
**PoolName** | Pointer to **string** | The name of the IPv4 pool, each IPv4 pool must have a unique name. | [optional] 
**PoolReadOnly** | Pointer to **int32** | The reservation status of the IPv4 pool. If set 1, the IP addresses it contains cannot be assigned. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**PoolClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**PoolClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamPoolEditInput

`func NewIpamPoolEditInput() *IpamPoolEditInput`

NewIpamPoolEditInput instantiates a new IpamPoolEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamPoolEditInputWithDefaults

`func NewIpamPoolEditInputWithDefaults() *IpamPoolEditInput`

NewIpamPoolEditInputWithDefaults instantiates a new IpamPoolEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolEndAddressAddr

`func (o *IpamPoolEditInput) GetPoolEndAddressAddr() string`

GetPoolEndAddressAddr returns the PoolEndAddressAddr field if non-nil, zero value otherwise.

### GetPoolEndAddressAddrOk

`func (o *IpamPoolEditInput) GetPoolEndAddressAddrOk() (*string, bool)`

GetPoolEndAddressAddrOk returns a tuple with the PoolEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEndAddressAddr

`func (o *IpamPoolEditInput) SetPoolEndAddressAddr(v string)`

SetPoolEndAddressAddr sets PoolEndAddressAddr field to given value.

### HasPoolEndAddressAddr

`func (o *IpamPoolEditInput) HasPoolEndAddressAddr() bool`

HasPoolEndAddressAddr returns a boolean if a field has been set.

### GetPoolId

`func (o *IpamPoolEditInput) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *IpamPoolEditInput) GetPoolIdOk() (*int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *IpamPoolEditInput) SetPoolId(v int32)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *IpamPoolEditInput) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolSize

`func (o *IpamPoolEditInput) GetPoolSize() int32`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *IpamPoolEditInput) GetPoolSizeOk() (*int32, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *IpamPoolEditInput) SetPoolSize(v int32)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *IpamPoolEditInput) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamPoolEditInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamPoolEditInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamPoolEditInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamPoolEditInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamPoolEditInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamPoolEditInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamPoolEditInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamPoolEditInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetPoolStartAddressAddr

`func (o *IpamPoolEditInput) GetPoolStartAddressAddr() string`

GetPoolStartAddressAddr returns the PoolStartAddressAddr field if non-nil, zero value otherwise.

### GetPoolStartAddressAddrOk

`func (o *IpamPoolEditInput) GetPoolStartAddressAddrOk() (*string, bool)`

GetPoolStartAddressAddrOk returns a tuple with the PoolStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStartAddressAddr

`func (o *IpamPoolEditInput) SetPoolStartAddressAddr(v string)`

SetPoolStartAddressAddr sets PoolStartAddressAddr field to given value.

### HasPoolStartAddressAddr

`func (o *IpamPoolEditInput) HasPoolStartAddressAddr() bool`

HasPoolStartAddressAddr returns a boolean if a field has been set.

### GetNetworkId

`func (o *IpamPoolEditInput) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *IpamPoolEditInput) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *IpamPoolEditInput) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *IpamPoolEditInput) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetPoolName

`func (o *IpamPoolEditInput) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *IpamPoolEditInput) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *IpamPoolEditInput) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *IpamPoolEditInput) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolReadOnly

`func (o *IpamPoolEditInput) GetPoolReadOnly() int32`

GetPoolReadOnly returns the PoolReadOnly field if non-nil, zero value otherwise.

### GetPoolReadOnlyOk

`func (o *IpamPoolEditInput) GetPoolReadOnlyOk() (*int32, bool)`

GetPoolReadOnlyOk returns a tuple with the PoolReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolReadOnly

`func (o *IpamPoolEditInput) SetPoolReadOnly(v int32)`

SetPoolReadOnly sets PoolReadOnly field to given value.

### HasPoolReadOnly

`func (o *IpamPoolEditInput) HasPoolReadOnly() bool`

HasPoolReadOnly returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *IpamPoolEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *IpamPoolEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *IpamPoolEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *IpamPoolEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetPoolClassName

`func (o *IpamPoolEditInput) GetPoolClassName() string`

GetPoolClassName returns the PoolClassName field if non-nil, zero value otherwise.

### GetPoolClassNameOk

`func (o *IpamPoolEditInput) GetPoolClassNameOk() (*string, bool)`

GetPoolClassNameOk returns a tuple with the PoolClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassName

`func (o *IpamPoolEditInput) SetPoolClassName(v string)`

SetPoolClassName sets PoolClassName field to given value.

### HasPoolClassName

`func (o *IpamPoolEditInput) HasPoolClassName() bool`

HasPoolClassName returns a boolean if a field has been set.

### GetPoolClassParameters

`func (o *IpamPoolEditInput) GetPoolClassParameters() []ApiClassParameterInputEntry`

GetPoolClassParameters returns the PoolClassParameters field if non-nil, zero value otherwise.

### GetPoolClassParametersOk

`func (o *IpamPoolEditInput) GetPoolClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetPoolClassParametersOk returns a tuple with the PoolClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassParameters

`func (o *IpamPoolEditInput) SetPoolClassParameters(v []ApiClassParameterInputEntry)`

SetPoolClassParameters sets PoolClassParameters field to given value.

### HasPoolClassParameters

`func (o *IpamPoolEditInput) HasPoolClassParameters() bool`

HasPoolClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamPoolEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamPoolEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamPoolEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamPoolEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


