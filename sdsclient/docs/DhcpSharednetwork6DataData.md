# DhcpSharednetwork6DataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Sharednetwork6DelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Sharednetwork6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 shared network. | [optional] 
**Sharednetwork6Name** | Pointer to **string** | The name of the DHCPv6 shared network. | [optional] 
**Sharednetwork6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**RowState** | Pointer to **string** | The object activation status: &lt;ul class&#x3D;dashed &gt;&lt;li&gt;                                                If set to &lt;b&gt;0&lt;/b&gt;, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                If set to &lt;b&gt;1&lt;/b&gt;, the object is enabled and managed.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_enabled&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 
**SmartArch** | Pointer to **string** | The type of the DHCPv6 smart architecture the object belongs to.&lt;table&gt;&lt;caption&gt;vdhcp6_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;The Single-Server smart architecture manages a single DHCPv6 server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;splitscope&lt;/td&gt;&lt;td &gt;The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;stateless&lt;/td&gt;&lt;td &gt;The Stateless smart architecture offers a limited number of options to the DHCP clients. The IP address is delivered thanks to the subnet gateway and it is impossible to create any ranges or statics or to retrieve any leases.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpSharednetwork6DataData

`func NewDhcpSharednetwork6DataData() *DhcpSharednetwork6DataData`

NewDhcpSharednetwork6DataData instantiates a new DhcpSharednetwork6DataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSharednetwork6DataDataWithDefaults

`func NewDhcpSharednetwork6DataDataWithDefaults() *DhcpSharednetwork6DataData`

NewDhcpSharednetwork6DataDataWithDefaults instantiates a new DhcpSharednetwork6DataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharednetwork6DelayedTime

`func (o *DhcpSharednetwork6DataData) GetSharednetwork6DelayedTime() string`

GetSharednetwork6DelayedTime returns the Sharednetwork6DelayedTime field if non-nil, zero value otherwise.

### GetSharednetwork6DelayedTimeOk

`func (o *DhcpSharednetwork6DataData) GetSharednetwork6DelayedTimeOk() (*string, bool)`

GetSharednetwork6DelayedTimeOk returns a tuple with the Sharednetwork6DelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6DelayedTime

`func (o *DhcpSharednetwork6DataData) SetSharednetwork6DelayedTime(v string)`

SetSharednetwork6DelayedTime sets Sharednetwork6DelayedTime field to given value.

### HasSharednetwork6DelayedTime

`func (o *DhcpSharednetwork6DataData) HasSharednetwork6DelayedTime() bool`

HasSharednetwork6DelayedTime returns a boolean if a field has been set.

### GetServer6Id

`func (o *DhcpSharednetwork6DataData) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DhcpSharednetwork6DataData) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DhcpSharednetwork6DataData) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DhcpSharednetwork6DataData) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DhcpSharednetwork6DataData) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DhcpSharednetwork6DataData) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DhcpSharednetwork6DataData) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DhcpSharednetwork6DataData) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DhcpSharednetwork6DataData) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DhcpSharednetwork6DataData) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DhcpSharednetwork6DataData) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DhcpSharednetwork6DataData) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetSharednetwork6Id

`func (o *DhcpSharednetwork6DataData) GetSharednetwork6Id() string`

GetSharednetwork6Id returns the Sharednetwork6Id field if non-nil, zero value otherwise.

### GetSharednetwork6IdOk

`func (o *DhcpSharednetwork6DataData) GetSharednetwork6IdOk() (*string, bool)`

GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Id

`func (o *DhcpSharednetwork6DataData) SetSharednetwork6Id(v string)`

SetSharednetwork6Id sets Sharednetwork6Id field to given value.

### HasSharednetwork6Id

`func (o *DhcpSharednetwork6DataData) HasSharednetwork6Id() bool`

HasSharednetwork6Id returns a boolean if a field has been set.

### GetSharednetwork6Name

`func (o *DhcpSharednetwork6DataData) GetSharednetwork6Name() string`

GetSharednetwork6Name returns the Sharednetwork6Name field if non-nil, zero value otherwise.

### GetSharednetwork6NameOk

`func (o *DhcpSharednetwork6DataData) GetSharednetwork6NameOk() (*string, bool)`

GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Name

`func (o *DhcpSharednetwork6DataData) SetSharednetwork6Name(v string)`

SetSharednetwork6Name sets Sharednetwork6Name field to given value.

### HasSharednetwork6Name

`func (o *DhcpSharednetwork6DataData) HasSharednetwork6Name() bool`

HasSharednetwork6Name returns a boolean if a field has been set.

### GetSharednetwork6Multistatus

`func (o *DhcpSharednetwork6DataData) GetSharednetwork6Multistatus() string`

GetSharednetwork6Multistatus returns the Sharednetwork6Multistatus field if non-nil, zero value otherwise.

### GetSharednetwork6MultistatusOk

`func (o *DhcpSharednetwork6DataData) GetSharednetwork6MultistatusOk() (*string, bool)`

GetSharednetwork6MultistatusOk returns a tuple with the Sharednetwork6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Multistatus

`func (o *DhcpSharednetwork6DataData) SetSharednetwork6Multistatus(v string)`

SetSharednetwork6Multistatus sets Sharednetwork6Multistatus field to given value.

### HasSharednetwork6Multistatus

`func (o *DhcpSharednetwork6DataData) HasSharednetwork6Multistatus() bool`

HasSharednetwork6Multistatus returns a boolean if a field has been set.

### GetRowState

`func (o *DhcpSharednetwork6DataData) GetRowState() string`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *DhcpSharednetwork6DataData) GetRowStateOk() (*string, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *DhcpSharednetwork6DataData) SetRowState(v string)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *DhcpSharednetwork6DataData) HasRowState() bool`

HasRowState returns a boolean if a field has been set.

### GetSmartArch

`func (o *DhcpSharednetwork6DataData) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DhcpSharednetwork6DataData) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DhcpSharednetwork6DataData) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DhcpSharednetwork6DataData) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpSharednetwork6DataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpSharednetwork6DataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpSharednetwork6DataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpSharednetwork6DataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


