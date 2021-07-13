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

// IpamNetwork6AddSuccessData struct for IpamNetwork6AddSuccessData
type IpamNetwork6AddSuccessData struct {
	// The database identifier (ID) of the object you added.
	Network6Id *string `json:"network6_id,omitempty"`
}

// NewIpamNetwork6AddSuccessData instantiates a new IpamNetwork6AddSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamNetwork6AddSuccessData() *IpamNetwork6AddSuccessData {
	this := IpamNetwork6AddSuccessData{}
	return &this
}

// NewIpamNetwork6AddSuccessDataWithDefaults instantiates a new IpamNetwork6AddSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamNetwork6AddSuccessDataWithDefaults() *IpamNetwork6AddSuccessData {
	this := IpamNetwork6AddSuccessData{}
	return &this
}

// GetNetwork6Id returns the Network6Id field value if set, zero value otherwise.
func (o *IpamNetwork6AddSuccessData) GetNetwork6Id() string {
	if o == nil || o.Network6Id == nil {
		var ret string
		return ret
	}
	return *o.Network6Id
}

// GetNetwork6IdOk returns a tuple with the Network6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddSuccessData) GetNetwork6IdOk() (*string, bool) {
	if o == nil || o.Network6Id == nil {
		return nil, false
	}
	return o.Network6Id, true
}

// HasNetwork6Id returns a boolean if a field has been set.
func (o *IpamNetwork6AddSuccessData) HasNetwork6Id() bool {
	if o != nil && o.Network6Id != nil {
		return true
	}

	return false
}

// SetNetwork6Id gets a reference to the given string and assigns it to the Network6Id field.
func (o *IpamNetwork6AddSuccessData) SetNetwork6Id(v string) {
	o.Network6Id = &v
}

func (o IpamNetwork6AddSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Network6Id != nil {
		toSerialize["network6_id"] = o.Network6Id
	}
	return json.Marshal(toSerialize)
}

type NullableIpamNetwork6AddSuccessData struct {
	value *IpamNetwork6AddSuccessData
	isSet bool
}

func (v NullableIpamNetwork6AddSuccessData) Get() *IpamNetwork6AddSuccessData {
	return v.value
}

func (v *NullableIpamNetwork6AddSuccessData) Set(val *IpamNetwork6AddSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamNetwork6AddSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamNetwork6AddSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamNetwork6AddSuccessData(val *IpamNetwork6AddSuccessData) *NullableIpamNetwork6AddSuccessData {
	return &NullableIpamNetwork6AddSuccessData{value: val, isSet: true}
}

func (v NullableIpamNetwork6AddSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamNetwork6AddSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


