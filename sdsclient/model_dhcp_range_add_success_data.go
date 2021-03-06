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

// DhcpRangeAddSuccessData struct for DhcpRangeAddSuccessData
type DhcpRangeAddSuccessData struct {
	// The database identifier (ID) of the object you added.
	RangeId *string `json:"range_id,omitempty"`
}

// NewDhcpRangeAddSuccessData instantiates a new DhcpRangeAddSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpRangeAddSuccessData() *DhcpRangeAddSuccessData {
	this := DhcpRangeAddSuccessData{}
	return &this
}

// NewDhcpRangeAddSuccessDataWithDefaults instantiates a new DhcpRangeAddSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpRangeAddSuccessDataWithDefaults() *DhcpRangeAddSuccessData {
	this := DhcpRangeAddSuccessData{}
	return &this
}

// GetRangeId returns the RangeId field value if set, zero value otherwise.
func (o *DhcpRangeAddSuccessData) GetRangeId() string {
	if o == nil || o.RangeId == nil {
		var ret string
		return ret
	}
	return *o.RangeId
}

// GetRangeIdOk returns a tuple with the RangeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpRangeAddSuccessData) GetRangeIdOk() (*string, bool) {
	if o == nil || o.RangeId == nil {
		return nil, false
	}
	return o.RangeId, true
}

// HasRangeId returns a boolean if a field has been set.
func (o *DhcpRangeAddSuccessData) HasRangeId() bool {
	if o != nil && o.RangeId != nil {
		return true
	}

	return false
}

// SetRangeId gets a reference to the given string and assigns it to the RangeId field.
func (o *DhcpRangeAddSuccessData) SetRangeId(v string) {
	o.RangeId = &v
}

func (o DhcpRangeAddSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RangeId != nil {
		toSerialize["range_id"] = o.RangeId
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpRangeAddSuccessData struct {
	value *DhcpRangeAddSuccessData
	isSet bool
}

func (v NullableDhcpRangeAddSuccessData) Get() *DhcpRangeAddSuccessData {
	return v.value
}

func (v *NullableDhcpRangeAddSuccessData) Set(val *DhcpRangeAddSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpRangeAddSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpRangeAddSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpRangeAddSuccessData(val *DhcpRangeAddSuccessData) *NullableDhcpRangeAddSuccessData {
	return &NullableDhcpRangeAddSuccessData{value: val, isSet: true}
}

func (v NullableDhcpRangeAddSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpRangeAddSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


