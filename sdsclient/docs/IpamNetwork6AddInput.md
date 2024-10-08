# IpamNetwork6AddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ParentNetwork6Id** | Pointer to **int32** | The database identifier (ID) of an existing IPv6 network you want to set as the parent of the IPv6 network you are adding. You can specify a subnet-type network to set up a network-based VLSM organization. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**Network6Addr** | Pointer to **string** | The start IP address of the IPv6 network, its first IP address. | [optional] 
**Network6EndAddr** | Pointer to **string** | The end IP address of the IPv6 network, its last IP address. | [optional] 
**Network6Prefix** | Pointer to **string** | The prefix of the IPv6 network, an integer that defines the number of address the network contains. | [optional] 
**AllowBlock6Creation** | Pointer to **int32** | Internal use. Not documented. | [optional] 
**AllowTreeReparenting** | Pointer to **int32** | A way to allow (1) or prevent (0) changing the parent of the network you are adding. Upon editing of the network, this parameter decides if you can associate it with a different parent network. | [optional] 
**Network6IsTerminal** | Pointer to **int32** | A way to determine if a network can contain other networks. If set to 1, the network is terminal and cannot contain other subnet-type networks. By essence, block-type networks are non-terminal and are always set to 0. | [optional] 
**Network6LockNetworkBroadcast** | Pointer to **int32** | A way to prevent &lt;b&gt;(1)&lt;/b&gt; users from assigning the broadcast IP address and network IP address of the network. | [optional] 
**PermitInvalid** | Pointer to **int32** | A way to authorize (&lt;b&gt;1&lt;/b&gt;) IPv6 networks overlapping within a space. | [optional] 
**PermitNoBlock6** | Pointer to **int32** | A way to force the creation of an IPv6 subnet-type network. If set to 1, you can create a subnet-type network even if no block-type network matching the start address exists. | [optional] 
**RelativePosition** | Pointer to **int32** | The position of the network within the hierarchy of networks of a VLSM organization. It calculates between &lt;b&gt;0&lt;/b&gt; and &lt;b&gt;n&lt;/b&gt; all the levels of the organization, its behavior depends on the value of the parameter &lt;b&gt;use_reversed_relative_position&lt;/b&gt;:&lt;ul&gt;&lt;li&gt; &lt;b&gt;use_reversed_relative_position&#x3D;0&lt;/b&gt; where &lt;b&gt;0&lt;/b&gt; indicates a block-type network at the highest level possible, in a space-based organization, it belongs to the top space. The levels increment from &lt;b&gt;0&lt;/b&gt; down to &lt;b&gt;n&lt;/b&gt;, the lowest level you set up, within networks or spaces.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; &lt;b&gt;use_reversed_relative_position&#x3D;1&lt;/b&gt; where &lt;b&gt;1&lt;/b&gt; indicates a network located at the lowest level of the organization, within networks or spaces. The levels increment from &lt;b&gt;0&lt;/b&gt; up to &lt;b&gt;n&lt;/b&gt;, the network at the highest level of the organization.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt; | [optional] 
**RowState** | Pointer to **string** | The object activation status.&lt;ul&gt;&lt;li&gt; If set to &lt;b&gt;0&lt;/b&gt;, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; If set to &lt;b&gt;1&lt;/b&gt;, the object is enabled and managed.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; If set to &lt;b&gt;2&lt;/b&gt;, the object is unmanaged, disabled or both depending on the context.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt;By default, &lt;b&gt;row_state&lt;/b&gt; is set to &lt;b&gt;1&lt;/b&gt; when an object is created. | [optional] 
**Network6Name** | Pointer to **string** | The name of the IPv6 network, each IPv6 network must have a unique name. | [optional] 
**NetworkLevel** | Pointer to **int32** | The level of the network within the space:&lt;ul&gt;&lt;li&gt; Set it to &lt;b&gt;0&lt;/b&gt; for a block-type network.&lt;br/&gt;&lt;/li&gt;&lt;li&gt; Set it to a value between &lt;b&gt;1&lt;/b&gt; and &lt;b&gt;n&lt;/b&gt; for a subnet-type network.&lt;br/&gt;&lt;/li&gt;&lt;/ul&gt;If you set a value between &lt;b&gt;2&lt;/b&gt; and &lt;b&gt;n&lt;/b&gt;, you are setting a network-based VLSM organization where non terminal subnet-type networks can contain other subnet-type networks. | [optional] 
**UseReversedRelativePosition** | Pointer to **int32** | A way to determine if the calculation of the parameter &lt;b&gt;relative_position&lt;/b&gt; should start from the top (&lt;b&gt;0&lt;/b&gt;) or the bottom (&lt;b&gt;1&lt;/b&gt;) of the VLSM organization. | [optional] 
**VlanId** | Pointer to **int32** | The database identifier (ID) of the VLAN you want to associate with the network. | [optional] 
**VlsmSpaceId** | Pointer to **int32** | The database identifier (ID) of a VLSM child space of the space specified in &lt;b&gt;space_id&lt;/b&gt;. If you specify an ID, the subnet-type network you are adding is duplicated as a VLSM block-type network in the child space, with the same name but a different ID. This parameter serves the same purpose as &lt;b&gt;vlsm_space_name&lt;/b&gt;. | [optional] 
**VlsmSpaceName** | Pointer to **string** | The name of a VLSM child space of the space specified in &lt;b&gt;space_id&lt;/b&gt;. If you specify a name, the subnet-type network you are adding is duplicated as a VLSM block-type network in the child space, with the same name but a different ID. This parameter serves the same purpose as &lt;b&gt;vlsm_space_id&lt;/b&gt;. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**Network6ClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Network6ClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamNetwork6AddInput

`func NewIpamNetwork6AddInput() *IpamNetwork6AddInput`

NewIpamNetwork6AddInput instantiates a new IpamNetwork6AddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamNetwork6AddInputWithDefaults

`func NewIpamNetwork6AddInputWithDefaults() *IpamNetwork6AddInput`

NewIpamNetwork6AddInputWithDefaults instantiates a new IpamNetwork6AddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParentNetwork6Id

`func (o *IpamNetwork6AddInput) GetParentNetwork6Id() int32`

GetParentNetwork6Id returns the ParentNetwork6Id field if non-nil, zero value otherwise.

### GetParentNetwork6IdOk

`func (o *IpamNetwork6AddInput) GetParentNetwork6IdOk() (*int32, bool)`

GetParentNetwork6IdOk returns a tuple with the ParentNetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentNetwork6Id

`func (o *IpamNetwork6AddInput) SetParentNetwork6Id(v int32)`

SetParentNetwork6Id sets ParentNetwork6Id field to given value.

### HasParentNetwork6Id

`func (o *IpamNetwork6AddInput) HasParentNetwork6Id() bool`

HasParentNetwork6Id returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamNetwork6AddInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamNetwork6AddInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamNetwork6AddInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamNetwork6AddInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamNetwork6AddInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamNetwork6AddInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamNetwork6AddInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamNetwork6AddInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetNetwork6Addr

`func (o *IpamNetwork6AddInput) GetNetwork6Addr() string`

GetNetwork6Addr returns the Network6Addr field if non-nil, zero value otherwise.

### GetNetwork6AddrOk

`func (o *IpamNetwork6AddInput) GetNetwork6AddrOk() (*string, bool)`

GetNetwork6AddrOk returns a tuple with the Network6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Addr

`func (o *IpamNetwork6AddInput) SetNetwork6Addr(v string)`

SetNetwork6Addr sets Network6Addr field to given value.

### HasNetwork6Addr

`func (o *IpamNetwork6AddInput) HasNetwork6Addr() bool`

HasNetwork6Addr returns a boolean if a field has been set.

### GetNetwork6EndAddr

`func (o *IpamNetwork6AddInput) GetNetwork6EndAddr() string`

GetNetwork6EndAddr returns the Network6EndAddr field if non-nil, zero value otherwise.

### GetNetwork6EndAddrOk

`func (o *IpamNetwork6AddInput) GetNetwork6EndAddrOk() (*string, bool)`

GetNetwork6EndAddrOk returns a tuple with the Network6EndAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6EndAddr

`func (o *IpamNetwork6AddInput) SetNetwork6EndAddr(v string)`

SetNetwork6EndAddr sets Network6EndAddr field to given value.

### HasNetwork6EndAddr

`func (o *IpamNetwork6AddInput) HasNetwork6EndAddr() bool`

HasNetwork6EndAddr returns a boolean if a field has been set.

### GetNetwork6Prefix

`func (o *IpamNetwork6AddInput) GetNetwork6Prefix() string`

GetNetwork6Prefix returns the Network6Prefix field if non-nil, zero value otherwise.

### GetNetwork6PrefixOk

`func (o *IpamNetwork6AddInput) GetNetwork6PrefixOk() (*string, bool)`

GetNetwork6PrefixOk returns a tuple with the Network6Prefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Prefix

`func (o *IpamNetwork6AddInput) SetNetwork6Prefix(v string)`

SetNetwork6Prefix sets Network6Prefix field to given value.

### HasNetwork6Prefix

`func (o *IpamNetwork6AddInput) HasNetwork6Prefix() bool`

HasNetwork6Prefix returns a boolean if a field has been set.

### GetAllowBlock6Creation

`func (o *IpamNetwork6AddInput) GetAllowBlock6Creation() int32`

GetAllowBlock6Creation returns the AllowBlock6Creation field if non-nil, zero value otherwise.

### GetAllowBlock6CreationOk

`func (o *IpamNetwork6AddInput) GetAllowBlock6CreationOk() (*int32, bool)`

GetAllowBlock6CreationOk returns a tuple with the AllowBlock6Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowBlock6Creation

`func (o *IpamNetwork6AddInput) SetAllowBlock6Creation(v int32)`

SetAllowBlock6Creation sets AllowBlock6Creation field to given value.

### HasAllowBlock6Creation

`func (o *IpamNetwork6AddInput) HasAllowBlock6Creation() bool`

HasAllowBlock6Creation returns a boolean if a field has been set.

### GetAllowTreeReparenting

`func (o *IpamNetwork6AddInput) GetAllowTreeReparenting() int32`

GetAllowTreeReparenting returns the AllowTreeReparenting field if non-nil, zero value otherwise.

### GetAllowTreeReparentingOk

`func (o *IpamNetwork6AddInput) GetAllowTreeReparentingOk() (*int32, bool)`

GetAllowTreeReparentingOk returns a tuple with the AllowTreeReparenting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowTreeReparenting

`func (o *IpamNetwork6AddInput) SetAllowTreeReparenting(v int32)`

SetAllowTreeReparenting sets AllowTreeReparenting field to given value.

### HasAllowTreeReparenting

`func (o *IpamNetwork6AddInput) HasAllowTreeReparenting() bool`

HasAllowTreeReparenting returns a boolean if a field has been set.

### GetNetwork6IsTerminal

`func (o *IpamNetwork6AddInput) GetNetwork6IsTerminal() int32`

GetNetwork6IsTerminal returns the Network6IsTerminal field if non-nil, zero value otherwise.

### GetNetwork6IsTerminalOk

`func (o *IpamNetwork6AddInput) GetNetwork6IsTerminalOk() (*int32, bool)`

GetNetwork6IsTerminalOk returns a tuple with the Network6IsTerminal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6IsTerminal

`func (o *IpamNetwork6AddInput) SetNetwork6IsTerminal(v int32)`

SetNetwork6IsTerminal sets Network6IsTerminal field to given value.

### HasNetwork6IsTerminal

`func (o *IpamNetwork6AddInput) HasNetwork6IsTerminal() bool`

HasNetwork6IsTerminal returns a boolean if a field has been set.

### GetNetwork6LockNetworkBroadcast

`func (o *IpamNetwork6AddInput) GetNetwork6LockNetworkBroadcast() int32`

GetNetwork6LockNetworkBroadcast returns the Network6LockNetworkBroadcast field if non-nil, zero value otherwise.

### GetNetwork6LockNetworkBroadcastOk

`func (o *IpamNetwork6AddInput) GetNetwork6LockNetworkBroadcastOk() (*int32, bool)`

GetNetwork6LockNetworkBroadcastOk returns a tuple with the Network6LockNetworkBroadcast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6LockNetworkBroadcast

`func (o *IpamNetwork6AddInput) SetNetwork6LockNetworkBroadcast(v int32)`

SetNetwork6LockNetworkBroadcast sets Network6LockNetworkBroadcast field to given value.

### HasNetwork6LockNetworkBroadcast

`func (o *IpamNetwork6AddInput) HasNetwork6LockNetworkBroadcast() bool`

HasNetwork6LockNetworkBroadcast returns a boolean if a field has been set.

### GetPermitInvalid

`func (o *IpamNetwork6AddInput) GetPermitInvalid() int32`

GetPermitInvalid returns the PermitInvalid field if non-nil, zero value otherwise.

### GetPermitInvalidOk

`func (o *IpamNetwork6AddInput) GetPermitInvalidOk() (*int32, bool)`

GetPermitInvalidOk returns a tuple with the PermitInvalid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitInvalid

`func (o *IpamNetwork6AddInput) SetPermitInvalid(v int32)`

SetPermitInvalid sets PermitInvalid field to given value.

### HasPermitInvalid

`func (o *IpamNetwork6AddInput) HasPermitInvalid() bool`

HasPermitInvalid returns a boolean if a field has been set.

### GetPermitNoBlock6

`func (o *IpamNetwork6AddInput) GetPermitNoBlock6() int32`

GetPermitNoBlock6 returns the PermitNoBlock6 field if non-nil, zero value otherwise.

### GetPermitNoBlock6Ok

`func (o *IpamNetwork6AddInput) GetPermitNoBlock6Ok() (*int32, bool)`

GetPermitNoBlock6Ok returns a tuple with the PermitNoBlock6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermitNoBlock6

`func (o *IpamNetwork6AddInput) SetPermitNoBlock6(v int32)`

SetPermitNoBlock6 sets PermitNoBlock6 field to given value.

### HasPermitNoBlock6

`func (o *IpamNetwork6AddInput) HasPermitNoBlock6() bool`

HasPermitNoBlock6 returns a boolean if a field has been set.

### GetRelativePosition

`func (o *IpamNetwork6AddInput) GetRelativePosition() int32`

GetRelativePosition returns the RelativePosition field if non-nil, zero value otherwise.

### GetRelativePositionOk

`func (o *IpamNetwork6AddInput) GetRelativePositionOk() (*int32, bool)`

GetRelativePositionOk returns a tuple with the RelativePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelativePosition

`func (o *IpamNetwork6AddInput) SetRelativePosition(v int32)`

SetRelativePosition sets RelativePosition field to given value.

### HasRelativePosition

`func (o *IpamNetwork6AddInput) HasRelativePosition() bool`

HasRelativePosition returns a boolean if a field has been set.

### GetRowState

`func (o *IpamNetwork6AddInput) GetRowState() string`

GetRowState returns the RowState field if non-nil, zero value otherwise.

### GetRowStateOk

`func (o *IpamNetwork6AddInput) GetRowStateOk() (*string, bool)`

GetRowStateOk returns a tuple with the RowState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRowState

`func (o *IpamNetwork6AddInput) SetRowState(v string)`

SetRowState sets RowState field to given value.

### HasRowState

`func (o *IpamNetwork6AddInput) HasRowState() bool`

HasRowState returns a boolean if a field has been set.

### GetNetwork6Name

`func (o *IpamNetwork6AddInput) GetNetwork6Name() string`

GetNetwork6Name returns the Network6Name field if non-nil, zero value otherwise.

### GetNetwork6NameOk

`func (o *IpamNetwork6AddInput) GetNetwork6NameOk() (*string, bool)`

GetNetwork6NameOk returns a tuple with the Network6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6Name

`func (o *IpamNetwork6AddInput) SetNetwork6Name(v string)`

SetNetwork6Name sets Network6Name field to given value.

### HasNetwork6Name

`func (o *IpamNetwork6AddInput) HasNetwork6Name() bool`

HasNetwork6Name returns a boolean if a field has been set.

### GetNetworkLevel

`func (o *IpamNetwork6AddInput) GetNetworkLevel() int32`

GetNetworkLevel returns the NetworkLevel field if non-nil, zero value otherwise.

### GetNetworkLevelOk

`func (o *IpamNetwork6AddInput) GetNetworkLevelOk() (*int32, bool)`

GetNetworkLevelOk returns a tuple with the NetworkLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkLevel

`func (o *IpamNetwork6AddInput) SetNetworkLevel(v int32)`

SetNetworkLevel sets NetworkLevel field to given value.

### HasNetworkLevel

`func (o *IpamNetwork6AddInput) HasNetworkLevel() bool`

HasNetworkLevel returns a boolean if a field has been set.

### GetUseReversedRelativePosition

`func (o *IpamNetwork6AddInput) GetUseReversedRelativePosition() int32`

GetUseReversedRelativePosition returns the UseReversedRelativePosition field if non-nil, zero value otherwise.

### GetUseReversedRelativePositionOk

`func (o *IpamNetwork6AddInput) GetUseReversedRelativePositionOk() (*int32, bool)`

GetUseReversedRelativePositionOk returns a tuple with the UseReversedRelativePosition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseReversedRelativePosition

`func (o *IpamNetwork6AddInput) SetUseReversedRelativePosition(v int32)`

SetUseReversedRelativePosition sets UseReversedRelativePosition field to given value.

### HasUseReversedRelativePosition

`func (o *IpamNetwork6AddInput) HasUseReversedRelativePosition() bool`

HasUseReversedRelativePosition returns a boolean if a field has been set.

### GetVlanId

`func (o *IpamNetwork6AddInput) GetVlanId() int32`

GetVlanId returns the VlanId field if non-nil, zero value otherwise.

### GetVlanIdOk

`func (o *IpamNetwork6AddInput) GetVlanIdOk() (*int32, bool)`

GetVlanIdOk returns a tuple with the VlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlanId

`func (o *IpamNetwork6AddInput) SetVlanId(v int32)`

SetVlanId sets VlanId field to given value.

### HasVlanId

`func (o *IpamNetwork6AddInput) HasVlanId() bool`

HasVlanId returns a boolean if a field has been set.

### GetVlsmSpaceId

`func (o *IpamNetwork6AddInput) GetVlsmSpaceId() int32`

GetVlsmSpaceId returns the VlsmSpaceId field if non-nil, zero value otherwise.

### GetVlsmSpaceIdOk

`func (o *IpamNetwork6AddInput) GetVlsmSpaceIdOk() (*int32, bool)`

GetVlsmSpaceIdOk returns a tuple with the VlsmSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceId

`func (o *IpamNetwork6AddInput) SetVlsmSpaceId(v int32)`

SetVlsmSpaceId sets VlsmSpaceId field to given value.

### HasVlsmSpaceId

`func (o *IpamNetwork6AddInput) HasVlsmSpaceId() bool`

HasVlsmSpaceId returns a boolean if a field has been set.

### GetVlsmSpaceName

`func (o *IpamNetwork6AddInput) GetVlsmSpaceName() string`

GetVlsmSpaceName returns the VlsmSpaceName field if non-nil, zero value otherwise.

### GetVlsmSpaceNameOk

`func (o *IpamNetwork6AddInput) GetVlsmSpaceNameOk() (*string, bool)`

GetVlsmSpaceNameOk returns a tuple with the VlsmSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVlsmSpaceName

`func (o *IpamNetwork6AddInput) SetVlsmSpaceName(v string)`

SetVlsmSpaceName sets VlsmSpaceName field to given value.

### HasVlsmSpaceName

`func (o *IpamNetwork6AddInput) HasVlsmSpaceName() bool`

HasVlsmSpaceName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *IpamNetwork6AddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *IpamNetwork6AddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *IpamNetwork6AddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *IpamNetwork6AddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetNetwork6ClassParameters

`func (o *IpamNetwork6AddInput) GetNetwork6ClassParameters() []ApiClassParameterInputEntry`

GetNetwork6ClassParameters returns the Network6ClassParameters field if non-nil, zero value otherwise.

### GetNetwork6ClassParametersOk

`func (o *IpamNetwork6AddInput) GetNetwork6ClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetNetwork6ClassParametersOk returns a tuple with the Network6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6ClassParameters

`func (o *IpamNetwork6AddInput) SetNetwork6ClassParameters(v []ApiClassParameterInputEntry)`

SetNetwork6ClassParameters sets Network6ClassParameters field to given value.

### HasNetwork6ClassParameters

`func (o *IpamNetwork6AddInput) HasNetwork6ClassParameters() bool`

HasNetwork6ClassParameters returns a boolean if a field has been set.

### GetNetwork6ClassName

`func (o *IpamNetwork6AddInput) GetNetwork6ClassName() string`

GetNetwork6ClassName returns the Network6ClassName field if non-nil, zero value otherwise.

### GetNetwork6ClassNameOk

`func (o *IpamNetwork6AddInput) GetNetwork6ClassNameOk() (*string, bool)`

GetNetwork6ClassNameOk returns a tuple with the Network6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork6ClassName

`func (o *IpamNetwork6AddInput) SetNetwork6ClassName(v string)`

SetNetwork6ClassName sets Network6ClassName field to given value.

### HasNetwork6ClassName

`func (o *IpamNetwork6AddInput) HasNetwork6ClassName() bool`

HasNetwork6ClassName returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamNetwork6AddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamNetwork6AddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamNetwork6AddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamNetwork6AddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


