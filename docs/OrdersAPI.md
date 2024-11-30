# \OrdersAPI

All URIs are relative to *https://api.schwabapi.com/trader/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelOrder**](OrdersAPI.md#CancelOrder) | **Delete** /accounts/{accountNumber}/orders/{orderId} | Cancel an order for a specific account
[**GetOrder**](OrdersAPI.md#GetOrder) | **Get** /accounts/{accountNumber}/orders/{orderId} | Get a specific order by its ID, for a specific account
[**GetOrdersByPathParam**](OrdersAPI.md#GetOrdersByPathParam) | **Get** /accounts/{accountNumber}/orders | Get all orders for a specific account.
[**GetOrdersByQueryParam**](OrdersAPI.md#GetOrdersByQueryParam) | **Get** /orders | Get all orders for all accounts
[**PlaceOrder**](OrdersAPI.md#PlaceOrder) | **Post** /accounts/{accountNumber}/orders | Place order for a specific account.
[**PreviewOrder**](OrdersAPI.md#PreviewOrder) | **Post** /accounts/{accountNumber}/previewOrder | Preview order for a specific account. **Coming Soon**.
[**ReplaceOrder**](OrdersAPI.md#ReplaceOrder) | **Put** /accounts/{accountNumber}/orders/{orderId} | Replace order for a specific account



## CancelOrder

> CancelOrder(ctx, accountNumber, orderId).Execute()

Cancel an order for a specific account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	accountNumber := "accountNumber_example" // string | The encrypted ID of the account
	orderId := int64(789) // int64 | The ID of the order being cancelled

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersAPI.CancelOrder(context.Background(), accountNumber, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.CancelOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | The encrypted ID of the account | 
**orderId** | **int64** | The ID of the order being cancelled | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrder

> Order GetOrder(ctx, accountNumber, orderId).Execute()

Get a specific order by its ID, for a specific account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	accountNumber := "accountNumber_example" // string | The encrypted ID of the account
	orderId := int64(789) // int64 | The ID of the order being retrieved.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetOrder(context.Background(), accountNumber, orderId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrder`: Order
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | The encrypted ID of the account | 
**orderId** | **int64** | The ID of the order being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Order**](Order.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersByPathParam

> []Order GetOrdersByPathParam(ctx, accountNumber).FromEnteredTime(fromEnteredTime).ToEnteredTime(toEnteredTime).MaxResults(maxResults).Status(status).Execute()

Get all orders for a specific account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	accountNumber := "accountNumber_example" // string | The encrypted ID of the account
	fromEnteredTime := "fromEnteredTime_example" // string | Specifies that no orders entered before this time should be returned.\\nValid ISO-8601 formats are :<br> <code>yyyy-MM-dd'T'HH:mm:ss.SSSZ</code>  Example fromEnteredTime is '2024-03-29T00:00:00.000Z'.\\n'toEnteredTime' must also be set.
	toEnteredTime := "toEnteredTime_example" // string | Specifies that no orders entered after this time should be returned.Valid\\nISO-8601 formats are :<br> <code>yyyy-MM-dd'T'HH:mm:ss.SSSZ</code>.  Example toEnteredTime is '2024-04-28T23:59:59.000Z'.\\n'fromEnteredTime' must also be set.
	maxResults := int64(789) // int64 | The max number of orders to retrieve. Default is 3000. (optional)
	status := openapiclient.apiOrderStatus("AWAITING_PARENT_ORDER") // ApiOrderStatus | Specifies that only orders of this status should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetOrdersByPathParam(context.Background(), accountNumber).FromEnteredTime(fromEnteredTime).ToEnteredTime(toEnteredTime).MaxResults(maxResults).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrdersByPathParam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersByPathParam`: []Order
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrdersByPathParam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | The encrypted ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersByPathParamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromEnteredTime** | **string** | Specifies that no orders entered before this time should be returned.\\nValid ISO-8601 formats are :&lt;br&gt; &lt;code&gt;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt;  Example fromEnteredTime is &#39;2024-03-29T00:00:00.000Z&#39;.\\n&#39;toEnteredTime&#39; must also be set. | 
 **toEnteredTime** | **string** | Specifies that no orders entered after this time should be returned.Valid\\nISO-8601 formats are :&lt;br&gt; &lt;code&gt;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt;.  Example toEnteredTime is &#39;2024-04-28T23:59:59.000Z&#39;.\\n&#39;fromEnteredTime&#39; must also be set. | 
 **maxResults** | **int64** | The max number of orders to retrieve. Default is 3000. | 
 **status** | [**ApiOrderStatus**](ApiOrderStatus.md) | Specifies that only orders of this status should be returned. | 

### Return type

[**[]Order**](Order.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrdersByQueryParam

> []Order GetOrdersByQueryParam(ctx).FromEnteredTime(fromEnteredTime).ToEnteredTime(toEnteredTime).MaxResults(maxResults).Status(status).Execute()

Get all orders for all accounts



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	fromEnteredTime := "fromEnteredTime_example" // string | Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are-\\nyyyy-MM-dd'T'HH:mm:ss.SSSZ Date must be within 60 days from today's date.\\n'toEnteredTime' must also be set.
	toEnteredTime := "toEnteredTime_example" // string | Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are -\\nyyyy-MM-dd'T'HH:mm:ss.SSSZ. 'fromEnteredTime' must also be set.
	maxResults := int64(789) // int64 | The max number of orders to retrieve. Default is 3000. (optional)
	status := openapiclient.apiOrderStatus("AWAITING_PARENT_ORDER") // ApiOrderStatus | Specifies that only orders of this status should be returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.GetOrdersByQueryParam(context.Background()).FromEnteredTime(fromEnteredTime).ToEnteredTime(toEnteredTime).MaxResults(maxResults).Status(status).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.GetOrdersByQueryParam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOrdersByQueryParam`: []Order
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.GetOrdersByQueryParam`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOrdersByQueryParamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fromEnteredTime** | **string** | Specifies that no orders entered before this time should be returned. Valid ISO-8601 formats are-\\nyyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ Date must be within 60 days from today&#39;s date.\\n&#39;toEnteredTime&#39; must also be set. | 
 **toEnteredTime** | **string** | Specifies that no orders entered after this time should be returned.Valid ISO-8601 formats are -\\nyyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ. &#39;fromEnteredTime&#39; must also be set. | 
 **maxResults** | **int64** | The max number of orders to retrieve. Default is 3000. | 
 **status** | [**ApiOrderStatus**](ApiOrderStatus.md) | Specifies that only orders of this status should be returned. | 

### Return type

[**[]Order**](Order.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PlaceOrder

> PlaceOrder(ctx, accountNumber).Body(body).Execute()

Place order for a specific account.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	accountNumber := "accountNumber_example" // string | The encrypted ID of the account
	body := *openapiclient.NewOrderRequest() // OrderRequest | The new Order Object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersAPI.PlaceOrder(context.Background(), accountNumber).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.PlaceOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | The encrypted ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPlaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OrderRequest**](OrderRequest.md) | The new Order Object. | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PreviewOrder

> PreviewOrder PreviewOrder(ctx, accountNumber).PreviewOrder(previewOrder).Execute()

Preview order for a specific account. **Coming Soon**.



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	accountNumber := "accountNumber_example" // string | The encrypted ID of the account
	previewOrder := *openapiclient.NewPreviewOrder() // PreviewOrder | The Order Object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OrdersAPI.PreviewOrder(context.Background(), accountNumber).PreviewOrder(previewOrder).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.PreviewOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PreviewOrder`: PreviewOrder
	fmt.Fprintf(os.Stdout, "Response from `OrdersAPI.PreviewOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | The encrypted ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiPreviewOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **previewOrder** | [**PreviewOrder**](PreviewOrder.md) | The Order Object. | 

### Return type

[**PreviewOrder**](PreviewOrder.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceOrder

> ReplaceOrder(ctx, accountNumber, orderId).Body(body).Execute()

Replace order for a specific account



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/kebroad/schwab-accounts-and-trading-production-go"
)

func main() {
	accountNumber := "accountNumber_example" // string | The encrypted ID of the account
	orderId := int64(789) // int64 | The ID of the order being retrieved.
	body := *openapiclient.NewOrderRequest() // OrderRequest | The Order Object.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.OrdersAPI.ReplaceOrder(context.Background(), accountNumber, orderId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OrdersAPI.ReplaceOrder``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | The encrypted ID of the account | 
**orderId** | **int64** | The ID of the order being retrieved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**OrderRequest**](OrderRequest.md) | The Order Object. | 

### Return type

 (empty response body)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

