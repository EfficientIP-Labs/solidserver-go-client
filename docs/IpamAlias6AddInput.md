# IpamAlias6AddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address6Hostaddr** | Pointer to **string** | todo(here) :ipam.alias6.add.input.address6_hostaddr. : IPv6 address | [optional] 
**Address6Id** | Pointer to **int32** | The database identifier (ID) of the IPv6 address, a unique numeric key value automatically incremented when you add an IPv6 address. Use the ID to specify the IPv6 address of your choice. | [optional] 
**Alias6Name** | Pointer to **string** | todo(here) :ipam.alias6.add.input.alias6_name. : String | [optional] 
**SpaceId** | Pointer to **int32** | The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 
**Alias6Type** | Pointer to **string** | The type of the alias. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewIpamAlias6AddInput

`func NewIpamAlias6AddInput() *IpamAlias6AddInput`

NewIpamAlias6AddInput instantiates a new IpamAlias6AddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIpamAlias6AddInputWithDefaults

`func NewIpamAlias6AddInputWithDefaults() *IpamAlias6AddInput`

NewIpamAlias6AddInputWithDefaults instantiates a new IpamAlias6AddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress6Hostaddr

`func (o *IpamAlias6AddInput) GetAddress6Hostaddr() string`

GetAddress6Hostaddr returns the Address6Hostaddr field if non-nil, zero value otherwise.

### GetAddress6HostaddrOk

`func (o *IpamAlias6AddInput) GetAddress6HostaddrOk() (*string, bool)`

GetAddress6HostaddrOk returns a tuple with the Address6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Hostaddr

`func (o *IpamAlias6AddInput) SetAddress6Hostaddr(v string)`

SetAddress6Hostaddr sets Address6Hostaddr field to given value.

### HasAddress6Hostaddr

`func (o *IpamAlias6AddInput) HasAddress6Hostaddr() bool`

HasAddress6Hostaddr returns a boolean if a field has been set.

### GetAddress6Id

`func (o *IpamAlias6AddInput) GetAddress6Id() int32`

GetAddress6Id returns the Address6Id field if non-nil, zero value otherwise.

### GetAddress6IdOk

`func (o *IpamAlias6AddInput) GetAddress6IdOk() (*int32, bool)`

GetAddress6IdOk returns a tuple with the Address6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress6Id

`func (o *IpamAlias6AddInput) SetAddress6Id(v int32)`

SetAddress6Id sets Address6Id field to given value.

### HasAddress6Id

`func (o *IpamAlias6AddInput) HasAddress6Id() bool`

HasAddress6Id returns a boolean if a field has been set.

### GetAlias6Name

`func (o *IpamAlias6AddInput) GetAlias6Name() string`

GetAlias6Name returns the Alias6Name field if non-nil, zero value otherwise.

### GetAlias6NameOk

`func (o *IpamAlias6AddInput) GetAlias6NameOk() (*string, bool)`

GetAlias6NameOk returns a tuple with the Alias6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias6Name

`func (o *IpamAlias6AddInput) SetAlias6Name(v string)`

SetAlias6Name sets Alias6Name field to given value.

### HasAlias6Name

`func (o *IpamAlias6AddInput) HasAlias6Name() bool`

HasAlias6Name returns a boolean if a field has been set.

### GetSpaceId

`func (o *IpamAlias6AddInput) GetSpaceId() int32`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *IpamAlias6AddInput) GetSpaceIdOk() (*int32, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *IpamAlias6AddInput) SetSpaceId(v int32)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *IpamAlias6AddInput) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceName

`func (o *IpamAlias6AddInput) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *IpamAlias6AddInput) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *IpamAlias6AddInput) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *IpamAlias6AddInput) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.

### GetAlias6Type

`func (o *IpamAlias6AddInput) GetAlias6Type() string`

GetAlias6Type returns the Alias6Type field if non-nil, zero value otherwise.

### GetAlias6TypeOk

`func (o *IpamAlias6AddInput) GetAlias6TypeOk() (*string, bool)`

GetAlias6TypeOk returns a tuple with the Alias6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias6Type

`func (o *IpamAlias6AddInput) SetAlias6Type(v string)`

SetAlias6Type sets Alias6Type field to given value.

### HasAlias6Type

`func (o *IpamAlias6AddInput) HasAlias6Type() bool`

HasAlias6Type returns a boolean if a field has been set.

### GetWarnings

`func (o *IpamAlias6AddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *IpamAlias6AddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *IpamAlias6AddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *IpamAlias6AddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


