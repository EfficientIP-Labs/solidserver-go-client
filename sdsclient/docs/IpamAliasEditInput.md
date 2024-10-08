# IpamAliasEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressHostaddr** | Pointer to **string** | The IP address. | [optional] 
**AddressId** | Pointer to **int32** | The database identifier (ID) of the IPv4 address, a unique numeric key value automatically incremented when you add an IPv4 address. Use the ID to specify the IPv4 address of your choice. | [optional] 
**AliasName** | Pointer to **string** | The name of the IPv4 alias. | [optional] 
**AliasId** | Pointer to **int32** | The database identifier (ID) of the IPv4 alias, a unique numeric key value automatically incremented when you add an IPv4 alias. Use the ID to specify the IPv4 alias of your choice. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**AliasType** | Pointer to **string** | The type of the alias. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamAliasEditInput

`func NewIpamAliasEditInput() *IpamAliasEditInput`

NewIpamAliasEditInput instantiates a new IpamAliasEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAliasEditInputWithDefaults

`func NewIpamAliasEditInputWithDefaults() *IpamAliasEditInput`

NewIpamAliasEditInputWithDefaults instantiates a new IpamAliasEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressHostaddr

`func (o *IpamAliasEditInput) GetAddressHostaddr() string`

GetAddressHostaddr returns the AddressHostaddr field if non-nil, zero value otherwise.

### GetAddressHostaddrOk

`func (o *IpamAliasEditInput) GetAddressHostaddrOk() (*string, bool)`

GetAddressHostaddrOk returns a tuple with the AddressHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressHostaddr

`func (o *IpamAliasEditInput) SetAddressHostaddr(v string)`

SetAddressHostaddr sets AddressHostaddr field to given value.

### HasAddressHostaddr

`func (o *IpamAliasEditInput) HasAddressHostaddr() bool`

HasAddressHostaddr returns a boolean if a field has been set.

### GetAddressId

`func (o *IpamAliasEditInput) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *IpamAliasEditInput) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *IpamAliasEditInput) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *IpamAliasEditInput) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetAliasName

`func (o *IpamAliasEditInput) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *IpamAliasEditInput) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *IpamAliasEditInput) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *IpamAliasEditInput) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetAliasId

`func (o *IpamAliasEditInput) GetAliasId() int32`

GetAliasId returns the AliasId field if non-nil, zero value otherwise.

### GetAliasIdOk

`func (o *IpamAliasEditInput) GetAliasIdOk() (*int32, bool)`

GetAliasIdOk returns a tuple with the AliasId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasId

`func (o *IpamAliasEditInput) SetAliasId(v int32)`

SetAliasId sets AliasId field to given value.

### HasAliasId

`func (o *IpamAliasEditInput) HasAliasId() bool`

HasAliasId returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAliasEditInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAliasEditInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAliasEditInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAliasEditInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAliasEditInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAliasEditInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAliasEditInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAliasEditInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetAliasType

`func (o *IpamAliasEditInput) GetAliasType() string`

GetAliasType returns the AliasType field if non-nil, zero value otherwise.

### GetAliasTypeOk

`func (o *IpamAliasEditInput) GetAliasTypeOk() (*string, bool)`

GetAliasTypeOk returns a tuple with the AliasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasType

`func (o *IpamAliasEditInput) SetAliasType(v string)`

SetAliasType sets AliasType field to given value.

### HasAliasType

`func (o *IpamAliasEditInput) HasAliasType() bool`

HasAliasType returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamAliasEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamAliasEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamAliasEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamAliasEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


