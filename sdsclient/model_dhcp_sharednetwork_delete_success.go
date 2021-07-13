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

// DhcpSharednetworkDeleteSuccess struct for DhcpSharednetworkDeleteSuccess
type DhcpSharednetworkDeleteSuccess struct {
	// state true/false indicate if action succeed
	Success *bool `json:"success,omitempty"`
	// List or notice/warning/error messages
	Messages *[]ApiMessageEntry `json:"messages,omitempty"`
	Data *[]DhcpSharednetworkDeleteSuccessData `json:"data,omitempty"`
}

// NewDhcpSharednetworkDeleteSuccess instantiates a new DhcpSharednetworkDeleteSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSharednetworkDeleteSuccess() *DhcpSharednetworkDeleteSuccess {
	this := DhcpSharednetworkDeleteSuccess{}
	return &this
}

// NewDhcpSharednetworkDeleteSuccessWithDefaults instantiates a new DhcpSharednetworkDeleteSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSharednetworkDeleteSuccessWithDefaults() *DhcpSharednetworkDeleteSuccess {
	this := DhcpSharednetworkDeleteSuccess{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DhcpSharednetworkDeleteSuccess) GetSuccess() bool {
	if o == nil || o.Success == nil {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDeleteSuccess) GetSuccessOk() (*bool, bool) {
	if o == nil || o.Success == nil {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DhcpSharednetworkDeleteSuccess) HasSuccess() bool {
	if o != nil && o.Success != nil {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DhcpSharednetworkDeleteSuccess) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DhcpSharednetworkDeleteSuccess) GetMessages() []ApiMessageEntry {
	if o == nil || o.Messages == nil {
		var ret []ApiMessageEntry
		return ret
	}
	return *o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDeleteSuccess) GetMessagesOk() (*[]ApiMessageEntry, bool) {
	if o == nil || o.Messages == nil {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DhcpSharednetworkDeleteSuccess) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ApiMessageEntry and assigns it to the Messages field.
func (o *DhcpSharednetworkDeleteSuccess) SetMessages(v []ApiMessageEntry) {
	o.Messages = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DhcpSharednetworkDeleteSuccess) GetData() []DhcpSharednetworkDeleteSuccessData {
	if o == nil || o.Data == nil {
		var ret []DhcpSharednetworkDeleteSuccessData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDeleteSuccess) GetDataOk() (*[]DhcpSharednetworkDeleteSuccessData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DhcpSharednetworkDeleteSuccess) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given []DhcpSharednetworkDeleteSuccessData and assigns it to the Data field.
func (o *DhcpSharednetworkDeleteSuccess) SetData(v []DhcpSharednetworkDeleteSuccessData) {
	o.Data = &v
}

func (o DhcpSharednetworkDeleteSuccess) MarshalJSON() ([]byte, error) {
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

type NullableDhcpSharednetworkDeleteSuccess struct {
	value *DhcpSharednetworkDeleteSuccess
	isSet bool
}

func (v NullableDhcpSharednetworkDeleteSuccess) Get() *DhcpSharednetworkDeleteSuccess {
	return v.value
}

func (v *NullableDhcpSharednetworkDeleteSuccess) Set(val *DhcpSharednetworkDeleteSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSharednetworkDeleteSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSharednetworkDeleteSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSharednetworkDeleteSuccess(val *DhcpSharednetworkDeleteSuccess) *NullableDhcpSharednetworkDeleteSuccess {
	return &NullableDhcpSharednetworkDeleteSuccess{value: val, isSet: true}
}

func (v NullableDhcpSharednetworkDeleteSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSharednetworkDeleteSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


