# DataInnerDhcpAclentry6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerClusterRole** | Pointer to **string** | The role of the server the object belongs to in the cluster, either &lt;b&gt;active (M)&lt;/b&gt;, &lt;b&gt;passive (B)&lt;/b&gt; or &lt;b&gt;N/A (#)&lt;/b&gt;. | [optional] 
**Aclentry6DelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;server6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv6 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Acl6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 ACL. | [optional] 
**Acl6Name** | Pointer to **string** | The name of the DHCPv6 ACL. | [optional] 
**Aclentry6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 ACL entry. | [optional] 
**Aclentry6Value** | Pointer to **string** | The value of the DHCPv6 ACL entry. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpAclentry6Data

`func NewDataInnerDhcpAclentry6Data() *DataInnerDhcpAclentry6Data`

NewDataInnerDhcpAclentry6Data instantiates a new DataInnerDhcpAclentry6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpAclentry6DataWithDefaults

`func NewDataInnerDhcpAclentry6DataWithDefaults() *DataInnerDhcpAclentry6Data`

NewDataInnerDhcpAclentry6DataWithDefaults instantiates a new DataInnerDhcpAclentry6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerClusterRole

`func (o *DataInnerDhcpAclentry6Data) GetServerClusterRole() string`

GetServerClusterRole returns the ServerClusterRole field if non-nil, zero value otherwise.

### GetServerClusterRoleOk

`func (o *DataInnerDhcpAclentry6Data) GetServerClusterRoleOk() (*string, bool)`

GetServerClusterRoleOk returns a tuple with the ServerClusterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClusterRole

`func (o *DataInnerDhcpAclentry6Data) SetServerClusterRole(v string)`

SetServerClusterRole sets ServerClusterRole field to given value.

### HasServerClusterRole

`func (o *DataInnerDhcpAclentry6Data) HasServerClusterRole() bool`

HasServerClusterRole returns a boolean if a field has been set.

### GetAclentry6DelayedTime

`func (o *DataInnerDhcpAclentry6Data) GetAclentry6DelayedTime() string`

GetAclentry6DelayedTime returns the Aclentry6DelayedTime field if non-nil, zero value otherwise.

### GetAclentry6DelayedTimeOk

`func (o *DataInnerDhcpAclentry6Data) GetAclentry6DelayedTimeOk() (*string, bool)`

GetAclentry6DelayedTimeOk returns a tuple with the Aclentry6DelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentry6DelayedTime

`func (o *DataInnerDhcpAclentry6Data) SetAclentry6DelayedTime(v string)`

SetAclentry6DelayedTime sets Aclentry6DelayedTime field to given value.

### HasAclentry6DelayedTime

`func (o *DataInnerDhcpAclentry6Data) HasAclentry6DelayedTime() bool`

HasAclentry6DelayedTime returns a boolean if a field has been set.

### GetServer6Id

`func (o *DataInnerDhcpAclentry6Data) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DataInnerDhcpAclentry6Data) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DataInnerDhcpAclentry6Data) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DataInnerDhcpAclentry6Data) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DataInnerDhcpAclentry6Data) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DataInnerDhcpAclentry6Data) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DataInnerDhcpAclentry6Data) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DataInnerDhcpAclentry6Data) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DataInnerDhcpAclentry6Data) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DataInnerDhcpAclentry6Data) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DataInnerDhcpAclentry6Data) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DataInnerDhcpAclentry6Data) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetAcl6Id

`func (o *DataInnerDhcpAclentry6Data) GetAcl6Id() string`

GetAcl6Id returns the Acl6Id field if non-nil, zero value otherwise.

### GetAcl6IdOk

`func (o *DataInnerDhcpAclentry6Data) GetAcl6IdOk() (*string, bool)`

GetAcl6IdOk returns a tuple with the Acl6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl6Id

`func (o *DataInnerDhcpAclentry6Data) SetAcl6Id(v string)`

SetAcl6Id sets Acl6Id field to given value.

### HasAcl6Id

`func (o *DataInnerDhcpAclentry6Data) HasAcl6Id() bool`

HasAcl6Id returns a boolean if a field has been set.

### GetAcl6Name

`func (o *DataInnerDhcpAclentry6Data) GetAcl6Name() string`

GetAcl6Name returns the Acl6Name field if non-nil, zero value otherwise.

### GetAcl6NameOk

`func (o *DataInnerDhcpAclentry6Data) GetAcl6NameOk() (*string, bool)`

GetAcl6NameOk returns a tuple with the Acl6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcl6Name

`func (o *DataInnerDhcpAclentry6Data) SetAcl6Name(v string)`

SetAcl6Name sets Acl6Name field to given value.

### HasAcl6Name

`func (o *DataInnerDhcpAclentry6Data) HasAcl6Name() bool`

HasAcl6Name returns a boolean if a field has been set.

### GetAclentry6Id

`func (o *DataInnerDhcpAclentry6Data) GetAclentry6Id() string`

GetAclentry6Id returns the Aclentry6Id field if non-nil, zero value otherwise.

### GetAclentry6IdOk

`func (o *DataInnerDhcpAclentry6Data) GetAclentry6IdOk() (*string, bool)`

GetAclentry6IdOk returns a tuple with the Aclentry6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentry6Id

`func (o *DataInnerDhcpAclentry6Data) SetAclentry6Id(v string)`

SetAclentry6Id sets Aclentry6Id field to given value.

### HasAclentry6Id

`func (o *DataInnerDhcpAclentry6Data) HasAclentry6Id() bool`

HasAclentry6Id returns a boolean if a field has been set.

### GetAclentry6Value

`func (o *DataInnerDhcpAclentry6Data) GetAclentry6Value() string`

GetAclentry6Value returns the Aclentry6Value field if non-nil, zero value otherwise.

### GetAclentry6ValueOk

`func (o *DataInnerDhcpAclentry6Data) GetAclentry6ValueOk() (*string, bool)`

GetAclentry6ValueOk returns a tuple with the Aclentry6Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAclentry6Value

`func (o *DataInnerDhcpAclentry6Data) SetAclentry6Value(v string)`

SetAclentry6Value sets Aclentry6Value field to given value.

### HasAclentry6Value

`func (o *DataInnerDhcpAclentry6Data) HasAclentry6Value() bool`

HasAclentry6Value returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpAclentry6Data) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpAclentry6Data) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpAclentry6Data) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpAclentry6Data) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


