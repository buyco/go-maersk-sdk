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

// Location Generally used to capture location-related data; also for locations without UN Location Codes.
type Location struct {
	// The name of the location.
	LocationName *string `json:"locationName,omitempty"`
	// Geographic coordinate that specifies the north–south position of a point on the Earth's surface.
	Latitude *string `json:"latitude,omitempty"`
	// Geographic coordinate that specifies the east–west position of a point on the Earth's surface.
	Longitude *string `json:"longitude,omitempty"`
	// The UN Location code specifying where the place is located.
	UNLocationCode *string `json:"UNLocationCode,omitempty"`
	// The code used for identifying the specific facility. This code is not the UN Location Code.
	FacilityCode *string `json:"facilityCode,omitempty"`
	// The provider used for identifying the facility Code
	FacilityCodeListProvider *string `json:"facilityCodeListProvider,omitempty"`
	Address *Address `json:"address,omitempty"`
}

// NewLocation instantiates a new Location object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocation() *Location {
	this := Location{}
	return &this
}

// NewLocationWithDefaults instantiates a new Location object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocationWithDefaults() *Location {
	this := Location{}
	return &this
}

// GetLocationName returns the LocationName field value if set, zero value otherwise.
func (o *Location) GetLocationName() string {
	if o == nil || o.LocationName == nil {
		var ret string
		return ret
	}
	return *o.LocationName
}

// GetLocationNameOk returns a tuple with the LocationName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLocationNameOk() (*string, bool) {
	if o == nil || o.LocationName == nil {
		return nil, false
	}
	return o.LocationName, true
}

// HasLocationName returns a boolean if a field has been set.
func (o *Location) HasLocationName() bool {
	if o != nil && o.LocationName != nil {
		return true
	}

	return false
}

// SetLocationName gets a reference to the given string and assigns it to the LocationName field.
func (o *Location) SetLocationName(v string) {
	o.LocationName = &v
}

// GetLatitude returns the Latitude field value if set, zero value otherwise.
func (o *Location) GetLatitude() string {
	if o == nil || o.Latitude == nil {
		var ret string
		return ret
	}
	return *o.Latitude
}

// GetLatitudeOk returns a tuple with the Latitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLatitudeOk() (*string, bool) {
	if o == nil || o.Latitude == nil {
		return nil, false
	}
	return o.Latitude, true
}

// HasLatitude returns a boolean if a field has been set.
func (o *Location) HasLatitude() bool {
	if o != nil && o.Latitude != nil {
		return true
	}

	return false
}

// SetLatitude gets a reference to the given string and assigns it to the Latitude field.
func (o *Location) SetLatitude(v string) {
	o.Latitude = &v
}

// GetLongitude returns the Longitude field value if set, zero value otherwise.
func (o *Location) GetLongitude() string {
	if o == nil || o.Longitude == nil {
		var ret string
		return ret
	}
	return *o.Longitude
}

// GetLongitudeOk returns a tuple with the Longitude field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetLongitudeOk() (*string, bool) {
	if o == nil || o.Longitude == nil {
		return nil, false
	}
	return o.Longitude, true
}

// HasLongitude returns a boolean if a field has been set.
func (o *Location) HasLongitude() bool {
	if o != nil && o.Longitude != nil {
		return true
	}

	return false
}

// SetLongitude gets a reference to the given string and assigns it to the Longitude field.
func (o *Location) SetLongitude(v string) {
	o.Longitude = &v
}

// GetUNLocationCode returns the UNLocationCode field value if set, zero value otherwise.
func (o *Location) GetUNLocationCode() string {
	if o == nil || o.UNLocationCode == nil {
		var ret string
		return ret
	}
	return *o.UNLocationCode
}

// GetUNLocationCodeOk returns a tuple with the UNLocationCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetUNLocationCodeOk() (*string, bool) {
	if o == nil || o.UNLocationCode == nil {
		return nil, false
	}
	return o.UNLocationCode, true
}

// HasUNLocationCode returns a boolean if a field has been set.
func (o *Location) HasUNLocationCode() bool {
	if o != nil && o.UNLocationCode != nil {
		return true
	}

	return false
}

// SetUNLocationCode gets a reference to the given string and assigns it to the UNLocationCode field.
func (o *Location) SetUNLocationCode(v string) {
	o.UNLocationCode = &v
}

// GetFacilityCode returns the FacilityCode field value if set, zero value otherwise.
func (o *Location) GetFacilityCode() string {
	if o == nil || o.FacilityCode == nil {
		var ret string
		return ret
	}
	return *o.FacilityCode
}

// GetFacilityCodeOk returns a tuple with the FacilityCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetFacilityCodeOk() (*string, bool) {
	if o == nil || o.FacilityCode == nil {
		return nil, false
	}
	return o.FacilityCode, true
}

// HasFacilityCode returns a boolean if a field has been set.
func (o *Location) HasFacilityCode() bool {
	if o != nil && o.FacilityCode != nil {
		return true
	}

	return false
}

// SetFacilityCode gets a reference to the given string and assigns it to the FacilityCode field.
func (o *Location) SetFacilityCode(v string) {
	o.FacilityCode = &v
}

// GetFacilityCodeListProvider returns the FacilityCodeListProvider field value if set, zero value otherwise.
func (o *Location) GetFacilityCodeListProvider() string {
	if o == nil || o.FacilityCodeListProvider == nil {
		var ret string
		return ret
	}
	return *o.FacilityCodeListProvider
}

// GetFacilityCodeListProviderOk returns a tuple with the FacilityCodeListProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetFacilityCodeListProviderOk() (*string, bool) {
	if o == nil || o.FacilityCodeListProvider == nil {
		return nil, false
	}
	return o.FacilityCodeListProvider, true
}

// HasFacilityCodeListProvider returns a boolean if a field has been set.
func (o *Location) HasFacilityCodeListProvider() bool {
	if o != nil && o.FacilityCodeListProvider != nil {
		return true
	}

	return false
}

// SetFacilityCodeListProvider gets a reference to the given string and assigns it to the FacilityCodeListProvider field.
func (o *Location) SetFacilityCodeListProvider(v string) {
	o.FacilityCodeListProvider = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *Location) GetAddress() Address {
	if o == nil || o.Address == nil {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Location) GetAddressOk() (*Address, bool) {
	if o == nil || o.Address == nil {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *Location) HasAddress() bool {
	if o != nil && o.Address != nil {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *Location) SetAddress(v Address) {
	o.Address = &v
}

func (o Location) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.LocationName != nil {
		toSerialize["locationName"] = o.LocationName
	}
	if o.Latitude != nil {
		toSerialize["latitude"] = o.Latitude
	}
	if o.Longitude != nil {
		toSerialize["longitude"] = o.Longitude
	}
	if o.UNLocationCode != nil {
		toSerialize["UNLocationCode"] = o.UNLocationCode
	}
	if o.FacilityCode != nil {
		toSerialize["facilityCode"] = o.FacilityCode
	}
	if o.FacilityCodeListProvider != nil {
		toSerialize["facilityCodeListProvider"] = o.FacilityCodeListProvider
	}
	if o.Address != nil {
		toSerialize["address"] = o.Address
	}
	return json.Marshal(toSerialize)
}

type NullableLocation struct {
	value *Location
	isSet bool
}

func (v NullableLocation) Get() *Location {
	return v.value
}

func (v *NullableLocation) Set(val *Location) {
	v.value = val
	v.isSet = true
}

func (v NullableLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocation(val *Location) *NullableLocation {
	return &NullableLocation{value: val, isSet: true}
}

func (v NullableLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


