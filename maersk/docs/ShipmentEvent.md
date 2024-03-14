# ShipmentEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID. | [optional] 
**EventType** | **string** |  | 
**EventDateTime** | **string** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. For Shipment Event, it is the same as eventCreatedDateTime in UTC timezone. | 
**EventCreatedDateTime** | **time.Time** | The UTC timestamp of when the event was created. | 
**EventClassifierCode** | **string** | Code for the event classifier, either PLN, ACT or EST. * PLN - Planned * ACT - Actual * EST - Estimated  | 
**References** | Pointer to [**[]EventReferencesInner**](EventReferencesInner.md) | References provided by the shipper or freight forwarder at the time of booking or at the time of providing shipping instruction. Carriers share it back when providing track and trace event updates, some are also printed on the B/L. Customers can use these references to track shipments in their internal systems. | [optional] 
**ShipmentEventTypeCode** | **string** | The status of the document in the process. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released)  Note: Version 1.1 replaces CONF (Confirmed) for RELS (Released) for documentTypeCode SRM (Shipment Release Message).  | 
**DocumentTypeCode** | **string** | The code to identify the type of information documentID points to. Can be one of the following values * CBR (Carrier Booking Request Reference) * BKG (Booking) * SHI (Shipping Instruction) * SRM (Shipment Release Message) * TRD (Transport Document) * ARN (Arrival Notice) * VGM (Verified Gross Mass) * CAS (Cargo Survey) * CUS (Customs Inspection) * DGD (Dangerous Goods Declaration) * OOG (Out of Gauge)  | 
**DocumentID** | **string** | The ID of the object defined by the Shipment Information Type. In some cases this is a UUID; in other cases this is a string.  | 
**Reason** | Pointer to **string** | Reason field in a Shipment event. This field can be used to explain why a specific event has been sent. | [optional] 

## Methods

### NewShipmentEvent

`func NewShipmentEvent(eventType string, eventDateTime string, eventCreatedDateTime time.Time, eventClassifierCode string, shipmentEventTypeCode string, documentTypeCode string, documentID string, ) *ShipmentEvent`

NewShipmentEvent instantiates a new ShipmentEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentEventWithDefaults

`func NewShipmentEventWithDefaults() *ShipmentEvent`

NewShipmentEventWithDefaults instantiates a new ShipmentEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *ShipmentEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *ShipmentEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *ShipmentEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *ShipmentEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventType

`func (o *ShipmentEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ShipmentEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ShipmentEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventDateTime

`func (o *ShipmentEvent) GetEventDateTime() string`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *ShipmentEvent) GetEventDateTimeOk() (*string, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *ShipmentEvent) SetEventDateTime(v string)`

SetEventDateTime sets EventDateTime field to given value.


### GetEventCreatedDateTime

`func (o *ShipmentEvent) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *ShipmentEvent) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *ShipmentEvent) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventClassifierCode

`func (o *ShipmentEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *ShipmentEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *ShipmentEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetReferences

`func (o *ShipmentEvent) GetReferences() []EventReferencesInner`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *ShipmentEvent) GetReferencesOk() (*[]EventReferencesInner, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *ShipmentEvent) SetReferences(v []EventReferencesInner)`

SetReferences sets References field to given value.

### HasReferences

`func (o *ShipmentEvent) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetShipmentEventTypeCode

`func (o *ShipmentEvent) GetShipmentEventTypeCode() string`

GetShipmentEventTypeCode returns the ShipmentEventTypeCode field if non-nil, zero value otherwise.

### GetShipmentEventTypeCodeOk

`func (o *ShipmentEvent) GetShipmentEventTypeCodeOk() (*string, bool)`

GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentEventTypeCode

`func (o *ShipmentEvent) SetShipmentEventTypeCode(v string)`

SetShipmentEventTypeCode sets ShipmentEventTypeCode field to given value.


### GetDocumentTypeCode

`func (o *ShipmentEvent) GetDocumentTypeCode() string`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *ShipmentEvent) GetDocumentTypeCodeOk() (*string, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *ShipmentEvent) SetDocumentTypeCode(v string)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetDocumentID

`func (o *ShipmentEvent) GetDocumentID() string`

GetDocumentID returns the DocumentID field if non-nil, zero value otherwise.

### GetDocumentIDOk

`func (o *ShipmentEvent) GetDocumentIDOk() (*string, bool)`

GetDocumentIDOk returns a tuple with the DocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentID

`func (o *ShipmentEvent) SetDocumentID(v string)`

SetDocumentID sets DocumentID field to given value.


### GetReason

`func (o *ShipmentEvent) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ShipmentEvent) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ShipmentEvent) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ShipmentEvent) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


