# AppPoolDataData

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
**PoolBestActiveNodes** | Pointer to **string** | The maximum number of active nodes with the lowest latency that must answer the queries made to the application FQDN. It is only returned if the parameter &lt;b&gt;apppool_lb_mode&lt;/b&gt; is set to &lt;b&gt;latency&lt;/b&gt;. | [optional] 
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

### NewAppPoolDataData

`func NewAppPoolDataData() *AppPoolDataData`

NewAppPoolDataData instantiates a new AppPoolDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPoolDataDataWithDefaults

`func NewAppPoolDataDataWithDefaults() *AppPoolDataData`

NewAppPoolDataDataWithDefaults instantiates a new AppPoolDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFqdn

`func (o *AppPoolDataData) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *AppPoolDataData) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *AppPoolDataData) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *AppPoolDataData) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetGslbserverId

`func (o *AppPoolDataData) GetGslbserverId() string`

GetGslbserverId returns the GslbserverId field if non-nil, zero value otherwise.

### GetGslbserverIdOk

`func (o *AppPoolDataData) GetGslbserverIdOk() (*string, bool)`

GetGslbserverIdOk returns a tuple with the GslbserverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverId

`func (o *AppPoolDataData) SetGslbserverId(v string)`

SetGslbserverId sets GslbserverId field to given value.

### HasGslbserverId

`func (o *AppPoolDataData) HasGslbserverId() bool`

HasGslbserverId returns a boolean if a field has been set.

### GetGslbserverName

`func (o *AppPoolDataData) GetGslbserverName() string`

GetGslbserverName returns the GslbserverName field if non-nil, zero value otherwise.

### GetGslbserverNameOk

`func (o *AppPoolDataData) GetGslbserverNameOk() (*string, bool)`

GetGslbserverNameOk returns a tuple with the GslbserverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverName

`func (o *AppPoolDataData) SetGslbserverName(v string)`

SetGslbserverName sets GslbserverName field to given value.

### HasGslbserverName

`func (o *AppPoolDataData) HasGslbserverName() bool`

HasGslbserverName returns a boolean if a field has been set.

### GetApplicationId

`func (o *AppPoolDataData) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AppPoolDataData) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AppPoolDataData) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AppPoolDataData) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *AppPoolDataData) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AppPoolDataData) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AppPoolDataData) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AppPoolDataData) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetPoolAffinitySessionTime

`func (o *AppPoolDataData) GetPoolAffinitySessionTime() string`

GetPoolAffinitySessionTime returns the PoolAffinitySessionTime field if non-nil, zero value otherwise.

### GetPoolAffinitySessionTimeOk

`func (o *AppPoolDataData) GetPoolAffinitySessionTimeOk() (*string, bool)`

GetPoolAffinitySessionTimeOk returns a tuple with the PoolAffinitySessionTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAffinitySessionTime

`func (o *AppPoolDataData) SetPoolAffinitySessionTime(v string)`

SetPoolAffinitySessionTime sets PoolAffinitySessionTime field to given value.

### HasPoolAffinitySessionTime

`func (o *AppPoolDataData) HasPoolAffinitySessionTime() bool`

HasPoolAffinitySessionTime returns a boolean if a field has been set.

### GetPoolAffinityState

`func (o *AppPoolDataData) GetPoolAffinityState() string`

GetPoolAffinityState returns the PoolAffinityState field if non-nil, zero value otherwise.

### GetPoolAffinityStateOk

`func (o *AppPoolDataData) GetPoolAffinityStateOk() (*string, bool)`

GetPoolAffinityStateOk returns a tuple with the PoolAffinityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolAffinityState

`func (o *AppPoolDataData) SetPoolAffinityState(v string)`

SetPoolAffinityState sets PoolAffinityState field to given value.

### HasPoolAffinityState

`func (o *AppPoolDataData) HasPoolAffinityState() bool`

HasPoolAffinityState returns a boolean if a field has been set.

### GetPoolBestActiveNodes

`func (o *AppPoolDataData) GetPoolBestActiveNodes() string`

GetPoolBestActiveNodes returns the PoolBestActiveNodes field if non-nil, zero value otherwise.

### GetPoolBestActiveNodesOk

`func (o *AppPoolDataData) GetPoolBestActiveNodesOk() (*string, bool)`

GetPoolBestActiveNodesOk returns a tuple with the PoolBestActiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolBestActiveNodes

`func (o *AppPoolDataData) SetPoolBestActiveNodes(v string)`

SetPoolBestActiveNodes sets PoolBestActiveNodes field to given value.

### HasPoolBestActiveNodes

`func (o *AppPoolDataData) HasPoolBestActiveNodes() bool`

HasPoolBestActiveNodes returns a boolean if a field has been set.

### GetPoolId

`func (o *AppPoolDataData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AppPoolDataData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AppPoolDataData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AppPoolDataData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolInactiveNodes

`func (o *AppPoolDataData) GetPoolInactiveNodes() string`

GetPoolInactiveNodes returns the PoolInactiveNodes field if non-nil, zero value otherwise.

### GetPoolInactiveNodesOk

`func (o *AppPoolDataData) GetPoolInactiveNodesOk() (*string, bool)`

GetPoolInactiveNodesOk returns a tuple with the PoolInactiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolInactiveNodes

`func (o *AppPoolDataData) SetPoolInactiveNodes(v string)`

SetPoolInactiveNodes sets PoolInactiveNodes field to given value.

### HasPoolInactiveNodes

`func (o *AppPoolDataData) HasPoolInactiveNodes() bool`

HasPoolInactiveNodes returns a boolean if a field has been set.

### GetPoolLbMode

`func (o *AppPoolDataData) GetPoolLbMode() string`

GetPoolLbMode returns the PoolLbMode field if non-nil, zero value otherwise.

### GetPoolLbModeOk

`func (o *AppPoolDataData) GetPoolLbModeOk() (*string, bool)`

GetPoolLbModeOk returns a tuple with the PoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLbMode

`func (o *AppPoolDataData) SetPoolLbMode(v string)`

SetPoolLbMode sets PoolLbMode field to given value.

### HasPoolLbMode

`func (o *AppPoolDataData) HasPoolLbMode() bool`

HasPoolLbMode returns a boolean if a field has been set.

### GetPoolName

`func (o *AppPoolDataData) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AppPoolDataData) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AppPoolDataData) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AppPoolDataData) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolTotalNodes

`func (o *AppPoolDataData) GetPoolTotalNodes() string`

GetPoolTotalNodes returns the PoolTotalNodes field if non-nil, zero value otherwise.

### GetPoolTotalNodesOk

`func (o *AppPoolDataData) GetPoolTotalNodesOk() (*string, bool)`

GetPoolTotalNodesOk returns a tuple with the PoolTotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolTotalNodes

`func (o *AppPoolDataData) SetPoolTotalNodes(v string)`

SetPoolTotalNodes sets PoolTotalNodes field to given value.

### HasPoolTotalNodes

`func (o *AppPoolDataData) HasPoolTotalNodes() bool`

HasPoolTotalNodes returns a boolean if a field has been set.

### GetPoolType

`func (o *AppPoolDataData) GetPoolType() string`

GetPoolType returns the PoolType field if non-nil, zero value otherwise.

### GetPoolTypeOk

`func (o *AppPoolDataData) GetPoolTypeOk() (*string, bool)`

GetPoolTypeOk returns a tuple with the PoolType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolType

`func (o *AppPoolDataData) SetPoolType(v string)`

SetPoolType sets PoolType field to given value.

### HasPoolType

`func (o *AppPoolDataData) HasPoolType() bool`

HasPoolType returns a boolean if a field has been set.

### GetPoolWeight

`func (o *AppPoolDataData) GetPoolWeight() string`

GetPoolWeight returns the PoolWeight field if non-nil, zero value otherwise.

### GetPoolWeightOk

`func (o *AppPoolDataData) GetPoolWeightOk() (*string, bool)`

GetPoolWeightOk returns a tuple with the PoolWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolWeight

`func (o *AppPoolDataData) SetPoolWeight(v string)`

SetPoolWeight sets PoolWeight field to given value.

### HasPoolWeight

`func (o *AppPoolDataData) HasPoolWeight() bool`

HasPoolWeight returns a boolean if a field has been set.

### GetPoolDelayedTime

`func (o *AppPoolDataData) GetPoolDelayedTime() string`

GetPoolDelayedTime returns the PoolDelayedTime field if non-nil, zero value otherwise.

### GetPoolDelayedTimeOk

`func (o *AppPoolDataData) GetPoolDelayedTimeOk() (*string, bool)`

GetPoolDelayedTimeOk returns a tuple with the PoolDelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolDelayedTime

`func (o *AppPoolDataData) SetPoolDelayedTime(v string)`

SetPoolDelayedTime sets PoolDelayedTime field to given value.

### HasPoolDelayedTime

`func (o *AppPoolDataData) HasPoolDelayedTime() bool`

HasPoolDelayedTime returns a boolean if a field has been set.

### GetGslbserverType

`func (o *AppPoolDataData) GetGslbserverType() string`

GetGslbserverType returns the GslbserverType field if non-nil, zero value otherwise.

### GetGslbserverTypeOk

`func (o *AppPoolDataData) GetGslbserverTypeOk() (*string, bool)`

GetGslbserverTypeOk returns a tuple with the GslbserverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverType

`func (o *AppPoolDataData) SetGslbserverType(v string)`

SetGslbserverType sets GslbserverType field to given value.

### HasGslbserverType

`func (o *AppPoolDataData) HasGslbserverType() bool`

HasGslbserverType returns a boolean if a field has been set.

### GetPoolMultistatus

`func (o *AppPoolDataData) GetPoolMultistatus() string`

GetPoolMultistatus returns the PoolMultistatus field if non-nil, zero value otherwise.

### GetPoolMultistatusOk

`func (o *AppPoolDataData) GetPoolMultistatusOk() (*string, bool)`

GetPoolMultistatusOk returns a tuple with the PoolMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMultistatus

`func (o *AppPoolDataData) SetPoolMultistatus(v string)`

SetPoolMultistatus sets PoolMultistatus field to given value.

### HasPoolMultistatus

`func (o *AppPoolDataData) HasPoolMultistatus() bool`

HasPoolMultistatus returns a boolean if a field has been set.

### GetParentApplicationId

`func (o *AppPoolDataData) GetParentApplicationId() string`

GetParentApplicationId returns the ParentApplicationId field if non-nil, zero value otherwise.

### GetParentApplicationIdOk

`func (o *AppPoolDataData) GetParentApplicationIdOk() (*string, bool)`

GetParentApplicationIdOk returns a tuple with the ParentApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentApplicationId

`func (o *AppPoolDataData) SetParentApplicationId(v string)`

SetParentApplicationId sets ParentApplicationId field to given value.

### HasParentApplicationId

`func (o *AppPoolDataData) HasParentApplicationId() bool`

HasParentApplicationId returns a boolean if a field has been set.

### GetTranslatedPoolLbMode

`func (o *AppPoolDataData) GetTranslatedPoolLbMode() string`

GetTranslatedPoolLbMode returns the TranslatedPoolLbMode field if non-nil, zero value otherwise.

### GetTranslatedPoolLbModeOk

`func (o *AppPoolDataData) GetTranslatedPoolLbModeOk() (*string, bool)`

GetTranslatedPoolLbModeOk returns a tuple with the TranslatedPoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedPoolLbMode

`func (o *AppPoolDataData) SetTranslatedPoolLbMode(v string)`

SetTranslatedPoolLbMode sets TranslatedPoolLbMode field to given value.

### HasTranslatedPoolLbMode

`func (o *AppPoolDataData) HasTranslatedPoolLbMode() bool`

HasTranslatedPoolLbMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


