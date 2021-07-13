# DeviceLinkAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Device1Name** | Pointer to **string** | The name of the device to which belongs the DM port or interface you want to link with &lt;b&gt;hostiface2_id&lt;/b&gt;. | [optional] 
**Device2Name** | Pointer to **string** | The name of the device to which belongs the DM port or interface you want to link with &lt;b&gt;hostiface1_id&lt;/b&gt;. | [optional] 
**Interface1Id** | Pointer to **int32** | The database identifier (ID) of the DM port or interface you want to link with &lt;b&gt;hostiface2_id&lt;/b&gt;. | [optional] 
**Interface1Name** | Pointer to **string** | The name of the DM port or interface you want to link with &lt;b&gt;hostiface2_id&lt;/b&gt;. | [optional] 
**Interface2Id** | Pointer to **int32** | The database identifier (ID) of the DM port or interface you want to link with &lt;b&gt;hostiface1_id&lt;/b&gt;. | [optional] 
**Interface2Name** | Pointer to **string** | The name of the DM port or interface you want to link with &lt;b&gt;hostiface1_id&lt;/b&gt;. | [optional] 
**LinkId** | Pointer to **int32** | The database identifier (ID) of the Device Manager port or interface link, a unique numeric key value automatically incremented when you add a link between a device and a port or interface. Use the ID to specify which port or interface link to edit. | [optional] 
**LinkAutoLink** | Pointer to **int32** | A way to determine if the link between two Device Manager devices is set automatically (&lt;b&gt;1&lt;/b&gt;) or manually (&lt;b&gt;0&lt;/b&gt;). | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewDeviceLinkAddInput

`func NewDeviceLinkAddInput() *DeviceLinkAddInput`

NewDeviceLinkAddInput instantiates a new DeviceLinkAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceLinkAddInputWithDefaults

`func NewDeviceLinkAddInputWithDefaults() *DeviceLinkAddInput`

NewDeviceLinkAddInputWithDefaults instantiates a new DeviceLinkAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDevice1Name

`func (o *DeviceLinkAddInput) GetDevice1Name() string`

GetDevice1Name returns the Device1Name field if non-nil, zero value otherwise.

### GetDevice1NameOk

`func (o *DeviceLinkAddInput) GetDevice1NameOk() (*string, bool)`

GetDevice1NameOk returns a tuple with the Device1Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice1Name

`func (o *DeviceLinkAddInput) SetDevice1Name(v string)`

SetDevice1Name sets Device1Name field to given value.

### HasDevice1Name

`func (o *DeviceLinkAddInput) HasDevice1Name() bool`

HasDevice1Name returns a boolean if a field has been set.

### GetDevice2Name

`func (o *DeviceLinkAddInput) GetDevice2Name() string`

GetDevice2Name returns the Device2Name field if non-nil, zero value otherwise.

### GetDevice2NameOk

`func (o *DeviceLinkAddInput) GetDevice2NameOk() (*string, bool)`

GetDevice2NameOk returns a tuple with the Device2Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice2Name

`func (o *DeviceLinkAddInput) SetDevice2Name(v string)`

SetDevice2Name sets Device2Name field to given value.

### HasDevice2Name

`func (o *DeviceLinkAddInput) HasDevice2Name() bool`

HasDevice2Name returns a boolean if a field has been set.

### GetInterface1Id

`func (o *DeviceLinkAddInput) GetInterface1Id() int32`

GetInterface1Id returns the Interface1Id field if non-nil, zero value otherwise.

### GetInterface1IdOk

`func (o *DeviceLinkAddInput) GetInterface1IdOk() (*int32, bool)`

GetInterface1IdOk returns a tuple with the Interface1Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface1Id

`func (o *DeviceLinkAddInput) SetInterface1Id(v int32)`

SetInterface1Id sets Interface1Id field to given value.

### HasInterface1Id

`func (o *DeviceLinkAddInput) HasInterface1Id() bool`

HasInterface1Id returns a boolean if a field has been set.

### GetInterface1Name

`func (o *DeviceLinkAddInput) GetInterface1Name() string`

GetInterface1Name returns the Interface1Name field if non-nil, zero value otherwise.

### GetInterface1NameOk

`func (o *DeviceLinkAddInput) GetInterface1NameOk() (*string, bool)`

GetInterface1NameOk returns a tuple with the Interface1Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface1Name

`func (o *DeviceLinkAddInput) SetInterface1Name(v string)`

SetInterface1Name sets Interface1Name field to given value.

### HasInterface1Name

`func (o *DeviceLinkAddInput) HasInterface1Name() bool`

HasInterface1Name returns a boolean if a field has been set.

### GetInterface2Id

`func (o *DeviceLinkAddInput) GetInterface2Id() int32`

GetInterface2Id returns the Interface2Id field if non-nil, zero value otherwise.

### GetInterface2IdOk

`func (o *DeviceLinkAddInput) GetInterface2IdOk() (*int32, bool)`

GetInterface2IdOk returns a tuple with the Interface2Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface2Id

`func (o *DeviceLinkAddInput) SetInterface2Id(v int32)`

SetInterface2Id sets Interface2Id field to given value.

### HasInterface2Id

`func (o *DeviceLinkAddInput) HasInterface2Id() bool`

HasInterface2Id returns a boolean if a field has been set.

### GetInterface2Name

`func (o *DeviceLinkAddInput) GetInterface2Name() string`

GetInterface2Name returns the Interface2Name field if non-nil, zero value otherwise.

### GetInterface2NameOk

`func (o *DeviceLinkAddInput) GetInterface2NameOk() (*string, bool)`

GetInterface2NameOk returns a tuple with the Interface2Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface2Name

`func (o *DeviceLinkAddInput) SetInterface2Name(v string)`

SetInterface2Name sets Interface2Name field to given value.

### HasInterface2Name

`func (o *DeviceLinkAddInput) HasInterface2Name() bool`

HasInterface2Name returns a boolean if a field has been set.

### GetLinkId

`func (o *DeviceLinkAddInput) GetLinkId() int32`

GetLinkId returns the LinkId field if non-nil, zero value otherwise.

### GetLinkIdOk

`func (o *DeviceLinkAddInput) GetLinkIdOk() (*int32, bool)`

GetLinkIdOk returns a tuple with the LinkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkId

`func (o *DeviceLinkAddInput) SetLinkId(v int32)`

SetLinkId sets LinkId field to given value.

### HasLinkId

`func (o *DeviceLinkAddInput) HasLinkId() bool`

HasLinkId returns a boolean if a field has been set.

### GetLinkAutoLink

`func (o *DeviceLinkAddInput) GetLinkAutoLink() int32`

GetLinkAutoLink returns the LinkAutoLink field if non-nil, zero value otherwise.

### GetLinkAutoLinkOk

`func (o *DeviceLinkAddInput) GetLinkAutoLinkOk() (*int32, bool)`

GetLinkAutoLinkOk returns a tuple with the LinkAutoLink field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinkAutoLink

`func (o *DeviceLinkAddInput) SetLinkAutoLink(v int32)`

SetLinkAutoLink sets LinkAutoLink field to given value.

### HasLinkAutoLink

`func (o *DeviceLinkAddInput) HasLinkAutoLink() bool`

HasLinkAutoLink returns a boolean if a field has been set.

### GetWarnings

`func (o *DeviceLinkAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *DeviceLinkAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *DeviceLinkAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *DeviceLinkAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


