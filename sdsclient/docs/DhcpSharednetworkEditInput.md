# DhcpSharednetworkEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. Use the ID to specify the DHCPv4 server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server. | [optional] 
**SharednetworkId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 shared network, a unique numeric key value automatically incremented when you add a DHCPv4 shared network. Use the ID to specify the DHCPv4 shared network of your choice. | [optional] 
**SharednetworkName** | Pointer to **string** | The name of the DHCPv4 shared network, each DHCPv4 shared network must have a unique name. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpSharednetworkEditInput

`func NewDhcpSharednetworkEditInput() *DhcpSharednetworkEditInput`

NewDhcpSharednetworkEditInput instantiates a new DhcpSharednetworkEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSharednetworkEditInputWithDefaults

`func NewDhcpSharednetworkEditInputWithDefaults() *DhcpSharednetworkEditInput`

NewDhcpSharednetworkEditInputWithDefaults instantiates a new DhcpSharednetworkEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DhcpSharednetworkEditInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpSharednetworkEditInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpSharednetworkEditInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpSharednetworkEditInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpSharednetworkEditInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpSharednetworkEditInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpSharednetworkEditInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpSharednetworkEditInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DhcpSharednetworkEditInput) GetSharednetworkId() int32`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DhcpSharednetworkEditInput) GetSharednetworkIdOk() (*int32, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DhcpSharednetworkEditInput) SetSharednetworkId(v int32)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DhcpSharednetworkEditInput) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DhcpSharednetworkEditInput) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DhcpSharednetworkEditInput) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DhcpSharednetworkEditInput) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DhcpSharednetworkEditInput) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DhcpSharednetworkEditInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DhcpSharednetworkEditInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DhcpSharednetworkEditInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DhcpSharednetworkEditInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpSharednetworkEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpSharednetworkEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpSharednetworkEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpSharednetworkEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


