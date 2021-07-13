# AppNodeDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **string** | The administrative status of the node. &lt;b&gt;1&lt;/b&gt; indicates the node is managed. &lt;b&gt;0&lt;/b&gt; indicates that is unmanaged and ignored. | [optional] 
**ApplicationFqdn** | Pointer to **string** | The FQDN of the application the object belongs to. | [optional] 
**GslbserverId** | Pointer to **string** | The database identifier (ID) of the GSLB server associated with the application the object belongs to. It is only returned for deployed applications. | [optional] 
**GslbserverName** | Pointer to **string** | The name of the GSLB server associated with the application the object belongs to. It is only returned for deployed applications. | [optional] 
**ApplicationId** | Pointer to **string** | The database identifier (ID) of the application the object belongs to. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application the object belongs to. | [optional] 
**HealthcheckFailback** | Pointer to **string** | The number of times, between &lt;b&gt;1&lt;/b&gt; and &lt;b&gt;10&lt;/b&gt;, before the parameter &lt;b&gt;appnode_status&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; (active) due to the health check result. By default, it is set to &lt;b&gt;3&lt;/b&gt;. | [optional] 
**HealthcheckFailover** | Pointer to **string** | The number of times, between &lt;b&gt;1&lt;/b&gt; and &lt;b&gt;10&lt;/b&gt;, before the parameter &lt;b&gt;appnode_status&lt;/b&gt; is set to &lt;b&gt;0&lt;/b&gt; (inactive) due to the health check result. By default, it is set to &lt;b&gt;3&lt;/b&gt;. | [optional] 
**HealthcheckFreq** | Pointer to **string** | The frequency to which the health check configured for the node is performed, in seconds. | [optional] 
**HealthcheckId** | Pointer to **string** | The database identifier (ID) of the health check configured for the node. | [optional] 
**HealthcheckName** | Pointer to **string** | The type of health check configured for the node. | [optional] 
**HealthcheckParams** | Pointer to **string** | The rest of the health check parameters configured, when relevant. | [optional] 
**HealthcheckTimeout** | Pointer to **string** | The number of seconds, between &lt;b&gt;1&lt;/b&gt; and &lt;b&gt;10&lt;/b&gt;, before the health check times out if the node is not responding. | [optional] 
**NodeId** | Pointer to **string** | The database identifier (ID) of the node. | [optional] 
**NodeAddress6Addr** | Pointer to **string** | The IPv6 address of the node. | [optional] 
**NodeAddressAddr** | Pointer to **string** | The IPv4 address of the node. | [optional] 
**NodeName** | Pointer to **string** | The name of the node. | [optional] 
**NodeStatus** | Pointer to **string** | The operational status of the node: | [optional] 
**NodeWeight** | Pointer to **string** | The weight of the node, an integer between &lt;b&gt;0&lt;/b&gt; and &lt;b&gt;255&lt;/b&gt;. &lt;b&gt;0&lt;/b&gt; indicates a backup node. | [optional] 
**PoolId** | Pointer to **string** | The database identifier (ID) of the pool. | [optional] 
**PoolLbMode** | Pointer to **string** | The load-balancing mode of the pool the object belongs to, either &lt;b&gt;weighted&lt;/b&gt;, &lt;b&gt;round-robin&lt;/b&gt; or &lt;b&gt;latency&lt;/b&gt;. | [optional] 
**PoolName** | Pointer to **string** | The name of the pool. | [optional] 
**NodeDelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**GslbserverType** | Pointer to **string** | The type of DNS server associated with the application the object belongs to. It is only returned for deployed applications. | [optional] 
**ParentApplicationId** | Pointer to **string** | The database identifier (ID) of the application the object belongs to. It is only returned for deployed applications. | [optional] 
**TranslatedHealthcheckName** | Pointer to **string** | The health check type, as displayed in the GUI. | [optional] 
**TranslatedPoolLbMode** | Pointer to **string** | The load-balancing mode of the pool the object belongs to, as displayed in the GUI. | [optional] 

## Methods

### NewAppNodeDataData

`func NewAppNodeDataData() *AppNodeDataData`

NewAppNodeDataData instantiates a new AppNodeDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppNodeDataDataWithDefaults

`func NewAppNodeDataDataWithDefaults() *AppNodeDataData`

NewAppNodeDataDataWithDefaults instantiates a new AppNodeDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *AppNodeDataData) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *AppNodeDataData) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *AppNodeDataData) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *AppNodeDataData) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetApplicationFqdn

`func (o *AppNodeDataData) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *AppNodeDataData) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *AppNodeDataData) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *AppNodeDataData) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetGslbserverId

`func (o *AppNodeDataData) GetGslbserverId() string`

GetGslbserverId returns the GslbserverId field if non-nil, zero value otherwise.

### GetGslbserverIdOk

`func (o *AppNodeDataData) GetGslbserverIdOk() (*string, bool)`

GetGslbserverIdOk returns a tuple with the GslbserverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverId

`func (o *AppNodeDataData) SetGslbserverId(v string)`

SetGslbserverId sets GslbserverId field to given value.

### HasGslbserverId

`func (o *AppNodeDataData) HasGslbserverId() bool`

HasGslbserverId returns a boolean if a field has been set.

### GetGslbserverName

`func (o *AppNodeDataData) GetGslbserverName() string`

GetGslbserverName returns the GslbserverName field if non-nil, zero value otherwise.

### GetGslbserverNameOk

`func (o *AppNodeDataData) GetGslbserverNameOk() (*string, bool)`

GetGslbserverNameOk returns a tuple with the GslbserverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverName

`func (o *AppNodeDataData) SetGslbserverName(v string)`

SetGslbserverName sets GslbserverName field to given value.

### HasGslbserverName

`func (o *AppNodeDataData) HasGslbserverName() bool`

HasGslbserverName returns a boolean if a field has been set.

### GetApplicationId

`func (o *AppNodeDataData) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AppNodeDataData) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AppNodeDataData) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AppNodeDataData) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *AppNodeDataData) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AppNodeDataData) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AppNodeDataData) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AppNodeDataData) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetHealthcheckFailback

`func (o *AppNodeDataData) GetHealthcheckFailback() string`

GetHealthcheckFailback returns the HealthcheckFailback field if non-nil, zero value otherwise.

### GetHealthcheckFailbackOk

`func (o *AppNodeDataData) GetHealthcheckFailbackOk() (*string, bool)`

GetHealthcheckFailbackOk returns a tuple with the HealthcheckFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckFailback

`func (o *AppNodeDataData) SetHealthcheckFailback(v string)`

SetHealthcheckFailback sets HealthcheckFailback field to given value.

### HasHealthcheckFailback

`func (o *AppNodeDataData) HasHealthcheckFailback() bool`

HasHealthcheckFailback returns a boolean if a field has been set.

### GetHealthcheckFailover

`func (o *AppNodeDataData) GetHealthcheckFailover() string`

GetHealthcheckFailover returns the HealthcheckFailover field if non-nil, zero value otherwise.

### GetHealthcheckFailoverOk

`func (o *AppNodeDataData) GetHealthcheckFailoverOk() (*string, bool)`

GetHealthcheckFailoverOk returns a tuple with the HealthcheckFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckFailover

`func (o *AppNodeDataData) SetHealthcheckFailover(v string)`

SetHealthcheckFailover sets HealthcheckFailover field to given value.

### HasHealthcheckFailover

`func (o *AppNodeDataData) HasHealthcheckFailover() bool`

HasHealthcheckFailover returns a boolean if a field has been set.

### GetHealthcheckFreq

`func (o *AppNodeDataData) GetHealthcheckFreq() string`

GetHealthcheckFreq returns the HealthcheckFreq field if non-nil, zero value otherwise.

### GetHealthcheckFreqOk

`func (o *AppNodeDataData) GetHealthcheckFreqOk() (*string, bool)`

GetHealthcheckFreqOk returns a tuple with the HealthcheckFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckFreq

`func (o *AppNodeDataData) SetHealthcheckFreq(v string)`

SetHealthcheckFreq sets HealthcheckFreq field to given value.

### HasHealthcheckFreq

`func (o *AppNodeDataData) HasHealthcheckFreq() bool`

HasHealthcheckFreq returns a boolean if a field has been set.

### GetHealthcheckId

`func (o *AppNodeDataData) GetHealthcheckId() string`

GetHealthcheckId returns the HealthcheckId field if non-nil, zero value otherwise.

### GetHealthcheckIdOk

`func (o *AppNodeDataData) GetHealthcheckIdOk() (*string, bool)`

GetHealthcheckIdOk returns a tuple with the HealthcheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckId

`func (o *AppNodeDataData) SetHealthcheckId(v string)`

SetHealthcheckId sets HealthcheckId field to given value.

### HasHealthcheckId

`func (o *AppNodeDataData) HasHealthcheckId() bool`

HasHealthcheckId returns a boolean if a field has been set.

### GetHealthcheckName

`func (o *AppNodeDataData) GetHealthcheckName() string`

GetHealthcheckName returns the HealthcheckName field if non-nil, zero value otherwise.

### GetHealthcheckNameOk

`func (o *AppNodeDataData) GetHealthcheckNameOk() (*string, bool)`

GetHealthcheckNameOk returns a tuple with the HealthcheckName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckName

`func (o *AppNodeDataData) SetHealthcheckName(v string)`

SetHealthcheckName sets HealthcheckName field to given value.

### HasHealthcheckName

`func (o *AppNodeDataData) HasHealthcheckName() bool`

HasHealthcheckName returns a boolean if a field has been set.

### GetHealthcheckParams

`func (o *AppNodeDataData) GetHealthcheckParams() string`

GetHealthcheckParams returns the HealthcheckParams field if non-nil, zero value otherwise.

### GetHealthcheckParamsOk

`func (o *AppNodeDataData) GetHealthcheckParamsOk() (*string, bool)`

GetHealthcheckParamsOk returns a tuple with the HealthcheckParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckParams

`func (o *AppNodeDataData) SetHealthcheckParams(v string)`

SetHealthcheckParams sets HealthcheckParams field to given value.

### HasHealthcheckParams

`func (o *AppNodeDataData) HasHealthcheckParams() bool`

HasHealthcheckParams returns a boolean if a field has been set.

### GetHealthcheckTimeout

`func (o *AppNodeDataData) GetHealthcheckTimeout() string`

GetHealthcheckTimeout returns the HealthcheckTimeout field if non-nil, zero value otherwise.

### GetHealthcheckTimeoutOk

`func (o *AppNodeDataData) GetHealthcheckTimeoutOk() (*string, bool)`

GetHealthcheckTimeoutOk returns a tuple with the HealthcheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckTimeout

`func (o *AppNodeDataData) SetHealthcheckTimeout(v string)`

SetHealthcheckTimeout sets HealthcheckTimeout field to given value.

### HasHealthcheckTimeout

`func (o *AppNodeDataData) HasHealthcheckTimeout() bool`

HasHealthcheckTimeout returns a boolean if a field has been set.

### GetNodeId

`func (o *AppNodeDataData) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *AppNodeDataData) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *AppNodeDataData) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *AppNodeDataData) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeAddress6Addr

`func (o *AppNodeDataData) GetNodeAddress6Addr() string`

GetNodeAddress6Addr returns the NodeAddress6Addr field if non-nil, zero value otherwise.

### GetNodeAddress6AddrOk

`func (o *AppNodeDataData) GetNodeAddress6AddrOk() (*string, bool)`

GetNodeAddress6AddrOk returns a tuple with the NodeAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddress6Addr

`func (o *AppNodeDataData) SetNodeAddress6Addr(v string)`

SetNodeAddress6Addr sets NodeAddress6Addr field to given value.

### HasNodeAddress6Addr

`func (o *AppNodeDataData) HasNodeAddress6Addr() bool`

HasNodeAddress6Addr returns a boolean if a field has been set.

### GetNodeAddressAddr

`func (o *AppNodeDataData) GetNodeAddressAddr() string`

GetNodeAddressAddr returns the NodeAddressAddr field if non-nil, zero value otherwise.

### GetNodeAddressAddrOk

`func (o *AppNodeDataData) GetNodeAddressAddrOk() (*string, bool)`

GetNodeAddressAddrOk returns a tuple with the NodeAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddressAddr

`func (o *AppNodeDataData) SetNodeAddressAddr(v string)`

SetNodeAddressAddr sets NodeAddressAddr field to given value.

### HasNodeAddressAddr

`func (o *AppNodeDataData) HasNodeAddressAddr() bool`

HasNodeAddressAddr returns a boolean if a field has been set.

### GetNodeName

`func (o *AppNodeDataData) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *AppNodeDataData) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *AppNodeDataData) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *AppNodeDataData) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeStatus

`func (o *AppNodeDataData) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *AppNodeDataData) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *AppNodeDataData) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *AppNodeDataData) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.

### GetNodeWeight

`func (o *AppNodeDataData) GetNodeWeight() string`

GetNodeWeight returns the NodeWeight field if non-nil, zero value otherwise.

### GetNodeWeightOk

`func (o *AppNodeDataData) GetNodeWeightOk() (*string, bool)`

GetNodeWeightOk returns a tuple with the NodeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeWeight

`func (o *AppNodeDataData) SetNodeWeight(v string)`

SetNodeWeight sets NodeWeight field to given value.

### HasNodeWeight

`func (o *AppNodeDataData) HasNodeWeight() bool`

HasNodeWeight returns a boolean if a field has been set.

### GetPoolId

`func (o *AppNodeDataData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AppNodeDataData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AppNodeDataData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AppNodeDataData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolLbMode

`func (o *AppNodeDataData) GetPoolLbMode() string`

GetPoolLbMode returns the PoolLbMode field if non-nil, zero value otherwise.

### GetPoolLbModeOk

`func (o *AppNodeDataData) GetPoolLbModeOk() (*string, bool)`

GetPoolLbModeOk returns a tuple with the PoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLbMode

`func (o *AppNodeDataData) SetPoolLbMode(v string)`

SetPoolLbMode sets PoolLbMode field to given value.

### HasPoolLbMode

`func (o *AppNodeDataData) HasPoolLbMode() bool`

HasPoolLbMode returns a boolean if a field has been set.

### GetPoolName

`func (o *AppNodeDataData) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AppNodeDataData) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AppNodeDataData) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AppNodeDataData) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetNodeDelayedTime

`func (o *AppNodeDataData) GetNodeDelayedTime() string`

GetNodeDelayedTime returns the NodeDelayedTime field if non-nil, zero value otherwise.

### GetNodeDelayedTimeOk

`func (o *AppNodeDataData) GetNodeDelayedTimeOk() (*string, bool)`

GetNodeDelayedTimeOk returns a tuple with the NodeDelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDelayedTime

`func (o *AppNodeDataData) SetNodeDelayedTime(v string)`

SetNodeDelayedTime sets NodeDelayedTime field to given value.

### HasNodeDelayedTime

`func (o *AppNodeDataData) HasNodeDelayedTime() bool`

HasNodeDelayedTime returns a boolean if a field has been set.

### GetGslbserverType

`func (o *AppNodeDataData) GetGslbserverType() string`

GetGslbserverType returns the GslbserverType field if non-nil, zero value otherwise.

### GetGslbserverTypeOk

`func (o *AppNodeDataData) GetGslbserverTypeOk() (*string, bool)`

GetGslbserverTypeOk returns a tuple with the GslbserverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverType

`func (o *AppNodeDataData) SetGslbserverType(v string)`

SetGslbserverType sets GslbserverType field to given value.

### HasGslbserverType

`func (o *AppNodeDataData) HasGslbserverType() bool`

HasGslbserverType returns a boolean if a field has been set.

### GetParentApplicationId

`func (o *AppNodeDataData) GetParentApplicationId() string`

GetParentApplicationId returns the ParentApplicationId field if non-nil, zero value otherwise.

### GetParentApplicationIdOk

`func (o *AppNodeDataData) GetParentApplicationIdOk() (*string, bool)`

GetParentApplicationIdOk returns a tuple with the ParentApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentApplicationId

`func (o *AppNodeDataData) SetParentApplicationId(v string)`

SetParentApplicationId sets ParentApplicationId field to given value.

### HasParentApplicationId

`func (o *AppNodeDataData) HasParentApplicationId() bool`

HasParentApplicationId returns a boolean if a field has been set.

### GetTranslatedHealthcheckName

`func (o *AppNodeDataData) GetTranslatedHealthcheckName() string`

GetTranslatedHealthcheckName returns the TranslatedHealthcheckName field if non-nil, zero value otherwise.

### GetTranslatedHealthcheckNameOk

`func (o *AppNodeDataData) GetTranslatedHealthcheckNameOk() (*string, bool)`

GetTranslatedHealthcheckNameOk returns a tuple with the TranslatedHealthcheckName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedHealthcheckName

`func (o *AppNodeDataData) SetTranslatedHealthcheckName(v string)`

SetTranslatedHealthcheckName sets TranslatedHealthcheckName field to given value.

### HasTranslatedHealthcheckName

`func (o *AppNodeDataData) HasTranslatedHealthcheckName() bool`

HasTranslatedHealthcheckName returns a boolean if a field has been set.

### GetTranslatedPoolLbMode

`func (o *AppNodeDataData) GetTranslatedPoolLbMode() string`

GetTranslatedPoolLbMode returns the TranslatedPoolLbMode field if non-nil, zero value otherwise.

### GetTranslatedPoolLbModeOk

`func (o *AppNodeDataData) GetTranslatedPoolLbModeOk() (*string, bool)`

GetTranslatedPoolLbModeOk returns a tuple with the TranslatedPoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedPoolLbMode

`func (o *AppNodeDataData) SetTranslatedPoolLbMode(v string)`

SetTranslatedPoolLbMode sets TranslatedPoolLbMode field to given value.

### HasTranslatedPoolLbMode

`func (o *AppNodeDataData) HasTranslatedPoolLbMode() bool`

HasTranslatedPoolLbMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


