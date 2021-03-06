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

// DhcpSharednetwork6DataData struct for DhcpSharednetwork6DataData
type DhcpSharednetwork6DataData struct {
	// The delay of creation/deletion status. <b>1</b> indicates that the object is not created/deleted yet.
	Sharednetwork6DelayedTime *string `json:"sharednetwork6_delayed_time,omitempty"`
	// The database identifier (ID) of the DHCPv6 server the object belongs to.
	Server6Id *string `json:"server6_id,omitempty"`
	// The name of the DHCPv6 server the object belongs to.
	Server6Name *string `json:"server6_name,omitempty"`
	// The type of the DHCPv6 server the object belongs to: <table><caption>dhcp6_type possible values</caption><br/><thead><tr><th>Type</th><th>Description</th></tr><br/></thead><br/><tbody><tr><td >ipm</td><td >EfficientIP DHCP server or EfficientIP DHCP Package</td></tr><tr><td >vdhcp</td><td >EfficientIP DHCP smart architecture</td></tr></tbody></table></p><br/>
	Server6Type *string `json:"server6_type,omitempty"`
	// The database identifier (ID) of the DHCPv6 shared network.
	Sharednetwork6Id *string `json:"sharednetwork6_id,omitempty"`
	// The name of the DHCPv6 shared network.
	Sharednetwork6Name *string `json:"sharednetwork6_name,omitempty"`
	// The Multi-status information is displayed as follows: <i>&lt;number-of-instances&gt;@&lt;message-number&gt;@&lt;multi-status-severity&gt;@&lt;module&gt;</i>. The different severity levels are:<br><b>Multi-status severity levels</b>    <table border=1>        <thead>        <tr >            <td><b>Message number</b></td>            <td><b>Severity</b></td>            <td><b>Description</b></td>        </tr>        </thead>        <tbody>        <tr  valign=middle>            <td>0 - 16</td>            <td>Emergency</td>            <td>The object configuration prevents the system from running properly. Action is required.</td>        </tr>        <tr  valign=middle>            <td>17 - 33</td>            <td>Critical</td>            <td>The object configuration is in critical conditions. Immediate action is recommended.</td>        </tr>        <tr  valign=middle>            <td>34 - 50</td>            <td>Error</td>            <td>The object configuration failed at some level. Action is recommended.</td>        </tr>        <tr  valign=middle>            <td>51 - 66</td>            <td>Warning</td>            <td>The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.</td>        </tr>        <tr  valign=middle>            <td>67 - 83</td>            <td>Notice</td>            <td>The object configuration is normal but undergoing events that might trigger errors. No immediate action required.</td>        </tr>        <tr  valign=middle>            <td>84 - 100</td>            <td>Informational</td>            <td>The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.</td>        </tr>        </tbody></table>
	Sharednetwork6Multistatus *string `json:"sharednetwork6_multistatus,omitempty"`
	// The object activation status: <ul class=dashed ><li>                                                If set to <b>0</b>, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.<br/>                                            </li><li>                                                If set to <b>1</b>, the object is enabled and managed.<br/>                                            </li></ul>By default, <b>row_enabled</b> is set to <b>1</b> when an object is created.
	RowState *string `json:"row_state,omitempty"`
	// The type of the DHCPv6 smart architecture the object belongs to.<table><caption>vdhcp6_arch possible values</caption><br/><thead><tr><th>Type</th><th>Description</th></tr><br/></thead><br/><tbody><tr><td >single</td><td >The Single-Server smart architecture manages a single DHCPv6 server.</td></tr><tr><td >splitscope</td><td >The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.</td></tr><tr><td >stateless</td><td >The Stateless smart architecture offers a limited number of options to the DHCP clients. The IP address is delivered thanks to the subnet gateway and it is impossible to create any ranges or statics or to retrieve any leases.</td></tr></tbody></table></p><br/>
	SmartArch *string `json:"smart_arch,omitempty"`
	// The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. <b>0</b> indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself.
	SmartParentId *string `json:"smart_parent_id,omitempty"`
}

// NewDhcpSharednetwork6DataData instantiates a new DhcpSharednetwork6DataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSharednetwork6DataData() *DhcpSharednetwork6DataData {
	this := DhcpSharednetwork6DataData{}
	return &this
}

// NewDhcpSharednetwork6DataDataWithDefaults instantiates a new DhcpSharednetwork6DataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSharednetwork6DataDataWithDefaults() *DhcpSharednetwork6DataData {
	this := DhcpSharednetwork6DataData{}
	return &this
}

// GetSharednetwork6DelayedTime returns the Sharednetwork6DelayedTime field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetSharednetwork6DelayedTime() string {
	if o == nil || o.Sharednetwork6DelayedTime == nil {
		var ret string
		return ret
	}
	return *o.Sharednetwork6DelayedTime
}

// GetSharednetwork6DelayedTimeOk returns a tuple with the Sharednetwork6DelayedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetSharednetwork6DelayedTimeOk() (*string, bool) {
	if o == nil || o.Sharednetwork6DelayedTime == nil {
		return nil, false
	}
	return o.Sharednetwork6DelayedTime, true
}

// HasSharednetwork6DelayedTime returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasSharednetwork6DelayedTime() bool {
	if o != nil && o.Sharednetwork6DelayedTime != nil {
		return true
	}

	return false
}

// SetSharednetwork6DelayedTime gets a reference to the given string and assigns it to the Sharednetwork6DelayedTime field.
func (o *DhcpSharednetwork6DataData) SetSharednetwork6DelayedTime(v string) {
	o.Sharednetwork6DelayedTime = &v
}

// GetServer6Id returns the Server6Id field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetServer6Id() string {
	if o == nil || o.Server6Id == nil {
		var ret string
		return ret
	}
	return *o.Server6Id
}

// GetServer6IdOk returns a tuple with the Server6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetServer6IdOk() (*string, bool) {
	if o == nil || o.Server6Id == nil {
		return nil, false
	}
	return o.Server6Id, true
}

// HasServer6Id returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasServer6Id() bool {
	if o != nil && o.Server6Id != nil {
		return true
	}

	return false
}

// SetServer6Id gets a reference to the given string and assigns it to the Server6Id field.
func (o *DhcpSharednetwork6DataData) SetServer6Id(v string) {
	o.Server6Id = &v
}

// GetServer6Name returns the Server6Name field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetServer6Name() string {
	if o == nil || o.Server6Name == nil {
		var ret string
		return ret
	}
	return *o.Server6Name
}

// GetServer6NameOk returns a tuple with the Server6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetServer6NameOk() (*string, bool) {
	if o == nil || o.Server6Name == nil {
		return nil, false
	}
	return o.Server6Name, true
}

// HasServer6Name returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasServer6Name() bool {
	if o != nil && o.Server6Name != nil {
		return true
	}

	return false
}

// SetServer6Name gets a reference to the given string and assigns it to the Server6Name field.
func (o *DhcpSharednetwork6DataData) SetServer6Name(v string) {
	o.Server6Name = &v
}

// GetServer6Type returns the Server6Type field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetServer6Type() string {
	if o == nil || o.Server6Type == nil {
		var ret string
		return ret
	}
	return *o.Server6Type
}

// GetServer6TypeOk returns a tuple with the Server6Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetServer6TypeOk() (*string, bool) {
	if o == nil || o.Server6Type == nil {
		return nil, false
	}
	return o.Server6Type, true
}

// HasServer6Type returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasServer6Type() bool {
	if o != nil && o.Server6Type != nil {
		return true
	}

	return false
}

// SetServer6Type gets a reference to the given string and assigns it to the Server6Type field.
func (o *DhcpSharednetwork6DataData) SetServer6Type(v string) {
	o.Server6Type = &v
}

// GetSharednetwork6Id returns the Sharednetwork6Id field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetSharednetwork6Id() string {
	if o == nil || o.Sharednetwork6Id == nil {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Id
}

// GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetSharednetwork6IdOk() (*string, bool) {
	if o == nil || o.Sharednetwork6Id == nil {
		return nil, false
	}
	return o.Sharednetwork6Id, true
}

// HasSharednetwork6Id returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasSharednetwork6Id() bool {
	if o != nil && o.Sharednetwork6Id != nil {
		return true
	}

	return false
}

// SetSharednetwork6Id gets a reference to the given string and assigns it to the Sharednetwork6Id field.
func (o *DhcpSharednetwork6DataData) SetSharednetwork6Id(v string) {
	o.Sharednetwork6Id = &v
}

// GetSharednetwork6Name returns the Sharednetwork6Name field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetSharednetwork6Name() string {
	if o == nil || o.Sharednetwork6Name == nil {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Name
}

// GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetSharednetwork6NameOk() (*string, bool) {
	if o == nil || o.Sharednetwork6Name == nil {
		return nil, false
	}
	return o.Sharednetwork6Name, true
}

// HasSharednetwork6Name returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasSharednetwork6Name() bool {
	if o != nil && o.Sharednetwork6Name != nil {
		return true
	}

	return false
}

// SetSharednetwork6Name gets a reference to the given string and assigns it to the Sharednetwork6Name field.
func (o *DhcpSharednetwork6DataData) SetSharednetwork6Name(v string) {
	o.Sharednetwork6Name = &v
}

// GetSharednetwork6Multistatus returns the Sharednetwork6Multistatus field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetSharednetwork6Multistatus() string {
	if o == nil || o.Sharednetwork6Multistatus == nil {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Multistatus
}

// GetSharednetwork6MultistatusOk returns a tuple with the Sharednetwork6Multistatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetSharednetwork6MultistatusOk() (*string, bool) {
	if o == nil || o.Sharednetwork6Multistatus == nil {
		return nil, false
	}
	return o.Sharednetwork6Multistatus, true
}

// HasSharednetwork6Multistatus returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasSharednetwork6Multistatus() bool {
	if o != nil && o.Sharednetwork6Multistatus != nil {
		return true
	}

	return false
}

// SetSharednetwork6Multistatus gets a reference to the given string and assigns it to the Sharednetwork6Multistatus field.
func (o *DhcpSharednetwork6DataData) SetSharednetwork6Multistatus(v string) {
	o.Sharednetwork6Multistatus = &v
}

// GetRowState returns the RowState field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetRowState() string {
	if o == nil || o.RowState == nil {
		var ret string
		return ret
	}
	return *o.RowState
}

// GetRowStateOk returns a tuple with the RowState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetRowStateOk() (*string, bool) {
	if o == nil || o.RowState == nil {
		return nil, false
	}
	return o.RowState, true
}

// HasRowState returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasRowState() bool {
	if o != nil && o.RowState != nil {
		return true
	}

	return false
}

// SetRowState gets a reference to the given string and assigns it to the RowState field.
func (o *DhcpSharednetwork6DataData) SetRowState(v string) {
	o.RowState = &v
}

// GetSmartArch returns the SmartArch field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetSmartArch() string {
	if o == nil || o.SmartArch == nil {
		var ret string
		return ret
	}
	return *o.SmartArch
}

// GetSmartArchOk returns a tuple with the SmartArch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetSmartArchOk() (*string, bool) {
	if o == nil || o.SmartArch == nil {
		return nil, false
	}
	return o.SmartArch, true
}

// HasSmartArch returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasSmartArch() bool {
	if o != nil && o.SmartArch != nil {
		return true
	}

	return false
}

// SetSmartArch gets a reference to the given string and assigns it to the SmartArch field.
func (o *DhcpSharednetwork6DataData) SetSmartArch(v string) {
	o.SmartArch = &v
}

// GetSmartParentId returns the SmartParentId field value if set, zero value otherwise.
func (o *DhcpSharednetwork6DataData) GetSmartParentId() string {
	if o == nil || o.SmartParentId == nil {
		var ret string
		return ret
	}
	return *o.SmartParentId
}

// GetSmartParentIdOk returns a tuple with the SmartParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetwork6DataData) GetSmartParentIdOk() (*string, bool) {
	if o == nil || o.SmartParentId == nil {
		return nil, false
	}
	return o.SmartParentId, true
}

// HasSmartParentId returns a boolean if a field has been set.
func (o *DhcpSharednetwork6DataData) HasSmartParentId() bool {
	if o != nil && o.SmartParentId != nil {
		return true
	}

	return false
}

// SetSmartParentId gets a reference to the given string and assigns it to the SmartParentId field.
func (o *DhcpSharednetwork6DataData) SetSmartParentId(v string) {
	o.SmartParentId = &v
}

func (o DhcpSharednetwork6DataData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Sharednetwork6DelayedTime != nil {
		toSerialize["sharednetwork6_delayed_time"] = o.Sharednetwork6DelayedTime
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
	if o.Sharednetwork6Id != nil {
		toSerialize["sharednetwork6_id"] = o.Sharednetwork6Id
	}
	if o.Sharednetwork6Name != nil {
		toSerialize["sharednetwork6_name"] = o.Sharednetwork6Name
	}
	if o.Sharednetwork6Multistatus != nil {
		toSerialize["sharednetwork6_multistatus"] = o.Sharednetwork6Multistatus
	}
	if o.RowState != nil {
		toSerialize["row_state"] = o.RowState
	}
	if o.SmartArch != nil {
		toSerialize["smart_arch"] = o.SmartArch
	}
	if o.SmartParentId != nil {
		toSerialize["smart_parent_id"] = o.SmartParentId
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpSharednetwork6DataData struct {
	value *DhcpSharednetwork6DataData
	isSet bool
}

func (v NullableDhcpSharednetwork6DataData) Get() *DhcpSharednetwork6DataData {
	return v.value
}

func (v *NullableDhcpSharednetwork6DataData) Set(val *DhcpSharednetwork6DataData) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSharednetwork6DataData) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSharednetwork6DataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSharednetwork6DataData(val *DhcpSharednetwork6DataData) *NullableDhcpSharednetwork6DataData {
	return &NullableDhcpSharednetwork6DataData{value: val, isSet: true}
}

func (v NullableDhcpSharednetwork6DataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSharednetwork6DataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


