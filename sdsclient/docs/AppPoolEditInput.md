# AppPoolEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFqdn** | Pointer to **string** | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. | [optional] 
**ApplicationId** | Pointer to **int32** | The database identifier (ID) of the application the object belongs to. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application the object belongs to. | [optional] 
**PoolId** | Pointer to **int32** | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. | [optional] 
**PoolLbMode** | Pointer to **string** | The load-balancing mode of the pool. | [optional] 
**PoolName** | Pointer to **string** | The name of the pool. | [optional] 
**PoolType** | Pointer to **string** | The type of the pool. | [optional] 
**PoolAffinitySessionTime** | Pointer to **int32** | The session duration, in seconds. You only need to set it if the parameter &lt;b&gt;pool_affinity_state&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt;. | [optional] 
**PoolAffinityState** | Pointer to **int32** | The session affinity activation status. | [optional] 
**GslbserverId** | Pointer to **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique numeric key value automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | [optional] 
**PoolBestActiveNodes** | Pointer to **int32** | The maximum number of active nodes with the lowest latency that must answer the queries made to the application FQDN. You only need to set it if you set the &lt;b&gt;pool_lb_mode&lt;/b&gt; to &lt;b&gt;latency&lt;/b&gt;. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewAppPoolEditInput

`func NewAppPoolEditInput() *AppPoolEditInput`

NewAppPoolEditInput instantiates a new AppPoolEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPoolEditInputWithDefaults

`func NewAppPoolEditInputWithDefaults() *AppPoolEditInput`

NewAppPoolEditInputWithDefaults instantiates a new AppPoolEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFqdn

`func (o *AppPoolEditInput) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *AppPoolEditInput) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *AppPoolEditInput) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *AppPoolEditInput) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetApplicationId

`func (o *AppPoolEditInput) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AppPoolEditInput) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AppPoolEditInput) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AppPoolEditInput) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *AppPoolEditInput) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AppPoolEditInput) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AppPoolEditInput) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AppPoolEditInput) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetPoolId

`func (o *AppPoolEditInput) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AppPoolEditInput) GetPoolIdOk() (*int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AppPoolEditInput) SetPoolId(v int32)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AppPoolEditInput) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolLbMode

`func (o *AppPoolEditInput) GetPoolLbMode() string`

GetPoolLbMode returns the PoolLbMode field if non-nil, zero value otherwise.

### GetPoolLbModeOk

`func (o *AppPoolEditInput) GetPoolLbModeOk() (*string, bool)`

GetPoolLbModeOk returns a tuple with the PoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLbMode

`func (o *AppPoolEditInput) SetPoolLbMode(v string)`

SetPoolLbMode sets PoolLbMode field to given value.

### HasPoolLbMode

`func (o *AppPoolEditInput) HasPoolLbMode() bool`

HasPoolLbMode returns a boolean if a field has been set.

### GetPoolName

`func (o *AppPoolEditInput) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AppPoolEditInput) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AppPoolEditInput) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AppPoolEditInput) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolType

`func (o *AppPoolEditInput) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *AppPoolEditInput) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *AppPoolEditInput) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *AppPoolEditInput) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetPoolAffinitySessionTime

`func (o *AppPoolEditInput) GetPoolAffinitySessionTime() int32`

GetPoolAffinitySessionTime returns the PoolAffinitySessionTime field if non-nil, zero value otherwise.

### GetPoolAffinitySessionTimeOk

`func (o *AppPoolEditInput) GetPoolAffinitySessionTimeOk() (*int32, bool)`

GetPoolAffinitySessionTimeOk returns a tuple with the PoolAffinitySessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAffinitySessionTime

`func (o *AppPoolEditInput) SetPoolAffinitySessionTime(v int32)`

SetPoolAffinitySessionTime sets PoolAffinitySessionTime field to given value.

### HasPoolAffinitySessionTime

`func (o *AppPoolEditInput) HasPoolAffinitySessionTime() bool`

HasPoolAffinitySessionTime returns a boolean if a field has been set.

### GetPoolAffinityState

`func (o *AppPoolEditInput) GetPoolAffinityState() int32`

GetPoolAffinityState returns the PoolAffinityState field if non-nil, zero value otherwise.

### GetPoolAffinityStateOk

`func (o *AppPoolEditInput) GetPoolAffinityStateOk() (*int32, bool)`

GetPoolAffinityStateOk returns a tuple with the PoolAffinityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAffinityState

`func (o *AppPoolEditInput) SetPoolAffinityState(v int32)`

SetPoolAffinityState sets PoolAffinityState field to given value.

### HasPoolAffinityState

`func (o *AppPoolEditInput) HasPoolAffinityState() bool`

HasPoolAffinityState returns a boolean if a field has been set.

### GetGslbserverId

`func (o *AppPoolEditInput) GetGslbserverId() int32`

GetGslbserverId returns the GslbserverId field if non-nil, zero value otherwise.

### GetGslbserverIdOk

`func (o *AppPoolEditInput) GetGslbserverIdOk() (*int32, bool)`

GetGslbserverIdOk returns a tuple with the GslbserverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverId

`func (o *AppPoolEditInput) SetGslbserverId(v int32)`

SetGslbserverId sets GslbserverId field to given value.

### HasGslbserverId

`func (o *AppPoolEditInput) HasGslbserverId() bool`

HasGslbserverId returns a boolean if a field has been set.

### GetPoolBestActiveNodes

`func (o *AppPoolEditInput) GetPoolBestActiveNodes() int32`

GetPoolBestActiveNodes returns the PoolBestActiveNodes field if non-nil, zero value otherwise.

### GetPoolBestActiveNodesOk

`func (o *AppPoolEditInput) GetPoolBestActiveNodesOk() (*int32, bool)`

GetPoolBestActiveNodesOk returns a tuple with the PoolBestActiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolBestActiveNodes

`func (o *AppPoolEditInput) SetPoolBestActiveNodes(v int32)`

SetPoolBestActiveNodes sets PoolBestActiveNodes field to given value.

### HasPoolBestActiveNodes

`func (o *AppPoolEditInput) HasPoolBestActiveNodes() bool`

HasPoolBestActiveNodes returns a boolean if a field has been set.

### GetWarnings

`func (o *AppPoolEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AppPoolEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AppPoolEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AppPoolEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


