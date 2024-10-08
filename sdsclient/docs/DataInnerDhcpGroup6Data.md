# DataInnerDhcpGroup6Data

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group6DelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**Server6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 server the object belongs to, it can be preceded by the class directory. | [optional] 
**Server6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 server the object belongs to. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server the object belongs to. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server the object belongs to. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server the object belongs to: &lt;table&gt;&lt;caption&gt;server6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv6 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Version** | Pointer to **string** | The version details of the DHCPv6 server the object belongs to. | [optional] 
**Group6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 group, it can be preceded by the class directory. | [optional] 
**Group6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 group. | [optional] 
**Group6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 group. | [optional] 
**Group6Name** | Pointer to **string** | The name of the DHCPv6 group. | [optional] 
**Server6Hostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server6_addr&lt;/b&gt; or &lt;b&gt;server6_addr6&lt;/b&gt;. | [optional] 
**Server6Addr6** | Pointer to **string** | The Management IP address of the DHCPv6 server the object belongs to, the IPv6 address configured when adding the server, in hexadecimal format. | [optional] 
**Server6Addr** | Pointer to **string** | The Management IP address of the DHCPv6 server the object belongs to, the IPv4 address configured when adding the server, in hexadecimal format. | [optional] 
**Group6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the server the object belongs to is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server the object belongs to. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 

## Methods

### NewDataInnerDhcpGroup6Data

`func NewDataInnerDhcpGroup6Data() *DataInnerDhcpGroup6Data`

NewDataInnerDhcpGroup6Data instantiates a new DataInnerDhcpGroup6Data object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpGroup6DataWithDefaults

`func NewDataInnerDhcpGroup6DataWithDefaults() *DataInnerDhcpGroup6Data`

NewDataInnerDhcpGroup6DataWithDefaults instantiates a new DataInnerDhcpGroup6Data object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup6DelayedTime

`func (o *DataInnerDhcpGroup6Data) GetGroup6DelayedTime() string`

GetGroup6DelayedTime returns the Group6DelayedTime field if non-nil, zero value otherwise.

### GetGroup6DelayedTimeOk

`func (o *DataInnerDhcpGroup6Data) GetGroup6DelayedTimeOk() (*string, bool)`

GetGroup6DelayedTimeOk returns a tuple with the Group6DelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6DelayedTime

`func (o *DataInnerDhcpGroup6Data) SetGroup6DelayedTime(v string)`

SetGroup6DelayedTime sets Group6DelayedTime field to given value.

### HasGroup6DelayedTime

`func (o *DataInnerDhcpGroup6Data) HasGroup6DelayedTime() bool`

HasGroup6DelayedTime returns a boolean if a field has been set.

### GetServer6ClassName

`func (o *DataInnerDhcpGroup6Data) GetServer6ClassName() string`

GetServer6ClassName returns the Server6ClassName field if non-nil, zero value otherwise.

### GetServer6ClassNameOk

`func (o *DataInnerDhcpGroup6Data) GetServer6ClassNameOk() (*string, bool)`

GetServer6ClassNameOk returns a tuple with the Server6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassName

`func (o *DataInnerDhcpGroup6Data) SetServer6ClassName(v string)`

SetServer6ClassName sets Server6ClassName field to given value.

### HasServer6ClassName

`func (o *DataInnerDhcpGroup6Data) HasServer6ClassName() bool`

HasServer6ClassName returns a boolean if a field has been set.

### GetServer6ClassParameters

`func (o *DataInnerDhcpGroup6Data) GetServer6ClassParameters() []ApiClassParameterOutputEntry`

GetServer6ClassParameters returns the Server6ClassParameters field if non-nil, zero value otherwise.

### GetServer6ClassParametersOk

`func (o *DataInnerDhcpGroup6Data) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassParameters

`func (o *DataInnerDhcpGroup6Data) SetServer6ClassParameters(v []ApiClassParameterOutputEntry)`

SetServer6ClassParameters sets Server6ClassParameters field to given value.

### HasServer6ClassParameters

`func (o *DataInnerDhcpGroup6Data) HasServer6ClassParameters() bool`

HasServer6ClassParameters returns a boolean if a field has been set.

### GetServer6Id

`func (o *DataInnerDhcpGroup6Data) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DataInnerDhcpGroup6Data) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DataInnerDhcpGroup6Data) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DataInnerDhcpGroup6Data) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DataInnerDhcpGroup6Data) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DataInnerDhcpGroup6Data) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DataInnerDhcpGroup6Data) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DataInnerDhcpGroup6Data) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6Type

`func (o *DataInnerDhcpGroup6Data) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DataInnerDhcpGroup6Data) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DataInnerDhcpGroup6Data) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DataInnerDhcpGroup6Data) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetServer6Version

`func (o *DataInnerDhcpGroup6Data) GetServer6Version() string`

GetServer6Version returns the Server6Version field if non-nil, zero value otherwise.

### GetServer6VersionOk

`func (o *DataInnerDhcpGroup6Data) GetServer6VersionOk() (*string, bool)`

GetServer6VersionOk returns a tuple with the Server6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Version

`func (o *DataInnerDhcpGroup6Data) SetServer6Version(v string)`

SetServer6Version sets Server6Version field to given value.

### HasServer6Version

`func (o *DataInnerDhcpGroup6Data) HasServer6Version() bool`

HasServer6Version returns a boolean if a field has been set.

### GetGroup6ClassName

`func (o *DataInnerDhcpGroup6Data) GetGroup6ClassName() string`

GetGroup6ClassName returns the Group6ClassName field if non-nil, zero value otherwise.

### GetGroup6ClassNameOk

`func (o *DataInnerDhcpGroup6Data) GetGroup6ClassNameOk() (*string, bool)`

GetGroup6ClassNameOk returns a tuple with the Group6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6ClassName

`func (o *DataInnerDhcpGroup6Data) SetGroup6ClassName(v string)`

SetGroup6ClassName sets Group6ClassName field to given value.

### HasGroup6ClassName

`func (o *DataInnerDhcpGroup6Data) HasGroup6ClassName() bool`

HasGroup6ClassName returns a boolean if a field has been set.

### GetGroup6ClassParameters

`func (o *DataInnerDhcpGroup6Data) GetGroup6ClassParameters() []ApiClassParameterOutputEntry`

GetGroup6ClassParameters returns the Group6ClassParameters field if non-nil, zero value otherwise.

### GetGroup6ClassParametersOk

`func (o *DataInnerDhcpGroup6Data) GetGroup6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetGroup6ClassParametersOk returns a tuple with the Group6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6ClassParameters

`func (o *DataInnerDhcpGroup6Data) SetGroup6ClassParameters(v []ApiClassParameterOutputEntry)`

SetGroup6ClassParameters sets Group6ClassParameters field to given value.

### HasGroup6ClassParameters

`func (o *DataInnerDhcpGroup6Data) HasGroup6ClassParameters() bool`

HasGroup6ClassParameters returns a boolean if a field has been set.

### GetGroup6Id

`func (o *DataInnerDhcpGroup6Data) GetGroup6Id() string`

GetGroup6Id returns the Group6Id field if non-nil, zero value otherwise.

### GetGroup6IdOk

`func (o *DataInnerDhcpGroup6Data) GetGroup6IdOk() (*string, bool)`

GetGroup6IdOk returns a tuple with the Group6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Id

`func (o *DataInnerDhcpGroup6Data) SetGroup6Id(v string)`

SetGroup6Id sets Group6Id field to given value.

### HasGroup6Id

`func (o *DataInnerDhcpGroup6Data) HasGroup6Id() bool`

HasGroup6Id returns a boolean if a field has been set.

### GetGroup6Name

`func (o *DataInnerDhcpGroup6Data) GetGroup6Name() string`

GetGroup6Name returns the Group6Name field if non-nil, zero value otherwise.

### GetGroup6NameOk

`func (o *DataInnerDhcpGroup6Data) GetGroup6NameOk() (*string, bool)`

GetGroup6NameOk returns a tuple with the Group6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Name

`func (o *DataInnerDhcpGroup6Data) SetGroup6Name(v string)`

SetGroup6Name sets Group6Name field to given value.

### HasGroup6Name

`func (o *DataInnerDhcpGroup6Data) HasGroup6Name() bool`

HasGroup6Name returns a boolean if a field has been set.

### GetServer6Hostaddr

`func (o *DataInnerDhcpGroup6Data) GetServer6Hostaddr() string`

GetServer6Hostaddr returns the Server6Hostaddr field if non-nil, zero value otherwise.

### GetServer6HostaddrOk

`func (o *DataInnerDhcpGroup6Data) GetServer6HostaddrOk() (*string, bool)`

GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Hostaddr

`func (o *DataInnerDhcpGroup6Data) SetServer6Hostaddr(v string)`

SetServer6Hostaddr sets Server6Hostaddr field to given value.

### HasServer6Hostaddr

`func (o *DataInnerDhcpGroup6Data) HasServer6Hostaddr() bool`

HasServer6Hostaddr returns a boolean if a field has been set.

### GetServer6Addr6

`func (o *DataInnerDhcpGroup6Data) GetServer6Addr6() string`

GetServer6Addr6 returns the Server6Addr6 field if non-nil, zero value otherwise.

### GetServer6Addr6Ok

`func (o *DataInnerDhcpGroup6Data) GetServer6Addr6Ok() (*string, bool)`

GetServer6Addr6Ok returns a tuple with the Server6Addr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr6

`func (o *DataInnerDhcpGroup6Data) SetServer6Addr6(v string)`

SetServer6Addr6 sets Server6Addr6 field to given value.

### HasServer6Addr6

`func (o *DataInnerDhcpGroup6Data) HasServer6Addr6() bool`

HasServer6Addr6 returns a boolean if a field has been set.

### GetServer6Addr

`func (o *DataInnerDhcpGroup6Data) GetServer6Addr() string`

GetServer6Addr returns the Server6Addr field if non-nil, zero value otherwise.

### GetServer6AddrOk

`func (o *DataInnerDhcpGroup6Data) GetServer6AddrOk() (*string, bool)`

GetServer6AddrOk returns a tuple with the Server6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr

`func (o *DataInnerDhcpGroup6Data) SetServer6Addr(v string)`

SetServer6Addr sets Server6Addr field to given value.

### HasServer6Addr

`func (o *DataInnerDhcpGroup6Data) HasServer6Addr() bool`

HasServer6Addr returns a boolean if a field has been set.

### GetGroup6Multistatus

`func (o *DataInnerDhcpGroup6Data) GetGroup6Multistatus() string`

GetGroup6Multistatus returns the Group6Multistatus field if non-nil, zero value otherwise.

### GetGroup6MultistatusOk

`func (o *DataInnerDhcpGroup6Data) GetGroup6MultistatusOk() (*string, bool)`

GetGroup6MultistatusOk returns a tuple with the Group6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup6Multistatus

`func (o *DataInnerDhcpGroup6Data) SetGroup6Multistatus(v string)`

SetGroup6Multistatus sets Group6Multistatus field to given value.

### HasGroup6Multistatus

`func (o *DataInnerDhcpGroup6Data) HasGroup6Multistatus() bool`

HasGroup6Multistatus returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpGroup6Data) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpGroup6Data) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpGroup6Data) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpGroup6Data) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpGroup6Data) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpGroup6Data) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpGroup6Data) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpGroup6Data) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


