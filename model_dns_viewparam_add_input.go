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

// DnsViewparamAddInput struct for DnsViewparamAddInput
type DnsViewparamAddInput struct {
	// The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice.
	ViewId *int32 `json:"view_id,omitempty"`
	// The name of the DNS option you want to add, edit or remove from the view. You can only set one option at a time. <ul class=dashed ><li>                                                To add or edit an option: specify its name in the parameter <b>param_key</b>, as follows <b>param_key=&lt;option-name&gt;</b>, and then specify its value in the parameter <b>param_value</b>.<br/>                                            </li><li>                                                To remove an option, specify its name in the parameter <b>param_key</b> and leave the parameter <b>param_value</b> empty.<br/>                                            </li></ul>To set several options, specify as many parameters (<b>param_key</b> and <b>param_value</b>) as you need.
	ViewparamKey *string `json:"viewparam_key,omitempty"`
	// A way to determine is the DNS view option is an array (<b>1</b>).
	ViewparamIsArray *int32 `json:"viewparam_is_array,omitempty"`
	// The value of the DNS option specified in the input <b>param_key</b>.<ul class=dashed ><li>                                                To add or edit an option value, specify its name in the parameter <b>param_key</b> and set its value as follows: <b>param_value=&lt;option-value&gt;</b> .<br/>                                            </li><li>                                                To remove an option value, specify its name in the parameter <b>param_key</b> and leave <b>param_value</b> empty: <b>param_value=</b> .<br/>                                            </li></ul>
	ViewparamValue *string `json:"viewparam_value,omitempty"`
}

// NewDnsViewparamAddInput instantiates a new DnsViewparamAddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsViewparamAddInput() *DnsViewparamAddInput {
	this := DnsViewparamAddInput{}
	return &this
}

// NewDnsViewparamAddInputWithDefaults instantiates a new DnsViewparamAddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsViewparamAddInputWithDefaults() *DnsViewparamAddInput {
	this := DnsViewparamAddInput{}
	return &this
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *DnsViewparamAddInput) GetViewId() int32 {
	if o == nil || o.ViewId == nil {
		var ret int32
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsViewparamAddInput) GetViewIdOk() (*int32, bool) {
	if o == nil || o.ViewId == nil {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *DnsViewparamAddInput) HasViewId() bool {
	if o != nil && o.ViewId != nil {
		return true
	}

	return false
}

// SetViewId gets a reference to the given int32 and assigns it to the ViewId field.
func (o *DnsViewparamAddInput) SetViewId(v int32) {
	o.ViewId = &v
}

// GetViewparamKey returns the ViewparamKey field value if set, zero value otherwise.
func (o *DnsViewparamAddInput) GetViewparamKey() string {
	if o == nil || o.ViewparamKey == nil {
		var ret string
		return ret
	}
	return *o.ViewparamKey
}

// GetViewparamKeyOk returns a tuple with the ViewparamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsViewparamAddInput) GetViewparamKeyOk() (*string, bool) {
	if o == nil || o.ViewparamKey == nil {
		return nil, false
	}
	return o.ViewparamKey, true
}

// HasViewparamKey returns a boolean if a field has been set.
func (o *DnsViewparamAddInput) HasViewparamKey() bool {
	if o != nil && o.ViewparamKey != nil {
		return true
	}

	return false
}

// SetViewparamKey gets a reference to the given string and assigns it to the ViewparamKey field.
func (o *DnsViewparamAddInput) SetViewparamKey(v string) {
	o.ViewparamKey = &v
}

// GetViewparamIsArray returns the ViewparamIsArray field value if set, zero value otherwise.
func (o *DnsViewparamAddInput) GetViewparamIsArray() int32 {
	if o == nil || o.ViewparamIsArray == nil {
		var ret int32
		return ret
	}
	return *o.ViewparamIsArray
}

// GetViewparamIsArrayOk returns a tuple with the ViewparamIsArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsViewparamAddInput) GetViewparamIsArrayOk() (*int32, bool) {
	if o == nil || o.ViewparamIsArray == nil {
		return nil, false
	}
	return o.ViewparamIsArray, true
}

// HasViewparamIsArray returns a boolean if a field has been set.
func (o *DnsViewparamAddInput) HasViewparamIsArray() bool {
	if o != nil && o.ViewparamIsArray != nil {
		return true
	}

	return false
}

// SetViewparamIsArray gets a reference to the given int32 and assigns it to the ViewparamIsArray field.
func (o *DnsViewparamAddInput) SetViewparamIsArray(v int32) {
	o.ViewparamIsArray = &v
}

// GetViewparamValue returns the ViewparamValue field value if set, zero value otherwise.
func (o *DnsViewparamAddInput) GetViewparamValue() string {
	if o == nil || o.ViewparamValue == nil {
		var ret string
		return ret
	}
	return *o.ViewparamValue
}

// GetViewparamValueOk returns a tuple with the ViewparamValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsViewparamAddInput) GetViewparamValueOk() (*string, bool) {
	if o == nil || o.ViewparamValue == nil {
		return nil, false
	}
	return o.ViewparamValue, true
}

// HasViewparamValue returns a boolean if a field has been set.
func (o *DnsViewparamAddInput) HasViewparamValue() bool {
	if o != nil && o.ViewparamValue != nil {
		return true
	}

	return false
}

// SetViewparamValue gets a reference to the given string and assigns it to the ViewparamValue field.
func (o *DnsViewparamAddInput) SetViewparamValue(v string) {
	o.ViewparamValue = &v
}

func (o DnsViewparamAddInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ViewId != nil {
		toSerialize["view_id"] = o.ViewId
	}
	if o.ViewparamKey != nil {
		toSerialize["viewparam_key"] = o.ViewparamKey
	}
	if o.ViewparamIsArray != nil {
		toSerialize["viewparam_is_array"] = o.ViewparamIsArray
	}
	if o.ViewparamValue != nil {
		toSerialize["viewparam_value"] = o.ViewparamValue
	}
	return json.Marshal(toSerialize)
}

type NullableDnsViewparamAddInput struct {
	value *DnsViewparamAddInput
	isSet bool
}

func (v NullableDnsViewparamAddInput) Get() *DnsViewparamAddInput {
	return v.value
}

func (v *NullableDnsViewparamAddInput) Set(val *DnsViewparamAddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsViewparamAddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsViewparamAddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsViewparamAddInput(val *DnsViewparamAddInput) *NullableDnsViewparamAddInput {
	return &NullableDnsViewparamAddInput{value: val, isSet: true}
}

func (v NullableDnsViewparamAddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsViewparamAddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


