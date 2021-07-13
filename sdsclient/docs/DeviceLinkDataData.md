# DeviceLinkDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LinkAddTime** | Pointer to **string** | The time at which the Device Manager port or interface has been added, in decimal UNIX date format. | [optional] 
**LinkAutoLink** | Pointer to **string** | A way to determine if the link between two Device Manager devices is set automatically (&lt;b&gt;1&lt;/b&gt;) or manually (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**DeviceId** | Pointer to **string** | The database identifier (ID) of the Device Manager device the object belongs to. | [optional] 
**DeviceName** | Pointer to **string** | The name of the Device Manager device. | [optional] 
**InterfaceId** | Pointer to **string** | The database identifier (ID) of the Device Manager port or interface. | [optional] 
**InterfaceMac** | Pointer to **string** | The MAC address associated with the Device Manager port or interface. | [optional] 
**InterfaceName** | Pointer to **string** | The name of the Device Manager port or interface. | [optional] 
**InterfaceType** | Pointer to **string** | A way to indicate if the object is either a &lt;b&gt;port&lt;/b&gt; or an &lt;b&gt;interface&lt;/b&gt;. | [optional] 
**LinkId** | Pointer to **string** | The database identifier (ID) of the Device Manager port or interface link. Use the ID to specify the port or interface link of your choice. | [optional] 

## Methods

### NewDeviceLinkDataData

`func NewDeviceLinkDataData() *DeviceLinkDataData`

NewDeviceLinkDataData instantiates a new DeviceLinkDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceLinkDataDataWithDefaults

`func NewDeviceLinkDataDataWithDefaults() *DeviceLinkDataData`

NewDeviceLinkDataDataWithDefaults instantiates a new DeviceLinkDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLinkAddTime

`func (o *DeviceLinkDataData) GetLinkAddTime() string`

GetLinkAddTime returns the LinkAddTime field if non-nil, zero value otherwise.

### GetLinkAddTimeOk

`func (o *DeviceLinkDataData) GetLinkAddTimeOk() (*string, bool)`

GetLinkAddTimeOk returns a tuple with the LinkAddTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAddTime

`func (o *DeviceLinkDataData) SetLinkAddTime(v string)`

SetLinkAddTime sets LinkAddTime field to given value.

### HasLinkAddTime

`func (o *DeviceLinkDataData) HasLinkAddTime() bool`

HasLinkAddTime returns a boolean if a field has been set.

### GetLinkAutoLink

`func (o *DeviceLinkDataData) GetLinkAutoLink() string`

GetLinkAutoLink returns the LinkAutoLink field if non-nil, zero value otherwise.

### GetLinkAutoLinkOk

`func (o *DeviceLinkDataData) GetLinkAutoLinkOk() (*string, bool)`

GetLinkAutoLinkOk returns a tuple with the LinkAutoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAutoLink

`func (o *DeviceLinkDataData) SetLinkAutoLink(v string)`

SetLinkAutoLink sets LinkAutoLink field to given value.

### HasLinkAutoLink

`func (o *DeviceLinkDataData) HasLinkAutoLink() bool`

HasLinkAutoLink returns a boolean if a field has been set.

### GetDeviceId

`func (o *DeviceLinkDataData) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceLinkDataData) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceLinkDataData) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceLinkDataData) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceName

`func (o *DeviceLinkDataData) GetDeviceName() string`

GetDeviceName returns the DeviceName field if non-nil, zero value otherwise.

### GetDeviceNameOk

`func (o *DeviceLinkDataData) GetDeviceNameOk() (*string, bool)`

GetDeviceNameOk returns a tuple with the DeviceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceName

`func (o *DeviceLinkDataData) SetDeviceName(v string)`

SetDeviceName sets DeviceName field to given value.

### HasDeviceName

`func (o *DeviceLinkDataData) HasDeviceName() bool`

HasDeviceName returns a boolean if a field has been set.

### GetInterfaceId

`func (o *DeviceLinkDataData) GetInterfaceId() string`

GetInterfaceId returns the InterfaceId field if non-nil, zero value otherwise.

### GetInterfaceIdOk

`func (o *DeviceLinkDataData) GetInterfaceIdOk() (*string, bool)`

GetInterfaceIdOk returns a tuple with the InterfaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceId

`func (o *DeviceLinkDataData) SetInterfaceId(v string)`

SetInterfaceId sets InterfaceId field to given value.

### HasInterfaceId

`func (o *DeviceLinkDataData) HasInterfaceId() bool`

HasInterfaceId returns a boolean if a field has been set.

### GetInterfaceMac

`func (o *DeviceLinkDataData) GetInterfaceMac() string`

GetInterfaceMac returns the InterfaceMac field if non-nil, zero value otherwise.

### GetInterfaceMacOk

`func (o *DeviceLinkDataData) GetInterfaceMacOk() (*string, bool)`

GetInterfaceMacOk returns a tuple with the InterfaceMac field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceMac

`func (o *DeviceLinkDataData) SetInterfaceMac(v string)`

SetInterfaceMac sets InterfaceMac field to given value.

### HasInterfaceMac

`func (o *DeviceLinkDataData) HasInterfaceMac() bool`

HasInterfaceMac returns a boolean if a field has been set.

### GetInterfaceName

`func (o *DeviceLinkDataData) GetInterfaceName() string`

GetInterfaceName returns the InterfaceName field if non-nil, zero value otherwise.

### GetInterfaceNameOk

`func (o *DeviceLinkDataData) GetInterfaceNameOk() (*string, bool)`

GetInterfaceNameOk returns a tuple with the InterfaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceName

`func (o *DeviceLinkDataData) SetInterfaceName(v string)`

SetInterfaceName sets InterfaceName field to given value.

### HasInterfaceName

`func (o *DeviceLinkDataData) HasInterfaceName() bool`

HasInterfaceName returns a boolean if a field has been set.

### GetInterfaceType

`func (o *DeviceLinkDataData) GetInterfaceType() string`

GetInterfaceType returns the InterfaceType field if non-nil, zero value otherwise.

### GetInterfaceTypeOk

`func (o *DeviceLinkDataData) GetInterfaceTypeOk() (*string, bool)`

GetInterfaceTypeOk returns a tuple with the InterfaceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaceType

`func (o *DeviceLinkDataData) SetInterfaceType(v string)`

SetInterfaceType sets InterfaceType field to given value.

### HasInterfaceType

`func (o *DeviceLinkDataData) HasInterfaceType() bool`

HasInterfaceType returns a boolean if a field has been set.

### GetLinkId

`func (o *DeviceLinkDataData) GetLinkId() string`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *DeviceLinkDataData) GetLinkIdOk() (*string, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *DeviceLinkDataData) SetLinkId(v string)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *DeviceLinkDataData) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


