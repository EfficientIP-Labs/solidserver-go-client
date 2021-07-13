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

// DhcpSharednetworkDataData struct for DhcpSharednetworkDataData
type DhcpSharednetworkDataData struct {
	// The delay of creation status. <b>1</b> indicates that the object is not created yet.
	SharednetworkDelayedCreateTime *string `json:"sharednetwork_delayed_create_time,omitempty"`
	// The delay of deletion status. <b>1</b> indicates that the object is not deleted yet.
	SharednetworkDelayedDeleteTime *string `json:"sharednetwork_delayed_delete_time,omitempty"`
	// The database identifier (ID) of the DHCPv4 server the object belongs to.
	ServerId *string `json:"server_id,omitempty"`
	// The name of the DHCPv4 server the object belongs to.
	ServerName *string `json:"server_name,omitempty"`
	// The type of the DHCPv4 server the object belongs to: <table><caption>dhcp_type possible values</caption><br/><thead><tr><th>Type</th><th>Description</th></tr><br/></thead><br/><tbody><tr><td >ipm</td><td >EfficientIP DHCP server or EfficientIP DHCP Package</td></tr><tr><td >msrpc</td><td >Microsoft DHCP server</td></tr><tr><td >dcs</td><td >Nominum DCS server</td></tr><tr><td >vdhcp</td><td >EfficientIP DHCP smart architecture</td></tr></tbody></table></p><br/>
	ServerType *string `json:"server_type,omitempty"`
	// The database identifier (ID) of the DHCPv4 shared network.
	SharednetworkId *string `json:"sharednetwork_id,omitempty"`
	// The name of the DHCPv4 shared network.
	SharednetworkName *string `json:"sharednetwork_name,omitempty"`
	// The Multi-status information is displayed as follows: <i>&lt;number-of-instances&gt;@&lt;message-number&gt;@&lt;multi-status-severity&gt;@&lt;module&gt;</i>. The different severity levels are:<br><b>Multi-status severity levels</b>    <table border=1>        <thead>        <tr >            <td><b>Message number</b></td>            <td><b>Severity</b></td>            <td><b>Description</b></td>        </tr>        </thead>        <tbody>        <tr  valign=middle>            <td>0 - 16</td>            <td>Emergency</td>            <td>The object configuration prevents the system from running properly. Action is required.</td>        </tr>        <tr  valign=middle>            <td>17 - 33</td>            <td>Critical</td>            <td>The object configuration is in critical conditions. Immediate action is recommended.</td>        </tr>        <tr  valign=middle>            <td>34 - 50</td>            <td>Error</td>            <td>The object configuration failed at some level. Action is recommended.</td>        </tr>        <tr  valign=middle>            <td>51 - 66</td>            <td>Warning</td>            <td>The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.</td>        </tr>        <tr  valign=middle>            <td>67 - 83</td>            <td>Notice</td>            <td>The object configuration is normal but undergoing events that might trigger errors. No immediate action required.</td>        </tr>        <tr  valign=middle>            <td>84 - 100</td>            <td>Informational</td>            <td>The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.</td>        </tr>        </tbody></table>
	SharednetworkMultistatus *string `json:"sharednetwork_multistatus,omitempty"`
	// The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. <b>0</b> indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself.
	SmartParentId *string `json:"smart_parent_id,omitempty"`
}

// NewDhcpSharednetworkDataData instantiates a new DhcpSharednetworkDataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpSharednetworkDataData() *DhcpSharednetworkDataData {
	this := DhcpSharednetworkDataData{}
	return &this
}

// NewDhcpSharednetworkDataDataWithDefaults instantiates a new DhcpSharednetworkDataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpSharednetworkDataDataWithDefaults() *DhcpSharednetworkDataData {
	this := DhcpSharednetworkDataData{}
	return &this
}

// GetSharednetworkDelayedCreateTime returns the SharednetworkDelayedCreateTime field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetSharednetworkDelayedCreateTime() string {
	if o == nil || o.SharednetworkDelayedCreateTime == nil {
		var ret string
		return ret
	}
	return *o.SharednetworkDelayedCreateTime
}

// GetSharednetworkDelayedCreateTimeOk returns a tuple with the SharednetworkDelayedCreateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetSharednetworkDelayedCreateTimeOk() (*string, bool) {
	if o == nil || o.SharednetworkDelayedCreateTime == nil {
		return nil, false
	}
	return o.SharednetworkDelayedCreateTime, true
}

// HasSharednetworkDelayedCreateTime returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasSharednetworkDelayedCreateTime() bool {
	if o != nil && o.SharednetworkDelayedCreateTime != nil {
		return true
	}

	return false
}

// SetSharednetworkDelayedCreateTime gets a reference to the given string and assigns it to the SharednetworkDelayedCreateTime field.
func (o *DhcpSharednetworkDataData) SetSharednetworkDelayedCreateTime(v string) {
	o.SharednetworkDelayedCreateTime = &v
}

// GetSharednetworkDelayedDeleteTime returns the SharednetworkDelayedDeleteTime field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetSharednetworkDelayedDeleteTime() string {
	if o == nil || o.SharednetworkDelayedDeleteTime == nil {
		var ret string
		return ret
	}
	return *o.SharednetworkDelayedDeleteTime
}

// GetSharednetworkDelayedDeleteTimeOk returns a tuple with the SharednetworkDelayedDeleteTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetSharednetworkDelayedDeleteTimeOk() (*string, bool) {
	if o == nil || o.SharednetworkDelayedDeleteTime == nil {
		return nil, false
	}
	return o.SharednetworkDelayedDeleteTime, true
}

// HasSharednetworkDelayedDeleteTime returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasSharednetworkDelayedDeleteTime() bool {
	if o != nil && o.SharednetworkDelayedDeleteTime != nil {
		return true
	}

	return false
}

// SetSharednetworkDelayedDeleteTime gets a reference to the given string and assigns it to the SharednetworkDelayedDeleteTime field.
func (o *DhcpSharednetworkDataData) SetSharednetworkDelayedDeleteTime(v string) {
	o.SharednetworkDelayedDeleteTime = &v
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetServerId() string {
	if o == nil || o.ServerId == nil {
		var ret string
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetServerIdOk() (*string, bool) {
	if o == nil || o.ServerId == nil {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasServerId() bool {
	if o != nil && o.ServerId != nil {
		return true
	}

	return false
}

// SetServerId gets a reference to the given string and assigns it to the ServerId field.
func (o *DhcpSharednetworkDataData) SetServerId(v string) {
	o.ServerId = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetServerName() string {
	if o == nil || o.ServerName == nil {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetServerNameOk() (*string, bool) {
	if o == nil || o.ServerName == nil {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasServerName() bool {
	if o != nil && o.ServerName != nil {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *DhcpSharednetworkDataData) SetServerName(v string) {
	o.ServerName = &v
}

// GetServerType returns the ServerType field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetServerType() string {
	if o == nil || o.ServerType == nil {
		var ret string
		return ret
	}
	return *o.ServerType
}

// GetServerTypeOk returns a tuple with the ServerType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetServerTypeOk() (*string, bool) {
	if o == nil || o.ServerType == nil {
		return nil, false
	}
	return o.ServerType, true
}

// HasServerType returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasServerType() bool {
	if o != nil && o.ServerType != nil {
		return true
	}

	return false
}

// SetServerType gets a reference to the given string and assigns it to the ServerType field.
func (o *DhcpSharednetworkDataData) SetServerType(v string) {
	o.ServerType = &v
}

// GetSharednetworkId returns the SharednetworkId field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetSharednetworkId() string {
	if o == nil || o.SharednetworkId == nil {
		var ret string
		return ret
	}
	return *o.SharednetworkId
}

// GetSharednetworkIdOk returns a tuple with the SharednetworkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetSharednetworkIdOk() (*string, bool) {
	if o == nil || o.SharednetworkId == nil {
		return nil, false
	}
	return o.SharednetworkId, true
}

// HasSharednetworkId returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasSharednetworkId() bool {
	if o != nil && o.SharednetworkId != nil {
		return true
	}

	return false
}

// SetSharednetworkId gets a reference to the given string and assigns it to the SharednetworkId field.
func (o *DhcpSharednetworkDataData) SetSharednetworkId(v string) {
	o.SharednetworkId = &v
}

// GetSharednetworkName returns the SharednetworkName field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetSharednetworkName() string {
	if o == nil || o.SharednetworkName == nil {
		var ret string
		return ret
	}
	return *o.SharednetworkName
}

// GetSharednetworkNameOk returns a tuple with the SharednetworkName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetSharednetworkNameOk() (*string, bool) {
	if o == nil || o.SharednetworkName == nil {
		return nil, false
	}
	return o.SharednetworkName, true
}

// HasSharednetworkName returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasSharednetworkName() bool {
	if o != nil && o.SharednetworkName != nil {
		return true
	}

	return false
}

// SetSharednetworkName gets a reference to the given string and assigns it to the SharednetworkName field.
func (o *DhcpSharednetworkDataData) SetSharednetworkName(v string) {
	o.SharednetworkName = &v
}

// GetSharednetworkMultistatus returns the SharednetworkMultistatus field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetSharednetworkMultistatus() string {
	if o == nil || o.SharednetworkMultistatus == nil {
		var ret string
		return ret
	}
	return *o.SharednetworkMultistatus
}

// GetSharednetworkMultistatusOk returns a tuple with the SharednetworkMultistatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetSharednetworkMultistatusOk() (*string, bool) {
	if o == nil || o.SharednetworkMultistatus == nil {
		return nil, false
	}
	return o.SharednetworkMultistatus, true
}

// HasSharednetworkMultistatus returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasSharednetworkMultistatus() bool {
	if o != nil && o.SharednetworkMultistatus != nil {
		return true
	}

	return false
}

// SetSharednetworkMultistatus gets a reference to the given string and assigns it to the SharednetworkMultistatus field.
func (o *DhcpSharednetworkDataData) SetSharednetworkMultistatus(v string) {
	o.SharednetworkMultistatus = &v
}

// GetSmartParentId returns the SmartParentId field value if set, zero value otherwise.
func (o *DhcpSharednetworkDataData) GetSmartParentId() string {
	if o == nil || o.SmartParentId == nil {
		var ret string
		return ret
	}
	return *o.SmartParentId
}

// GetSmartParentIdOk returns a tuple with the SmartParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpSharednetworkDataData) GetSmartParentIdOk() (*string, bool) {
	if o == nil || o.SmartParentId == nil {
		return nil, false
	}
	return o.SmartParentId, true
}

// HasSmartParentId returns a boolean if a field has been set.
func (o *DhcpSharednetworkDataData) HasSmartParentId() bool {
	if o != nil && o.SmartParentId != nil {
		return true
	}

	return false
}

// SetSmartParentId gets a reference to the given string and assigns it to the SmartParentId field.
func (o *DhcpSharednetworkDataData) SetSmartParentId(v string) {
	o.SmartParentId = &v
}

func (o DhcpSharednetworkDataData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SharednetworkDelayedCreateTime != nil {
		toSerialize["sharednetwork_delayed_create_time"] = o.SharednetworkDelayedCreateTime
	}
	if o.SharednetworkDelayedDeleteTime != nil {
		toSerialize["sharednetwork_delayed_delete_time"] = o.SharednetworkDelayedDeleteTime
	}
	if o.ServerId != nil {
		toSerialize["server_id"] = o.ServerId
	}
	if o.ServerName != nil {
		toSerialize["server_name"] = o.ServerName
	}
	if o.ServerType != nil {
		toSerialize["server_type"] = o.ServerType
	}
	if o.SharednetworkId != nil {
		toSerialize["sharednetwork_id"] = o.SharednetworkId
	}
	if o.SharednetworkName != nil {
		toSerialize["sharednetwork_name"] = o.SharednetworkName
	}
	if o.SharednetworkMultistatus != nil {
		toSerialize["sharednetwork_multistatus"] = o.SharednetworkMultistatus
	}
	if o.SmartParentId != nil {
		toSerialize["smart_parent_id"] = o.SmartParentId
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpSharednetworkDataData struct {
	value *DhcpSharednetworkDataData
	isSet bool
}

func (v NullableDhcpSharednetworkDataData) Get() *DhcpSharednetworkDataData {
	return v.value
}

func (v *NullableDhcpSharednetworkDataData) Set(val *DhcpSharednetworkDataData) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpSharednetworkDataData) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpSharednetworkDataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpSharednetworkDataData(val *DhcpSharednetworkDataData) *NullableDhcpSharednetworkDataData {
	return &NullableDhcpSharednetworkDataData{value: val, isSet: true}
}

func (v NullableDhcpSharednetworkDataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpSharednetworkDataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

