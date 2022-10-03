# ShipmentEventAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventType** | Pointer to **string** |  | [optional] 
**ShipmentEventTypeCode** | **string** | The status of the document in the process. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released)  Note: Version 1.1 replaces CONF (Confirmed) for RELS (Released) for documentTypeCode SRM (Shipment Release Message).  | 
**DocumentTypeCode** | **string** | The code to identify the type of information documentID points to. Can be one of the following values * CBR (Carrier Booking Request Reference) * BKG (Booking) * SHI (Shipping Instruction) * SRM (Shipment Release Message) * TRD (Transport Document) * ARN (Arrival Notice) * VGM (Verified Gross Mass) * CAS (Cargo Survey) * CUS (Customs Inspection) * DGD (Dangerous Goods Declaration) * OOG (Out of Gauge)  | 
**DocumentID** | **string** | The ID of the object defined by the Shipment Information Type. In some cases this is a UUID; in other cases this is a string.  | 

## Methods

### NewShipmentEventAllOf

`func NewShipmentEventAllOf(shipmentEventTypeCode string, documentTypeCode string, documentID string, ) *ShipmentEventAllOf`

NewShipmentEventAllOf instantiates a new ShipmentEventAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentEventAllOfWithDefaults

`func NewShipmentEventAllOfWithDefaults() *ShipmentEventAllOf`

NewShipmentEventAllOfWithDefaults instantiates a new ShipmentEventAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventType

`func (o *ShipmentEventAllOf) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ShipmentEventAllOf) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ShipmentEventAllOf) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ShipmentEventAllOf) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetShipmentEventTypeCode

`func (o *ShipmentEventAllOf) GetShipmentEventTypeCode() string`

GetShipmentEventTypeCode returns the ShipmentEventTypeCode field if non-nil, zero value otherwise.

### GetShipmentEventTypeCodeOk

`func (o *ShipmentEventAllOf) GetShipmentEventTypeCodeOk() (*string, bool)`

GetShipmentEventTypeCodeOk returns a tuple with the ShipmentEventTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentEventTypeCode

`func (o *ShipmentEventAllOf) SetShipmentEventTypeCode(v string)`

SetShipmentEventTypeCode sets ShipmentEventTypeCode field to given value.


### GetDocumentTypeCode

`func (o *ShipmentEventAllOf) GetDocumentTypeCode() string`

GetDocumentTypeCode returns the DocumentTypeCode field if non-nil, zero value otherwise.

### GetDocumentTypeCodeOk

`func (o *ShipmentEventAllOf) GetDocumentTypeCodeOk() (*string, bool)`

GetDocumentTypeCodeOk returns a tuple with the DocumentTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentTypeCode

`func (o *ShipmentEventAllOf) SetDocumentTypeCode(v string)`

SetDocumentTypeCode sets DocumentTypeCode field to given value.


### GetDocumentID

`func (o *ShipmentEventAllOf) GetDocumentID() string`

GetDocumentID returns the DocumentID field if non-nil, zero value otherwise.

### GetDocumentIDOk

`func (o *ShipmentEventAllOf) GetDocumentIDOk() (*string, bool)`

GetDocumentIDOk returns a tuple with the DocumentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentID

`func (o *ShipmentEventAllOf) SetDocumentID(v string)`

SetDocumentID sets DocumentID field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


