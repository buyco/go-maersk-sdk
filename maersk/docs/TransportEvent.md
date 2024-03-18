# TransportEvent

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventID** | Pointer to **string** | The unique identifier for the Equipment Event ID/Transport Event ID/Shipment Event ID. | [optional] 
**EventType** | **string** |  | 
**EventDateTime** | **string** | The local date and time, where the event took place or when the event will take place, in ISO 8601 format. For Shipment Event, it is the same as eventCreatedDateTime in UTC timezone. | 
**EventCreatedDateTime** | **time.Time** | The UTC timestamp of when the event was created. | 
**EventClassifierCode** | **string** | Code for the event classifier, either PLN, ACT or EST. * PLN - Planned * ACT - Actual * EST - Estimated  | 
**References** | Pointer to [**[]EventReferencesInner**](EventReferencesInner.md) | References provided by the shipper or freight forwarder at the time of booking or at the time of providing shipping instruction. Carriers share it back when providing track and trace event updates, some are also printed on the B/L. Customers can use these references to track shipments in their internal systems. | [optional] 
**TransportEventTypeCode** | Pointer to **string** | Identifier for type of Transport event - ARRI (Arrived) - DEPA (Departed)  | [optional] 
**DelayReasonCode** | Pointer to **string** | Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/  | [optional] 
**ChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage. | [optional] 
**DocumentReferences** | Pointer to [**[]DocumentReferencesInner**](DocumentReferencesInner.md) | An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The documentReferenceType-field is used to describe where the documentReferenceValue-field is pointing to. | [optional] 
**TransportCall** | [**TransportCall**](TransportCall.md) |  | 

## Methods

### NewTransportEvent

`func NewTransportEvent(eventType string, eventDateTime string, eventCreatedDateTime time.Time, eventClassifierCode string, transportCall TransportCall, ) *TransportEvent`

NewTransportEvent instantiates a new TransportEvent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportEventWithDefaults

`func NewTransportEventWithDefaults() *TransportEvent`

NewTransportEventWithDefaults instantiates a new TransportEvent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventID

`func (o *TransportEvent) GetEventID() string`

GetEventID returns the EventID field if non-nil, zero value otherwise.

### GetEventIDOk

`func (o *TransportEvent) GetEventIDOk() (*string, bool)`

GetEventIDOk returns a tuple with the EventID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventID

`func (o *TransportEvent) SetEventID(v string)`

SetEventID sets EventID field to given value.

### HasEventID

`func (o *TransportEvent) HasEventID() bool`

HasEventID returns a boolean if a field has been set.

### GetEventType

`func (o *TransportEvent) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransportEvent) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransportEvent) SetEventType(v string)`

SetEventType sets EventType field to given value.


### GetEventDateTime

`func (o *TransportEvent) GetEventDateTime() string`

GetEventDateTime returns the EventDateTime field if non-nil, zero value otherwise.

### GetEventDateTimeOk

`func (o *TransportEvent) GetEventDateTimeOk() (*string, bool)`

GetEventDateTimeOk returns a tuple with the EventDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventDateTime

`func (o *TransportEvent) SetEventDateTime(v string)`

SetEventDateTime sets EventDateTime field to given value.


### GetEventCreatedDateTime

`func (o *TransportEvent) GetEventCreatedDateTime() time.Time`

GetEventCreatedDateTime returns the EventCreatedDateTime field if non-nil, zero value otherwise.

### GetEventCreatedDateTimeOk

`func (o *TransportEvent) GetEventCreatedDateTimeOk() (*time.Time, bool)`

GetEventCreatedDateTimeOk returns a tuple with the EventCreatedDateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventCreatedDateTime

`func (o *TransportEvent) SetEventCreatedDateTime(v time.Time)`

SetEventCreatedDateTime sets EventCreatedDateTime field to given value.


### GetEventClassifierCode

`func (o *TransportEvent) GetEventClassifierCode() string`

GetEventClassifierCode returns the EventClassifierCode field if non-nil, zero value otherwise.

### GetEventClassifierCodeOk

`func (o *TransportEvent) GetEventClassifierCodeOk() (*string, bool)`

GetEventClassifierCodeOk returns a tuple with the EventClassifierCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventClassifierCode

`func (o *TransportEvent) SetEventClassifierCode(v string)`

SetEventClassifierCode sets EventClassifierCode field to given value.


### GetReferences

`func (o *TransportEvent) GetReferences() []EventReferencesInner`

GetReferences returns the References field if non-nil, zero value otherwise.

### GetReferencesOk

`func (o *TransportEvent) GetReferencesOk() (*[]EventReferencesInner, bool)`

GetReferencesOk returns a tuple with the References field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferences

`func (o *TransportEvent) SetReferences(v []EventReferencesInner)`

SetReferences sets References field to given value.

### HasReferences

`func (o *TransportEvent) HasReferences() bool`

HasReferences returns a boolean if a field has been set.

### GetTransportEventTypeCode

`func (o *TransportEvent) GetTransportEventTypeCode() string`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *TransportEvent) GetTransportEventTypeCodeOk() (*string, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *TransportEvent) SetTransportEventTypeCode(v string)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.

### HasTransportEventTypeCode

`func (o *TransportEvent) HasTransportEventTypeCode() bool`

HasTransportEventTypeCode returns a boolean if a field has been set.

### GetDelayReasonCode

`func (o *TransportEvent) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *TransportEvent) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *TransportEvent) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *TransportEvent) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetChangeRemark

`func (o *TransportEvent) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *TransportEvent) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *TransportEvent) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *TransportEvent) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetDocumentReferences

`func (o *TransportEvent) GetDocumentReferences() []DocumentReferencesInner`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *TransportEvent) GetDocumentReferencesOk() (*[]DocumentReferencesInner, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *TransportEvent) SetDocumentReferences(v []DocumentReferencesInner)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *TransportEvent) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.

### GetTransportCall

`func (o *TransportEvent) GetTransportCall() TransportCall`

GetTransportCall returns the TransportCall field if non-nil, zero value otherwise.

### GetTransportCallOk

`func (o *TransportEvent) GetTransportCallOk() (*TransportCall, bool)`

GetTransportCallOk returns a tuple with the TransportCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportCall

`func (o *TransportEvent) SetTransportCall(v TransportCall)`

SetTransportCall sets TransportCall field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


