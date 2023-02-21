/*
Track & Trace Events

Retrieve Track & Trace Events based on DCSA Interface standard v.2.2  This service provides shippers and consignees visibility to Shipment, Equipment and Transport events for shipments booked with A.P. Moller-Maersk A/S using standards set by the Digital Container Shipping Association.\\ <https://dcsa.org/>

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package maersk

import (
	"encoding/json"
)

// CreateAccessToken200Response struct for CreateAccessToken200Response
type CreateAccessToken200Response struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	IdToken     string `json:"id_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int32  `json:"expires_in"`
}

// NewCreateAccessToken200Response instantiates a new CreateAccessToken200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccessToken200Response(accessToken string, scope string, idToken string, tokenType string, expiresIn int32) *CreateAccessToken200Response {
	this := CreateAccessToken200Response{}
	this.AccessToken = accessToken
	this.Scope = scope
	this.IdToken = idToken
	this.TokenType = tokenType
	this.ExpiresIn = expiresIn
	return &this
}

// NewCreateAccessToken200ResponseWithDefaults instantiates a new CreateAccessToken200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccessToken200ResponseWithDefaults() *CreateAccessToken200Response {
	this := CreateAccessToken200Response{}
	return &this
}

// GetAccessToken returns the AccessToken field value
func (o *CreateAccessToken200Response) GetAccessToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AccessToken
}

// GetAccessTokenOk returns a tuple with the AccessToken field value
// and a boolean to check if the value has been set.
func (o *CreateAccessToken200Response) GetAccessTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AccessToken, true
}

// SetAccessToken sets field value
func (o *CreateAccessToken200Response) SetAccessToken(v string) {
	o.AccessToken = v
}

// GetScope returns the Scope field value
func (o *CreateAccessToken200Response) GetScope() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Scope
}

// GetScopeOk returns a tuple with the Scope field value
// and a boolean to check if the value has been set.
func (o *CreateAccessToken200Response) GetScopeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Scope, true
}

// SetScope sets field value
func (o *CreateAccessToken200Response) SetScope(v string) {
	o.Scope = v
}

// GetIdToken returns the IdToken field value
func (o *CreateAccessToken200Response) GetIdToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IdToken
}

// GetIdTokenOk returns a tuple with the IdToken field value
// and a boolean to check if the value has been set.
func (o *CreateAccessToken200Response) GetIdTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IdToken, true
}

// SetIdToken sets field value
func (o *CreateAccessToken200Response) SetIdToken(v string) {
	o.IdToken = v
}

// GetTokenType returns the TokenType field value
func (o *CreateAccessToken200Response) GetTokenType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TokenType
}

// GetTokenTypeOk returns a tuple with the TokenType field value
// and a boolean to check if the value has been set.
func (o *CreateAccessToken200Response) GetTokenTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TokenType, true
}

// SetTokenType sets field value
func (o *CreateAccessToken200Response) SetTokenType(v string) {
	o.TokenType = v
}

// GetExpiresIn returns the ExpiresIn field value
func (o *CreateAccessToken200Response) GetExpiresIn() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value
// and a boolean to check if the value has been set.
func (o *CreateAccessToken200Response) GetExpiresInOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExpiresIn, true
}

// SetExpiresIn sets field value
func (o *CreateAccessToken200Response) SetExpiresIn(v int32) {
	o.ExpiresIn = v
}

func (o CreateAccessToken200Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["access_token"] = o.AccessToken
	}
	if true {
		toSerialize["scope"] = o.Scope
	}
	if true {
		toSerialize["id_token"] = o.IdToken
	}
	if true {
		toSerialize["token_type"] = o.TokenType
	}
	if true {
		toSerialize["expires_in"] = o.ExpiresIn
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAccessToken200Response struct {
	value *CreateAccessToken200Response
	isSet bool
}

func (v NullableCreateAccessToken200Response) Get() *CreateAccessToken200Response {
	return v.value
}

func (v *NullableCreateAccessToken200Response) Set(val *CreateAccessToken200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccessToken200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccessToken200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccessToken200Response(val *CreateAccessToken200Response) *NullableCreateAccessToken200Response {
	return &NullableCreateAccessToken200Response{value: val, isSet: true}
}

func (v NullableCreateAccessToken200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccessToken200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}