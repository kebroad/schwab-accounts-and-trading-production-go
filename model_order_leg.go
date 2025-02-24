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

// checks if the OrderLeg type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OrderLeg{}

// OrderLeg struct for OrderLeg
type OrderLeg struct {
	AskPrice *float64 `json:"askPrice,omitempty"`
	BidPrice *float64 `json:"bidPrice,omitempty"`
	LastPrice *float64 `json:"lastPrice,omitempty"`
	MarkPrice *float64 `json:"markPrice,omitempty"`
	ProjectedCommission *float64 `json:"projectedCommission,omitempty"`
	Quantity *float64 `json:"quantity,omitempty"`
	FinalSymbol *string `json:"finalSymbol,omitempty"`
	LegId *float32 `json:"legId,omitempty"`
	AssetType *AssetType `json:"assetType,omitempty"`
	Instruction *Instruction `json:"instruction,omitempty"`
}

// NewOrderLeg instantiates a new OrderLeg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOrderLeg() *OrderLeg {
	this := OrderLeg{}
	return &this
}

// NewOrderLegWithDefaults instantiates a new OrderLeg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOrderLegWithDefaults() *OrderLeg {
	this := OrderLeg{}
	return &this
}

// GetAskPrice returns the AskPrice field value if set, zero value otherwise.
func (o *OrderLeg) GetAskPrice() float64 {
	if o == nil || IsNil(o.AskPrice) {
		var ret float64
		return ret
	}
	return *o.AskPrice
}

// GetAskPriceOk returns a tuple with the AskPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetAskPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.AskPrice) {
		return nil, false
	}
	return o.AskPrice, true
}

// HasAskPrice returns a boolean if a field has been set.
func (o *OrderLeg) HasAskPrice() bool {
	if o != nil && !IsNil(o.AskPrice) {
		return true
	}

	return false
}

// SetAskPrice gets a reference to the given float64 and assigns it to the AskPrice field.
func (o *OrderLeg) SetAskPrice(v float64) {
	o.AskPrice = &v
}

// GetBidPrice returns the BidPrice field value if set, zero value otherwise.
func (o *OrderLeg) GetBidPrice() float64 {
	if o == nil || IsNil(o.BidPrice) {
		var ret float64
		return ret
	}
	return *o.BidPrice
}

// GetBidPriceOk returns a tuple with the BidPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetBidPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.BidPrice) {
		return nil, false
	}
	return o.BidPrice, true
}

// HasBidPrice returns a boolean if a field has been set.
func (o *OrderLeg) HasBidPrice() bool {
	if o != nil && !IsNil(o.BidPrice) {
		return true
	}

	return false
}

// SetBidPrice gets a reference to the given float64 and assigns it to the BidPrice field.
func (o *OrderLeg) SetBidPrice(v float64) {
	o.BidPrice = &v
}

// GetLastPrice returns the LastPrice field value if set, zero value otherwise.
func (o *OrderLeg) GetLastPrice() float64 {
	if o == nil || IsNil(o.LastPrice) {
		var ret float64
		return ret
	}
	return *o.LastPrice
}

// GetLastPriceOk returns a tuple with the LastPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetLastPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.LastPrice) {
		return nil, false
	}
	return o.LastPrice, true
}

// HasLastPrice returns a boolean if a field has been set.
func (o *OrderLeg) HasLastPrice() bool {
	if o != nil && !IsNil(o.LastPrice) {
		return true
	}

	return false
}

// SetLastPrice gets a reference to the given float64 and assigns it to the LastPrice field.
func (o *OrderLeg) SetLastPrice(v float64) {
	o.LastPrice = &v
}

// GetMarkPrice returns the MarkPrice field value if set, zero value otherwise.
func (o *OrderLeg) GetMarkPrice() float64 {
	if o == nil || IsNil(o.MarkPrice) {
		var ret float64
		return ret
	}
	return *o.MarkPrice
}

// GetMarkPriceOk returns a tuple with the MarkPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetMarkPriceOk() (*float64, bool) {
	if o == nil || IsNil(o.MarkPrice) {
		return nil, false
	}
	return o.MarkPrice, true
}

// HasMarkPrice returns a boolean if a field has been set.
func (o *OrderLeg) HasMarkPrice() bool {
	if o != nil && !IsNil(o.MarkPrice) {
		return true
	}

	return false
}

// SetMarkPrice gets a reference to the given float64 and assigns it to the MarkPrice field.
func (o *OrderLeg) SetMarkPrice(v float64) {
	o.MarkPrice = &v
}

// GetProjectedCommission returns the ProjectedCommission field value if set, zero value otherwise.
func (o *OrderLeg) GetProjectedCommission() float64 {
	if o == nil || IsNil(o.ProjectedCommission) {
		var ret float64
		return ret
	}
	return *o.ProjectedCommission
}

// GetProjectedCommissionOk returns a tuple with the ProjectedCommission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetProjectedCommissionOk() (*float64, bool) {
	if o == nil || IsNil(o.ProjectedCommission) {
		return nil, false
	}
	return o.ProjectedCommission, true
}

// HasProjectedCommission returns a boolean if a field has been set.
func (o *OrderLeg) HasProjectedCommission() bool {
	if o != nil && !IsNil(o.ProjectedCommission) {
		return true
	}

	return false
}

// SetProjectedCommission gets a reference to the given float64 and assigns it to the ProjectedCommission field.
func (o *OrderLeg) SetProjectedCommission(v float64) {
	o.ProjectedCommission = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *OrderLeg) GetQuantity() float64 {
	if o == nil || IsNil(o.Quantity) {
		var ret float64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetQuantityOk() (*float64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// HasQuantity returns a boolean if a field has been set.
func (o *OrderLeg) HasQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given float64 and assigns it to the Quantity field.
func (o *OrderLeg) SetQuantity(v float64) {
	o.Quantity = &v
}

// GetFinalSymbol returns the FinalSymbol field value if set, zero value otherwise.
func (o *OrderLeg) GetFinalSymbol() string {
	if o == nil || IsNil(o.FinalSymbol) {
		var ret string
		return ret
	}
	return *o.FinalSymbol
}

// GetFinalSymbolOk returns a tuple with the FinalSymbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetFinalSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.FinalSymbol) {
		return nil, false
	}
	return o.FinalSymbol, true
}

// HasFinalSymbol returns a boolean if a field has been set.
func (o *OrderLeg) HasFinalSymbol() bool {
	if o != nil && !IsNil(o.FinalSymbol) {
		return true
	}

	return false
}

// SetFinalSymbol gets a reference to the given string and assigns it to the FinalSymbol field.
func (o *OrderLeg) SetFinalSymbol(v string) {
	o.FinalSymbol = &v
}

// GetLegId returns the LegId field value if set, zero value otherwise.
func (o *OrderLeg) GetLegId() float32 {
	if o == nil || IsNil(o.LegId) {
		var ret float32
		return ret
	}
	return *o.LegId
}

// GetLegIdOk returns a tuple with the LegId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetLegIdOk() (*float32, bool) {
	if o == nil || IsNil(o.LegId) {
		return nil, false
	}
	return o.LegId, true
}

// HasLegId returns a boolean if a field has been set.
func (o *OrderLeg) HasLegId() bool {
	if o != nil && !IsNil(o.LegId) {
		return true
	}

	return false
}

// SetLegId gets a reference to the given float32 and assigns it to the LegId field.
func (o *OrderLeg) SetLegId(v float32) {
	o.LegId = &v
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *OrderLeg) GetAssetType() AssetType {
	if o == nil || IsNil(o.AssetType) {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || IsNil(o.AssetType) {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *OrderLeg) HasAssetType() bool {
	if o != nil && !IsNil(o.AssetType) {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *OrderLeg) SetAssetType(v AssetType) {
	o.AssetType = &v
}

// GetInstruction returns the Instruction field value if set, zero value otherwise.
func (o *OrderLeg) GetInstruction() Instruction {
	if o == nil || IsNil(o.Instruction) {
		var ret Instruction
		return ret
	}
	return *o.Instruction
}

// GetInstructionOk returns a tuple with the Instruction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OrderLeg) GetInstructionOk() (*Instruction, bool) {
	if o == nil || IsNil(o.Instruction) {
		return nil, false
	}
	return o.Instruction, true
}

// HasInstruction returns a boolean if a field has been set.
func (o *OrderLeg) HasInstruction() bool {
	if o != nil && !IsNil(o.Instruction) {
		return true
	}

	return false
}

// SetInstruction gets a reference to the given Instruction and assigns it to the Instruction field.
func (o *OrderLeg) SetInstruction(v Instruction) {
	o.Instruction = &v
}

func (o OrderLeg) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OrderLeg) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AskPrice) {
		toSerialize["askPrice"] = o.AskPrice
	}
	if !IsNil(o.BidPrice) {
		toSerialize["bidPrice"] = o.BidPrice
	}
	if !IsNil(o.LastPrice) {
		toSerialize["lastPrice"] = o.LastPrice
	}
	if !IsNil(o.MarkPrice) {
		toSerialize["markPrice"] = o.MarkPrice
	}
	if !IsNil(o.ProjectedCommission) {
		toSerialize["projectedCommission"] = o.ProjectedCommission
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}
	if !IsNil(o.FinalSymbol) {
		toSerialize["finalSymbol"] = o.FinalSymbol
	}
	if !IsNil(o.LegId) {
		toSerialize["legId"] = o.LegId
	}
	if !IsNil(o.AssetType) {
		toSerialize["assetType"] = o.AssetType
	}
	if !IsNil(o.Instruction) {
		toSerialize["instruction"] = o.Instruction
	}
	return toSerialize, nil
}

type NullableOrderLeg struct {
	value *OrderLeg
	isSet bool
}

func (v NullableOrderLeg) Get() *OrderLeg {
	return v.value
}

func (v *NullableOrderLeg) Set(val *OrderLeg) {
	v.value = val
	v.isSet = true
}

func (v NullableOrderLeg) IsSet() bool {
	return v.isSet
}

func (v *NullableOrderLeg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrderLeg(val *OrderLeg) *NullableOrderLeg {
	return &NullableOrderLeg{value: val, isSet: true}
}

func (v NullableOrderLeg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrderLeg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


