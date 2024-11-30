# TransactionMutualFund

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FundFamilyName** | Pointer to **string** |  | [optional] 
**FundFamilySymbol** | Pointer to **string** |  | [optional] 
**FundGroup** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**ExchangeCutoffTime** | Pointer to **time.Time** |  | [optional] 
**PurchaseCutoffTime** | Pointer to **time.Time** |  | [optional] 
**RedemptionCutoffTime** | Pointer to **time.Time** |  | [optional] 
**AssetType** | **string** |  | 
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewTransactionMutualFund

`func NewTransactionMutualFund(assetType string, ) *TransactionMutualFund`

NewTransactionMutualFund instantiates a new TransactionMutualFund object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionMutualFundWithDefaults

`func NewTransactionMutualFundWithDefaults() *TransactionMutualFund`

NewTransactionMutualFundWithDefaults instantiates a new TransactionMutualFund object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFundFamilyName

`func (o *TransactionMutualFund) GetFundFamilyName() string`

GetFundFamilyName returns the FundFamilyName field if non-nil, zero value otherwise.

### GetFundFamilyNameOk

`func (o *TransactionMutualFund) GetFundFamilyNameOk() (*string, bool)`

GetFundFamilyNameOk returns a tuple with the FundFamilyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundFamilyName

`func (o *TransactionMutualFund) SetFundFamilyName(v string)`

SetFundFamilyName sets FundFamilyName field to given value.

### HasFundFamilyName

`func (o *TransactionMutualFund) HasFundFamilyName() bool`

HasFundFamilyName returns a boolean if a field has been set.

### GetFundFamilySymbol

`func (o *TransactionMutualFund) GetFundFamilySymbol() string`

GetFundFamilySymbol returns the FundFamilySymbol field if non-nil, zero value otherwise.

### GetFundFamilySymbolOk

`func (o *TransactionMutualFund) GetFundFamilySymbolOk() (*string, bool)`

GetFundFamilySymbolOk returns a tuple with the FundFamilySymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundFamilySymbol

`func (o *TransactionMutualFund) SetFundFamilySymbol(v string)`

SetFundFamilySymbol sets FundFamilySymbol field to given value.

### HasFundFamilySymbol

`func (o *TransactionMutualFund) HasFundFamilySymbol() bool`

HasFundFamilySymbol returns a boolean if a field has been set.

### GetFundGroup

`func (o *TransactionMutualFund) GetFundGroup() string`

GetFundGroup returns the FundGroup field if non-nil, zero value otherwise.

### GetFundGroupOk

`func (o *TransactionMutualFund) GetFundGroupOk() (*string, bool)`

GetFundGroupOk returns a tuple with the FundGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFundGroup

`func (o *TransactionMutualFund) SetFundGroup(v string)`

SetFundGroup sets FundGroup field to given value.

### HasFundGroup

`func (o *TransactionMutualFund) HasFundGroup() bool`

HasFundGroup returns a boolean if a field has been set.

### GetType

`func (o *TransactionMutualFund) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionMutualFund) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionMutualFund) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionMutualFund) HasType() bool`

HasType returns a boolean if a field has been set.

### GetExchangeCutoffTime

`func (o *TransactionMutualFund) GetExchangeCutoffTime() time.Time`

GetExchangeCutoffTime returns the ExchangeCutoffTime field if non-nil, zero value otherwise.

### GetExchangeCutoffTimeOk

`func (o *TransactionMutualFund) GetExchangeCutoffTimeOk() (*time.Time, bool)`

GetExchangeCutoffTimeOk returns a tuple with the ExchangeCutoffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExchangeCutoffTime

`func (o *TransactionMutualFund) SetExchangeCutoffTime(v time.Time)`

SetExchangeCutoffTime sets ExchangeCutoffTime field to given value.

### HasExchangeCutoffTime

`func (o *TransactionMutualFund) HasExchangeCutoffTime() bool`

HasExchangeCutoffTime returns a boolean if a field has been set.

### GetPurchaseCutoffTime

`func (o *TransactionMutualFund) GetPurchaseCutoffTime() time.Time`

GetPurchaseCutoffTime returns the PurchaseCutoffTime field if non-nil, zero value otherwise.

### GetPurchaseCutoffTimeOk

`func (o *TransactionMutualFund) GetPurchaseCutoffTimeOk() (*time.Time, bool)`

GetPurchaseCutoffTimeOk returns a tuple with the PurchaseCutoffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseCutoffTime

`func (o *TransactionMutualFund) SetPurchaseCutoffTime(v time.Time)`

SetPurchaseCutoffTime sets PurchaseCutoffTime field to given value.

### HasPurchaseCutoffTime

`func (o *TransactionMutualFund) HasPurchaseCutoffTime() bool`

HasPurchaseCutoffTime returns a boolean if a field has been set.

### GetRedemptionCutoffTime

`func (o *TransactionMutualFund) GetRedemptionCutoffTime() time.Time`

GetRedemptionCutoffTime returns the RedemptionCutoffTime field if non-nil, zero value otherwise.

### GetRedemptionCutoffTimeOk

`func (o *TransactionMutualFund) GetRedemptionCutoffTimeOk() (*time.Time, bool)`

GetRedemptionCutoffTimeOk returns a tuple with the RedemptionCutoffTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedemptionCutoffTime

`func (o *TransactionMutualFund) SetRedemptionCutoffTime(v time.Time)`

SetRedemptionCutoffTime sets RedemptionCutoffTime field to given value.

### HasRedemptionCutoffTime

`func (o *TransactionMutualFund) HasRedemptionCutoffTime() bool`

HasRedemptionCutoffTime returns a boolean if a field has been set.

### GetAssetType

`func (o *TransactionMutualFund) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *TransactionMutualFund) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *TransactionMutualFund) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCusip

`func (o *TransactionMutualFund) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *TransactionMutualFund) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *TransactionMutualFund) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *TransactionMutualFund) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *TransactionMutualFund) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransactionMutualFund) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransactionMutualFund) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TransactionMutualFund) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionMutualFund) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionMutualFund) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionMutualFund) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionMutualFund) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentId

`func (o *TransactionMutualFund) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *TransactionMutualFund) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *TransactionMutualFund) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *TransactionMutualFund) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetNetChange

`func (o *TransactionMutualFund) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *TransactionMutualFund) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *TransactionMutualFund) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *TransactionMutualFund) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


