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

// VlanVlanEditSuccessData struct for VlanVlanEditSuccessData
type VlanVlanEditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	VlanId *string `json:"vlan_id,omitempty"`
}

// NewVlanVlanEditSuccessData instantiates a new VlanVlanEditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVlanVlanEditSuccessData() *VlanVlanEditSuccessData {
	this := VlanVlanEditSuccessData{}
	return &this
}

// NewVlanVlanEditSuccessDataWithDefaults instantiates a new VlanVlanEditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVlanVlanEditSuccessDataWithDefaults() *VlanVlanEditSuccessData {
	this := VlanVlanEditSuccessData{}
	return &this
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *VlanVlanEditSuccessData) GetVlanId() string {
	if o == nil || o.VlanId == nil {
		var ret string
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VlanVlanEditSuccessData) GetVlanIdOk() (*string, bool) {
	if o == nil || o.VlanId == nil {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *VlanVlanEditSuccessData) HasVlanId() bool {
	if o != nil && o.VlanId != nil {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given string and assigns it to the VlanId field.
func (o *VlanVlanEditSuccessData) SetVlanId(v string) {
	o.VlanId = &v
}

func (o VlanVlanEditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.VlanId != nil {
		toSerialize["vlan_id"] = o.VlanId
	}
	return json.Marshal(toSerialize)
}

type NullableVlanVlanEditSuccessData struct {
	value *VlanVlanEditSuccessData
	isSet bool
}

func (v NullableVlanVlanEditSuccessData) Get() *VlanVlanEditSuccessData {
	return v.value
}

func (v *NullableVlanVlanEditSuccessData) Set(val *VlanVlanEditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableVlanVlanEditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableVlanVlanEditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVlanVlanEditSuccessData(val *VlanVlanEditSuccessData) *NullableVlanVlanEditSuccessData {
	return &NullableVlanVlanEditSuccessData{value: val, isSet: true}
}

func (v NullableVlanVlanEditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVlanVlanEditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

