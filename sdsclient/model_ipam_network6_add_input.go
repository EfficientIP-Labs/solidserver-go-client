/*
SOLIDserver API

OpenAPI 3.0.2 API definition for SOLIDserver service from EfficientIP.<p>Copyright © 2000-2024 EfficientIP</p><p><em>All specifications and information regarding the products in this document are subject to change without notice and should not be construed as a commitment by EfficientIP. EfficientIP assumes no responsibility or liability for any mistakes or inaccuracies that may appear in this document. All statements and recommendations in this document are believed to be accurate but are presented without warranty. Users must take full responsibility for their application of any product.</em></p><p><em>This document aims at detailing EfficientIP proprietary solutions. As our solutions rely on several third-party products, created by other companies or organizations, it may redirect readers to third-party websites and documentation for further information. In such a case, EfficientIP cannot be liable or expected to provide said information on products they do maintain or created.</em></p><p>Generated (Friday 4th of October 2024 03:41:11 PM)</p>

API version: 2.0
Contact: support-api@efficientip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdsclient

import (
	"encoding/json"
)

// checks if the IpamNetwork6AddInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamNetwork6AddInput{}

// IpamNetwork6AddInput struct for IpamNetwork6AddInput
type IpamNetwork6AddInput struct {
	// The database identifier (ID) of an existing IPv6 network you want to set as the parent of the IPv6 network you are adding. You can specify a subnet-type network to set up a network-based VLSM organization.
	ParentNetwork6Id *int32 `json:"parent_network6_id,omitempty"`
	// The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice.
	SpaceId *int32 `json:"space_id,omitempty"`
	// The name of the space.
	SpaceName *string `json:"space_name,omitempty"`
	// The start IP address of the IPv6 network, its first IP address.
	Network6Addr *string `json:"network6_addr,omitempty"`
	// The end IP address of the IPv6 network, its last IP address.
	Network6EndAddr *string `json:"network6_end_addr,omitempty"`
	// The prefix of the IPv6 network, an integer that defines the number of address the network contains.
	Network6Prefix *string `json:"network6_prefix,omitempty"`
	// Internal use. Not documented.
	AllowBlock6Creation *int32 `json:"allow_block6_creation,omitempty"`
	// A way to allow (1) or prevent (0) changing the parent of the network you are adding. Upon editing of the network, this parameter decides if you can associate it with a different parent network.
	AllowTreeReparenting *int32 `json:"allow_tree_reparenting,omitempty"`
	// A way to determine if a network can contain other networks. If set to 1, the network is terminal and cannot contain other subnet-type networks. By essence, block-type networks are non-terminal and are always set to 0.
	Network6IsTerminal *int32 `json:"network6_is_terminal,omitempty"`
	// A way to prevent <b>(1)</b> users from assigning the broadcast IP address and network IP address of the network.
	Network6LockNetworkBroadcast *int32 `json:"network6_lock_network_broadcast,omitempty"`
	// A way to authorize (<b>1</b>) IPv6 networks overlapping within a space.
	PermitInvalid *int32 `json:"permit_invalid,omitempty"`
	// A way to force the creation of an IPv6 subnet-type network. If set to 1, you can create a subnet-type network even if no block-type network matching the start address exists.
	PermitNoBlock6 *int32 `json:"permit_no_block6,omitempty"`
	// The position of the network within the hierarchy of networks of a VLSM organization. It calculates between <b>0</b> and <b>n</b> all the levels of the organization, its behavior depends on the value of the parameter <b>use_reversed_relative_position</b>:<ul><li> <b>use_reversed_relative_position=0</b> where <b>0</b> indicates a block-type network at the highest level possible, in a space-based organization, it belongs to the top space. The levels increment from <b>0</b> down to <b>n</b>, the lowest level you set up, within networks or spaces.<br/></li><li> <b>use_reversed_relative_position=1</b> where <b>1</b> indicates a network located at the lowest level of the organization, within networks or spaces. The levels increment from <b>0</b> up to <b>n</b>, the network at the highest level of the organization.<br/></li></ul>
	RelativePosition *int32 `json:"relative_position,omitempty"`
	// The object activation status.<ul><li> If set to <b>0</b>, the object is present in the database but ignored, i.e. it cannot be managed, counted or listed. This status is applied on objects deleted from the GUI.<br/></li><li> If set to <b>1</b>, the object is enabled and managed.<br/></li><li> If set to <b>2</b>, the object is unmanaged, disabled or both depending on the context.<br/></li></ul>By default, <b>row_state</b> is set to <b>1</b> when an object is created.
	RowState *string `json:"row_state,omitempty"`
	// The name of the IPv6 network, each IPv6 network must have a unique name.
	Network6Name *string `json:"network6_name,omitempty"`
	// The level of the network within the space:<ul><li> Set it to <b>0</b> for a block-type network.<br/></li><li> Set it to a value between <b>1</b> and <b>n</b> for a subnet-type network.<br/></li></ul>If you set a value between <b>2</b> and <b>n</b>, you are setting a network-based VLSM organization where non terminal subnet-type networks can contain other subnet-type networks.
	NetworkLevel *int32 `json:"network_level,omitempty"`
	// A way to determine if the calculation of the parameter <b>relative_position</b> should start from the top (<b>0</b>) or the bottom (<b>1</b>) of the VLSM organization.
	UseReversedRelativePosition *int32 `json:"use_reversed_relative_position,omitempty"`
	// The database identifier (ID) of the VLAN you want to associate with the network.
	VlanId *int32 `json:"vlan_id,omitempty"`
	// The database identifier (ID) of a VLSM child space of the space specified in <b>space_id</b>. If you specify an ID, the subnet-type network you are adding is duplicated as a VLSM block-type network in the child space, with the same name but a different ID. This parameter serves the same purpose as <b>vlsm_space_name</b>.
	VlsmSpaceId *int32 `json:"vlsm_space_id,omitempty"`
	// The name of a VLSM child space of the space specified in <b>space_id</b>. If you specify a name, the subnet-type network you are adding is duplicated as a VLSM block-type network in the child space, with the same name but a different ID. This parameter serves the same purpose as <b>vlsm_space_id</b>.
	VlsmSpaceName *string `json:"vlsm_space_name,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete []string `json:"class_parameters_to_delete,omitempty"`
	// class parameters in json format
	Network6ClassParameters []ApiClassParameterInputEntry `json:"network6_class_parameters,omitempty"`
	// The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	Network6ClassName *string `json:"network6_class_name,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewIpamNetwork6AddInput instantiates a new IpamNetwork6AddInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamNetwork6AddInput() *IpamNetwork6AddInput {
	this := IpamNetwork6AddInput{}
	return &this
}

// NewIpamNetwork6AddInputWithDefaults instantiates a new IpamNetwork6AddInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamNetwork6AddInputWithDefaults() *IpamNetwork6AddInput {
	this := IpamNetwork6AddInput{}
	return &this
}

// GetParentNetwork6Id returns the ParentNetwork6Id field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetParentNetwork6Id() int32 {
	if o == nil || IsNil(o.ParentNetwork6Id) {
		var ret int32
		return ret
	}
	return *o.ParentNetwork6Id
}

// GetParentNetwork6IdOk returns a tuple with the ParentNetwork6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetParentNetwork6IdOk() (*int32, bool) {
	if o == nil || IsNil(o.ParentNetwork6Id) {
		return nil, false
	}
	return o.ParentNetwork6Id, true
}

// HasParentNetwork6Id returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasParentNetwork6Id() bool {
	if o != nil && !IsNil(o.ParentNetwork6Id) {
		return true
	}

	return false
}

// SetParentNetwork6Id gets a reference to the given int32 and assigns it to the ParentNetwork6Id field.
func (o *IpamNetwork6AddInput) SetParentNetwork6Id(v int32) {
	o.ParentNetwork6Id = &v
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetSpaceId() int32 {
	if o == nil || IsNil(o.SpaceId) {
		var ret int32
		return ret
	}
	return *o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetSpaceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SpaceId) {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasSpaceId() bool {
	if o != nil && !IsNil(o.SpaceId) {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given int32 and assigns it to the SpaceId field.
func (o *IpamNetwork6AddInput) SetSpaceId(v int32) {
	o.SpaceId = &v
}

// GetSpaceName returns the SpaceName field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetSpaceName() string {
	if o == nil || IsNil(o.SpaceName) {
		var ret string
		return ret
	}
	return *o.SpaceName
}

// GetSpaceNameOk returns a tuple with the SpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetSpaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SpaceName) {
		return nil, false
	}
	return o.SpaceName, true
}

// HasSpaceName returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasSpaceName() bool {
	if o != nil && !IsNil(o.SpaceName) {
		return true
	}

	return false
}

// SetSpaceName gets a reference to the given string and assigns it to the SpaceName field.
func (o *IpamNetwork6AddInput) SetSpaceName(v string) {
	o.SpaceName = &v
}

// GetNetwork6Addr returns the Network6Addr field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetwork6Addr() string {
	if o == nil || IsNil(o.Network6Addr) {
		var ret string
		return ret
	}
	return *o.Network6Addr
}

// GetNetwork6AddrOk returns a tuple with the Network6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetwork6AddrOk() (*string, bool) {
	if o == nil || IsNil(o.Network6Addr) {
		return nil, false
	}
	return o.Network6Addr, true
}

// HasNetwork6Addr returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetwork6Addr() bool {
	if o != nil && !IsNil(o.Network6Addr) {
		return true
	}

	return false
}

// SetNetwork6Addr gets a reference to the given string and assigns it to the Network6Addr field.
func (o *IpamNetwork6AddInput) SetNetwork6Addr(v string) {
	o.Network6Addr = &v
}

// GetNetwork6EndAddr returns the Network6EndAddr field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetwork6EndAddr() string {
	if o == nil || IsNil(o.Network6EndAddr) {
		var ret string
		return ret
	}
	return *o.Network6EndAddr
}

// GetNetwork6EndAddrOk returns a tuple with the Network6EndAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetwork6EndAddrOk() (*string, bool) {
	if o == nil || IsNil(o.Network6EndAddr) {
		return nil, false
	}
	return o.Network6EndAddr, true
}

// HasNetwork6EndAddr returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetwork6EndAddr() bool {
	if o != nil && !IsNil(o.Network6EndAddr) {
		return true
	}

	return false
}

// SetNetwork6EndAddr gets a reference to the given string and assigns it to the Network6EndAddr field.
func (o *IpamNetwork6AddInput) SetNetwork6EndAddr(v string) {
	o.Network6EndAddr = &v
}

// GetNetwork6Prefix returns the Network6Prefix field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetwork6Prefix() string {
	if o == nil || IsNil(o.Network6Prefix) {
		var ret string
		return ret
	}
	return *o.Network6Prefix
}

// GetNetwork6PrefixOk returns a tuple with the Network6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetwork6PrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Network6Prefix) {
		return nil, false
	}
	return o.Network6Prefix, true
}

// HasNetwork6Prefix returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetwork6Prefix() bool {
	if o != nil && !IsNil(o.Network6Prefix) {
		return true
	}

	return false
}

// SetNetwork6Prefix gets a reference to the given string and assigns it to the Network6Prefix field.
func (o *IpamNetwork6AddInput) SetNetwork6Prefix(v string) {
	o.Network6Prefix = &v
}

// GetAllowBlock6Creation returns the AllowBlock6Creation field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetAllowBlock6Creation() int32 {
	if o == nil || IsNil(o.AllowBlock6Creation) {
		var ret int32
		return ret
	}
	return *o.AllowBlock6Creation
}

// GetAllowBlock6CreationOk returns a tuple with the AllowBlock6Creation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetAllowBlock6CreationOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowBlock6Creation) {
		return nil, false
	}
	return o.AllowBlock6Creation, true
}

// HasAllowBlock6Creation returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasAllowBlock6Creation() bool {
	if o != nil && !IsNil(o.AllowBlock6Creation) {
		return true
	}

	return false
}

// SetAllowBlock6Creation gets a reference to the given int32 and assigns it to the AllowBlock6Creation field.
func (o *IpamNetwork6AddInput) SetAllowBlock6Creation(v int32) {
	o.AllowBlock6Creation = &v
}

// GetAllowTreeReparenting returns the AllowTreeReparenting field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetAllowTreeReparenting() int32 {
	if o == nil || IsNil(o.AllowTreeReparenting) {
		var ret int32
		return ret
	}
	return *o.AllowTreeReparenting
}

// GetAllowTreeReparentingOk returns a tuple with the AllowTreeReparenting field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetAllowTreeReparentingOk() (*int32, bool) {
	if o == nil || IsNil(o.AllowTreeReparenting) {
		return nil, false
	}
	return o.AllowTreeReparenting, true
}

// HasAllowTreeReparenting returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasAllowTreeReparenting() bool {
	if o != nil && !IsNil(o.AllowTreeReparenting) {
		return true
	}

	return false
}

// SetAllowTreeReparenting gets a reference to the given int32 and assigns it to the AllowTreeReparenting field.
func (o *IpamNetwork6AddInput) SetAllowTreeReparenting(v int32) {
	o.AllowTreeReparenting = &v
}

// GetNetwork6IsTerminal returns the Network6IsTerminal field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetwork6IsTerminal() int32 {
	if o == nil || IsNil(o.Network6IsTerminal) {
		var ret int32
		return ret
	}
	return *o.Network6IsTerminal
}

// GetNetwork6IsTerminalOk returns a tuple with the Network6IsTerminal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetwork6IsTerminalOk() (*int32, bool) {
	if o == nil || IsNil(o.Network6IsTerminal) {
		return nil, false
	}
	return o.Network6IsTerminal, true
}

// HasNetwork6IsTerminal returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetwork6IsTerminal() bool {
	if o != nil && !IsNil(o.Network6IsTerminal) {
		return true
	}

	return false
}

// SetNetwork6IsTerminal gets a reference to the given int32 and assigns it to the Network6IsTerminal field.
func (o *IpamNetwork6AddInput) SetNetwork6IsTerminal(v int32) {
	o.Network6IsTerminal = &v
}

// GetNetwork6LockNetworkBroadcast returns the Network6LockNetworkBroadcast field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetwork6LockNetworkBroadcast() int32 {
	if o == nil || IsNil(o.Network6LockNetworkBroadcast) {
		var ret int32
		return ret
	}
	return *o.Network6LockNetworkBroadcast
}

// GetNetwork6LockNetworkBroadcastOk returns a tuple with the Network6LockNetworkBroadcast field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetwork6LockNetworkBroadcastOk() (*int32, bool) {
	if o == nil || IsNil(o.Network6LockNetworkBroadcast) {
		return nil, false
	}
	return o.Network6LockNetworkBroadcast, true
}

// HasNetwork6LockNetworkBroadcast returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetwork6LockNetworkBroadcast() bool {
	if o != nil && !IsNil(o.Network6LockNetworkBroadcast) {
		return true
	}

	return false
}

// SetNetwork6LockNetworkBroadcast gets a reference to the given int32 and assigns it to the Network6LockNetworkBroadcast field.
func (o *IpamNetwork6AddInput) SetNetwork6LockNetworkBroadcast(v int32) {
	o.Network6LockNetworkBroadcast = &v
}

// GetPermitInvalid returns the PermitInvalid field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetPermitInvalid() int32 {
	if o == nil || IsNil(o.PermitInvalid) {
		var ret int32
		return ret
	}
	return *o.PermitInvalid
}

// GetPermitInvalidOk returns a tuple with the PermitInvalid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetPermitInvalidOk() (*int32, bool) {
	if o == nil || IsNil(o.PermitInvalid) {
		return nil, false
	}
	return o.PermitInvalid, true
}

// HasPermitInvalid returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasPermitInvalid() bool {
	if o != nil && !IsNil(o.PermitInvalid) {
		return true
	}

	return false
}

// SetPermitInvalid gets a reference to the given int32 and assigns it to the PermitInvalid field.
func (o *IpamNetwork6AddInput) SetPermitInvalid(v int32) {
	o.PermitInvalid = &v
}

// GetPermitNoBlock6 returns the PermitNoBlock6 field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetPermitNoBlock6() int32 {
	if o == nil || IsNil(o.PermitNoBlock6) {
		var ret int32
		return ret
	}
	return *o.PermitNoBlock6
}

// GetPermitNoBlock6Ok returns a tuple with the PermitNoBlock6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetPermitNoBlock6Ok() (*int32, bool) {
	if o == nil || IsNil(o.PermitNoBlock6) {
		return nil, false
	}
	return o.PermitNoBlock6, true
}

// HasPermitNoBlock6 returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasPermitNoBlock6() bool {
	if o != nil && !IsNil(o.PermitNoBlock6) {
		return true
	}

	return false
}

// SetPermitNoBlock6 gets a reference to the given int32 and assigns it to the PermitNoBlock6 field.
func (o *IpamNetwork6AddInput) SetPermitNoBlock6(v int32) {
	o.PermitNoBlock6 = &v
}

// GetRelativePosition returns the RelativePosition field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetRelativePosition() int32 {
	if o == nil || IsNil(o.RelativePosition) {
		var ret int32
		return ret
	}
	return *o.RelativePosition
}

// GetRelativePositionOk returns a tuple with the RelativePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetRelativePositionOk() (*int32, bool) {
	if o == nil || IsNil(o.RelativePosition) {
		return nil, false
	}
	return o.RelativePosition, true
}

// HasRelativePosition returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasRelativePosition() bool {
	if o != nil && !IsNil(o.RelativePosition) {
		return true
	}

	return false
}

// SetRelativePosition gets a reference to the given int32 and assigns it to the RelativePosition field.
func (o *IpamNetwork6AddInput) SetRelativePosition(v int32) {
	o.RelativePosition = &v
}

// GetRowState returns the RowState field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetRowState() string {
	if o == nil || IsNil(o.RowState) {
		var ret string
		return ret
	}
	return *o.RowState
}

// GetRowStateOk returns a tuple with the RowState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetRowStateOk() (*string, bool) {
	if o == nil || IsNil(o.RowState) {
		return nil, false
	}
	return o.RowState, true
}

// HasRowState returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasRowState() bool {
	if o != nil && !IsNil(o.RowState) {
		return true
	}

	return false
}

// SetRowState gets a reference to the given string and assigns it to the RowState field.
func (o *IpamNetwork6AddInput) SetRowState(v string) {
	o.RowState = &v
}

// GetNetwork6Name returns the Network6Name field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetwork6Name() string {
	if o == nil || IsNil(o.Network6Name) {
		var ret string
		return ret
	}
	return *o.Network6Name
}

// GetNetwork6NameOk returns a tuple with the Network6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetwork6NameOk() (*string, bool) {
	if o == nil || IsNil(o.Network6Name) {
		return nil, false
	}
	return o.Network6Name, true
}

// HasNetwork6Name returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetwork6Name() bool {
	if o != nil && !IsNil(o.Network6Name) {
		return true
	}

	return false
}

// SetNetwork6Name gets a reference to the given string and assigns it to the Network6Name field.
func (o *IpamNetwork6AddInput) SetNetwork6Name(v string) {
	o.Network6Name = &v
}

// GetNetworkLevel returns the NetworkLevel field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetworkLevel() int32 {
	if o == nil || IsNil(o.NetworkLevel) {
		var ret int32
		return ret
	}
	return *o.NetworkLevel
}

// GetNetworkLevelOk returns a tuple with the NetworkLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetworkLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.NetworkLevel) {
		return nil, false
	}
	return o.NetworkLevel, true
}

// HasNetworkLevel returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetworkLevel() bool {
	if o != nil && !IsNil(o.NetworkLevel) {
		return true
	}

	return false
}

// SetNetworkLevel gets a reference to the given int32 and assigns it to the NetworkLevel field.
func (o *IpamNetwork6AddInput) SetNetworkLevel(v int32) {
	o.NetworkLevel = &v
}

// GetUseReversedRelativePosition returns the UseReversedRelativePosition field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetUseReversedRelativePosition() int32 {
	if o == nil || IsNil(o.UseReversedRelativePosition) {
		var ret int32
		return ret
	}
	return *o.UseReversedRelativePosition
}

// GetUseReversedRelativePositionOk returns a tuple with the UseReversedRelativePosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetUseReversedRelativePositionOk() (*int32, bool) {
	if o == nil || IsNil(o.UseReversedRelativePosition) {
		return nil, false
	}
	return o.UseReversedRelativePosition, true
}

// HasUseReversedRelativePosition returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasUseReversedRelativePosition() bool {
	if o != nil && !IsNil(o.UseReversedRelativePosition) {
		return true
	}

	return false
}

// SetUseReversedRelativePosition gets a reference to the given int32 and assigns it to the UseReversedRelativePosition field.
func (o *IpamNetwork6AddInput) SetUseReversedRelativePosition(v int32) {
	o.UseReversedRelativePosition = &v
}

// GetVlanId returns the VlanId field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetVlanId() int32 {
	if o == nil || IsNil(o.VlanId) {
		var ret int32
		return ret
	}
	return *o.VlanId
}

// GetVlanIdOk returns a tuple with the VlanId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetVlanIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlanId) {
		return nil, false
	}
	return o.VlanId, true
}

// HasVlanId returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasVlanId() bool {
	if o != nil && !IsNil(o.VlanId) {
		return true
	}

	return false
}

// SetVlanId gets a reference to the given int32 and assigns it to the VlanId field.
func (o *IpamNetwork6AddInput) SetVlanId(v int32) {
	o.VlanId = &v
}

// GetVlsmSpaceId returns the VlsmSpaceId field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetVlsmSpaceId() int32 {
	if o == nil || IsNil(o.VlsmSpaceId) {
		var ret int32
		return ret
	}
	return *o.VlsmSpaceId
}

// GetVlsmSpaceIdOk returns a tuple with the VlsmSpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetVlsmSpaceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.VlsmSpaceId) {
		return nil, false
	}
	return o.VlsmSpaceId, true
}

// HasVlsmSpaceId returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasVlsmSpaceId() bool {
	if o != nil && !IsNil(o.VlsmSpaceId) {
		return true
	}

	return false
}

// SetVlsmSpaceId gets a reference to the given int32 and assigns it to the VlsmSpaceId field.
func (o *IpamNetwork6AddInput) SetVlsmSpaceId(v int32) {
	o.VlsmSpaceId = &v
}

// GetVlsmSpaceName returns the VlsmSpaceName field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetVlsmSpaceName() string {
	if o == nil || IsNil(o.VlsmSpaceName) {
		var ret string
		return ret
	}
	return *o.VlsmSpaceName
}

// GetVlsmSpaceNameOk returns a tuple with the VlsmSpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetVlsmSpaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.VlsmSpaceName) {
		return nil, false
	}
	return o.VlsmSpaceName, true
}

// HasVlsmSpaceName returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasVlsmSpaceName() bool {
	if o != nil && !IsNil(o.VlsmSpaceName) {
		return true
	}

	return false
}

// SetVlsmSpaceName gets a reference to the given string and assigns it to the VlsmSpaceName field.
func (o *IpamNetwork6AddInput) SetVlsmSpaceName(v string) {
	o.VlsmSpaceName = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetClassParametersToDelete() []string {
	if o == nil || IsNil(o.ClassParametersToDelete) {
		var ret []string
		return ret
	}
	return o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetClassParametersToDeleteOk() ([]string, bool) {
	if o == nil || IsNil(o.ClassParametersToDelete) {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasClassParametersToDelete() bool {
	if o != nil && !IsNil(o.ClassParametersToDelete) {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *IpamNetwork6AddInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = v
}

// GetNetwork6ClassParameters returns the Network6ClassParameters field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetwork6ClassParameters() []ApiClassParameterInputEntry {
	if o == nil || IsNil(o.Network6ClassParameters) {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return o.Network6ClassParameters
}

// GetNetwork6ClassParametersOk returns a tuple with the Network6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetwork6ClassParametersOk() ([]ApiClassParameterInputEntry, bool) {
	if o == nil || IsNil(o.Network6ClassParameters) {
		return nil, false
	}
	return o.Network6ClassParameters, true
}

// HasNetwork6ClassParameters returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetwork6ClassParameters() bool {
	if o != nil && !IsNil(o.Network6ClassParameters) {
		return true
	}

	return false
}

// SetNetwork6ClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the Network6ClassParameters field.
func (o *IpamNetwork6AddInput) SetNetwork6ClassParameters(v []ApiClassParameterInputEntry) {
	o.Network6ClassParameters = v
}

// GetNetwork6ClassName returns the Network6ClassName field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetNetwork6ClassName() string {
	if o == nil || IsNil(o.Network6ClassName) {
		var ret string
		return ret
	}
	return *o.Network6ClassName
}

// GetNetwork6ClassNameOk returns a tuple with the Network6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetNetwork6ClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.Network6ClassName) {
		return nil, false
	}
	return o.Network6ClassName, true
}

// HasNetwork6ClassName returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasNetwork6ClassName() bool {
	if o != nil && !IsNil(o.Network6ClassName) {
		return true
	}

	return false
}

// SetNetwork6ClassName gets a reference to the given string and assigns it to the Network6ClassName field.
func (o *IpamNetwork6AddInput) SetNetwork6ClassName(v string) {
	o.Network6ClassName = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *IpamNetwork6AddInput) GetWarnings() string {
	if o == nil || IsNil(o.Warnings) {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamNetwork6AddInput) GetWarningsOk() (*string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *IpamNetwork6AddInput) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *IpamNetwork6AddInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o IpamNetwork6AddInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamNetwork6AddInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ParentNetwork6Id) {
		toSerialize["parent_network6_id"] = o.ParentNetwork6Id
	}
	if !IsNil(o.SpaceId) {
		toSerialize["space_id"] = o.SpaceId
	}
	if !IsNil(o.SpaceName) {
		toSerialize["space_name"] = o.SpaceName
	}
	if !IsNil(o.Network6Addr) {
		toSerialize["network6_addr"] = o.Network6Addr
	}
	if !IsNil(o.Network6EndAddr) {
		toSerialize["network6_end_addr"] = o.Network6EndAddr
	}
	if !IsNil(o.Network6Prefix) {
		toSerialize["network6_prefix"] = o.Network6Prefix
	}
	if !IsNil(o.AllowBlock6Creation) {
		toSerialize["allow_block6_creation"] = o.AllowBlock6Creation
	}
	if !IsNil(o.AllowTreeReparenting) {
		toSerialize["allow_tree_reparenting"] = o.AllowTreeReparenting
	}
	if !IsNil(o.Network6IsTerminal) {
		toSerialize["network6_is_terminal"] = o.Network6IsTerminal
	}
	if !IsNil(o.Network6LockNetworkBroadcast) {
		toSerialize["network6_lock_network_broadcast"] = o.Network6LockNetworkBroadcast
	}
	if !IsNil(o.PermitInvalid) {
		toSerialize["permit_invalid"] = o.PermitInvalid
	}
	if !IsNil(o.PermitNoBlock6) {
		toSerialize["permit_no_block6"] = o.PermitNoBlock6
	}
	if !IsNil(o.RelativePosition) {
		toSerialize["relative_position"] = o.RelativePosition
	}
	if !IsNil(o.RowState) {
		toSerialize["row_state"] = o.RowState
	}
	if !IsNil(o.Network6Name) {
		toSerialize["network6_name"] = o.Network6Name
	}
	if !IsNil(o.NetworkLevel) {
		toSerialize["network_level"] = o.NetworkLevel
	}
	if !IsNil(o.UseReversedRelativePosition) {
		toSerialize["use_reversed_relative_position"] = o.UseReversedRelativePosition
	}
	if !IsNil(o.VlanId) {
		toSerialize["vlan_id"] = o.VlanId
	}
	if !IsNil(o.VlsmSpaceId) {
		toSerialize["vlsm_space_id"] = o.VlsmSpaceId
	}
	if !IsNil(o.VlsmSpaceName) {
		toSerialize["vlsm_space_name"] = o.VlsmSpaceName
	}
	if !IsNil(o.ClassParametersToDelete) {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if !IsNil(o.Network6ClassParameters) {
		toSerialize["network6_class_parameters"] = o.Network6ClassParameters
	}
	if !IsNil(o.Network6ClassName) {
		toSerialize["network6_class_name"] = o.Network6ClassName
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableIpamNetwork6AddInput struct {
	value *IpamNetwork6AddInput
	isSet bool
}

func (v NullableIpamNetwork6AddInput) Get() *IpamNetwork6AddInput {
	return v.value
}

func (v *NullableIpamNetwork6AddInput) Set(val *IpamNetwork6AddInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamNetwork6AddInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamNetwork6AddInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamNetwork6AddInput(val *IpamNetwork6AddInput) *NullableIpamNetwork6AddInput {
	return &NullableIpamNetwork6AddInput{value: val, isSet: true}
}

func (v NullableIpamNetwork6AddInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamNetwork6AddInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
