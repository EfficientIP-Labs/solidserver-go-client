# DataInnerAppApplicationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationClassName** | Pointer to **string** | The name of the class applied to the object, it can be preceded by the class directory. | [optional] 
**ApplicationClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the application and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;...&lt;/b&gt; . | [optional] 
**ApplicationFqdn** | Pointer to **string** | The FQDN of the application. | [optional] 
**GslbserverId** | Pointer to **string** | The database identifier (ID) of the GSLB server associated with the application. It is only returned for deployed applications. | [optional] 
**GslbserverList** | Pointer to **string** | The name of all the GSLB servers associated with the application. It lists the name of each server separated by a comma. | [optional] 
**GslbserverName** | Pointer to **string** | The name of the GSLB server associated with the application. It is only returned for deployed applications. | [optional] 
**GslbserverStatus** | Pointer to **string** | The status of the GSLB server associated with the application, either OK (&lt;b&gt;1&lt;/b&gt;), GSLB Stopped (&lt;b&gt;2&lt;/b&gt;), GSLB Invalid Credentials (&lt;b&gt;4&lt;/b&gt;) or GSLB Timeout (&lt;b&gt;5&lt;/b&gt;). It is only returned for deployed applications. | [optional] 
**ApplicationId** | Pointer to **string** | The database identifier (ID) of the application. | [optional] 
**ApplicationInactiveNodes** | Pointer to **string** | The number of nodes of the application that are &lt;b&gt;Inactive&lt;/b&gt; (down). | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application. | [optional] 
**ApplicationTotalNodes** | Pointer to **string** | The number of nodes of the application. | [optional] 
**ApplicationDelayedTime** | Pointer to **string** | The delay of creation/deletion status. &lt;b&gt;1&lt;/b&gt; indicates that the object is not created/deleted yet. | [optional] 
**GslbserverType** | Pointer to **string** | The type of DNS server associated with the application. It is only returned for deployed applications. | [optional] 
**ApplicationMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ParentApplicationId** | Pointer to **string** | The database identifier (ID) of the application. It is only returned for deployed applications. | [optional] 
**ParentApplicationName** | Pointer to **string** | The name of the application. It is only returned for deployed applications. | [optional] 

## Methods

### NewDataInnerAppApplicationData

`func NewDataInnerAppApplicationData() *DataInnerAppApplicationData`

NewDataInnerAppApplicationData instantiates a new DataInnerAppApplicationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerAppApplicationDataWithDefaults

`func NewDataInnerAppApplicationDataWithDefaults() *DataInnerAppApplicationData`

NewDataInnerAppApplicationDataWithDefaults instantiates a new DataInnerAppApplicationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationClassName

`func (o *DataInnerAppApplicationData) GetApplicationClassName() string`

GetApplicationClassName returns the ApplicationClassName field if non-nil, zero value otherwise.

### GetApplicationClassNameOk

`func (o *DataInnerAppApplicationData) GetApplicationClassNameOk() (*string, bool)`

GetApplicationClassNameOk returns a tuple with the ApplicationClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationClassName

`func (o *DataInnerAppApplicationData) SetApplicationClassName(v string)`

SetApplicationClassName sets ApplicationClassName field to given value.

### HasApplicationClassName

`func (o *DataInnerAppApplicationData) HasApplicationClassName() bool`

HasApplicationClassName returns a boolean if a field has been set.

### GetApplicationClassParameters

`func (o *DataInnerAppApplicationData) GetApplicationClassParameters() []ApiClassParameterOutputEntry`

GetApplicationClassParameters returns the ApplicationClassParameters field if non-nil, zero value otherwise.

### GetApplicationClassParametersOk

`func (o *DataInnerAppApplicationData) GetApplicationClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetApplicationClassParametersOk returns a tuple with the ApplicationClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationClassParameters

`func (o *DataInnerAppApplicationData) SetApplicationClassParameters(v []ApiClassParameterOutputEntry)`

SetApplicationClassParameters sets ApplicationClassParameters field to given value.

### HasApplicationClassParameters

`func (o *DataInnerAppApplicationData) HasApplicationClassParameters() bool`

HasApplicationClassParameters returns a boolean if a field has been set.

### GetApplicationFqdn

`func (o *DataInnerAppApplicationData) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *DataInnerAppApplicationData) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *DataInnerAppApplicationData) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *DataInnerAppApplicationData) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetGslbserverId

`func (o *DataInnerAppApplicationData) GetGslbserverId() string`

GetGslbserverId returns the GslbserverId field if non-nil, zero value otherwise.

### GetGslbserverIdOk

`func (o *DataInnerAppApplicationData) GetGslbserverIdOk() (*string, bool)`

GetGslbserverIdOk returns a tuple with the GslbserverId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverId

`func (o *DataInnerAppApplicationData) SetGslbserverId(v string)`

SetGslbserverId sets GslbserverId field to given value.

### HasGslbserverId

`func (o *DataInnerAppApplicationData) HasGslbserverId() bool`

HasGslbserverId returns a boolean if a field has been set.

### GetGslbserverList

`func (o *DataInnerAppApplicationData) GetGslbserverList() string`

GetGslbserverList returns the GslbserverList field if non-nil, zero value otherwise.

### GetGslbserverListOk

`func (o *DataInnerAppApplicationData) GetGslbserverListOk() (*string, bool)`

GetGslbserverListOk returns a tuple with the GslbserverList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverList

`func (o *DataInnerAppApplicationData) SetGslbserverList(v string)`

SetGslbserverList sets GslbserverList field to given value.

### HasGslbserverList

`func (o *DataInnerAppApplicationData) HasGslbserverList() bool`

HasGslbserverList returns a boolean if a field has been set.

### GetGslbserverName

`func (o *DataInnerAppApplicationData) GetGslbserverName() string`

GetGslbserverName returns the GslbserverName field if non-nil, zero value otherwise.

### GetGslbserverNameOk

`func (o *DataInnerAppApplicationData) GetGslbserverNameOk() (*string, bool)`

GetGslbserverNameOk returns a tuple with the GslbserverName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverName

`func (o *DataInnerAppApplicationData) SetGslbserverName(v string)`

SetGslbserverName sets GslbserverName field to given value.

### HasGslbserverName

`func (o *DataInnerAppApplicationData) HasGslbserverName() bool`

HasGslbserverName returns a boolean if a field has been set.

### GetGslbserverStatus

`func (o *DataInnerAppApplicationData) GetGslbserverStatus() string`

GetGslbserverStatus returns the GslbserverStatus field if non-nil, zero value otherwise.

### GetGslbserverStatusOk

`func (o *DataInnerAppApplicationData) GetGslbserverStatusOk() (*string, bool)`

GetGslbserverStatusOk returns a tuple with the GslbserverStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverStatus

`func (o *DataInnerAppApplicationData) SetGslbserverStatus(v string)`

SetGslbserverStatus sets GslbserverStatus field to given value.

### HasGslbserverStatus

`func (o *DataInnerAppApplicationData) HasGslbserverStatus() bool`

HasGslbserverStatus returns a boolean if a field has been set.

### GetApplicationId

`func (o *DataInnerAppApplicationData) GetApplicationId() string`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *DataInnerAppApplicationData) GetApplicationIdOk() (*string, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *DataInnerAppApplicationData) SetApplicationId(v string)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *DataInnerAppApplicationData) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationInactiveNodes

`func (o *DataInnerAppApplicationData) GetApplicationInactiveNodes() string`

GetApplicationInactiveNodes returns the ApplicationInactiveNodes field if non-nil, zero value otherwise.

### GetApplicationInactiveNodesOk

`func (o *DataInnerAppApplicationData) GetApplicationInactiveNodesOk() (*string, bool)`

GetApplicationInactiveNodesOk returns a tuple with the ApplicationInactiveNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationInactiveNodes

`func (o *DataInnerAppApplicationData) SetApplicationInactiveNodes(v string)`

SetApplicationInactiveNodes sets ApplicationInactiveNodes field to given value.

### HasApplicationInactiveNodes

`func (o *DataInnerAppApplicationData) HasApplicationInactiveNodes() bool`

HasApplicationInactiveNodes returns a boolean if a field has been set.

### GetApplicationName

`func (o *DataInnerAppApplicationData) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *DataInnerAppApplicationData) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *DataInnerAppApplicationData) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *DataInnerAppApplicationData) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetApplicationTotalNodes

`func (o *DataInnerAppApplicationData) GetApplicationTotalNodes() string`

GetApplicationTotalNodes returns the ApplicationTotalNodes field if non-nil, zero value otherwise.

### GetApplicationTotalNodesOk

`func (o *DataInnerAppApplicationData) GetApplicationTotalNodesOk() (*string, bool)`

GetApplicationTotalNodesOk returns a tuple with the ApplicationTotalNodes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationTotalNodes

`func (o *DataInnerAppApplicationData) SetApplicationTotalNodes(v string)`

SetApplicationTotalNodes sets ApplicationTotalNodes field to given value.

### HasApplicationTotalNodes

`func (o *DataInnerAppApplicationData) HasApplicationTotalNodes() bool`

HasApplicationTotalNodes returns a boolean if a field has been set.

### GetApplicationDelayedTime

`func (o *DataInnerAppApplicationData) GetApplicationDelayedTime() string`

GetApplicationDelayedTime returns the ApplicationDelayedTime field if non-nil, zero value otherwise.

### GetApplicationDelayedTimeOk

`func (o *DataInnerAppApplicationData) GetApplicationDelayedTimeOk() (*string, bool)`

GetApplicationDelayedTimeOk returns a tuple with the ApplicationDelayedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationDelayedTime

`func (o *DataInnerAppApplicationData) SetApplicationDelayedTime(v string)`

SetApplicationDelayedTime sets ApplicationDelayedTime field to given value.

### HasApplicationDelayedTime

`func (o *DataInnerAppApplicationData) HasApplicationDelayedTime() bool`

HasApplicationDelayedTime returns a boolean if a field has been set.

### GetGslbserverType

`func (o *DataInnerAppApplicationData) GetGslbserverType() string`

GetGslbserverType returns the GslbserverType field if non-nil, zero value otherwise.

### GetGslbserverTypeOk

`func (o *DataInnerAppApplicationData) GetGslbserverTypeOk() (*string, bool)`

GetGslbserverTypeOk returns a tuple with the GslbserverType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverType

`func (o *DataInnerAppApplicationData) SetGslbserverType(v string)`

SetGslbserverType sets GslbserverType field to given value.

### HasGslbserverType

`func (o *DataInnerAppApplicationData) HasGslbserverType() bool`

HasGslbserverType returns a boolean if a field has been set.

### GetApplicationMultistatus

`func (o *DataInnerAppApplicationData) GetApplicationMultistatus() string`

GetApplicationMultistatus returns the ApplicationMultistatus field if non-nil, zero value otherwise.

### GetApplicationMultistatusOk

`func (o *DataInnerAppApplicationData) GetApplicationMultistatusOk() (*string, bool)`

GetApplicationMultistatusOk returns a tuple with the ApplicationMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationMultistatus

`func (o *DataInnerAppApplicationData) SetApplicationMultistatus(v string)`

SetApplicationMultistatus sets ApplicationMultistatus field to given value.

### HasApplicationMultistatus

`func (o *DataInnerAppApplicationData) HasApplicationMultistatus() bool`

HasApplicationMultistatus returns a boolean if a field has been set.

### GetParentApplicationId

`func (o *DataInnerAppApplicationData) GetParentApplicationId() string`

GetParentApplicationId returns the ParentApplicationId field if non-nil, zero value otherwise.

### GetParentApplicationIdOk

`func (o *DataInnerAppApplicationData) GetParentApplicationIdOk() (*string, bool)`

GetParentApplicationIdOk returns a tuple with the ParentApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentApplicationId

`func (o *DataInnerAppApplicationData) SetParentApplicationId(v string)`

SetParentApplicationId sets ParentApplicationId field to given value.

### HasParentApplicationId

`func (o *DataInnerAppApplicationData) HasParentApplicationId() bool`

HasParentApplicationId returns a boolean if a field has been set.

### GetParentApplicationName

`func (o *DataInnerAppApplicationData) GetParentApplicationName() string`

GetParentApplicationName returns the ParentApplicationName field if non-nil, zero value otherwise.

### GetParentApplicationNameOk

`func (o *DataInnerAppApplicationData) GetParentApplicationNameOk() (*string, bool)`

GetParentApplicationNameOk returns a tuple with the ParentApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentApplicationName

`func (o *DataInnerAppApplicationData) SetParentApplicationName(v string)`

SetParentApplicationName sets ParentApplicationName field to given value.

### HasParentApplicationName

`func (o *DataInnerAppApplicationData) HasParentApplicationName() bool`

HasParentApplicationName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


