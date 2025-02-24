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

// checks if the UserPreference type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPreference{}

// UserPreference struct for UserPreference
type UserPreference struct {
	Accounts []UserPreferenceAccount `json:"accounts,omitempty"`
	StreamerInfo []StreamerInfo `json:"streamerInfo,omitempty"`
	Offers []Offer `json:"offers,omitempty"`
}

// NewUserPreference instantiates a new UserPreference object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPreference() *UserPreference {
	this := UserPreference{}
	return &this
}

// NewUserPreferenceWithDefaults instantiates a new UserPreference object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPreferenceWithDefaults() *UserPreference {
	this := UserPreference{}
	return &this
}

// GetAccounts returns the Accounts field value if set, zero value otherwise.
func (o *UserPreference) GetAccounts() []UserPreferenceAccount {
	if o == nil || IsNil(o.Accounts) {
		var ret []UserPreferenceAccount
		return ret
	}
	return o.Accounts
}

// GetAccountsOk returns a tuple with the Accounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreference) GetAccountsOk() ([]UserPreferenceAccount, bool) {
	if o == nil || IsNil(o.Accounts) {
		return nil, false
	}
	return o.Accounts, true
}

// HasAccounts returns a boolean if a field has been set.
func (o *UserPreference) HasAccounts() bool {
	if o != nil && !IsNil(o.Accounts) {
		return true
	}

	return false
}

// SetAccounts gets a reference to the given []UserPreferenceAccount and assigns it to the Accounts field.
func (o *UserPreference) SetAccounts(v []UserPreferenceAccount) {
	o.Accounts = v
}

// GetStreamerInfo returns the StreamerInfo field value if set, zero value otherwise.
func (o *UserPreference) GetStreamerInfo() []StreamerInfo {
	if o == nil || IsNil(o.StreamerInfo) {
		var ret []StreamerInfo
		return ret
	}
	return o.StreamerInfo
}

// GetStreamerInfoOk returns a tuple with the StreamerInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreference) GetStreamerInfoOk() ([]StreamerInfo, bool) {
	if o == nil || IsNil(o.StreamerInfo) {
		return nil, false
	}
	return o.StreamerInfo, true
}

// HasStreamerInfo returns a boolean if a field has been set.
func (o *UserPreference) HasStreamerInfo() bool {
	if o != nil && !IsNil(o.StreamerInfo) {
		return true
	}

	return false
}

// SetStreamerInfo gets a reference to the given []StreamerInfo and assigns it to the StreamerInfo field.
func (o *UserPreference) SetStreamerInfo(v []StreamerInfo) {
	o.StreamerInfo = v
}

// GetOffers returns the Offers field value if set, zero value otherwise.
func (o *UserPreference) GetOffers() []Offer {
	if o == nil || IsNil(o.Offers) {
		var ret []Offer
		return ret
	}
	return o.Offers
}

// GetOffersOk returns a tuple with the Offers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPreference) GetOffersOk() ([]Offer, bool) {
	if o == nil || IsNil(o.Offers) {
		return nil, false
	}
	return o.Offers, true
}

// HasOffers returns a boolean if a field has been set.
func (o *UserPreference) HasOffers() bool {
	if o != nil && !IsNil(o.Offers) {
		return true
	}

	return false
}

// SetOffers gets a reference to the given []Offer and assigns it to the Offers field.
func (o *UserPreference) SetOffers(v []Offer) {
	o.Offers = v
}

func (o UserPreference) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPreference) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Accounts) {
		toSerialize["accounts"] = o.Accounts
	}
	if !IsNil(o.StreamerInfo) {
		toSerialize["streamerInfo"] = o.StreamerInfo
	}
	if !IsNil(o.Offers) {
		toSerialize["offers"] = o.Offers
	}
	return toSerialize, nil
}

type NullableUserPreference struct {
	value *UserPreference
	isSet bool
}

func (v NullableUserPreference) Get() *UserPreference {
	return v.value
}

func (v *NullableUserPreference) Set(val *UserPreference) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPreference) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPreference) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPreference(val *UserPreference) *NullableUserPreference {
	return &NullableUserPreference{value: val, isSet: true}
}

func (v NullableUserPreference) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPreference) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


