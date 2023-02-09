# EventReferencesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReferenceType** | **string** | The reference type codes defined by DCSA. * FF (Freight Forwarder’s Reference) * SI (Shipper’s Reference) * PO (Purchase Order Reference) * CR (Customer’s Reference) * AAO (Consignee’s Reference) * EQ (Equipment Reference)  | 
**ReferenceValue** | **string** | The actual value of the reference. | 

## Methods

### NewEventReferencesInner

`func NewEventReferencesInner(referenceType string, referenceValue string, ) *EventReferencesInner`

NewEventReferencesInner instantiates a new EventReferencesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEventReferencesInnerWithDefaults

`func NewEventReferencesInnerWithDefaults() *EventReferencesInner`

NewEventReferencesInnerWithDefaults instantiates a new EventReferencesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReferenceType

`func (o *EventReferencesInner) GetReferenceType() string`

GetReferenceType returns the ReferenceType field if non-nil, zero value otherwise.

### GetReferenceTypeOk

`func (o *EventReferencesInner) GetReferenceTypeOk() (*string, bool)`

GetReferenceTypeOk returns a tuple with the ReferenceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceType

`func (o *EventReferencesInner) SetReferenceType(v string)`

SetReferenceType sets ReferenceType field to given value.


### GetReferenceValue

`func (o *EventReferencesInner) GetReferenceValue() string`

GetReferenceValue returns the ReferenceValue field if non-nil, zero value otherwise.

### GetReferenceValueOk

`func (o *EventReferencesInner) GetReferenceValueOk() (*string, bool)`

GetReferenceValueOk returns a tuple with the ReferenceValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceValue

`func (o *EventReferencesInner) SetReferenceValue(v string)`

SetReferenceValue sets ReferenceValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


