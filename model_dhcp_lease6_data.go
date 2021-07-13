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

// DhcpLease6Data struct for DhcpLease6Data
type DhcpLease6Data struct {
	// state true/false indicate if action succeed
	Success *bool `json:"success,omitempty"`
	// List or notice/warning/error messages
	Messages *[]ApiMessageEntry `json:"messages,omitempty"`
	Data *[]DhcpLease6DataData `json:"data,omitempty"`
}

// NewDhcpLease6Data instantiates a new DhcpLease6Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpLease6Data() *DhcpLease6Data {
	this := DhcpLease6Data{}
	return &this
}

// NewDhcpLease6DataWithDefaults instantiates a new DhcpLease6Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpLease6DataWithDefaults() *DhcpLease6Data {
	this := DhcpLease6Data{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DhcpLease6Data) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpLease6Data) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DhcpLease6Data) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DhcpLease6Data) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DhcpLease6Data) GetMessages() []ApiMessageEntry {
	if o == nil || o.Messages == nil {
		var ret []ApiMessageEntry
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpLease6Data) GetMessagesOk() (*[]ApiMessageEntry, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DhcpLease6Data) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ApiMessageEntry and assigns it to the Messages field.
func (o *DhcpLease6Data) SetMessages(v []ApiMessageEntry) {
	o.Messages = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DhcpLease6Data) GetData() []DhcpLease6DataData {
	if o == nil || o.Data == nil {
		var ret []DhcpLease6DataData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpLease6Data) GetDataOk() (*[]DhcpLease6DataData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DhcpLease6Data) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []DhcpLease6DataData and assigns it to the Data field.
func (o *DhcpLease6Data) SetData(v []DhcpLease6DataData) {
	o.Data = &v
}

func (o DhcpLease6Data) MarshalJSON() ([]byte, error) {
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

type NullableDhcpLease6Data struct {
	value *DhcpLease6Data
	isSet bool
}

func (v NullableDhcpLease6Data) Get() *DhcpLease6Data {
	return v.value
}

func (v *NullableDhcpLease6Data) Set(val *DhcpLease6Data) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpLease6Data) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpLease6Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpLease6Data(val *DhcpLease6Data) *NullableDhcpLease6Data {
	return &NullableDhcpLease6Data{value: val, isSet: true}
}

func (v NullableDhcpLease6Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpLease6Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


