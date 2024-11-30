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

// checks if the CommissionValue type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommissionValue{}

// CommissionValue struct for CommissionValue
type CommissionValue struct {
	Value *float64 `json:"value,omitempty"`
	Type *FeeType `json:"type,omitempty"`
}

// NewCommissionValue instantiates a new CommissionValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommissionValue() *CommissionValue {
	this := CommissionValue{}
	return &this
}

// NewCommissionValueWithDefaults instantiates a new CommissionValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommissionValueWithDefaults() *CommissionValue {
	this := CommissionValue{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *CommissionValue) GetValue() float64 {
	if o == nil || IsNil(o.Value) {
		var ret float64
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionValue) GetValueOk() (*float64, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *CommissionValue) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given float64 and assigns it to the Value field.
func (o *CommissionValue) SetValue(v float64) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CommissionValue) GetType() FeeType {
	if o == nil || IsNil(o.Type) {
		var ret FeeType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommissionValue) GetTypeOk() (*FeeType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CommissionValue) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given FeeType and assigns it to the Type field.
func (o *CommissionValue) SetType(v FeeType) {
	o.Type = &v
}

func (o CommissionValue) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommissionValue) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	return toSerialize, nil
}

type NullableCommissionValue struct {
	value *CommissionValue
	isSet bool
}

func (v NullableCommissionValue) Get() *CommissionValue {
	return v.value
}

func (v *NullableCommissionValue) Set(val *CommissionValue) {
	v.value = val
	v.isSet = true
}

func (v NullableCommissionValue) IsSet() bool {
	return v.isSet
}

func (v *NullableCommissionValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommissionValue(val *CommissionValue) *NullableCommissionValue {
	return &NullableCommissionValue{value: val, isSet: true}
}

func (v NullableCommissionValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommissionValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


