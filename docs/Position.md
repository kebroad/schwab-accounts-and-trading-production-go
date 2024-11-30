# Position

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortQuantity** | Pointer to **float64** |  | [optional] 
**AveragePrice** | Pointer to **float64** |  | [optional] 
**CurrentDayProfitLoss** | Pointer to **float64** |  | [optional] 
**CurrentDayProfitLossPercentage** | Pointer to **float64** |  | [optional] 
**LongQuantity** | Pointer to **float64** |  | [optional] 
**SettledLongQuantity** | Pointer to **float64** |  | [optional] 
**SettledShortQuantity** | Pointer to **float64** |  | [optional] 
**AgedQuantity** | Pointer to **float64** |  | [optional] 
**Instrument** | Pointer to **interface{}** |  | [optional] 
**MarketValue** | Pointer to **float64** |  | [optional] 
**MaintenanceRequirement** | Pointer to **float64** |  | [optional] 
**AverageLongPrice** | Pointer to **float64** |  | [optional] 
**AverageShortPrice** | Pointer to **float64** |  | [optional] 
**TaxLotAverageLongPrice** | Pointer to **float64** |  | [optional] 
**TaxLotAverageShortPrice** | Pointer to **float64** |  | [optional] 
**LongOpenProfitLoss** | Pointer to **float64** |  | [optional] 
**ShortOpenProfitLoss** | Pointer to **float64** |  | [optional] 
**PreviousSessionLongQuantity** | Pointer to **float64** |  | [optional] 
**PreviousSessionShortQuantity** | Pointer to **float64** |  | [optional] 
**CurrentDayCost** | Pointer to **float64** |  | [optional] 

## Methods

### NewPosition

`func NewPosition() *Position`

NewPosition instantiates a new Position object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPositionWithDefaults

`func NewPositionWithDefaults() *Position`

NewPositionWithDefaults instantiates a new Position object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortQuantity

`func (o *Position) GetShortQuantity() float64`

GetShortQuantity returns the ShortQuantity field if non-nil, zero value otherwise.

### GetShortQuantityOk

`func (o *Position) GetShortQuantityOk() (*float64, bool)`

GetShortQuantityOk returns a tuple with the ShortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortQuantity

`func (o *Position) SetShortQuantity(v float64)`

SetShortQuantity sets ShortQuantity field to given value.

### HasShortQuantity

`func (o *Position) HasShortQuantity() bool`

HasShortQuantity returns a boolean if a field has been set.

### GetAveragePrice

`func (o *Position) GetAveragePrice() float64`

GetAveragePrice returns the AveragePrice field if non-nil, zero value otherwise.

### GetAveragePriceOk

`func (o *Position) GetAveragePriceOk() (*float64, bool)`

GetAveragePriceOk returns a tuple with the AveragePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAveragePrice

`func (o *Position) SetAveragePrice(v float64)`

SetAveragePrice sets AveragePrice field to given value.

### HasAveragePrice

`func (o *Position) HasAveragePrice() bool`

HasAveragePrice returns a boolean if a field has been set.

### GetCurrentDayProfitLoss

`func (o *Position) GetCurrentDayProfitLoss() float64`

GetCurrentDayProfitLoss returns the CurrentDayProfitLoss field if non-nil, zero value otherwise.

### GetCurrentDayProfitLossOk

`func (o *Position) GetCurrentDayProfitLossOk() (*float64, bool)`

GetCurrentDayProfitLossOk returns a tuple with the CurrentDayProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDayProfitLoss

`func (o *Position) SetCurrentDayProfitLoss(v float64)`

SetCurrentDayProfitLoss sets CurrentDayProfitLoss field to given value.

### HasCurrentDayProfitLoss

`func (o *Position) HasCurrentDayProfitLoss() bool`

HasCurrentDayProfitLoss returns a boolean if a field has been set.

### GetCurrentDayProfitLossPercentage

`func (o *Position) GetCurrentDayProfitLossPercentage() float64`

GetCurrentDayProfitLossPercentage returns the CurrentDayProfitLossPercentage field if non-nil, zero value otherwise.

### GetCurrentDayProfitLossPercentageOk

`func (o *Position) GetCurrentDayProfitLossPercentageOk() (*float64, bool)`

GetCurrentDayProfitLossPercentageOk returns a tuple with the CurrentDayProfitLossPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDayProfitLossPercentage

`func (o *Position) SetCurrentDayProfitLossPercentage(v float64)`

SetCurrentDayProfitLossPercentage sets CurrentDayProfitLossPercentage field to given value.

### HasCurrentDayProfitLossPercentage

`func (o *Position) HasCurrentDayProfitLossPercentage() bool`

HasCurrentDayProfitLossPercentage returns a boolean if a field has been set.

### GetLongQuantity

`func (o *Position) GetLongQuantity() float64`

GetLongQuantity returns the LongQuantity field if non-nil, zero value otherwise.

### GetLongQuantityOk

`func (o *Position) GetLongQuantityOk() (*float64, bool)`

GetLongQuantityOk returns a tuple with the LongQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongQuantity

`func (o *Position) SetLongQuantity(v float64)`

SetLongQuantity sets LongQuantity field to given value.

### HasLongQuantity

`func (o *Position) HasLongQuantity() bool`

HasLongQuantity returns a boolean if a field has been set.

### GetSettledLongQuantity

`func (o *Position) GetSettledLongQuantity() float64`

GetSettledLongQuantity returns the SettledLongQuantity field if non-nil, zero value otherwise.

### GetSettledLongQuantityOk

`func (o *Position) GetSettledLongQuantityOk() (*float64, bool)`

GetSettledLongQuantityOk returns a tuple with the SettledLongQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledLongQuantity

`func (o *Position) SetSettledLongQuantity(v float64)`

SetSettledLongQuantity sets SettledLongQuantity field to given value.

### HasSettledLongQuantity

`func (o *Position) HasSettledLongQuantity() bool`

HasSettledLongQuantity returns a boolean if a field has been set.

### GetSettledShortQuantity

`func (o *Position) GetSettledShortQuantity() float64`

GetSettledShortQuantity returns the SettledShortQuantity field if non-nil, zero value otherwise.

### GetSettledShortQuantityOk

`func (o *Position) GetSettledShortQuantityOk() (*float64, bool)`

GetSettledShortQuantityOk returns a tuple with the SettledShortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettledShortQuantity

`func (o *Position) SetSettledShortQuantity(v float64)`

SetSettledShortQuantity sets SettledShortQuantity field to given value.

### HasSettledShortQuantity

`func (o *Position) HasSettledShortQuantity() bool`

HasSettledShortQuantity returns a boolean if a field has been set.

### GetAgedQuantity

`func (o *Position) GetAgedQuantity() float64`

GetAgedQuantity returns the AgedQuantity field if non-nil, zero value otherwise.

### GetAgedQuantityOk

`func (o *Position) GetAgedQuantityOk() (*float64, bool)`

GetAgedQuantityOk returns a tuple with the AgedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgedQuantity

`func (o *Position) SetAgedQuantity(v float64)`

SetAgedQuantity sets AgedQuantity field to given value.

### HasAgedQuantity

`func (o *Position) HasAgedQuantity() bool`

HasAgedQuantity returns a boolean if a field has been set.

### GetInstrument

`func (o *Position) GetInstrument() interface{}`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *Position) GetInstrumentOk() (*interface{}, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *Position) SetInstrument(v interface{})`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *Position) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### SetInstrumentNil

`func (o *Position) SetInstrumentNil(b bool)`

 SetInstrumentNil sets the value for Instrument to be an explicit nil

### UnsetInstrument
`func (o *Position) UnsetInstrument()`

UnsetInstrument ensures that no value is present for Instrument, not even an explicit nil
### GetMarketValue

`func (o *Position) GetMarketValue() float64`

GetMarketValue returns the MarketValue field if non-nil, zero value otherwise.

### GetMarketValueOk

`func (o *Position) GetMarketValueOk() (*float64, bool)`

GetMarketValueOk returns a tuple with the MarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketValue

`func (o *Position) SetMarketValue(v float64)`

SetMarketValue sets MarketValue field to given value.

### HasMarketValue

`func (o *Position) HasMarketValue() bool`

HasMarketValue returns a boolean if a field has been set.

### GetMaintenanceRequirement

`func (o *Position) GetMaintenanceRequirement() float64`

GetMaintenanceRequirement returns the MaintenanceRequirement field if non-nil, zero value otherwise.

### GetMaintenanceRequirementOk

`func (o *Position) GetMaintenanceRequirementOk() (*float64, bool)`

GetMaintenanceRequirementOk returns a tuple with the MaintenanceRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceRequirement

`func (o *Position) SetMaintenanceRequirement(v float64)`

SetMaintenanceRequirement sets MaintenanceRequirement field to given value.

### HasMaintenanceRequirement

`func (o *Position) HasMaintenanceRequirement() bool`

HasMaintenanceRequirement returns a boolean if a field has been set.

### GetAverageLongPrice

`func (o *Position) GetAverageLongPrice() float64`

GetAverageLongPrice returns the AverageLongPrice field if non-nil, zero value otherwise.

### GetAverageLongPriceOk

`func (o *Position) GetAverageLongPriceOk() (*float64, bool)`

GetAverageLongPriceOk returns a tuple with the AverageLongPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageLongPrice

`func (o *Position) SetAverageLongPrice(v float64)`

SetAverageLongPrice sets AverageLongPrice field to given value.

### HasAverageLongPrice

`func (o *Position) HasAverageLongPrice() bool`

HasAverageLongPrice returns a boolean if a field has been set.

### GetAverageShortPrice

`func (o *Position) GetAverageShortPrice() float64`

GetAverageShortPrice returns the AverageShortPrice field if non-nil, zero value otherwise.

### GetAverageShortPriceOk

`func (o *Position) GetAverageShortPriceOk() (*float64, bool)`

GetAverageShortPriceOk returns a tuple with the AverageShortPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageShortPrice

`func (o *Position) SetAverageShortPrice(v float64)`

SetAverageShortPrice sets AverageShortPrice field to given value.

### HasAverageShortPrice

`func (o *Position) HasAverageShortPrice() bool`

HasAverageShortPrice returns a boolean if a field has been set.

### GetTaxLotAverageLongPrice

`func (o *Position) GetTaxLotAverageLongPrice() float64`

GetTaxLotAverageLongPrice returns the TaxLotAverageLongPrice field if non-nil, zero value otherwise.

### GetTaxLotAverageLongPriceOk

`func (o *Position) GetTaxLotAverageLongPriceOk() (*float64, bool)`

GetTaxLotAverageLongPriceOk returns a tuple with the TaxLotAverageLongPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotAverageLongPrice

`func (o *Position) SetTaxLotAverageLongPrice(v float64)`

SetTaxLotAverageLongPrice sets TaxLotAverageLongPrice field to given value.

### HasTaxLotAverageLongPrice

`func (o *Position) HasTaxLotAverageLongPrice() bool`

HasTaxLotAverageLongPrice returns a boolean if a field has been set.

### GetTaxLotAverageShortPrice

`func (o *Position) GetTaxLotAverageShortPrice() float64`

GetTaxLotAverageShortPrice returns the TaxLotAverageShortPrice field if non-nil, zero value otherwise.

### GetTaxLotAverageShortPriceOk

`func (o *Position) GetTaxLotAverageShortPriceOk() (*float64, bool)`

GetTaxLotAverageShortPriceOk returns a tuple with the TaxLotAverageShortPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotAverageShortPrice

`func (o *Position) SetTaxLotAverageShortPrice(v float64)`

SetTaxLotAverageShortPrice sets TaxLotAverageShortPrice field to given value.

### HasTaxLotAverageShortPrice

`func (o *Position) HasTaxLotAverageShortPrice() bool`

HasTaxLotAverageShortPrice returns a boolean if a field has been set.

### GetLongOpenProfitLoss

`func (o *Position) GetLongOpenProfitLoss() float64`

GetLongOpenProfitLoss returns the LongOpenProfitLoss field if non-nil, zero value otherwise.

### GetLongOpenProfitLossOk

`func (o *Position) GetLongOpenProfitLossOk() (*float64, bool)`

GetLongOpenProfitLossOk returns a tuple with the LongOpenProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongOpenProfitLoss

`func (o *Position) SetLongOpenProfitLoss(v float64)`

SetLongOpenProfitLoss sets LongOpenProfitLoss field to given value.

### HasLongOpenProfitLoss

`func (o *Position) HasLongOpenProfitLoss() bool`

HasLongOpenProfitLoss returns a boolean if a field has been set.

### GetShortOpenProfitLoss

`func (o *Position) GetShortOpenProfitLoss() float64`

GetShortOpenProfitLoss returns the ShortOpenProfitLoss field if non-nil, zero value otherwise.

### GetShortOpenProfitLossOk

`func (o *Position) GetShortOpenProfitLossOk() (*float64, bool)`

GetShortOpenProfitLossOk returns a tuple with the ShortOpenProfitLoss field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOpenProfitLoss

`func (o *Position) SetShortOpenProfitLoss(v float64)`

SetShortOpenProfitLoss sets ShortOpenProfitLoss field to given value.

### HasShortOpenProfitLoss

`func (o *Position) HasShortOpenProfitLoss() bool`

HasShortOpenProfitLoss returns a boolean if a field has been set.

### GetPreviousSessionLongQuantity

`func (o *Position) GetPreviousSessionLongQuantity() float64`

GetPreviousSessionLongQuantity returns the PreviousSessionLongQuantity field if non-nil, zero value otherwise.

### GetPreviousSessionLongQuantityOk

`func (o *Position) GetPreviousSessionLongQuantityOk() (*float64, bool)`

GetPreviousSessionLongQuantityOk returns a tuple with the PreviousSessionLongQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSessionLongQuantity

`func (o *Position) SetPreviousSessionLongQuantity(v float64)`

SetPreviousSessionLongQuantity sets PreviousSessionLongQuantity field to given value.

### HasPreviousSessionLongQuantity

`func (o *Position) HasPreviousSessionLongQuantity() bool`

HasPreviousSessionLongQuantity returns a boolean if a field has been set.

### GetPreviousSessionShortQuantity

`func (o *Position) GetPreviousSessionShortQuantity() float64`

GetPreviousSessionShortQuantity returns the PreviousSessionShortQuantity field if non-nil, zero value otherwise.

### GetPreviousSessionShortQuantityOk

`func (o *Position) GetPreviousSessionShortQuantityOk() (*float64, bool)`

GetPreviousSessionShortQuantityOk returns a tuple with the PreviousSessionShortQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviousSessionShortQuantity

`func (o *Position) SetPreviousSessionShortQuantity(v float64)`

SetPreviousSessionShortQuantity sets PreviousSessionShortQuantity field to given value.

### HasPreviousSessionShortQuantity

`func (o *Position) HasPreviousSessionShortQuantity() bool`

HasPreviousSessionShortQuantity returns a boolean if a field has been set.

### GetCurrentDayCost

`func (o *Position) GetCurrentDayCost() float64`

GetCurrentDayCost returns the CurrentDayCost field if non-nil, zero value otherwise.

### GetCurrentDayCostOk

`func (o *Position) GetCurrentDayCostOk() (*float64, bool)`

GetCurrentDayCostOk returns a tuple with the CurrentDayCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentDayCost

`func (o *Position) SetCurrentDayCost(v float64)`

SetCurrentDayCost sets CurrentDayCost field to given value.

### HasCurrentDayCost

`func (o *Position) HasCurrentDayCost() bool`

HasCurrentDayCost returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


