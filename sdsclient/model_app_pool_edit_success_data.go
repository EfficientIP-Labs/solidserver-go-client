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

// AppPoolEditSuccessData struct for AppPoolEditSuccessData
type AppPoolEditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	PoolId *string `json:"pool_id,omitempty"`
}

// NewAppPoolEditSuccessData instantiates a new AppPoolEditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPoolEditSuccessData() *AppPoolEditSuccessData {
	this := AppPoolEditSuccessData{}
	return &this
}

// NewAppPoolEditSuccessDataWithDefaults instantiates a new AppPoolEditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPoolEditSuccessDataWithDefaults() *AppPoolEditSuccessData {
	this := AppPoolEditSuccessData{}
	return &this
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *AppPoolEditSuccessData) GetPoolId() string {
	if o == nil || o.PoolId == nil {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPoolEditSuccessData) GetPoolIdOk() (*string, bool) {
	if o == nil || o.PoolId == nil {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *AppPoolEditSuccessData) HasPoolId() bool {
	if o != nil && o.PoolId != nil {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *AppPoolEditSuccessData) SetPoolId(v string) {
	o.PoolId = &v
}

func (o AppPoolEditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.PoolId != nil {
		toSerialize["pool_id"] = o.PoolId
	}
	return json.Marshal(toSerialize)
}

type NullableAppPoolEditSuccessData struct {
	value *AppPoolEditSuccessData
	isSet bool
}

func (v NullableAppPoolEditSuccessData) Get() *AppPoolEditSuccessData {
	return v.value
}

func (v *NullableAppPoolEditSuccessData) Set(val *AppPoolEditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPoolEditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPoolEditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPoolEditSuccessData(val *AppPoolEditSuccessData) *NullableAppPoolEditSuccessData {
	return &NullableAppPoolEditSuccessData{value: val, isSet: true}
}

func (v NullableAppPoolEditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPoolEditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


