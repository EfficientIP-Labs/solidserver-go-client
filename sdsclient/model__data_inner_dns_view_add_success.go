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

// checks if the DataInnerDnsViewAddSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerDnsViewAddSuccess{}

// DataInnerDnsViewAddSuccess struct for DataInnerDnsViewAddSuccess
type DataInnerDnsViewAddSuccess struct {
	// The database identifier (ID) of the object you added.
	ViewId *string `json:"view_id,omitempty"`
}

// NewDataInnerDnsViewAddSuccess instantiates a new DataInnerDnsViewAddSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerDnsViewAddSuccess() *DataInnerDnsViewAddSuccess {
	this := DataInnerDnsViewAddSuccess{}
	return &this
}

// NewDataInnerDnsViewAddSuccessWithDefaults instantiates a new DataInnerDnsViewAddSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerDnsViewAddSuccessWithDefaults() *DataInnerDnsViewAddSuccess {
	this := DataInnerDnsViewAddSuccess{}
	return &this
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *DataInnerDnsViewAddSuccess) GetViewId() string {
	if o == nil || IsNil(o.ViewId) {
		var ret string
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDnsViewAddSuccess) GetViewIdOk() (*string, bool) {
	if o == nil || IsNil(o.ViewId) {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *DataInnerDnsViewAddSuccess) HasViewId() bool {
	if o != nil && !IsNil(o.ViewId) {
		return true
	}

	return false
}

// SetViewId gets a reference to the given string and assigns it to the ViewId field.
func (o *DataInnerDnsViewAddSuccess) SetViewId(v string) {
	o.ViewId = &v
}

func (o DataInnerDnsViewAddSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerDnsViewAddSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViewId) {
		toSerialize["view_id"] = o.ViewId
	}
	return toSerialize, nil
}

type NullableDataInnerDnsViewAddSuccess struct {
	value *DataInnerDnsViewAddSuccess
	isSet bool
}

func (v NullableDataInnerDnsViewAddSuccess) Get() *DataInnerDnsViewAddSuccess {
	return v.value
}

func (v *NullableDataInnerDnsViewAddSuccess) Set(val *DataInnerDnsViewAddSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerDnsViewAddSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerDnsViewAddSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerDnsViewAddSuccess(val *DataInnerDnsViewAddSuccess) *NullableDataInnerDnsViewAddSuccess {
	return &NullableDataInnerDnsViewAddSuccess{value: val, isSet: true}
}

func (v NullableDataInnerDnsViewAddSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerDnsViewAddSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}