# AppNodeAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFqdn** | Pointer to **string** | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. | [optional] 
**ApplicationId** | Pointer to **int32** | The database identifier (ID) of the application the object belongs to. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application the object belongs to. | [optional] 
**PoolId** | Pointer to **int32** | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. | [optional] 
**PoolName** | Pointer to **string** | The name of the pool. | [optional] 
**NodeHostaddr** | Pointer to **string** | The IPv4 or IPv6 address of the node. | [optional] 
**NodeName** | Pointer to **string** | The name of the node. | [optional] 
**AdminStatus** | Pointer to **int32** | The administrative status of the node, managed (1) or unmanaged (0). | [optional] 
**GslbserverId** | Pointer to **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique identifier automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | [optional] 
**HealthcheckName** | Pointer to **string** | The type of health check of the node. | [optional] 
**NodeStatus** | Pointer to **string** | Internal use. Not documented. | [optional] 
**NodeWeight** | Pointer to **string** | The weight of the node, it sets which node is used first within the pool. It must be an integer between 0 and 255, where 0 sets a node as backup. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewAppNodeAddInput

`func NewAppNodeAddInput() *AppNodeAddInput`

NewAppNodeAddInput instantiates a new AppNodeAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppNodeAddInputWithDefaults

`func NewAppNodeAddInputWithDefaults() *AppNodeAddInput`

NewAppNodeAddInputWithDefaults instantiates a new AppNodeAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFqdn

`func (o *AppNodeAddInput) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *AppNodeAddInput) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *AppNodeAddInput) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *AppNodeAddInput) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetApplicationId

`func (o *AppNodeAddInput) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AppNodeAddInput) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AppNodeAddInput) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AppNodeAddInput) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationName

`func (o *AppNodeAddInput) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AppNodeAddInput) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AppNodeAddInput) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AppNodeAddInput) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetPoolId

`func (o *AppNodeAddInput) GetPoolId() int32`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *AppNodeAddInput) GetPoolIdOk() (*int32, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *AppNodeAddInput) SetPoolId(v int32)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *AppNodeAddInput) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *AppNodeAddInput) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *AppNodeAddInput) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *AppNodeAddInput) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *AppNodeAddInput) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetNodeHostaddr

`func (o *AppNodeAddInput) GetNodeHostaddr() string`

GetNodeHostaddr returns the NodeHostaddr field if non-nil, zero value otherwise.

### GetNodeHostaddrOk

`func (o *AppNodeAddInput) GetNodeHostaddrOk() (*string, bool)`

GetNodeHostaddrOk returns a tuple with the NodeHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeHostaddr

`func (o *AppNodeAddInput) SetNodeHostaddr(v string)`

SetNodeHostaddr sets NodeHostaddr field to given value.

### HasNodeHostaddr

`func (o *AppNodeAddInput) HasNodeHostaddr() bool`

HasNodeHostaddr returns a boolean if a field has been set.

### GetNodeName

`func (o *AppNodeAddInput) GetNodeName() string`

GetNodeName returns the NodeName field if non-nil, zero value otherwise.

### GetNodeNameOk

`func (o *AppNodeAddInput) GetNodeNameOk() (*string, bool)`

GetNodeNameOk returns a tuple with the NodeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeName

`func (o *AppNodeAddInput) SetNodeName(v string)`

SetNodeName sets NodeName field to given value.

### HasNodeName

`func (o *AppNodeAddInput) HasNodeName() bool`

HasNodeName returns a boolean if a field has been set.

### GetAdminStatus

`func (o *AppNodeAddInput) GetAdminStatus() int32`

GetAdminStatus returns the AdminStatus field if non-nil, zero value otherwise.

### GetAdminStatusOk

`func (o *AppNodeAddInput) GetAdminStatusOk() (*int32, bool)`

GetAdminStatusOk returns a tuple with the AdminStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminStatus

`func (o *AppNodeAddInput) SetAdminStatus(v int32)`

SetAdminStatus sets AdminStatus field to given value.

### HasAdminStatus

`func (o *AppNodeAddInput) HasAdminStatus() bool`

HasAdminStatus returns a boolean if a field has been set.

### GetGslbserverId

`func (o *AppNodeAddInput) GetGslbserverId() int32`

GetGslbserverId returns the GslbserverId field if non-nil, zero value otherwise.

### GetGslbserverIdOk

`func (o *AppNodeAddInput) GetGslbserverIdOk() (*int32, bool)`

GetGslbserverIdOk returns a tuple with the GslbserverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverId

`func (o *AppNodeAddInput) SetGslbserverId(v int32)`

SetGslbserverId sets GslbserverId field to given value.

### HasGslbserverId

`func (o *AppNodeAddInput) HasGslbserverId() bool`

HasGslbserverId returns a boolean if a field has been set.

### GetHealthcheckName

`func (o *AppNodeAddInput) GetHealthcheckName() string`

GetHealthcheckName returns the HealthcheckName field if non-nil, zero value otherwise.

### GetHealthcheckNameOk

`func (o *AppNodeAddInput) GetHealthcheckNameOk() (*string, bool)`

GetHealthcheckNameOk returns a tuple with the HealthcheckName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheckName

`func (o *AppNodeAddInput) SetHealthcheckName(v string)`

SetHealthcheckName sets HealthcheckName field to given value.

### HasHealthcheckName

`func (o *AppNodeAddInput) HasHealthcheckName() bool`

HasHealthcheckName returns a boolean if a field has been set.

### GetNodeStatus

`func (o *AppNodeAddInput) GetNodeStatus() string`

GetNodeStatus returns the NodeStatus field if non-nil, zero value otherwise.

### GetNodeStatusOk

`func (o *AppNodeAddInput) GetNodeStatusOk() (*string, bool)`

GetNodeStatusOk returns a tuple with the NodeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeStatus

`func (o *AppNodeAddInput) SetNodeStatus(v string)`

SetNodeStatus sets NodeStatus field to given value.

### HasNodeStatus

`func (o *AppNodeAddInput) HasNodeStatus() bool`

HasNodeStatus returns a boolean if a field has been set.

### GetNodeWeight

`func (o *AppNodeAddInput) GetNodeWeight() string`

GetNodeWeight returns the NodeWeight field if non-nil, zero value otherwise.

### GetNodeWeightOk

`func (o *AppNodeAddInput) GetNodeWeightOk() (*string, bool)`

GetNodeWeightOk returns a tuple with the NodeWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeWeight

`func (o *AppNodeAddInput) SetNodeWeight(v string)`

SetNodeWeight sets NodeWeight field to given value.

### HasNodeWeight

`func (o *AppNodeAddInput) HasNodeWeight() bool`

HasNodeWeight returns a boolean if a field has been set.

### GetWarnings

`func (o *AppNodeAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AppNodeAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AppNodeAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AppNodeAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


