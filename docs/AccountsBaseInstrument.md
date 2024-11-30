# AccountsBaseInstrument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetType** | **string** |  | 
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewAccountsBaseInstrument

`func NewAccountsBaseInstrument(assetType string, ) *AccountsBaseInstrument`

NewAccountsBaseInstrument instantiates a new AccountsBaseInstrument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsBaseInstrumentWithDefaults

`func NewAccountsBaseInstrumentWithDefaults() *AccountsBaseInstrument`

NewAccountsBaseInstrumentWithDefaults instantiates a new AccountsBaseInstrument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetType

`func (o *AccountsBaseInstrument) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *AccountsBaseInstrument) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *AccountsBaseInstrument) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCusip

`func (o *AccountsBaseInstrument) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *AccountsBaseInstrument) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *AccountsBaseInstrument) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *AccountsBaseInstrument) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *AccountsBaseInstrument) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *AccountsBaseInstrument) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *AccountsBaseInstrument) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *AccountsBaseInstrument) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *AccountsBaseInstrument) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AccountsBaseInstrument) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AccountsBaseInstrument) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AccountsBaseInstrument) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentId

`func (o *AccountsBaseInstrument) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *AccountsBaseInstrument) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *AccountsBaseInstrument) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *AccountsBaseInstrument) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetNetChange

`func (o *AccountsBaseInstrument) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *AccountsBaseInstrument) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *AccountsBaseInstrument) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *AccountsBaseInstrument) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


