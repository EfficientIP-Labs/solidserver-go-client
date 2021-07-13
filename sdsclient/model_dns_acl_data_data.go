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

// DnsAclDataData struct for DnsAclDataData
type DnsAclDataData struct {
	// The database identifier (ID) of the DNS server the object belongs to.
	ServerId *string `json:"server_id,omitempty"`
	// The database identifier (ID) of the DNS ACL.
	AclId *string `json:"acl_id,omitempty"`
	// The name of the DNS ACL.
	AclName *string `json:"acl_name,omitempty"`
	// The values of the DNS ACL in order of priority, as follows: &lt;value_1&gt;;&lt;value_2&gt;... .
	AclValue *string `json:"acl_value,omitempty"`
}

// NewDnsAclDataData instantiates a new DnsAclDataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsAclDataData() *DnsAclDataData {
	this := DnsAclDataData{}
	return &this
}

// NewDnsAclDataDataWithDefaults instantiates a new DnsAclDataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsAclDataDataWithDefaults() *DnsAclDataData {
	this := DnsAclDataData{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *DnsAclDataData) GetServerId() string {
	if o == nil || o.ServerId == nil {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAclDataData) GetServerIdOk() (*string, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *DnsAclDataData) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *DnsAclDataData) SetServerId(v string) {
	o.ServerId = &v
}

// GetAclId returns the AclId field value if set, zero value otherwise.
func (o *DnsAclDataData) GetAclId() string {
	if o == nil || o.AclId == nil {
		var ret string
		return ret
	}
	return *o.AclId
}

// GetAclIdOk returns a tuple with the AclId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAclDataData) GetAclIdOk() (*string, bool) {
	if o == nil || o.AclId == nil {
		return nil, false
	}
	return o.AclId, true
}

// HasAclId returns a boolean if a field has been set.
func (o *DnsAclDataData) HasAclId() bool {
	if o != nil && o.AclId != nil {
		return true
	}

	return false
}

// SetAclId gets a reference to the given string and assigns it to the AclId field.
func (o *DnsAclDataData) SetAclId(v string) {
	o.AclId = &v
}

// GetAclName returns the AclName field value if set, zero value otherwise.
func (o *DnsAclDataData) GetAclName() string {
	if o == nil || o.AclName == nil {
		var ret string
		return ret
	}
	return *o.AclName
}

// GetAclNameOk returns a tuple with the AclName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAclDataData) GetAclNameOk() (*string, bool) {
	if o == nil || o.AclName == nil {
		return nil, false
	}
	return o.AclName, true
}

// HasAclName returns a boolean if a field has been set.
func (o *DnsAclDataData) HasAclName() bool {
	if o != nil && o.AclName != nil {
		return true
	}

	return false
}

// SetAclName gets a reference to the given string and assigns it to the AclName field.
func (o *DnsAclDataData) SetAclName(v string) {
	o.AclName = &v
}

// GetAclValue returns the AclValue field value if set, zero value otherwise.
func (o *DnsAclDataData) GetAclValue() string {
	if o == nil || o.AclValue == nil {
		var ret string
		return ret
	}
	return *o.AclValue
}

// GetAclValueOk returns a tuple with the AclValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsAclDataData) GetAclValueOk() (*string, bool) {
	if o == nil || o.AclValue == nil {
		return nil, false
	}
	return o.AclValue, true
}

// HasAclValue returns a boolean if a field has been set.
func (o *DnsAclDataData) HasAclValue() bool {
	if o != nil && o.AclValue != nil {
		return true
	}

	return false
}

// SetAclValue gets a reference to the given string and assigns it to the AclValue field.
func (o *DnsAclDataData) SetAclValue(v string) {
	o.AclValue = &v
}

func (o DnsAclDataData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ServerId != nil {
		toSerialize["server_id"] = o.ServerId
	}
	if o.AclId != nil {
		toSerialize["acl_id"] = o.AclId
	}
	if o.AclName != nil {
		toSerialize["acl_name"] = o.AclName
	}
	if o.AclValue != nil {
		toSerialize["acl_value"] = o.AclValue
	}
	return json.Marshal(toSerialize)
}

type NullableDnsAclDataData struct {
	value *DnsAclDataData
	isSet bool
}

func (v NullableDnsAclDataData) Get() *DnsAclDataData {
	return v.value
}

func (v *NullableDnsAclDataData) Set(val *DnsAclDataData) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsAclDataData) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsAclDataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsAclDataData(val *DnsAclDataData) *NullableDnsAclDataData {
	return &NullableDnsAclDataData{value: val, isSet: true}
}

func (v NullableDnsAclDataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsAclDataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

