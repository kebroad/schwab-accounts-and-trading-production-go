# TransactionOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**OptionDeliverables** | Pointer to [**[]TransactionAPIOptionDeliverable**](TransactionAPIOptionDeliverable.md) |  | [optional] 
**OptionPremiumMultiplier** | Pointer to **int64** |  | [optional] 
**PutCall** | Pointer to **string** |  | [optional] 
**StrikePrice** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UnderlyingSymbol** | Pointer to **string** |  | [optional] 
**UnderlyingCusip** | Pointer to **string** |  | [optional] 
**Deliverable** | Pointer to **interface{}** |  | [optional] 
**AssetType** | **string** |  | 
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewTransactionOption

`func NewTransactionOption(assetType string, ) *TransactionOption`

NewTransactionOption instantiates a new TransactionOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionOptionWithDefaults

`func NewTransactionOptionWithDefaults() *TransactionOption`

NewTransactionOptionWithDefaults instantiates a new TransactionOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationDate

`func (o *TransactionOption) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *TransactionOption) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *TransactionOption) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *TransactionOption) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetOptionDeliverables

`func (o *TransactionOption) GetOptionDeliverables() []TransactionAPIOptionDeliverable`

GetOptionDeliverables returns the OptionDeliverables field if non-nil, zero value otherwise.

### GetOptionDeliverablesOk

`func (o *TransactionOption) GetOptionDeliverablesOk() (*[]TransactionAPIOptionDeliverable, bool)`

GetOptionDeliverablesOk returns a tuple with the OptionDeliverables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionDeliverables

`func (o *TransactionOption) SetOptionDeliverables(v []TransactionAPIOptionDeliverable)`

SetOptionDeliverables sets OptionDeliverables field to given value.

### HasOptionDeliverables

`func (o *TransactionOption) HasOptionDeliverables() bool`

HasOptionDeliverables returns a boolean if a field has been set.

### GetOptionPremiumMultiplier

`func (o *TransactionOption) GetOptionPremiumMultiplier() int64`

GetOptionPremiumMultiplier returns the OptionPremiumMultiplier field if non-nil, zero value otherwise.

### GetOptionPremiumMultiplierOk

`func (o *TransactionOption) GetOptionPremiumMultiplierOk() (*int64, bool)`

GetOptionPremiumMultiplierOk returns a tuple with the OptionPremiumMultiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionPremiumMultiplier

`func (o *TransactionOption) SetOptionPremiumMultiplier(v int64)`

SetOptionPremiumMultiplier sets OptionPremiumMultiplier field to given value.

### HasOptionPremiumMultiplier

`func (o *TransactionOption) HasOptionPremiumMultiplier() bool`

HasOptionPremiumMultiplier returns a boolean if a field has been set.

### GetPutCall

`func (o *TransactionOption) GetPutCall() string`

GetPutCall returns the PutCall field if non-nil, zero value otherwise.

### GetPutCallOk

`func (o *TransactionOption) GetPutCallOk() (*string, bool)`

GetPutCallOk returns a tuple with the PutCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPutCall

`func (o *TransactionOption) SetPutCall(v string)`

SetPutCall sets PutCall field to given value.

### HasPutCall

`func (o *TransactionOption) HasPutCall() bool`

HasPutCall returns a boolean if a field has been set.

### GetStrikePrice

`func (o *TransactionOption) GetStrikePrice() float64`

GetStrikePrice returns the StrikePrice field if non-nil, zero value otherwise.

### GetStrikePriceOk

`func (o *TransactionOption) GetStrikePriceOk() (*float64, bool)`

GetStrikePriceOk returns a tuple with the StrikePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrikePrice

`func (o *TransactionOption) SetStrikePrice(v float64)`

SetStrikePrice sets StrikePrice field to given value.

### HasStrikePrice

`func (o *TransactionOption) HasStrikePrice() bool`

HasStrikePrice returns a boolean if a field has been set.

### GetType

`func (o *TransactionOption) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionOption) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionOption) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionOption) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnderlyingSymbol

`func (o *TransactionOption) GetUnderlyingSymbol() string`

GetUnderlyingSymbol returns the UnderlyingSymbol field if non-nil, zero value otherwise.

### GetUnderlyingSymbolOk

`func (o *TransactionOption) GetUnderlyingSymbolOk() (*string, bool)`

GetUnderlyingSymbolOk returns a tuple with the UnderlyingSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingSymbol

`func (o *TransactionOption) SetUnderlyingSymbol(v string)`

SetUnderlyingSymbol sets UnderlyingSymbol field to given value.

### HasUnderlyingSymbol

`func (o *TransactionOption) HasUnderlyingSymbol() bool`

HasUnderlyingSymbol returns a boolean if a field has been set.

### GetUnderlyingCusip

`func (o *TransactionOption) GetUnderlyingCusip() string`

GetUnderlyingCusip returns the UnderlyingCusip field if non-nil, zero value otherwise.

### GetUnderlyingCusipOk

`func (o *TransactionOption) GetUnderlyingCusipOk() (*string, bool)`

GetUnderlyingCusipOk returns a tuple with the UnderlyingCusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnderlyingCusip

`func (o *TransactionOption) SetUnderlyingCusip(v string)`

SetUnderlyingCusip sets UnderlyingCusip field to given value.

### HasUnderlyingCusip

`func (o *TransactionOption) HasUnderlyingCusip() bool`

HasUnderlyingCusip returns a boolean if a field has been set.

### GetDeliverable

`func (o *TransactionOption) GetDeliverable() interface{}`

GetDeliverable returns the Deliverable field if non-nil, zero value otherwise.

### GetDeliverableOk

`func (o *TransactionOption) GetDeliverableOk() (*interface{}, bool)`

GetDeliverableOk returns a tuple with the Deliverable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverable

`func (o *TransactionOption) SetDeliverable(v interface{})`

SetDeliverable sets Deliverable field to given value.

### HasDeliverable

`func (o *TransactionOption) HasDeliverable() bool`

HasDeliverable returns a boolean if a field has been set.

### SetDeliverableNil

`func (o *TransactionOption) SetDeliverableNil(b bool)`

 SetDeliverableNil sets the value for Deliverable to be an explicit nil

### UnsetDeliverable
`func (o *TransactionOption) UnsetDeliverable()`

UnsetDeliverable ensures that no value is present for Deliverable, not even an explicit nil
### GetAssetType

`func (o *TransactionOption) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *TransactionOption) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *TransactionOption) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCusip

`func (o *TransactionOption) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *TransactionOption) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *TransactionOption) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *TransactionOption) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *TransactionOption) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransactionOption) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransactionOption) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TransactionOption) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionOption) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionOption) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionOption) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionOption) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentId

`func (o *TransactionOption) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *TransactionOption) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *TransactionOption) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *TransactionOption) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetNetChange

`func (o *TransactionOption) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *TransactionOption) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *TransactionOption) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *TransactionOption) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


