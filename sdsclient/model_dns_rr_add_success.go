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

// checks if the DnsRrAddSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRrAddSuccess{}

// DnsRrAddSuccess struct for DnsRrAddSuccess
type DnsRrAddSuccess struct {
	// state true/false indicate if action succeed
	Success *bool `json:"success,omitempty"`
	// List or notice/warning/error messages
	Messages []ApiMessageEntry          `json:"messages,omitempty"`
	Data     []DataInnerDnsRrAddSuccess `json:"data,omitempty"`
}

// NewDnsRrAddSuccess instantiates a new DnsRrAddSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRrAddSuccess() *DnsRrAddSuccess {
	this := DnsRrAddSuccess{}
	return &this
}

// NewDnsRrAddSuccessWithDefaults instantiates a new DnsRrAddSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRrAddSuccessWithDefaults() *DnsRrAddSuccess {
	this := DnsRrAddSuccess{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *DnsRrAddSuccess) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddSuccess) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// HasSuccess returns a boolean if a field has been set.
func (o *DnsRrAddSuccess) HasSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *DnsRrAddSuccess) SetSuccess(v bool) {
	o.Success = &v
}

// GetMessages returns the Messages field value if set, zero value otherwise.
func (o *DnsRrAddSuccess) GetMessages() []ApiMessageEntry {
	if o == nil || IsNil(o.Messages) {
		var ret []ApiMessageEntry
		return ret
	}
	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddSuccess) GetMessagesOk() ([]ApiMessageEntry, bool) {
	if o == nil || IsNil(o.Messages) {
		return nil, false
	}
	return o.Messages, true
}

// HasMessages returns a boolean if a field has been set.
func (o *DnsRrAddSuccess) HasMessages() bool {
	if o != nil && !IsNil(o.Messages) {
		return true
	}

	return false
}

// SetMessages gets a reference to the given []ApiMessageEntry and assigns it to the Messages field.
func (o *DnsRrAddSuccess) SetMessages(v []ApiMessageEntry) {
	o.Messages = v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *DnsRrAddSuccess) GetData() []DataInnerDnsRrAddSuccess {
	if o == nil || IsNil(o.Data) {
		var ret []DataInnerDnsRrAddSuccess
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddSuccess) GetDataOk() ([]DataInnerDnsRrAddSuccess, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *DnsRrAddSuccess) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []DataInnerDnsRrAddSuccess and assigns it to the Data field.
func (o *DnsRrAddSuccess) SetData(v []DataInnerDnsRrAddSuccess) {
	o.Data = v
}

func (o DnsRrAddSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRrAddSuccess) ToMap() (map[string]interface{}, error) {
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

type NullableDnsRrAddSuccess struct {
	value *DnsRrAddSuccess
	isSet bool
}

func (v NullableDnsRrAddSuccess) Get() *DnsRrAddSuccess {
	return v.value
}

func (v *NullableDnsRrAddSuccess) Set(val *DnsRrAddSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRrAddSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRrAddSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRrAddSuccess(val *DnsRrAddSuccess) *NullableDnsRrAddSuccess {
	return &NullableDnsRrAddSuccess{value: val, isSet: true}
}

func (v NullableDnsRrAddSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRrAddSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
