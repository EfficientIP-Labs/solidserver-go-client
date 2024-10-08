/*
SOLIDserver API

OpenAPI 3.0.2 API definition for SOLIDserver service from EfficientIP.<p>Copyright © 2000-2024 EfficientIP</p><p><em>All specifications and information regarding the products in this document are subject to change without notice and should not be construed as a commitment by EfficientIP. EfficientIP assumes no responsibility or liability for any mistakes or inaccuracies that may appear in this document. All statements and recommendations in this document are believed to be accurate but are presented without warranty. Users must take full responsibility for their application of any product.</em></p><p><em>This document aims at detailing EfficientIP proprietary solutions. As our solutions rely on several third-party products, created by other companies or organizations, it may redirect readers to third-party websites and documentation for further information. In such a case, EfficientIP cannot be liable or expected to provide said information on products they do maintain or created.</em></p><p>Generated (Friday 4th of October 2024 03:41:11 PM)</p>

API version: 2.0
Contact: support-api@efficientip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdsclient

import (
	"encoding/json"
)

// checks if the DataInnerIpamAliasEditSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerIpamAliasEditSuccess{}

// DataInnerIpamAliasEditSuccess struct for DataInnerIpamAliasEditSuccess
type DataInnerIpamAliasEditSuccess struct {
	// The database identifier (ID) of the object you edited.
	AliasId *string `json:"alias_id,omitempty"`
}

// NewDataInnerIpamAliasEditSuccess instantiates a new DataInnerIpamAliasEditSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerIpamAliasEditSuccess() *DataInnerIpamAliasEditSuccess {
	this := DataInnerIpamAliasEditSuccess{}
	return &this
}

// NewDataInnerIpamAliasEditSuccessWithDefaults instantiates a new DataInnerIpamAliasEditSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerIpamAliasEditSuccessWithDefaults() *DataInnerIpamAliasEditSuccess {
	this := DataInnerIpamAliasEditSuccess{}
	return &this
}

// GetAliasId returns the AliasId field value if set, zero value otherwise.
func (o *DataInnerIpamAliasEditSuccess) GetAliasId() string {
	if o == nil || IsNil(o.AliasId) {
		var ret string
		return ret
	}
	return *o.AliasId
}

// GetAliasIdOk returns a tuple with the AliasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerIpamAliasEditSuccess) GetAliasIdOk() (*string, bool) {
	if o == nil || IsNil(o.AliasId) {
		return nil, false
	}
	return o.AliasId, true
}

// HasAliasId returns a boolean if a field has been set.
func (o *DataInnerIpamAliasEditSuccess) HasAliasId() bool {
	if o != nil && !IsNil(o.AliasId) {
		return true
	}

	return false
}

// SetAliasId gets a reference to the given string and assigns it to the AliasId field.
func (o *DataInnerIpamAliasEditSuccess) SetAliasId(v string) {
	o.AliasId = &v
}

func (o DataInnerIpamAliasEditSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerIpamAliasEditSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AliasId) {
		toSerialize["alias_id"] = o.AliasId
	}
	return toSerialize, nil
}

type NullableDataInnerIpamAliasEditSuccess struct {
	value *DataInnerIpamAliasEditSuccess
	isSet bool
}

func (v NullableDataInnerIpamAliasEditSuccess) Get() *DataInnerIpamAliasEditSuccess {
	return v.value
}

func (v *NullableDataInnerIpamAliasEditSuccess) Set(val *DataInnerIpamAliasEditSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerIpamAliasEditSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerIpamAliasEditSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerIpamAliasEditSuccess(val *DataInnerIpamAliasEditSuccess) *NullableDataInnerIpamAliasEditSuccess {
	return &NullableDataInnerIpamAliasEditSuccess{value: val, isSet: true}
}

func (v NullableDataInnerIpamAliasEditSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerIpamAliasEditSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
