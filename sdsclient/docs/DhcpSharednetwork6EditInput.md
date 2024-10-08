# DhcpSharednetwork6EditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 server, a unique numeric key value automatically incremented when you add a DHCPv6 server. Use the ID to specify the DHCPv6 server of your choice. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server. | [optional] 
**Sharednetwork6Id** | Pointer to **int32** | The database identifier (ID) of the DHCPv6 shared network, a unique numeric key value automatically incremented when you add a DHCPv6 shared network. Use the ID to specify the DHCPv6 shared network of your choice. | [optional] 
**Sharednetwork6Name** | Pointer to **string** | The name of the DHCPv6 shared network. | [optional] 
**Server6Hostaddr** | Pointer to **string** | The IP address of the DHCP server. | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDhcpSharednetwork6EditInput

`func NewDhcpSharednetwork6EditInput() *DhcpSharednetwork6EditInput`

NewDhcpSharednetwork6EditInput instantiates a new DhcpSharednetwork6EditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpSharednetwork6EditInputWithDefaults

`func NewDhcpSharednetwork6EditInputWithDefaults() *DhcpSharednetwork6EditInput`

NewDhcpSharednetwork6EditInputWithDefaults instantiates a new DhcpSharednetwork6EditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer6Id

`func (o *DhcpSharednetwork6EditInput) GetServer6Id() int32`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DhcpSharednetwork6EditInput) GetServer6IdOk() (*int32, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DhcpSharednetwork6EditInput) SetServer6Id(v int32)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DhcpSharednetwork6EditInput) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6Name

`func (o *DhcpSharednetwork6EditInput) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DhcpSharednetwork6EditInput) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DhcpSharednetwork6EditInput) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DhcpSharednetwork6EditInput) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetSharednetwork6Id

`func (o *DhcpSharednetwork6EditInput) GetSharednetwork6Id() int32`

GetSharednetwork6Id returns the Sharednetwork6Id field if non-nil, zero value otherwise.

### GetSharednetwork6IdOk

`func (o *DhcpSharednetwork6EditInput) GetSharednetwork6IdOk() (*int32, bool)`

GetSharednetwork6IdOk returns a tuple with the Sharednetwork6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Id

`func (o *DhcpSharednetwork6EditInput) SetSharednetwork6Id(v int32)`

SetSharednetwork6Id sets Sharednetwork6Id field to given value.

### HasSharednetwork6Id

`func (o *DhcpSharednetwork6EditInput) HasSharednetwork6Id() bool`

HasSharednetwork6Id returns a boolean if a field has been set.

### GetSharednetwork6Name

`func (o *DhcpSharednetwork6EditInput) GetSharednetwork6Name() string`

GetSharednetwork6Name returns the Sharednetwork6Name field if non-nil, zero value otherwise.

### GetSharednetwork6NameOk

`func (o *DhcpSharednetwork6EditInput) GetSharednetwork6NameOk() (*string, bool)`

GetSharednetwork6NameOk returns a tuple with the Sharednetwork6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharednetwork6Name

`func (o *DhcpSharednetwork6EditInput) SetSharednetwork6Name(v string)`

SetSharednetwork6Name sets Sharednetwork6Name field to given value.

### HasSharednetwork6Name

`func (o *DhcpSharednetwork6EditInput) HasSharednetwork6Name() bool`

HasSharednetwork6Name returns a boolean if a field has been set.

### GetServer6Hostaddr

`func (o *DhcpSharednetwork6EditInput) GetServer6Hostaddr() string`

GetServer6Hostaddr returns the Server6Hostaddr field if non-nil, zero value otherwise.

### GetServer6HostaddrOk

`func (o *DhcpSharednetwork6EditInput) GetServer6HostaddrOk() (*string, bool)`

GetServer6HostaddrOk returns a tuple with the Server6Hostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Hostaddr

`func (o *DhcpSharednetwork6EditInput) SetServer6Hostaddr(v string)`

SetServer6Hostaddr sets Server6Hostaddr field to given value.

### HasServer6Hostaddr

`func (o *DhcpSharednetwork6EditInput) HasServer6Hostaddr() bool`

HasServer6Hostaddr returns a boolean if a field has been set.

### GetWarnings

`func (o *DhcpSharednetwork6EditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DhcpSharednetwork6EditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DhcpSharednetwork6EditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DhcpSharednetwork6EditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


