/*
Track & Trace Events

Retrieve Track & Trace Events based on DCSA Interface standard v.2.2  This service provides shippers and consignees visibility to Shipment, Equipment and Transport events for shipments booked with A.P. Moller-Maersk A/S using standards set by the Digital Container Shipping Association.\\ <https://dcsa.org/>

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package api

import (
	"encoding/json"
	"time"
)

// ShipmentEvent The shipment event entity is a specialization of the event entity to support specification of data that only applies to a shipment event.
type ShipmentEvent struct {
	// The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID.
	EventID   *string `json:"eventID,omitempty"`
	EventType string  `json:"eventType"`
	// The local date and time, where the event took place or when the event will take place, in ISO 8601 format. For Shipment Event, it is the same as eventCreatedDateTime in UTC timezone.
	EventDateTime time.Time `json:"eventDateTime"`
	// The UTC timestamp of when the event was created.
	EventCreatedDateTime time.Time `json:"eventCreatedDateTime"`
	// Code for the event classifier, either PLN, ACT or EST. * PLN - Planned * ACT - Actual * EST - Estimated
	EventClassifierCode string `json:"eventClassifierCode"`
	// References provided by the shipper or freight forwarder at the time of booking or at the time of providing shipping instruction. Carriers share it back when providing track and trace event updates, some are also printed on the B/L. Customers can use these references to track shipments in their internal systems.
	References []EventReferencesInner `json:"references,omitempty"`
	// The status of the document in the process. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released)  Note: Version 1.1 replaces CONF (Confirmed) for RELS (Released) for documentTypeCode SRM (Shipment Release Message).
	ShipmentEventTypeCode string `json:"shipmentEventTypeCode"`
	// The code to identify the type of information documentID points to. Can be one of the following values * CBR (Carrier Booking Request Reference) * BKG (Booking) * SHI (Shipping Instruction) * SRM (Shipment Release Message) * TRD (Transport Document) * ARN (Arrival Notice) * VGM (Verified Gross Mass) * CAS (Cargo Survey) * CUS (Customs Inspection) * DGD (Dangerous Goods Declaration) * OOG (Out of Gauge)
	DocumentTypeCode string `json:"documentTypeCode"`
	// The ID of the object defined by the Shipment Information Type. In some cases this is a UUID; in other cases this is a string.
	DocumentID string `json:"documentID"`
	// Reason field in a Shipment event. This field can be used to explain why a specific event has been sent.
	Reason *string `json:"reason,omitempty"`
}

// NewShipmentEvent instantiates a new ShipmentEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentEvent(eventType string, eventDateTime time.Time, eventCreatedDateTime time.Time, eventClassifierCode string, shipmentEventTypeCode string, documentTypeCode string, documentID string) *ShipmentEvent {
	this := ShipmentEvent{}
	this.EventType = eventType
	this.EventDateTime = eventDateTime
	this.EventCreatedDateTime = eventCreatedDateTime
	this.EventClassifierCode = eventClassifierCode
	this.ShipmentEventTypeCode = shipmentEventTypeCode
	this.DocumentTypeCode = documentTypeCode
	this.DocumentID = documentID
	return &this
}

// NewShipmentEventWithDefaults instantiates a new ShipmentEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentEventWithDefaults() *ShipmentEvent {
	this := ShipmentEvent{}
	return &this
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *ShipmentEvent) GetEventID() string {
	if o == nil || o.EventID == nil {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventIDOk() (*string, bool) {
	if o == nil || o.EventID == nil {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *ShipmentEvent) HasEventID() bool {
	if o != nil && o.EventID != nil {
		return true
	}

	return false
}

// SetEventID gets a reference to the given string and assigns it to the EventID field.
func (o *ShipmentEvent) SetEventID(v string) {
	o.EventID = &v
}

// GetEventType returns the EventType field value
func (o *ShipmentEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *ShipmentEvent) SetEventType(v string) {
	o.EventType = v
}

// GetEventDateTime returns the EventDateTime field value
func (o *ShipmentEvent) GetEventDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *ShipmentEvent) SetEventDateTime(v time.Time) {
	o.EventDateTime = v
}

// GetEventCreatedDateTime returns the EventCreatedDateTime field value
func (o *ShipmentEvent) GetEventCreatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventCreatedDateTime
}

// GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCreatedDateTime, true
}

// SetEventCreatedDateTime sets field value
func (o *ShipmentEvent) SetEventCreatedDateTime(v time.Time) {
	o.EventCreatedDateTime = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *ShipmentEvent) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *ShipmentEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *ShipmentEvent) GetReferences() []EventReferencesInner {
	if o == nil || o.References == nil {
		var ret []EventReferencesInner
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetReferencesOk() ([]EventReferencesInner, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *ShipmentEvent) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []EventReferencesInner and assigns it to the References field.
func (o *ShipmentEvent) SetReferences(v []EventReferencesInner) {
	o.References = v
}

// GetShipmentEventTypeCode returns the ShipmentEventTypeCode field value
func (o *ShipmentEvent) GetShipmentEventTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentEventTypeCode
}

// GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetShipmentEventTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentEventTypeCode, true
}

// SetShipmentEventTypeCode sets field value
func (o *ShipmentEvent) SetShipmentEventTypeCode(v string) {
	o.ShipmentEventTypeCode = v
}

// GetDocumentTypeCode returns the DocumentTypeCode field value
func (o *ShipmentEvent) GetDocumentTypeCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentTypeCode
}

// GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetDocumentTypeCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentTypeCode, true
}

// SetDocumentTypeCode sets field value
func (o *ShipmentEvent) SetDocumentTypeCode(v string) {
	o.DocumentTypeCode = v
}

// GetDocumentID returns the DocumentID field value
func (o *ShipmentEvent) GetDocumentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DocumentID
}

// GetDocumentIDOk returns a tuple with the DocumentID field value
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetDocumentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DocumentID, true
}

// SetDocumentID sets field value
func (o *ShipmentEvent) SetDocumentID(v string) {
	o.DocumentID = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *ShipmentEvent) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentEvent) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *ShipmentEvent) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *ShipmentEvent) SetReason(v string) {
	o.Reason = &v
}

func (o ShipmentEvent) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.EventID != nil {
		toSerialize["eventID"] = o.EventID
	}
	if true {
		toSerialize["eventType"] = o.EventType
	}
	if true {
		toSerialize["eventDateTime"] = o.EventDateTime
	}
	if true {
		toSerialize["eventCreatedDateTime"] = o.EventCreatedDateTime
	}
	if true {
		toSerialize["eventClassifierCode"] = o.EventClassifierCode
	}
	if o.References != nil {
		toSerialize["references"] = o.References
	}
	if true {
		toSerialize["shipmentEventTypeCode"] = o.ShipmentEventTypeCode
	}
	if true {
		toSerialize["documentTypeCode"] = o.DocumentTypeCode
	}
	if true {
		toSerialize["documentID"] = o.DocumentID
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableShipmentEvent struct {
	value *ShipmentEvent
	isSet bool
}

func (v NullableShipmentEvent) Get() *ShipmentEvent {
	return v.value
}

func (v *NullableShipmentEvent) Set(val *ShipmentEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentEvent(val *ShipmentEvent) *NullableShipmentEvent {
	return &NullableShipmentEvent{value: val, isSet: true}
}

func (v NullableShipmentEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}