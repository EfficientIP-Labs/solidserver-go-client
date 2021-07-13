# DnsViewparamAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ViewId** | Pointer to **int32** | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. | [optional] 
**ViewparamKey** | Pointer to **string** | The name of the DNS option you want to add, edit or remove from the view. You can only set one option at a time. &lt;ul class&#x3D;dashed &gt;&lt;li&gt;                                                To add or edit an option: specify its name in the parameter &lt;b&gt;param_key&lt;/b&gt;, as follows &lt;b&gt;param_key&#x3D;&amp;lt;option-name&amp;gt;&lt;/b&gt;, and then specify its value in the parameter &lt;b&gt;param_value&lt;/b&gt;.&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                To remove an option, specify its name in the parameter &lt;b&gt;param_key&lt;/b&gt; and leave the parameter &lt;b&gt;param_value&lt;/b&gt; empty.&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt;To set several options, specify as many parameters (&lt;b&gt;param_key&lt;/b&gt; and &lt;b&gt;param_value&lt;/b&gt;) as you need. | [optional] 
**ViewparamIsArray** | Pointer to **int32** | A way to determine is the DNS view option is an array (&lt;b&gt;1&lt;/b&gt;). | [optional] 
**ViewparamValue** | Pointer to **string** | The value of the DNS option specified in the input &lt;b&gt;param_key&lt;/b&gt;.&lt;ul class&#x3D;dashed &gt;&lt;li&gt;                                                To add or edit an option value, specify its name in the parameter &lt;b&gt;param_key&lt;/b&gt; and set its value as follows: &lt;b&gt;param_value&#x3D;&amp;lt;option-value&amp;gt;&lt;/b&gt; .&lt;br/&gt;                                            &lt;/li&gt;&lt;li&gt;                                                To remove an option value, specify its name in the parameter &lt;b&gt;param_key&lt;/b&gt; and leave &lt;b&gt;param_value&lt;/b&gt; empty: &lt;b&gt;param_value&#x3D;&lt;/b&gt; .&lt;br/&gt;                                            &lt;/li&gt;&lt;/ul&gt; | [optional] 

## Methods

### NewDnsViewparamAddInput

`func NewDnsViewparamAddInput() *DnsViewparamAddInput`

NewDnsViewparamAddInput instantiates a new DnsViewparamAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsViewparamAddInputWithDefaults

`func NewDnsViewparamAddInputWithDefaults() *DnsViewparamAddInput`

NewDnsViewparamAddInputWithDefaults instantiates a new DnsViewparamAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetViewId

`func (o *DnsViewparamAddInput) GetViewId() int32`

GetViewId returns the ViewId field if non-nil, zero value otherwise.

### GetViewIdOk

`func (o *DnsViewparamAddInput) GetViewIdOk() (*int32, bool)`

GetViewIdOk returns a tuple with the ViewId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewId

`func (o *DnsViewparamAddInput) SetViewId(v int32)`

SetViewId sets ViewId field to given value.

### HasViewId

`func (o *DnsViewparamAddInput) HasViewId() bool`

HasViewId returns a boolean if a field has been set.

### GetViewparamKey

`func (o *DnsViewparamAddInput) GetViewparamKey() string`

GetViewparamKey returns the ViewparamKey field if non-nil, zero value otherwise.

### GetViewparamKeyOk

`func (o *DnsViewparamAddInput) GetViewparamKeyOk() (*string, bool)`

GetViewparamKeyOk returns a tuple with the ViewparamKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewparamKey

`func (o *DnsViewparamAddInput) SetViewparamKey(v string)`

SetViewparamKey sets ViewparamKey field to given value.

### HasViewparamKey

`func (o *DnsViewparamAddInput) HasViewparamKey() bool`

HasViewparamKey returns a boolean if a field has been set.

### GetViewparamIsArray

`func (o *DnsViewparamAddInput) GetViewparamIsArray() int32`

GetViewparamIsArray returns the ViewparamIsArray field if non-nil, zero value otherwise.

### GetViewparamIsArrayOk

`func (o *DnsViewparamAddInput) GetViewparamIsArrayOk() (*int32, bool)`

GetViewparamIsArrayOk returns a tuple with the ViewparamIsArray field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewparamIsArray

`func (o *DnsViewparamAddInput) SetViewparamIsArray(v int32)`

SetViewparamIsArray sets ViewparamIsArray field to given value.

### HasViewparamIsArray

`func (o *DnsViewparamAddInput) HasViewparamIsArray() bool`

HasViewparamIsArray returns a boolean if a field has been set.

### GetViewparamValue

`func (o *DnsViewparamAddInput) GetViewparamValue() string`

GetViewparamValue returns the ViewparamValue field if non-nil, zero value otherwise.

### GetViewparamValueOk

`func (o *DnsViewparamAddInput) GetViewparamValueOk() (*string, bool)`

GetViewparamValueOk returns a tuple with the ViewparamValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewparamValue

`func (o *DnsViewparamAddInput) SetViewparamValue(v string)`

SetViewparamValue sets ViewparamValue field to given value.

### HasViewparamValue

`func (o *DnsViewparamAddInput) HasViewparamValue() bool`

HasViewparamValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


