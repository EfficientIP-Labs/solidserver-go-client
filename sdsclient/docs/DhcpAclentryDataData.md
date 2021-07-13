# DhcpAclentryDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclentryDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**AclentryDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**AclId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 ACL. | [optional] 
**AclName** | Pointer to **string** | The name of the DHCPv4 ACL. | [optional] 
**AclentryId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 ACL entry. | [optional] 
**AclentryValue** | Pointer to **string** | The value of the DHCPv4 ACL entry. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpAclentryDataData

`func NewDhcpAclentryDataData() *DhcpAclentryDataData`

NewDhcpAclentryDataData instantiates a new DhcpAclentryDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpAclentryDataDataWithDefaults

`func NewDhcpAclentryDataDataWithDefaults() *DhcpAclentryDataData`

NewDhcpAclentryDataDataWithDefaults instantiates a new DhcpAclentryDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclentryDelayedCreateTime

`func (o *DhcpAclentryDataData) GetAclentryDelayedCreateTime() string`

GetAclentryDelayedCreateTime returns the AclentryDelayedCreateTime field if non-nil, zero value otherwise.

### GetAclentryDelayedCreateTimeOk

`func (o *DhcpAclentryDataData) GetAclentryDelayedCreateTimeOk() (*string, bool)`

GetAclentryDelayedCreateTimeOk returns a tuple with the AclentryDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryDelayedCreateTime

`func (o *DhcpAclentryDataData) SetAclentryDelayedCreateTime(v string)`

SetAclentryDelayedCreateTime sets AclentryDelayedCreateTime field to given value.

### HasAclentryDelayedCreateTime

`func (o *DhcpAclentryDataData) HasAclentryDelayedCreateTime() bool`

HasAclentryDelayedCreateTime returns a boolean if a field has been set.

### GetAclentryDelayedDeleteTime

`func (o *DhcpAclentryDataData) GetAclentryDelayedDeleteTime() string`

GetAclentryDelayedDeleteTime returns the AclentryDelayedDeleteTime field if non-nil, zero value otherwise.

### GetAclentryDelayedDeleteTimeOk

`func (o *DhcpAclentryDataData) GetAclentryDelayedDeleteTimeOk() (*string, bool)`

GetAclentryDelayedDeleteTimeOk returns a tuple with the AclentryDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryDelayedDeleteTime

`func (o *DhcpAclentryDataData) SetAclentryDelayedDeleteTime(v string)`

SetAclentryDelayedDeleteTime sets AclentryDelayedDeleteTime field to given value.

### HasAclentryDelayedDeleteTime

`func (o *DhcpAclentryDataData) HasAclentryDelayedDeleteTime() bool`

HasAclentryDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpAclentryDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpAclentryDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpAclentryDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpAclentryDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpAclentryDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpAclentryDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpAclentryDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpAclentryDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpAclentryDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpAclentryDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpAclentryDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpAclentryDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetAclId

`func (o *DhcpAclentryDataData) GetAclId() string`

GetAclId returns the AclId field if non-nil, zero value otherwise.

### GetAclIdOk

`func (o *DhcpAclentryDataData) GetAclIdOk() (*string, bool)`

GetAclIdOk returns a tuple with the AclId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclId

`func (o *DhcpAclentryDataData) SetAclId(v string)`

SetAclId sets AclId field to given value.

### HasAclId

`func (o *DhcpAclentryDataData) HasAclId() bool`

HasAclId returns a boolean if a field has been set.

### GetAclName

`func (o *DhcpAclentryDataData) GetAclName() string`

GetAclName returns the AclName field if non-nil, zero value otherwise.

### GetAclNameOk

`func (o *DhcpAclentryDataData) GetAclNameOk() (*string, bool)`

GetAclNameOk returns a tuple with the AclName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclName

`func (o *DhcpAclentryDataData) SetAclName(v string)`

SetAclName sets AclName field to given value.

### HasAclName

`func (o *DhcpAclentryDataData) HasAclName() bool`

HasAclName returns a boolean if a field has been set.

### GetAclentryId

`func (o *DhcpAclentryDataData) GetAclentryId() string`

GetAclentryId returns the AclentryId field if non-nil, zero value otherwise.

### GetAclentryIdOk

`func (o *DhcpAclentryDataData) GetAclentryIdOk() (*string, bool)`

GetAclentryIdOk returns a tuple with the AclentryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryId

`func (o *DhcpAclentryDataData) SetAclentryId(v string)`

SetAclentryId sets AclentryId field to given value.

### HasAclentryId

`func (o *DhcpAclentryDataData) HasAclentryId() bool`

HasAclentryId returns a boolean if a field has been set.

### GetAclentryValue

`func (o *DhcpAclentryDataData) GetAclentryValue() string`

GetAclentryValue returns the AclentryValue field if non-nil, zero value otherwise.

### GetAclentryValueOk

`func (o *DhcpAclentryDataData) GetAclentryValueOk() (*string, bool)`

GetAclentryValueOk returns a tuple with the AclentryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryValue

`func (o *DhcpAclentryDataData) SetAclentryValue(v string)`

SetAclentryValue sets AclentryValue field to given value.

### HasAclentryValue

`func (o *DhcpAclentryDataData) HasAclentryValue() bool`

HasAclentryValue returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpAclentryDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpAclentryDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpAclentryDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpAclentryDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


