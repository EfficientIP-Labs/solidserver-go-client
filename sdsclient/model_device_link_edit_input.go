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

// checks if the DeviceLinkEditInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceLinkEditInput{}

// DeviceLinkEditInput struct for DeviceLinkEditInput
type DeviceLinkEditInput struct {
	// The name of the device to which belongs the DM port or interface you want to link with <b>interface2_id</b>.
	Device1Name *string `json:"device1_name,omitempty"`
	// The name of the device to which belongs the DM port or interface you want to link with <b>interface1_id</b>.
	Device2Name *string `json:"device2_name,omitempty"`
	// The database identifier (ID) of the DM port or interface you want to link with <b>interface2_id</b>.
	Interface1Id *int32 `json:"interface1_id,omitempty"`
	// The name of the DM port or interface you want to link with <b>interface2_id</b>.
	Interface1Name *string `json:"interface1_name,omitempty"`
	// The database identifier (ID) of the DM port or interface you want to link with <b>interface1_id</b>.
	Interface2Id *int32 `json:"interface2_id,omitempty"`
	// The name of the DM port or interface you want to link with <b>interface1_id</b>.
	Interface2Name *string `json:"interface2_name,omitempty"`
	// The database identifier (ID) of the Device Manager port or interface link, a unique numeric key value automatically incremented when you add a link between a device and a port or interface. Use the ID to specify the port or interface link of your choice.
	LinkId *int32 `json:"link_id,omitempty"`
	// A way to determine if the link between two Device Manager devices is set automatically (<b>1</b>) or manually (<b>0</b>).
	LinkAutoLink *int32 `json:"link_auto_link,omitempty"`
	// A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules.
	Warnings *string `json:"warnings,omitempty"`
}

// NewDeviceLinkEditInput instantiates a new DeviceLinkEditInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceLinkEditInput() *DeviceLinkEditInput {
	this := DeviceLinkEditInput{}
	return &this
}

// NewDeviceLinkEditInputWithDefaults instantiates a new DeviceLinkEditInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceLinkEditInputWithDefaults() *DeviceLinkEditInput {
	this := DeviceLinkEditInput{}
	return &this
}

// GetDevice1Name returns the Device1Name field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetDevice1Name() string {
	if o == nil || IsNil(o.Device1Name) {
		var ret string
		return ret
	}
	return *o.Device1Name
}

// GetDevice1NameOk returns a tuple with the Device1Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetDevice1NameOk() (*string, bool) {
	if o == nil || IsNil(o.Device1Name) {
		return nil, false
	}
	return o.Device1Name, true
}

// HasDevice1Name returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasDevice1Name() bool {
	if o != nil && !IsNil(o.Device1Name) {
		return true
	}

	return false
}

// SetDevice1Name gets a reference to the given string and assigns it to the Device1Name field.
func (o *DeviceLinkEditInput) SetDevice1Name(v string) {
	o.Device1Name = &v
}

// GetDevice2Name returns the Device2Name field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetDevice2Name() string {
	if o == nil || IsNil(o.Device2Name) {
		var ret string
		return ret
	}
	return *o.Device2Name
}

// GetDevice2NameOk returns a tuple with the Device2Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetDevice2NameOk() (*string, bool) {
	if o == nil || IsNil(o.Device2Name) {
		return nil, false
	}
	return o.Device2Name, true
}

// HasDevice2Name returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasDevice2Name() bool {
	if o != nil && !IsNil(o.Device2Name) {
		return true
	}

	return false
}

// SetDevice2Name gets a reference to the given string and assigns it to the Device2Name field.
func (o *DeviceLinkEditInput) SetDevice2Name(v string) {
	o.Device2Name = &v
}

// GetInterface1Id returns the Interface1Id field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetInterface1Id() int32 {
	if o == nil || IsNil(o.Interface1Id) {
		var ret int32
		return ret
	}
	return *o.Interface1Id
}

// GetInterface1IdOk returns a tuple with the Interface1Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetInterface1IdOk() (*int32, bool) {
	if o == nil || IsNil(o.Interface1Id) {
		return nil, false
	}
	return o.Interface1Id, true
}

// HasInterface1Id returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasInterface1Id() bool {
	if o != nil && !IsNil(o.Interface1Id) {
		return true
	}

	return false
}

// SetInterface1Id gets a reference to the given int32 and assigns it to the Interface1Id field.
func (o *DeviceLinkEditInput) SetInterface1Id(v int32) {
	o.Interface1Id = &v
}

// GetInterface1Name returns the Interface1Name field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetInterface1Name() string {
	if o == nil || IsNil(o.Interface1Name) {
		var ret string
		return ret
	}
	return *o.Interface1Name
}

// GetInterface1NameOk returns a tuple with the Interface1Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetInterface1NameOk() (*string, bool) {
	if o == nil || IsNil(o.Interface1Name) {
		return nil, false
	}
	return o.Interface1Name, true
}

// HasInterface1Name returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasInterface1Name() bool {
	if o != nil && !IsNil(o.Interface1Name) {
		return true
	}

	return false
}

// SetInterface1Name gets a reference to the given string and assigns it to the Interface1Name field.
func (o *DeviceLinkEditInput) SetInterface1Name(v string) {
	o.Interface1Name = &v
}

// GetInterface2Id returns the Interface2Id field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetInterface2Id() int32 {
	if o == nil || IsNil(o.Interface2Id) {
		var ret int32
		return ret
	}
	return *o.Interface2Id
}

// GetInterface2IdOk returns a tuple with the Interface2Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetInterface2IdOk() (*int32, bool) {
	if o == nil || IsNil(o.Interface2Id) {
		return nil, false
	}
	return o.Interface2Id, true
}

// HasInterface2Id returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasInterface2Id() bool {
	if o != nil && !IsNil(o.Interface2Id) {
		return true
	}

	return false
}

// SetInterface2Id gets a reference to the given int32 and assigns it to the Interface2Id field.
func (o *DeviceLinkEditInput) SetInterface2Id(v int32) {
	o.Interface2Id = &v
}

// GetInterface2Name returns the Interface2Name field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetInterface2Name() string {
	if o == nil || IsNil(o.Interface2Name) {
		var ret string
		return ret
	}
	return *o.Interface2Name
}

// GetInterface2NameOk returns a tuple with the Interface2Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetInterface2NameOk() (*string, bool) {
	if o == nil || IsNil(o.Interface2Name) {
		return nil, false
	}
	return o.Interface2Name, true
}

// HasInterface2Name returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasInterface2Name() bool {
	if o != nil && !IsNil(o.Interface2Name) {
		return true
	}

	return false
}

// SetInterface2Name gets a reference to the given string and assigns it to the Interface2Name field.
func (o *DeviceLinkEditInput) SetInterface2Name(v string) {
	o.Interface2Name = &v
}

// GetLinkId returns the LinkId field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetLinkId() int32 {
	if o == nil || IsNil(o.LinkId) {
		var ret int32
		return ret
	}
	return *o.LinkId
}

// GetLinkIdOk returns a tuple with the LinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetLinkIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkId) {
		return nil, false
	}
	return o.LinkId, true
}

// HasLinkId returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasLinkId() bool {
	if o != nil && !IsNil(o.LinkId) {
		return true
	}

	return false
}

// SetLinkId gets a reference to the given int32 and assigns it to the LinkId field.
func (o *DeviceLinkEditInput) SetLinkId(v int32) {
	o.LinkId = &v
}

// GetLinkAutoLink returns the LinkAutoLink field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetLinkAutoLink() int32 {
	if o == nil || IsNil(o.LinkAutoLink) {
		var ret int32
		return ret
	}
	return *o.LinkAutoLink
}

// GetLinkAutoLinkOk returns a tuple with the LinkAutoLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetLinkAutoLinkOk() (*int32, bool) {
	if o == nil || IsNil(o.LinkAutoLink) {
		return nil, false
	}
	return o.LinkAutoLink, true
}

// HasLinkAutoLink returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasLinkAutoLink() bool {
	if o != nil && !IsNil(o.LinkAutoLink) {
		return true
	}

	return false
}

// SetLinkAutoLink gets a reference to the given int32 and assigns it to the LinkAutoLink field.
func (o *DeviceLinkEditInput) SetLinkAutoLink(v int32) {
	o.LinkAutoLink = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DeviceLinkEditInput) GetWarnings() string {
	if o == nil || IsNil(o.Warnings) {
		var ret string
		return ret
	}
	return *o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceLinkEditInput) GetWarningsOk() (*string, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DeviceLinkEditInput) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given string and assigns it to the Warnings field.
func (o *DeviceLinkEditInput) SetWarnings(v string) {
	o.Warnings = &v
}

func (o DeviceLinkEditInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceLinkEditInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Device1Name) {
		toSerialize["device1_name"] = o.Device1Name
	}
	if !IsNil(o.Device2Name) {
		toSerialize["device2_name"] = o.Device2Name
	}
	if !IsNil(o.Interface1Id) {
		toSerialize["interface1_id"] = o.Interface1Id
	}
	if !IsNil(o.Interface1Name) {
		toSerialize["interface1_name"] = o.Interface1Name
	}
	if !IsNil(o.Interface2Id) {
		toSerialize["interface2_id"] = o.Interface2Id
	}
	if !IsNil(o.Interface2Name) {
		toSerialize["interface2_name"] = o.Interface2Name
	}
	if !IsNil(o.LinkId) {
		toSerialize["link_id"] = o.LinkId
	}
	if !IsNil(o.LinkAutoLink) {
		toSerialize["link_auto_link"] = o.LinkAutoLink
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableDeviceLinkEditInput struct {
	value *DeviceLinkEditInput
	isSet bool
}

func (v NullableDeviceLinkEditInput) Get() *DeviceLinkEditInput {
	return v.value
}

func (v *NullableDeviceLinkEditInput) Set(val *DeviceLinkEditInput) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceLinkEditInput) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceLinkEditInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceLinkEditInput(val *DeviceLinkEditInput) *NullableDeviceLinkEditInput {
	return &NullableDeviceLinkEditInput{value: val, isSet: true}
}

func (v NullableDeviceLinkEditInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceLinkEditInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
