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

// DhcpAcl6Data struct for DhcpAcl6Data
type DhcpAcl6Data struct {
	// state true/false indicate if action succeed
	Success *bool `json:"success,omitempty"`
	// List or notice/warning/error messages
	Messages *[]ApiMessageEntry `json:"messages,omitempty"`
	Data *[]DhcpAcl6DataData `json:"data,omitempty"`
}

// NewDhcpAcl6Data instantiates a new DhcpAcl6Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpAcl6Data() *DhcpAcl6Data {
	this := DhcpAcl6Data{}
	return &this
}

// NewDhcpAcl6DataWithDefaults instantiates a new DhcpAcl6Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpAcl6DataWithDefaults() *DhcpAcl6Data {
	this := DhcpAcl6Data{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DhcpAcl6Data) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpAcl6Data) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DhcpAcl6Data) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DhcpAcl6Data) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DhcpAcl6Data) GetMessages() []ApiMessageEntry {
	if o == nil || o.Messages == nil {
		var ret []ApiMessageEntry
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpAcl6Data) GetMessagesOk() (*[]ApiMessageEntry, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DhcpAcl6Data) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ApiMessageEntry and assigns it to the Messages field.
func (o *DhcpAcl6Data) SetMessages(v []ApiMessageEntry) {
	o.Messages = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DhcpAcl6Data) GetData() []DhcpAcl6DataData {
	if o == nil || o.Data == nil {
		var ret []DhcpAcl6DataData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpAcl6Data) GetDataOk() (*[]DhcpAcl6DataData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DhcpAcl6Data) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []DhcpAcl6DataData and assigns it to the Data field.
func (o *DhcpAcl6Data) SetData(v []DhcpAcl6DataData) {
	o.Data = &v
}

func (o DhcpAcl6Data) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Success != nil {
		toSerialize["success"] = o.Success
	}
	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpAcl6Data struct {
	value *DhcpAcl6Data
	isSet bool
}

func (v NullableDhcpAcl6Data) Get() *DhcpAcl6Data {
	return v.value
}

func (v *NullableDhcpAcl6Data) Set(val *DhcpAcl6Data) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpAcl6Data) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpAcl6Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpAcl6Data(val *DhcpAcl6Data) *NullableDhcpAcl6Data {
	return &NullableDhcpAcl6Data{value: val, isSet: true}
}

func (v NullableDhcpAcl6Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpAcl6Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

