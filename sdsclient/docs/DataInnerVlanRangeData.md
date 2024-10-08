# DataInnerVlanRangeData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSupportVxlan** | Pointer to **string** | The type of virtual network used by the domain the range belongs to, VXLAN (&lt;b&gt;1&lt;/b&gt;) or VLAN (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**DomainClassName** | Pointer to **string** | The name of the class applied to the VLAN domain the object belongs to, it can be preceded by the class directory. | [optional] 
**DomainClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the VLAN domain the object belongs to. | [optional] 
**DomainDescription** | Pointer to **string** | The description of the VLAN domain the object belongs to. | [optional] 
**DomainEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN domain the object belongs to. | [optional] 
**DomainId** | Pointer to **string** | The database identifier (ID) of the VLAN domain the object belongs to. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain the object belongs to. | [optional] 
**DomainStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN domain the object belongs to. | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class applied to the VLAN range, it can be preceded by the class directory. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the VLAN range. | [optional] 
**RangeDescription** | Pointer to **string** | The description of the VLAN range. | [optional] 
**RangeDisableOverlapping** | Pointer to **string** | The overlapping restriction status of the VLAN range. &lt;b&gt;1&lt;/b&gt; indicates that when creating VLANs in the range, their IDs should not overlap. | [optional] 
**RangeEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN range. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the VLAN range. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range. | [optional] 
**RangeStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN range. | [optional] 

## Methods

### NewDataInnerVlanRangeData

`func NewDataInnerVlanRangeData() *DataInnerVlanRangeData`

NewDataInnerVlanRangeData instantiates a new DataInnerVlanRangeData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerVlanRangeDataWithDefaults

`func NewDataInnerVlanRangeDataWithDefaults() *DataInnerVlanRangeData`

NewDataInnerVlanRangeDataWithDefaults instantiates a new DataInnerVlanRangeData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainSupportVxlan

`func (o *DataInnerVlanRangeData) GetDomainSupportVxlan() string`

GetDomainSupportVxlan returns the DomainSupportVxlan field if non-nil, zero value otherwise.

### GetDomainSupportVxlanOk

`func (o *DataInnerVlanRangeData) GetDomainSupportVxlanOk() (*string, bool)`

GetDomainSupportVxlanOk returns a tuple with the DomainSupportVxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSupportVxlan

`func (o *DataInnerVlanRangeData) SetDomainSupportVxlan(v string)`

SetDomainSupportVxlan sets DomainSupportVxlan field to given value.

### HasDomainSupportVxlan

`func (o *DataInnerVlanRangeData) HasDomainSupportVxlan() bool`

HasDomainSupportVxlan returns a boolean if a field has been set.

### GetDomainClassName

`func (o *DataInnerVlanRangeData) GetDomainClassName() string`

GetDomainClassName returns the DomainClassName field if non-nil, zero value otherwise.

### GetDomainClassNameOk

`func (o *DataInnerVlanRangeData) GetDomainClassNameOk() (*string, bool)`

GetDomainClassNameOk returns a tuple with the DomainClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassName

`func (o *DataInnerVlanRangeData) SetDomainClassName(v string)`

SetDomainClassName sets DomainClassName field to given value.

### HasDomainClassName

`func (o *DataInnerVlanRangeData) HasDomainClassName() bool`

HasDomainClassName returns a boolean if a field has been set.

### GetDomainClassParameters

`func (o *DataInnerVlanRangeData) GetDomainClassParameters() []ApiClassParameterOutputEntry`

GetDomainClassParameters returns the DomainClassParameters field if non-nil, zero value otherwise.

### GetDomainClassParametersOk

`func (o *DataInnerVlanRangeData) GetDomainClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetDomainClassParametersOk returns a tuple with the DomainClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassParameters

`func (o *DataInnerVlanRangeData) SetDomainClassParameters(v []ApiClassParameterOutputEntry)`

SetDomainClassParameters sets DomainClassParameters field to given value.

### HasDomainClassParameters

`func (o *DataInnerVlanRangeData) HasDomainClassParameters() bool`

HasDomainClassParameters returns a boolean if a field has been set.

### GetDomainDescription

`func (o *DataInnerVlanRangeData) GetDomainDescription() string`

GetDomainDescription returns the DomainDescription field if non-nil, zero value otherwise.

### GetDomainDescriptionOk

`func (o *DataInnerVlanRangeData) GetDomainDescriptionOk() (*string, bool)`

GetDomainDescriptionOk returns a tuple with the DomainDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDescription

`func (o *DataInnerVlanRangeData) SetDomainDescription(v string)`

SetDomainDescription sets DomainDescription field to given value.

### HasDomainDescription

`func (o *DataInnerVlanRangeData) HasDomainDescription() bool`

HasDomainDescription returns a boolean if a field has been set.

### GetDomainEndVlanId

`func (o *DataInnerVlanRangeData) GetDomainEndVlanId() string`

GetDomainEndVlanId returns the DomainEndVlanId field if non-nil, zero value otherwise.

### GetDomainEndVlanIdOk

`func (o *DataInnerVlanRangeData) GetDomainEndVlanIdOk() (*string, bool)`

GetDomainEndVlanIdOk returns a tuple with the DomainEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEndVlanId

`func (o *DataInnerVlanRangeData) SetDomainEndVlanId(v string)`

SetDomainEndVlanId sets DomainEndVlanId field to given value.

### HasDomainEndVlanId

`func (o *DataInnerVlanRangeData) HasDomainEndVlanId() bool`

HasDomainEndVlanId returns a boolean if a field has been set.

### GetDomainId

`func (o *DataInnerVlanRangeData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DataInnerVlanRangeData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DataInnerVlanRangeData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DataInnerVlanRangeData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *DataInnerVlanRangeData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DataInnerVlanRangeData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DataInnerVlanRangeData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DataInnerVlanRangeData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainStartVlanId

`func (o *DataInnerVlanRangeData) GetDomainStartVlanId() string`

GetDomainStartVlanId returns the DomainStartVlanId field if non-nil, zero value otherwise.

### GetDomainStartVlanIdOk

`func (o *DataInnerVlanRangeData) GetDomainStartVlanIdOk() (*string, bool)`

GetDomainStartVlanIdOk returns a tuple with the DomainStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStartVlanId

`func (o *DataInnerVlanRangeData) SetDomainStartVlanId(v string)`

SetDomainStartVlanId sets DomainStartVlanId field to given value.

### HasDomainStartVlanId

`func (o *DataInnerVlanRangeData) HasDomainStartVlanId() bool`

HasDomainStartVlanId returns a boolean if a field has been set.

### GetRangeClassName

`func (o *DataInnerVlanRangeData) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *DataInnerVlanRangeData) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *DataInnerVlanRangeData) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *DataInnerVlanRangeData) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *DataInnerVlanRangeData) GetRangeClassParameters() []ApiClassParameterOutputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *DataInnerVlanRangeData) GetRangeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *DataInnerVlanRangeData) SetRangeClassParameters(v []ApiClassParameterOutputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *DataInnerVlanRangeData) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetRangeDescription

`func (o *DataInnerVlanRangeData) GetRangeDescription() string`

GetRangeDescription returns the RangeDescription field if non-nil, zero value otherwise.

### GetRangeDescriptionOk

`func (o *DataInnerVlanRangeData) GetRangeDescriptionOk() (*string, bool)`

GetRangeDescriptionOk returns a tuple with the RangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDescription

`func (o *DataInnerVlanRangeData) SetRangeDescription(v string)`

SetRangeDescription sets RangeDescription field to given value.

### HasRangeDescription

`func (o *DataInnerVlanRangeData) HasRangeDescription() bool`

HasRangeDescription returns a boolean if a field has been set.

### GetRangeDisableOverlapping

`func (o *DataInnerVlanRangeData) GetRangeDisableOverlapping() string`

GetRangeDisableOverlapping returns the RangeDisableOverlapping field if non-nil, zero value otherwise.

### GetRangeDisableOverlappingOk

`func (o *DataInnerVlanRangeData) GetRangeDisableOverlappingOk() (*string, bool)`

GetRangeDisableOverlappingOk returns a tuple with the RangeDisableOverlapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDisableOverlapping

`func (o *DataInnerVlanRangeData) SetRangeDisableOverlapping(v string)`

SetRangeDisableOverlapping sets RangeDisableOverlapping field to given value.

### HasRangeDisableOverlapping

`func (o *DataInnerVlanRangeData) HasRangeDisableOverlapping() bool`

HasRangeDisableOverlapping returns a boolean if a field has been set.

### GetRangeEndVlanId

`func (o *DataInnerVlanRangeData) GetRangeEndVlanId() string`

GetRangeEndVlanId returns the RangeEndVlanId field if non-nil, zero value otherwise.

### GetRangeEndVlanIdOk

`func (o *DataInnerVlanRangeData) GetRangeEndVlanIdOk() (*string, bool)`

GetRangeEndVlanIdOk returns a tuple with the RangeEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndVlanId

`func (o *DataInnerVlanRangeData) SetRangeEndVlanId(v string)`

SetRangeEndVlanId sets RangeEndVlanId field to given value.

### HasRangeEndVlanId

`func (o *DataInnerVlanRangeData) HasRangeEndVlanId() bool`

HasRangeEndVlanId returns a boolean if a field has been set.

### GetRangeId

`func (o *DataInnerVlanRangeData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *DataInnerVlanRangeData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *DataInnerVlanRangeData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *DataInnerVlanRangeData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *DataInnerVlanRangeData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *DataInnerVlanRangeData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *DataInnerVlanRangeData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *DataInnerVlanRangeData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetRangeStartVlanId

`func (o *DataInnerVlanRangeData) GetRangeStartVlanId() string`

GetRangeStartVlanId returns the RangeStartVlanId field if non-nil, zero value otherwise.

### GetRangeStartVlanIdOk

`func (o *DataInnerVlanRangeData) GetRangeStartVlanIdOk() (*string, bool)`

GetRangeStartVlanIdOk returns a tuple with the RangeStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartVlanId

`func (o *DataInnerVlanRangeData) SetRangeStartVlanId(v string)`

SetRangeStartVlanId sets RangeStartVlanId field to given value.

### HasRangeStartVlanId

`func (o *DataInnerVlanRangeData) HasRangeStartVlanId() bool`

HasRangeStartVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


