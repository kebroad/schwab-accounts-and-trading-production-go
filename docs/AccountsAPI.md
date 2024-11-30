# \AccountsAPI

All URIs are relative to *https://api.schwabapi.com/trader/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccount**](AccountsAPI.md#GetAccount) | **Get** /accounts/{accountNumber} | Get a specific account balance and positions for the logged in user.
[**GetAccountNumbers**](AccountsAPI.md#GetAccountNumbers) | **Get** /accounts/accountNumbers | Get list of account numbers and their encrypted values
[**GetAccounts**](AccountsAPI.md#GetAccounts) | **Get** /accounts | Get linked account(s) balances and positions for the logged in user.



## GetAccount

> Account GetAccount(ctx, accountNumber).Fields(fields).Execute()

Get a specific account balance and positions for the logged in user.



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
	fields := "fields_example" // string | This allows one to determine\\nwhich fields they want returned. Possible values in this String can be:\\n<br><code>positions</code><br> Example:<br><code>fields=positions</code> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccount(context.Background(), accountNumber).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccount`: Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accountNumber** | **string** | The encrypted ID of the account | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **string** | This allows one to determine\\nwhich fields they want returned. Possible values in this String can be:\\n&lt;br&gt;&lt;code&gt;positions&lt;/code&gt;&lt;br&gt; Example:&lt;br&gt;&lt;code&gt;fields&#x3D;positions&lt;/code&gt; | 

### Return type

[**Account**](Account.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccountNumbers

> []AccountNumberHash GetAccountNumbers(ctx).Execute()

Get list of account numbers and their encrypted values



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccountNumbers(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccountNumbers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccountNumbers`: []AccountNumberHash
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccountNumbers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountNumbersRequest struct via the builder pattern


### Return type

[**[]AccountNumberHash**](AccountNumberHash.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccounts

> []Account GetAccounts(ctx).Fields(fields).Execute()

Get linked account(s) balances and positions for the logged in user.



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
	fields := "fields_example" // string | This allows one to determine which fields they want returned. Possible value in this String can be:\\n<br><code>positions</code><br> Example:<br><code>fields=positions</code> (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AccountsAPI.GetAccounts(context.Background()).Fields(fields).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AccountsAPI.GetAccounts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccounts`: []Account
	fmt.Fprintf(os.Stdout, "Response from `AccountsAPI.GetAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fields** | **string** | This allows one to determine which fields they want returned. Possible value in this String can be:\\n&lt;br&gt;&lt;code&gt;positions&lt;/code&gt;&lt;br&gt; Example:&lt;br&gt;&lt;code&gt;fields&#x3D;positions&lt;/code&gt; | 

### Return type

[**[]Account**](Account.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

