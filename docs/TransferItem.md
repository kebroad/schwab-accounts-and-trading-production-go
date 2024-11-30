# TransferItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instrument** | Pointer to **interface{}** |  | [optional] 
**Amount** | Pointer to **float64** |  | [optional] 
**Cost** | Pointer to **float64** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**FeeType** | Pointer to **string** |  | [optional] 
**PositionEffect** | Pointer to **string** |  | [optional] 

## Methods

### NewTransferItem

`func NewTransferItem() *TransferItem`

NewTransferItem instantiates a new TransferItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferItemWithDefaults

`func NewTransferItemWithDefaults() *TransferItem`

NewTransferItemWithDefaults instantiates a new TransferItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstrument

`func (o *TransferItem) GetInstrument() interface{}`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *TransferItem) GetInstrumentOk() (*interface{}, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *TransferItem) SetInstrument(v interface{})`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *TransferItem) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### SetInstrumentNil

`func (o *TransferItem) SetInstrumentNil(b bool)`

 SetInstrumentNil sets the value for Instrument to be an explicit nil

### UnsetInstrument
`func (o *TransferItem) UnsetInstrument()`

UnsetInstrument ensures that no value is present for Instrument, not even an explicit nil
### GetAmount

`func (o *TransferItem) GetAmount() float64`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferItem) GetAmountOk() (*float64, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferItem) SetAmount(v float64)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *TransferItem) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetCost

`func (o *TransferItem) GetCost() float64`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *TransferItem) GetCostOk() (*float64, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *TransferItem) SetCost(v float64)`

SetCost sets Cost field to given value.

### HasCost

`func (o *TransferItem) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetPrice

`func (o *TransferItem) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TransferItem) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TransferItem) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TransferItem) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetFeeType

`func (o *TransferItem) GetFeeType() string`

GetFeeType returns the FeeType field if non-nil, zero value otherwise.

### GetFeeTypeOk

`func (o *TransferItem) GetFeeTypeOk() (*string, bool)`

GetFeeTypeOk returns a tuple with the FeeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeType

`func (o *TransferItem) SetFeeType(v string)`

SetFeeType sets FeeType field to given value.

### HasFeeType

`func (o *TransferItem) HasFeeType() bool`

HasFeeType returns a boolean if a field has been set.

### GetPositionEffect

`func (o *TransferItem) GetPositionEffect() string`

GetPositionEffect returns the PositionEffect field if non-nil, zero value otherwise.

### GetPositionEffectOk

`func (o *TransferItem) GetPositionEffectOk() (*string, bool)`

GetPositionEffectOk returns a tuple with the PositionEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionEffect

`func (o *TransferItem) SetPositionEffect(v string)`

SetPositionEffect sets PositionEffect field to given value.

### HasPositionEffect

`func (o *TransferItem) HasPositionEffect() bool`

HasPositionEffect returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


