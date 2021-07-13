# IpamPool6DataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pool6EndHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;end_ip6_addr&lt;/b&gt;. | [optional] 
**EndAddress6Addr** | Pointer to **string** | The last IP address of the IPv6 pool, in hexadecimal format. | [optional] 
**Pool6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentNetwork6ClassName** | Pointer to **string** | The name of the class applied to the parent of the IPv6 network the object belongs to, it can be preceded by the class directory. | [optional] 
**ParentNetwork6Id** | Pointer to **string** | The database identifier (ID) of the parent IPv6 network. It identifies the parent of the IPv6 network the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the network the object belongs to has no parent network. | [optional] 
**ParentNetwork6Name** | Pointer to **string** | The name of the parent IPv6 network. &lt;b&gt;#&lt;/b&gt; indicates that the network the object belongs to has no parent network. | [optional] 
**ParentNetwork6Prefix** | Pointer to **string** | The prefix of the parent of the IPv6 network the object belongs to. | [optional] 
**Pool6ClassName** | Pointer to **string** | The name of the class applied to the IPv6 pool, it can be preceded by the class directory. | [optional] 
**Pool6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv6 pool and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Pool6EndAddress6Addr** | Pointer to **string** | The last IP address of the IPv6 pool, in hexadecimal format. | [optional] 
**Pool6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 pool. | [optional] 
**Pool6Name** | Pointer to **string** | The name of the IPv6 pool. | [optional] 
**Pool6ReadOnly** | Pointer to **string** | The reservation status of the IPv6 pool. If set &lt;b&gt;1&lt;/b&gt;, the IP addresses it contains cannot be assigned. | [optional] 
**Pool6Size** | Pointer to **string** | The number of IP addresses the IPv6 pool contains. | [optional] 
**Pool6StartAddress6Addr** | Pointer to **string** | The first IP address of the IPv6 pool, in hexadecimal format. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space the object belongs to. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceIsTemplate** | Pointer to **string** | The template status of the space the object belongs to. If the space is used as template (1), all the IPv4 networks, pools and IP addresses it contains are also used as template. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**Pool6StartHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;start_ip6_addr&lt;/b&gt;. | [optional] 
**StartAddress6Addr** | Pointer to **string** | The first IP address of the IPv6 pool, in hexadecimal format. | [optional] 
**Network6ClassName** | Pointer to **string** | The name of the class applied to the IPv6 network the object belongs to, it can be preceded by the class directory. | [optional] 
**Network6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv6 network the object belongs to and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Network6EndAddress6Addr** | Pointer to **string** | The last IP address of the IPv6 network the object belongs to. | [optional] 
**Network6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 network the object belongs to. | [optional] 
**Network6Name** | Pointer to **string** | The name of the IPv6 network the object belongs to. | [optional] 
**Network6Prefix** | Pointer to **string** | The prefix of the IPv6 network the object belongs to. | [optional] 
**Network6StartAddress6Addr** | Pointer to **string** | The first IP address of the IPv6 network the object belongs to. | [optional] 
**VlsmBlock6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 VLSM block-type network duplicated, in a VLSM child space, from the network the pool belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the parent of the network the pool belongs to is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmNetwork6Id** | Pointer to **string** | The database identifier (ID) of the IPv6 subnet-type network, located in the VLSM parent space, from which the parent of the network the pool belongs to was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the parent of the network the pool belongs to is not a VLSM block-type network duplicated from a parent space. | [optional] 

## Methods

### NewIpamPool6DataData

`func NewIpamPool6DataData() *IpamPool6DataData`

NewIpamPool6DataData instantiates a new IpamPool6DataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamPool6DataDataWithDefaults

`func NewIpamPool6DataDataWithDefaults() *IpamPool6DataData`

NewIpamPool6DataDataWithDefaults instantiates a new IpamPool6DataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPool6EndHostaddr

`func (o *IpamPool6DataData) GetPool6EndHostaddr() string`

GetPool6EndHostaddr returns the Pool6EndHostaddr field if non-nil, zero value otherwise.

### GetPool6EndHostaddrOk

`func (o *IpamPool6DataData) GetPool6EndHostaddrOk() (*string, bool)`

GetPool6EndHostaddrOk returns a tuple with the Pool6EndHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6EndHostaddr

`func (o *IpamPool6DataData) SetPool6EndHostaddr(v string)`

SetPool6EndHostaddr sets Pool6EndHostaddr field to given value.

### HasPool6EndHostaddr

`func (o *IpamPool6DataData) HasPool6EndHostaddr() bool`

HasPool6EndHostaddr returns a boolean if a field has been set.

### GetEndAddress6Addr

`func (o *IpamPool6DataData) GetEndAddress6Addr() string`

GetEndAddress6Addr returns the EndAddress6Addr field if non-nil, zero value otherwise.

### GetEndAddress6AddrOk

`func (o *IpamPool6DataData) GetEndAddress6AddrOk() (*string, bool)`

GetEndAddress6AddrOk returns a tuple with the EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndAddress6Addr

`func (o *IpamPool6DataData) SetEndAddress6Addr(v string)`

SetEndAddress6Addr sets EndAddress6Addr field to given value.

### HasEndAddress6Addr

`func (o *IpamPool6DataData) HasEndAddress6Addr() bool`

HasEndAddress6Addr returns a boolean if a field has been set.

### GetPool6Multistatus

`func (o *IpamPool6DataData) GetPool6Multistatus() string`

GetPool6Multistatus returns the Pool6Multistatus field if non-nil, zero value otherwise.

### GetPool6MultistatusOk

`func (o *IpamPool6DataData) GetPool6MultistatusOk() (*string, bool)`

GetPool6MultistatusOk returns a tuple with the Pool6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Multistatus

`func (o *IpamPool6DataData) SetPool6Multistatus(v string)`

SetPool6Multistatus sets Pool6Multistatus field to given value.

### HasPool6Multistatus

`func (o *IpamPool6DataData) HasPool6Multistatus() bool`

HasPool6Multistatus returns a boolean if a field has been set.

### GetParentNetwork6ClassName

`func (o *IpamPool6DataData) GetParentNetwork6ClassName() string`

GetParentNetwork6ClassName returns the ParentNetwork6ClassName field if non-nil, zero value otherwise.

### GetParentNetwork6ClassNameOk

`func (o *IpamPool6DataData) GetParentNetwork6ClassNameOk() (*string, bool)`

GetParentNetwork6ClassNameOk returns a tuple with the ParentNetwork6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6ClassName

`func (o *IpamPool6DataData) SetParentNetwork6ClassName(v string)`

SetParentNetwork6ClassName sets ParentNetwork6ClassName field to given value.

### HasParentNetwork6ClassName

`func (o *IpamPool6DataData) HasParentNetwork6ClassName() bool`

HasParentNetwork6ClassName returns a boolean if a field has been set.

### GetParentNetwork6Id

`func (o *IpamPool6DataData) GetParentNetwork6Id() string`

GetParentNetwork6Id returns the ParentNetwork6Id field if non-nil, zero value otherwise.

### GetParentNetwork6IdOk

`func (o *IpamPool6DataData) GetParentNetwork6IdOk() (*string, bool)`

GetParentNetwork6IdOk returns a tuple with the ParentNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Id

`func (o *IpamPool6DataData) SetParentNetwork6Id(v string)`

SetParentNetwork6Id sets ParentNetwork6Id field to given value.

### HasParentNetwork6Id

`func (o *IpamPool6DataData) HasParentNetwork6Id() bool`

HasParentNetwork6Id returns a boolean if a field has been set.

### GetParentNetwork6Name

`func (o *IpamPool6DataData) GetParentNetwork6Name() string`

GetParentNetwork6Name returns the ParentNetwork6Name field if non-nil, zero value otherwise.

### GetParentNetwork6NameOk

`func (o *IpamPool6DataData) GetParentNetwork6NameOk() (*string, bool)`

GetParentNetwork6NameOk returns a tuple with the ParentNetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Name

`func (o *IpamPool6DataData) SetParentNetwork6Name(v string)`

SetParentNetwork6Name sets ParentNetwork6Name field to given value.

### HasParentNetwork6Name

`func (o *IpamPool6DataData) HasParentNetwork6Name() bool`

HasParentNetwork6Name returns a boolean if a field has been set.

### GetParentNetwork6Prefix

`func (o *IpamPool6DataData) GetParentNetwork6Prefix() string`

GetParentNetwork6Prefix returns the ParentNetwork6Prefix field if non-nil, zero value otherwise.

### GetParentNetwork6PrefixOk

`func (o *IpamPool6DataData) GetParentNetwork6PrefixOk() (*string, bool)`

GetParentNetwork6PrefixOk returns a tuple with the ParentNetwork6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Prefix

`func (o *IpamPool6DataData) SetParentNetwork6Prefix(v string)`

SetParentNetwork6Prefix sets ParentNetwork6Prefix field to given value.

### HasParentNetwork6Prefix

`func (o *IpamPool6DataData) HasParentNetwork6Prefix() bool`

HasParentNetwork6Prefix returns a boolean if a field has been set.

### GetPool6ClassName

`func (o *IpamPool6DataData) GetPool6ClassName() string`

GetPool6ClassName returns the Pool6ClassName field if non-nil, zero value otherwise.

### GetPool6ClassNameOk

`func (o *IpamPool6DataData) GetPool6ClassNameOk() (*string, bool)`

GetPool6ClassNameOk returns a tuple with the Pool6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ClassName

`func (o *IpamPool6DataData) SetPool6ClassName(v string)`

SetPool6ClassName sets Pool6ClassName field to given value.

### HasPool6ClassName

`func (o *IpamPool6DataData) HasPool6ClassName() bool`

HasPool6ClassName returns a boolean if a field has been set.

### GetPool6ClassParameters

`func (o *IpamPool6DataData) GetPool6ClassParameters() []ApiClassParameterOutputEntry`

GetPool6ClassParameters returns the Pool6ClassParameters field if non-nil, zero value otherwise.

### GetPool6ClassParametersOk

`func (o *IpamPool6DataData) GetPool6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetPool6ClassParametersOk returns a tuple with the Pool6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ClassParameters

`func (o *IpamPool6DataData) SetPool6ClassParameters(v []ApiClassParameterOutputEntry)`

SetPool6ClassParameters sets Pool6ClassParameters field to given value.

### HasPool6ClassParameters

`func (o *IpamPool6DataData) HasPool6ClassParameters() bool`

HasPool6ClassParameters returns a boolean if a field has been set.

### GetPool6EndAddress6Addr

`func (o *IpamPool6DataData) GetPool6EndAddress6Addr() string`

GetPool6EndAddress6Addr returns the Pool6EndAddress6Addr field if non-nil, zero value otherwise.

### GetPool6EndAddress6AddrOk

`func (o *IpamPool6DataData) GetPool6EndAddress6AddrOk() (*string, bool)`

GetPool6EndAddress6AddrOk returns a tuple with the Pool6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6EndAddress6Addr

`func (o *IpamPool6DataData) SetPool6EndAddress6Addr(v string)`

SetPool6EndAddress6Addr sets Pool6EndAddress6Addr field to given value.

### HasPool6EndAddress6Addr

`func (o *IpamPool6DataData) HasPool6EndAddress6Addr() bool`

HasPool6EndAddress6Addr returns a boolean if a field has been set.

### GetPool6Id

`func (o *IpamPool6DataData) GetPool6Id() string`

GetPool6Id returns the Pool6Id field if non-nil, zero value otherwise.

### GetPool6IdOk

`func (o *IpamPool6DataData) GetPool6IdOk() (*string, bool)`

GetPool6IdOk returns a tuple with the Pool6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Id

`func (o *IpamPool6DataData) SetPool6Id(v string)`

SetPool6Id sets Pool6Id field to given value.

### HasPool6Id

`func (o *IpamPool6DataData) HasPool6Id() bool`

HasPool6Id returns a boolean if a field has been set.

### GetPool6Name

`func (o *IpamPool6DataData) GetPool6Name() string`

GetPool6Name returns the Pool6Name field if non-nil, zero value otherwise.

### GetPool6NameOk

`func (o *IpamPool6DataData) GetPool6NameOk() (*string, bool)`

GetPool6NameOk returns a tuple with the Pool6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Name

`func (o *IpamPool6DataData) SetPool6Name(v string)`

SetPool6Name sets Pool6Name field to given value.

### HasPool6Name

`func (o *IpamPool6DataData) HasPool6Name() bool`

HasPool6Name returns a boolean if a field has been set.

### GetPool6ReadOnly

`func (o *IpamPool6DataData) GetPool6ReadOnly() string`

GetPool6ReadOnly returns the Pool6ReadOnly field if non-nil, zero value otherwise.

### GetPool6ReadOnlyOk

`func (o *IpamPool6DataData) GetPool6ReadOnlyOk() (*string, bool)`

GetPool6ReadOnlyOk returns a tuple with the Pool6ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6ReadOnly

`func (o *IpamPool6DataData) SetPool6ReadOnly(v string)`

SetPool6ReadOnly sets Pool6ReadOnly field to given value.

### HasPool6ReadOnly

`func (o *IpamPool6DataData) HasPool6ReadOnly() bool`

HasPool6ReadOnly returns a boolean if a field has been set.

### GetPool6Size

`func (o *IpamPool6DataData) GetPool6Size() string`

GetPool6Size returns the Pool6Size field if non-nil, zero value otherwise.

### GetPool6SizeOk

`func (o *IpamPool6DataData) GetPool6SizeOk() (*string, bool)`

GetPool6SizeOk returns a tuple with the Pool6Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6Size

`func (o *IpamPool6DataData) SetPool6Size(v string)`

SetPool6Size sets Pool6Size field to given value.

### HasPool6Size

`func (o *IpamPool6DataData) HasPool6Size() bool`

HasPool6Size returns a boolean if a field has been set.

### GetPool6StartAddress6Addr

`func (o *IpamPool6DataData) GetPool6StartAddress6Addr() string`

GetPool6StartAddress6Addr returns the Pool6StartAddress6Addr field if non-nil, zero value otherwise.

### GetPool6StartAddress6AddrOk

`func (o *IpamPool6DataData) GetPool6StartAddress6AddrOk() (*string, bool)`

GetPool6StartAddress6AddrOk returns a tuple with the Pool6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6StartAddress6Addr

`func (o *IpamPool6DataData) SetPool6StartAddress6Addr(v string)`

SetPool6StartAddress6Addr sets Pool6StartAddress6Addr field to given value.

### HasPool6StartAddress6Addr

`func (o *IpamPool6DataData) HasPool6StartAddress6Addr() bool`

HasPool6StartAddress6Addr returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *IpamPool6DataData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *IpamPool6DataData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *IpamPool6DataData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *IpamPool6DataData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *IpamPool6DataData) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *IpamPool6DataData) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *IpamPool6DataData) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *IpamPool6DataData) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *IpamPool6DataData) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *IpamPool6DataData) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *IpamPool6DataData) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *IpamPool6DataData) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamPool6DataData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamPool6DataData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamPool6DataData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamPool6DataData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceIsTemplate

`func (o *IpamPool6DataData) GetSpaceIsTemplate() string`

GetSpaceIsTemplate returns the SpaceIsTemplate field if non-nil, zero value otherwise.

### GetSpaceIsTemplateOk

`func (o *IpamPool6DataData) GetSpaceIsTemplateOk() (*string, bool)`

GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceIsTemplate

`func (o *IpamPool6DataData) SetSpaceIsTemplate(v string)`

SetSpaceIsTemplate sets SpaceIsTemplate field to given value.

### HasSpaceIsTemplate

`func (o *IpamPool6DataData) HasSpaceIsTemplate() bool`

HasSpaceIsTemplate returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamPool6DataData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamPool6DataData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamPool6DataData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamPool6DataData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetPool6StartHostaddr

`func (o *IpamPool6DataData) GetPool6StartHostaddr() string`

GetPool6StartHostaddr returns the Pool6StartHostaddr field if non-nil, zero value otherwise.

### GetPool6StartHostaddrOk

`func (o *IpamPool6DataData) GetPool6StartHostaddrOk() (*string, bool)`

GetPool6StartHostaddrOk returns a tuple with the Pool6StartHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool6StartHostaddr

`func (o *IpamPool6DataData) SetPool6StartHostaddr(v string)`

SetPool6StartHostaddr sets Pool6StartHostaddr field to given value.

### HasPool6StartHostaddr

`func (o *IpamPool6DataData) HasPool6StartHostaddr() bool`

HasPool6StartHostaddr returns a boolean if a field has been set.

### GetStartAddress6Addr

`func (o *IpamPool6DataData) GetStartAddress6Addr() string`

GetStartAddress6Addr returns the StartAddress6Addr field if non-nil, zero value otherwise.

### GetStartAddress6AddrOk

`func (o *IpamPool6DataData) GetStartAddress6AddrOk() (*string, bool)`

GetStartAddress6AddrOk returns a tuple with the StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAddress6Addr

`func (o *IpamPool6DataData) SetStartAddress6Addr(v string)`

SetStartAddress6Addr sets StartAddress6Addr field to given value.

### HasStartAddress6Addr

`func (o *IpamPool6DataData) HasStartAddress6Addr() bool`

HasStartAddress6Addr returns a boolean if a field has been set.

### GetNetwork6ClassName

`func (o *IpamPool6DataData) GetNetwork6ClassName() string`

GetNetwork6ClassName returns the Network6ClassName field if non-nil, zero value otherwise.

### GetNetwork6ClassNameOk

`func (o *IpamPool6DataData) GetNetwork6ClassNameOk() (*string, bool)`

GetNetwork6ClassNameOk returns a tuple with the Network6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6ClassName

`func (o *IpamPool6DataData) SetNetwork6ClassName(v string)`

SetNetwork6ClassName sets Network6ClassName field to given value.

### HasNetwork6ClassName

`func (o *IpamPool6DataData) HasNetwork6ClassName() bool`

HasNetwork6ClassName returns a boolean if a field has been set.

### GetNetwork6ClassParameters

`func (o *IpamPool6DataData) GetNetwork6ClassParameters() []ApiClassParameterOutputEntry`

GetNetwork6ClassParameters returns the Network6ClassParameters field if non-nil, zero value otherwise.

### GetNetwork6ClassParametersOk

`func (o *IpamPool6DataData) GetNetwork6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetNetwork6ClassParametersOk returns a tuple with the Network6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6ClassParameters

`func (o *IpamPool6DataData) SetNetwork6ClassParameters(v []ApiClassParameterOutputEntry)`

SetNetwork6ClassParameters sets Network6ClassParameters field to given value.

### HasNetwork6ClassParameters

`func (o *IpamPool6DataData) HasNetwork6ClassParameters() bool`

HasNetwork6ClassParameters returns a boolean if a field has been set.

### GetNetwork6EndAddress6Addr

`func (o *IpamPool6DataData) GetNetwork6EndAddress6Addr() string`

GetNetwork6EndAddress6Addr returns the Network6EndAddress6Addr field if non-nil, zero value otherwise.

### GetNetwork6EndAddress6AddrOk

`func (o *IpamPool6DataData) GetNetwork6EndAddress6AddrOk() (*string, bool)`

GetNetwork6EndAddress6AddrOk returns a tuple with the Network6EndAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6EndAddress6Addr

`func (o *IpamPool6DataData) SetNetwork6EndAddress6Addr(v string)`

SetNetwork6EndAddress6Addr sets Network6EndAddress6Addr field to given value.

### HasNetwork6EndAddress6Addr

`func (o *IpamPool6DataData) HasNetwork6EndAddress6Addr() bool`

HasNetwork6EndAddress6Addr returns a boolean if a field has been set.

### GetNetwork6Id

`func (o *IpamPool6DataData) GetNetwork6Id() string`

GetNetwork6Id returns the Network6Id field if non-nil, zero value otherwise.

### GetNetwork6IdOk

`func (o *IpamPool6DataData) GetNetwork6IdOk() (*string, bool)`

GetNetwork6IdOk returns a tuple with the Network6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Id

`func (o *IpamPool6DataData) SetNetwork6Id(v string)`

SetNetwork6Id sets Network6Id field to given value.

### HasNetwork6Id

`func (o *IpamPool6DataData) HasNetwork6Id() bool`

HasNetwork6Id returns a boolean if a field has been set.

### GetNetwork6Name

`func (o *IpamPool6DataData) GetNetwork6Name() string`

GetNetwork6Name returns the Network6Name field if non-nil, zero value otherwise.

### GetNetwork6NameOk

`func (o *IpamPool6DataData) GetNetwork6NameOk() (*string, bool)`

GetNetwork6NameOk returns a tuple with the Network6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Name

`func (o *IpamPool6DataData) SetNetwork6Name(v string)`

SetNetwork6Name sets Network6Name field to given value.

### HasNetwork6Name

`func (o *IpamPool6DataData) HasNetwork6Name() bool`

HasNetwork6Name returns a boolean if a field has been set.

### GetNetwork6Prefix

`func (o *IpamPool6DataData) GetNetwork6Prefix() string`

GetNetwork6Prefix returns the Network6Prefix field if non-nil, zero value otherwise.

### GetNetwork6PrefixOk

`func (o *IpamPool6DataData) GetNetwork6PrefixOk() (*string, bool)`

GetNetwork6PrefixOk returns a tuple with the Network6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Prefix

`func (o *IpamPool6DataData) SetNetwork6Prefix(v string)`

SetNetwork6Prefix sets Network6Prefix field to given value.

### HasNetwork6Prefix

`func (o *IpamPool6DataData) HasNetwork6Prefix() bool`

HasNetwork6Prefix returns a boolean if a field has been set.

### GetNetwork6StartAddress6Addr

`func (o *IpamPool6DataData) GetNetwork6StartAddress6Addr() string`

GetNetwork6StartAddress6Addr returns the Network6StartAddress6Addr field if non-nil, zero value otherwise.

### GetNetwork6StartAddress6AddrOk

`func (o *IpamPool6DataData) GetNetwork6StartAddress6AddrOk() (*string, bool)`

GetNetwork6StartAddress6AddrOk returns a tuple with the Network6StartAddress6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6StartAddress6Addr

`func (o *IpamPool6DataData) SetNetwork6StartAddress6Addr(v string)`

SetNetwork6StartAddress6Addr sets Network6StartAddress6Addr field to given value.

### HasNetwork6StartAddress6Addr

`func (o *IpamPool6DataData) HasNetwork6StartAddress6Addr() bool`

HasNetwork6StartAddress6Addr returns a boolean if a field has been set.

### GetVlsmBlock6Id

`func (o *IpamPool6DataData) GetVlsmBlock6Id() string`

GetVlsmBlock6Id returns the VlsmBlock6Id field if non-nil, zero value otherwise.

### GetVlsmBlock6IdOk

`func (o *IpamPool6DataData) GetVlsmBlock6IdOk() (*string, bool)`

GetVlsmBlock6IdOk returns a tuple with the VlsmBlock6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmBlock6Id

`func (o *IpamPool6DataData) SetVlsmBlock6Id(v string)`

SetVlsmBlock6Id sets VlsmBlock6Id field to given value.

### HasVlsmBlock6Id

`func (o *IpamPool6DataData) HasVlsmBlock6Id() bool`

HasVlsmBlock6Id returns a boolean if a field has been set.

### GetVlsmNetwork6Id

`func (o *IpamPool6DataData) GetVlsmNetwork6Id() string`

GetVlsmNetwork6Id returns the VlsmNetwork6Id field if non-nil, zero value otherwise.

### GetVlsmNetwork6IdOk

`func (o *IpamPool6DataData) GetVlsmNetwork6IdOk() (*string, bool)`

GetVlsmNetwork6IdOk returns a tuple with the VlsmNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmNetwork6Id

`func (o *IpamPool6DataData) SetVlsmNetwork6Id(v string)`

SetVlsmNetwork6Id sets VlsmNetwork6Id field to given value.

### HasVlsmNetwork6Id

`func (o *IpamPool6DataData) HasVlsmNetwork6Id() bool`

HasVlsmNetwork6Id returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


