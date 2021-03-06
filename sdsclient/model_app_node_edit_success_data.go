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

// AppNodeEditSuccessData struct for AppNodeEditSuccessData
type AppNodeEditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	NodeId *string `json:"node_id,omitempty"`
}

// NewAppNodeEditSuccessData instantiates a new AppNodeEditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppNodeEditSuccessData() *AppNodeEditSuccessData {
	this := AppNodeEditSuccessData{}
	return &this
}

// NewAppNodeEditSuccessDataWithDefaults instantiates a new AppNodeEditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppNodeEditSuccessDataWithDefaults() *AppNodeEditSuccessData {
	this := AppNodeEditSuccessData{}
	return &this
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *AppNodeEditSuccessData) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeEditSuccessData) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *AppNodeEditSuccessData) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *AppNodeEditSuccessData) SetNodeId(v string) {
	o.NodeId = &v
}

func (o AppNodeEditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeId != nil {
		toSerialize["node_id"] = o.NodeId
	}
	return json.Marshal(toSerialize)
}

type NullableAppNodeEditSuccessData struct {
	value *AppNodeEditSuccessData
	isSet bool
}

func (v NullableAppNodeEditSuccessData) Get() *AppNodeEditSuccessData {
	return v.value
}

func (v *NullableAppNodeEditSuccessData) Set(val *AppNodeEditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppNodeEditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppNodeEditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppNodeEditSuccessData(val *AppNodeEditSuccessData) *NullableAppNodeEditSuccessData {
	return &NullableAppNodeEditSuccessData{value: val, isSet: true}
}

func (v NullableAppNodeEditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppNodeEditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


