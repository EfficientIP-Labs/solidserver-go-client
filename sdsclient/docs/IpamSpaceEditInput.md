# IpamSpaceEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space, each space must have a unique name. | [optional] 
**ParentSpaceId** | Pointer to **int32** | The database identifier (ID) of an existing space you want to set as the VLSM parent of the space you are editing. This sets up a space-based VLSM organization. | [optional] 
**ParentSpaceName** | Pointer to **string** | The name of an existing space you want to set as the VLSM parent of the space you are editing. This sets up a space-based VLSM organization. | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamSpaceEditInput

`func NewIpamSpaceEditInput() *IpamSpaceEditInput`

NewIpamSpaceEditInput instantiates a new IpamSpaceEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamSpaceEditInputWithDefaults

`func NewIpamSpaceEditInputWithDefaults() *IpamSpaceEditInput`

NewIpamSpaceEditInputWithDefaults instantiates a new IpamSpaceEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceId

`func (o *IpamSpaceEditInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamSpaceEditInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamSpaceEditInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamSpaceEditInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamSpaceEditInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamSpaceEditInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamSpaceEditInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamSpaceEditInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetParentSpaceId

`func (o *IpamSpaceEditInput) GetParentSpaceId() int32`

GetParentSpaceId returns the ParentSpaceId field if non-nil, zero value otherwise.

### GetParentSpaceIdOk

`func (o *IpamSpaceEditInput) GetParentSpaceIdOk() (*int32, bool)`

GetParentSpaceIdOk returns a tuple with the ParentSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceId

`func (o *IpamSpaceEditInput) SetParentSpaceId(v int32)`

SetParentSpaceId sets ParentSpaceId field to given value.

### HasParentSpaceId

`func (o *IpamSpaceEditInput) HasParentSpaceId() bool`

HasParentSpaceId returns a boolean if a field has been set.

### GetParentSpaceName

`func (o *IpamSpaceEditInput) GetParentSpaceName() string`

GetParentSpaceName returns the ParentSpaceName field if non-nil, zero value otherwise.

### GetParentSpaceNameOk

`func (o *IpamSpaceEditInput) GetParentSpaceNameOk() (*string, bool)`

GetParentSpaceNameOk returns a tuple with the ParentSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceName

`func (o *IpamSpaceEditInput) SetParentSpaceName(v string)`

SetParentSpaceName sets ParentSpaceName field to given value.

### HasParentSpaceName

`func (o *IpamSpaceEditInput) HasParentSpaceName() bool`

HasParentSpaceName returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *IpamSpaceEditInput) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *IpamSpaceEditInput) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *IpamSpaceEditInput) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *IpamSpaceEditInput) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *IpamSpaceEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *IpamSpaceEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *IpamSpaceEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *IpamSpaceEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *IpamSpaceEditInput) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *IpamSpaceEditInput) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *IpamSpaceEditInput) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *IpamSpaceEditInput) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *IpamSpaceEditInput) GetSpaceClassParameters() []ApiClassParameterInputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *IpamSpaceEditInput) GetSpaceClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *IpamSpaceEditInput) SetSpaceClassParameters(v []ApiClassParameterInputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *IpamSpaceEditInput) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamSpaceEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamSpaceEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamSpaceEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamSpaceEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


