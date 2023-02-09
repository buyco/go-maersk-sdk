/*
Track & Trace Events

Retrieve Track & Trace Events based on DCSA Interface standard v.2.2  This service provides shippers and consignees visibility to Shipment, Equipment and Transport events for shipments booked with A.P. Moller-Maersk A/S using standards set by the Digital Container Shipping Association.\\ <https://dcsa.org/>

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
)

// Address Address related information
type Address struct {
	// Name of the address
	AddressName *string `json:"addressName,omitempty"`
	// The name of the street of the party’s address.
	StreetName *string `json:"streetName,omitempty"`
	// The number of the street of the party’s address.
	StreetNumber *string `json:"streetNumber,omitempty"`
	// The floor of the party’s street number.
	Floor *string `json:"floor,omitempty"`
	// The postal code of the party’s address.
	PostCode *string `json:"postCode,omitempty"`
	// The city name of the party’s address.
	CityName *string `json:"cityName,omitempty"`
	// The state/region of the party’s address.
	StateRegion *string `json:"stateRegion,omitempty"`
	// The country of the party’s address.
	Country *string `json:"country,omitempty"`
}

// NewAddress instantiates a new Address object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddress() *Address {
	this := Address{}
	return &this
}

// NewAddressWithDefaults instantiates a new Address object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressWithDefaults() *Address {
	this := Address{}
	return &this
}

// GetAddressName returns the AddressName field value if set, zero value otherwise.
func (o *Address) GetAddressName() string {
	if o == nil || o.AddressName == nil {
		var ret string
		return ret
	}
	return *o.AddressName
}

// GetAddressNameOk returns a tuple with the AddressName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetAddressNameOk() (*string, bool) {
	if o == nil || o.AddressName == nil {
		return nil, false
	}
	return o.AddressName, true
}

// HasAddressName returns a boolean if a field has been set.
func (o *Address) HasAddressName() bool {
	if o != nil && o.AddressName != nil {
		return true
	}

	return false
}

// SetAddressName gets a reference to the given string and assigns it to the AddressName field.
func (o *Address) SetAddressName(v string) {
	o.AddressName = &v
}

// GetStreetName returns the StreetName field value if set, zero value otherwise.
func (o *Address) GetStreetName() string {
	if o == nil || o.StreetName == nil {
		var ret string
		return ret
	}
	return *o.StreetName
}

// GetStreetNameOk returns a tuple with the StreetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStreetNameOk() (*string, bool) {
	if o == nil || o.StreetName == nil {
		return nil, false
	}
	return o.StreetName, true
}

// HasStreetName returns a boolean if a field has been set.
func (o *Address) HasStreetName() bool {
	if o != nil && o.StreetName != nil {
		return true
	}

	return false
}

// SetStreetName gets a reference to the given string and assigns it to the StreetName field.
func (o *Address) SetStreetName(v string) {
	o.StreetName = &v
}

// GetStreetNumber returns the StreetNumber field value if set, zero value otherwise.
func (o *Address) GetStreetNumber() string {
	if o == nil || o.StreetNumber == nil {
		var ret string
		return ret
	}
	return *o.StreetNumber
}

// GetStreetNumberOk returns a tuple with the StreetNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStreetNumberOk() (*string, bool) {
	if o == nil || o.StreetNumber == nil {
		return nil, false
	}
	return o.StreetNumber, true
}

// HasStreetNumber returns a boolean if a field has been set.
func (o *Address) HasStreetNumber() bool {
	if o != nil && o.StreetNumber != nil {
		return true
	}

	return false
}

// SetStreetNumber gets a reference to the given string and assigns it to the StreetNumber field.
func (o *Address) SetStreetNumber(v string) {
	o.StreetNumber = &v
}

// GetFloor returns the Floor field value if set, zero value otherwise.
func (o *Address) GetFloor() string {
	if o == nil || o.Floor == nil {
		var ret string
		return ret
	}
	return *o.Floor
}

// GetFloorOk returns a tuple with the Floor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetFloorOk() (*string, bool) {
	if o == nil || o.Floor == nil {
		return nil, false
	}
	return o.Floor, true
}

// HasFloor returns a boolean if a field has been set.
func (o *Address) HasFloor() bool {
	if o != nil && o.Floor != nil {
		return true
	}

	return false
}

// SetFloor gets a reference to the given string and assigns it to the Floor field.
func (o *Address) SetFloor(v string) {
	o.Floor = &v
}

// GetPostCode returns the PostCode field value if set, zero value otherwise.
func (o *Address) GetPostCode() string {
	if o == nil || o.PostCode == nil {
		var ret string
		return ret
	}
	return *o.PostCode
}

// GetPostCodeOk returns a tuple with the PostCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetPostCodeOk() (*string, bool) {
	if o == nil || o.PostCode == nil {
		return nil, false
	}
	return o.PostCode, true
}

// HasPostCode returns a boolean if a field has been set.
func (o *Address) HasPostCode() bool {
	if o != nil && o.PostCode != nil {
		return true
	}

	return false
}

// SetPostCode gets a reference to the given string and assigns it to the PostCode field.
func (o *Address) SetPostCode(v string) {
	o.PostCode = &v
}

// GetCityName returns the CityName field value if set, zero value otherwise.
func (o *Address) GetCityName() string {
	if o == nil || o.CityName == nil {
		var ret string
		return ret
	}
	return *o.CityName
}

// GetCityNameOk returns a tuple with the CityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCityNameOk() (*string, bool) {
	if o == nil || o.CityName == nil {
		return nil, false
	}
	return o.CityName, true
}

// HasCityName returns a boolean if a field has been set.
func (o *Address) HasCityName() bool {
	if o != nil && o.CityName != nil {
		return true
	}

	return false
}

// SetCityName gets a reference to the given string and assigns it to the CityName field.
func (o *Address) SetCityName(v string) {
	o.CityName = &v
}

// GetStateRegion returns the StateRegion field value if set, zero value otherwise.
func (o *Address) GetStateRegion() string {
	if o == nil || o.StateRegion == nil {
		var ret string
		return ret
	}
	return *o.StateRegion
}

// GetStateRegionOk returns a tuple with the StateRegion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetStateRegionOk() (*string, bool) {
	if o == nil || o.StateRegion == nil {
		return nil, false
	}
	return o.StateRegion, true
}

// HasStateRegion returns a boolean if a field has been set.
func (o *Address) HasStateRegion() bool {
	if o != nil && o.StateRegion != nil {
		return true
	}

	return false
}

// SetStateRegion gets a reference to the given string and assigns it to the StateRegion field.
func (o *Address) SetStateRegion(v string) {
	o.StateRegion = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *Address) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Address) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *Address) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *Address) SetCountry(v string) {
	o.Country = &v
}

func (o Address) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddressName != nil {
		toSerialize["addressName"] = o.AddressName
	}
	if o.StreetName != nil {
		toSerialize["streetName"] = o.StreetName
	}
	if o.StreetNumber != nil {
		toSerialize["streetNumber"] = o.StreetNumber
	}
	if o.Floor != nil {
		toSerialize["floor"] = o.Floor
	}
	if o.PostCode != nil {
		toSerialize["postCode"] = o.PostCode
	}
	if o.CityName != nil {
		toSerialize["cityName"] = o.CityName
	}
	if o.StateRegion != nil {
		toSerialize["stateRegion"] = o.StateRegion
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	return json.Marshal(toSerialize)
}

type NullableAddress struct {
	value *Address
	isSet bool
}

func (v NullableAddress) Get() *Address {
	return v.value
}

func (v *NullableAddress) Set(val *Address) {
	v.value = val
	v.isSet = true
}

func (v NullableAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddress(val *Address) *NullableAddress {
	return &NullableAddress{value: val, isSet: true}
}

func (v NullableAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
