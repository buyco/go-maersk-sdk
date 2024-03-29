/*
Track & Trace Events

Retrieve Track & Trace Events based on DCSA Interface standard v.2.2  This service provides shippers and consignees visibility to Shipment, Equipment and Transport events for shipments booked with A.P. Moller-Maersk A/S using standards set by the Digital Container Shipping Association.\\ <https://dcsa.org/>

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package maersk

import (
	"encoding/json"
	"time"
)

// TransportEvent The transport event entity is a specialization of the event entity to support specification of data that only applies to a transport event.
type TransportEvent struct {
	// The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID.
	EventID   *string `json:"eventID,omitempty"`
	EventType string  `json:"eventType"`
	// The local date and time, where the event took place or when the event will take place, in ISO 8601 format. For Shipment Event, it is the same as eventCreatedDateTime in UTC timezone.
	EventDateTime string `json:"eventDateTime"`
	// The UTC timestamp of when the event was created.
	EventCreatedDateTime time.Time `json:"eventCreatedDateTime"`
	// Code for the event classifier, either PLN, ACT or EST. * PLN - Planned * ACT - Actual * EST - Estimated
	EventClassifierCode string `json:"eventClassifierCode"`
	// References provided by the shipper or freight forwarder at the time of booking or at the time of providing shipping instruction. Carriers share it back when providing track and trace event updates, some are also printed on the B/L. Customers can use these references to track shipments in their internal systems.
	References []EventReferencesInner `json:"references,omitempty"`
	// Identifier for type of Transport event - ARRI (Arrived) - DEPA (Departed)
	TransportEventTypeCode *string `json:"transportEventTypeCode,omitempty"`
	// Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/
	DelayReasonCode *string `json:"delayReasonCode,omitempty"`
	// Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage.
	ChangeRemark *string `json:"changeRemark,omitempty"`
	// An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The documentReferenceType-field is used to describe where the documentReferenceValue-field is pointing to.
	DocumentReferences []DocumentReferencesInner `json:"documentReferences,omitempty"`
	TransportCall      TransportCall             `json:"transportCall"`
}

// NewTransportEvent instantiates a new TransportEvent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransportEvent(eventType string, eventDateTime string, eventCreatedDateTime time.Time, eventClassifierCode string, transportCall TransportCall) *TransportEvent {
	this := TransportEvent{}
	this.EventType = eventType
	this.EventDateTime = eventDateTime
	this.EventCreatedDateTime = eventCreatedDateTime
	this.EventClassifierCode = eventClassifierCode
	this.TransportCall = transportCall
	return &this
}

// NewTransportEventWithDefaults instantiates a new TransportEvent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransportEventWithDefaults() *TransportEvent {
	this := TransportEvent{}
	return &this
}

// GetEventID returns the EventID field value if set, zero value otherwise.
func (o *TransportEvent) GetEventID() string {
	if o == nil || o.EventID == nil {
		var ret string
		return ret
	}
	return *o.EventID
}

// GetEventIDOk returns a tuple with the EventID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventIDOk() (*string, bool) {
	if o == nil || o.EventID == nil {
		return nil, false
	}
	return o.EventID, true
}

// HasEventID returns a boolean if a field has been set.
func (o *TransportEvent) HasEventID() bool {
	if o != nil && o.EventID != nil {
		return true
	}

	return false
}

// SetEventID gets a reference to the given string and assigns it to the EventID field.
func (o *TransportEvent) SetEventID(v string) {
	o.EventID = &v
}

// GetEventType returns the EventType field value
func (o *TransportEvent) GetEventType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventType, true
}

// SetEventType sets field value
func (o *TransportEvent) SetEventType(v string) {
	o.EventType = v
}

// GetEventDateTime returns the EventDateTime field value
func (o *TransportEvent) GetEventDateTime() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventDateTime
}

// GetEventDateTimeOk returns a tuple with the EventDateTime field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventDateTimeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventDateTime, true
}

// SetEventDateTime sets field value
func (o *TransportEvent) SetEventDateTime(v string) {
	o.EventDateTime = v
}

// GetEventCreatedDateTime returns the EventCreatedDateTime field value
func (o *TransportEvent) GetEventCreatedDateTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.EventCreatedDateTime
}

// GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventCreatedDateTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventCreatedDateTime, true
}

// SetEventCreatedDateTime sets field value
func (o *TransportEvent) SetEventCreatedDateTime(v time.Time) {
	o.EventCreatedDateTime = v
}

// GetEventClassifierCode returns the EventClassifierCode field value
func (o *TransportEvent) GetEventClassifierCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventClassifierCode
}

// GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetEventClassifierCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventClassifierCode, true
}

// SetEventClassifierCode sets field value
func (o *TransportEvent) SetEventClassifierCode(v string) {
	o.EventClassifierCode = v
}

// GetReferences returns the References field value if set, zero value otherwise.
func (o *TransportEvent) GetReferences() []EventReferencesInner {
	if o == nil || o.References == nil {
		var ret []EventReferencesInner
		return ret
	}
	return o.References
}

// GetReferencesOk returns a tuple with the References field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetReferencesOk() ([]EventReferencesInner, bool) {
	if o == nil || o.References == nil {
		return nil, false
	}
	return o.References, true
}

// HasReferences returns a boolean if a field has been set.
func (o *TransportEvent) HasReferences() bool {
	if o != nil && o.References != nil {
		return true
	}

	return false
}

// SetReferences gets a reference to the given []EventReferencesInner and assigns it to the References field.
func (o *TransportEvent) SetReferences(v []EventReferencesInner) {
	o.References = v
}

// GetTransportEventTypeCode returns the TransportEventTypeCode field value if set, zero value otherwise.
func (o *TransportEvent) GetTransportEventTypeCode() string {
	if o == nil || o.TransportEventTypeCode == nil {
		var ret string
		return ret
	}
	return *o.TransportEventTypeCode
}

// GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetTransportEventTypeCodeOk() (*string, bool) {
	if o == nil || o.TransportEventTypeCode == nil {
		return nil, false
	}
	return o.TransportEventTypeCode, true
}

// HasTransportEventTypeCode returns a boolean if a field has been set.
func (o *TransportEvent) HasTransportEventTypeCode() bool {
	if o != nil && o.TransportEventTypeCode != nil {
		return true
	}

	return false
}

// SetTransportEventTypeCode gets a reference to the given string and assigns it to the TransportEventTypeCode field.
func (o *TransportEvent) SetTransportEventTypeCode(v string) {
	o.TransportEventTypeCode = &v
}

// GetDelayReasonCode returns the DelayReasonCode field value if set, zero value otherwise.
func (o *TransportEvent) GetDelayReasonCode() string {
	if o == nil || o.DelayReasonCode == nil {
		var ret string
		return ret
	}
	return *o.DelayReasonCode
}

// GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetDelayReasonCodeOk() (*string, bool) {
	if o == nil || o.DelayReasonCode == nil {
		return nil, false
	}
	return o.DelayReasonCode, true
}

// HasDelayReasonCode returns a boolean if a field has been set.
func (o *TransportEvent) HasDelayReasonCode() bool {
	if o != nil && o.DelayReasonCode != nil {
		return true
	}

	return false
}

// SetDelayReasonCode gets a reference to the given string and assigns it to the DelayReasonCode field.
func (o *TransportEvent) SetDelayReasonCode(v string) {
	o.DelayReasonCode = &v
}

// GetChangeRemark returns the ChangeRemark field value if set, zero value otherwise.
func (o *TransportEvent) GetChangeRemark() string {
	if o == nil || o.ChangeRemark == nil {
		var ret string
		return ret
	}
	return *o.ChangeRemark
}

// GetChangeRemarkOk returns a tuple with the ChangeRemark field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetChangeRemarkOk() (*string, bool) {
	if o == nil || o.ChangeRemark == nil {
		return nil, false
	}
	return o.ChangeRemark, true
}

// HasChangeRemark returns a boolean if a field has been set.
func (o *TransportEvent) HasChangeRemark() bool {
	if o != nil && o.ChangeRemark != nil {
		return true
	}

	return false
}

// SetChangeRemark gets a reference to the given string and assigns it to the ChangeRemark field.
func (o *TransportEvent) SetChangeRemark(v string) {
	o.ChangeRemark = &v
}

// GetDocumentReferences returns the DocumentReferences field value if set, zero value otherwise.
func (o *TransportEvent) GetDocumentReferences() []DocumentReferencesInner {
	if o == nil || o.DocumentReferences == nil {
		var ret []DocumentReferencesInner
		return ret
	}
	return o.DocumentReferences
}

// GetDocumentReferencesOk returns a tuple with the DocumentReferences field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetDocumentReferencesOk() ([]DocumentReferencesInner, bool) {
	if o == nil || o.DocumentReferences == nil {
		return nil, false
	}
	return o.DocumentReferences, true
}

// HasDocumentReferences returns a boolean if a field has been set.
func (o *TransportEvent) HasDocumentReferences() bool {
	if o != nil && o.DocumentReferences != nil {
		return true
	}

	return false
}

// SetDocumentReferences gets a reference to the given []DocumentReferencesInner and assigns it to the DocumentReferences field.
func (o *TransportEvent) SetDocumentReferences(v []DocumentReferencesInner) {
	o.DocumentReferences = v
}

// GetTransportCall returns the TransportCall field value
func (o *TransportEvent) GetTransportCall() TransportCall {
	if o == nil {
		var ret TransportCall
		return ret
	}

	return o.TransportCall
}

// GetTransportCallOk returns a tuple with the TransportCall field value
// and a boolean to check if the value has been set.
func (o *TransportEvent) GetTransportCallOk() (*TransportCall, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TransportCall, true
}

// SetTransportCall sets field value
func (o *TransportEvent) SetTransportCall(v TransportCall) {
	o.TransportCall = v
}

func (o TransportEvent) MarshalJSON() ([]byte, error) {
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
	if o.TransportEventTypeCode != nil {
		toSerialize["transportEventTypeCode"] = o.TransportEventTypeCode
	}
	if o.DelayReasonCode != nil {
		toSerialize["delayReasonCode"] = o.DelayReasonCode
	}
	if o.ChangeRemark != nil {
		toSerialize["changeRemark"] = o.ChangeRemark
	}
	if o.DocumentReferences != nil {
		toSerialize["documentReferences"] = o.DocumentReferences
	}
	if true {
		toSerialize["transportCall"] = o.TransportCall
	}
	return json.Marshal(toSerialize)
}

type NullableTransportEvent struct {
	value *TransportEvent
	isSet bool
}

func (v NullableTransportEvent) Get() *TransportEvent {
	return v.value
}

func (v *NullableTransportEvent) Set(val *TransportEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTransportEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTransportEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransportEvent(val *TransportEvent) *NullableTransportEvent {
	return &NullableTransportEvent{value: val, isSet: true}
}

func (v NullableTransportEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransportEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
