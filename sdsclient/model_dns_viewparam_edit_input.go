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

// checks if the DnsViewparamEditInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsViewparamEditInput{}

// DnsViewparamEditInput struct for DnsViewparamEditInput
type DnsViewparamEditInput struct {
	// The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice.
	ViewId *int32 `json:"view_id,omitempty"`
	// The name of the DNS option you want to add, edit or remove from the view. You can only set one option at a time.<ul><li> To add or edit an option: specify its name in the parameter <b>viewparam_key</b>, as follows <b>viewparam_key=&lt;option-name&gt;</b>, and then specify its value in the parameter <b>viewparam_value</b>.<br/></li><li> To remove an option, specify its name in the parameter <b>viewparam_key</b> and leave the parameter <b>viewparam_value</b> empty.<br/></li></ul>To set several options, specify as many parameters (<b>viewparam_key</b> and <b>viewparam_value</b>) as you need.
	ViewparamKey *string `json:"viewparam_key,omitempty"`
	// A way to determine is the DNS view option is an array (<b>1</b>).
	ViewparamIsArray *int32 `json:"viewparam_is_array,omitempty"`
	// The value of the DNS option specified in the input <b>viewparam_key</b>.<ul><li> To add or edit an option value, specify its name in the parameter <b>viewparam_key</b> and set its value as follows: <b>viewparam_value=&lt;option-value&gt;</b> .<br/></li><li> To remove an option value, specify its name in the parameter <b>viewparam_key</b> and leave <b>viewparam_value</b> empty: <b>viewparam_value=</b> .<br/></li></ul>
	ViewparamValue *string `json:"viewparam_value,omitempty"`
}

// NewDnsViewparamEditInput instantiates a new DnsViewparamEditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsViewparamEditInput() *DnsViewparamEditInput {
	this := DnsViewparamEditInput{}
	return &this
}

// NewDnsViewparamEditInputWithDefaults instantiates a new DnsViewparamEditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsViewparamEditInputWithDefaults() *DnsViewparamEditInput {
	this := DnsViewparamEditInput{}
	return &this
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *DnsViewparamEditInput) GetViewId() int32 {
	if o == nil || IsNil(o.ViewId) {
		var ret int32
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsViewparamEditInput) GetViewIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ViewId) {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *DnsViewparamEditInput) HasViewId() bool {
	if o != nil && !IsNil(o.ViewId) {
		return true
	}

	return false
}

// SetViewId gets a reference to the given int32 and assigns it to the ViewId field.
func (o *DnsViewparamEditInput) SetViewId(v int32) {
	o.ViewId = &v
}

// GetViewparamKey returns the ViewparamKey field value if set, zero value otherwise.
func (o *DnsViewparamEditInput) GetViewparamKey() string {
	if o == nil || IsNil(o.ViewparamKey) {
		var ret string
		return ret
	}
	return *o.ViewparamKey
}

// GetViewparamKeyOk returns a tuple with the ViewparamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsViewparamEditInput) GetViewparamKeyOk() (*string, bool) {
	if o == nil || IsNil(o.ViewparamKey) {
		return nil, false
	}
	return o.ViewparamKey, true
}

// HasViewparamKey returns a boolean if a field has been set.
func (o *DnsViewparamEditInput) HasViewparamKey() bool {
	if o != nil && !IsNil(o.ViewparamKey) {
		return true
	}

	return false
}

// SetViewparamKey gets a reference to the given string and assigns it to the ViewparamKey field.
func (o *DnsViewparamEditInput) SetViewparamKey(v string) {
	o.ViewparamKey = &v
}

// GetViewparamIsArray returns the ViewparamIsArray field value if set, zero value otherwise.
func (o *DnsViewparamEditInput) GetViewparamIsArray() int32 {
	if o == nil || IsNil(o.ViewparamIsArray) {
		var ret int32
		return ret
	}
	return *o.ViewparamIsArray
}

// GetViewparamIsArrayOk returns a tuple with the ViewparamIsArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsViewparamEditInput) GetViewparamIsArrayOk() (*int32, bool) {
	if o == nil || IsNil(o.ViewparamIsArray) {
		return nil, false
	}
	return o.ViewparamIsArray, true
}

// HasViewparamIsArray returns a boolean if a field has been set.
func (o *DnsViewparamEditInput) HasViewparamIsArray() bool {
	if o != nil && !IsNil(o.ViewparamIsArray) {
		return true
	}

	return false
}

// SetViewparamIsArray gets a reference to the given int32 and assigns it to the ViewparamIsArray field.
func (o *DnsViewparamEditInput) SetViewparamIsArray(v int32) {
	o.ViewparamIsArray = &v
}

// GetViewparamValue returns the ViewparamValue field value if set, zero value otherwise.
func (o *DnsViewparamEditInput) GetViewparamValue() string {
	if o == nil || IsNil(o.ViewparamValue) {
		var ret string
		return ret
	}
	return *o.ViewparamValue
}

// GetViewparamValueOk returns a tuple with the ViewparamValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsViewparamEditInput) GetViewparamValueOk() (*string, bool) {
	if o == nil || IsNil(o.ViewparamValue) {
		return nil, false
	}
	return o.ViewparamValue, true
}

// HasViewparamValue returns a boolean if a field has been set.
func (o *DnsViewparamEditInput) HasViewparamValue() bool {
	if o != nil && !IsNil(o.ViewparamValue) {
		return true
	}

	return false
}

// SetViewparamValue gets a reference to the given string and assigns it to the ViewparamValue field.
func (o *DnsViewparamEditInput) SetViewparamValue(v string) {
	o.ViewparamValue = &v
}

func (o DnsViewparamEditInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsViewparamEditInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ViewId) {
		toSerialize["view_id"] = o.ViewId
	}
	if !IsNil(o.ViewparamKey) {
		toSerialize["viewparam_key"] = o.ViewparamKey
	}
	if !IsNil(o.ViewparamIsArray) {
		toSerialize["viewparam_is_array"] = o.ViewparamIsArray
	}
	if !IsNil(o.ViewparamValue) {
		toSerialize["viewparam_value"] = o.ViewparamValue
	}
	return toSerialize, nil
}

type NullableDnsViewparamEditInput struct {
	value *DnsViewparamEditInput
	isSet bool
}

func (v NullableDnsViewparamEditInput) Get() *DnsViewparamEditInput {
	return v.value
}

func (v *NullableDnsViewparamEditInput) Set(val *DnsViewparamEditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsViewparamEditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsViewparamEditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsViewparamEditInput(val *DnsViewparamEditInput) *NullableDnsViewparamEditInput {
	return &NullableDnsViewparamEditInput{value: val, isSet: true}
}

func (v NullableDnsViewparamEditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsViewparamEditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
