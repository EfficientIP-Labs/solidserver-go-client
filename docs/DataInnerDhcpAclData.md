# DataInnerDhcpAclData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerClusterRole** | Pointer to **string** | The role of the server the object belongs to in the cluster, either &lt;b&gt;active (M)&lt;/b&gt;, &lt;b&gt;passive (B)&lt;/b&gt; or &lt;b&gt;N/A (#)&lt;/b&gt;. | [optional] 
**AclDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**AclDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server the object belongs to. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server the object belongs to: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft Windows DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv4 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**AclId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 ACL. | [optional] 
**AclLeaselimit** | Pointer to **string** | The lease limit of the DHCPv4 ACL. | [optional] 
**AclMatch** | Pointer to **string** | The ACL rule associated with the DHCPv4 ACL, as follows: &lt;b&gt;&amp;lt;match if (substring(option agent.remote-id,0,6) &#x3D; dslam1);&amp;gt;&lt;/b&gt; | [optional] 
**AclName** | Pointer to **string** | The name of the DHCPv4 ACL. | [optional] 
**AclSpawnwith** | Pointer to **string** | The spawning class associated with the DHCPv4 ACL. | [optional] 
**AclStatement** | Pointer to **string** | The statement associated with the DHCPv4 ACL. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpAclData

`func NewDataInnerDhcpAclData() *DataInnerDhcpAclData`

NewDataInnerDhcpAclData instantiates a new DataInnerDhcpAclData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpAclDataWithDefaults

`func NewDataInnerDhcpAclDataWithDefaults() *DataInnerDhcpAclData`

NewDataInnerDhcpAclDataWithDefaults instantiates a new DataInnerDhcpAclData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerClusterRole

`func (o *DataInnerDhcpAclData) GetServerClusterRole() string`

GetServerClusterRole returns the ServerClusterRole field if non-nil, zero value otherwise.

### GetServerClusterRoleOk

`func (o *DataInnerDhcpAclData) GetServerClusterRoleOk() (*string, bool)`

GetServerClusterRoleOk returns a tuple with the ServerClusterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClusterRole

`func (o *DataInnerDhcpAclData) SetServerClusterRole(v string)`

SetServerClusterRole sets ServerClusterRole field to given value.

### HasServerClusterRole

`func (o *DataInnerDhcpAclData) HasServerClusterRole() bool`

HasServerClusterRole returns a boolean if a field has been set.

### GetAclDelayedCreateTime

`func (o *DataInnerDhcpAclData) GetAclDelayedCreateTime() string`

GetAclDelayedCreateTime returns the AclDelayedCreateTime field if non-nil, zero value otherwise.

### GetAclDelayedCreateTimeOk

`func (o *DataInnerDhcpAclData) GetAclDelayedCreateTimeOk() (*string, bool)`

GetAclDelayedCreateTimeOk returns a tuple with the AclDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclDelayedCreateTime

`func (o *DataInnerDhcpAclData) SetAclDelayedCreateTime(v string)`

SetAclDelayedCreateTime sets AclDelayedCreateTime field to given value.

### HasAclDelayedCreateTime

`func (o *DataInnerDhcpAclData) HasAclDelayedCreateTime() bool`

HasAclDelayedCreateTime returns a boolean if a field has been set.

### GetAclDelayedDeleteTime

`func (o *DataInnerDhcpAclData) GetAclDelayedDeleteTime() string`

GetAclDelayedDeleteTime returns the AclDelayedDeleteTime field if non-nil, zero value otherwise.

### GetAclDelayedDeleteTimeOk

`func (o *DataInnerDhcpAclData) GetAclDelayedDeleteTimeOk() (*string, bool)`

GetAclDelayedDeleteTimeOk returns a tuple with the AclDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclDelayedDeleteTime

`func (o *DataInnerDhcpAclData) SetAclDelayedDeleteTime(v string)`

SetAclDelayedDeleteTime sets AclDelayedDeleteTime field to given value.

### HasAclDelayedDeleteTime

`func (o *DataInnerDhcpAclData) HasAclDelayedDeleteTime() bool`

HasAclDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDhcpAclData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDhcpAclData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDhcpAclData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDhcpAclData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDhcpAclData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDhcpAclData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDhcpAclData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDhcpAclData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDhcpAclData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDhcpAclData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDhcpAclData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDhcpAclData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetAclId

`func (o *DataInnerDhcpAclData) GetAclId() string`

GetAclId returns the AclId field if non-nil, zero value otherwise.

### GetAclIdOk

`func (o *DataInnerDhcpAclData) GetAclIdOk() (*string, bool)`

GetAclIdOk returns a tuple with the AclId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclId

`func (o *DataInnerDhcpAclData) SetAclId(v string)`

SetAclId sets AclId field to given value.

### HasAclId

`func (o *DataInnerDhcpAclData) HasAclId() bool`

HasAclId returns a boolean if a field has been set.

### GetAclLeaselimit

`func (o *DataInnerDhcpAclData) GetAclLeaselimit() string`

GetAclLeaselimit returns the AclLeaselimit field if non-nil, zero value otherwise.

### GetAclLeaselimitOk

`func (o *DataInnerDhcpAclData) GetAclLeaselimitOk() (*string, bool)`

GetAclLeaselimitOk returns a tuple with the AclLeaselimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclLeaselimit

`func (o *DataInnerDhcpAclData) SetAclLeaselimit(v string)`

SetAclLeaselimit sets AclLeaselimit field to given value.

### HasAclLeaselimit

`func (o *DataInnerDhcpAclData) HasAclLeaselimit() bool`

HasAclLeaselimit returns a boolean if a field has been set.

### GetAclMatch

`func (o *DataInnerDhcpAclData) GetAclMatch() string`

GetAclMatch returns the AclMatch field if non-nil, zero value otherwise.

### GetAclMatchOk

`func (o *DataInnerDhcpAclData) GetAclMatchOk() (*string, bool)`

GetAclMatchOk returns a tuple with the AclMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclMatch

`func (o *DataInnerDhcpAclData) SetAclMatch(v string)`

SetAclMatch sets AclMatch field to given value.

### HasAclMatch

`func (o *DataInnerDhcpAclData) HasAclMatch() bool`

HasAclMatch returns a boolean if a field has been set.

### GetAclName

`func (o *DataInnerDhcpAclData) GetAclName() string`

GetAclName returns the AclName field if non-nil, zero value otherwise.

### GetAclNameOk

`func (o *DataInnerDhcpAclData) GetAclNameOk() (*string, bool)`

GetAclNameOk returns a tuple with the AclName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclName

`func (o *DataInnerDhcpAclData) SetAclName(v string)`

SetAclName sets AclName field to given value.

### HasAclName

`func (o *DataInnerDhcpAclData) HasAclName() bool`

HasAclName returns a boolean if a field has been set.

### GetAclSpawnwith

`func (o *DataInnerDhcpAclData) GetAclSpawnwith() string`

GetAclSpawnwith returns the AclSpawnwith field if non-nil, zero value otherwise.

### GetAclSpawnwithOk

`func (o *DataInnerDhcpAclData) GetAclSpawnwithOk() (*string, bool)`

GetAclSpawnwithOk returns a tuple with the AclSpawnwith field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclSpawnwith

`func (o *DataInnerDhcpAclData) SetAclSpawnwith(v string)`

SetAclSpawnwith sets AclSpawnwith field to given value.

### HasAclSpawnwith

`func (o *DataInnerDhcpAclData) HasAclSpawnwith() bool`

HasAclSpawnwith returns a boolean if a field has been set.

### GetAclStatement

`func (o *DataInnerDhcpAclData) GetAclStatement() string`

GetAclStatement returns the AclStatement field if non-nil, zero value otherwise.

### GetAclStatementOk

`func (o *DataInnerDhcpAclData) GetAclStatementOk() (*string, bool)`

GetAclStatementOk returns a tuple with the AclStatement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclStatement

`func (o *DataInnerDhcpAclData) SetAclStatement(v string)`

SetAclStatement sets AclStatement field to given value.

### HasAclStatement

`func (o *DataInnerDhcpAclData) HasAclStatement() bool`

HasAclStatement returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpAclData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpAclData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpAclData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpAclData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


