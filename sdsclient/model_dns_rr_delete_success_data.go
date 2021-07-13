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

// DnsRrDeleteSuccessData struct for DnsRrDeleteSuccessData
type DnsRrDeleteSuccessData struct {
	// The database identifier (ID) of the object you deleted.
	RrId *string `json:"rr_id,omitempty"`
}

// NewDnsRrDeleteSuccessData instantiates a new DnsRrDeleteSuccessData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRrDeleteSuccessData() *DnsRrDeleteSuccessData {
	this := DnsRrDeleteSuccessData{}
	return &this
}

// NewDnsRrDeleteSuccessDataWithDefaults instantiates a new DnsRrDeleteSuccessData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRrDeleteSuccessDataWithDefaults() *DnsRrDeleteSuccessData {
	this := DnsRrDeleteSuccessData{}
	return &this
}

// GetRrId returns the RrId field value if set, zero value otherwise.
func (o *DnsRrDeleteSuccessData) GetRrId() string {
	if o == nil || o.RrId == nil {
		var ret string
		return ret
	}
	return *o.RrId
}

// GetRrIdOk returns a tuple with the RrId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrDeleteSuccessData) GetRrIdOk() (*string, bool) {
	if o == nil || o.RrId == nil {
		return nil, false
	}
	return o.RrId, true
}

// HasRrId returns a boolean if a field has been set.
func (o *DnsRrDeleteSuccessData) HasRrId() bool {
	if o != nil && o.RrId != nil {
		return true
	}

	return false
}

// SetRrId gets a reference to the given string and assigns it to the RrId field.
func (o *DnsRrDeleteSuccessData) SetRrId(v string) {
	o.RrId = &v
}

func (o DnsRrDeleteSuccessData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.RrId != nil {
		toSerialize["rr_id"] = o.RrId
	}
	return json.Marshal(toSerialize)
}

type NullableDnsRrDeleteSuccessData struct {
	value *DnsRrDeleteSuccessData
	isSet bool
}

func (v NullableDnsRrDeleteSuccessData) Get() *DnsRrDeleteSuccessData {
	return v.value
}

func (v *NullableDnsRrDeleteSuccessData) Set(val *DnsRrDeleteSuccessData) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRrDeleteSuccessData) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRrDeleteSuccessData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRrDeleteSuccessData(val *DnsRrDeleteSuccessData) *NullableDnsRrDeleteSuccessData {
	return &NullableDnsRrDeleteSuccessData{value: val, isSet: true}
}

func (v NullableDnsRrDeleteSuccessData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRrDeleteSuccessData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


