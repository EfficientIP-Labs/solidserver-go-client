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

// DhcpStatic6AddFailed struct for DhcpStatic6AddFailed
type DhcpStatic6AddFailed struct {
	// state true/false indicate if action succeed
	Success *bool `json:"success,omitempty"`
	// List or notice/warning/error messages
	Messages *[]ApiMessageEntry `json:"messages,omitempty"`
	Data *[]map[string]interface{} `json:"data,omitempty"`
}

// NewDhcpStatic6AddFailed instantiates a new DhcpStatic6AddFailed object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpStatic6AddFailed() *DhcpStatic6AddFailed {
	this := DhcpStatic6AddFailed{}
	return &this
}

// NewDhcpStatic6AddFailedWithDefaults instantiates a new DhcpStatic6AddFailed object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpStatic6AddFailedWithDefaults() *DhcpStatic6AddFailed {
	this := DhcpStatic6AddFailed{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DhcpStatic6AddFailed) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6AddFailed) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DhcpStatic6AddFailed) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DhcpStatic6AddFailed) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DhcpStatic6AddFailed) GetMessages() []ApiMessageEntry {
	if o == nil || o.Messages == nil {
		var ret []ApiMessageEntry
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6AddFailed) GetMessagesOk() (*[]ApiMessageEntry, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DhcpStatic6AddFailed) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ApiMessageEntry and assigns it to the Messages field.
func (o *DhcpStatic6AddFailed) SetMessages(v []ApiMessageEntry) {
	o.Messages = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DhcpStatic6AddFailed) GetData() []map[string]interface{} {
	if o == nil || o.Data == nil {
		var ret []map[string]interface{}
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpStatic6AddFailed) GetDataOk() (*[]map[string]interface{}, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DhcpStatic6AddFailed) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []map[string]interface{} and assigns it to the Data field.
func (o *DhcpStatic6AddFailed) SetData(v []map[string]interface{}) {
	o.Data = &v
}

func (o DhcpStatic6AddFailed) MarshalJSON() ([]byte, error) {
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

type NullableDhcpStatic6AddFailed struct {
	value *DhcpStatic6AddFailed
	isSet bool
}

func (v NullableDhcpStatic6AddFailed) Get() *DhcpStatic6AddFailed {
	return v.value
}

func (v *NullableDhcpStatic6AddFailed) Set(val *DhcpStatic6AddFailed) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpStatic6AddFailed) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpStatic6AddFailed) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpStatic6AddFailed(val *DhcpStatic6AddFailed) *NullableDhcpStatic6AddFailed {
	return &NullableDhcpStatic6AddFailed{value: val, isSet: true}
}

func (v NullableDhcpStatic6AddFailed) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpStatic6AddFailed) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

