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

// checks if the DataInnerVlanRangeDeleteSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerVlanRangeDeleteSuccess{}

// DataInnerVlanRangeDeleteSuccess struct for DataInnerVlanRangeDeleteSuccess
type DataInnerVlanRangeDeleteSuccess struct {
	// The database identifier (ID) of the object you deleted.
	RangeId *string `json:"range_id,omitempty"`
}

// NewDataInnerVlanRangeDeleteSuccess instantiates a new DataInnerVlanRangeDeleteSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerVlanRangeDeleteSuccess() *DataInnerVlanRangeDeleteSuccess {
	this := DataInnerVlanRangeDeleteSuccess{}
	return &this
}

// NewDataInnerVlanRangeDeleteSuccessWithDefaults instantiates a new DataInnerVlanRangeDeleteSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerVlanRangeDeleteSuccessWithDefaults() *DataInnerVlanRangeDeleteSuccess {
	this := DataInnerVlanRangeDeleteSuccess{}
	return &this
}

// GetRangeId returns the RangeId field value if set, zero value otherwise.
func (o *DataInnerVlanRangeDeleteSuccess) GetRangeId() string {
	if o == nil || IsNil(o.RangeId) {
		var ret string
		return ret
	}
	return *o.RangeId
}

// GetRangeIdOk returns a tuple with the RangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerVlanRangeDeleteSuccess) GetRangeIdOk() (*string, bool) {
	if o == nil || IsNil(o.RangeId) {
		return nil, false
	}
	return o.RangeId, true
}

// HasRangeId returns a boolean if a field has been set.
func (o *DataInnerVlanRangeDeleteSuccess) HasRangeId() bool {
	if o != nil && !IsNil(o.RangeId) {
		return true
	}

	return false
}

// SetRangeId gets a reference to the given string and assigns it to the RangeId field.
func (o *DataInnerVlanRangeDeleteSuccess) SetRangeId(v string) {
	o.RangeId = &v
}

func (o DataInnerVlanRangeDeleteSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerVlanRangeDeleteSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.RangeId) {
		toSerialize["range_id"] = o.RangeId
	}
	return toSerialize, nil
}

type NullableDataInnerVlanRangeDeleteSuccess struct {
	value *DataInnerVlanRangeDeleteSuccess
	isSet bool
}

func (v NullableDataInnerVlanRangeDeleteSuccess) Get() *DataInnerVlanRangeDeleteSuccess {
	return v.value
}

func (v *NullableDataInnerVlanRangeDeleteSuccess) Set(val *DataInnerVlanRangeDeleteSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerVlanRangeDeleteSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerVlanRangeDeleteSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerVlanRangeDeleteSuccess(val *DataInnerVlanRangeDeleteSuccess) *NullableDataInnerVlanRangeDeleteSuccess {
	return &NullableDataInnerVlanRangeDeleteSuccess{value: val, isSet: true}
}

func (v NullableDataInnerVlanRangeDeleteSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerVlanRangeDeleteSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}