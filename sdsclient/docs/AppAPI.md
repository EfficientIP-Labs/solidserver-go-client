# \AppAPI

All URIs are relative to *https://sds-ip-or-name:443/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppApplicationAdd**](AppAPI.md#AppApplicationAdd) | **Post** /app/application/add | Add an application
[**AppApplicationCount**](AppAPI.md#AppApplicationCount) | **Get** /app/application/count | Count the number of applications
[**AppApplicationDelete**](AppAPI.md#AppApplicationDelete) | **Delete** /app/application/delete | Delete an application
[**AppApplicationEdit**](AppAPI.md#AppApplicationEdit) | **Put** /app/application/edit | Edit an application
[**AppApplicationInfo**](AppAPI.md#AppApplicationInfo) | **Get** /app/application/info | Display the properties of an application
[**AppApplicationList**](AppAPI.md#AppApplicationList) | **Get** /app/application/list | List the applications
[**AppNodeAdd**](AppAPI.md#AppNodeAdd) | **Post** /app/node/add | Add a node
[**AppNodeCount**](AppAPI.md#AppNodeCount) | **Get** /app/node/count | Count the number of nodes
[**AppNodeDelete**](AppAPI.md#AppNodeDelete) | **Delete** /app/node/delete | Delete a node
[**AppNodeEdit**](AppAPI.md#AppNodeEdit) | **Put** /app/node/edit | Edit a node
[**AppNodeInfo**](AppAPI.md#AppNodeInfo) | **Get** /app/node/info | Display the properties of a node
[**AppNodeList**](AppAPI.md#AppNodeList) | **Get** /app/node/list | List the nodes
[**AppPoolAdd**](AppAPI.md#AppPoolAdd) | **Post** /app/pool/add | Add a pool
[**AppPoolCount**](AppAPI.md#AppPoolCount) | **Get** /app/pool/count | Count the number of pools
[**AppPoolDelete**](AppAPI.md#AppPoolDelete) | **Delete** /app/pool/delete | Delete a pool
[**AppPoolEdit**](AppAPI.md#AppPoolEdit) | **Put** /app/pool/edit | Edit a pool
[**AppPoolInfo**](AppAPI.md#AppPoolInfo) | **Get** /app/pool/info | Display the properties of a pool
[**AppPoolList**](AppAPI.md#AppPoolList) | **Get** /app/pool/list | List the pools



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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	appApplicationAddInput := *openapiclient.NewAppApplicationAddInput() // AppApplicationAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppApplicationAdd(context.Background()).AppApplicationAddInput(appApplicationAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppApplicationAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppApplicationAdd`: AppApplicationAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppApplicationAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appApplicationAddInput** | [**AppApplicationAddInput**](AppApplicationAddInput.md) |  | 

### Return type

[**AppApplicationAddSuccess**](AppApplicationAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	where := "where_example" // string | A clause that allows you to filter the result. You can include any output parameter of the service <object>/list of the object in this clause.To filter the result using class parameters, you must tag them first. (optional)
	tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppApplicationCount(context.Background()).Where(where).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppApplicationCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppApplicationCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppApplicationCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service &lt;object&gt;/list of the object in this clause.To filter the result using class parameters, you must tag them first. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	applicationId := int32(56) // int32 | The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice. (optional)
	applicationFqdn := "applicationFqdn_example" // string | The Fully Qualified Domain Name (FQDN) of the application. (optional)
	applicationName := "applicationName_example" // string | The name of the application. (optional)
	gslbserverId := int32(56) // int32 | The database identifier (ID) of the GSLB server associated with the application, a unique numeric key value automatically incremented when you add the server. Use it to identify the GSLB server of your choice. (optional)
	gslbserverName := "gslbserverName_example" // string | The name of the GSLB server associated with the application. (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppApplicationDelete(context.Background()).ApplicationId(applicationId).ApplicationFqdn(applicationFqdn).ApplicationName(applicationName).GslbserverId(gslbserverId).GslbserverName(gslbserverName).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppApplicationDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppApplicationDelete`: AppApplicationDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppApplicationDelete`: %v\n", resp)
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
 **gslbserverId** | **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique numeric key value automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | 
 **gslbserverName** | **string** | The name of the GSLB server associated with the application. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**AppApplicationDeleteSuccess**](AppApplicationDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	appApplicationEditInput := *openapiclient.NewAppApplicationEditInput() // AppApplicationEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppApplicationEdit(context.Background()).AppApplicationEditInput(appApplicationEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppApplicationEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppApplicationEdit`: AppApplicationEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppApplicationEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appApplicationEditInput** | [**AppApplicationEditInput**](AppApplicationEditInput.md) |  | 

### Return type

[**AppApplicationEditSuccess**](AppApplicationEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	applicationId := int32(56) // int32 | The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppApplicationInfo(context.Background()).ApplicationId(applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppApplicationInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppApplicationInfo`: AppApplicationData
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppApplicationInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationId** | **int32** | The database identifier (ID) of the application, a unique numeric key value automatically incremented when you add an application. Use the ID to specify the application of your choice. | 

### Return type

[**AppApplicationData**](AppApplicationData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppApplicationList

> AppApplicationData AppApplicationList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the applications



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	orderby := "orderby_example" // string | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
	select_ := "select__example" // string | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service <object>/list.If you specify several parameters they must be separated by a comma as follows: <b>select=&lt;param1&gt;<b>,</b>&lt;param2&gt;<b>,</b>...</b> .If the call includes the clause <b>where</b>, all the parameters it contains must be specified in the statement <b>select</b>.If the call includes the clause <b>orderby</b>, all the parameters it contains must be specified in the statement <b>select</b>.To include class parameters in the statement, you must tag them first. (optional)
	where := "where_example" // string | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first<br/>.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause is case insensitive and must be encoded in URL format. (optional)
	limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
	offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)
	tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppApplicationList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppApplicationList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppApplicationList`: AppApplicationData
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppApplicationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **select_** | **string** | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service &lt;object&gt;/list.If you specify several parameters they must be separated by a comma as follows: &lt;b&gt;select&#x3D;&amp;lt;param1&amp;gt;&lt;b&gt;,&lt;/b&gt;&amp;lt;param2&amp;gt;&lt;b&gt;,&lt;/b&gt;...&lt;/b&gt; .If the call includes the clause &lt;b&gt;where&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.If the call includes the clause &lt;b&gt;orderby&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.To include class parameters in the statement, you must tag them first. | 
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first&lt;br/&gt;.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause is case insensitive and must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**AppApplicationData**](AppApplicationData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	appNodeAddInput := *openapiclient.NewAppNodeAddInput() // AppNodeAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppNodeAdd(context.Background()).AppNodeAddInput(appNodeAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppNodeAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppNodeAdd`: AppNodeAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppNodeAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appNodeAddInput** | [**AppNodeAddInput**](AppNodeAddInput.md) |  | 

### Return type

[**AppNodeAddSuccess**](AppNodeAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	where := "where_example" // string | A clause that allows you to filter the result. You can include any output parameter of the service <object>/list of the object in this clause.To filter the result using class parameters, you must tag them first. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppNodeCount(context.Background()).Where(where).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppNodeCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppNodeCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppNodeCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service &lt;object&gt;/list of the object in this clause.To filter the result using class parameters, you must tag them first. | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	applicationFqdn := "applicationFqdn_example" // string | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. (optional)
	applicationId := int32(56) // int32 | The database identifier (ID) of the application the object belongs to. (optional)
	applicationName := "applicationName_example" // string | The name of the application the object belongs to. (optional)
	nodeId := int32(56) // int32 | The database identifier (ID) of the node, a unique numeric key value automatically incremented when you add a node. Use the ID to specify the node of your choice. (optional)
	poolId := int32(56) // int32 | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. (optional)
	poolName := "poolName_example" // string | The name of the pool. (optional)
	nodeName := "nodeName_example" // string | The name of the node. (optional)
	gslbserverId := int32(56) // int32 | The database identifier (ID) of the GSLB server associated with the application, a unique numeric key value automatically incremented when you add the server. Use it to identify the GSLB server of your choice. (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppNodeDelete(context.Background()).ApplicationFqdn(applicationFqdn).ApplicationId(applicationId).ApplicationName(applicationName).NodeId(nodeId).PoolId(poolId).PoolName(poolName).NodeName(nodeName).GslbserverId(gslbserverId).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppNodeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppNodeDelete`: AppNodeDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppNodeDelete`: %v\n", resp)
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
 **gslbserverId** | **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique numeric key value automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**AppNodeDeleteSuccess**](AppNodeDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	appNodeEditInput := *openapiclient.NewAppNodeEditInput() // AppNodeEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppNodeEdit(context.Background()).AppNodeEditInput(appNodeEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppNodeEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppNodeEdit`: AppNodeEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppNodeEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appNodeEditInput** | [**AppNodeEditInput**](AppNodeEditInput.md) |  | 

### Return type

[**AppNodeEditSuccess**](AppNodeEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	nodeId := int32(56) // int32 | The database identifier (ID) of the node, a unique numeric key value automatically incremented when you add a node. Use the ID to specify the node of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppNodeInfo(context.Background()).NodeId(nodeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppNodeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppNodeInfo`: AppNodeData
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppNodeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **nodeId** | **int32** | The database identifier (ID) of the node, a unique numeric key value automatically incremented when you add a node. Use the ID to specify the node of your choice. | 

### Return type

[**AppNodeData**](AppNodeData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppNodeList

> AppNodeData AppNodeList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Execute()

List the nodes



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	orderby := "orderby_example" // string | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
	select_ := "select__example" // string | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service <object>/list.If you specify several parameters they must be separated by a comma as follows: <b>select=&lt;param1&gt;<b>,</b>&lt;param2&gt;<b>,</b>...</b> .If the call includes the clause <b>where</b>, all the parameters it contains must be specified in the statement <b>select</b>.If the call includes the clause <b>orderby</b>, all the parameters it contains must be specified in the statement <b>select</b>.To include class parameters in the statement, you must tag them first. (optional)
	where := "where_example" // string | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first<br/>.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause is case insensitive and must be encoded in URL format. (optional)
	limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
	offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppNodeList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppNodeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppNodeList`: AppNodeData
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppNodeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppNodeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **select_** | **string** | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service &lt;object&gt;/list.If you specify several parameters they must be separated by a comma as follows: &lt;b&gt;select&#x3D;&amp;lt;param1&amp;gt;&lt;b&gt;,&lt;/b&gt;&amp;lt;param2&amp;gt;&lt;b&gt;,&lt;/b&gt;...&lt;/b&gt; .If the call includes the clause &lt;b&gt;where&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.If the call includes the clause &lt;b&gt;orderby&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.To include class parameters in the statement, you must tag them first. | 
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first&lt;br/&gt;.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause is case insensitive and must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 

### Return type

[**AppNodeData**](AppNodeData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	appPoolAddInput := *openapiclient.NewAppPoolAddInput() // AppPoolAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppPoolAdd(context.Background()).AppPoolAddInput(appPoolAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppPoolAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPoolAdd`: AppPoolAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppPoolAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appPoolAddInput** | [**AppPoolAddInput**](AppPoolAddInput.md) |  | 

### Return type

[**AppPoolAddSuccess**](AppPoolAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	where := "where_example" // string | A clause that allows you to filter the result. You can include any output parameter of the service <object>/list of the object in this clause.To filter the result using class parameters, you must tag them first. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppPoolCount(context.Background()).Where(where).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppPoolCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPoolCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppPoolCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service &lt;object&gt;/list of the object in this clause.To filter the result using class parameters, you must tag them first. | 

### Return type

[**ApiCountResponseSuccess**](ApiCountResponseSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	applicationFqdn := "applicationFqdn_example" // string | The Fully Qualified Domain Name (FQDN) of the application the object belongs to. (optional)
	applicationId := int32(56) // int32 | The database identifier (ID) of the application the object belongs to. (optional)
	applicationName := "applicationName_example" // string | The name of the application the object belongs to. (optional)
	poolId := int32(56) // int32 | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. (optional)
	poolName := "poolName_example" // string | The name of the pool. (optional)
	gslbserverId := int32(56) // int32 | The database identifier (ID) of the GSLB server associated with the application, a unique numeric key value automatically incremented when you add the server. Use it to identify the GSLB server of your choice. (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppPoolDelete(context.Background()).ApplicationFqdn(applicationFqdn).ApplicationId(applicationId).ApplicationName(applicationName).PoolId(poolId).PoolName(poolName).GslbserverId(gslbserverId).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppPoolDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPoolDelete`: AppPoolDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppPoolDelete`: %v\n", resp)
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
 **gslbserverId** | **int32** | The database identifier (ID) of the GSLB server associated with the application, a unique numeric key value automatically incremented when you add the server. Use it to identify the GSLB server of your choice. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**AppPoolDeleteSuccess**](AppPoolDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	appPoolEditInput := *openapiclient.NewAppPoolEditInput() // AppPoolEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppPoolEdit(context.Background()).AppPoolEditInput(appPoolEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppPoolEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPoolEdit`: AppPoolEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppPoolEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **appPoolEditInput** | [**AppPoolEditInput**](AppPoolEditInput.md) |  | 

### Return type

[**AppPoolEditSuccess**](AppPoolEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

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
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	poolId := int32(56) // int32 | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppPoolInfo(context.Background()).PoolId(poolId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppPoolInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPoolInfo`: AppPoolData
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppPoolInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **poolId** | **int32** | The database identifier (ID) of the pool, a unique numeric key value automatically incremented when you add a pool. Use the ID to specify the pool of your choice. | 

### Return type

[**AppPoolData**](AppPoolData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AppPoolList

> AppPoolData AppPoolList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Execute()

List the pools



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
	orderby := "orderby_example" // string | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
	select_ := "select__example" // string | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service <object>/list.If you specify several parameters they must be separated by a comma as follows: <b>select=&lt;param1&gt;<b>,</b>&lt;param2&gt;<b>,</b>...</b> .If the call includes the clause <b>where</b>, all the parameters it contains must be specified in the statement <b>select</b>.If the call includes the clause <b>orderby</b>, all the parameters it contains must be specified in the statement <b>select</b>.To include class parameters in the statement, you must tag them first. (optional)
	where := "where_example" // string | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first<br/>.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause is case insensitive and must be encoded in URL format. (optional)
	limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
	offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppAPI.AppPoolList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppAPI.AppPoolList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppPoolList`: AppPoolData
	fmt.Fprintf(os.Stdout, "Response from `AppAPI.AppPoolList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppPoolListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **select_** | **string** | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service &lt;object&gt;/list.If you specify several parameters they must be separated by a comma as follows: &lt;b&gt;select&#x3D;&amp;lt;param1&amp;gt;&lt;b&gt;,&lt;/b&gt;&amp;lt;param2&amp;gt;&lt;b&gt;,&lt;/b&gt;...&lt;/b&gt; .If the call includes the clause &lt;b&gt;where&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.If the call includes the clause &lt;b&gt;orderby&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.To include class parameters in the statement, you must tag them first. | 
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first&lt;br/&gt;.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause is case insensitive and must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 

### Return type

[**AppPoolData**](AppPoolData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

