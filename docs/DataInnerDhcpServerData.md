# DataInnerDhcpServerData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServerCiscoLogin** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerCiscoPassword** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerCiscoRootPassword** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerCiscoUseSsh** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ClusterHbHostaddr** | Pointer to **string** | The IP address of the heartbeat of the server, a direct link between the servers of the cluster. | [optional] 
**ClusterPeerServerId** | Pointer to **string** | The database identifier (ID) of the other server of the cluster. | [optional] 
**ServerClusterRole** | Pointer to **string** | The role of the server in the cluster, either &lt;b&gt;active (M)&lt;/b&gt;, &lt;b&gt;passive (B)&lt;/b&gt; or &lt;b&gt;N/A (#)&lt;/b&gt;. | [optional] 
**ClusterSshKeyringId** | Pointer to **string** | The database identifier (ID) of the SSH key dedicated to the cluster communication. | [optional] 
**ClusterVipPhysHostaddr** | Pointer to **string** | The local physical IP address of the VIP the cluster relies on. | [optional] 
**ConnectionprofileName** | Pointer to **string** | The name of the connection profile used as connection method for the DHCPv4 server. | [optional] 
**ServerClassName** | Pointer to **string** | The name of the class applied to the DHCPv4 server, it can be preceded by the class directory. | [optional] 
**ServerClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv4 server. | [optional] 
**ServerComment** | Pointer to **string** | The description of the DHCPv4 server. | [optional] 
**ServerId** | Pointer to **string** | The database identifier (ID) of the DHCPv4 server, a unique numeric key value automatically incremented when you add a DHCPv4 server. | [optional] 
**ServerLocaltime** | Pointer to **string** | The local time on the DHCPv4 server, in decimal UNIX date format. | [optional] 
**ServerName** | Pointer to **string** | The name of the DHCPv4 server. | [optional] 
**ServerState** | Pointer to **string** | The status of the DHCPv4 server: &lt;table&gt;&lt;caption&gt;server_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ER&lt;/td&gt;&lt;td &gt;The license used in SOLIDserver is not compliant with the added server: the license is invalid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ES&lt;/td&gt;&lt;td &gt;The server configuration could not be parsed properly.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ET&lt;/td&gt;&lt;td &gt;The server does not answer anymore due to a scheduled configuration of the server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IS&lt;/td&gt;&lt;td &gt;There was a setting error during the server declaration. For instance, some settings were added to a server that does not support them or a smart architecture is not managing any physical server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IC&lt;/td&gt;&lt;td &gt;The SSL credentials are invalid&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IP&lt;/td&gt;&lt;td &gt;The account used to add the Microsoft Windows DHCP server does not have sufficient privileges to manage it.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;LS&lt;/td&gt;&lt;td &gt;The server configuration is not viable.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;N&lt;/td&gt;&lt;td &gt;The server does not have a status as it has not synchronized yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;Y&lt;/td&gt;&lt;td &gt;The server is operational.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerSynching** | Pointer to **string** | The synchronization status of the DHCPv4 server. &lt;b&gt;1&lt;/b&gt; indicates that the server is currently being synchronized. | [optional] 
**ServerType** | Pointer to **string** | The type of the DHCPv4 server: &lt;table&gt;&lt;caption&gt;server_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP or EfficientIP Package server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;msrpc&lt;/td&gt;&lt;td &gt;Microsoft Windows DHCP server&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCPv4 smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**ServerUboottime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerVersion** | Pointer to **string** | The version details of the DHCPv4 server. | [optional] 
**ServerHostaddr** | Pointer to **string** | The human readable version of the parameter &lt;b&gt;server_addr&lt;/b&gt; or &lt;b&gt;server_addr6&lt;/b&gt;. | [optional] 
**ServerAddr6** | Pointer to **string** | The Management IP address of the DHCPv4 server, the IPv6 address configured when adding the server, in hexadecimal format. | [optional] 
**ServerAddr** | Pointer to **string** | The Management IP address of the DHCPv4 server, the IPv4 address configured when adding the server, in hexadecimal format. | [optional] 
**ServerHttpsLogin** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIsPackage** | Pointer to **string** | The DHCPv4 server package information. &lt;b&gt;Y&lt;/b&gt; for an EfficientIP Package server, &lt;b&gt;N&lt;/b&gt; for an appliance or virtual machine, &lt;b&gt;U&lt;/b&gt; the package information is irrelevant. For servers with a &lt;b&gt;server_type&lt;/b&gt; set to &lt;b&gt;ipm&lt;/b&gt;, &lt;b&gt;U&lt;/b&gt; indicates either EfficientIP Packages or appliances/virtual machines. | [optional] 
**ServerIpmdhcpProtocol** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerIsolated** | Pointer to **string** | A way to determine if the server can update any other module &lt;b&gt;(1)&lt;/b&gt;. | [optional] 
**ServerMsUseSsl** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ServerMsrpcDomain** | Pointer to **string** | The domain name of the Microsoft Windows DHCP server. | [optional] 
**ServerMsrpcLogin** | Pointer to **string** | The login of the Microsoft Windows DHCP server. | [optional] 
**ServerMultistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**Ref1ServerName** | Pointer to **string** | The name of the Master or Single DHCPv4 server within the smart architecture. | [optional] 
**Ref2ServerName** | Pointer to **string** | Internal use. Not documented. | [optional] 
**ReverseProxyConf** | Pointer to **string** | The URL of the HTTP(S) reverse proxy server that forwards client requests to the DHCPv4 server, if you configured one. | [optional] 
**ServiceAddressAddr** | Pointer to **string** | The Service IP address of the DHCPv4 server, the IPv4 address configured when adding the server. | [optional] 
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
**SmartArch** | Pointer to **string** | The type of the DHCPv4 smart architecture: &lt;table&gt;&lt;caption&gt;smart_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;masterslave&lt;/td&gt;&lt;td &gt;The One-to-One smart architecture sets a pair of DHCP servers in a Master/Backup configuration.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;star&lt;/td&gt;&lt;td &gt;The One-to-Many smart architecture sets a multi-site failover configuration at the cost of n-servers+1.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;splitscope&lt;/td&gt;&lt;td &gt;The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;The Single-Server smart architecture manages a single DHCP server.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
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

### NewDataInnerDhcpServerData

`func NewDataInnerDhcpServerData() *DataInnerDhcpServerData`

NewDataInnerDhcpServerData instantiates a new DataInnerDhcpServerData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataInnerDhcpServerDataWithDefaults

`func NewDataInnerDhcpServerDataWithDefaults() *DataInnerDhcpServerData`

NewDataInnerDhcpServerDataWithDefaults instantiates a new DataInnerDhcpServerData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServerCiscoLogin

`func (o *DataInnerDhcpServerData) GetServerCiscoLogin() string`

GetServerCiscoLogin returns the ServerCiscoLogin field if non-nil, zero value otherwise.

### GetServerCiscoLoginOk

`func (o *DataInnerDhcpServerData) GetServerCiscoLoginOk() (*string, bool)`

GetServerCiscoLoginOk returns a tuple with the ServerCiscoLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCiscoLogin

`func (o *DataInnerDhcpServerData) SetServerCiscoLogin(v string)`

SetServerCiscoLogin sets ServerCiscoLogin field to given value.

### HasServerCiscoLogin

`func (o *DataInnerDhcpServerData) HasServerCiscoLogin() bool`

HasServerCiscoLogin returns a boolean if a field has been set.

### GetServerCiscoPassword

`func (o *DataInnerDhcpServerData) GetServerCiscoPassword() string`

GetServerCiscoPassword returns the ServerCiscoPassword field if non-nil, zero value otherwise.

### GetServerCiscoPasswordOk

`func (o *DataInnerDhcpServerData) GetServerCiscoPasswordOk() (*string, bool)`

GetServerCiscoPasswordOk returns a tuple with the ServerCiscoPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCiscoPassword

`func (o *DataInnerDhcpServerData) SetServerCiscoPassword(v string)`

SetServerCiscoPassword sets ServerCiscoPassword field to given value.

### HasServerCiscoPassword

`func (o *DataInnerDhcpServerData) HasServerCiscoPassword() bool`

HasServerCiscoPassword returns a boolean if a field has been set.

### GetServerCiscoRootPassword

`func (o *DataInnerDhcpServerData) GetServerCiscoRootPassword() string`

GetServerCiscoRootPassword returns the ServerCiscoRootPassword field if non-nil, zero value otherwise.

### GetServerCiscoRootPasswordOk

`func (o *DataInnerDhcpServerData) GetServerCiscoRootPasswordOk() (*string, bool)`

GetServerCiscoRootPasswordOk returns a tuple with the ServerCiscoRootPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCiscoRootPassword

`func (o *DataInnerDhcpServerData) SetServerCiscoRootPassword(v string)`

SetServerCiscoRootPassword sets ServerCiscoRootPassword field to given value.

### HasServerCiscoRootPassword

`func (o *DataInnerDhcpServerData) HasServerCiscoRootPassword() bool`

HasServerCiscoRootPassword returns a boolean if a field has been set.

### GetServerCiscoUseSsh

`func (o *DataInnerDhcpServerData) GetServerCiscoUseSsh() string`

GetServerCiscoUseSsh returns the ServerCiscoUseSsh field if non-nil, zero value otherwise.

### GetServerCiscoUseSshOk

`func (o *DataInnerDhcpServerData) GetServerCiscoUseSshOk() (*string, bool)`

GetServerCiscoUseSshOk returns a tuple with the ServerCiscoUseSsh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerCiscoUseSsh

`func (o *DataInnerDhcpServerData) SetServerCiscoUseSsh(v string)`

SetServerCiscoUseSsh sets ServerCiscoUseSsh field to given value.

### HasServerCiscoUseSsh

`func (o *DataInnerDhcpServerData) HasServerCiscoUseSsh() bool`

HasServerCiscoUseSsh returns a boolean if a field has been set.

### GetClusterHbHostaddr

`func (o *DataInnerDhcpServerData) GetClusterHbHostaddr() string`

GetClusterHbHostaddr returns the ClusterHbHostaddr field if non-nil, zero value otherwise.

### GetClusterHbHostaddrOk

`func (o *DataInnerDhcpServerData) GetClusterHbHostaddrOk() (*string, bool)`

GetClusterHbHostaddrOk returns a tuple with the ClusterHbHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterHbHostaddr

`func (o *DataInnerDhcpServerData) SetClusterHbHostaddr(v string)`

SetClusterHbHostaddr sets ClusterHbHostaddr field to given value.

### HasClusterHbHostaddr

`func (o *DataInnerDhcpServerData) HasClusterHbHostaddr() bool`

HasClusterHbHostaddr returns a boolean if a field has been set.

### GetClusterPeerServerId

`func (o *DataInnerDhcpServerData) GetClusterPeerServerId() string`

GetClusterPeerServerId returns the ClusterPeerServerId field if non-nil, zero value otherwise.

### GetClusterPeerServerIdOk

`func (o *DataInnerDhcpServerData) GetClusterPeerServerIdOk() (*string, bool)`

GetClusterPeerServerIdOk returns a tuple with the ClusterPeerServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPeerServerId

`func (o *DataInnerDhcpServerData) SetClusterPeerServerId(v string)`

SetClusterPeerServerId sets ClusterPeerServerId field to given value.

### HasClusterPeerServerId

`func (o *DataInnerDhcpServerData) HasClusterPeerServerId() bool`

HasClusterPeerServerId returns a boolean if a field has been set.

### GetServerClusterRole

`func (o *DataInnerDhcpServerData) GetServerClusterRole() string`

GetServerClusterRole returns the ServerClusterRole field if non-nil, zero value otherwise.

### GetServerClusterRoleOk

`func (o *DataInnerDhcpServerData) GetServerClusterRoleOk() (*string, bool)`

GetServerClusterRoleOk returns a tuple with the ServerClusterRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClusterRole

`func (o *DataInnerDhcpServerData) SetServerClusterRole(v string)`

SetServerClusterRole sets ServerClusterRole field to given value.

### HasServerClusterRole

`func (o *DataInnerDhcpServerData) HasServerClusterRole() bool`

HasServerClusterRole returns a boolean if a field has been set.

### GetClusterSshKeyringId

`func (o *DataInnerDhcpServerData) GetClusterSshKeyringId() string`

GetClusterSshKeyringId returns the ClusterSshKeyringId field if non-nil, zero value otherwise.

### GetClusterSshKeyringIdOk

`func (o *DataInnerDhcpServerData) GetClusterSshKeyringIdOk() (*string, bool)`

GetClusterSshKeyringIdOk returns a tuple with the ClusterSshKeyringId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterSshKeyringId

`func (o *DataInnerDhcpServerData) SetClusterSshKeyringId(v string)`

SetClusterSshKeyringId sets ClusterSshKeyringId field to given value.

### HasClusterSshKeyringId

`func (o *DataInnerDhcpServerData) HasClusterSshKeyringId() bool`

HasClusterSshKeyringId returns a boolean if a field has been set.

### GetClusterVipPhysHostaddr

`func (o *DataInnerDhcpServerData) GetClusterVipPhysHostaddr() string`

GetClusterVipPhysHostaddr returns the ClusterVipPhysHostaddr field if non-nil, zero value otherwise.

### GetClusterVipPhysHostaddrOk

`func (o *DataInnerDhcpServerData) GetClusterVipPhysHostaddrOk() (*string, bool)`

GetClusterVipPhysHostaddrOk returns a tuple with the ClusterVipPhysHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterVipPhysHostaddr

`func (o *DataInnerDhcpServerData) SetClusterVipPhysHostaddr(v string)`

SetClusterVipPhysHostaddr sets ClusterVipPhysHostaddr field to given value.

### HasClusterVipPhysHostaddr

`func (o *DataInnerDhcpServerData) HasClusterVipPhysHostaddr() bool`

HasClusterVipPhysHostaddr returns a boolean if a field has been set.

### GetConnectionprofileName

`func (o *DataInnerDhcpServerData) GetConnectionprofileName() string`

GetConnectionprofileName returns the ConnectionprofileName field if non-nil, zero value otherwise.

### GetConnectionprofileNameOk

`func (o *DataInnerDhcpServerData) GetConnectionprofileNameOk() (*string, bool)`

GetConnectionprofileNameOk returns a tuple with the ConnectionprofileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionprofileName

`func (o *DataInnerDhcpServerData) SetConnectionprofileName(v string)`

SetConnectionprofileName sets ConnectionprofileName field to given value.

### HasConnectionprofileName

`func (o *DataInnerDhcpServerData) HasConnectionprofileName() bool`

HasConnectionprofileName returns a boolean if a field has been set.

### GetServerClassName

`func (o *DataInnerDhcpServerData) GetServerClassName() string`

GetServerClassName returns the ServerClassName field if non-nil, zero value otherwise.

### GetServerClassNameOk

`func (o *DataInnerDhcpServerData) GetServerClassNameOk() (*string, bool)`

GetServerClassNameOk returns a tuple with the ServerClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassName

`func (o *DataInnerDhcpServerData) SetServerClassName(v string)`

SetServerClassName sets ServerClassName field to given value.

### HasServerClassName

`func (o *DataInnerDhcpServerData) HasServerClassName() bool`

HasServerClassName returns a boolean if a field has been set.

### GetServerClassParameters

`func (o *DataInnerDhcpServerData) GetServerClassParameters() []ApiClassParameterOutputEntry`

GetServerClassParameters returns the ServerClassParameters field if non-nil, zero value otherwise.

### GetServerClassParametersOk

`func (o *DataInnerDhcpServerData) GetServerClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServerClassParametersOk returns a tuple with the ServerClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerClassParameters

`func (o *DataInnerDhcpServerData) SetServerClassParameters(v []ApiClassParameterOutputEntry)`

SetServerClassParameters sets ServerClassParameters field to given value.

### HasServerClassParameters

`func (o *DataInnerDhcpServerData) HasServerClassParameters() bool`

HasServerClassParameters returns a boolean if a field has been set.

### GetServerComment

`func (o *DataInnerDhcpServerData) GetServerComment() string`

GetServerComment returns the ServerComment field if non-nil, zero value otherwise.

### GetServerCommentOk

`func (o *DataInnerDhcpServerData) GetServerCommentOk() (*string, bool)`

GetServerCommentOk returns a tuple with the ServerComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerComment

`func (o *DataInnerDhcpServerData) SetServerComment(v string)`

SetServerComment sets ServerComment field to given value.

### HasServerComment

`func (o *DataInnerDhcpServerData) HasServerComment() bool`

HasServerComment returns a boolean if a field has been set.

### GetServerId

`func (o *DataInnerDhcpServerData) GetServerId() string`

GetServerId returns the ServerId field if non-nil, zero value otherwise.

### GetServerIdOk

`func (o *DataInnerDhcpServerData) GetServerIdOk() (*string, bool)`

GetServerIdOk returns a tuple with the ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerId

`func (o *DataInnerDhcpServerData) SetServerId(v string)`

SetServerId sets ServerId field to given value.

### HasServerId

`func (o *DataInnerDhcpServerData) HasServerId() bool`

HasServerId returns a boolean if a field has been set.

### GetServerLocaltime

`func (o *DataInnerDhcpServerData) GetServerLocaltime() string`

GetServerLocaltime returns the ServerLocaltime field if non-nil, zero value otherwise.

### GetServerLocaltimeOk

`func (o *DataInnerDhcpServerData) GetServerLocaltimeOk() (*string, bool)`

GetServerLocaltimeOk returns a tuple with the ServerLocaltime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerLocaltime

`func (o *DataInnerDhcpServerData) SetServerLocaltime(v string)`

SetServerLocaltime sets ServerLocaltime field to given value.

### HasServerLocaltime

`func (o *DataInnerDhcpServerData) HasServerLocaltime() bool`

HasServerLocaltime returns a boolean if a field has been set.

### GetServerName

`func (o *DataInnerDhcpServerData) GetServerName() string`

GetServerName returns the ServerName field if non-nil, zero value otherwise.

### GetServerNameOk

`func (o *DataInnerDhcpServerData) GetServerNameOk() (*string, bool)`

GetServerNameOk returns a tuple with the ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerName

`func (o *DataInnerDhcpServerData) SetServerName(v string)`

SetServerName sets ServerName field to given value.

### HasServerName

`func (o *DataInnerDhcpServerData) HasServerName() bool`

HasServerName returns a boolean if a field has been set.

### GetServerState

`func (o *DataInnerDhcpServerData) GetServerState() string`

GetServerState returns the ServerState field if non-nil, zero value otherwise.

### GetServerStateOk

`func (o *DataInnerDhcpServerData) GetServerStateOk() (*string, bool)`

GetServerStateOk returns a tuple with the ServerState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerState

`func (o *DataInnerDhcpServerData) SetServerState(v string)`

SetServerState sets ServerState field to given value.

### HasServerState

`func (o *DataInnerDhcpServerData) HasServerState() bool`

HasServerState returns a boolean if a field has been set.

### GetServerSynching

`func (o *DataInnerDhcpServerData) GetServerSynching() string`

GetServerSynching returns the ServerSynching field if non-nil, zero value otherwise.

### GetServerSynchingOk

`func (o *DataInnerDhcpServerData) GetServerSynchingOk() (*string, bool)`

GetServerSynchingOk returns a tuple with the ServerSynching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSynching

`func (o *DataInnerDhcpServerData) SetServerSynching(v string)`

SetServerSynching sets ServerSynching field to given value.

### HasServerSynching

`func (o *DataInnerDhcpServerData) HasServerSynching() bool`

HasServerSynching returns a boolean if a field has been set.

### GetServerType

`func (o *DataInnerDhcpServerData) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *DataInnerDhcpServerData) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *DataInnerDhcpServerData) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *DataInnerDhcpServerData) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetServerUboottime

`func (o *DataInnerDhcpServerData) GetServerUboottime() string`

GetServerUboottime returns the ServerUboottime field if non-nil, zero value otherwise.

### GetServerUboottimeOk

`func (o *DataInnerDhcpServerData) GetServerUboottimeOk() (*string, bool)`

GetServerUboottimeOk returns a tuple with the ServerUboottime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerUboottime

`func (o *DataInnerDhcpServerData) SetServerUboottime(v string)`

SetServerUboottime sets ServerUboottime field to given value.

### HasServerUboottime

`func (o *DataInnerDhcpServerData) HasServerUboottime() bool`

HasServerUboottime returns a boolean if a field has been set.

### GetServerVersion

`func (o *DataInnerDhcpServerData) GetServerVersion() string`

GetServerVersion returns the ServerVersion field if non-nil, zero value otherwise.

### GetServerVersionOk

`func (o *DataInnerDhcpServerData) GetServerVersionOk() (*string, bool)`

GetServerVersionOk returns a tuple with the ServerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerVersion

`func (o *DataInnerDhcpServerData) SetServerVersion(v string)`

SetServerVersion sets ServerVersion field to given value.

### HasServerVersion

`func (o *DataInnerDhcpServerData) HasServerVersion() bool`

HasServerVersion returns a boolean if a field has been set.

### GetServerHostaddr

`func (o *DataInnerDhcpServerData) GetServerHostaddr() string`

GetServerHostaddr returns the ServerHostaddr field if non-nil, zero value otherwise.

### GetServerHostaddrOk

`func (o *DataInnerDhcpServerData) GetServerHostaddrOk() (*string, bool)`

GetServerHostaddrOk returns a tuple with the ServerHostaddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHostaddr

`func (o *DataInnerDhcpServerData) SetServerHostaddr(v string)`

SetServerHostaddr sets ServerHostaddr field to given value.

### HasServerHostaddr

`func (o *DataInnerDhcpServerData) HasServerHostaddr() bool`

HasServerHostaddr returns a boolean if a field has been set.

### GetServerAddr6

`func (o *DataInnerDhcpServerData) GetServerAddr6() string`

GetServerAddr6 returns the ServerAddr6 field if non-nil, zero value otherwise.

### GetServerAddr6Ok

`func (o *DataInnerDhcpServerData) GetServerAddr6Ok() (*string, bool)`

GetServerAddr6Ok returns a tuple with the ServerAddr6 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr6

`func (o *DataInnerDhcpServerData) SetServerAddr6(v string)`

SetServerAddr6 sets ServerAddr6 field to given value.

### HasServerAddr6

`func (o *DataInnerDhcpServerData) HasServerAddr6() bool`

HasServerAddr6 returns a boolean if a field has been set.

### GetServerAddr

`func (o *DataInnerDhcpServerData) GetServerAddr() string`

GetServerAddr returns the ServerAddr field if non-nil, zero value otherwise.

### GetServerAddrOk

`func (o *DataInnerDhcpServerData) GetServerAddrOk() (*string, bool)`

GetServerAddrOk returns a tuple with the ServerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerAddr

`func (o *DataInnerDhcpServerData) SetServerAddr(v string)`

SetServerAddr sets ServerAddr field to given value.

### HasServerAddr

`func (o *DataInnerDhcpServerData) HasServerAddr() bool`

HasServerAddr returns a boolean if a field has been set.

### GetServerHttpsLogin

`func (o *DataInnerDhcpServerData) GetServerHttpsLogin() string`

GetServerHttpsLogin returns the ServerHttpsLogin field if non-nil, zero value otherwise.

### GetServerHttpsLoginOk

`func (o *DataInnerDhcpServerData) GetServerHttpsLoginOk() (*string, bool)`

GetServerHttpsLoginOk returns a tuple with the ServerHttpsLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerHttpsLogin

`func (o *DataInnerDhcpServerData) SetServerHttpsLogin(v string)`

SetServerHttpsLogin sets ServerHttpsLogin field to given value.

### HasServerHttpsLogin

`func (o *DataInnerDhcpServerData) HasServerHttpsLogin() bool`

HasServerHttpsLogin returns a boolean if a field has been set.

### GetServerIsPackage

`func (o *DataInnerDhcpServerData) GetServerIsPackage() string`

GetServerIsPackage returns the ServerIsPackage field if non-nil, zero value otherwise.

### GetServerIsPackageOk

`func (o *DataInnerDhcpServerData) GetServerIsPackageOk() (*string, bool)`

GetServerIsPackageOk returns a tuple with the ServerIsPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIsPackage

`func (o *DataInnerDhcpServerData) SetServerIsPackage(v string)`

SetServerIsPackage sets ServerIsPackage field to given value.

### HasServerIsPackage

`func (o *DataInnerDhcpServerData) HasServerIsPackage() bool`

HasServerIsPackage returns a boolean if a field has been set.

### GetServerIpmdhcpProtocol

`func (o *DataInnerDhcpServerData) GetServerIpmdhcpProtocol() string`

GetServerIpmdhcpProtocol returns the ServerIpmdhcpProtocol field if non-nil, zero value otherwise.

### GetServerIpmdhcpProtocolOk

`func (o *DataInnerDhcpServerData) GetServerIpmdhcpProtocolOk() (*string, bool)`

GetServerIpmdhcpProtocolOk returns a tuple with the ServerIpmdhcpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIpmdhcpProtocol

`func (o *DataInnerDhcpServerData) SetServerIpmdhcpProtocol(v string)`

SetServerIpmdhcpProtocol sets ServerIpmdhcpProtocol field to given value.

### HasServerIpmdhcpProtocol

`func (o *DataInnerDhcpServerData) HasServerIpmdhcpProtocol() bool`

HasServerIpmdhcpProtocol returns a boolean if a field has been set.

### GetServerIsolated

`func (o *DataInnerDhcpServerData) GetServerIsolated() string`

GetServerIsolated returns the ServerIsolated field if non-nil, zero value otherwise.

### GetServerIsolatedOk

`func (o *DataInnerDhcpServerData) GetServerIsolatedOk() (*string, bool)`

GetServerIsolatedOk returns a tuple with the ServerIsolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerIsolated

`func (o *DataInnerDhcpServerData) SetServerIsolated(v string)`

SetServerIsolated sets ServerIsolated field to given value.

### HasServerIsolated

`func (o *DataInnerDhcpServerData) HasServerIsolated() bool`

HasServerIsolated returns a boolean if a field has been set.

### GetServerMsUseSsl

`func (o *DataInnerDhcpServerData) GetServerMsUseSsl() string`

GetServerMsUseSsl returns the ServerMsUseSsl field if non-nil, zero value otherwise.

### GetServerMsUseSslOk

`func (o *DataInnerDhcpServerData) GetServerMsUseSslOk() (*string, bool)`

GetServerMsUseSslOk returns a tuple with the ServerMsUseSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMsUseSsl

`func (o *DataInnerDhcpServerData) SetServerMsUseSsl(v string)`

SetServerMsUseSsl sets ServerMsUseSsl field to given value.

### HasServerMsUseSsl

`func (o *DataInnerDhcpServerData) HasServerMsUseSsl() bool`

HasServerMsUseSsl returns a boolean if a field has been set.

### GetServerMsrpcDomain

`func (o *DataInnerDhcpServerData) GetServerMsrpcDomain() string`

GetServerMsrpcDomain returns the ServerMsrpcDomain field if non-nil, zero value otherwise.

### GetServerMsrpcDomainOk

`func (o *DataInnerDhcpServerData) GetServerMsrpcDomainOk() (*string, bool)`

GetServerMsrpcDomainOk returns a tuple with the ServerMsrpcDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMsrpcDomain

`func (o *DataInnerDhcpServerData) SetServerMsrpcDomain(v string)`

SetServerMsrpcDomain sets ServerMsrpcDomain field to given value.

### HasServerMsrpcDomain

`func (o *DataInnerDhcpServerData) HasServerMsrpcDomain() bool`

HasServerMsrpcDomain returns a boolean if a field has been set.

### GetServerMsrpcLogin

`func (o *DataInnerDhcpServerData) GetServerMsrpcLogin() string`

GetServerMsrpcLogin returns the ServerMsrpcLogin field if non-nil, zero value otherwise.

### GetServerMsrpcLoginOk

`func (o *DataInnerDhcpServerData) GetServerMsrpcLoginOk() (*string, bool)`

GetServerMsrpcLoginOk returns a tuple with the ServerMsrpcLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMsrpcLogin

`func (o *DataInnerDhcpServerData) SetServerMsrpcLogin(v string)`

SetServerMsrpcLogin sets ServerMsrpcLogin field to given value.

### HasServerMsrpcLogin

`func (o *DataInnerDhcpServerData) HasServerMsrpcLogin() bool`

HasServerMsrpcLogin returns a boolean if a field has been set.

### GetServerMultistatus

`func (o *DataInnerDhcpServerData) GetServerMultistatus() string`

GetServerMultistatus returns the ServerMultistatus field if non-nil, zero value otherwise.

### GetServerMultistatusOk

`func (o *DataInnerDhcpServerData) GetServerMultistatusOk() (*string, bool)`

GetServerMultistatusOk returns a tuple with the ServerMultistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerMultistatus

`func (o *DataInnerDhcpServerData) SetServerMultistatus(v string)`

SetServerMultistatus sets ServerMultistatus field to given value.

### HasServerMultistatus

`func (o *DataInnerDhcpServerData) HasServerMultistatus() bool`

HasServerMultistatus returns a boolean if a field has been set.

### GetRef1ServerName

`func (o *DataInnerDhcpServerData) GetRef1ServerName() string`

GetRef1ServerName returns the Ref1ServerName field if non-nil, zero value otherwise.

### GetRef1ServerNameOk

`func (o *DataInnerDhcpServerData) GetRef1ServerNameOk() (*string, bool)`

GetRef1ServerNameOk returns a tuple with the Ref1ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef1ServerName

`func (o *DataInnerDhcpServerData) SetRef1ServerName(v string)`

SetRef1ServerName sets Ref1ServerName field to given value.

### HasRef1ServerName

`func (o *DataInnerDhcpServerData) HasRef1ServerName() bool`

HasRef1ServerName returns a boolean if a field has been set.

### GetRef2ServerName

`func (o *DataInnerDhcpServerData) GetRef2ServerName() string`

GetRef2ServerName returns the Ref2ServerName field if non-nil, zero value otherwise.

### GetRef2ServerNameOk

`func (o *DataInnerDhcpServerData) GetRef2ServerNameOk() (*string, bool)`

GetRef2ServerNameOk returns a tuple with the Ref2ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef2ServerName

`func (o *DataInnerDhcpServerData) SetRef2ServerName(v string)`

SetRef2ServerName sets Ref2ServerName field to given value.

### HasRef2ServerName

`func (o *DataInnerDhcpServerData) HasRef2ServerName() bool`

HasRef2ServerName returns a boolean if a field has been set.

### GetReverseProxyConf

`func (o *DataInnerDhcpServerData) GetReverseProxyConf() string`

GetReverseProxyConf returns the ReverseProxyConf field if non-nil, zero value otherwise.

### GetReverseProxyConfOk

`func (o *DataInnerDhcpServerData) GetReverseProxyConfOk() (*string, bool)`

GetReverseProxyConfOk returns a tuple with the ReverseProxyConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseProxyConf

`func (o *DataInnerDhcpServerData) SetReverseProxyConf(v string)`

SetReverseProxyConf sets ReverseProxyConf field to given value.

### HasReverseProxyConf

`func (o *DataInnerDhcpServerData) HasReverseProxyConf() bool`

HasReverseProxyConf returns a boolean if a field has been set.

### GetServiceAddressAddr

`func (o *DataInnerDhcpServerData) GetServiceAddressAddr() string`

GetServiceAddressAddr returns the ServiceAddressAddr field if non-nil, zero value otherwise.

### GetServiceAddressAddrOk

`func (o *DataInnerDhcpServerData) GetServiceAddressAddrOk() (*string, bool)`

GetServiceAddressAddrOk returns a tuple with the ServiceAddressAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAddressAddr

`func (o *DataInnerDhcpServerData) SetServiceAddressAddr(v string)`

SetServiceAddressAddr sets ServiceAddressAddr field to given value.

### HasServiceAddressAddr

`func (o *DataInnerDhcpServerData) HasServiceAddressAddr() bool`

HasServiceAddressAddr returns a boolean if a field has been set.

### GetServerSnmpId

`func (o *DataInnerDhcpServerData) GetServerSnmpId() string`

GetServerSnmpId returns the ServerSnmpId field if non-nil, zero value otherwise.

### GetServerSnmpIdOk

`func (o *DataInnerDhcpServerData) GetServerSnmpIdOk() (*string, bool)`

GetServerSnmpIdOk returns a tuple with the ServerSnmpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpId

`func (o *DataInnerDhcpServerData) SetServerSnmpId(v string)`

SetServerSnmpId sets ServerSnmpId field to given value.

### HasServerSnmpId

`func (o *DataInnerDhcpServerData) HasServerSnmpId() bool`

HasServerSnmpId returns a boolean if a field has been set.

### GetServerSnmpPort

`func (o *DataInnerDhcpServerData) GetServerSnmpPort() string`

GetServerSnmpPort returns the ServerSnmpPort field if non-nil, zero value otherwise.

### GetServerSnmpPortOk

`func (o *DataInnerDhcpServerData) GetServerSnmpPortOk() (*string, bool)`

GetServerSnmpPortOk returns a tuple with the ServerSnmpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpPort

`func (o *DataInnerDhcpServerData) SetServerSnmpPort(v string)`

SetServerSnmpPort sets ServerSnmpPort field to given value.

### HasServerSnmpPort

`func (o *DataInnerDhcpServerData) HasServerSnmpPort() bool`

HasServerSnmpPort returns a boolean if a field has been set.

### GetServerSnmpProfileId

`func (o *DataInnerDhcpServerData) GetServerSnmpProfileId() string`

GetServerSnmpProfileId returns the ServerSnmpProfileId field if non-nil, zero value otherwise.

### GetServerSnmpProfileIdOk

`func (o *DataInnerDhcpServerData) GetServerSnmpProfileIdOk() (*string, bool)`

GetServerSnmpProfileIdOk returns a tuple with the ServerSnmpProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpProfileId

`func (o *DataInnerDhcpServerData) SetServerSnmpProfileId(v string)`

SetServerSnmpProfileId sets ServerSnmpProfileId field to given value.

### HasServerSnmpProfileId

`func (o *DataInnerDhcpServerData) HasServerSnmpProfileId() bool`

HasServerSnmpProfileId returns a boolean if a field has been set.

### GetServerSnmpRetry

`func (o *DataInnerDhcpServerData) GetServerSnmpRetry() string`

GetServerSnmpRetry returns the ServerSnmpRetry field if non-nil, zero value otherwise.

### GetServerSnmpRetryOk

`func (o *DataInnerDhcpServerData) GetServerSnmpRetryOk() (*string, bool)`

GetServerSnmpRetryOk returns a tuple with the ServerSnmpRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpRetry

`func (o *DataInnerDhcpServerData) SetServerSnmpRetry(v string)`

SetServerSnmpRetry sets ServerSnmpRetry field to given value.

### HasServerSnmpRetry

`func (o *DataInnerDhcpServerData) HasServerSnmpRetry() bool`

HasServerSnmpRetry returns a boolean if a field has been set.

### GetServerSnmpTimeout

`func (o *DataInnerDhcpServerData) GetServerSnmpTimeout() string`

GetServerSnmpTimeout returns the ServerSnmpTimeout field if non-nil, zero value otherwise.

### GetServerSnmpTimeoutOk

`func (o *DataInnerDhcpServerData) GetServerSnmpTimeoutOk() (*string, bool)`

GetServerSnmpTimeoutOk returns a tuple with the ServerSnmpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpTimeout

`func (o *DataInnerDhcpServerData) SetServerSnmpTimeout(v string)`

SetServerSnmpTimeout sets ServerSnmpTimeout field to given value.

### HasServerSnmpTimeout

`func (o *DataInnerDhcpServerData) HasServerSnmpTimeout() bool`

HasServerSnmpTimeout returns a boolean if a field has been set.

### GetServerSnmpUseTcp

`func (o *DataInnerDhcpServerData) GetServerSnmpUseTcp() string`

GetServerSnmpUseTcp returns the ServerSnmpUseTcp field if non-nil, zero value otherwise.

### GetServerSnmpUseTcpOk

`func (o *DataInnerDhcpServerData) GetServerSnmpUseTcpOk() (*string, bool)`

GetServerSnmpUseTcpOk returns a tuple with the ServerSnmpUseTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerSnmpUseTcp

`func (o *DataInnerDhcpServerData) SetServerSnmpUseTcp(v string)`

SetServerSnmpUseTcp sets ServerSnmpUseTcp field to given value.

### HasServerSnmpUseTcp

`func (o *DataInnerDhcpServerData) HasServerSnmpUseTcp() bool`

HasServerSnmpUseTcp returns a boolean if a field has been set.

### GetServerStatEnabled

`func (o *DataInnerDhcpServerData) GetServerStatEnabled() string`

GetServerStatEnabled returns the ServerStatEnabled field if non-nil, zero value otherwise.

### GetServerStatEnabledOk

`func (o *DataInnerDhcpServerData) GetServerStatEnabledOk() (*string, bool)`

GetServerStatEnabledOk returns a tuple with the ServerStatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatEnabled

`func (o *DataInnerDhcpServerData) SetServerStatEnabled(v string)`

SetServerStatEnabled sets ServerStatEnabled field to given value.

### HasServerStatEnabled

`func (o *DataInnerDhcpServerData) HasServerStatEnabled() bool`

HasServerStatEnabled returns a boolean if a field has been set.

### GetServerStatNiceness

`func (o *DataInnerDhcpServerData) GetServerStatNiceness() string`

GetServerStatNiceness returns the ServerStatNiceness field if non-nil, zero value otherwise.

### GetServerStatNicenessOk

`func (o *DataInnerDhcpServerData) GetServerStatNicenessOk() (*string, bool)`

GetServerStatNicenessOk returns a tuple with the ServerStatNiceness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatNiceness

`func (o *DataInnerDhcpServerData) SetServerStatNiceness(v string)`

SetServerStatNiceness sets ServerStatNiceness field to given value.

### HasServerStatNiceness

`func (o *DataInnerDhcpServerData) HasServerStatNiceness() bool`

HasServerStatNiceness returns a boolean if a field has been set.

### GetServerStatPeriod

`func (o *DataInnerDhcpServerData) GetServerStatPeriod() string`

GetServerStatPeriod returns the ServerStatPeriod field if non-nil, zero value otherwise.

### GetServerStatPeriodOk

`func (o *DataInnerDhcpServerData) GetServerStatPeriodOk() (*string, bool)`

GetServerStatPeriodOk returns a tuple with the ServerStatPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatPeriod

`func (o *DataInnerDhcpServerData) SetServerStatPeriod(v string)`

SetServerStatPeriod sets ServerStatPeriod field to given value.

### HasServerStatPeriod

`func (o *DataInnerDhcpServerData) HasServerStatPeriod() bool`

HasServerStatPeriod returns a boolean if a field has been set.

### GetServerStatTime

`func (o *DataInnerDhcpServerData) GetServerStatTime() string`

GetServerStatTime returns the ServerStatTime field if non-nil, zero value otherwise.

### GetServerStatTimeOk

`func (o *DataInnerDhcpServerData) GetServerStatTimeOk() (*string, bool)`

GetServerStatTimeOk returns a tuple with the ServerStatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerStatTime

`func (o *DataInnerDhcpServerData) SetServerStatTime(v string)`

SetServerStatTime sets ServerStatTime field to given value.

### HasServerStatTime

`func (o *DataInnerDhcpServerData) HasServerStatTime() bool`

HasServerStatTime returns a boolean if a field has been set.

### GetServerTcpPort

`func (o *DataInnerDhcpServerData) GetServerTcpPort() string`

GetServerTcpPort returns the ServerTcpPort field if non-nil, zero value otherwise.

### GetServerTcpPortOk

`func (o *DataInnerDhcpServerData) GetServerTcpPortOk() (*string, bool)`

GetServerTcpPortOk returns a tuple with the ServerTcpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerTcpPort

`func (o *DataInnerDhcpServerData) SetServerTcpPort(v string)`

SetServerTcpPort sets ServerTcpPort field to given value.

### HasServerTcpPort

`func (o *DataInnerDhcpServerData) HasServerTcpPort() bool`

HasServerTcpPort returns a boolean if a field has been set.

### GetTotalSmartMembers

`func (o *DataInnerDhcpServerData) GetTotalSmartMembers() string`

GetTotalSmartMembers returns the TotalSmartMembers field if non-nil, zero value otherwise.

### GetTotalSmartMembersOk

`func (o *DataInnerDhcpServerData) GetTotalSmartMembersOk() (*string, bool)`

GetTotalSmartMembersOk returns a tuple with the TotalSmartMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSmartMembers

`func (o *DataInnerDhcpServerData) SetTotalSmartMembers(v string)`

SetTotalSmartMembers sets TotalSmartMembers field to given value.

### HasTotalSmartMembers

`func (o *DataInnerDhcpServerData) HasTotalSmartMembers() bool`

HasTotalSmartMembers returns a boolean if a field has been set.

### GetSmartArch

`func (o *DataInnerDhcpServerData) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DataInnerDhcpServerData) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DataInnerDhcpServerData) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DataInnerDhcpServerData) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartMembersName

`func (o *DataInnerDhcpServerData) GetSmartMembersName() string`

GetSmartMembersName returns the SmartMembersName field if non-nil, zero value otherwise.

### GetSmartMembersNameOk

`func (o *DataInnerDhcpServerData) GetSmartMembersNameOk() (*string, bool)`

GetSmartMembersNameOk returns a tuple with the SmartMembersName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartMembersName

`func (o *DataInnerDhcpServerData) SetSmartMembersName(v string)`

SetSmartMembersName sets SmartMembersName field to given value.

### HasSmartMembersName

`func (o *DataInnerDhcpServerData) HasSmartMembersName() bool`

HasSmartMembersName returns a boolean if a field has been set.

### GetSmartParam1

`func (o *DataInnerDhcpServerData) GetSmartParam1() string`

GetSmartParam1 returns the SmartParam1 field if non-nil, zero value otherwise.

### GetSmartParam1Ok

`func (o *DataInnerDhcpServerData) GetSmartParam1Ok() (*string, bool)`

GetSmartParam1Ok returns a tuple with the SmartParam1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParam1

`func (o *DataInnerDhcpServerData) SetSmartParam1(v string)`

SetSmartParam1 sets SmartParam1 field to given value.

### HasSmartParam1

`func (o *DataInnerDhcpServerData) HasSmartParam1() bool`

HasSmartParam1 returns a boolean if a field has been set.

### GetSmartParentArch

`func (o *DataInnerDhcpServerData) GetSmartParentArch() string`

GetSmartParentArch returns the SmartParentArch field if non-nil, zero value otherwise.

### GetSmartParentArchOk

`func (o *DataInnerDhcpServerData) GetSmartParentArchOk() (*string, bool)`

GetSmartParentArchOk returns a tuple with the SmartParentArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentArch

`func (o *DataInnerDhcpServerData) SetSmartParentArch(v string)`

SetSmartParentArch sets SmartParentArch field to given value.

### HasSmartParentArch

`func (o *DataInnerDhcpServerData) HasSmartParentArch() bool`

HasSmartParentArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DataInnerDhcpServerData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DataInnerDhcpServerData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DataInnerDhcpServerData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DataInnerDhcpServerData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DataInnerDhcpServerData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DataInnerDhcpServerData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DataInnerDhcpServerData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DataInnerDhcpServerData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.

### GetSmartRef1ServerId

`func (o *DataInnerDhcpServerData) GetSmartRef1ServerId() string`

GetSmartRef1ServerId returns the SmartRef1ServerId field if non-nil, zero value otherwise.

### GetSmartRef1ServerIdOk

`func (o *DataInnerDhcpServerData) GetSmartRef1ServerIdOk() (*string, bool)`

GetSmartRef1ServerIdOk returns a tuple with the SmartRef1ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef1ServerId

`func (o *DataInnerDhcpServerData) SetSmartRef1ServerId(v string)`

SetSmartRef1ServerId sets SmartRef1ServerId field to given value.

### HasSmartRef1ServerId

`func (o *DataInnerDhcpServerData) HasSmartRef1ServerId() bool`

HasSmartRef1ServerId returns a boolean if a field has been set.

### GetSmartRef1ServerName

`func (o *DataInnerDhcpServerData) GetSmartRef1ServerName() string`

GetSmartRef1ServerName returns the SmartRef1ServerName field if non-nil, zero value otherwise.

### GetSmartRef1ServerNameOk

`func (o *DataInnerDhcpServerData) GetSmartRef1ServerNameOk() (*string, bool)`

GetSmartRef1ServerNameOk returns a tuple with the SmartRef1ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef1ServerName

`func (o *DataInnerDhcpServerData) SetSmartRef1ServerName(v string)`

SetSmartRef1ServerName sets SmartRef1ServerName field to given value.

### HasSmartRef1ServerName

`func (o *DataInnerDhcpServerData) HasSmartRef1ServerName() bool`

HasSmartRef1ServerName returns a boolean if a field has been set.

### GetSmartRef2ServerId

`func (o *DataInnerDhcpServerData) GetSmartRef2ServerId() string`

GetSmartRef2ServerId returns the SmartRef2ServerId field if non-nil, zero value otherwise.

### GetSmartRef2ServerIdOk

`func (o *DataInnerDhcpServerData) GetSmartRef2ServerIdOk() (*string, bool)`

GetSmartRef2ServerIdOk returns a tuple with the SmartRef2ServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef2ServerId

`func (o *DataInnerDhcpServerData) SetSmartRef2ServerId(v string)`

SetSmartRef2ServerId sets SmartRef2ServerId field to given value.

### HasSmartRef2ServerId

`func (o *DataInnerDhcpServerData) HasSmartRef2ServerId() bool`

HasSmartRef2ServerId returns a boolean if a field has been set.

### GetSmartRef2ServerName

`func (o *DataInnerDhcpServerData) GetSmartRef2ServerName() string`

GetSmartRef2ServerName returns the SmartRef2ServerName field if non-nil, zero value otherwise.

### GetSmartRef2ServerNameOk

`func (o *DataInnerDhcpServerData) GetSmartRef2ServerNameOk() (*string, bool)`

GetSmartRef2ServerNameOk returns a tuple with the SmartRef2ServerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef2ServerName

`func (o *DataInnerDhcpServerData) SetSmartRef2ServerName(v string)`

SetSmartRef2ServerName sets SmartRef2ServerName field to given value.

### HasSmartRef2ServerName

`func (o *DataInnerDhcpServerData) HasSmartRef2ServerName() bool`

HasSmartRef2ServerName returns a boolean if a field has been set.

### GetServerWindhcpProtocol

`func (o *DataInnerDhcpServerData) GetServerWindhcpProtocol() string`

GetServerWindhcpProtocol returns the ServerWindhcpProtocol field if non-nil, zero value otherwise.

### GetServerWindhcpProtocolOk

`func (o *DataInnerDhcpServerData) GetServerWindhcpProtocolOk() (*string, bool)`

GetServerWindhcpProtocolOk returns a tuple with the ServerWindhcpProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerWindhcpProtocol

`func (o *DataInnerDhcpServerData) SetServerWindhcpProtocol(v string)`

SetServerWindhcpProtocol sets ServerWindhcpProtocol field to given value.

### HasServerWindhcpProtocol

`func (o *DataInnerDhcpServerData) HasServerWindhcpProtocol() bool`

HasServerWindhcpProtocol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


