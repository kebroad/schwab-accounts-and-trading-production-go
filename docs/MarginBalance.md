# MarginBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableFunds** | Pointer to **float64** |  | [optional] 
**AvailableFundsNonMarginableTrade** | Pointer to **float64** |  | [optional] 
**BuyingPower** | Pointer to **float64** |  | [optional] 
**BuyingPowerNonMarginableTrade** | Pointer to **float64** |  | [optional] 
**DayTradingBuyingPower** | Pointer to **float64** |  | [optional] 
**DayTradingBuyingPowerCall** | Pointer to **float64** |  | [optional] 
**Equity** | Pointer to **float64** |  | [optional] 
**EquityPercentage** | Pointer to **float64** |  | [optional] 
**LongMarginValue** | Pointer to **float64** |  | [optional] 
**MaintenanceCall** | Pointer to **float64** |  | [optional] 
**MaintenanceRequirement** | Pointer to **float64** |  | [optional] 
**MarginBalance** | Pointer to **float64** |  | [optional] 
**RegTCall** | Pointer to **float64** |  | [optional] 
**ShortBalance** | Pointer to **float64** |  | [optional] 
**ShortMarginValue** | Pointer to **float64** |  | [optional] 
**Sma** | Pointer to **float64** |  | [optional] 
**IsInCall** | Pointer to **float64** |  | [optional] 
**StockBuyingPower** | Pointer to **float64** |  | [optional] 
**OptionBuyingPower** | Pointer to **float64** |  | [optional] 

## Methods

### NewMarginBalance

`func NewMarginBalance() *MarginBalance`

NewMarginBalance instantiates a new MarginBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginBalanceWithDefaults

`func NewMarginBalanceWithDefaults() *MarginBalance`

NewMarginBalanceWithDefaults instantiates a new MarginBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableFunds

`func (o *MarginBalance) GetAvailableFunds() float64`

GetAvailableFunds returns the AvailableFunds field if non-nil, zero value otherwise.

### GetAvailableFundsOk

`func (o *MarginBalance) GetAvailableFundsOk() (*float64, bool)`

GetAvailableFundsOk returns a tuple with the AvailableFunds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFunds

`func (o *MarginBalance) SetAvailableFunds(v float64)`

SetAvailableFunds sets AvailableFunds field to given value.

### HasAvailableFunds

`func (o *MarginBalance) HasAvailableFunds() bool`

HasAvailableFunds returns a boolean if a field has been set.

### GetAvailableFundsNonMarginableTrade

`func (o *MarginBalance) GetAvailableFundsNonMarginableTrade() float64`

GetAvailableFundsNonMarginableTrade returns the AvailableFundsNonMarginableTrade field if non-nil, zero value otherwise.

### GetAvailableFundsNonMarginableTradeOk

`func (o *MarginBalance) GetAvailableFundsNonMarginableTradeOk() (*float64, bool)`

GetAvailableFundsNonMarginableTradeOk returns a tuple with the AvailableFundsNonMarginableTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFundsNonMarginableTrade

`func (o *MarginBalance) SetAvailableFundsNonMarginableTrade(v float64)`

SetAvailableFundsNonMarginableTrade sets AvailableFundsNonMarginableTrade field to given value.

### HasAvailableFundsNonMarginableTrade

`func (o *MarginBalance) HasAvailableFundsNonMarginableTrade() bool`

HasAvailableFundsNonMarginableTrade returns a boolean if a field has been set.

### GetBuyingPower

`func (o *MarginBalance) GetBuyingPower() float64`

GetBuyingPower returns the BuyingPower field if non-nil, zero value otherwise.

### GetBuyingPowerOk

`func (o *MarginBalance) GetBuyingPowerOk() (*float64, bool)`

GetBuyingPowerOk returns a tuple with the BuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPower

`func (o *MarginBalance) SetBuyingPower(v float64)`

SetBuyingPower sets BuyingPower field to given value.

### HasBuyingPower

`func (o *MarginBalance) HasBuyingPower() bool`

HasBuyingPower returns a boolean if a field has been set.

### GetBuyingPowerNonMarginableTrade

`func (o *MarginBalance) GetBuyingPowerNonMarginableTrade() float64`

GetBuyingPowerNonMarginableTrade returns the BuyingPowerNonMarginableTrade field if non-nil, zero value otherwise.

### GetBuyingPowerNonMarginableTradeOk

`func (o *MarginBalance) GetBuyingPowerNonMarginableTradeOk() (*float64, bool)`

GetBuyingPowerNonMarginableTradeOk returns a tuple with the BuyingPowerNonMarginableTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPowerNonMarginableTrade

`func (o *MarginBalance) SetBuyingPowerNonMarginableTrade(v float64)`

SetBuyingPowerNonMarginableTrade sets BuyingPowerNonMarginableTrade field to given value.

### HasBuyingPowerNonMarginableTrade

`func (o *MarginBalance) HasBuyingPowerNonMarginableTrade() bool`

HasBuyingPowerNonMarginableTrade returns a boolean if a field has been set.

### GetDayTradingBuyingPower

`func (o *MarginBalance) GetDayTradingBuyingPower() float64`

GetDayTradingBuyingPower returns the DayTradingBuyingPower field if non-nil, zero value otherwise.

### GetDayTradingBuyingPowerOk

`func (o *MarginBalance) GetDayTradingBuyingPowerOk() (*float64, bool)`

GetDayTradingBuyingPowerOk returns a tuple with the DayTradingBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingBuyingPower

`func (o *MarginBalance) SetDayTradingBuyingPower(v float64)`

SetDayTradingBuyingPower sets DayTradingBuyingPower field to given value.

### HasDayTradingBuyingPower

`func (o *MarginBalance) HasDayTradingBuyingPower() bool`

HasDayTradingBuyingPower returns a boolean if a field has been set.

### GetDayTradingBuyingPowerCall

`func (o *MarginBalance) GetDayTradingBuyingPowerCall() float64`

GetDayTradingBuyingPowerCall returns the DayTradingBuyingPowerCall field if non-nil, zero value otherwise.

### GetDayTradingBuyingPowerCallOk

`func (o *MarginBalance) GetDayTradingBuyingPowerCallOk() (*float64, bool)`

GetDayTradingBuyingPowerCallOk returns a tuple with the DayTradingBuyingPowerCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingBuyingPowerCall

`func (o *MarginBalance) SetDayTradingBuyingPowerCall(v float64)`

SetDayTradingBuyingPowerCall sets DayTradingBuyingPowerCall field to given value.

### HasDayTradingBuyingPowerCall

`func (o *MarginBalance) HasDayTradingBuyingPowerCall() bool`

HasDayTradingBuyingPowerCall returns a boolean if a field has been set.

### GetEquity

`func (o *MarginBalance) GetEquity() float64`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *MarginBalance) GetEquityOk() (*float64, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *MarginBalance) SetEquity(v float64)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *MarginBalance) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetEquityPercentage

`func (o *MarginBalance) GetEquityPercentage() float64`

GetEquityPercentage returns the EquityPercentage field if non-nil, zero value otherwise.

### GetEquityPercentageOk

`func (o *MarginBalance) GetEquityPercentageOk() (*float64, bool)`

GetEquityPercentageOk returns a tuple with the EquityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPercentage

`func (o *MarginBalance) SetEquityPercentage(v float64)`

SetEquityPercentage sets EquityPercentage field to given value.

### HasEquityPercentage

`func (o *MarginBalance) HasEquityPercentage() bool`

HasEquityPercentage returns a boolean if a field has been set.

### GetLongMarginValue

`func (o *MarginBalance) GetLongMarginValue() float64`

GetLongMarginValue returns the LongMarginValue field if non-nil, zero value otherwise.

### GetLongMarginValueOk

`func (o *MarginBalance) GetLongMarginValueOk() (*float64, bool)`

GetLongMarginValueOk returns a tuple with the LongMarginValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMarginValue

`func (o *MarginBalance) SetLongMarginValue(v float64)`

SetLongMarginValue sets LongMarginValue field to given value.

### HasLongMarginValue

`func (o *MarginBalance) HasLongMarginValue() bool`

HasLongMarginValue returns a boolean if a field has been set.

### GetMaintenanceCall

`func (o *MarginBalance) GetMaintenanceCall() float64`

GetMaintenanceCall returns the MaintenanceCall field if non-nil, zero value otherwise.

### GetMaintenanceCallOk

`func (o *MarginBalance) GetMaintenanceCallOk() (*float64, bool)`

GetMaintenanceCallOk returns a tuple with the MaintenanceCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCall

`func (o *MarginBalance) SetMaintenanceCall(v float64)`

SetMaintenanceCall sets MaintenanceCall field to given value.

### HasMaintenanceCall

`func (o *MarginBalance) HasMaintenanceCall() bool`

HasMaintenanceCall returns a boolean if a field has been set.

### GetMaintenanceRequirement

`func (o *MarginBalance) GetMaintenanceRequirement() float64`

GetMaintenanceRequirement returns the MaintenanceRequirement field if non-nil, zero value otherwise.

### GetMaintenanceRequirementOk

`func (o *MarginBalance) GetMaintenanceRequirementOk() (*float64, bool)`

GetMaintenanceRequirementOk returns a tuple with the MaintenanceRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceRequirement

`func (o *MarginBalance) SetMaintenanceRequirement(v float64)`

SetMaintenanceRequirement sets MaintenanceRequirement field to given value.

### HasMaintenanceRequirement

`func (o *MarginBalance) HasMaintenanceRequirement() bool`

HasMaintenanceRequirement returns a boolean if a field has been set.

### GetMarginBalance

`func (o *MarginBalance) GetMarginBalance() float64`

GetMarginBalance returns the MarginBalance field if non-nil, zero value otherwise.

### GetMarginBalanceOk

`func (o *MarginBalance) GetMarginBalanceOk() (*float64, bool)`

GetMarginBalanceOk returns a tuple with the MarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBalance

`func (o *MarginBalance) SetMarginBalance(v float64)`

SetMarginBalance sets MarginBalance field to given value.

### HasMarginBalance

`func (o *MarginBalance) HasMarginBalance() bool`

HasMarginBalance returns a boolean if a field has been set.

### GetRegTCall

`func (o *MarginBalance) GetRegTCall() float64`

GetRegTCall returns the RegTCall field if non-nil, zero value otherwise.

### GetRegTCallOk

`func (o *MarginBalance) GetRegTCallOk() (*float64, bool)`

GetRegTCallOk returns a tuple with the RegTCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegTCall

`func (o *MarginBalance) SetRegTCall(v float64)`

SetRegTCall sets RegTCall field to given value.

### HasRegTCall

`func (o *MarginBalance) HasRegTCall() bool`

HasRegTCall returns a boolean if a field has been set.

### GetShortBalance

`func (o *MarginBalance) GetShortBalance() float64`

GetShortBalance returns the ShortBalance field if non-nil, zero value otherwise.

### GetShortBalanceOk

`func (o *MarginBalance) GetShortBalanceOk() (*float64, bool)`

GetShortBalanceOk returns a tuple with the ShortBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortBalance

`func (o *MarginBalance) SetShortBalance(v float64)`

SetShortBalance sets ShortBalance field to given value.

### HasShortBalance

`func (o *MarginBalance) HasShortBalance() bool`

HasShortBalance returns a boolean if a field has been set.

### GetShortMarginValue

`func (o *MarginBalance) GetShortMarginValue() float64`

GetShortMarginValue returns the ShortMarginValue field if non-nil, zero value otherwise.

### GetShortMarginValueOk

`func (o *MarginBalance) GetShortMarginValueOk() (*float64, bool)`

GetShortMarginValueOk returns a tuple with the ShortMarginValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMarginValue

`func (o *MarginBalance) SetShortMarginValue(v float64)`

SetShortMarginValue sets ShortMarginValue field to given value.

### HasShortMarginValue

`func (o *MarginBalance) HasShortMarginValue() bool`

HasShortMarginValue returns a boolean if a field has been set.

### GetSma

`func (o *MarginBalance) GetSma() float64`

GetSma returns the Sma field if non-nil, zero value otherwise.

### GetSmaOk

`func (o *MarginBalance) GetSmaOk() (*float64, bool)`

GetSmaOk returns a tuple with the Sma field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSma

`func (o *MarginBalance) SetSma(v float64)`

SetSma sets Sma field to given value.

### HasSma

`func (o *MarginBalance) HasSma() bool`

HasSma returns a boolean if a field has been set.

### GetIsInCall

`func (o *MarginBalance) GetIsInCall() float64`

GetIsInCall returns the IsInCall field if non-nil, zero value otherwise.

### GetIsInCallOk

`func (o *MarginBalance) GetIsInCallOk() (*float64, bool)`

GetIsInCallOk returns a tuple with the IsInCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInCall

`func (o *MarginBalance) SetIsInCall(v float64)`

SetIsInCall sets IsInCall field to given value.

### HasIsInCall

`func (o *MarginBalance) HasIsInCall() bool`

HasIsInCall returns a boolean if a field has been set.

### GetStockBuyingPower

`func (o *MarginBalance) GetStockBuyingPower() float64`

GetStockBuyingPower returns the StockBuyingPower field if non-nil, zero value otherwise.

### GetStockBuyingPowerOk

`func (o *MarginBalance) GetStockBuyingPowerOk() (*float64, bool)`

GetStockBuyingPowerOk returns a tuple with the StockBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockBuyingPower

`func (o *MarginBalance) SetStockBuyingPower(v float64)`

SetStockBuyingPower sets StockBuyingPower field to given value.

### HasStockBuyingPower

`func (o *MarginBalance) HasStockBuyingPower() bool`

HasStockBuyingPower returns a boolean if a field has been set.

### GetOptionBuyingPower

`func (o *MarginBalance) GetOptionBuyingPower() float64`

GetOptionBuyingPower returns the OptionBuyingPower field if non-nil, zero value otherwise.

### GetOptionBuyingPowerOk

`func (o *MarginBalance) GetOptionBuyingPowerOk() (*float64, bool)`

GetOptionBuyingPowerOk returns a tuple with the OptionBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionBuyingPower

`func (o *MarginBalance) SetOptionBuyingPower(v float64)`

SetOptionBuyingPower sets OptionBuyingPower field to given value.

### HasOptionBuyingPower

`func (o *MarginBalance) HasOptionBuyingPower() bool`

HasOptionBuyingPower returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


