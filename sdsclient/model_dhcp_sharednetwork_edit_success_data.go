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

// DhcpSharednetworkEditSuccessData struct for DhcpSharednetworkEditSuccessData
type DhcpSharednetworkEditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	SharednetworkId *string `json:"sharednetwork_id,omitempty"`
}

// NewDhcpSharednetworkEditSuccessData instantiates a new DhcpSharednetworkEditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSharednetworkEditSuccessData() *DhcpSharednetworkEditSuccessData {
	this := DhcpSharednetworkEditSuccessData{}
	return &this
}

// NewDhcpSharednetworkEditSuccessDataWithDefaults instantiates a new DhcpSharednetworkEditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSharednetworkEditSuccessDataWithDefaults() *DhcpSharednetworkEditSuccessData {
	this := DhcpSharednetworkEditSuccessData{}
	return &this
}

// GetSharednetworkId returns the SharednetworkId field value if set, zero value otherwise.
func (o *DhcpSharednetworkEditSuccessData) GetSharednetworkId() string {
	if o == nil || o.SharednetworkId == nil {
		var ret string
		return ret
	}
	return *o.SharednetworkId
}

// GetSharednetworkIdOk returns a tuple with the SharednetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkEditSuccessData) GetSharednetworkIdOk() (*string, bool) {
	if o == nil || o.SharednetworkId == nil {
		return nil, false
	}
	return o.SharednetworkId, true
}

// HasSharednetworkId returns a boolean if a field has been set.
func (o *DhcpSharednetworkEditSuccessData) HasSharednetworkId() bool {
	if o != nil && o.SharednetworkId != nil {
		return true
	}

	return false
}

// SetSharednetworkId gets a reference to the given string and assigns it to the SharednetworkId field.
func (o *DhcpSharednetworkEditSuccessData) SetSharednetworkId(v string) {
	o.SharednetworkId = &v
}

func (o DhcpSharednetworkEditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SharednetworkId != nil {
		toSerialize["sharednetwork_id"] = o.SharednetworkId
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpSharednetworkEditSuccessData struct {
	value *DhcpSharednetworkEditSuccessData
	isSet bool
}

func (v NullableDhcpSharednetworkEditSuccessData) Get() *DhcpSharednetworkEditSuccessData {
	return v.value
}

func (v *NullableDhcpSharednetworkEditSuccessData) Set(val *DhcpSharednetworkEditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSharednetworkEditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSharednetworkEditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSharednetworkEditSuccessData(val *DhcpSharednetworkEditSuccessData) *NullableDhcpSharednetworkEditSuccessData {
	return &NullableDhcpSharednetworkEditSuccessData{value: val, isSet: true}
}

func (v NullableDhcpSharednetworkEditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSharednetworkEditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


