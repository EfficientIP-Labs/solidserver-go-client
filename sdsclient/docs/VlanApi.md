# \VlanApi

All URIs are relative to *https://vmdev-ct-fb11:443/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**VlanDomainAdd**](VlanApi.md#VlanDomainAdd) | **Post** /vlan/domain/add | Add a VLAN domain
[**VlanDomainCount**](VlanApi.md#VlanDomainCount) | **Get** /vlan/domain/count | Count the number of VLAN domains
[**VlanDomainDelete**](VlanApi.md#VlanDomainDelete) | **Delete** /vlan/domain/delete | Delete a VLAN domain
[**VlanDomainEdit**](VlanApi.md#VlanDomainEdit) | **Put** /vlan/domain/edit | Edit a VLAN domain
[**VlanDomainInfo**](VlanApi.md#VlanDomainInfo) | **Get** /vlan/domain/info | Display the properties of a VLAN domain
[**VlanDomainList**](VlanApi.md#VlanDomainList) | **Get** /vlan/domain/list | List the VLAN domains
[**VlanRangeAdd**](VlanApi.md#VlanRangeAdd) | **Post** /vlan/range/add | Add a VLAN range
[**VlanRangeCount**](VlanApi.md#VlanRangeCount) | **Get** /vlan/range/count | Count the number of VLAN ranges
[**VlanRangeDelete**](VlanApi.md#VlanRangeDelete) | **Delete** /vlan/range/delete | Delete a VLAN range
[**VlanRangeEdit**](VlanApi.md#VlanRangeEdit) | **Put** /vlan/range/edit | Edit a VLAN range
[**VlanRangeInfo**](VlanApi.md#VlanRangeInfo) | **Get** /vlan/range/info | Display the properties of a VLAN range
[**VlanRangeList**](VlanApi.md#VlanRangeList) | **Get** /vlan/range/list | List the VLAN ranges
[**VlanVlanAdd**](VlanApi.md#VlanVlanAdd) | **Post** /vlan/vlan/add | Add a VLAN
[**VlanVlanCount**](VlanApi.md#VlanVlanCount) | **Get** /vlan/vlan/count | Count the number of VLANs
[**VlanVlanDelete**](VlanApi.md#VlanVlanDelete) | **Delete** /vlan/vlan/delete | Delete a VLAN
[**VlanVlanEdit**](VlanApi.md#VlanVlanEdit) | **Put** /vlan/vlan/edit | Edit a VLAN
[**VlanVlanInfo**](VlanApi.md#VlanVlanInfo) | **Get** /vlan/vlan/info | Display the properties of a VLAN
[**VlanVlanList**](VlanApi.md#VlanVlanList) | **Get** /vlan/vlan/list | List the VLANs



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
    openapiclient "./openapi"
)

func main() {
    vlanDomainAddInput := *openapiclient.NewVlanDomainAddInput() // VlanDomainAddInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanDomainAdd(context.Background()).VlanDomainAddInput(vlanDomainAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanDomainAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanDomainAdd`: VlanDomainAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanDomainAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanDomainAddInput** | [**VlanDomainAddInput**](VlanDomainAddInput.md) |  | 

### Return type

[**VlanDomainAddSuccess**](vlan_domain_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanDomainCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanDomainCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanDomainCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanDomainCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainCountRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    domainId := int32(56) // int32 | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. (optional)
    domainName := "domainName_example" // string | The name of the VLAN domain. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanDomainDelete(context.Background()).DomainId(domainId).DomainName(domainName).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanDomainDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanDomainDelete`: VlanDomainDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanDomainDelete`: %v\n", resp)
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

[**VlanDomainDeleteSuccess**](vlan_domain_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    vlanDomainEditInput := *openapiclient.NewVlanDomainEditInput() // VlanDomainEditInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanDomainEdit(context.Background()).VlanDomainEditInput(vlanDomainEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanDomainEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanDomainEdit`: VlanDomainEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanDomainEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanDomainEditInput** | [**VlanDomainEditInput**](VlanDomainEditInput.md) |  | 

### Return type

[**VlanDomainEditSuccess**](vlan_domain_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    domainId := int32(56) // int32 | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanDomainInfo(context.Background()).DomainId(domainId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanDomainInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanDomainInfo`: VlanDomainData
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanDomainInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **domainId** | **int32** | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. | 

### Return type

[**VlanDomainData**](vlan_domain_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanDomainList

> VlanDomainData VlanDomainList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the VLAN domains



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanDomainList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanDomainList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanDomainList`: VlanDomainData
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanDomainList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanDomainListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**VlanDomainData**](vlan_domain_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    vlanRangeAddInput := *openapiclient.NewVlanRangeAddInput() // VlanRangeAddInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanRangeAdd(context.Background()).VlanRangeAddInput(vlanRangeAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanRangeAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanRangeAdd`: VlanRangeAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanRangeAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanRangeAddInput** | [**VlanRangeAddInput**](VlanRangeAddInput.md) |  | 

### Return type

[**VlanRangeAddSuccess**](vlan_range_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanRangeCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanRangeCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanRangeCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanRangeCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeCountRequest struct via the builder pattern


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
    openapiclient "./openapi"
)

func main() {
    domainId := int32(56) // int32 | The database identifier (ID) of the VLAN domain, a unique numeric key value automatically incremented when you add a VLAN domain. Use the ID to specify the VLAN domain of your choice. (optional)
    domainName := "domainName_example" // string | The name of the VLAN domain. (optional)
    rangeId := int32(56) // int32 | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. (optional)
    rangeName := "rangeName_example" // string | The name of the VLAN range. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanRangeDelete(context.Background()).DomainId(domainId).DomainName(domainName).RangeId(rangeId).RangeName(rangeName).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanRangeDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanRangeDelete`: VlanRangeDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanRangeDelete`: %v\n", resp)
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

[**VlanRangeDeleteSuccess**](vlan_range_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    vlanRangeEditInput := *openapiclient.NewVlanRangeEditInput() // VlanRangeEditInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanRangeEdit(context.Background()).VlanRangeEditInput(vlanRangeEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanRangeEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanRangeEdit`: VlanRangeEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanRangeEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanRangeEditInput** | [**VlanRangeEditInput**](VlanRangeEditInput.md) |  | 

### Return type

[**VlanRangeEditSuccess**](vlan_range_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanRangeInfo

> VlanRangeData VlanRangeInfo(ctx).RangeId(rangeId).Execute()

Display the properties of a VLAN range



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    rangeId := int32(56) // int32 | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanRangeInfo(context.Background()).RangeId(rangeId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanRangeInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanRangeInfo`: VlanRangeData
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanRangeInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rangeId** | **int32** | The database identifier (ID) of the VLAN range, a unique numeric key value automatically incremented when you add a VLAN range. Use the ID to specify the VLAN range of your choice. | 

### Return type

[**VlanRangeData**](vlan_range_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanRangeList

> VlanRangeData VlanRangeList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the VLAN ranges



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanRangeList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanRangeList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanRangeList`: VlanRangeData
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanRangeList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanRangeListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**VlanRangeData**](vlan_range_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    vlanVlanAddInput := *openapiclient.NewVlanVlanAddInput() // VlanVlanAddInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanVlanAdd(context.Background()).VlanVlanAddInput(vlanVlanAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanVlanAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanVlanAdd`: VlanVlanAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanVlanAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanVlanAddInput** | [**VlanVlanAddInput**](VlanVlanAddInput.md) |  | 

### Return type

[**VlanVlanAddSuccess**](vlan_vlan_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service <b>*_list</b> of the object in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is no longer possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanVlanCount(context.Background()).Where(where).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanVlanCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanVlanCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanVlanCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanCountRequest struct via the builder pattern


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
    openapiclient "./openapi"
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
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanVlanDelete(context.Background()).DomainId(domainId).DomainName(domainName).VlanId(vlanId).VlanVlanId(vlanVlanId).RangeId(rangeId).RangeName(rangeName).VlanName(vlanName).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanVlanDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanVlanDelete`: VlanVlanDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanVlanDelete`: %v\n", resp)
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

[**VlanVlanDeleteSuccess**](vlan_vlan_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

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
    openapiclient "./openapi"
)

func main() {
    vlanVlanEditInput := *openapiclient.NewVlanVlanEditInput() // VlanVlanEditInput |  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanVlanEdit(context.Background()).VlanVlanEditInput(vlanVlanEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanVlanEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanVlanEdit`: VlanVlanEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanVlanEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanVlanEditInput** | [**VlanVlanEditInput**](VlanVlanEditInput.md) |  | 

### Return type

[**VlanVlanEditSuccess**](vlan_vlan_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanVlanInfo

> VlanVlanData VlanVlanInfo(ctx).VlanId(vlanId).Execute()

Display the properties of a VLAN



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    vlanId := int32(56) // int32 | The database identifier (ID) of the VLAN, a unique numeric key value automatically incremented when you add a VLAN. Use the ID to specify the VLAN of your choice. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanVlanInfo(context.Background()).VlanId(vlanId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanVlanInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanVlanInfo`: VlanVlanData
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanVlanInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vlanId** | **int32** | The database identifier (ID) of the VLAN, a unique numeric key value automatically incremented when you add a VLAN. Use the ID to specify the VLAN of your choice. | 

### Return type

[**VlanVlanData**](vlan_vlan_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## VlanVlanList

> VlanVlanData VlanVlanList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()

List the VLANs



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    orderby := "orderby_example" // string | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: <b>&lt;parameter&gt;='&lt;value&gt;'</b>. The clause must be encoded in URL format.You can add the optional keyword <b>ASC</b> (ascending) or <b>DESC</b> (descending) after each parameter. If not specified, <b>ASC</b> is used by default. The order of the parameters specified is set using their value's name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. (optional)
    where := "where_example" // string | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first<i>                                            It is not possible to use the structure <b>&lt;object-name&gt;_class_parameters like &lt;value&gt;</b> directly in the clause <b>WHERE</b>.<br/>                                        </i>. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : <b>&lt;parameter&gt;='&lt;value&gt;'</b> or <b>&lt;parameter&gt; IS NOT NULL</b>. The clause must be encoded in URL format. (optional)
    limit := int32(56) // int32 | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter <b>limit</b> must be specified in <b>lowercase</b>. (optional)
    offset := int32(56) // int32 | The number of rows to skip in the service output.The input parameter <b>offset</b> must be specified in <b>lowercase</b>. (optional)
    tags := "tags_example" // string | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format <object-type>.<parameter>, e.g. site.decription . (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.VlanApi.VlanVlanList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Tags(tags).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `VlanApi.VlanVlanList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `VlanVlanList`: VlanVlanData
    fmt.Fprintf(os.Stdout, "Response from `VlanApi.VlanVlanList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiVlanVlanListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **tags** | **string** | The list of class parameters to tag in the call, separated by a comma. Each parameter must be specified following the format &lt;object-type&gt;.&lt;parameter&gt;, e.g. site.decription . | 

### Return type

[**VlanVlanData**](vlan_vlan_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

