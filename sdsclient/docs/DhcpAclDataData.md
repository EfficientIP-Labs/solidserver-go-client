# DhcpAclDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AclDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**AclDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**AclId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 ACL. | [optional] 
**AclMatch** | Pointer to **string** | The ACL rule associated with the DHCPv4 ACL, as follows: &lt;b&gt;&amp;lt;match if (substring(option agent.remote-id,0,6) &#x3D; dslam1);&amp;gt;&lt;/b&gt; | [optional] 
**AclName** | Pointer to **string** | The name of the DHCPv4 ACL. | [optional] 
**AclSpawnwith** | Pointer to **string** | The spawning class associated with the DHCPv4 ACL. | [optional] 
**AclStatement** | Pointer to **string** | The statement associated with the DHCPv4 ACL. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpAclDataData

`func NewDhcpAclDataData() *DhcpAclDataData`

NewDhcpAclDataData instantiates a new DhcpAclDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpAclDataDataWithDefaults

`func NewDhcpAclDataDataWithDefaults() *DhcpAclDataData`

NewDhcpAclDataDataWithDefaults instantiates a new DhcpAclDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAclDelayedCreateTime

`func (o *DhcpAclDataData) GetAclDelayedCreateTime() string`

GetAclDelayedCreateTime returns the AclDelayedCreateTime field if non-nil, zero value otherwise.

### GetAclDelayedCreateTimeOk

`func (o *DhcpAclDataData) GetAclDelayedCreateTimeOk() (*string, bool)`

GetAclDelayedCreateTimeOk returns a tuple with the AclDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclDelayedCreateTime

`func (o *DhcpAclDataData) SetAclDelayedCreateTime(v string)`

SetAclDelayedCreateTime sets AclDelayedCreateTime field to given value.

### HasAclDelayedCreateTime

`func (o *DhcpAclDataData) HasAclDelayedCreateTime() bool`

HasAclDelayedCreateTime returns a boolean if a field has been set.

### GetAclDelayedDeleteTime

`func (o *DhcpAclDataData) GetAclDelayedDeleteTime() string`

GetAclDelayedDeleteTime returns the AclDelayedDeleteTime field if non-nil, zero value otherwise.

### GetAclDelayedDeleteTimeOk

`func (o *DhcpAclDataData) GetAclDelayedDeleteTimeOk() (*string, bool)`

GetAclDelayedDeleteTimeOk returns a tuple with the AclDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclDelayedDeleteTime

`func (o *DhcpAclDataData) SetAclDelayedDeleteTime(v string)`

SetAclDelayedDeleteTime sets AclDelayedDeleteTime field to given value.

### HasAclDelayedDeleteTime

`func (o *DhcpAclDataData) HasAclDelayedDeleteTime() bool`

HasAclDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpAclDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpAclDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpAclDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpAclDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpAclDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpAclDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpAclDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpAclDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpAclDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpAclDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpAclDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpAclDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetAclId

`func (o *DhcpAclDataData) GetAclId() string`

GetAclId returns the AclId field if non-nil, zero value otherwise.

### GetAclIdOk

`func (o *DhcpAclDataData) GetAclIdOk() (*string, bool)`

GetAclIdOk returns a tuple with the AclId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclId

`func (o *DhcpAclDataData) SetAclId(v string)`

SetAclId sets AclId field to given value.

### HasAclId

`func (o *DhcpAclDataData) HasAclId() bool`

HasAclId returns a boolean if a field has been set.

### GetAclMatch

`func (o *DhcpAclDataData) GetAclMatch() string`

GetAclMatch returns the AclMatch field if non-nil, zero value otherwise.

### GetAclMatchOk

`func (o *DhcpAclDataData) GetAclMatchOk() (*string, bool)`

GetAclMatchOk returns a tuple with the AclMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclMatch

`func (o *DhcpAclDataData) SetAclMatch(v string)`

SetAclMatch sets AclMatch field to given value.

### HasAclMatch

`func (o *DhcpAclDataData) HasAclMatch() bool`

HasAclMatch returns a boolean if a field has been set.

### GetAclName

`func (o *DhcpAclDataData) GetAclName() string`

GetAclName returns the AclName field if non-nil, zero value otherwise.

### GetAclNameOk

`func (o *DhcpAclDataData) GetAclNameOk() (*string, bool)`

GetAclNameOk returns a tuple with the AclName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclName

`func (o *DhcpAclDataData) SetAclName(v string)`

SetAclName sets AclName field to given value.

### HasAclName

`func (o *DhcpAclDataData) HasAclName() bool`

HasAclName returns a boolean if a field has been set.

### GetAclSpawnwith

`func (o *DhcpAclDataData) GetAclSpawnwith() string`

GetAclSpawnwith returns the AclSpawnwith field if non-nil, zero value otherwise.

### GetAclSpawnwithOk

`func (o *DhcpAclDataData) GetAclSpawnwithOk() (*string, bool)`

GetAclSpawnwithOk returns a tuple with the AclSpawnwith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclSpawnwith

`func (o *DhcpAclDataData) SetAclSpawnwith(v string)`

SetAclSpawnwith sets AclSpawnwith field to given value.

### HasAclSpawnwith

`func (o *DhcpAclDataData) HasAclSpawnwith() bool`

HasAclSpawnwith returns a boolean if a field has been set.

### GetAclStatement

`func (o *DhcpAclDataData) GetAclStatement() string`

GetAclStatement returns the AclStatement field if non-nil, zero value otherwise.

### GetAclStatementOk

`func (o *DhcpAclDataData) GetAclStatementOk() (*string, bool)`

GetAclStatementOk returns a tuple with the AclStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclStatement

`func (o *DhcpAclDataData) SetAclStatement(v string)`

SetAclStatement sets AclStatement field to given value.

### HasAclStatement

`func (o *DhcpAclDataData) HasAclStatement() bool`

HasAclStatement returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpAclDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpAclDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpAclDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpAclDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


