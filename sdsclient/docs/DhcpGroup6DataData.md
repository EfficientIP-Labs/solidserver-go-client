# DhcpGroup6DataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group6DelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**Server6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 server the object belongs to, it can be preceded by the class directory. | [optional] 
**Server6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 server the object belongs to and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;dhcp6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Version** | Pointer to **string** | The version details of the DHCPv6 server the object belongs to. | [optional] 
**Group6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 group, it can be preceded by the class directory. | [optional] 
**Group6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 group and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Group6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 group. | [optional] 
**Group6Name** | Pointer to **string** | The name of the DHCPv6 group. | [optional] 
**Server6Addr** | Pointer to **string** | The IP address of the DHCP server the object belongs to, in hexadecimal format. | [optional] 
**Group6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDhcpGroup6DataData

`func NewDhcpGroup6DataData() *DhcpGroup6DataData`

NewDhcpGroup6DataData instantiates a new DhcpGroup6DataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpGroup6DataDataWithDefaults

`func NewDhcpGroup6DataDataWithDefaults() *DhcpGroup6DataData`

NewDhcpGroup6DataDataWithDefaults instantiates a new DhcpGroup6DataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup6DelayedTime

`func (o *DhcpGroup6DataData) GetGroup6DelayedTime() string`

GetGroup6DelayedTime returns the Group6DelayedTime field if non-nil, zero value otherwise.

### GetGroup6DelayedTimeOk

`func (o *DhcpGroup6DataData) GetGroup6DelayedTimeOk() (*string, bool)`

GetGroup6DelayedTimeOk returns a tuple with the Group6DelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6DelayedTime

`func (o *DhcpGroup6DataData) SetGroup6DelayedTime(v string)`

SetGroup6DelayedTime sets Group6DelayedTime field to given value.

### HasGroup6DelayedTime

`func (o *DhcpGroup6DataData) HasGroup6DelayedTime() bool`

HasGroup6DelayedTime returns a boolean if a field has been set.

### GetServer6ClassName

`func (o *DhcpGroup6DataData) GetServer6ClassName() string`

GetServer6ClassName returns the Server6ClassName field if non-nil, zero value otherwise.

### GetServer6ClassNameOk

`func (o *DhcpGroup6DataData) GetServer6ClassNameOk() (*string, bool)`

GetServer6ClassNameOk returns a tuple with the Server6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassName

`func (o *DhcpGroup6DataData) SetServer6ClassName(v string)`

SetServer6ClassName sets Server6ClassName field to given value.

### HasServer6ClassName

`func (o *DhcpGroup6DataData) HasServer6ClassName() bool`

HasServer6ClassName returns a boolean if a field has been set.

### GetServer6ClassParameters

`func (o *DhcpGroup6DataData) GetServer6ClassParameters() []ApiClassParameterOutputEntry`

GetServer6ClassParameters returns the Server6ClassParameters field if non-nil, zero value otherwise.

### GetServer6ClassParametersOk

`func (o *DhcpGroup6DataData) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassParameters

`func (o *DhcpGroup6DataData) SetServer6ClassParameters(v []ApiClassParameterOutputEntry)`

SetServer6ClassParameters sets Server6ClassParameters field to given value.

### HasServer6ClassParameters

`func (o *DhcpGroup6DataData) HasServer6ClassParameters() bool`

HasServer6ClassParameters returns a boolean if a field has been set.

### GetServer6Id

`func (o *DhcpGroup6DataData) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DhcpGroup6DataData) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DhcpGroup6DataData) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DhcpGroup6DataData) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DhcpGroup6DataData) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DhcpGroup6DataData) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DhcpGroup6DataData) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DhcpGroup6DataData) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DhcpGroup6DataData) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DhcpGroup6DataData) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DhcpGroup6DataData) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DhcpGroup6DataData) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetServer6Version

`func (o *DhcpGroup6DataData) GetServer6Version() string`

GetServer6Version returns the Server6Version field if non-nil, zero value otherwise.

### GetServer6VersionOk

`func (o *DhcpGroup6DataData) GetServer6VersionOk() (*string, bool)`

GetServer6VersionOk returns a tuple with the Server6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Version

`func (o *DhcpGroup6DataData) SetServer6Version(v string)`

SetServer6Version sets Server6Version field to given value.

### HasServer6Version

`func (o *DhcpGroup6DataData) HasServer6Version() bool`

HasServer6Version returns a boolean if a field has been set.

### GetGroup6ClassName

`func (o *DhcpGroup6DataData) GetGroup6ClassName() string`

GetGroup6ClassName returns the Group6ClassName field if non-nil, zero value otherwise.

### GetGroup6ClassNameOk

`func (o *DhcpGroup6DataData) GetGroup6ClassNameOk() (*string, bool)`

GetGroup6ClassNameOk returns a tuple with the Group6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6ClassName

`func (o *DhcpGroup6DataData) SetGroup6ClassName(v string)`

SetGroup6ClassName sets Group6ClassName field to given value.

### HasGroup6ClassName

`func (o *DhcpGroup6DataData) HasGroup6ClassName() bool`

HasGroup6ClassName returns a boolean if a field has been set.

### GetGroup6ClassParameters

`func (o *DhcpGroup6DataData) GetGroup6ClassParameters() []ApiClassParameterOutputEntry`

GetGroup6ClassParameters returns the Group6ClassParameters field if non-nil, zero value otherwise.

### GetGroup6ClassParametersOk

`func (o *DhcpGroup6DataData) GetGroup6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetGroup6ClassParametersOk returns a tuple with the Group6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6ClassParameters

`func (o *DhcpGroup6DataData) SetGroup6ClassParameters(v []ApiClassParameterOutputEntry)`

SetGroup6ClassParameters sets Group6ClassParameters field to given value.

### HasGroup6ClassParameters

`func (o *DhcpGroup6DataData) HasGroup6ClassParameters() bool`

HasGroup6ClassParameters returns a boolean if a field has been set.

### GetGroup6Id

`func (o *DhcpGroup6DataData) GetGroup6Id() string`

GetGroup6Id returns the Group6Id field if non-nil, zero value otherwise.

### GetGroup6IdOk

`func (o *DhcpGroup6DataData) GetGroup6IdOk() (*string, bool)`

GetGroup6IdOk returns a tuple with the Group6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Id

`func (o *DhcpGroup6DataData) SetGroup6Id(v string)`

SetGroup6Id sets Group6Id field to given value.

### HasGroup6Id

`func (o *DhcpGroup6DataData) HasGroup6Id() bool`

HasGroup6Id returns a boolean if a field has been set.

### GetGroup6Name

`func (o *DhcpGroup6DataData) GetGroup6Name() string`

GetGroup6Name returns the Group6Name field if non-nil, zero value otherwise.

### GetGroup6NameOk

`func (o *DhcpGroup6DataData) GetGroup6NameOk() (*string, bool)`

GetGroup6NameOk returns a tuple with the Group6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Name

`func (o *DhcpGroup6DataData) SetGroup6Name(v string)`

SetGroup6Name sets Group6Name field to given value.

### HasGroup6Name

`func (o *DhcpGroup6DataData) HasGroup6Name() bool`

HasGroup6Name returns a boolean if a field has been set.

### GetServer6Addr

`func (o *DhcpGroup6DataData) GetServer6Addr() string`

GetServer6Addr returns the Server6Addr field if non-nil, zero value otherwise.

### GetServer6AddrOk

`func (o *DhcpGroup6DataData) GetServer6AddrOk() (*string, bool)`

GetServer6AddrOk returns a tuple with the Server6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr

`func (o *DhcpGroup6DataData) SetServer6Addr(v string)`

SetServer6Addr sets Server6Addr field to given value.

### HasServer6Addr

`func (o *DhcpGroup6DataData) HasServer6Addr() bool`

HasServer6Addr returns a boolean if a field has been set.

### GetGroup6Multistatus

`func (o *DhcpGroup6DataData) GetGroup6Multistatus() string`

GetGroup6Multistatus returns the Group6Multistatus field if non-nil, zero value otherwise.

### GetGroup6MultistatusOk

`func (o *DhcpGroup6DataData) GetGroup6MultistatusOk() (*string, bool)`

GetGroup6MultistatusOk returns a tuple with the Group6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Multistatus

`func (o *DhcpGroup6DataData) SetGroup6Multistatus(v string)`

SetGroup6Multistatus sets Group6Multistatus field to given value.

### HasGroup6Multistatus

`func (o *DhcpGroup6DataData) HasGroup6Multistatus() bool`

HasGroup6Multistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpGroup6DataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpGroup6DataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpGroup6DataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpGroup6DataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DhcpGroup6DataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DhcpGroup6DataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DhcpGroup6DataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DhcpGroup6DataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


