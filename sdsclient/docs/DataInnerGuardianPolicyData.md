# DataInnerGuardianPolicyData

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

### NewDataInnerGuardianPolicyData

`func NewDataInnerGuardianPolicyData() *DataInnerGuardianPolicyData`

NewDataInnerGuardianPolicyData instantiates a new DataInnerGuardianPolicyData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerGuardianPolicyDataWithDefaults

`func NewDataInnerGuardianPolicyDataWithDefaults() *DataInnerGuardianPolicyData`

NewDataInnerGuardianPolicyDataWithDefaults instantiates a new DataInnerGuardianPolicyData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyDelayedTime

`func (o *DataInnerGuardianPolicyData) GetPolicyDelayedTime() string`

GetPolicyDelayedTime returns the PolicyDelayedTime field if non-nil, zero value otherwise.

### GetPolicyDelayedTimeOk

`func (o *DataInnerGuardianPolicyData) GetPolicyDelayedTimeOk() (*string, bool)`

GetPolicyDelayedTimeOk returns a tuple with the PolicyDelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDelayedTime

`func (o *DataInnerGuardianPolicyData) SetPolicyDelayedTime(v string)`

SetPolicyDelayedTime sets PolicyDelayedTime field to given value.

### HasPolicyDelayedTime

`func (o *DataInnerGuardianPolicyData) HasPolicyDelayedTime() bool`

HasPolicyDelayedTime returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerGuardianPolicyData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerGuardianPolicyData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerGuardianPolicyData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerGuardianPolicyData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerGuardianPolicyData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerGuardianPolicyData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerGuardianPolicyData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerGuardianPolicyData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetPolicyDescription

`func (o *DataInnerGuardianPolicyData) GetPolicyDescription() string`

GetPolicyDescription returns the PolicyDescription field if non-nil, zero value otherwise.

### GetPolicyDescriptionOk

`func (o *DataInnerGuardianPolicyData) GetPolicyDescriptionOk() (*string, bool)`

GetPolicyDescriptionOk returns a tuple with the PolicyDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyDescription

`func (o *DataInnerGuardianPolicyData) SetPolicyDescription(v string)`

SetPolicyDescription sets PolicyDescription field to given value.

### HasPolicyDescription

`func (o *DataInnerGuardianPolicyData) HasPolicyDescription() bool`

HasPolicyDescription returns a boolean if a field has been set.

### GetPolicyId

`func (o *DataInnerGuardianPolicyData) GetPolicyId() string`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *DataInnerGuardianPolicyData) GetPolicyIdOk() (*string, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *DataInnerGuardianPolicyData) SetPolicyId(v string)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *DataInnerGuardianPolicyData) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *DataInnerGuardianPolicyData) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *DataInnerGuardianPolicyData) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *DataInnerGuardianPolicyData) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *DataInnerGuardianPolicyData) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetPolicyReadonly

`func (o *DataInnerGuardianPolicyData) GetPolicyReadonly() string`

GetPolicyReadonly returns the PolicyReadonly field if non-nil, zero value otherwise.

### GetPolicyReadonlyOk

`func (o *DataInnerGuardianPolicyData) GetPolicyReadonlyOk() (*string, bool)`

GetPolicyReadonlyOk returns a tuple with the PolicyReadonly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyReadonly

`func (o *DataInnerGuardianPolicyData) SetPolicyReadonly(v string)`

SetPolicyReadonly sets PolicyReadonly field to given value.

### HasPolicyReadonly

`func (o *DataInnerGuardianPolicyData) HasPolicyReadonly() bool`

HasPolicyReadonly returns a boolean if a field has been set.

### GetParentPolicyId

`func (o *DataInnerGuardianPolicyData) GetParentPolicyId() string`

GetParentPolicyId returns the ParentPolicyId field if non-nil, zero value otherwise.

### GetParentPolicyIdOk

`func (o *DataInnerGuardianPolicyData) GetParentPolicyIdOk() (*string, bool)`

GetParentPolicyIdOk returns a tuple with the ParentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentPolicyId

`func (o *DataInnerGuardianPolicyData) SetParentPolicyId(v string)`

SetParentPolicyId sets ParentPolicyId field to given value.

### HasParentPolicyId

`func (o *DataInnerGuardianPolicyData) HasParentPolicyId() bool`

HasParentPolicyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


