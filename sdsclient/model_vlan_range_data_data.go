/*
 * SOLIDserver API
 *
 * OpenAPI 3.0.2 API definition for SOLIDserver service from EfficientIP.<p>Copyright © 2000-2021 EfficientIP</p><p><em>All specifications and information regarding the products in  this document are subject to change without notice and should not be  construed as a commitment by EfficientIP. EfficientIP assumes no  responsibility or liability for any mistakes or inaccuracies that may appear  in this document. All statements and recommendations in this document are  believed to be accurate but are presented without warranty. Users must take  full responsibility for their application of any product.</em></p><p>Generated (Monday 14th of June 2021 12:30:34 PM)</p>
 *
 * API version: 2.0
 * Contact: support-api@efficientip.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdsclient

import (
	"encoding/json"
)

// VlanRangeDataData struct for VlanRangeDataData
type VlanRangeDataData struct {
	// The type of virtual network used by the domain the range belongs to, VXLAN (<b>1</b>) or VLAN (<b>0</b>).
	DomainSupportVxlan *string `json:"domain_support_vxlan,omitempty"`
	// The name of the class applied to the VLAN domain the object belongs to, it can be preceded by the class directory.
	DomainClassName *string `json:"domain_class_name,omitempty"`
	// #general.output.class_parameters#
	DomainClassParameters *[]ApiClassParameterOutputEntry `json:"domain_class_parameters,omitempty"`
	// The description of the VLAN domain the object belongs to.
	DomainDescription *string `json:"domain_description,omitempty"`
	// The VLAN identifier (ID) of the last VLAN in the VLAN domain the object belongs to.
	DomainEndVlanId *string `json:"domain_end_vlan_id,omitempty"`
	// The database identifier (ID) of the VLAN domain the object belongs to.
	DomainId *string `json:"domain_id,omitempty"`
	// The name of the VLAN domain the object belongs to.
	DomainName *string `json:"domain_name,omitempty"`
	// The VLAN identifier (ID) of the first VLAN in the VLAN domain the object belongs to.
	DomainStartVlanId *string `json:"domain_start_vlan_id,omitempty"`
	// The name of the class applied to the VLAN range, it can be preceded by the class directory.
	RangeClassName *string `json:"range_class_name,omitempty"`
	// #general.output.class_parameters#
	RangeClassParameters *[]ApiClassParameterOutputEntry `json:"range_class_parameters,omitempty"`
	// The description of the VLAN range.
	RangeDescription *string `json:"range_description,omitempty"`
	// The overlapping restriction status of the VLAN range. <b>1</b> indicates that when creating VLANs in the range, their IDs should not overlap.
	RangeDisableOverlapping *string `json:"range_disable_overlapping,omitempty"`
	// The VLAN identifier (ID) of the last VLAN in the VLAN range.
	RangeEndVlanId *string `json:"range_end_vlan_id,omitempty"`
	// The database identifier (ID) of the VLAN range.
	RangeId *string `json:"range_id,omitempty"`
	// The name of the VLAN range.
	RangeName *string `json:"range_name,omitempty"`
	// The VLAN identifier (ID) of the first VLAN in the VLAN range.
	RangeStartVlanId *string `json:"range_start_vlan_id,omitempty"`
}

// NewVlanRangeDataData instantiates a new VlanRangeDataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlanRangeDataData() *VlanRangeDataData {
	this := VlanRangeDataData{}
	return &this
}

// NewVlanRangeDataDataWithDefaults instantiates a new VlanRangeDataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanRangeDataDataWithDefaults() *VlanRangeDataData {
	this := VlanRangeDataData{}
	return &this
}

// GetDomainSupportVxlan returns the DomainSupportVxlan field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetDomainSupportVxlan() string {
	if o == nil || o.DomainSupportVxlan == nil {
		var ret string
		return ret
	}
	return *o.DomainSupportVxlan
}

// GetDomainSupportVxlanOk returns a tuple with the DomainSupportVxlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetDomainSupportVxlanOk() (*string, bool) {
	if o == nil || o.DomainSupportVxlan == nil {
		return nil, false
	}
	return o.DomainSupportVxlan, true
}

// HasDomainSupportVxlan returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasDomainSupportVxlan() bool {
	if o != nil && o.DomainSupportVxlan != nil {
		return true
	}

	return false
}

// SetDomainSupportVxlan gets a reference to the given string and assigns it to the DomainSupportVxlan field.
func (o *VlanRangeDataData) SetDomainSupportVxlan(v string) {
	o.DomainSupportVxlan = &v
}

// GetDomainClassName returns the DomainClassName field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetDomainClassName() string {
	if o == nil || o.DomainClassName == nil {
		var ret string
		return ret
	}
	return *o.DomainClassName
}

// GetDomainClassNameOk returns a tuple with the DomainClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetDomainClassNameOk() (*string, bool) {
	if o == nil || o.DomainClassName == nil {
		return nil, false
	}
	return o.DomainClassName, true
}

// HasDomainClassName returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasDomainClassName() bool {
	if o != nil && o.DomainClassName != nil {
		return true
	}

	return false
}

// SetDomainClassName gets a reference to the given string and assigns it to the DomainClassName field.
func (o *VlanRangeDataData) SetDomainClassName(v string) {
	o.DomainClassName = &v
}

// GetDomainClassParameters returns the DomainClassParameters field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetDomainClassParameters() []ApiClassParameterOutputEntry {
	if o == nil || o.DomainClassParameters == nil {
		var ret []ApiClassParameterOutputEntry
		return ret
	}
	return *o.DomainClassParameters
}

// GetDomainClassParametersOk returns a tuple with the DomainClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetDomainClassParametersOk() (*[]ApiClassParameterOutputEntry, bool) {
	if o == nil || o.DomainClassParameters == nil {
		return nil, false
	}
	return o.DomainClassParameters, true
}

// HasDomainClassParameters returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasDomainClassParameters() bool {
	if o != nil && o.DomainClassParameters != nil {
		return true
	}

	return false
}

// SetDomainClassParameters gets a reference to the given []ApiClassParameterOutputEntry and assigns it to the DomainClassParameters field.
func (o *VlanRangeDataData) SetDomainClassParameters(v []ApiClassParameterOutputEntry) {
	o.DomainClassParameters = &v
}

// GetDomainDescription returns the DomainDescription field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetDomainDescription() string {
	if o == nil || o.DomainDescription == nil {
		var ret string
		return ret
	}
	return *o.DomainDescription
}

// GetDomainDescriptionOk returns a tuple with the DomainDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetDomainDescriptionOk() (*string, bool) {
	if o == nil || o.DomainDescription == nil {
		return nil, false
	}
	return o.DomainDescription, true
}

// HasDomainDescription returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasDomainDescription() bool {
	if o != nil && o.DomainDescription != nil {
		return true
	}

	return false
}

// SetDomainDescription gets a reference to the given string and assigns it to the DomainDescription field.
func (o *VlanRangeDataData) SetDomainDescription(v string) {
	o.DomainDescription = &v
}

// GetDomainEndVlanId returns the DomainEndVlanId field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetDomainEndVlanId() string {
	if o == nil || o.DomainEndVlanId == nil {
		var ret string
		return ret
	}
	return *o.DomainEndVlanId
}

// GetDomainEndVlanIdOk returns a tuple with the DomainEndVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetDomainEndVlanIdOk() (*string, bool) {
	if o == nil || o.DomainEndVlanId == nil {
		return nil, false
	}
	return o.DomainEndVlanId, true
}

// HasDomainEndVlanId returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasDomainEndVlanId() bool {
	if o != nil && o.DomainEndVlanId != nil {
		return true
	}

	return false
}

// SetDomainEndVlanId gets a reference to the given string and assigns it to the DomainEndVlanId field.
func (o *VlanRangeDataData) SetDomainEndVlanId(v string) {
	o.DomainEndVlanId = &v
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetDomainId() string {
	if o == nil || o.DomainId == nil {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetDomainIdOk() (*string, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasDomainId() bool {
	if o != nil && o.DomainId != nil {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *VlanRangeDataData) SetDomainId(v string) {
	o.DomainId = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *VlanRangeDataData) SetDomainName(v string) {
	o.DomainName = &v
}

// GetDomainStartVlanId returns the DomainStartVlanId field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetDomainStartVlanId() string {
	if o == nil || o.DomainStartVlanId == nil {
		var ret string
		return ret
	}
	return *o.DomainStartVlanId
}

// GetDomainStartVlanIdOk returns a tuple with the DomainStartVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetDomainStartVlanIdOk() (*string, bool) {
	if o == nil || o.DomainStartVlanId == nil {
		return nil, false
	}
	return o.DomainStartVlanId, true
}

// HasDomainStartVlanId returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasDomainStartVlanId() bool {
	if o != nil && o.DomainStartVlanId != nil {
		return true
	}

	return false
}

// SetDomainStartVlanId gets a reference to the given string and assigns it to the DomainStartVlanId field.
func (o *VlanRangeDataData) SetDomainStartVlanId(v string) {
	o.DomainStartVlanId = &v
}

// GetRangeClassName returns the RangeClassName field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetRangeClassName() string {
	if o == nil || o.RangeClassName == nil {
		var ret string
		return ret
	}
	return *o.RangeClassName
}

// GetRangeClassNameOk returns a tuple with the RangeClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetRangeClassNameOk() (*string, bool) {
	if o == nil || o.RangeClassName == nil {
		return nil, false
	}
	return o.RangeClassName, true
}

// HasRangeClassName returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasRangeClassName() bool {
	if o != nil && o.RangeClassName != nil {
		return true
	}

	return false
}

// SetRangeClassName gets a reference to the given string and assigns it to the RangeClassName field.
func (o *VlanRangeDataData) SetRangeClassName(v string) {
	o.RangeClassName = &v
}

// GetRangeClassParameters returns the RangeClassParameters field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetRangeClassParameters() []ApiClassParameterOutputEntry {
	if o == nil || o.RangeClassParameters == nil {
		var ret []ApiClassParameterOutputEntry
		return ret
	}
	return *o.RangeClassParameters
}

// GetRangeClassParametersOk returns a tuple with the RangeClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetRangeClassParametersOk() (*[]ApiClassParameterOutputEntry, bool) {
	if o == nil || o.RangeClassParameters == nil {
		return nil, false
	}
	return o.RangeClassParameters, true
}

// HasRangeClassParameters returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasRangeClassParameters() bool {
	if o != nil && o.RangeClassParameters != nil {
		return true
	}

	return false
}

// SetRangeClassParameters gets a reference to the given []ApiClassParameterOutputEntry and assigns it to the RangeClassParameters field.
func (o *VlanRangeDataData) SetRangeClassParameters(v []ApiClassParameterOutputEntry) {
	o.RangeClassParameters = &v
}

// GetRangeDescription returns the RangeDescription field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetRangeDescription() string {
	if o == nil || o.RangeDescription == nil {
		var ret string
		return ret
	}
	return *o.RangeDescription
}

// GetRangeDescriptionOk returns a tuple with the RangeDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetRangeDescriptionOk() (*string, bool) {
	if o == nil || o.RangeDescription == nil {
		return nil, false
	}
	return o.RangeDescription, true
}

// HasRangeDescription returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasRangeDescription() bool {
	if o != nil && o.RangeDescription != nil {
		return true
	}

	return false
}

// SetRangeDescription gets a reference to the given string and assigns it to the RangeDescription field.
func (o *VlanRangeDataData) SetRangeDescription(v string) {
	o.RangeDescription = &v
}

// GetRangeDisableOverlapping returns the RangeDisableOverlapping field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetRangeDisableOverlapping() string {
	if o == nil || o.RangeDisableOverlapping == nil {
		var ret string
		return ret
	}
	return *o.RangeDisableOverlapping
}

// GetRangeDisableOverlappingOk returns a tuple with the RangeDisableOverlapping field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetRangeDisableOverlappingOk() (*string, bool) {
	if o == nil || o.RangeDisableOverlapping == nil {
		return nil, false
	}
	return o.RangeDisableOverlapping, true
}

// HasRangeDisableOverlapping returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasRangeDisableOverlapping() bool {
	if o != nil && o.RangeDisableOverlapping != nil {
		return true
	}

	return false
}

// SetRangeDisableOverlapping gets a reference to the given string and assigns it to the RangeDisableOverlapping field.
func (o *VlanRangeDataData) SetRangeDisableOverlapping(v string) {
	o.RangeDisableOverlapping = &v
}

// GetRangeEndVlanId returns the RangeEndVlanId field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetRangeEndVlanId() string {
	if o == nil || o.RangeEndVlanId == nil {
		var ret string
		return ret
	}
	return *o.RangeEndVlanId
}

// GetRangeEndVlanIdOk returns a tuple with the RangeEndVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetRangeEndVlanIdOk() (*string, bool) {
	if o == nil || o.RangeEndVlanId == nil {
		return nil, false
	}
	return o.RangeEndVlanId, true
}

// HasRangeEndVlanId returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasRangeEndVlanId() bool {
	if o != nil && o.RangeEndVlanId != nil {
		return true
	}

	return false
}

// SetRangeEndVlanId gets a reference to the given string and assigns it to the RangeEndVlanId field.
func (o *VlanRangeDataData) SetRangeEndVlanId(v string) {
	o.RangeEndVlanId = &v
}

// GetRangeId returns the RangeId field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetRangeId() string {
	if o == nil || o.RangeId == nil {
		var ret string
		return ret
	}
	return *o.RangeId
}

// GetRangeIdOk returns a tuple with the RangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetRangeIdOk() (*string, bool) {
	if o == nil || o.RangeId == nil {
		return nil, false
	}
	return o.RangeId, true
}

// HasRangeId returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasRangeId() bool {
	if o != nil && o.RangeId != nil {
		return true
	}

	return false
}

// SetRangeId gets a reference to the given string and assigns it to the RangeId field.
func (o *VlanRangeDataData) SetRangeId(v string) {
	o.RangeId = &v
}

// GetRangeName returns the RangeName field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetRangeName() string {
	if o == nil || o.RangeName == nil {
		var ret string
		return ret
	}
	return *o.RangeName
}

// GetRangeNameOk returns a tuple with the RangeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetRangeNameOk() (*string, bool) {
	if o == nil || o.RangeName == nil {
		return nil, false
	}
	return o.RangeName, true
}

// HasRangeName returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasRangeName() bool {
	if o != nil && o.RangeName != nil {
		return true
	}

	return false
}

// SetRangeName gets a reference to the given string and assigns it to the RangeName field.
func (o *VlanRangeDataData) SetRangeName(v string) {
	o.RangeName = &v
}

// GetRangeStartVlanId returns the RangeStartVlanId field value if set, zero value otherwise.
func (o *VlanRangeDataData) GetRangeStartVlanId() string {
	if o == nil || o.RangeStartVlanId == nil {
		var ret string
		return ret
	}
	return *o.RangeStartVlanId
}

// GetRangeStartVlanIdOk returns a tuple with the RangeStartVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanRangeDataData) GetRangeStartVlanIdOk() (*string, bool) {
	if o == nil || o.RangeStartVlanId == nil {
		return nil, false
	}
	return o.RangeStartVlanId, true
}

// HasRangeStartVlanId returns a boolean if a field has been set.
func (o *VlanRangeDataData) HasRangeStartVlanId() bool {
	if o != nil && o.RangeStartVlanId != nil {
		return true
	}

	return false
}

// SetRangeStartVlanId gets a reference to the given string and assigns it to the RangeStartVlanId field.
func (o *VlanRangeDataData) SetRangeStartVlanId(v string) {
	o.RangeStartVlanId = &v
}

func (o VlanRangeDataData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainSupportVxlan != nil {
		toSerialize["domain_support_vxlan"] = o.DomainSupportVxlan
	}
	if o.DomainClassName != nil {
		toSerialize["domain_class_name"] = o.DomainClassName
	}
	if o.DomainClassParameters != nil {
		toSerialize["domain_class_parameters"] = o.DomainClassParameters
	}
	if o.DomainDescription != nil {
		toSerialize["domain_description"] = o.DomainDescription
	}
	if o.DomainEndVlanId != nil {
		toSerialize["domain_end_vlan_id"] = o.DomainEndVlanId
	}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	if o.DomainName != nil {
		toSerialize["domain_name"] = o.DomainName
	}
	if o.DomainStartVlanId != nil {
		toSerialize["domain_start_vlan_id"] = o.DomainStartVlanId
	}
	if o.RangeClassName != nil {
		toSerialize["range_class_name"] = o.RangeClassName
	}
	if o.RangeClassParameters != nil {
		toSerialize["range_class_parameters"] = o.RangeClassParameters
	}
	if o.RangeDescription != nil {
		toSerialize["range_description"] = o.RangeDescription
	}
	if o.RangeDisableOverlapping != nil {
		toSerialize["range_disable_overlapping"] = o.RangeDisableOverlapping
	}
	if o.RangeEndVlanId != nil {
		toSerialize["range_end_vlan_id"] = o.RangeEndVlanId
	}
	if o.RangeId != nil {
		toSerialize["range_id"] = o.RangeId
	}
	if o.RangeName != nil {
		toSerialize["range_name"] = o.RangeName
	}
	if o.RangeStartVlanId != nil {
		toSerialize["range_start_vlan_id"] = o.RangeStartVlanId
	}
	return json.Marshal(toSerialize)
}

type NullableVlanRangeDataData struct {
	value *VlanRangeDataData
	isSet bool
}

func (v NullableVlanRangeDataData) Get() *VlanRangeDataData {
	return v.value
}

func (v *NullableVlanRangeDataData) Set(val *VlanRangeDataData) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanRangeDataData) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanRangeDataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanRangeDataData(val *VlanRangeDataData) *NullableVlanRangeDataData {
	return &NullableVlanRangeDataData{value: val, isSet: true}
}

func (v NullableVlanRangeDataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanRangeDataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


