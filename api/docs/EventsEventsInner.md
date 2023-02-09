# EventsEventsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID. | [optional] 
**EventType** | **string** |  | 
**EventDateTime** | **time.Time** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. For Shipment Event, it is the same as eventCreatedDateTime in UTC timezone. | 
**EventCreatedDateTime** | **time.Time** | The UTC timestamp of when the event was created. | 
**EventClassifierCode** | **string** | Code for the event classifier, either PLN, ACT or EST. * PLN - Planned * ACT - Actual * EST - Estimated  | 
**References** | Pointer to [**[]EventReferencesInner**](EventReferencesInner.md) | References provided by the shipper or freight forwarder at the time of booking or at the time of providing shipping instruction. Carriers share it back when providing track and trace event updates, some are also printed on the B/L. Customers can use these references to track shipments in their internal systems. | [optional] 
**ShipmentEventTypeCode** | **string** | The status of the document in the process. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released)  Note: Version 1.1 replaces CONF (Confirmed) for RELS (Released) for documentTypeCode SRM (Shipment Release Message).  | 
**DocumentTypeCode** | **string** | The code to identify the type of information documentID points to. Can be one of the following values * CBR (Carrier Booking Request Reference) * BKG (Booking) * SHI (Shipping Instruction) * SRM (Shipment Release Message) * TRD (Transport Document) * ARN (Arrival Notice) * VGM (Verified Gross Mass) * CAS (Cargo Survey) * CUS (Customs Inspection) * DGD (Dangerous Goods Declaration) * OOG (Out of Gauge)  | 
**DocumentID** | **string** | The ID of the object defined by the Shipment Information Type. In some cases this is a UUID; in other cases this is a string.  | 
**Reason** | Pointer to **string** | Reason field in a Shipment event. This field can be used to explain why a specific event has been sent. | [optional] 
**TransportEventTypeCode** | Pointer to **string** | Identifier for type of Transport event - ARRI (Arrived) - DEPA (Departed)  | [optional] 
**DelayReasonCode** | Pointer to **string** | Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/  | [optional] 
**ChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage. | [optional] 
**DocumentReferences** | Pointer to [**[]DocumentReferencesInner**](DocumentReferencesInner.md) | An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The documentReferenceType-field is used to describe where the documentReferenceValue-field is pointing to. | [optional] 
**TransportCall** | [**TransportCall**](TransportCall.md) |  | 
**EquipmentEventTypeCode** | Pointer to **string** | Unique identifier for equipmentEventTypeCode. * LOAD (Loaded) * DISC (Discharged) * GTIN (Gated in) * GTOT (Gated out) * STUF (Stuffed) * STRP (Stripped) * PICK (Pick-up) * DROP (Drop-off) * RSEA (Resealed) * RMVD (Removed) * INSP (Inspected)  | [optional] 
**EquipmentReference** | Pointer to **string** | The unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. According to ISO 6346, a container identification code consists of a 4-letter prefix and a 7-digit number (composed of a 3-letter owner code, a category identifier, a serial number, and a check-digit). If a container does not comply with ISO 6346, it is suggested to follow Recommendation #2 “Container with non-ISO identification” from SMDG.  | [optional] 
**ISOEquipmentCode** | Pointer to **string** | Unique code for the different equipment size/type used for transporting commodities. The code is a concatenation of ISO Equipment Size Code and ISO Equipment Type Code A and follows the ISO 6346 standard. | [optional] 
**EmptyIndicatorCode** | **string** | Code to denote whether the equipment is empty or laden. | 
**EventLocation** | Pointer to [**Location**](Location.md) |  | [optional] 
**Seals** | Pointer to [**[]Seal**](Seal.md) |  | [optional] 

## Methods

### NewEventsEventsInner

`func NewEventsEventsInner(eventType string, eventDateTime time.Time, eventCreatedDateTime time.Time, eventClassifierCode string, shipmentEventTypeCode string, documentTypeCode string, documentID string, transportCall TransportCall, emptyIndicatorCode string, ) *EventsEventsInner`

NewEventsEventsInner instantiates a new EventsEventsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventsEventsInnerWithDefaults

`func NewEventsEventsInnerWithDefaults() *EventsEventsInner`

NewEventsEventsInnerWithDefaults instantiates a new EventsEventsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *EventsEventsInner) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *EventsEventsInner) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *EventsEventsInner) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *EventsEventsInner) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventType

`func (o *EventsEventsInner) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *EventsEventsInner) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *EventsEventsInner) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventDateTime

`func (o *EventsEventsInner) GetEventDateTime() time.Time`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *EventsEventsInner) GetEventDateTimeOk() (*time.Time, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *EventsEventsInner) SetEventDateTime(v time.Time)`

SetEventDateTime sets EventDateTime field to given value.


### GetEventCreatedDateTime

`func (o *EventsEventsInner) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *EventsEventsInner) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *EventsEventsInner) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventClassifierCode

`func (o *EventsEventsInner) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *EventsEventsInner) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *EventsEventsInner) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetReferences

`func (o *EventsEventsInner) GetReferences() []EventReferencesInner`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *EventsEventsInner) GetReferencesOk() (*[]EventReferencesInner, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *EventsEventsInner) SetReferences(v []EventReferencesInner)`

SetReferences sets References field to given value.

### HasReferences

`func (o *EventsEventsInner) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetShipmentEventTypeCode

`func (o *EventsEventsInner) GetShipmentEventTypeCode() string`

GetShipmentEventTypeCode returns the ShipmentEventTypeCode field if non-nil, zero value otherwise.

### GetShipmentEventTypeCodeOk

`func (o *EventsEventsInner) GetShipmentEventTypeCodeOk() (*string, bool)`

GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentEventTypeCode

`func (o *EventsEventsInner) SetShipmentEventTypeCode(v string)`

SetShipmentEventTypeCode sets ShipmentEventTypeCode field to given value.


### GetDocumentTypeCode

`func (o *EventsEventsInner) GetDocumentTypeCode() string`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *EventsEventsInner) GetDocumentTypeCodeOk() (*string, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *EventsEventsInner) SetDocumentTypeCode(v string)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetDocumentID

`func (o *EventsEventsInner) GetDocumentID() string`

GetDocumentID returns the DocumentID field if non-nil, zero value otherwise.

### GetDocumentIDOk

`func (o *EventsEventsInner) GetDocumentIDOk() (*string, bool)`

GetDocumentIDOk returns a tuple with the DocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentID

`func (o *EventsEventsInner) SetDocumentID(v string)`

SetDocumentID sets DocumentID field to given value.


### GetReason

`func (o *EventsEventsInner) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *EventsEventsInner) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *EventsEventsInner) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *EventsEventsInner) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetTransportEventTypeCode

`func (o *EventsEventsInner) GetTransportEventTypeCode() string`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *EventsEventsInner) GetTransportEventTypeCodeOk() (*string, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *EventsEventsInner) SetTransportEventTypeCode(v string)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.

### HasTransportEventTypeCode

`func (o *EventsEventsInner) HasTransportEventTypeCode() bool`

HasTransportEventTypeCode returns a boolean if a field has been set.

### GetDelayReasonCode

`func (o *EventsEventsInner) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *EventsEventsInner) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *EventsEventsInner) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *EventsEventsInner) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetChangeRemark

`func (o *EventsEventsInner) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *EventsEventsInner) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *EventsEventsInner) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *EventsEventsInner) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetDocumentReferences

`func (o *EventsEventsInner) GetDocumentReferences() []DocumentReferencesInner`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *EventsEventsInner) GetDocumentReferencesOk() (*[]DocumentReferencesInner, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *EventsEventsInner) SetDocumentReferences(v []DocumentReferencesInner)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *EventsEventsInner) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetTransportCall

`func (o *EventsEventsInner) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *EventsEventsInner) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *EventsEventsInner) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.


### GetEquipmentEventTypeCode

`func (o *EventsEventsInner) GetEquipmentEventTypeCode() string`

GetEquipmentEventTypeCode returns the EquipmentEventTypeCode field if non-nil, zero value otherwise.

### GetEquipmentEventTypeCodeOk

`func (o *EventsEventsInner) GetEquipmentEventTypeCodeOk() (*string, bool)`

GetEquipmentEventTypeCodeOk returns a tuple with the EquipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentEventTypeCode

`func (o *EventsEventsInner) SetEquipmentEventTypeCode(v string)`

SetEquipmentEventTypeCode sets EquipmentEventTypeCode field to given value.

### HasEquipmentEventTypeCode

`func (o *EventsEventsInner) HasEquipmentEventTypeCode() bool`

HasEquipmentEventTypeCode returns a boolean if a field has been set.

### GetEquipmentReference

`func (o *EventsEventsInner) GetEquipmentReference() string`

GetEquipmentReference returns the EquipmentReference field if non-nil, zero value otherwise.

### GetEquipmentReferenceOk

`func (o *EventsEventsInner) GetEquipmentReferenceOk() (*string, bool)`

GetEquipmentReferenceOk returns a tuple with the EquipmentReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquipmentReference

`func (o *EventsEventsInner) SetEquipmentReference(v string)`

SetEquipmentReference sets EquipmentReference field to given value.

### HasEquipmentReference

`func (o *EventsEventsInner) HasEquipmentReference() bool`

HasEquipmentReference returns a boolean if a field has been set.

### GetISOEquipmentCode

`func (o *EventsEventsInner) GetISOEquipmentCode() string`

GetISOEquipmentCode returns the ISOEquipmentCode field if non-nil, zero value otherwise.

### GetISOEquipmentCodeOk

`func (o *EventsEventsInner) GetISOEquipmentCodeOk() (*string, bool)`

GetISOEquipmentCodeOk returns a tuple with the ISOEquipmentCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetISOEquipmentCode

`func (o *EventsEventsInner) SetISOEquipmentCode(v string)`

SetISOEquipmentCode sets ISOEquipmentCode field to given value.

### HasISOEquipmentCode

`func (o *EventsEventsInner) HasISOEquipmentCode() bool`

HasISOEquipmentCode returns a boolean if a field has been set.

### GetEmptyIndicatorCode

`func (o *EventsEventsInner) GetEmptyIndicatorCode() string`

GetEmptyIndicatorCode returns the EmptyIndicatorCode field if non-nil, zero value otherwise.

### GetEmptyIndicatorCodeOk

`func (o *EventsEventsInner) GetEmptyIndicatorCodeOk() (*string, bool)`

GetEmptyIndicatorCodeOk returns a tuple with the EmptyIndicatorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmptyIndicatorCode

`func (o *EventsEventsInner) SetEmptyIndicatorCode(v string)`

SetEmptyIndicatorCode sets EmptyIndicatorCode field to given value.


### GetEventLocation

`func (o *EventsEventsInner) GetEventLocation() Location`

GetEventLocation returns the EventLocation field if non-nil, zero value otherwise.

### GetEventLocationOk

`func (o *EventsEventsInner) GetEventLocationOk() (*Location, bool)`

GetEventLocationOk returns a tuple with the EventLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventLocation

`func (o *EventsEventsInner) SetEventLocation(v Location)`

SetEventLocation sets EventLocation field to given value.

### HasEventLocation

`func (o *EventsEventsInner) HasEventLocation() bool`

HasEventLocation returns a boolean if a field has been set.

### GetSeals

`func (o *EventsEventsInner) GetSeals() []Seal`

GetSeals returns the Seals field if non-nil, zero value otherwise.

### GetSealsOk

`func (o *EventsEventsInner) GetSealsOk() (*[]Seal, bool)`

GetSealsOk returns a tuple with the Seals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeals

`func (o *EventsEventsInner) SetSeals(v []Seal)`

SetSeals sets Seals field to given value.

### HasSeals

`func (o *EventsEventsInner) HasSeals() bool`

HasSeals returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


