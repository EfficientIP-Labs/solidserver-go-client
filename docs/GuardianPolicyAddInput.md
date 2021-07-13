# GuardianPolicyAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PolicyName** | Pointer to **string** | The name of the policy. | [optional] 
**Description** | Pointer to **string** | The description of the policy. | [optional] 
**ServerList** | Pointer to **string** | The name of all the Guardian servers associated with the policy. You can specify one or more names. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewGuardianPolicyAddInput

`func NewGuardianPolicyAddInput() *GuardianPolicyAddInput`

NewGuardianPolicyAddInput instantiates a new GuardianPolicyAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardianPolicyAddInputWithDefaults

`func NewGuardianPolicyAddInputWithDefaults() *GuardianPolicyAddInput`

NewGuardianPolicyAddInputWithDefaults instantiates a new GuardianPolicyAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPolicyName

`func (o *GuardianPolicyAddInput) GetPolicyName() string`

GetPolicyName returns the PolicyName field if non-nil, zero value otherwise.

### GetPolicyNameOk

`func (o *GuardianPolicyAddInput) GetPolicyNameOk() (*string, bool)`

GetPolicyNameOk returns a tuple with the PolicyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicyName

`func (o *GuardianPolicyAddInput) SetPolicyName(v string)`

SetPolicyName sets PolicyName field to given value.

### HasPolicyName

`func (o *GuardianPolicyAddInput) HasPolicyName() bool`

HasPolicyName returns a boolean if a field has been set.

### GetDescription

`func (o *GuardianPolicyAddInput) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuardianPolicyAddInput) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuardianPolicyAddInput) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuardianPolicyAddInput) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetServerList

`func (o *GuardianPolicyAddInput) GetServerList() string`

GetServerList returns the ServerList field if non-nil, zero value otherwise.

### GetServerListOk

`func (o *GuardianPolicyAddInput) GetServerListOk() (*string, bool)`

GetServerListOk returns a tuple with the ServerList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerList

`func (o *GuardianPolicyAddInput) SetServerList(v string)`

SetServerList sets ServerList field to given value.

### HasServerList

`func (o *GuardianPolicyAddInput) HasServerList() bool`

HasServerList returns a boolean if a field has been set.

### GetWarnings

`func (o *GuardianPolicyAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *GuardianPolicyAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *GuardianPolicyAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *GuardianPolicyAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


