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

// checks if the DataInnerIpamNetwork6AddSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerIpamNetwork6AddSuccess{}

// DataInnerIpamNetwork6AddSuccess struct for DataInnerIpamNetwork6AddSuccess
type DataInnerIpamNetwork6AddSuccess struct {
	// The database identifier (ID) of the object you added.
	Network6Id *string `json:"network6_id,omitempty"`
}

// NewDataInnerIpamNetwork6AddSuccess instantiates a new DataInnerIpamNetwork6AddSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerIpamNetwork6AddSuccess() *DataInnerIpamNetwork6AddSuccess {
	this := DataInnerIpamNetwork6AddSuccess{}
	return &this
}

// NewDataInnerIpamNetwork6AddSuccessWithDefaults instantiates a new DataInnerIpamNetwork6AddSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerIpamNetwork6AddSuccessWithDefaults() *DataInnerIpamNetwork6AddSuccess {
	this := DataInnerIpamNetwork6AddSuccess{}
	return &this
}

// GetNetwork6Id returns the Network6Id field value if set, zero value otherwise.
func (o *DataInnerIpamNetwork6AddSuccess) GetNetwork6Id() string {
	if o == nil || IsNil(o.Network6Id) {
		var ret string
		return ret
	}
	return *o.Network6Id
}

// GetNetwork6IdOk returns a tuple with the Network6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerIpamNetwork6AddSuccess) GetNetwork6IdOk() (*string, bool) {
	if o == nil || IsNil(o.Network6Id) {
		return nil, false
	}
	return o.Network6Id, true
}

// HasNetwork6Id returns a boolean if a field has been set.
func (o *DataInnerIpamNetwork6AddSuccess) HasNetwork6Id() bool {
	if o != nil && !IsNil(o.Network6Id) {
		return true
	}

	return false
}

// SetNetwork6Id gets a reference to the given string and assigns it to the Network6Id field.
func (o *DataInnerIpamNetwork6AddSuccess) SetNetwork6Id(v string) {
	o.Network6Id = &v
}

func (o DataInnerIpamNetwork6AddSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerIpamNetwork6AddSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Network6Id) {
		toSerialize["network6_id"] = o.Network6Id
	}
	return toSerialize, nil
}

type NullableDataInnerIpamNetwork6AddSuccess struct {
	value *DataInnerIpamNetwork6AddSuccess
	isSet bool
}

func (v NullableDataInnerIpamNetwork6AddSuccess) Get() *DataInnerIpamNetwork6AddSuccess {
	return v.value
}

func (v *NullableDataInnerIpamNetwork6AddSuccess) Set(val *DataInnerIpamNetwork6AddSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerIpamNetwork6AddSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerIpamNetwork6AddSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerIpamNetwork6AddSuccess(val *DataInnerIpamNetwork6AddSuccess) *NullableDataInnerIpamNetwork6AddSuccess {
	return &NullableDataInnerIpamNetwork6AddSuccess{value: val, isSet: true}
}

func (v NullableDataInnerIpamNetwork6AddSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerIpamNetwork6AddSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
