# GuardianPolicyDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyDelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the Guardian server associated with the policy. It is only returned for deployed policies. | [optional] 
**ServerName** | Pointer to **string** | The name of the Guardian server associated with the policy. It is only returned for deployed policies. | [optional] 
**PolicyDescription** | Pointer to **string** | The description of the policy. | [optional] 
**PolicyId** | Pointer to **string** | The database identifier (ID) of the policy. | [optional] 
**PolicyName** | Pointer to **string** | The name of the policy. | [optional] 
**PolicyReadonly** | Pointer to **string** | The read only status of the policy. If set to &lt;b&gt;1&lt;/b&gt;, the policy cannot be edited. | [optional] 
**ParentPolicyId** | Pointer to **string** | The database identifier (ID) of the policy. It is only returned for deployed policies. | [optional] 

## Methods

### NewGuardianPolicyDataData

`func NewGuardianPolicyDataData() *GuardianPolicyDataData`

NewGuardianPolicyDataData instantiates a new GuardianPolicyDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardianPolicyDataDataWithDefaults

`func NewGuardianPolicyDataDataWithDefaults() *GuardianPolicyDataData`

NewGuardianPolicyDataDataWithDefaults instantiates a new GuardianPolicyDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyDelayedTime

`func (o *GuardianPolicyDataData) GetPolicyDelayedTime() string`

GetPolicyDelayedTime returns the PolicyDelayedTime field if non-nil, zero value otherwise.

### GetPolicyDelayedTimeOk

`func (o *GuardianPolicyDataData) GetPolicyDelayedTimeOk() (*string, bool)`

GetPolicyDelayedTimeOk returns a tuple with the PolicyDelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDelayedTime

`func (o *GuardianPolicyDataData) SetPolicyDelayedTime(v string)`

SetPolicyDelayedTime sets PolicyDelayedTime field to given value.

### HasPolicyDelayedTime

`func (o *GuardianPolicyDataData) HasPolicyDelayedTime() bool`

HasPolicyDelayedTime returns a boolean if a field has been set.

### GetServerId

`func (o *GuardianPolicyDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *GuardianPolicyDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *GuardianPolicyDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *GuardianPolicyDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *GuardianPolicyDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *GuardianPolicyDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *GuardianPolicyDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *GuardianPolicyDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetPolicyDescription

`func (o *GuardianPolicyDataData) GetPolicyDescription() string`

GetPolicyDescription returns the PolicyDescription field if non-nil, zero value otherwise.

### GetPolicyDescriptionOk

`func (o *GuardianPolicyDataData) GetPolicyDescriptionOk() (*string, bool)`

GetPolicyDescriptionOk returns a tuple with the PolicyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDescription

`func (o *GuardianPolicyDataData) SetPolicyDescription(v string)`

SetPolicyDescription sets PolicyDescription field to given value.

### HasPolicyDescription

`func (o *GuardianPolicyDataData) HasPolicyDescription() bool`

HasPolicyDescription returns a boolean if a field has been set.

### GetPolicyId

`func (o *GuardianPolicyDataData) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *GuardianPolicyDataData) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *GuardianPolicyDataData) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *GuardianPolicyDataData) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *GuardianPolicyDataData) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *GuardianPolicyDataData) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *GuardianPolicyDataData) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *GuardianPolicyDataData) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyReadonly

`func (o *GuardianPolicyDataData) GetPolicyReadonly() string`

GetPolicyReadonly returns the PolicyReadonly field if non-nil, zero value otherwise.

### GetPolicyReadonlyOk

`func (o *GuardianPolicyDataData) GetPolicyReadonlyOk() (*string, bool)`

GetPolicyReadonlyOk returns a tuple with the PolicyReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReadonly

`func (o *GuardianPolicyDataData) SetPolicyReadonly(v string)`

SetPolicyReadonly sets PolicyReadonly field to given value.

### HasPolicyReadonly

`func (o *GuardianPolicyDataData) HasPolicyReadonly() bool`

HasPolicyReadonly returns a boolean if a field has been set.

### GetParentPolicyId

`func (o *GuardianPolicyDataData) GetParentPolicyId() string`

GetParentPolicyId returns the ParentPolicyId field if non-nil, zero value otherwise.

### GetParentPolicyIdOk

`func (o *GuardianPolicyDataData) GetParentPolicyIdOk() (*string, bool)`

GetParentPolicyIdOk returns a tuple with the ParentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPolicyId

`func (o *GuardianPolicyDataData) SetParentPolicyId(v string)`

SetParentPolicyId sets ParentPolicyId field to given value.

### HasParentPolicyId

`func (o *GuardianPolicyDataData) HasParentPolicyId() bool`

HasParentPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


