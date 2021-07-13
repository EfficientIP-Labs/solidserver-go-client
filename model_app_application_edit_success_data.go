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

// AppApplicationEditSuccessData struct for AppApplicationEditSuccessData
type AppApplicationEditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	ApplicationId *string `json:"application_id,omitempty"`
}

// NewAppApplicationEditSuccessData instantiates a new AppApplicationEditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppApplicationEditSuccessData() *AppApplicationEditSuccessData {
	this := AppApplicationEditSuccessData{}
	return &this
}

// NewAppApplicationEditSuccessDataWithDefaults instantiates a new AppApplicationEditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppApplicationEditSuccessDataWithDefaults() *AppApplicationEditSuccessData {
	this := AppApplicationEditSuccessData{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *AppApplicationEditSuccessData) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditSuccessData) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AppApplicationEditSuccessData) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *AppApplicationEditSuccessData) SetApplicationId(v string) {
	o.ApplicationId = &v
}

func (o AppApplicationEditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	return json.Marshal(toSerialize)
}

type NullableAppApplicationEditSuccessData struct {
	value *AppApplicationEditSuccessData
	isSet bool
}

func (v NullableAppApplicationEditSuccessData) Get() *AppApplicationEditSuccessData {
	return v.value
}

func (v *NullableAppApplicationEditSuccessData) Set(val *AppApplicationEditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppApplicationEditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppApplicationEditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppApplicationEditSuccessData(val *AppApplicationEditSuccessData) *NullableAppApplicationEditSuccessData {
	return &NullableAppApplicationEditSuccessData{value: val, isSet: true}
}

func (v NullableAppApplicationEditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppApplicationEditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


