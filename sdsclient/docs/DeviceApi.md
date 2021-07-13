# \DeviceApi

All URIs are relative to *https://vmdev-ct-fb11:443/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeviceDeviceAdd**](DeviceApi.md#DeviceDeviceAdd) | **Post** /device/device/add | Add a Device Manager device
[**DeviceDeviceCount**](DeviceApi.md#DeviceDeviceCount) | **Get** /device/device/count | Count the number of Device Manager devices
[**DeviceDeviceDelete**](DeviceApi.md#DeviceDeviceDelete) | **Delete** /device/device/delete | Delete a Device Manager device
[**DeviceDeviceEdit**](DeviceApi.md#DeviceDeviceEdit) | **Put** /device/device/edit | Add a Device Manager device
[**DeviceDeviceInfo**](DeviceApi.md#DeviceDeviceInfo) | **Get** /device/device/info | Display the properties of a Device Manager device
[**DeviceDeviceList**](DeviceApi.md#DeviceDeviceList) | **Get** /device/device/list | List the Device Manager devices
[**DeviceInterfaceAdd**](DeviceApi.md#DeviceInterfaceAdd) | **Post** /device/interface/add | Add a Device Manager port or interface
[**DeviceInterfaceCount**](DeviceApi.md#DeviceInterfaceCount) | **Get** /device/interface/count | Count the number Device Manager ports &amp;amp; interfaces
[**DeviceInterfaceDelete**](DeviceApi.md#DeviceInterfaceDelete) | **Delete** /device/interface/delete | Delete a Device Manager port or interface
[**DeviceInterfaceEdit**](DeviceApi.md#DeviceInterfaceEdit) | **Put** /device/interface/edit | Add a Device Manager port or interface
[**DeviceInterfaceInfo**](DeviceApi.md#DeviceInterfaceInfo) | **Get** /device/interface/info | Display the properties of a Device Manager port or interface
[**DeviceInterfaceList**](DeviceApi.md#DeviceInterfaceList) | **Get** /device/interface/list | List the Device Manager ports &amp;amp; interfaces
[**DeviceLinkAdd**](DeviceApi.md#DeviceLinkAdd) | **Post** /device/link/add | Link two Device Manager devices using their ports and/or interfaces
[**DeviceLinkCount**](DeviceApi.md#DeviceLinkCount) | **Get** /device/link/count | Count the number of links between Device Manager devices
[**DeviceLinkDelete**](DeviceApi.md#DeviceLinkDelete) | **Delete** /device/link/delete | Delete a link between two Device Manager devices
[**DeviceLinkEdit**](DeviceApi.md#DeviceLinkEdit) | **Put** /device/link/edit | Link two Device Manager devices using their ports and/or interfaces
[**DeviceLinkList**](DeviceApi.md#DeviceLinkList) | **Get** /device/link/list | List Device Manager ports &amp;amp; interfaces



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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceDeviceAddInput := *sdsclient.NewDeviceDeviceAddInput() // DeviceDeviceAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceDeviceAdd(context.Background()).DeviceDeviceAddInput(deviceDeviceAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceDeviceAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceDeviceAdd`: DeviceDeviceAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceDeviceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceDeviceAddInput** | [**DeviceDeviceAddInput**](DeviceDeviceAddInput.md) |  | 

### Return type

[**DeviceDeviceAddSuccess**](device_device_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceDeviceCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceDeviceCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceDeviceCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceDeviceCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceCountRequest struct via the builder pattern


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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceId := int32(56) // int32 | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. (optional)
    deviceName := "deviceName_example" // string | The name of the Device Manager device. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceDeviceDelete(context.Background()).DeviceId(deviceId).DeviceName(deviceName).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceDeviceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceDeviceDelete`: DeviceDeviceDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceDeviceDelete`: %v\n", resp)
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

[**DeviceDeviceDeleteSuccess**](device_device_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceDeviceEditInput := *sdsclient.NewDeviceDeviceEditInput() // DeviceDeviceEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceDeviceEdit(context.Background()).DeviceDeviceEditInput(deviceDeviceEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceDeviceEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceDeviceEdit`: DeviceDeviceEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceDeviceEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceDeviceEditInput** | [**DeviceDeviceEditInput**](DeviceDeviceEditInput.md) |  | 

### Return type

[**DeviceDeviceEditSuccess**](device_device_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceId := int32(56) // int32 | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceDeviceInfo(context.Background()).DeviceId(deviceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceDeviceInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceDeviceInfo`: DeviceDeviceData
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceDeviceInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceId** | **int32** | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. | 

### Return type

[**DeviceDeviceData**](device_device_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceDeviceList

> DeviceDeviceData DeviceDeviceList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the Device Manager devices



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
    resp, r, err := api_client.DeviceApi.DeviceDeviceList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceDeviceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceDeviceList`: DeviceDeviceData
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceDeviceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceDeviceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**DeviceDeviceData**](device_device_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceInterfaceAddInput := *sdsclient.NewDeviceInterfaceAddInput() // DeviceInterfaceAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceInterfaceAdd(context.Background()).DeviceInterfaceAddInput(deviceInterfaceAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceInterfaceAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceInterfaceAdd`: DeviceInterfaceAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceInterfaceAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceInterfaceAddInput** | [**DeviceInterfaceAddInput**](DeviceInterfaceAddInput.md) |  | 

### Return type

[**DeviceInterfaceAddSuccess**](device_interface_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceInterfaceCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceInterfaceCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceInterfaceCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceInterfaceCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceCountRequest struct via the builder pattern


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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceId := int32(56) // int32 | The database identifier (ID) of the Device Manager device, a unique numeric key value automatically incremented when you add a device. Use the ID to specify the device of your choice. (optional)
    deviceName := "deviceName_example" // string | The name of the Device Manager device. (optional)
    interfaceId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. (optional)
    interfaceName := "interfaceName_example" // string | The name of the Device Manager port or interface. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceInterfaceDelete(context.Background()).DeviceId(deviceId).DeviceName(deviceName).InterfaceId(interfaceId).InterfaceName(interfaceName).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceInterfaceDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceInterfaceDelete`: DeviceInterfaceDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceInterfaceDelete`: %v\n", resp)
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

[**DeviceInterfaceDeleteSuccess**](device_interface_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceInterfaceEditInput := *sdsclient.NewDeviceInterfaceEditInput() // DeviceInterfaceEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceInterfaceEdit(context.Background()).DeviceInterfaceEditInput(deviceInterfaceEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceInterfaceEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceInterfaceEdit`: DeviceInterfaceEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceInterfaceEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceInterfaceEditInput** | [**DeviceInterfaceEditInput**](DeviceInterfaceEditInput.md) |  | 

### Return type

[**DeviceInterfaceEditSuccess**](device_interface_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceInterfaceInfo

> DeviceInterfaceData DeviceInterfaceInfo(ctx).InterfaceId(interfaceId).Execute()

Display the properties of a Device Manager port or interface



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
    interfaceId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceInterfaceInfo(context.Background()).InterfaceId(interfaceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceInterfaceInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceInterfaceInfo`: DeviceInterfaceData
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceInterfaceInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **interfaceId** | **int32** | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. | 

### Return type

[**DeviceInterfaceData**](device_interface_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceInterfaceList

> DeviceInterfaceData DeviceInterfaceList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the Device Manager ports &amp; interfaces



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
    resp, r, err := api_client.DeviceApi.DeviceInterfaceList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceInterfaceList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceInterfaceList`: DeviceInterfaceData
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceInterfaceList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceInterfaceListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**DeviceInterfaceData**](device_interface_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceLinkAddInput := *sdsclient.NewDeviceLinkAddInput() // DeviceLinkAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceLinkAdd(context.Background()).DeviceLinkAddInput(deviceLinkAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceLinkAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceLinkAdd`: DeviceLinkAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceLinkAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceLinkAddInput** | [**DeviceLinkAddInput**](DeviceLinkAddInput.md) |  | 

### Return type

[**DeviceLinkAddSuccess**](device_link_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    interfaceId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceLinkCount(context.Background()).Where(where).InterfaceId(interfaceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceLinkCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceLinkCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceLinkCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service &lt;b&gt;*_list&lt;/b&gt; of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is no longer possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **interfaceId** | **int32** | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. | 

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    device1Name := "device1Name_example" // string | The name of the device to which belongs the DM port or interface you want to unlink from <b>hostiface2_id</b>. (optional)
    device2Name := "device2Name_example" // string | The name of the device to which belongs the DM port or interface you want to unlink from <b>hostiface1_id</b>. (optional)
    interface1Id := int32(56) // int32 | The database identifier (ID) of the DM port or interface you want to unlink from <b>hostiface2_id</b>. (optional)
    interface1Name := "interface1Name_example" // string | The name of the DM port or interface you want to unlink from <b>hostiface2_id</b>. (optional)
    interface2Id := int32(56) // int32 | The database identifier (ID) of the DM port or interface you want to unlink from <b>hostiface1_id</b>. (optional)
    interface2Name := "interface2Name_example" // string | The name of the DM port or interface you want to unlink from <b>hostiface1_id</b>. (optional)
    linkId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface link, a unique numeric key value automatically incremented when you add a link between a device and a port or interface. Use the ID to specify the port or interface link of your choice. (optional)
    linkAutoLink := int32(56) // int32 | A way to determine if the link between two Device Manager devices is set automatically (<b>1</b>) or manually (<b>0</b>). (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceLinkDelete(context.Background()).Device1Name(device1Name).Device2Name(device2Name).Interface1Id(interface1Id).Interface1Name(interface1Name).Interface2Id(interface2Id).Interface2Name(interface2Name).LinkId(linkId).LinkAutoLink(linkAutoLink).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceLinkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceLinkDelete`: DeviceLinkDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceLinkDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device1Name** | **string** | The name of the device to which belongs the DM port or interface you want to unlink from &lt;b&gt;hostiface2_id&lt;/b&gt;. | 
 **device2Name** | **string** | The name of the device to which belongs the DM port or interface you want to unlink from &lt;b&gt;hostiface1_id&lt;/b&gt;. | 
 **interface1Id** | **int32** | The database identifier (ID) of the DM port or interface you want to unlink from &lt;b&gt;hostiface2_id&lt;/b&gt;. | 
 **interface1Name** | **string** | The name of the DM port or interface you want to unlink from &lt;b&gt;hostiface2_id&lt;/b&gt;. | 
 **interface2Id** | **int32** | The database identifier (ID) of the DM port or interface you want to unlink from &lt;b&gt;hostiface1_id&lt;/b&gt;. | 
 **interface2Name** | **string** | The name of the DM port or interface you want to unlink from &lt;b&gt;hostiface1_id&lt;/b&gt;. | 
 **linkId** | **int32** | The database identifier (ID) of the Device Manager port or interface link, a unique numeric key value automatically incremented when you add a link between a device and a port or interface. Use the ID to specify the port or interface link of your choice. | 
 **linkAutoLink** | **int32** | A way to determine if the link between two Device Manager devices is set automatically (&lt;b&gt;1&lt;/b&gt;) or manually (&lt;b&gt;0&lt;/b&gt;). | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**DeviceLinkDeleteSuccess**](device_link_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    "github.com/efficientip-labs/solidserver-go-client/sdsclient"
)

func main() {
    deviceLinkEditInput := *sdsclient.NewDeviceLinkEditInput() // DeviceLinkEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceLinkEdit(context.Background()).DeviceLinkEditInput(deviceLinkEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceLinkEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceLinkEdit`: DeviceLinkEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceLinkEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deviceLinkEditInput** | [**DeviceLinkEditInput**](DeviceLinkEditInput.md) |  | 

### Return type

[**DeviceLinkEditSuccess**](device_link_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeviceLinkList

> DeviceLinkData DeviceLinkList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).InterfaceId(interfaceId).Execute()

List Device Manager ports &amp; interfaces



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
    interfaceId := int32(56) // int32 | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceApi.DeviceLinkList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).InterfaceId(interfaceId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceApi.DeviceLinkList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeviceLinkList`: DeviceLinkData
    fmt.Fprintf(os.Stdout, "Response from `DeviceApi.DeviceLinkList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeviceLinkListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **interfaceId** | **int32** | The database identifier (ID) of the Device Manager port or interface, a unique numeric key value automatically incremented when you add a port or interface. Use the ID to specify the port or interface of your choice. | 

### Return type

[**DeviceLinkData**](device_link_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

