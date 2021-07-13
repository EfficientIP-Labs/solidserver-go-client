# DnsAclAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server. | [optional] 
**AclName** | Pointer to **string** | The name of the DNS ACL, each DNS ACL must have a unique name. | [optional] 
**AclValue** | Pointer to **string** | The values of the DNS ACL in order of priority, as follows: &amp;lt;value_1&amp;gt;;&amp;lt;value_2&amp;gt;... . | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DNS server. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDnsAclAddInput

`func NewDnsAclAddInput() *DnsAclAddInput`

NewDnsAclAddInput instantiates a new DnsAclAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsAclAddInputWithDefaults

`func NewDnsAclAddInputWithDefaults() *DnsAclAddInput`

NewDnsAclAddInputWithDefaults instantiates a new DnsAclAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DnsAclAddInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsAclAddInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsAclAddInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsAclAddInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsAclAddInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsAclAddInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsAclAddInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsAclAddInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetAclName

`func (o *DnsAclAddInput) GetAclName() string`

GetAclName returns the AclName field if non-nil, zero value otherwise.

### GetAclNameOk

`func (o *DnsAclAddInput) GetAclNameOk() (*string, bool)`

GetAclNameOk returns a tuple with the AclName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclName

`func (o *DnsAclAddInput) SetAclName(v string)`

SetAclName sets AclName field to given value.

### HasAclName

`func (o *DnsAclAddInput) HasAclName() bool`

HasAclName returns a boolean if a field has been set.

### GetAclValue

`func (o *DnsAclAddInput) GetAclValue() string`

GetAclValue returns the AclValue field if non-nil, zero value otherwise.

### GetAclValueOk

`func (o *DnsAclAddInput) GetAclValueOk() (*string, bool)`

GetAclValueOk returns a tuple with the AclValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclValue

`func (o *DnsAclAddInput) SetAclValue(v string)`

SetAclValue sets AclValue field to given value.

### HasAclValue

`func (o *DnsAclAddInput) HasAclValue() bool`

HasAclValue returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DnsAclAddInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DnsAclAddInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DnsAclAddInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DnsAclAddInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetWarnings

`func (o *DnsAclAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DnsAclAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DnsAclAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DnsAclAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


