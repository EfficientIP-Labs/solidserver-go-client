# VlanRangeAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainId** | Pointer to **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range, each VLAN range must have a unique name. | [optional] 
**RangeDescription** | Pointer to **string** | The description of the VLAN range. | [optional] 
**RangeDisableOverlapping** | Pointer to **int32** | The overlapping restriction status of the VLAN range. Set it to &lt;b&gt;1&lt;/b&gt; to prevent VLAN ID overlapping in the range. | [optional] 
**RangeEndVlanId** | Pointer to **int32** | The VLAN identifier (ID) of an existing VLAN you want to set as the last VLAN in the VLAN range. | [optional] 
**RangeStartVlanId** | Pointer to **int32** | The VLAN identifier (ID) of an existing VLAN you want to set as the first VLAN in the VLAN range. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewVlanRangeAddInput

`func NewVlanRangeAddInput() *VlanRangeAddInput`

NewVlanRangeAddInput instantiates a new VlanRangeAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanRangeAddInputWithDefaults

`func NewVlanRangeAddInputWithDefaults() *VlanRangeAddInput`

NewVlanRangeAddInputWithDefaults instantiates a new VlanRangeAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainId

`func (o *VlanRangeAddInput) GetDomainId() int32`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *VlanRangeAddInput) GetDomainIdOk() (*int32, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *VlanRangeAddInput) SetDomainId(v int32)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *VlanRangeAddInput) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *VlanRangeAddInput) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VlanRangeAddInput) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VlanRangeAddInput) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *VlanRangeAddInput) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetRangeName

`func (o *VlanRangeAddInput) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *VlanRangeAddInput) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *VlanRangeAddInput) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *VlanRangeAddInput) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetRangeDescription

`func (o *VlanRangeAddInput) GetRangeDescription() string`

GetRangeDescription returns the RangeDescription field if non-nil, zero value otherwise.

### GetRangeDescriptionOk

`func (o *VlanRangeAddInput) GetRangeDescriptionOk() (*string, bool)`

GetRangeDescriptionOk returns a tuple with the RangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDescription

`func (o *VlanRangeAddInput) SetRangeDescription(v string)`

SetRangeDescription sets RangeDescription field to given value.

### HasRangeDescription

`func (o *VlanRangeAddInput) HasRangeDescription() bool`

HasRangeDescription returns a boolean if a field has been set.

### GetRangeDisableOverlapping

`func (o *VlanRangeAddInput) GetRangeDisableOverlapping() int32`

GetRangeDisableOverlapping returns the RangeDisableOverlapping field if non-nil, zero value otherwise.

### GetRangeDisableOverlappingOk

`func (o *VlanRangeAddInput) GetRangeDisableOverlappingOk() (*int32, bool)`

GetRangeDisableOverlappingOk returns a tuple with the RangeDisableOverlapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDisableOverlapping

`func (o *VlanRangeAddInput) SetRangeDisableOverlapping(v int32)`

SetRangeDisableOverlapping sets RangeDisableOverlapping field to given value.

### HasRangeDisableOverlapping

`func (o *VlanRangeAddInput) HasRangeDisableOverlapping() bool`

HasRangeDisableOverlapping returns a boolean if a field has been set.

### GetRangeEndVlanId

`func (o *VlanRangeAddInput) GetRangeEndVlanId() int32`

GetRangeEndVlanId returns the RangeEndVlanId field if non-nil, zero value otherwise.

### GetRangeEndVlanIdOk

`func (o *VlanRangeAddInput) GetRangeEndVlanIdOk() (*int32, bool)`

GetRangeEndVlanIdOk returns a tuple with the RangeEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndVlanId

`func (o *VlanRangeAddInput) SetRangeEndVlanId(v int32)`

SetRangeEndVlanId sets RangeEndVlanId field to given value.

### HasRangeEndVlanId

`func (o *VlanRangeAddInput) HasRangeEndVlanId() bool`

HasRangeEndVlanId returns a boolean if a field has been set.

### GetRangeStartVlanId

`func (o *VlanRangeAddInput) GetRangeStartVlanId() int32`

GetRangeStartVlanId returns the RangeStartVlanId field if non-nil, zero value otherwise.

### GetRangeStartVlanIdOk

`func (o *VlanRangeAddInput) GetRangeStartVlanIdOk() (*int32, bool)`

GetRangeStartVlanIdOk returns a tuple with the RangeStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartVlanId

`func (o *VlanRangeAddInput) SetRangeStartVlanId(v int32)`

SetRangeStartVlanId sets RangeStartVlanId field to given value.

### HasRangeStartVlanId

`func (o *VlanRangeAddInput) HasRangeStartVlanId() bool`

HasRangeStartVlanId returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *VlanRangeAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *VlanRangeAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *VlanRangeAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *VlanRangeAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetRangeClassName

`func (o *VlanRangeAddInput) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *VlanRangeAddInput) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *VlanRangeAddInput) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *VlanRangeAddInput) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *VlanRangeAddInput) GetRangeClassParameters() []ApiClassParameterInputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *VlanRangeAddInput) GetRangeClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *VlanRangeAddInput) SetRangeClassParameters(v []ApiClassParameterInputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *VlanRangeAddInput) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *VlanRangeAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *VlanRangeAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *VlanRangeAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *VlanRangeAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


