# AccountOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OptionDeliverables** | Pointer to [**[]AccountAPIOptionDeliverable**](AccountAPIOptionDeliverable.md) |  | [optional] 
**PutCall** | Pointer to **string** |  | [optional] 
**OptionMultiplier** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UnderlyingSymbol** | Pointer to **string** |  | [optional] 
**AssetType** | **string** |  | 
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewAccountOption

`func NewAccountOption(assetType string, ) *AccountOption`

NewAccountOption instantiates a new AccountOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountOptionWithDefaults

`func NewAccountOptionWithDefaults() *AccountOption`

NewAccountOptionWithDefaults instantiates a new AccountOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptionDeliverables

`func (o *AccountOption) GetOptionDeliverables() []AccountAPIOptionDeliverable`

GetOptionDeliverables returns the OptionDeliverables field if non-nil, zero value otherwise.

### GetOptionDeliverablesOk

`func (o *AccountOption) GetOptionDeliverablesOk() (*[]AccountAPIOptionDeliverable, bool)`

GetOptionDeliverablesOk returns a tuple with the OptionDeliverables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDeliverables

`func (o *AccountOption) SetOptionDeliverables(v []AccountAPIOptionDeliverable)`

SetOptionDeliverables sets OptionDeliverables field to given value.

### HasOptionDeliverables

`func (o *AccountOption) HasOptionDeliverables() bool`

HasOptionDeliverables returns a boolean if a field has been set.

### GetPutCall

`func (o *AccountOption) GetPutCall() string`

GetPutCall returns the PutCall field if non-nil, zero value otherwise.

### GetPutCallOk

`func (o *AccountOption) GetPutCallOk() (*string, bool)`

GetPutCallOk returns a tuple with the PutCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutCall

`func (o *AccountOption) SetPutCall(v string)`

SetPutCall sets PutCall field to given value.

### HasPutCall

`func (o *AccountOption) HasPutCall() bool`

HasPutCall returns a boolean if a field has been set.

### GetOptionMultiplier

`func (o *AccountOption) GetOptionMultiplier() int32`

GetOptionMultiplier returns the OptionMultiplier field if non-nil, zero value otherwise.

### GetOptionMultiplierOk

`func (o *AccountOption) GetOptionMultiplierOk() (*int32, bool)`

GetOptionMultiplierOk returns a tuple with the OptionMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionMultiplier

`func (o *AccountOption) SetOptionMultiplier(v int32)`

SetOptionMultiplier sets OptionMultiplier field to given value.

### HasOptionMultiplier

`func (o *AccountOption) HasOptionMultiplier() bool`

HasOptionMultiplier returns a boolean if a field has been set.

### GetType

`func (o *AccountOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AccountOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AccountOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AccountOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnderlyingSymbol

`func (o *AccountOption) GetUnderlyingSymbol() string`

GetUnderlyingSymbol returns the UnderlyingSymbol field if non-nil, zero value otherwise.

### GetUnderlyingSymbolOk

`func (o *AccountOption) GetUnderlyingSymbolOk() (*string, bool)`

GetUnderlyingSymbolOk returns a tuple with the UnderlyingSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSymbol

`func (o *AccountOption) SetUnderlyingSymbol(v string)`

SetUnderlyingSymbol sets UnderlyingSymbol field to given value.

### HasUnderlyingSymbol

`func (o *AccountOption) HasUnderlyingSymbol() bool`

HasUnderlyingSymbol returns a boolean if a field has been set.

### GetAssetType

`func (o *AccountOption) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AccountOption) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AccountOption) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCusip

`func (o *AccountOption) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *AccountOption) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *AccountOption) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *AccountOption) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *AccountOption) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountOption) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountOption) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountOption) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *AccountOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentId

`func (o *AccountOption) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *AccountOption) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *AccountOption) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *AccountOption) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetNetChange

`func (o *AccountOption) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *AccountOption) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *AccountOption) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *AccountOption) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


