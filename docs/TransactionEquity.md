# TransactionEquity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**AssetType** | **string** |  | 
**Cusip** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**NetChange** | Pointer to **float64** |  | [optional] 

## Methods

### NewTransactionEquity

`func NewTransactionEquity(assetType string, ) *TransactionEquity`

NewTransactionEquity instantiates a new TransactionEquity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionEquityWithDefaults

`func NewTransactionEquityWithDefaults() *TransactionEquity`

NewTransactionEquityWithDefaults instantiates a new TransactionEquity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *TransactionEquity) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransactionEquity) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransactionEquity) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *TransactionEquity) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAssetType

`func (o *TransactionEquity) GetAssetType() string`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *TransactionEquity) GetAssetTypeOk() (*string, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *TransactionEquity) SetAssetType(v string)`

SetAssetType sets AssetType field to given value.


### GetCusip

`func (o *TransactionEquity) GetCusip() string`

GetCusip returns the Cusip field if non-nil, zero value otherwise.

### GetCusipOk

`func (o *TransactionEquity) GetCusipOk() (*string, bool)`

GetCusipOk returns a tuple with the Cusip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCusip

`func (o *TransactionEquity) SetCusip(v string)`

SetCusip sets Cusip field to given value.

### HasCusip

`func (o *TransactionEquity) HasCusip() bool`

HasCusip returns a boolean if a field has been set.

### GetSymbol

`func (o *TransactionEquity) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *TransactionEquity) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *TransactionEquity) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *TransactionEquity) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetDescription

`func (o *TransactionEquity) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransactionEquity) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransactionEquity) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransactionEquity) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetInstrumentId

`func (o *TransactionEquity) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *TransactionEquity) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *TransactionEquity) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *TransactionEquity) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetNetChange

`func (o *TransactionEquity) GetNetChange() float64`

GetNetChange returns the NetChange field if non-nil, zero value otherwise.

### GetNetChangeOk

`func (o *TransactionEquity) GetNetChangeOk() (*float64, bool)`

GetNetChangeOk returns a tuple with the NetChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetChange

`func (o *TransactionEquity) SetNetChange(v float64)`

SetNetChange sets NetChange field to given value.

### HasNetChange

`func (o *TransactionEquity) HasNetChange() bool`

HasNetChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


