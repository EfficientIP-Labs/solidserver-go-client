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

// DhcpScopeEditSuccessData struct for DhcpScopeEditSuccessData
type DhcpScopeEditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	ScopeId *string `json:"scope_id,omitempty"`
}

// NewDhcpScopeEditSuccessData instantiates a new DhcpScopeEditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpScopeEditSuccessData() *DhcpScopeEditSuccessData {
	this := DhcpScopeEditSuccessData{}
	return &this
}

// NewDhcpScopeEditSuccessDataWithDefaults instantiates a new DhcpScopeEditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpScopeEditSuccessDataWithDefaults() *DhcpScopeEditSuccessData {
	this := DhcpScopeEditSuccessData{}
	return &this
}

// GetScopeId returns the ScopeId field value if set, zero value otherwise.
func (o *DhcpScopeEditSuccessData) GetScopeId() string {
	if o == nil || o.ScopeId == nil {
		var ret string
		return ret
	}
	return *o.ScopeId
}

// GetScopeIdOk returns a tuple with the ScopeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScopeEditSuccessData) GetScopeIdOk() (*string, bool) {
	if o == nil || o.ScopeId == nil {
		return nil, false
	}
	return o.ScopeId, true
}

// HasScopeId returns a boolean if a field has been set.
func (o *DhcpScopeEditSuccessData) HasScopeId() bool {
	if o != nil && o.ScopeId != nil {
		return true
	}

	return false
}

// SetScopeId gets a reference to the given string and assigns it to the ScopeId field.
func (o *DhcpScopeEditSuccessData) SetScopeId(v string) {
	o.ScopeId = &v
}

func (o DhcpScopeEditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ScopeId != nil {
		toSerialize["scope_id"] = o.ScopeId
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpScopeEditSuccessData struct {
	value *DhcpScopeEditSuccessData
	isSet bool
}

func (v NullableDhcpScopeEditSuccessData) Get() *DhcpScopeEditSuccessData {
	return v.value
}

func (v *NullableDhcpScopeEditSuccessData) Set(val *DhcpScopeEditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpScopeEditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpScopeEditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpScopeEditSuccessData(val *DhcpScopeEditSuccessData) *NullableDhcpScopeEditSuccessData {
	return &NullableDhcpScopeEditSuccessData{value: val, isSet: true}
}

func (v NullableDhcpScopeEditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpScopeEditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


