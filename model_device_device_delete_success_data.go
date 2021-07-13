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

// DeviceDeviceDeleteSuccessData struct for DeviceDeviceDeleteSuccessData
type DeviceDeviceDeleteSuccessData struct {
	// The database identifier (ID) of the object you deleted.
	DeviceId *string `json:"device_id,omitempty"`
}

// NewDeviceDeviceDeleteSuccessData instantiates a new DeviceDeviceDeleteSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceDeviceDeleteSuccessData() *DeviceDeviceDeleteSuccessData {
	this := DeviceDeviceDeleteSuccessData{}
	return &this
}

// NewDeviceDeviceDeleteSuccessDataWithDefaults instantiates a new DeviceDeviceDeleteSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceDeviceDeleteSuccessDataWithDefaults() *DeviceDeviceDeleteSuccessData {
	this := DeviceDeviceDeleteSuccessData{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *DeviceDeviceDeleteSuccessData) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceDeviceDeleteSuccessData) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *DeviceDeviceDeleteSuccessData) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *DeviceDeviceDeleteSuccessData) SetDeviceId(v string) {
	o.DeviceId = &v
}

func (o DeviceDeviceDeleteSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["device_id"] = o.DeviceId
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceDeviceDeleteSuccessData struct {
	value *DeviceDeviceDeleteSuccessData
	isSet bool
}

func (v NullableDeviceDeviceDeleteSuccessData) Get() *DeviceDeviceDeleteSuccessData {
	return v.value
}

func (v *NullableDeviceDeviceDeleteSuccessData) Set(val *DeviceDeviceDeleteSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceDeviceDeleteSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceDeviceDeleteSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceDeviceDeleteSuccessData(val *DeviceDeviceDeleteSuccessData) *NullableDeviceDeviceDeleteSuccessData {
	return &NullableDeviceDeviceDeleteSuccessData{value: val, isSet: true}
}

func (v NullableDeviceDeviceDeleteSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceDeviceDeleteSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


