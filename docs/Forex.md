# Forex

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**BaseCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**CounterCurrency** | Pointer to [**Currency**](Currency.md) |  | [optional] 
**AssetType** | **string** |  | 
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewForex

`func NewForex(assetType string, ) *Forex`

NewForex instantiates a new Forex object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForexWithDefaults

`func NewForexWithDefaults() *Forex`

NewForexWithDefaults instantiates a new Forex object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *Forex) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Forex) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Forex) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Forex) HasType() bool`

HasType returns a boolean if a field has been set.

### GetBaseCurrency

`func (o *Forex) GetBaseCurrency() Currency`

GetBaseCurrency returns the BaseCurrency field if non-nil, zero value otherwise.

### GetBaseCurrencyOk

`func (o *Forex) GetBaseCurrencyOk() (*Currency, bool)`

GetBaseCurrencyOk returns a tuple with the BaseCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseCurrency

`func (o *Forex) SetBaseCurrency(v Currency)`

SetBaseCurrency sets BaseCurrency field to given value.

### HasBaseCurrency

`func (o *Forex) HasBaseCurrency() bool`

HasBaseCurrency returns a boolean if a field has been set.

### GetCounterCurrency

`func (o *Forex) GetCounterCurrency() Currency`

GetCounterCurrency returns the CounterCurrency field if non-nil, zero value otherwise.

### GetCounterCurrencyOk

`func (o *Forex) GetCounterCurrencyOk() (*Currency, bool)`

GetCounterCurrencyOk returns a tuple with the CounterCurrency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterCurrency

`func (o *Forex) SetCounterCurrency(v Currency)`

SetCounterCurrency sets CounterCurrency field to given value.

### HasCounterCurrency

`func (o *Forex) HasCounterCurrency() bool`

HasCounterCurrency returns a boolean if a field has been set.

### GetAssetType

`func (o *Forex) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *Forex) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *Forex) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCusip

`func (o *Forex) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *Forex) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *Forex) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *Forex) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *Forex) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *Forex) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *Forex) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *Forex) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *Forex) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Forex) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Forex) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Forex) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentId

`func (o *Forex) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *Forex) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *Forex) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *Forex) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetNetChange

`func (o *Forex) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *Forex) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *Forex) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *Forex) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


