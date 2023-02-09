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

// Seal Addresses the seal-related information associated with the shipment equipment. A seal is put on a shipment equipment once it is loaded. This seal is meant to stay on until the shipment equipment reaches its final destination.
type Seal struct {
	// Identifies a seal affixed to the container.
	SealNumber string `json:"sealNumber"`
	// The source of the seal, namely who has affixed the seal. This attribute links to the Seal Source ID defined in the Seal Source reference data entity. * CAR (Carrier) * SHI (Shipper) * PHY (Phytosanitary) * VET (Veterinary) * CUS (Customs) 
	SealSource *string `json:"sealSource,omitempty"`
	// The type of seal. This attribute links to the Seal Type ID defined in the Seal Type reference data entity. * KLP (Keyless padlock) * BLT (Bolt) * WIR (Wire) 
	SealType string `json:"sealType"`
}

// NewSeal instantiates a new Seal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeal(sealNumber string, sealType string) *Seal {
	this := Seal{}
	this.SealNumber = sealNumber
	this.SealType = sealType
	return &this
}

// NewSealWithDefaults instantiates a new Seal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSealWithDefaults() *Seal {
	this := Seal{}
	return &this
}

// GetSealNumber returns the SealNumber field value
func (o *Seal) GetSealNumber() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SealNumber
}

// GetSealNumberOk returns a tuple with the SealNumber field value
// and a boolean to check if the value has been set.
func (o *Seal) GetSealNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SealNumber, true
}

// SetSealNumber sets field value
func (o *Seal) SetSealNumber(v string) {
	o.SealNumber = v
}

// GetSealSource returns the SealSource field value if set, zero value otherwise.
func (o *Seal) GetSealSource() string {
	if o == nil || o.SealSource == nil {
		var ret string
		return ret
	}
	return *o.SealSource
}

// GetSealSourceOk returns a tuple with the SealSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Seal) GetSealSourceOk() (*string, bool) {
	if o == nil || o.SealSource == nil {
		return nil, false
	}
	return o.SealSource, true
}

// HasSealSource returns a boolean if a field has been set.
func (o *Seal) HasSealSource() bool {
	if o != nil && o.SealSource != nil {
		return true
	}

	return false
}

// SetSealSource gets a reference to the given string and assigns it to the SealSource field.
func (o *Seal) SetSealSource(v string) {
	o.SealSource = &v
}

// GetSealType returns the SealType field value
func (o *Seal) GetSealType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SealType
}

// GetSealTypeOk returns a tuple with the SealType field value
// and a boolean to check if the value has been set.
func (o *Seal) GetSealTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SealType, true
}

// SetSealType sets field value
func (o *Seal) SetSealType(v string) {
	o.SealType = v
}

func (o Seal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sealNumber"] = o.SealNumber
	}
	if o.SealSource != nil {
		toSerialize["sealSource"] = o.SealSource
	}
	if true {
		toSerialize["sealType"] = o.SealType
	}
	return json.Marshal(toSerialize)
}

type NullableSeal struct {
	value *Seal
	isSet bool
}

func (v NullableSeal) Get() *Seal {
	return v.value
}

func (v *NullableSeal) Set(val *Seal) {
	v.value = val
	v.isSet = true
}

func (v NullableSeal) IsSet() bool {
	return v.isSet
}

func (v *NullableSeal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeal(val *Seal) *NullableSeal {
	return &NullableSeal{value: val, isSet: true}
}

func (v NullableSeal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


