# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationName** | Pointer to **string** | The name of the location. | [optional] 
**Latitude** | Pointer to **string** | Geographic coordinate that specifies the north–south position of a point on the Earth&#39;s surface. | [optional] 
**Longitude** | Pointer to **string** | Geographic coordinate that specifies the east–west position of a point on the Earth&#39;s surface. | [optional] 
**UNLocationCode** | Pointer to **string** | The UN Location code specifying where the place is located. | [optional] 
**FacilityCode** | Pointer to **string** | The code used for identifying the specific facility. This code is not the UN Location Code. | [optional] 
**FacilityCodeListProvider** | Pointer to **string** | The provider used for identifying the facility Code | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationName

`func (o *Location) GetLocationName() string`

GetLocationName returns the LocationName field if non-nil, zero value otherwise.

### GetLocationNameOk

`func (o *Location) GetLocationNameOk() (*string, bool)`

GetLocationNameOk returns a tuple with the LocationName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationName

`func (o *Location) SetLocationName(v string)`

SetLocationName sets LocationName field to given value.

### HasLocationName

`func (o *Location) HasLocationName() bool`

HasLocationName returns a boolean if a field has been set.

### GetLatitude

`func (o *Location) GetLatitude() string`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *Location) GetLatitudeOk() (*string, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *Location) SetLatitude(v string)`

SetLatitude sets Latitude field to given value.

### HasLatitude

`func (o *Location) HasLatitude() bool`

HasLatitude returns a boolean if a field has been set.

### GetLongitude

`func (o *Location) GetLongitude() string`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *Location) GetLongitudeOk() (*string, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *Location) SetLongitude(v string)`

SetLongitude sets Longitude field to given value.

### HasLongitude

`func (o *Location) HasLongitude() bool`

HasLongitude returns a boolean if a field has been set.

### GetUNLocationCode

`func (o *Location) GetUNLocationCode() string`

GetUNLocationCode returns the UNLocationCode field if non-nil, zero value otherwise.

### GetUNLocationCodeOk

`func (o *Location) GetUNLocationCodeOk() (*string, bool)`

GetUNLocationCodeOk returns a tuple with the UNLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLocationCode

`func (o *Location) SetUNLocationCode(v string)`

SetUNLocationCode sets UNLocationCode field to given value.

### HasUNLocationCode

`func (o *Location) HasUNLocationCode() bool`

HasUNLocationCode returns a boolean if a field has been set.

### GetFacilityCode

`func (o *Location) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *Location) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *Location) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.

### HasFacilityCode

`func (o *Location) HasFacilityCode() bool`

HasFacilityCode returns a boolean if a field has been set.

### GetFacilityCodeListProvider

`func (o *Location) GetFacilityCodeListProvider() string`

GetFacilityCodeListProvider returns the FacilityCodeListProvider field if non-nil, zero value otherwise.

### GetFacilityCodeListProviderOk

`func (o *Location) GetFacilityCodeListProviderOk() (*string, bool)`

GetFacilityCodeListProviderOk returns a tuple with the FacilityCodeListProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCodeListProvider

`func (o *Location) SetFacilityCodeListProvider(v string)`

SetFacilityCodeListProvider sets FacilityCodeListProvider field to given value.

### HasFacilityCodeListProvider

`func (o *Location) HasFacilityCodeListProvider() bool`

HasFacilityCodeListProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


