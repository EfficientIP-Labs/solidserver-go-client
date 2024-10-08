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

// checks if the DataInnerIpamAlias6DeleteSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerIpamAlias6DeleteSuccess{}

// DataInnerIpamAlias6DeleteSuccess struct for DataInnerIpamAlias6DeleteSuccess
type DataInnerIpamAlias6DeleteSuccess struct {
	// The database identifier (ID) of the object you deleted.
	Alias6Id *string `json:"alias6_id,omitempty"`
}

// NewDataInnerIpamAlias6DeleteSuccess instantiates a new DataInnerIpamAlias6DeleteSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerIpamAlias6DeleteSuccess() *DataInnerIpamAlias6DeleteSuccess {
	this := DataInnerIpamAlias6DeleteSuccess{}
	return &this
}

// NewDataInnerIpamAlias6DeleteSuccessWithDefaults instantiates a new DataInnerIpamAlias6DeleteSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerIpamAlias6DeleteSuccessWithDefaults() *DataInnerIpamAlias6DeleteSuccess {
	this := DataInnerIpamAlias6DeleteSuccess{}
	return &this
}

// GetAlias6Id returns the Alias6Id field value if set, zero value otherwise.
func (o *DataInnerIpamAlias6DeleteSuccess) GetAlias6Id() string {
	if o == nil || IsNil(o.Alias6Id) {
		var ret string
		return ret
	}
	return *o.Alias6Id
}

// GetAlias6IdOk returns a tuple with the Alias6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerIpamAlias6DeleteSuccess) GetAlias6IdOk() (*string, bool) {
	if o == nil || IsNil(o.Alias6Id) {
		return nil, false
	}
	return o.Alias6Id, true
}

// HasAlias6Id returns a boolean if a field has been set.
func (o *DataInnerIpamAlias6DeleteSuccess) HasAlias6Id() bool {
	if o != nil && !IsNil(o.Alias6Id) {
		return true
	}

	return false
}

// SetAlias6Id gets a reference to the given string and assigns it to the Alias6Id field.
func (o *DataInnerIpamAlias6DeleteSuccess) SetAlias6Id(v string) {
	o.Alias6Id = &v
}

func (o DataInnerIpamAlias6DeleteSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerIpamAlias6DeleteSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias6Id) {
		toSerialize["alias6_id"] = o.Alias6Id
	}
	return toSerialize, nil
}

type NullableDataInnerIpamAlias6DeleteSuccess struct {
	value *DataInnerIpamAlias6DeleteSuccess
	isSet bool
}

func (v NullableDataInnerIpamAlias6DeleteSuccess) Get() *DataInnerIpamAlias6DeleteSuccess {
	return v.value
}

func (v *NullableDataInnerIpamAlias6DeleteSuccess) Set(val *DataInnerIpamAlias6DeleteSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerIpamAlias6DeleteSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerIpamAlias6DeleteSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerIpamAlias6DeleteSuccess(val *DataInnerIpamAlias6DeleteSuccess) *NullableDataInnerIpamAlias6DeleteSuccess {
	return &NullableDataInnerIpamAlias6DeleteSuccess{value: val, isSet: true}
}

func (v NullableDataInnerIpamAlias6DeleteSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerIpamAlias6DeleteSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
