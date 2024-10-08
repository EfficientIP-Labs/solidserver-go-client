# VlanDomainEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | Pointer to **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain, each VLAN domain must have a unique name. | [optional] 
**DomainSupportVxlan** | Pointer to **int32** | The type of virtual network used by the domain. Set it to &lt;b&gt;1&lt;/b&gt; to use VXLAN or &lt;b&gt;0&lt;/b&gt; to use VLAN. | [optional] 
**DomainDescription** | Pointer to **string** | The description of the VLAN domain. | [optional] 
**DomainEndVlanId** | Pointer to **int32** | The VLAN identifier (ID) of a VLAN, a numeric value between 1 and 4094. Use the ID to specify the last VLAN in the VLAN domain. | [optional] 
**DomainStartVlanId** | Pointer to **int32** | The VLAN identifier (ID) of a VLAN, a numeric value between 1 and 4094. Use the ID to specify the first VLAN in the VLAN domain. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**DomainClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**DomainClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewVlanDomainEditInput

`func NewVlanDomainEditInput() *VlanDomainEditInput`

NewVlanDomainEditInput instantiates a new VlanDomainEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanDomainEditInputWithDefaults

`func NewVlanDomainEditInputWithDefaults() *VlanDomainEditInput`

NewVlanDomainEditInputWithDefaults instantiates a new VlanDomainEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *VlanDomainEditInput) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *VlanDomainEditInput) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *VlanDomainEditInput) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *VlanDomainEditInput) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *VlanDomainEditInput) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VlanDomainEditInput) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VlanDomainEditInput) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *VlanDomainEditInput) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainSupportVxlan

`func (o *VlanDomainEditInput) GetDomainSupportVxlan() int32`

GetDomainSupportVxlan returns the DomainSupportVxlan field if non-nil, zero value otherwise.

### GetDomainSupportVxlanOk

`func (o *VlanDomainEditInput) GetDomainSupportVxlanOk() (*int32, bool)`

GetDomainSupportVxlanOk returns a tuple with the DomainSupportVxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSupportVxlan

`func (o *VlanDomainEditInput) SetDomainSupportVxlan(v int32)`

SetDomainSupportVxlan sets DomainSupportVxlan field to given value.

### HasDomainSupportVxlan

`func (o *VlanDomainEditInput) HasDomainSupportVxlan() bool`

HasDomainSupportVxlan returns a boolean if a field has been set.

### GetDomainDescription

`func (o *VlanDomainEditInput) GetDomainDescription() string`

GetDomainDescription returns the DomainDescription field if non-nil, zero value otherwise.

### GetDomainDescriptionOk

`func (o *VlanDomainEditInput) GetDomainDescriptionOk() (*string, bool)`

GetDomainDescriptionOk returns a tuple with the DomainDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDescription

`func (o *VlanDomainEditInput) SetDomainDescription(v string)`

SetDomainDescription sets DomainDescription field to given value.

### HasDomainDescription

`func (o *VlanDomainEditInput) HasDomainDescription() bool`

HasDomainDescription returns a boolean if a field has been set.

### GetDomainEndVlanId

`func (o *VlanDomainEditInput) GetDomainEndVlanId() int32`

GetDomainEndVlanId returns the DomainEndVlanId field if non-nil, zero value otherwise.

### GetDomainEndVlanIdOk

`func (o *VlanDomainEditInput) GetDomainEndVlanIdOk() (*int32, bool)`

GetDomainEndVlanIdOk returns a tuple with the DomainEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEndVlanId

`func (o *VlanDomainEditInput) SetDomainEndVlanId(v int32)`

SetDomainEndVlanId sets DomainEndVlanId field to given value.

### HasDomainEndVlanId

`func (o *VlanDomainEditInput) HasDomainEndVlanId() bool`

HasDomainEndVlanId returns a boolean if a field has been set.

### GetDomainStartVlanId

`func (o *VlanDomainEditInput) GetDomainStartVlanId() int32`

GetDomainStartVlanId returns the DomainStartVlanId field if non-nil, zero value otherwise.

### GetDomainStartVlanIdOk

`func (o *VlanDomainEditInput) GetDomainStartVlanIdOk() (*int32, bool)`

GetDomainStartVlanIdOk returns a tuple with the DomainStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStartVlanId

`func (o *VlanDomainEditInput) SetDomainStartVlanId(v int32)`

SetDomainStartVlanId sets DomainStartVlanId field to given value.

### HasDomainStartVlanId

`func (o *VlanDomainEditInput) HasDomainStartVlanId() bool`

HasDomainStartVlanId returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *VlanDomainEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *VlanDomainEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *VlanDomainEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *VlanDomainEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetDomainClassName

`func (o *VlanDomainEditInput) GetDomainClassName() string`

GetDomainClassName returns the DomainClassName field if non-nil, zero value otherwise.

### GetDomainClassNameOk

`func (o *VlanDomainEditInput) GetDomainClassNameOk() (*string, bool)`

GetDomainClassNameOk returns a tuple with the DomainClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassName

`func (o *VlanDomainEditInput) SetDomainClassName(v string)`

SetDomainClassName sets DomainClassName field to given value.

### HasDomainClassName

`func (o *VlanDomainEditInput) HasDomainClassName() bool`

HasDomainClassName returns a boolean if a field has been set.

### GetDomainClassParameters

`func (o *VlanDomainEditInput) GetDomainClassParameters() []ApiClassParameterInputEntry`

GetDomainClassParameters returns the DomainClassParameters field if non-nil, zero value otherwise.

### GetDomainClassParametersOk

`func (o *VlanDomainEditInput) GetDomainClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetDomainClassParametersOk returns a tuple with the DomainClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassParameters

`func (o *VlanDomainEditInput) SetDomainClassParameters(v []ApiClassParameterInputEntry)`

SetDomainClassParameters sets DomainClassParameters field to given value.

### HasDomainClassParameters

`func (o *VlanDomainEditInput) HasDomainClassParameters() bool`

HasDomainClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *VlanDomainEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VlanDomainEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VlanDomainEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *VlanDomainEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


