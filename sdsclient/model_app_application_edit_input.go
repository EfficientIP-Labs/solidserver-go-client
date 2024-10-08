/*
SOLIDserver API

OpenAPI 3.0.2 API definition for SOLIDserver service from EfficientIP.<p>Copyright © 2000-2024 EfficientIP</p><p><em>All specifications and information regarding the products in this document are subject to change without notice and should not be construed as a commitment by EfficientIP. EfficientIP assumes no responsibility or liability for any mistakes or inaccuracies that may appear in this document. All statements and recommendations in this document are believed to be accurate but are presented without warranty. Users must take full responsibility for their application of any product.</em></p><p><em>This document aims at detailing EfficientIP proprietary solutions. As our solutions rely on several third-party products, created by other companies or organizations, it may redirect readers to third-party websites and documentation for further information. In such a case, EfficientIP cannot be liable or expected to provide said information on products they do maintain or created.</em></p><p>Generated (Friday 4th of October 2024 03:41:11 PM)</p>

API version: 2.0
Contact: support-api@efficientip.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdsclient

import (
	"encoding/json"
)

// checks if the AppApplicationEditInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppApplicationEditInput{}

// AppApplicationEditInput struct for AppApplicationEditInput
type AppApplicationEditInput struct {
	// The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice.
	ApplicationId *int32 `json:"application_id,omitempty"`
	// The Fully Qualified Domain Name (FQDN) of the application.
	ApplicationFqdn *string `json:"application_fqdn,omitempty"`
	// The name of the application.
	ApplicationName *string `json:"application_name,omitempty"`
	// The name of all the GSLB servers associated with the application. You can specify one or more names.
	GslbserverList *string `json:"gslbserver_list,omitempty"`
	// The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	ApplicationClassName *string `json:"application_class_name,omitempty"`
	// class parameters in json format
	ApplicationClassParameters []ApiClassParameterInputEntry `json:"application_class_parameters,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete []string `json:"class_parameters_to_delete,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewAppApplicationEditInput instantiates a new AppApplicationEditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppApplicationEditInput() *AppApplicationEditInput {
	this := AppApplicationEditInput{}
	return &this
}

// NewAppApplicationEditInputWithDefaults instantiates a new AppApplicationEditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppApplicationEditInputWithDefaults() *AppApplicationEditInput {
	this := AppApplicationEditInput{}
	return &this
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *AppApplicationEditInput) GetApplicationId() int32 {
	if o == nil || IsNil(o.ApplicationId) {
		var ret int32
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditInput) GetApplicationIdOk() (*int32, bool) {
	if o == nil || IsNil(o.ApplicationId) {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AppApplicationEditInput) HasApplicationId() bool {
	if o != nil && !IsNil(o.ApplicationId) {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given int32 and assigns it to the ApplicationId field.
func (o *AppApplicationEditInput) SetApplicationId(v int32) {
	o.ApplicationId = &v
}

// GetApplicationFqdn returns the ApplicationFqdn field value if set, zero value otherwise.
func (o *AppApplicationEditInput) GetApplicationFqdn() string {
	if o == nil || IsNil(o.ApplicationFqdn) {
		var ret string
		return ret
	}
	return *o.ApplicationFqdn
}

// GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditInput) GetApplicationFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationFqdn) {
		return nil, false
	}
	return o.ApplicationFqdn, true
}

// HasApplicationFqdn returns a boolean if a field has been set.
func (o *AppApplicationEditInput) HasApplicationFqdn() bool {
	if o != nil && !IsNil(o.ApplicationFqdn) {
		return true
	}

	return false
}

// SetApplicationFqdn gets a reference to the given string and assigns it to the ApplicationFqdn field.
func (o *AppApplicationEditInput) SetApplicationFqdn(v string) {
	o.ApplicationFqdn = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *AppApplicationEditInput) GetApplicationName() string {
	if o == nil || IsNil(o.ApplicationName) {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditInput) GetApplicationNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationName) {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *AppApplicationEditInput) HasApplicationName() bool {
	if o != nil && !IsNil(o.ApplicationName) {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *AppApplicationEditInput) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetGslbserverList returns the GslbserverList field value if set, zero value otherwise.
func (o *AppApplicationEditInput) GetGslbserverList() string {
	if o == nil || IsNil(o.GslbserverList) {
		var ret string
		return ret
	}
	return *o.GslbserverList
}

// GetGslbserverListOk returns a tuple with the GslbserverList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditInput) GetGslbserverListOk() (*string, bool) {
	if o == nil || IsNil(o.GslbserverList) {
		return nil, false
	}
	return o.GslbserverList, true
}

// HasGslbserverList returns a boolean if a field has been set.
func (o *AppApplicationEditInput) HasGslbserverList() bool {
	if o != nil && !IsNil(o.GslbserverList) {
		return true
	}

	return false
}

// SetGslbserverList gets a reference to the given string and assigns it to the GslbserverList field.
func (o *AppApplicationEditInput) SetGslbserverList(v string) {
	o.GslbserverList = &v
}

// GetApplicationClassName returns the ApplicationClassName field value if set, zero value otherwise.
func (o *AppApplicationEditInput) GetApplicationClassName() string {
	if o == nil || IsNil(o.ApplicationClassName) {
		var ret string
		return ret
	}
	return *o.ApplicationClassName
}

// GetApplicationClassNameOk returns a tuple with the ApplicationClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditInput) GetApplicationClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.ApplicationClassName) {
		return nil, false
	}
	return o.ApplicationClassName, true
}

// HasApplicationClassName returns a boolean if a field has been set.
func (o *AppApplicationEditInput) HasApplicationClassName() bool {
	if o != nil && !IsNil(o.ApplicationClassName) {
		return true
	}

	return false
}

// SetApplicationClassName gets a reference to the given string and assigns it to the ApplicationClassName field.
func (o *AppApplicationEditInput) SetApplicationClassName(v string) {
	o.ApplicationClassName = &v
}

// GetApplicationClassParameters returns the ApplicationClassParameters field value if set, zero value otherwise.
func (o *AppApplicationEditInput) GetApplicationClassParameters() []ApiClassParameterInputEntry {
	if o == nil || IsNil(o.ApplicationClassParameters) {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return o.ApplicationClassParameters
}

// GetApplicationClassParametersOk returns a tuple with the ApplicationClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditInput) GetApplicationClassParametersOk() ([]ApiClassParameterInputEntry, bool) {
	if o == nil || IsNil(o.ApplicationClassParameters) {
		return nil, false
	}
	return o.ApplicationClassParameters, true
}

// HasApplicationClassParameters returns a boolean if a field has been set.
func (o *AppApplicationEditInput) HasApplicationClassParameters() bool {
	if o != nil && !IsNil(o.ApplicationClassParameters) {
		return true
	}

	return false
}

// SetApplicationClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the ApplicationClassParameters field.
func (o *AppApplicationEditInput) SetApplicationClassParameters(v []ApiClassParameterInputEntry) {
	o.ApplicationClassParameters = v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *AppApplicationEditInput) GetClassParametersToDelete() []string {
	if o == nil || IsNil(o.ClassParametersToDelete) {
		var ret []string
		return ret
	}
	return o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditInput) GetClassParametersToDeleteOk() ([]string, bool) {
	if o == nil || IsNil(o.ClassParametersToDelete) {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *AppApplicationEditInput) HasClassParametersToDelete() bool {
	if o != nil && !IsNil(o.ClassParametersToDelete) {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *AppApplicationEditInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AppApplicationEditInput) GetWarnings() string {
	if o == nil || IsNil(o.Warnings) {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppApplicationEditInput) GetWarningsOk() (*string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AppApplicationEditInput) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *AppApplicationEditInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o AppApplicationEditInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppApplicationEditInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApplicationId) {
		toSerialize["application_id"] = o.ApplicationId
	}
	if !IsNil(o.ApplicationFqdn) {
		toSerialize["application_fqdn"] = o.ApplicationFqdn
	}
	if !IsNil(o.ApplicationName) {
		toSerialize["application_name"] = o.ApplicationName
	}
	if !IsNil(o.GslbserverList) {
		toSerialize["gslbserver_list"] = o.GslbserverList
	}
	if !IsNil(o.ApplicationClassName) {
		toSerialize["application_class_name"] = o.ApplicationClassName
	}
	if !IsNil(o.ApplicationClassParameters) {
		toSerialize["application_class_parameters"] = o.ApplicationClassParameters
	}
	if !IsNil(o.ClassParametersToDelete) {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAppApplicationEditInput struct {
	value *AppApplicationEditInput
	isSet bool
}

func (v NullableAppApplicationEditInput) Get() *AppApplicationEditInput {
	return v.value
}

func (v *NullableAppApplicationEditInput) Set(val *AppApplicationEditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableAppApplicationEditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableAppApplicationEditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppApplicationEditInput(val *AppApplicationEditInput) *NullableAppApplicationEditInput {
	return &NullableAppApplicationEditInput{value: val, isSet: true}
}

func (v NullableAppApplicationEditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppApplicationEditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
