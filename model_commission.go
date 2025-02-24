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
)

// checks if the Commission type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Commission{}

// Commission struct for Commission
type Commission struct {
	CommissionLegs []CommissionLeg `json:"commissionLegs,omitempty"`
}

// NewCommission instantiates a new Commission object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommission() *Commission {
	this := Commission{}
	return &this
}

// NewCommissionWithDefaults instantiates a new Commission object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommissionWithDefaults() *Commission {
	this := Commission{}
	return &this
}

// GetCommissionLegs returns the CommissionLegs field value if set, zero value otherwise.
func (o *Commission) GetCommissionLegs() []CommissionLeg {
	if o == nil || IsNil(o.CommissionLegs) {
		var ret []CommissionLeg
		return ret
	}
	return o.CommissionLegs
}

// GetCommissionLegsOk returns a tuple with the CommissionLegs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Commission) GetCommissionLegsOk() ([]CommissionLeg, bool) {
	if o == nil || IsNil(o.CommissionLegs) {
		return nil, false
	}
	return o.CommissionLegs, true
}

// HasCommissionLegs returns a boolean if a field has been set.
func (o *Commission) HasCommissionLegs() bool {
	if o != nil && !IsNil(o.CommissionLegs) {
		return true
	}

	return false
}

// SetCommissionLegs gets a reference to the given []CommissionLeg and assigns it to the CommissionLegs field.
func (o *Commission) SetCommissionLegs(v []CommissionLeg) {
	o.CommissionLegs = v
}

func (o Commission) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Commission) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CommissionLegs) {
		toSerialize["commissionLegs"] = o.CommissionLegs
	}
	return toSerialize, nil
}

type NullableCommission struct {
	value *Commission
	isSet bool
}

func (v NullableCommission) Get() *Commission {
	return v.value
}

func (v *NullableCommission) Set(val *Commission) {
	v.value = val
	v.isSet = true
}

func (v NullableCommission) IsSet() bool {
	return v.isSet
}

func (v *NullableCommission) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommission(val *Commission) *NullableCommission {
	return &NullableCommission{value: val, isSet: true}
}

func (v NullableCommission) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommission) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


