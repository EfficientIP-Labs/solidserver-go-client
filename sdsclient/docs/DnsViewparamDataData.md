# DnsViewparamDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewparamValue** | Pointer to **string** | The value of the DNS option set on the view. | [optional] 
**ViewparamDelayedCreateTime** | Pointer to **string** | The delay of creation status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created yet. | [optional] 
**ViewparamDelayedDeleteTime** | Pointer to **string** | The delay of deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not deleted yet. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DNS server the object belongs to. | [optional] 
**ServerName** | Pointer to **string** | The name of the DNS server the object belongs to. | [optional] 
**ViewId** | Pointer to **string** | The database identifier (ID) of the DNS view. | [optional] 
**ViewName** | Pointer to **string** | The name of the DNS view. | [optional] 
**Oid** | Pointer to **string** | The database identifier (ID) of the DNS view param the object belongs to. | [optional] 
**ViewparamKey** | Pointer to **string** | The name of the DNS option set on the view. | [optional] 

## Methods

### NewDnsViewparamDataData

`func NewDnsViewparamDataData() *DnsViewparamDataData`

NewDnsViewparamDataData instantiates a new DnsViewparamDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsViewparamDataDataWithDefaults

`func NewDnsViewparamDataDataWithDefaults() *DnsViewparamDataData`

NewDnsViewparamDataDataWithDefaults instantiates a new DnsViewparamDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewparamValue

`func (o *DnsViewparamDataData) GetViewparamValue() string`

GetViewparamValue returns the ViewparamValue field if non-nil, zero value otherwise.

### GetViewparamValueOk

`func (o *DnsViewparamDataData) GetViewparamValueOk() (*string, bool)`

GetViewparamValueOk returns a tuple with the ViewparamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewparamValue

`func (o *DnsViewparamDataData) SetViewparamValue(v string)`

SetViewparamValue sets ViewparamValue field to given value.

### HasViewparamValue

`func (o *DnsViewparamDataData) HasViewparamValue() bool`

HasViewparamValue returns a boolean if a field has been set.

### GetViewparamDelayedCreateTime

`func (o *DnsViewparamDataData) GetViewparamDelayedCreateTime() string`

GetViewparamDelayedCreateTime returns the ViewparamDelayedCreateTime field if non-nil, zero value otherwise.

### GetViewparamDelayedCreateTimeOk

`func (o *DnsViewparamDataData) GetViewparamDelayedCreateTimeOk() (*string, bool)`

GetViewparamDelayedCreateTimeOk returns a tuple with the ViewparamDelayedCreateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewparamDelayedCreateTime

`func (o *DnsViewparamDataData) SetViewparamDelayedCreateTime(v string)`

SetViewparamDelayedCreateTime sets ViewparamDelayedCreateTime field to given value.

### HasViewparamDelayedCreateTime

`func (o *DnsViewparamDataData) HasViewparamDelayedCreateTime() bool`

HasViewparamDelayedCreateTime returns a boolean if a field has been set.

### GetViewparamDelayedDeleteTime

`func (o *DnsViewparamDataData) GetViewparamDelayedDeleteTime() string`

GetViewparamDelayedDeleteTime returns the ViewparamDelayedDeleteTime field if non-nil, zero value otherwise.

### GetViewparamDelayedDeleteTimeOk

`func (o *DnsViewparamDataData) GetViewparamDelayedDeleteTimeOk() (*string, bool)`

GetViewparamDelayedDeleteTimeOk returns a tuple with the ViewparamDelayedDeleteTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewparamDelayedDeleteTime

`func (o *DnsViewparamDataData) SetViewparamDelayedDeleteTime(v string)`

SetViewparamDelayedDeleteTime sets ViewparamDelayedDeleteTime field to given value.

### HasViewparamDelayedDeleteTime

`func (o *DnsViewparamDataData) HasViewparamDelayedDeleteTime() bool`

HasViewparamDelayedDeleteTime returns a boolean if a field has been set.

### GetServerId

`func (o *DnsViewparamDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DnsViewparamDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DnsViewparamDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DnsViewparamDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerName

`func (o *DnsViewparamDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DnsViewparamDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DnsViewparamDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DnsViewparamDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetViewId

`func (o *DnsViewparamDataData) GetViewId() string`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DnsViewparamDataData) GetViewIdOk() (*string, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DnsViewparamDataData) SetViewId(v string)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DnsViewparamDataData) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewName

`func (o *DnsViewparamDataData) GetViewName() string`

GetViewName returns the ViewName field if non-nil, zero value otherwise.

### GetViewNameOk

`func (o *DnsViewparamDataData) GetViewNameOk() (*string, bool)`

GetViewNameOk returns a tuple with the ViewName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewName

`func (o *DnsViewparamDataData) SetViewName(v string)`

SetViewName sets ViewName field to given value.

### HasViewName

`func (o *DnsViewparamDataData) HasViewName() bool`

HasViewName returns a boolean if a field has been set.

### GetOid

`func (o *DnsViewparamDataData) GetOid() string`

GetOid returns the Oid field if non-nil, zero value otherwise.

### GetOidOk

`func (o *DnsViewparamDataData) GetOidOk() (*string, bool)`

GetOidOk returns a tuple with the Oid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOid

`func (o *DnsViewparamDataData) SetOid(v string)`

SetOid sets Oid field to given value.

### HasOid

`func (o *DnsViewparamDataData) HasOid() bool`

HasOid returns a boolean if a field has been set.

### GetViewparamKey

`func (o *DnsViewparamDataData) GetViewparamKey() string`

GetViewparamKey returns the ViewparamKey field if non-nil, zero value otherwise.

### GetViewparamKeyOk

`func (o *DnsViewparamDataData) GetViewparamKeyOk() (*string, bool)`

GetViewparamKeyOk returns a tuple with the ViewparamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewparamKey

`func (o *DnsViewparamDataData) SetViewparamKey(v string)`

SetViewparamKey sets ViewparamKey field to given value.

### HasViewparamKey

`func (o *DnsViewparamDataData) HasViewparamKey() bool`

HasViewparamKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


