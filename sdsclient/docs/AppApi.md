# \AppApi

All URIs are relative to *https://vmdev-ct-fb11:443/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppApplicationAdd**](AppApi.md#AppApplicationAdd) | **Post** /app/application/add | Add an application
[**AppApplicationCount**](AppApi.md#AppApplicationCount) | **Get** /app/application/count | Count the number of applications
[**AppApplicationDelete**](AppApi.md#AppApplicationDelete) | **Delete** /app/application/delete | Delete an application
[**AppApplicationEdit**](AppApi.md#AppApplicationEdit) | **Put** /app/application/edit | Edit an application
[**AppApplicationInfo**](AppApi.md#AppApplicationInfo) | **Get** /app/application/info | Display the properties of an application
[**AppApplicationList**](AppApi.md#AppApplicationList) | **Get** /app/application/list | List the applications
[**AppNodeAdd**](AppApi.md#AppNodeAdd) | **Post** /app/node/add | Add a node
[**AppNodeCount**](AppApi.md#AppNodeCount) | **Get** /app/node/count | Count the number of nodes
[**AppNodeDelete**](AppApi.md#AppNodeDelete) | **Delete** /app/node/delete | Delete a node
[**AppNodeEdit**](AppApi.md#AppNodeEdit) | **Put** /app/node/edit | Edit a node
[**AppNodeInfo**](AppApi.md#AppNodeInfo) | **Get** /app/node/info | Display the properties of a node
[**AppNodeList**](AppApi.md#AppNodeList) | **Get** /app/node/list | List the nodes
[**AppPoolAdd**](AppApi.md#AppPoolAdd) | **Post** /app/pool/add | Add a pool
[**AppPoolCount**](AppApi.md#AppPoolCount) | **Get** /app/pool/count | Count the number of pools
[**AppPoolDelete**](AppApi.md#AppPoolDelete) | **Delete** /app/pool/delete | Delete a pool
[**AppPoolEdit**](AppApi.md#AppPoolEdit) | **Put** /app/pool/edit | Edit a pool
[**AppPoolInfo**](AppApi.md#AppPoolInfo) | **Get** /app/pool/info | Display the properties of a pool
[**AppPoolList**](AppApi.md#AppPoolList) | **Get** /app/pool/list | List the pools



## AppApplicationAdd

> AppApplicationAddSuccess AppApplicationAdd(ctx).AppApplicationAddInput(appApplicationAddInput).Execute()

Add an application



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
    appApplicationAddInput := *sdsclient.NewAppApplicationAddInput() // AppApplicationAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppApplicationAdd(context.Background()).AppApplicationAddInput(appApplicationAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppApplicationAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApplicationAdd`: AppApplicationAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppApplicationAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appApplicationAddInput** | [**AppApplicationAddInput**](AppApplicationAddInput.md) |  | 

### Return type

[**AppApplicationAddSuccess**](app_application_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppApplicationCount

> ApiCountResponseSuccess AppApplicationCount(ctx).Where(where).Tags(tags).Execute()

Count the number of applications



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
    resp, r, err := api_client.AppApi.AppApplicationCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppApplicationCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApplicationCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppApplicationCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationCountRequest struct via the builder pattern


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


## AppApplicationDelete

> AppApplicationDeleteSuccess AppApplicationDelete(ctx).ApplicationId(applicationId).ApplicationFqdn(applicationFqdn).ApplicationName(applicationName).GslbserverId(gslbserverId).GslbserverName(gslbserverName).Warnings(warnings).Execute()

Delete an application



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
    applicationId := int32(56) // int32 | The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice. (optional)
    applicationFqdn := "applicationFqdn_example" // string | The Fully Qualified Domain Name (FQDN) of the application. (optional)
    applicationName := "applicationName_example" // string | The name of the application. (optional)
    gslbserverId := int32(56) // int32 | The database identifier (ID) of the GSLB server associated with the application, a unique identifier automatically incremented when you add the server. Use it to identify the GSLB server of your choice. (optional)
    gslbserverName := "gslbserverName_example" // string | The name of the GSLB server associated with the application. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppApplicationDelete(context.Background()).ApplicationId(applicationId).ApplicationFqdn(applicationFqdn).ApplicationName(applicationName).GslbserverId(gslbserverId).GslbserverName(gslbserverName).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppApplicationDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApplicationDelete`: AppApplicationDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppApplicationDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **int32** | The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice. | 
 **applicationFqdn** | **string** | The Fully Qualified Domain Name (FQDN) of the application. | 
 **applicationName** | **string** | The name of the application. | 
 **gslbserverId** | **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique identifier automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | 
 **gslbserverName** | **string** | The name of the GSLB server associated with the application. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**AppApplicationDeleteSuccess**](app_application_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppApplicationEdit

> AppApplicationEditSuccess AppApplicationEdit(ctx).AppApplicationEditInput(appApplicationEditInput).Execute()

Edit an application



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
    appApplicationEditInput := *sdsclient.NewAppApplicationEditInput() // AppApplicationEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppApplicationEdit(context.Background()).AppApplicationEditInput(appApplicationEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppApplicationEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApplicationEdit`: AppApplicationEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppApplicationEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appApplicationEditInput** | [**AppApplicationEditInput**](AppApplicationEditInput.md) |  | 

### Return type

[**AppApplicationEditSuccess**](app_application_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppApplicationInfo

> AppApplicationData AppApplicationInfo(ctx).ApplicationId(applicationId).Execute()

Display the properties of an application



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
    applicationId := int32(56) // int32 | The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppApplicationInfo(context.Background()).ApplicationId(applicationId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppApplicationInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApplicationInfo`: AppApplicationData
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppApplicationInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **int32** | The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice. | 

### Return type

[**AppApplicationData**](app_application_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppApplicationList

> AppApplicationData AppApplicationList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the applications



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
    resp, r, err := api_client.AppApi.AppApplicationList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppApplicationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppApplicationList`: AppApplicationData
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppApplicationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**AppApplicationData**](app_application_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppNodeAdd

> AppNodeAddSuccess AppNodeAdd(ctx).AppNodeAddInput(appNodeAddInput).Execute()

Add a node



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
    appNodeAddInput := *sdsclient.NewAppNodeAddInput() // AppNodeAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppNodeAdd(context.Background()).AppNodeAddInput(appNodeAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppNodeAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppNodeAdd`: AppNodeAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppNodeAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appNodeAddInput** | [**AppNodeAddInput**](AppNodeAddInput.md) |  | 

### Return type

[**AppNodeAddSuccess**](app_node_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppNodeCount

> ApiCountResponseSuccess AppNodeCount(ctx).Where(where).Execute()

Count the number of nodes



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
    resp, r, err := api_client.AppApi.AppNodeCount(context.Background()).Where(where).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppNodeCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppNodeCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppNodeCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeCountRequest struct via the builder pattern


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


## AppNodeDelete

> AppNodeDeleteSuccess AppNodeDelete(ctx).ApplicationFqdn(applicationFqdn).ApplicationId(applicationId).ApplicationName(applicationName).NodeId(nodeId).PoolId(poolId).PoolName(poolName).NodeName(nodeName).GslbserverId(gslbserverId).Warnings(warnings).Execute()

Delete a node



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
    applicationFqdn := "applicationFqdn_example" // string | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. (optional)
    applicationId := int32(56) // int32 | The database identifier (ID) of the application the object belongs to. (optional)
    applicationName := "applicationName_example" // string | The name of the application the object belongs to. (optional)
    nodeId := int32(56) // int32 | The database identifier (ID) of the node, a unique numeric key value automatically incremented when you add a node. Use the ID to specify the node of your choice. (optional)
    poolId := int32(56) // int32 | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. (optional)
    poolName := "poolName_example" // string | The name of the pool. (optional)
    nodeName := "nodeName_example" // string | The name of the node. (optional)
    gslbserverId := int32(56) // int32 | The database identifier (ID) of the GSLB server associated with the application, a unique identifier automatically incremented when you add the server. Use it to identify the GSLB server of your choice. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppNodeDelete(context.Background()).ApplicationFqdn(applicationFqdn).ApplicationId(applicationId).ApplicationName(applicationName).NodeId(nodeId).PoolId(poolId).PoolName(poolName).NodeName(nodeName).GslbserverId(gslbserverId).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppNodeDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppNodeDelete`: AppNodeDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppNodeDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationFqdn** | **string** | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. | 
 **applicationId** | **int32** | The database identifier (ID) of the application the object belongs to. | 
 **applicationName** | **string** | The name of the application the object belongs to. | 
 **nodeId** | **int32** | The database identifier (ID) of the node, a unique numeric key value automatically incremented when you add a node. Use the ID to specify the node of your choice. | 
 **poolId** | **int32** | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. | 
 **poolName** | **string** | The name of the pool. | 
 **nodeName** | **string** | The name of the node. | 
 **gslbserverId** | **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique identifier automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**AppNodeDeleteSuccess**](app_node_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppNodeEdit

> AppNodeEditSuccess AppNodeEdit(ctx).AppNodeEditInput(appNodeEditInput).Execute()

Edit a node



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
    appNodeEditInput := *sdsclient.NewAppNodeEditInput() // AppNodeEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppNodeEdit(context.Background()).AppNodeEditInput(appNodeEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppNodeEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppNodeEdit`: AppNodeEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppNodeEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appNodeEditInput** | [**AppNodeEditInput**](AppNodeEditInput.md) |  | 

### Return type

[**AppNodeEditSuccess**](app_node_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppNodeInfo

> AppNodeData AppNodeInfo(ctx).NodeId(nodeId).Execute()

Display the properties of a node



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
    nodeId := int32(56) // int32 | The database identifier (ID) of the node, a unique numeric key value automatically incremented when you add a node. Use the ID to specify the node of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppNodeInfo(context.Background()).NodeId(nodeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppNodeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppNodeInfo`: AppNodeData
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppNodeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **int32** | The database identifier (ID) of the node, a unique numeric key value automatically incremented when you add a node. Use the ID to specify the node of your choice. | 

### Return type

[**AppNodeData**](app_node_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppNodeList

> AppNodeData AppNodeList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()

List the nodes



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
    resp, r, err := api_client.AppApi.AppNodeList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppNodeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppNodeList`: AppNodeData
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppNodeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 

### Return type

[**AppNodeData**](app_node_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPoolAdd

> AppPoolAddSuccess AppPoolAdd(ctx).AppPoolAddInput(appPoolAddInput).Execute()

Add a pool



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
    appPoolAddInput := *sdsclient.NewAppPoolAddInput() // AppPoolAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppPoolAdd(context.Background()).AppPoolAddInput(appPoolAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppPoolAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPoolAdd`: AppPoolAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppPoolAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appPoolAddInput** | [**AppPoolAddInput**](AppPoolAddInput.md) |  | 

### Return type

[**AppPoolAddSuccess**](app_pool_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPoolCount

> ApiCountResponseSuccess AppPoolCount(ctx).Where(where).Execute()

Count the number of pools



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
    resp, r, err := api_client.AppApi.AppPoolCount(context.Background()).Where(where).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppPoolCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPoolCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppPoolCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolCountRequest struct via the builder pattern


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


## AppPoolDelete

> AppPoolDeleteSuccess AppPoolDelete(ctx).ApplicationFqdn(applicationFqdn).ApplicationId(applicationId).ApplicationName(applicationName).PoolId(poolId).PoolName(poolName).GslbserverId(gslbserverId).Warnings(warnings).Execute()

Delete a pool



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
    applicationFqdn := "applicationFqdn_example" // string | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. (optional)
    applicationId := int32(56) // int32 | The database identifier (ID) of the application the object belongs to. (optional)
    applicationName := "applicationName_example" // string | The name of the application the object belongs to. (optional)
    poolId := int32(56) // int32 | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. (optional)
    poolName := "poolName_example" // string | The name of the pool. (optional)
    gslbserverId := int32(56) // int32 | The database identifier (ID) of the GSLB server associated with the application, a unique identifier automatically incremented when you add the server. Use it to identify the GSLB server of your choice. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppPoolDelete(context.Background()).ApplicationFqdn(applicationFqdn).ApplicationId(applicationId).ApplicationName(applicationName).PoolId(poolId).PoolName(poolName).GslbserverId(gslbserverId).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppPoolDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPoolDelete`: AppPoolDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppPoolDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationFqdn** | **string** | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. | 
 **applicationId** | **int32** | The database identifier (ID) of the application the object belongs to. | 
 **applicationName** | **string** | The name of the application the object belongs to. | 
 **poolId** | **int32** | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. | 
 **poolName** | **string** | The name of the pool. | 
 **gslbserverId** | **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique identifier automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**AppPoolDeleteSuccess**](app_pool_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPoolEdit

> AppPoolEditSuccess AppPoolEdit(ctx).AppPoolEditInput(appPoolEditInput).Execute()

Edit a pool



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
    appPoolEditInput := *sdsclient.NewAppPoolEditInput() // AppPoolEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppPoolEdit(context.Background()).AppPoolEditInput(appPoolEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppPoolEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPoolEdit`: AppPoolEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppPoolEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appPoolEditInput** | [**AppPoolEditInput**](AppPoolEditInput.md) |  | 

### Return type

[**AppPoolEditSuccess**](app_pool_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPoolInfo

> AppPoolData AppPoolInfo(ctx).PoolId(poolId).Execute()

Display the properties of a pool



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
    poolId := int32(56) // int32 | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.AppApi.AppPoolInfo(context.Background()).PoolId(poolId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppPoolInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPoolInfo`: AppPoolData
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppPoolInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **int32** | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. | 

### Return type

[**AppPoolData**](app_pool_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPoolList

> AppPoolData AppPoolList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()

List the pools



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
    resp, r, err := api_client.AppApi.AppPoolList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `AppApi.AppPoolList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AppPoolList`: AppPoolData
    fmt.Fprintf(os.Stdout, "Response from `AppApi.AppPoolList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 

### Return type

[**AppPoolData**](app_pool_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

