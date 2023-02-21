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

// CreateAccessToken400Response struct for CreateAccessToken400Response
type CreateAccessToken400Response struct {
	ErrorDescription *string `json:"error_description,omitempty"`
	Error            *string `json:"error,omitempty"`
}

// NewCreateAccessToken400Response instantiates a new CreateAccessToken400Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccessToken400Response() *CreateAccessToken400Response {
	this := CreateAccessToken400Response{}
	return &this
}

// NewCreateAccessToken400ResponseWithDefaults instantiates a new CreateAccessToken400Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccessToken400ResponseWithDefaults() *CreateAccessToken400Response {
	this := CreateAccessToken400Response{}
	return &this
}

// GetErrorDescription returns the ErrorDescription field value if set, zero value otherwise.
func (o *CreateAccessToken400Response) GetErrorDescription() string {
	if o == nil || o.ErrorDescription == nil {
		var ret string
		return ret
	}
	return *o.ErrorDescription
}

// GetErrorDescriptionOk returns a tuple with the ErrorDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessToken400Response) GetErrorDescriptionOk() (*string, bool) {
	if o == nil || o.ErrorDescription == nil {
		return nil, false
	}
	return o.ErrorDescription, true
}

// HasErrorDescription returns a boolean if a field has been set.
func (o *CreateAccessToken400Response) HasErrorDescription() bool {
	if o != nil && o.ErrorDescription != nil {
		return true
	}

	return false
}

// SetErrorDescription gets a reference to the given string and assigns it to the ErrorDescription field.
func (o *CreateAccessToken400Response) SetErrorDescription(v string) {
	o.ErrorDescription = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CreateAccessToken400Response) GetError() string {
	if o == nil || o.Error == nil {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessToken400Response) GetErrorOk() (*string, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CreateAccessToken400Response) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *CreateAccessToken400Response) SetError(v string) {
	o.Error = &v
}

func (o CreateAccessToken400Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ErrorDescription != nil {
		toSerialize["error_description"] = o.ErrorDescription
	}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAccessToken400Response struct {
	value *CreateAccessToken400Response
	isSet bool
}

func (v NullableCreateAccessToken400Response) Get() *CreateAccessToken400Response {
	return v.value
}

func (v *NullableCreateAccessToken400Response) Set(val *CreateAccessToken400Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccessToken400Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccessToken400Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccessToken400Response(val *CreateAccessToken400Response) *NullableCreateAccessToken400Response {
	return &NullableCreateAccessToken400Response{value: val, isSet: true}
}

func (v NullableCreateAccessToken400Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccessToken400Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}