# VlanVlanEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | Pointer to **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain. | [optional] 
**VlanId** | Pointer to **int32** | The database identifier (ID) of the VLAN, a unique numeric key value automatically incremented when you add a VLAN. Use the ID to specify the VLAN of your choice. | [optional] 
**VlanVlanId** | Pointer to **int32** | The VLAN identifier (ID) of the VLAN, a unique numeric key value within a VLAN domain. Use the ID to specify the VLAN of your choice. | [optional] 
**RangeId** | Pointer to **int32** | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range. | [optional] 
**VlanName** | Pointer to **string** | The name of the VLAN, each VLAN must have a unique name. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**VlanClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**VlanClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewVlanVlanEditInput

`func NewVlanVlanEditInput() *VlanVlanEditInput`

NewVlanVlanEditInput instantiates a new VlanVlanEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanVlanEditInputWithDefaults

`func NewVlanVlanEditInputWithDefaults() *VlanVlanEditInput`

NewVlanVlanEditInputWithDefaults instantiates a new VlanVlanEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *VlanVlanEditInput) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *VlanVlanEditInput) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *VlanVlanEditInput) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *VlanVlanEditInput) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *VlanVlanEditInput) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VlanVlanEditInput) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VlanVlanEditInput) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *VlanVlanEditInput) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetVlanId

`func (o *VlanVlanEditInput) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VlanVlanEditInput) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VlanVlanEditInput) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *VlanVlanEditInput) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanVlanId

`func (o *VlanVlanEditInput) GetVlanVlanId() int32`

GetVlanVlanId returns the VlanVlanId field if non-nil, zero value otherwise.

### GetVlanVlanIdOk

`func (o *VlanVlanEditInput) GetVlanVlanIdOk() (*int32, bool)`

GetVlanVlanIdOk returns a tuple with the VlanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVlanId

`func (o *VlanVlanEditInput) SetVlanVlanId(v int32)`

SetVlanVlanId sets VlanVlanId field to given value.

### HasVlanVlanId

`func (o *VlanVlanEditInput) HasVlanVlanId() bool`

HasVlanVlanId returns a boolean if a field has been set.

### GetRangeId

`func (o *VlanVlanEditInput) GetRangeId() int32`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *VlanVlanEditInput) GetRangeIdOk() (*int32, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *VlanVlanEditInput) SetRangeId(v int32)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *VlanVlanEditInput) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *VlanVlanEditInput) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *VlanVlanEditInput) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *VlanVlanEditInput) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *VlanVlanEditInput) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetVlanName

`func (o *VlanVlanEditInput) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *VlanVlanEditInput) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *VlanVlanEditInput) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *VlanVlanEditInput) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *VlanVlanEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *VlanVlanEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *VlanVlanEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *VlanVlanEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetVlanClassName

`func (o *VlanVlanEditInput) GetVlanClassName() string`

GetVlanClassName returns the VlanClassName field if non-nil, zero value otherwise.

### GetVlanClassNameOk

`func (o *VlanVlanEditInput) GetVlanClassNameOk() (*string, bool)`

GetVlanClassNameOk returns a tuple with the VlanClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanClassName

`func (o *VlanVlanEditInput) SetVlanClassName(v string)`

SetVlanClassName sets VlanClassName field to given value.

### HasVlanClassName

`func (o *VlanVlanEditInput) HasVlanClassName() bool`

HasVlanClassName returns a boolean if a field has been set.

### GetVlanClassParameters

`func (o *VlanVlanEditInput) GetVlanClassParameters() []ApiClassParameterInputEntry`

GetVlanClassParameters returns the VlanClassParameters field if non-nil, zero value otherwise.

### GetVlanClassParametersOk

`func (o *VlanVlanEditInput) GetVlanClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetVlanClassParametersOk returns a tuple with the VlanClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanClassParameters

`func (o *VlanVlanEditInput) SetVlanClassParameters(v []ApiClassParameterInputEntry)`

SetVlanClassParameters sets VlanClassParameters field to given value.

### HasVlanClassParameters

`func (o *VlanVlanEditInput) HasVlanClassParameters() bool`

HasVlanClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *VlanVlanEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VlanVlanEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VlanVlanEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *VlanVlanEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


