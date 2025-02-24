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

// checks if the OrderActivity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderActivity{}

// OrderActivity struct for OrderActivity
type OrderActivity struct {
	ActivityType *string `json:"activityType,omitempty"`
	ExecutionType *string `json:"executionType,omitempty"`
	Quantity *float64 `json:"quantity,omitempty"`
	OrderRemainingQuantity *float64 `json:"orderRemainingQuantity,omitempty"`
	ExecutionLegs []ExecutionLeg `json:"executionLegs,omitempty"`
}

// NewOrderActivity instantiates a new OrderActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderActivity() *OrderActivity {
	this := OrderActivity{}
	return &this
}

// NewOrderActivityWithDefaults instantiates a new OrderActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderActivityWithDefaults() *OrderActivity {
	this := OrderActivity{}
	return &this
}

// GetActivityType returns the ActivityType field value if set, zero value otherwise.
func (o *OrderActivity) GetActivityType() string {
	if o == nil || IsNil(o.ActivityType) {
		var ret string
		return ret
	}
	return *o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderActivity) GetActivityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityType) {
		return nil, false
	}
	return o.ActivityType, true
}

// HasActivityType returns a boolean if a field has been set.
func (o *OrderActivity) HasActivityType() bool {
	if o != nil && !IsNil(o.ActivityType) {
		return true
	}

	return false
}

// SetActivityType gets a reference to the given string and assigns it to the ActivityType field.
func (o *OrderActivity) SetActivityType(v string) {
	o.ActivityType = &v
}

// GetExecutionType returns the ExecutionType field value if set, zero value otherwise.
func (o *OrderActivity) GetExecutionType() string {
	if o == nil || IsNil(o.ExecutionType) {
		var ret string
		return ret
	}
	return *o.ExecutionType
}

// GetExecutionTypeOk returns a tuple with the ExecutionType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderActivity) GetExecutionTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ExecutionType) {
		return nil, false
	}
	return o.ExecutionType, true
}

// HasExecutionType returns a boolean if a field has been set.
func (o *OrderActivity) HasExecutionType() bool {
	if o != nil && !IsNil(o.ExecutionType) {
		return true
	}

	return false
}

// SetExecutionType gets a reference to the given string and assigns it to the ExecutionType field.
func (o *OrderActivity) SetExecutionType(v string) {
	o.ExecutionType = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderActivity) GetQuantity() float64 {
	if o == nil || IsNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderActivity) GetQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderActivity) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *OrderActivity) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetOrderRemainingQuantity returns the OrderRemainingQuantity field value if set, zero value otherwise.
func (o *OrderActivity) GetOrderRemainingQuantity() float64 {
	if o == nil || IsNil(o.OrderRemainingQuantity) {
		var ret float64
		return ret
	}
	return *o.OrderRemainingQuantity
}

// GetOrderRemainingQuantityOk returns a tuple with the OrderRemainingQuantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderActivity) GetOrderRemainingQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.OrderRemainingQuantity) {
		return nil, false
	}
	return o.OrderRemainingQuantity, true
}

// HasOrderRemainingQuantity returns a boolean if a field has been set.
func (o *OrderActivity) HasOrderRemainingQuantity() bool {
	if o != nil && !IsNil(o.OrderRemainingQuantity) {
		return true
	}

	return false
}

// SetOrderRemainingQuantity gets a reference to the given float64 and assigns it to the OrderRemainingQuantity field.
func (o *OrderActivity) SetOrderRemainingQuantity(v float64) {
	o.OrderRemainingQuantity = &v
}

// GetExecutionLegs returns the ExecutionLegs field value if set, zero value otherwise.
func (o *OrderActivity) GetExecutionLegs() []ExecutionLeg {
	if o == nil || IsNil(o.ExecutionLegs) {
		var ret []ExecutionLeg
		return ret
	}
	return o.ExecutionLegs
}

// GetExecutionLegsOk returns a tuple with the ExecutionLegs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderActivity) GetExecutionLegsOk() ([]ExecutionLeg, bool) {
	if o == nil || IsNil(o.ExecutionLegs) {
		return nil, false
	}
	return o.ExecutionLegs, true
}

// HasExecutionLegs returns a boolean if a field has been set.
func (o *OrderActivity) HasExecutionLegs() bool {
	if o != nil && !IsNil(o.ExecutionLegs) {
		return true
	}

	return false
}

// SetExecutionLegs gets a reference to the given []ExecutionLeg and assigns it to the ExecutionLegs field.
func (o *OrderActivity) SetExecutionLegs(v []ExecutionLeg) {
	o.ExecutionLegs = v
}

func (o OrderActivity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderActivity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ActivityType) {
		toSerialize["activityType"] = o.ActivityType
	}
	if !IsNil(o.ExecutionType) {
		toSerialize["executionType"] = o.ExecutionType
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.OrderRemainingQuantity) {
		toSerialize["orderRemainingQuantity"] = o.OrderRemainingQuantity
	}
	if !IsNil(o.ExecutionLegs) {
		toSerialize["executionLegs"] = o.ExecutionLegs
	}
	return toSerialize, nil
}

type NullableOrderActivity struct {
	value *OrderActivity
	isSet bool
}

func (v NullableOrderActivity) Get() *OrderActivity {
	return v.value
}

func (v *NullableOrderActivity) Set(val *OrderActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderActivity(val *OrderActivity) *NullableOrderActivity {
	return &NullableOrderActivity{value: val, isSet: true}
}

func (v NullableOrderActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


