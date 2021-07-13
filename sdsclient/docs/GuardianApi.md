# \GuardianApi

All URIs are relative to *https://SDS-URL:443/api/v2.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GuardianPolicyAdd**](GuardianApi.md#GuardianPolicyAdd) | **Post** /guardian/policy/add | Add a policy
[**GuardianPolicyCount**](GuardianApi.md#GuardianPolicyCount) | **Get** /guardian/policy/count | Count the number of policies
[**GuardianPolicyDelete**](GuardianApi.md#GuardianPolicyDelete) | **Delete** /guardian/policy/delete | Delete a policy
[**GuardianPolicyEdit**](GuardianApi.md#GuardianPolicyEdit) | **Put** /guardian/policy/edit | Edit a policy
[**GuardianPolicyInfo**](GuardianApi.md#GuardianPolicyInfo) | **Get** /guardian/policy/info | Display the properties of a policy
[**GuardianPolicyList**](GuardianApi.md#GuardianPolicyList) | **Get** /guardian/policy/list | List the policies



## GuardianPolicyAdd

> GuardianPolicyAddSuccess GuardianPolicyAdd(ctx).GuardianPolicyAddInput(guardianPolicyAddInput).Execute()

Add a policy



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
    guardianPolicyAddInput := *sdsclient.NewGuardianPolicyAddInput() // GuardianPolicyAddInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.GuardianApi.GuardianPolicyAdd(context.Background()).GuardianPolicyAddInput(guardianPolicyAddInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuardianApi.GuardianPolicyAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GuardianPolicyAdd`: GuardianPolicyAddSuccess
    fmt.Fprintf(os.Stdout, "Response from `GuardianApi.GuardianPolicyAdd`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardianPolicyAddRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guardianPolicyAddInput** | [**GuardianPolicyAddInput**](GuardianPolicyAddInput.md) |  | 

### Return type

[**GuardianPolicyAddSuccess**](guardian_policy_add_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GuardianPolicyCount

> ApiCountResponseSuccess GuardianPolicyCount(ctx).Where(where).Execute()

Count the number of policies



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
    resp, r, err := api_client.GuardianApi.GuardianPolicyCount(context.Background()).Where(where).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuardianApi.GuardianPolicyCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GuardianPolicyCount`: ApiCountResponseSuccess
    fmt.Fprintf(os.Stdout, "Response from `GuardianApi.GuardianPolicyCount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardianPolicyCountRequest struct via the builder pattern


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


## GuardianPolicyDelete

> GuardianPolicyDeleteSuccess GuardianPolicyDelete(ctx).ServerId(serverId).ServerName(serverName).PolicyId(policyId).PolicyName(policyName).Warnings(warnings).Execute()

Delete a policy



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
    serverId := int32(56) // int32 | The database identifier (ID) of the Guardian server associated with the policy, a unique identifier automatically incremented when you add the server. Use it to identify the Guardian server of your choice. (optional)
    serverName := "serverName_example" // string | The name of the Guardian server associated with the policy. (optional)
    policyId := int32(56) // int32 | The database identifier (ID) of the policy, a unique identifier automatically incremented when you add the policy. Use it to identify the policy of your choice. (optional)
    policyName := "policyName_example" // string | The name of the policy. (optional)
    warnings := "warnings_example" // string | A way to bypass <b>(accept) </b>any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.GuardianApi.GuardianPolicyDelete(context.Background()).ServerId(serverId).ServerName(serverName).PolicyId(policyId).PolicyName(policyName).Warnings(warnings).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuardianApi.GuardianPolicyDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GuardianPolicyDelete`: GuardianPolicyDeleteSuccess
    fmt.Fprintf(os.Stdout, "Response from `GuardianApi.GuardianPolicyDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardianPolicyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverId** | **int32** | The database identifier (ID) of the Guardian server associated with the policy, a unique identifier automatically incremented when you add the server. Use it to identify the Guardian server of your choice. | 
 **serverName** | **string** | The name of the Guardian server associated with the policy. | 
 **policyId** | **int32** | The database identifier (ID) of the policy, a unique identifier automatically incremented when you add the policy. Use it to identify the policy of your choice. | 
 **policyName** | **string** | The name of the policy. | 
 **warnings** | **string** | A way to bypass &lt;b&gt;(accept) &lt;/b&gt;any enabled rule that would return warning messages. If the service returns an error message, you cannot bypass the enabled rules. | 

### Return type

[**GuardianPolicyDeleteSuccess**](guardian_policy_delete_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GuardianPolicyEdit

> GuardianPolicyEditSuccess GuardianPolicyEdit(ctx).GuardianPolicyEditInput(guardianPolicyEditInput).Execute()

Edit a policy



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
    guardianPolicyEditInput := *sdsclient.NewGuardianPolicyEditInput() // GuardianPolicyEditInput |  (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.GuardianApi.GuardianPolicyEdit(context.Background()).GuardianPolicyEditInput(guardianPolicyEditInput).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuardianApi.GuardianPolicyEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GuardianPolicyEdit`: GuardianPolicyEditSuccess
    fmt.Fprintf(os.Stdout, "Response from `GuardianApi.GuardianPolicyEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardianPolicyEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **guardianPolicyEditInput** | [**GuardianPolicyEditInput**](GuardianPolicyEditInput.md) |  | 

### Return type

[**GuardianPolicyEditSuccess**](guardian_policy_edit_success.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GuardianPolicyInfo

> GuardianPolicyData GuardianPolicyInfo(ctx).PolicyId(policyId).Execute()

Display the properties of a policy



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
    policyId := int32(56) // int32 | The database identifier (ID) of the policy, a unique identifier automatically incremented when you add the policy. Use it to identify the policy of your choice. (optional)

    configuration := sdsclient.NewConfiguration()
    api_client := sdsclient.NewAPIClient(configuration)
    resp, r, err := api_client.GuardianApi.GuardianPolicyInfo(context.Background()).PolicyId(policyId).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuardianApi.GuardianPolicyInfo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GuardianPolicyInfo`: GuardianPolicyData
    fmt.Fprintf(os.Stdout, "Response from `GuardianApi.GuardianPolicyInfo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardianPolicyInfoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **policyId** | **int32** | The database identifier (ID) of the policy, a unique identifier automatically incremented when you add the policy. Use it to identify the policy of your choice. | 

### Return type

[**GuardianPolicyData**](guardian_policy_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GuardianPolicyList

> GuardianPolicyData GuardianPolicyList(ctx).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()

List the policies



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
    resp, r, err := api_client.GuardianApi.GuardianPolicyList(context.Background()).Orderby(orderby).Where(where).Limit(limit).Offset(offset).Execute()
    if err.Error() != "" {
        fmt.Fprintf(os.Stderr, "Error when calling `GuardianApi.GuardianPolicyList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GuardianPolicyList`: GuardianPolicyData
    fmt.Fprintf(os.Stdout, "Response from `GuardianApi.GuardianPolicyList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGuardianPolicyListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **orderby** | **string** | A clause that allows to sort the result. You can include any output parameter of the service in this clause, except class parameters.To sort the result using class parameters, you must tag them first. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as follows: &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt;. The clause must be encoded in URL format.You can add the optional keyword &lt;b&gt;ASC&lt;/b&gt; (ascending) or &lt;b&gt;DESC&lt;/b&gt; (descending) after each parameter. If not specified, &lt;b&gt;ASC&lt;/b&gt; is used by default. The order of the parameters specified is set using their value&#39;s name or ordinal number. Each parameter value is compared from one row to the next. If all the parameters of the rows are equal, they are returned in an implementation-dependent order. | 
 **where** | **string** | A clause that allows to filter the result. You can include any output parameter of the service in this clause, except class parameters.To filter the result using class parameters, you must tag them first&lt;i&gt;                                            It is not possible to use the structure &lt;b&gt;&amp;lt;object-name&amp;gt;_class_parameters like &amp;lt;value&amp;gt;&lt;/b&gt; directly in the clause &lt;b&gt;WHERE&lt;/b&gt;.&lt;br/&gt;                                        &lt;/i&gt;. For more details, refer to the section .The parameters and their value must be specified following the operators and syntax of the SQL standard, as in the following examples : &lt;b&gt;&amp;lt;parameter&amp;gt;&#x3D;&#39;&amp;lt;value&amp;gt;&#39;&lt;/b&gt; or &lt;b&gt;&amp;lt;parameter&amp;gt; IS NOT NULL&lt;/b&gt;. The clause must be encoded in URL format. | 
 **limit** | **int32** | The maximum number of results to be returned. Depending on the user resources and the database content, it can return less results than the value you have specified.The input parameter &lt;b&gt;limit&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 
 **offset** | **int32** | The number of rows to skip in the service output.The input parameter &lt;b&gt;offset&lt;/b&gt; must be specified in &lt;b&gt;lowercase&lt;/b&gt;. | 

### Return type

[**GuardianPolicyData**](guardian_policy_data.md)

### Authorization

[BasicAuth](../README.md#BasicAuth), [HttpHeaderLoginKey](../README.md#HttpHeaderLoginKey), [HttpHeaderPasswordKey](../README.md#HttpHeaderPasswordKey)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

