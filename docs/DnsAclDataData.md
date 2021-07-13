# DnsAclDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server the object belongs to. | [optional] 
**AclId** | Pointer to **string** | The database identifier (ID) of the DNS ACL. | [optional] 
**AclName** | Pointer to **string** | The name of the DNS ACL. | [optional] 
**AclValue** | Pointer to **string** | The values of the DNS ACL in order of priority, as follows: &amp;lt;value_1&amp;gt;;&amp;lt;value_2&amp;gt;... . | [optional] 

## Methods

### NewDnsAclDataData

`func NewDnsAclDataData() *DnsAclDataData`

NewDnsAclDataData instantiates a new DnsAclDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsAclDataDataWithDefaults

`func NewDnsAclDataDataWithDefaults() *DnsAclDataData`

NewDnsAclDataDataWithDefaults instantiates a new DnsAclDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DnsAclDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsAclDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsAclDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsAclDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetAclId

`func (o *DnsAclDataData) GetAclId() string`

GetAclId returns the AclId field if non-nil, zero value otherwise.

### GetAclIdOk

`func (o *DnsAclDataData) GetAclIdOk() (*string, bool)`

GetAclIdOk returns a tuple with the AclId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclId

`func (o *DnsAclDataData) SetAclId(v string)`

SetAclId sets AclId field to given value.

### HasAclId

`func (o *DnsAclDataData) HasAclId() bool`

HasAclId returns a boolean if a field has been set.

### GetAclName

`func (o *DnsAclDataData) GetAclName() string`

GetAclName returns the AclName field if non-nil, zero value otherwise.

### GetAclNameOk

`func (o *DnsAclDataData) GetAclNameOk() (*string, bool)`

GetAclNameOk returns a tuple with the AclName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclName

`func (o *DnsAclDataData) SetAclName(v string)`

SetAclName sets AclName field to given value.

### HasAclName

`func (o *DnsAclDataData) HasAclName() bool`

HasAclName returns a boolean if a field has been set.

### GetAclValue

`func (o *DnsAclDataData) GetAclValue() string`

GetAclValue returns the AclValue field if non-nil, zero value otherwise.

### GetAclValueOk

`func (o *DnsAclDataData) GetAclValueOk() (*string, bool)`

GetAclValueOk returns a tuple with the AclValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclValue

`func (o *DnsAclDataData) SetAclValue(v string)`

SetAclValue sets AclValue field to given value.

### HasAclValue

`func (o *DnsAclDataData) HasAclValue() bool`

HasAclValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


