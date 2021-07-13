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

// IpamAlias6EditSuccessData struct for IpamAlias6EditSuccessData
type IpamAlias6EditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	Alias6Id *string `json:"alias6_id,omitempty"`
}

// NewIpamAlias6EditSuccessData instantiates a new IpamAlias6EditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamAlias6EditSuccessData() *IpamAlias6EditSuccessData {
	this := IpamAlias6EditSuccessData{}
	return &this
}

// NewIpamAlias6EditSuccessDataWithDefaults instantiates a new IpamAlias6EditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamAlias6EditSuccessDataWithDefaults() *IpamAlias6EditSuccessData {
	this := IpamAlias6EditSuccessData{}
	return &this
}

// GetAlias6Id returns the Alias6Id field value if set, zero value otherwise.
func (o *IpamAlias6EditSuccessData) GetAlias6Id() string {
	if o == nil || o.Alias6Id == nil {
		var ret string
		return ret
	}
	return *o.Alias6Id
}

// GetAlias6IdOk returns a tuple with the Alias6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAlias6EditSuccessData) GetAlias6IdOk() (*string, bool) {
	if o == nil || o.Alias6Id == nil {
		return nil, false
	}
	return o.Alias6Id, true
}

// HasAlias6Id returns a boolean if a field has been set.
func (o *IpamAlias6EditSuccessData) HasAlias6Id() bool {
	if o != nil && o.Alias6Id != nil {
		return true
	}

	return false
}

// SetAlias6Id gets a reference to the given string and assigns it to the Alias6Id field.
func (o *IpamAlias6EditSuccessData) SetAlias6Id(v string) {
	o.Alias6Id = &v
}

func (o IpamAlias6EditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Alias6Id != nil {
		toSerialize["alias6_id"] = o.Alias6Id
	}
	return json.Marshal(toSerialize)
}

type NullableIpamAlias6EditSuccessData struct {
	value *IpamAlias6EditSuccessData
	isSet bool
}

func (v NullableIpamAlias6EditSuccessData) Get() *IpamAlias6EditSuccessData {
	return v.value
}

func (v *NullableIpamAlias6EditSuccessData) Set(val *IpamAlias6EditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamAlias6EditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamAlias6EditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamAlias6EditSuccessData(val *IpamAlias6EditSuccessData) *NullableIpamAlias6EditSuccessData {
	return &NullableIpamAlias6EditSuccessData{value: val, isSet: true}
}

func (v NullableIpamAlias6EditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamAlias6EditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

