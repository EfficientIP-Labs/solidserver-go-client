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

// checks if the DataInnerDhcpSharednetwork6Data type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DataInnerDhcpSharednetwork6Data{}

// DataInnerDhcpSharednetwork6Data struct for DataInnerDhcpSharednetwork6Data
type DataInnerDhcpSharednetwork6Data struct {
	// The delay of creation/deletion status. <b>1</b> indicates that the object is not created/deleted yet.
	Sharednetwork6DelayedTime *string `json:"sharednetwork6_delayed_time,omitempty"`
	// The database identifier (ID) of the DHCPv6 server the object belongs to.
	Server6Id *string `json:"server6_id,omitempty"`
	// The name of the DHCPv6 server the object belongs to.
	Server6Name *string `json:"server6_name,omitempty"`
	// The type of the DHCPv6 server the object belongs to: <table><caption>server6_type possible values</caption><br/><thead><tr><th>Type</th><th>Description</th></tr><br/></thead><br/><tbody><tr><td >ipm</td><td >EfficientIP or EfficientIP Package server</td></tr><tr><td >vdhcp</td><td >EfficientIP DHCPv6 smart architecture</td></tr></tbody></table></p><br/>
	Server6Type *string `json:"server6_type,omitempty"`
	// The database identifier (ID) of the DHCPv6 shared network.
	Sharednetwork6Id *string `json:"sharednetwork6_id,omitempty"`
	// The name of the DHCPv6 shared network.
	Sharednetwork6Name *string `json:"sharednetwork6_name,omitempty"`
	// The Multi-status information is displayed as follows: <i>&lt;number-of-instances&gt;@&lt;message-number&gt;@&lt;multi-status-severity&gt;@&lt;module&gt;</i>. The different severity levels are:<br><b>Multi-status severity levels</b>    <table border=1>        <thead>        <tr >            <td><b>Message number</b></td>            <td><b>Severity</b></td>            <td><b>Description</b></td>        </tr>        </thead>        <tbody>        <tr  valign=middle>            <td>0 - 16</td>            <td>Emergency</td>            <td>The object configuration prevents the system from running properly. Action is required.</td>        </tr>        <tr  valign=middle>            <td>17 - 33</td>            <td>Critical</td>            <td>The object configuration is in critical conditions. Immediate action is recommended.</td>        </tr>        <tr  valign=middle>            <td>34 - 50</td>            <td>Error</td>            <td>The object configuration failed at some level. Action is recommended.</td>        </tr>        <tr  valign=middle>            <td>51 - 66</td>            <td>Warning</td>            <td>The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.</td>        </tr>        <tr  valign=middle>            <td>67 - 83</td>            <td>Notice</td>            <td>The object configuration is normal but undergoing events that might trigger errors. No immediate action required.</td>        </tr>        <tr  valign=middle>            <td>84 - 100</td>            <td>Informational</td>            <td>The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.</td>        </tr>        </tbody></table>
	Sharednetwork6Multistatus *string `json:"sharednetwork6_multistatus,omitempty"`
	// The type of the DHCPv6 smart architecture the object belongs to.<table><caption>smart_arch possible values</caption><br/><thead><tr><th>Type</th><th>Description</th></tr><br/></thead><br/><tbody><tr><td >single</td><td >The Single-Server smart architecture manages a single DHCPv6 server.</td></tr><tr><td >splitscope</td><td >The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.</td></tr><tr><td >stateless</td><td >The Stateless smart architecture offers a limited number of options to the DHCP clients. The IP address is delivered thanks to the subnet gateway and it is impossible to create any ranges or statics or to retrieve any leases.</td></tr></tbody></table></p><br/>
	SmartArch *string `json:"smart_arch,omitempty"`
	// The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. <b>0</b> indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself.
	SmartParentId *string `json:"smart_parent_id,omitempty"`
}

// NewDataInnerDhcpSharednetwork6Data instantiates a new DataInnerDhcpSharednetwork6Data object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataInnerDhcpSharednetwork6Data() *DataInnerDhcpSharednetwork6Data {
	this := DataInnerDhcpSharednetwork6Data{}
	return &this
}

// NewDataInnerDhcpSharednetwork6DataWithDefaults instantiates a new DataInnerDhcpSharednetwork6Data object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataInnerDhcpSharednetwork6DataWithDefaults() *DataInnerDhcpSharednetwork6Data {
	this := DataInnerDhcpSharednetwork6Data{}
	return &this
}

// GetSharednetwork6DelayedTime returns the Sharednetwork6DelayedTime field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetSharednetwork6DelayedTime() string {
	if o == nil || IsNil(o.Sharednetwork6DelayedTime) {
		var ret string
		return ret
	}
	return *o.Sharednetwork6DelayedTime
}

// GetSharednetwork6DelayedTimeOk returns a tuple with the Sharednetwork6DelayedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetSharednetwork6DelayedTimeOk() (*string, bool) {
	if o == nil || IsNil(o.Sharednetwork6DelayedTime) {
		return nil, false
	}
	return o.Sharednetwork6DelayedTime, true
}

// HasSharednetwork6DelayedTime returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasSharednetwork6DelayedTime() bool {
	if o != nil && !IsNil(o.Sharednetwork6DelayedTime) {
		return true
	}

	return false
}

// SetSharednetwork6DelayedTime gets a reference to the given string and assigns it to the Sharednetwork6DelayedTime field.
func (o *DataInnerDhcpSharednetwork6Data) SetSharednetwork6DelayedTime(v string) {
	o.Sharednetwork6DelayedTime = &v
}

// GetServer6Id returns the Server6Id field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetServer6Id() string {
	if o == nil || IsNil(o.Server6Id) {
		var ret string
		return ret
	}
	return *o.Server6Id
}

// GetServer6IdOk returns a tuple with the Server6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetServer6IdOk() (*string, bool) {
	if o == nil || IsNil(o.Server6Id) {
		return nil, false
	}
	return o.Server6Id, true
}

// HasServer6Id returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasServer6Id() bool {
	if o != nil && !IsNil(o.Server6Id) {
		return true
	}

	return false
}

// SetServer6Id gets a reference to the given string and assigns it to the Server6Id field.
func (o *DataInnerDhcpSharednetwork6Data) SetServer6Id(v string) {
	o.Server6Id = &v
}

// GetServer6Name returns the Server6Name field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetServer6Name() string {
	if o == nil || IsNil(o.Server6Name) {
		var ret string
		return ret
	}
	return *o.Server6Name
}

// GetServer6NameOk returns a tuple with the Server6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetServer6NameOk() (*string, bool) {
	if o == nil || IsNil(o.Server6Name) {
		return nil, false
	}
	return o.Server6Name, true
}

// HasServer6Name returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasServer6Name() bool {
	if o != nil && !IsNil(o.Server6Name) {
		return true
	}

	return false
}

// SetServer6Name gets a reference to the given string and assigns it to the Server6Name field.
func (o *DataInnerDhcpSharednetwork6Data) SetServer6Name(v string) {
	o.Server6Name = &v
}

// GetServer6Type returns the Server6Type field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetServer6Type() string {
	if o == nil || IsNil(o.Server6Type) {
		var ret string
		return ret
	}
	return *o.Server6Type
}

// GetServer6TypeOk returns a tuple with the Server6Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetServer6TypeOk() (*string, bool) {
	if o == nil || IsNil(o.Server6Type) {
		return nil, false
	}
	return o.Server6Type, true
}

// HasServer6Type returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasServer6Type() bool {
	if o != nil && !IsNil(o.Server6Type) {
		return true
	}

	return false
}

// SetServer6Type gets a reference to the given string and assigns it to the Server6Type field.
func (o *DataInnerDhcpSharednetwork6Data) SetServer6Type(v string) {
	o.Server6Type = &v
}

// GetSharednetwork6Id returns the Sharednetwork6Id field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetSharednetwork6Id() string {
	if o == nil || IsNil(o.Sharednetwork6Id) {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Id
}

// GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetSharednetwork6IdOk() (*string, bool) {
	if o == nil || IsNil(o.Sharednetwork6Id) {
		return nil, false
	}
	return o.Sharednetwork6Id, true
}

// HasSharednetwork6Id returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasSharednetwork6Id() bool {
	if o != nil && !IsNil(o.Sharednetwork6Id) {
		return true
	}

	return false
}

// SetSharednetwork6Id gets a reference to the given string and assigns it to the Sharednetwork6Id field.
func (o *DataInnerDhcpSharednetwork6Data) SetSharednetwork6Id(v string) {
	o.Sharednetwork6Id = &v
}

// GetSharednetwork6Name returns the Sharednetwork6Name field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetSharednetwork6Name() string {
	if o == nil || IsNil(o.Sharednetwork6Name) {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Name
}

// GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetSharednetwork6NameOk() (*string, bool) {
	if o == nil || IsNil(o.Sharednetwork6Name) {
		return nil, false
	}
	return o.Sharednetwork6Name, true
}

// HasSharednetwork6Name returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasSharednetwork6Name() bool {
	if o != nil && !IsNil(o.Sharednetwork6Name) {
		return true
	}

	return false
}

// SetSharednetwork6Name gets a reference to the given string and assigns it to the Sharednetwork6Name field.
func (o *DataInnerDhcpSharednetwork6Data) SetSharednetwork6Name(v string) {
	o.Sharednetwork6Name = &v
}

// GetSharednetwork6Multistatus returns the Sharednetwork6Multistatus field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetSharednetwork6Multistatus() string {
	if o == nil || IsNil(o.Sharednetwork6Multistatus) {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Multistatus
}

// GetSharednetwork6MultistatusOk returns a tuple with the Sharednetwork6Multistatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetSharednetwork6MultistatusOk() (*string, bool) {
	if o == nil || IsNil(o.Sharednetwork6Multistatus) {
		return nil, false
	}
	return o.Sharednetwork6Multistatus, true
}

// HasSharednetwork6Multistatus returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasSharednetwork6Multistatus() bool {
	if o != nil && !IsNil(o.Sharednetwork6Multistatus) {
		return true
	}

	return false
}

// SetSharednetwork6Multistatus gets a reference to the given string and assigns it to the Sharednetwork6Multistatus field.
func (o *DataInnerDhcpSharednetwork6Data) SetSharednetwork6Multistatus(v string) {
	o.Sharednetwork6Multistatus = &v
}

// GetSmartArch returns the SmartArch field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetSmartArch() string {
	if o == nil || IsNil(o.SmartArch) {
		var ret string
		return ret
	}
	return *o.SmartArch
}

// GetSmartArchOk returns a tuple with the SmartArch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetSmartArchOk() (*string, bool) {
	if o == nil || IsNil(o.SmartArch) {
		return nil, false
	}
	return o.SmartArch, true
}

// HasSmartArch returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasSmartArch() bool {
	if o != nil && !IsNil(o.SmartArch) {
		return true
	}

	return false
}

// SetSmartArch gets a reference to the given string and assigns it to the SmartArch field.
func (o *DataInnerDhcpSharednetwork6Data) SetSmartArch(v string) {
	o.SmartArch = &v
}

// GetSmartParentId returns the SmartParentId field value if set, zero value otherwise.
func (o *DataInnerDhcpSharednetwork6Data) GetSmartParentId() string {
	if o == nil || IsNil(o.SmartParentId) {
		var ret string
		return ret
	}
	return *o.SmartParentId
}

// GetSmartParentIdOk returns a tuple with the SmartParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataInnerDhcpSharednetwork6Data) GetSmartParentIdOk() (*string, bool) {
	if o == nil || IsNil(o.SmartParentId) {
		return nil, false
	}
	return o.SmartParentId, true
}

// HasSmartParentId returns a boolean if a field has been set.
func (o *DataInnerDhcpSharednetwork6Data) HasSmartParentId() bool {
	if o != nil && !IsNil(o.SmartParentId) {
		return true
	}

	return false
}

// SetSmartParentId gets a reference to the given string and assigns it to the SmartParentId field.
func (o *DataInnerDhcpSharednetwork6Data) SetSmartParentId(v string) {
	o.SmartParentId = &v
}

func (o DataInnerDhcpSharednetwork6Data) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataInnerDhcpSharednetwork6Data) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Sharednetwork6DelayedTime) {
		toSerialize["sharednetwork6_delayed_time"] = o.Sharednetwork6DelayedTime
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
	if !IsNil(o.Sharednetwork6Id) {
		toSerialize["sharednetwork6_id"] = o.Sharednetwork6Id
	}
	if !IsNil(o.Sharednetwork6Name) {
		toSerialize["sharednetwork6_name"] = o.Sharednetwork6Name
	}
	if !IsNil(o.Sharednetwork6Multistatus) {
		toSerialize["sharednetwork6_multistatus"] = o.Sharednetwork6Multistatus
	}
	if !IsNil(o.SmartArch) {
		toSerialize["smart_arch"] = o.SmartArch
	}
	if !IsNil(o.SmartParentId) {
		toSerialize["smart_parent_id"] = o.SmartParentId
	}
	return toSerialize, nil
}

type NullableDataInnerDhcpSharednetwork6Data struct {
	value *DataInnerDhcpSharednetwork6Data
	isSet bool
}

func (v NullableDataInnerDhcpSharednetwork6Data) Get() *DataInnerDhcpSharednetwork6Data {
	return v.value
}

func (v *NullableDataInnerDhcpSharednetwork6Data) Set(val *DataInnerDhcpSharednetwork6Data) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInnerDhcpSharednetwork6Data) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInnerDhcpSharednetwork6Data) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInnerDhcpSharednetwork6Data(val *DataInnerDhcpSharednetwork6Data) *NullableDataInnerDhcpSharednetwork6Data {
	return &NullableDataInnerDhcpSharednetwork6Data{value: val, isSet: true}
}

func (v NullableDataInnerDhcpSharednetwork6Data) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInnerDhcpSharednetwork6Data) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
