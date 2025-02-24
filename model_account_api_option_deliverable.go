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

// checks if the AccountAPIOptionDeliverable type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountAPIOptionDeliverable{}

// AccountAPIOptionDeliverable struct for AccountAPIOptionDeliverable
type AccountAPIOptionDeliverable struct {
	Symbol *string `json:"symbol,omitempty"`
	DeliverableUnits *float64 `json:"deliverableUnits,omitempty"`
	ApiCurrencyType *string `json:"apiCurrencyType,omitempty"`
	AssetType *AssetType `json:"assetType,omitempty"`
}

// NewAccountAPIOptionDeliverable instantiates a new AccountAPIOptionDeliverable object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountAPIOptionDeliverable() *AccountAPIOptionDeliverable {
	this := AccountAPIOptionDeliverable{}
	return &this
}

// NewAccountAPIOptionDeliverableWithDefaults instantiates a new AccountAPIOptionDeliverable object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountAPIOptionDeliverableWithDefaults() *AccountAPIOptionDeliverable {
	this := AccountAPIOptionDeliverable{}
	return &this
}

// GetSymbol returns the Symbol field value if set, zero value otherwise.
func (o *AccountAPIOptionDeliverable) GetSymbol() string {
	if o == nil || IsNil(o.Symbol) {
		var ret string
		return ret
	}
	return *o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAPIOptionDeliverable) GetSymbolOk() (*string, bool) {
	if o == nil || IsNil(o.Symbol) {
		return nil, false
	}
	return o.Symbol, true
}

// HasSymbol returns a boolean if a field has been set.
func (o *AccountAPIOptionDeliverable) HasSymbol() bool {
	if o != nil && !IsNil(o.Symbol) {
		return true
	}

	return false
}

// SetSymbol gets a reference to the given string and assigns it to the Symbol field.
func (o *AccountAPIOptionDeliverable) SetSymbol(v string) {
	o.Symbol = &v
}

// GetDeliverableUnits returns the DeliverableUnits field value if set, zero value otherwise.
func (o *AccountAPIOptionDeliverable) GetDeliverableUnits() float64 {
	if o == nil || IsNil(o.DeliverableUnits) {
		var ret float64
		return ret
	}
	return *o.DeliverableUnits
}

// GetDeliverableUnitsOk returns a tuple with the DeliverableUnits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAPIOptionDeliverable) GetDeliverableUnitsOk() (*float64, bool) {
	if o == nil || IsNil(o.DeliverableUnits) {
		return nil, false
	}
	return o.DeliverableUnits, true
}

// HasDeliverableUnits returns a boolean if a field has been set.
func (o *AccountAPIOptionDeliverable) HasDeliverableUnits() bool {
	if o != nil && !IsNil(o.DeliverableUnits) {
		return true
	}

	return false
}

// SetDeliverableUnits gets a reference to the given float64 and assigns it to the DeliverableUnits field.
func (o *AccountAPIOptionDeliverable) SetDeliverableUnits(v float64) {
	o.DeliverableUnits = &v
}

// GetApiCurrencyType returns the ApiCurrencyType field value if set, zero value otherwise.
func (o *AccountAPIOptionDeliverable) GetApiCurrencyType() string {
	if o == nil || IsNil(o.ApiCurrencyType) {
		var ret string
		return ret
	}
	return *o.ApiCurrencyType
}

// GetApiCurrencyTypeOk returns a tuple with the ApiCurrencyType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAPIOptionDeliverable) GetApiCurrencyTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ApiCurrencyType) {
		return nil, false
	}
	return o.ApiCurrencyType, true
}

// HasApiCurrencyType returns a boolean if a field has been set.
func (o *AccountAPIOptionDeliverable) HasApiCurrencyType() bool {
	if o != nil && !IsNil(o.ApiCurrencyType) {
		return true
	}

	return false
}

// SetApiCurrencyType gets a reference to the given string and assigns it to the ApiCurrencyType field.
func (o *AccountAPIOptionDeliverable) SetApiCurrencyType(v string) {
	o.ApiCurrencyType = &v
}

// GetAssetType returns the AssetType field value if set, zero value otherwise.
func (o *AccountAPIOptionDeliverable) GetAssetType() AssetType {
	if o == nil || IsNil(o.AssetType) {
		var ret AssetType
		return ret
	}
	return *o.AssetType
}

// GetAssetTypeOk returns a tuple with the AssetType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountAPIOptionDeliverable) GetAssetTypeOk() (*AssetType, bool) {
	if o == nil || IsNil(o.AssetType) {
		return nil, false
	}
	return o.AssetType, true
}

// HasAssetType returns a boolean if a field has been set.
func (o *AccountAPIOptionDeliverable) HasAssetType() bool {
	if o != nil && !IsNil(o.AssetType) {
		return true
	}

	return false
}

// SetAssetType gets a reference to the given AssetType and assigns it to the AssetType field.
func (o *AccountAPIOptionDeliverable) SetAssetType(v AssetType) {
	o.AssetType = &v
}

func (o AccountAPIOptionDeliverable) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountAPIOptionDeliverable) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Symbol) {
		toSerialize["symbol"] = o.Symbol
	}
	if !IsNil(o.DeliverableUnits) {
		toSerialize["deliverableUnits"] = o.DeliverableUnits
	}
	if !IsNil(o.ApiCurrencyType) {
		toSerialize["apiCurrencyType"] = o.ApiCurrencyType
	}
	if !IsNil(o.AssetType) {
		toSerialize["assetType"] = o.AssetType
	}
	return toSerialize, nil
}

type NullableAccountAPIOptionDeliverable struct {
	value *AccountAPIOptionDeliverable
	isSet bool
}

func (v NullableAccountAPIOptionDeliverable) Get() *AccountAPIOptionDeliverable {
	return v.value
}

func (v *NullableAccountAPIOptionDeliverable) Set(val *AccountAPIOptionDeliverable) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountAPIOptionDeliverable) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountAPIOptionDeliverable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountAPIOptionDeliverable(val *AccountAPIOptionDeliverable) *NullableAccountAPIOptionDeliverable {
	return &NullableAccountAPIOptionDeliverable{value: val, isSet: true}
}

func (v NullableAccountAPIOptionDeliverable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountAPIOptionDeliverable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


