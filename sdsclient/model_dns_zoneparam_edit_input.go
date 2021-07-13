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

// DnsZoneparamEditInput struct for DnsZoneparamEditInput
type DnsZoneparamEditInput struct {
	// The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice.
	ZoneId *int32 `json:"zone_id,omitempty"`
	// The name of the DNS option you want to add, edit or remove from the zone. You can only set one option at a time. <ul class=dashed ><li>                                                To add or edit an option: specify its name in the parameter <b>param_key</b>, as follows <b>param_key=&lt;option-name&gt;</b>, and then specify its value in the parameter <b>param_value</b>.<br/>                                            </li><li>                                                To remove an option, specify its name in the parameter <b>param_key</b> and leave the parameter <b>param_value</b> empty.<br/>                                            </li></ul>To set several options, specify as many parameters (<b>param_key</b> and <b>param_value</b>) as you need.
	ZoneparamKey *string `json:"zoneparam_key,omitempty"`
	// A way to determine is the DNS zone option is an array (<b>1</b>).
	ZoneparamIsArray *int32 `json:"zoneparam_is_array,omitempty"`
	// The value of the DNS option specified in the input <b>param_key</b>.<ul class=dashed ><li>                                                To add or edit an option value, specify its name in the parameter <b>param_key</b> and set its value as follows: <b>param_value=&lt;option-value&gt;</b> .<br/>                                            </li><li>                                                To remove an option value, specify its name in the parameter <b>param_key</b> and leave <b>param_value</b> empty: <b>param_value=</b> .<br/>                                            </li></ul>
	ZoneparamValue *string `json:"zoneparam_value,omitempty"`
}

// NewDnsZoneparamEditInput instantiates a new DnsZoneparamEditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsZoneparamEditInput() *DnsZoneparamEditInput {
	this := DnsZoneparamEditInput{}
	return &this
}

// NewDnsZoneparamEditInputWithDefaults instantiates a new DnsZoneparamEditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsZoneparamEditInputWithDefaults() *DnsZoneparamEditInput {
	this := DnsZoneparamEditInput{}
	return &this
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *DnsZoneparamEditInput) GetZoneId() int32 {
	if o == nil || o.ZoneId == nil {
		var ret int32
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsZoneparamEditInput) GetZoneIdOk() (*int32, bool) {
	if o == nil || o.ZoneId == nil {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *DnsZoneparamEditInput) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given int32 and assigns it to the ZoneId field.
func (o *DnsZoneparamEditInput) SetZoneId(v int32) {
	o.ZoneId = &v
}

// GetZoneparamKey returns the ZoneparamKey field value if set, zero value otherwise.
func (o *DnsZoneparamEditInput) GetZoneparamKey() string {
	if o == nil || o.ZoneparamKey == nil {
		var ret string
		return ret
	}
	return *o.ZoneparamKey
}

// GetZoneparamKeyOk returns a tuple with the ZoneparamKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsZoneparamEditInput) GetZoneparamKeyOk() (*string, bool) {
	if o == nil || o.ZoneparamKey == nil {
		return nil, false
	}
	return o.ZoneparamKey, true
}

// HasZoneparamKey returns a boolean if a field has been set.
func (o *DnsZoneparamEditInput) HasZoneparamKey() bool {
	if o != nil && o.ZoneparamKey != nil {
		return true
	}

	return false
}

// SetZoneparamKey gets a reference to the given string and assigns it to the ZoneparamKey field.
func (o *DnsZoneparamEditInput) SetZoneparamKey(v string) {
	o.ZoneparamKey = &v
}

// GetZoneparamIsArray returns the ZoneparamIsArray field value if set, zero value otherwise.
func (o *DnsZoneparamEditInput) GetZoneparamIsArray() int32 {
	if o == nil || o.ZoneparamIsArray == nil {
		var ret int32
		return ret
	}
	return *o.ZoneparamIsArray
}

// GetZoneparamIsArrayOk returns a tuple with the ZoneparamIsArray field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsZoneparamEditInput) GetZoneparamIsArrayOk() (*int32, bool) {
	if o == nil || o.ZoneparamIsArray == nil {
		return nil, false
	}
	return o.ZoneparamIsArray, true
}

// HasZoneparamIsArray returns a boolean if a field has been set.
func (o *DnsZoneparamEditInput) HasZoneparamIsArray() bool {
	if o != nil && o.ZoneparamIsArray != nil {
		return true
	}

	return false
}

// SetZoneparamIsArray gets a reference to the given int32 and assigns it to the ZoneparamIsArray field.
func (o *DnsZoneparamEditInput) SetZoneparamIsArray(v int32) {
	o.ZoneparamIsArray = &v
}

// GetZoneparamValue returns the ZoneparamValue field value if set, zero value otherwise.
func (o *DnsZoneparamEditInput) GetZoneparamValue() string {
	if o == nil || o.ZoneparamValue == nil {
		var ret string
		return ret
	}
	return *o.ZoneparamValue
}

// GetZoneparamValueOk returns a tuple with the ZoneparamValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsZoneparamEditInput) GetZoneparamValueOk() (*string, bool) {
	if o == nil || o.ZoneparamValue == nil {
		return nil, false
	}
	return o.ZoneparamValue, true
}

// HasZoneparamValue returns a boolean if a field has been set.
func (o *DnsZoneparamEditInput) HasZoneparamValue() bool {
	if o != nil && o.ZoneparamValue != nil {
		return true
	}

	return false
}

// SetZoneparamValue gets a reference to the given string and assigns it to the ZoneparamValue field.
func (o *DnsZoneparamEditInput) SetZoneparamValue(v string) {
	o.ZoneparamValue = &v
}

func (o DnsZoneparamEditInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneId != nil {
		toSerialize["zone_id"] = o.ZoneId
	}
	if o.ZoneparamKey != nil {
		toSerialize["zoneparam_key"] = o.ZoneparamKey
	}
	if o.ZoneparamIsArray != nil {
		toSerialize["zoneparam_is_array"] = o.ZoneparamIsArray
	}
	if o.ZoneparamValue != nil {
		toSerialize["zoneparam_value"] = o.ZoneparamValue
	}
	return json.Marshal(toSerialize)
}

type NullableDnsZoneparamEditInput struct {
	value *DnsZoneparamEditInput
	isSet bool
}

func (v NullableDnsZoneparamEditInput) Get() *DnsZoneparamEditInput {
	return v.value
}

func (v *NullableDnsZoneparamEditInput) Set(val *DnsZoneparamEditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsZoneparamEditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsZoneparamEditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsZoneparamEditInput(val *DnsZoneparamEditInput) *NullableDnsZoneparamEditInput {
	return &NullableDnsZoneparamEditInput{value: val, isSet: true}
}

func (v NullableDnsZoneparamEditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsZoneparamEditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


