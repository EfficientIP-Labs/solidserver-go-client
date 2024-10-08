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

// checks if the DataInnerIpamNetworkDeleteSuccess type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerIpamNetworkDeleteSuccess{}

// DataInnerIpamNetworkDeleteSuccess struct for DataInnerIpamNetworkDeleteSuccess
type DataInnerIpamNetworkDeleteSuccess struct {
	// The database identifier (ID) of the object you deleted.
	NetworkId *string `json:"network_id,omitempty"`
}

// NewDataInnerIpamNetworkDeleteSuccess instantiates a new DataInnerIpamNetworkDeleteSuccess object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerIpamNetworkDeleteSuccess() *DataInnerIpamNetworkDeleteSuccess {
	this := DataInnerIpamNetworkDeleteSuccess{}
	return &this
}

// NewDataInnerIpamNetworkDeleteSuccessWithDefaults instantiates a new DataInnerIpamNetworkDeleteSuccess object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerIpamNetworkDeleteSuccessWithDefaults() *DataInnerIpamNetworkDeleteSuccess {
	this := DataInnerIpamNetworkDeleteSuccess{}
	return &this
}

// GetNetworkId returns the NetworkId field value if set, zero value otherwise.
func (o *DataInnerIpamNetworkDeleteSuccess) GetNetworkId() string {
	if o == nil || IsNil(o.NetworkId) {
		var ret string
		return ret
	}
	return *o.NetworkId
}

// GetNetworkIdOk returns a tuple with the NetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerIpamNetworkDeleteSuccess) GetNetworkIdOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkId) {
		return nil, false
	}
	return o.NetworkId, true
}

// HasNetworkId returns a boolean if a field has been set.
func (o *DataInnerIpamNetworkDeleteSuccess) HasNetworkId() bool {
	if o != nil && !IsNil(o.NetworkId) {
		return true
	}

	return false
}

// SetNetworkId gets a reference to the given string and assigns it to the NetworkId field.
func (o *DataInnerIpamNetworkDeleteSuccess) SetNetworkId(v string) {
	o.NetworkId = &v
}

func (o DataInnerIpamNetworkDeleteSuccess) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerIpamNetworkDeleteSuccess) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkId) {
		toSerialize["network_id"] = o.NetworkId
	}
	return toSerialize, nil
}

type NullableDataInnerIpamNetworkDeleteSuccess struct {
	value *DataInnerIpamNetworkDeleteSuccess
	isSet bool
}

func (v NullableDataInnerIpamNetworkDeleteSuccess) Get() *DataInnerIpamNetworkDeleteSuccess {
	return v.value
}

func (v *NullableDataInnerIpamNetworkDeleteSuccess) Set(val *DataInnerIpamNetworkDeleteSuccess) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerIpamNetworkDeleteSuccess) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerIpamNetworkDeleteSuccess) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerIpamNetworkDeleteSuccess(val *DataInnerIpamNetworkDeleteSuccess) *NullableDataInnerIpamNetworkDeleteSuccess {
	return &NullableDataInnerIpamNetworkDeleteSuccess{value: val, isSet: true}
}

func (v NullableDataInnerIpamNetworkDeleteSuccess) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerIpamNetworkDeleteSuccess) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}