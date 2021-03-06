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

// IpamAliasAddSuccessData struct for IpamAliasAddSuccessData
type IpamAliasAddSuccessData struct {
	// The database identifier (ID) of the object you added.
	AliasId *string `json:"alias_id,omitempty"`
}

// NewIpamAliasAddSuccessData instantiates a new IpamAliasAddSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamAliasAddSuccessData() *IpamAliasAddSuccessData {
	this := IpamAliasAddSuccessData{}
	return &this
}

// NewIpamAliasAddSuccessDataWithDefaults instantiates a new IpamAliasAddSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamAliasAddSuccessDataWithDefaults() *IpamAliasAddSuccessData {
	this := IpamAliasAddSuccessData{}
	return &this
}

// GetAliasId returns the AliasId field value if set, zero value otherwise.
func (o *IpamAliasAddSuccessData) GetAliasId() string {
	if o == nil || o.AliasId == nil {
		var ret string
		return ret
	}
	return *o.AliasId
}

// GetAliasIdOk returns a tuple with the AliasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAliasAddSuccessData) GetAliasIdOk() (*string, bool) {
	if o == nil || o.AliasId == nil {
		return nil, false
	}
	return o.AliasId, true
}

// HasAliasId returns a boolean if a field has been set.
func (o *IpamAliasAddSuccessData) HasAliasId() bool {
	if o != nil && o.AliasId != nil {
		return true
	}

	return false
}

// SetAliasId gets a reference to the given string and assigns it to the AliasId field.
func (o *IpamAliasAddSuccessData) SetAliasId(v string) {
	o.AliasId = &v
}

func (o IpamAliasAddSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AliasId != nil {
		toSerialize["alias_id"] = o.AliasId
	}
	return json.Marshal(toSerialize)
}

type NullableIpamAliasAddSuccessData struct {
	value *IpamAliasAddSuccessData
	isSet bool
}

func (v NullableIpamAliasAddSuccessData) Get() *IpamAliasAddSuccessData {
	return v.value
}

func (v *NullableIpamAliasAddSuccessData) Set(val *IpamAliasAddSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamAliasAddSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamAliasAddSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamAliasAddSuccessData(val *IpamAliasAddSuccessData) *NullableIpamAliasAddSuccessData {
	return &NullableIpamAliasAddSuccessData{value: val, isSet: true}
}

func (v NullableIpamAliasAddSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamAliasAddSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


