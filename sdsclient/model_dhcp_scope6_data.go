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

// checks if the DhcpScope6Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DhcpScope6Data{}

// DhcpScope6Data struct for DhcpScope6Data
type DhcpScope6Data struct {
	// state true/false indicate if action succeed
	Success *bool `json:"success,omitempty"`
	// List or notice/warning/error messages
	Messages []ApiMessageEntry         `json:"messages,omitempty"`
	Data     []DataInnerDhcpScope6Data `json:"data,omitempty"`
}

// NewDhcpScope6Data instantiates a new DhcpScope6Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpScope6Data() *DhcpScope6Data {
	this := DhcpScope6Data{}
	return &this
}

// NewDhcpScope6DataWithDefaults instantiates a new DhcpScope6Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpScope6DataWithDefaults() *DhcpScope6Data {
	this := DhcpScope6Data{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DhcpScope6Data) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6Data) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DhcpScope6Data) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DhcpScope6Data) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DhcpScope6Data) GetMessages() []ApiMessageEntry {
	if o == nil || IsNil(o.Messages) {
		var ret []ApiMessageEntry
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6Data) GetMessagesOk() ([]ApiMessageEntry, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DhcpScope6Data) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ApiMessageEntry and assigns it to the Messages field.
func (o *DhcpScope6Data) SetMessages(v []ApiMessageEntry) {
	o.Messages = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DhcpScope6Data) GetData() []DataInnerDhcpScope6Data {
	if o == nil || IsNil(o.Data) {
		var ret []DataInnerDhcpScope6Data
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6Data) GetDataOk() ([]DataInnerDhcpScope6Data, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DhcpScope6Data) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DataInnerDhcpScope6Data and assigns it to the Data field.
func (o *DhcpScope6Data) SetData(v []DataInnerDhcpScope6Data) {
	o.Data = v
}

func (o DhcpScope6Data) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DhcpScope6Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.Messages) {
		toSerialize["messages"] = o.Messages
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableDhcpScope6Data struct {
	value *DhcpScope6Data
	isSet bool
}

func (v NullableDhcpScope6Data) Get() *DhcpScope6Data {
	return v.value
}

func (v *NullableDhcpScope6Data) Set(val *DhcpScope6Data) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpScope6Data) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpScope6Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpScope6Data(val *DhcpScope6Data) *NullableDhcpScope6Data {
	return &NullableDhcpScope6Data{value: val, isSet: true}
}

func (v NullableDhcpScope6Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpScope6Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
