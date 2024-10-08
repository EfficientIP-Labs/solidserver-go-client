# DnsZoneparamEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ZoneId** | Pointer to **int32** | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. | [optional] 
**ZoneparamKey** | Pointer to **string** | The name of the DNS option you want to add, edit or remove from the zone. You can only set one option at a time.&lt;ul&gt;&lt;li&gt; To add or edit an option: specify its name in the parameter &lt;b&gt;zoneparam_key&lt;/b&gt;, as follows &lt;b&gt;zoneparam_key&#x3D;&amp;lt;option-name&amp;gt;&lt;/b&gt;, and then specify its value in the parameter &lt;b&gt;zoneparam_value&lt;/b&gt;.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; To remove an option, specify its name in the parameter &lt;b&gt;zoneparam_key&lt;/b&gt; and leave the parameter &lt;b&gt;zoneparam_value&lt;/b&gt; empty.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt;To set several options, specify as many parameters (&lt;b&gt;zoneparam_key&lt;/b&gt; and &lt;b&gt;zoneparam_value&lt;/b&gt;) as you need. | [optional] 
**ZoneparamIsArray** | Pointer to **int32** | A way to determine is the DNS zone option is an array (&lt;b&gt;1&lt;/b&gt;). | [optional] 
**ZoneparamValue** | Pointer to **string** | The value of the DNS option specified in the input &lt;b&gt;zoneparam_key&lt;/b&gt;.&lt;ul&gt;&lt;li&gt; To add or edit an option value, specify its name in the parameter &lt;b&gt;zoneparam_key&lt;/b&gt; and set its value as follows: &lt;b&gt;zoneparam_value&#x3D;&amp;lt;option-value&amp;gt;&lt;/b&gt; .&lt;br/&gt;&lt;/li&gt;&lt;li&gt; To remove an option value, specify its name in the parameter &lt;b&gt;zoneparam_key&lt;/b&gt; and leave &lt;b&gt;zoneparam_value&lt;/b&gt; empty: &lt;b&gt;zoneparam_value&#x3D;&lt;/b&gt; .&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 

## Methods

### NewDnsZoneparamEditInput

`func NewDnsZoneparamEditInput() *DnsZoneparamEditInput`

NewDnsZoneparamEditInput instantiates a new DnsZoneparamEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsZoneparamEditInputWithDefaults

`func NewDnsZoneparamEditInputWithDefaults() *DnsZoneparamEditInput`

NewDnsZoneparamEditInputWithDefaults instantiates a new DnsZoneparamEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetZoneId

`func (o *DnsZoneparamEditInput) GetZoneId() int32`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *DnsZoneparamEditInput) GetZoneIdOk() (*int32, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *DnsZoneparamEditInput) SetZoneId(v int32)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *DnsZoneparamEditInput) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetZoneparamKey

`func (o *DnsZoneparamEditInput) GetZoneparamKey() string`

GetZoneparamKey returns the ZoneparamKey field if non-nil, zero value otherwise.

### GetZoneparamKeyOk

`func (o *DnsZoneparamEditInput) GetZoneparamKeyOk() (*string, bool)`

GetZoneparamKeyOk returns a tuple with the ZoneparamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneparamKey

`func (o *DnsZoneparamEditInput) SetZoneparamKey(v string)`

SetZoneparamKey sets ZoneparamKey field to given value.

### HasZoneparamKey

`func (o *DnsZoneparamEditInput) HasZoneparamKey() bool`

HasZoneparamKey returns a boolean if a field has been set.

### GetZoneparamIsArray

`func (o *DnsZoneparamEditInput) GetZoneparamIsArray() int32`

GetZoneparamIsArray returns the ZoneparamIsArray field if non-nil, zero value otherwise.

### GetZoneparamIsArrayOk

`func (o *DnsZoneparamEditInput) GetZoneparamIsArrayOk() (*int32, bool)`

GetZoneparamIsArrayOk returns a tuple with the ZoneparamIsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneparamIsArray

`func (o *DnsZoneparamEditInput) SetZoneparamIsArray(v int32)`

SetZoneparamIsArray sets ZoneparamIsArray field to given value.

### HasZoneparamIsArray

`func (o *DnsZoneparamEditInput) HasZoneparamIsArray() bool`

HasZoneparamIsArray returns a boolean if a field has been set.

### GetZoneparamValue

`func (o *DnsZoneparamEditInput) GetZoneparamValue() string`

GetZoneparamValue returns the ZoneparamValue field if non-nil, zero value otherwise.

### GetZoneparamValueOk

`func (o *DnsZoneparamEditInput) GetZoneparamValueOk() (*string, bool)`

GetZoneparamValueOk returns a tuple with the ZoneparamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneparamValue

`func (o *DnsZoneparamEditInput) SetZoneparamValue(v string)`

SetZoneparamValue sets ZoneparamValue field to given value.

### HasZoneparamValue

`func (o *DnsZoneparamEditInput) HasZoneparamValue() bool`

HasZoneparamValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


