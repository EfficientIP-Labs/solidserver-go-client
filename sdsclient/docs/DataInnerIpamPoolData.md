# DataInnerIpamPoolData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PoolEndHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;pool_end_ip_addr&lt;/b&gt;. | [optional] 
**PoolEndIpAddr** | Pointer to **string** | The last IP address of the IPv4 pool, in hexadecimal format. | [optional] 
**PoolMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentNetworkClassName** | Pointer to **string** | The name of the class applied to the parent of the IPv4 network the object belongs to, it can be preceded by the class directory. | [optional] 
**ParentNetworkId** | Pointer to **string** | The database identifier (ID) of the parent IPv4 network. It identifies the parent of the IPv4 network the object belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the network the object belongs to has no parent network. | [optional] 
**ParentNetworkName** | Pointer to **string** | The name of the parent IPv4 network:&lt;ul class&#x3D;dashed &gt;&lt;li&gt; &lt;b&gt;#&lt;/b&gt; indicates that the network the object belongs to has no parent network.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; &lt;b&gt;Default&lt;/b&gt; indicates that the network the object belongs to is in an orphan network.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**ParentNetworkSize** | Pointer to **string** | The number of IP addresses of the parent of the network the object belongs to. | [optional] 
**PoolClassName** | Pointer to **string** | The name of the class applied to the IPv4 pool, it can be preceded by the class directory. | [optional] 
**PoolClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv4 pool. | [optional] 
**PoolEndAddressAddr** | Pointer to **string** | The last IP address of the IPv4 pool, in hexadecimal format. | [optional] 
**PoolId** | Pointer to **string** | The database identifier (ID) of the IPv4 pool. | [optional] 
**PoolName** | Pointer to **string** | The name of the IPv4 pool. | [optional] 
**PoolReadOnly** | Pointer to **string** | The reservation status of the IPv4 pool. If set &lt;b&gt;1&lt;/b&gt;, the IP addresses it contains cannot be assigned. | [optional] 
**PoolSize** | Pointer to **string** | The number of IP addresses the IPv4 pool contains. | [optional] 
**PoolStartAddressAddr** | Pointer to **string** | The first IP address of the IPv4 pool, in hexadecimal format. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space the object belongs to, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the space the object belongs to. | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space the object belongs to. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space the object belongs to, a unique numeric key value automatically incremented when you add a space. | [optional] 
**SpaceIsTemplate** | Pointer to **string** | The template status of the space the object belongs to. If the space is used as template (1), all the IPv4 networks, pools and IP addresses it contains are also used as template. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space the object belongs to. | [optional] 
**PoolStartHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;pool_start_ip_addr&lt;/b&gt;. | [optional] 
**PoolStartIpAddr** | Pointer to **string** | The first IP address of the IPv4 pool, in hexadecimal format. | [optional] 
**NetworkClassName** | Pointer to **string** | The name of the class applied to the IPv4 network the object belongs to, it can be preceded by the class directory. | [optional] 
**NetworkClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the IPv4 network the object belongs to. | [optional] 
**NetworkEndAddressAddr** | Pointer to **string** | The last IP address of the IPv4 network the object belongs to. | [optional] 
**NetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 network the object belongs to. | [optional] 
**NetworkName** | Pointer to **string** | The name of the IPv4 network the object belongs to. &lt;b&gt;Default&lt;/b&gt; indicates that the network the object belongs to is an orphan network. | [optional] 
**NetworkSize** | Pointer to **string** | The number of IP addresses the network the object belongs to contains. | [optional] 
**NetworkStartAddressAddr** | Pointer to **string** | The first IP address of the IPv4 network the object belongs to. | [optional] 
**PoolOperationDetailsAddedOn** | Pointer to **string** | The creation date of the IPv4 pool, in decimal UNIX date format. | [optional] 
**PoolOperationDetailsCallStack** | Pointer to **string** | The call stack of the IPv4 pool operation details, as follows: &lt;b&gt;&amp;lt;service1&amp;gt;&amp;amp;&amp;lt;service2&amp;gt;&amp;amp;&amp;lt;service3&amp;gt;&lt;/b&gt;... . | [optional] 
**PoolOperationDetailsOriginModule** | Pointer to **string** | The name of the module where the IPv4 pool addition originated. | [optional] 
**PoolOperationDetailsRequestedBy** | Pointer to **string** | The login of the user who requested the IPv4 pool. | [optional] 
**PoolOperationDetailsAddedBy** | Pointer to **string** | The login of the user who added the IPv4 pool. | [optional] 
**PoolOperationDetailsLastUpdatedOn** | Pointer to **string** | The last time the IPv4 pool was updated, in decimal UNIX date format. | [optional] 
**VlsmBlockId** | Pointer to **string** | The database identifier (ID) of the IPv4 VLSM block-type network duplicated, in a VLSM child space, from the network the pool belongs to. &lt;b&gt;0&lt;/b&gt; indicates that the parent of the network the pool belongs to is not duplicated as a VLSM block-type network in a child space. | [optional] 
**VlsmNetworkId** | Pointer to **string** | The database identifier (ID) of the IPv4 subnet-type network, located in the VLSM parent space, from which the parent of the network the pool belongs to was duplicated. &lt;b&gt;0&lt;/b&gt; indicates that the parent of the network the pool belongs to is not a VLSM block-type network duplicated from a parent space. | [optional] 

## Methods

### NewDataInnerIpamPoolData

`func NewDataInnerIpamPoolData() *DataInnerIpamPoolData`

NewDataInnerIpamPoolData instantiates a new DataInnerIpamPoolData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerIpamPoolDataWithDefaults

`func NewDataInnerIpamPoolDataWithDefaults() *DataInnerIpamPoolData`

NewDataInnerIpamPoolDataWithDefaults instantiates a new DataInnerIpamPoolData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPoolEndHostaddr

`func (o *DataInnerIpamPoolData) GetPoolEndHostaddr() string`

GetPoolEndHostaddr returns the PoolEndHostaddr field if non-nil, zero value otherwise.

### GetPoolEndHostaddrOk

`func (o *DataInnerIpamPoolData) GetPoolEndHostaddrOk() (*string, bool)`

GetPoolEndHostaddrOk returns a tuple with the PoolEndHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEndHostaddr

`func (o *DataInnerIpamPoolData) SetPoolEndHostaddr(v string)`

SetPoolEndHostaddr sets PoolEndHostaddr field to given value.

### HasPoolEndHostaddr

`func (o *DataInnerIpamPoolData) HasPoolEndHostaddr() bool`

HasPoolEndHostaddr returns a boolean if a field has been set.

### GetPoolEndIpAddr

`func (o *DataInnerIpamPoolData) GetPoolEndIpAddr() string`

GetPoolEndIpAddr returns the PoolEndIpAddr field if non-nil, zero value otherwise.

### GetPoolEndIpAddrOk

`func (o *DataInnerIpamPoolData) GetPoolEndIpAddrOk() (*string, bool)`

GetPoolEndIpAddrOk returns a tuple with the PoolEndIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEndIpAddr

`func (o *DataInnerIpamPoolData) SetPoolEndIpAddr(v string)`

SetPoolEndIpAddr sets PoolEndIpAddr field to given value.

### HasPoolEndIpAddr

`func (o *DataInnerIpamPoolData) HasPoolEndIpAddr() bool`

HasPoolEndIpAddr returns a boolean if a field has been set.

### GetPoolMultistatus

`func (o *DataInnerIpamPoolData) GetPoolMultistatus() string`

GetPoolMultistatus returns the PoolMultistatus field if non-nil, zero value otherwise.

### GetPoolMultistatusOk

`func (o *DataInnerIpamPoolData) GetPoolMultistatusOk() (*string, bool)`

GetPoolMultistatusOk returns a tuple with the PoolMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolMultistatus

`func (o *DataInnerIpamPoolData) SetPoolMultistatus(v string)`

SetPoolMultistatus sets PoolMultistatus field to given value.

### HasPoolMultistatus

`func (o *DataInnerIpamPoolData) HasPoolMultistatus() bool`

HasPoolMultistatus returns a boolean if a field has been set.

### GetParentNetworkClassName

`func (o *DataInnerIpamPoolData) GetParentNetworkClassName() string`

GetParentNetworkClassName returns the ParentNetworkClassName field if non-nil, zero value otherwise.

### GetParentNetworkClassNameOk

`func (o *DataInnerIpamPoolData) GetParentNetworkClassNameOk() (*string, bool)`

GetParentNetworkClassNameOk returns a tuple with the ParentNetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkClassName

`func (o *DataInnerIpamPoolData) SetParentNetworkClassName(v string)`

SetParentNetworkClassName sets ParentNetworkClassName field to given value.

### HasParentNetworkClassName

`func (o *DataInnerIpamPoolData) HasParentNetworkClassName() bool`

HasParentNetworkClassName returns a boolean if a field has been set.

### GetParentNetworkId

`func (o *DataInnerIpamPoolData) GetParentNetworkId() string`

GetParentNetworkId returns the ParentNetworkId field if non-nil, zero value otherwise.

### GetParentNetworkIdOk

`func (o *DataInnerIpamPoolData) GetParentNetworkIdOk() (*string, bool)`

GetParentNetworkIdOk returns a tuple with the ParentNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkId

`func (o *DataInnerIpamPoolData) SetParentNetworkId(v string)`

SetParentNetworkId sets ParentNetworkId field to given value.

### HasParentNetworkId

`func (o *DataInnerIpamPoolData) HasParentNetworkId() bool`

HasParentNetworkId returns a boolean if a field has been set.

### GetParentNetworkName

`func (o *DataInnerIpamPoolData) GetParentNetworkName() string`

GetParentNetworkName returns the ParentNetworkName field if non-nil, zero value otherwise.

### GetParentNetworkNameOk

`func (o *DataInnerIpamPoolData) GetParentNetworkNameOk() (*string, bool)`

GetParentNetworkNameOk returns a tuple with the ParentNetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkName

`func (o *DataInnerIpamPoolData) SetParentNetworkName(v string)`

SetParentNetworkName sets ParentNetworkName field to given value.

### HasParentNetworkName

`func (o *DataInnerIpamPoolData) HasParentNetworkName() bool`

HasParentNetworkName returns a boolean if a field has been set.

### GetParentNetworkSize

`func (o *DataInnerIpamPoolData) GetParentNetworkSize() string`

GetParentNetworkSize returns the ParentNetworkSize field if non-nil, zero value otherwise.

### GetParentNetworkSizeOk

`func (o *DataInnerIpamPoolData) GetParentNetworkSizeOk() (*string, bool)`

GetParentNetworkSizeOk returns a tuple with the ParentNetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkSize

`func (o *DataInnerIpamPoolData) SetParentNetworkSize(v string)`

SetParentNetworkSize sets ParentNetworkSize field to given value.

### HasParentNetworkSize

`func (o *DataInnerIpamPoolData) HasParentNetworkSize() bool`

HasParentNetworkSize returns a boolean if a field has been set.

### GetPoolClassName

`func (o *DataInnerIpamPoolData) GetPoolClassName() string`

GetPoolClassName returns the PoolClassName field if non-nil, zero value otherwise.

### GetPoolClassNameOk

`func (o *DataInnerIpamPoolData) GetPoolClassNameOk() (*string, bool)`

GetPoolClassNameOk returns a tuple with the PoolClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassName

`func (o *DataInnerIpamPoolData) SetPoolClassName(v string)`

SetPoolClassName sets PoolClassName field to given value.

### HasPoolClassName

`func (o *DataInnerIpamPoolData) HasPoolClassName() bool`

HasPoolClassName returns a boolean if a field has been set.

### GetPoolClassParameters

`func (o *DataInnerIpamPoolData) GetPoolClassParameters() []ApiClassParameterOutputEntry`

GetPoolClassParameters returns the PoolClassParameters field if non-nil, zero value otherwise.

### GetPoolClassParametersOk

`func (o *DataInnerIpamPoolData) GetPoolClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetPoolClassParametersOk returns a tuple with the PoolClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolClassParameters

`func (o *DataInnerIpamPoolData) SetPoolClassParameters(v []ApiClassParameterOutputEntry)`

SetPoolClassParameters sets PoolClassParameters field to given value.

### HasPoolClassParameters

`func (o *DataInnerIpamPoolData) HasPoolClassParameters() bool`

HasPoolClassParameters returns a boolean if a field has been set.

### GetPoolEndAddressAddr

`func (o *DataInnerIpamPoolData) GetPoolEndAddressAddr() string`

GetPoolEndAddressAddr returns the PoolEndAddressAddr field if non-nil, zero value otherwise.

### GetPoolEndAddressAddrOk

`func (o *DataInnerIpamPoolData) GetPoolEndAddressAddrOk() (*string, bool)`

GetPoolEndAddressAddrOk returns a tuple with the PoolEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolEndAddressAddr

`func (o *DataInnerIpamPoolData) SetPoolEndAddressAddr(v string)`

SetPoolEndAddressAddr sets PoolEndAddressAddr field to given value.

### HasPoolEndAddressAddr

`func (o *DataInnerIpamPoolData) HasPoolEndAddressAddr() bool`

HasPoolEndAddressAddr returns a boolean if a field has been set.

### GetPoolId

`func (o *DataInnerIpamPoolData) GetPoolId() string`

GetPoolId returns the PoolId field if non-nil, zero value otherwise.

### GetPoolIdOk

`func (o *DataInnerIpamPoolData) GetPoolIdOk() (*string, bool)`

GetPoolIdOk returns a tuple with the PoolId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolId

`func (o *DataInnerIpamPoolData) SetPoolId(v string)`

SetPoolId sets PoolId field to given value.

### HasPoolId

`func (o *DataInnerIpamPoolData) HasPoolId() bool`

HasPoolId returns a boolean if a field has been set.

### GetPoolName

`func (o *DataInnerIpamPoolData) GetPoolName() string`

GetPoolName returns the PoolName field if non-nil, zero value otherwise.

### GetPoolNameOk

`func (o *DataInnerIpamPoolData) GetPoolNameOk() (*string, bool)`

GetPoolNameOk returns a tuple with the PoolName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolName

`func (o *DataInnerIpamPoolData) SetPoolName(v string)`

SetPoolName sets PoolName field to given value.

### HasPoolName

`func (o *DataInnerIpamPoolData) HasPoolName() bool`

HasPoolName returns a boolean if a field has been set.

### GetPoolReadOnly

`func (o *DataInnerIpamPoolData) GetPoolReadOnly() string`

GetPoolReadOnly returns the PoolReadOnly field if non-nil, zero value otherwise.

### GetPoolReadOnlyOk

`func (o *DataInnerIpamPoolData) GetPoolReadOnlyOk() (*string, bool)`

GetPoolReadOnlyOk returns a tuple with the PoolReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolReadOnly

`func (o *DataInnerIpamPoolData) SetPoolReadOnly(v string)`

SetPoolReadOnly sets PoolReadOnly field to given value.

### HasPoolReadOnly

`func (o *DataInnerIpamPoolData) HasPoolReadOnly() bool`

HasPoolReadOnly returns a boolean if a field has been set.

### GetPoolSize

`func (o *DataInnerIpamPoolData) GetPoolSize() string`

GetPoolSize returns the PoolSize field if non-nil, zero value otherwise.

### GetPoolSizeOk

`func (o *DataInnerIpamPoolData) GetPoolSizeOk() (*string, bool)`

GetPoolSizeOk returns a tuple with the PoolSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolSize

`func (o *DataInnerIpamPoolData) SetPoolSize(v string)`

SetPoolSize sets PoolSize field to given value.

### HasPoolSize

`func (o *DataInnerIpamPoolData) HasPoolSize() bool`

HasPoolSize returns a boolean if a field has been set.

### GetPoolStartAddressAddr

`func (o *DataInnerIpamPoolData) GetPoolStartAddressAddr() string`

GetPoolStartAddressAddr returns the PoolStartAddressAddr field if non-nil, zero value otherwise.

### GetPoolStartAddressAddrOk

`func (o *DataInnerIpamPoolData) GetPoolStartAddressAddrOk() (*string, bool)`

GetPoolStartAddressAddrOk returns a tuple with the PoolStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStartAddressAddr

`func (o *DataInnerIpamPoolData) SetPoolStartAddressAddr(v string)`

SetPoolStartAddressAddr sets PoolStartAddressAddr field to given value.

### HasPoolStartAddressAddr

`func (o *DataInnerIpamPoolData) HasPoolStartAddressAddr() bool`

HasPoolStartAddressAddr returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *DataInnerIpamPoolData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *DataInnerIpamPoolData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *DataInnerIpamPoolData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *DataInnerIpamPoolData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *DataInnerIpamPoolData) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *DataInnerIpamPoolData) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *DataInnerIpamPoolData) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *DataInnerIpamPoolData) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *DataInnerIpamPoolData) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *DataInnerIpamPoolData) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *DataInnerIpamPoolData) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *DataInnerIpamPoolData) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *DataInnerIpamPoolData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DataInnerIpamPoolData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DataInnerIpamPoolData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DataInnerIpamPoolData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceIsTemplate

`func (o *DataInnerIpamPoolData) GetSpaceIsTemplate() string`

GetSpaceIsTemplate returns the SpaceIsTemplate field if non-nil, zero value otherwise.

### GetSpaceIsTemplateOk

`func (o *DataInnerIpamPoolData) GetSpaceIsTemplateOk() (*string, bool)`

GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceIsTemplate

`func (o *DataInnerIpamPoolData) SetSpaceIsTemplate(v string)`

SetSpaceIsTemplate sets SpaceIsTemplate field to given value.

### HasSpaceIsTemplate

`func (o *DataInnerIpamPoolData) HasSpaceIsTemplate() bool`

HasSpaceIsTemplate returns a boolean if a field has been set.

### GetSpaceName

`func (o *DataInnerIpamPoolData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *DataInnerIpamPoolData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *DataInnerIpamPoolData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *DataInnerIpamPoolData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetPoolStartHostaddr

`func (o *DataInnerIpamPoolData) GetPoolStartHostaddr() string`

GetPoolStartHostaddr returns the PoolStartHostaddr field if non-nil, zero value otherwise.

### GetPoolStartHostaddrOk

`func (o *DataInnerIpamPoolData) GetPoolStartHostaddrOk() (*string, bool)`

GetPoolStartHostaddrOk returns a tuple with the PoolStartHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStartHostaddr

`func (o *DataInnerIpamPoolData) SetPoolStartHostaddr(v string)`

SetPoolStartHostaddr sets PoolStartHostaddr field to given value.

### HasPoolStartHostaddr

`func (o *DataInnerIpamPoolData) HasPoolStartHostaddr() bool`

HasPoolStartHostaddr returns a boolean if a field has been set.

### GetPoolStartIpAddr

`func (o *DataInnerIpamPoolData) GetPoolStartIpAddr() string`

GetPoolStartIpAddr returns the PoolStartIpAddr field if non-nil, zero value otherwise.

### GetPoolStartIpAddrOk

`func (o *DataInnerIpamPoolData) GetPoolStartIpAddrOk() (*string, bool)`

GetPoolStartIpAddrOk returns a tuple with the PoolStartIpAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolStartIpAddr

`func (o *DataInnerIpamPoolData) SetPoolStartIpAddr(v string)`

SetPoolStartIpAddr sets PoolStartIpAddr field to given value.

### HasPoolStartIpAddr

`func (o *DataInnerIpamPoolData) HasPoolStartIpAddr() bool`

HasPoolStartIpAddr returns a boolean if a field has been set.

### GetNetworkClassName

`func (o *DataInnerIpamPoolData) GetNetworkClassName() string`

GetNetworkClassName returns the NetworkClassName field if non-nil, zero value otherwise.

### GetNetworkClassNameOk

`func (o *DataInnerIpamPoolData) GetNetworkClassNameOk() (*string, bool)`

GetNetworkClassNameOk returns a tuple with the NetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassName

`func (o *DataInnerIpamPoolData) SetNetworkClassName(v string)`

SetNetworkClassName sets NetworkClassName field to given value.

### HasNetworkClassName

`func (o *DataInnerIpamPoolData) HasNetworkClassName() bool`

HasNetworkClassName returns a boolean if a field has been set.

### GetNetworkClassParameters

`func (o *DataInnerIpamPoolData) GetNetworkClassParameters() []ApiClassParameterOutputEntry`

GetNetworkClassParameters returns the NetworkClassParameters field if non-nil, zero value otherwise.

### GetNetworkClassParametersOk

`func (o *DataInnerIpamPoolData) GetNetworkClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetNetworkClassParametersOk returns a tuple with the NetworkClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassParameters

`func (o *DataInnerIpamPoolData) SetNetworkClassParameters(v []ApiClassParameterOutputEntry)`

SetNetworkClassParameters sets NetworkClassParameters field to given value.

### HasNetworkClassParameters

`func (o *DataInnerIpamPoolData) HasNetworkClassParameters() bool`

HasNetworkClassParameters returns a boolean if a field has been set.

### GetNetworkEndAddressAddr

`func (o *DataInnerIpamPoolData) GetNetworkEndAddressAddr() string`

GetNetworkEndAddressAddr returns the NetworkEndAddressAddr field if non-nil, zero value otherwise.

### GetNetworkEndAddressAddrOk

`func (o *DataInnerIpamPoolData) GetNetworkEndAddressAddrOk() (*string, bool)`

GetNetworkEndAddressAddrOk returns a tuple with the NetworkEndAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndAddressAddr

`func (o *DataInnerIpamPoolData) SetNetworkEndAddressAddr(v string)`

SetNetworkEndAddressAddr sets NetworkEndAddressAddr field to given value.

### HasNetworkEndAddressAddr

`func (o *DataInnerIpamPoolData) HasNetworkEndAddressAddr() bool`

HasNetworkEndAddressAddr returns a boolean if a field has been set.

### GetNetworkId

`func (o *DataInnerIpamPoolData) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *DataInnerIpamPoolData) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *DataInnerIpamPoolData) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *DataInnerIpamPoolData) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkName

`func (o *DataInnerIpamPoolData) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *DataInnerIpamPoolData) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *DataInnerIpamPoolData) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *DataInnerIpamPoolData) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetNetworkSize

`func (o *DataInnerIpamPoolData) GetNetworkSize() string`

GetNetworkSize returns the NetworkSize field if non-nil, zero value otherwise.

### GetNetworkSizeOk

`func (o *DataInnerIpamPoolData) GetNetworkSizeOk() (*string, bool)`

GetNetworkSizeOk returns a tuple with the NetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSize

`func (o *DataInnerIpamPoolData) SetNetworkSize(v string)`

SetNetworkSize sets NetworkSize field to given value.

### HasNetworkSize

`func (o *DataInnerIpamPoolData) HasNetworkSize() bool`

HasNetworkSize returns a boolean if a field has been set.

### GetNetworkStartAddressAddr

`func (o *DataInnerIpamPoolData) GetNetworkStartAddressAddr() string`

GetNetworkStartAddressAddr returns the NetworkStartAddressAddr field if non-nil, zero value otherwise.

### GetNetworkStartAddressAddrOk

`func (o *DataInnerIpamPoolData) GetNetworkStartAddressAddrOk() (*string, bool)`

GetNetworkStartAddressAddrOk returns a tuple with the NetworkStartAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkStartAddressAddr

`func (o *DataInnerIpamPoolData) SetNetworkStartAddressAddr(v string)`

SetNetworkStartAddressAddr sets NetworkStartAddressAddr field to given value.

### HasNetworkStartAddressAddr

`func (o *DataInnerIpamPoolData) HasNetworkStartAddressAddr() bool`

HasNetworkStartAddressAddr returns a boolean if a field has been set.

### GetPoolOperationDetailsAddedOn

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsAddedOn() string`

GetPoolOperationDetailsAddedOn returns the PoolOperationDetailsAddedOn field if non-nil, zero value otherwise.

### GetPoolOperationDetailsAddedOnOk

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsAddedOnOk() (*string, bool)`

GetPoolOperationDetailsAddedOnOk returns a tuple with the PoolOperationDetailsAddedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolOperationDetailsAddedOn

`func (o *DataInnerIpamPoolData) SetPoolOperationDetailsAddedOn(v string)`

SetPoolOperationDetailsAddedOn sets PoolOperationDetailsAddedOn field to given value.

### HasPoolOperationDetailsAddedOn

`func (o *DataInnerIpamPoolData) HasPoolOperationDetailsAddedOn() bool`

HasPoolOperationDetailsAddedOn returns a boolean if a field has been set.

### GetPoolOperationDetailsCallStack

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsCallStack() string`

GetPoolOperationDetailsCallStack returns the PoolOperationDetailsCallStack field if non-nil, zero value otherwise.

### GetPoolOperationDetailsCallStackOk

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsCallStackOk() (*string, bool)`

GetPoolOperationDetailsCallStackOk returns a tuple with the PoolOperationDetailsCallStack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolOperationDetailsCallStack

`func (o *DataInnerIpamPoolData) SetPoolOperationDetailsCallStack(v string)`

SetPoolOperationDetailsCallStack sets PoolOperationDetailsCallStack field to given value.

### HasPoolOperationDetailsCallStack

`func (o *DataInnerIpamPoolData) HasPoolOperationDetailsCallStack() bool`

HasPoolOperationDetailsCallStack returns a boolean if a field has been set.

### GetPoolOperationDetailsOriginModule

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsOriginModule() string`

GetPoolOperationDetailsOriginModule returns the PoolOperationDetailsOriginModule field if non-nil, zero value otherwise.

### GetPoolOperationDetailsOriginModuleOk

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsOriginModuleOk() (*string, bool)`

GetPoolOperationDetailsOriginModuleOk returns a tuple with the PoolOperationDetailsOriginModule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolOperationDetailsOriginModule

`func (o *DataInnerIpamPoolData) SetPoolOperationDetailsOriginModule(v string)`

SetPoolOperationDetailsOriginModule sets PoolOperationDetailsOriginModule field to given value.

### HasPoolOperationDetailsOriginModule

`func (o *DataInnerIpamPoolData) HasPoolOperationDetailsOriginModule() bool`

HasPoolOperationDetailsOriginModule returns a boolean if a field has been set.

### GetPoolOperationDetailsRequestedBy

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsRequestedBy() string`

GetPoolOperationDetailsRequestedBy returns the PoolOperationDetailsRequestedBy field if non-nil, zero value otherwise.

### GetPoolOperationDetailsRequestedByOk

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsRequestedByOk() (*string, bool)`

GetPoolOperationDetailsRequestedByOk returns a tuple with the PoolOperationDetailsRequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolOperationDetailsRequestedBy

`func (o *DataInnerIpamPoolData) SetPoolOperationDetailsRequestedBy(v string)`

SetPoolOperationDetailsRequestedBy sets PoolOperationDetailsRequestedBy field to given value.

### HasPoolOperationDetailsRequestedBy

`func (o *DataInnerIpamPoolData) HasPoolOperationDetailsRequestedBy() bool`

HasPoolOperationDetailsRequestedBy returns a boolean if a field has been set.

### GetPoolOperationDetailsAddedBy

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsAddedBy() string`

GetPoolOperationDetailsAddedBy returns the PoolOperationDetailsAddedBy field if non-nil, zero value otherwise.

### GetPoolOperationDetailsAddedByOk

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsAddedByOk() (*string, bool)`

GetPoolOperationDetailsAddedByOk returns a tuple with the PoolOperationDetailsAddedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolOperationDetailsAddedBy

`func (o *DataInnerIpamPoolData) SetPoolOperationDetailsAddedBy(v string)`

SetPoolOperationDetailsAddedBy sets PoolOperationDetailsAddedBy field to given value.

### HasPoolOperationDetailsAddedBy

`func (o *DataInnerIpamPoolData) HasPoolOperationDetailsAddedBy() bool`

HasPoolOperationDetailsAddedBy returns a boolean if a field has been set.

### GetPoolOperationDetailsLastUpdatedOn

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsLastUpdatedOn() string`

GetPoolOperationDetailsLastUpdatedOn returns the PoolOperationDetailsLastUpdatedOn field if non-nil, zero value otherwise.

### GetPoolOperationDetailsLastUpdatedOnOk

`func (o *DataInnerIpamPoolData) GetPoolOperationDetailsLastUpdatedOnOk() (*string, bool)`

GetPoolOperationDetailsLastUpdatedOnOk returns a tuple with the PoolOperationDetailsLastUpdatedOn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolOperationDetailsLastUpdatedOn

`func (o *DataInnerIpamPoolData) SetPoolOperationDetailsLastUpdatedOn(v string)`

SetPoolOperationDetailsLastUpdatedOn sets PoolOperationDetailsLastUpdatedOn field to given value.

### HasPoolOperationDetailsLastUpdatedOn

`func (o *DataInnerIpamPoolData) HasPoolOperationDetailsLastUpdatedOn() bool`

HasPoolOperationDetailsLastUpdatedOn returns a boolean if a field has been set.

### GetVlsmBlockId

`func (o *DataInnerIpamPoolData) GetVlsmBlockId() string`

GetVlsmBlockId returns the VlsmBlockId field if non-nil, zero value otherwise.

### GetVlsmBlockIdOk

`func (o *DataInnerIpamPoolData) GetVlsmBlockIdOk() (*string, bool)`

GetVlsmBlockIdOk returns a tuple with the VlsmBlockId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmBlockId

`func (o *DataInnerIpamPoolData) SetVlsmBlockId(v string)`

SetVlsmBlockId sets VlsmBlockId field to given value.

### HasVlsmBlockId

`func (o *DataInnerIpamPoolData) HasVlsmBlockId() bool`

HasVlsmBlockId returns a boolean if a field has been set.

### GetVlsmNetworkId

`func (o *DataInnerIpamPoolData) GetVlsmNetworkId() string`

GetVlsmNetworkId returns the VlsmNetworkId field if non-nil, zero value otherwise.

### GetVlsmNetworkIdOk

`func (o *DataInnerIpamPoolData) GetVlsmNetworkIdOk() (*string, bool)`

GetVlsmNetworkIdOk returns a tuple with the VlsmNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmNetworkId

`func (o *DataInnerIpamPoolData) SetVlsmNetworkId(v string)`

SetVlsmNetworkId sets VlsmNetworkId field to given value.

### HasVlsmNetworkId

`func (o *DataInnerIpamPoolData) HasVlsmNetworkId() bool`

HasVlsmNetworkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


