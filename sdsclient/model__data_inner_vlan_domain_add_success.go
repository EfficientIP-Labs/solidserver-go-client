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

// checks if the DataInnerVlanDomainAddSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerVlanDomainAddSuccess{}

// DataInnerVlanDomainAddSuccess struct for DataInnerVlanDomainAddSuccess
type DataInnerVlanDomainAddSuccess struct {
	// The database identifier (ID) of the object you added.
	DomainId *string `json:"domain_id,omitempty"`
}

// NewDataInnerVlanDomainAddSuccess instantiates a new DataInnerVlanDomainAddSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerVlanDomainAddSuccess() *DataInnerVlanDomainAddSuccess {
	this := DataInnerVlanDomainAddSuccess{}
	return &this
}

// NewDataInnerVlanDomainAddSuccessWithDefaults instantiates a new DataInnerVlanDomainAddSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerVlanDomainAddSuccessWithDefaults() *DataInnerVlanDomainAddSuccess {
	this := DataInnerVlanDomainAddSuccess{}
	return &this
}

// GetDomainId returns the DomainId field value if set, zero value otherwise.
func (o *DataInnerVlanDomainAddSuccess) GetDomainId() string {
	if o == nil || IsNil(o.DomainId) {
		var ret string
		return ret
	}
	return *o.DomainId
}

// GetDomainIdOk returns a tuple with the DomainId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerVlanDomainAddSuccess) GetDomainIdOk() (*string, bool) {
	if o == nil || IsNil(o.DomainId) {
		return nil, false
	}
	return o.DomainId, true
}

// HasDomainId returns a boolean if a field has been set.
func (o *DataInnerVlanDomainAddSuccess) HasDomainId() bool {
	if o != nil && !IsNil(o.DomainId) {
		return true
	}

	return false
}

// SetDomainId gets a reference to the given string and assigns it to the DomainId field.
func (o *DataInnerVlanDomainAddSuccess) SetDomainId(v string) {
	o.DomainId = &v
}

func (o DataInnerVlanDomainAddSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerVlanDomainAddSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DomainId) {
		toSerialize["domain_id"] = o.DomainId
	}
	return toSerialize, nil
}

type NullableDataInnerVlanDomainAddSuccess struct {
	value *DataInnerVlanDomainAddSuccess
	isSet bool
}

func (v NullableDataInnerVlanDomainAddSuccess) Get() *DataInnerVlanDomainAddSuccess {
	return v.value
}

func (v *NullableDataInnerVlanDomainAddSuccess) Set(val *DataInnerVlanDomainAddSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerVlanDomainAddSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerVlanDomainAddSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerVlanDomainAddSuccess(val *DataInnerVlanDomainAddSuccess) *NullableDataInnerVlanDomainAddSuccess {
	return &NullableDataInnerVlanDomainAddSuccess{value: val, isSet: true}
}

func (v NullableDataInnerVlanDomainAddSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerVlanDomainAddSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}