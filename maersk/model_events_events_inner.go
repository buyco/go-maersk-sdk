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
	"fmt"
)

// EventsEventsInner struct for EventsEventsInner
type EventsEventsInner struct {
	EquipmentEvent *EquipmentEvent
	ShipmentEvent *ShipmentEvent
	TransportEvent *TransportEvent
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventsEventsInner) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discriminator lookup.")
	}

	// check if the discriminator value is 'EquipmentEvent'
	if jsonDict["eventType"] == "EquipmentEvent" {
		// try to unmarshal JSON data into EquipmentEvent
		err = json.Unmarshal(data, &dst.EquipmentEvent);
		if err == nil {
			jsonEquipmentEvent, _ := json.Marshal(dst.EquipmentEvent)
			if string(jsonEquipmentEvent) == "{}" { // empty struct
				dst.EquipmentEvent = nil
			} else {
				return nil // data stored in dst.EquipmentEvent, return on the first match
			}
		} else {
			dst.EquipmentEvent = nil
		}
	}

	// check if the discriminator value is 'ShipmentEvent'
	if jsonDict["eventType"] == "ShipmentEvent" {
		// try to unmarshal JSON data into ShipmentEvent
		err = json.Unmarshal(data, &dst.ShipmentEvent);
		if err == nil {
			jsonShipmentEvent, _ := json.Marshal(dst.ShipmentEvent)
			if string(jsonShipmentEvent) == "{}" { // empty struct
				dst.ShipmentEvent = nil
			} else {
				return nil // data stored in dst.ShipmentEvent, return on the first match
			}
		} else {
			dst.ShipmentEvent = nil
		}
	}

	// check if the discriminator value is 'TransportEvent'
	if jsonDict["eventType"] == "TransportEvent" {
		// try to unmarshal JSON data into TransportEvent
		err = json.Unmarshal(data, &dst.TransportEvent);
		if err == nil {
			jsonTransportEvent, _ := json.Marshal(dst.TransportEvent)
			if string(jsonTransportEvent) == "{}" { // empty struct
				dst.TransportEvent = nil
			} else {
				return nil // data stored in dst.TransportEvent, return on the first match
			}
		} else {
			dst.TransportEvent = nil
		}
	}

	// try to unmarshal JSON data into EquipmentEvent
	err = json.Unmarshal(data, &dst.EquipmentEvent);
	if err == nil {
		jsonEquipmentEvent, _ := json.Marshal(dst.EquipmentEvent)
		if string(jsonEquipmentEvent) == "{}" { // empty struct
			dst.EquipmentEvent = nil
		} else {
			return nil // data stored in dst.EquipmentEvent, return on the first match
		}
	} else {
		dst.EquipmentEvent = nil
	}

	// try to unmarshal JSON data into ShipmentEvent
	err = json.Unmarshal(data, &dst.ShipmentEvent);
	if err == nil {
		jsonShipmentEvent, _ := json.Marshal(dst.ShipmentEvent)
		if string(jsonShipmentEvent) == "{}" { // empty struct
			dst.ShipmentEvent = nil
		} else {
			return nil // data stored in dst.ShipmentEvent, return on the first match
		}
	} else {
		dst.ShipmentEvent = nil
	}

	// try to unmarshal JSON data into TransportEvent
	err = json.Unmarshal(data, &dst.TransportEvent);
	if err == nil {
		jsonTransportEvent, _ := json.Marshal(dst.TransportEvent)
		if string(jsonTransportEvent) == "{}" { // empty struct
			dst.TransportEvent = nil
		} else {
			return nil // data stored in dst.TransportEvent, return on the first match
		}
	} else {
		dst.TransportEvent = nil
	}

	return fmt.Errorf("Data failed to match schemas in anyOf(EventsEventsInner)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventsEventsInner) MarshalJSON() ([]byte, error) {
	if src.EquipmentEvent != nil {
		return json.Marshal(&src.EquipmentEvent)
	}

	if src.ShipmentEvent != nil {
		return json.Marshal(&src.ShipmentEvent)
	}

	if src.TransportEvent != nil {
		return json.Marshal(&src.TransportEvent)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventsEventsInner struct {
	value *EventsEventsInner
	isSet bool
}

func (v NullableEventsEventsInner) Get() *EventsEventsInner {
	return v.value
}

func (v *NullableEventsEventsInner) Set(val *EventsEventsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableEventsEventsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableEventsEventsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventsEventsInner(val *EventsEventsInner) *NullableEventsEventsInner {
	return &NullableEventsEventsInner{value: val, isSet: true}
}

func (v NullableEventsEventsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventsEventsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


