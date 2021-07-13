# VlanDomainDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DomainSupportVxlan** | Pointer to **string** | The type of virtual network used by the domain, VXLAN (&lt;b&gt;1&lt;/b&gt;) or VLAN (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**DomainClassName** | Pointer to **string** | The name of the class applied to the VLAN domain, it can be preceded by the class directory. | [optional] 
**DomainClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**DomainDescription** | Pointer to **string** | The description of the VLAN domain. | [optional] 
**DomainEndVlanId** | Pointer to **string** | The VLAN identifier (ID) of the last VLAN in the VLAN domain. | [optional] 
**DomainId** | Pointer to **string** | The database identifier (ID) of the VLAN domain. | [optional] 
**DomainName** | Pointer to **string** | The name of the VLAN domain. | [optional] 
**DomainStartVlanId** | Pointer to **string** | The VLAN identifier (ID) of the first VLAN in the VLAN domain. | [optional] 

## Methods

### NewVlanDomainDataData

`func NewVlanDomainDataData() *VlanDomainDataData`

NewVlanDomainDataData instantiates a new VlanDomainDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVlanDomainDataDataWithDefaults

`func NewVlanDomainDataDataWithDefaults() *VlanDomainDataData`

NewVlanDomainDataDataWithDefaults instantiates a new VlanDomainDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDomainSupportVxlan

`func (o *VlanDomainDataData) GetDomainSupportVxlan() string`

GetDomainSupportVxlan returns the DomainSupportVxlan field if non-nil, zero value otherwise.

### GetDomainSupportVxlanOk

`func (o *VlanDomainDataData) GetDomainSupportVxlanOk() (*string, bool)`

GetDomainSupportVxlanOk returns a tuple with the DomainSupportVxlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainSupportVxlan

`func (o *VlanDomainDataData) SetDomainSupportVxlan(v string)`

SetDomainSupportVxlan sets DomainSupportVxlan field to given value.

### HasDomainSupportVxlan

`func (o *VlanDomainDataData) HasDomainSupportVxlan() bool`

HasDomainSupportVxlan returns a boolean if a field has been set.

### GetDomainClassName

`func (o *VlanDomainDataData) GetDomainClassName() string`

GetDomainClassName returns the DomainClassName field if non-nil, zero value otherwise.

### GetDomainClassNameOk

`func (o *VlanDomainDataData) GetDomainClassNameOk() (*string, bool)`

GetDomainClassNameOk returns a tuple with the DomainClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassName

`func (o *VlanDomainDataData) SetDomainClassName(v string)`

SetDomainClassName sets DomainClassName field to given value.

### HasDomainClassName

`func (o *VlanDomainDataData) HasDomainClassName() bool`

HasDomainClassName returns a boolean if a field has been set.

### GetDomainClassParameters

`func (o *VlanDomainDataData) GetDomainClassParameters() []ApiClassParameterOutputEntry`

GetDomainClassParameters returns the DomainClassParameters field if non-nil, zero value otherwise.

### GetDomainClassParametersOk

`func (o *VlanDomainDataData) GetDomainClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetDomainClassParametersOk returns a tuple with the DomainClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainClassParameters

`func (o *VlanDomainDataData) SetDomainClassParameters(v []ApiClassParameterOutputEntry)`

SetDomainClassParameters sets DomainClassParameters field to given value.

### HasDomainClassParameters

`func (o *VlanDomainDataData) HasDomainClassParameters() bool`

HasDomainClassParameters returns a boolean if a field has been set.

### GetDomainDescription

`func (o *VlanDomainDataData) GetDomainDescription() string`

GetDomainDescription returns the DomainDescription field if non-nil, zero value otherwise.

### GetDomainDescriptionOk

`func (o *VlanDomainDataData) GetDomainDescriptionOk() (*string, bool)`

GetDomainDescriptionOk returns a tuple with the DomainDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainDescription

`func (o *VlanDomainDataData) SetDomainDescription(v string)`

SetDomainDescription sets DomainDescription field to given value.

### HasDomainDescription

`func (o *VlanDomainDataData) HasDomainDescription() bool`

HasDomainDescription returns a boolean if a field has been set.

### GetDomainEndVlanId

`func (o *VlanDomainDataData) GetDomainEndVlanId() string`

GetDomainEndVlanId returns the DomainEndVlanId field if non-nil, zero value otherwise.

### GetDomainEndVlanIdOk

`func (o *VlanDomainDataData) GetDomainEndVlanIdOk() (*string, bool)`

GetDomainEndVlanIdOk returns a tuple with the DomainEndVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainEndVlanId

`func (o *VlanDomainDataData) SetDomainEndVlanId(v string)`

SetDomainEndVlanId sets DomainEndVlanId field to given value.

### HasDomainEndVlanId

`func (o *VlanDomainDataData) HasDomainEndVlanId() bool`

HasDomainEndVlanId returns a boolean if a field has been set.

### GetDomainId

`func (o *VlanDomainDataData) GetDomainId() string`

GetDomainId returns the DomainId field if non-nil, zero value otherwise.

### GetDomainIdOk

`func (o *VlanDomainDataData) GetDomainIdOk() (*string, bool)`

GetDomainIdOk returns a tuple with the DomainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainId

`func (o *VlanDomainDataData) SetDomainId(v string)`

SetDomainId sets DomainId field to given value.

### HasDomainId

`func (o *VlanDomainDataData) HasDomainId() bool`

HasDomainId returns a boolean if a field has been set.

### GetDomainName

`func (o *VlanDomainDataData) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *VlanDomainDataData) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *VlanDomainDataData) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *VlanDomainDataData) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetDomainStartVlanId

`func (o *VlanDomainDataData) GetDomainStartVlanId() string`

GetDomainStartVlanId returns the DomainStartVlanId field if non-nil, zero value otherwise.

### GetDomainStartVlanIdOk

`func (o *VlanDomainDataData) GetDomainStartVlanIdOk() (*string, bool)`

GetDomainStartVlanIdOk returns a tuple with the DomainStartVlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainStartVlanId

`func (o *VlanDomainDataData) SetDomainStartVlanId(v string)`

SetDomainStartVlanId sets DomainStartVlanId field to given value.

### HasDomainStartVlanId

`func (o *VlanDomainDataData) HasDomainStartVlanId() bool`

HasDomainStartVlanId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


