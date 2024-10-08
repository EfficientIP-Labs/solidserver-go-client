# DataInnerDhcpAclentryData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerClusterRole** | Pointer to **string** | The role of the server the object belongs to in the cluster, either &lt;b&gt;active (M)&lt;/b&gt;, &lt;b&gt;passive (B)&lt;/b&gt; or &lt;b&gt;N/A (#)&lt;/b&gt;. | [optional] 
**AclentryDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**AclentryDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft Windows DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv4 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**AclId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 ACL. | [optional] 
**AclName** | Pointer to **string** | The name of the DHCPv4 ACL. | [optional] 
**AclentryId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 ACL entry. | [optional] 
**AclentryLeaselimit** | Pointer to **string** | The lease limit of the DHCPv4 ACL entry. | [optional] 
**AclentryValue** | Pointer to **string** | The value of the DHCPv4 ACL entry. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpAclentryData

`func NewDataInnerDhcpAclentryData() *DataInnerDhcpAclentryData`

NewDataInnerDhcpAclentryData instantiates a new DataInnerDhcpAclentryData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpAclentryDataWithDefaults

`func NewDataInnerDhcpAclentryDataWithDefaults() *DataInnerDhcpAclentryData`

NewDataInnerDhcpAclentryDataWithDefaults instantiates a new DataInnerDhcpAclentryData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerClusterRole

`func (o *DataInnerDhcpAclentryData) GetServerClusterRole() string`

GetServerClusterRole returns the ServerClusterRole field if non-nil, zero value otherwise.

### GetServerClusterRoleOk

`func (o *DataInnerDhcpAclentryData) GetServerClusterRoleOk() (*string, bool)`

GetServerClusterRoleOk returns a tuple with the ServerClusterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClusterRole

`func (o *DataInnerDhcpAclentryData) SetServerClusterRole(v string)`

SetServerClusterRole sets ServerClusterRole field to given value.

### HasServerClusterRole

`func (o *DataInnerDhcpAclentryData) HasServerClusterRole() bool`

HasServerClusterRole returns a boolean if a field has been set.

### GetAclentryDelayedCreateTime

`func (o *DataInnerDhcpAclentryData) GetAclentryDelayedCreateTime() string`

GetAclentryDelayedCreateTime returns the AclentryDelayedCreateTime field if non-nil, zero value otherwise.

### GetAclentryDelayedCreateTimeOk

`func (o *DataInnerDhcpAclentryData) GetAclentryDelayedCreateTimeOk() (*string, bool)`

GetAclentryDelayedCreateTimeOk returns a tuple with the AclentryDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryDelayedCreateTime

`func (o *DataInnerDhcpAclentryData) SetAclentryDelayedCreateTime(v string)`

SetAclentryDelayedCreateTime sets AclentryDelayedCreateTime field to given value.

### HasAclentryDelayedCreateTime

`func (o *DataInnerDhcpAclentryData) HasAclentryDelayedCreateTime() bool`

HasAclentryDelayedCreateTime returns a boolean if a field has been set.

### GetAclentryDelayedDeleteTime

`func (o *DataInnerDhcpAclentryData) GetAclentryDelayedDeleteTime() string`

GetAclentryDelayedDeleteTime returns the AclentryDelayedDeleteTime field if non-nil, zero value otherwise.

### GetAclentryDelayedDeleteTimeOk

`func (o *DataInnerDhcpAclentryData) GetAclentryDelayedDeleteTimeOk() (*string, bool)`

GetAclentryDelayedDeleteTimeOk returns a tuple with the AclentryDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryDelayedDeleteTime

`func (o *DataInnerDhcpAclentryData) SetAclentryDelayedDeleteTime(v string)`

SetAclentryDelayedDeleteTime sets AclentryDelayedDeleteTime field to given value.

### HasAclentryDelayedDeleteTime

`func (o *DataInnerDhcpAclentryData) HasAclentryDelayedDeleteTime() bool`

HasAclentryDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDhcpAclentryData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDhcpAclentryData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDhcpAclentryData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDhcpAclentryData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDhcpAclentryData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDhcpAclentryData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDhcpAclentryData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDhcpAclentryData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDhcpAclentryData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDhcpAclentryData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDhcpAclentryData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDhcpAclentryData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetAclId

`func (o *DataInnerDhcpAclentryData) GetAclId() string`

GetAclId returns the AclId field if non-nil, zero value otherwise.

### GetAclIdOk

`func (o *DataInnerDhcpAclentryData) GetAclIdOk() (*string, bool)`

GetAclIdOk returns a tuple with the AclId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclId

`func (o *DataInnerDhcpAclentryData) SetAclId(v string)`

SetAclId sets AclId field to given value.

### HasAclId

`func (o *DataInnerDhcpAclentryData) HasAclId() bool`

HasAclId returns a boolean if a field has been set.

### GetAclName

`func (o *DataInnerDhcpAclentryData) GetAclName() string`

GetAclName returns the AclName field if non-nil, zero value otherwise.

### GetAclNameOk

`func (o *DataInnerDhcpAclentryData) GetAclNameOk() (*string, bool)`

GetAclNameOk returns a tuple with the AclName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclName

`func (o *DataInnerDhcpAclentryData) SetAclName(v string)`

SetAclName sets AclName field to given value.

### HasAclName

`func (o *DataInnerDhcpAclentryData) HasAclName() bool`

HasAclName returns a boolean if a field has been set.

### GetAclentryId

`func (o *DataInnerDhcpAclentryData) GetAclentryId() string`

GetAclentryId returns the AclentryId field if non-nil, zero value otherwise.

### GetAclentryIdOk

`func (o *DataInnerDhcpAclentryData) GetAclentryIdOk() (*string, bool)`

GetAclentryIdOk returns a tuple with the AclentryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryId

`func (o *DataInnerDhcpAclentryData) SetAclentryId(v string)`

SetAclentryId sets AclentryId field to given value.

### HasAclentryId

`func (o *DataInnerDhcpAclentryData) HasAclentryId() bool`

HasAclentryId returns a boolean if a field has been set.

### GetAclentryLeaselimit

`func (o *DataInnerDhcpAclentryData) GetAclentryLeaselimit() string`

GetAclentryLeaselimit returns the AclentryLeaselimit field if non-nil, zero value otherwise.

### GetAclentryLeaselimitOk

`func (o *DataInnerDhcpAclentryData) GetAclentryLeaselimitOk() (*string, bool)`

GetAclentryLeaselimitOk returns a tuple with the AclentryLeaselimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryLeaselimit

`func (o *DataInnerDhcpAclentryData) SetAclentryLeaselimit(v string)`

SetAclentryLeaselimit sets AclentryLeaselimit field to given value.

### HasAclentryLeaselimit

`func (o *DataInnerDhcpAclentryData) HasAclentryLeaselimit() bool`

HasAclentryLeaselimit returns a boolean if a field has been set.

### GetAclentryValue

`func (o *DataInnerDhcpAclentryData) GetAclentryValue() string`

GetAclentryValue returns the AclentryValue field if non-nil, zero value otherwise.

### GetAclentryValueOk

`func (o *DataInnerDhcpAclentryData) GetAclentryValueOk() (*string, bool)`

GetAclentryValueOk returns a tuple with the AclentryValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentryValue

`func (o *DataInnerDhcpAclentryData) SetAclentryValue(v string)`

SetAclentryValue sets AclentryValue field to given value.

### HasAclentryValue

`func (o *DataInnerDhcpAclentryData) HasAclentryValue() bool`

HasAclentryValue returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpAclentryData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpAclentryData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpAclentryData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpAclentryData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


