# AppApplicationEditInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationId** | Pointer to **int32** | The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice. | [optional] 
**ApplicationFqdn** | Pointer to **string** | The Fully Qualified Domain Name (FQDN) of the application. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application. | [optional] 
**GslbserverList** | Pointer to **string** | The name of all the GSLB servers associated with the application. You can specify one or more names. | [optional] 
**ApplicationClassName** | Pointer to **string** | The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**ApplicationClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewAppApplicationEditInput

`func NewAppApplicationEditInput() *AppApplicationEditInput`

NewAppApplicationEditInput instantiates a new AppApplicationEditInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApplicationEditInputWithDefaults

`func NewAppApplicationEditInputWithDefaults() *AppApplicationEditInput`

NewAppApplicationEditInputWithDefaults instantiates a new AppApplicationEditInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationId

`func (o *AppApplicationEditInput) GetApplicationId() int32`

GetApplicationId returns the ApplicationId field if non-nil, zero value otherwise.

### GetApplicationIdOk

`func (o *AppApplicationEditInput) GetApplicationIdOk() (*int32, bool)`

GetApplicationIdOk returns a tuple with the ApplicationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationId

`func (o *AppApplicationEditInput) SetApplicationId(v int32)`

SetApplicationId sets ApplicationId field to given value.

### HasApplicationId

`func (o *AppApplicationEditInput) HasApplicationId() bool`

HasApplicationId returns a boolean if a field has been set.

### GetApplicationFqdn

`func (o *AppApplicationEditInput) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *AppApplicationEditInput) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *AppApplicationEditInput) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *AppApplicationEditInput) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetApplicationName

`func (o *AppApplicationEditInput) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AppApplicationEditInput) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AppApplicationEditInput) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AppApplicationEditInput) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetGslbserverList

`func (o *AppApplicationEditInput) GetGslbserverList() string`

GetGslbserverList returns the GslbserverList field if non-nil, zero value otherwise.

### GetGslbserverListOk

`func (o *AppApplicationEditInput) GetGslbserverListOk() (*string, bool)`

GetGslbserverListOk returns a tuple with the GslbserverList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverList

`func (o *AppApplicationEditInput) SetGslbserverList(v string)`

SetGslbserverList sets GslbserverList field to given value.

### HasGslbserverList

`func (o *AppApplicationEditInput) HasGslbserverList() bool`

HasGslbserverList returns a boolean if a field has been set.

### GetApplicationClassName

`func (o *AppApplicationEditInput) GetApplicationClassName() string`

GetApplicationClassName returns the ApplicationClassName field if non-nil, zero value otherwise.

### GetApplicationClassNameOk

`func (o *AppApplicationEditInput) GetApplicationClassNameOk() (*string, bool)`

GetApplicationClassNameOk returns a tuple with the ApplicationClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationClassName

`func (o *AppApplicationEditInput) SetApplicationClassName(v string)`

SetApplicationClassName sets ApplicationClassName field to given value.

### HasApplicationClassName

`func (o *AppApplicationEditInput) HasApplicationClassName() bool`

HasApplicationClassName returns a boolean if a field has been set.

### GetApplicationClassParameters

`func (o *AppApplicationEditInput) GetApplicationClassParameters() []ApiClassParameterInputEntry`

GetApplicationClassParameters returns the ApplicationClassParameters field if non-nil, zero value otherwise.

### GetApplicationClassParametersOk

`func (o *AppApplicationEditInput) GetApplicationClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetApplicationClassParametersOk returns a tuple with the ApplicationClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationClassParameters

`func (o *AppApplicationEditInput) SetApplicationClassParameters(v []ApiClassParameterInputEntry)`

SetApplicationClassParameters sets ApplicationClassParameters field to given value.

### HasApplicationClassParameters

`func (o *AppApplicationEditInput) HasApplicationClassParameters() bool`

HasApplicationClassParameters returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *AppApplicationEditInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *AppApplicationEditInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *AppApplicationEditInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *AppApplicationEditInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetWarnings

`func (o *AppApplicationEditInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AppApplicationEditInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AppApplicationEditInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AppApplicationEditInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


