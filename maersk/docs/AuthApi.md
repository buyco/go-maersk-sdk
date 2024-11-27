# \AuthApi

All URIs are relative to *https://api.maersk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessToken**](AuthApi.md#CreateAccessToken) | **Post** /customer-identity/oauth/v2/access_token | 



## CreateAccessToken

> CreateAccessToken200Response CreateAccessToken(ctx).ConsumerKey(consumerKey).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()



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
    consumerKey := "consumerKey_example" // string | The Consumer Key issued for your registered application.
    grantType := "grantType_example" // string | 
    clientId := "clientId_example" // string | 
    clientSecret := "clientSecret_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AuthApi.CreateAccessToken(context.Background()).ConsumerKey(consumerKey).GrantType(grantType).ClientId(clientId).ClientSecret(clientSecret).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuthApi.CreateAccessToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccessToken`: CreateAccessToken200Response
    fmt.Fprintf(os.Stdout, "Response from `AuthApi.CreateAccessToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumerKey** | **string** | The Consumer Key issued for your registered application. | 
 **grantType** | **string** |  | 
 **clientId** | **string** |  | 
 **clientSecret** | **string** |  | 

### Return type

[**CreateAccessToken200Response**](CreateAccessToken200Response.md)

### Authorization

[ConsumerKey](../README.md#ConsumerKey)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

