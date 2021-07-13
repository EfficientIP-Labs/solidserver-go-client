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

// DnsZoneEditSuccessData struct for DnsZoneEditSuccessData
type DnsZoneEditSuccessData struct {
	// The database identifier (ID) of the object you edited.
	ZoneId *string `json:"zone_id,omitempty"`
}

// NewDnsZoneEditSuccessData instantiates a new DnsZoneEditSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsZoneEditSuccessData() *DnsZoneEditSuccessData {
	this := DnsZoneEditSuccessData{}
	return &this
}

// NewDnsZoneEditSuccessDataWithDefaults instantiates a new DnsZoneEditSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsZoneEditSuccessDataWithDefaults() *DnsZoneEditSuccessData {
	this := DnsZoneEditSuccessData{}
	return &this
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *DnsZoneEditSuccessData) GetZoneId() string {
	if o == nil || o.ZoneId == nil {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsZoneEditSuccessData) GetZoneIdOk() (*string, bool) {
	if o == nil || o.ZoneId == nil {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *DnsZoneEditSuccessData) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *DnsZoneEditSuccessData) SetZoneId(v string) {
	o.ZoneId = &v
}

func (o DnsZoneEditSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneId != nil {
		toSerialize["zone_id"] = o.ZoneId
	}
	return json.Marshal(toSerialize)
}

type NullableDnsZoneEditSuccessData struct {
	value *DnsZoneEditSuccessData
	isSet bool
}

func (v NullableDnsZoneEditSuccessData) Get() *DnsZoneEditSuccessData {
	return v.value
}

func (v *NullableDnsZoneEditSuccessData) Set(val *DnsZoneEditSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsZoneEditSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsZoneEditSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsZoneEditSuccessData(val *DnsZoneEditSuccessData) *NullableDnsZoneEditSuccessData {
	return &NullableDnsZoneEditSuccessData{value: val, isSet: true}
}

func (v NullableDnsZoneEditSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsZoneEditSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


