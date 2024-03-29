/*
Track & Trace Events

Retrieve Track & Trace Events based on DCSA Interface standard v.2.2  This service provides shippers and consignees visibility to Shipment, Equipment and Transport events for shipments booked with A.P. Moller-Maersk A/S using standards set by the Digital Container Shipping Association.\\ <https://dcsa.org/>

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package maersk

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
)

type AuthApi interface {

	/*
		CreateAccessToken Method for CreateAccessToken

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ApiCreateAccessTokenRequest
	*/
	CreateAccessToken(ctx context.Context) ApiCreateAccessTokenRequest

	// CreateAccessTokenExecute executes the request
	//  @return CreateAccessToken200Response
	CreateAccessTokenExecute(r ApiCreateAccessTokenRequest) (*CreateAccessToken200Response, *http.Response, error)
}

// AuthApiService AuthApi service
type AuthApiService service

type ApiCreateAccessTokenRequest struct {
	ctx          context.Context
	ApiService   AuthApi
	grantType    *string
	clientId     *string
	clientSecret *string
}

func (r ApiCreateAccessTokenRequest) GrantType(grantType string) ApiCreateAccessTokenRequest {
	r.grantType = &grantType
	return r
}

func (r ApiCreateAccessTokenRequest) ClientId(clientId string) ApiCreateAccessTokenRequest {
	r.clientId = &clientId
	return r
}

func (r ApiCreateAccessTokenRequest) ClientSecret(clientSecret string) ApiCreateAccessTokenRequest {
	r.clientSecret = &clientSecret
	return r
}

func (r ApiCreateAccessTokenRequest) Execute() (*CreateAccessToken200Response, *http.Response, error) {
	return r.ApiService.CreateAccessTokenExecute(r)
}

/*
CreateAccessToken Method for CreateAccessToken

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ApiCreateAccessTokenRequest
*/
func (a *AuthApiService) CreateAccessToken(ctx context.Context) ApiCreateAccessTokenRequest {
	return ApiCreateAccessTokenRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return CreateAccessToken200Response
func (a *AuthApiService) CreateAccessTokenExecute(r ApiCreateAccessTokenRequest) (*CreateAccessToken200Response, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue *CreateAccessToken200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AuthApiService.CreateAccessToken")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/oauth2/access_token"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.grantType == nil {
		return localVarReturnValue, nil, reportError("grantType is required and must be specified")
	}
	if r.clientId == nil {
		return localVarReturnValue, nil, reportError("clientId is required and must be specified")
	}
	if r.clientSecret == nil {
		return localVarReturnValue, nil, reportError("clientSecret is required and must be specified")
	}

	localVarQueryParams.Add("grant_type", parameterToString(*r.grantType, ""))
	localVarQueryParams.Add("client_id", parameterToString(*r.clientId, ""))
	localVarQueryParams.Add("client_secret", parameterToString(*r.clientSecret, ""))
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v CreateAccessToken400Response
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
