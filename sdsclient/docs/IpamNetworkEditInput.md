# IpamNetworkEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentNetworkId** | Pointer to **int32** | The database identifier (ID) of an existing IPv4 network you want to set as the parent of the IPv4 network you are editing. You can specify a subnet-type network to set up a network-based VLSM organization. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**NetworkAddr** | Pointer to **string** | The start IP address of the IPv4 network, its first IP address. | [optional] 
**NetworkEndAddr** | Pointer to **string** | The end IP address of the IPv4 network, its last IP address. | [optional] 
**NetworkId** | Pointer to **int32** | The database identifier (ID) of the IPv4 network, a unique numeric key value automatically incremented when you add an IPv4 network. Use the ID to specify which IPv4 network to edit. | [optional] 
**NetworkMask** | Pointer to **string** | The netmask of the IPv4 network. It is expressed in dot-decimal notation and defines the number of addresses the network contains. | [optional] 
**NetworkPrefix** | Pointer to **string** | The prefix of the IPv4 network, an integer that defines the number of addresses the network contains. | [optional] 
**NetworkSize** | Pointer to **int32** | The size of the IPv4 network, the number of IP addresses it contains. | [optional] 
**AllowTreeReparenting** | Pointer to **int32** | A way to allow (1) or prevent (0) changing the parent of the network you are adding. Upon edition of the network, this parameter decides if you can associate it with a different parent network. | [optional] 
**NetworkIsTerminal** | Pointer to **int32** | A way to determine if a network can contain other networks. If set to 1, the network is terminal and cannot contain other subnet-type networks. By essence, block-type networks are non-terminal and are always set to 0. | [optional] 
**NetworkLockNetworkBroadcast** | Pointer to **int32** | A way to prevent &lt;b&gt;(1)&lt;/b&gt; users from assigning the broadcast IP address and network IP address of the network. | [optional] 
**PermitInvalid** | Pointer to **int32** | A way to authorize (&lt;b&gt;1&lt;/b&gt;) IPv4 networks overlapping within a space. | [optional] 
**PermitNoBlock** | Pointer to **int32** | A way to force the creation of an IPv4 subnet-type network. If set to &lt;b&gt;1&lt;/b&gt;, you can create a subnet-type network even if no block-type network matching the start address exists. | [optional] 
**RelativePosition** | Pointer to **int32** | The position of the network within the hierarchy of networks of a VLSM organization. It calculates between &lt;b&gt;0&lt;/b&gt; and &lt;b&gt;n&lt;/b&gt; all the levels of the organization, its behavior depends on the value of the parameter &lt;b&gt;use_reversed_relative_position&lt;/b&gt;:&lt;ul class&#x3D;dashed &gt;&lt;li&gt;&lt;b&gt;use_reversed_relative_position&#x3D;0&lt;/b&gt; where &lt;b&gt;0&lt;/b&gt; indicates a block-type network at the highest level possible, in a space-based organization, it belongs to the top space. The levels increment from &lt;b&gt;0&lt;/b&gt; down to &lt;b&gt;n&lt;/b&gt;, the lowest level you set up, within networks or spaces.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;&lt;b&gt;use_reversed_relative_position&#x3D;1&lt;/b&gt; where &lt;b&gt;1&lt;/b&gt; indicates a network located at the lowest level of the organization, within networks or spaces. The levels increment from &lt;b&gt;0&lt;/b&gt; up to &lt;b&gt;n&lt;/b&gt;, the network at the highest level of the organization.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt; | [optional] 
**RowState** | Pointer to **string** | The object activation status.&lt;ul class&#x3D;dashed &gt;&lt;li&gt;                                                If set to &lt;b&gt;0&lt;/b&gt;, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                If set to &lt;b&gt;1&lt;/b&gt;, the object is enabled and managed.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                If set to &lt;b&gt;2&lt;/b&gt;, the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_enabled&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 
**NetworkLevel** | Pointer to **int32** | The level of the network within the space:&lt;ul class&#x3D;dashed &gt;&lt;li&gt;                                                Set it to &lt;b&gt;0&lt;/b&gt; for a block-type network.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                Set it to a value between &lt;b&gt;1&lt;/b&gt; and &lt;b&gt;n&lt;/b&gt; for a subnet-type network.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt;If you set a value between &lt;b&gt;2&lt;/b&gt; and &lt;b&gt;n&lt;/b&gt;, you are setting a network-based VLSM organization where non terminal subnet-type networks can contain other subnet-type networks. | [optional] 
**NetworkName** | Pointer to **string** | The name of the IPv4 network, each IPv4 network must have a unique name. | [optional] 
**UseReversedRelativePosition** | Pointer to **int32** | A way to determine if the calculation of the parameter &lt;b&gt;relative_position&lt;/b&gt; should start from the top (&lt;b&gt;0&lt;/b&gt;) or the bottom (&lt;b&gt;1&lt;/b&gt;) of the VLSM organization. | [optional] 
**VlanId** | Pointer to **int32** | The database identifier (ID) of the VLAN you want to associate with the network. | [optional] 
**VlsmSpaceId** | Pointer to **int32** | The database identifier (ID) of a VLSM child space of the space specified in &lt;b&gt;site_id&lt;/b&gt;. If you specify an ID, the subnet-type network you are editing is duplicated as a VLSM block-type network in the child space, with the same name but a different ID. This parameter serves the same purpose as &lt;b&gt;vlsm_site_name&lt;/b&gt;. | [optional] 
**VlsmSpaceName** | Pointer to **string** | The name of a VLSM child space of the space specified in &lt;b&gt;site_id&lt;/b&gt;. If you specify a name, the subnet-type network you are editing is duplicated as a VLSM block-type network in the child space, with the same name but a different ID. This parameter serves the same purpose as &lt;b&gt;vlsm_site_id&lt;/b&gt;. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**NetworkClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**NetworkClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamNetworkEditInput

`func NewIpamNetworkEditInput() *IpamNetworkEditInput`

NewIpamNetworkEditInput instantiates a new IpamNetworkEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamNetworkEditInputWithDefaults

`func NewIpamNetworkEditInputWithDefaults() *IpamNetworkEditInput`

NewIpamNetworkEditInputWithDefaults instantiates a new IpamNetworkEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentNetworkId

`func (o *IpamNetworkEditInput) GetParentNetworkId() int32`

GetParentNetworkId returns the ParentNetworkId field if non-nil, zero value otherwise.

### GetParentNetworkIdOk

`func (o *IpamNetworkEditInput) GetParentNetworkIdOk() (*int32, bool)`

GetParentNetworkIdOk returns a tuple with the ParentNetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetworkId

`func (o *IpamNetworkEditInput) SetParentNetworkId(v int32)`

SetParentNetworkId sets ParentNetworkId field to given value.

### HasParentNetworkId

`func (o *IpamNetworkEditInput) HasParentNetworkId() bool`

HasParentNetworkId returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamNetworkEditInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamNetworkEditInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamNetworkEditInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamNetworkEditInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamNetworkEditInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamNetworkEditInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamNetworkEditInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamNetworkEditInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetNetworkAddr

`func (o *IpamNetworkEditInput) GetNetworkAddr() string`

GetNetworkAddr returns the NetworkAddr field if non-nil, zero value otherwise.

### GetNetworkAddrOk

`func (o *IpamNetworkEditInput) GetNetworkAddrOk() (*string, bool)`

GetNetworkAddrOk returns a tuple with the NetworkAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkAddr

`func (o *IpamNetworkEditInput) SetNetworkAddr(v string)`

SetNetworkAddr sets NetworkAddr field to given value.

### HasNetworkAddr

`func (o *IpamNetworkEditInput) HasNetworkAddr() bool`

HasNetworkAddr returns a boolean if a field has been set.

### GetNetworkEndAddr

`func (o *IpamNetworkEditInput) GetNetworkEndAddr() string`

GetNetworkEndAddr returns the NetworkEndAddr field if non-nil, zero value otherwise.

### GetNetworkEndAddrOk

`func (o *IpamNetworkEditInput) GetNetworkEndAddrOk() (*string, bool)`

GetNetworkEndAddrOk returns a tuple with the NetworkEndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkEndAddr

`func (o *IpamNetworkEditInput) SetNetworkEndAddr(v string)`

SetNetworkEndAddr sets NetworkEndAddr field to given value.

### HasNetworkEndAddr

`func (o *IpamNetworkEditInput) HasNetworkEndAddr() bool`

HasNetworkEndAddr returns a boolean if a field has been set.

### GetNetworkId

`func (o *IpamNetworkEditInput) GetNetworkId() int32`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *IpamNetworkEditInput) GetNetworkIdOk() (*int32, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *IpamNetworkEditInput) SetNetworkId(v int32)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *IpamNetworkEditInput) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetNetworkMask

`func (o *IpamNetworkEditInput) GetNetworkMask() string`

GetNetworkMask returns the NetworkMask field if non-nil, zero value otherwise.

### GetNetworkMaskOk

`func (o *IpamNetworkEditInput) GetNetworkMaskOk() (*string, bool)`

GetNetworkMaskOk returns a tuple with the NetworkMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkMask

`func (o *IpamNetworkEditInput) SetNetworkMask(v string)`

SetNetworkMask sets NetworkMask field to given value.

### HasNetworkMask

`func (o *IpamNetworkEditInput) HasNetworkMask() bool`

HasNetworkMask returns a boolean if a field has been set.

### GetNetworkPrefix

`func (o *IpamNetworkEditInput) GetNetworkPrefix() string`

GetNetworkPrefix returns the NetworkPrefix field if non-nil, zero value otherwise.

### GetNetworkPrefixOk

`func (o *IpamNetworkEditInput) GetNetworkPrefixOk() (*string, bool)`

GetNetworkPrefixOk returns a tuple with the NetworkPrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPrefix

`func (o *IpamNetworkEditInput) SetNetworkPrefix(v string)`

SetNetworkPrefix sets NetworkPrefix field to given value.

### HasNetworkPrefix

`func (o *IpamNetworkEditInput) HasNetworkPrefix() bool`

HasNetworkPrefix returns a boolean if a field has been set.

### GetNetworkSize

`func (o *IpamNetworkEditInput) GetNetworkSize() int32`

GetNetworkSize returns the NetworkSize field if non-nil, zero value otherwise.

### GetNetworkSizeOk

`func (o *IpamNetworkEditInput) GetNetworkSizeOk() (*int32, bool)`

GetNetworkSizeOk returns a tuple with the NetworkSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkSize

`func (o *IpamNetworkEditInput) SetNetworkSize(v int32)`

SetNetworkSize sets NetworkSize field to given value.

### HasNetworkSize

`func (o *IpamNetworkEditInput) HasNetworkSize() bool`

HasNetworkSize returns a boolean if a field has been set.

### GetAllowTreeReparenting

`func (o *IpamNetworkEditInput) GetAllowTreeReparenting() int32`

GetAllowTreeReparenting returns the AllowTreeReparenting field if non-nil, zero value otherwise.

### GetAllowTreeReparentingOk

`func (o *IpamNetworkEditInput) GetAllowTreeReparentingOk() (*int32, bool)`

GetAllowTreeReparentingOk returns a tuple with the AllowTreeReparenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTreeReparenting

`func (o *IpamNetworkEditInput) SetAllowTreeReparenting(v int32)`

SetAllowTreeReparenting sets AllowTreeReparenting field to given value.

### HasAllowTreeReparenting

`func (o *IpamNetworkEditInput) HasAllowTreeReparenting() bool`

HasAllowTreeReparenting returns a boolean if a field has been set.

### GetNetworkIsTerminal

`func (o *IpamNetworkEditInput) GetNetworkIsTerminal() int32`

GetNetworkIsTerminal returns the NetworkIsTerminal field if non-nil, zero value otherwise.

### GetNetworkIsTerminalOk

`func (o *IpamNetworkEditInput) GetNetworkIsTerminalOk() (*int32, bool)`

GetNetworkIsTerminalOk returns a tuple with the NetworkIsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkIsTerminal

`func (o *IpamNetworkEditInput) SetNetworkIsTerminal(v int32)`

SetNetworkIsTerminal sets NetworkIsTerminal field to given value.

### HasNetworkIsTerminal

`func (o *IpamNetworkEditInput) HasNetworkIsTerminal() bool`

HasNetworkIsTerminal returns a boolean if a field has been set.

### GetNetworkLockNetworkBroadcast

`func (o *IpamNetworkEditInput) GetNetworkLockNetworkBroadcast() int32`

GetNetworkLockNetworkBroadcast returns the NetworkLockNetworkBroadcast field if non-nil, zero value otherwise.

### GetNetworkLockNetworkBroadcastOk

`func (o *IpamNetworkEditInput) GetNetworkLockNetworkBroadcastOk() (*int32, bool)`

GetNetworkLockNetworkBroadcastOk returns a tuple with the NetworkLockNetworkBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLockNetworkBroadcast

`func (o *IpamNetworkEditInput) SetNetworkLockNetworkBroadcast(v int32)`

SetNetworkLockNetworkBroadcast sets NetworkLockNetworkBroadcast field to given value.

### HasNetworkLockNetworkBroadcast

`func (o *IpamNetworkEditInput) HasNetworkLockNetworkBroadcast() bool`

HasNetworkLockNetworkBroadcast returns a boolean if a field has been set.

### GetPermitInvalid

`func (o *IpamNetworkEditInput) GetPermitInvalid() int32`

GetPermitInvalid returns the PermitInvalid field if non-nil, zero value otherwise.

### GetPermitInvalidOk

`func (o *IpamNetworkEditInput) GetPermitInvalidOk() (*int32, bool)`

GetPermitInvalidOk returns a tuple with the PermitInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitInvalid

`func (o *IpamNetworkEditInput) SetPermitInvalid(v int32)`

SetPermitInvalid sets PermitInvalid field to given value.

### HasPermitInvalid

`func (o *IpamNetworkEditInput) HasPermitInvalid() bool`

HasPermitInvalid returns a boolean if a field has been set.

### GetPermitNoBlock

`func (o *IpamNetworkEditInput) GetPermitNoBlock() int32`

GetPermitNoBlock returns the PermitNoBlock field if non-nil, zero value otherwise.

### GetPermitNoBlockOk

`func (o *IpamNetworkEditInput) GetPermitNoBlockOk() (*int32, bool)`

GetPermitNoBlockOk returns a tuple with the PermitNoBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitNoBlock

`func (o *IpamNetworkEditInput) SetPermitNoBlock(v int32)`

SetPermitNoBlock sets PermitNoBlock field to given value.

### HasPermitNoBlock

`func (o *IpamNetworkEditInput) HasPermitNoBlock() bool`

HasPermitNoBlock returns a boolean if a field has been set.

### GetRelativePosition

`func (o *IpamNetworkEditInput) GetRelativePosition() int32`

GetRelativePosition returns the RelativePosition field if non-nil, zero value otherwise.

### GetRelativePositionOk

`func (o *IpamNetworkEditInput) GetRelativePositionOk() (*int32, bool)`

GetRelativePositionOk returns a tuple with the RelativePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePosition

`func (o *IpamNetworkEditInput) SetRelativePosition(v int32)`

SetRelativePosition sets RelativePosition field to given value.

### HasRelativePosition

`func (o *IpamNetworkEditInput) HasRelativePosition() bool`

HasRelativePosition returns a boolean if a field has been set.

### GetRowState

`func (o *IpamNetworkEditInput) GetRowState() string`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *IpamNetworkEditInput) GetRowStateOk() (*string, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *IpamNetworkEditInput) SetRowState(v string)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *IpamNetworkEditInput) HasRowState() bool`

HasRowState returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *IpamNetworkEditInput) GetNetworkLevel() int32`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *IpamNetworkEditInput) GetNetworkLevelOk() (*int32, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *IpamNetworkEditInput) SetNetworkLevel(v int32)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *IpamNetworkEditInput) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetNetworkName

`func (o *IpamNetworkEditInput) GetNetworkName() string`

GetNetworkName returns the NetworkName field if non-nil, zero value otherwise.

### GetNetworkNameOk

`func (o *IpamNetworkEditInput) GetNetworkNameOk() (*string, bool)`

GetNetworkNameOk returns a tuple with the NetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkName

`func (o *IpamNetworkEditInput) SetNetworkName(v string)`

SetNetworkName sets NetworkName field to given value.

### HasNetworkName

`func (o *IpamNetworkEditInput) HasNetworkName() bool`

HasNetworkName returns a boolean if a field has been set.

### GetUseReversedRelativePosition

`func (o *IpamNetworkEditInput) GetUseReversedRelativePosition() int32`

GetUseReversedRelativePosition returns the UseReversedRelativePosition field if non-nil, zero value otherwise.

### GetUseReversedRelativePositionOk

`func (o *IpamNetworkEditInput) GetUseReversedRelativePositionOk() (*int32, bool)`

GetUseReversedRelativePositionOk returns a tuple with the UseReversedRelativePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReversedRelativePosition

`func (o *IpamNetworkEditInput) SetUseReversedRelativePosition(v int32)`

SetUseReversedRelativePosition sets UseReversedRelativePosition field to given value.

### HasUseReversedRelativePosition

`func (o *IpamNetworkEditInput) HasUseReversedRelativePosition() bool`

HasUseReversedRelativePosition returns a boolean if a field has been set.

### GetVlanId

`func (o *IpamNetworkEditInput) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *IpamNetworkEditInput) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *IpamNetworkEditInput) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *IpamNetworkEditInput) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlsmSpaceId

`func (o *IpamNetworkEditInput) GetVlsmSpaceId() int32`

GetVlsmSpaceId returns the VlsmSpaceId field if non-nil, zero value otherwise.

### GetVlsmSpaceIdOk

`func (o *IpamNetworkEditInput) GetVlsmSpaceIdOk() (*int32, bool)`

GetVlsmSpaceIdOk returns a tuple with the VlsmSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceId

`func (o *IpamNetworkEditInput) SetVlsmSpaceId(v int32)`

SetVlsmSpaceId sets VlsmSpaceId field to given value.

### HasVlsmSpaceId

`func (o *IpamNetworkEditInput) HasVlsmSpaceId() bool`

HasVlsmSpaceId returns a boolean if a field has been set.

### GetVlsmSpaceName

`func (o *IpamNetworkEditInput) GetVlsmSpaceName() string`

GetVlsmSpaceName returns the VlsmSpaceName field if non-nil, zero value otherwise.

### GetVlsmSpaceNameOk

`func (o *IpamNetworkEditInput) GetVlsmSpaceNameOk() (*string, bool)`

GetVlsmSpaceNameOk returns a tuple with the VlsmSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceName

`func (o *IpamNetworkEditInput) SetVlsmSpaceName(v string)`

SetVlsmSpaceName sets VlsmSpaceName field to given value.

### HasVlsmSpaceName

`func (o *IpamNetworkEditInput) HasVlsmSpaceName() bool`

HasVlsmSpaceName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *IpamNetworkEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *IpamNetworkEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *IpamNetworkEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *IpamNetworkEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetNetworkClassParameters

`func (o *IpamNetworkEditInput) GetNetworkClassParameters() []ApiClassParameterInputEntry`

GetNetworkClassParameters returns the NetworkClassParameters field if non-nil, zero value otherwise.

### GetNetworkClassParametersOk

`func (o *IpamNetworkEditInput) GetNetworkClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetNetworkClassParametersOk returns a tuple with the NetworkClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassParameters

`func (o *IpamNetworkEditInput) SetNetworkClassParameters(v []ApiClassParameterInputEntry)`

SetNetworkClassParameters sets NetworkClassParameters field to given value.

### HasNetworkClassParameters

`func (o *IpamNetworkEditInput) HasNetworkClassParameters() bool`

HasNetworkClassParameters returns a boolean if a field has been set.

### GetNetworkClassName

`func (o *IpamNetworkEditInput) GetNetworkClassName() string`

GetNetworkClassName returns the NetworkClassName field if non-nil, zero value otherwise.

### GetNetworkClassNameOk

`func (o *IpamNetworkEditInput) GetNetworkClassNameOk() (*string, bool)`

GetNetworkClassNameOk returns a tuple with the NetworkClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkClassName

`func (o *IpamNetworkEditInput) SetNetworkClassName(v string)`

SetNetworkClassName sets NetworkClassName field to given value.

### HasNetworkClassName

`func (o *IpamNetworkEditInput) HasNetworkClassName() bool`

HasNetworkClassName returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamNetworkEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamNetworkEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamNetworkEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamNetworkEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


