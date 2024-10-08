# DataInnerIpamSpaceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpaceMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentSpaceClassName** | Pointer to **string** | The name of the class applied to the VLSM parent space, it can be preceded by the class directory. | [optional] 
**ParentSpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the VLSM parent space and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;...&lt;/b&gt; . | [optional] 
**ParentSpaceId** | Pointer to **string** | The database identifier (ID) of the VLSM parent space. &lt;b&gt;0&lt;/b&gt; indicates that space has no parent space. | [optional] 
**ParentSpaceName** | Pointer to **string** | The name of the VLSM parent space. &lt;b&gt;#&lt;/b&gt; indicates that space has no parent space. | [optional] 
**SpaceClassName** | Pointer to **string** | The name of the class applied to the space, it can be preceded by the class directory. | [optional] 
**SpaceClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the space. | [optional] 
**SpaceDescription** | Pointer to **string** | The description of the space. | [optional] 
**SpaceId** | Pointer to **string** | The database identifier (ID) of the space. | [optional] 
**SpaceIsTemplate** | Pointer to **string** | The template status of the space. If the space is used as template (&lt;b&gt;1&lt;/b&gt;), all the IPv4 networks, pools and IP addresses it contains are also used as template. | [optional] 
**SpaceName** | Pointer to **string** | The name of the space. | [optional] 

## Methods

### NewDataInnerIpamSpaceData

`func NewDataInnerIpamSpaceData() *DataInnerIpamSpaceData`

NewDataInnerIpamSpaceData instantiates a new DataInnerIpamSpaceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerIpamSpaceDataWithDefaults

`func NewDataInnerIpamSpaceDataWithDefaults() *DataInnerIpamSpaceData`

NewDataInnerIpamSpaceDataWithDefaults instantiates a new DataInnerIpamSpaceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpaceMultistatus

`func (o *DataInnerIpamSpaceData) GetSpaceMultistatus() string`

GetSpaceMultistatus returns the SpaceMultistatus field if non-nil, zero value otherwise.

### GetSpaceMultistatusOk

`func (o *DataInnerIpamSpaceData) GetSpaceMultistatusOk() (*string, bool)`

GetSpaceMultistatusOk returns a tuple with the SpaceMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceMultistatus

`func (o *DataInnerIpamSpaceData) SetSpaceMultistatus(v string)`

SetSpaceMultistatus sets SpaceMultistatus field to given value.

### HasSpaceMultistatus

`func (o *DataInnerIpamSpaceData) HasSpaceMultistatus() bool`

HasSpaceMultistatus returns a boolean if a field has been set.

### GetParentSpaceClassName

`func (o *DataInnerIpamSpaceData) GetParentSpaceClassName() string`

GetParentSpaceClassName returns the ParentSpaceClassName field if non-nil, zero value otherwise.

### GetParentSpaceClassNameOk

`func (o *DataInnerIpamSpaceData) GetParentSpaceClassNameOk() (*string, bool)`

GetParentSpaceClassNameOk returns a tuple with the ParentSpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceClassName

`func (o *DataInnerIpamSpaceData) SetParentSpaceClassName(v string)`

SetParentSpaceClassName sets ParentSpaceClassName field to given value.

### HasParentSpaceClassName

`func (o *DataInnerIpamSpaceData) HasParentSpaceClassName() bool`

HasParentSpaceClassName returns a boolean if a field has been set.

### GetParentSpaceClassParameters

`func (o *DataInnerIpamSpaceData) GetParentSpaceClassParameters() []ApiClassParameterOutputEntry`

GetParentSpaceClassParameters returns the ParentSpaceClassParameters field if non-nil, zero value otherwise.

### GetParentSpaceClassParametersOk

`func (o *DataInnerIpamSpaceData) GetParentSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetParentSpaceClassParametersOk returns a tuple with the ParentSpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceClassParameters

`func (o *DataInnerIpamSpaceData) SetParentSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetParentSpaceClassParameters sets ParentSpaceClassParameters field to given value.

### HasParentSpaceClassParameters

`func (o *DataInnerIpamSpaceData) HasParentSpaceClassParameters() bool`

HasParentSpaceClassParameters returns a boolean if a field has been set.

### GetParentSpaceId

`func (o *DataInnerIpamSpaceData) GetParentSpaceId() string`

GetParentSpaceId returns the ParentSpaceId field if non-nil, zero value otherwise.

### GetParentSpaceIdOk

`func (o *DataInnerIpamSpaceData) GetParentSpaceIdOk() (*string, bool)`

GetParentSpaceIdOk returns a tuple with the ParentSpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceId

`func (o *DataInnerIpamSpaceData) SetParentSpaceId(v string)`

SetParentSpaceId sets ParentSpaceId field to given value.

### HasParentSpaceId

`func (o *DataInnerIpamSpaceData) HasParentSpaceId() bool`

HasParentSpaceId returns a boolean if a field has been set.

### GetParentSpaceName

`func (o *DataInnerIpamSpaceData) GetParentSpaceName() string`

GetParentSpaceName returns the ParentSpaceName field if non-nil, zero value otherwise.

### GetParentSpaceNameOk

`func (o *DataInnerIpamSpaceData) GetParentSpaceNameOk() (*string, bool)`

GetParentSpaceNameOk returns a tuple with the ParentSpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSpaceName

`func (o *DataInnerIpamSpaceData) SetParentSpaceName(v string)`

SetParentSpaceName sets ParentSpaceName field to given value.

### HasParentSpaceName

`func (o *DataInnerIpamSpaceData) HasParentSpaceName() bool`

HasParentSpaceName returns a boolean if a field has been set.

### GetSpaceClassName

`func (o *DataInnerIpamSpaceData) GetSpaceClassName() string`

GetSpaceClassName returns the SpaceClassName field if non-nil, zero value otherwise.

### GetSpaceClassNameOk

`func (o *DataInnerIpamSpaceData) GetSpaceClassNameOk() (*string, bool)`

GetSpaceClassNameOk returns a tuple with the SpaceClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassName

`func (o *DataInnerIpamSpaceData) SetSpaceClassName(v string)`

SetSpaceClassName sets SpaceClassName field to given value.

### HasSpaceClassName

`func (o *DataInnerIpamSpaceData) HasSpaceClassName() bool`

HasSpaceClassName returns a boolean if a field has been set.

### GetSpaceClassParameters

`func (o *DataInnerIpamSpaceData) GetSpaceClassParameters() []ApiClassParameterOutputEntry`

GetSpaceClassParameters returns the SpaceClassParameters field if non-nil, zero value otherwise.

### GetSpaceClassParametersOk

`func (o *DataInnerIpamSpaceData) GetSpaceClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceClassParameters

`func (o *DataInnerIpamSpaceData) SetSpaceClassParameters(v []ApiClassParameterOutputEntry)`

SetSpaceClassParameters sets SpaceClassParameters field to given value.

### HasSpaceClassParameters

`func (o *DataInnerIpamSpaceData) HasSpaceClassParameters() bool`

HasSpaceClassParameters returns a boolean if a field has been set.

### GetSpaceDescription

`func (o *DataInnerIpamSpaceData) GetSpaceDescription() string`

GetSpaceDescription returns the SpaceDescription field if non-nil, zero value otherwise.

### GetSpaceDescriptionOk

`func (o *DataInnerIpamSpaceData) GetSpaceDescriptionOk() (*string, bool)`

GetSpaceDescriptionOk returns a tuple with the SpaceDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceDescription

`func (o *DataInnerIpamSpaceData) SetSpaceDescription(v string)`

SetSpaceDescription sets SpaceDescription field to given value.

### HasSpaceDescription

`func (o *DataInnerIpamSpaceData) HasSpaceDescription() bool`

HasSpaceDescription returns a boolean if a field has been set.

### GetSpaceId

`func (o *DataInnerIpamSpaceData) GetSpaceId() string`

GetSpaceId returns the SpaceId field if non-nil, zero value otherwise.

### GetSpaceIdOk

`func (o *DataInnerIpamSpaceData) GetSpaceIdOk() (*string, bool)`

GetSpaceIdOk returns a tuple with the SpaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceId

`func (o *DataInnerIpamSpaceData) SetSpaceId(v string)`

SetSpaceId sets SpaceId field to given value.

### HasSpaceId

`func (o *DataInnerIpamSpaceData) HasSpaceId() bool`

HasSpaceId returns a boolean if a field has been set.

### GetSpaceIsTemplate

`func (o *DataInnerIpamSpaceData) GetSpaceIsTemplate() string`

GetSpaceIsTemplate returns the SpaceIsTemplate field if non-nil, zero value otherwise.

### GetSpaceIsTemplateOk

`func (o *DataInnerIpamSpaceData) GetSpaceIsTemplateOk() (*string, bool)`

GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceIsTemplate

`func (o *DataInnerIpamSpaceData) SetSpaceIsTemplate(v string)`

SetSpaceIsTemplate sets SpaceIsTemplate field to given value.

### HasSpaceIsTemplate

`func (o *DataInnerIpamSpaceData) HasSpaceIsTemplate() bool`

HasSpaceIsTemplate returns a boolean if a field has been set.

### GetSpaceName

`func (o *DataInnerIpamSpaceData) GetSpaceName() string`

GetSpaceName returns the SpaceName field if non-nil, zero value otherwise.

### GetSpaceNameOk

`func (o *DataInnerIpamSpaceData) GetSpaceNameOk() (*string, bool)`

GetSpaceNameOk returns a tuple with the SpaceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpaceName

`func (o *DataInnerIpamSpaceData) SetSpaceName(v string)`

SetSpaceName sets SpaceName field to given value.

### HasSpaceName

`func (o *DataInnerIpamSpaceData) HasSpaceName() bool`

HasSpaceName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


