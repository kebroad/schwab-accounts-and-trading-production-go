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

// checks if the Offer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Offer{}

// Offer struct for Offer
type Offer struct {
	Level2Permissions *bool `json:"level2Permissions,omitempty"`
	MktDataPermission *string `json:"mktDataPermission,omitempty"`
}

// NewOffer instantiates a new Offer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOffer() *Offer {
	this := Offer{}
	var level2Permissions bool = false
	this.Level2Permissions = &level2Permissions
	return &this
}

// NewOfferWithDefaults instantiates a new Offer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOfferWithDefaults() *Offer {
	this := Offer{}
	var level2Permissions bool = false
	this.Level2Permissions = &level2Permissions
	return &this
}

// GetLevel2Permissions returns the Level2Permissions field value if set, zero value otherwise.
func (o *Offer) GetLevel2Permissions() bool {
	if o == nil || IsNil(o.Level2Permissions) {
		var ret bool
		return ret
	}
	return *o.Level2Permissions
}

// GetLevel2PermissionsOk returns a tuple with the Level2Permissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offer) GetLevel2PermissionsOk() (*bool, bool) {
	if o == nil || IsNil(o.Level2Permissions) {
		return nil, false
	}
	return o.Level2Permissions, true
}

// HasLevel2Permissions returns a boolean if a field has been set.
func (o *Offer) HasLevel2Permissions() bool {
	if o != nil && !IsNil(o.Level2Permissions) {
		return true
	}

	return false
}

// SetLevel2Permissions gets a reference to the given bool and assigns it to the Level2Permissions field.
func (o *Offer) SetLevel2Permissions(v bool) {
	o.Level2Permissions = &v
}

// GetMktDataPermission returns the MktDataPermission field value if set, zero value otherwise.
func (o *Offer) GetMktDataPermission() string {
	if o == nil || IsNil(o.MktDataPermission) {
		var ret string
		return ret
	}
	return *o.MktDataPermission
}

// GetMktDataPermissionOk returns a tuple with the MktDataPermission field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Offer) GetMktDataPermissionOk() (*string, bool) {
	if o == nil || IsNil(o.MktDataPermission) {
		return nil, false
	}
	return o.MktDataPermission, true
}

// HasMktDataPermission returns a boolean if a field has been set.
func (o *Offer) HasMktDataPermission() bool {
	if o != nil && !IsNil(o.MktDataPermission) {
		return true
	}

	return false
}

// SetMktDataPermission gets a reference to the given string and assigns it to the MktDataPermission field.
func (o *Offer) SetMktDataPermission(v string) {
	o.MktDataPermission = &v
}

func (o Offer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Offer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Level2Permissions) {
		toSerialize["level2Permissions"] = o.Level2Permissions
	}
	if !IsNil(o.MktDataPermission) {
		toSerialize["mktDataPermission"] = o.MktDataPermission
	}
	return toSerialize, nil
}

type NullableOffer struct {
	value *Offer
	isSet bool
}

func (v NullableOffer) Get() *Offer {
	return v.value
}

func (v *NullableOffer) Set(val *Offer) {
	v.value = val
	v.isSet = true
}

func (v NullableOffer) IsSet() bool {
	return v.isSet
}

func (v *NullableOffer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOffer(val *Offer) *NullableOffer {
	return &NullableOffer{value: val, isSet: true}
}

func (v NullableOffer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOffer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


