# AppApplicationAddInput

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationFqdn** | Pointer to **string** | The Fully Qualified Domain Name (FQDN) of the application. | [optional] 
**ApplicationName** | Pointer to **string** | The name of the application. | [optional] 
**GslbserverList** | Pointer to **string** | The name of all the GSLB servers associated with the application. You can specify one or more names. | [optional] 
**ApplicationClassName** | Pointer to **string** | The name of the class to apply to the object you are adding. You must specify the class file directory, e.g. &lt;b&gt;my_directory/my_class.class&lt;/b&gt; . You cannot use the classes &lt;b&gt;global&lt;/b&gt; and &lt;b&gt;default&lt;/b&gt;, they are reserved by the system. | [optional] 
**ApplicationClassParameters** | Pointer to [**[]ApiClassParameterInputEntry**](ApiClassParameterInputEntry.md) | class parameters in json format | [optional] 
**ClassParametersToDelete** | Pointer to **[]string** | class parameters you want to delete from the object | [optional] 
**Warnings** | Pointer to **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | [optional] 

## Methods

### NewAppApplicationAddInput

`func NewAppApplicationAddInput() *AppApplicationAddInput`

NewAppApplicationAddInput instantiates a new AppApplicationAddInput object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppApplicationAddInputWithDefaults

`func NewAppApplicationAddInputWithDefaults() *AppApplicationAddInput`

NewAppApplicationAddInputWithDefaults instantiates a new AppApplicationAddInput object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationFqdn

`func (o *AppApplicationAddInput) GetApplicationFqdn() string`

GetApplicationFqdn returns the ApplicationFqdn field if non-nil, zero value otherwise.

### GetApplicationFqdnOk

`func (o *AppApplicationAddInput) GetApplicationFqdnOk() (*string, bool)`

GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationFqdn

`func (o *AppApplicationAddInput) SetApplicationFqdn(v string)`

SetApplicationFqdn sets ApplicationFqdn field to given value.

### HasApplicationFqdn

`func (o *AppApplicationAddInput) HasApplicationFqdn() bool`

HasApplicationFqdn returns a boolean if a field has been set.

### GetApplicationName

`func (o *AppApplicationAddInput) GetApplicationName() string`

GetApplicationName returns the ApplicationName field if non-nil, zero value otherwise.

### GetApplicationNameOk

`func (o *AppApplicationAddInput) GetApplicationNameOk() (*string, bool)`

GetApplicationNameOk returns a tuple with the ApplicationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationName

`func (o *AppApplicationAddInput) SetApplicationName(v string)`

SetApplicationName sets ApplicationName field to given value.

### HasApplicationName

`func (o *AppApplicationAddInput) HasApplicationName() bool`

HasApplicationName returns a boolean if a field has been set.

### GetGslbserverList

`func (o *AppApplicationAddInput) GetGslbserverList() string`

GetGslbserverList returns the GslbserverList field if non-nil, zero value otherwise.

### GetGslbserverListOk

`func (o *AppApplicationAddInput) GetGslbserverListOk() (*string, bool)`

GetGslbserverListOk returns a tuple with the GslbserverList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGslbserverList

`func (o *AppApplicationAddInput) SetGslbserverList(v string)`

SetGslbserverList sets GslbserverList field to given value.

### HasGslbserverList

`func (o *AppApplicationAddInput) HasGslbserverList() bool`

HasGslbserverList returns a boolean if a field has been set.

### GetApplicationClassName

`func (o *AppApplicationAddInput) GetApplicationClassName() string`

GetApplicationClassName returns the ApplicationClassName field if non-nil, zero value otherwise.

### GetApplicationClassNameOk

`func (o *AppApplicationAddInput) GetApplicationClassNameOk() (*string, bool)`

GetApplicationClassNameOk returns a tuple with the ApplicationClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationClassName

`func (o *AppApplicationAddInput) SetApplicationClassName(v string)`

SetApplicationClassName sets ApplicationClassName field to given value.

### HasApplicationClassName

`func (o *AppApplicationAddInput) HasApplicationClassName() bool`

HasApplicationClassName returns a boolean if a field has been set.

### GetApplicationClassParameters

`func (o *AppApplicationAddInput) GetApplicationClassParameters() []ApiClassParameterInputEntry`

GetApplicationClassParameters returns the ApplicationClassParameters field if non-nil, zero value otherwise.

### GetApplicationClassParametersOk

`func (o *AppApplicationAddInput) GetApplicationClassParametersOk() (*[]ApiClassParameterInputEntry, bool)`

GetApplicationClassParametersOk returns a tuple with the ApplicationClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationClassParameters

`func (o *AppApplicationAddInput) SetApplicationClassParameters(v []ApiClassParameterInputEntry)`

SetApplicationClassParameters sets ApplicationClassParameters field to given value.

### HasApplicationClassParameters

`func (o *AppApplicationAddInput) HasApplicationClassParameters() bool`

HasApplicationClassParameters returns a boolean if a field has been set.

### GetClassParametersToDelete

`func (o *AppApplicationAddInput) GetClassParametersToDelete() []string`

GetClassParametersToDelete returns the ClassParametersToDelete field if non-nil, zero value otherwise.

### GetClassParametersToDeleteOk

`func (o *AppApplicationAddInput) GetClassParametersToDeleteOk() (*[]string, bool)`

GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassParametersToDelete

`func (o *AppApplicationAddInput) SetClassParametersToDelete(v []string)`

SetClassParametersToDelete sets ClassParametersToDelete field to given value.

### HasClassParametersToDelete

`func (o *AppApplicationAddInput) HasClassParametersToDelete() bool`

HasClassParametersToDelete returns a boolean if a field has been set.

### GetWarnings

`func (o *AppApplicationAddInput) GetWarnings() string`

GetWarnings returns the Warnings field if non-nil, zero value otherwise.

### GetWarningsOk

`func (o *AppApplicationAddInput) GetWarningsOk() (*string, bool)`

GetWarningsOk returns a tuple with the Warnings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarnings

`func (o *AppApplicationAddInput) SetWarnings(v string)`

SetWarnings sets Warnings field to given value.

### HasWarnings

`func (o *AppApplicationAddInput) HasWarnings() bool`

HasWarnings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


