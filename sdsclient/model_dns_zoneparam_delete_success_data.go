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

// DnsZoneparamDeleteSuccessData struct for DnsZoneparamDeleteSuccessData
type DnsZoneparamDeleteSuccessData struct {
	// The database identifier (ID) of the object you deleted.
	ZoneparamId *string `json:"zoneparam_id,omitempty"`
}

// NewDnsZoneparamDeleteSuccessData instantiates a new DnsZoneparamDeleteSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsZoneparamDeleteSuccessData() *DnsZoneparamDeleteSuccessData {
	this := DnsZoneparamDeleteSuccessData{}
	return &this
}

// NewDnsZoneparamDeleteSuccessDataWithDefaults instantiates a new DnsZoneparamDeleteSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsZoneparamDeleteSuccessDataWithDefaults() *DnsZoneparamDeleteSuccessData {
	this := DnsZoneparamDeleteSuccessData{}
	return &this
}

// GetZoneparamId returns the ZoneparamId field value if set, zero value otherwise.
func (o *DnsZoneparamDeleteSuccessData) GetZoneparamId() string {
	if o == nil || o.ZoneparamId == nil {
		var ret string
		return ret
	}
	return *o.ZoneparamId
}

// GetZoneparamIdOk returns a tuple with the ZoneparamId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsZoneparamDeleteSuccessData) GetZoneparamIdOk() (*string, bool) {
	if o == nil || o.ZoneparamId == nil {
		return nil, false
	}
	return o.ZoneparamId, true
}

// HasZoneparamId returns a boolean if a field has been set.
func (o *DnsZoneparamDeleteSuccessData) HasZoneparamId() bool {
	if o != nil && o.ZoneparamId != nil {
		return true
	}

	return false
}

// SetZoneparamId gets a reference to the given string and assigns it to the ZoneparamId field.
func (o *DnsZoneparamDeleteSuccessData) SetZoneparamId(v string) {
	o.ZoneparamId = &v
}

func (o DnsZoneparamDeleteSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ZoneparamId != nil {
		toSerialize["zoneparam_id"] = o.ZoneparamId
	}
	return json.Marshal(toSerialize)
}

type NullableDnsZoneparamDeleteSuccessData struct {
	value *DnsZoneparamDeleteSuccessData
	isSet bool
}

func (v NullableDnsZoneparamDeleteSuccessData) Get() *DnsZoneparamDeleteSuccessData {
	return v.value
}

func (v *NullableDnsZoneparamDeleteSuccessData) Set(val *DnsZoneparamDeleteSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsZoneparamDeleteSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsZoneparamDeleteSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsZoneparamDeleteSuccessData(val *DnsZoneparamDeleteSuccessData) *NullableDnsZoneparamDeleteSuccessData {
	return &NullableDnsZoneparamDeleteSuccessData{value: val, isSet: true}
}

func (v NullableDnsZoneparamDeleteSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsZoneparamDeleteSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


