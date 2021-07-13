# IpamPoolDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolEndHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;end_ip_addr&lt;/b&gt;. | [optional] 
**PoolEndAddressAddr** | Pointer to **string** | The last IP address of the IPv4 pool, in hexadecimal format. | [optional] 
**PoolMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentNetworkClassName** | Pointer to **string** | The name of the class applied to the parent of the IPv4 network the object belongs to, it can be preceded by the class directory. | [optional] 
**ParentNetworkId** | Pointer to **string** | The database identifier (ID) of the parent IPv4 network. It identifies the parent of the IPv4 network the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the network the object belongs to has no parent network. | [optional] 
**ParentNetworkName** | Pointer to **string** | The name of the parent IPv4 network:&lt;ul class&#x3D;dashed &gt;&lt;li&gt;&lt;b&gt;#&lt;/b&gt; indicates that the network the object belongs to has no parent network.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;&lt;b&gt;Default&lt;/b&gt; indicates that the network the object belongs to is in an orphan network.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt; | [optional] 
**ParentNetworkSize** | Pointer to **string** | The number of IP addresses of the parent of the network the object belongs to. | [optional] 
**PoolClassName** | Pointer to **string** | The name of the class applied to the IPv4 pool, it can be preceded by the class directory. | [optional] 
**PoolClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**PoolId** | Pointer to **string** | The database identifier (ID) of the IPv4 pool. | [optional] 
**PoolName** | Pointer to **string** | The name of the IPv4 pool. | [optional] 
**PoolReadOnly** | Pointer to **string** | The reservation status of the IPv4 pool. If set &lt;b&gt;1&lt;/b&gt;, the IP addresses it contains cannot be assigned. | [optional] 
**PoolSize** | Pointer to **string** | The number of IP addresses the IPv4 pool contains. | [optional] 
**PoolStartAddressAddr** | Pointer to **string** | The first IP address of the IPv4 pool, in hexadecimal format. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space the object belongs to. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceIsTemplate** | Pointer to **string** | The template status of the space the object belongs to. If the space is used as template (1), all the IPv4 networks, pools and IP addresses it contains are also used as template. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**PoolStartHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;start_ip_addr&lt;/b&gt;. | [optional] 
**NetworkClassName** | Pointer to **string** | The name of the class applied to the IPv4 network the object belongs to, it can be preceded by the class directory. | [optional] 
**NetworkClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**NetworkEndAddressAddr** | Pointer to **string** | The last IP address of the IPv4 network the object belongs to. | [optional] 
**NetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 network the object belongs to. | [optional] 
**NetworkName** | Pointer to **string** | The name of the IPv4 network the object belongs to. &lt;b&gt;Default&lt;/b&gt; indicates that the network the object belongs to is an orphan network. | [optional] 
**NetworkSize** | Pointer to **string** | The number of IP addresses the network the object belongs to contains. | [optional] 
**NetworkStartAddressAddr** | Pointer to **string** | The first IP address of the IPv4 network the object belongs to. | [optional] 
**VlsmBlockId** | Pointer to **string** | The database identifier (ID) of the IPv4 VLSM block-type network duplicated, in a VLSM child space, from the network the pool belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the parent of the network the pool belongs to is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmNetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 subnet-type network, located in the VLSM parent space, from which the parent of the network the pool belongs to was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the parent of the network the pool belongs to is not a VLSM block-type network duplicated from a parent space. | [optional] 

## Methods

### NewIpamPoolDataData

`func NewIpamPoolDataData() *IpamPoolDataData`

NewIpamPoolDataData instantiates a new IpamPoolDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamPoolDataDataWithDefaults

`func NewIpamPoolDataDataWithDefaults() *IpamPoolDataData`

NewIpamPoolDataDataWithDefaults instantiates a new IpamPoolDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolEndHostaddr

`func (o *IpamPoolDataData) GetPoolEndHostaddr() string`

GetPoolEndHostaddr returns the PoolEndHostaddr field if non-nil, zero value otherwise.

### GetPoolEndHostaddrOk

`func (o *IpamPoolDataData) GetPoolEndHostaddrOk() (*string, bool)`

GetPoolEndHostaddrOk returns a tuple with the PoolEndHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEndHostaddr

`func (o *IpamPoolDataData) SetPoolEndHostaddr(v string)`

SetPoolEndHostaddr sets PoolEndHostaddr field to given value.

### HasPoolEndHostaddr

`func (o *IpamPoolDataData) HasPoolEndHostaddr() bool`

HasPoolEndHostaddr returns a boolean if a field has been set.

### GetPoolEndAddressAddr

`func (o *IpamPoolDataData) GetPoolEndAddressAddr() string`

GetPoolEndAddressAddr returns the PoolEndAddressAddr field if non-nil, zero value otherwise.

### GetPoolEndAddressAddrOk

`func (o *IpamPoolDataData) GetPoolEndAddressAddrOk() (*string, bool)`

GetPoolEndAddressAddrOk returns a tuple with the PoolEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEndAddressAddr

`func (o *IpamPoolDataData) SetPoolEndAddressAddr(v string)`

SetPoolEndAddressAddr sets PoolEndAddressAddr field to given value.

### HasPoolEndAddressAddr

`func (o *IpamPoolDataData) HasPoolEndAddressAddr() bool`

HasPoolEndAddressAddr returns a boolean if a field has been set.

### GetPoolMultistatus

`func (o *IpamPoolDataData) GetPoolMultistatus() string`

GetPoolMultistatus returns the PoolMultistatus field if non-nil, zero value otherwise.

### GetPoolMultistatusOk

`func (o *IpamPoolDataData) GetPoolMultistatusOk() (*string, bool)`

GetPoolMultistatusOk returns a tuple with the PoolMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMultistatus

`func (o *IpamPoolDataData) SetPoolMultistatus(v string)`

SetPoolMultistatus sets PoolMultistatus field to given value.

### HasPoolMultistatus

`func (o *IpamPoolDataData) HasPoolMultistatus() bool`

HasPoolMultistatus returns a boolean if a field has been set.

### GetParentNetworkClassName

`func (o *IpamPoolDataData) GetParentNetworkClassName() string`

GetParentNetworkClassName returns the ParentNetworkClassName field if non-nil, zero value otherwise.

### GetParentNetworkClassNameOk

`func (o *IpamPoolDataData) GetParentNetworkClassNameOk() (*string, bool)`

GetParentNetworkClassNameOk returns a tuple with the ParentNetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkClassName

`func (o *IpamPoolDataData) SetParentNetworkClassName(v string)`

SetParentNetworkClassName sets ParentNetworkClassName field to given value.

### HasParentNetworkClassName

`func (o *IpamPoolDataData) HasParentNetworkClassName() bool`

HasParentNetworkClassName returns a boolean if a field has been set.

### GetParentNetworkId

`func (o *IpamPoolDataData) GetParentNetworkId() string`

GetParentNetworkId returns the ParentNetworkId field if non-nil, zero value otherwise.

### GetParentNetworkIdOk

`func (o *IpamPoolDataData) GetParentNetworkIdOk() (*string, bool)`

GetParentNetworkIdOk returns a tuple with the ParentNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkId

`func (o *IpamPoolDataData) SetParentNetworkId(v string)`

SetParentNetworkId sets ParentNetworkId field to given value.

### HasParentNetworkId

`func (o *IpamPoolDataData) HasParentNetworkId() bool`

HasParentNetworkId returns a boolean if a field has been set.

### GetParentNetworkName

`func (o *IpamPoolDataData) GetParentNetworkName() string`

GetParentNetworkName returns the ParentNetworkName field if non-nil, zero value otherwise.

### GetParentNetworkNameOk

`func (o *IpamPoolDataData) GetParentNetworkNameOk() (*string, bool)`

GetParentNetworkNameOk returns a tuple with the ParentNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkName

`func (o *IpamPoolDataData) SetParentNetworkName(v string)`

SetParentNetworkName sets ParentNetworkName field to given value.

### HasParentNetworkName

`func (o *IpamPoolDataData) HasParentNetworkName() bool`

HasParentNetworkName returns a boolean if a field has been set.

### GetParentNetworkSize

`func (o *IpamPoolDataData) GetParentNetworkSize() string`

GetParentNetworkSize returns the ParentNetworkSize field if non-nil, zero value otherwise.

### GetParentNetworkSizeOk

`func (o *IpamPoolDataData) GetParentNetworkSizeOk() (*string, bool)`

GetParentNetworkSizeOk returns a tuple with the ParentNetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkSize

`func (o *IpamPoolDataData) SetParentNetworkSize(v string)`

SetParentNetworkSize sets ParentNetworkSize field to given value.

### HasParentNetworkSize

`func (o *IpamPoolDataData) HasParentNetworkSize() bool`

HasParentNetworkSize returns a boolean if a field has been set.

### GetPoolClassName

`func (o *IpamPoolDataData) GetPoolClassName() string`

GetPoolClassName returns the PoolClassName field if non-nil, zero value otherwise.

### GetPoolClassNameOk

`func (o *IpamPoolDataData) GetPoolClassNameOk() (*string, bool)`

GetPoolClassNameOk returns a tuple with the PoolClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassName

`func (o *IpamPoolDataData) SetPoolClassName(v string)`

SetPoolClassName sets PoolClassName field to given value.

### HasPoolClassName

`func (o *IpamPoolDataData) HasPoolClassName() bool`

HasPoolClassName returns a boolean if a field has been set.

### GetPoolClassParameters

`func (o *IpamPoolDataData) GetPoolClassParameters() []ApiClassParameterOutputEntry`

GetPoolClassParameters returns the PoolClassParameters field if non-nil, zero value otherwise.

### GetPoolClassParametersOk

`func (o *IpamPoolDataData) GetPoolClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetPoolClassParametersOk returns a tuple with the PoolClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassParameters

`func (o *IpamPoolDataData) SetPoolClassParameters(v []ApiClassParameterOutputEntry)`

SetPoolClassParameters sets PoolClassParameters field to given value.

### HasPoolClassParameters

`func (o *IpamPoolDataData) HasPoolClassParameters() bool`

HasPoolClassParameters returns a boolean if a field has been set.

### GetPoolId

`func (o *IpamPoolDataData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *IpamPoolDataData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *IpamPoolDataData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *IpamPoolDataData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *IpamPoolDataData) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *IpamPoolDataData) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *IpamPoolDataData) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *IpamPoolDataData) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolReadOnly

`func (o *IpamPoolDataData) GetPoolReadOnly() string`

GetPoolReadOnly returns the PoolReadOnly field if non-nil, zero value otherwise.

### GetPoolReadOnlyOk

`func (o *IpamPoolDataData) GetPoolReadOnlyOk() (*string, bool)`

GetPoolReadOnlyOk returns a tuple with the PoolReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolReadOnly

`func (o *IpamPoolDataData) SetPoolReadOnly(v string)`

SetPoolReadOnly sets PoolReadOnly field to given value.

### HasPoolReadOnly

`func (o *IpamPoolDataData) HasPoolReadOnly() bool`

HasPoolReadOnly returns a boolean if a field has been set.

### GetPoolSize

`func (o *IpamPoolDataData) GetPoolSize() string`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *IpamPoolDataData) GetPoolSizeOk() (*string, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *IpamPoolDataData) SetPoolSize(v string)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *IpamPoolDataData) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.

### GetPoolStartAddressAddr

`func (o *IpamPoolDataData) GetPoolStartAddressAddr() string`

GetPoolStartAddressAddr returns the PoolStartAddressAddr field if non-nil, zero value otherwise.

### GetPoolStartAddressAddrOk

`func (o *IpamPoolDataData) GetPoolStartAddressAddrOk() (*string, bool)`

GetPoolStartAddressAddrOk returns a tuple with the PoolStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStartAddressAddr

`func (o *IpamPoolDataData) SetPoolStartAddressAddr(v string)`

SetPoolStartAddressAddr sets PoolStartAddressAddr field to given value.

### HasPoolStartAddressAddr

`func (o *IpamPoolDataData) HasPoolStartAddressAddr() bool`

HasPoolStartAddressAddr returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *IpamPoolDataData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *IpamPoolDataData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *IpamPoolDataData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *IpamPoolDataData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *IpamPoolDataData) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *IpamPoolDataData) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *IpamPoolDataData) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *IpamPoolDataData) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *IpamPoolDataData) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *IpamPoolDataData) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *IpamPoolDataData) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *IpamPoolDataData) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamPoolDataData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamPoolDataData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamPoolDataData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamPoolDataData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceIsTemplate

`func (o *IpamPoolDataData) GetSpaceIsTemplate() string`

GetSpaceIsTemplate returns the SpaceIsTemplate field if non-nil, zero value otherwise.

### GetSpaceIsTemplateOk

`func (o *IpamPoolDataData) GetSpaceIsTemplateOk() (*string, bool)`

GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceIsTemplate

`func (o *IpamPoolDataData) SetSpaceIsTemplate(v string)`

SetSpaceIsTemplate sets SpaceIsTemplate field to given value.

### HasSpaceIsTemplate

`func (o *IpamPoolDataData) HasSpaceIsTemplate() bool`

HasSpaceIsTemplate returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamPoolDataData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamPoolDataData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamPoolDataData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamPoolDataData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetPoolStartHostaddr

`func (o *IpamPoolDataData) GetPoolStartHostaddr() string`

GetPoolStartHostaddr returns the PoolStartHostaddr field if non-nil, zero value otherwise.

### GetPoolStartHostaddrOk

`func (o *IpamPoolDataData) GetPoolStartHostaddrOk() (*string, bool)`

GetPoolStartHostaddrOk returns a tuple with the PoolStartHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStartHostaddr

`func (o *IpamPoolDataData) SetPoolStartHostaddr(v string)`

SetPoolStartHostaddr sets PoolStartHostaddr field to given value.

### HasPoolStartHostaddr

`func (o *IpamPoolDataData) HasPoolStartHostaddr() bool`

HasPoolStartHostaddr returns a boolean if a field has been set.

### GetNetworkClassName

`func (o *IpamPoolDataData) GetNetworkClassName() string`

GetNetworkClassName returns the NetworkClassName field if non-nil, zero value otherwise.

### GetNetworkClassNameOk

`func (o *IpamPoolDataData) GetNetworkClassNameOk() (*string, bool)`

GetNetworkClassNameOk returns a tuple with the NetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassName

`func (o *IpamPoolDataData) SetNetworkClassName(v string)`

SetNetworkClassName sets NetworkClassName field to given value.

### HasNetworkClassName

`func (o *IpamPoolDataData) HasNetworkClassName() bool`

HasNetworkClassName returns a boolean if a field has been set.

### GetNetworkClassParameters

`func (o *IpamPoolDataData) GetNetworkClassParameters() []ApiClassParameterOutputEntry`

GetNetworkClassParameters returns the NetworkClassParameters field if non-nil, zero value otherwise.

### GetNetworkClassParametersOk

`func (o *IpamPoolDataData) GetNetworkClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetNetworkClassParametersOk returns a tuple with the NetworkClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassParameters

`func (o *IpamPoolDataData) SetNetworkClassParameters(v []ApiClassParameterOutputEntry)`

SetNetworkClassParameters sets NetworkClassParameters field to given value.

### HasNetworkClassParameters

`func (o *IpamPoolDataData) HasNetworkClassParameters() bool`

HasNetworkClassParameters returns a boolean if a field has been set.

### GetNetworkEndAddressAddr

`func (o *IpamPoolDataData) GetNetworkEndAddressAddr() string`

GetNetworkEndAddressAddr returns the NetworkEndAddressAddr field if non-nil, zero value otherwise.

### GetNetworkEndAddressAddrOk

`func (o *IpamPoolDataData) GetNetworkEndAddressAddrOk() (*string, bool)`

GetNetworkEndAddressAddrOk returns a tuple with the NetworkEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndAddressAddr

`func (o *IpamPoolDataData) SetNetworkEndAddressAddr(v string)`

SetNetworkEndAddressAddr sets NetworkEndAddressAddr field to given value.

### HasNetworkEndAddressAddr

`func (o *IpamPoolDataData) HasNetworkEndAddressAddr() bool`

HasNetworkEndAddressAddr returns a boolean if a field has been set.

### GetNetworkId

`func (o *IpamPoolDataData) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *IpamPoolDataData) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *IpamPoolDataData) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *IpamPoolDataData) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *IpamPoolDataData) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *IpamPoolDataData) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *IpamPoolDataData) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *IpamPoolDataData) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkSize

`func (o *IpamPoolDataData) GetNetworkSize() string`

GetNetworkSize returns the NetworkSize field if non-nil, zero value otherwise.

### GetNetworkSizeOk

`func (o *IpamPoolDataData) GetNetworkSizeOk() (*string, bool)`

GetNetworkSizeOk returns a tuple with the NetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSize

`func (o *IpamPoolDataData) SetNetworkSize(v string)`

SetNetworkSize sets NetworkSize field to given value.

### HasNetworkSize

`func (o *IpamPoolDataData) HasNetworkSize() bool`

HasNetworkSize returns a boolean if a field has been set.

### GetNetworkStartAddressAddr

`func (o *IpamPoolDataData) GetNetworkStartAddressAddr() string`

GetNetworkStartAddressAddr returns the NetworkStartAddressAddr field if non-nil, zero value otherwise.

### GetNetworkStartAddressAddrOk

`func (o *IpamPoolDataData) GetNetworkStartAddressAddrOk() (*string, bool)`

GetNetworkStartAddressAddrOk returns a tuple with the NetworkStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStartAddressAddr

`func (o *IpamPoolDataData) SetNetworkStartAddressAddr(v string)`

SetNetworkStartAddressAddr sets NetworkStartAddressAddr field to given value.

### HasNetworkStartAddressAddr

`func (o *IpamPoolDataData) HasNetworkStartAddressAddr() bool`

HasNetworkStartAddressAddr returns a boolean if a field has been set.

### GetVlsmBlockId

`func (o *IpamPoolDataData) GetVlsmBlockId() string`

GetVlsmBlockId returns the VlsmBlockId field if non-nil, zero value otherwise.

### GetVlsmBlockIdOk

`func (o *IpamPoolDataData) GetVlsmBlockIdOk() (*string, bool)`

GetVlsmBlockIdOk returns a tuple with the VlsmBlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmBlockId

`func (o *IpamPoolDataData) SetVlsmBlockId(v string)`

SetVlsmBlockId sets VlsmBlockId field to given value.

### HasVlsmBlockId

`func (o *IpamPoolDataData) HasVlsmBlockId() bool`

HasVlsmBlockId returns a boolean if a field has been set.

### GetVlsmNetworkId

`func (o *IpamPoolDataData) GetVlsmNetworkId() string`

GetVlsmNetworkId returns the VlsmNetworkId field if non-nil, zero value otherwise.

### GetVlsmNetworkIdOk

`func (o *IpamPoolDataData) GetVlsmNetworkIdOk() (*string, bool)`

GetVlsmNetworkIdOk returns a tuple with the VlsmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmNetworkId

`func (o *IpamPoolDataData) SetVlsmNetworkId(v string)`

SetVlsmNetworkId sets VlsmNetworkId field to given value.

### HasVlsmNetworkId

`func (o *IpamPoolDataData) HasVlsmNetworkId() bool`

HasVlsmNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


