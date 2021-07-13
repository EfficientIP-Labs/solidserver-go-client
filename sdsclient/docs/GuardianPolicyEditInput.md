# GuardianPolicyEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyId** | Pointer to **int32** | The database identifier (ID) of the policy, a unique identifier automatically incremented when you add the policy. Use it to identify the policy of your choice. | [optional] 
**PolicyName** | Pointer to **string** | The name of the policy. | [optional] 
**Description** | Pointer to **string** | The description of the policy. | [optional] 
**ServerList** | Pointer to **string** | The name of all the Guardian servers associated with the policy. You can specify one or more names. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewGuardianPolicyEditInput

`func NewGuardianPolicyEditInput() *GuardianPolicyEditInput`

NewGuardianPolicyEditInput instantiates a new GuardianPolicyEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardianPolicyEditInputWithDefaults

`func NewGuardianPolicyEditInputWithDefaults() *GuardianPolicyEditInput`

NewGuardianPolicyEditInputWithDefaults instantiates a new GuardianPolicyEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyId

`func (o *GuardianPolicyEditInput) GetPolicyId() int32`

GetPolicyId returns the PolicyId field if non-nil, zero value otherwise.

### GetPolicyIdOk

`func (o *GuardianPolicyEditInput) GetPolicyIdOk() (*int32, bool)`

GetPolicyIdOk returns a tuple with the PolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyId

`func (o *GuardianPolicyEditInput) SetPolicyId(v int32)`

SetPolicyId sets PolicyId field to given value.

### HasPolicyId

`func (o *GuardianPolicyEditInput) HasPolicyId() bool`

HasPolicyId returns a boolean if a field has been set.

### GetPolicyName

`func (o *GuardianPolicyEditInput) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *GuardianPolicyEditInput) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *GuardianPolicyEditInput) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *GuardianPolicyEditInput) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetDescription

`func (o *GuardianPolicyEditInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuardianPolicyEditInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuardianPolicyEditInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuardianPolicyEditInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServerList

`func (o *GuardianPolicyEditInput) GetServerList() string`

GetServerList returns the ServerList field if non-nil, zero value otherwise.

### GetServerListOk

`func (o *GuardianPolicyEditInput) GetServerListOk() (*string, bool)`

GetServerListOk returns a tuple with the ServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerList

`func (o *GuardianPolicyEditInput) SetServerList(v string)`

SetServerList sets ServerList field to given value.

### HasServerList

`func (o *GuardianPolicyEditInput) HasServerList() bool`

HasServerList returns a boolean if a field has been set.

### GetWarnings

`func (o *GuardianPolicyEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *GuardianPolicyEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *GuardianPolicyEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *GuardianPolicyEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


