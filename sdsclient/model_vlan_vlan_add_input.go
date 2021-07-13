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

// VlanVlanAddInput struct for VlanVlanAddInput
type VlanVlanAddInput struct {
	// The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice.
	DomainId *int32 `json:"domain_id,omitempty"`
	// The name of the VLAN domain.
	DomainName *string `json:"domain_name,omitempty"`
	// The VLAN identifier (ID) of the VLAN, a unique numeric key value within a VLAN domain. Use the ID to specify which VLAN to edit.
	VlanVlanId *int32 `json:"vlan_vlan_id,omitempty"`
	// The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice.
	RangeId *int32 `json:"range_id,omitempty"`
	// The name of the VLAN range.
	RangeName *string `json:"range_name,omitempty"`
	// The name of the VLAN, each VLAN must have a unique name.
	VlanName *string `json:"vlan_name,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete *[]string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	VlanClassName *string `json:"vlan_class_name,omitempty"`
	// class parameters in json format
	VlanClassParameters *[]ApiClassParameterInputEntry `json:"vlan_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewVlanVlanAddInput instantiates a new VlanVlanAddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlanVlanAddInput() *VlanVlanAddInput {
	this := VlanVlanAddInput{}
	return &this
}

// NewVlanVlanAddInputWithDefaults instantiates a new VlanVlanAddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanVlanAddInputWithDefaults() *VlanVlanAddInput {
	this := VlanVlanAddInput{}
	return &this
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetDomainId() int32 {
	if o == nil || o.DomainId == nil {
		var ret int32
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetDomainIdOk() (*int32, bool) {
	if o == nil || o.DomainId == nil {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasDomainId() bool {
	if o != nil && o.DomainId != nil {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given int32 and assigns it to the DomainId field.
func (o *VlanVlanAddInput) SetDomainId(v int32) {
	o.DomainId = &v
}

// GetDomainName returns the DomainName field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetDomainName() string {
	if o == nil || o.DomainName == nil {
		var ret string
		return ret
	}
	return *o.DomainName
}

// GetDomainNameOk returns a tuple with the DomainName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetDomainNameOk() (*string, bool) {
	if o == nil || o.DomainName == nil {
		return nil, false
	}
	return o.DomainName, true
}

// HasDomainName returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasDomainName() bool {
	if o != nil && o.DomainName != nil {
		return true
	}

	return false
}

// SetDomainName gets a reference to the given string and assigns it to the DomainName field.
func (o *VlanVlanAddInput) SetDomainName(v string) {
	o.DomainName = &v
}

// GetVlanVlanId returns the VlanVlanId field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetVlanVlanId() int32 {
	if o == nil || o.VlanVlanId == nil {
		var ret int32
		return ret
	}
	return *o.VlanVlanId
}

// GetVlanVlanIdOk returns a tuple with the VlanVlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetVlanVlanIdOk() (*int32, bool) {
	if o == nil || o.VlanVlanId == nil {
		return nil, false
	}
	return o.VlanVlanId, true
}

// HasVlanVlanId returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasVlanVlanId() bool {
	if o != nil && o.VlanVlanId != nil {
		return true
	}

	return false
}

// SetVlanVlanId gets a reference to the given int32 and assigns it to the VlanVlanId field.
func (o *VlanVlanAddInput) SetVlanVlanId(v int32) {
	o.VlanVlanId = &v
}

// GetRangeId returns the RangeId field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetRangeId() int32 {
	if o == nil || o.RangeId == nil {
		var ret int32
		return ret
	}
	return *o.RangeId
}

// GetRangeIdOk returns a tuple with the RangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetRangeIdOk() (*int32, bool) {
	if o == nil || o.RangeId == nil {
		return nil, false
	}
	return o.RangeId, true
}

// HasRangeId returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasRangeId() bool {
	if o != nil && o.RangeId != nil {
		return true
	}

	return false
}

// SetRangeId gets a reference to the given int32 and assigns it to the RangeId field.
func (o *VlanVlanAddInput) SetRangeId(v int32) {
	o.RangeId = &v
}

// GetRangeName returns the RangeName field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetRangeName() string {
	if o == nil || o.RangeName == nil {
		var ret string
		return ret
	}
	return *o.RangeName
}

// GetRangeNameOk returns a tuple with the RangeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetRangeNameOk() (*string, bool) {
	if o == nil || o.RangeName == nil {
		return nil, false
	}
	return o.RangeName, true
}

// HasRangeName returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasRangeName() bool {
	if o != nil && o.RangeName != nil {
		return true
	}

	return false
}

// SetRangeName gets a reference to the given string and assigns it to the RangeName field.
func (o *VlanVlanAddInput) SetRangeName(v string) {
	o.RangeName = &v
}

// GetVlanName returns the VlanName field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetVlanName() string {
	if o == nil || o.VlanName == nil {
		var ret string
		return ret
	}
	return *o.VlanName
}

// GetVlanNameOk returns a tuple with the VlanName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetVlanNameOk() (*string, bool) {
	if o == nil || o.VlanName == nil {
		return nil, false
	}
	return o.VlanName, true
}

// HasVlanName returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasVlanName() bool {
	if o != nil && o.VlanName != nil {
		return true
	}

	return false
}

// SetVlanName gets a reference to the given string and assigns it to the VlanName field.
func (o *VlanVlanAddInput) SetVlanName(v string) {
	o.VlanName = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetClassParametersToDelete() []string {
	if o == nil || o.ClassParametersToDelete == nil {
		var ret []string
		return ret
	}
	return *o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetClassParametersToDeleteOk() (*[]string, bool) {
	if o == nil || o.ClassParametersToDelete == nil {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasClassParametersToDelete() bool {
	if o != nil && o.ClassParametersToDelete != nil {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *VlanVlanAddInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = &v
}

// GetVlanClassName returns the VlanClassName field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetVlanClassName() string {
	if o == nil || o.VlanClassName == nil {
		var ret string
		return ret
	}
	return *o.VlanClassName
}

// GetVlanClassNameOk returns a tuple with the VlanClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetVlanClassNameOk() (*string, bool) {
	if o == nil || o.VlanClassName == nil {
		return nil, false
	}
	return o.VlanClassName, true
}

// HasVlanClassName returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasVlanClassName() bool {
	if o != nil && o.VlanClassName != nil {
		return true
	}

	return false
}

// SetVlanClassName gets a reference to the given string and assigns it to the VlanClassName field.
func (o *VlanVlanAddInput) SetVlanClassName(v string) {
	o.VlanClassName = &v
}

// GetVlanClassParameters returns the VlanClassParameters field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetVlanClassParameters() []ApiClassParameterInputEntry {
	if o == nil || o.VlanClassParameters == nil {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return *o.VlanClassParameters
}

// GetVlanClassParametersOk returns a tuple with the VlanClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetVlanClassParametersOk() (*[]ApiClassParameterInputEntry, bool) {
	if o == nil || o.VlanClassParameters == nil {
		return nil, false
	}
	return o.VlanClassParameters, true
}

// HasVlanClassParameters returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasVlanClassParameters() bool {
	if o != nil && o.VlanClassParameters != nil {
		return true
	}

	return false
}

// SetVlanClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the VlanClassParameters field.
func (o *VlanVlanAddInput) SetVlanClassParameters(v []ApiClassParameterInputEntry) {
	o.VlanClassParameters = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *VlanVlanAddInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanAddInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *VlanVlanAddInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *VlanVlanAddInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o VlanVlanAddInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DomainId != nil {
		toSerialize["domain_id"] = o.DomainId
	}
	if o.DomainName != nil {
		toSerialize["domain_name"] = o.DomainName
	}
	if o.VlanVlanId != nil {
		toSerialize["vlan_vlan_id"] = o.VlanVlanId
	}
	if o.RangeId != nil {
		toSerialize["range_id"] = o.RangeId
	}
	if o.RangeName != nil {
		toSerialize["range_name"] = o.RangeName
	}
	if o.VlanName != nil {
		toSerialize["vlan_name"] = o.VlanName
	}
	if o.ClassParametersToDelete != nil {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if o.VlanClassName != nil {
		toSerialize["vlan_class_name"] = o.VlanClassName
	}
	if o.VlanClassParameters != nil {
		toSerialize["vlan_class_parameters"] = o.VlanClassParameters
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableVlanVlanAddInput struct {
	value *VlanVlanAddInput
	isSet bool
}

func (v NullableVlanVlanAddInput) Get() *VlanVlanAddInput {
	return v.value
}

func (v *NullableVlanVlanAddInput) Set(val *VlanVlanAddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanVlanAddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanVlanAddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanVlanAddInput(val *VlanVlanAddInput) *NullableVlanVlanAddInput {
	return &NullableVlanVlanAddInput{value: val, isSet: true}
}

func (v NullableVlanVlanAddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanVlanAddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


