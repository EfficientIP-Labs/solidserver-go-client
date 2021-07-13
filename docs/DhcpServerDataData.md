# DhcpServerDataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerCiscoLogin** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerCiscoPassword** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerCiscoRootPassword** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerCiscoUseSsh** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | #general.output.class_parameters# | [optional] 
**ServerComment** | Pointer to **string** | The description of the DHCPv4 server. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. | [optional] 
**ServerLocaltime** | Pointer to **string** | The local time on the DHCPv4 server, in decimal UNIX date format. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server. | [optional] 
**ServerState** | Pointer to **string** | The status of the DHCPv4 server: &lt;table&gt;&lt;caption&gt;dhcp_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ER&lt;/td&gt;&lt;td &gt;The license used in SOLIDserver is not compliant with the added server: the license is invalid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ES&lt;/td&gt;&lt;td &gt;The server configuration could not be parsed properly.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ET&lt;/td&gt;&lt;td &gt;The server does not answer anymore due to a scheduled configuration of the server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IS&lt;/td&gt;&lt;td &gt;There was a setting error during the server declaration. For instance, some settings were added to a server that does not support them or a smart architecture is not managing any physical server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IC&lt;/td&gt;&lt;td &gt;The SSL credentials are invalid&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IP&lt;/td&gt;&lt;td &gt;The account used to add the Agentless DHCP server does not have sufficient privileges to manage it.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;LS&lt;/td&gt;&lt;td &gt;The server configuration is not viable.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;N&lt;/td&gt;&lt;td &gt;The server does not have a status as it has not synchronized yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;Y&lt;/td&gt;&lt;td &gt;The server is operational.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerSynching** | Pointer to **string** | The synchronization status of the DHCPv4 server. &lt;b&gt;1&lt;/b&gt; indicates that the server is currently being synchronized. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server: &lt;table&gt;&lt;caption&gt;dhcp_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;dcs&lt;/td&gt;&lt;td &gt;Nominum DCS server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerUboottime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server. | [optional] 
**ServerAddr** | Pointer to **string** | The IP address of the DHCP server, in hexadecimal format. | [optional] 
**ServerHttpsLogin** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIsPackage** | Pointer to **string** | The DHCP server package information. &lt;b&gt;Y&lt;/b&gt; for an EfficientIP DHCP Package server, &lt;b&gt;N&lt;/b&gt; for an appliance or virtual machine, &lt;b&gt;U&lt;/b&gt; the package information is irrelevant. For servers with a &lt;b&gt;dhcp_type&lt;/b&gt; set to &lt;b&gt;ipm&lt;/b&gt;, &lt;b&gt;U&lt;/b&gt; indicates either EfficientIP DHCP Packages or appliances/virtual machines. | [optional] 
**ServerIpmdhcpProtocol** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIsolated** | Pointer to **string** | A way to determine if the server can update any other module &lt;b&gt;(1)&lt;/b&gt;. | [optional] 
**ServerMsUseSsl** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerMsrpcDomain** | Pointer to **string** | The domain name of the Microsoft DHCP server. | [optional] 
**ServerMsrpcLogin** | Pointer to **string** | The login of the Microsoft DHCP server. | [optional] 
**ServerMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**Ref1ServerName** | Pointer to **string** | The name of the Master or Single DHCPv4 server within the smart architecture. | [optional] 
**Ref2ServerName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ReverseProxyConf** | Pointer to **string** | TODO:dhcp_server_list.output.reverse_proxy_conf | [optional] 
**ServerSnmpId** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerSnmpPort** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerSnmpProfileId** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerSnmpRetry** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerSnmpTimeout** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerSnmpUseTcp** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatEnabled** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatNiceness** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatPeriod** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerStatTime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerTcpPort** | Pointer to **string** | Internal use. Not documented. | [optional] 
**TotalSmartMembers** | Pointer to **string** | The total number of servers managed by the DHCPv4 smart architecture. | [optional] 
**SmartArch** | Pointer to **string** | The type of the DHCPv4 smart architecture: &lt;table&gt;&lt;caption&gt;vdhcp_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;masterslave&lt;/td&gt;&lt;td &gt;The One-to-One smart architecture sets a pair of DHCP servers in a Master/Backup configuration.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;star&lt;/td&gt;&lt;td &gt;The One-to-Many smart architecture sets a multi-site failover configuration at the cost of n-servers+1.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;splitscope&lt;/td&gt;&lt;td &gt;The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;The Single-Server smart architecture manages a single DHCP server.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SmartMembersName** | Pointer to **string** | The list of the servers managed by the DHCPv4 smart architecture, as follows: &lt;b&gt;&amp;lt;dhcp_name&amp;gt;,&amp;lt;dhcp_name&amp;gt;,...&lt;/b&gt; . | [optional] 
**SmartParam1** | Pointer to **string** | Internal use. Not documented. | [optional] 
**SmartParentArch** | Pointer to **string** | The type of the DHCPv4 smart architecture managing the DHCPv4 server. No value indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture managing the DHCPv4 server. &lt;b&gt;0&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv4 smart architecture managing the DHCPv4 server. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartRef1ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 smart architecture the server belongs to. | [optional] 
**SmartRef1ServerName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**SmartRef2ServerId** | Pointer to **string** | Internal use. Not documented. | [optional] 
**SmartRef2ServerName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerWindhcpProtocol** | Pointer to **string** | Internal use. Not documented. | [optional] 

## Methods

### NewDhcpServerDataData

`func NewDhcpServerDataData() *DhcpServerDataData`

NewDhcpServerDataData instantiates a new DhcpServerDataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpServerDataDataWithDefaults

`func NewDhcpServerDataDataWithDefaults() *DhcpServerDataData`

NewDhcpServerDataDataWithDefaults instantiates a new DhcpServerDataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerCiscoLogin

`func (o *DhcpServerDataData) GetServerCiscoLogin() string`

GetServerCiscoLogin returns the ServerCiscoLogin field if non-nil, zero value otherwise.

### GetServerCiscoLoginOk

`func (o *DhcpServerDataData) GetServerCiscoLoginOk() (*string, bool)`

GetServerCiscoLoginOk returns a tuple with the ServerCiscoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCiscoLogin

`func (o *DhcpServerDataData) SetServerCiscoLogin(v string)`

SetServerCiscoLogin sets ServerCiscoLogin field to given value.

### HasServerCiscoLogin

`func (o *DhcpServerDataData) HasServerCiscoLogin() bool`

HasServerCiscoLogin returns a boolean if a field has been set.

### GetServerCiscoPassword

`func (o *DhcpServerDataData) GetServerCiscoPassword() string`

GetServerCiscoPassword returns the ServerCiscoPassword field if non-nil, zero value otherwise.

### GetServerCiscoPasswordOk

`func (o *DhcpServerDataData) GetServerCiscoPasswordOk() (*string, bool)`

GetServerCiscoPasswordOk returns a tuple with the ServerCiscoPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCiscoPassword

`func (o *DhcpServerDataData) SetServerCiscoPassword(v string)`

SetServerCiscoPassword sets ServerCiscoPassword field to given value.

### HasServerCiscoPassword

`func (o *DhcpServerDataData) HasServerCiscoPassword() bool`

HasServerCiscoPassword returns a boolean if a field has been set.

### GetServerCiscoRootPassword

`func (o *DhcpServerDataData) GetServerCiscoRootPassword() string`

GetServerCiscoRootPassword returns the ServerCiscoRootPassword field if non-nil, zero value otherwise.

### GetServerCiscoRootPasswordOk

`func (o *DhcpServerDataData) GetServerCiscoRootPasswordOk() (*string, bool)`

GetServerCiscoRootPasswordOk returns a tuple with the ServerCiscoRootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCiscoRootPassword

`func (o *DhcpServerDataData) SetServerCiscoRootPassword(v string)`

SetServerCiscoRootPassword sets ServerCiscoRootPassword field to given value.

### HasServerCiscoRootPassword

`func (o *DhcpServerDataData) HasServerCiscoRootPassword() bool`

HasServerCiscoRootPassword returns a boolean if a field has been set.

### GetServerCiscoUseSsh

`func (o *DhcpServerDataData) GetServerCiscoUseSsh() string`

GetServerCiscoUseSsh returns the ServerCiscoUseSsh field if non-nil, zero value otherwise.

### GetServerCiscoUseSshOk

`func (o *DhcpServerDataData) GetServerCiscoUseSshOk() (*string, bool)`

GetServerCiscoUseSshOk returns a tuple with the ServerCiscoUseSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCiscoUseSsh

`func (o *DhcpServerDataData) SetServerCiscoUseSsh(v string)`

SetServerCiscoUseSsh sets ServerCiscoUseSsh field to given value.

### HasServerCiscoUseSsh

`func (o *DhcpServerDataData) HasServerCiscoUseSsh() bool`

HasServerCiscoUseSsh returns a boolean if a field has been set.

### GetServerClassName

`func (o *DhcpServerDataData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DhcpServerDataData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DhcpServerDataData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DhcpServerDataData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DhcpServerDataData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DhcpServerDataData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DhcpServerDataData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DhcpServerDataData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerComment

`func (o *DhcpServerDataData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DhcpServerDataData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DhcpServerDataData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DhcpServerDataData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerId

`func (o *DhcpServerDataData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DhcpServerDataData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DhcpServerDataData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DhcpServerDataData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerLocaltime

`func (o *DhcpServerDataData) GetServerLocaltime() string`

GetServerLocaltime returns the ServerLocaltime field if non-nil, zero value otherwise.

### GetServerLocaltimeOk

`func (o *DhcpServerDataData) GetServerLocaltimeOk() (*string, bool)`

GetServerLocaltimeOk returns a tuple with the ServerLocaltime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLocaltime

`func (o *DhcpServerDataData) SetServerLocaltime(v string)`

SetServerLocaltime sets ServerLocaltime field to given value.

### HasServerLocaltime

`func (o *DhcpServerDataData) HasServerLocaltime() bool`

HasServerLocaltime returns a boolean if a field has been set.

### GetServerName

`func (o *DhcpServerDataData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DhcpServerDataData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DhcpServerDataData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DhcpServerDataData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerState

`func (o *DhcpServerDataData) GetServerState() string`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *DhcpServerDataData) GetServerStateOk() (*string, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *DhcpServerDataData) SetServerState(v string)`

SetServerState sets ServerState field to given value.

### HasServerState

`func (o *DhcpServerDataData) HasServerState() bool`

HasServerState returns a boolean if a field has been set.

### GetServerSynching

`func (o *DhcpServerDataData) GetServerSynching() string`

GetServerSynching returns the ServerSynching field if non-nil, zero value otherwise.

### GetServerSynchingOk

`func (o *DhcpServerDataData) GetServerSynchingOk() (*string, bool)`

GetServerSynchingOk returns a tuple with the ServerSynching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSynching

`func (o *DhcpServerDataData) SetServerSynching(v string)`

SetServerSynching sets ServerSynching field to given value.

### HasServerSynching

`func (o *DhcpServerDataData) HasServerSynching() bool`

HasServerSynching returns a boolean if a field has been set.

### GetServerType

`func (o *DhcpServerDataData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DhcpServerDataData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DhcpServerDataData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DhcpServerDataData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerUboottime

`func (o *DhcpServerDataData) GetServerUboottime() string`

GetServerUboottime returns the ServerUboottime field if non-nil, zero value otherwise.

### GetServerUboottimeOk

`func (o *DhcpServerDataData) GetServerUboottimeOk() (*string, bool)`

GetServerUboottimeOk returns a tuple with the ServerUboottime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUboottime

`func (o *DhcpServerDataData) SetServerUboottime(v string)`

SetServerUboottime sets ServerUboottime field to given value.

### HasServerUboottime

`func (o *DhcpServerDataData) HasServerUboottime() bool`

HasServerUboottime returns a boolean if a field has been set.

### GetServerVersion

`func (o *DhcpServerDataData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DhcpServerDataData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DhcpServerDataData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DhcpServerDataData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerAddr

`func (o *DhcpServerDataData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DhcpServerDataData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DhcpServerDataData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DhcpServerDataData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetServerHttpsLogin

`func (o *DhcpServerDataData) GetServerHttpsLogin() string`

GetServerHttpsLogin returns the ServerHttpsLogin field if non-nil, zero value otherwise.

### GetServerHttpsLoginOk

`func (o *DhcpServerDataData) GetServerHttpsLoginOk() (*string, bool)`

GetServerHttpsLoginOk returns a tuple with the ServerHttpsLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHttpsLogin

`func (o *DhcpServerDataData) SetServerHttpsLogin(v string)`

SetServerHttpsLogin sets ServerHttpsLogin field to given value.

### HasServerHttpsLogin

`func (o *DhcpServerDataData) HasServerHttpsLogin() bool`

HasServerHttpsLogin returns a boolean if a field has been set.

### GetServerIsPackage

`func (o *DhcpServerDataData) GetServerIsPackage() string`

GetServerIsPackage returns the ServerIsPackage field if non-nil, zero value otherwise.

### GetServerIsPackageOk

`func (o *DhcpServerDataData) GetServerIsPackageOk() (*string, bool)`

GetServerIsPackageOk returns a tuple with the ServerIsPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIsPackage

`func (o *DhcpServerDataData) SetServerIsPackage(v string)`

SetServerIsPackage sets ServerIsPackage field to given value.

### HasServerIsPackage

`func (o *DhcpServerDataData) HasServerIsPackage() bool`

HasServerIsPackage returns a boolean if a field has been set.

### GetServerIpmdhcpProtocol

`func (o *DhcpServerDataData) GetServerIpmdhcpProtocol() string`

GetServerIpmdhcpProtocol returns the ServerIpmdhcpProtocol field if non-nil, zero value otherwise.

### GetServerIpmdhcpProtocolOk

`func (o *DhcpServerDataData) GetServerIpmdhcpProtocolOk() (*string, bool)`

GetServerIpmdhcpProtocolOk returns a tuple with the ServerIpmdhcpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmdhcpProtocol

`func (o *DhcpServerDataData) SetServerIpmdhcpProtocol(v string)`

SetServerIpmdhcpProtocol sets ServerIpmdhcpProtocol field to given value.

### HasServerIpmdhcpProtocol

`func (o *DhcpServerDataData) HasServerIpmdhcpProtocol() bool`

HasServerIpmdhcpProtocol returns a boolean if a field has been set.

### GetServerIsolated

`func (o *DhcpServerDataData) GetServerIsolated() string`

GetServerIsolated returns the ServerIsolated field if non-nil, zero value otherwise.

### GetServerIsolatedOk

`func (o *DhcpServerDataData) GetServerIsolatedOk() (*string, bool)`

GetServerIsolatedOk returns a tuple with the ServerIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIsolated

`func (o *DhcpServerDataData) SetServerIsolated(v string)`

SetServerIsolated sets ServerIsolated field to given value.

### HasServerIsolated

`func (o *DhcpServerDataData) HasServerIsolated() bool`

HasServerIsolated returns a boolean if a field has been set.

### GetServerMsUseSsl

`func (o *DhcpServerDataData) GetServerMsUseSsl() string`

GetServerMsUseSsl returns the ServerMsUseSsl field if non-nil, zero value otherwise.

### GetServerMsUseSslOk

`func (o *DhcpServerDataData) GetServerMsUseSslOk() (*string, bool)`

GetServerMsUseSslOk returns a tuple with the ServerMsUseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMsUseSsl

`func (o *DhcpServerDataData) SetServerMsUseSsl(v string)`

SetServerMsUseSsl sets ServerMsUseSsl field to given value.

### HasServerMsUseSsl

`func (o *DhcpServerDataData) HasServerMsUseSsl() bool`

HasServerMsUseSsl returns a boolean if a field has been set.

### GetServerMsrpcDomain

`func (o *DhcpServerDataData) GetServerMsrpcDomain() string`

GetServerMsrpcDomain returns the ServerMsrpcDomain field if non-nil, zero value otherwise.

### GetServerMsrpcDomainOk

`func (o *DhcpServerDataData) GetServerMsrpcDomainOk() (*string, bool)`

GetServerMsrpcDomainOk returns a tuple with the ServerMsrpcDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMsrpcDomain

`func (o *DhcpServerDataData) SetServerMsrpcDomain(v string)`

SetServerMsrpcDomain sets ServerMsrpcDomain field to given value.

### HasServerMsrpcDomain

`func (o *DhcpServerDataData) HasServerMsrpcDomain() bool`

HasServerMsrpcDomain returns a boolean if a field has been set.

### GetServerMsrpcLogin

`func (o *DhcpServerDataData) GetServerMsrpcLogin() string`

GetServerMsrpcLogin returns the ServerMsrpcLogin field if non-nil, zero value otherwise.

### GetServerMsrpcLoginOk

`func (o *DhcpServerDataData) GetServerMsrpcLoginOk() (*string, bool)`

GetServerMsrpcLoginOk returns a tuple with the ServerMsrpcLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMsrpcLogin

`func (o *DhcpServerDataData) SetServerMsrpcLogin(v string)`

SetServerMsrpcLogin sets ServerMsrpcLogin field to given value.

### HasServerMsrpcLogin

`func (o *DhcpServerDataData) HasServerMsrpcLogin() bool`

HasServerMsrpcLogin returns a boolean if a field has been set.

### GetServerMultistatus

`func (o *DhcpServerDataData) GetServerMultistatus() string`

GetServerMultistatus returns the ServerMultistatus field if non-nil, zero value otherwise.

### GetServerMultistatusOk

`func (o *DhcpServerDataData) GetServerMultistatusOk() (*string, bool)`

GetServerMultistatusOk returns a tuple with the ServerMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMultistatus

`func (o *DhcpServerDataData) SetServerMultistatus(v string)`

SetServerMultistatus sets ServerMultistatus field to given value.

### HasServerMultistatus

`func (o *DhcpServerDataData) HasServerMultistatus() bool`

HasServerMultistatus returns a boolean if a field has been set.

### GetRef1ServerName

`func (o *DhcpServerDataData) GetRef1ServerName() string`

GetRef1ServerName returns the Ref1ServerName field if non-nil, zero value otherwise.

### GetRef1ServerNameOk

`func (o *DhcpServerDataData) GetRef1ServerNameOk() (*string, bool)`

GetRef1ServerNameOk returns a tuple with the Ref1ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef1ServerName

`func (o *DhcpServerDataData) SetRef1ServerName(v string)`

SetRef1ServerName sets Ref1ServerName field to given value.

### HasRef1ServerName

`func (o *DhcpServerDataData) HasRef1ServerName() bool`

HasRef1ServerName returns a boolean if a field has been set.

### GetRef2ServerName

`func (o *DhcpServerDataData) GetRef2ServerName() string`

GetRef2ServerName returns the Ref2ServerName field if non-nil, zero value otherwise.

### GetRef2ServerNameOk

`func (o *DhcpServerDataData) GetRef2ServerNameOk() (*string, bool)`

GetRef2ServerNameOk returns a tuple with the Ref2ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef2ServerName

`func (o *DhcpServerDataData) SetRef2ServerName(v string)`

SetRef2ServerName sets Ref2ServerName field to given value.

### HasRef2ServerName

`func (o *DhcpServerDataData) HasRef2ServerName() bool`

HasRef2ServerName returns a boolean if a field has been set.

### GetReverseProxyConf

`func (o *DhcpServerDataData) GetReverseProxyConf() string`

GetReverseProxyConf returns the ReverseProxyConf field if non-nil, zero value otherwise.

### GetReverseProxyConfOk

`func (o *DhcpServerDataData) GetReverseProxyConfOk() (*string, bool)`

GetReverseProxyConfOk returns a tuple with the ReverseProxyConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseProxyConf

`func (o *DhcpServerDataData) SetReverseProxyConf(v string)`

SetReverseProxyConf sets ReverseProxyConf field to given value.

### HasReverseProxyConf

`func (o *DhcpServerDataData) HasReverseProxyConf() bool`

HasReverseProxyConf returns a boolean if a field has been set.

### GetServerSnmpId

`func (o *DhcpServerDataData) GetServerSnmpId() string`

GetServerSnmpId returns the ServerSnmpId field if non-nil, zero value otherwise.

### GetServerSnmpIdOk

`func (o *DhcpServerDataData) GetServerSnmpIdOk() (*string, bool)`

GetServerSnmpIdOk returns a tuple with the ServerSnmpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpId

`func (o *DhcpServerDataData) SetServerSnmpId(v string)`

SetServerSnmpId sets ServerSnmpId field to given value.

### HasServerSnmpId

`func (o *DhcpServerDataData) HasServerSnmpId() bool`

HasServerSnmpId returns a boolean if a field has been set.

### GetServerSnmpPort

`func (o *DhcpServerDataData) GetServerSnmpPort() string`

GetServerSnmpPort returns the ServerSnmpPort field if non-nil, zero value otherwise.

### GetServerSnmpPortOk

`func (o *DhcpServerDataData) GetServerSnmpPortOk() (*string, bool)`

GetServerSnmpPortOk returns a tuple with the ServerSnmpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpPort

`func (o *DhcpServerDataData) SetServerSnmpPort(v string)`

SetServerSnmpPort sets ServerSnmpPort field to given value.

### HasServerSnmpPort

`func (o *DhcpServerDataData) HasServerSnmpPort() bool`

HasServerSnmpPort returns a boolean if a field has been set.

### GetServerSnmpProfileId

`func (o *DhcpServerDataData) GetServerSnmpProfileId() string`

GetServerSnmpProfileId returns the ServerSnmpProfileId field if non-nil, zero value otherwise.

### GetServerSnmpProfileIdOk

`func (o *DhcpServerDataData) GetServerSnmpProfileIdOk() (*string, bool)`

GetServerSnmpProfileIdOk returns a tuple with the ServerSnmpProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpProfileId

`func (o *DhcpServerDataData) SetServerSnmpProfileId(v string)`

SetServerSnmpProfileId sets ServerSnmpProfileId field to given value.

### HasServerSnmpProfileId

`func (o *DhcpServerDataData) HasServerSnmpProfileId() bool`

HasServerSnmpProfileId returns a boolean if a field has been set.

### GetServerSnmpRetry

`func (o *DhcpServerDataData) GetServerSnmpRetry() string`

GetServerSnmpRetry returns the ServerSnmpRetry field if non-nil, zero value otherwise.

### GetServerSnmpRetryOk

`func (o *DhcpServerDataData) GetServerSnmpRetryOk() (*string, bool)`

GetServerSnmpRetryOk returns a tuple with the ServerSnmpRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpRetry

`func (o *DhcpServerDataData) SetServerSnmpRetry(v string)`

SetServerSnmpRetry sets ServerSnmpRetry field to given value.

### HasServerSnmpRetry

`func (o *DhcpServerDataData) HasServerSnmpRetry() bool`

HasServerSnmpRetry returns a boolean if a field has been set.

### GetServerSnmpTimeout

`func (o *DhcpServerDataData) GetServerSnmpTimeout() string`

GetServerSnmpTimeout returns the ServerSnmpTimeout field if non-nil, zero value otherwise.

### GetServerSnmpTimeoutOk

`func (o *DhcpServerDataData) GetServerSnmpTimeoutOk() (*string, bool)`

GetServerSnmpTimeoutOk returns a tuple with the ServerSnmpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpTimeout

`func (o *DhcpServerDataData) SetServerSnmpTimeout(v string)`

SetServerSnmpTimeout sets ServerSnmpTimeout field to given value.

### HasServerSnmpTimeout

`func (o *DhcpServerDataData) HasServerSnmpTimeout() bool`

HasServerSnmpTimeout returns a boolean if a field has been set.

### GetServerSnmpUseTcp

`func (o *DhcpServerDataData) GetServerSnmpUseTcp() string`

GetServerSnmpUseTcp returns the ServerSnmpUseTcp field if non-nil, zero value otherwise.

### GetServerSnmpUseTcpOk

`func (o *DhcpServerDataData) GetServerSnmpUseTcpOk() (*string, bool)`

GetServerSnmpUseTcpOk returns a tuple with the ServerSnmpUseTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpUseTcp

`func (o *DhcpServerDataData) SetServerSnmpUseTcp(v string)`

SetServerSnmpUseTcp sets ServerSnmpUseTcp field to given value.

### HasServerSnmpUseTcp

`func (o *DhcpServerDataData) HasServerSnmpUseTcp() bool`

HasServerSnmpUseTcp returns a boolean if a field has been set.

### GetServerStatEnabled

`func (o *DhcpServerDataData) GetServerStatEnabled() string`

GetServerStatEnabled returns the ServerStatEnabled field if non-nil, zero value otherwise.

### GetServerStatEnabledOk

`func (o *DhcpServerDataData) GetServerStatEnabledOk() (*string, bool)`

GetServerStatEnabledOk returns a tuple with the ServerStatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatEnabled

`func (o *DhcpServerDataData) SetServerStatEnabled(v string)`

SetServerStatEnabled sets ServerStatEnabled field to given value.

### HasServerStatEnabled

`func (o *DhcpServerDataData) HasServerStatEnabled() bool`

HasServerStatEnabled returns a boolean if a field has been set.

### GetServerStatNiceness

`func (o *DhcpServerDataData) GetServerStatNiceness() string`

GetServerStatNiceness returns the ServerStatNiceness field if non-nil, zero value otherwise.

### GetServerStatNicenessOk

`func (o *DhcpServerDataData) GetServerStatNicenessOk() (*string, bool)`

GetServerStatNicenessOk returns a tuple with the ServerStatNiceness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatNiceness

`func (o *DhcpServerDataData) SetServerStatNiceness(v string)`

SetServerStatNiceness sets ServerStatNiceness field to given value.

### HasServerStatNiceness

`func (o *DhcpServerDataData) HasServerStatNiceness() bool`

HasServerStatNiceness returns a boolean if a field has been set.

### GetServerStatPeriod

`func (o *DhcpServerDataData) GetServerStatPeriod() string`

GetServerStatPeriod returns the ServerStatPeriod field if non-nil, zero value otherwise.

### GetServerStatPeriodOk

`func (o *DhcpServerDataData) GetServerStatPeriodOk() (*string, bool)`

GetServerStatPeriodOk returns a tuple with the ServerStatPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatPeriod

`func (o *DhcpServerDataData) SetServerStatPeriod(v string)`

SetServerStatPeriod sets ServerStatPeriod field to given value.

### HasServerStatPeriod

`func (o *DhcpServerDataData) HasServerStatPeriod() bool`

HasServerStatPeriod returns a boolean if a field has been set.

### GetServerStatTime

`func (o *DhcpServerDataData) GetServerStatTime() string`

GetServerStatTime returns the ServerStatTime field if non-nil, zero value otherwise.

### GetServerStatTimeOk

`func (o *DhcpServerDataData) GetServerStatTimeOk() (*string, bool)`

GetServerStatTimeOk returns a tuple with the ServerStatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatTime

`func (o *DhcpServerDataData) SetServerStatTime(v string)`

SetServerStatTime sets ServerStatTime field to given value.

### HasServerStatTime

`func (o *DhcpServerDataData) HasServerStatTime() bool`

HasServerStatTime returns a boolean if a field has been set.

### GetServerTcpPort

`func (o *DhcpServerDataData) GetServerTcpPort() string`

GetServerTcpPort returns the ServerTcpPort field if non-nil, zero value otherwise.

### GetServerTcpPortOk

`func (o *DhcpServerDataData) GetServerTcpPortOk() (*string, bool)`

GetServerTcpPortOk returns a tuple with the ServerTcpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTcpPort

`func (o *DhcpServerDataData) SetServerTcpPort(v string)`

SetServerTcpPort sets ServerTcpPort field to given value.

### HasServerTcpPort

`func (o *DhcpServerDataData) HasServerTcpPort() bool`

HasServerTcpPort returns a boolean if a field has been set.

### GetTotalSmartMembers

`func (o *DhcpServerDataData) GetTotalSmartMembers() string`

GetTotalSmartMembers returns the TotalSmartMembers field if non-nil, zero value otherwise.

### GetTotalSmartMembersOk

`func (o *DhcpServerDataData) GetTotalSmartMembersOk() (*string, bool)`

GetTotalSmartMembersOk returns a tuple with the TotalSmartMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSmartMembers

`func (o *DhcpServerDataData) SetTotalSmartMembers(v string)`

SetTotalSmartMembers sets TotalSmartMembers field to given value.

### HasTotalSmartMembers

`func (o *DhcpServerDataData) HasTotalSmartMembers() bool`

HasTotalSmartMembers returns a boolean if a field has been set.

### GetSmartArch

`func (o *DhcpServerDataData) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DhcpServerDataData) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DhcpServerDataData) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DhcpServerDataData) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartMembersName

`func (o *DhcpServerDataData) GetSmartMembersName() string`

GetSmartMembersName returns the SmartMembersName field if non-nil, zero value otherwise.

### GetSmartMembersNameOk

`func (o *DhcpServerDataData) GetSmartMembersNameOk() (*string, bool)`

GetSmartMembersNameOk returns a tuple with the SmartMembersName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartMembersName

`func (o *DhcpServerDataData) SetSmartMembersName(v string)`

SetSmartMembersName sets SmartMembersName field to given value.

### HasSmartMembersName

`func (o *DhcpServerDataData) HasSmartMembersName() bool`

HasSmartMembersName returns a boolean if a field has been set.

### GetSmartParam1

`func (o *DhcpServerDataData) GetSmartParam1() string`

GetSmartParam1 returns the SmartParam1 field if non-nil, zero value otherwise.

### GetSmartParam1Ok

`func (o *DhcpServerDataData) GetSmartParam1Ok() (*string, bool)`

GetSmartParam1Ok returns a tuple with the SmartParam1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParam1

`func (o *DhcpServerDataData) SetSmartParam1(v string)`

SetSmartParam1 sets SmartParam1 field to given value.

### HasSmartParam1

`func (o *DhcpServerDataData) HasSmartParam1() bool`

HasSmartParam1 returns a boolean if a field has been set.

### GetSmartParentArch

`func (o *DhcpServerDataData) GetSmartParentArch() string`

GetSmartParentArch returns the SmartParentArch field if non-nil, zero value otherwise.

### GetSmartParentArchOk

`func (o *DhcpServerDataData) GetSmartParentArchOk() (*string, bool)`

GetSmartParentArchOk returns a tuple with the SmartParentArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentArch

`func (o *DhcpServerDataData) SetSmartParentArch(v string)`

SetSmartParentArch sets SmartParentArch field to given value.

### HasSmartParentArch

`func (o *DhcpServerDataData) HasSmartParentArch() bool`

HasSmartParentArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpServerDataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpServerDataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpServerDataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpServerDataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DhcpServerDataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DhcpServerDataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DhcpServerDataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DhcpServerDataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.

### GetSmartRef1ServerId

`func (o *DhcpServerDataData) GetSmartRef1ServerId() string`

GetSmartRef1ServerId returns the SmartRef1ServerId field if non-nil, zero value otherwise.

### GetSmartRef1ServerIdOk

`func (o *DhcpServerDataData) GetSmartRef1ServerIdOk() (*string, bool)`

GetSmartRef1ServerIdOk returns a tuple with the SmartRef1ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef1ServerId

`func (o *DhcpServerDataData) SetSmartRef1ServerId(v string)`

SetSmartRef1ServerId sets SmartRef1ServerId field to given value.

### HasSmartRef1ServerId

`func (o *DhcpServerDataData) HasSmartRef1ServerId() bool`

HasSmartRef1ServerId returns a boolean if a field has been set.

### GetSmartRef1ServerName

`func (o *DhcpServerDataData) GetSmartRef1ServerName() string`

GetSmartRef1ServerName returns the SmartRef1ServerName field if non-nil, zero value otherwise.

### GetSmartRef1ServerNameOk

`func (o *DhcpServerDataData) GetSmartRef1ServerNameOk() (*string, bool)`

GetSmartRef1ServerNameOk returns a tuple with the SmartRef1ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef1ServerName

`func (o *DhcpServerDataData) SetSmartRef1ServerName(v string)`

SetSmartRef1ServerName sets SmartRef1ServerName field to given value.

### HasSmartRef1ServerName

`func (o *DhcpServerDataData) HasSmartRef1ServerName() bool`

HasSmartRef1ServerName returns a boolean if a field has been set.

### GetSmartRef2ServerId

`func (o *DhcpServerDataData) GetSmartRef2ServerId() string`

GetSmartRef2ServerId returns the SmartRef2ServerId field if non-nil, zero value otherwise.

### GetSmartRef2ServerIdOk

`func (o *DhcpServerDataData) GetSmartRef2ServerIdOk() (*string, bool)`

GetSmartRef2ServerIdOk returns a tuple with the SmartRef2ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef2ServerId

`func (o *DhcpServerDataData) SetSmartRef2ServerId(v string)`

SetSmartRef2ServerId sets SmartRef2ServerId field to given value.

### HasSmartRef2ServerId

`func (o *DhcpServerDataData) HasSmartRef2ServerId() bool`

HasSmartRef2ServerId returns a boolean if a field has been set.

### GetSmartRef2ServerName

`func (o *DhcpServerDataData) GetSmartRef2ServerName() string`

GetSmartRef2ServerName returns the SmartRef2ServerName field if non-nil, zero value otherwise.

### GetSmartRef2ServerNameOk

`func (o *DhcpServerDataData) GetSmartRef2ServerNameOk() (*string, bool)`

GetSmartRef2ServerNameOk returns a tuple with the SmartRef2ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef2ServerName

`func (o *DhcpServerDataData) SetSmartRef2ServerName(v string)`

SetSmartRef2ServerName sets SmartRef2ServerName field to given value.

### HasSmartRef2ServerName

`func (o *DhcpServerDataData) HasSmartRef2ServerName() bool`

HasSmartRef2ServerName returns a boolean if a field has been set.

### GetServerWindhcpProtocol

`func (o *DhcpServerDataData) GetServerWindhcpProtocol() string`

GetServerWindhcpProtocol returns the ServerWindhcpProtocol field if non-nil, zero value otherwise.

### GetServerWindhcpProtocolOk

`func (o *DhcpServerDataData) GetServerWindhcpProtocolOk() (*string, bool)`

GetServerWindhcpProtocolOk returns a tuple with the ServerWindhcpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWindhcpProtocol

`func (o *DhcpServerDataData) SetServerWindhcpProtocol(v string)`

SetServerWindhcpProtocol sets ServerWindhcpProtocol field to given value.

### HasServerWindhcpProtocol

`func (o *DhcpServerDataData) HasServerWindhcpProtocol() bool`

HasServerWindhcpProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


