# ApiMessageEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Extras** | Pointer to **map[string]string** |  | [optional] 

## Methods

### NewApiMessageEntry

`func NewApiMessageEntry() *ApiMessageEntry`

NewApiMessageEntry instantiates a new ApiMessageEntry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiMessageEntryWithDefaults

`func NewApiMessageEntryWithDefaults() *ApiMessageEntry`

NewApiMessageEntryWithDefaults instantiates a new ApiMessageEntry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ApiMessageEntry) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ApiMessageEntry) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ApiMessageEntry) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *ApiMessageEntry) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMsg

`func (o *ApiMessageEntry) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ApiMessageEntry) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ApiMessageEntry) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ApiMessageEntry) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetType

`func (o *ApiMessageEntry) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApiMessageEntry) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApiMessageEntry) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApiMessageEntry) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExtras

`func (o *ApiMessageEntry) GetExtras() map[string]string`

GetExtras returns the Extras field if non-nil, zero value otherwise.

### GetExtrasOk

`func (o *ApiMessageEntry) GetExtrasOk() (*map[string]string, bool)`

GetExtrasOk returns a tuple with the Extras field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtras

`func (o *ApiMessageEntry) SetExtras(v map[string]string)`

SetExtras sets Extras field to given value.

### HasExtras

`func (o *ApiMessageEntry) HasExtras() bool`

HasExtras returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


