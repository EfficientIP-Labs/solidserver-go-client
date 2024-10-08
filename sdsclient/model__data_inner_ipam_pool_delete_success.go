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

// checks if the DataInnerIpamPoolDeleteSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerIpamPoolDeleteSuccess{}

// DataInnerIpamPoolDeleteSuccess struct for DataInnerIpamPoolDeleteSuccess
type DataInnerIpamPoolDeleteSuccess struct {
	// The database identifier (ID) of the object you deleted.
	PoolId *string `json:"pool_id,omitempty"`
}

// NewDataInnerIpamPoolDeleteSuccess instantiates a new DataInnerIpamPoolDeleteSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerIpamPoolDeleteSuccess() *DataInnerIpamPoolDeleteSuccess {
	this := DataInnerIpamPoolDeleteSuccess{}
	return &this
}

// NewDataInnerIpamPoolDeleteSuccessWithDefaults instantiates a new DataInnerIpamPoolDeleteSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerIpamPoolDeleteSuccessWithDefaults() *DataInnerIpamPoolDeleteSuccess {
	this := DataInnerIpamPoolDeleteSuccess{}
	return &this
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *DataInnerIpamPoolDeleteSuccess) GetPoolId() string {
	if o == nil || IsNil(o.PoolId) {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerIpamPoolDeleteSuccess) GetPoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.PoolId) {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *DataInnerIpamPoolDeleteSuccess) HasPoolId() bool {
	if o != nil && !IsNil(o.PoolId) {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *DataInnerIpamPoolDeleteSuccess) SetPoolId(v string) {
	o.PoolId = &v
}

func (o DataInnerIpamPoolDeleteSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerIpamPoolDeleteSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PoolId) {
		toSerialize["pool_id"] = o.PoolId
	}
	return toSerialize, nil
}

type NullableDataInnerIpamPoolDeleteSuccess struct {
	value *DataInnerIpamPoolDeleteSuccess
	isSet bool
}

func (v NullableDataInnerIpamPoolDeleteSuccess) Get() *DataInnerIpamPoolDeleteSuccess {
	return v.value
}

func (v *NullableDataInnerIpamPoolDeleteSuccess) Set(val *DataInnerIpamPoolDeleteSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerIpamPoolDeleteSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerIpamPoolDeleteSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerIpamPoolDeleteSuccess(val *DataInnerIpamPoolDeleteSuccess) *NullableDataInnerIpamPoolDeleteSuccess {
	return &NullableDataInnerIpamPoolDeleteSuccess{value: val, isSet: true}
}

func (v NullableDataInnerIpamPoolDeleteSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerIpamPoolDeleteSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}