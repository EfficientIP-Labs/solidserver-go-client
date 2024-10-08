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

// checks if the DnsRrAddInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DnsRrAddInput{}

// DnsRrAddInput struct for DnsRrAddInput
type DnsRrAddInput struct {
	// The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice.
	ServerId *int32 `json:"server_id,omitempty"`
	// The name of the DNS server.
	ServerName *string `json:"server_name,omitempty"`
	// The IP address of the DNS server.
	ServerHostaddr *string `json:"server_hostaddr,omitempty"`
	// The full name of the DNS resource record. Specify it as value, it can either follow the format <b>&lt;rr-name&gt;.&lt;existing-zone-name&gt;.&lt;extension&gt;</b> or be <b>.</b> (a dot).
	RrName *string `json:"rr_name,omitempty"`
	// The type of the DNS resource record.<table><caption>rr_type possible values</caption><br/><thead><tr><th>Value</th><th>Record type description</th></tr><br/></thead><br/><tbody><tr><td >SOA</td><td >Start of Authority. Defines the zone name, an email contact and various time and refresh values applicable to the zone. It is automatically generated upon creation of a zone and cannot be added manually.</td></tr><tr><td >NS</td><td >Name Server. Defines the authoritative name server(s) for the domain (defined by the SOA record) or the subdomain. The NS record that indicates which server has authority over a zone is automatically generated upon the creation of a zone, once the server has been synchronized.</td></tr><tr><td >A</td><td >IPv4 Address. An IPv4 address for a host.</td></tr><tr><td >PTR</td><td >Pointer Record. Address Resolution, from an IP address (IPv4 or IPv6) to a host. Used in reverse mapping.</td></tr><tr><td >AAAA</td><td >IPv6 Address. An IPv6 address for a host.</td></tr><tr><td >CNAME</td><td >Canonical Name. An alias name for a host.</td></tr><tr><td >MX</td><td >Mail Exchange. The mail server/exchanger that services this zone.</td></tr><tr><td >SRV</td><td >Services record. Defines services available in the zone, for example, LDAP, HTTP, etc...</td></tr><tr><td >DNAME</td><td >Delegation of Reverse Names. Delegation of reverse addresses primarily in IPv6. (Deprecated, use the CNAME RR instead)</td></tr><tr><td >TXT</td><td >Text. Information associated with a name.</td></tr><tr><td >DS</td><td >Delegation Signer, a DNSSEC related RR used to verify the validity of the ZSK of a subdomain.</td></tr><tr><td >DNSKEY</td><td >DNS Key. It contains the public cryptographic key used to sign the zone with DNSSEC.</td></tr><tr><td >65534</td><td >A private type record automatically added to the zone once it is signed with DNSSEC.</td></tr><tr><td >HINFO</td><td >System Information. Information about a host: hardware type and operating system description.</td></tr><tr><td >MINFO</td><td >Mailbox mail list Information. Defines the mail administrator for a mail list and optionally a mailbox to receive error messages relating to the mail list.</td></tr><tr><td >AFSDB</td><td >AFS Database. Location of the AFS servers.</td></tr><tr><td >WKS</td><td >Well-Known Service. Defines the services and protocols supported by a host. (Deprecated, use the SRV RR instead)</td></tr><tr><td >NAPTR</td><td >Naming Authority Pointer Record. General purpose definition of rule set to be used by applications e.g. VoIP.</td></tr><tr><td >NSAP</td><td >Network Service Access Point. Defines record (equivalent of an A record) maps a host name to an endpoint address.</td></tr></tbody></table></p><br/>Note that the parameter is not case sensitive, you could type in <b>A</b> or <b>a</b>.
	RrType *string `json:"rr_type,omitempty"`
	// The first or only value required for the DNS resource record, as detailed in the service description.
	RrValue1 *string `json:"rr_value1,omitempty"`
	// A way to check the values of the DNS resource record before upon addition (<b>1)</b> in order to create a record with the same name but with different values.
	CheckValue *string `json:"check_value,omitempty"`
	// The database identifier (ID) of the DNS view, a unique numeric key value automatically incremented when you add a DNS view. Use the ID to specify the DNS view of your choice.
	ViewId *int32 `json:"view_id,omitempty"`
	// The name of the DNS view.
	ViewName *string `json:"view_name,omitempty"`
	// The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice.
	ZoneId *int32 `json:"zone_id,omitempty"`
	// The name of the DNS zone the object belongs to.
	ZoneName *string `json:"zone_name,omitempty"`
	// The database identifier (ID) of the space associated with the DNS zone the record belongs to.
	ZoneSpaceId *int32 `json:"zone_space_id,omitempty"`
	// The shortname of the DNS resource record.
	RrGlue *string `json:"rr_glue,omitempty"`
	// The time to live of the DNS resource record, in seconds.
	RrTtl *int32 `json:"rr_ttl,omitempty"`
	// The second value of the DNS resource record, depending on its type, as detailed in the service description.
	RrValue2 *string `json:"rr_value2,omitempty"`
	// The third value of the DNS resource record, depending on its type, as detailed in the service description.
	RrValue3 *string `json:"rr_value3,omitempty"`
	// The fourth value of the DNS resource record, depending on its type, as detailed in the service description.
	RrValue4 *string `json:"rr_value4,omitempty"`
	// The fifth value of the DNS resource record, depending on its type, as detailed in the service description.
	RrValue5 *string `json:"rr_value5,omitempty"`
	// The sixth value of the DNS resource record, depending on its type, as detailed in the service description.
	RrValue6 *string `json:"rr_value6,omitempty"`
	// The seventh value of the DNS resource record, depending on its type, as detailed in the service description.
	RrValue7 *string `json:"rr_value7,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete []string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	RrClassName *string `json:"rr_class_name,omitempty"`
	// class parameters in json format
	RrClassParameters []ApiClassParameterInputEntry `json:"rr_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewDnsRrAddInput instantiates a new DnsRrAddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDnsRrAddInput() *DnsRrAddInput {
	this := DnsRrAddInput{}
	return &this
}

// NewDnsRrAddInputWithDefaults instantiates a new DnsRrAddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDnsRrAddInputWithDefaults() *DnsRrAddInput {
	this := DnsRrAddInput{}
	return &this
}

// GetServerId returns the ServerId field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetServerId() int32 {
	if o == nil || IsNil(o.ServerId) {
		var ret int32
		return ret
	}
	return *o.ServerId
}

// GetServerIdOk returns a tuple with the ServerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetServerIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ServerId) {
		return nil, false
	}
	return o.ServerId, true
}

// HasServerId returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasServerId() bool {
	if o != nil && !IsNil(o.ServerId) {
		return true
	}

	return false
}

// SetServerId gets a reference to the given int32 and assigns it to the ServerId field.
func (o *DnsRrAddInput) SetServerId(v int32) {
	o.ServerId = &v
}

// GetServerName returns the ServerName field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetServerName() string {
	if o == nil || IsNil(o.ServerName) {
		var ret string
		return ret
	}
	return *o.ServerName
}

// GetServerNameOk returns a tuple with the ServerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetServerNameOk() (*string, bool) {
	if o == nil || IsNil(o.ServerName) {
		return nil, false
	}
	return o.ServerName, true
}

// HasServerName returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasServerName() bool {
	if o != nil && !IsNil(o.ServerName) {
		return true
	}

	return false
}

// SetServerName gets a reference to the given string and assigns it to the ServerName field.
func (o *DnsRrAddInput) SetServerName(v string) {
	o.ServerName = &v
}

// GetServerHostaddr returns the ServerHostaddr field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetServerHostaddr() string {
	if o == nil || IsNil(o.ServerHostaddr) {
		var ret string
		return ret
	}
	return *o.ServerHostaddr
}

// GetServerHostaddrOk returns a tuple with the ServerHostaddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetServerHostaddrOk() (*string, bool) {
	if o == nil || IsNil(o.ServerHostaddr) {
		return nil, false
	}
	return o.ServerHostaddr, true
}

// HasServerHostaddr returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasServerHostaddr() bool {
	if o != nil && !IsNil(o.ServerHostaddr) {
		return true
	}

	return false
}

// SetServerHostaddr gets a reference to the given string and assigns it to the ServerHostaddr field.
func (o *DnsRrAddInput) SetServerHostaddr(v string) {
	o.ServerHostaddr = &v
}

// GetRrName returns the RrName field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrName() string {
	if o == nil || IsNil(o.RrName) {
		var ret string
		return ret
	}
	return *o.RrName
}

// GetRrNameOk returns a tuple with the RrName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrNameOk() (*string, bool) {
	if o == nil || IsNil(o.RrName) {
		return nil, false
	}
	return o.RrName, true
}

// HasRrName returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrName() bool {
	if o != nil && !IsNil(o.RrName) {
		return true
	}

	return false
}

// SetRrName gets a reference to the given string and assigns it to the RrName field.
func (o *DnsRrAddInput) SetRrName(v string) {
	o.RrName = &v
}

// GetRrType returns the RrType field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrType() string {
	if o == nil || IsNil(o.RrType) {
		var ret string
		return ret
	}
	return *o.RrType
}

// GetRrTypeOk returns a tuple with the RrType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RrType) {
		return nil, false
	}
	return o.RrType, true
}

// HasRrType returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrType() bool {
	if o != nil && !IsNil(o.RrType) {
		return true
	}

	return false
}

// SetRrType gets a reference to the given string and assigns it to the RrType field.
func (o *DnsRrAddInput) SetRrType(v string) {
	o.RrType = &v
}

// GetRrValue1 returns the RrValue1 field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrValue1() string {
	if o == nil || IsNil(o.RrValue1) {
		var ret string
		return ret
	}
	return *o.RrValue1
}

// GetRrValue1Ok returns a tuple with the RrValue1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrValue1Ok() (*string, bool) {
	if o == nil || IsNil(o.RrValue1) {
		return nil, false
	}
	return o.RrValue1, true
}

// HasRrValue1 returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrValue1() bool {
	if o != nil && !IsNil(o.RrValue1) {
		return true
	}

	return false
}

// SetRrValue1 gets a reference to the given string and assigns it to the RrValue1 field.
func (o *DnsRrAddInput) SetRrValue1(v string) {
	o.RrValue1 = &v
}

// GetCheckValue returns the CheckValue field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetCheckValue() string {
	if o == nil || IsNil(o.CheckValue) {
		var ret string
		return ret
	}
	return *o.CheckValue
}

// GetCheckValueOk returns a tuple with the CheckValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetCheckValueOk() (*string, bool) {
	if o == nil || IsNil(o.CheckValue) {
		return nil, false
	}
	return o.CheckValue, true
}

// HasCheckValue returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasCheckValue() bool {
	if o != nil && !IsNil(o.CheckValue) {
		return true
	}

	return false
}

// SetCheckValue gets a reference to the given string and assigns it to the CheckValue field.
func (o *DnsRrAddInput) SetCheckValue(v string) {
	o.CheckValue = &v
}

// GetViewId returns the ViewId field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetViewId() int32 {
	if o == nil || IsNil(o.ViewId) {
		var ret int32
		return ret
	}
	return *o.ViewId
}

// GetViewIdOk returns a tuple with the ViewId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetViewIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ViewId) {
		return nil, false
	}
	return o.ViewId, true
}

// HasViewId returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasViewId() bool {
	if o != nil && !IsNil(o.ViewId) {
		return true
	}

	return false
}

// SetViewId gets a reference to the given int32 and assigns it to the ViewId field.
func (o *DnsRrAddInput) SetViewId(v int32) {
	o.ViewId = &v
}

// GetViewName returns the ViewName field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetViewName() string {
	if o == nil || IsNil(o.ViewName) {
		var ret string
		return ret
	}
	return *o.ViewName
}

// GetViewNameOk returns a tuple with the ViewName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetViewNameOk() (*string, bool) {
	if o == nil || IsNil(o.ViewName) {
		return nil, false
	}
	return o.ViewName, true
}

// HasViewName returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasViewName() bool {
	if o != nil && !IsNil(o.ViewName) {
		return true
	}

	return false
}

// SetViewName gets a reference to the given string and assigns it to the ViewName field.
func (o *DnsRrAddInput) SetViewName(v string) {
	o.ViewName = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetZoneId() int32 {
	if o == nil || IsNil(o.ZoneId) {
		var ret int32
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetZoneIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ZoneId) {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasZoneId() bool {
	if o != nil && !IsNil(o.ZoneId) {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given int32 and assigns it to the ZoneId field.
func (o *DnsRrAddInput) SetZoneId(v int32) {
	o.ZoneId = &v
}

// GetZoneName returns the ZoneName field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetZoneName() string {
	if o == nil || IsNil(o.ZoneName) {
		var ret string
		return ret
	}
	return *o.ZoneName
}

// GetZoneNameOk returns a tuple with the ZoneName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetZoneNameOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneName) {
		return nil, false
	}
	return o.ZoneName, true
}

// HasZoneName returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasZoneName() bool {
	if o != nil && !IsNil(o.ZoneName) {
		return true
	}

	return false
}

// SetZoneName gets a reference to the given string and assigns it to the ZoneName field.
func (o *DnsRrAddInput) SetZoneName(v string) {
	o.ZoneName = &v
}

// GetZoneSpaceId returns the ZoneSpaceId field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetZoneSpaceId() int32 {
	if o == nil || IsNil(o.ZoneSpaceId) {
		var ret int32
		return ret
	}
	return *o.ZoneSpaceId
}

// GetZoneSpaceIdOk returns a tuple with the ZoneSpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetZoneSpaceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ZoneSpaceId) {
		return nil, false
	}
	return o.ZoneSpaceId, true
}

// HasZoneSpaceId returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasZoneSpaceId() bool {
	if o != nil && !IsNil(o.ZoneSpaceId) {
		return true
	}

	return false
}

// SetZoneSpaceId gets a reference to the given int32 and assigns it to the ZoneSpaceId field.
func (o *DnsRrAddInput) SetZoneSpaceId(v int32) {
	o.ZoneSpaceId = &v
}

// GetRrGlue returns the RrGlue field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrGlue() string {
	if o == nil || IsNil(o.RrGlue) {
		var ret string
		return ret
	}
	return *o.RrGlue
}

// GetRrGlueOk returns a tuple with the RrGlue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrGlueOk() (*string, bool) {
	if o == nil || IsNil(o.RrGlue) {
		return nil, false
	}
	return o.RrGlue, true
}

// HasRrGlue returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrGlue() bool {
	if o != nil && !IsNil(o.RrGlue) {
		return true
	}

	return false
}

// SetRrGlue gets a reference to the given string and assigns it to the RrGlue field.
func (o *DnsRrAddInput) SetRrGlue(v string) {
	o.RrGlue = &v
}

// GetRrTtl returns the RrTtl field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrTtl() int32 {
	if o == nil || IsNil(o.RrTtl) {
		var ret int32
		return ret
	}
	return *o.RrTtl
}

// GetRrTtlOk returns a tuple with the RrTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.RrTtl) {
		return nil, false
	}
	return o.RrTtl, true
}

// HasRrTtl returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrTtl() bool {
	if o != nil && !IsNil(o.RrTtl) {
		return true
	}

	return false
}

// SetRrTtl gets a reference to the given int32 and assigns it to the RrTtl field.
func (o *DnsRrAddInput) SetRrTtl(v int32) {
	o.RrTtl = &v
}

// GetRrValue2 returns the RrValue2 field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrValue2() string {
	if o == nil || IsNil(o.RrValue2) {
		var ret string
		return ret
	}
	return *o.RrValue2
}

// GetRrValue2Ok returns a tuple with the RrValue2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrValue2Ok() (*string, bool) {
	if o == nil || IsNil(o.RrValue2) {
		return nil, false
	}
	return o.RrValue2, true
}

// HasRrValue2 returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrValue2() bool {
	if o != nil && !IsNil(o.RrValue2) {
		return true
	}

	return false
}

// SetRrValue2 gets a reference to the given string and assigns it to the RrValue2 field.
func (o *DnsRrAddInput) SetRrValue2(v string) {
	o.RrValue2 = &v
}

// GetRrValue3 returns the RrValue3 field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrValue3() string {
	if o == nil || IsNil(o.RrValue3) {
		var ret string
		return ret
	}
	return *o.RrValue3
}

// GetRrValue3Ok returns a tuple with the RrValue3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrValue3Ok() (*string, bool) {
	if o == nil || IsNil(o.RrValue3) {
		return nil, false
	}
	return o.RrValue3, true
}

// HasRrValue3 returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrValue3() bool {
	if o != nil && !IsNil(o.RrValue3) {
		return true
	}

	return false
}

// SetRrValue3 gets a reference to the given string and assigns it to the RrValue3 field.
func (o *DnsRrAddInput) SetRrValue3(v string) {
	o.RrValue3 = &v
}

// GetRrValue4 returns the RrValue4 field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrValue4() string {
	if o == nil || IsNil(o.RrValue4) {
		var ret string
		return ret
	}
	return *o.RrValue4
}

// GetRrValue4Ok returns a tuple with the RrValue4 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrValue4Ok() (*string, bool) {
	if o == nil || IsNil(o.RrValue4) {
		return nil, false
	}
	return o.RrValue4, true
}

// HasRrValue4 returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrValue4() bool {
	if o != nil && !IsNil(o.RrValue4) {
		return true
	}

	return false
}

// SetRrValue4 gets a reference to the given string and assigns it to the RrValue4 field.
func (o *DnsRrAddInput) SetRrValue4(v string) {
	o.RrValue4 = &v
}

// GetRrValue5 returns the RrValue5 field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrValue5() string {
	if o == nil || IsNil(o.RrValue5) {
		var ret string
		return ret
	}
	return *o.RrValue5
}

// GetRrValue5Ok returns a tuple with the RrValue5 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrValue5Ok() (*string, bool) {
	if o == nil || IsNil(o.RrValue5) {
		return nil, false
	}
	return o.RrValue5, true
}

// HasRrValue5 returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrValue5() bool {
	if o != nil && !IsNil(o.RrValue5) {
		return true
	}

	return false
}

// SetRrValue5 gets a reference to the given string and assigns it to the RrValue5 field.
func (o *DnsRrAddInput) SetRrValue5(v string) {
	o.RrValue5 = &v
}

// GetRrValue6 returns the RrValue6 field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrValue6() string {
	if o == nil || IsNil(o.RrValue6) {
		var ret string
		return ret
	}
	return *o.RrValue6
}

// GetRrValue6Ok returns a tuple with the RrValue6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrValue6Ok() (*string, bool) {
	if o == nil || IsNil(o.RrValue6) {
		return nil, false
	}
	return o.RrValue6, true
}

// HasRrValue6 returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrValue6() bool {
	if o != nil && !IsNil(o.RrValue6) {
		return true
	}

	return false
}

// SetRrValue6 gets a reference to the given string and assigns it to the RrValue6 field.
func (o *DnsRrAddInput) SetRrValue6(v string) {
	o.RrValue6 = &v
}

// GetRrValue7 returns the RrValue7 field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrValue7() string {
	if o == nil || IsNil(o.RrValue7) {
		var ret string
		return ret
	}
	return *o.RrValue7
}

// GetRrValue7Ok returns a tuple with the RrValue7 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrValue7Ok() (*string, bool) {
	if o == nil || IsNil(o.RrValue7) {
		return nil, false
	}
	return o.RrValue7, true
}

// HasRrValue7 returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrValue7() bool {
	if o != nil && !IsNil(o.RrValue7) {
		return true
	}

	return false
}

// SetRrValue7 gets a reference to the given string and assigns it to the RrValue7 field.
func (o *DnsRrAddInput) SetRrValue7(v string) {
	o.RrValue7 = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetClassParametersToDelete() []string {
	if o == nil || IsNil(o.ClassParametersToDelete) {
		var ret []string
		return ret
	}
	return o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetClassParametersToDeleteOk() ([]string, bool) {
	if o == nil || IsNil(o.ClassParametersToDelete) {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasClassParametersToDelete() bool {
	if o != nil && !IsNil(o.ClassParametersToDelete) {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *DnsRrAddInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = v
}

// GetRrClassName returns the RrClassName field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrClassName() string {
	if o == nil || IsNil(o.RrClassName) {
		var ret string
		return ret
	}
	return *o.RrClassName
}

// GetRrClassNameOk returns a tuple with the RrClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.RrClassName) {
		return nil, false
	}
	return o.RrClassName, true
}

// HasRrClassName returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrClassName() bool {
	if o != nil && !IsNil(o.RrClassName) {
		return true
	}

	return false
}

// SetRrClassName gets a reference to the given string and assigns it to the RrClassName field.
func (o *DnsRrAddInput) SetRrClassName(v string) {
	o.RrClassName = &v
}

// GetRrClassParameters returns the RrClassParameters field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetRrClassParameters() []ApiClassParameterInputEntry {
	if o == nil || IsNil(o.RrClassParameters) {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return o.RrClassParameters
}

// GetRrClassParametersOk returns a tuple with the RrClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetRrClassParametersOk() ([]ApiClassParameterInputEntry, bool) {
	if o == nil || IsNil(o.RrClassParameters) {
		return nil, false
	}
	return o.RrClassParameters, true
}

// HasRrClassParameters returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasRrClassParameters() bool {
	if o != nil && !IsNil(o.RrClassParameters) {
		return true
	}

	return false
}

// SetRrClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the RrClassParameters field.
func (o *DnsRrAddInput) SetRrClassParameters(v []ApiClassParameterInputEntry) {
	o.RrClassParameters = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DnsRrAddInput) GetWarnings() string {
	if o == nil || IsNil(o.Warnings) {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DnsRrAddInput) GetWarningsOk() (*string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DnsRrAddInput) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *DnsRrAddInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o DnsRrAddInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DnsRrAddInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServerId) {
		toSerialize["server_id"] = o.ServerId
	}
	if !IsNil(o.ServerName) {
		toSerialize["server_name"] = o.ServerName
	}
	if !IsNil(o.ServerHostaddr) {
		toSerialize["server_hostaddr"] = o.ServerHostaddr
	}
	if !IsNil(o.RrName) {
		toSerialize["rr_name"] = o.RrName
	}
	if !IsNil(o.RrType) {
		toSerialize["rr_type"] = o.RrType
	}
	if !IsNil(o.RrValue1) {
		toSerialize["rr_value1"] = o.RrValue1
	}
	if !IsNil(o.CheckValue) {
		toSerialize["check_value"] = o.CheckValue
	}
	if !IsNil(o.ViewId) {
		toSerialize["view_id"] = o.ViewId
	}
	if !IsNil(o.ViewName) {
		toSerialize["view_name"] = o.ViewName
	}
	if !IsNil(o.ZoneId) {
		toSerialize["zone_id"] = o.ZoneId
	}
	if !IsNil(o.ZoneName) {
		toSerialize["zone_name"] = o.ZoneName
	}
	if !IsNil(o.ZoneSpaceId) {
		toSerialize["zone_space_id"] = o.ZoneSpaceId
	}
	if !IsNil(o.RrGlue) {
		toSerialize["rr_glue"] = o.RrGlue
	}
	if !IsNil(o.RrTtl) {
		toSerialize["rr_ttl"] = o.RrTtl
	}
	if !IsNil(o.RrValue2) {
		toSerialize["rr_value2"] = o.RrValue2
	}
	if !IsNil(o.RrValue3) {
		toSerialize["rr_value3"] = o.RrValue3
	}
	if !IsNil(o.RrValue4) {
		toSerialize["rr_value4"] = o.RrValue4
	}
	if !IsNil(o.RrValue5) {
		toSerialize["rr_value5"] = o.RrValue5
	}
	if !IsNil(o.RrValue6) {
		toSerialize["rr_value6"] = o.RrValue6
	}
	if !IsNil(o.RrValue7) {
		toSerialize["rr_value7"] = o.RrValue7
	}
	if !IsNil(o.ClassParametersToDelete) {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if !IsNil(o.RrClassName) {
		toSerialize["rr_class_name"] = o.RrClassName
	}
	if !IsNil(o.RrClassParameters) {
		toSerialize["rr_class_parameters"] = o.RrClassParameters
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableDnsRrAddInput struct {
	value *DnsRrAddInput
	isSet bool
}

func (v NullableDnsRrAddInput) Get() *DnsRrAddInput {
	return v.value
}

func (v *NullableDnsRrAddInput) Set(val *DnsRrAddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsRrAddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsRrAddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsRrAddInput(val *DnsRrAddInput) *NullableDnsRrAddInput {
	return &NullableDnsRrAddInput{value: val, isSet: true}
}

func (v NullableDnsRrAddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsRrAddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
