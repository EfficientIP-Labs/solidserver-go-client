# IpamAliasAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressHostaddr** | Pointer to **string** | todo(here) :ipam.alias.add.input.address_hostaddr. : IPv4 address | [optional] 
**AddressId** | Pointer to **int32** | The database identifier (ID) of the IPv4 address, a unique numeric key value automatically incremented when you add an IPv4 address. Use the ID to specify the IPv4 address of your choice. | [optional] 
**AliasName** | Pointer to **string** | todo(here) :ipam.alias.add.input.alias_name. : String | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**AliasType** | Pointer to **string** | The type of the alias. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamAliasAddInput

`func NewIpamAliasAddInput() *IpamAliasAddInput`

NewIpamAliasAddInput instantiates a new IpamAliasAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAliasAddInputWithDefaults

`func NewIpamAliasAddInputWithDefaults() *IpamAliasAddInput`

NewIpamAliasAddInputWithDefaults instantiates a new IpamAliasAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddressHostaddr

`func (o *IpamAliasAddInput) GetAddressHostaddr() string`

GetAddressHostaddr returns the AddressHostaddr field if non-nil, zero value otherwise.

### GetAddressHostaddrOk

`func (o *IpamAliasAddInput) GetAddressHostaddrOk() (*string, bool)`

GetAddressHostaddrOk returns a tuple with the AddressHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressHostaddr

`func (o *IpamAliasAddInput) SetAddressHostaddr(v string)`

SetAddressHostaddr sets AddressHostaddr field to given value.

### HasAddressHostaddr

`func (o *IpamAliasAddInput) HasAddressHostaddr() bool`

HasAddressHostaddr returns a boolean if a field has been set.

### GetAddressId

`func (o *IpamAliasAddInput) GetAddressId() int32`

GetAddressId returns the AddressId field if non-nil, zero value otherwise.

### GetAddressIdOk

`func (o *IpamAliasAddInput) GetAddressIdOk() (*int32, bool)`

GetAddressIdOk returns a tuple with the AddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressId

`func (o *IpamAliasAddInput) SetAddressId(v int32)`

SetAddressId sets AddressId field to given value.

### HasAddressId

`func (o *IpamAliasAddInput) HasAddressId() bool`

HasAddressId returns a boolean if a field has been set.

### GetAliasName

`func (o *IpamAliasAddInput) GetAliasName() string`

GetAliasName returns the AliasName field if non-nil, zero value otherwise.

### GetAliasNameOk

`func (o *IpamAliasAddInput) GetAliasNameOk() (*string, bool)`

GetAliasNameOk returns a tuple with the AliasName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasName

`func (o *IpamAliasAddInput) SetAliasName(v string)`

SetAliasName sets AliasName field to given value.

### HasAliasName

`func (o *IpamAliasAddInput) HasAliasName() bool`

HasAliasName returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAliasAddInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAliasAddInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAliasAddInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAliasAddInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAliasAddInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAliasAddInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAliasAddInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAliasAddInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetAliasType

`func (o *IpamAliasAddInput) GetAliasType() string`

GetAliasType returns the AliasType field if non-nil, zero value otherwise.

### GetAliasTypeOk

`func (o *IpamAliasAddInput) GetAliasTypeOk() (*string, bool)`

GetAliasTypeOk returns a tuple with the AliasType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAliasType

`func (o *IpamAliasAddInput) SetAliasType(v string)`

SetAliasType sets AliasType field to given value.

### HasAliasType

`func (o *IpamAliasAddInput) HasAliasType() bool`

HasAliasType returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamAliasAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamAliasAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamAliasAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamAliasAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


