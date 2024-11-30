# \UserPreferenceAPI

All URIs are relative to *https://api.schwabapi.com/trader/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserPreference**](UserPreferenceAPI.md#GetUserPreference) | **Get** /userPreference | Get user preference information for the logged in user.



## GetUserPreference

> []UserPreference GetUserPreference(ctx).Execute()

Get user preference information for the logged in user.



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
	resp, r, err := apiClient.UserPreferenceAPI.GetUserPreference(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserPreferenceAPI.GetUserPreference``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserPreference`: []UserPreference
	fmt.Fprintf(os.Stdout, "Response from `UserPreferenceAPI.GetUserPreference`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPreferenceRequest struct via the builder pattern


### Return type

[**[]UserPreference**](UserPreference.md)

### Authorization

[oauth](../README.md#oauth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

