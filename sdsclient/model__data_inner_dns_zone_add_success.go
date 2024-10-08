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

// checks if the DataInnerDnsZoneAddSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerDnsZoneAddSuccess{}

// DataInnerDnsZoneAddSuccess struct for DataInnerDnsZoneAddSuccess
type DataInnerDnsZoneAddSuccess struct {
	// The database identifier (ID) of the object you added.
	ZoneId *string `json:"zone_id,omitempty"`
}

// NewDataInnerDnsZoneAddSuccess instantiates a new DataInnerDnsZoneAddSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerDnsZoneAddSuccess() *DataInnerDnsZoneAddSuccess {
	this := DataInnerDnsZoneAddSuccess{}
	return &this
}

// NewDataInnerDnsZoneAddSuccessWithDefaults instantiates a new DataInnerDnsZoneAddSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerDnsZoneAddSuccessWithDefaults() *DataInnerDnsZoneAddSuccess {
	this := DataInnerDnsZoneAddSuccess{}
	return &this
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *DataInnerDnsZoneAddSuccess) GetZoneId() string {
	if o == nil || IsNil(o.ZoneId) {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDnsZoneAddSuccess) GetZoneIdOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *DataInnerDnsZoneAddSuccess) HasZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *DataInnerDnsZoneAddSuccess) SetZoneId(v string) {
	o.ZoneId = &v
}

func (o DataInnerDnsZoneAddSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerDnsZoneAddSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ZoneId) {
		toSerialize["zone_id"] = o.ZoneId
	}
	return toSerialize, nil
}

type NullableDataInnerDnsZoneAddSuccess struct {
	value *DataInnerDnsZoneAddSuccess
	isSet bool
}

func (v NullableDataInnerDnsZoneAddSuccess) Get() *DataInnerDnsZoneAddSuccess {
	return v.value
}

func (v *NullableDataInnerDnsZoneAddSuccess) Set(val *DataInnerDnsZoneAddSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerDnsZoneAddSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerDnsZoneAddSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerDnsZoneAddSuccess(val *DataInnerDnsZoneAddSuccess) *NullableDataInnerDnsZoneAddSuccess {
	return &NullableDataInnerDnsZoneAddSuccess{value: val, isSet: true}
}

func (v NullableDataInnerDnsZoneAddSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerDnsZoneAddSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}