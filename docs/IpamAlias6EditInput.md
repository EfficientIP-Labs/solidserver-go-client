# IpamAlias6EditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address6Hostaddr** | Pointer to **string** | todo(here) :ipam.alias6.edit.input.address6_hostaddr. : IPv6 address | [optional] 
**Address6Id** | Pointer to **int32** | The database identifier (ID) of the IPv6 address, a unique numeric key value automatically incremented when you add an IPv6 address. Use the ID to specify the IPv6 address of your choice. | [optional] 
**Alias6Name** | Pointer to **string** | todo(here) :ipam.alias6.edit.input.alias6_name. : String | [optional] 
**Alias6Id** | Pointer to **int32** | The database identifier (ID) of the IPv6 alias, a unique numeric key value automatically incremented when you add an IPv6 alias. Use the ID to specify which IPv6 alias to edit. | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**Alias6Type** | Pointer to **string** | The type of the alias. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamAlias6EditInput

`func NewIpamAlias6EditInput() *IpamAlias6EditInput`

NewIpamAlias6EditInput instantiates a new IpamAlias6EditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAlias6EditInputWithDefaults

`func NewIpamAlias6EditInputWithDefaults() *IpamAlias6EditInput`

NewIpamAlias6EditInputWithDefaults instantiates a new IpamAlias6EditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress6Hostaddr

`func (o *IpamAlias6EditInput) GetAddress6Hostaddr() string`

GetAddress6Hostaddr returns the Address6Hostaddr field if non-nil, zero value otherwise.

### GetAddress6HostaddrOk

`func (o *IpamAlias6EditInput) GetAddress6HostaddrOk() (*string, bool)`

GetAddress6HostaddrOk returns a tuple with the Address6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Hostaddr

`func (o *IpamAlias6EditInput) SetAddress6Hostaddr(v string)`

SetAddress6Hostaddr sets Address6Hostaddr field to given value.

### HasAddress6Hostaddr

`func (o *IpamAlias6EditInput) HasAddress6Hostaddr() bool`

HasAddress6Hostaddr returns a boolean if a field has been set.

### GetAddress6Id

`func (o *IpamAlias6EditInput) GetAddress6Id() int32`

GetAddress6Id returns the Address6Id field if non-nil, zero value otherwise.

### GetAddress6IdOk

`func (o *IpamAlias6EditInput) GetAddress6IdOk() (*int32, bool)`

GetAddress6IdOk returns a tuple with the Address6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Id

`func (o *IpamAlias6EditInput) SetAddress6Id(v int32)`

SetAddress6Id sets Address6Id field to given value.

### HasAddress6Id

`func (o *IpamAlias6EditInput) HasAddress6Id() bool`

HasAddress6Id returns a boolean if a field has been set.

### GetAlias6Name

`func (o *IpamAlias6EditInput) GetAlias6Name() string`

GetAlias6Name returns the Alias6Name field if non-nil, zero value otherwise.

### GetAlias6NameOk

`func (o *IpamAlias6EditInput) GetAlias6NameOk() (*string, bool)`

GetAlias6NameOk returns a tuple with the Alias6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias6Name

`func (o *IpamAlias6EditInput) SetAlias6Name(v string)`

SetAlias6Name sets Alias6Name field to given value.

### HasAlias6Name

`func (o *IpamAlias6EditInput) HasAlias6Name() bool`

HasAlias6Name returns a boolean if a field has been set.

### GetAlias6Id

`func (o *IpamAlias6EditInput) GetAlias6Id() int32`

GetAlias6Id returns the Alias6Id field if non-nil, zero value otherwise.

### GetAlias6IdOk

`func (o *IpamAlias6EditInput) GetAlias6IdOk() (*int32, bool)`

GetAlias6IdOk returns a tuple with the Alias6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias6Id

`func (o *IpamAlias6EditInput) SetAlias6Id(v int32)`

SetAlias6Id sets Alias6Id field to given value.

### HasAlias6Id

`func (o *IpamAlias6EditInput) HasAlias6Id() bool`

HasAlias6Id returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAlias6EditInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAlias6EditInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAlias6EditInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAlias6EditInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAlias6EditInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAlias6EditInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAlias6EditInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAlias6EditInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetAlias6Type

`func (o *IpamAlias6EditInput) GetAlias6Type() string`

GetAlias6Type returns the Alias6Type field if non-nil, zero value otherwise.

### GetAlias6TypeOk

`func (o *IpamAlias6EditInput) GetAlias6TypeOk() (*string, bool)`

GetAlias6TypeOk returns a tuple with the Alias6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias6Type

`func (o *IpamAlias6EditInput) SetAlias6Type(v string)`

SetAlias6Type sets Alias6Type field to given value.

### HasAlias6Type

`func (o *IpamAlias6EditInput) HasAlias6Type() bool`

HasAlias6Type returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamAlias6EditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamAlias6EditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamAlias6EditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamAlias6EditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


