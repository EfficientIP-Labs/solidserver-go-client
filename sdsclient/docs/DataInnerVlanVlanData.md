# DataInnerVlanVlanData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FreeEndVlanId** | Pointer to **string** | For free VLAN IDs (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;), it returns the last VLAN of a range of VLANs that are not assigned yet. The first VLAN in that range is returned in the parameter &lt;b&gt;free_start_vlan_id&lt;/b&gt;. | [optional] 
**FreeStartVlanId** | Pointer to **string** | For free VLAN IDs (&lt;b&gt;type&lt;/b&gt;&lt;b&gt;free&lt;/b&gt;), it returns the first VLAN of the range of VLANs that are not assigned yet. The last VLAN in that range is returned in the parameter &lt;b&gt;free_end_vlan_id&lt;/b&gt;. | [optional] 
**DomainSupportVxlan** | Pointer to **string** | The type of virtual network used by the domain the VLAN belongs to, VXLAN (&lt;b&gt;1&lt;/b&gt;) or VLAN (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**VlanType** | Pointer to **string** | The type of the VLAN (&lt;b&gt;free&lt;/b&gt; or &lt;b&gt;used&lt;/b&gt;). | [optional] 
**DomainClassName** | Pointer to **string** | The name of the class applied to the VLAN domain the object belongs to, it can be preceded by the class directory. | [optional] 
**DomainClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters of the domain the VLAN belongs to. | [optional] 
**DomainDescription** | Pointer to **string** | The description of the VLAN domain the object belongs to. | [optional] 
**DomainEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN domain the object belongs to. | [optional] 
**DomainId** | Pointer to **string** | The database identifier (ID) of the VLAN domain the object belongs to. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain the object belongs to. | [optional] 
**DomainStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN domain the object belongs to. | [optional] 
**RangeClassName** | Pointer to **string** | The name of the class applied to the VLAN range the object belongs to, it can be preceded by the class directory. | [optional] 
**RangeClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters of the range the VLAN belongs to. | [optional] 
**RangeDescription** | Pointer to **string** | The description of the VLAN range the object belongs to. | [optional] 
**RangeEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN range the object belongs to. | [optional] 
**RangeId** | Pointer to **string** | The database identifier (ID) of the VLAN range the object belongs to. | [optional] 
**RangeName** | Pointer to **string** | The name of the VLAN range the object belongs to. | [optional] 
**RangeStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN range the object belongs to. | [optional] 
**VlanClassName** | Pointer to **string** | The name of the class applied to the object, it can be preceded by the class directory. | [optional] 
**VlanClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters of the VLAN. | [optional] 
**VlanId** | Pointer to **string** | The database identifier (ID) of the VLAN. | [optional] 
**VlanName** | Pointer to **string** | The name of the VLAN. | [optional] 
**VlanVlanId** | Pointer to **string** | The VLAN identifier (ID) of the VLAN. | [optional] 

## Methods

### NewDataInnerVlanVlanData

`func NewDataInnerVlanVlanData() *DataInnerVlanVlanData`

NewDataInnerVlanVlanData instantiates a new DataInnerVlanVlanData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerVlanVlanDataWithDefaults

`func NewDataInnerVlanVlanDataWithDefaults() *DataInnerVlanVlanData`

NewDataInnerVlanVlanDataWithDefaults instantiates a new DataInnerVlanVlanData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFreeEndVlanId

`func (o *DataInnerVlanVlanData) GetFreeEndVlanId() string`

GetFreeEndVlanId returns the FreeEndVlanId field if non-nil, zero value otherwise.

### GetFreeEndVlanIdOk

`func (o *DataInnerVlanVlanData) GetFreeEndVlanIdOk() (*string, bool)`

GetFreeEndVlanIdOk returns a tuple with the FreeEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeEndVlanId

`func (o *DataInnerVlanVlanData) SetFreeEndVlanId(v string)`

SetFreeEndVlanId sets FreeEndVlanId field to given value.

### HasFreeEndVlanId

`func (o *DataInnerVlanVlanData) HasFreeEndVlanId() bool`

HasFreeEndVlanId returns a boolean if a field has been set.

### GetFreeStartVlanId

`func (o *DataInnerVlanVlanData) GetFreeStartVlanId() string`

GetFreeStartVlanId returns the FreeStartVlanId field if non-nil, zero value otherwise.

### GetFreeStartVlanIdOk

`func (o *DataInnerVlanVlanData) GetFreeStartVlanIdOk() (*string, bool)`

GetFreeStartVlanIdOk returns a tuple with the FreeStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeStartVlanId

`func (o *DataInnerVlanVlanData) SetFreeStartVlanId(v string)`

SetFreeStartVlanId sets FreeStartVlanId field to given value.

### HasFreeStartVlanId

`func (o *DataInnerVlanVlanData) HasFreeStartVlanId() bool`

HasFreeStartVlanId returns a boolean if a field has been set.

### GetDomainSupportVxlan

`func (o *DataInnerVlanVlanData) GetDomainSupportVxlan() string`

GetDomainSupportVxlan returns the DomainSupportVxlan field if non-nil, zero value otherwise.

### GetDomainSupportVxlanOk

`func (o *DataInnerVlanVlanData) GetDomainSupportVxlanOk() (*string, bool)`

GetDomainSupportVxlanOk returns a tuple with the DomainSupportVxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSupportVxlan

`func (o *DataInnerVlanVlanData) SetDomainSupportVxlan(v string)`

SetDomainSupportVxlan sets DomainSupportVxlan field to given value.

### HasDomainSupportVxlan

`func (o *DataInnerVlanVlanData) HasDomainSupportVxlan() bool`

HasDomainSupportVxlan returns a boolean if a field has been set.

### GetVlanType

`func (o *DataInnerVlanVlanData) GetVlanType() string`

GetVlanType returns the VlanType field if non-nil, zero value otherwise.

### GetVlanTypeOk

`func (o *DataInnerVlanVlanData) GetVlanTypeOk() (*string, bool)`

GetVlanTypeOk returns a tuple with the VlanType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanType

`func (o *DataInnerVlanVlanData) SetVlanType(v string)`

SetVlanType sets VlanType field to given value.

### HasVlanType

`func (o *DataInnerVlanVlanData) HasVlanType() bool`

HasVlanType returns a boolean if a field has been set.

### GetDomainClassName

`func (o *DataInnerVlanVlanData) GetDomainClassName() string`

GetDomainClassName returns the DomainClassName field if non-nil, zero value otherwise.

### GetDomainClassNameOk

`func (o *DataInnerVlanVlanData) GetDomainClassNameOk() (*string, bool)`

GetDomainClassNameOk returns a tuple with the DomainClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassName

`func (o *DataInnerVlanVlanData) SetDomainClassName(v string)`

SetDomainClassName sets DomainClassName field to given value.

### HasDomainClassName

`func (o *DataInnerVlanVlanData) HasDomainClassName() bool`

HasDomainClassName returns a boolean if a field has been set.

### GetDomainClassParameters

`func (o *DataInnerVlanVlanData) GetDomainClassParameters() []ApiClassParameterOutputEntry`

GetDomainClassParameters returns the DomainClassParameters field if non-nil, zero value otherwise.

### GetDomainClassParametersOk

`func (o *DataInnerVlanVlanData) GetDomainClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetDomainClassParametersOk returns a tuple with the DomainClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassParameters

`func (o *DataInnerVlanVlanData) SetDomainClassParameters(v []ApiClassParameterOutputEntry)`

SetDomainClassParameters sets DomainClassParameters field to given value.

### HasDomainClassParameters

`func (o *DataInnerVlanVlanData) HasDomainClassParameters() bool`

HasDomainClassParameters returns a boolean if a field has been set.

### GetDomainDescription

`func (o *DataInnerVlanVlanData) GetDomainDescription() string`

GetDomainDescription returns the DomainDescription field if non-nil, zero value otherwise.

### GetDomainDescriptionOk

`func (o *DataInnerVlanVlanData) GetDomainDescriptionOk() (*string, bool)`

GetDomainDescriptionOk returns a tuple with the DomainDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDescription

`func (o *DataInnerVlanVlanData) SetDomainDescription(v string)`

SetDomainDescription sets DomainDescription field to given value.

### HasDomainDescription

`func (o *DataInnerVlanVlanData) HasDomainDescription() bool`

HasDomainDescription returns a boolean if a field has been set.

### GetDomainEndVlanId

`func (o *DataInnerVlanVlanData) GetDomainEndVlanId() string`

GetDomainEndVlanId returns the DomainEndVlanId field if non-nil, zero value otherwise.

### GetDomainEndVlanIdOk

`func (o *DataInnerVlanVlanData) GetDomainEndVlanIdOk() (*string, bool)`

GetDomainEndVlanIdOk returns a tuple with the DomainEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEndVlanId

`func (o *DataInnerVlanVlanData) SetDomainEndVlanId(v string)`

SetDomainEndVlanId sets DomainEndVlanId field to given value.

### HasDomainEndVlanId

`func (o *DataInnerVlanVlanData) HasDomainEndVlanId() bool`

HasDomainEndVlanId returns a boolean if a field has been set.

### GetDomainId

`func (o *DataInnerVlanVlanData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *DataInnerVlanVlanData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *DataInnerVlanVlanData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *DataInnerVlanVlanData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *DataInnerVlanVlanData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *DataInnerVlanVlanData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *DataInnerVlanVlanData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *DataInnerVlanVlanData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainStartVlanId

`func (o *DataInnerVlanVlanData) GetDomainStartVlanId() string`

GetDomainStartVlanId returns the DomainStartVlanId field if non-nil, zero value otherwise.

### GetDomainStartVlanIdOk

`func (o *DataInnerVlanVlanData) GetDomainStartVlanIdOk() (*string, bool)`

GetDomainStartVlanIdOk returns a tuple with the DomainStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStartVlanId

`func (o *DataInnerVlanVlanData) SetDomainStartVlanId(v string)`

SetDomainStartVlanId sets DomainStartVlanId field to given value.

### HasDomainStartVlanId

`func (o *DataInnerVlanVlanData) HasDomainStartVlanId() bool`

HasDomainStartVlanId returns a boolean if a field has been set.

### GetRangeClassName

`func (o *DataInnerVlanVlanData) GetRangeClassName() string`

GetRangeClassName returns the RangeClassName field if non-nil, zero value otherwise.

### GetRangeClassNameOk

`func (o *DataInnerVlanVlanData) GetRangeClassNameOk() (*string, bool)`

GetRangeClassNameOk returns a tuple with the RangeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassName

`func (o *DataInnerVlanVlanData) SetRangeClassName(v string)`

SetRangeClassName sets RangeClassName field to given value.

### HasRangeClassName

`func (o *DataInnerVlanVlanData) HasRangeClassName() bool`

HasRangeClassName returns a boolean if a field has been set.

### GetRangeClassParameters

`func (o *DataInnerVlanVlanData) GetRangeClassParameters() []ApiClassParameterOutputEntry`

GetRangeClassParameters returns the RangeClassParameters field if non-nil, zero value otherwise.

### GetRangeClassParametersOk

`func (o *DataInnerVlanVlanData) GetRangeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetRangeClassParametersOk returns a tuple with the RangeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeClassParameters

`func (o *DataInnerVlanVlanData) SetRangeClassParameters(v []ApiClassParameterOutputEntry)`

SetRangeClassParameters sets RangeClassParameters field to given value.

### HasRangeClassParameters

`func (o *DataInnerVlanVlanData) HasRangeClassParameters() bool`

HasRangeClassParameters returns a boolean if a field has been set.

### GetRangeDescription

`func (o *DataInnerVlanVlanData) GetRangeDescription() string`

GetRangeDescription returns the RangeDescription field if non-nil, zero value otherwise.

### GetRangeDescriptionOk

`func (o *DataInnerVlanVlanData) GetRangeDescriptionOk() (*string, bool)`

GetRangeDescriptionOk returns a tuple with the RangeDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeDescription

`func (o *DataInnerVlanVlanData) SetRangeDescription(v string)`

SetRangeDescription sets RangeDescription field to given value.

### HasRangeDescription

`func (o *DataInnerVlanVlanData) HasRangeDescription() bool`

HasRangeDescription returns a boolean if a field has been set.

### GetRangeEndVlanId

`func (o *DataInnerVlanVlanData) GetRangeEndVlanId() string`

GetRangeEndVlanId returns the RangeEndVlanId field if non-nil, zero value otherwise.

### GetRangeEndVlanIdOk

`func (o *DataInnerVlanVlanData) GetRangeEndVlanIdOk() (*string, bool)`

GetRangeEndVlanIdOk returns a tuple with the RangeEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeEndVlanId

`func (o *DataInnerVlanVlanData) SetRangeEndVlanId(v string)`

SetRangeEndVlanId sets RangeEndVlanId field to given value.

### HasRangeEndVlanId

`func (o *DataInnerVlanVlanData) HasRangeEndVlanId() bool`

HasRangeEndVlanId returns a boolean if a field has been set.

### GetRangeId

`func (o *DataInnerVlanVlanData) GetRangeId() string`

GetRangeId returns the RangeId field if non-nil, zero value otherwise.

### GetRangeIdOk

`func (o *DataInnerVlanVlanData) GetRangeIdOk() (*string, bool)`

GetRangeIdOk returns a tuple with the RangeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeId

`func (o *DataInnerVlanVlanData) SetRangeId(v string)`

SetRangeId sets RangeId field to given value.

### HasRangeId

`func (o *DataInnerVlanVlanData) HasRangeId() bool`

HasRangeId returns a boolean if a field has been set.

### GetRangeName

`func (o *DataInnerVlanVlanData) GetRangeName() string`

GetRangeName returns the RangeName field if non-nil, zero value otherwise.

### GetRangeNameOk

`func (o *DataInnerVlanVlanData) GetRangeNameOk() (*string, bool)`

GetRangeNameOk returns a tuple with the RangeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeName

`func (o *DataInnerVlanVlanData) SetRangeName(v string)`

SetRangeName sets RangeName field to given value.

### HasRangeName

`func (o *DataInnerVlanVlanData) HasRangeName() bool`

HasRangeName returns a boolean if a field has been set.

### GetRangeStartVlanId

`func (o *DataInnerVlanVlanData) GetRangeStartVlanId() string`

GetRangeStartVlanId returns the RangeStartVlanId field if non-nil, zero value otherwise.

### GetRangeStartVlanIdOk

`func (o *DataInnerVlanVlanData) GetRangeStartVlanIdOk() (*string, bool)`

GetRangeStartVlanIdOk returns a tuple with the RangeStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRangeStartVlanId

`func (o *DataInnerVlanVlanData) SetRangeStartVlanId(v string)`

SetRangeStartVlanId sets RangeStartVlanId field to given value.

### HasRangeStartVlanId

`func (o *DataInnerVlanVlanData) HasRangeStartVlanId() bool`

HasRangeStartVlanId returns a boolean if a field has been set.

### GetVlanClassName

`func (o *DataInnerVlanVlanData) GetVlanClassName() string`

GetVlanClassName returns the VlanClassName field if non-nil, zero value otherwise.

### GetVlanClassNameOk

`func (o *DataInnerVlanVlanData) GetVlanClassNameOk() (*string, bool)`

GetVlanClassNameOk returns a tuple with the VlanClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanClassName

`func (o *DataInnerVlanVlanData) SetVlanClassName(v string)`

SetVlanClassName sets VlanClassName field to given value.

### HasVlanClassName

`func (o *DataInnerVlanVlanData) HasVlanClassName() bool`

HasVlanClassName returns a boolean if a field has been set.

### GetVlanClassParameters

`func (o *DataInnerVlanVlanData) GetVlanClassParameters() []ApiClassParameterOutputEntry`

GetVlanClassParameters returns the VlanClassParameters field if non-nil, zero value otherwise.

### GetVlanClassParametersOk

`func (o *DataInnerVlanVlanData) GetVlanClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetVlanClassParametersOk returns a tuple with the VlanClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanClassParameters

`func (o *DataInnerVlanVlanData) SetVlanClassParameters(v []ApiClassParameterOutputEntry)`

SetVlanClassParameters sets VlanClassParameters field to given value.

### HasVlanClassParameters

`func (o *DataInnerVlanVlanData) HasVlanClassParameters() bool`

HasVlanClassParameters returns a boolean if a field has been set.

### GetVlanId

`func (o *DataInnerVlanVlanData) GetVlanId() string`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *DataInnerVlanVlanData) GetVlanIdOk() (*string, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *DataInnerVlanVlanData) SetVlanId(v string)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *DataInnerVlanVlanData) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlanName

`func (o *DataInnerVlanVlanData) GetVlanName() string`

GetVlanName returns the VlanName field if non-nil, zero value otherwise.

### GetVlanNameOk

`func (o *DataInnerVlanVlanData) GetVlanNameOk() (*string, bool)`

GetVlanNameOk returns a tuple with the VlanName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanName

`func (o *DataInnerVlanVlanData) SetVlanName(v string)`

SetVlanName sets VlanName field to given value.

### HasVlanName

`func (o *DataInnerVlanVlanData) HasVlanName() bool`

HasVlanName returns a boolean if a field has been set.

### GetVlanVlanId

`func (o *DataInnerVlanVlanData) GetVlanVlanId() string`

GetVlanVlanId returns the VlanVlanId field if non-nil, zero value otherwise.

### GetVlanVlanIdOk

`func (o *DataInnerVlanVlanData) GetVlanVlanIdOk() (*string, bool)`

GetVlanVlanIdOk returns a tuple with the VlanVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanVlanId

`func (o *DataInnerVlanVlanData) SetVlanVlanId(v string)`

SetVlanVlanId sets VlanVlanId field to given value.

### HasVlanVlanId

`func (o *DataInnerVlanVlanData) HasVlanVlanId() bool`

HasVlanVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


