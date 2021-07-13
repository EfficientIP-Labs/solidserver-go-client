# VlanVlanDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeEndVlanId** | Pointer to **string** | For free VLAN IDs (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;), it returns the last VLAN of a range of VLANs that are not assigned yet. The first VLAN in that range is returned in the parameter &lt;b&gt;free_start_vlan_id&lt;/b&gt;. | [optional] 
**FreeStartVlanId** | Pointer to **string** | For free VLAN IDs (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;), it returns the first VLAN of the range of VLANs that are not assigned yet. The last VLAN in that range is returned in the parameter &lt;b&gt;free_end_vlan_id&lt;/b&gt;. | [optional] 
**DomainSupportVxlan** | Pointer to **string** | The type of virtual network used by the domain the VLAN belongs to, VXLAN (&lt;b&gt;1&lt;/b&gt;) or VLAN (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**VlanType** | Pointer to **string** | The type of the VLAN (&lt;b&gt;free&lt;/b&gt; or &lt;b&gt;used&lt;/b&gt;). | [optional] 
**DomainClassName** | Pointer to **string** | The name of the class applied to the VLAN domain the object belongs to, it can be preceded by the class directory. | [optional] 
**DomainClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**DomainDescription** | Pointer to **string** | The description of the VLAN domain the object belongs to. | [optional] 
**DomainEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN domain the object belongs to. | [optional] 
**DomainId** | Pointer to **string** | The database identifier (ID) of the VLAN domain the object belongs to. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain the object belongs to. | [optional] 
**DomainStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN domain the object belongs to. | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class applied to the VLAN range the object belongs to, it can be preceded by the class directory. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**RangeDescription** | Pointer to **string** | The description of the VLAN range the object belongs to. | [optional] 
**RangeEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN range the object belongs to. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the VLAN range the object belongs to. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range the object belongs to. | [optional] 
**RangeStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN range the object belongs to. | [optional] 
**VlanClassName** | Pointer to **string** | The name of the class applied to the object. | [optional] 
**VlanClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**VlanId** | Pointer to **string** | The database identifier (ID) of the VLAN. | [optional] 
**VlanName** | Pointer to **string** | The name of the VLAN. | [optional] 
**VlanVlanId** | Pointer to **string** | The VLAN identifier (ID) of the VLAN. | [optional] 

## Methods

### NewVlanVlanDataData

`func NewVlanVlanDataData() *VlanVlanDataData`

NewVlanVlanDataData instantiates a new VlanVlanDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanVlanDataDataWithDefaults

`func NewVlanVlanDataDataWithDefaults() *VlanVlanDataData`

NewVlanVlanDataDataWithDefaults instantiates a new VlanVlanDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeEndVlanId

`func (o *VlanVlanDataData) GetFreeEndVlanId() string`

GetFreeEndVlanId returns the FreeEndVlanId field if non-nil, zero value otherwise.

### GetFreeEndVlanIdOk

`func (o *VlanVlanDataData) GetFreeEndVlanIdOk() (*string, bool)`

GetFreeEndVlanIdOk returns a tuple with the FreeEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeEndVlanId

`func (o *VlanVlanDataData) SetFreeEndVlanId(v string)`

SetFreeEndVlanId sets FreeEndVlanId field to given value.

### HasFreeEndVlanId

`func (o *VlanVlanDataData) HasFreeEndVlanId() bool`

HasFreeEndVlanId returns a boolean if a field has been set.

### GetFreeStartVlanId

`func (o *VlanVlanDataData) GetFreeStartVlanId() string`

GetFreeStartVlanId returns the FreeStartVlanId field if non-nil, zero value otherwise.

### GetFreeStartVlanIdOk

`func (o *VlanVlanDataData) GetFreeStartVlanIdOk() (*string, bool)`

GetFreeStartVlanIdOk returns a tuple with the FreeStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStartVlanId

`func (o *VlanVlanDataData) SetFreeStartVlanId(v string)`

SetFreeStartVlanId sets FreeStartVlanId field to given value.

### HasFreeStartVlanId

`func (o *VlanVlanDataData) HasFreeStartVlanId() bool`

HasFreeStartVlanId returns a boolean if a field has been set.

### GetDomainSupportVxlan

`func (o *VlanVlanDataData) GetDomainSupportVxlan() string`

GetDomainSupportVxlan returns the DomainSupportVxlan field if non-nil, zero value otherwise.

### GetDomainSupportVxlanOk

`func (o *VlanVlanDataData) GetDomainSupportVxlanOk() (*string, bool)`

GetDomainSupportVxlanOk returns a tuple with the DomainSupportVxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSupportVxlan

`func (o *VlanVlanDataData) SetDomainSupportVxlan(v string)`

SetDomainSupportVxlan sets DomainSupportVxlan field to given value.

### HasDomainSupportVxlan

`func (o *VlanVlanDataData) HasDomainSupportVxlan() bool`

HasDomainSupportVxlan returns a boolean if a field has been set.

### GetVlanType

`func (o *VlanVlanDataData) GetVlanType() string`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *VlanVlanDataData) GetVlanTypeOk() (*string, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *VlanVlanDataData) SetVlanType(v string)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *VlanVlanDataData) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetDomainClassName

`func (o *VlanVlanDataData) GetDomainClassName() string`

GetDomainClassName returns the DomainClassName field if non-nil, zero value otherwise.

### GetDomainClassNameOk

`func (o *VlanVlanDataData) GetDomainClassNameOk() (*string, bool)`

GetDomainClassNameOk returns a tuple with the DomainClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassName

`func (o *VlanVlanDataData) SetDomainClassName(v string)`

SetDomainClassName sets DomainClassName field to given value.

### HasDomainClassName

`func (o *VlanVlanDataData) HasDomainClassName() bool`

HasDomainClassName returns a boolean if a field has been set.

### GetDomainClassParameters

`func (o *VlanVlanDataData) GetDomainClassParameters() []ApiClassParameterOutputEntry`

GetDomainClassParameters returns the DomainClassParameters field if non-nil, zero value otherwise.

### GetDomainClassParametersOk

`func (o *VlanVlanDataData) GetDomainClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetDomainClassParametersOk returns a tuple with the DomainClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassParameters

`func (o *VlanVlanDataData) SetDomainClassParameters(v []ApiClassParameterOutputEntry)`

SetDomainClassParameters sets DomainClassParameters field to given value.

### HasDomainClassParameters

`func (o *VlanVlanDataData) HasDomainClassParameters() bool`

HasDomainClassParameters returns a boolean if a field has been set.

### GetDomainDescription

`func (o *VlanVlanDataData) GetDomainDescription() string`

GetDomainDescription returns the DomainDescription field if non-nil, zero value otherwise.

### GetDomainDescriptionOk

`func (o *VlanVlanDataData) GetDomainDescriptionOk() (*string, bool)`

GetDomainDescriptionOk returns a tuple with the DomainDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDescription

`func (o *VlanVlanDataData) SetDomainDescription(v string)`

SetDomainDescription sets DomainDescription field to given value.

### HasDomainDescription

`func (o *VlanVlanDataData) HasDomainDescription() bool`

HasDomainDescription returns a boolean if a field has been set.

### GetDomainEndVlanId

`func (o *VlanVlanDataData) GetDomainEndVlanId() string`

GetDomainEndVlanId returns the DomainEndVlanId field if non-nil, zero value otherwise.

### GetDomainEndVlanIdOk

`func (o *VlanVlanDataData) GetDomainEndVlanIdOk() (*string, bool)`

GetDomainEndVlanIdOk returns a tuple with the DomainEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEndVlanId

`func (o *VlanVlanDataData) SetDomainEndVlanId(v string)`

SetDomainEndVlanId sets DomainEndVlanId field to given value.

### HasDomainEndVlanId

`func (o *VlanVlanDataData) HasDomainEndVlanId() bool`

HasDomainEndVlanId returns a boolean if a field has been set.

### GetDomainId

`func (o *VlanVlanDataData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *VlanVlanDataData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *VlanVlanDataData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *VlanVlanDataData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *VlanVlanDataData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VlanVlanDataData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VlanVlanDataData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *VlanVlanDataData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainStartVlanId

`func (o *VlanVlanDataData) GetDomainStartVlanId() string`

GetDomainStartVlanId returns the DomainStartVlanId field if non-nil, zero value otherwise.

### GetDomainStartVlanIdOk

`func (o *VlanVlanDataData) GetDomainStartVlanIdOk() (*string, bool)`

GetDomainStartVlanIdOk returns a tuple with the DomainStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStartVlanId

`func (o *VlanVlanDataData) SetDomainStartVlanId(v string)`

SetDomainStartVlanId sets DomainStartVlanId field to given value.

### HasDomainStartVlanId

`func (o *VlanVlanDataData) HasDomainStartVlanId() bool`

HasDomainStartVlanId returns a boolean if a field has been set.

### GetRangeClassName

`func (o *VlanVlanDataData) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *VlanVlanDataData) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *VlanVlanDataData) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *VlanVlanDataData) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *VlanVlanDataData) GetRangeClassParameters() []ApiClassParameterOutputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *VlanVlanDataData) GetRangeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *VlanVlanDataData) SetRangeClassParameters(v []ApiClassParameterOutputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *VlanVlanDataData) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetRangeDescription

`func (o *VlanVlanDataData) GetRangeDescription() string`

GetRangeDescription returns the RangeDescription field if non-nil, zero value otherwise.

### GetRangeDescriptionOk

`func (o *VlanVlanDataData) GetRangeDescriptionOk() (*string, bool)`

GetRangeDescriptionOk returns a tuple with the RangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDescription

`func (o *VlanVlanDataData) SetRangeDescription(v string)`

SetRangeDescription sets RangeDescription field to given value.

### HasRangeDescription

`func (o *VlanVlanDataData) HasRangeDescription() bool`

HasRangeDescription returns a boolean if a field has been set.

### GetRangeEndVlanId

`func (o *VlanVlanDataData) GetRangeEndVlanId() string`

GetRangeEndVlanId returns the RangeEndVlanId field if non-nil, zero value otherwise.

### GetRangeEndVlanIdOk

`func (o *VlanVlanDataData) GetRangeEndVlanIdOk() (*string, bool)`

GetRangeEndVlanIdOk returns a tuple with the RangeEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndVlanId

`func (o *VlanVlanDataData) SetRangeEndVlanId(v string)`

SetRangeEndVlanId sets RangeEndVlanId field to given value.

### HasRangeEndVlanId

`func (o *VlanVlanDataData) HasRangeEndVlanId() bool`

HasRangeEndVlanId returns a boolean if a field has been set.

### GetRangeId

`func (o *VlanVlanDataData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *VlanVlanDataData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *VlanVlanDataData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *VlanVlanDataData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *VlanVlanDataData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *VlanVlanDataData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *VlanVlanDataData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *VlanVlanDataData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetRangeStartVlanId

`func (o *VlanVlanDataData) GetRangeStartVlanId() string`

GetRangeStartVlanId returns the RangeStartVlanId field if non-nil, zero value otherwise.

### GetRangeStartVlanIdOk

`func (o *VlanVlanDataData) GetRangeStartVlanIdOk() (*string, bool)`

GetRangeStartVlanIdOk returns a tuple with the RangeStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartVlanId

`func (o *VlanVlanDataData) SetRangeStartVlanId(v string)`

SetRangeStartVlanId sets RangeStartVlanId field to given value.

### HasRangeStartVlanId

`func (o *VlanVlanDataData) HasRangeStartVlanId() bool`

HasRangeStartVlanId returns a boolean if a field has been set.

### GetVlanClassName

`func (o *VlanVlanDataData) GetVlanClassName() string`

GetVlanClassName returns the VlanClassName field if non-nil, zero value otherwise.

### GetVlanClassNameOk

`func (o *VlanVlanDataData) GetVlanClassNameOk() (*string, bool)`

GetVlanClassNameOk returns a tuple with the VlanClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanClassName

`func (o *VlanVlanDataData) SetVlanClassName(v string)`

SetVlanClassName sets VlanClassName field to given value.

### HasVlanClassName

`func (o *VlanVlanDataData) HasVlanClassName() bool`

HasVlanClassName returns a boolean if a field has been set.

### GetVlanClassParameters

`func (o *VlanVlanDataData) GetVlanClassParameters() []ApiClassParameterOutputEntry`

GetVlanClassParameters returns the VlanClassParameters field if non-nil, zero value otherwise.

### GetVlanClassParametersOk

`func (o *VlanVlanDataData) GetVlanClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetVlanClassParametersOk returns a tuple with the VlanClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanClassParameters

`func (o *VlanVlanDataData) SetVlanClassParameters(v []ApiClassParameterOutputEntry)`

SetVlanClassParameters sets VlanClassParameters field to given value.

### HasVlanClassParameters

`func (o *VlanVlanDataData) HasVlanClassParameters() bool`

HasVlanClassParameters returns a boolean if a field has been set.

### GetVlanId

`func (o *VlanVlanDataData) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *VlanVlanDataData) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *VlanVlanDataData) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *VlanVlanDataData) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanName

`func (o *VlanVlanDataData) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *VlanVlanDataData) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *VlanVlanDataData) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *VlanVlanDataData) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetVlanVlanId

`func (o *VlanVlanDataData) GetVlanVlanId() string`

GetVlanVlanId returns the VlanVlanId field if non-nil, zero value otherwise.

### GetVlanVlanIdOk

`func (o *VlanVlanDataData) GetVlanVlanIdOk() (*string, bool)`

GetVlanVlanIdOk returns a tuple with the VlanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVlanId

`func (o *VlanVlanDataData) SetVlanVlanId(v string)`

SetVlanVlanId sets VlanVlanId field to given value.

### HasVlanVlanId

`func (o *VlanVlanDataData) HasVlanVlanId() bool`

HasVlanVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


