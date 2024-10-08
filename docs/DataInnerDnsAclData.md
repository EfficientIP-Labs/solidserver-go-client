# DataInnerDnsAclData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server the object belongs to. | [optional] 
**AclId** | Pointer to **string** | The database identifier (ID) of the DNS ACL. | [optional] 
**AclName** | Pointer to **string** | The name of the DNS ACL. | [optional] 
**AclValue** | Pointer to **string** | The values of the DNS ACL in order of priority, as follows: &amp;lt;value_1&amp;gt;;&amp;lt;value_2&amp;gt;... . | [optional] 

## Methods

### NewDataInnerDnsAclData

`func NewDataInnerDnsAclData() *DataInnerDnsAclData`

NewDataInnerDnsAclData instantiates a new DataInnerDnsAclData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDnsAclDataWithDefaults

`func NewDataInnerDnsAclDataWithDefaults() *DataInnerDnsAclData`

NewDataInnerDnsAclDataWithDefaults instantiates a new DataInnerDnsAclData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DataInnerDnsAclData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDnsAclData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDnsAclData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDnsAclData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetAclId

`func (o *DataInnerDnsAclData) GetAclId() string`

GetAclId returns the AclId field if non-nil, zero value otherwise.

### GetAclIdOk

`func (o *DataInnerDnsAclData) GetAclIdOk() (*string, bool)`

GetAclIdOk returns a tuple with the AclId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclId

`func (o *DataInnerDnsAclData) SetAclId(v string)`

SetAclId sets AclId field to given value.

### HasAclId

`func (o *DataInnerDnsAclData) HasAclId() bool`

HasAclId returns a boolean if a field has been set.

### GetAclName

`func (o *DataInnerDnsAclData) GetAclName() string`

GetAclName returns the AclName field if non-nil, zero value otherwise.

### GetAclNameOk

`func (o *DataInnerDnsAclData) GetAclNameOk() (*string, bool)`

GetAclNameOk returns a tuple with the AclName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclName

`func (o *DataInnerDnsAclData) SetAclName(v string)`

SetAclName sets AclName field to given value.

### HasAclName

`func (o *DataInnerDnsAclData) HasAclName() bool`

HasAclName returns a boolean if a field has been set.

### GetAclValue

`func (o *DataInnerDnsAclData) GetAclValue() string`

GetAclValue returns the AclValue field if non-nil, zero value otherwise.

### GetAclValueOk

`func (o *DataInnerDnsAclData) GetAclValueOk() (*string, bool)`

GetAclValueOk returns a tuple with the AclValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclValue

`func (o *DataInnerDnsAclData) SetAclValue(v string)`

SetAclValue sets AclValue field to given value.

### HasAclValue

`func (o *DataInnerDnsAclData) HasAclValue() bool`

HasAclValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


