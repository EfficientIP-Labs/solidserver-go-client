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

// checks if the DataInnerDhcpRange6DeleteSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerDhcpRange6DeleteSuccess{}

// DataInnerDhcpRange6DeleteSuccess struct for DataInnerDhcpRange6DeleteSuccess
type DataInnerDhcpRange6DeleteSuccess struct {
	// The database identifier (ID) of the object you deleted.
	Range6Id *string `json:"range6_id,omitempty"`
}

// NewDataInnerDhcpRange6DeleteSuccess instantiates a new DataInnerDhcpRange6DeleteSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerDhcpRange6DeleteSuccess() *DataInnerDhcpRange6DeleteSuccess {
	this := DataInnerDhcpRange6DeleteSuccess{}
	return &this
}

// NewDataInnerDhcpRange6DeleteSuccessWithDefaults instantiates a new DataInnerDhcpRange6DeleteSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerDhcpRange6DeleteSuccessWithDefaults() *DataInnerDhcpRange6DeleteSuccess {
	this := DataInnerDhcpRange6DeleteSuccess{}
	return &this
}

// GetRange6Id returns the Range6Id field value if set, zero value otherwise.
func (o *DataInnerDhcpRange6DeleteSuccess) GetRange6Id() string {
	if o == nil || IsNil(o.Range6Id) {
		var ret string
		return ret
	}
	return *o.Range6Id
}

// GetRange6IdOk returns a tuple with the Range6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpRange6DeleteSuccess) GetRange6IdOk() (*string, bool) {
	if o == nil || IsNil(o.Range6Id) {
		return nil, false
	}
	return o.Range6Id, true
}

// HasRange6Id returns a boolean if a field has been set.
func (o *DataInnerDhcpRange6DeleteSuccess) HasRange6Id() bool {
	if o != nil && !IsNil(o.Range6Id) {
		return true
	}

	return false
}

// SetRange6Id gets a reference to the given string and assigns it to the Range6Id field.
func (o *DataInnerDhcpRange6DeleteSuccess) SetRange6Id(v string) {
	o.Range6Id = &v
}

func (o DataInnerDhcpRange6DeleteSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerDhcpRange6DeleteSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Range6Id) {
		toSerialize["range6_id"] = o.Range6Id
	}
	return toSerialize, nil
}

type NullableDataInnerDhcpRange6DeleteSuccess struct {
	value *DataInnerDhcpRange6DeleteSuccess
	isSet bool
}

func (v NullableDataInnerDhcpRange6DeleteSuccess) Get() *DataInnerDhcpRange6DeleteSuccess {
	return v.value
}

func (v *NullableDataInnerDhcpRange6DeleteSuccess) Set(val *DataInnerDhcpRange6DeleteSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerDhcpRange6DeleteSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerDhcpRange6DeleteSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerDhcpRange6DeleteSuccess(val *DataInnerDhcpRange6DeleteSuccess) *NullableDataInnerDhcpRange6DeleteSuccess {
	return &NullableDataInnerDhcpRange6DeleteSuccess{value: val, isSet: true}
}

func (v NullableDataInnerDhcpRange6DeleteSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerDhcpRange6DeleteSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}