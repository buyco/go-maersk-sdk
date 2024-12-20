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

// CreateAccessToken401Response struct for CreateAccessToken401Response
type CreateAccessToken401Response struct {
	Error []CreateAccessToken401ResponseErrorInner `json:"error,omitempty"`
}

// NewCreateAccessToken401Response instantiates a new CreateAccessToken401Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateAccessToken401Response() *CreateAccessToken401Response {
	this := CreateAccessToken401Response{}
	return &this
}

// NewCreateAccessToken401ResponseWithDefaults instantiates a new CreateAccessToken401Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateAccessToken401ResponseWithDefaults() *CreateAccessToken401Response {
	this := CreateAccessToken401Response{}
	return &this
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *CreateAccessToken401Response) GetError() []CreateAccessToken401ResponseErrorInner {
	if o == nil || o.Error == nil {
		var ret []CreateAccessToken401ResponseErrorInner
		return ret
	}
	return o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateAccessToken401Response) GetErrorOk() ([]CreateAccessToken401ResponseErrorInner, bool) {
	if o == nil || o.Error == nil {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *CreateAccessToken401Response) HasError() bool {
	if o != nil && o.Error != nil {
		return true
	}

	return false
}

// SetError gets a reference to the given []CreateAccessToken401ResponseErrorInner and assigns it to the Error field.
func (o *CreateAccessToken401Response) SetError(v []CreateAccessToken401ResponseErrorInner) {
	o.Error = v
}

func (o CreateAccessToken401Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Error != nil {
		toSerialize["error"] = o.Error
	}
	return json.Marshal(toSerialize)
}

type NullableCreateAccessToken401Response struct {
	value *CreateAccessToken401Response
	isSet bool
}

func (v NullableCreateAccessToken401Response) Get() *CreateAccessToken401Response {
	return v.value
}

func (v *NullableCreateAccessToken401Response) Set(val *CreateAccessToken401Response) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateAccessToken401Response) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateAccessToken401Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateAccessToken401Response(val *CreateAccessToken401Response) *NullableCreateAccessToken401Response {
	return &NullableCreateAccessToken401Response{value: val, isSet: true}
}

func (v NullableCreateAccessToken401Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateAccessToken401Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
