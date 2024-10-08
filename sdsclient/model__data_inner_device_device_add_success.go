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

// checks if the DataInnerDeviceDeviceAddSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerDeviceDeviceAddSuccess{}

// DataInnerDeviceDeviceAddSuccess struct for DataInnerDeviceDeviceAddSuccess
type DataInnerDeviceDeviceAddSuccess struct {
	// The database identifier (ID) of the object you added.
	DeviceId *string `json:"device_id,omitempty"`
}

// NewDataInnerDeviceDeviceAddSuccess instantiates a new DataInnerDeviceDeviceAddSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerDeviceDeviceAddSuccess() *DataInnerDeviceDeviceAddSuccess {
	this := DataInnerDeviceDeviceAddSuccess{}
	return &this
}

// NewDataInnerDeviceDeviceAddSuccessWithDefaults instantiates a new DataInnerDeviceDeviceAddSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerDeviceDeviceAddSuccessWithDefaults() *DataInnerDeviceDeviceAddSuccess {
	this := DataInnerDeviceDeviceAddSuccess{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DataInnerDeviceDeviceAddSuccess) GetDeviceId() string {
	if o == nil || IsNil(o.DeviceId) {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDeviceDeviceAddSuccess) GetDeviceIdOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceId) {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DataInnerDeviceDeviceAddSuccess) HasDeviceId() bool {
	if o != nil && !IsNil(o.DeviceId) {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DataInnerDeviceDeviceAddSuccess) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o DataInnerDeviceDeviceAddSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerDeviceDeviceAddSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DeviceId) {
		toSerialize["device_id"] = o.DeviceId
	}
	return toSerialize, nil
}

type NullableDataInnerDeviceDeviceAddSuccess struct {
	value *DataInnerDeviceDeviceAddSuccess
	isSet bool
}

func (v NullableDataInnerDeviceDeviceAddSuccess) Get() *DataInnerDeviceDeviceAddSuccess {
	return v.value
}

func (v *NullableDataInnerDeviceDeviceAddSuccess) Set(val *DataInnerDeviceDeviceAddSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerDeviceDeviceAddSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerDeviceDeviceAddSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerDeviceDeviceAddSuccess(val *DataInnerDeviceDeviceAddSuccess) *NullableDataInnerDeviceDeviceAddSuccess {
	return &NullableDataInnerDeviceDeviceAddSuccess{value: val, isSet: true}
}

func (v NullableDataInnerDeviceDeviceAddSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerDeviceDeviceAddSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}