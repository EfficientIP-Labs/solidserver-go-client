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

// checks if the DataInnerDeviceLinkEditSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerDeviceLinkEditSuccess{}

// DataInnerDeviceLinkEditSuccess struct for DataInnerDeviceLinkEditSuccess
type DataInnerDeviceLinkEditSuccess struct {
	// The database identifier (ID) of the object you edited.
	LinkId *string `json:"link_id,omitempty"`
}

// NewDataInnerDeviceLinkEditSuccess instantiates a new DataInnerDeviceLinkEditSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerDeviceLinkEditSuccess() *DataInnerDeviceLinkEditSuccess {
	this := DataInnerDeviceLinkEditSuccess{}
	return &this
}

// NewDataInnerDeviceLinkEditSuccessWithDefaults instantiates a new DataInnerDeviceLinkEditSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerDeviceLinkEditSuccessWithDefaults() *DataInnerDeviceLinkEditSuccess {
	this := DataInnerDeviceLinkEditSuccess{}
	return &this
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *DataInnerDeviceLinkEditSuccess) GetLinkId() string {
	if o == nil || IsNil(o.LinkId) {
		var ret string
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDeviceLinkEditSuccess) GetLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *DataInnerDeviceLinkEditSuccess) HasLinkId() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given string and assigns it to the LinkId field.
func (o *DataInnerDeviceLinkEditSuccess) SetLinkId(v string) {
	o.LinkId = &v
}

func (o DataInnerDeviceLinkEditSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerDeviceLinkEditSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.LinkId) {
		toSerialize["link_id"] = o.LinkId
	}
	return toSerialize, nil
}

type NullableDataInnerDeviceLinkEditSuccess struct {
	value *DataInnerDeviceLinkEditSuccess
	isSet bool
}

func (v NullableDataInnerDeviceLinkEditSuccess) Get() *DataInnerDeviceLinkEditSuccess {
	return v.value
}

func (v *NullableDataInnerDeviceLinkEditSuccess) Set(val *DataInnerDeviceLinkEditSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerDeviceLinkEditSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerDeviceLinkEditSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerDeviceLinkEditSuccess(val *DataInnerDeviceLinkEditSuccess) *NullableDataInnerDeviceLinkEditSuccess {
	return &NullableDataInnerDeviceLinkEditSuccess{value: val, isSet: true}
}

func (v NullableDataInnerDeviceLinkEditSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerDeviceLinkEditSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}