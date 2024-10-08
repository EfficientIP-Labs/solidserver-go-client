# DataInnerAppPoolData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFqdn** | Pointer to **string** | The FQDN of the application the object belongs to. | [optional] 
**GslbserverId** | Pointer to **string** | The database identifier (ID) of the GSLB server associated with the application the object belongs to. It is only returned for deployed applications. | [optional] 
**GslbserverName** | Pointer to **string** | The name of the GSLB server associated with the application the object belongs to. It is only returned for deployed applications. | [optional] 
**ApplicationId** | Pointer to **string** | The database identifier (ID) of the application the object belongs to. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application the object belongs to. | [optional] 
**PoolAffinitySessionTime** | Pointer to **string** | The session duration, in seconds. It is only returned if the parameter &lt;b&gt;affinity_state&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt;. | [optional] 
**PoolAffinityState** | Pointer to **string** | The session affinity activation status. | [optional] 
**PoolBestActiveNodes** | Pointer to **string** | The maximum number of active nodes with the lowest latency that must answer the queries made to the application FQDN. It is only returned if the parameter &lt;b&gt;pool_lb_mode&lt;/b&gt; is set to &lt;b&gt;latency&lt;/b&gt;. | [optional] 
**PoolId** | Pointer to **string** | The database identifier (ID) of the pool. | [optional] 
**PoolInactiveNodes** | Pointer to **string** | The number nodes of the pool that are &lt;b&gt;Inactive&lt;/b&gt; (down). | [optional] 
**PoolLbMode** | Pointer to **string** | The load-balancing mode of the pool, either &lt;b&gt;weighted&lt;/b&gt;, &lt;b&gt;round-robin&lt;/b&gt; or &lt;b&gt;latency&lt;/b&gt;. | [optional] 
**PoolName** | Pointer to **string** | The name of the pool. | [optional] 
**PoolTotalNodes** | Pointer to **string** | The number of nodes of the pool. | [optional] 
**PoolType** | Pointer to **string** | The type of the pool, &lt;b&gt;ipv4&lt;/b&gt; or &lt;b&gt;ipv6&lt;/b&gt;. | [optional] 
**PoolWeight** | Pointer to **string** | Not documented. Internal use. | [optional] 
**PoolDelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**GslbserverType** | Pointer to **string** | The type of DNS server associated with the application the object belongs to. It is only returned for deployed applications. | [optional] 
**PoolMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentApplicationId** | Pointer to **string** | The database identifier (ID) of the application the object belongs to. It is only returned for deployed applications. | [optional] 
**TranslatedPoolLbMode** | Pointer to **string** | The load-balancing mode of the pool, as displayed in the GUI. | [optional] 

## Methods

### NewDataInnerAppPoolData

`func NewDataInnerAppPoolData() *DataInnerAppPoolData`

NewDataInnerAppPoolData instantiates a new DataInnerAppPoolData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerAppPoolDataWithDefaults

`func NewDataInnerAppPoolDataWithDefaults() *DataInnerAppPoolData`

NewDataInnerAppPoolDataWithDefaults instantiates a new DataInnerAppPoolData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFqdn

`func (o *DataInnerAppPoolData) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *DataInnerAppPoolData) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *DataInnerAppPoolData) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *DataInnerAppPoolData) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetGslbserverId

`func (o *DataInnerAppPoolData) GetGslbserverId() string`

GetGslbserverId returns the GslbserverId field if non-nil, zero value otherwise.

### GetGslbserverIdOk

`func (o *DataInnerAppPoolData) GetGslbserverIdOk() (*string, bool)`

GetGslbserverIdOk returns a tuple with the GslbserverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverId

`func (o *DataInnerAppPoolData) SetGslbserverId(v string)`

SetGslbserverId sets GslbserverId field to given value.

### HasGslbserverId

`func (o *DataInnerAppPoolData) HasGslbserverId() bool`

HasGslbserverId returns a boolean if a field has been set.

### GetGslbserverName

`func (o *DataInnerAppPoolData) GetGslbserverName() string`

GetGslbserverName returns the GslbserverName field if non-nil, zero value otherwise.

### GetGslbserverNameOk

`func (o *DataInnerAppPoolData) GetGslbserverNameOk() (*string, bool)`

GetGslbserverNameOk returns a tuple with the GslbserverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverName

`func (o *DataInnerAppPoolData) SetGslbserverName(v string)`

SetGslbserverName sets GslbserverName field to given value.

### HasGslbserverName

`func (o *DataInnerAppPoolData) HasGslbserverName() bool`

HasGslbserverName returns a boolean if a field has been set.

### GetApplicationId

`func (o *DataInnerAppPoolData) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *DataInnerAppPoolData) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *DataInnerAppPoolData) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *DataInnerAppPoolData) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *DataInnerAppPoolData) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *DataInnerAppPoolData) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *DataInnerAppPoolData) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *DataInnerAppPoolData) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetPoolAffinitySessionTime

`func (o *DataInnerAppPoolData) GetPoolAffinitySessionTime() string`

GetPoolAffinitySessionTime returns the PoolAffinitySessionTime field if non-nil, zero value otherwise.

### GetPoolAffinitySessionTimeOk

`func (o *DataInnerAppPoolData) GetPoolAffinitySessionTimeOk() (*string, bool)`

GetPoolAffinitySessionTimeOk returns a tuple with the PoolAffinitySessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAffinitySessionTime

`func (o *DataInnerAppPoolData) SetPoolAffinitySessionTime(v string)`

SetPoolAffinitySessionTime sets PoolAffinitySessionTime field to given value.

### HasPoolAffinitySessionTime

`func (o *DataInnerAppPoolData) HasPoolAffinitySessionTime() bool`

HasPoolAffinitySessionTime returns a boolean if a field has been set.

### GetPoolAffinityState

`func (o *DataInnerAppPoolData) GetPoolAffinityState() string`

GetPoolAffinityState returns the PoolAffinityState field if non-nil, zero value otherwise.

### GetPoolAffinityStateOk

`func (o *DataInnerAppPoolData) GetPoolAffinityStateOk() (*string, bool)`

GetPoolAffinityStateOk returns a tuple with the PoolAffinityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAffinityState

`func (o *DataInnerAppPoolData) SetPoolAffinityState(v string)`

SetPoolAffinityState sets PoolAffinityState field to given value.

### HasPoolAffinityState

`func (o *DataInnerAppPoolData) HasPoolAffinityState() bool`

HasPoolAffinityState returns a boolean if a field has been set.

### GetPoolBestActiveNodes

`func (o *DataInnerAppPoolData) GetPoolBestActiveNodes() string`

GetPoolBestActiveNodes returns the PoolBestActiveNodes field if non-nil, zero value otherwise.

### GetPoolBestActiveNodesOk

`func (o *DataInnerAppPoolData) GetPoolBestActiveNodesOk() (*string, bool)`

GetPoolBestActiveNodesOk returns a tuple with the PoolBestActiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolBestActiveNodes

`func (o *DataInnerAppPoolData) SetPoolBestActiveNodes(v string)`

SetPoolBestActiveNodes sets PoolBestActiveNodes field to given value.

### HasPoolBestActiveNodes

`func (o *DataInnerAppPoolData) HasPoolBestActiveNodes() bool`

HasPoolBestActiveNodes returns a boolean if a field has been set.

### GetPoolId

`func (o *DataInnerAppPoolData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *DataInnerAppPoolData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *DataInnerAppPoolData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *DataInnerAppPoolData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolInactiveNodes

`func (o *DataInnerAppPoolData) GetPoolInactiveNodes() string`

GetPoolInactiveNodes returns the PoolInactiveNodes field if non-nil, zero value otherwise.

### GetPoolInactiveNodesOk

`func (o *DataInnerAppPoolData) GetPoolInactiveNodesOk() (*string, bool)`

GetPoolInactiveNodesOk returns a tuple with the PoolInactiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolInactiveNodes

`func (o *DataInnerAppPoolData) SetPoolInactiveNodes(v string)`

SetPoolInactiveNodes sets PoolInactiveNodes field to given value.

### HasPoolInactiveNodes

`func (o *DataInnerAppPoolData) HasPoolInactiveNodes() bool`

HasPoolInactiveNodes returns a boolean if a field has been set.

### GetPoolLbMode

`func (o *DataInnerAppPoolData) GetPoolLbMode() string`

GetPoolLbMode returns the PoolLbMode field if non-nil, zero value otherwise.

### GetPoolLbModeOk

`func (o *DataInnerAppPoolData) GetPoolLbModeOk() (*string, bool)`

GetPoolLbModeOk returns a tuple with the PoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLbMode

`func (o *DataInnerAppPoolData) SetPoolLbMode(v string)`

SetPoolLbMode sets PoolLbMode field to given value.

### HasPoolLbMode

`func (o *DataInnerAppPoolData) HasPoolLbMode() bool`

HasPoolLbMode returns a boolean if a field has been set.

### GetPoolName

`func (o *DataInnerAppPoolData) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *DataInnerAppPoolData) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *DataInnerAppPoolData) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *DataInnerAppPoolData) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolTotalNodes

`func (o *DataInnerAppPoolData) GetPoolTotalNodes() string`

GetPoolTotalNodes returns the PoolTotalNodes field if non-nil, zero value otherwise.

### GetPoolTotalNodesOk

`func (o *DataInnerAppPoolData) GetPoolTotalNodesOk() (*string, bool)`

GetPoolTotalNodesOk returns a tuple with the PoolTotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTotalNodes

`func (o *DataInnerAppPoolData) SetPoolTotalNodes(v string)`

SetPoolTotalNodes sets PoolTotalNodes field to given value.

### HasPoolTotalNodes

`func (o *DataInnerAppPoolData) HasPoolTotalNodes() bool`

HasPoolTotalNodes returns a boolean if a field has been set.

### GetPoolType

`func (o *DataInnerAppPoolData) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *DataInnerAppPoolData) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *DataInnerAppPoolData) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *DataInnerAppPoolData) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetPoolWeight

`func (o *DataInnerAppPoolData) GetPoolWeight() string`

GetPoolWeight returns the PoolWeight field if non-nil, zero value otherwise.

### GetPoolWeightOk

`func (o *DataInnerAppPoolData) GetPoolWeightOk() (*string, bool)`

GetPoolWeightOk returns a tuple with the PoolWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolWeight

`func (o *DataInnerAppPoolData) SetPoolWeight(v string)`

SetPoolWeight sets PoolWeight field to given value.

### HasPoolWeight

`func (o *DataInnerAppPoolData) HasPoolWeight() bool`

HasPoolWeight returns a boolean if a field has been set.

### GetPoolDelayedTime

`func (o *DataInnerAppPoolData) GetPoolDelayedTime() string`

GetPoolDelayedTime returns the PoolDelayedTime field if non-nil, zero value otherwise.

### GetPoolDelayedTimeOk

`func (o *DataInnerAppPoolData) GetPoolDelayedTimeOk() (*string, bool)`

GetPoolDelayedTimeOk returns a tuple with the PoolDelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolDelayedTime

`func (o *DataInnerAppPoolData) SetPoolDelayedTime(v string)`

SetPoolDelayedTime sets PoolDelayedTime field to given value.

### HasPoolDelayedTime

`func (o *DataInnerAppPoolData) HasPoolDelayedTime() bool`

HasPoolDelayedTime returns a boolean if a field has been set.

### GetGslbserverType

`func (o *DataInnerAppPoolData) GetGslbserverType() string`

GetGslbserverType returns the GslbserverType field if non-nil, zero value otherwise.

### GetGslbserverTypeOk

`func (o *DataInnerAppPoolData) GetGslbserverTypeOk() (*string, bool)`

GetGslbserverTypeOk returns a tuple with the GslbserverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverType

`func (o *DataInnerAppPoolData) SetGslbserverType(v string)`

SetGslbserverType sets GslbserverType field to given value.

### HasGslbserverType

`func (o *DataInnerAppPoolData) HasGslbserverType() bool`

HasGslbserverType returns a boolean if a field has been set.

### GetPoolMultistatus

`func (o *DataInnerAppPoolData) GetPoolMultistatus() string`

GetPoolMultistatus returns the PoolMultistatus field if non-nil, zero value otherwise.

### GetPoolMultistatusOk

`func (o *DataInnerAppPoolData) GetPoolMultistatusOk() (*string, bool)`

GetPoolMultistatusOk returns a tuple with the PoolMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMultistatus

`func (o *DataInnerAppPoolData) SetPoolMultistatus(v string)`

SetPoolMultistatus sets PoolMultistatus field to given value.

### HasPoolMultistatus

`func (o *DataInnerAppPoolData) HasPoolMultistatus() bool`

HasPoolMultistatus returns a boolean if a field has been set.

### GetParentApplicationId

`func (o *DataInnerAppPoolData) GetParentApplicationId() string`

GetParentApplicationId returns the ParentApplicationId field if non-nil, zero value otherwise.

### GetParentApplicationIdOk

`func (o *DataInnerAppPoolData) GetParentApplicationIdOk() (*string, bool)`

GetParentApplicationIdOk returns a tuple with the ParentApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentApplicationId

`func (o *DataInnerAppPoolData) SetParentApplicationId(v string)`

SetParentApplicationId sets ParentApplicationId field to given value.

### HasParentApplicationId

`func (o *DataInnerAppPoolData) HasParentApplicationId() bool`

HasParentApplicationId returns a boolean if a field has been set.

### GetTranslatedPoolLbMode

`func (o *DataInnerAppPoolData) GetTranslatedPoolLbMode() string`

GetTranslatedPoolLbMode returns the TranslatedPoolLbMode field if non-nil, zero value otherwise.

### GetTranslatedPoolLbModeOk

`func (o *DataInnerAppPoolData) GetTranslatedPoolLbModeOk() (*string, bool)`

GetTranslatedPoolLbModeOk returns a tuple with the TranslatedPoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedPoolLbMode

`func (o *DataInnerAppPoolData) SetTranslatedPoolLbMode(v string)`

SetTranslatedPoolLbMode sets TranslatedPoolLbMode field to given value.

### HasTranslatedPoolLbMode

`func (o *DataInnerAppPoolData) HasTranslatedPoolLbMode() bool`

HasTranslatedPoolLbMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


