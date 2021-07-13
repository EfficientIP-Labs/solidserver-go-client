# DhcpServer6DataData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server6ClassName** | Pointer to **string** | The name of the class applied to the DHCPv6 server, it can be preceded by the class directory. | [optional] 
**Server6ClassParameters** | Pointer to [**[]ApiClassParameterOutputEntry**](ApiClassParameterOutputEntry.md) | The class parameters applied to the DHCPv6 server and their value: &lt;b&gt;&amp;lt;class-parameter1&amp;gt;&#x3D;&amp;lt;value1&amp;gt;&amp;amp;&amp;lt;class-parameter2&amp;gt;&#x3D;&amp;lt;value2&amp;gt;&amp;amp;&lt;/b&gt;... . | [optional] 
**Server6Comment** | Pointer to **string** | The description of the DHCPv6 server. | [optional] 
**Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 server, a unique numeric key value automatically incremented when you add a DHCPv6 server. | [optional] 
**Server6LastRefreshTime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6Name** | Pointer to **string** | The name of the DHCPv6 server. | [optional] 
**Server6State** | Pointer to **string** | The status of the DHCPv6 server: &lt;table&gt;&lt;caption&gt;dhcp6_state possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Status&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ER&lt;/td&gt;&lt;td &gt;The license used in SOLIDserver is not compliant with the added server: the license is invalid.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ES&lt;/td&gt;&lt;td &gt;The server configuration could not be parsed properly.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;ET&lt;/td&gt;&lt;td &gt;The server does not answer anymore due to a scheduled configuration of the server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IS&lt;/td&gt;&lt;td &gt;There was a setting error during the server declaration. For instance, some settings were added to a server that does not support them or a smart architecture is not managing any physical server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IC&lt;/td&gt;&lt;td &gt;The SSL credentials are invalid&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;IP&lt;/td&gt;&lt;td &gt;The account used to add the Agentless DHCP server does not have sufficient privileges to manage it.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;LS&lt;/td&gt;&lt;td &gt;The server configuration is not viable.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;N&lt;/td&gt;&lt;td &gt;The server does not have a status as it has not synchronized yet.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;Y&lt;/td&gt;&lt;td &gt;The server is operational.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Synching** | Pointer to **string** | The synchronization status of the DHCPv6 server. &lt;b&gt;1&lt;/b&gt; indicates that the server is currently being synchronized. | [optional] 
**Server6Type** | Pointer to **string** | The type of the DHCPv6 server: &lt;table&gt;&lt;caption&gt;dhcp6_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;ipm&lt;/td&gt;&lt;td &gt;EfficientIP DHCP server or EfficientIP DHCP Package&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;vdhcp&lt;/td&gt;&lt;td &gt;EfficientIP DHCP smart architecture&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**Server6Uboottime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6Version** | Pointer to **string** | The version details of the DHCPv6 server. | [optional] 
**Server6Addr** | Pointer to **string** | The IP address of the DHCP server, in hexadecimal format. | [optional] 
**Server6HttpsLogin** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6IsPackage** | Pointer to **string** | The DHCPv6 server package information. &lt;b&gt;Y&lt;/b&gt; for an EfficientIP DHCP Package server, &lt;b&gt;N&lt;/b&gt; for an appliance or virtual machine, &lt;b&gt;U&lt;/b&gt; the package information is irrelevant. For servers with a &lt;b&gt;dhcp6_type&lt;/b&gt; set to &lt;b&gt;ipm&lt;/b&gt;, &lt;b&gt;U&lt;/b&gt; indicates either EfficientIP DHCP Packages or appliances/virtual machines. | [optional] 
**Server6Isolated** | Pointer to **string** | A way to determine if the server can update any other module &lt;b&gt;(1)&lt;/b&gt;. | [optional] 
**Server6Multistatus** | Pointer to **string** | The Multi-status information is displayed as follows: &lt;i&gt;&amp;lt;number-of-instances&amp;gt;@&amp;lt;message-number&amp;gt;@&amp;lt;multi-status-severity&amp;gt;@&amp;lt;module&amp;gt;&lt;/i&gt;. The different severity levels are:&lt;br&gt;&lt;b&gt;Multi-status severity levels&lt;/b&gt;    &lt;table border&#x3D;1&gt;        &lt;thead&gt;        &lt;tr &gt;            &lt;td&gt;&lt;b&gt;Message number&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Severity&lt;/b&gt;&lt;/td&gt;            &lt;td&gt;&lt;b&gt;Description&lt;/b&gt;&lt;/td&gt;        &lt;/tr&gt;        &lt;/thead&gt;        &lt;tbody&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;0 - 16&lt;/td&gt;            &lt;td&gt;Emergency&lt;/td&gt;            &lt;td&gt;The object configuration prevents the system from running properly. Action is required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;17 - 33&lt;/td&gt;            &lt;td&gt;Critical&lt;/td&gt;            &lt;td&gt;The object configuration is in critical conditions. Immediate action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;34 - 50&lt;/td&gt;            &lt;td&gt;Error&lt;/td&gt;            &lt;td&gt;The object configuration failed at some level. Action is recommended.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;51 - 66&lt;/td&gt;            &lt;td&gt;Warning&lt;/td&gt;            &lt;td&gt;The object configuration triggers error messages if no action is taken. Action to be taken at your discretion.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;67 - 83&lt;/td&gt;            &lt;td&gt;Notice&lt;/td&gt;            &lt;td&gt;The object configuration is normal but undergoing events that might trigger errors. No immediate action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;tr  valign&#x3D;middle&gt;            &lt;td&gt;84 - 100&lt;/td&gt;            &lt;td&gt;Informational&lt;/td&gt;            &lt;td&gt;The object configuration is normal, operational messages (might inform you about potential incompatibilities with other modules, etc). No action required.&lt;/td&gt;        &lt;/tr&gt;        &lt;/tbody&gt;&lt;/table&gt; | [optional] 
**ReverseProxyConf** | Pointer to **string** | TODO:dhcp6_server6_list.output.reverse_proxy_conf | [optional] 
**Server6SnmpId** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6SnmpPort** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6SnmpProfileId** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6SnmpRetry** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6SnmpTimeout** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6SnmpUseTcp** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6StatEnabled** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6StatNiceness** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6StatPeriod** | Pointer to **string** | Internal use. Not documented. | [optional] 
**Server6StatTime** | Pointer to **string** | Internal use. Not documented. | [optional] 
**TotalSmartMembers** | Pointer to **string** | The total number of servers managed by the DHCPv6 smart architecture. | [optional] 
**SmartArch** | Pointer to **string** | The type of the DHCPv6 smart architecture.&lt;table&gt;&lt;caption&gt;vdhcp6_arch possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Type&lt;/th&gt;&lt;th&gt;Description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;single&lt;/td&gt;&lt;td &gt;The Single-Server smart architecture manages a single DHCPv6 server.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;splitscope&lt;/td&gt;&lt;td &gt;The Split-Scope smart architecture sets a pair of DHCP servers in a configuration where the two scopes listen to the same subnet, but the range of addresses is divided.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;stateless&lt;/td&gt;&lt;td &gt;The Stateless smart architecture offers a limited number of options to the DHCP clients. The IP address is delivered thanks to the subnet gateway and it is impossible to create any ranges or statics or to retrieve any leases.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt; | [optional] 
**SmartMembersName** | Pointer to **string** | The list of the servers managed by the DHCPv6 smart architecture, as follows: &lt;b&gt;&amp;lt;dhcp6_name&amp;gt;,&amp;lt;dhcp6_name&amp;gt;,...&lt;/b&gt; . | [optional] 
**SmartParam1** | Pointer to **string** | Internal use. Not documented. | [optional] 
**SmartParentArch** | Pointer to **string** | The type of the DHCPv6 smart architecture managing the DHCPv6 server. No value indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentId** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture managing the DHCPv6 server. &lt;b&gt;0&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartParentName** | Pointer to **string** | The name of the DHCPv6 smart architecture managing the DHCPv6 server. &lt;b&gt;#&lt;/b&gt; indicates that the server is not managed by a smart architecture or is a smart architecture itself. | [optional] 
**SmartRef1Server6Id** | Pointer to **string** | The database identifier (ID) of the DHCPv6 smart architecture the server belongs to. | [optional] 
**SmartRef1Server6Name** | Pointer to **string** | Internal use. Not documented. | [optional] 
**SmartRef2Server6Id** | Pointer to **string** | Internal use. Not documented. | [optional] 
**SmartRef2Server6Name** | Pointer to **string** | Internal use. Not documented. | [optional] 

## Methods

### NewDhcpServer6DataData

`func NewDhcpServer6DataData() *DhcpServer6DataData`

NewDhcpServer6DataData instantiates a new DhcpServer6DataData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDhcpServer6DataDataWithDefaults

`func NewDhcpServer6DataDataWithDefaults() *DhcpServer6DataData`

NewDhcpServer6DataDataWithDefaults instantiates a new DhcpServer6DataData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer6ClassName

`func (o *DhcpServer6DataData) GetServer6ClassName() string`

GetServer6ClassName returns the Server6ClassName field if non-nil, zero value otherwise.

### GetServer6ClassNameOk

`func (o *DhcpServer6DataData) GetServer6ClassNameOk() (*string, bool)`

GetServer6ClassNameOk returns a tuple with the Server6ClassName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassName

`func (o *DhcpServer6DataData) SetServer6ClassName(v string)`

SetServer6ClassName sets Server6ClassName field to given value.

### HasServer6ClassName

`func (o *DhcpServer6DataData) HasServer6ClassName() bool`

HasServer6ClassName returns a boolean if a field has been set.

### GetServer6ClassParameters

`func (o *DhcpServer6DataData) GetServer6ClassParameters() []ApiClassParameterOutputEntry`

GetServer6ClassParameters returns the Server6ClassParameters field if non-nil, zero value otherwise.

### GetServer6ClassParametersOk

`func (o *DhcpServer6DataData) GetServer6ClassParametersOk() (*[]ApiClassParameterOutputEntry, bool)`

GetServer6ClassParametersOk returns a tuple with the Server6ClassParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6ClassParameters

`func (o *DhcpServer6DataData) SetServer6ClassParameters(v []ApiClassParameterOutputEntry)`

SetServer6ClassParameters sets Server6ClassParameters field to given value.

### HasServer6ClassParameters

`func (o *DhcpServer6DataData) HasServer6ClassParameters() bool`

HasServer6ClassParameters returns a boolean if a field has been set.

### GetServer6Comment

`func (o *DhcpServer6DataData) GetServer6Comment() string`

GetServer6Comment returns the Server6Comment field if non-nil, zero value otherwise.

### GetServer6CommentOk

`func (o *DhcpServer6DataData) GetServer6CommentOk() (*string, bool)`

GetServer6CommentOk returns a tuple with the Server6Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Comment

`func (o *DhcpServer6DataData) SetServer6Comment(v string)`

SetServer6Comment sets Server6Comment field to given value.

### HasServer6Comment

`func (o *DhcpServer6DataData) HasServer6Comment() bool`

HasServer6Comment returns a boolean if a field has been set.

### GetServer6Id

`func (o *DhcpServer6DataData) GetServer6Id() string`

GetServer6Id returns the Server6Id field if non-nil, zero value otherwise.

### GetServer6IdOk

`func (o *DhcpServer6DataData) GetServer6IdOk() (*string, bool)`

GetServer6IdOk returns a tuple with the Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Id

`func (o *DhcpServer6DataData) SetServer6Id(v string)`

SetServer6Id sets Server6Id field to given value.

### HasServer6Id

`func (o *DhcpServer6DataData) HasServer6Id() bool`

HasServer6Id returns a boolean if a field has been set.

### GetServer6LastRefreshTime

`func (o *DhcpServer6DataData) GetServer6LastRefreshTime() string`

GetServer6LastRefreshTime returns the Server6LastRefreshTime field if non-nil, zero value otherwise.

### GetServer6LastRefreshTimeOk

`func (o *DhcpServer6DataData) GetServer6LastRefreshTimeOk() (*string, bool)`

GetServer6LastRefreshTimeOk returns a tuple with the Server6LastRefreshTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6LastRefreshTime

`func (o *DhcpServer6DataData) SetServer6LastRefreshTime(v string)`

SetServer6LastRefreshTime sets Server6LastRefreshTime field to given value.

### HasServer6LastRefreshTime

`func (o *DhcpServer6DataData) HasServer6LastRefreshTime() bool`

HasServer6LastRefreshTime returns a boolean if a field has been set.

### GetServer6Name

`func (o *DhcpServer6DataData) GetServer6Name() string`

GetServer6Name returns the Server6Name field if non-nil, zero value otherwise.

### GetServer6NameOk

`func (o *DhcpServer6DataData) GetServer6NameOk() (*string, bool)`

GetServer6NameOk returns a tuple with the Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Name

`func (o *DhcpServer6DataData) SetServer6Name(v string)`

SetServer6Name sets Server6Name field to given value.

### HasServer6Name

`func (o *DhcpServer6DataData) HasServer6Name() bool`

HasServer6Name returns a boolean if a field has been set.

### GetServer6State

`func (o *DhcpServer6DataData) GetServer6State() string`

GetServer6State returns the Server6State field if non-nil, zero value otherwise.

### GetServer6StateOk

`func (o *DhcpServer6DataData) GetServer6StateOk() (*string, bool)`

GetServer6StateOk returns a tuple with the Server6State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6State

`func (o *DhcpServer6DataData) SetServer6State(v string)`

SetServer6State sets Server6State field to given value.

### HasServer6State

`func (o *DhcpServer6DataData) HasServer6State() bool`

HasServer6State returns a boolean if a field has been set.

### GetServer6Synching

`func (o *DhcpServer6DataData) GetServer6Synching() string`

GetServer6Synching returns the Server6Synching field if non-nil, zero value otherwise.

### GetServer6SynchingOk

`func (o *DhcpServer6DataData) GetServer6SynchingOk() (*string, bool)`

GetServer6SynchingOk returns a tuple with the Server6Synching field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Synching

`func (o *DhcpServer6DataData) SetServer6Synching(v string)`

SetServer6Synching sets Server6Synching field to given value.

### HasServer6Synching

`func (o *DhcpServer6DataData) HasServer6Synching() bool`

HasServer6Synching returns a boolean if a field has been set.

### GetServer6Type

`func (o *DhcpServer6DataData) GetServer6Type() string`

GetServer6Type returns the Server6Type field if non-nil, zero value otherwise.

### GetServer6TypeOk

`func (o *DhcpServer6DataData) GetServer6TypeOk() (*string, bool)`

GetServer6TypeOk returns a tuple with the Server6Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Type

`func (o *DhcpServer6DataData) SetServer6Type(v string)`

SetServer6Type sets Server6Type field to given value.

### HasServer6Type

`func (o *DhcpServer6DataData) HasServer6Type() bool`

HasServer6Type returns a boolean if a field has been set.

### GetServer6Uboottime

`func (o *DhcpServer6DataData) GetServer6Uboottime() string`

GetServer6Uboottime returns the Server6Uboottime field if non-nil, zero value otherwise.

### GetServer6UboottimeOk

`func (o *DhcpServer6DataData) GetServer6UboottimeOk() (*string, bool)`

GetServer6UboottimeOk returns a tuple with the Server6Uboottime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Uboottime

`func (o *DhcpServer6DataData) SetServer6Uboottime(v string)`

SetServer6Uboottime sets Server6Uboottime field to given value.

### HasServer6Uboottime

`func (o *DhcpServer6DataData) HasServer6Uboottime() bool`

HasServer6Uboottime returns a boolean if a field has been set.

### GetServer6Version

`func (o *DhcpServer6DataData) GetServer6Version() string`

GetServer6Version returns the Server6Version field if non-nil, zero value otherwise.

### GetServer6VersionOk

`func (o *DhcpServer6DataData) GetServer6VersionOk() (*string, bool)`

GetServer6VersionOk returns a tuple with the Server6Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Version

`func (o *DhcpServer6DataData) SetServer6Version(v string)`

SetServer6Version sets Server6Version field to given value.

### HasServer6Version

`func (o *DhcpServer6DataData) HasServer6Version() bool`

HasServer6Version returns a boolean if a field has been set.

### GetServer6Addr

`func (o *DhcpServer6DataData) GetServer6Addr() string`

GetServer6Addr returns the Server6Addr field if non-nil, zero value otherwise.

### GetServer6AddrOk

`func (o *DhcpServer6DataData) GetServer6AddrOk() (*string, bool)`

GetServer6AddrOk returns a tuple with the Server6Addr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Addr

`func (o *DhcpServer6DataData) SetServer6Addr(v string)`

SetServer6Addr sets Server6Addr field to given value.

### HasServer6Addr

`func (o *DhcpServer6DataData) HasServer6Addr() bool`

HasServer6Addr returns a boolean if a field has been set.

### GetServer6HttpsLogin

`func (o *DhcpServer6DataData) GetServer6HttpsLogin() string`

GetServer6HttpsLogin returns the Server6HttpsLogin field if non-nil, zero value otherwise.

### GetServer6HttpsLoginOk

`func (o *DhcpServer6DataData) GetServer6HttpsLoginOk() (*string, bool)`

GetServer6HttpsLoginOk returns a tuple with the Server6HttpsLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6HttpsLogin

`func (o *DhcpServer6DataData) SetServer6HttpsLogin(v string)`

SetServer6HttpsLogin sets Server6HttpsLogin field to given value.

### HasServer6HttpsLogin

`func (o *DhcpServer6DataData) HasServer6HttpsLogin() bool`

HasServer6HttpsLogin returns a boolean if a field has been set.

### GetServer6IsPackage

`func (o *DhcpServer6DataData) GetServer6IsPackage() string`

GetServer6IsPackage returns the Server6IsPackage field if non-nil, zero value otherwise.

### GetServer6IsPackageOk

`func (o *DhcpServer6DataData) GetServer6IsPackageOk() (*string, bool)`

GetServer6IsPackageOk returns a tuple with the Server6IsPackage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6IsPackage

`func (o *DhcpServer6DataData) SetServer6IsPackage(v string)`

SetServer6IsPackage sets Server6IsPackage field to given value.

### HasServer6IsPackage

`func (o *DhcpServer6DataData) HasServer6IsPackage() bool`

HasServer6IsPackage returns a boolean if a field has been set.

### GetServer6Isolated

`func (o *DhcpServer6DataData) GetServer6Isolated() string`

GetServer6Isolated returns the Server6Isolated field if non-nil, zero value otherwise.

### GetServer6IsolatedOk

`func (o *DhcpServer6DataData) GetServer6IsolatedOk() (*string, bool)`

GetServer6IsolatedOk returns a tuple with the Server6Isolated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Isolated

`func (o *DhcpServer6DataData) SetServer6Isolated(v string)`

SetServer6Isolated sets Server6Isolated field to given value.

### HasServer6Isolated

`func (o *DhcpServer6DataData) HasServer6Isolated() bool`

HasServer6Isolated returns a boolean if a field has been set.

### GetServer6Multistatus

`func (o *DhcpServer6DataData) GetServer6Multistatus() string`

GetServer6Multistatus returns the Server6Multistatus field if non-nil, zero value otherwise.

### GetServer6MultistatusOk

`func (o *DhcpServer6DataData) GetServer6MultistatusOk() (*string, bool)`

GetServer6MultistatusOk returns a tuple with the Server6Multistatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6Multistatus

`func (o *DhcpServer6DataData) SetServer6Multistatus(v string)`

SetServer6Multistatus sets Server6Multistatus field to given value.

### HasServer6Multistatus

`func (o *DhcpServer6DataData) HasServer6Multistatus() bool`

HasServer6Multistatus returns a boolean if a field has been set.

### GetReverseProxyConf

`func (o *DhcpServer6DataData) GetReverseProxyConf() string`

GetReverseProxyConf returns the ReverseProxyConf field if non-nil, zero value otherwise.

### GetReverseProxyConfOk

`func (o *DhcpServer6DataData) GetReverseProxyConfOk() (*string, bool)`

GetReverseProxyConfOk returns a tuple with the ReverseProxyConf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReverseProxyConf

`func (o *DhcpServer6DataData) SetReverseProxyConf(v string)`

SetReverseProxyConf sets ReverseProxyConf field to given value.

### HasReverseProxyConf

`func (o *DhcpServer6DataData) HasReverseProxyConf() bool`

HasReverseProxyConf returns a boolean if a field has been set.

### GetServer6SnmpId

`func (o *DhcpServer6DataData) GetServer6SnmpId() string`

GetServer6SnmpId returns the Server6SnmpId field if non-nil, zero value otherwise.

### GetServer6SnmpIdOk

`func (o *DhcpServer6DataData) GetServer6SnmpIdOk() (*string, bool)`

GetServer6SnmpIdOk returns a tuple with the Server6SnmpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6SnmpId

`func (o *DhcpServer6DataData) SetServer6SnmpId(v string)`

SetServer6SnmpId sets Server6SnmpId field to given value.

### HasServer6SnmpId

`func (o *DhcpServer6DataData) HasServer6SnmpId() bool`

HasServer6SnmpId returns a boolean if a field has been set.

### GetServer6SnmpPort

`func (o *DhcpServer6DataData) GetServer6SnmpPort() string`

GetServer6SnmpPort returns the Server6SnmpPort field if non-nil, zero value otherwise.

### GetServer6SnmpPortOk

`func (o *DhcpServer6DataData) GetServer6SnmpPortOk() (*string, bool)`

GetServer6SnmpPortOk returns a tuple with the Server6SnmpPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6SnmpPort

`func (o *DhcpServer6DataData) SetServer6SnmpPort(v string)`

SetServer6SnmpPort sets Server6SnmpPort field to given value.

### HasServer6SnmpPort

`func (o *DhcpServer6DataData) HasServer6SnmpPort() bool`

HasServer6SnmpPort returns a boolean if a field has been set.

### GetServer6SnmpProfileId

`func (o *DhcpServer6DataData) GetServer6SnmpProfileId() string`

GetServer6SnmpProfileId returns the Server6SnmpProfileId field if non-nil, zero value otherwise.

### GetServer6SnmpProfileIdOk

`func (o *DhcpServer6DataData) GetServer6SnmpProfileIdOk() (*string, bool)`

GetServer6SnmpProfileIdOk returns a tuple with the Server6SnmpProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6SnmpProfileId

`func (o *DhcpServer6DataData) SetServer6SnmpProfileId(v string)`

SetServer6SnmpProfileId sets Server6SnmpProfileId field to given value.

### HasServer6SnmpProfileId

`func (o *DhcpServer6DataData) HasServer6SnmpProfileId() bool`

HasServer6SnmpProfileId returns a boolean if a field has been set.

### GetServer6SnmpRetry

`func (o *DhcpServer6DataData) GetServer6SnmpRetry() string`

GetServer6SnmpRetry returns the Server6SnmpRetry field if non-nil, zero value otherwise.

### GetServer6SnmpRetryOk

`func (o *DhcpServer6DataData) GetServer6SnmpRetryOk() (*string, bool)`

GetServer6SnmpRetryOk returns a tuple with the Server6SnmpRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6SnmpRetry

`func (o *DhcpServer6DataData) SetServer6SnmpRetry(v string)`

SetServer6SnmpRetry sets Server6SnmpRetry field to given value.

### HasServer6SnmpRetry

`func (o *DhcpServer6DataData) HasServer6SnmpRetry() bool`

HasServer6SnmpRetry returns a boolean if a field has been set.

### GetServer6SnmpTimeout

`func (o *DhcpServer6DataData) GetServer6SnmpTimeout() string`

GetServer6SnmpTimeout returns the Server6SnmpTimeout field if non-nil, zero value otherwise.

### GetServer6SnmpTimeoutOk

`func (o *DhcpServer6DataData) GetServer6SnmpTimeoutOk() (*string, bool)`

GetServer6SnmpTimeoutOk returns a tuple with the Server6SnmpTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6SnmpTimeout

`func (o *DhcpServer6DataData) SetServer6SnmpTimeout(v string)`

SetServer6SnmpTimeout sets Server6SnmpTimeout field to given value.

### HasServer6SnmpTimeout

`func (o *DhcpServer6DataData) HasServer6SnmpTimeout() bool`

HasServer6SnmpTimeout returns a boolean if a field has been set.

### GetServer6SnmpUseTcp

`func (o *DhcpServer6DataData) GetServer6SnmpUseTcp() string`

GetServer6SnmpUseTcp returns the Server6SnmpUseTcp field if non-nil, zero value otherwise.

### GetServer6SnmpUseTcpOk

`func (o *DhcpServer6DataData) GetServer6SnmpUseTcpOk() (*string, bool)`

GetServer6SnmpUseTcpOk returns a tuple with the Server6SnmpUseTcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6SnmpUseTcp

`func (o *DhcpServer6DataData) SetServer6SnmpUseTcp(v string)`

SetServer6SnmpUseTcp sets Server6SnmpUseTcp field to given value.

### HasServer6SnmpUseTcp

`func (o *DhcpServer6DataData) HasServer6SnmpUseTcp() bool`

HasServer6SnmpUseTcp returns a boolean if a field has been set.

### GetServer6StatEnabled

`func (o *DhcpServer6DataData) GetServer6StatEnabled() string`

GetServer6StatEnabled returns the Server6StatEnabled field if non-nil, zero value otherwise.

### GetServer6StatEnabledOk

`func (o *DhcpServer6DataData) GetServer6StatEnabledOk() (*string, bool)`

GetServer6StatEnabledOk returns a tuple with the Server6StatEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6StatEnabled

`func (o *DhcpServer6DataData) SetServer6StatEnabled(v string)`

SetServer6StatEnabled sets Server6StatEnabled field to given value.

### HasServer6StatEnabled

`func (o *DhcpServer6DataData) HasServer6StatEnabled() bool`

HasServer6StatEnabled returns a boolean if a field has been set.

### GetServer6StatNiceness

`func (o *DhcpServer6DataData) GetServer6StatNiceness() string`

GetServer6StatNiceness returns the Server6StatNiceness field if non-nil, zero value otherwise.

### GetServer6StatNicenessOk

`func (o *DhcpServer6DataData) GetServer6StatNicenessOk() (*string, bool)`

GetServer6StatNicenessOk returns a tuple with the Server6StatNiceness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6StatNiceness

`func (o *DhcpServer6DataData) SetServer6StatNiceness(v string)`

SetServer6StatNiceness sets Server6StatNiceness field to given value.

### HasServer6StatNiceness

`func (o *DhcpServer6DataData) HasServer6StatNiceness() bool`

HasServer6StatNiceness returns a boolean if a field has been set.

### GetServer6StatPeriod

`func (o *DhcpServer6DataData) GetServer6StatPeriod() string`

GetServer6StatPeriod returns the Server6StatPeriod field if non-nil, zero value otherwise.

### GetServer6StatPeriodOk

`func (o *DhcpServer6DataData) GetServer6StatPeriodOk() (*string, bool)`

GetServer6StatPeriodOk returns a tuple with the Server6StatPeriod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6StatPeriod

`func (o *DhcpServer6DataData) SetServer6StatPeriod(v string)`

SetServer6StatPeriod sets Server6StatPeriod field to given value.

### HasServer6StatPeriod

`func (o *DhcpServer6DataData) HasServer6StatPeriod() bool`

HasServer6StatPeriod returns a boolean if a field has been set.

### GetServer6StatTime

`func (o *DhcpServer6DataData) GetServer6StatTime() string`

GetServer6StatTime returns the Server6StatTime field if non-nil, zero value otherwise.

### GetServer6StatTimeOk

`func (o *DhcpServer6DataData) GetServer6StatTimeOk() (*string, bool)`

GetServer6StatTimeOk returns a tuple with the Server6StatTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer6StatTime

`func (o *DhcpServer6DataData) SetServer6StatTime(v string)`

SetServer6StatTime sets Server6StatTime field to given value.

### HasServer6StatTime

`func (o *DhcpServer6DataData) HasServer6StatTime() bool`

HasServer6StatTime returns a boolean if a field has been set.

### GetTotalSmartMembers

`func (o *DhcpServer6DataData) GetTotalSmartMembers() string`

GetTotalSmartMembers returns the TotalSmartMembers field if non-nil, zero value otherwise.

### GetTotalSmartMembersOk

`func (o *DhcpServer6DataData) GetTotalSmartMembersOk() (*string, bool)`

GetTotalSmartMembersOk returns a tuple with the TotalSmartMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSmartMembers

`func (o *DhcpServer6DataData) SetTotalSmartMembers(v string)`

SetTotalSmartMembers sets TotalSmartMembers field to given value.

### HasTotalSmartMembers

`func (o *DhcpServer6DataData) HasTotalSmartMembers() bool`

HasTotalSmartMembers returns a boolean if a field has been set.

### GetSmartArch

`func (o *DhcpServer6DataData) GetSmartArch() string`

GetSmartArch returns the SmartArch field if non-nil, zero value otherwise.

### GetSmartArchOk

`func (o *DhcpServer6DataData) GetSmartArchOk() (*string, bool)`

GetSmartArchOk returns a tuple with the SmartArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartArch

`func (o *DhcpServer6DataData) SetSmartArch(v string)`

SetSmartArch sets SmartArch field to given value.

### HasSmartArch

`func (o *DhcpServer6DataData) HasSmartArch() bool`

HasSmartArch returns a boolean if a field has been set.

### GetSmartMembersName

`func (o *DhcpServer6DataData) GetSmartMembersName() string`

GetSmartMembersName returns the SmartMembersName field if non-nil, zero value otherwise.

### GetSmartMembersNameOk

`func (o *DhcpServer6DataData) GetSmartMembersNameOk() (*string, bool)`

GetSmartMembersNameOk returns a tuple with the SmartMembersName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartMembersName

`func (o *DhcpServer6DataData) SetSmartMembersName(v string)`

SetSmartMembersName sets SmartMembersName field to given value.

### HasSmartMembersName

`func (o *DhcpServer6DataData) HasSmartMembersName() bool`

HasSmartMembersName returns a boolean if a field has been set.

### GetSmartParam1

`func (o *DhcpServer6DataData) GetSmartParam1() string`

GetSmartParam1 returns the SmartParam1 field if non-nil, zero value otherwise.

### GetSmartParam1Ok

`func (o *DhcpServer6DataData) GetSmartParam1Ok() (*string, bool)`

GetSmartParam1Ok returns a tuple with the SmartParam1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParam1

`func (o *DhcpServer6DataData) SetSmartParam1(v string)`

SetSmartParam1 sets SmartParam1 field to given value.

### HasSmartParam1

`func (o *DhcpServer6DataData) HasSmartParam1() bool`

HasSmartParam1 returns a boolean if a field has been set.

### GetSmartParentArch

`func (o *DhcpServer6DataData) GetSmartParentArch() string`

GetSmartParentArch returns the SmartParentArch field if non-nil, zero value otherwise.

### GetSmartParentArchOk

`func (o *DhcpServer6DataData) GetSmartParentArchOk() (*string, bool)`

GetSmartParentArchOk returns a tuple with the SmartParentArch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentArch

`func (o *DhcpServer6DataData) SetSmartParentArch(v string)`

SetSmartParentArch sets SmartParentArch field to given value.

### HasSmartParentArch

`func (o *DhcpServer6DataData) HasSmartParentArch() bool`

HasSmartParentArch returns a boolean if a field has been set.

### GetSmartParentId

`func (o *DhcpServer6DataData) GetSmartParentId() string`

GetSmartParentId returns the SmartParentId field if non-nil, zero value otherwise.

### GetSmartParentIdOk

`func (o *DhcpServer6DataData) GetSmartParentIdOk() (*string, bool)`

GetSmartParentIdOk returns a tuple with the SmartParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentId

`func (o *DhcpServer6DataData) SetSmartParentId(v string)`

SetSmartParentId sets SmartParentId field to given value.

### HasSmartParentId

`func (o *DhcpServer6DataData) HasSmartParentId() bool`

HasSmartParentId returns a boolean if a field has been set.

### GetSmartParentName

`func (o *DhcpServer6DataData) GetSmartParentName() string`

GetSmartParentName returns the SmartParentName field if non-nil, zero value otherwise.

### GetSmartParentNameOk

`func (o *DhcpServer6DataData) GetSmartParentNameOk() (*string, bool)`

GetSmartParentNameOk returns a tuple with the SmartParentName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartParentName

`func (o *DhcpServer6DataData) SetSmartParentName(v string)`

SetSmartParentName sets SmartParentName field to given value.

### HasSmartParentName

`func (o *DhcpServer6DataData) HasSmartParentName() bool`

HasSmartParentName returns a boolean if a field has been set.

### GetSmartRef1Server6Id

`func (o *DhcpServer6DataData) GetSmartRef1Server6Id() string`

GetSmartRef1Server6Id returns the SmartRef1Server6Id field if non-nil, zero value otherwise.

### GetSmartRef1Server6IdOk

`func (o *DhcpServer6DataData) GetSmartRef1Server6IdOk() (*string, bool)`

GetSmartRef1Server6IdOk returns a tuple with the SmartRef1Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef1Server6Id

`func (o *DhcpServer6DataData) SetSmartRef1Server6Id(v string)`

SetSmartRef1Server6Id sets SmartRef1Server6Id field to given value.

### HasSmartRef1Server6Id

`func (o *DhcpServer6DataData) HasSmartRef1Server6Id() bool`

HasSmartRef1Server6Id returns a boolean if a field has been set.

### GetSmartRef1Server6Name

`func (o *DhcpServer6DataData) GetSmartRef1Server6Name() string`

GetSmartRef1Server6Name returns the SmartRef1Server6Name field if non-nil, zero value otherwise.

### GetSmartRef1Server6NameOk

`func (o *DhcpServer6DataData) GetSmartRef1Server6NameOk() (*string, bool)`

GetSmartRef1Server6NameOk returns a tuple with the SmartRef1Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef1Server6Name

`func (o *DhcpServer6DataData) SetSmartRef1Server6Name(v string)`

SetSmartRef1Server6Name sets SmartRef1Server6Name field to given value.

### HasSmartRef1Server6Name

`func (o *DhcpServer6DataData) HasSmartRef1Server6Name() bool`

HasSmartRef1Server6Name returns a boolean if a field has been set.

### GetSmartRef2Server6Id

`func (o *DhcpServer6DataData) GetSmartRef2Server6Id() string`

GetSmartRef2Server6Id returns the SmartRef2Server6Id field if non-nil, zero value otherwise.

### GetSmartRef2Server6IdOk

`func (o *DhcpServer6DataData) GetSmartRef2Server6IdOk() (*string, bool)`

GetSmartRef2Server6IdOk returns a tuple with the SmartRef2Server6Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef2Server6Id

`func (o *DhcpServer6DataData) SetSmartRef2Server6Id(v string)`

SetSmartRef2Server6Id sets SmartRef2Server6Id field to given value.

### HasSmartRef2Server6Id

`func (o *DhcpServer6DataData) HasSmartRef2Server6Id() bool`

HasSmartRef2Server6Id returns a boolean if a field has been set.

### GetSmartRef2Server6Name

`func (o *DhcpServer6DataData) GetSmartRef2Server6Name() string`

GetSmartRef2Server6Name returns the SmartRef2Server6Name field if non-nil, zero value otherwise.

### GetSmartRef2Server6NameOk

`func (o *DhcpServer6DataData) GetSmartRef2Server6NameOk() (*string, bool)`

GetSmartRef2Server6NameOk returns a tuple with the SmartRef2Server6Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSmartRef2Server6Name

`func (o *DhcpServer6DataData) SetSmartRef2Server6Name(v string)`

SetSmartRef2Server6Name sets SmartRef2Server6Name field to given value.

### HasSmartRef2Server6Name

`func (o *DhcpServer6DataData) HasSmartRef2Server6Name() bool`

HasSmartRef2Server6Name returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


