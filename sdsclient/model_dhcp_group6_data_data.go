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

// DhcpGroup6DataData struct for DhcpGroup6DataData
type DhcpGroup6DataData struct {
	// The delay of creation/deletion status. <b>1</b> indicates that the object is not created/deleted yet.
	Group6DelayedTime *string `json:"group6_delayed_time,omitempty"`
	// The name of the class applied to the DHCPv6 server the object belongs to, it can be preceded by the class directory.
	Server6ClassName *string `json:"server6_class_name,omitempty"`
	// The class parameters applied to the DHCPv6 server the object belongs to and their value: <b>&lt;class-parameter1&gt;=&lt;value1&gt;&amp;&lt;class-parameter2&gt;=&lt;value2&gt;&amp;</b>... .
	Server6ClassParameters *[]ApiClassParameterOutputEntry `json:"server6_class_parameters,omitempty"`
	// The database identifier (ID) of the DHCPv6 server the object belongs to.
	Server6Id *string `json:"server6_id,omitempty"`
	// The name of the DHCPv6 server the object belongs to.
	Server6Name *string `json:"server6_name,omitempty"`
	// The type of the DHCPv6 server the object belongs to: <table><caption>dhcp6_type possible values</caption><br/><thead><tr><th>Type</th><th>Description</th></tr><br/></thead><br/><tbody><tr><td >ipm</td><td >EfficientIP DHCP server or EfficientIP DHCP Package</td></tr><tr><td >vdhcp</td><td >EfficientIP DHCP smart architecture</td></tr></tbody></table></p><br/>
	Server6Type *string `json:"server6_type,omitempty"`
	// The version details of the DHCPv6 server the object belongs to.
	Server6Version *string `json:"server6_version,omitempty"`
	// The name of the class applied to the DHCPv6 group, it can be preceded by the class directory.
	Group6ClassName *string `json:"group6_class_name,omitempty"`
	// The class parameters applied to the DHCPv6 group and their value: <b>&lt;class-parameter1&gt;=&lt;value1&gt;&amp;&lt;class-parameter2&gt;=&lt;value2&gt;&amp;</b>... .
	Group6ClassParameters *[]ApiClassParameterOutputEntry `json:"group6_class_parameters,omitempty"`
	// The database identifier (ID) of the DHCPv6 group.
	Group6Id *string `json:"group6_id,omitempty"`
	// The name of the DHCPv6 group.
	Group6Name *string `json:"group6_name,omitempty"`
	// The IP address of the DHCP server the object belongs to, in hexadecimal format.
	Server6Addr *string `json:"server6_addr,omitempty"`
	// The Multi-status information is displayed as follows: <i>&lt;number-of-instances&gt;@&lt;message-number&gt;@&lt;multi-status-severity&gt;@&lt;module&gt;</i>. The different severity levels are:<br><b>Multi-status severity levels</b>    <table border=1>        <thead>        <tr >            <td><b>Message number</b></td>            <td><b>Severity</b></td>            <td><b>Description</b></td>        </tr>        </thead>        <tbody>        <tr  valign=middle>            <td>0 - 16</td>            <td>Emergency</td>            <td>The object configuration prevents the system from running properly. Action is required.</td>        </tr>        <tr  valign=middle>            <td>17 - 33</td>            <td>Critical</td>            <td>The object configuration is in critical conditions. Immediate action is recommended.</td>        </tr>        <tr  valign=middle>            <td>34 - 50</td>            <td>Error</td>            <td>The object configuration failed at some level. Action is recommended.</td>        </tr>        <tr  valign=middle>            <td>51 - 66</td>            <td>Warning</td>            <td>The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.</td>        </tr>        <tr  valign=middle>            <td>67 - 83</td>            <td>Notice</td>            <td>The object configuration is normal but undergoing events that might trigger errors. No immediate action required.</td>        </tr>        <tr  valign=middle>            <td>84 - 100</td>            <td>Informational</td>            <td>The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.</td>        </tr>        </tbody></table>
	Group6Multistatus *string `json:"group6_multistatus,omitempty"`
	// The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. <b>0</b> indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself.
	SmartParentId *string `json:"smart_parent_id,omitempty"`
	// The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. <b>#</b> indicates that the server is not managed by a smart architecture or is a smart architecture itself.
	SmartParentName *string `json:"smart_parent_name,omitempty"`
}

// NewDhcpGroup6DataData instantiates a new DhcpGroup6DataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpGroup6DataData() *DhcpGroup6DataData {
	this := DhcpGroup6DataData{}
	return &this
}

// NewDhcpGroup6DataDataWithDefaults instantiates a new DhcpGroup6DataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpGroup6DataDataWithDefaults() *DhcpGroup6DataData {
	this := DhcpGroup6DataData{}
	return &this
}

// GetGroup6DelayedTime returns the Group6DelayedTime field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetGroup6DelayedTime() string {
	if o == nil || o.Group6DelayedTime == nil {
		var ret string
		return ret
	}
	return *o.Group6DelayedTime
}

// GetGroup6DelayedTimeOk returns a tuple with the Group6DelayedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetGroup6DelayedTimeOk() (*string, bool) {
	if o == nil || o.Group6DelayedTime == nil {
		return nil, false
	}
	return o.Group6DelayedTime, true
}

// HasGroup6DelayedTime returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasGroup6DelayedTime() bool {
	if o != nil && o.Group6DelayedTime != nil {
		return true
	}

	return false
}

// SetGroup6DelayedTime gets a reference to the given string and assigns it to the Group6DelayedTime field.
func (o *DhcpGroup6DataData) SetGroup6DelayedTime(v string) {
	o.Group6DelayedTime = &v
}

// GetServer6ClassName returns the Server6ClassName field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetServer6ClassName() string {
	if o == nil || o.Server6ClassName == nil {
		var ret string
		return ret
	}
	return *o.Server6ClassName
}

// GetServer6ClassNameOk returns a tuple with the Server6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetServer6ClassNameOk() (*string, bool) {
	if o == nil || o.Server6ClassName == nil {
		return nil, false
	}
	return o.Server6ClassName, true
}

// HasServer6ClassName returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasServer6ClassName() bool {
	if o != nil && o.Server6ClassName != nil {
		return true
	}

	return false
}

// SetServer6ClassName gets a reference to the given string and assigns it to the Server6ClassName field.
func (o *DhcpGroup6DataData) SetServer6ClassName(v string) {
	o.Server6ClassName = &v
}

// GetServer6ClassParameters returns the Server6ClassParameters field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetServer6ClassParameters() []ApiClassParameterOutputEntry {
	if o == nil || o.Server6ClassParameters == nil {
		var ret []ApiClassParameterOutputEntry
		return ret
	}
	return *o.Server6ClassParameters
}

// GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool) {
	if o == nil || o.Server6ClassParameters == nil {
		return nil, false
	}
	return o.Server6ClassParameters, true
}

// HasServer6ClassParameters returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasServer6ClassParameters() bool {
	if o != nil && o.Server6ClassParameters != nil {
		return true
	}

	return false
}

// SetServer6ClassParameters gets a reference to the given []ApiClassParameterOutputEntry and assigns it to the Server6ClassParameters field.
func (o *DhcpGroup6DataData) SetServer6ClassParameters(v []ApiClassParameterOutputEntry) {
	o.Server6ClassParameters = &v
}

// GetServer6Id returns the Server6Id field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetServer6Id() string {
	if o == nil || o.Server6Id == nil {
		var ret string
		return ret
	}
	return *o.Server6Id
}

// GetServer6IdOk returns a tuple with the Server6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetServer6IdOk() (*string, bool) {
	if o == nil || o.Server6Id == nil {
		return nil, false
	}
	return o.Server6Id, true
}

// HasServer6Id returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasServer6Id() bool {
	if o != nil && o.Server6Id != nil {
		return true
	}

	return false
}

// SetServer6Id gets a reference to the given string and assigns it to the Server6Id field.
func (o *DhcpGroup6DataData) SetServer6Id(v string) {
	o.Server6Id = &v
}

// GetServer6Name returns the Server6Name field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetServer6Name() string {
	if o == nil || o.Server6Name == nil {
		var ret string
		return ret
	}
	return *o.Server6Name
}

// GetServer6NameOk returns a tuple with the Server6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetServer6NameOk() (*string, bool) {
	if o == nil || o.Server6Name == nil {
		return nil, false
	}
	return o.Server6Name, true
}

// HasServer6Name returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasServer6Name() bool {
	if o != nil && o.Server6Name != nil {
		return true
	}

	return false
}

// SetServer6Name gets a reference to the given string and assigns it to the Server6Name field.
func (o *DhcpGroup6DataData) SetServer6Name(v string) {
	o.Server6Name = &v
}

// GetServer6Type returns the Server6Type field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetServer6Type() string {
	if o == nil || o.Server6Type == nil {
		var ret string
		return ret
	}
	return *o.Server6Type
}

// GetServer6TypeOk returns a tuple with the Server6Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetServer6TypeOk() (*string, bool) {
	if o == nil || o.Server6Type == nil {
		return nil, false
	}
	return o.Server6Type, true
}

// HasServer6Type returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasServer6Type() bool {
	if o != nil && o.Server6Type != nil {
		return true
	}

	return false
}

// SetServer6Type gets a reference to the given string and assigns it to the Server6Type field.
func (o *DhcpGroup6DataData) SetServer6Type(v string) {
	o.Server6Type = &v
}

// GetServer6Version returns the Server6Version field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetServer6Version() string {
	if o == nil || o.Server6Version == nil {
		var ret string
		return ret
	}
	return *o.Server6Version
}

// GetServer6VersionOk returns a tuple with the Server6Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetServer6VersionOk() (*string, bool) {
	if o == nil || o.Server6Version == nil {
		return nil, false
	}
	return o.Server6Version, true
}

// HasServer6Version returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasServer6Version() bool {
	if o != nil && o.Server6Version != nil {
		return true
	}

	return false
}

// SetServer6Version gets a reference to the given string and assigns it to the Server6Version field.
func (o *DhcpGroup6DataData) SetServer6Version(v string) {
	o.Server6Version = &v
}

// GetGroup6ClassName returns the Group6ClassName field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetGroup6ClassName() string {
	if o == nil || o.Group6ClassName == nil {
		var ret string
		return ret
	}
	return *o.Group6ClassName
}

// GetGroup6ClassNameOk returns a tuple with the Group6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetGroup6ClassNameOk() (*string, bool) {
	if o == nil || o.Group6ClassName == nil {
		return nil, false
	}
	return o.Group6ClassName, true
}

// HasGroup6ClassName returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasGroup6ClassName() bool {
	if o != nil && o.Group6ClassName != nil {
		return true
	}

	return false
}

// SetGroup6ClassName gets a reference to the given string and assigns it to the Group6ClassName field.
func (o *DhcpGroup6DataData) SetGroup6ClassName(v string) {
	o.Group6ClassName = &v
}

// GetGroup6ClassParameters returns the Group6ClassParameters field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetGroup6ClassParameters() []ApiClassParameterOutputEntry {
	if o == nil || o.Group6ClassParameters == nil {
		var ret []ApiClassParameterOutputEntry
		return ret
	}
	return *o.Group6ClassParameters
}

// GetGroup6ClassParametersOk returns a tuple with the Group6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetGroup6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool) {
	if o == nil || o.Group6ClassParameters == nil {
		return nil, false
	}
	return o.Group6ClassParameters, true
}

// HasGroup6ClassParameters returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasGroup6ClassParameters() bool {
	if o != nil && o.Group6ClassParameters != nil {
		return true
	}

	return false
}

// SetGroup6ClassParameters gets a reference to the given []ApiClassParameterOutputEntry and assigns it to the Group6ClassParameters field.
func (o *DhcpGroup6DataData) SetGroup6ClassParameters(v []ApiClassParameterOutputEntry) {
	o.Group6ClassParameters = &v
}

// GetGroup6Id returns the Group6Id field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetGroup6Id() string {
	if o == nil || o.Group6Id == nil {
		var ret string
		return ret
	}
	return *o.Group6Id
}

// GetGroup6IdOk returns a tuple with the Group6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetGroup6IdOk() (*string, bool) {
	if o == nil || o.Group6Id == nil {
		return nil, false
	}
	return o.Group6Id, true
}

// HasGroup6Id returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasGroup6Id() bool {
	if o != nil && o.Group6Id != nil {
		return true
	}

	return false
}

// SetGroup6Id gets a reference to the given string and assigns it to the Group6Id field.
func (o *DhcpGroup6DataData) SetGroup6Id(v string) {
	o.Group6Id = &v
}

// GetGroup6Name returns the Group6Name field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetGroup6Name() string {
	if o == nil || o.Group6Name == nil {
		var ret string
		return ret
	}
	return *o.Group6Name
}

// GetGroup6NameOk returns a tuple with the Group6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetGroup6NameOk() (*string, bool) {
	if o == nil || o.Group6Name == nil {
		return nil, false
	}
	return o.Group6Name, true
}

// HasGroup6Name returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasGroup6Name() bool {
	if o != nil && o.Group6Name != nil {
		return true
	}

	return false
}

// SetGroup6Name gets a reference to the given string and assigns it to the Group6Name field.
func (o *DhcpGroup6DataData) SetGroup6Name(v string) {
	o.Group6Name = &v
}

// GetServer6Addr returns the Server6Addr field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetServer6Addr() string {
	if o == nil || o.Server6Addr == nil {
		var ret string
		return ret
	}
	return *o.Server6Addr
}

// GetServer6AddrOk returns a tuple with the Server6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetServer6AddrOk() (*string, bool) {
	if o == nil || o.Server6Addr == nil {
		return nil, false
	}
	return o.Server6Addr, true
}

// HasServer6Addr returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasServer6Addr() bool {
	if o != nil && o.Server6Addr != nil {
		return true
	}

	return false
}

// SetServer6Addr gets a reference to the given string and assigns it to the Server6Addr field.
func (o *DhcpGroup6DataData) SetServer6Addr(v string) {
	o.Server6Addr = &v
}

// GetGroup6Multistatus returns the Group6Multistatus field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetGroup6Multistatus() string {
	if o == nil || o.Group6Multistatus == nil {
		var ret string
		return ret
	}
	return *o.Group6Multistatus
}

// GetGroup6MultistatusOk returns a tuple with the Group6Multistatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetGroup6MultistatusOk() (*string, bool) {
	if o == nil || o.Group6Multistatus == nil {
		return nil, false
	}
	return o.Group6Multistatus, true
}

// HasGroup6Multistatus returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasGroup6Multistatus() bool {
	if o != nil && o.Group6Multistatus != nil {
		return true
	}

	return false
}

// SetGroup6Multistatus gets a reference to the given string and assigns it to the Group6Multistatus field.
func (o *DhcpGroup6DataData) SetGroup6Multistatus(v string) {
	o.Group6Multistatus = &v
}

// GetSmartParentId returns the SmartParentId field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetSmartParentId() string {
	if o == nil || o.SmartParentId == nil {
		var ret string
		return ret
	}
	return *o.SmartParentId
}

// GetSmartParentIdOk returns a tuple with the SmartParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetSmartParentIdOk() (*string, bool) {
	if o == nil || o.SmartParentId == nil {
		return nil, false
	}
	return o.SmartParentId, true
}

// HasSmartParentId returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasSmartParentId() bool {
	if o != nil && o.SmartParentId != nil {
		return true
	}

	return false
}

// SetSmartParentId gets a reference to the given string and assigns it to the SmartParentId field.
func (o *DhcpGroup6DataData) SetSmartParentId(v string) {
	o.SmartParentId = &v
}

// GetSmartParentName returns the SmartParentName field value if set, zero value otherwise.
func (o *DhcpGroup6DataData) GetSmartParentName() string {
	if o == nil || o.SmartParentName == nil {
		var ret string
		return ret
	}
	return *o.SmartParentName
}

// GetSmartParentNameOk returns a tuple with the SmartParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpGroup6DataData) GetSmartParentNameOk() (*string, bool) {
	if o == nil || o.SmartParentName == nil {
		return nil, false
	}
	return o.SmartParentName, true
}

// HasSmartParentName returns a boolean if a field has been set.
func (o *DhcpGroup6DataData) HasSmartParentName() bool {
	if o != nil && o.SmartParentName != nil {
		return true
	}

	return false
}

// SetSmartParentName gets a reference to the given string and assigns it to the SmartParentName field.
func (o *DhcpGroup6DataData) SetSmartParentName(v string) {
	o.SmartParentName = &v
}

func (o DhcpGroup6DataData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Group6DelayedTime != nil {
		toSerialize["group6_delayed_time"] = o.Group6DelayedTime
	}
	if o.Server6ClassName != nil {
		toSerialize["server6_class_name"] = o.Server6ClassName
	}
	if o.Server6ClassParameters != nil {
		toSerialize["server6_class_parameters"] = o.Server6ClassParameters
	}
	if o.Server6Id != nil {
		toSerialize["server6_id"] = o.Server6Id
	}
	if o.Server6Name != nil {
		toSerialize["server6_name"] = o.Server6Name
	}
	if o.Server6Type != nil {
		toSerialize["server6_type"] = o.Server6Type
	}
	if o.Server6Version != nil {
		toSerialize["server6_version"] = o.Server6Version
	}
	if o.Group6ClassName != nil {
		toSerialize["group6_class_name"] = o.Group6ClassName
	}
	if o.Group6ClassParameters != nil {
		toSerialize["group6_class_parameters"] = o.Group6ClassParameters
	}
	if o.Group6Id != nil {
		toSerialize["group6_id"] = o.Group6Id
	}
	if o.Group6Name != nil {
		toSerialize["group6_name"] = o.Group6Name
	}
	if o.Server6Addr != nil {
		toSerialize["server6_addr"] = o.Server6Addr
	}
	if o.Group6Multistatus != nil {
		toSerialize["group6_multistatus"] = o.Group6Multistatus
	}
	if o.SmartParentId != nil {
		toSerialize["smart_parent_id"] = o.SmartParentId
	}
	if o.SmartParentName != nil {
		toSerialize["smart_parent_name"] = o.SmartParentName
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpGroup6DataData struct {
	value *DhcpGroup6DataData
	isSet bool
}

func (v NullableDhcpGroup6DataData) Get() *DhcpGroup6DataData {
	return v.value
}

func (v *NullableDhcpGroup6DataData) Set(val *DhcpGroup6DataData) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpGroup6DataData) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpGroup6DataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpGroup6DataData(val *DhcpGroup6DataData) *NullableDhcpGroup6DataData {
	return &NullableDhcpGroup6DataData{value: val, isSet: true}
}

func (v NullableDhcpGroup6DataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpGroup6DataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


