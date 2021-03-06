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

// IpamSpaceAddSuccessData struct for IpamSpaceAddSuccessData
type IpamSpaceAddSuccessData struct {
	// The database identifier (ID) of the object you added.
	SpaceId *string `json:"space_id,omitempty"`
}

// NewIpamSpaceAddSuccessData instantiates a new IpamSpaceAddSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamSpaceAddSuccessData() *IpamSpaceAddSuccessData {
	this := IpamSpaceAddSuccessData{}
	return &this
}

// NewIpamSpaceAddSuccessDataWithDefaults instantiates a new IpamSpaceAddSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamSpaceAddSuccessDataWithDefaults() *IpamSpaceAddSuccessData {
	this := IpamSpaceAddSuccessData{}
	return &this
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *IpamSpaceAddSuccessData) GetSpaceId() string {
	if o == nil || o.SpaceId == nil {
		var ret string
		return ret
	}
	return *o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceAddSuccessData) GetSpaceIdOk() (*string, bool) {
	if o == nil || o.SpaceId == nil {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *IpamSpaceAddSuccessData) HasSpaceId() bool {
	if o != nil && o.SpaceId != nil {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given string and assigns it to the SpaceId field.
func (o *IpamSpaceAddSuccessData) SetSpaceId(v string) {
	o.SpaceId = &v
}

func (o IpamSpaceAddSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpaceId != nil {
		toSerialize["space_id"] = o.SpaceId
	}
	return json.Marshal(toSerialize)
}

type NullableIpamSpaceAddSuccessData struct {
	value *IpamSpaceAddSuccessData
	isSet bool
}

func (v NullableIpamSpaceAddSuccessData) Get() *IpamSpaceAddSuccessData {
	return v.value
}

func (v *NullableIpamSpaceAddSuccessData) Set(val *IpamSpaceAddSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamSpaceAddSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamSpaceAddSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamSpaceAddSuccessData(val *IpamSpaceAddSuccessData) *NullableIpamSpaceAddSuccessData {
	return &NullableIpamSpaceAddSuccessData{value: val, isSet: true}
}

func (v NullableIpamSpaceAddSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamSpaceAddSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


