# \DeviceAPI

All URIs are relative to *https://sds-ip-or-name:443/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceDeviceAdd**](DeviceAPI.md#DeviceDeviceAdd) | **Post** /device/device/add | Add a Device Manager device
[**DeviceDeviceCount**](DeviceAPI.md#DeviceDeviceCount) | **Get** /device/device/count | Count the number of Device Manager devices
[**DeviceDeviceDelete**](DeviceAPI.md#DeviceDeviceDelete) | **Delete** /device/device/delete | Delete a Device Manager device
[**DeviceDeviceEdit**](DeviceAPI.md#DeviceDeviceEdit) | **Put** /device/device/edit | Add a Device Manager device
[**DeviceDeviceInfo**](DeviceAPI.md#DeviceDeviceInfo) | **Get** /device/device/info | Display the properties of a Device Manager device
[**DeviceDeviceList**](DeviceAPI.md#DeviceDeviceList) | **Get** /device/device/list | List the Device Manager devices
[**DeviceInterfaceAdd**](DeviceAPI.md#DeviceInterfaceAdd) | **Post** /device/interface/add | Add a Device Manager port or interface
[**DeviceInterfaceCount**](DeviceAPI.md#DeviceInterfaceCount) | **Get** /device/interface/count | Count the number Device Manager ports &amp;amp; interfaces
[**DeviceInterfaceDelete**](DeviceAPI.md#DeviceInterfaceDelete) | **Delete** /device/interface/delete | Delete a Device Manager port or interface
[**DeviceInterfaceEdit**](DeviceAPI.md#DeviceInterfaceEdit) | **Put** /device/interface/edit | Add a Device Manager port or interface
[**DeviceInterfaceInfo**](DeviceAPI.md#DeviceInterfaceInfo) | **Get** /device/interface/info | Display the properties of a Device Manager port or interface
[**DeviceInterfaceList**](DeviceAPI.md#DeviceInterfaceList) | **Get** /device/interface/list | List the Device Manager ports &amp;amp; interfaces
[**DeviceLinkAdd**](DeviceAPI.md#DeviceLinkAdd) | **Post** /device/link/add | Link two Device Manager devices using their ports and/or interfaces
[**DeviceLinkCount**](DeviceAPI.md#DeviceLinkCount) | **Get** /device/link/count | Count the number of links between Device Manager devices
[**DeviceLinkDelete**](DeviceAPI.md#DeviceLinkDelete) | **Delete** /device/link/delete | Delete a link between two Device Manager devices
[**DeviceLinkEdit**](DeviceAPI.md#DeviceLinkEdit) | **Put** /device/link/edit | Link two Device Manager devices using their ports and/or interfaces
[**DeviceLinkList**](DeviceAPI.md#DeviceLinkList) | **Get** /device/link/list | List Device Manager ports &amp;amp; interfaces



## DeviceDeviceAdd

> DeviceDeviceAddSuccess DeviceDeviceAdd(ctx).DeviceDeviceAddInput(deviceDeviceAddInput).Execute()

Add a Device Manager device



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
	deviceDeviceAddInput := *openapiclient.NewDeviceDeviceAddInput() // DeviceDeviceAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceDeviceAdd(context.Background()).DeviceDeviceAddInput(deviceDeviceAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceDeviceAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceDeviceAdd`: DeviceDeviceAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceDeviceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceDeviceAddInput** | [**DeviceDeviceAddInput**](DeviceDeviceAddInput.md) |  | 

### Return type

[**DeviceDeviceAddSuccess**](DeviceDeviceAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceDeviceCount

> ApiCountResponseSuccess DeviceDeviceCount(ctx).Where(where).Tags(tags).Execute()

Count the number of Device Manager devices



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
	resp, r, err := apiClient.DeviceAPI.DeviceDeviceCount(context.Background()).Where(where).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceDeviceCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceDeviceCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceDeviceCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceCountRequest struct via the builder pattern


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


## DeviceDeviceDelete

> DeviceDeviceDeleteSuccess DeviceDeviceDelete(ctx).DeviceId(deviceId).DeviceName(deviceName).Warnings(warnings).Execute()

Delete a Device Manager device



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
	deviceId := int32(56) // int32 | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. (optional)
	deviceName := "deviceName_example" // string | The name of the Device Manager device. (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceDeviceDelete(context.Background()).DeviceId(deviceId).DeviceName(deviceName).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceDeviceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceDeviceDelete`: DeviceDeviceDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceDeviceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **int32** | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. | 
 **deviceName** | **string** | The name of the Device Manager device. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**DeviceDeviceDeleteSuccess**](DeviceDeviceDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceDeviceEdit

> DeviceDeviceEditSuccess DeviceDeviceEdit(ctx).DeviceDeviceEditInput(deviceDeviceEditInput).Execute()

Add a Device Manager device



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
	deviceDeviceEditInput := *openapiclient.NewDeviceDeviceEditInput() // DeviceDeviceEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceDeviceEdit(context.Background()).DeviceDeviceEditInput(deviceDeviceEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceDeviceEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceDeviceEdit`: DeviceDeviceEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceDeviceEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceDeviceEditInput** | [**DeviceDeviceEditInput**](DeviceDeviceEditInput.md) |  | 

### Return type

[**DeviceDeviceEditSuccess**](DeviceDeviceEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceDeviceInfo

> DeviceDeviceData DeviceDeviceInfo(ctx).DeviceId(deviceId).Execute()

Display the properties of a Device Manager device



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
	deviceId := int32(56) // int32 | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceDeviceInfo(context.Background()).DeviceId(deviceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceDeviceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceDeviceInfo`: DeviceDeviceData
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceDeviceInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **int32** | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. | 

### Return type

[**DeviceDeviceData**](DeviceDeviceData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceDeviceList

> DeviceDeviceData DeviceDeviceList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the Device Manager devices



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
	resp, r, err := apiClient.DeviceAPI.DeviceDeviceList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceDeviceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceDeviceList`: DeviceDeviceData
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceDeviceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **select_** | **string** | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service &lt;object&gt;/list.If you specify several parameters they must be separated by a comma as follows: &lt;b&gt;select&#x3D;&amp;lt;param1&amp;gt;&lt;b&gt;,&lt;/b&gt;&amp;lt;param2&amp;gt;&lt;b&gt;,&lt;/b&gt;...&lt;/b&gt; .If the call includes the clause &lt;b&gt;where&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.If the call includes the clause &lt;b&gt;orderby&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.To include class parameters in the statement, you must tag them first. | 
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first&lt;br/&gt;.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause is case insensitive and must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**DeviceDeviceData**](DeviceDeviceData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceInterfaceAdd

> DeviceInterfaceAddSuccess DeviceInterfaceAdd(ctx).DeviceInterfaceAddInput(deviceInterfaceAddInput).Execute()

Add a Device Manager port or interface



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
	deviceInterfaceAddInput := *openapiclient.NewDeviceInterfaceAddInput() // DeviceInterfaceAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceInterfaceAdd(context.Background()).DeviceInterfaceAddInput(deviceInterfaceAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceInterfaceAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInterfaceAdd`: DeviceInterfaceAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceInterfaceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceInterfaceAddInput** | [**DeviceInterfaceAddInput**](DeviceInterfaceAddInput.md) |  | 

### Return type

[**DeviceInterfaceAddSuccess**](DeviceInterfaceAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceInterfaceCount

> ApiCountResponseSuccess DeviceInterfaceCount(ctx).Where(where).Tags(tags).Execute()

Count the number Device Manager ports &amp; interfaces



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
	resp, r, err := apiClient.DeviceAPI.DeviceInterfaceCount(context.Background()).Where(where).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceInterfaceCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInterfaceCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceInterfaceCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceCountRequest struct via the builder pattern


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


## DeviceInterfaceDelete

> DeviceInterfaceDeleteSuccess DeviceInterfaceDelete(ctx).DeviceId(deviceId).DeviceName(deviceName).InterfaceId(interfaceId).InterfaceName(interfaceName).Warnings(warnings).Execute()

Delete a Device Manager port or interface



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
	deviceId := int32(56) // int32 | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. (optional)
	deviceName := "deviceName_example" // string | The name of the Device Manager device. (optional)
	interfaceId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. (optional)
	interfaceName := "interfaceName_example" // string | The name of the Device Manager port or interface. (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceInterfaceDelete(context.Background()).DeviceId(deviceId).DeviceName(deviceName).InterfaceId(interfaceId).InterfaceName(interfaceName).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceInterfaceDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInterfaceDelete`: DeviceInterfaceDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceInterfaceDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **int32** | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. | 
 **deviceName** | **string** | The name of the Device Manager device. | 
 **interfaceId** | **int32** | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. | 
 **interfaceName** | **string** | The name of the Device Manager port or interface. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**DeviceInterfaceDeleteSuccess**](DeviceInterfaceDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceInterfaceEdit

> DeviceInterfaceEditSuccess DeviceInterfaceEdit(ctx).DeviceInterfaceEditInput(deviceInterfaceEditInput).Execute()

Add a Device Manager port or interface



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
	deviceInterfaceEditInput := *openapiclient.NewDeviceInterfaceEditInput() // DeviceInterfaceEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceInterfaceEdit(context.Background()).DeviceInterfaceEditInput(deviceInterfaceEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceInterfaceEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInterfaceEdit`: DeviceInterfaceEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceInterfaceEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceInterfaceEditInput** | [**DeviceInterfaceEditInput**](DeviceInterfaceEditInput.md) |  | 

### Return type

[**DeviceInterfaceEditSuccess**](DeviceInterfaceEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceInterfaceInfo

> DeviceInterfaceData DeviceInterfaceInfo(ctx).NoParentClassParam(noParentClassParam).InterfaceId(interfaceId).Execute()

Display the properties of a Device Manager port or interface



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
	noParentClassParam := int32(56) // int32 | A way to exclude the class parameter details of the parent of the object in the output parameter dedicated to its class parameters. (optional)
	interfaceId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceInterfaceInfo(context.Background()).NoParentClassParam(noParentClassParam).InterfaceId(interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceInterfaceInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInterfaceInfo`: DeviceInterfaceData
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceInterfaceInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noParentClassParam** | **int32** | A way to exclude the class parameter details of the parent of the object in the output parameter dedicated to its class parameters. | 
 **interfaceId** | **int32** | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. | 

### Return type

[**DeviceInterfaceData**](DeviceInterfaceData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceInterfaceList

> DeviceInterfaceData DeviceInterfaceList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).NoParentClassParam(noParentClassParam).Tags(tags).Execute()

List the Device Manager ports &amp; interfaces



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
	noParentClassParam := int32(56) // int32 | A way to exclude the class parameter details of the parent of the object in the output parameter dedicated to its class parameters. (optional)
	tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceInterfaceList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).NoParentClassParam(noParentClassParam).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceInterfaceList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceInterfaceList`: DeviceInterfaceData
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceInterfaceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **select_** | **string** | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service &lt;object&gt;/list.If you specify several parameters they must be separated by a comma as follows: &lt;b&gt;select&#x3D;&amp;lt;param1&amp;gt;&lt;b&gt;,&lt;/b&gt;&amp;lt;param2&amp;gt;&lt;b&gt;,&lt;/b&gt;...&lt;/b&gt; .If the call includes the clause &lt;b&gt;where&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.If the call includes the clause &lt;b&gt;orderby&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.To include class parameters in the statement, you must tag them first. | 
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first&lt;br/&gt;.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause is case insensitive and must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **noParentClassParam** | **int32** | A way to exclude the class parameter details of the parent of the object in the output parameter dedicated to its class parameters. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**DeviceInterfaceData**](DeviceInterfaceData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLinkAdd

> DeviceLinkAddSuccess DeviceLinkAdd(ctx).DeviceLinkAddInput(deviceLinkAddInput).Execute()

Link two Device Manager devices using their ports and/or interfaces



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
	deviceLinkAddInput := *openapiclient.NewDeviceLinkAddInput() // DeviceLinkAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceLinkAdd(context.Background()).DeviceLinkAddInput(deviceLinkAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceLinkAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceLinkAdd`: DeviceLinkAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceLinkAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceLinkAddInput** | [**DeviceLinkAddInput**](DeviceLinkAddInput.md) |  | 

### Return type

[**DeviceLinkAddSuccess**](DeviceLinkAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLinkCount

> ApiCountResponseSuccess DeviceLinkCount(ctx).Where(where).InterfaceId(interfaceId).Execute()

Count the number of links between Device Manager devices



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
	interfaceId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceLinkCount(context.Background()).Where(where).InterfaceId(interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceLinkCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceLinkCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceLinkCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service &lt;object&gt;/list of the object in this clause.To filter the result using class parameters, you must tag them first. | 
 **interfaceId** | **int32** | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. | 

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


## DeviceLinkDelete

> DeviceLinkDeleteSuccess DeviceLinkDelete(ctx).Device1Name(device1Name).Device2Name(device2Name).Interface1Id(interface1Id).Interface1Name(interface1Name).Interface2Id(interface2Id).Interface2Name(interface2Name).LinkId(linkId).LinkAutoLink(linkAutoLink).Warnings(warnings).Execute()

Delete a link between two Device Manager devices



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
	device1Name := "device1Name_example" // string | The name of the device to which belongs the DM port or interface you want to unlink from <b>interface2_id</b>. (optional)
	device2Name := "device2Name_example" // string | The name of the device to which belongs the DM port or interface you want to unlink from <b>interface1_id</b>. (optional)
	interface1Id := int32(56) // int32 | The database identifier (ID) of the DM port or interface you want to unlink from <b>interface2_id</b>. (optional)
	interface1Name := "interface1Name_example" // string | The name of the DM port or interface you want to unlink from <b>interface2_id</b>. (optional)
	interface2Id := int32(56) // int32 | The database identifier (ID) of the DM port or interface you want to unlink from <b>interface1_id</b>. (optional)
	interface2Name := "interface2Name_example" // string | The name of the DM port or interface you want to unlink from <b>interface1_id</b>. (optional)
	linkId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface link, a unique numeric key value automatically incremented when you add a link between a device and a port or interface. Use the ID to specify the port or interface link of your choice. (optional)
	linkAutoLink := int32(56) // int32 | A way to determine if the link between two Device Manager devices is set automatically (<b>1</b>) or manually (<b>0</b>). (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceLinkDelete(context.Background()).Device1Name(device1Name).Device2Name(device2Name).Interface1Id(interface1Id).Interface1Name(interface1Name).Interface2Id(interface2Id).Interface2Name(interface2Name).LinkId(linkId).LinkAutoLink(linkAutoLink).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceLinkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceLinkDelete`: DeviceLinkDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceLinkDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device1Name** | **string** | The name of the device to which belongs the DM port or interface you want to unlink from &lt;b&gt;interface2_id&lt;/b&gt;. | 
 **device2Name** | **string** | The name of the device to which belongs the DM port or interface you want to unlink from &lt;b&gt;interface1_id&lt;/b&gt;. | 
 **interface1Id** | **int32** | The database identifier (ID) of the DM port or interface you want to unlink from &lt;b&gt;interface2_id&lt;/b&gt;. | 
 **interface1Name** | **string** | The name of the DM port or interface you want to unlink from &lt;b&gt;interface2_id&lt;/b&gt;. | 
 **interface2Id** | **int32** | The database identifier (ID) of the DM port or interface you want to unlink from &lt;b&gt;interface1_id&lt;/b&gt;. | 
 **interface2Name** | **string** | The name of the DM port or interface you want to unlink from &lt;b&gt;interface1_id&lt;/b&gt;. | 
 **linkId** | **int32** | The database identifier (ID) of the Device Manager port or interface link, a unique numeric key value automatically incremented when you add a link between a device and a port or interface. Use the ID to specify the port or interface link of your choice. | 
 **linkAutoLink** | **int32** | A way to determine if the link between two Device Manager devices is set automatically (&lt;b&gt;1&lt;/b&gt;) or manually (&lt;b&gt;0&lt;/b&gt;). | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**DeviceLinkDeleteSuccess**](DeviceLinkDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLinkEdit

> DeviceLinkEditSuccess DeviceLinkEdit(ctx).DeviceLinkEditInput(deviceLinkEditInput).Execute()

Link two Device Manager devices using their ports and/or interfaces



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
	deviceLinkEditInput := *openapiclient.NewDeviceLinkEditInput() // DeviceLinkEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceLinkEdit(context.Background()).DeviceLinkEditInput(deviceLinkEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceLinkEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceLinkEdit`: DeviceLinkEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceLinkEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceLinkEditInput** | [**DeviceLinkEditInput**](DeviceLinkEditInput.md) |  | 

### Return type

[**DeviceLinkEditSuccess**](DeviceLinkEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLinkList

> DeviceLinkData DeviceLinkList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).InterfaceId(interfaceId).Execute()

List Device Manager ports &amp; interfaces



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
	interfaceId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DeviceAPI.DeviceLinkList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).InterfaceId(interfaceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DeviceAPI.DeviceLinkList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeviceLinkList`: DeviceLinkData
	fmt.Fprintf(os.Stdout, "Response from `DeviceAPI.DeviceLinkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **select_** | **string** | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service &lt;object&gt;/list.If you specify several parameters they must be separated by a comma as follows: &lt;b&gt;select&#x3D;&amp;lt;param1&amp;gt;&lt;b&gt;,&lt;/b&gt;&amp;lt;param2&amp;gt;&lt;b&gt;,&lt;/b&gt;...&lt;/b&gt; .If the call includes the clause &lt;b&gt;where&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.If the call includes the clause &lt;b&gt;orderby&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.To include class parameters in the statement, you must tag them first. | 
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first&lt;br/&gt;.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause is case insensitive and must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **interfaceId** | **int32** | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. | 

### Return type

[**DeviceLinkData**](DeviceLinkData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

