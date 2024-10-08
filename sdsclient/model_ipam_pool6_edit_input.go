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

// checks if the IpamPool6EditInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IpamPool6EditInput{}

// IpamPool6EditInput struct for IpamPool6EditInput
type IpamPool6EditInput struct {
	// The last IP address of the pool.
	Pool6EndIpAddr *string `json:"pool6_end_ip_addr,omitempty"`
	// The database identifier (ID) of the IPv6 pool, a unique numeric key value automatically incremented when you add an IPv6 pool. Use the ID to specify the IPv6 pool of your choice.
	Pool6Id *int32 `json:"pool6_id,omitempty"`
	// The database identifier (ID) of the space, a unique numeric key value automatically incremented when you add a space. Use the ID to specify the space of your choice.
	SpaceId *int32 `json:"space_id,omitempty"`
	// The name of the space.
	SpaceName *string `json:"space_name,omitempty"`
	// The first IP address of the pool.
	Pool6StartIpAddr *string `json:"pool6_start_ip_addr,omitempty"`
	// The database identifier (ID) of the IPv6 network, a unique numeric key value automatically incremented when you add an IPv6 network. Use the ID to specify the IPv6 network of your choice.
	Network6Id *int32 `json:"network6_id,omitempty"`
	// The name of the IPv6 pool, each IPv6 pool must have a unique name.
	Pool6Name *string `json:"pool6_name,omitempty"`
	// The reservation status of the IPv6 pool. If set 1, the IP addresses it contains cannot be assigned.
	Pool6ReadOnly *int32 `json:"pool6_read_only,omitempty"`
	// class parameters you want to delete from the object
	ClassParametersToDelete []string `json:"class_parameters_to_delete,omitempty"`
	// The name of the class to apply to the object you are editing. You must specify the class file directory, e.g. <b>my_directory/my_class.class</b> . You cannot use the classes <b>global</b> and <b>default</b>, they are reserved by the system.
	Pool6ClassName *string `json:"pool6_class_name,omitempty"`
	// class parameters in json format
	Pool6ClassParameters []ApiClassParameterInputEntry `json:"pool6_class_parameters,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewIpamPool6EditInput instantiates a new IpamPool6EditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpamPool6EditInput() *IpamPool6EditInput {
	this := IpamPool6EditInput{}
	return &this
}

// NewIpamPool6EditInputWithDefaults instantiates a new IpamPool6EditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpamPool6EditInputWithDefaults() *IpamPool6EditInput {
	this := IpamPool6EditInput{}
	return &this
}

// GetPool6EndIpAddr returns the Pool6EndIpAddr field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetPool6EndIpAddr() string {
	if o == nil || IsNil(o.Pool6EndIpAddr) {
		var ret string
		return ret
	}
	return *o.Pool6EndIpAddr
}

// GetPool6EndIpAddrOk returns a tuple with the Pool6EndIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetPool6EndIpAddrOk() (*string, bool) {
	if o == nil || IsNil(o.Pool6EndIpAddr) {
		return nil, false
	}
	return o.Pool6EndIpAddr, true
}

// HasPool6EndIpAddr returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasPool6EndIpAddr() bool {
	if o != nil && !IsNil(o.Pool6EndIpAddr) {
		return true
	}

	return false
}

// SetPool6EndIpAddr gets a reference to the given string and assigns it to the Pool6EndIpAddr field.
func (o *IpamPool6EditInput) SetPool6EndIpAddr(v string) {
	o.Pool6EndIpAddr = &v
}

// GetPool6Id returns the Pool6Id field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetPool6Id() int32 {
	if o == nil || IsNil(o.Pool6Id) {
		var ret int32
		return ret
	}
	return *o.Pool6Id
}

// GetPool6IdOk returns a tuple with the Pool6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetPool6IdOk() (*int32, bool) {
	if o == nil || IsNil(o.Pool6Id) {
		return nil, false
	}
	return o.Pool6Id, true
}

// HasPool6Id returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasPool6Id() bool {
	if o != nil && !IsNil(o.Pool6Id) {
		return true
	}

	return false
}

// SetPool6Id gets a reference to the given int32 and assigns it to the Pool6Id field.
func (o *IpamPool6EditInput) SetPool6Id(v int32) {
	o.Pool6Id = &v
}

// GetSpaceId returns the SpaceId field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetSpaceId() int32 {
	if o == nil || IsNil(o.SpaceId) {
		var ret int32
		return ret
	}
	return *o.SpaceId
}

// GetSpaceIdOk returns a tuple with the SpaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetSpaceIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SpaceId) {
		return nil, false
	}
	return o.SpaceId, true
}

// HasSpaceId returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasSpaceId() bool {
	if o != nil && !IsNil(o.SpaceId) {
		return true
	}

	return false
}

// SetSpaceId gets a reference to the given int32 and assigns it to the SpaceId field.
func (o *IpamPool6EditInput) SetSpaceId(v int32) {
	o.SpaceId = &v
}

// GetSpaceName returns the SpaceName field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetSpaceName() string {
	if o == nil || IsNil(o.SpaceName) {
		var ret string
		return ret
	}
	return *o.SpaceName
}

// GetSpaceNameOk returns a tuple with the SpaceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetSpaceNameOk() (*string, bool) {
	if o == nil || IsNil(o.SpaceName) {
		return nil, false
	}
	return o.SpaceName, true
}

// HasSpaceName returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasSpaceName() bool {
	if o != nil && !IsNil(o.SpaceName) {
		return true
	}

	return false
}

// SetSpaceName gets a reference to the given string and assigns it to the SpaceName field.
func (o *IpamPool6EditInput) SetSpaceName(v string) {
	o.SpaceName = &v
}

// GetPool6StartIpAddr returns the Pool6StartIpAddr field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetPool6StartIpAddr() string {
	if o == nil || IsNil(o.Pool6StartIpAddr) {
		var ret string
		return ret
	}
	return *o.Pool6StartIpAddr
}

// GetPool6StartIpAddrOk returns a tuple with the Pool6StartIpAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetPool6StartIpAddrOk() (*string, bool) {
	if o == nil || IsNil(o.Pool6StartIpAddr) {
		return nil, false
	}
	return o.Pool6StartIpAddr, true
}

// HasPool6StartIpAddr returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasPool6StartIpAddr() bool {
	if o != nil && !IsNil(o.Pool6StartIpAddr) {
		return true
	}

	return false
}

// SetPool6StartIpAddr gets a reference to the given string and assigns it to the Pool6StartIpAddr field.
func (o *IpamPool6EditInput) SetPool6StartIpAddr(v string) {
	o.Pool6StartIpAddr = &v
}

// GetNetwork6Id returns the Network6Id field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetNetwork6Id() int32 {
	if o == nil || IsNil(o.Network6Id) {
		var ret int32
		return ret
	}
	return *o.Network6Id
}

// GetNetwork6IdOk returns a tuple with the Network6Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetNetwork6IdOk() (*int32, bool) {
	if o == nil || IsNil(o.Network6Id) {
		return nil, false
	}
	return o.Network6Id, true
}

// HasNetwork6Id returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasNetwork6Id() bool {
	if o != nil && !IsNil(o.Network6Id) {
		return true
	}

	return false
}

// SetNetwork6Id gets a reference to the given int32 and assigns it to the Network6Id field.
func (o *IpamPool6EditInput) SetNetwork6Id(v int32) {
	o.Network6Id = &v
}

// GetPool6Name returns the Pool6Name field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetPool6Name() string {
	if o == nil || IsNil(o.Pool6Name) {
		var ret string
		return ret
	}
	return *o.Pool6Name
}

// GetPool6NameOk returns a tuple with the Pool6Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetPool6NameOk() (*string, bool) {
	if o == nil || IsNil(o.Pool6Name) {
		return nil, false
	}
	return o.Pool6Name, true
}

// HasPool6Name returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasPool6Name() bool {
	if o != nil && !IsNil(o.Pool6Name) {
		return true
	}

	return false
}

// SetPool6Name gets a reference to the given string and assigns it to the Pool6Name field.
func (o *IpamPool6EditInput) SetPool6Name(v string) {
	o.Pool6Name = &v
}

// GetPool6ReadOnly returns the Pool6ReadOnly field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetPool6ReadOnly() int32 {
	if o == nil || IsNil(o.Pool6ReadOnly) {
		var ret int32
		return ret
	}
	return *o.Pool6ReadOnly
}

// GetPool6ReadOnlyOk returns a tuple with the Pool6ReadOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetPool6ReadOnlyOk() (*int32, bool) {
	if o == nil || IsNil(o.Pool6ReadOnly) {
		return nil, false
	}
	return o.Pool6ReadOnly, true
}

// HasPool6ReadOnly returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasPool6ReadOnly() bool {
	if o != nil && !IsNil(o.Pool6ReadOnly) {
		return true
	}

	return false
}

// SetPool6ReadOnly gets a reference to the given int32 and assigns it to the Pool6ReadOnly field.
func (o *IpamPool6EditInput) SetPool6ReadOnly(v int32) {
	o.Pool6ReadOnly = &v
}

// GetClassParametersToDelete returns the ClassParametersToDelete field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetClassParametersToDelete() []string {
	if o == nil || IsNil(o.ClassParametersToDelete) {
		var ret []string
		return ret
	}
	return o.ClassParametersToDelete
}

// GetClassParametersToDeleteOk returns a tuple with the ClassParametersToDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetClassParametersToDeleteOk() ([]string, bool) {
	if o == nil || IsNil(o.ClassParametersToDelete) {
		return nil, false
	}
	return o.ClassParametersToDelete, true
}

// HasClassParametersToDelete returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasClassParametersToDelete() bool {
	if o != nil && !IsNil(o.ClassParametersToDelete) {
		return true
	}

	return false
}

// SetClassParametersToDelete gets a reference to the given []string and assigns it to the ClassParametersToDelete field.
func (o *IpamPool6EditInput) SetClassParametersToDelete(v []string) {
	o.ClassParametersToDelete = v
}

// GetPool6ClassName returns the Pool6ClassName field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetPool6ClassName() string {
	if o == nil || IsNil(o.Pool6ClassName) {
		var ret string
		return ret
	}
	return *o.Pool6ClassName
}

// GetPool6ClassNameOk returns a tuple with the Pool6ClassName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetPool6ClassNameOk() (*string, bool) {
	if o == nil || IsNil(o.Pool6ClassName) {
		return nil, false
	}
	return o.Pool6ClassName, true
}

// HasPool6ClassName returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasPool6ClassName() bool {
	if o != nil && !IsNil(o.Pool6ClassName) {
		return true
	}

	return false
}

// SetPool6ClassName gets a reference to the given string and assigns it to the Pool6ClassName field.
func (o *IpamPool6EditInput) SetPool6ClassName(v string) {
	o.Pool6ClassName = &v
}

// GetPool6ClassParameters returns the Pool6ClassParameters field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetPool6ClassParameters() []ApiClassParameterInputEntry {
	if o == nil || IsNil(o.Pool6ClassParameters) {
		var ret []ApiClassParameterInputEntry
		return ret
	}
	return o.Pool6ClassParameters
}

// GetPool6ClassParametersOk returns a tuple with the Pool6ClassParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetPool6ClassParametersOk() ([]ApiClassParameterInputEntry, bool) {
	if o == nil || IsNil(o.Pool6ClassParameters) {
		return nil, false
	}
	return o.Pool6ClassParameters, true
}

// HasPool6ClassParameters returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasPool6ClassParameters() bool {
	if o != nil && !IsNil(o.Pool6ClassParameters) {
		return true
	}

	return false
}

// SetPool6ClassParameters gets a reference to the given []ApiClassParameterInputEntry and assigns it to the Pool6ClassParameters field.
func (o *IpamPool6EditInput) SetPool6ClassParameters(v []ApiClassParameterInputEntry) {
	o.Pool6ClassParameters = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *IpamPool6EditInput) GetWarnings() string {
	if o == nil || IsNil(o.Warnings) {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IpamPool6EditInput) GetWarningsOk() (*string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *IpamPool6EditInput) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *IpamPool6EditInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o IpamPool6EditInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IpamPool6EditInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pool6EndIpAddr) {
		toSerialize["pool6_end_ip_addr"] = o.Pool6EndIpAddr
	}
	if !IsNil(o.Pool6Id) {
		toSerialize["pool6_id"] = o.Pool6Id
	}
	if !IsNil(o.SpaceId) {
		toSerialize["space_id"] = o.SpaceId
	}
	if !IsNil(o.SpaceName) {
		toSerialize["space_name"] = o.SpaceName
	}
	if !IsNil(o.Pool6StartIpAddr) {
		toSerialize["pool6_start_ip_addr"] = o.Pool6StartIpAddr
	}
	if !IsNil(o.Network6Id) {
		toSerialize["network6_id"] = o.Network6Id
	}
	if !IsNil(o.Pool6Name) {
		toSerialize["pool6_name"] = o.Pool6Name
	}
	if !IsNil(o.Pool6ReadOnly) {
		toSerialize["pool6_read_only"] = o.Pool6ReadOnly
	}
	if !IsNil(o.ClassParametersToDelete) {
		toSerialize["class_parameters_to_delete"] = o.ClassParametersToDelete
	}
	if !IsNil(o.Pool6ClassName) {
		toSerialize["pool6_class_name"] = o.Pool6ClassName
	}
	if !IsNil(o.Pool6ClassParameters) {
		toSerialize["pool6_class_parameters"] = o.Pool6ClassParameters
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableIpamPool6EditInput struct {
	value *IpamPool6EditInput
	isSet bool
}

func (v NullableIpamPool6EditInput) Get() *IpamPool6EditInput {
	return v.value
}

func (v *NullableIpamPool6EditInput) Set(val *IpamPool6EditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableIpamPool6EditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableIpamPool6EditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpamPool6EditInput(val *IpamPool6EditInput) *NullableIpamPool6EditInput {
	return &NullableIpamPool6EditInput{value: val, isSet: true}
}

func (v NullableIpamPool6EditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpamPool6EditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
