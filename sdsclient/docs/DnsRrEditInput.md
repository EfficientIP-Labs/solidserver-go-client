# DnsRrEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DNS server. | [optional] 
**RrId** | Pointer to **int32** | The database identifier (ID) of the DNS resource record, a unique numeric key value automatically incremented when you add a DNS RR. Use the ID to specify which record to edit. | [optional] 
**RrName** | Pointer to **string** | The full name of the DNS resource record. Specify it as value, it can either follow the format &lt;b&gt;&amp;lt;rr-name&amp;gt;.&amp;lt;existing-zone-name&amp;gt;.&amp;lt;extension&amp;gt;&lt;/b&gt; or be &lt;b&gt;.&lt;/b&gt; (a dot). | [optional] 
**RrType** | Pointer to **string** | The type of the DNS resource record.&lt;table&gt;&lt;caption&gt;rr_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Value&lt;/th&gt;&lt;th&gt;Record type description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;SOA&lt;/td&gt;&lt;td &gt;Start of Authority. Defines the zone name, an email contact and various time and refresh values applicable to the zone. It is automatically generated upon creation of a zone and cannot be added manually.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NS&lt;/td&gt;&lt;td &gt;Name Server. Defines the authoritative name server(s) for the domain (defined by the SOA record) or the subdomain. The NS record that indicates which server has authority over a zone is automatically generated upon the creation of a zone, once the server has been synchronized.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;A&lt;/td&gt;&lt;td &gt;IPv4 Address. An IPv4 address for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;PTR&lt;/td&gt;&lt;td &gt;Pointer Record. Address Resolution, from an IP address (IPv4 or IPv6) to a host. Used in reverse mapping.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;AAAA&lt;/td&gt;&lt;td &gt;IPv6 Address. An IPv6 address for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;CNAME&lt;/td&gt;&lt;td &gt;Canonical Name. An alias name for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;MX&lt;/td&gt;&lt;td &gt;Mail Exchange. The mail server/exchanger that services this zone.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;SRV&lt;/td&gt;&lt;td &gt;Services record. Defines services available in the zone, for example, LDAP, HTTP, etc...&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DNAME&lt;/td&gt;&lt;td &gt;Delegation of Reverse Names. Delegation of reverse addresses primarily in IPv6. (Deprecated, use the CNAME RR instead)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;TXT&lt;/td&gt;&lt;td &gt;Text. Information associated with a name.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DS&lt;/td&gt;&lt;td &gt;Delegation Signer, a DNSSEC related RR used to verify the validity of the ZSK of a subdomain.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DNSKEY&lt;/td&gt;&lt;td &gt;DNS Key. It contains the public cryptographic key used to sign the zone with DNSSEC.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;65534&lt;/td&gt;&lt;td &gt;A private type record automatically added to the zone once it is signed with DNSSEC.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;HINFO&lt;/td&gt;&lt;td &gt;System Information. Information about a host: hardware type and operating system description.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;MINFO&lt;/td&gt;&lt;td &gt;Mailbox mail list Information. Defines the mail administrator for a mail list and optionally a mailbox to receive error messages relating to the mail list.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;AFSDB&lt;/td&gt;&lt;td &gt;AFS Database. Location of the AFS servers.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;WKS&lt;/td&gt;&lt;td &gt;Well-Known Service. Defines the services and protocols supported by a host. (Deprecated, use the SRV RR instead)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NAPTR&lt;/td&gt;&lt;td &gt;Naming Authority Pointer Record. General purpose definition of rule set to be used by applications e.g. VoIP.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NSAP&lt;/td&gt;&lt;td &gt;Network Service Access Point. Defines record (equivalent of an A record) maps a host name to an endpoint address.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt;Note that the parameter is not case sensitive, you could specify &lt;b&gt;A&lt;/b&gt; or &lt;b&gt;a&lt;/b&gt;. | [optional] 
**RrValue1** | Pointer to **string** | The first or only value required for the DNS resource record, as detailed in the service description. | [optional] 
**CheckValue** | Pointer to **string** | A way to check the values of the DNS resource record before upon addition (&lt;b&gt;1)&lt;/b&gt; in order to create a record with the same name but with different values. | [optional] 
**ViewId** | Pointer to **int32** | The database identifier (ID) of the DNS view, a unique numeric key value automatically incremented when you add a DNS view. Use the ID to specify the DNS view of your choice. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view. | [optional] 
**ZoneId** | Pointer to **int32** | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. | [optional] 
**ZoneName** | Pointer to **string** | The name of the DNS zone the object belongs to. | [optional] 
**ZoneSpaceId** | Pointer to **int32** | The database identifier (ID) of the space associated with the DNS zone the record belongs to. | [optional] 
**RrGlue** | Pointer to **string** | The shortname of the DNS resource record. | [optional] 
**RrTtl** | Pointer to **int32** | The time to live of the DNS resource record, in seconds. | [optional] 
**RrValue2** | Pointer to **string** | The second value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue3** | Pointer to **string** | The third value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue4** | Pointer to **string** | The fourth value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue5** | Pointer to **string** | The fifth value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue6** | Pointer to **string** | The sixth value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**RrValue7** | Pointer to **string** | The seventh value of the DNS resource record, depending on its type, as detailed in the service description. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**RrClassName** | Pointer to **string** | TODO:dns_rr_add.input.rr_class_name | [optional] 
**RrClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDnsRrEditInput

`func NewDnsRrEditInput() *DnsRrEditInput`

NewDnsRrEditInput instantiates a new DnsRrEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsRrEditInputWithDefaults

`func NewDnsRrEditInputWithDefaults() *DnsRrEditInput`

NewDnsRrEditInputWithDefaults instantiates a new DnsRrEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DnsRrEditInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsRrEditInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsRrEditInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsRrEditInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsRrEditInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsRrEditInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsRrEditInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsRrEditInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DnsRrEditInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DnsRrEditInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DnsRrEditInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DnsRrEditInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetRrId

`func (o *DnsRrEditInput) GetRrId() int32`

GetRrId returns the RrId field if non-nil, zero value otherwise.

### GetRrIdOk

`func (o *DnsRrEditInput) GetRrIdOk() (*int32, bool)`

GetRrIdOk returns a tuple with the RrId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrId

`func (o *DnsRrEditInput) SetRrId(v int32)`

SetRrId sets RrId field to given value.

### HasRrId

`func (o *DnsRrEditInput) HasRrId() bool`

HasRrId returns a boolean if a field has been set.

### GetRrName

`func (o *DnsRrEditInput) GetRrName() string`

GetRrName returns the RrName field if non-nil, zero value otherwise.

### GetRrNameOk

`func (o *DnsRrEditInput) GetRrNameOk() (*string, bool)`

GetRrNameOk returns a tuple with the RrName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrName

`func (o *DnsRrEditInput) SetRrName(v string)`

SetRrName sets RrName field to given value.

### HasRrName

`func (o *DnsRrEditInput) HasRrName() bool`

HasRrName returns a boolean if a field has been set.

### GetRrType

`func (o *DnsRrEditInput) GetRrType() string`

GetRrType returns the RrType field if non-nil, zero value otherwise.

### GetRrTypeOk

`func (o *DnsRrEditInput) GetRrTypeOk() (*string, bool)`

GetRrTypeOk returns a tuple with the RrType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrType

`func (o *DnsRrEditInput) SetRrType(v string)`

SetRrType sets RrType field to given value.

### HasRrType

`func (o *DnsRrEditInput) HasRrType() bool`

HasRrType returns a boolean if a field has been set.

### GetRrValue1

`func (o *DnsRrEditInput) GetRrValue1() string`

GetRrValue1 returns the RrValue1 field if non-nil, zero value otherwise.

### GetRrValue1Ok

`func (o *DnsRrEditInput) GetRrValue1Ok() (*string, bool)`

GetRrValue1Ok returns a tuple with the RrValue1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue1

`func (o *DnsRrEditInput) SetRrValue1(v string)`

SetRrValue1 sets RrValue1 field to given value.

### HasRrValue1

`func (o *DnsRrEditInput) HasRrValue1() bool`

HasRrValue1 returns a boolean if a field has been set.

### GetCheckValue

`func (o *DnsRrEditInput) GetCheckValue() string`

GetCheckValue returns the CheckValue field if non-nil, zero value otherwise.

### GetCheckValueOk

`func (o *DnsRrEditInput) GetCheckValueOk() (*string, bool)`

GetCheckValueOk returns a tuple with the CheckValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckValue

`func (o *DnsRrEditInput) SetCheckValue(v string)`

SetCheckValue sets CheckValue field to given value.

### HasCheckValue

`func (o *DnsRrEditInput) HasCheckValue() bool`

HasCheckValue returns a boolean if a field has been set.

### GetViewId

`func (o *DnsRrEditInput) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DnsRrEditInput) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DnsRrEditInput) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DnsRrEditInput) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewName

`func (o *DnsRrEditInput) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DnsRrEditInput) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DnsRrEditInput) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DnsRrEditInput) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetZoneId

`func (o *DnsRrEditInput) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnsRrEditInput) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnsRrEditInput) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DnsRrEditInput) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneName

`func (o *DnsRrEditInput) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DnsRrEditInput) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DnsRrEditInput) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *DnsRrEditInput) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetZoneSpaceId

`func (o *DnsRrEditInput) GetZoneSpaceId() int32`

GetZoneSpaceId returns the ZoneSpaceId field if non-nil, zero value otherwise.

### GetZoneSpaceIdOk

`func (o *DnsRrEditInput) GetZoneSpaceIdOk() (*int32, bool)`

GetZoneSpaceIdOk returns a tuple with the ZoneSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneSpaceId

`func (o *DnsRrEditInput) SetZoneSpaceId(v int32)`

SetZoneSpaceId sets ZoneSpaceId field to given value.

### HasZoneSpaceId

`func (o *DnsRrEditInput) HasZoneSpaceId() bool`

HasZoneSpaceId returns a boolean if a field has been set.

### GetRrGlue

`func (o *DnsRrEditInput) GetRrGlue() string`

GetRrGlue returns the RrGlue field if non-nil, zero value otherwise.

### GetRrGlueOk

`func (o *DnsRrEditInput) GetRrGlueOk() (*string, bool)`

GetRrGlueOk returns a tuple with the RrGlue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrGlue

`func (o *DnsRrEditInput) SetRrGlue(v string)`

SetRrGlue sets RrGlue field to given value.

### HasRrGlue

`func (o *DnsRrEditInput) HasRrGlue() bool`

HasRrGlue returns a boolean if a field has been set.

### GetRrTtl

`func (o *DnsRrEditInput) GetRrTtl() int32`

GetRrTtl returns the RrTtl field if non-nil, zero value otherwise.

### GetRrTtlOk

`func (o *DnsRrEditInput) GetRrTtlOk() (*int32, bool)`

GetRrTtlOk returns a tuple with the RrTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrTtl

`func (o *DnsRrEditInput) SetRrTtl(v int32)`

SetRrTtl sets RrTtl field to given value.

### HasRrTtl

`func (o *DnsRrEditInput) HasRrTtl() bool`

HasRrTtl returns a boolean if a field has been set.

### GetRrValue2

`func (o *DnsRrEditInput) GetRrValue2() string`

GetRrValue2 returns the RrValue2 field if non-nil, zero value otherwise.

### GetRrValue2Ok

`func (o *DnsRrEditInput) GetRrValue2Ok() (*string, bool)`

GetRrValue2Ok returns a tuple with the RrValue2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue2

`func (o *DnsRrEditInput) SetRrValue2(v string)`

SetRrValue2 sets RrValue2 field to given value.

### HasRrValue2

`func (o *DnsRrEditInput) HasRrValue2() bool`

HasRrValue2 returns a boolean if a field has been set.

### GetRrValue3

`func (o *DnsRrEditInput) GetRrValue3() string`

GetRrValue3 returns the RrValue3 field if non-nil, zero value otherwise.

### GetRrValue3Ok

`func (o *DnsRrEditInput) GetRrValue3Ok() (*string, bool)`

GetRrValue3Ok returns a tuple with the RrValue3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue3

`func (o *DnsRrEditInput) SetRrValue3(v string)`

SetRrValue3 sets RrValue3 field to given value.

### HasRrValue3

`func (o *DnsRrEditInput) HasRrValue3() bool`

HasRrValue3 returns a boolean if a field has been set.

### GetRrValue4

`func (o *DnsRrEditInput) GetRrValue4() string`

GetRrValue4 returns the RrValue4 field if non-nil, zero value otherwise.

### GetRrValue4Ok

`func (o *DnsRrEditInput) GetRrValue4Ok() (*string, bool)`

GetRrValue4Ok returns a tuple with the RrValue4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue4

`func (o *DnsRrEditInput) SetRrValue4(v string)`

SetRrValue4 sets RrValue4 field to given value.

### HasRrValue4

`func (o *DnsRrEditInput) HasRrValue4() bool`

HasRrValue4 returns a boolean if a field has been set.

### GetRrValue5

`func (o *DnsRrEditInput) GetRrValue5() string`

GetRrValue5 returns the RrValue5 field if non-nil, zero value otherwise.

### GetRrValue5Ok

`func (o *DnsRrEditInput) GetRrValue5Ok() (*string, bool)`

GetRrValue5Ok returns a tuple with the RrValue5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue5

`func (o *DnsRrEditInput) SetRrValue5(v string)`

SetRrValue5 sets RrValue5 field to given value.

### HasRrValue5

`func (o *DnsRrEditInput) HasRrValue5() bool`

HasRrValue5 returns a boolean if a field has been set.

### GetRrValue6

`func (o *DnsRrEditInput) GetRrValue6() string`

GetRrValue6 returns the RrValue6 field if non-nil, zero value otherwise.

### GetRrValue6Ok

`func (o *DnsRrEditInput) GetRrValue6Ok() (*string, bool)`

GetRrValue6Ok returns a tuple with the RrValue6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue6

`func (o *DnsRrEditInput) SetRrValue6(v string)`

SetRrValue6 sets RrValue6 field to given value.

### HasRrValue6

`func (o *DnsRrEditInput) HasRrValue6() bool`

HasRrValue6 returns a boolean if a field has been set.

### GetRrValue7

`func (o *DnsRrEditInput) GetRrValue7() string`

GetRrValue7 returns the RrValue7 field if non-nil, zero value otherwise.

### GetRrValue7Ok

`func (o *DnsRrEditInput) GetRrValue7Ok() (*string, bool)`

GetRrValue7Ok returns a tuple with the RrValue7 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrValue7

`func (o *DnsRrEditInput) SetRrValue7(v string)`

SetRrValue7 sets RrValue7 field to given value.

### HasRrValue7

`func (o *DnsRrEditInput) HasRrValue7() bool`

HasRrValue7 returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DnsRrEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DnsRrEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DnsRrEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DnsRrEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetRrClassName

`func (o *DnsRrEditInput) GetRrClassName() string`

GetRrClassName returns the RrClassName field if non-nil, zero value otherwise.

### GetRrClassNameOk

`func (o *DnsRrEditInput) GetRrClassNameOk() (*string, bool)`

GetRrClassNameOk returns a tuple with the RrClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrClassName

`func (o *DnsRrEditInput) SetRrClassName(v string)`

SetRrClassName sets RrClassName field to given value.

### HasRrClassName

`func (o *DnsRrEditInput) HasRrClassName() bool`

HasRrClassName returns a boolean if a field has been set.

### GetRrClassParameters

`func (o *DnsRrEditInput) GetRrClassParameters() []ApiClassParameterInputEntry`

GetRrClassParameters returns the RrClassParameters field if non-nil, zero value otherwise.

### GetRrClassParametersOk

`func (o *DnsRrEditInput) GetRrClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetRrClassParametersOk returns a tuple with the RrClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRrClassParameters

`func (o *DnsRrEditInput) SetRrClassParameters(v []ApiClassParameterInputEntry)`

SetRrClassParameters sets RrClassParameters field to given value.

### HasRrClassParameters

`func (o *DnsRrEditInput) HasRrClassParameters() bool`

HasRrClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DnsRrEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DnsRrEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DnsRrEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DnsRrEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


