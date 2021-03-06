/*
 * SOLIDserver API
 *
 * OpenAPI 3.0.2 API definition for SOLIDserver service from EfficientIP.<p>Copyright © 2000-2021 EfficientIP</p><p><em>All specifications and information regarding the products in  this document are subject to change without notice and should not be  construed as a commitment by EfficientIP. EfficientIP assumes no  responsibility or liability for any mistakes or inaccuracies that may appear  in this document. All statements and recommendations in this document are  believed to be accurate but are presented without warranty. Users must take  full responsibility for their application of any product.</em></p><p>Generated (Monday 14th of June 2021 12:30:34 PM)</p>
 *
 * API version: 2.0
 * Contact: support-api@efficientip.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdsclient

import (
	"encoding/json"
)

// IpamSpaceEditInput struct for IpamSpaceEditInput
type IpamSpaceEditInput struct {
	// The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify which space to edit.
	SpaceId *int32 `json:"space_id,omitempty"`
	// The name of the space, each space must have a unique name.
	SpaceName *string `json:"space_name,omitempty"`
	// The database identifier (ID) of an existing space you want to set as the VLSM parent of the space you are editing. This sets up a space-based VLSM organization.
	ParentSpaceId *int32 `json:"parent_space_id,omitempty"`
	// The name of an existing space you want to set as the VLSM parent of the space you are editing. This sets up a space-based VLSM organization.
	ParentSpaceName *string `json:"parent_space_name,omitempty"`
	// The description of the space.
	SpaceDescription *string `json:"space_description,omitempty"`
	// The template status of the space you are editing. If the space is used as template (<b>1</b>), all the IPv4 networks, pools and IP addresses it contains are also used as template. You can only set this parameter once, you cannot edit its value.
	SpaceIsTemplate *int32 `json:"space_is_template,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete *[]string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	SpaceClassName *string `json:"space_class_name,omitempty"`
	// class parameters in json format
	SpaceClassParameters *[]ApiClassParameterInputEntry `json:"space_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewIpamSpaceEditInput instantiates a new IpamSpaceEditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamSpaceEditInput() *IpamSpaceEditInput {
	this := IpamSpaceEditInput{}
	return &this
}

// NewIpamSpaceEditInputWithDefaults instantiates a new IpamSpaceEditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamSpaceEditInputWithDefaults() *IpamSpaceEditInput {
	this := IpamSpaceEditInput{}
	return &this
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetSpaceId() int32 {
	if o == nil || o.SpaceId == nil {
		var ret int32
		return ret
	}
	return *o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetSpaceIdOk() (*int32, bool) {
	if o == nil || o.SpaceId == nil {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasSpaceId() bool {
	if o != nil && o.SpaceId != nil {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given int32 and assigns it to the SpaceId field.
func (o *IpamSpaceEditInput) SetSpaceId(v int32) {
	o.SpaceId = &v
}

// GetSpaceName returns the SpaceName field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetSpaceName() string {
	if o == nil || o.SpaceName == nil {
		var ret string
		return ret
	}
	return *o.SpaceName
}

// GetSpaceNameOk returns a tuple with the SpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetSpaceNameOk() (*string, bool) {
	if o == nil || o.SpaceName == nil {
		return nil, false
	}
	return o.SpaceName, true
}

// HasSpaceName returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasSpaceName() bool {
	if o != nil && o.SpaceName != nil {
		return true
	}

	return false
}

// SetSpaceName gets a reference to the given string and assigns it to the SpaceName field.
func (o *IpamSpaceEditInput) SetSpaceName(v string) {
	o.SpaceName = &v
}

// GetParentSpaceId returns the ParentSpaceId field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetParentSpaceId() int32 {
	if o == nil || o.ParentSpaceId == nil {
		var ret int32
		return ret
	}
	return *o.ParentSpaceId
}

// GetParentSpaceIdOk returns a tuple with the ParentSpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetParentSpaceIdOk() (*int32, bool) {
	if o == nil || o.ParentSpaceId == nil {
		return nil, false
	}
	return o.ParentSpaceId, true
}

// HasParentSpaceId returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasParentSpaceId() bool {
	if o != nil && o.ParentSpaceId != nil {
		return true
	}

	return false
}

// SetParentSpaceId gets a reference to the given int32 and assigns it to the ParentSpaceId field.
func (o *IpamSpaceEditInput) SetParentSpaceId(v int32) {
	o.ParentSpaceId = &v
}

// GetParentSpaceName returns the ParentSpaceName field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetParentSpaceName() string {
	if o == nil || o.ParentSpaceName == nil {
		var ret string
		return ret
	}
	return *o.ParentSpaceName
}

// GetParentSpaceNameOk returns a tuple with the ParentSpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetParentSpaceNameOk() (*string, bool) {
	if o == nil || o.ParentSpaceName == nil {
		return nil, false
	}
	return o.ParentSpaceName, true
}

// HasParentSpaceName returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasParentSpaceName() bool {
	if o != nil && o.ParentSpaceName != nil {
		return true
	}

	return false
}

// SetParentSpaceName gets a reference to the given string and assigns it to the ParentSpaceName field.
func (o *IpamSpaceEditInput) SetParentSpaceName(v string) {
	o.ParentSpaceName = &v
}

// GetSpaceDescription returns the SpaceDescription field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetSpaceDescription() string {
	if o == nil || o.SpaceDescription == nil {
		var ret string
		return ret
	}
	return *o.SpaceDescription
}

// GetSpaceDescriptionOk returns a tuple with the SpaceDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetSpaceDescriptionOk() (*string, bool) {
	if o == nil || o.SpaceDescription == nil {
		return nil, false
	}
	return o.SpaceDescription, true
}

// HasSpaceDescription returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasSpaceDescription() bool {
	if o != nil && o.SpaceDescription != nil {
		return true
	}

	return false
}

// SetSpaceDescription gets a reference to the given string and assigns it to the SpaceDescription field.
func (o *IpamSpaceEditInput) SetSpaceDescription(v string) {
	o.SpaceDescription = &v
}

// GetSpaceIsTemplate returns the SpaceIsTemplate field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetSpaceIsTemplate() int32 {
	if o == nil || o.SpaceIsTemplate == nil {
		var ret int32
		return ret
	}
	return *o.SpaceIsTemplate
}

// GetSpaceIsTemplateOk returns a tuple with the SpaceIsTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetSpaceIsTemplateOk() (*int32, bool) {
	if o == nil || o.SpaceIsTemplate == nil {
		return nil, false
	}
	return o.SpaceIsTemplate, true
}

// HasSpaceIsTemplate returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasSpaceIsTemplate() bool {
	if o != nil && o.SpaceIsTemplate != nil {
		return true
	}

	return false
}

// SetSpaceIsTemplate gets a reference to the given int32 and assigns it to the SpaceIsTemplate field.
func (o *IpamSpaceEditInput) SetSpaceIsTemplate(v int32) {
	o.SpaceIsTemplate = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetClassParametersToDelete() []string {
	if o == nil || o.ClassParametersToDelete == nil {
		var ret []string
		return ret
	}
	return *o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetClassParametersToDeleteOk() (*[]string, bool) {
	if o == nil || o.ClassParametersToDelete == nil {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasClassParametersToDelete() bool {
	if o != nil && o.ClassParametersToDelete != nil {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *IpamSpaceEditInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = &v
}

// GetSpaceClassName returns the SpaceClassName field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetSpaceClassName() string {
	if o == nil || o.SpaceClassName == nil {
		var ret string
		return ret
	}
	return *o.SpaceClassName
}

// GetSpaceClassNameOk returns a tuple with the SpaceClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetSpaceClassNameOk() (*string, bool) {
	if o == nil || o.SpaceClassName == nil {
		return nil, false
	}
	return o.SpaceClassName, true
}

// HasSpaceClassName returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasSpaceClassName() bool {
	if o != nil && o.SpaceClassName != nil {
		return true
	}

	return false
}

// SetSpaceClassName gets a reference to the given string and assigns it to the SpaceClassName field.
func (o *IpamSpaceEditInput) SetSpaceClassName(v string) {
	o.SpaceClassName = &v
}

// GetSpaceClassParameters returns the SpaceClassParameters field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetSpaceClassParameters() []ApiClassParameterInputEntry {
	if o == nil || o.SpaceClassParameters == nil {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return *o.SpaceClassParameters
}

// GetSpaceClassParametersOk returns a tuple with the SpaceClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetSpaceClassParametersOk() (*[]ApiClassParameterInputEntry, bool) {
	if o == nil || o.SpaceClassParameters == nil {
		return nil, false
	}
	return o.SpaceClassParameters, true
}

// HasSpaceClassParameters returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasSpaceClassParameters() bool {
	if o != nil && o.SpaceClassParameters != nil {
		return true
	}

	return false
}

// SetSpaceClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the SpaceClassParameters field.
func (o *IpamSpaceEditInput) SetSpaceClassParameters(v []ApiClassParameterInputEntry) {
	o.SpaceClassParameters = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *IpamSpaceEditInput) GetWarnings() string {
	if o == nil || o.Warnings == nil {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamSpaceEditInput) GetWarningsOk() (*string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *IpamSpaceEditInput) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *IpamSpaceEditInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o IpamSpaceEditInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SpaceId != nil {
		toSerialize["space_id"] = o.SpaceId
	}
	if o.SpaceName != nil {
		toSerialize["space_name"] = o.SpaceName
	}
	if o.ParentSpaceId != nil {
		toSerialize["parent_space_id"] = o.ParentSpaceId
	}
	if o.ParentSpaceName != nil {
		toSerialize["parent_space_name"] = o.ParentSpaceName
	}
	if o.SpaceDescription != nil {
		toSerialize["space_description"] = o.SpaceDescription
	}
	if o.SpaceIsTemplate != nil {
		toSerialize["space_is_template"] = o.SpaceIsTemplate
	}
	if o.ClassParametersToDelete != nil {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if o.SpaceClassName != nil {
		toSerialize["space_class_name"] = o.SpaceClassName
	}
	if o.SpaceClassParameters != nil {
		toSerialize["space_class_parameters"] = o.SpaceClassParameters
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	return json.Marshal(toSerialize)
}

type NullableIpamSpaceEditInput struct {
	value *IpamSpaceEditInput
	isSet bool
}

func (v NullableIpamSpaceEditInput) Get() *IpamSpaceEditInput {
	return v.value
}

func (v *NullableIpamSpaceEditInput) Set(val *IpamSpaceEditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamSpaceEditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamSpaceEditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamSpaceEditInput(val *IpamSpaceEditInput) *NullableIpamSpaceEditInput {
	return &NullableIpamSpaceEditInput{value: val, isSet: true}
}

func (v NullableIpamSpaceEditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamSpaceEditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


