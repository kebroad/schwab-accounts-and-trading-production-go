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

// checks if the UserPreferenceAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPreferenceAccount{}

// UserPreferenceAccount struct for UserPreferenceAccount
type UserPreferenceAccount struct {
	AccountNumber *string `json:"accountNumber,omitempty"`
	PrimaryAccount *bool `json:"primaryAccount,omitempty"`
	Type *string `json:"type,omitempty"`
	NickName *string `json:"nickName,omitempty"`
	// Green | Blue
	AccountColor *string `json:"accountColor,omitempty"`
	DisplayAcctId *string `json:"displayAcctId,omitempty"`
	AutoPositionEffect *bool `json:"autoPositionEffect,omitempty"`
}

// NewUserPreferenceAccount instantiates a new UserPreferenceAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPreferenceAccount() *UserPreferenceAccount {
	this := UserPreferenceAccount{}
	var primaryAccount bool = false
	this.PrimaryAccount = &primaryAccount
	var autoPositionEffect bool = false
	this.AutoPositionEffect = &autoPositionEffect
	return &this
}

// NewUserPreferenceAccountWithDefaults instantiates a new UserPreferenceAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPreferenceAccountWithDefaults() *UserPreferenceAccount {
	this := UserPreferenceAccount{}
	var primaryAccount bool = false
	this.PrimaryAccount = &primaryAccount
	var autoPositionEffect bool = false
	this.AutoPositionEffect = &autoPositionEffect
	return &this
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *UserPreferenceAccount) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAccount) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// HasAccountNumber returns a boolean if a field has been set.
func (o *UserPreferenceAccount) HasAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *UserPreferenceAccount) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetPrimaryAccount returns the PrimaryAccount field value if set, zero value otherwise.
func (o *UserPreferenceAccount) GetPrimaryAccount() bool {
	if o == nil || IsNil(o.PrimaryAccount) {
		var ret bool
		return ret
	}
	return *o.PrimaryAccount
}

// GetPrimaryAccountOk returns a tuple with the PrimaryAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAccount) GetPrimaryAccountOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryAccount) {
		return nil, false
	}
	return o.PrimaryAccount, true
}

// HasPrimaryAccount returns a boolean if a field has been set.
func (o *UserPreferenceAccount) HasPrimaryAccount() bool {
	if o != nil && !IsNil(o.PrimaryAccount) {
		return true
	}

	return false
}

// SetPrimaryAccount gets a reference to the given bool and assigns it to the PrimaryAccount field.
func (o *UserPreferenceAccount) SetPrimaryAccount(v bool) {
	o.PrimaryAccount = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserPreferenceAccount) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAccount) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserPreferenceAccount) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserPreferenceAccount) SetType(v string) {
	o.Type = &v
}

// GetNickName returns the NickName field value if set, zero value otherwise.
func (o *UserPreferenceAccount) GetNickName() string {
	if o == nil || IsNil(o.NickName) {
		var ret string
		return ret
	}
	return *o.NickName
}

// GetNickNameOk returns a tuple with the NickName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAccount) GetNickNameOk() (*string, bool) {
	if o == nil || IsNil(o.NickName) {
		return nil, false
	}
	return o.NickName, true
}

// HasNickName returns a boolean if a field has been set.
func (o *UserPreferenceAccount) HasNickName() bool {
	if o != nil && !IsNil(o.NickName) {
		return true
	}

	return false
}

// SetNickName gets a reference to the given string and assigns it to the NickName field.
func (o *UserPreferenceAccount) SetNickName(v string) {
	o.NickName = &v
}

// GetAccountColor returns the AccountColor field value if set, zero value otherwise.
func (o *UserPreferenceAccount) GetAccountColor() string {
	if o == nil || IsNil(o.AccountColor) {
		var ret string
		return ret
	}
	return *o.AccountColor
}

// GetAccountColorOk returns a tuple with the AccountColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAccount) GetAccountColorOk() (*string, bool) {
	if o == nil || IsNil(o.AccountColor) {
		return nil, false
	}
	return o.AccountColor, true
}

// HasAccountColor returns a boolean if a field has been set.
func (o *UserPreferenceAccount) HasAccountColor() bool {
	if o != nil && !IsNil(o.AccountColor) {
		return true
	}

	return false
}

// SetAccountColor gets a reference to the given string and assigns it to the AccountColor field.
func (o *UserPreferenceAccount) SetAccountColor(v string) {
	o.AccountColor = &v
}

// GetDisplayAcctId returns the DisplayAcctId field value if set, zero value otherwise.
func (o *UserPreferenceAccount) GetDisplayAcctId() string {
	if o == nil || IsNil(o.DisplayAcctId) {
		var ret string
		return ret
	}
	return *o.DisplayAcctId
}

// GetDisplayAcctIdOk returns a tuple with the DisplayAcctId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAccount) GetDisplayAcctIdOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayAcctId) {
		return nil, false
	}
	return o.DisplayAcctId, true
}

// HasDisplayAcctId returns a boolean if a field has been set.
func (o *UserPreferenceAccount) HasDisplayAcctId() bool {
	if o != nil && !IsNil(o.DisplayAcctId) {
		return true
	}

	return false
}

// SetDisplayAcctId gets a reference to the given string and assigns it to the DisplayAcctId field.
func (o *UserPreferenceAccount) SetDisplayAcctId(v string) {
	o.DisplayAcctId = &v
}

// GetAutoPositionEffect returns the AutoPositionEffect field value if set, zero value otherwise.
func (o *UserPreferenceAccount) GetAutoPositionEffect() bool {
	if o == nil || IsNil(o.AutoPositionEffect) {
		var ret bool
		return ret
	}
	return *o.AutoPositionEffect
}

// GetAutoPositionEffectOk returns a tuple with the AutoPositionEffect field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreferenceAccount) GetAutoPositionEffectOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoPositionEffect) {
		return nil, false
	}
	return o.AutoPositionEffect, true
}

// HasAutoPositionEffect returns a boolean if a field has been set.
func (o *UserPreferenceAccount) HasAutoPositionEffect() bool {
	if o != nil && !IsNil(o.AutoPositionEffect) {
		return true
	}

	return false
}

// SetAutoPositionEffect gets a reference to the given bool and assigns it to the AutoPositionEffect field.
func (o *UserPreferenceAccount) SetAutoPositionEffect(v bool) {
	o.AutoPositionEffect = &v
}

func (o UserPreferenceAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPreferenceAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.PrimaryAccount) {
		toSerialize["primaryAccount"] = o.PrimaryAccount
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.NickName) {
		toSerialize["nickName"] = o.NickName
	}
	if !IsNil(o.AccountColor) {
		toSerialize["accountColor"] = o.AccountColor
	}
	if !IsNil(o.DisplayAcctId) {
		toSerialize["displayAcctId"] = o.DisplayAcctId
	}
	if !IsNil(o.AutoPositionEffect) {
		toSerialize["autoPositionEffect"] = o.AutoPositionEffect
	}
	return toSerialize, nil
}

type NullableUserPreferenceAccount struct {
	value *UserPreferenceAccount
	isSet bool
}

func (v NullableUserPreferenceAccount) Get() *UserPreferenceAccount {
	return v.value
}

func (v *NullableUserPreferenceAccount) Set(val *UserPreferenceAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPreferenceAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPreferenceAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPreferenceAccount(val *UserPreferenceAccount) *NullableUserPreferenceAccount {
	return &NullableUserPreferenceAccount{value: val, isSet: true}
}

func (v NullableUserPreferenceAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPreferenceAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


