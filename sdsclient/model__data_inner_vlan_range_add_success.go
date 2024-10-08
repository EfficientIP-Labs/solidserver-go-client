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

// checks if the DataInnerVlanRangeAddSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerVlanRangeAddSuccess{}

// DataInnerVlanRangeAddSuccess struct for DataInnerVlanRangeAddSuccess
type DataInnerVlanRangeAddSuccess struct {
	// The database identifier (ID) of the object you added.
	RangeId *string `json:"range_id,omitempty"`
}

// NewDataInnerVlanRangeAddSuccess instantiates a new DataInnerVlanRangeAddSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerVlanRangeAddSuccess() *DataInnerVlanRangeAddSuccess {
	this := DataInnerVlanRangeAddSuccess{}
	return &this
}

// NewDataInnerVlanRangeAddSuccessWithDefaults instantiates a new DataInnerVlanRangeAddSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerVlanRangeAddSuccessWithDefaults() *DataInnerVlanRangeAddSuccess {
	this := DataInnerVlanRangeAddSuccess{}
	return &this
}

// GetRangeId returns the RangeId field value if set, zero value otherwise.
func (o *DataInnerVlanRangeAddSuccess) GetRangeId() string {
	if o == nil || IsNil(o.RangeId) {
		var ret string
		return ret
	}
	return *o.RangeId
}

// GetRangeIdOk returns a tuple with the RangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerVlanRangeAddSuccess) GetRangeIdOk() (*string, bool) {
	if o == nil || IsNil(o.RangeId) {
		return nil, false
	}
	return o.RangeId, true
}

// HasRangeId returns a boolean if a field has been set.
func (o *DataInnerVlanRangeAddSuccess) HasRangeId() bool {
	if o != nil && !IsNil(o.RangeId) {
		return true
	}

	return false
}

// SetRangeId gets a reference to the given string and assigns it to the RangeId field.
func (o *DataInnerVlanRangeAddSuccess) SetRangeId(v string) {
	o.RangeId = &v
}

func (o DataInnerVlanRangeAddSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerVlanRangeAddSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RangeId) {
		toSerialize["range_id"] = o.RangeId
	}
	return toSerialize, nil
}

type NullableDataInnerVlanRangeAddSuccess struct {
	value *DataInnerVlanRangeAddSuccess
	isSet bool
}

func (v NullableDataInnerVlanRangeAddSuccess) Get() *DataInnerVlanRangeAddSuccess {
	return v.value
}

func (v *NullableDataInnerVlanRangeAddSuccess) Set(val *DataInnerVlanRangeAddSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerVlanRangeAddSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerVlanRangeAddSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerVlanRangeAddSuccess(val *DataInnerVlanRangeAddSuccess) *NullableDataInnerVlanRangeAddSuccess {
	return &NullableDataInnerVlanRangeAddSuccess{value: val, isSet: true}
}

func (v NullableDataInnerVlanRangeAddSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerVlanRangeAddSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
