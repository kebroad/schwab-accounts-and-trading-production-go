# AccountFixedIncome

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaturityDate** | Pointer to **time.Time** |  | [optional] 
**Factor** | Pointer to **float64** |  | [optional] 
**VariableRate** | Pointer to **float64** |  | [optional] 
**AssetType** | **string** |  | 
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewAccountFixedIncome

`func NewAccountFixedIncome(assetType string, ) *AccountFixedIncome`

NewAccountFixedIncome instantiates a new AccountFixedIncome object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountFixedIncomeWithDefaults

`func NewAccountFixedIncomeWithDefaults() *AccountFixedIncome`

NewAccountFixedIncomeWithDefaults instantiates a new AccountFixedIncome object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaturityDate

`func (o *AccountFixedIncome) GetMaturityDate() time.Time`

GetMaturityDate returns the MaturityDate field if non-nil, zero value otherwise.

### GetMaturityDateOk

`func (o *AccountFixedIncome) GetMaturityDateOk() (*time.Time, bool)`

GetMaturityDateOk returns a tuple with the MaturityDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturityDate

`func (o *AccountFixedIncome) SetMaturityDate(v time.Time)`

SetMaturityDate sets MaturityDate field to given value.

### HasMaturityDate

`func (o *AccountFixedIncome) HasMaturityDate() bool`

HasMaturityDate returns a boolean if a field has been set.

### GetFactor

`func (o *AccountFixedIncome) GetFactor() float64`

GetFactor returns the Factor field if non-nil, zero value otherwise.

### GetFactorOk

`func (o *AccountFixedIncome) GetFactorOk() (*float64, bool)`

GetFactorOk returns a tuple with the Factor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFactor

`func (o *AccountFixedIncome) SetFactor(v float64)`

SetFactor sets Factor field to given value.

### HasFactor

`func (o *AccountFixedIncome) HasFactor() bool`

HasFactor returns a boolean if a field has been set.

### GetVariableRate

`func (o *AccountFixedIncome) GetVariableRate() float64`

GetVariableRate returns the VariableRate field if non-nil, zero value otherwise.

### GetVariableRateOk

`func (o *AccountFixedIncome) GetVariableRateOk() (*float64, bool)`

GetVariableRateOk returns a tuple with the VariableRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariableRate

`func (o *AccountFixedIncome) SetVariableRate(v float64)`

SetVariableRate sets VariableRate field to given value.

### HasVariableRate

`func (o *AccountFixedIncome) HasVariableRate() bool`

HasVariableRate returns a boolean if a field has been set.

### GetAssetType

`func (o *AccountFixedIncome) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AccountFixedIncome) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AccountFixedIncome) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCusip

`func (o *AccountFixedIncome) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *AccountFixedIncome) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *AccountFixedIncome) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *AccountFixedIncome) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *AccountFixedIncome) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountFixedIncome) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountFixedIncome) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountFixedIncome) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *AccountFixedIncome) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountFixedIncome) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountFixedIncome) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountFixedIncome) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentId

`func (o *AccountFixedIncome) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *AccountFixedIncome) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *AccountFixedIncome) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *AccountFixedIncome) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetNetChange

`func (o *AccountFixedIncome) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *AccountFixedIncome) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *AccountFixedIncome) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *AccountFixedIncome) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


