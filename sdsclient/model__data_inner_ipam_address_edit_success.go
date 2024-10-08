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

// checks if the DataInnerIpamAddressEditSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerIpamAddressEditSuccess{}

// DataInnerIpamAddressEditSuccess struct for DataInnerIpamAddressEditSuccess
type DataInnerIpamAddressEditSuccess struct {
	// The database identifier (ID) of the object you edited.
	AddressId *string `json:"address_id,omitempty"`
}

// NewDataInnerIpamAddressEditSuccess instantiates a new DataInnerIpamAddressEditSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerIpamAddressEditSuccess() *DataInnerIpamAddressEditSuccess {
	this := DataInnerIpamAddressEditSuccess{}
	return &this
}

// NewDataInnerIpamAddressEditSuccessWithDefaults instantiates a new DataInnerIpamAddressEditSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerIpamAddressEditSuccessWithDefaults() *DataInnerIpamAddressEditSuccess {
	this := DataInnerIpamAddressEditSuccess{}
	return &this
}

// GetAddressId returns the AddressId field value if set, zero value otherwise.
func (o *DataInnerIpamAddressEditSuccess) GetAddressId() string {
	if o == nil || IsNil(o.AddressId) {
		var ret string
		return ret
	}
	return *o.AddressId
}

// GetAddressIdOk returns a tuple with the AddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerIpamAddressEditSuccess) GetAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.AddressId) {
		return nil, false
	}
	return o.AddressId, true
}

// HasAddressId returns a boolean if a field has been set.
func (o *DataInnerIpamAddressEditSuccess) HasAddressId() bool {
	if o != nil && !IsNil(o.AddressId) {
		return true
	}

	return false
}

// SetAddressId gets a reference to the given string and assigns it to the AddressId field.
func (o *DataInnerIpamAddressEditSuccess) SetAddressId(v string) {
	o.AddressId = &v
}

func (o DataInnerIpamAddressEditSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerIpamAddressEditSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AddressId) {
		toSerialize["address_id"] = o.AddressId
	}
	return toSerialize, nil
}

type NullableDataInnerIpamAddressEditSuccess struct {
	value *DataInnerIpamAddressEditSuccess
	isSet bool
}

func (v NullableDataInnerIpamAddressEditSuccess) Get() *DataInnerIpamAddressEditSuccess {
	return v.value
}

func (v *NullableDataInnerIpamAddressEditSuccess) Set(val *DataInnerIpamAddressEditSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerIpamAddressEditSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerIpamAddressEditSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerIpamAddressEditSuccess(val *DataInnerIpamAddressEditSuccess) *NullableDataInnerIpamAddressEditSuccess {
	return &NullableDataInnerIpamAddressEditSuccess{value: val, isSet: true}
}

func (v NullableDataInnerIpamAddressEditSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerIpamAddressEditSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
