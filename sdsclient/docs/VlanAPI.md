# \VlanAPI

All URIs are relative to *https://sds-ip-or-name:443/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VlanDomainAdd**](VlanAPI.md#VlanDomainAdd) | **Post** /vlan/domain/add | Add a VLAN domain
[**VlanDomainCount**](VlanAPI.md#VlanDomainCount) | **Get** /vlan/domain/count | Count the number of VLAN domains
[**VlanDomainDelete**](VlanAPI.md#VlanDomainDelete) | **Delete** /vlan/domain/delete | Delete a VLAN domain
[**VlanDomainEdit**](VlanAPI.md#VlanDomainEdit) | **Put** /vlan/domain/edit | Edit a VLAN domain
[**VlanDomainInfo**](VlanAPI.md#VlanDomainInfo) | **Get** /vlan/domain/info | Display the properties of a VLAN domain
[**VlanDomainList**](VlanAPI.md#VlanDomainList) | **Get** /vlan/domain/list | List the VLAN domains
[**VlanRangeAdd**](VlanAPI.md#VlanRangeAdd) | **Post** /vlan/range/add | Add a VLAN range
[**VlanRangeCount**](VlanAPI.md#VlanRangeCount) | **Get** /vlan/range/count | Count the number of VLAN ranges
[**VlanRangeDelete**](VlanAPI.md#VlanRangeDelete) | **Delete** /vlan/range/delete | Delete a VLAN range
[**VlanRangeEdit**](VlanAPI.md#VlanRangeEdit) | **Put** /vlan/range/edit | Edit a VLAN range
[**VlanRangeInfo**](VlanAPI.md#VlanRangeInfo) | **Get** /vlan/range/info | Display the properties of a VLAN range
[**VlanRangeList**](VlanAPI.md#VlanRangeList) | **Get** /vlan/range/list | List the VLAN ranges
[**VlanVlanAdd**](VlanAPI.md#VlanVlanAdd) | **Post** /vlan/vlan/add | Add a VLAN
[**VlanVlanCount**](VlanAPI.md#VlanVlanCount) | **Get** /vlan/vlan/count | Count the number of VLANs
[**VlanVlanDelete**](VlanAPI.md#VlanVlanDelete) | **Delete** /vlan/vlan/delete | Delete a VLAN
[**VlanVlanEdit**](VlanAPI.md#VlanVlanEdit) | **Put** /vlan/vlan/edit | Edit a VLAN
[**VlanVlanInfo**](VlanAPI.md#VlanVlanInfo) | **Get** /vlan/vlan/info | Display the properties of a VLAN
[**VlanVlanList**](VlanAPI.md#VlanVlanList) | **Get** /vlan/vlan/list | List the VLANs



## VlanDomainAdd

> VlanDomainAddSuccess VlanDomainAdd(ctx).VlanDomainAddInput(vlanDomainAddInput).Execute()

Add a VLAN domain



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
	vlanDomainAddInput := *openapiclient.NewVlanDomainAddInput() // VlanDomainAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanDomainAdd(context.Background()).VlanDomainAddInput(vlanDomainAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanDomainAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanDomainAdd`: VlanDomainAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanDomainAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanDomainAddInput** | [**VlanDomainAddInput**](VlanDomainAddInput.md) |  | 

### Return type

[**VlanDomainAddSuccess**](VlanDomainAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanDomainCount

> ApiCountResponseSuccess VlanDomainCount(ctx).Where(where).Tags(tags).Execute()

Count the number of VLAN domains



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
	resp, r, err := apiClient.VlanAPI.VlanDomainCount(context.Background()).Where(where).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanDomainCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanDomainCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanDomainCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainCountRequest struct via the builder pattern


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


## VlanDomainDelete

> VlanDomainDeleteSuccess VlanDomainDelete(ctx).DomainId(domainId).DomainName(domainName).Warnings(warnings).Execute()

Delete a VLAN domain



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
	domainId := int32(56) // int32 | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. (optional)
	domainName := "domainName_example" // string | The name of the VLAN domain. (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanDomainDelete(context.Background()).DomainId(domainId).DomainName(domainName).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanDomainDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanDomainDelete`: VlanDomainDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanDomainDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | 
 **domainName** | **string** | The name of the VLAN domain. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**VlanDomainDeleteSuccess**](VlanDomainDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanDomainEdit

> VlanDomainEditSuccess VlanDomainEdit(ctx).VlanDomainEditInput(vlanDomainEditInput).Execute()

Edit a VLAN domain



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
	vlanDomainEditInput := *openapiclient.NewVlanDomainEditInput() // VlanDomainEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanDomainEdit(context.Background()).VlanDomainEditInput(vlanDomainEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanDomainEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanDomainEdit`: VlanDomainEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanDomainEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanDomainEditInput** | [**VlanDomainEditInput**](VlanDomainEditInput.md) |  | 

### Return type

[**VlanDomainEditSuccess**](VlanDomainEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanDomainInfo

> VlanDomainData VlanDomainInfo(ctx).DomainId(domainId).Execute()

Display the properties of a VLAN domain



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
	domainId := int32(56) // int32 | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanDomainInfo(context.Background()).DomainId(domainId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanDomainInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanDomainInfo`: VlanDomainData
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanDomainInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | 

### Return type

[**VlanDomainData**](VlanDomainData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanDomainList

> VlanDomainData VlanDomainList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the VLAN domains



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
	resp, r, err := apiClient.VlanAPI.VlanDomainList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanDomainList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanDomainList`: VlanDomainData
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanDomainList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows you to sort the result. You can include any output parameter of the service in this clause.To sort the result using class parameters, you must tag them first.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **select_** | **string** | A statement that allows you to specify which column(s), i.e. parameter, is returned by the service.The statement can contain any output parameter of the service &lt;object&gt;/list.If you specify several parameters they must be separated by a comma as follows: &lt;b&gt;select&#x3D;&amp;lt;param1&amp;gt;&lt;b&gt;,&lt;/b&gt;&amp;lt;param2&amp;gt;&lt;b&gt;,&lt;/b&gt;...&lt;/b&gt; .If the call includes the clause &lt;b&gt;where&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.If the call includes the clause &lt;b&gt;orderby&lt;/b&gt;, all the parameters it contains must be specified in the statement &lt;b&gt;select&lt;/b&gt;.To include class parameters in the statement, you must tag them first. | 
 **where** | **string** | A clause that allows you to filter the result. You can include any output parameter of the service in this clause.To filter the result using class parameters, you must tag them first&lt;br/&gt;.The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause is case insensitive and must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**VlanDomainData**](VlanDomainData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanRangeAdd

> VlanRangeAddSuccess VlanRangeAdd(ctx).VlanRangeAddInput(vlanRangeAddInput).Execute()

Add a VLAN range



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
	vlanRangeAddInput := *openapiclient.NewVlanRangeAddInput() // VlanRangeAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanRangeAdd(context.Background()).VlanRangeAddInput(vlanRangeAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanRangeAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanRangeAdd`: VlanRangeAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanRangeAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanRangeAddInput** | [**VlanRangeAddInput**](VlanRangeAddInput.md) |  | 

### Return type

[**VlanRangeAddSuccess**](VlanRangeAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanRangeCount

> ApiCountResponseSuccess VlanRangeCount(ctx).Where(where).Tags(tags).Execute()

Count the number of VLAN ranges



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
	resp, r, err := apiClient.VlanAPI.VlanRangeCount(context.Background()).Where(where).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanRangeCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanRangeCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanRangeCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeCountRequest struct via the builder pattern


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


## VlanRangeDelete

> VlanRangeDeleteSuccess VlanRangeDelete(ctx).DomainId(domainId).DomainName(domainName).RangeId(rangeId).RangeName(rangeName).Warnings(warnings).Execute()

Delete a VLAN range



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
	domainId := int32(56) // int32 | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. (optional)
	domainName := "domainName_example" // string | The name of the VLAN domain. (optional)
	rangeId := int32(56) // int32 | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. (optional)
	rangeName := "rangeName_example" // string | The name of the VLAN range. (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanRangeDelete(context.Background()).DomainId(domainId).DomainName(domainName).RangeId(rangeId).RangeName(rangeName).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanRangeDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanRangeDelete`: VlanRangeDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanRangeDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | 
 **domainName** | **string** | The name of the VLAN domain. | 
 **rangeId** | **int32** | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. | 
 **rangeName** | **string** | The name of the VLAN range. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**VlanRangeDeleteSuccess**](VlanRangeDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanRangeEdit

> VlanRangeEditSuccess VlanRangeEdit(ctx).VlanRangeEditInput(vlanRangeEditInput).Execute()

Edit a VLAN range



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
	vlanRangeEditInput := *openapiclient.NewVlanRangeEditInput() // VlanRangeEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanRangeEdit(context.Background()).VlanRangeEditInput(vlanRangeEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanRangeEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanRangeEdit`: VlanRangeEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanRangeEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanRangeEditInput** | [**VlanRangeEditInput**](VlanRangeEditInput.md) |  | 

### Return type

[**VlanRangeEditSuccess**](VlanRangeEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanRangeInfo

> VlanRangeData VlanRangeInfo(ctx).NoParentClassParam(noParentClassParam).RangeId(rangeId).Execute()

Display the properties of a VLAN range



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
	rangeId := int32(56) // int32 | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanRangeInfo(context.Background()).NoParentClassParam(noParentClassParam).RangeId(rangeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanRangeInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanRangeInfo`: VlanRangeData
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanRangeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noParentClassParam** | **int32** | A way to exclude the class parameter details of the parent of the object in the output parameter dedicated to its class parameters. | 
 **rangeId** | **int32** | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. | 

### Return type

[**VlanRangeData**](VlanRangeData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanRangeList

> VlanRangeData VlanRangeList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).NoParentClassParam(noParentClassParam).Tags(tags).Execute()

List the VLAN ranges



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
	resp, r, err := apiClient.VlanAPI.VlanRangeList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).NoParentClassParam(noParentClassParam).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanRangeList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanRangeList`: VlanRangeData
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanRangeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeListRequest struct via the builder pattern


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

[**VlanRangeData**](VlanRangeData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanVlanAdd

> VlanVlanAddSuccess VlanVlanAdd(ctx).VlanVlanAddInput(vlanVlanAddInput).Execute()

Add a VLAN



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
	vlanVlanAddInput := *openapiclient.NewVlanVlanAddInput() // VlanVlanAddInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanVlanAdd(context.Background()).VlanVlanAddInput(vlanVlanAddInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanVlanAdd``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanVlanAdd`: VlanVlanAddSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanVlanAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanVlanAddInput** | [**VlanVlanAddInput**](VlanVlanAddInput.md) |  | 

### Return type

[**VlanVlanAddSuccess**](VlanVlanAddSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanVlanCount

> ApiCountResponseSuccess VlanVlanCount(ctx).Where(where).Tags(tags).Execute()

Count the number of VLANs



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
	resp, r, err := apiClient.VlanAPI.VlanVlanCount(context.Background()).Where(where).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanVlanCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanVlanCount`: ApiCountResponseSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanVlanCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanCountRequest struct via the builder pattern


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


## VlanVlanDelete

> VlanVlanDeleteSuccess VlanVlanDelete(ctx).DomainId(domainId).DomainName(domainName).VlanId(vlanId).VlanVlanId(vlanVlanId).RangeId(rangeId).RangeName(rangeName).VlanName(vlanName).Warnings(warnings).Execute()

Delete a VLAN



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
	domainId := int32(56) // int32 | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. (optional)
	domainName := "domainName_example" // string | The name of the VLAN domain. (optional)
	vlanId := int32(56) // int32 | The database identifier (ID) of the VLAN, a unique numeric key value automatically incremented when you add a VLAN. Use the ID to specify the VLAN of your choice. (optional)
	vlanVlanId := int32(56) // int32 | The VLAN identifier (ID) of the VLAN, a unique numeric key value within a VLAN domain. Use the ID to specify the VLAN of your choice. (optional)
	rangeId := int32(56) // int32 | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. (optional)
	rangeName := "rangeName_example" // string | The name of the VLAN range. (optional)
	vlanName := "vlanName_example" // string | The name of the VLAN. (optional)
	warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanVlanDelete(context.Background()).DomainId(domainId).DomainName(domainName).VlanId(vlanId).VlanVlanId(vlanVlanId).RangeId(rangeId).RangeName(rangeName).VlanName(vlanName).Warnings(warnings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanVlanDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanVlanDelete`: VlanVlanDeleteSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanVlanDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | 
 **domainName** | **string** | The name of the VLAN domain. | 
 **vlanId** | **int32** | The database identifier (ID) of the VLAN, a unique numeric key value automatically incremented when you add a VLAN. Use the ID to specify the VLAN of your choice. | 
 **vlanVlanId** | **int32** | The VLAN identifier (ID) of the VLAN, a unique numeric key value within a VLAN domain. Use the ID to specify the VLAN of your choice. | 
 **rangeId** | **int32** | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. | 
 **rangeName** | **string** | The name of the VLAN range. | 
 **vlanName** | **string** | The name of the VLAN. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**VlanVlanDeleteSuccess**](VlanVlanDeleteSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanVlanEdit

> VlanVlanEditSuccess VlanVlanEdit(ctx).VlanVlanEditInput(vlanVlanEditInput).Execute()

Edit a VLAN



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
	vlanVlanEditInput := *openapiclient.NewVlanVlanEditInput() // VlanVlanEditInput |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanVlanEdit(context.Background()).VlanVlanEditInput(vlanVlanEditInput).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanVlanEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanVlanEdit`: VlanVlanEditSuccess
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanVlanEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanVlanEditInput** | [**VlanVlanEditInput**](VlanVlanEditInput.md) |  | 

### Return type

[**VlanVlanEditSuccess**](VlanVlanEditSuccess.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanVlanInfo

> VlanVlanData VlanVlanInfo(ctx).NoParentClassParam(noParentClassParam).VlanId(vlanId).Execute()

Display the properties of a VLAN



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
	vlanId := int32(56) // int32 | The database identifier (ID) of the VLAN, a unique numeric key value automatically incremented when you add a VLAN. Use the ID to specify the VLAN of your choice. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.VlanAPI.VlanVlanInfo(context.Background()).NoParentClassParam(noParentClassParam).VlanId(vlanId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanVlanInfo``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanVlanInfo`: VlanVlanData
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanVlanInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **noParentClassParam** | **int32** | A way to exclude the class parameter details of the parent of the object in the output parameter dedicated to its class parameters. | 
 **vlanId** | **int32** | The database identifier (ID) of the VLAN, a unique numeric key value automatically incremented when you add a VLAN. Use the ID to specify the VLAN of your choice. | 

### Return type

[**VlanVlanData**](VlanVlanData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanVlanList

> VlanVlanData VlanVlanList(ctx).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).NoParentClassParam(noParentClassParam).Tags(tags).Execute()

List the VLANs



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
	resp, r, err := apiClient.VlanAPI.VlanVlanList(context.Background()).Orderby(orderby).Select_(select_).Where(where).Limit(limit).Offset(offset).NoParentClassParam(noParentClassParam).Tags(tags).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `VlanAPI.VlanVlanList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `VlanVlanList`: VlanVlanData
	fmt.Fprintf(os.Stdout, "Response from `VlanAPI.VlanVlanList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanListRequest struct via the builder pattern


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

[**VlanVlanData**](VlanVlanData.md)

### Authorization

[HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey), [BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [EipApiTokenAuth](../README.md#EipApiTokenAuth), [EipApiTokenTime](../README.md#EipApiTokenTime)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

