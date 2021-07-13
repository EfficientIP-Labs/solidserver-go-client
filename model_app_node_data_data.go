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

// AppNodeDataData struct for AppNodeDataData
type AppNodeDataData struct {
	// The administrative status of the node. <b>1</b> indicates the node is managed. <b>0</b> indicates that is unmanaged and ignored.
	AdminStatus *string `json:"admin_status,omitempty"`
	// The FQDN of the application the object belongs to.
	ApplicationFqdn *string `json:"application_fqdn,omitempty"`
	// The database identifier (ID) of the GSLB server associated with the application the object belongs to. It is only returned for deployed applications.
	GslbserverId *string `json:"gslbserver_id,omitempty"`
	// The name of the GSLB server associated with the application the object belongs to. It is only returned for deployed applications.
	GslbserverName *string `json:"gslbserver_name,omitempty"`
	// The database identifier (ID) of the application the object belongs to.
	ApplicationId *string `json:"application_id,omitempty"`
	// The name of the application the object belongs to.
	ApplicationName *string `json:"application_name,omitempty"`
	// The number of times, between <b>1</b> and <b>10</b>, before the parameter <b>appnode_status</b> is set to <b>1</b> (active) due to the health check result. By default, it is set to <b>3</b>.
	HealthcheckFailback *string `json:"healthcheck_failback,omitempty"`
	// The number of times, between <b>1</b> and <b>10</b>, before the parameter <b>appnode_status</b> is set to <b>0</b> (inactive) due to the health check result. By default, it is set to <b>3</b>.
	HealthcheckFailover *string `json:"healthcheck_failover,omitempty"`
	// The frequency to which the health check configured for the node is performed, in seconds.
	HealthcheckFreq *string `json:"healthcheck_freq,omitempty"`
	// The database identifier (ID) of the health check configured for the node.
	HealthcheckId *string `json:"healthcheck_id,omitempty"`
	// The type of health check configured for the node.
	HealthcheckName *string `json:"healthcheck_name,omitempty"`
	// The rest of the health check parameters configured, when relevant.
	HealthcheckParams *string `json:"healthcheck_params,omitempty"`
	// The number of seconds, between <b>1</b> and <b>10</b>, before the health check times out if the node is not responding.
	HealthcheckTimeout *string `json:"healthcheck_timeout,omitempty"`
	// The database identifier (ID) of the node.
	NodeId *string `json:"node_id,omitempty"`
	// The IPv6 address of the node.
	NodeAddress6Addr *string `json:"node_address6_addr,omitempty"`
	// The IPv4 address of the node.
	NodeAddressAddr *string `json:"node_address_addr,omitempty"`
	// The name of the node.
	NodeName *string `json:"node_name,omitempty"`
	// The operational status of the node:
	NodeStatus *string `json:"node_status,omitempty"`
	// The weight of the node, an integer between <b>0</b> and <b>255</b>. <b>0</b> indicates a backup node.
	NodeWeight *string `json:"node_weight,omitempty"`
	// The database identifier (ID) of the pool.
	PoolId *string `json:"pool_id,omitempty"`
	// The load-balancing mode of the pool the object belongs to, either <b>weighted</b>, <b>round-robin</b> or <b>latency</b>.
	PoolLbMode *string `json:"pool_lb_mode,omitempty"`
	// The name of the pool.
	PoolName *string `json:"pool_name,omitempty"`
	// The delay of creation/deletion status. <b>1</b> indicates that the object is not created/deleted yet.
	NodeDelayedTime *string `json:"node_delayed_time,omitempty"`
	// The type of DNS server associated with the application the object belongs to. It is only returned for deployed applications.
	GslbserverType *string `json:"gslbserver_type,omitempty"`
	// The database identifier (ID) of the application the object belongs to. It is only returned for deployed applications.
	ParentApplicationId *string `json:"parent_application_id,omitempty"`
	// The health check type, as displayed in the GUI.
	TranslatedHealthcheckName *string `json:"translated_healthcheck_name,omitempty"`
	// The load-balancing mode of the pool the object belongs to, as displayed in the GUI.
	TranslatedPoolLbMode *string `json:"translated_pool_lb_mode,omitempty"`
}

// NewAppNodeDataData instantiates a new AppNodeDataData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppNodeDataData() *AppNodeDataData {
	this := AppNodeDataData{}
	return &this
}

// NewAppNodeDataDataWithDefaults instantiates a new AppNodeDataData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppNodeDataDataWithDefaults() *AppNodeDataData {
	this := AppNodeDataData{}
	return &this
}

// GetAdminStatus returns the AdminStatus field value if set, zero value otherwise.
func (o *AppNodeDataData) GetAdminStatus() string {
	if o == nil || o.AdminStatus == nil {
		var ret string
		return ret
	}
	return *o.AdminStatus
}

// GetAdminStatusOk returns a tuple with the AdminStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetAdminStatusOk() (*string, bool) {
	if o == nil || o.AdminStatus == nil {
		return nil, false
	}
	return o.AdminStatus, true
}

// HasAdminStatus returns a boolean if a field has been set.
func (o *AppNodeDataData) HasAdminStatus() bool {
	if o != nil && o.AdminStatus != nil {
		return true
	}

	return false
}

// SetAdminStatus gets a reference to the given string and assigns it to the AdminStatus field.
func (o *AppNodeDataData) SetAdminStatus(v string) {
	o.AdminStatus = &v
}

// GetApplicationFqdn returns the ApplicationFqdn field value if set, zero value otherwise.
func (o *AppNodeDataData) GetApplicationFqdn() string {
	if o == nil || o.ApplicationFqdn == nil {
		var ret string
		return ret
	}
	return *o.ApplicationFqdn
}

// GetApplicationFqdnOk returns a tuple with the ApplicationFqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetApplicationFqdnOk() (*string, bool) {
	if o == nil || o.ApplicationFqdn == nil {
		return nil, false
	}
	return o.ApplicationFqdn, true
}

// HasApplicationFqdn returns a boolean if a field has been set.
func (o *AppNodeDataData) HasApplicationFqdn() bool {
	if o != nil && o.ApplicationFqdn != nil {
		return true
	}

	return false
}

// SetApplicationFqdn gets a reference to the given string and assigns it to the ApplicationFqdn field.
func (o *AppNodeDataData) SetApplicationFqdn(v string) {
	o.ApplicationFqdn = &v
}

// GetGslbserverId returns the GslbserverId field value if set, zero value otherwise.
func (o *AppNodeDataData) GetGslbserverId() string {
	if o == nil || o.GslbserverId == nil {
		var ret string
		return ret
	}
	return *o.GslbserverId
}

// GetGslbserverIdOk returns a tuple with the GslbserverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetGslbserverIdOk() (*string, bool) {
	if o == nil || o.GslbserverId == nil {
		return nil, false
	}
	return o.GslbserverId, true
}

// HasGslbserverId returns a boolean if a field has been set.
func (o *AppNodeDataData) HasGslbserverId() bool {
	if o != nil && o.GslbserverId != nil {
		return true
	}

	return false
}

// SetGslbserverId gets a reference to the given string and assigns it to the GslbserverId field.
func (o *AppNodeDataData) SetGslbserverId(v string) {
	o.GslbserverId = &v
}

// GetGslbserverName returns the GslbserverName field value if set, zero value otherwise.
func (o *AppNodeDataData) GetGslbserverName() string {
	if o == nil || o.GslbserverName == nil {
		var ret string
		return ret
	}
	return *o.GslbserverName
}

// GetGslbserverNameOk returns a tuple with the GslbserverName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetGslbserverNameOk() (*string, bool) {
	if o == nil || o.GslbserverName == nil {
		return nil, false
	}
	return o.GslbserverName, true
}

// HasGslbserverName returns a boolean if a field has been set.
func (o *AppNodeDataData) HasGslbserverName() bool {
	if o != nil && o.GslbserverName != nil {
		return true
	}

	return false
}

// SetGslbserverName gets a reference to the given string and assigns it to the GslbserverName field.
func (o *AppNodeDataData) SetGslbserverName(v string) {
	o.GslbserverName = &v
}

// GetApplicationId returns the ApplicationId field value if set, zero value otherwise.
func (o *AppNodeDataData) GetApplicationId() string {
	if o == nil || o.ApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ApplicationId
}

// GetApplicationIdOk returns a tuple with the ApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetApplicationIdOk() (*string, bool) {
	if o == nil || o.ApplicationId == nil {
		return nil, false
	}
	return o.ApplicationId, true
}

// HasApplicationId returns a boolean if a field has been set.
func (o *AppNodeDataData) HasApplicationId() bool {
	if o != nil && o.ApplicationId != nil {
		return true
	}

	return false
}

// SetApplicationId gets a reference to the given string and assigns it to the ApplicationId field.
func (o *AppNodeDataData) SetApplicationId(v string) {
	o.ApplicationId = &v
}

// GetApplicationName returns the ApplicationName field value if set, zero value otherwise.
func (o *AppNodeDataData) GetApplicationName() string {
	if o == nil || o.ApplicationName == nil {
		var ret string
		return ret
	}
	return *o.ApplicationName
}

// GetApplicationNameOk returns a tuple with the ApplicationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetApplicationNameOk() (*string, bool) {
	if o == nil || o.ApplicationName == nil {
		return nil, false
	}
	return o.ApplicationName, true
}

// HasApplicationName returns a boolean if a field has been set.
func (o *AppNodeDataData) HasApplicationName() bool {
	if o != nil && o.ApplicationName != nil {
		return true
	}

	return false
}

// SetApplicationName gets a reference to the given string and assigns it to the ApplicationName field.
func (o *AppNodeDataData) SetApplicationName(v string) {
	o.ApplicationName = &v
}

// GetHealthcheckFailback returns the HealthcheckFailback field value if set, zero value otherwise.
func (o *AppNodeDataData) GetHealthcheckFailback() string {
	if o == nil || o.HealthcheckFailback == nil {
		var ret string
		return ret
	}
	return *o.HealthcheckFailback
}

// GetHealthcheckFailbackOk returns a tuple with the HealthcheckFailback field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetHealthcheckFailbackOk() (*string, bool) {
	if o == nil || o.HealthcheckFailback == nil {
		return nil, false
	}
	return o.HealthcheckFailback, true
}

// HasHealthcheckFailback returns a boolean if a field has been set.
func (o *AppNodeDataData) HasHealthcheckFailback() bool {
	if o != nil && o.HealthcheckFailback != nil {
		return true
	}

	return false
}

// SetHealthcheckFailback gets a reference to the given string and assigns it to the HealthcheckFailback field.
func (o *AppNodeDataData) SetHealthcheckFailback(v string) {
	o.HealthcheckFailback = &v
}

// GetHealthcheckFailover returns the HealthcheckFailover field value if set, zero value otherwise.
func (o *AppNodeDataData) GetHealthcheckFailover() string {
	if o == nil || o.HealthcheckFailover == nil {
		var ret string
		return ret
	}
	return *o.HealthcheckFailover
}

// GetHealthcheckFailoverOk returns a tuple with the HealthcheckFailover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetHealthcheckFailoverOk() (*string, bool) {
	if o == nil || o.HealthcheckFailover == nil {
		return nil, false
	}
	return o.HealthcheckFailover, true
}

// HasHealthcheckFailover returns a boolean if a field has been set.
func (o *AppNodeDataData) HasHealthcheckFailover() bool {
	if o != nil && o.HealthcheckFailover != nil {
		return true
	}

	return false
}

// SetHealthcheckFailover gets a reference to the given string and assigns it to the HealthcheckFailover field.
func (o *AppNodeDataData) SetHealthcheckFailover(v string) {
	o.HealthcheckFailover = &v
}

// GetHealthcheckFreq returns the HealthcheckFreq field value if set, zero value otherwise.
func (o *AppNodeDataData) GetHealthcheckFreq() string {
	if o == nil || o.HealthcheckFreq == nil {
		var ret string
		return ret
	}
	return *o.HealthcheckFreq
}

// GetHealthcheckFreqOk returns a tuple with the HealthcheckFreq field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetHealthcheckFreqOk() (*string, bool) {
	if o == nil || o.HealthcheckFreq == nil {
		return nil, false
	}
	return o.HealthcheckFreq, true
}

// HasHealthcheckFreq returns a boolean if a field has been set.
func (o *AppNodeDataData) HasHealthcheckFreq() bool {
	if o != nil && o.HealthcheckFreq != nil {
		return true
	}

	return false
}

// SetHealthcheckFreq gets a reference to the given string and assigns it to the HealthcheckFreq field.
func (o *AppNodeDataData) SetHealthcheckFreq(v string) {
	o.HealthcheckFreq = &v
}

// GetHealthcheckId returns the HealthcheckId field value if set, zero value otherwise.
func (o *AppNodeDataData) GetHealthcheckId() string {
	if o == nil || o.HealthcheckId == nil {
		var ret string
		return ret
	}
	return *o.HealthcheckId
}

// GetHealthcheckIdOk returns a tuple with the HealthcheckId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetHealthcheckIdOk() (*string, bool) {
	if o == nil || o.HealthcheckId == nil {
		return nil, false
	}
	return o.HealthcheckId, true
}

// HasHealthcheckId returns a boolean if a field has been set.
func (o *AppNodeDataData) HasHealthcheckId() bool {
	if o != nil && o.HealthcheckId != nil {
		return true
	}

	return false
}

// SetHealthcheckId gets a reference to the given string and assigns it to the HealthcheckId field.
func (o *AppNodeDataData) SetHealthcheckId(v string) {
	o.HealthcheckId = &v
}

// GetHealthcheckName returns the HealthcheckName field value if set, zero value otherwise.
func (o *AppNodeDataData) GetHealthcheckName() string {
	if o == nil || o.HealthcheckName == nil {
		var ret string
		return ret
	}
	return *o.HealthcheckName
}

// GetHealthcheckNameOk returns a tuple with the HealthcheckName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetHealthcheckNameOk() (*string, bool) {
	if o == nil || o.HealthcheckName == nil {
		return nil, false
	}
	return o.HealthcheckName, true
}

// HasHealthcheckName returns a boolean if a field has been set.
func (o *AppNodeDataData) HasHealthcheckName() bool {
	if o != nil && o.HealthcheckName != nil {
		return true
	}

	return false
}

// SetHealthcheckName gets a reference to the given string and assigns it to the HealthcheckName field.
func (o *AppNodeDataData) SetHealthcheckName(v string) {
	o.HealthcheckName = &v
}

// GetHealthcheckParams returns the HealthcheckParams field value if set, zero value otherwise.
func (o *AppNodeDataData) GetHealthcheckParams() string {
	if o == nil || o.HealthcheckParams == nil {
		var ret string
		return ret
	}
	return *o.HealthcheckParams
}

// GetHealthcheckParamsOk returns a tuple with the HealthcheckParams field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetHealthcheckParamsOk() (*string, bool) {
	if o == nil || o.HealthcheckParams == nil {
		return nil, false
	}
	return o.HealthcheckParams, true
}

// HasHealthcheckParams returns a boolean if a field has been set.
func (o *AppNodeDataData) HasHealthcheckParams() bool {
	if o != nil && o.HealthcheckParams != nil {
		return true
	}

	return false
}

// SetHealthcheckParams gets a reference to the given string and assigns it to the HealthcheckParams field.
func (o *AppNodeDataData) SetHealthcheckParams(v string) {
	o.HealthcheckParams = &v
}

// GetHealthcheckTimeout returns the HealthcheckTimeout field value if set, zero value otherwise.
func (o *AppNodeDataData) GetHealthcheckTimeout() string {
	if o == nil || o.HealthcheckTimeout == nil {
		var ret string
		return ret
	}
	return *o.HealthcheckTimeout
}

// GetHealthcheckTimeoutOk returns a tuple with the HealthcheckTimeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetHealthcheckTimeoutOk() (*string, bool) {
	if o == nil || o.HealthcheckTimeout == nil {
		return nil, false
	}
	return o.HealthcheckTimeout, true
}

// HasHealthcheckTimeout returns a boolean if a field has been set.
func (o *AppNodeDataData) HasHealthcheckTimeout() bool {
	if o != nil && o.HealthcheckTimeout != nil {
		return true
	}

	return false
}

// SetHealthcheckTimeout gets a reference to the given string and assigns it to the HealthcheckTimeout field.
func (o *AppNodeDataData) SetHealthcheckTimeout(v string) {
	o.HealthcheckTimeout = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *AppNodeDataData) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *AppNodeDataData) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *AppNodeDataData) SetNodeId(v string) {
	o.NodeId = &v
}

// GetNodeAddress6Addr returns the NodeAddress6Addr field value if set, zero value otherwise.
func (o *AppNodeDataData) GetNodeAddress6Addr() string {
	if o == nil || o.NodeAddress6Addr == nil {
		var ret string
		return ret
	}
	return *o.NodeAddress6Addr
}

// GetNodeAddress6AddrOk returns a tuple with the NodeAddress6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetNodeAddress6AddrOk() (*string, bool) {
	if o == nil || o.NodeAddress6Addr == nil {
		return nil, false
	}
	return o.NodeAddress6Addr, true
}

// HasNodeAddress6Addr returns a boolean if a field has been set.
func (o *AppNodeDataData) HasNodeAddress6Addr() bool {
	if o != nil && o.NodeAddress6Addr != nil {
		return true
	}

	return false
}

// SetNodeAddress6Addr gets a reference to the given string and assigns it to the NodeAddress6Addr field.
func (o *AppNodeDataData) SetNodeAddress6Addr(v string) {
	o.NodeAddress6Addr = &v
}

// GetNodeAddressAddr returns the NodeAddressAddr field value if set, zero value otherwise.
func (o *AppNodeDataData) GetNodeAddressAddr() string {
	if o == nil || o.NodeAddressAddr == nil {
		var ret string
		return ret
	}
	return *o.NodeAddressAddr
}

// GetNodeAddressAddrOk returns a tuple with the NodeAddressAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetNodeAddressAddrOk() (*string, bool) {
	if o == nil || o.NodeAddressAddr == nil {
		return nil, false
	}
	return o.NodeAddressAddr, true
}

// HasNodeAddressAddr returns a boolean if a field has been set.
func (o *AppNodeDataData) HasNodeAddressAddr() bool {
	if o != nil && o.NodeAddressAddr != nil {
		return true
	}

	return false
}

// SetNodeAddressAddr gets a reference to the given string and assigns it to the NodeAddressAddr field.
func (o *AppNodeDataData) SetNodeAddressAddr(v string) {
	o.NodeAddressAddr = &v
}

// GetNodeName returns the NodeName field value if set, zero value otherwise.
func (o *AppNodeDataData) GetNodeName() string {
	if o == nil || o.NodeName == nil {
		var ret string
		return ret
	}
	return *o.NodeName
}

// GetNodeNameOk returns a tuple with the NodeName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetNodeNameOk() (*string, bool) {
	if o == nil || o.NodeName == nil {
		return nil, false
	}
	return o.NodeName, true
}

// HasNodeName returns a boolean if a field has been set.
func (o *AppNodeDataData) HasNodeName() bool {
	if o != nil && o.NodeName != nil {
		return true
	}

	return false
}

// SetNodeName gets a reference to the given string and assigns it to the NodeName field.
func (o *AppNodeDataData) SetNodeName(v string) {
	o.NodeName = &v
}

// GetNodeStatus returns the NodeStatus field value if set, zero value otherwise.
func (o *AppNodeDataData) GetNodeStatus() string {
	if o == nil || o.NodeStatus == nil {
		var ret string
		return ret
	}
	return *o.NodeStatus
}

// GetNodeStatusOk returns a tuple with the NodeStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetNodeStatusOk() (*string, bool) {
	if o == nil || o.NodeStatus == nil {
		return nil, false
	}
	return o.NodeStatus, true
}

// HasNodeStatus returns a boolean if a field has been set.
func (o *AppNodeDataData) HasNodeStatus() bool {
	if o != nil && o.NodeStatus != nil {
		return true
	}

	return false
}

// SetNodeStatus gets a reference to the given string and assigns it to the NodeStatus field.
func (o *AppNodeDataData) SetNodeStatus(v string) {
	o.NodeStatus = &v
}

// GetNodeWeight returns the NodeWeight field value if set, zero value otherwise.
func (o *AppNodeDataData) GetNodeWeight() string {
	if o == nil || o.NodeWeight == nil {
		var ret string
		return ret
	}
	return *o.NodeWeight
}

// GetNodeWeightOk returns a tuple with the NodeWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetNodeWeightOk() (*string, bool) {
	if o == nil || o.NodeWeight == nil {
		return nil, false
	}
	return o.NodeWeight, true
}

// HasNodeWeight returns a boolean if a field has been set.
func (o *AppNodeDataData) HasNodeWeight() bool {
	if o != nil && o.NodeWeight != nil {
		return true
	}

	return false
}

// SetNodeWeight gets a reference to the given string and assigns it to the NodeWeight field.
func (o *AppNodeDataData) SetNodeWeight(v string) {
	o.NodeWeight = &v
}

// GetPoolId returns the PoolId field value if set, zero value otherwise.
func (o *AppNodeDataData) GetPoolId() string {
	if o == nil || o.PoolId == nil {
		var ret string
		return ret
	}
	return *o.PoolId
}

// GetPoolIdOk returns a tuple with the PoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetPoolIdOk() (*string, bool) {
	if o == nil || o.PoolId == nil {
		return nil, false
	}
	return o.PoolId, true
}

// HasPoolId returns a boolean if a field has been set.
func (o *AppNodeDataData) HasPoolId() bool {
	if o != nil && o.PoolId != nil {
		return true
	}

	return false
}

// SetPoolId gets a reference to the given string and assigns it to the PoolId field.
func (o *AppNodeDataData) SetPoolId(v string) {
	o.PoolId = &v
}

// GetPoolLbMode returns the PoolLbMode field value if set, zero value otherwise.
func (o *AppNodeDataData) GetPoolLbMode() string {
	if o == nil || o.PoolLbMode == nil {
		var ret string
		return ret
	}
	return *o.PoolLbMode
}

// GetPoolLbModeOk returns a tuple with the PoolLbMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetPoolLbModeOk() (*string, bool) {
	if o == nil || o.PoolLbMode == nil {
		return nil, false
	}
	return o.PoolLbMode, true
}

// HasPoolLbMode returns a boolean if a field has been set.
func (o *AppNodeDataData) HasPoolLbMode() bool {
	if o != nil && o.PoolLbMode != nil {
		return true
	}

	return false
}

// SetPoolLbMode gets a reference to the given string and assigns it to the PoolLbMode field.
func (o *AppNodeDataData) SetPoolLbMode(v string) {
	o.PoolLbMode = &v
}

// GetPoolName returns the PoolName field value if set, zero value otherwise.
func (o *AppNodeDataData) GetPoolName() string {
	if o == nil || o.PoolName == nil {
		var ret string
		return ret
	}
	return *o.PoolName
}

// GetPoolNameOk returns a tuple with the PoolName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetPoolNameOk() (*string, bool) {
	if o == nil || o.PoolName == nil {
		return nil, false
	}
	return o.PoolName, true
}

// HasPoolName returns a boolean if a field has been set.
func (o *AppNodeDataData) HasPoolName() bool {
	if o != nil && o.PoolName != nil {
		return true
	}

	return false
}

// SetPoolName gets a reference to the given string and assigns it to the PoolName field.
func (o *AppNodeDataData) SetPoolName(v string) {
	o.PoolName = &v
}

// GetNodeDelayedTime returns the NodeDelayedTime field value if set, zero value otherwise.
func (o *AppNodeDataData) GetNodeDelayedTime() string {
	if o == nil || o.NodeDelayedTime == nil {
		var ret string
		return ret
	}
	return *o.NodeDelayedTime
}

// GetNodeDelayedTimeOk returns a tuple with the NodeDelayedTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetNodeDelayedTimeOk() (*string, bool) {
	if o == nil || o.NodeDelayedTime == nil {
		return nil, false
	}
	return o.NodeDelayedTime, true
}

// HasNodeDelayedTime returns a boolean if a field has been set.
func (o *AppNodeDataData) HasNodeDelayedTime() bool {
	if o != nil && o.NodeDelayedTime != nil {
		return true
	}

	return false
}

// SetNodeDelayedTime gets a reference to the given string and assigns it to the NodeDelayedTime field.
func (o *AppNodeDataData) SetNodeDelayedTime(v string) {
	o.NodeDelayedTime = &v
}

// GetGslbserverType returns the GslbserverType field value if set, zero value otherwise.
func (o *AppNodeDataData) GetGslbserverType() string {
	if o == nil || o.GslbserverType == nil {
		var ret string
		return ret
	}
	return *o.GslbserverType
}

// GetGslbserverTypeOk returns a tuple with the GslbserverType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetGslbserverTypeOk() (*string, bool) {
	if o == nil || o.GslbserverType == nil {
		return nil, false
	}
	return o.GslbserverType, true
}

// HasGslbserverType returns a boolean if a field has been set.
func (o *AppNodeDataData) HasGslbserverType() bool {
	if o != nil && o.GslbserverType != nil {
		return true
	}

	return false
}

// SetGslbserverType gets a reference to the given string and assigns it to the GslbserverType field.
func (o *AppNodeDataData) SetGslbserverType(v string) {
	o.GslbserverType = &v
}

// GetParentApplicationId returns the ParentApplicationId field value if set, zero value otherwise.
func (o *AppNodeDataData) GetParentApplicationId() string {
	if o == nil || o.ParentApplicationId == nil {
		var ret string
		return ret
	}
	return *o.ParentApplicationId
}

// GetParentApplicationIdOk returns a tuple with the ParentApplicationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetParentApplicationIdOk() (*string, bool) {
	if o == nil || o.ParentApplicationId == nil {
		return nil, false
	}
	return o.ParentApplicationId, true
}

// HasParentApplicationId returns a boolean if a field has been set.
func (o *AppNodeDataData) HasParentApplicationId() bool {
	if o != nil && o.ParentApplicationId != nil {
		return true
	}

	return false
}

// SetParentApplicationId gets a reference to the given string and assigns it to the ParentApplicationId field.
func (o *AppNodeDataData) SetParentApplicationId(v string) {
	o.ParentApplicationId = &v
}

// GetTranslatedHealthcheckName returns the TranslatedHealthcheckName field value if set, zero value otherwise.
func (o *AppNodeDataData) GetTranslatedHealthcheckName() string {
	if o == nil || o.TranslatedHealthcheckName == nil {
		var ret string
		return ret
	}
	return *o.TranslatedHealthcheckName
}

// GetTranslatedHealthcheckNameOk returns a tuple with the TranslatedHealthcheckName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetTranslatedHealthcheckNameOk() (*string, bool) {
	if o == nil || o.TranslatedHealthcheckName == nil {
		return nil, false
	}
	return o.TranslatedHealthcheckName, true
}

// HasTranslatedHealthcheckName returns a boolean if a field has been set.
func (o *AppNodeDataData) HasTranslatedHealthcheckName() bool {
	if o != nil && o.TranslatedHealthcheckName != nil {
		return true
	}

	return false
}

// SetTranslatedHealthcheckName gets a reference to the given string and assigns it to the TranslatedHealthcheckName field.
func (o *AppNodeDataData) SetTranslatedHealthcheckName(v string) {
	o.TranslatedHealthcheckName = &v
}

// GetTranslatedPoolLbMode returns the TranslatedPoolLbMode field value if set, zero value otherwise.
func (o *AppNodeDataData) GetTranslatedPoolLbMode() string {
	if o == nil || o.TranslatedPoolLbMode == nil {
		var ret string
		return ret
	}
	return *o.TranslatedPoolLbMode
}

// GetTranslatedPoolLbModeOk returns a tuple with the TranslatedPoolLbMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppNodeDataData) GetTranslatedPoolLbModeOk() (*string, bool) {
	if o == nil || o.TranslatedPoolLbMode == nil {
		return nil, false
	}
	return o.TranslatedPoolLbMode, true
}

// HasTranslatedPoolLbMode returns a boolean if a field has been set.
func (o *AppNodeDataData) HasTranslatedPoolLbMode() bool {
	if o != nil && o.TranslatedPoolLbMode != nil {
		return true
	}

	return false
}

// SetTranslatedPoolLbMode gets a reference to the given string and assigns it to the TranslatedPoolLbMode field.
func (o *AppNodeDataData) SetTranslatedPoolLbMode(v string) {
	o.TranslatedPoolLbMode = &v
}

func (o AppNodeDataData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AdminStatus != nil {
		toSerialize["admin_status"] = o.AdminStatus
	}
	if o.ApplicationFqdn != nil {
		toSerialize["application_fqdn"] = o.ApplicationFqdn
	}
	if o.GslbserverId != nil {
		toSerialize["gslbserver_id"] = o.GslbserverId
	}
	if o.GslbserverName != nil {
		toSerialize["gslbserver_name"] = o.GslbserverName
	}
	if o.ApplicationId != nil {
		toSerialize["application_id"] = o.ApplicationId
	}
	if o.ApplicationName != nil {
		toSerialize["application_name"] = o.ApplicationName
	}
	if o.HealthcheckFailback != nil {
		toSerialize["healthcheck_failback"] = o.HealthcheckFailback
	}
	if o.HealthcheckFailover != nil {
		toSerialize["healthcheck_failover"] = o.HealthcheckFailover
	}
	if o.HealthcheckFreq != nil {
		toSerialize["healthcheck_freq"] = o.HealthcheckFreq
	}
	if o.HealthcheckId != nil {
		toSerialize["healthcheck_id"] = o.HealthcheckId
	}
	if o.HealthcheckName != nil {
		toSerialize["healthcheck_name"] = o.HealthcheckName
	}
	if o.HealthcheckParams != nil {
		toSerialize["healthcheck_params"] = o.HealthcheckParams
	}
	if o.HealthcheckTimeout != nil {
		toSerialize["healthcheck_timeout"] = o.HealthcheckTimeout
	}
	if o.NodeId != nil {
		toSerialize["node_id"] = o.NodeId
	}
	if o.NodeAddress6Addr != nil {
		toSerialize["node_address6_addr"] = o.NodeAddress6Addr
	}
	if o.NodeAddressAddr != nil {
		toSerialize["node_address_addr"] = o.NodeAddressAddr
	}
	if o.NodeName != nil {
		toSerialize["node_name"] = o.NodeName
	}
	if o.NodeStatus != nil {
		toSerialize["node_status"] = o.NodeStatus
	}
	if o.NodeWeight != nil {
		toSerialize["node_weight"] = o.NodeWeight
	}
	if o.PoolId != nil {
		toSerialize["pool_id"] = o.PoolId
	}
	if o.PoolLbMode != nil {
		toSerialize["pool_lb_mode"] = o.PoolLbMode
	}
	if o.PoolName != nil {
		toSerialize["pool_name"] = o.PoolName
	}
	if o.NodeDelayedTime != nil {
		toSerialize["node_delayed_time"] = o.NodeDelayedTime
	}
	if o.GslbserverType != nil {
		toSerialize["gslbserver_type"] = o.GslbserverType
	}
	if o.ParentApplicationId != nil {
		toSerialize["parent_application_id"] = o.ParentApplicationId
	}
	if o.TranslatedHealthcheckName != nil {
		toSerialize["translated_healthcheck_name"] = o.TranslatedHealthcheckName
	}
	if o.TranslatedPoolLbMode != nil {
		toSerialize["translated_pool_lb_mode"] = o.TranslatedPoolLbMode
	}
	return json.Marshal(toSerialize)
}

type NullableAppNodeDataData struct {
	value *AppNodeDataData
	isSet bool
}

func (v NullableAppNodeDataData) Get() *AppNodeDataData {
	return v.value
}

func (v *NullableAppNodeDataData) Set(val *AppNodeDataData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppNodeDataData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppNodeDataData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppNodeDataData(val *AppNodeDataData) *NullableAppNodeDataData {
	return &NullableAppNodeDataData{value: val, isSet: true}
}

func (v NullableAppNodeDataData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppNodeDataData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


