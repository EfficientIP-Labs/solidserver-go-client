# \DnsApi

All URIs are relative to *https://SDS-URL:443/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DnsAclAdd**](DnsApi.md#DnsAclAdd) | **Post** /dns/acl/add | Add a DNS ACL
[**DnsAclCount**](DnsApi.md#DnsAclCount) | **Get** /dns/acl/count | Count the number of DNS ACLs
[**DnsAclDelete**](DnsApi.md#DnsAclDelete) | **Delete** /dns/acl/delete | Delete a DNS ACL
[**DnsAclEdit**](DnsApi.md#DnsAclEdit) | **Put** /dns/acl/edit | Edit a DNS ACL
[**DnsAclInfo**](DnsApi.md#DnsAclInfo) | **Get** /dns/acl/info | Display the properties of a DNS ACL
[**DnsAclList**](DnsApi.md#DnsAclList) | **Get** /dns/acl/list | List the DNS ACLs
[**DnsRrAdd**](DnsApi.md#DnsRrAdd) | **Post** /dns/rr/add | Add a resource record
[**DnsRrCount**](DnsApi.md#DnsRrCount) | **Get** /dns/rr/count | Count the number of resource records
[**DnsRrDelete**](DnsApi.md#DnsRrDelete) | **Delete** /dns/rr/delete | Delete a resource record
[**DnsRrEdit**](DnsApi.md#DnsRrEdit) | **Put** /dns/rr/edit | Edit a resource record
[**DnsRrInfo**](DnsApi.md#DnsRrInfo) | **Get** /dns/rr/info | Display the properties of a resource record
[**DnsRrList**](DnsApi.md#DnsRrList) | **Get** /dns/rr/list | List the resource records
[**DnsServerCount**](DnsApi.md#DnsServerCount) | **Get** /dns/server/count | Count the number of DNS servers
[**DnsServerInfo**](DnsApi.md#DnsServerInfo) | **Get** /dns/server/info | Display the properties of a DNS server
[**DnsServerList**](DnsApi.md#DnsServerList) | **Get** /dns/server/list | List the DNS servers
[**DnsViewAdd**](DnsApi.md#DnsViewAdd) | **Post** /dns/view/add | Add a view
[**DnsViewCount**](DnsApi.md#DnsViewCount) | **Get** /dns/view/count | Count the number of views
[**DnsViewDelete**](DnsApi.md#DnsViewDelete) | **Delete** /dns/view/delete | Delete a view
[**DnsViewEdit**](DnsApi.md#DnsViewEdit) | **Put** /dns/view/edit | Edit a view
[**DnsViewInfo**](DnsApi.md#DnsViewInfo) | **Get** /dns/view/info | Display the properties of a view
[**DnsViewList**](DnsApi.md#DnsViewList) | **Get** /dns/view/list | List the views
[**DnsViewparamAdd**](DnsApi.md#DnsViewparamAdd) | **Post** /dns/viewparam/add | Add a DNS option on a view
[**DnsViewparamCount**](DnsApi.md#DnsViewparamCount) | **Get** /dns/viewparam/count | Count the number of DNS options of a view
[**DnsViewparamDelete**](DnsApi.md#DnsViewparamDelete) | **Delete** /dns/viewparam/delete | Delete a DNS option from a view
[**DnsViewparamEdit**](DnsApi.md#DnsViewparamEdit) | **Put** /dns/viewparam/edit | Edit a DNS option on a view
[**DnsViewparamInfo**](DnsApi.md#DnsViewparamInfo) | **Get** /dns/viewparam/info | Display the properties of a DNS option set on a view
[**DnsViewparamList**](DnsApi.md#DnsViewparamList) | **Get** /dns/viewparam/list | List the DNS options of a view
[**DnsZoneAdd**](DnsApi.md#DnsZoneAdd) | **Post** /dns/zone/add | Add a zone
[**DnsZoneCount**](DnsApi.md#DnsZoneCount) | **Get** /dns/zone/count | Count the number of zones
[**DnsZoneDelete**](DnsApi.md#DnsZoneDelete) | **Delete** /dns/zone/delete | Delete a zone
[**DnsZoneEdit**](DnsApi.md#DnsZoneEdit) | **Put** /dns/zone/edit | Edit a zone
[**DnsZoneInfo**](DnsApi.md#DnsZoneInfo) | **Get** /dns/zone/info | Display the properties of a zone
[**DnsZoneList**](DnsApi.md#DnsZoneList) | **Get** /dns/zone/list | List the DNS options of a zone
[**DnsZoneparamAdd**](DnsApi.md#DnsZoneparamAdd) | **Post** /dns/zoneparam/add | Add a DNS option on a zone
[**DnsZoneparamCount**](DnsApi.md#DnsZoneparamCount) | **Get** /dns/zoneparam/count | Count the number of DNS options of a zone
[**DnsZoneparamDelete**](DnsApi.md#DnsZoneparamDelete) | **Delete** /dns/zoneparam/delete | Delete a DNS option from a zone
[**DnsZoneparamEdit**](DnsApi.md#DnsZoneparamEdit) | **Put** /dns/zoneparam/edit | Edit a DNS option on a zone
[**DnsZoneparamInfo**](DnsApi.md#DnsZoneparamInfo) | **Get** /dns/zoneparam/info | Display the properties of a DNS option set on a zone
[**DnsZoneparamList**](DnsApi.md#DnsZoneparamList) | **Get** /dns/zoneparam/list | List the DNS options of a zone



## DnsAclAdd

> DnsAclAddSuccess DnsAclAdd(ctx).DnsAclAddInput(dnsAclAddInput).Execute()

Add a DNS ACL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsAclAddInput := *sdsclient.NewDnsAclAddInput() // DnsAclAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsAclAdd(context.Background()).DnsAclAddInput(dnsAclAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsAclAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsAclAdd`: DnsAclAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsAclAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsAclAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsAclAddInput** | [**DnsAclAddInput**](DnsAclAddInput.md) |  | 

### Return type

[**DnsAclAddSuccess**](dns_acl_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsAclCount

> ApiCountResponseSuccess DnsAclCount(ctx).Where(where).Execute()

Count the number of DNS ACLs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsAclCount(context.Background()).Where(where).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsAclCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsAclCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsAclCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsAclCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service &lt;b&gt;*_list&lt;/b&gt; of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is no longer possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsAclDelete

> DnsAclDeleteSuccess DnsAclDelete(ctx).AclId(aclId).ServerId(serverId).ServerName(serverName).AclName(aclName).ServerHostaddr(serverHostaddr).Warnings(warnings).Execute()

Delete a DNS ACL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    aclId := int32(56) // int32 | The database identifier (ID) of the DNS ACL, a unique numeric key value automatically incremented when you add a DNS ACL. Use the ID to specify the DNS ACL of your choice. (optional)
    serverId := int32(56) // int32 | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. (optional)
    serverName := "serverName_example" // string | The name of the DNS server. (optional)
    aclName := "aclName_example" // string | The name of the DNS ACL. (optional)
    serverHostaddr := "serverHostaddr_example" // string | The IP address of the DNS server. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsAclDelete(context.Background()).AclId(aclId).ServerId(serverId).ServerName(serverName).AclName(aclName).ServerHostaddr(serverHostaddr).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsAclDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsAclDelete`: DnsAclDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsAclDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsAclDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclId** | **int32** | The database identifier (ID) of the DNS ACL, a unique numeric key value automatically incremented when you add a DNS ACL. Use the ID to specify the DNS ACL of your choice. | 
 **serverId** | **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | 
 **serverName** | **string** | The name of the DNS server. | 
 **aclName** | **string** | The name of the DNS ACL. | 
 **serverHostaddr** | **string** | The IP address of the DNS server. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**DnsAclDeleteSuccess**](dns_acl_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsAclEdit

> DnsAclEditSuccess DnsAclEdit(ctx).DnsAclEditInput(dnsAclEditInput).Execute()

Edit a DNS ACL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsAclEditInput := *sdsclient.NewDnsAclEditInput() // DnsAclEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsAclEdit(context.Background()).DnsAclEditInput(dnsAclEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsAclEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsAclEdit`: DnsAclEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsAclEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsAclEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsAclEditInput** | [**DnsAclEditInput**](DnsAclEditInput.md) |  | 

### Return type

[**DnsAclEditSuccess**](dns_acl_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsAclInfo

> DnsAclData DnsAclInfo(ctx).AclId(aclId).Execute()

Display the properties of a DNS ACL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    aclId := int32(56) // int32 | The database identifier (ID) of the DNS ACL, a unique numeric key value automatically incremented when you add a DNS ACL. Use the ID to specify the DNS ACL of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsAclInfo(context.Background()).AclId(aclId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsAclInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsAclInfo`: DnsAclData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsAclInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsAclInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **aclId** | **int32** | The database identifier (ID) of the DNS ACL, a unique numeric key value automatically incremented when you add a DNS ACL. Use the ID to specify the DNS ACL of your choice. | 

### Return type

[**DnsAclData**](dns_acl_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsAclList

> DnsAclData DnsAclList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()

List the DNS ACLs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsAclList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsAclList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsAclList`: DnsAclData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsAclList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsAclListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 

### Return type

[**DnsAclData**](dns_acl_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsRrAdd

> DnsRrAddSuccess DnsRrAdd(ctx).DnsRrAddInput(dnsRrAddInput).Execute()

Add a resource record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsRrAddInput := *sdsclient.NewDnsRrAddInput() // DnsRrAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsRrAdd(context.Background()).DnsRrAddInput(dnsRrAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsRrAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRrAdd`: DnsRrAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsRrAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsRrAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsRrAddInput** | [**DnsRrAddInput**](DnsRrAddInput.md) |  | 

### Return type

[**DnsRrAddSuccess**](dns_rr_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsRrCount

> ApiCountResponseSuccess DnsRrCount(ctx).Where(where).Tags(tags).Execute()

Count the number of resource records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsRrCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsRrCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRrCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsRrCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsRrCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service &lt;b&gt;*_list&lt;/b&gt; of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is no longer possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsRrDelete

> DnsRrDeleteSuccess DnsRrDelete(ctx).ServerId(serverId).ServerName(serverName).ZoneId(zoneId).ZoneName(zoneName).ServerHostaddr(serverHostaddr).RrId(rrId).RrName(rrName).ViewId(viewId).ViewName(viewName).ZoneSpaceId(zoneSpaceId).RrType(rrType).RrValue1(rrValue1).RrValue2(rrValue2).RrValue3(rrValue3).RrValue4(rrValue4).RrValue5(rrValue5).RrValue6(rrValue6).RrValue7(rrValue7).Warnings(warnings).Execute()

Delete a resource record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    serverId := int32(56) // int32 | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. (optional)
    serverName := "serverName_example" // string | The name of the DNS server. (optional)
    zoneId := int32(56) // int32 | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. (optional)
    zoneName := "zoneName_example" // string | The name of the DNS zone the object belongs to. (optional)
    serverHostaddr := "serverHostaddr_example" // string | The IP address of the DNS server. (optional)
    rrId := int32(56) // int32 | The database identifier (ID) of the DNS resource record, a unique numeric key value automatically incremented when you add a DNS RR. Use the ID to specify the DNS RR of your choice. (optional)
    rrName := "rrName_example" // string | The name of the DNS resource record. (optional)
    viewId := int32(56) // int32 | The database identifier (ID) of the DNS view, a unique numeric key value automatically incremented when you add a DNS view. Use the ID to specify the DNS view of your choice. (optional)
    viewName := "viewName_example" // string | The name of the DNS view. (optional)
    zoneSpaceId := int32(56) // int32 | The database identifier (ID) of the space associated with the DNS zone the record belongs to. (optional)
    rrType := "rrType_example" // string | The type of the DNS resource record.<table><caption>rr_type possible values</caption><br/><thead><tr><th>Value</th><th>Record type description</th></tr><br/></thead><br/><tbody><tr><td >SOA</td><td >Start of Authority. Defines the zone name, an email contact and various time and refresh values applicable to the zone. It is automatically generated upon creation of a zone and cannot be added manually.</td></tr><tr><td >NS</td><td >Name Server. Defines the authoritative name server(s) for the domain (defined by the SOA record) or the subdomain. The NS record that indicates which server has authority over a zone is automatically generated upon the creation of a zone, once the server has been synchronized.</td></tr><tr><td >A</td><td >IPv4 Address. An IPv4 address for a host.</td></tr><tr><td >PTR</td><td >Pointer Record. Address Resolution, from an IP address (IPv4 or IPv6) to a host. Used in reverse mapping.</td></tr><tr><td >AAAA</td><td >IPv6 Address. An IPv6 address for a host.</td></tr><tr><td >CNAME</td><td >Canonical Name. An alias name for a host.</td></tr><tr><td >MX</td><td >Mail Exchange. The mail server/exchanger that services this zone.</td></tr><tr><td >SRV</td><td >Services record. Defines services available in the zone, for example, LDAP, HTTP, etc...</td></tr><tr><td >DNAME</td><td >Delegation of Reverse Names. Delegation of reverse addresses primarily in IPv6. (Deprecated, use the CNAME RR instead)</td></tr><tr><td >TXT</td><td >Text. Information associated with a name.</td></tr><tr><td >DS</td><td >Delegation Signer, a DNSSEC related RR used to verify the validity of the ZSK of a subdomain.</td></tr><tr><td >DNSKEY</td><td >DNS Key. It contains the public cryptographic key used to sign the zone with DNSSEC.</td></tr><tr><td >65534</td><td >A private type record automatically added to the zone once it is signed with DNSSEC.</td></tr><tr><td >HINFO</td><td >System Information. Information about a host: hardware type and operating system description.</td></tr><tr><td >MINFO</td><td >Mailbox mail list Information. Defines the mail administrator for a mail list and optionally a mailbox to receive error messages relating to the mail list.</td></tr><tr><td >AFSDB</td><td >AFS Database. Location of the AFS servers.</td></tr><tr><td >WKS</td><td >Well-Known Service. Defines the services and protocols supported by a host. (Deprecated, use the SRV RR instead)</td></tr><tr><td >NAPTR</td><td >Naming Authority Pointer Record. General purpose definition of rule set to be used by applications e.g. VoIP.</td></tr><tr><td >NSAP</td><td >Network Service Access Point. Defines record (equivalent of an A record) maps a host name to an endpoint address.</td></tr></tbody></table></p><br/>Note that the parameter is not case sensitive, you could specify <b>A</b> or <b>a</b>. (optional)
    rrValue1 := "rrValue1_example" // string | The first or only value required for the DNS resource record, as detailed in the service description (optional)
    rrValue2 := "rrValue2_example" // string | The second value of the DNS resource record, depending on its type, as detailed in the service description (optional)
    rrValue3 := "rrValue3_example" // string | The third value of the DNS resource record, depending on its type, as detailed in the service description (optional)
    rrValue4 := "rrValue4_example" // string | The fourth value of the DNS resource record, depending on its type, as detailed in the service description (optional)
    rrValue5 := "rrValue5_example" // string | The fifth value of the DNS resource record, depending on its type, as detailed in the service description (optional)
    rrValue6 := "rrValue6_example" // string | The sixth value of the DNS resource record, depending on its type, as detailed in the service description (optional)
    rrValue7 := "rrValue7_example" // string | The seventh value of the DNS resource record, depending on its type, as detailed in the service description (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsRrDelete(context.Background()).ServerId(serverId).ServerName(serverName).ZoneId(zoneId).ZoneName(zoneName).ServerHostaddr(serverHostaddr).RrId(rrId).RrName(rrName).ViewId(viewId).ViewName(viewName).ZoneSpaceId(zoneSpaceId).RrType(rrType).RrValue1(rrValue1).RrValue2(rrValue2).RrValue3(rrValue3).RrValue4(rrValue4).RrValue5(rrValue5).RrValue6(rrValue6).RrValue7(rrValue7).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsRrDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRrDelete`: DnsRrDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsRrDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsRrDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | 
 **serverName** | **string** | The name of the DNS server. | 
 **zoneId** | **int32** | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. | 
 **zoneName** | **string** | The name of the DNS zone the object belongs to. | 
 **serverHostaddr** | **string** | The IP address of the DNS server. | 
 **rrId** | **int32** | The database identifier (ID) of the DNS resource record, a unique numeric key value automatically incremented when you add a DNS RR. Use the ID to specify the DNS RR of your choice. | 
 **rrName** | **string** | The name of the DNS resource record. | 
 **viewId** | **int32** | The database identifier (ID) of the DNS view, a unique numeric key value automatically incremented when you add a DNS view. Use the ID to specify the DNS view of your choice. | 
 **viewName** | **string** | The name of the DNS view. | 
 **zoneSpaceId** | **int32** | The database identifier (ID) of the space associated with the DNS zone the record belongs to. | 
 **rrType** | **string** | The type of the DNS resource record.&lt;table&gt;&lt;caption&gt;rr_type possible values&lt;/caption&gt;&lt;br/&gt;&lt;thead&gt;&lt;tr&gt;&lt;th&gt;Value&lt;/th&gt;&lt;th&gt;Record type description&lt;/th&gt;&lt;/tr&gt;&lt;br/&gt;&lt;/thead&gt;&lt;br/&gt;&lt;tbody&gt;&lt;tr&gt;&lt;td &gt;SOA&lt;/td&gt;&lt;td &gt;Start of Authority. Defines the zone name, an email contact and various time and refresh values applicable to the zone. It is automatically generated upon creation of a zone and cannot be added manually.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NS&lt;/td&gt;&lt;td &gt;Name Server. Defines the authoritative name server(s) for the domain (defined by the SOA record) or the subdomain. The NS record that indicates which server has authority over a zone is automatically generated upon the creation of a zone, once the server has been synchronized.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;A&lt;/td&gt;&lt;td &gt;IPv4 Address. An IPv4 address for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;PTR&lt;/td&gt;&lt;td &gt;Pointer Record. Address Resolution, from an IP address (IPv4 or IPv6) to a host. Used in reverse mapping.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;AAAA&lt;/td&gt;&lt;td &gt;IPv6 Address. An IPv6 address for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;CNAME&lt;/td&gt;&lt;td &gt;Canonical Name. An alias name for a host.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;MX&lt;/td&gt;&lt;td &gt;Mail Exchange. The mail server/exchanger that services this zone.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;SRV&lt;/td&gt;&lt;td &gt;Services record. Defines services available in the zone, for example, LDAP, HTTP, etc...&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DNAME&lt;/td&gt;&lt;td &gt;Delegation of Reverse Names. Delegation of reverse addresses primarily in IPv6. (Deprecated, use the CNAME RR instead)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;TXT&lt;/td&gt;&lt;td &gt;Text. Information associated with a name.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DS&lt;/td&gt;&lt;td &gt;Delegation Signer, a DNSSEC related RR used to verify the validity of the ZSK of a subdomain.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;DNSKEY&lt;/td&gt;&lt;td &gt;DNS Key. It contains the public cryptographic key used to sign the zone with DNSSEC.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;65534&lt;/td&gt;&lt;td &gt;A private type record automatically added to the zone once it is signed with DNSSEC.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;HINFO&lt;/td&gt;&lt;td &gt;System Information. Information about a host: hardware type and operating system description.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;MINFO&lt;/td&gt;&lt;td &gt;Mailbox mail list Information. Defines the mail administrator for a mail list and optionally a mailbox to receive error messages relating to the mail list.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;AFSDB&lt;/td&gt;&lt;td &gt;AFS Database. Location of the AFS servers.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;WKS&lt;/td&gt;&lt;td &gt;Well-Known Service. Defines the services and protocols supported by a host. (Deprecated, use the SRV RR instead)&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NAPTR&lt;/td&gt;&lt;td &gt;Naming Authority Pointer Record. General purpose definition of rule set to be used by applications e.g. VoIP.&lt;/td&gt;&lt;/tr&gt;&lt;tr&gt;&lt;td &gt;NSAP&lt;/td&gt;&lt;td &gt;Network Service Access Point. Defines record (equivalent of an A record) maps a host name to an endpoint address.&lt;/td&gt;&lt;/tr&gt;&lt;/tbody&gt;&lt;/table&gt;&lt;/p&gt;&lt;br/&gt;Note that the parameter is not case sensitive, you could specify &lt;b&gt;A&lt;/b&gt; or &lt;b&gt;a&lt;/b&gt;. | 
 **rrValue1** | **string** | The first or only value required for the DNS resource record, as detailed in the service description | 
 **rrValue2** | **string** | The second value of the DNS resource record, depending on its type, as detailed in the service description | 
 **rrValue3** | **string** | The third value of the DNS resource record, depending on its type, as detailed in the service description | 
 **rrValue4** | **string** | The fourth value of the DNS resource record, depending on its type, as detailed in the service description | 
 **rrValue5** | **string** | The fifth value of the DNS resource record, depending on its type, as detailed in the service description | 
 **rrValue6** | **string** | The sixth value of the DNS resource record, depending on its type, as detailed in the service description | 
 **rrValue7** | **string** | The seventh value of the DNS resource record, depending on its type, as detailed in the service description | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**DnsRrDeleteSuccess**](dns_rr_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsRrEdit

> DnsRrEditSuccess DnsRrEdit(ctx).DnsRrEditInput(dnsRrEditInput).Execute()

Edit a resource record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsRrEditInput := *sdsclient.NewDnsRrEditInput() // DnsRrEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsRrEdit(context.Background()).DnsRrEditInput(dnsRrEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsRrEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRrEdit`: DnsRrEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsRrEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsRrEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsRrEditInput** | [**DnsRrEditInput**](DnsRrEditInput.md) |  | 

### Return type

[**DnsRrEditSuccess**](dns_rr_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsRrInfo

> DnsRrData DnsRrInfo(ctx).RrId(rrId).Execute()

Display the properties of a resource record



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    rrId := int32(56) // int32 | The database identifier (ID) of the DNS resource record, a unique numeric key value automatically incremented when you add a DNS RR. Use the ID to specify the DNS RR of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsRrInfo(context.Background()).RrId(rrId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsRrInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRrInfo`: DnsRrData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsRrInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsRrInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rrId** | **int32** | The database identifier (ID) of the DNS resource record, a unique numeric key value automatically incremented when you add a DNS RR. Use the ID to specify the DNS RR of your choice. | 

### Return type

[**DnsRrData**](dns_rr_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsRrList

> DnsRrData DnsRrList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the resource records



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsRrList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsRrList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsRrList`: DnsRrData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsRrList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsRrListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**DnsRrData**](dns_rr_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsServerCount

> ApiCountResponseSuccess DnsServerCount(ctx).Where(where).Tags(tags).Execute()

Count the number of DNS servers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsServerCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsServerCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsServerCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsServerCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsServerCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service &lt;b&gt;*_list&lt;/b&gt; of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is no longer possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsServerInfo

> DnsServerData DnsServerInfo(ctx).ServerId(serverId).Execute()

Display the properties of a DNS server



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    serverId := int32(56) // int32 | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsServerInfo(context.Background()).ServerId(serverId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsServerInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsServerInfo`: DnsServerData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsServerInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsServerInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | 

### Return type

[**DnsServerData**](dns_server_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsServerList

> DnsServerData DnsServerList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the DNS servers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsServerList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsServerList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsServerList`: DnsServerData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsServerList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsServerListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**DnsServerData**](dns_server_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewAdd

> DnsViewAddSuccess DnsViewAdd(ctx).DnsViewAddInput(dnsViewAddInput).Execute()

Add a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsViewAddInput := *sdsclient.NewDnsViewAddInput() // DnsViewAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewAdd(context.Background()).DnsViewAddInput(dnsViewAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewAdd`: DnsViewAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsViewAddInput** | [**DnsViewAddInput**](DnsViewAddInput.md) |  | 

### Return type

[**DnsViewAddSuccess**](dns_view_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewCount

> ApiCountResponseSuccess DnsViewCount(ctx).Where(where).Tags(tags).Execute()

Count the number of views



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service &lt;b&gt;*_list&lt;/b&gt; of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is no longer possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewDelete

> DnsViewDeleteSuccess DnsViewDelete(ctx).ServerId(serverId).ServerName(serverName).ViewId(viewId).ViewName(viewName).ServerHostaddr(serverHostaddr).Warnings(warnings).Execute()

Delete a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    serverId := int32(56) // int32 | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. (optional)
    serverName := "serverName_example" // string | The name of the DNS server. (optional)
    viewId := int32(56) // int32 | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. (optional)
    viewName := "viewName_example" // string | The name of the DNS view. (optional)
    serverHostaddr := "serverHostaddr_example" // string | The IP address of the DNS server. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewDelete(context.Background()).ServerId(serverId).ServerName(serverName).ViewId(viewId).ViewName(viewName).ServerHostaddr(serverHostaddr).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewDelete`: DnsViewDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | 
 **serverName** | **string** | The name of the DNS server. | 
 **viewId** | **int32** | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. | 
 **viewName** | **string** | The name of the DNS view. | 
 **serverHostaddr** | **string** | The IP address of the DNS server. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**DnsViewDeleteSuccess**](dns_view_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewEdit

> DnsViewEditSuccess DnsViewEdit(ctx).DnsViewEditInput(dnsViewEditInput).Execute()

Edit a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsViewEditInput := *sdsclient.NewDnsViewEditInput() // DnsViewEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewEdit(context.Background()).DnsViewEditInput(dnsViewEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewEdit`: DnsViewEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsViewEditInput** | [**DnsViewEditInput**](DnsViewEditInput.md) |  | 

### Return type

[**DnsViewEditSuccess**](dns_view_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewInfo

> DnsViewData DnsViewInfo(ctx).ViewId(viewId).Execute()

Display the properties of a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    viewId := int32(56) // int32 | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewInfo(context.Background()).ViewId(viewId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewInfo`: DnsViewData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewId** | **int32** | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. | 

### Return type

[**DnsViewData**](dns_view_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewList

> DnsViewData DnsViewList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the views



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewList`: DnsViewData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**DnsViewData**](dns_view_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewparamAdd

> DnsViewparamAddSuccess DnsViewparamAdd(ctx).DnsViewparamAddInput(dnsViewparamAddInput).Execute()

Add a DNS option on a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsViewparamAddInput := *sdsclient.NewDnsViewparamAddInput() // DnsViewparamAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewparamAdd(context.Background()).DnsViewparamAddInput(dnsViewparamAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewparamAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewparamAdd`: DnsViewparamAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewparamAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewparamAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsViewparamAddInput** | [**DnsViewparamAddInput**](DnsViewparamAddInput.md) |  | 

### Return type

[**DnsViewparamAddSuccess**](dns_viewparam_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewparamCount

> ApiCountResponseSuccess DnsViewparamCount(ctx).Where(where).Execute()

Count the number of DNS options of a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewparamCount(context.Background()).Where(where).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewparamCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewparamCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewparamCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewparamCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service &lt;b&gt;*_list&lt;/b&gt; of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is no longer possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewparamDelete

> DnsViewparamDeleteSuccess DnsViewparamDelete(ctx).ViewId(viewId).ViewparamKey(viewparamKey).Execute()

Delete a DNS option from a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    viewId := int32(56) // int32 | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. (optional)
    viewparamKey := "viewparamKey_example" // string | The name of the DNS option that you want to remove from the view: <b>param_key=&lt;option-name&gt;</b>. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewparamDelete(context.Background()).ViewId(viewId).ViewparamKey(viewparamKey).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewparamDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewparamDelete`: DnsViewparamDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewparamDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewparamDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewId** | **int32** | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. | 
 **viewparamKey** | **string** | The name of the DNS option that you want to remove from the view: &lt;b&gt;param_key&#x3D;&amp;lt;option-name&amp;gt;&lt;/b&gt;. | 

### Return type

[**DnsViewparamDeleteSuccess**](dns_viewparam_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewparamEdit

> DnsViewparamEditSuccess DnsViewparamEdit(ctx).DnsViewparamEditInput(dnsViewparamEditInput).Execute()

Edit a DNS option on a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsViewparamEditInput := *sdsclient.NewDnsViewparamEditInput() // DnsViewparamEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewparamEdit(context.Background()).DnsViewparamEditInput(dnsViewparamEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewparamEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewparamEdit`: DnsViewparamEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewparamEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewparamEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsViewparamEditInput** | [**DnsViewparamEditInput**](DnsViewparamEditInput.md) |  | 

### Return type

[**DnsViewparamEditSuccess**](dns_viewparam_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewparamInfo

> DnsViewparamData DnsViewparamInfo(ctx).ViewId(viewId).Execute()

Display the properties of a DNS option set on a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    viewId := int32(56) // int32 | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewparamInfo(context.Background()).ViewId(viewId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewparamInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewparamInfo`: DnsViewparamData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewparamInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewparamInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **viewId** | **int32** | The database identifier (ID) of the DNS view. Use the ID to specify the DNS view of your choice. | 

### Return type

[**DnsViewparamData**](dns_viewparam_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsViewparamList

> DnsViewparamData DnsViewparamList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()

List the DNS options of a view



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsViewparamList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsViewparamList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsViewparamList`: DnsViewparamData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsViewparamList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsViewparamListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 

### Return type

[**DnsViewparamData**](dns_viewparam_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneAdd

> DnsZoneAddSuccess DnsZoneAdd(ctx).DnsZoneAddInput(dnsZoneAddInput).Execute()

Add a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsZoneAddInput := *sdsclient.NewDnsZoneAddInput() // DnsZoneAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneAdd(context.Background()).DnsZoneAddInput(dnsZoneAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneAdd`: DnsZoneAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsZoneAddInput** | [**DnsZoneAddInput**](DnsZoneAddInput.md) |  | 

### Return type

[**DnsZoneAddSuccess**](dns_zone_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneCount

> ApiCountResponseSuccess DnsZoneCount(ctx).Where(where).Tags(tags).Execute()

Count the number of zones



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service &lt;b&gt;*_list&lt;/b&gt; of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is no longer possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneDelete

> DnsZoneDeleteSuccess DnsZoneDelete(ctx).ServerId(serverId).ServerName(serverName).ViewId(viewId).ViewName(viewName).ZoneId(zoneId).ZoneName(zoneName).ServerHostaddr(serverHostaddr).Warnings(warnings).Execute()

Delete a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    serverId := int32(56) // int32 | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. (optional)
    serverName := "serverName_example" // string | The name of the DNS server. (optional)
    viewId := int32(56) // int32 | The database identifier (ID) of the DNS view the object belongs to. (optional)
    viewName := "viewName_example" // string | The name of the DNS view the object belongs to. (optional)
    zoneId := int32(56) // int32 | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. (optional)
    zoneName := "zoneName_example" // string | The name of the DNS zone the object belongs to. (optional)
    serverHostaddr := "serverHostaddr_example" // string | The IP address of the DNS server. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneDelete(context.Background()).ServerId(serverId).ServerName(serverName).ViewId(viewId).ViewName(viewName).ZoneId(zoneId).ZoneName(zoneName).ServerHostaddr(serverHostaddr).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneDelete`: DnsZoneDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **int32** | The database identifier (ID) of the DNS server, a unique numeric key value automatically incremented when you add a DNS server. Use the ID to specify the DNS server of your choice. | 
 **serverName** | **string** | The name of the DNS server. | 
 **viewId** | **int32** | The database identifier (ID) of the DNS view the object belongs to. | 
 **viewName** | **string** | The name of the DNS view the object belongs to. | 
 **zoneId** | **int32** | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. | 
 **zoneName** | **string** | The name of the DNS zone the object belongs to. | 
 **serverHostaddr** | **string** | The IP address of the DNS server. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**DnsZoneDeleteSuccess**](dns_zone_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneEdit

> DnsZoneEditSuccess DnsZoneEdit(ctx).DnsZoneEditInput(dnsZoneEditInput).Execute()

Edit a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsZoneEditInput := *sdsclient.NewDnsZoneEditInput() // DnsZoneEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneEdit(context.Background()).DnsZoneEditInput(dnsZoneEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneEdit`: DnsZoneEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsZoneEditInput** | [**DnsZoneEditInput**](DnsZoneEditInput.md) |  | 

### Return type

[**DnsZoneEditSuccess**](dns_zone_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneInfo

> DnsZoneData DnsZoneInfo(ctx).ZoneId(zoneId).Execute()

Display the properties of a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    zoneId := int32(56) // int32 | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneInfo(context.Background()).ZoneId(zoneId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneInfo`: DnsZoneData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int32** | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. | 

### Return type

[**DnsZoneData**](dns_zone_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneList

> DnsZoneData DnsZoneList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the DNS options of a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneList`: DnsZoneData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**DnsZoneData**](dns_zone_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneparamAdd

> DnsZoneparamAddSuccess DnsZoneparamAdd(ctx).DnsZoneparamAddInput(dnsZoneparamAddInput).Execute()

Add a DNS option on a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsZoneparamAddInput := *sdsclient.NewDnsZoneparamAddInput() // DnsZoneparamAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneparamAdd(context.Background()).DnsZoneparamAddInput(dnsZoneparamAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneparamAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneparamAdd`: DnsZoneparamAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneparamAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneparamAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsZoneparamAddInput** | [**DnsZoneparamAddInput**](DnsZoneparamAddInput.md) |  | 

### Return type

[**DnsZoneparamAddSuccess**](dns_zoneparam_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneparamCount

> ApiCountResponseSuccess DnsZoneparamCount(ctx).Where(where).Execute()

Count the number of DNS options of a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneparamCount(context.Background()).Where(where).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneparamCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneparamCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneparamCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneparamCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service &lt;b&gt;*_list&lt;/b&gt; of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is no longer possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneparamDelete

> DnsZoneparamDeleteSuccess DnsZoneparamDelete(ctx).ZoneId(zoneId).ZoneparamKey(zoneparamKey).Execute()

Delete a DNS option from a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    zoneId := int32(56) // int32 | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. (optional)
    zoneparamKey := "zoneparamKey_example" // string | The name of the DNS option that you want to remove from the zone: <b>param_key=&lt;option-name&gt;</b>. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneparamDelete(context.Background()).ZoneId(zoneId).ZoneparamKey(zoneparamKey).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneparamDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneparamDelete`: DnsZoneparamDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneparamDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneparamDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int32** | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. | 
 **zoneparamKey** | **string** | The name of the DNS option that you want to remove from the zone: &lt;b&gt;param_key&#x3D;&amp;lt;option-name&amp;gt;&lt;/b&gt;. | 

### Return type

[**DnsZoneparamDeleteSuccess**](dns_zoneparam_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneparamEdit

> DnsZoneparamEditSuccess DnsZoneparamEdit(ctx).DnsZoneparamEditInput(dnsZoneparamEditInput).Execute()

Edit a DNS option on a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    dnsZoneparamEditInput := *sdsclient.NewDnsZoneparamEditInput() // DnsZoneparamEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneparamEdit(context.Background()).DnsZoneparamEditInput(dnsZoneparamEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneparamEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneparamEdit`: DnsZoneparamEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneparamEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneparamEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dnsZoneparamEditInput** | [**DnsZoneparamEditInput**](DnsZoneparamEditInput.md) |  | 

### Return type

[**DnsZoneparamEditSuccess**](dns_zoneparam_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneparamInfo

> DnsZoneparamData DnsZoneparamInfo(ctx).ZoneId(zoneId).Execute()

Display the properties of a DNS option set on a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    zoneId := int32(56) // int32 | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneparamInfo(context.Background()).ZoneId(zoneId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneparamInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneparamInfo`: DnsZoneparamData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneparamInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneparamInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **zoneId** | **int32** | The database identifier (ID) of the DNS zone, a unique numeric key value automatically incremented when you add a DNS zone. Use the ID to specify the DNS zone of your choice. | 

### Return type

[**DnsZoneparamData**](dns_zoneparam_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DnsZoneparamList

> DnsZoneparamData DnsZoneparamList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()

List the DNS options of a zone



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DnsApi.DnsZoneparamList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DnsApi.DnsZoneparamList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DnsZoneparamList`: DnsZoneparamData
    fmt.Fprintf(os.Stdout, "Response from `DnsApi.DnsZoneparamList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDnsZoneparamListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 

### Return type

[**DnsZoneparamData**](dns_zoneparam_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

