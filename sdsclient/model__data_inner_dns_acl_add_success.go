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

// checks if the DataInnerDnsAclAddSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerDnsAclAddSuccess{}

// DataInnerDnsAclAddSuccess struct for DataInnerDnsAclAddSuccess
type DataInnerDnsAclAddSuccess struct {
	// The database identifier (ID) of the object you added.
	AclId *string `json:"acl_id,omitempty"`
}

// NewDataInnerDnsAclAddSuccess instantiates a new DataInnerDnsAclAddSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerDnsAclAddSuccess() *DataInnerDnsAclAddSuccess {
	this := DataInnerDnsAclAddSuccess{}
	return &this
}

// NewDataInnerDnsAclAddSuccessWithDefaults instantiates a new DataInnerDnsAclAddSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerDnsAclAddSuccessWithDefaults() *DataInnerDnsAclAddSuccess {
	this := DataInnerDnsAclAddSuccess{}
	return &this
}

// GetAclId returns the AclId field value if set, zero value otherwise.
func (o *DataInnerDnsAclAddSuccess) GetAclId() string {
	if o == nil || IsNil(o.AclId) {
		var ret string
		return ret
	}
	return *o.AclId
}

// GetAclIdOk returns a tuple with the AclId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDnsAclAddSuccess) GetAclIdOk() (*string, bool) {
	if o == nil || IsNil(o.AclId) {
		return nil, false
	}
	return o.AclId, true
}

// HasAclId returns a boolean if a field has been set.
func (o *DataInnerDnsAclAddSuccess) HasAclId() bool {
	if o != nil && !IsNil(o.AclId) {
		return true
	}

	return false
}

// SetAclId gets a reference to the given string and assigns it to the AclId field.
func (o *DataInnerDnsAclAddSuccess) SetAclId(v string) {
	o.AclId = &v
}

func (o DataInnerDnsAclAddSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerDnsAclAddSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AclId) {
		toSerialize["acl_id"] = o.AclId
	}
	return toSerialize, nil
}

type NullableDataInnerDnsAclAddSuccess struct {
	value *DataInnerDnsAclAddSuccess
	isSet bool
}

func (v NullableDataInnerDnsAclAddSuccess) Get() *DataInnerDnsAclAddSuccess {
	return v.value
}

func (v *NullableDataInnerDnsAclAddSuccess) Set(val *DataInnerDnsAclAddSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerDnsAclAddSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerDnsAclAddSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerDnsAclAddSuccess(val *DataInnerDnsAclAddSuccess) *NullableDataInnerDnsAclAddSuccess {
	return &NullableDataInnerDnsAclAddSuccess{value: val, isSet: true}
}

func (v NullableDataInnerDnsAclAddSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerDnsAclAddSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
