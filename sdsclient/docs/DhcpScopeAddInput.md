# DhcpScopeAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. Use the ID to specify the DHCPv4 server of your choice. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server. | [optional] 
**ScopeNetAddr** | Pointer to **string** | The first IP address of the DHCPv4 scope. | [optional] 
**ScopeNetMask** | Pointer to **string** | The netmask of the DHCPv4 scope. It is expressed in dot-decimal notation and defines the number of addresses the scope contains. | [optional] 
**ServerHostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**FailoverId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 failover channel, a unique numeric key value automatically incremented when you add a DHCPv4 failover channel. Use the ID to specify the DHCPv4 failover channel of your choice. | [optional] 
**FailoverName** | Pointer to **string** | The name of the DHCPv4 failover channel. | [optional] 
**ScopeName** | Pointer to **string** | The name of the DHCPv4 scope, each DHCPv4 scope must have a unique name. | [optional] 
**ScopeSpaceId** | Pointer to **int32** | The database identifier (ID) of an existing space you want to associate with the DHCPv4 scope. | [optional] 
**ScopeSpaceName** | Pointer to **string** | The name of an existing space you want to associate with the DHCPv4 scope. | [optional] 
**SharednetworkId** | Pointer to **int32** | The database identifier (ID) of the DHCPv4 shared network, a unique numeric key value automatically incremented when you add a DHCPv4 shared network. Use the ID to specify the DHCPv4 shared network of your choice. | [optional] 
**SharednetworkName** | Pointer to **string** | The name of the DHCPv4 shared network. | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**ScopeClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**ScopeClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpScopeAddInput

`func NewDhcpScopeAddInput() *DhcpScopeAddInput`

NewDhcpScopeAddInput instantiates a new DhcpScopeAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpScopeAddInputWithDefaults

`func NewDhcpScopeAddInputWithDefaults() *DhcpScopeAddInput`

NewDhcpScopeAddInputWithDefaults instantiates a new DhcpScopeAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerId

`func (o *DhcpScopeAddInput) GetServerId() int32`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpScopeAddInput) GetServerIdOk() (*int32, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpScopeAddInput) SetServerId(v int32)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpScopeAddInput) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpScopeAddInput) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpScopeAddInput) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpScopeAddInput) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpScopeAddInput) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetScopeNetAddr

`func (o *DhcpScopeAddInput) GetScopeNetAddr() string`

GetScopeNetAddr returns the ScopeNetAddr field if non-nil, zero value otherwise.

### GetScopeNetAddrOk

`func (o *DhcpScopeAddInput) GetScopeNetAddrOk() (*string, bool)`

GetScopeNetAddrOk returns a tuple with the ScopeNetAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetAddr

`func (o *DhcpScopeAddInput) SetScopeNetAddr(v string)`

SetScopeNetAddr sets ScopeNetAddr field to given value.

### HasScopeNetAddr

`func (o *DhcpScopeAddInput) HasScopeNetAddr() bool`

HasScopeNetAddr returns a boolean if a field has been set.

### GetScopeNetMask

`func (o *DhcpScopeAddInput) GetScopeNetMask() string`

GetScopeNetMask returns the ScopeNetMask field if non-nil, zero value otherwise.

### GetScopeNetMaskOk

`func (o *DhcpScopeAddInput) GetScopeNetMaskOk() (*string, bool)`

GetScopeNetMaskOk returns a tuple with the ScopeNetMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeNetMask

`func (o *DhcpScopeAddInput) SetScopeNetMask(v string)`

SetScopeNetMask sets ScopeNetMask field to given value.

### HasScopeNetMask

`func (o *DhcpScopeAddInput) HasScopeNetMask() bool`

HasScopeNetMask returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DhcpScopeAddInput) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DhcpScopeAddInput) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DhcpScopeAddInput) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DhcpScopeAddInput) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetFailoverId

`func (o *DhcpScopeAddInput) GetFailoverId() int32`

GetFailoverId returns the FailoverId field if non-nil, zero value otherwise.

### GetFailoverIdOk

`func (o *DhcpScopeAddInput) GetFailoverIdOk() (*int32, bool)`

GetFailoverIdOk returns a tuple with the FailoverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverId

`func (o *DhcpScopeAddInput) SetFailoverId(v int32)`

SetFailoverId sets FailoverId field to given value.

### HasFailoverId

`func (o *DhcpScopeAddInput) HasFailoverId() bool`

HasFailoverId returns a boolean if a field has been set.

### GetFailoverName

`func (o *DhcpScopeAddInput) GetFailoverName() string`

GetFailoverName returns the FailoverName field if non-nil, zero value otherwise.

### GetFailoverNameOk

`func (o *DhcpScopeAddInput) GetFailoverNameOk() (*string, bool)`

GetFailoverNameOk returns a tuple with the FailoverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailoverName

`func (o *DhcpScopeAddInput) SetFailoverName(v string)`

SetFailoverName sets FailoverName field to given value.

### HasFailoverName

`func (o *DhcpScopeAddInput) HasFailoverName() bool`

HasFailoverName returns a boolean if a field has been set.

### GetScopeName

`func (o *DhcpScopeAddInput) GetScopeName() string`

GetScopeName returns the ScopeName field if non-nil, zero value otherwise.

### GetScopeNameOk

`func (o *DhcpScopeAddInput) GetScopeNameOk() (*string, bool)`

GetScopeNameOk returns a tuple with the ScopeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeName

`func (o *DhcpScopeAddInput) SetScopeName(v string)`

SetScopeName sets ScopeName field to given value.

### HasScopeName

`func (o *DhcpScopeAddInput) HasScopeName() bool`

HasScopeName returns a boolean if a field has been set.

### GetScopeSpaceId

`func (o *DhcpScopeAddInput) GetScopeSpaceId() int32`

GetScopeSpaceId returns the ScopeSpaceId field if non-nil, zero value otherwise.

### GetScopeSpaceIdOk

`func (o *DhcpScopeAddInput) GetScopeSpaceIdOk() (*int32, bool)`

GetScopeSpaceIdOk returns a tuple with the ScopeSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceId

`func (o *DhcpScopeAddInput) SetScopeSpaceId(v int32)`

SetScopeSpaceId sets ScopeSpaceId field to given value.

### HasScopeSpaceId

`func (o *DhcpScopeAddInput) HasScopeSpaceId() bool`

HasScopeSpaceId returns a boolean if a field has been set.

### GetScopeSpaceName

`func (o *DhcpScopeAddInput) GetScopeSpaceName() string`

GetScopeSpaceName returns the ScopeSpaceName field if non-nil, zero value otherwise.

### GetScopeSpaceNameOk

`func (o *DhcpScopeAddInput) GetScopeSpaceNameOk() (*string, bool)`

GetScopeSpaceNameOk returns a tuple with the ScopeSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeSpaceName

`func (o *DhcpScopeAddInput) SetScopeSpaceName(v string)`

SetScopeSpaceName sets ScopeSpaceName field to given value.

### HasScopeSpaceName

`func (o *DhcpScopeAddInput) HasScopeSpaceName() bool`

HasScopeSpaceName returns a boolean if a field has been set.

### GetSharednetworkId

`func (o *DhcpScopeAddInput) GetSharednetworkId() int32`

GetSharednetworkId returns the SharednetworkId field if non-nil, zero value otherwise.

### GetSharednetworkIdOk

`func (o *DhcpScopeAddInput) GetSharednetworkIdOk() (*int32, bool)`

GetSharednetworkIdOk returns a tuple with the SharednetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkId

`func (o *DhcpScopeAddInput) SetSharednetworkId(v int32)`

SetSharednetworkId sets SharednetworkId field to given value.

### HasSharednetworkId

`func (o *DhcpScopeAddInput) HasSharednetworkId() bool`

HasSharednetworkId returns a boolean if a field has been set.

### GetSharednetworkName

`func (o *DhcpScopeAddInput) GetSharednetworkName() string`

GetSharednetworkName returns the SharednetworkName field if non-nil, zero value otherwise.

### GetSharednetworkNameOk

`func (o *DhcpScopeAddInput) GetSharednetworkNameOk() (*string, bool)`

GetSharednetworkNameOk returns a tuple with the SharednetworkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetworkName

`func (o *DhcpScopeAddInput) SetSharednetworkName(v string)`

SetSharednetworkName sets SharednetworkName field to given value.

### HasSharednetworkName

`func (o *DhcpScopeAddInput) HasSharednetworkName() bool`

HasSharednetworkName returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *DhcpScopeAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *DhcpScopeAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *DhcpScopeAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *DhcpScopeAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetScopeClassName

`func (o *DhcpScopeAddInput) GetScopeClassName() string`

GetScopeClassName returns the ScopeClassName field if non-nil, zero value otherwise.

### GetScopeClassNameOk

`func (o *DhcpScopeAddInput) GetScopeClassNameOk() (*string, bool)`

GetScopeClassNameOk returns a tuple with the ScopeClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassName

`func (o *DhcpScopeAddInput) SetScopeClassName(v string)`

SetScopeClassName sets ScopeClassName field to given value.

### HasScopeClassName

`func (o *DhcpScopeAddInput) HasScopeClassName() bool`

HasScopeClassName returns a boolean if a field has been set.

### GetScopeClassParameters

`func (o *DhcpScopeAddInput) GetScopeClassParameters() []ApiClassParameterInputEntry`

GetScopeClassParameters returns the ScopeClassParameters field if non-nil, zero value otherwise.

### GetScopeClassParametersOk

`func (o *DhcpScopeAddInput) GetScopeClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetScopeClassParametersOk returns a tuple with the ScopeClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopeClassParameters

`func (o *DhcpScopeAddInput) SetScopeClassParameters(v []ApiClassParameterInputEntry)`

SetScopeClassParameters sets ScopeClassParameters field to given value.

### HasScopeClassParameters

`func (o *DhcpScopeAddInput) HasScopeClassParameters() bool`

HasScopeClassParameters returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpScopeAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpScopeAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpScopeAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpScopeAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


