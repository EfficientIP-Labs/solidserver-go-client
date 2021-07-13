# VlanRangeDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSupportVxlan** | Pointer to **string** | The type of virtual network used by the domain the range belongs to, VXLAN (&lt;b&gt;1&lt;/b&gt;) or VLAN (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**DomainClassName** | Pointer to **string** | The name of the class applied to the VLAN domain the object belongs to, it can be preceded by the class directory. | [optional] 
**DomainClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**DomainDescription** | Pointer to **string** | The description of the VLAN domain the object belongs to. | [optional] 
**DomainEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN domain the object belongs to. | [optional] 
**DomainId** | Pointer to **string** | The database identifier (ID) of the VLAN domain the object belongs to. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain the object belongs to. | [optional] 
**DomainStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN domain the object belongs to. | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class applied to the VLAN range, it can be preceded by the class directory. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**RangeDescription** | Pointer to **string** | The description of the VLAN range. | [optional] 
**RangeDisableOverlapping** | Pointer to **string** | The overlapping restriction status of the VLAN range. &lt;b&gt;1&lt;/b&gt; indicates that when creating VLANs in the range, their IDs should not overlap. | [optional] 
**RangeEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN range. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the VLAN range. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range. | [optional] 
**RangeStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN range. | [optional] 

## Methods

### NewVlanRangeDataData

`func NewVlanRangeDataData() *VlanRangeDataData`

NewVlanRangeDataData instantiates a new VlanRangeDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanRangeDataDataWithDefaults

`func NewVlanRangeDataDataWithDefaults() *VlanRangeDataData`

NewVlanRangeDataDataWithDefaults instantiates a new VlanRangeDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainSupportVxlan

`func (o *VlanRangeDataData) GetDomainSupportVxlan() string`

GetDomainSupportVxlan returns the DomainSupportVxlan field if non-nil, zero value otherwise.

### GetDomainSupportVxlanOk

`func (o *VlanRangeDataData) GetDomainSupportVxlanOk() (*string, bool)`

GetDomainSupportVxlanOk returns a tuple with the DomainSupportVxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSupportVxlan

`func (o *VlanRangeDataData) SetDomainSupportVxlan(v string)`

SetDomainSupportVxlan sets DomainSupportVxlan field to given value.

### HasDomainSupportVxlan

`func (o *VlanRangeDataData) HasDomainSupportVxlan() bool`

HasDomainSupportVxlan returns a boolean if a field has been set.

### GetDomainClassName

`func (o *VlanRangeDataData) GetDomainClassName() string`

GetDomainClassName returns the DomainClassName field if non-nil, zero value otherwise.

### GetDomainClassNameOk

`func (o *VlanRangeDataData) GetDomainClassNameOk() (*string, bool)`

GetDomainClassNameOk returns a tuple with the DomainClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassName

`func (o *VlanRangeDataData) SetDomainClassName(v string)`

SetDomainClassName sets DomainClassName field to given value.

### HasDomainClassName

`func (o *VlanRangeDataData) HasDomainClassName() bool`

HasDomainClassName returns a boolean if a field has been set.

### GetDomainClassParameters

`func (o *VlanRangeDataData) GetDomainClassParameters() []ApiClassParameterOutputEntry`

GetDomainClassParameters returns the DomainClassParameters field if non-nil, zero value otherwise.

### GetDomainClassParametersOk

`func (o *VlanRangeDataData) GetDomainClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetDomainClassParametersOk returns a tuple with the DomainClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassParameters

`func (o *VlanRangeDataData) SetDomainClassParameters(v []ApiClassParameterOutputEntry)`

SetDomainClassParameters sets DomainClassParameters field to given value.

### HasDomainClassParameters

`func (o *VlanRangeDataData) HasDomainClassParameters() bool`

HasDomainClassParameters returns a boolean if a field has been set.

### GetDomainDescription

`func (o *VlanRangeDataData) GetDomainDescription() string`

GetDomainDescription returns the DomainDescription field if non-nil, zero value otherwise.

### GetDomainDescriptionOk

`func (o *VlanRangeDataData) GetDomainDescriptionOk() (*string, bool)`

GetDomainDescriptionOk returns a tuple with the DomainDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDescription

`func (o *VlanRangeDataData) SetDomainDescription(v string)`

SetDomainDescription sets DomainDescription field to given value.

### HasDomainDescription

`func (o *VlanRangeDataData) HasDomainDescription() bool`

HasDomainDescription returns a boolean if a field has been set.

### GetDomainEndVlanId

`func (o *VlanRangeDataData) GetDomainEndVlanId() string`

GetDomainEndVlanId returns the DomainEndVlanId field if non-nil, zero value otherwise.

### GetDomainEndVlanIdOk

`func (o *VlanRangeDataData) GetDomainEndVlanIdOk() (*string, bool)`

GetDomainEndVlanIdOk returns a tuple with the DomainEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEndVlanId

`func (o *VlanRangeDataData) SetDomainEndVlanId(v string)`

SetDomainEndVlanId sets DomainEndVlanId field to given value.

### HasDomainEndVlanId

`func (o *VlanRangeDataData) HasDomainEndVlanId() bool`

HasDomainEndVlanId returns a boolean if a field has been set.

### GetDomainId

`func (o *VlanRangeDataData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *VlanRangeDataData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *VlanRangeDataData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *VlanRangeDataData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *VlanRangeDataData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VlanRangeDataData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VlanRangeDataData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *VlanRangeDataData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainStartVlanId

`func (o *VlanRangeDataData) GetDomainStartVlanId() string`

GetDomainStartVlanId returns the DomainStartVlanId field if non-nil, zero value otherwise.

### GetDomainStartVlanIdOk

`func (o *VlanRangeDataData) GetDomainStartVlanIdOk() (*string, bool)`

GetDomainStartVlanIdOk returns a tuple with the DomainStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStartVlanId

`func (o *VlanRangeDataData) SetDomainStartVlanId(v string)`

SetDomainStartVlanId sets DomainStartVlanId field to given value.

### HasDomainStartVlanId

`func (o *VlanRangeDataData) HasDomainStartVlanId() bool`

HasDomainStartVlanId returns a boolean if a field has been set.

### GetRangeClassName

`func (o *VlanRangeDataData) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *VlanRangeDataData) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *VlanRangeDataData) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *VlanRangeDataData) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *VlanRangeDataData) GetRangeClassParameters() []ApiClassParameterOutputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *VlanRangeDataData) GetRangeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *VlanRangeDataData) SetRangeClassParameters(v []ApiClassParameterOutputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *VlanRangeDataData) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetRangeDescription

`func (o *VlanRangeDataData) GetRangeDescription() string`

GetRangeDescription returns the RangeDescription field if non-nil, zero value otherwise.

### GetRangeDescriptionOk

`func (o *VlanRangeDataData) GetRangeDescriptionOk() (*string, bool)`

GetRangeDescriptionOk returns a tuple with the RangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDescription

`func (o *VlanRangeDataData) SetRangeDescription(v string)`

SetRangeDescription sets RangeDescription field to given value.

### HasRangeDescription

`func (o *VlanRangeDataData) HasRangeDescription() bool`

HasRangeDescription returns a boolean if a field has been set.

### GetRangeDisableOverlapping

`func (o *VlanRangeDataData) GetRangeDisableOverlapping() string`

GetRangeDisableOverlapping returns the RangeDisableOverlapping field if non-nil, zero value otherwise.

### GetRangeDisableOverlappingOk

`func (o *VlanRangeDataData) GetRangeDisableOverlappingOk() (*string, bool)`

GetRangeDisableOverlappingOk returns a tuple with the RangeDisableOverlapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDisableOverlapping

`func (o *VlanRangeDataData) SetRangeDisableOverlapping(v string)`

SetRangeDisableOverlapping sets RangeDisableOverlapping field to given value.

### HasRangeDisableOverlapping

`func (o *VlanRangeDataData) HasRangeDisableOverlapping() bool`

HasRangeDisableOverlapping returns a boolean if a field has been set.

### GetRangeEndVlanId

`func (o *VlanRangeDataData) GetRangeEndVlanId() string`

GetRangeEndVlanId returns the RangeEndVlanId field if non-nil, zero value otherwise.

### GetRangeEndVlanIdOk

`func (o *VlanRangeDataData) GetRangeEndVlanIdOk() (*string, bool)`

GetRangeEndVlanIdOk returns a tuple with the RangeEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndVlanId

`func (o *VlanRangeDataData) SetRangeEndVlanId(v string)`

SetRangeEndVlanId sets RangeEndVlanId field to given value.

### HasRangeEndVlanId

`func (o *VlanRangeDataData) HasRangeEndVlanId() bool`

HasRangeEndVlanId returns a boolean if a field has been set.

### GetRangeId

`func (o *VlanRangeDataData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *VlanRangeDataData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *VlanRangeDataData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *VlanRangeDataData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *VlanRangeDataData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *VlanRangeDataData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *VlanRangeDataData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *VlanRangeDataData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetRangeStartVlanId

`func (o *VlanRangeDataData) GetRangeStartVlanId() string`

GetRangeStartVlanId returns the RangeStartVlanId field if non-nil, zero value otherwise.

### GetRangeStartVlanIdOk

`func (o *VlanRangeDataData) GetRangeStartVlanIdOk() (*string, bool)`

GetRangeStartVlanIdOk returns a tuple with the RangeStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartVlanId

`func (o *VlanRangeDataData) SetRangeStartVlanId(v string)`

SetRangeStartVlanId sets RangeStartVlanId field to given value.

### HasRangeStartVlanId

`func (o *VlanRangeDataData) HasRangeStartVlanId() bool`

HasRangeStartVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


