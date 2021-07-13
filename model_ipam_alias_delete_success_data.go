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

// IpamAliasDeleteSuccessData struct for IpamAliasDeleteSuccessData
type IpamAliasDeleteSuccessData struct {
	// The database identifier (ID) of the object you deleted.
	AliasId *string `json:"alias_id,omitempty"`
}

// NewIpamAliasDeleteSuccessData instantiates a new IpamAliasDeleteSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamAliasDeleteSuccessData() *IpamAliasDeleteSuccessData {
	this := IpamAliasDeleteSuccessData{}
	return &this
}

// NewIpamAliasDeleteSuccessDataWithDefaults instantiates a new IpamAliasDeleteSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamAliasDeleteSuccessDataWithDefaults() *IpamAliasDeleteSuccessData {
	this := IpamAliasDeleteSuccessData{}
	return &this
}

// GetAliasId returns the AliasId field value if set, zero value otherwise.
func (o *IpamAliasDeleteSuccessData) GetAliasId() string {
	if o == nil || o.AliasId == nil {
		var ret string
		return ret
	}
	return *o.AliasId
}

// GetAliasIdOk returns a tuple with the AliasId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamAliasDeleteSuccessData) GetAliasIdOk() (*string, bool) {
	if o == nil || o.AliasId == nil {
		return nil, false
	}
	return o.AliasId, true
}

// HasAliasId returns a boolean if a field has been set.
func (o *IpamAliasDeleteSuccessData) HasAliasId() bool {
	if o != nil && o.AliasId != nil {
		return true
	}

	return false
}

// SetAliasId gets a reference to the given string and assigns it to the AliasId field.
func (o *IpamAliasDeleteSuccessData) SetAliasId(v string) {
	o.AliasId = &v
}

func (o IpamAliasDeleteSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AliasId != nil {
		toSerialize["alias_id"] = o.AliasId
	}
	return json.Marshal(toSerialize)
}

type NullableIpamAliasDeleteSuccessData struct {
	value *IpamAliasDeleteSuccessData
	isSet bool
}

func (v NullableIpamAliasDeleteSuccessData) Get() *IpamAliasDeleteSuccessData {
	return v.value
}

func (v *NullableIpamAliasDeleteSuccessData) Set(val *IpamAliasDeleteSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamAliasDeleteSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamAliasDeleteSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamAliasDeleteSuccessData(val *IpamAliasDeleteSuccessData) *NullableIpamAliasDeleteSuccessData {
	return &NullableIpamAliasDeleteSuccessData{value: val, isSet: true}
}

func (v NullableIpamAliasDeleteSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamAliasDeleteSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

