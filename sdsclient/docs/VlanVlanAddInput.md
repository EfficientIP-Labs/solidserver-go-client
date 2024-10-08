# VlanVlanAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | Pointer to **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain. | [optional] 
**VlanVlanId** | Pointer to **int32** | The VLAN identifier (ID) of the VLAN, a unique numeric key value within a VLAN domain. Use the ID to specify the VLAN of your choice. | [optional] 
**RangeId** | Pointer to **int32** | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range. | [optional] 
**VlanName** | Pointer to **string** | The name of the VLAN, each VLAN must have a unique name. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**VlanClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**VlanClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewVlanVlanAddInput

`func NewVlanVlanAddInput() *VlanVlanAddInput`

NewVlanVlanAddInput instantiates a new VlanVlanAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanVlanAddInputWithDefaults

`func NewVlanVlanAddInputWithDefaults() *VlanVlanAddInput`

NewVlanVlanAddInputWithDefaults instantiates a new VlanVlanAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *VlanVlanAddInput) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *VlanVlanAddInput) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *VlanVlanAddInput) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *VlanVlanAddInput) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *VlanVlanAddInput) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VlanVlanAddInput) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VlanVlanAddInput) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *VlanVlanAddInput) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetVlanVlanId

`func (o *VlanVlanAddInput) GetVlanVlanId() int32`

GetVlanVlanId returns the VlanVlanId field if non-nil, zero value otherwise.

### GetVlanVlanIdOk

`func (o *VlanVlanAddInput) GetVlanVlanIdOk() (*int32, bool)`

GetVlanVlanIdOk returns a tuple with the VlanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVlanId

`func (o *VlanVlanAddInput) SetVlanVlanId(v int32)`

SetVlanVlanId sets VlanVlanId field to given value.

### HasVlanVlanId

`func (o *VlanVlanAddInput) HasVlanVlanId() bool`

HasVlanVlanId returns a boolean if a field has been set.

### GetRangeId

`func (o *VlanVlanAddInput) GetRangeId() int32`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *VlanVlanAddInput) GetRangeIdOk() (*int32, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *VlanVlanAddInput) SetRangeId(v int32)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *VlanVlanAddInput) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *VlanVlanAddInput) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *VlanVlanAddInput) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *VlanVlanAddInput) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *VlanVlanAddInput) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetVlanName

`func (o *VlanVlanAddInput) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *VlanVlanAddInput) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *VlanVlanAddInput) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *VlanVlanAddInput) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *VlanVlanAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *VlanVlanAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *VlanVlanAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *VlanVlanAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetVlanClassName

`func (o *VlanVlanAddInput) GetVlanClassName() string`

GetVlanClassName returns the VlanClassName field if non-nil, zero value otherwise.

### GetVlanClassNameOk

`func (o *VlanVlanAddInput) GetVlanClassNameOk() (*string, bool)`

GetVlanClassNameOk returns a tuple with the VlanClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanClassName

`func (o *VlanVlanAddInput) SetVlanClassName(v string)`

SetVlanClassName sets VlanClassName field to given value.

### HasVlanClassName

`func (o *VlanVlanAddInput) HasVlanClassName() bool`

HasVlanClassName returns a boolean if a field has been set.

### GetVlanClassParameters

`func (o *VlanVlanAddInput) GetVlanClassParameters() []ApiClassParameterInputEntry`

GetVlanClassParameters returns the VlanClassParameters field if non-nil, zero value otherwise.

### GetVlanClassParametersOk

`func (o *VlanVlanAddInput) GetVlanClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetVlanClassParametersOk returns a tuple with the VlanClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanClassParameters

`func (o *VlanVlanAddInput) SetVlanClassParameters(v []ApiClassParameterInputEntry)`

SetVlanClassParameters sets VlanClassParameters field to given value.

### HasVlanClassParameters

`func (o *VlanVlanAddInput) HasVlanClassParameters() bool`

HasVlanClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *VlanVlanAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VlanVlanAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VlanVlanAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *VlanVlanAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


