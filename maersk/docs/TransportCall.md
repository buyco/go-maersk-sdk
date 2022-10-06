# TransportCall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransportCallID** | Pointer to **string** | The unique identifier for a transport call | [optional] 
**CarrierServiceCode** | Pointer to **string** | The code of the service for which the schedule details are published. | [optional] 
**CarrierVoyageNumber** | Pointer to **string** | \\*\\*\\*do not use\\*\\*\\* this field is replaced by &#x60;exportVoyageNumber&#x60; from Version 1.1 and will be removed by the API from next major version  The vessel operator-specific identifier of the Voyage.\\ The export voyage is returned when there are multiple voyages.  | [optional] 
**ExportVoyageNumber** | Pointer to **string** | The vessel operator-specific identifier of the export Voyage. | [optional] 
**ImportVoyageNumber** | Pointer to **string** | The vessel operator-specific identifier of the import Voyage. | [optional] 
**TransportCallSequenceNumber** | Pointer to **int32** | Transport operator&#39;s key that uniquely identifies each individual call. This key is essential to distinguish between two separate calls at the same location within one voyage. | [optional] 
**UNLocationCode** | Pointer to **string** | The UN Location code specifying where the place is located. | [optional] 
**FacilityCode** | Pointer to **string** | The code used for identifying the specific facility. This code is not the UN Location Code. | [optional] 
**FacilityCodeListProvider** | Pointer to **string** | The provider used for identifying the facility Code | [optional] 
**FacilityTypeCode** | Pointer to **string** | A specialized version of the facilityCode to be used in TransportCalls. The code to identify the specific type of facility. * BOCR (Border crossing) * CLOC (Customer location) * COFS (Container freight station) * COYA (Deprecated - use OFFD intead) * OFFD (Off dock storage) * DEPO (Depot) * INTE (Inland terminal) * POTE (Port terminal) * RAMP (Ramp)  | [optional] 
**OtherFacility** | Pointer to **string** | An alternative way to capture the facility when no standardized DCSA facility code can be found. | [optional] 
**ModeOfTransport** | **string** | The code specifying the mode of transport. | 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Vessel** | Pointer to [**Vessel**](Vessel.md) |  | [optional] 

## Methods

### NewTransportCall

`func NewTransportCall(modeOfTransport string, ) *TransportCall`

NewTransportCall instantiates a new TransportCall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportCallWithDefaults

`func NewTransportCallWithDefaults() *TransportCall`

NewTransportCallWithDefaults instantiates a new TransportCall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransportCallID

`func (o *TransportCall) GetTransportCallID() string`

GetTransportCallID returns the TransportCallID field if non-nil, zero value otherwise.

### GetTransportCallIDOk

`func (o *TransportCall) GetTransportCallIDOk() (*string, bool)`

GetTransportCallIDOk returns a tuple with the TransportCallID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCallID

`func (o *TransportCall) SetTransportCallID(v string)`

SetTransportCallID sets TransportCallID field to given value.

### HasTransportCallID

`func (o *TransportCall) HasTransportCallID() bool`

HasTransportCallID returns a boolean if a field has been set.

### GetCarrierServiceCode

`func (o *TransportCall) GetCarrierServiceCode() string`

GetCarrierServiceCode returns the CarrierServiceCode field if non-nil, zero value otherwise.

### GetCarrierServiceCodeOk

`func (o *TransportCall) GetCarrierServiceCodeOk() (*string, bool)`

GetCarrierServiceCodeOk returns a tuple with the CarrierServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierServiceCode

`func (o *TransportCall) SetCarrierServiceCode(v string)`

SetCarrierServiceCode sets CarrierServiceCode field to given value.

### HasCarrierServiceCode

`func (o *TransportCall) HasCarrierServiceCode() bool`

HasCarrierServiceCode returns a boolean if a field has been set.

### GetCarrierVoyageNumber

`func (o *TransportCall) GetCarrierVoyageNumber() string`

GetCarrierVoyageNumber returns the CarrierVoyageNumber field if non-nil, zero value otherwise.

### GetCarrierVoyageNumberOk

`func (o *TransportCall) GetCarrierVoyageNumberOk() (*string, bool)`

GetCarrierVoyageNumberOk returns a tuple with the CarrierVoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierVoyageNumber

`func (o *TransportCall) SetCarrierVoyageNumber(v string)`

SetCarrierVoyageNumber sets CarrierVoyageNumber field to given value.

### HasCarrierVoyageNumber

`func (o *TransportCall) HasCarrierVoyageNumber() bool`

HasCarrierVoyageNumber returns a boolean if a field has been set.

### GetExportVoyageNumber

`func (o *TransportCall) GetExportVoyageNumber() string`

GetExportVoyageNumber returns the ExportVoyageNumber field if non-nil, zero value otherwise.

### GetExportVoyageNumberOk

`func (o *TransportCall) GetExportVoyageNumberOk() (*string, bool)`

GetExportVoyageNumberOk returns a tuple with the ExportVoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportVoyageNumber

`func (o *TransportCall) SetExportVoyageNumber(v string)`

SetExportVoyageNumber sets ExportVoyageNumber field to given value.

### HasExportVoyageNumber

`func (o *TransportCall) HasExportVoyageNumber() bool`

HasExportVoyageNumber returns a boolean if a field has been set.

### GetImportVoyageNumber

`func (o *TransportCall) GetImportVoyageNumber() string`

GetImportVoyageNumber returns the ImportVoyageNumber field if non-nil, zero value otherwise.

### GetImportVoyageNumberOk

`func (o *TransportCall) GetImportVoyageNumberOk() (*string, bool)`

GetImportVoyageNumberOk returns a tuple with the ImportVoyageNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportVoyageNumber

`func (o *TransportCall) SetImportVoyageNumber(v string)`

SetImportVoyageNumber sets ImportVoyageNumber field to given value.

### HasImportVoyageNumber

`func (o *TransportCall) HasImportVoyageNumber() bool`

HasImportVoyageNumber returns a boolean if a field has been set.

### GetTransportCallSequenceNumber

`func (o *TransportCall) GetTransportCallSequenceNumber() int32`

GetTransportCallSequenceNumber returns the TransportCallSequenceNumber field if non-nil, zero value otherwise.

### GetTransportCallSequenceNumberOk

`func (o *TransportCall) GetTransportCallSequenceNumberOk() (*int32, bool)`

GetTransportCallSequenceNumberOk returns a tuple with the TransportCallSequenceNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCallSequenceNumber

`func (o *TransportCall) SetTransportCallSequenceNumber(v int32)`

SetTransportCallSequenceNumber sets TransportCallSequenceNumber field to given value.

### HasTransportCallSequenceNumber

`func (o *TransportCall) HasTransportCallSequenceNumber() bool`

HasTransportCallSequenceNumber returns a boolean if a field has been set.

### GetUNLocationCode

`func (o *TransportCall) GetUNLocationCode() string`

GetUNLocationCode returns the UNLocationCode field if non-nil, zero value otherwise.

### GetUNLocationCodeOk

`func (o *TransportCall) GetUNLocationCodeOk() (*string, bool)`

GetUNLocationCodeOk returns a tuple with the UNLocationCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUNLocationCode

`func (o *TransportCall) SetUNLocationCode(v string)`

SetUNLocationCode sets UNLocationCode field to given value.

### HasUNLocationCode

`func (o *TransportCall) HasUNLocationCode() bool`

HasUNLocationCode returns a boolean if a field has been set.

### GetFacilityCode

`func (o *TransportCall) GetFacilityCode() string`

GetFacilityCode returns the FacilityCode field if non-nil, zero value otherwise.

### GetFacilityCodeOk

`func (o *TransportCall) GetFacilityCodeOk() (*string, bool)`

GetFacilityCodeOk returns a tuple with the FacilityCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCode

`func (o *TransportCall) SetFacilityCode(v string)`

SetFacilityCode sets FacilityCode field to given value.

### HasFacilityCode

`func (o *TransportCall) HasFacilityCode() bool`

HasFacilityCode returns a boolean if a field has been set.

### GetFacilityCodeListProvider

`func (o *TransportCall) GetFacilityCodeListProvider() string`

GetFacilityCodeListProvider returns the FacilityCodeListProvider field if non-nil, zero value otherwise.

### GetFacilityCodeListProviderOk

`func (o *TransportCall) GetFacilityCodeListProviderOk() (*string, bool)`

GetFacilityCodeListProviderOk returns a tuple with the FacilityCodeListProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityCodeListProvider

`func (o *TransportCall) SetFacilityCodeListProvider(v string)`

SetFacilityCodeListProvider sets FacilityCodeListProvider field to given value.

### HasFacilityCodeListProvider

`func (o *TransportCall) HasFacilityCodeListProvider() bool`

HasFacilityCodeListProvider returns a boolean if a field has been set.

### GetFacilityTypeCode

`func (o *TransportCall) GetFacilityTypeCode() string`

GetFacilityTypeCode returns the FacilityTypeCode field if non-nil, zero value otherwise.

### GetFacilityTypeCodeOk

`func (o *TransportCall) GetFacilityTypeCodeOk() (*string, bool)`

GetFacilityTypeCodeOk returns a tuple with the FacilityTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityTypeCode

`func (o *TransportCall) SetFacilityTypeCode(v string)`

SetFacilityTypeCode sets FacilityTypeCode field to given value.

### HasFacilityTypeCode

`func (o *TransportCall) HasFacilityTypeCode() bool`

HasFacilityTypeCode returns a boolean if a field has been set.

### GetOtherFacility

`func (o *TransportCall) GetOtherFacility() string`

GetOtherFacility returns the OtherFacility field if non-nil, zero value otherwise.

### GetOtherFacilityOk

`func (o *TransportCall) GetOtherFacilityOk() (*string, bool)`

GetOtherFacilityOk returns a tuple with the OtherFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherFacility

`func (o *TransportCall) SetOtherFacility(v string)`

SetOtherFacility sets OtherFacility field to given value.

### HasOtherFacility

`func (o *TransportCall) HasOtherFacility() bool`

HasOtherFacility returns a boolean if a field has been set.

### GetModeOfTransport

`func (o *TransportCall) GetModeOfTransport() string`

GetModeOfTransport returns the ModeOfTransport field if non-nil, zero value otherwise.

### GetModeOfTransportOk

`func (o *TransportCall) GetModeOfTransportOk() (*string, bool)`

GetModeOfTransportOk returns a tuple with the ModeOfTransport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModeOfTransport

`func (o *TransportCall) SetModeOfTransport(v string)`

SetModeOfTransport sets ModeOfTransport field to given value.


### GetLocation

`func (o *TransportCall) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *TransportCall) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *TransportCall) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *TransportCall) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetVessel

`func (o *TransportCall) GetVessel() Vessel`

GetVessel returns the Vessel field if non-nil, zero value otherwise.

### GetVesselOk

`func (o *TransportCall) GetVesselOk() (*Vessel, bool)`

GetVesselOk returns a tuple with the Vessel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVessel

`func (o *TransportCall) SetVessel(v Vessel)`

SetVessel sets Vessel field to given value.

### HasVessel

`func (o *TransportCall) HasVessel() bool`

HasVessel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


