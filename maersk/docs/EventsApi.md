# \EventsApi

All URIs are relative to *https://api.maersk.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListEvents**](EventsApi.md#ListEvents) | **Get** /track-and-trace-private/events | Find events.



## ListEvents

> Events ListEvents(ctx).ConsumerKey(consumerKey).Authorization(authorization).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).EquipmentReference(equipmentReference).EventType(eventType).EventCreatedDateTime(eventCreatedDateTime).ShipmentEventTypeCode(shipmentEventTypeCode).TransportEventTypeCode(transportEventTypeCode).EquipmentEventTypeCode(equipmentEventTypeCode).Limit(limit).Cursor(cursor).APIVersion(aPIVersion).Execute()

Find events.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    consumerKey := "consumerKey_example" // string | The Consumer Key issued for your registered application.
    authorization := "authorization_example" // string | Bearer JWT
    carrierBookingReference := "VAS000001" // string | A set of unique characters provided by carrier to identify a booking. Specifying this filter will only return events related to this particular carrierBookingReference.  (optional)
    transportDocumentReference := "260029935" // string | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment. Specifying this filter will only return events related to this particular transportDocumentReference  (optional)
    equipmentReference := "APZU4812090" // string | Will filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. Specifying this filter will only return events related to this particular equipmentReference  (optional)
    eventType := []string{"EventType_example"} // []string | The type of event(s) to filter by. Possible values are - SHIPMENT (Shipment events) - TRANSPORT (Transport events) - EQUIPMENT (Equipment events)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR operator is used. For example, [eventType=SHIPMENT,EQUIPMENT] matches both Shipment and Equipment events.\\ Default value is all event types.  (optional) (default to ["SHIPMENT","TRANSPORT","EQUIPMENT"])
    eventCreatedDateTime := "gte:2021-04-01T00:00:00Z" // string | Limit the result based on a UTC date. It is possible to use operators on this query parameter. This is done by adding an operator at the beginning of the value followed by a colon:\\ eventCreatedDateTime = **gte**:2021-04-01T00:00:00Z\\ would result in all events created >= 2021-04-01T00:00:00Z\\ The following operators are supported - gte: (>= Greater than or equal) - gt: (> Greater than) - lte: (<= Less than or equal) - lt: (< Less than) - eq: (= Equal to)  If no operator is provided, a **strictly equal** is used (this is equivalent to **eq:** operator).  (optional)
    shipmentEventTypeCode := []string{"ShipmentEventTypeCode_example"} // []string | The status of the document in the process to filter by. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example, [shipmentEventTypeCode=RECE,DRFT] matches **both** Received (RECE) and Drafted (DRFT) shipment events.\\ Default is all shipmentEventTypeCodes.\\ This filter is only relevant when filtering on ShipmentEvents  **Note: Version 1.1 replaces CONF (Confirmed) for RELS (Released) for documentTypeCode SRM (Shipment Release Message).**  (optional) (default to ["RECE","DRFT","PENA","PENU","REJE","APPR","ISSU","SURR","SUBM","VOID","CONF","REQS","CMPL","HOLD","RELS"])
    transportEventTypeCode := []string{"TransportEventTypeCode_example"} // []string | Identifier for type of Transport event to filter by - ARRI (Arrived) - DEPA (Departed)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR operator is used. For example, [transportEventTypeCode=ARRI,DEPA} matches **both** Arrived (ARRI) and Departed (DEPA) transport events.\\ Default is all transportEventTypeCodes.\\ This filter is only relevant when filtering on TransportEvents  (optional) (default to ["ARRI","DEPA"])
    equipmentEventTypeCode := []string{"EquipmentEventTypeCode_example"} // []string | Unique identifier for equipmentEventTypeCode. * LOAD (Loaded) * DISC (Discharged) * GTIN (Gated in) * GTOT (Gated out) * STUF (Stuffed) * STRP (Stripped) * PICK (Pick-up) * DROP (Drop-off) * RSEA (Resealed) * RMVD (Removed) * INSP (Inspected)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR operator is used. For example, [equipmentEventTypeCode=GTIN,GTOT] matches **both** Gated in (GTIN) and Gated out (GTOT) equipment events.\\ Default is all equipmentEventTypeCodes.\\ This filter is only relevant when filtering on EquipmentEvents  (optional) (default to ["LOAD","DISC","GTIN","GTOT","STUF","STRP","PICK","DROP","RSEA","RMVD","INSP"])
    limit := int32(100) // int32 | Maximum number of items to return. (optional) (default to 100)
    cursor := "1" // string | A server generated value to specify a specific point in a collection result, used for pagination.  The current, previous, next, first and last pages are available in the response headers.  For the initial request to the service, this parameter should be null or 1. (optional)
    aPIVersion := "1" // string | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EventsApi.ListEvents(context.Background()).ConsumerKey(consumerKey).Authorization(authorization).CarrierBookingReference(carrierBookingReference).TransportDocumentReference(transportDocumentReference).EquipmentReference(equipmentReference).EventType(eventType).EventCreatedDateTime(eventCreatedDateTime).ShipmentEventTypeCode(shipmentEventTypeCode).TransportEventTypeCode(transportEventTypeCode).EquipmentEventTypeCode(equipmentEventTypeCode).Limit(limit).Cursor(cursor).APIVersion(aPIVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventsApi.ListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvents`: Events
    fmt.Fprintf(os.Stdout, "Response from `EventsApi.ListEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **consumerKey** | **string** | The Consumer Key issued for your registered application. | 
 **authorization** | **string** | Bearer JWT | 
 **carrierBookingReference** | **string** | A set of unique characters provided by carrier to identify a booking. Specifying this filter will only return events related to this particular carrierBookingReference.  | 
 **transportDocumentReference** | **string** | A unique number reference allocated by the shipping line to the transport document and the main number used for the tracking of the status of the shipment. Specifying this filter will only return events related to this particular transportDocumentReference  | 
 **equipmentReference** | **string** | Will filter by the unique identifier for the equipment, which should follow the BIC ISO Container Identification Number where possible. Specifying this filter will only return events related to this particular equipmentReference  | 
 **eventType** | **[]string** | The type of event(s) to filter by. Possible values are - SHIPMENT (Shipment events) - TRANSPORT (Transport events) - EQUIPMENT (Equipment events)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR operator is used. For example, [eventType&#x3D;SHIPMENT,EQUIPMENT] matches both Shipment and Equipment events.\\ Default value is all event types.  | [default to [&quot;SHIPMENT&quot;,&quot;TRANSPORT&quot;,&quot;EQUIPMENT&quot;]]
 **eventCreatedDateTime** | **string** | Limit the result based on a UTC date. It is possible to use operators on this query parameter. This is done by adding an operator at the beginning of the value followed by a colon:\\ eventCreatedDateTime &#x3D; **gte**:2021-04-01T00:00:00Z\\ would result in all events created &gt;&#x3D; 2021-04-01T00:00:00Z\\ The following operators are supported - gte: (&gt;&#x3D; Greater than or equal) - gt: (&gt; Greater than) - lte: (&lt;&#x3D; Less than or equal) - lt: (&lt; Less than) - eq: (&#x3D; Equal to)  If no operator is provided, a **strictly equal** is used (this is equivalent to **eq:** operator).  | 
 **shipmentEventTypeCode** | **[]string** | The status of the document in the process to filter by. Possible values are - RECE (Received) - DRFT (Drafted) - PENA (Pending Approval) - PENU (Pending Update) - REJE (Rejected) - APPR (Approved) - ISSU (Issued) - SURR (Surrendered) - SUBM (Submitted) - VOID (Void) - CONF (Confirmed) - REQS (Requested) - CMPL (Completed) - HOLD (On Hold) - RELS (Released)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR-operator is used. For example, [shipmentEventTypeCode&#x3D;RECE,DRFT] matches **both** Received (RECE) and Drafted (DRFT) shipment events.\\ Default is all shipmentEventTypeCodes.\\ This filter is only relevant when filtering on ShipmentEvents  **Note: Version 1.1 replaces CONF (Confirmed) for RELS (Released) for documentTypeCode SRM (Shipment Release Message).**  | [default to [&quot;RECE&quot;,&quot;DRFT&quot;,&quot;PENA&quot;,&quot;PENU&quot;,&quot;REJE&quot;,&quot;APPR&quot;,&quot;ISSU&quot;,&quot;SURR&quot;,&quot;SUBM&quot;,&quot;VOID&quot;,&quot;CONF&quot;,&quot;REQS&quot;,&quot;CMPL&quot;,&quot;HOLD&quot;,&quot;RELS&quot;]]
 **transportEventTypeCode** | **[]string** | Identifier for type of Transport event to filter by - ARRI (Arrived) - DEPA (Departed)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR operator is used. For example, [transportEventTypeCode&#x3D;ARRI,DEPA} matches **both** Arrived (ARRI) and Departed (DEPA) transport events.\\ Default is all transportEventTypeCodes.\\ This filter is only relevant when filtering on TransportEvents  | [default to [&quot;ARRI&quot;,&quot;DEPA&quot;]]
 **equipmentEventTypeCode** | **[]string** | Unique identifier for equipmentEventTypeCode. * LOAD (Loaded) * DISC (Discharged) * GTIN (Gated in) * GTOT (Gated out) * STUF (Stuffed) * STRP (Stripped) * PICK (Pick-up) * DROP (Drop-off) * RSEA (Resealed) * RMVD (Removed) * INSP (Inspected)  It is possible to select multiple values by comma (,) separating them. For multiple values the OR operator is used. For example, [equipmentEventTypeCode&#x3D;GTIN,GTOT] matches **both** Gated in (GTIN) and Gated out (GTOT) equipment events.\\ Default is all equipmentEventTypeCodes.\\ This filter is only relevant when filtering on EquipmentEvents  | [default to [&quot;LOAD&quot;,&quot;DISC&quot;,&quot;GTIN&quot;,&quot;GTOT&quot;,&quot;STUF&quot;,&quot;STRP&quot;,&quot;PICK&quot;,&quot;DROP&quot;,&quot;RSEA&quot;,&quot;RMVD&quot;,&quot;INSP&quot;]]
 **limit** | **int32** | Maximum number of items to return. | [default to 100]
 **cursor** | **string** | A server generated value to specify a specific point in a collection result, used for pagination.  The current, previous, next, first and last pages are available in the response headers.  For the initial request to the service, this parameter should be null or 1. | 
 **aPIVersion** | **string** | An API-Version header MAY be added to the request (optional); if added it MUST only contain MAJOR version. API-Version header MUST be aligned with the URI version. | 

### Return type

[**Events**](Events.md)

### Authorization

[ConsumerKey](../README.md#ConsumerKey), [oAuth2](../README.md#oAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, application/stream+json, text/event-stream

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

