# AppPoolAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFqdn** | Pointer to **string** | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. | [optional] 
**ApplicationId** | Pointer to **int32** | The database identifier (ID) of the application the object belongs to. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application the object belongs to. | [optional] 
**PoolLbMode** | Pointer to **string** | The load-balancing mode of the pool. | [optional] 
**PoolName** | Pointer to **string** | The name of the pool. | [optional] 
**PoolType** | Pointer to **string** | The type of the pool. | [optional] 
**PoolAffinitySessionTime** | Pointer to **int32** | The session duration, in seconds. You only need to set it if the parameter &lt;b&gt;pool_affinity_state&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt;. | [optional] 
**PoolAffinityState** | Pointer to **int32** | The session affinity activation status. | [optional] 
**GslbserverId** | Pointer to **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique numeric key value automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | [optional] 
**PoolBestActiveNodes** | Pointer to **int32** | The maximum number of active nodes with the lowest latency that must answer the queries made to the application FQDN. You only need to set it if you set the &lt;b&gt;pool_lb_mode&lt;/b&gt; to &lt;b&gt;latency&lt;/b&gt;. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewAppPoolAddInput

`func NewAppPoolAddInput() *AppPoolAddInput`

NewAppPoolAddInput instantiates a new AppPoolAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPoolAddInputWithDefaults

`func NewAppPoolAddInputWithDefaults() *AppPoolAddInput`

NewAppPoolAddInputWithDefaults instantiates a new AppPoolAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFqdn

`func (o *AppPoolAddInput) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *AppPoolAddInput) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *AppPoolAddInput) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *AppPoolAddInput) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetApplicationId

`func (o *AppPoolAddInput) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AppPoolAddInput) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AppPoolAddInput) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AppPoolAddInput) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *AppPoolAddInput) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AppPoolAddInput) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AppPoolAddInput) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AppPoolAddInput) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetPoolLbMode

`func (o *AppPoolAddInput) GetPoolLbMode() string`

GetPoolLbMode returns the PoolLbMode field if non-nil, zero value otherwise.

### GetPoolLbModeOk

`func (o *AppPoolAddInput) GetPoolLbModeOk() (*string, bool)`

GetPoolLbModeOk returns a tuple with the PoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLbMode

`func (o *AppPoolAddInput) SetPoolLbMode(v string)`

SetPoolLbMode sets PoolLbMode field to given value.

### HasPoolLbMode

`func (o *AppPoolAddInput) HasPoolLbMode() bool`

HasPoolLbMode returns a boolean if a field has been set.

### GetPoolName

`func (o *AppPoolAddInput) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AppPoolAddInput) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AppPoolAddInput) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AppPoolAddInput) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolType

`func (o *AppPoolAddInput) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *AppPoolAddInput) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *AppPoolAddInput) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *AppPoolAddInput) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetPoolAffinitySessionTime

`func (o *AppPoolAddInput) GetPoolAffinitySessionTime() int32`

GetPoolAffinitySessionTime returns the PoolAffinitySessionTime field if non-nil, zero value otherwise.

### GetPoolAffinitySessionTimeOk

`func (o *AppPoolAddInput) GetPoolAffinitySessionTimeOk() (*int32, bool)`

GetPoolAffinitySessionTimeOk returns a tuple with the PoolAffinitySessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAffinitySessionTime

`func (o *AppPoolAddInput) SetPoolAffinitySessionTime(v int32)`

SetPoolAffinitySessionTime sets PoolAffinitySessionTime field to given value.

### HasPoolAffinitySessionTime

`func (o *AppPoolAddInput) HasPoolAffinitySessionTime() bool`

HasPoolAffinitySessionTime returns a boolean if a field has been set.

### GetPoolAffinityState

`func (o *AppPoolAddInput) GetPoolAffinityState() int32`

GetPoolAffinityState returns the PoolAffinityState field if non-nil, zero value otherwise.

### GetPoolAffinityStateOk

`func (o *AppPoolAddInput) GetPoolAffinityStateOk() (*int32, bool)`

GetPoolAffinityStateOk returns a tuple with the PoolAffinityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAffinityState

`func (o *AppPoolAddInput) SetPoolAffinityState(v int32)`

SetPoolAffinityState sets PoolAffinityState field to given value.

### HasPoolAffinityState

`func (o *AppPoolAddInput) HasPoolAffinityState() bool`

HasPoolAffinityState returns a boolean if a field has been set.

### GetGslbserverId

`func (o *AppPoolAddInput) GetGslbserverId() int32`

GetGslbserverId returns the GslbserverId field if non-nil, zero value otherwise.

### GetGslbserverIdOk

`func (o *AppPoolAddInput) GetGslbserverIdOk() (*int32, bool)`

GetGslbserverIdOk returns a tuple with the GslbserverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverId

`func (o *AppPoolAddInput) SetGslbserverId(v int32)`

SetGslbserverId sets GslbserverId field to given value.

### HasGslbserverId

`func (o *AppPoolAddInput) HasGslbserverId() bool`

HasGslbserverId returns a boolean if a field has been set.

### GetPoolBestActiveNodes

`func (o *AppPoolAddInput) GetPoolBestActiveNodes() int32`

GetPoolBestActiveNodes returns the PoolBestActiveNodes field if non-nil, zero value otherwise.

### GetPoolBestActiveNodesOk

`func (o *AppPoolAddInput) GetPoolBestActiveNodesOk() (*int32, bool)`

GetPoolBestActiveNodesOk returns a tuple with the PoolBestActiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolBestActiveNodes

`func (o *AppPoolAddInput) SetPoolBestActiveNodes(v int32)`

SetPoolBestActiveNodes sets PoolBestActiveNodes field to given value.

### HasPoolBestActiveNodes

`func (o *AppPoolAddInput) HasPoolBestActiveNodes() bool`

HasPoolBestActiveNodes returns a boolean if a field has been set.

### GetWarnings

`func (o *AppPoolAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AppPoolAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AppPoolAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AppPoolAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


