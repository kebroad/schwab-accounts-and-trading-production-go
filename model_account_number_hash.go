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

// checks if the AccountNumberHash type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountNumberHash{}

// AccountNumberHash struct for AccountNumberHash
type AccountNumberHash struct {
	AccountNumber *string `json:"accountNumber,omitempty"`
	HashValue *string `json:"hashValue,omitempty"`
}

// NewAccountNumberHash instantiates a new AccountNumberHash object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountNumberHash() *AccountNumberHash {
	this := AccountNumberHash{}
	return &this
}

// NewAccountNumberHashWithDefaults instantiates a new AccountNumberHash object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountNumberHashWithDefaults() *AccountNumberHash {
	this := AccountNumberHash{}
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *AccountNumberHash) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNumberHash) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *AccountNumberHash) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *AccountNumberHash) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetHashValue returns the HashValue field value if set, zero value otherwise.
func (o *AccountNumberHash) GetHashValue() string {
	if o == nil || IsNil(o.HashValue) {
		var ret string
		return ret
	}
	return *o.HashValue
}

// GetHashValueOk returns a tuple with the HashValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountNumberHash) GetHashValueOk() (*string, bool) {
	if o == nil || IsNil(o.HashValue) {
		return nil, false
	}
	return o.HashValue, true
}

// HasHashValue returns a boolean if a field has been set.
func (o *AccountNumberHash) HasHashValue() bool {
	if o != nil && !IsNil(o.HashValue) {
		return true
	}

	return false
}

// SetHashValue gets a reference to the given string and assigns it to the HashValue field.
func (o *AccountNumberHash) SetHashValue(v string) {
	o.HashValue = &v
}

func (o AccountNumberHash) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountNumberHash) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.HashValue) {
		toSerialize["hashValue"] = o.HashValue
	}
	return toSerialize, nil
}

type NullableAccountNumberHash struct {
	value *AccountNumberHash
	isSet bool
}

func (v NullableAccountNumberHash) Get() *AccountNumberHash {
	return v.value
}

func (v *NullableAccountNumberHash) Set(val *AccountNumberHash) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountNumberHash) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountNumberHash) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountNumberHash(val *AccountNumberHash) *NullableAccountNumberHash {
	return &NullableAccountNumberHash{value: val, isSet: true}
}

func (v NullableAccountNumberHash) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountNumberHash) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


