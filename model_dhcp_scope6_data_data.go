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

// DhcpScope6DataData struct for DhcpScope6DataData
type DhcpScope6DataData struct {
	// The delay of creation/deletion status. <b>1</b> indicates that the object is not created/deleted yet.
	Scope6DelayedTime *string `json:"scope6_delayed_time,omitempty"`
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
	// The database identifier (ID) of the DHCPv6 failover channel associated with the object.
	Failover6Id *string `json:"failover6_id,omitempty"`
	// The name of the DHCPv6 failover channel associated with the object.
	Failover6Name *string `json:"failover6_name,omitempty"`
	// The name of the class applied to the DHCPv6 scope, it can be preceded by the class directory.
	Scope6ClassName *string `json:"scope6_class_name,omitempty"`
	// The class parameters applied to the DHCPv6 scope and their value: <b>&lt;class-parameter1&gt;=&lt;value1&gt;&amp;&lt;class-parameter2&gt;=&lt;value2&gt;&amp;</b>... .
	Scope6ClassParameters *[]ApiClassParameterOutputEntry `json:"scope6_class_parameters,omitempty"`
	// The last IP address of the DHCPv6 scope, in hexadecimal format.
	Scope6EndAddress6Addr *string `json:"scope6_end_address6_addr,omitempty"`
	// The database identifier (ID) of the DHCPv6 scope.
	Scope6Id *string `json:"scope6_id,omitempty"`
	// The name of the DHCPv6 scope.
	Scope6Name *string `json:"scope6_name,omitempty"`
	// The prefix of the DHCPv6 scope.
	Scope6Prefix *string `json:"scope6_prefix,omitempty"`
	// The database identifier (ID) of the space associated with the DHCPv6 scope.
	Scope6SpaceId *string `json:"scope6_space_id,omitempty"`
	// The name of the space associated with the DHCPv6 scope.
	Scope6SpaceName *string `json:"scope6_space_name,omitempty"`
	// The number of IP addresses the DHCPv6 scope contains.
	Scope6Size *string `json:"scope6_size,omitempty"`
	// Internal use. Not documented.
	Scope6SortName *string `json:"scope6_sort_name,omitempty"`
	// The first IP address of the DHCPv6 scope, in hexadecimal format.
	Scope6StartAddress6Addr *string `json:"scope6_start_address6_addr,omitempty"`
	// The database identifier (ID) of the DHCPv6 shared network the object belongs to.
	Sharednetwork6Id *string `json:"sharednetwork6_id,omitempty"`
	// The name of the DHCPv6 shared network the object belongs to.
	Sharednetwork6Name *string `json:"sharednetwork6_name,omitempty"`
	// The IP address of the DHCP server the object belongs to, in hexadecimal format.
	Server6Addr *string `json:"server6_addr,omitempty"`
	// The Multi-status information is displayed as follows: <i>&lt;number-of-instances&gt;@&lt;message-number&gt;@&lt;multi-status-severity&gt;@&lt;module&gt;</i>. The different severity levels are:<br><b>Multi-status severity levels</b>    <table border=1>        <thead>        <tr >            <td><b>Message number</b></td>            <td><b>Severity</b></td>            <td><b>Description</b></td>        </tr>        </thead>        <tbody>        <tr  valign=middle>            <td>0 - 16</td>            <td>Emergency</td>            <td>The object configuration prevents the system from running properly. Action is required.</td>        </tr>        <tr  valign=middle>            <td>17 - 33</td>            <td>Critical</td>            <td>The object configuration is in critical conditions. Immediate action is recommended.</td>        </tr>        <tr  valign=middle>            <td>34 - 50</td>            <td>Error</td>            <td>The object configuration failed at some level. Action is recommended.</td>        </tr>        <tr  valign=middle>            <td>51 - 66</td>            <td>Warning</td>            <td>The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.</td>        </tr>        <tr  valign=middle>            <td>67 - 83</td>            <td>Notice</td>            <td>The object configuration is normal but undergoing events that might trigger errors. No immediate action required.</td>        </tr>        <tr  valign=middle>            <td>84 - 100</td>            <td>Informational</td>            <td>The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.</td>        </tr>        </tbody></table>
	Scope6Multistatus *string `json:"scope6_multistatus,omitempty"`
	// The type of the DHCPv6 smart architecture the object belongs to.<table><caption>vdhcp6_arch possible values</caption><br/><thead><tr><th>Type</th><th>Description</th></tr><br/></thead><br/><tbody><tr><td >single</td><td >The Single-Server smart architecture manages a single DHCPv6 server.</td></tr><tr><td >splitscope</td><td >The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.</td></tr><tr><td >stateless</td><td >The Stateless smart architecture offers a limited number of options to the DHCP clients. The IP address is delivered thanks to the subnet gateway and it is impossible to create any ranges or statics or to retrieve any leases.</td></tr></tbody></table></p><br/>
	SmartArch *string `json:"smart_arch,omitempty"`
	// The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. <b>0</b> indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself.
	SmartParentId *string `json:"smart_parent_id,omitempty"`
	// The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. <b>#</b> indicates that the server is not managed by a smart architecture or is a smart architecture itself.
	SmartParentName *string `json:"smart_parent_name,omitempty"`
}

// NewDhcpScope6DataData instantiates a new DhcpScope6DataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDhcpScope6DataData() *DhcpScope6DataData {
	this := DhcpScope6DataData{}
	return &this
}

// NewDhcpScope6DataDataWithDefaults instantiates a new DhcpScope6DataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDhcpScope6DataDataWithDefaults() *DhcpScope6DataData {
	this := DhcpScope6DataData{}
	return &this
}

// GetScope6DelayedTime returns the Scope6DelayedTime field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6DelayedTime() string {
	if o == nil || o.Scope6DelayedTime == nil {
		var ret string
		return ret
	}
	return *o.Scope6DelayedTime
}

// GetScope6DelayedTimeOk returns a tuple with the Scope6DelayedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6DelayedTimeOk() (*string, bool) {
	if o == nil || o.Scope6DelayedTime == nil {
		return nil, false
	}
	return o.Scope6DelayedTime, true
}

// HasScope6DelayedTime returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6DelayedTime() bool {
	if o != nil && o.Scope6DelayedTime != nil {
		return true
	}

	return false
}

// SetScope6DelayedTime gets a reference to the given string and assigns it to the Scope6DelayedTime field.
func (o *DhcpScope6DataData) SetScope6DelayedTime(v string) {
	o.Scope6DelayedTime = &v
}

// GetServer6ClassName returns the Server6ClassName field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetServer6ClassName() string {
	if o == nil || o.Server6ClassName == nil {
		var ret string
		return ret
	}
	return *o.Server6ClassName
}

// GetServer6ClassNameOk returns a tuple with the Server6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetServer6ClassNameOk() (*string, bool) {
	if o == nil || o.Server6ClassName == nil {
		return nil, false
	}
	return o.Server6ClassName, true
}

// HasServer6ClassName returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasServer6ClassName() bool {
	if o != nil && o.Server6ClassName != nil {
		return true
	}

	return false
}

// SetServer6ClassName gets a reference to the given string and assigns it to the Server6ClassName field.
func (o *DhcpScope6DataData) SetServer6ClassName(v string) {
	o.Server6ClassName = &v
}

// GetServer6ClassParameters returns the Server6ClassParameters field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetServer6ClassParameters() []ApiClassParameterOutputEntry {
	if o == nil || o.Server6ClassParameters == nil {
		var ret []ApiClassParameterOutputEntry
		return ret
	}
	return *o.Server6ClassParameters
}

// GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool) {
	if o == nil || o.Server6ClassParameters == nil {
		return nil, false
	}
	return o.Server6ClassParameters, true
}

// HasServer6ClassParameters returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasServer6ClassParameters() bool {
	if o != nil && o.Server6ClassParameters != nil {
		return true
	}

	return false
}

// SetServer6ClassParameters gets a reference to the given []ApiClassParameterOutputEntry and assigns it to the Server6ClassParameters field.
func (o *DhcpScope6DataData) SetServer6ClassParameters(v []ApiClassParameterOutputEntry) {
	o.Server6ClassParameters = &v
}

// GetServer6Id returns the Server6Id field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetServer6Id() string {
	if o == nil || o.Server6Id == nil {
		var ret string
		return ret
	}
	return *o.Server6Id
}

// GetServer6IdOk returns a tuple with the Server6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetServer6IdOk() (*string, bool) {
	if o == nil || o.Server6Id == nil {
		return nil, false
	}
	return o.Server6Id, true
}

// HasServer6Id returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasServer6Id() bool {
	if o != nil && o.Server6Id != nil {
		return true
	}

	return false
}

// SetServer6Id gets a reference to the given string and assigns it to the Server6Id field.
func (o *DhcpScope6DataData) SetServer6Id(v string) {
	o.Server6Id = &v
}

// GetServer6Name returns the Server6Name field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetServer6Name() string {
	if o == nil || o.Server6Name == nil {
		var ret string
		return ret
	}
	return *o.Server6Name
}

// GetServer6NameOk returns a tuple with the Server6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetServer6NameOk() (*string, bool) {
	if o == nil || o.Server6Name == nil {
		return nil, false
	}
	return o.Server6Name, true
}

// HasServer6Name returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasServer6Name() bool {
	if o != nil && o.Server6Name != nil {
		return true
	}

	return false
}

// SetServer6Name gets a reference to the given string and assigns it to the Server6Name field.
func (o *DhcpScope6DataData) SetServer6Name(v string) {
	o.Server6Name = &v
}

// GetServer6Type returns the Server6Type field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetServer6Type() string {
	if o == nil || o.Server6Type == nil {
		var ret string
		return ret
	}
	return *o.Server6Type
}

// GetServer6TypeOk returns a tuple with the Server6Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetServer6TypeOk() (*string, bool) {
	if o == nil || o.Server6Type == nil {
		return nil, false
	}
	return o.Server6Type, true
}

// HasServer6Type returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasServer6Type() bool {
	if o != nil && o.Server6Type != nil {
		return true
	}

	return false
}

// SetServer6Type gets a reference to the given string and assigns it to the Server6Type field.
func (o *DhcpScope6DataData) SetServer6Type(v string) {
	o.Server6Type = &v
}

// GetServer6Version returns the Server6Version field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetServer6Version() string {
	if o == nil || o.Server6Version == nil {
		var ret string
		return ret
	}
	return *o.Server6Version
}

// GetServer6VersionOk returns a tuple with the Server6Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetServer6VersionOk() (*string, bool) {
	if o == nil || o.Server6Version == nil {
		return nil, false
	}
	return o.Server6Version, true
}

// HasServer6Version returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasServer6Version() bool {
	if o != nil && o.Server6Version != nil {
		return true
	}

	return false
}

// SetServer6Version gets a reference to the given string and assigns it to the Server6Version field.
func (o *DhcpScope6DataData) SetServer6Version(v string) {
	o.Server6Version = &v
}

// GetFailover6Id returns the Failover6Id field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetFailover6Id() string {
	if o == nil || o.Failover6Id == nil {
		var ret string
		return ret
	}
	return *o.Failover6Id
}

// GetFailover6IdOk returns a tuple with the Failover6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetFailover6IdOk() (*string, bool) {
	if o == nil || o.Failover6Id == nil {
		return nil, false
	}
	return o.Failover6Id, true
}

// HasFailover6Id returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasFailover6Id() bool {
	if o != nil && o.Failover6Id != nil {
		return true
	}

	return false
}

// SetFailover6Id gets a reference to the given string and assigns it to the Failover6Id field.
func (o *DhcpScope6DataData) SetFailover6Id(v string) {
	o.Failover6Id = &v
}

// GetFailover6Name returns the Failover6Name field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetFailover6Name() string {
	if o == nil || o.Failover6Name == nil {
		var ret string
		return ret
	}
	return *o.Failover6Name
}

// GetFailover6NameOk returns a tuple with the Failover6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetFailover6NameOk() (*string, bool) {
	if o == nil || o.Failover6Name == nil {
		return nil, false
	}
	return o.Failover6Name, true
}

// HasFailover6Name returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasFailover6Name() bool {
	if o != nil && o.Failover6Name != nil {
		return true
	}

	return false
}

// SetFailover6Name gets a reference to the given string and assigns it to the Failover6Name field.
func (o *DhcpScope6DataData) SetFailover6Name(v string) {
	o.Failover6Name = &v
}

// GetScope6ClassName returns the Scope6ClassName field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6ClassName() string {
	if o == nil || o.Scope6ClassName == nil {
		var ret string
		return ret
	}
	return *o.Scope6ClassName
}

// GetScope6ClassNameOk returns a tuple with the Scope6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6ClassNameOk() (*string, bool) {
	if o == nil || o.Scope6ClassName == nil {
		return nil, false
	}
	return o.Scope6ClassName, true
}

// HasScope6ClassName returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6ClassName() bool {
	if o != nil && o.Scope6ClassName != nil {
		return true
	}

	return false
}

// SetScope6ClassName gets a reference to the given string and assigns it to the Scope6ClassName field.
func (o *DhcpScope6DataData) SetScope6ClassName(v string) {
	o.Scope6ClassName = &v
}

// GetScope6ClassParameters returns the Scope6ClassParameters field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6ClassParameters() []ApiClassParameterOutputEntry {
	if o == nil || o.Scope6ClassParameters == nil {
		var ret []ApiClassParameterOutputEntry
		return ret
	}
	return *o.Scope6ClassParameters
}

// GetScope6ClassParametersOk returns a tuple with the Scope6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool) {
	if o == nil || o.Scope6ClassParameters == nil {
		return nil, false
	}
	return o.Scope6ClassParameters, true
}

// HasScope6ClassParameters returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6ClassParameters() bool {
	if o != nil && o.Scope6ClassParameters != nil {
		return true
	}

	return false
}

// SetScope6ClassParameters gets a reference to the given []ApiClassParameterOutputEntry and assigns it to the Scope6ClassParameters field.
func (o *DhcpScope6DataData) SetScope6ClassParameters(v []ApiClassParameterOutputEntry) {
	o.Scope6ClassParameters = &v
}

// GetScope6EndAddress6Addr returns the Scope6EndAddress6Addr field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6EndAddress6Addr() string {
	if o == nil || o.Scope6EndAddress6Addr == nil {
		var ret string
		return ret
	}
	return *o.Scope6EndAddress6Addr
}

// GetScope6EndAddress6AddrOk returns a tuple with the Scope6EndAddress6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6EndAddress6AddrOk() (*string, bool) {
	if o == nil || o.Scope6EndAddress6Addr == nil {
		return nil, false
	}
	return o.Scope6EndAddress6Addr, true
}

// HasScope6EndAddress6Addr returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6EndAddress6Addr() bool {
	if o != nil && o.Scope6EndAddress6Addr != nil {
		return true
	}

	return false
}

// SetScope6EndAddress6Addr gets a reference to the given string and assigns it to the Scope6EndAddress6Addr field.
func (o *DhcpScope6DataData) SetScope6EndAddress6Addr(v string) {
	o.Scope6EndAddress6Addr = &v
}

// GetScope6Id returns the Scope6Id field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6Id() string {
	if o == nil || o.Scope6Id == nil {
		var ret string
		return ret
	}
	return *o.Scope6Id
}

// GetScope6IdOk returns a tuple with the Scope6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6IdOk() (*string, bool) {
	if o == nil || o.Scope6Id == nil {
		return nil, false
	}
	return o.Scope6Id, true
}

// HasScope6Id returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6Id() bool {
	if o != nil && o.Scope6Id != nil {
		return true
	}

	return false
}

// SetScope6Id gets a reference to the given string and assigns it to the Scope6Id field.
func (o *DhcpScope6DataData) SetScope6Id(v string) {
	o.Scope6Id = &v
}

// GetScope6Name returns the Scope6Name field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6Name() string {
	if o == nil || o.Scope6Name == nil {
		var ret string
		return ret
	}
	return *o.Scope6Name
}

// GetScope6NameOk returns a tuple with the Scope6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6NameOk() (*string, bool) {
	if o == nil || o.Scope6Name == nil {
		return nil, false
	}
	return o.Scope6Name, true
}

// HasScope6Name returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6Name() bool {
	if o != nil && o.Scope6Name != nil {
		return true
	}

	return false
}

// SetScope6Name gets a reference to the given string and assigns it to the Scope6Name field.
func (o *DhcpScope6DataData) SetScope6Name(v string) {
	o.Scope6Name = &v
}

// GetScope6Prefix returns the Scope6Prefix field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6Prefix() string {
	if o == nil || o.Scope6Prefix == nil {
		var ret string
		return ret
	}
	return *o.Scope6Prefix
}

// GetScope6PrefixOk returns a tuple with the Scope6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6PrefixOk() (*string, bool) {
	if o == nil || o.Scope6Prefix == nil {
		return nil, false
	}
	return o.Scope6Prefix, true
}

// HasScope6Prefix returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6Prefix() bool {
	if o != nil && o.Scope6Prefix != nil {
		return true
	}

	return false
}

// SetScope6Prefix gets a reference to the given string and assigns it to the Scope6Prefix field.
func (o *DhcpScope6DataData) SetScope6Prefix(v string) {
	o.Scope6Prefix = &v
}

// GetScope6SpaceId returns the Scope6SpaceId field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6SpaceId() string {
	if o == nil || o.Scope6SpaceId == nil {
		var ret string
		return ret
	}
	return *o.Scope6SpaceId
}

// GetScope6SpaceIdOk returns a tuple with the Scope6SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6SpaceIdOk() (*string, bool) {
	if o == nil || o.Scope6SpaceId == nil {
		return nil, false
	}
	return o.Scope6SpaceId, true
}

// HasScope6SpaceId returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6SpaceId() bool {
	if o != nil && o.Scope6SpaceId != nil {
		return true
	}

	return false
}

// SetScope6SpaceId gets a reference to the given string and assigns it to the Scope6SpaceId field.
func (o *DhcpScope6DataData) SetScope6SpaceId(v string) {
	o.Scope6SpaceId = &v
}

// GetScope6SpaceName returns the Scope6SpaceName field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6SpaceName() string {
	if o == nil || o.Scope6SpaceName == nil {
		var ret string
		return ret
	}
	return *o.Scope6SpaceName
}

// GetScope6SpaceNameOk returns a tuple with the Scope6SpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6SpaceNameOk() (*string, bool) {
	if o == nil || o.Scope6SpaceName == nil {
		return nil, false
	}
	return o.Scope6SpaceName, true
}

// HasScope6SpaceName returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6SpaceName() bool {
	if o != nil && o.Scope6SpaceName != nil {
		return true
	}

	return false
}

// SetScope6SpaceName gets a reference to the given string and assigns it to the Scope6SpaceName field.
func (o *DhcpScope6DataData) SetScope6SpaceName(v string) {
	o.Scope6SpaceName = &v
}

// GetScope6Size returns the Scope6Size field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6Size() string {
	if o == nil || o.Scope6Size == nil {
		var ret string
		return ret
	}
	return *o.Scope6Size
}

// GetScope6SizeOk returns a tuple with the Scope6Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6SizeOk() (*string, bool) {
	if o == nil || o.Scope6Size == nil {
		return nil, false
	}
	return o.Scope6Size, true
}

// HasScope6Size returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6Size() bool {
	if o != nil && o.Scope6Size != nil {
		return true
	}

	return false
}

// SetScope6Size gets a reference to the given string and assigns it to the Scope6Size field.
func (o *DhcpScope6DataData) SetScope6Size(v string) {
	o.Scope6Size = &v
}

// GetScope6SortName returns the Scope6SortName field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6SortName() string {
	if o == nil || o.Scope6SortName == nil {
		var ret string
		return ret
	}
	return *o.Scope6SortName
}

// GetScope6SortNameOk returns a tuple with the Scope6SortName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6SortNameOk() (*string, bool) {
	if o == nil || o.Scope6SortName == nil {
		return nil, false
	}
	return o.Scope6SortName, true
}

// HasScope6SortName returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6SortName() bool {
	if o != nil && o.Scope6SortName != nil {
		return true
	}

	return false
}

// SetScope6SortName gets a reference to the given string and assigns it to the Scope6SortName field.
func (o *DhcpScope6DataData) SetScope6SortName(v string) {
	o.Scope6SortName = &v
}

// GetScope6StartAddress6Addr returns the Scope6StartAddress6Addr field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6StartAddress6Addr() string {
	if o == nil || o.Scope6StartAddress6Addr == nil {
		var ret string
		return ret
	}
	return *o.Scope6StartAddress6Addr
}

// GetScope6StartAddress6AddrOk returns a tuple with the Scope6StartAddress6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6StartAddress6AddrOk() (*string, bool) {
	if o == nil || o.Scope6StartAddress6Addr == nil {
		return nil, false
	}
	return o.Scope6StartAddress6Addr, true
}

// HasScope6StartAddress6Addr returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6StartAddress6Addr() bool {
	if o != nil && o.Scope6StartAddress6Addr != nil {
		return true
	}

	return false
}

// SetScope6StartAddress6Addr gets a reference to the given string and assigns it to the Scope6StartAddress6Addr field.
func (o *DhcpScope6DataData) SetScope6StartAddress6Addr(v string) {
	o.Scope6StartAddress6Addr = &v
}

// GetSharednetwork6Id returns the Sharednetwork6Id field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetSharednetwork6Id() string {
	if o == nil || o.Sharednetwork6Id == nil {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Id
}

// GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetSharednetwork6IdOk() (*string, bool) {
	if o == nil || o.Sharednetwork6Id == nil {
		return nil, false
	}
	return o.Sharednetwork6Id, true
}

// HasSharednetwork6Id returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasSharednetwork6Id() bool {
	if o != nil && o.Sharednetwork6Id != nil {
		return true
	}

	return false
}

// SetSharednetwork6Id gets a reference to the given string and assigns it to the Sharednetwork6Id field.
func (o *DhcpScope6DataData) SetSharednetwork6Id(v string) {
	o.Sharednetwork6Id = &v
}

// GetSharednetwork6Name returns the Sharednetwork6Name field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetSharednetwork6Name() string {
	if o == nil || o.Sharednetwork6Name == nil {
		var ret string
		return ret
	}
	return *o.Sharednetwork6Name
}

// GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetSharednetwork6NameOk() (*string, bool) {
	if o == nil || o.Sharednetwork6Name == nil {
		return nil, false
	}
	return o.Sharednetwork6Name, true
}

// HasSharednetwork6Name returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasSharednetwork6Name() bool {
	if o != nil && o.Sharednetwork6Name != nil {
		return true
	}

	return false
}

// SetSharednetwork6Name gets a reference to the given string and assigns it to the Sharednetwork6Name field.
func (o *DhcpScope6DataData) SetSharednetwork6Name(v string) {
	o.Sharednetwork6Name = &v
}

// GetServer6Addr returns the Server6Addr field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetServer6Addr() string {
	if o == nil || o.Server6Addr == nil {
		var ret string
		return ret
	}
	return *o.Server6Addr
}

// GetServer6AddrOk returns a tuple with the Server6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetServer6AddrOk() (*string, bool) {
	if o == nil || o.Server6Addr == nil {
		return nil, false
	}
	return o.Server6Addr, true
}

// HasServer6Addr returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasServer6Addr() bool {
	if o != nil && o.Server6Addr != nil {
		return true
	}

	return false
}

// SetServer6Addr gets a reference to the given string and assigns it to the Server6Addr field.
func (o *DhcpScope6DataData) SetServer6Addr(v string) {
	o.Server6Addr = &v
}

// GetScope6Multistatus returns the Scope6Multistatus field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetScope6Multistatus() string {
	if o == nil || o.Scope6Multistatus == nil {
		var ret string
		return ret
	}
	return *o.Scope6Multistatus
}

// GetScope6MultistatusOk returns a tuple with the Scope6Multistatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetScope6MultistatusOk() (*string, bool) {
	if o == nil || o.Scope6Multistatus == nil {
		return nil, false
	}
	return o.Scope6Multistatus, true
}

// HasScope6Multistatus returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasScope6Multistatus() bool {
	if o != nil && o.Scope6Multistatus != nil {
		return true
	}

	return false
}

// SetScope6Multistatus gets a reference to the given string and assigns it to the Scope6Multistatus field.
func (o *DhcpScope6DataData) SetScope6Multistatus(v string) {
	o.Scope6Multistatus = &v
}

// GetSmartArch returns the SmartArch field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetSmartArch() string {
	if o == nil || o.SmartArch == nil {
		var ret string
		return ret
	}
	return *o.SmartArch
}

// GetSmartArchOk returns a tuple with the SmartArch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetSmartArchOk() (*string, bool) {
	if o == nil || o.SmartArch == nil {
		return nil, false
	}
	return o.SmartArch, true
}

// HasSmartArch returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasSmartArch() bool {
	if o != nil && o.SmartArch != nil {
		return true
	}

	return false
}

// SetSmartArch gets a reference to the given string and assigns it to the SmartArch field.
func (o *DhcpScope6DataData) SetSmartArch(v string) {
	o.SmartArch = &v
}

// GetSmartParentId returns the SmartParentId field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetSmartParentId() string {
	if o == nil || o.SmartParentId == nil {
		var ret string
		return ret
	}
	return *o.SmartParentId
}

// GetSmartParentIdOk returns a tuple with the SmartParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetSmartParentIdOk() (*string, bool) {
	if o == nil || o.SmartParentId == nil {
		return nil, false
	}
	return o.SmartParentId, true
}

// HasSmartParentId returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasSmartParentId() bool {
	if o != nil && o.SmartParentId != nil {
		return true
	}

	return false
}

// SetSmartParentId gets a reference to the given string and assigns it to the SmartParentId field.
func (o *DhcpScope6DataData) SetSmartParentId(v string) {
	o.SmartParentId = &v
}

// GetSmartParentName returns the SmartParentName field value if set, zero value otherwise.
func (o *DhcpScope6DataData) GetSmartParentName() string {
	if o == nil || o.SmartParentName == nil {
		var ret string
		return ret
	}
	return *o.SmartParentName
}

// GetSmartParentNameOk returns a tuple with the SmartParentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DhcpScope6DataData) GetSmartParentNameOk() (*string, bool) {
	if o == nil || o.SmartParentName == nil {
		return nil, false
	}
	return o.SmartParentName, true
}

// HasSmartParentName returns a boolean if a field has been set.
func (o *DhcpScope6DataData) HasSmartParentName() bool {
	if o != nil && o.SmartParentName != nil {
		return true
	}

	return false
}

// SetSmartParentName gets a reference to the given string and assigns it to the SmartParentName field.
func (o *DhcpScope6DataData) SetSmartParentName(v string) {
	o.SmartParentName = &v
}

func (o DhcpScope6DataData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Scope6DelayedTime != nil {
		toSerialize["scope6_delayed_time"] = o.Scope6DelayedTime
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
	if o.Failover6Id != nil {
		toSerialize["failover6_id"] = o.Failover6Id
	}
	if o.Failover6Name != nil {
		toSerialize["failover6_name"] = o.Failover6Name
	}
	if o.Scope6ClassName != nil {
		toSerialize["scope6_class_name"] = o.Scope6ClassName
	}
	if o.Scope6ClassParameters != nil {
		toSerialize["scope6_class_parameters"] = o.Scope6ClassParameters
	}
	if o.Scope6EndAddress6Addr != nil {
		toSerialize["scope6_end_address6_addr"] = o.Scope6EndAddress6Addr
	}
	if o.Scope6Id != nil {
		toSerialize["scope6_id"] = o.Scope6Id
	}
	if o.Scope6Name != nil {
		toSerialize["scope6_name"] = o.Scope6Name
	}
	if o.Scope6Prefix != nil {
		toSerialize["scope6_prefix"] = o.Scope6Prefix
	}
	if o.Scope6SpaceId != nil {
		toSerialize["scope6_space_id"] = o.Scope6SpaceId
	}
	if o.Scope6SpaceName != nil {
		toSerialize["scope6_space_name"] = o.Scope6SpaceName
	}
	if o.Scope6Size != nil {
		toSerialize["scope6_size"] = o.Scope6Size
	}
	if o.Scope6SortName != nil {
		toSerialize["scope6_sort_name"] = o.Scope6SortName
	}
	if o.Scope6StartAddress6Addr != nil {
		toSerialize["scope6_start_address6_addr"] = o.Scope6StartAddress6Addr
	}
	if o.Sharednetwork6Id != nil {
		toSerialize["sharednetwork6_id"] = o.Sharednetwork6Id
	}
	if o.Sharednetwork6Name != nil {
		toSerialize["sharednetwork6_name"] = o.Sharednetwork6Name
	}
	if o.Server6Addr != nil {
		toSerialize["server6_addr"] = o.Server6Addr
	}
	if o.Scope6Multistatus != nil {
		toSerialize["scope6_multistatus"] = o.Scope6Multistatus
	}
	if o.SmartArch != nil {
		toSerialize["smart_arch"] = o.SmartArch
	}
	if o.SmartParentId != nil {
		toSerialize["smart_parent_id"] = o.SmartParentId
	}
	if o.SmartParentName != nil {
		toSerialize["smart_parent_name"] = o.SmartParentName
	}
	return json.Marshal(toSerialize)
}

type NullableDhcpScope6DataData struct {
	value *DhcpScope6DataData
	isSet bool
}

func (v NullableDhcpScope6DataData) Get() *DhcpScope6DataData {
	return v.value
}

func (v *NullableDhcpScope6DataData) Set(val *DhcpScope6DataData) {
	v.value = val
	v.isSet = true
}

func (v NullableDhcpScope6DataData) IsSet() bool {
	return v.isSet
}

func (v *NullableDhcpScope6DataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDhcpScope6DataData(val *DhcpScope6DataData) *NullableDhcpScope6DataData {
	return &NullableDhcpScope6DataData{value: val, isSet: true}
}

func (v NullableDhcpScope6DataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDhcpScope6DataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


