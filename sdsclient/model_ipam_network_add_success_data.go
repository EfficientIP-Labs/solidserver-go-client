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

// IpamNetworkAddSuccessData struct for IpamNetworkAddSuccessData
type IpamNetworkAddSuccessData struct {
	// The database identifier (ID) of the object you added.
	NetworkId *string `json:"network_id,omitempty"`
}

// NewIpamNetworkAddSuccessData instantiates a new IpamNetworkAddSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamNetworkAddSuccessData() *IpamNetworkAddSuccessData {
	this := IpamNetworkAddSuccessData{}
	return &this
}

// NewIpamNetworkAddSuccessDataWithDefaults instantiates a new IpamNetworkAddSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamNetworkAddSuccessDataWithDefaults() *IpamNetworkAddSuccessData {
	this := IpamNetworkAddSuccessData{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *IpamNetworkAddSuccessData) GetNetworkId() string {
	if o == nil || o.NetworkId == nil {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetworkAddSuccessData) GetNetworkIdOk() (*string, bool) {
	if o == nil || o.NetworkId == nil {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *IpamNetworkAddSuccessData) HasNetworkId() bool {
	if o != nil && o.NetworkId != nil {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *IpamNetworkAddSuccessData) SetNetworkId(v string) {
	o.NetworkId = &v
}

func (o IpamNetworkAddSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NetworkId != nil {
		toSerialize["network_id"] = o.NetworkId
	}
	return json.Marshal(toSerialize)
}

type NullableIpamNetworkAddSuccessData struct {
	value *IpamNetworkAddSuccessData
	isSet bool
}

func (v NullableIpamNetworkAddSuccessData) Get() *IpamNetworkAddSuccessData {
	return v.value
}

func (v *NullableIpamNetworkAddSuccessData) Set(val *IpamNetworkAddSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamNetworkAddSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamNetworkAddSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamNetworkAddSuccessData(val *IpamNetworkAddSuccessData) *NullableIpamNetworkAddSuccessData {
	return &NullableIpamNetworkAddSuccessData{value: val, isSet: true}
}

func (v NullableIpamNetworkAddSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamNetworkAddSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


