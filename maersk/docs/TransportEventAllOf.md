# TransportEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**TransportEventTypeCode** | Pointer to **string** | Identifier for type of Transport event - ARRI (Arrived) - DEPA (Departed)  | [optional] 
**DelayReasonCode** | Pointer to **string** | Reason code for the delay. The SMDG-Delay-Reason-Codes are used for this attribute. The code list can be found at http://www.smdg.org/smdg-code-lists/  | [optional] 
**ChangeRemark** | Pointer to **string** | Free text information provided by the vessel operator regarding the reasons for the change in schedule and/or plans to mitigate schedule slippage. | [optional] 
**DocumentReferences** | Pointer to [**[]DocumentReferencesInner**](DocumentReferencesInner.md) | An optional list of key-value (documentReferenceType-documentReferenceValue) pairs representing links to objects relevant to the event. The documentReferenceType-field is used to describe where the documentReferenceValue-field is pointing to. | [optional] 

## Methods

### NewTransportEventAllOf

`func NewTransportEventAllOf() *TransportEventAllOf`

NewTransportEventAllOf instantiates a new TransportEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransportEventAllOfWithDefaults

`func NewTransportEventAllOfWithDefaults() *TransportEventAllOf`

NewTransportEventAllOfWithDefaults instantiates a new TransportEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *TransportEventAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *TransportEventAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *TransportEventAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *TransportEventAllOf) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetTransportEventTypeCode

`func (o *TransportEventAllOf) GetTransportEventTypeCode() string`

GetTransportEventTypeCode returns the TransportEventTypeCode field if non-nil, zero value otherwise.

### GetTransportEventTypeCodeOk

`func (o *TransportEventAllOf) GetTransportEventTypeCodeOk() (*string, bool)`

GetTransportEventTypeCodeOk returns a tuple with the TransportEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransportEventTypeCode

`func (o *TransportEventAllOf) SetTransportEventTypeCode(v string)`

SetTransportEventTypeCode sets TransportEventTypeCode field to given value.

### HasTransportEventTypeCode

`func (o *TransportEventAllOf) HasTransportEventTypeCode() bool`

HasTransportEventTypeCode returns a boolean if a field has been set.

### GetDelayReasonCode

`func (o *TransportEventAllOf) GetDelayReasonCode() string`

GetDelayReasonCode returns the DelayReasonCode field if non-nil, zero value otherwise.

### GetDelayReasonCodeOk

`func (o *TransportEventAllOf) GetDelayReasonCodeOk() (*string, bool)`

GetDelayReasonCodeOk returns a tuple with the DelayReasonCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayReasonCode

`func (o *TransportEventAllOf) SetDelayReasonCode(v string)`

SetDelayReasonCode sets DelayReasonCode field to given value.

### HasDelayReasonCode

`func (o *TransportEventAllOf) HasDelayReasonCode() bool`

HasDelayReasonCode returns a boolean if a field has been set.

### GetChangeRemark

`func (o *TransportEventAllOf) GetChangeRemark() string`

GetChangeRemark returns the ChangeRemark field if non-nil, zero value otherwise.

### GetChangeRemarkOk

`func (o *TransportEventAllOf) GetChangeRemarkOk() (*string, bool)`

GetChangeRemarkOk returns a tuple with the ChangeRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChangeRemark

`func (o *TransportEventAllOf) SetChangeRemark(v string)`

SetChangeRemark sets ChangeRemark field to given value.

### HasChangeRemark

`func (o *TransportEventAllOf) HasChangeRemark() bool`

HasChangeRemark returns a boolean if a field has been set.

### GetDocumentReferences

`func (o *TransportEventAllOf) GetDocumentReferences() []DocumentReferencesInner`

GetDocumentReferences returns the DocumentReferences field if non-nil, zero value otherwise.

### GetDocumentReferencesOk

`func (o *TransportEventAllOf) GetDocumentReferencesOk() (*[]DocumentReferencesInner, bool)`

GetDocumentReferencesOk returns a tuple with the DocumentReferences field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentReferences

`func (o *TransportEventAllOf) SetDocumentReferences(v []DocumentReferencesInner)`

SetDocumentReferences sets DocumentReferences field to given value.

### HasDocumentReferences

`func (o *TransportEventAllOf) HasDocumentReferences() bool`

HasDocumentReferences returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


