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

// checks if the DataInnerDhcpAclentry6Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerDhcpAclentry6Data{}

// DataInnerDhcpAclentry6Data struct for DataInnerDhcpAclentry6Data
type DataInnerDhcpAclentry6Data struct {
	// The role of the server the object belongs to in the cluster, either <b>active (M)</b>, <b>passive (B)</b> or <b>N/A (#)</b>.
	ServerClusterRole *string `json:"server_cluster_role,omitempty"`
	// The delay of creation/deletion status. <b>1</b> indicates that the object is not created/deleted yet.
	Aclentry6DelayedTime *string `json:"aclentry6_delayed_time,omitempty"`
	// The database identifier (ID) of the DHCPv6 server the object belongs to.
	Server6Id *string `json:"server6_id,omitempty"`
	// The name of the DHCPv6 server the object belongs to.
	Server6Name *string `json:"server6_name,omitempty"`
	// The type of the DHCPv6 server the object belongs to: <table><caption>server6_type possible values</caption><br/><thead><tr><th>Type</th><th>Description</th></tr><br/></thead><br/><tbody><tr><td >ipm</td><td >EfficientIP or EfficientIP Package server</td></tr><tr><td >vdhcp</td><td >EfficientIP DHCPv6 smart architecture</td></tr></tbody></table></p><br/>
	Server6Type *string `json:"server6_type,omitempty"`
	// The database identifier (ID) of the DHCPv6 ACL.
	Acl6Id *string `json:"acl6_id,omitempty"`
	// The name of the DHCPv6 ACL.
	Acl6Name *string `json:"acl6_name,omitempty"`
	// The database identifier (ID) of the DHCPv6 ACL entry.
	Aclentry6Id *string `json:"aclentry6_id,omitempty"`
	// The value of the DHCPv6 ACL entry.
	Aclentry6Value *string `json:"aclentry6_value,omitempty"`
	// The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. <b>0</b> indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself.
	SmartParentId *string `json:"smart_parent_id,omitempty"`
}

// NewDataInnerDhcpAclentry6Data instantiates a new DataInnerDhcpAclentry6Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerDhcpAclentry6Data() *DataInnerDhcpAclentry6Data {
	this := DataInnerDhcpAclentry6Data{}
	return &this
}

// NewDataInnerDhcpAclentry6DataWithDefaults instantiates a new DataInnerDhcpAclentry6Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerDhcpAclentry6DataWithDefaults() *DataInnerDhcpAclentry6Data {
	this := DataInnerDhcpAclentry6Data{}
	return &this
}

// GetServerClusterRole returns the ServerClusterRole field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetServerClusterRole() string {
	if o == nil || IsNil(o.ServerClusterRole) {
		var ret string
		return ret
	}
	return *o.ServerClusterRole
}

// GetServerClusterRoleOk returns a tuple with the ServerClusterRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetServerClusterRoleOk() (*string, bool) {
	if o == nil || IsNil(o.ServerClusterRole) {
		return nil, false
	}
	return o.ServerClusterRole, true
}

// HasServerClusterRole returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasServerClusterRole() bool {
	if o != nil && !IsNil(o.ServerClusterRole) {
		return true
	}

	return false
}

// SetServerClusterRole gets a reference to the given string and assigns it to the ServerClusterRole field.
func (o *DataInnerDhcpAclentry6Data) SetServerClusterRole(v string) {
	o.ServerClusterRole = &v
}

// GetAclentry6DelayedTime returns the Aclentry6DelayedTime field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetAclentry6DelayedTime() string {
	if o == nil || IsNil(o.Aclentry6DelayedTime) {
		var ret string
		return ret
	}
	return *o.Aclentry6DelayedTime
}

// GetAclentry6DelayedTimeOk returns a tuple with the Aclentry6DelayedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetAclentry6DelayedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Aclentry6DelayedTime) {
		return nil, false
	}
	return o.Aclentry6DelayedTime, true
}

// HasAclentry6DelayedTime returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasAclentry6DelayedTime() bool {
	if o != nil && !IsNil(o.Aclentry6DelayedTime) {
		return true
	}

	return false
}

// SetAclentry6DelayedTime gets a reference to the given string and assigns it to the Aclentry6DelayedTime field.
func (o *DataInnerDhcpAclentry6Data) SetAclentry6DelayedTime(v string) {
	o.Aclentry6DelayedTime = &v
}

// GetServer6Id returns the Server6Id field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetServer6Id() string {
	if o == nil || IsNil(o.Server6Id) {
		var ret string
		return ret
	}
	return *o.Server6Id
}

// GetServer6IdOk returns a tuple with the Server6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetServer6IdOk() (*string, bool) {
	if o == nil || IsNil(o.Server6Id) {
		return nil, false
	}
	return o.Server6Id, true
}

// HasServer6Id returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasServer6Id() bool {
	if o != nil && !IsNil(o.Server6Id) {
		return true
	}

	return false
}

// SetServer6Id gets a reference to the given string and assigns it to the Server6Id field.
func (o *DataInnerDhcpAclentry6Data) SetServer6Id(v string) {
	o.Server6Id = &v
}

// GetServer6Name returns the Server6Name field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetServer6Name() string {
	if o == nil || IsNil(o.Server6Name) {
		var ret string
		return ret
	}
	return *o.Server6Name
}

// GetServer6NameOk returns a tuple with the Server6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetServer6NameOk() (*string, bool) {
	if o == nil || IsNil(o.Server6Name) {
		return nil, false
	}
	return o.Server6Name, true
}

// HasServer6Name returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasServer6Name() bool {
	if o != nil && !IsNil(o.Server6Name) {
		return true
	}

	return false
}

// SetServer6Name gets a reference to the given string and assigns it to the Server6Name field.
func (o *DataInnerDhcpAclentry6Data) SetServer6Name(v string) {
	o.Server6Name = &v
}

// GetServer6Type returns the Server6Type field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetServer6Type() string {
	if o == nil || IsNil(o.Server6Type) {
		var ret string
		return ret
	}
	return *o.Server6Type
}

// GetServer6TypeOk returns a tuple with the Server6Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetServer6TypeOk() (*string, bool) {
	if o == nil || IsNil(o.Server6Type) {
		return nil, false
	}
	return o.Server6Type, true
}

// HasServer6Type returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasServer6Type() bool {
	if o != nil && !IsNil(o.Server6Type) {
		return true
	}

	return false
}

// SetServer6Type gets a reference to the given string and assigns it to the Server6Type field.
func (o *DataInnerDhcpAclentry6Data) SetServer6Type(v string) {
	o.Server6Type = &v
}

// GetAcl6Id returns the Acl6Id field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetAcl6Id() string {
	if o == nil || IsNil(o.Acl6Id) {
		var ret string
		return ret
	}
	return *o.Acl6Id
}

// GetAcl6IdOk returns a tuple with the Acl6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetAcl6IdOk() (*string, bool) {
	if o == nil || IsNil(o.Acl6Id) {
		return nil, false
	}
	return o.Acl6Id, true
}

// HasAcl6Id returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasAcl6Id() bool {
	if o != nil && !IsNil(o.Acl6Id) {
		return true
	}

	return false
}

// SetAcl6Id gets a reference to the given string and assigns it to the Acl6Id field.
func (o *DataInnerDhcpAclentry6Data) SetAcl6Id(v string) {
	o.Acl6Id = &v
}

// GetAcl6Name returns the Acl6Name field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetAcl6Name() string {
	if o == nil || IsNil(o.Acl6Name) {
		var ret string
		return ret
	}
	return *o.Acl6Name
}

// GetAcl6NameOk returns a tuple with the Acl6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetAcl6NameOk() (*string, bool) {
	if o == nil || IsNil(o.Acl6Name) {
		return nil, false
	}
	return o.Acl6Name, true
}

// HasAcl6Name returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasAcl6Name() bool {
	if o != nil && !IsNil(o.Acl6Name) {
		return true
	}

	return false
}

// SetAcl6Name gets a reference to the given string and assigns it to the Acl6Name field.
func (o *DataInnerDhcpAclentry6Data) SetAcl6Name(v string) {
	o.Acl6Name = &v
}

// GetAclentry6Id returns the Aclentry6Id field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetAclentry6Id() string {
	if o == nil || IsNil(o.Aclentry6Id) {
		var ret string
		return ret
	}
	return *o.Aclentry6Id
}

// GetAclentry6IdOk returns a tuple with the Aclentry6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetAclentry6IdOk() (*string, bool) {
	if o == nil || IsNil(o.Aclentry6Id) {
		return nil, false
	}
	return o.Aclentry6Id, true
}

// HasAclentry6Id returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasAclentry6Id() bool {
	if o != nil && !IsNil(o.Aclentry6Id) {
		return true
	}

	return false
}

// SetAclentry6Id gets a reference to the given string and assigns it to the Aclentry6Id field.
func (o *DataInnerDhcpAclentry6Data) SetAclentry6Id(v string) {
	o.Aclentry6Id = &v
}

// GetAclentry6Value returns the Aclentry6Value field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetAclentry6Value() string {
	if o == nil || IsNil(o.Aclentry6Value) {
		var ret string
		return ret
	}
	return *o.Aclentry6Value
}

// GetAclentry6ValueOk returns a tuple with the Aclentry6Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetAclentry6ValueOk() (*string, bool) {
	if o == nil || IsNil(o.Aclentry6Value) {
		return nil, false
	}
	return o.Aclentry6Value, true
}

// HasAclentry6Value returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasAclentry6Value() bool {
	if o != nil && !IsNil(o.Aclentry6Value) {
		return true
	}

	return false
}

// SetAclentry6Value gets a reference to the given string and assigns it to the Aclentry6Value field.
func (o *DataInnerDhcpAclentry6Data) SetAclentry6Value(v string) {
	o.Aclentry6Value = &v
}

// GetSmartParentId returns the SmartParentId field value if set, zero value otherwise.
func (o *DataInnerDhcpAclentry6Data) GetSmartParentId() string {
	if o == nil || IsNil(o.SmartParentId) {
		var ret string
		return ret
	}
	return *o.SmartParentId
}

// GetSmartParentIdOk returns a tuple with the SmartParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpAclentry6Data) GetSmartParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SmartParentId) {
		return nil, false
	}
	return o.SmartParentId, true
}

// HasSmartParentId returns a boolean if a field has been set.
func (o *DataInnerDhcpAclentry6Data) HasSmartParentId() bool {
	if o != nil && !IsNil(o.SmartParentId) {
		return true
	}

	return false
}

// SetSmartParentId gets a reference to the given string and assigns it to the SmartParentId field.
func (o *DataInnerDhcpAclentry6Data) SetSmartParentId(v string) {
	o.SmartParentId = &v
}

func (o DataInnerDhcpAclentry6Data) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerDhcpAclentry6Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerClusterRole) {
		toSerialize["server_cluster_role"] = o.ServerClusterRole
	}
	if !IsNil(o.Aclentry6DelayedTime) {
		toSerialize["aclentry6_delayed_time"] = o.Aclentry6DelayedTime
	}
	if !IsNil(o.Server6Id) {
		toSerialize["server6_id"] = o.Server6Id
	}
	if !IsNil(o.Server6Name) {
		toSerialize["server6_name"] = o.Server6Name
	}
	if !IsNil(o.Server6Type) {
		toSerialize["server6_type"] = o.Server6Type
	}
	if !IsNil(o.Acl6Id) {
		toSerialize["acl6_id"] = o.Acl6Id
	}
	if !IsNil(o.Acl6Name) {
		toSerialize["acl6_name"] = o.Acl6Name
	}
	if !IsNil(o.Aclentry6Id) {
		toSerialize["aclentry6_id"] = o.Aclentry6Id
	}
	if !IsNil(o.Aclentry6Value) {
		toSerialize["aclentry6_value"] = o.Aclentry6Value
	}
	if !IsNil(o.SmartParentId) {
		toSerialize["smart_parent_id"] = o.SmartParentId
	}
	return toSerialize, nil
}

type NullableDataInnerDhcpAclentry6Data struct {
	value *DataInnerDhcpAclentry6Data
	isSet bool
}

func (v NullableDataInnerDhcpAclentry6Data) Get() *DataInnerDhcpAclentry6Data {
	return v.value
}

func (v *NullableDataInnerDhcpAclentry6Data) Set(val *DataInnerDhcpAclentry6Data) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerDhcpAclentry6Data) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerDhcpAclentry6Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerDhcpAclentry6Data(val *DataInnerDhcpAclentry6Data) *NullableDataInnerDhcpAclentry6Data {
	return &NullableDataInnerDhcpAclentry6Data{value: val, isSet: true}
}

func (v NullableDataInnerDhcpAclentry6Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerDhcpAclentry6Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
