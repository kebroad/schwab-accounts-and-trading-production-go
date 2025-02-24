/*
Trader API - Account Access and User Preferences

Schwab Trader API access to Account, Order entry and User Preferences

API version: 1.0.0
Contact: TraderAPI@Schwab.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"fmt"
)

// ApiOrderStatus the model 'ApiOrderStatus'
type ApiOrderStatus string

// List of apiOrderStatus
const (
	APIORDERSTATUS_AWAITING_PARENT_ORDER ApiOrderStatus = "AWAITING_PARENT_ORDER"
	APIORDERSTATUS_AWAITING_CONDITION ApiOrderStatus = "AWAITING_CONDITION"
	APIORDERSTATUS_AWAITING_STOP_CONDITION ApiOrderStatus = "AWAITING_STOP_CONDITION"
	APIORDERSTATUS_AWAITING_MANUAL_REVIEW ApiOrderStatus = "AWAITING_MANUAL_REVIEW"
	APIORDERSTATUS_ACCEPTED ApiOrderStatus = "ACCEPTED"
	APIORDERSTATUS_AWAITING_UR_OUT ApiOrderStatus = "AWAITING_UR_OUT"
	APIORDERSTATUS_PENDING_ACTIVATION ApiOrderStatus = "PENDING_ACTIVATION"
	APIORDERSTATUS_QUEUED ApiOrderStatus = "QUEUED"
	APIORDERSTATUS_WORKING ApiOrderStatus = "WORKING"
	APIORDERSTATUS_REJECTED ApiOrderStatus = "REJECTED"
	APIORDERSTATUS_PENDING_CANCEL ApiOrderStatus = "PENDING_CANCEL"
	APIORDERSTATUS_CANCELED ApiOrderStatus = "CANCELED"
	APIORDERSTATUS_PENDING_REPLACE ApiOrderStatus = "PENDING_REPLACE"
	APIORDERSTATUS_REPLACED ApiOrderStatus = "REPLACED"
	APIORDERSTATUS_FILLED ApiOrderStatus = "FILLED"
	APIORDERSTATUS_EXPIRED ApiOrderStatus = "EXPIRED"
	APIORDERSTATUS_NEW ApiOrderStatus = "NEW"
	APIORDERSTATUS_AWAITING_RELEASE_TIME ApiOrderStatus = "AWAITING_RELEASE_TIME"
	APIORDERSTATUS_PENDING_ACKNOWLEDGEMENT ApiOrderStatus = "PENDING_ACKNOWLEDGEMENT"
	APIORDERSTATUS_PENDING_RECALL ApiOrderStatus = "PENDING_RECALL"
	APIORDERSTATUS_UNKNOWN ApiOrderStatus = "UNKNOWN"
)

// All allowed values of ApiOrderStatus enum
var AllowedApiOrderStatusEnumValues = []ApiOrderStatus{
	"AWAITING_PARENT_ORDER",
	"AWAITING_CONDITION",
	"AWAITING_STOP_CONDITION",
	"AWAITING_MANUAL_REVIEW",
	"ACCEPTED",
	"AWAITING_UR_OUT",
	"PENDING_ACTIVATION",
	"QUEUED",
	"WORKING",
	"REJECTED",
	"PENDING_CANCEL",
	"CANCELED",
	"PENDING_REPLACE",
	"REPLACED",
	"FILLED",
	"EXPIRED",
	"NEW",
	"AWAITING_RELEASE_TIME",
	"PENDING_ACKNOWLEDGEMENT",
	"PENDING_RECALL",
	"UNKNOWN",
}

func (v *ApiOrderStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ApiOrderStatus(value)
	for _, existing := range AllowedApiOrderStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ApiOrderStatus", value)
}

// NewApiOrderStatusFromValue returns a pointer to a valid ApiOrderStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewApiOrderStatusFromValue(v string) (*ApiOrderStatus, error) {
	ev := ApiOrderStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ApiOrderStatus: valid values are %v", v, AllowedApiOrderStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ApiOrderStatus) IsValid() bool {
	for _, existing := range AllowedApiOrderStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to apiOrderStatus value
func (v ApiOrderStatus) Ptr() *ApiOrderStatus {
	return &v
}

type NullableApiOrderStatus struct {
	value *ApiOrderStatus
	isSet bool
}

func (v NullableApiOrderStatus) Get() *ApiOrderStatus {
	return v.value
}

func (v *NullableApiOrderStatus) Set(val *ApiOrderStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableApiOrderStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableApiOrderStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiOrderStatus(val *ApiOrderStatus) *NullableApiOrderStatus {
	return &NullableApiOrderStatus{value: val, isSet: true}
}

func (v NullableApiOrderStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiOrderStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

