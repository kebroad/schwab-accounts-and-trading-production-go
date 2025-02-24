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

// StopPriceLinkBasis the model 'StopPriceLinkBasis'
type StopPriceLinkBasis string

// List of stopPriceLinkBasis
const (
	STOPPRICELINKBASIS_MANUAL StopPriceLinkBasis = "MANUAL"
	STOPPRICELINKBASIS_BASE StopPriceLinkBasis = "BASE"
	STOPPRICELINKBASIS_TRIGGER StopPriceLinkBasis = "TRIGGER"
	STOPPRICELINKBASIS_LAST StopPriceLinkBasis = "LAST"
	STOPPRICELINKBASIS_BID StopPriceLinkBasis = "BID"
	STOPPRICELINKBASIS_ASK StopPriceLinkBasis = "ASK"
	STOPPRICELINKBASIS_ASK_BID StopPriceLinkBasis = "ASK_BID"
	STOPPRICELINKBASIS_MARK StopPriceLinkBasis = "MARK"
	STOPPRICELINKBASIS_AVERAGE StopPriceLinkBasis = "AVERAGE"
)

// All allowed values of StopPriceLinkBasis enum
var AllowedStopPriceLinkBasisEnumValues = []StopPriceLinkBasis{
	"MANUAL",
	"BASE",
	"TRIGGER",
	"LAST",
	"BID",
	"ASK",
	"ASK_BID",
	"MARK",
	"AVERAGE",
}

func (v *StopPriceLinkBasis) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StopPriceLinkBasis(value)
	for _, existing := range AllowedStopPriceLinkBasisEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StopPriceLinkBasis", value)
}

// NewStopPriceLinkBasisFromValue returns a pointer to a valid StopPriceLinkBasis
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStopPriceLinkBasisFromValue(v string) (*StopPriceLinkBasis, error) {
	ev := StopPriceLinkBasis(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StopPriceLinkBasis: valid values are %v", v, AllowedStopPriceLinkBasisEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StopPriceLinkBasis) IsValid() bool {
	for _, existing := range AllowedStopPriceLinkBasisEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to stopPriceLinkBasis value
func (v StopPriceLinkBasis) Ptr() *StopPriceLinkBasis {
	return &v
}

type NullableStopPriceLinkBasis struct {
	value *StopPriceLinkBasis
	isSet bool
}

func (v NullableStopPriceLinkBasis) Get() *StopPriceLinkBasis {
	return v.value
}

func (v *NullableStopPriceLinkBasis) Set(val *StopPriceLinkBasis) {
	v.value = val
	v.isSet = true
}

func (v NullableStopPriceLinkBasis) IsSet() bool {
	return v.isSet
}

func (v *NullableStopPriceLinkBasis) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStopPriceLinkBasis(val *StopPriceLinkBasis) *NullableStopPriceLinkBasis {
	return &NullableStopPriceLinkBasis{value: val, isSet: true}
}

func (v NullableStopPriceLinkBasis) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStopPriceLinkBasis) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

