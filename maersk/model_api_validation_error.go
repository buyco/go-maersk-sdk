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

// ApiValidationError struct for ApiValidationError
type ApiValidationError struct {
	// The field that has failed validation.
	Field string `json:"field"`
	// The value that has failed validation.
	RejectedValue string `json:"rejectedValue"`
	// The reason and advice for failed validation.
	Message string `json:"message"`
}

// NewApiValidationError instantiates a new ApiValidationError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiValidationError(field string, rejectedValue string, message string) *ApiValidationError {
	this := ApiValidationError{}
	this.Field = field
	this.RejectedValue = rejectedValue
	this.Message = message
	return &this
}

// NewApiValidationErrorWithDefaults instantiates a new ApiValidationError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiValidationErrorWithDefaults() *ApiValidationError {
	this := ApiValidationError{}
	return &this
}

// GetField returns the Field field value
func (o *ApiValidationError) GetField() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Field
}

// GetFieldOk returns a tuple with the Field field value
// and a boolean to check if the value has been set.
func (o *ApiValidationError) GetFieldOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Field, true
}

// SetField sets field value
func (o *ApiValidationError) SetField(v string) {
	o.Field = v
}

// GetRejectedValue returns the RejectedValue field value
func (o *ApiValidationError) GetRejectedValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RejectedValue
}

// GetRejectedValueOk returns a tuple with the RejectedValue field value
// and a boolean to check if the value has been set.
func (o *ApiValidationError) GetRejectedValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RejectedValue, true
}

// SetRejectedValue sets field value
func (o *ApiValidationError) SetRejectedValue(v string) {
	o.RejectedValue = v
}

// GetMessage returns the Message field value
func (o *ApiValidationError) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *ApiValidationError) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *ApiValidationError) SetMessage(v string) {
	o.Message = v
}

func (o ApiValidationError) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["field"] = o.Field
	}
	if true {
		toSerialize["rejectedValue"] = o.RejectedValue
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableApiValidationError struct {
	value *ApiValidationError
	isSet bool
}

func (v NullableApiValidationError) Get() *ApiValidationError {
	return v.value
}

func (v *NullableApiValidationError) Set(val *ApiValidationError) {
	v.value = val
	v.isSet = true
}

func (v NullableApiValidationError) IsSet() bool {
	return v.isSet
}

func (v *NullableApiValidationError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiValidationError(val *ApiValidationError) *NullableApiValidationError {
	return &NullableApiValidationError{value: val, isSet: true}
}

func (v NullableApiValidationError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiValidationError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
