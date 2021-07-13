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

// DhcpSharednetwork6EditSuccessData struct for DhcpSharednetwork6EditSuccessData
type DhcpSharednetwork6EditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	Sharednetwork6Id *string `json:"sharednetwork6_id,omitempty"`
}

// NewDhcpSharednetwork6EditSuccessData instantiates a new DhcpSharednetwork6EditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSharednetwork6EditSuccessData() *DhcpSharednetwork6EditSuccessData {
	this := DhcpSharednetwork6EditSuccessData{}
	return &this
}

// NewDhcpSharednetwork6EditSuccessDataWithDefaults instantiates a new DhcpSharednetwork6EditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSharednetwork6EditSuccessDataWithDefaults() *DhcpSharednetwork6EditSuccessData {
	this := DhcpSharednetwork6EditSuccessData{}
	return &this
}

// GetSharednetwork6Id returns the Sharednetwork6Id field value if set, zero value otherwise.
func (o *DhcpSharednetwork6EditSuccessData) GetSharednetwork6Id() string {
	if o == nil || o.Sharednetwork6Id == nil {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Id
}

// GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6EditSuccessData) GetSharednetwork6IdOk() (*string, bool) {
	if o == nil || o.Sharednetwork6Id == nil {
		return nil, false
	}
	return o.Sharednetwork6Id, true
}

// HasSharednetwork6Id returns a boolean if a field has been set.
func (o *DhcpSharednetwork6EditSuccessData) HasSharednetwork6Id() bool {
	if o != nil && o.Sharednetwork6Id != nil {
		return true
	}

	return false
}

// SetSharednetwork6Id gets a reference to the given string and assigns it to the Sharednetwork6Id field.
func (o *DhcpSharednetwork6EditSuccessData) SetSharednetwork6Id(v string) {
	o.Sharednetwork6Id = &v
}

func (o DhcpSharednetwork6EditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sharednetwork6Id != nil {
		toSerialize["sharednetwork6_id"] = o.Sharednetwork6Id
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpSharednetwork6EditSuccessData struct {
	value *DhcpSharednetwork6EditSuccessData
	isSet bool
}

func (v NullableDhcpSharednetwork6EditSuccessData) Get() *DhcpSharednetwork6EditSuccessData {
	return v.value
}

func (v *NullableDhcpSharednetwork6EditSuccessData) Set(val *DhcpSharednetwork6EditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSharednetwork6EditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSharednetwork6EditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSharednetwork6EditSuccessData(val *DhcpSharednetwork6EditSuccessData) *NullableDhcpSharednetwork6EditSuccessData {
	return &NullableDhcpSharednetwork6EditSuccessData{value: val, isSet: true}
}

func (v NullableDhcpSharednetwork6EditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSharednetwork6EditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

