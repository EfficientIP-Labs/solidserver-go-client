# DataInnerAppNodeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminStatus** | Pointer to **string** | The administrative status of the node. &lt;b&gt;1&lt;/b&gt; indicates the node is managed. &lt;b&gt;0&lt;/b&gt; indicates that is unmanaged and ignored. | [optional] 
**ApplicationFqdn** | Pointer to **string** | The FQDN of the application the object belongs to. | [optional] 
**GslbserverId** | Pointer to **string** | The database identifier (ID) of the GSLB server associated with the application the object belongs to. It is only returned for deployed applications. | [optional] 
**GslbserverName** | Pointer to **string** | The name of the GSLB server associated with the application the object belongs to. It is only returned for deployed applications. | [optional] 
**ApplicationId** | Pointer to **string** | The database identifier (ID) of the application the object belongs to. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application the object belongs to. | [optional] 
**HealthcheckFailback** | Pointer to **string** | The number of times, between &lt;b&gt;1&lt;/b&gt; and &lt;b&gt;10&lt;/b&gt;, before the parameter &lt;b&gt;node_status&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; (active) due to the health check result. By default, it is set to &lt;b&gt;3&lt;/b&gt;. | [optional] 
**HealthcheckFailover** | Pointer to **string** | The number of times, between &lt;b&gt;1&lt;/b&gt; and &lt;b&gt;10&lt;/b&gt;, before the parameter &lt;b&gt;node_status&lt;/b&gt; is set to &lt;b&gt;0&lt;/b&gt; (inactive) due to the health check result. By default, it is set to &lt;b&gt;3&lt;/b&gt;. | [optional] 
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

### NewDataInnerAppNodeData

`func NewDataInnerAppNodeData() *DataInnerAppNodeData`

NewDataInnerAppNodeData instantiates a new DataInnerAppNodeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerAppNodeDataWithDefaults

`func NewDataInnerAppNodeDataWithDefaults() *DataInnerAppNodeData`

NewDataInnerAppNodeDataWithDefaults instantiates a new DataInnerAppNodeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminStatus

`func (o *DataInnerAppNodeData) GetAdminStatus() string`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *DataInnerAppNodeData) GetAdminStatusOk() (*string, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *DataInnerAppNodeData) SetAdminStatus(v string)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *DataInnerAppNodeData) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetApplicationFqdn

`func (o *DataInnerAppNodeData) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *DataInnerAppNodeData) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *DataInnerAppNodeData) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *DataInnerAppNodeData) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetGslbserverId

`func (o *DataInnerAppNodeData) GetGslbserverId() string`

GetGslbserverId returns the GslbserverId field if non-nil, zero value otherwise.

### GetGslbserverIdOk

`func (o *DataInnerAppNodeData) GetGslbserverIdOk() (*string, bool)`

GetGslbserverIdOk returns a tuple with the GslbserverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverId

`func (o *DataInnerAppNodeData) SetGslbserverId(v string)`

SetGslbserverId sets GslbserverId field to given value.

### HasGslbserverId

`func (o *DataInnerAppNodeData) HasGslbserverId() bool`

HasGslbserverId returns a boolean if a field has been set.

### GetGslbserverName

`func (o *DataInnerAppNodeData) GetGslbserverName() string`

GetGslbserverName returns the GslbserverName field if non-nil, zero value otherwise.

### GetGslbserverNameOk

`func (o *DataInnerAppNodeData) GetGslbserverNameOk() (*string, bool)`

GetGslbserverNameOk returns a tuple with the GslbserverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverName

`func (o *DataInnerAppNodeData) SetGslbserverName(v string)`

SetGslbserverName sets GslbserverName field to given value.

### HasGslbserverName

`func (o *DataInnerAppNodeData) HasGslbserverName() bool`

HasGslbserverName returns a boolean if a field has been set.

### GetApplicationId

`func (o *DataInnerAppNodeData) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *DataInnerAppNodeData) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *DataInnerAppNodeData) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *DataInnerAppNodeData) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *DataInnerAppNodeData) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *DataInnerAppNodeData) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *DataInnerAppNodeData) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *DataInnerAppNodeData) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetHealthcheckFailback

`func (o *DataInnerAppNodeData) GetHealthcheckFailback() string`

GetHealthcheckFailback returns the HealthcheckFailback field if non-nil, zero value otherwise.

### GetHealthcheckFailbackOk

`func (o *DataInnerAppNodeData) GetHealthcheckFailbackOk() (*string, bool)`

GetHealthcheckFailbackOk returns a tuple with the HealthcheckFailback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckFailback

`func (o *DataInnerAppNodeData) SetHealthcheckFailback(v string)`

SetHealthcheckFailback sets HealthcheckFailback field to given value.

### HasHealthcheckFailback

`func (o *DataInnerAppNodeData) HasHealthcheckFailback() bool`

HasHealthcheckFailback returns a boolean if a field has been set.

### GetHealthcheckFailover

`func (o *DataInnerAppNodeData) GetHealthcheckFailover() string`

GetHealthcheckFailover returns the HealthcheckFailover field if non-nil, zero value otherwise.

### GetHealthcheckFailoverOk

`func (o *DataInnerAppNodeData) GetHealthcheckFailoverOk() (*string, bool)`

GetHealthcheckFailoverOk returns a tuple with the HealthcheckFailover field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckFailover

`func (o *DataInnerAppNodeData) SetHealthcheckFailover(v string)`

SetHealthcheckFailover sets HealthcheckFailover field to given value.

### HasHealthcheckFailover

`func (o *DataInnerAppNodeData) HasHealthcheckFailover() bool`

HasHealthcheckFailover returns a boolean if a field has been set.

### GetHealthcheckFreq

`func (o *DataInnerAppNodeData) GetHealthcheckFreq() string`

GetHealthcheckFreq returns the HealthcheckFreq field if non-nil, zero value otherwise.

### GetHealthcheckFreqOk

`func (o *DataInnerAppNodeData) GetHealthcheckFreqOk() (*string, bool)`

GetHealthcheckFreqOk returns a tuple with the HealthcheckFreq field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckFreq

`func (o *DataInnerAppNodeData) SetHealthcheckFreq(v string)`

SetHealthcheckFreq sets HealthcheckFreq field to given value.

### HasHealthcheckFreq

`func (o *DataInnerAppNodeData) HasHealthcheckFreq() bool`

HasHealthcheckFreq returns a boolean if a field has been set.

### GetHealthcheckId

`func (o *DataInnerAppNodeData) GetHealthcheckId() string`

GetHealthcheckId returns the HealthcheckId field if non-nil, zero value otherwise.

### GetHealthcheckIdOk

`func (o *DataInnerAppNodeData) GetHealthcheckIdOk() (*string, bool)`

GetHealthcheckIdOk returns a tuple with the HealthcheckId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckId

`func (o *DataInnerAppNodeData) SetHealthcheckId(v string)`

SetHealthcheckId sets HealthcheckId field to given value.

### HasHealthcheckId

`func (o *DataInnerAppNodeData) HasHealthcheckId() bool`

HasHealthcheckId returns a boolean if a field has been set.

### GetHealthcheckName

`func (o *DataInnerAppNodeData) GetHealthcheckName() string`

GetHealthcheckName returns the HealthcheckName field if non-nil, zero value otherwise.

### GetHealthcheckNameOk

`func (o *DataInnerAppNodeData) GetHealthcheckNameOk() (*string, bool)`

GetHealthcheckNameOk returns a tuple with the HealthcheckName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckName

`func (o *DataInnerAppNodeData) SetHealthcheckName(v string)`

SetHealthcheckName sets HealthcheckName field to given value.

### HasHealthcheckName

`func (o *DataInnerAppNodeData) HasHealthcheckName() bool`

HasHealthcheckName returns a boolean if a field has been set.

### GetHealthcheckParams

`func (o *DataInnerAppNodeData) GetHealthcheckParams() string`

GetHealthcheckParams returns the HealthcheckParams field if non-nil, zero value otherwise.

### GetHealthcheckParamsOk

`func (o *DataInnerAppNodeData) GetHealthcheckParamsOk() (*string, bool)`

GetHealthcheckParamsOk returns a tuple with the HealthcheckParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckParams

`func (o *DataInnerAppNodeData) SetHealthcheckParams(v string)`

SetHealthcheckParams sets HealthcheckParams field to given value.

### HasHealthcheckParams

`func (o *DataInnerAppNodeData) HasHealthcheckParams() bool`

HasHealthcheckParams returns a boolean if a field has been set.

### GetHealthcheckTimeout

`func (o *DataInnerAppNodeData) GetHealthcheckTimeout() string`

GetHealthcheckTimeout returns the HealthcheckTimeout field if non-nil, zero value otherwise.

### GetHealthcheckTimeoutOk

`func (o *DataInnerAppNodeData) GetHealthcheckTimeoutOk() (*string, bool)`

GetHealthcheckTimeoutOk returns a tuple with the HealthcheckTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckTimeout

`func (o *DataInnerAppNodeData) SetHealthcheckTimeout(v string)`

SetHealthcheckTimeout sets HealthcheckTimeout field to given value.

### HasHealthcheckTimeout

`func (o *DataInnerAppNodeData) HasHealthcheckTimeout() bool`

HasHealthcheckTimeout returns a boolean if a field has been set.

### GetNodeId

`func (o *DataInnerAppNodeData) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DataInnerAppNodeData) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DataInnerAppNodeData) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *DataInnerAppNodeData) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetNodeAddress6Addr

`func (o *DataInnerAppNodeData) GetNodeAddress6Addr() string`

GetNodeAddress6Addr returns the NodeAddress6Addr field if non-nil, zero value otherwise.

### GetNodeAddress6AddrOk

`func (o *DataInnerAppNodeData) GetNodeAddress6AddrOk() (*string, bool)`

GetNodeAddress6AddrOk returns a tuple with the NodeAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddress6Addr

`func (o *DataInnerAppNodeData) SetNodeAddress6Addr(v string)`

SetNodeAddress6Addr sets NodeAddress6Addr field to given value.

### HasNodeAddress6Addr

`func (o *DataInnerAppNodeData) HasNodeAddress6Addr() bool`

HasNodeAddress6Addr returns a boolean if a field has been set.

### GetNodeAddressAddr

`func (o *DataInnerAppNodeData) GetNodeAddressAddr() string`

GetNodeAddressAddr returns the NodeAddressAddr field if non-nil, zero value otherwise.

### GetNodeAddressAddrOk

`func (o *DataInnerAppNodeData) GetNodeAddressAddrOk() (*string, bool)`

GetNodeAddressAddrOk returns a tuple with the NodeAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeAddressAddr

`func (o *DataInnerAppNodeData) SetNodeAddressAddr(v string)`

SetNodeAddressAddr sets NodeAddressAddr field to given value.

### HasNodeAddressAddr

`func (o *DataInnerAppNodeData) HasNodeAddressAddr() bool`

HasNodeAddressAddr returns a boolean if a field has been set.

### GetNodeName

`func (o *DataInnerAppNodeData) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *DataInnerAppNodeData) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *DataInnerAppNodeData) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *DataInnerAppNodeData) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetNodeStatus

`func (o *DataInnerAppNodeData) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *DataInnerAppNodeData) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *DataInnerAppNodeData) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *DataInnerAppNodeData) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.

### GetNodeWeight

`func (o *DataInnerAppNodeData) GetNodeWeight() string`

GetNodeWeight returns the NodeWeight field if non-nil, zero value otherwise.

### GetNodeWeightOk

`func (o *DataInnerAppNodeData) GetNodeWeightOk() (*string, bool)`

GetNodeWeightOk returns a tuple with the NodeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeWeight

`func (o *DataInnerAppNodeData) SetNodeWeight(v string)`

SetNodeWeight sets NodeWeight field to given value.

### HasNodeWeight

`func (o *DataInnerAppNodeData) HasNodeWeight() bool`

HasNodeWeight returns a boolean if a field has been set.

### GetPoolId

`func (o *DataInnerAppNodeData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *DataInnerAppNodeData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *DataInnerAppNodeData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *DataInnerAppNodeData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolLbMode

`func (o *DataInnerAppNodeData) GetPoolLbMode() string`

GetPoolLbMode returns the PoolLbMode field if non-nil, zero value otherwise.

### GetPoolLbModeOk

`func (o *DataInnerAppNodeData) GetPoolLbModeOk() (*string, bool)`

GetPoolLbModeOk returns a tuple with the PoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolLbMode

`func (o *DataInnerAppNodeData) SetPoolLbMode(v string)`

SetPoolLbMode sets PoolLbMode field to given value.

### HasPoolLbMode

`func (o *DataInnerAppNodeData) HasPoolLbMode() bool`

HasPoolLbMode returns a boolean if a field has been set.

### GetPoolName

`func (o *DataInnerAppNodeData) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *DataInnerAppNodeData) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *DataInnerAppNodeData) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *DataInnerAppNodeData) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetNodeDelayedTime

`func (o *DataInnerAppNodeData) GetNodeDelayedTime() string`

GetNodeDelayedTime returns the NodeDelayedTime field if non-nil, zero value otherwise.

### GetNodeDelayedTimeOk

`func (o *DataInnerAppNodeData) GetNodeDelayedTimeOk() (*string, bool)`

GetNodeDelayedTimeOk returns a tuple with the NodeDelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeDelayedTime

`func (o *DataInnerAppNodeData) SetNodeDelayedTime(v string)`

SetNodeDelayedTime sets NodeDelayedTime field to given value.

### HasNodeDelayedTime

`func (o *DataInnerAppNodeData) HasNodeDelayedTime() bool`

HasNodeDelayedTime returns a boolean if a field has been set.

### GetGslbserverType

`func (o *DataInnerAppNodeData) GetGslbserverType() string`

GetGslbserverType returns the GslbserverType field if non-nil, zero value otherwise.

### GetGslbserverTypeOk

`func (o *DataInnerAppNodeData) GetGslbserverTypeOk() (*string, bool)`

GetGslbserverTypeOk returns a tuple with the GslbserverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverType

`func (o *DataInnerAppNodeData) SetGslbserverType(v string)`

SetGslbserverType sets GslbserverType field to given value.

### HasGslbserverType

`func (o *DataInnerAppNodeData) HasGslbserverType() bool`

HasGslbserverType returns a boolean if a field has been set.

### GetParentApplicationId

`func (o *DataInnerAppNodeData) GetParentApplicationId() string`

GetParentApplicationId returns the ParentApplicationId field if non-nil, zero value otherwise.

### GetParentApplicationIdOk

`func (o *DataInnerAppNodeData) GetParentApplicationIdOk() (*string, bool)`

GetParentApplicationIdOk returns a tuple with the ParentApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentApplicationId

`func (o *DataInnerAppNodeData) SetParentApplicationId(v string)`

SetParentApplicationId sets ParentApplicationId field to given value.

### HasParentApplicationId

`func (o *DataInnerAppNodeData) HasParentApplicationId() bool`

HasParentApplicationId returns a boolean if a field has been set.

### GetTranslatedHealthcheckName

`func (o *DataInnerAppNodeData) GetTranslatedHealthcheckName() string`

GetTranslatedHealthcheckName returns the TranslatedHealthcheckName field if non-nil, zero value otherwise.

### GetTranslatedHealthcheckNameOk

`func (o *DataInnerAppNodeData) GetTranslatedHealthcheckNameOk() (*string, bool)`

GetTranslatedHealthcheckNameOk returns a tuple with the TranslatedHealthcheckName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedHealthcheckName

`func (o *DataInnerAppNodeData) SetTranslatedHealthcheckName(v string)`

SetTranslatedHealthcheckName sets TranslatedHealthcheckName field to given value.

### HasTranslatedHealthcheckName

`func (o *DataInnerAppNodeData) HasTranslatedHealthcheckName() bool`

HasTranslatedHealthcheckName returns a boolean if a field has been set.

### GetTranslatedPoolLbMode

`func (o *DataInnerAppNodeData) GetTranslatedPoolLbMode() string`

GetTranslatedPoolLbMode returns the TranslatedPoolLbMode field if non-nil, zero value otherwise.

### GetTranslatedPoolLbModeOk

`func (o *DataInnerAppNodeData) GetTranslatedPoolLbModeOk() (*string, bool)`

GetTranslatedPoolLbModeOk returns a tuple with the TranslatedPoolLbMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslatedPoolLbMode

`func (o *DataInnerAppNodeData) SetTranslatedPoolLbMode(v string)`

SetTranslatedPoolLbMode sets TranslatedPoolLbMode field to given value.

### HasTranslatedPoolLbMode

`func (o *DataInnerAppNodeData) HasTranslatedPoolLbMode() bool`

HasTranslatedPoolLbMode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


