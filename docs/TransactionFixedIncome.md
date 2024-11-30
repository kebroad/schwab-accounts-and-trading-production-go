# TransactionFixedIncome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**MaturityDate** | Pointer to **time.Time** |  | [optional] 
**Factor** | Pointer to **float64** |  | [optional] 
**Multiplier** | Pointer to **float64** |  | [optional] 
**VariableRate** | Pointer to **float64** |  | [optional] 
**AssetType** | **string** |  | 
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewTransactionFixedIncome

`func NewTransactionFixedIncome(assetType string, ) *TransactionFixedIncome`

NewTransactionFixedIncome instantiates a new TransactionFixedIncome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionFixedIncomeWithDefaults

`func NewTransactionFixedIncomeWithDefaults() *TransactionFixedIncome`

NewTransactionFixedIncomeWithDefaults instantiates a new TransactionFixedIncome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionFixedIncome) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionFixedIncome) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionFixedIncome) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionFixedIncome) HasType() bool`

HasType returns a boolean if a field has been set.

### GetMaturityDate

`func (o *TransactionFixedIncome) GetMaturityDate() time.Time`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *TransactionFixedIncome) GetMaturityDateOk() (*time.Time, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *TransactionFixedIncome) SetMaturityDate(v time.Time)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *TransactionFixedIncome) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetFactor

`func (o *TransactionFixedIncome) GetFactor() float64`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *TransactionFixedIncome) GetFactorOk() (*float64, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *TransactionFixedIncome) SetFactor(v float64)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *TransactionFixedIncome) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetMultiplier

`func (o *TransactionFixedIncome) GetMultiplier() float64`

GetMultiplier returns the Multiplier field if non-nil, zero value otherwise.

### GetMultiplierOk

`func (o *TransactionFixedIncome) GetMultiplierOk() (*float64, bool)`

GetMultiplierOk returns a tuple with the Multiplier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiplier

`func (o *TransactionFixedIncome) SetMultiplier(v float64)`

SetMultiplier sets Multiplier field to given value.

### HasMultiplier

`func (o *TransactionFixedIncome) HasMultiplier() bool`

HasMultiplier returns a boolean if a field has been set.

### GetVariableRate

`func (o *TransactionFixedIncome) GetVariableRate() float64`

GetVariableRate returns the VariableRate field if non-nil, zero value otherwise.

### GetVariableRateOk

`func (o *TransactionFixedIncome) GetVariableRateOk() (*float64, bool)`

GetVariableRateOk returns a tuple with the VariableRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableRate

`func (o *TransactionFixedIncome) SetVariableRate(v float64)`

SetVariableRate sets VariableRate field to given value.

### HasVariableRate

`func (o *TransactionFixedIncome) HasVariableRate() bool`

HasVariableRate returns a boolean if a field has been set.

### GetAssetType

`func (o *TransactionFixedIncome) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *TransactionFixedIncome) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *TransactionFixedIncome) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCusip

`func (o *TransactionFixedIncome) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *TransactionFixedIncome) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *TransactionFixedIncome) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *TransactionFixedIncome) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *TransactionFixedIncome) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransactionFixedIncome) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransactionFixedIncome) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TransactionFixedIncome) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionFixedIncome) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionFixedIncome) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionFixedIncome) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionFixedIncome) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentId

`func (o *TransactionFixedIncome) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *TransactionFixedIncome) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *TransactionFixedIncome) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *TransactionFixedIncome) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetNetChange

`func (o *TransactionFixedIncome) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *TransactionFixedIncome) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *TransactionFixedIncome) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *TransactionFixedIncome) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


