# DnsZoneparamDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneparamValue** | Pointer to **string** | The value of the DNS option set on the zone. | [optional] 
**ZoneparamDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**ZoneparamDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server the object belongs to. | [optional] 
**ZoneId** | Pointer to **string** | The database identifier (ID) of the DNS zone. | [optional] 
**ZoneName** | Pointer to **string** | The name of the DNS zone. | [optional] 
**Oid** | Pointer to **string** | The database identifier (ID) of the DNS zone param the object belongs to. | [optional] 
**ZoneparamKey** | Pointer to **string** | The name of the DNS option set on the zone. | [optional] 

## Methods

### NewDnsZoneparamDataData

`func NewDnsZoneparamDataData() *DnsZoneparamDataData`

NewDnsZoneparamDataData instantiates a new DnsZoneparamDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneparamDataDataWithDefaults

`func NewDnsZoneparamDataDataWithDefaults() *DnsZoneparamDataData`

NewDnsZoneparamDataDataWithDefaults instantiates a new DnsZoneparamDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneparamValue

`func (o *DnsZoneparamDataData) GetZoneparamValue() string`

GetZoneparamValue returns the ZoneparamValue field if non-nil, zero value otherwise.

### GetZoneparamValueOk

`func (o *DnsZoneparamDataData) GetZoneparamValueOk() (*string, bool)`

GetZoneparamValueOk returns a tuple with the ZoneparamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneparamValue

`func (o *DnsZoneparamDataData) SetZoneparamValue(v string)`

SetZoneparamValue sets ZoneparamValue field to given value.

### HasZoneparamValue

`func (o *DnsZoneparamDataData) HasZoneparamValue() bool`

HasZoneparamValue returns a boolean if a field has been set.

### GetZoneparamDelayedCreateTime

`func (o *DnsZoneparamDataData) GetZoneparamDelayedCreateTime() string`

GetZoneparamDelayedCreateTime returns the ZoneparamDelayedCreateTime field if non-nil, zero value otherwise.

### GetZoneparamDelayedCreateTimeOk

`func (o *DnsZoneparamDataData) GetZoneparamDelayedCreateTimeOk() (*string, bool)`

GetZoneparamDelayedCreateTimeOk returns a tuple with the ZoneparamDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneparamDelayedCreateTime

`func (o *DnsZoneparamDataData) SetZoneparamDelayedCreateTime(v string)`

SetZoneparamDelayedCreateTime sets ZoneparamDelayedCreateTime field to given value.

### HasZoneparamDelayedCreateTime

`func (o *DnsZoneparamDataData) HasZoneparamDelayedCreateTime() bool`

HasZoneparamDelayedCreateTime returns a boolean if a field has been set.

### GetZoneparamDelayedDeleteTime

`func (o *DnsZoneparamDataData) GetZoneparamDelayedDeleteTime() string`

GetZoneparamDelayedDeleteTime returns the ZoneparamDelayedDeleteTime field if non-nil, zero value otherwise.

### GetZoneparamDelayedDeleteTimeOk

`func (o *DnsZoneparamDataData) GetZoneparamDelayedDeleteTimeOk() (*string, bool)`

GetZoneparamDelayedDeleteTimeOk returns a tuple with the ZoneparamDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneparamDelayedDeleteTime

`func (o *DnsZoneparamDataData) SetZoneparamDelayedDeleteTime(v string)`

SetZoneparamDelayedDeleteTime sets ZoneparamDelayedDeleteTime field to given value.

### HasZoneparamDelayedDeleteTime

`func (o *DnsZoneparamDataData) HasZoneparamDelayedDeleteTime() bool`

HasZoneparamDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DnsZoneparamDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsZoneparamDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsZoneparamDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsZoneparamDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsZoneparamDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsZoneparamDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsZoneparamDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsZoneparamDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetZoneId

`func (o *DnsZoneparamDataData) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnsZoneparamDataData) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnsZoneparamDataData) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DnsZoneparamDataData) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneName

`func (o *DnsZoneparamDataData) GetZoneName() string`

GetZoneName returns the ZoneName field if non-nil, zero value otherwise.

### GetZoneNameOk

`func (o *DnsZoneparamDataData) GetZoneNameOk() (*string, bool)`

GetZoneNameOk returns a tuple with the ZoneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneName

`func (o *DnsZoneparamDataData) SetZoneName(v string)`

SetZoneName sets ZoneName field to given value.

### HasZoneName

`func (o *DnsZoneparamDataData) HasZoneName() bool`

HasZoneName returns a boolean if a field has been set.

### GetOid

`func (o *DnsZoneparamDataData) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *DnsZoneparamDataData) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *DnsZoneparamDataData) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *DnsZoneparamDataData) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetZoneparamKey

`func (o *DnsZoneparamDataData) GetZoneparamKey() string`

GetZoneparamKey returns the ZoneparamKey field if non-nil, zero value otherwise.

### GetZoneparamKeyOk

`func (o *DnsZoneparamDataData) GetZoneparamKeyOk() (*string, bool)`

GetZoneparamKeyOk returns a tuple with the ZoneparamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneparamKey

`func (o *DnsZoneparamDataData) SetZoneparamKey(v string)`

SetZoneparamKey sets ZoneparamKey field to given value.

### HasZoneparamKey

`func (o *DnsZoneparamDataData) HasZoneparamKey() bool`

HasZoneparamKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


