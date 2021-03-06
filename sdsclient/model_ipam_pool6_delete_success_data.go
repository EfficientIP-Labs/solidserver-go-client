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

// IpamPool6DeleteSuccessData struct for IpamPool6DeleteSuccessData
type IpamPool6DeleteSuccessData struct {
	// The database identifier (ID) of the object you deleted.
	Pool6Id *string `json:"pool6_id,omitempty"`
}

// NewIpamPool6DeleteSuccessData instantiates a new IpamPool6DeleteSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamPool6DeleteSuccessData() *IpamPool6DeleteSuccessData {
	this := IpamPool6DeleteSuccessData{}
	return &this
}

// NewIpamPool6DeleteSuccessDataWithDefaults instantiates a new IpamPool6DeleteSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamPool6DeleteSuccessDataWithDefaults() *IpamPool6DeleteSuccessData {
	this := IpamPool6DeleteSuccessData{}
	return &this
}

// GetPool6Id returns the Pool6Id field value if set, zero value otherwise.
func (o *IpamPool6DeleteSuccessData) GetPool6Id() string {
	if o == nil || o.Pool6Id == nil {
		var ret string
		return ret
	}
	return *o.Pool6Id
}

// GetPool6IdOk returns a tuple with the Pool6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6DeleteSuccessData) GetPool6IdOk() (*string, bool) {
	if o == nil || o.Pool6Id == nil {
		return nil, false
	}
	return o.Pool6Id, true
}

// HasPool6Id returns a boolean if a field has been set.
func (o *IpamPool6DeleteSuccessData) HasPool6Id() bool {
	if o != nil && o.Pool6Id != nil {
		return true
	}

	return false
}

// SetPool6Id gets a reference to the given string and assigns it to the Pool6Id field.
func (o *IpamPool6DeleteSuccessData) SetPool6Id(v string) {
	o.Pool6Id = &v
}

func (o IpamPool6DeleteSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Pool6Id != nil {
		toSerialize["pool6_id"] = o.Pool6Id
	}
	return json.Marshal(toSerialize)
}

type NullableIpamPool6DeleteSuccessData struct {
	value *IpamPool6DeleteSuccessData
	isSet bool
}

func (v NullableIpamPool6DeleteSuccessData) Get() *IpamPool6DeleteSuccessData {
	return v.value
}

func (v *NullableIpamPool6DeleteSuccessData) Set(val *IpamPool6DeleteSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamPool6DeleteSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamPool6DeleteSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamPool6DeleteSuccessData(val *IpamPool6DeleteSuccessData) *NullableIpamPool6DeleteSuccessData {
	return &NullableIpamPool6DeleteSuccessData{value: val, isSet: true}
}

func (v NullableIpamPool6DeleteSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamPool6DeleteSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


