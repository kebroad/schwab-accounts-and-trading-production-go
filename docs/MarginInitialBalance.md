# MarginInitialBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedInterest** | Pointer to **float64** |  | [optional] 
**AvailableFundsNonMarginableTrade** | Pointer to **float64** |  | [optional] 
**BondValue** | Pointer to **float64** |  | [optional] 
**BuyingPower** | Pointer to **float64** |  | [optional] 
**CashBalance** | Pointer to **float64** |  | [optional] 
**CashAvailableForTrading** | Pointer to **float64** |  | [optional] 
**CashReceipts** | Pointer to **float64** |  | [optional] 
**DayTradingBuyingPower** | Pointer to **float64** |  | [optional] 
**DayTradingBuyingPowerCall** | Pointer to **float64** |  | [optional] 
**DayTradingEquityCall** | Pointer to **float64** |  | [optional] 
**Equity** | Pointer to **float64** |  | [optional] 
**EquityPercentage** | Pointer to **float64** |  | [optional] 
**LiquidationValue** | Pointer to **float64** |  | [optional] 
**LongMarginValue** | Pointer to **float64** |  | [optional] 
**LongOptionMarketValue** | Pointer to **float64** |  | [optional] 
**LongStockValue** | Pointer to **float64** |  | [optional] 
**MaintenanceCall** | Pointer to **float64** |  | [optional] 
**MaintenanceRequirement** | Pointer to **float64** |  | [optional] 
**Margin** | Pointer to **float64** |  | [optional] 
**MarginEquity** | Pointer to **float64** |  | [optional] 
**MoneyMarketFund** | Pointer to **float64** |  | [optional] 
**MutualFundValue** | Pointer to **float64** |  | [optional] 
**RegTCall** | Pointer to **float64** |  | [optional] 
**ShortMarginValue** | Pointer to **float64** |  | [optional] 
**ShortOptionMarketValue** | Pointer to **float64** |  | [optional] 
**ShortStockValue** | Pointer to **float64** |  | [optional] 
**TotalCash** | Pointer to **float64** |  | [optional] 
**IsInCall** | Pointer to **float64** |  | [optional] 
**UnsettledCash** | Pointer to **float64** |  | [optional] 
**PendingDeposits** | Pointer to **float64** |  | [optional] 
**MarginBalance** | Pointer to **float64** |  | [optional] 
**ShortBalance** | Pointer to **float64** |  | [optional] 
**AccountValue** | Pointer to **float64** |  | [optional] 

## Methods

### NewMarginInitialBalance

`func NewMarginInitialBalance() *MarginInitialBalance`

NewMarginInitialBalance instantiates a new MarginInitialBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginInitialBalanceWithDefaults

`func NewMarginInitialBalanceWithDefaults() *MarginInitialBalance`

NewMarginInitialBalanceWithDefaults instantiates a new MarginInitialBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedInterest

`func (o *MarginInitialBalance) GetAccruedInterest() float64`

GetAccruedInterest returns the AccruedInterest field if non-nil, zero value otherwise.

### GetAccruedInterestOk

`func (o *MarginInitialBalance) GetAccruedInterestOk() (*float64, bool)`

GetAccruedInterestOk returns a tuple with the AccruedInterest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedInterest

`func (o *MarginInitialBalance) SetAccruedInterest(v float64)`

SetAccruedInterest sets AccruedInterest field to given value.

### HasAccruedInterest

`func (o *MarginInitialBalance) HasAccruedInterest() bool`

HasAccruedInterest returns a boolean if a field has been set.

### GetAvailableFundsNonMarginableTrade

`func (o *MarginInitialBalance) GetAvailableFundsNonMarginableTrade() float64`

GetAvailableFundsNonMarginableTrade returns the AvailableFundsNonMarginableTrade field if non-nil, zero value otherwise.

### GetAvailableFundsNonMarginableTradeOk

`func (o *MarginInitialBalance) GetAvailableFundsNonMarginableTradeOk() (*float64, bool)`

GetAvailableFundsNonMarginableTradeOk returns a tuple with the AvailableFundsNonMarginableTrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableFundsNonMarginableTrade

`func (o *MarginInitialBalance) SetAvailableFundsNonMarginableTrade(v float64)`

SetAvailableFundsNonMarginableTrade sets AvailableFundsNonMarginableTrade field to given value.

### HasAvailableFundsNonMarginableTrade

`func (o *MarginInitialBalance) HasAvailableFundsNonMarginableTrade() bool`

HasAvailableFundsNonMarginableTrade returns a boolean if a field has been set.

### GetBondValue

`func (o *MarginInitialBalance) GetBondValue() float64`

GetBondValue returns the BondValue field if non-nil, zero value otherwise.

### GetBondValueOk

`func (o *MarginInitialBalance) GetBondValueOk() (*float64, bool)`

GetBondValueOk returns a tuple with the BondValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBondValue

`func (o *MarginInitialBalance) SetBondValue(v float64)`

SetBondValue sets BondValue field to given value.

### HasBondValue

`func (o *MarginInitialBalance) HasBondValue() bool`

HasBondValue returns a boolean if a field has been set.

### GetBuyingPower

`func (o *MarginInitialBalance) GetBuyingPower() float64`

GetBuyingPower returns the BuyingPower field if non-nil, zero value otherwise.

### GetBuyingPowerOk

`func (o *MarginInitialBalance) GetBuyingPowerOk() (*float64, bool)`

GetBuyingPowerOk returns a tuple with the BuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyingPower

`func (o *MarginInitialBalance) SetBuyingPower(v float64)`

SetBuyingPower sets BuyingPower field to given value.

### HasBuyingPower

`func (o *MarginInitialBalance) HasBuyingPower() bool`

HasBuyingPower returns a boolean if a field has been set.

### GetCashBalance

`func (o *MarginInitialBalance) GetCashBalance() float64`

GetCashBalance returns the CashBalance field if non-nil, zero value otherwise.

### GetCashBalanceOk

`func (o *MarginInitialBalance) GetCashBalanceOk() (*float64, bool)`

GetCashBalanceOk returns a tuple with the CashBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashBalance

`func (o *MarginInitialBalance) SetCashBalance(v float64)`

SetCashBalance sets CashBalance field to given value.

### HasCashBalance

`func (o *MarginInitialBalance) HasCashBalance() bool`

HasCashBalance returns a boolean if a field has been set.

### GetCashAvailableForTrading

`func (o *MarginInitialBalance) GetCashAvailableForTrading() float64`

GetCashAvailableForTrading returns the CashAvailableForTrading field if non-nil, zero value otherwise.

### GetCashAvailableForTradingOk

`func (o *MarginInitialBalance) GetCashAvailableForTradingOk() (*float64, bool)`

GetCashAvailableForTradingOk returns a tuple with the CashAvailableForTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAvailableForTrading

`func (o *MarginInitialBalance) SetCashAvailableForTrading(v float64)`

SetCashAvailableForTrading sets CashAvailableForTrading field to given value.

### HasCashAvailableForTrading

`func (o *MarginInitialBalance) HasCashAvailableForTrading() bool`

HasCashAvailableForTrading returns a boolean if a field has been set.

### GetCashReceipts

`func (o *MarginInitialBalance) GetCashReceipts() float64`

GetCashReceipts returns the CashReceipts field if non-nil, zero value otherwise.

### GetCashReceiptsOk

`func (o *MarginInitialBalance) GetCashReceiptsOk() (*float64, bool)`

GetCashReceiptsOk returns a tuple with the CashReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashReceipts

`func (o *MarginInitialBalance) SetCashReceipts(v float64)`

SetCashReceipts sets CashReceipts field to given value.

### HasCashReceipts

`func (o *MarginInitialBalance) HasCashReceipts() bool`

HasCashReceipts returns a boolean if a field has been set.

### GetDayTradingBuyingPower

`func (o *MarginInitialBalance) GetDayTradingBuyingPower() float64`

GetDayTradingBuyingPower returns the DayTradingBuyingPower field if non-nil, zero value otherwise.

### GetDayTradingBuyingPowerOk

`func (o *MarginInitialBalance) GetDayTradingBuyingPowerOk() (*float64, bool)`

GetDayTradingBuyingPowerOk returns a tuple with the DayTradingBuyingPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingBuyingPower

`func (o *MarginInitialBalance) SetDayTradingBuyingPower(v float64)`

SetDayTradingBuyingPower sets DayTradingBuyingPower field to given value.

### HasDayTradingBuyingPower

`func (o *MarginInitialBalance) HasDayTradingBuyingPower() bool`

HasDayTradingBuyingPower returns a boolean if a field has been set.

### GetDayTradingBuyingPowerCall

`func (o *MarginInitialBalance) GetDayTradingBuyingPowerCall() float64`

GetDayTradingBuyingPowerCall returns the DayTradingBuyingPowerCall field if non-nil, zero value otherwise.

### GetDayTradingBuyingPowerCallOk

`func (o *MarginInitialBalance) GetDayTradingBuyingPowerCallOk() (*float64, bool)`

GetDayTradingBuyingPowerCallOk returns a tuple with the DayTradingBuyingPowerCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingBuyingPowerCall

`func (o *MarginInitialBalance) SetDayTradingBuyingPowerCall(v float64)`

SetDayTradingBuyingPowerCall sets DayTradingBuyingPowerCall field to given value.

### HasDayTradingBuyingPowerCall

`func (o *MarginInitialBalance) HasDayTradingBuyingPowerCall() bool`

HasDayTradingBuyingPowerCall returns a boolean if a field has been set.

### GetDayTradingEquityCall

`func (o *MarginInitialBalance) GetDayTradingEquityCall() float64`

GetDayTradingEquityCall returns the DayTradingEquityCall field if non-nil, zero value otherwise.

### GetDayTradingEquityCallOk

`func (o *MarginInitialBalance) GetDayTradingEquityCallOk() (*float64, bool)`

GetDayTradingEquityCallOk returns a tuple with the DayTradingEquityCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDayTradingEquityCall

`func (o *MarginInitialBalance) SetDayTradingEquityCall(v float64)`

SetDayTradingEquityCall sets DayTradingEquityCall field to given value.

### HasDayTradingEquityCall

`func (o *MarginInitialBalance) HasDayTradingEquityCall() bool`

HasDayTradingEquityCall returns a boolean if a field has been set.

### GetEquity

`func (o *MarginInitialBalance) GetEquity() float64`

GetEquity returns the Equity field if non-nil, zero value otherwise.

### GetEquityOk

`func (o *MarginInitialBalance) GetEquityOk() (*float64, bool)`

GetEquityOk returns a tuple with the Equity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquity

`func (o *MarginInitialBalance) SetEquity(v float64)`

SetEquity sets Equity field to given value.

### HasEquity

`func (o *MarginInitialBalance) HasEquity() bool`

HasEquity returns a boolean if a field has been set.

### GetEquityPercentage

`func (o *MarginInitialBalance) GetEquityPercentage() float64`

GetEquityPercentage returns the EquityPercentage field if non-nil, zero value otherwise.

### GetEquityPercentageOk

`func (o *MarginInitialBalance) GetEquityPercentageOk() (*float64, bool)`

GetEquityPercentageOk returns a tuple with the EquityPercentage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEquityPercentage

`func (o *MarginInitialBalance) SetEquityPercentage(v float64)`

SetEquityPercentage sets EquityPercentage field to given value.

### HasEquityPercentage

`func (o *MarginInitialBalance) HasEquityPercentage() bool`

HasEquityPercentage returns a boolean if a field has been set.

### GetLiquidationValue

`func (o *MarginInitialBalance) GetLiquidationValue() float64`

GetLiquidationValue returns the LiquidationValue field if non-nil, zero value otherwise.

### GetLiquidationValueOk

`func (o *MarginInitialBalance) GetLiquidationValueOk() (*float64, bool)`

GetLiquidationValueOk returns a tuple with the LiquidationValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiquidationValue

`func (o *MarginInitialBalance) SetLiquidationValue(v float64)`

SetLiquidationValue sets LiquidationValue field to given value.

### HasLiquidationValue

`func (o *MarginInitialBalance) HasLiquidationValue() bool`

HasLiquidationValue returns a boolean if a field has been set.

### GetLongMarginValue

`func (o *MarginInitialBalance) GetLongMarginValue() float64`

GetLongMarginValue returns the LongMarginValue field if non-nil, zero value otherwise.

### GetLongMarginValueOk

`func (o *MarginInitialBalance) GetLongMarginValueOk() (*float64, bool)`

GetLongMarginValueOk returns a tuple with the LongMarginValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongMarginValue

`func (o *MarginInitialBalance) SetLongMarginValue(v float64)`

SetLongMarginValue sets LongMarginValue field to given value.

### HasLongMarginValue

`func (o *MarginInitialBalance) HasLongMarginValue() bool`

HasLongMarginValue returns a boolean if a field has been set.

### GetLongOptionMarketValue

`func (o *MarginInitialBalance) GetLongOptionMarketValue() float64`

GetLongOptionMarketValue returns the LongOptionMarketValue field if non-nil, zero value otherwise.

### GetLongOptionMarketValueOk

`func (o *MarginInitialBalance) GetLongOptionMarketValueOk() (*float64, bool)`

GetLongOptionMarketValueOk returns a tuple with the LongOptionMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongOptionMarketValue

`func (o *MarginInitialBalance) SetLongOptionMarketValue(v float64)`

SetLongOptionMarketValue sets LongOptionMarketValue field to given value.

### HasLongOptionMarketValue

`func (o *MarginInitialBalance) HasLongOptionMarketValue() bool`

HasLongOptionMarketValue returns a boolean if a field has been set.

### GetLongStockValue

`func (o *MarginInitialBalance) GetLongStockValue() float64`

GetLongStockValue returns the LongStockValue field if non-nil, zero value otherwise.

### GetLongStockValueOk

`func (o *MarginInitialBalance) GetLongStockValueOk() (*float64, bool)`

GetLongStockValueOk returns a tuple with the LongStockValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongStockValue

`func (o *MarginInitialBalance) SetLongStockValue(v float64)`

SetLongStockValue sets LongStockValue field to given value.

### HasLongStockValue

`func (o *MarginInitialBalance) HasLongStockValue() bool`

HasLongStockValue returns a boolean if a field has been set.

### GetMaintenanceCall

`func (o *MarginInitialBalance) GetMaintenanceCall() float64`

GetMaintenanceCall returns the MaintenanceCall field if non-nil, zero value otherwise.

### GetMaintenanceCallOk

`func (o *MarginInitialBalance) GetMaintenanceCallOk() (*float64, bool)`

GetMaintenanceCallOk returns a tuple with the MaintenanceCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceCall

`func (o *MarginInitialBalance) SetMaintenanceCall(v float64)`

SetMaintenanceCall sets MaintenanceCall field to given value.

### HasMaintenanceCall

`func (o *MarginInitialBalance) HasMaintenanceCall() bool`

HasMaintenanceCall returns a boolean if a field has been set.

### GetMaintenanceRequirement

`func (o *MarginInitialBalance) GetMaintenanceRequirement() float64`

GetMaintenanceRequirement returns the MaintenanceRequirement field if non-nil, zero value otherwise.

### GetMaintenanceRequirementOk

`func (o *MarginInitialBalance) GetMaintenanceRequirementOk() (*float64, bool)`

GetMaintenanceRequirementOk returns a tuple with the MaintenanceRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaintenanceRequirement

`func (o *MarginInitialBalance) SetMaintenanceRequirement(v float64)`

SetMaintenanceRequirement sets MaintenanceRequirement field to given value.

### HasMaintenanceRequirement

`func (o *MarginInitialBalance) HasMaintenanceRequirement() bool`

HasMaintenanceRequirement returns a boolean if a field has been set.

### GetMargin

`func (o *MarginInitialBalance) GetMargin() float64`

GetMargin returns the Margin field if non-nil, zero value otherwise.

### GetMarginOk

`func (o *MarginInitialBalance) GetMarginOk() (*float64, bool)`

GetMarginOk returns a tuple with the Margin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMargin

`func (o *MarginInitialBalance) SetMargin(v float64)`

SetMargin sets Margin field to given value.

### HasMargin

`func (o *MarginInitialBalance) HasMargin() bool`

HasMargin returns a boolean if a field has been set.

### GetMarginEquity

`func (o *MarginInitialBalance) GetMarginEquity() float64`

GetMarginEquity returns the MarginEquity field if non-nil, zero value otherwise.

### GetMarginEquityOk

`func (o *MarginInitialBalance) GetMarginEquityOk() (*float64, bool)`

GetMarginEquityOk returns a tuple with the MarginEquity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginEquity

`func (o *MarginInitialBalance) SetMarginEquity(v float64)`

SetMarginEquity sets MarginEquity field to given value.

### HasMarginEquity

`func (o *MarginInitialBalance) HasMarginEquity() bool`

HasMarginEquity returns a boolean if a field has been set.

### GetMoneyMarketFund

`func (o *MarginInitialBalance) GetMoneyMarketFund() float64`

GetMoneyMarketFund returns the MoneyMarketFund field if non-nil, zero value otherwise.

### GetMoneyMarketFundOk

`func (o *MarginInitialBalance) GetMoneyMarketFundOk() (*float64, bool)`

GetMoneyMarketFundOk returns a tuple with the MoneyMarketFund field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoneyMarketFund

`func (o *MarginInitialBalance) SetMoneyMarketFund(v float64)`

SetMoneyMarketFund sets MoneyMarketFund field to given value.

### HasMoneyMarketFund

`func (o *MarginInitialBalance) HasMoneyMarketFund() bool`

HasMoneyMarketFund returns a boolean if a field has been set.

### GetMutualFundValue

`func (o *MarginInitialBalance) GetMutualFundValue() float64`

GetMutualFundValue returns the MutualFundValue field if non-nil, zero value otherwise.

### GetMutualFundValueOk

`func (o *MarginInitialBalance) GetMutualFundValueOk() (*float64, bool)`

GetMutualFundValueOk returns a tuple with the MutualFundValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMutualFundValue

`func (o *MarginInitialBalance) SetMutualFundValue(v float64)`

SetMutualFundValue sets MutualFundValue field to given value.

### HasMutualFundValue

`func (o *MarginInitialBalance) HasMutualFundValue() bool`

HasMutualFundValue returns a boolean if a field has been set.

### GetRegTCall

`func (o *MarginInitialBalance) GetRegTCall() float64`

GetRegTCall returns the RegTCall field if non-nil, zero value otherwise.

### GetRegTCallOk

`func (o *MarginInitialBalance) GetRegTCallOk() (*float64, bool)`

GetRegTCallOk returns a tuple with the RegTCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegTCall

`func (o *MarginInitialBalance) SetRegTCall(v float64)`

SetRegTCall sets RegTCall field to given value.

### HasRegTCall

`func (o *MarginInitialBalance) HasRegTCall() bool`

HasRegTCall returns a boolean if a field has been set.

### GetShortMarginValue

`func (o *MarginInitialBalance) GetShortMarginValue() float64`

GetShortMarginValue returns the ShortMarginValue field if non-nil, zero value otherwise.

### GetShortMarginValueOk

`func (o *MarginInitialBalance) GetShortMarginValueOk() (*float64, bool)`

GetShortMarginValueOk returns a tuple with the ShortMarginValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMarginValue

`func (o *MarginInitialBalance) SetShortMarginValue(v float64)`

SetShortMarginValue sets ShortMarginValue field to given value.

### HasShortMarginValue

`func (o *MarginInitialBalance) HasShortMarginValue() bool`

HasShortMarginValue returns a boolean if a field has been set.

### GetShortOptionMarketValue

`func (o *MarginInitialBalance) GetShortOptionMarketValue() float64`

GetShortOptionMarketValue returns the ShortOptionMarketValue field if non-nil, zero value otherwise.

### GetShortOptionMarketValueOk

`func (o *MarginInitialBalance) GetShortOptionMarketValueOk() (*float64, bool)`

GetShortOptionMarketValueOk returns a tuple with the ShortOptionMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortOptionMarketValue

`func (o *MarginInitialBalance) SetShortOptionMarketValue(v float64)`

SetShortOptionMarketValue sets ShortOptionMarketValue field to given value.

### HasShortOptionMarketValue

`func (o *MarginInitialBalance) HasShortOptionMarketValue() bool`

HasShortOptionMarketValue returns a boolean if a field has been set.

### GetShortStockValue

`func (o *MarginInitialBalance) GetShortStockValue() float64`

GetShortStockValue returns the ShortStockValue field if non-nil, zero value otherwise.

### GetShortStockValueOk

`func (o *MarginInitialBalance) GetShortStockValueOk() (*float64, bool)`

GetShortStockValueOk returns a tuple with the ShortStockValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortStockValue

`func (o *MarginInitialBalance) SetShortStockValue(v float64)`

SetShortStockValue sets ShortStockValue field to given value.

### HasShortStockValue

`func (o *MarginInitialBalance) HasShortStockValue() bool`

HasShortStockValue returns a boolean if a field has been set.

### GetTotalCash

`func (o *MarginInitialBalance) GetTotalCash() float64`

GetTotalCash returns the TotalCash field if non-nil, zero value otherwise.

### GetTotalCashOk

`func (o *MarginInitialBalance) GetTotalCashOk() (*float64, bool)`

GetTotalCashOk returns a tuple with the TotalCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCash

`func (o *MarginInitialBalance) SetTotalCash(v float64)`

SetTotalCash sets TotalCash field to given value.

### HasTotalCash

`func (o *MarginInitialBalance) HasTotalCash() bool`

HasTotalCash returns a boolean if a field has been set.

### GetIsInCall

`func (o *MarginInitialBalance) GetIsInCall() float64`

GetIsInCall returns the IsInCall field if non-nil, zero value otherwise.

### GetIsInCallOk

`func (o *MarginInitialBalance) GetIsInCallOk() (*float64, bool)`

GetIsInCallOk returns a tuple with the IsInCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInCall

`func (o *MarginInitialBalance) SetIsInCall(v float64)`

SetIsInCall sets IsInCall field to given value.

### HasIsInCall

`func (o *MarginInitialBalance) HasIsInCall() bool`

HasIsInCall returns a boolean if a field has been set.

### GetUnsettledCash

`func (o *MarginInitialBalance) GetUnsettledCash() float64`

GetUnsettledCash returns the UnsettledCash field if non-nil, zero value otherwise.

### GetUnsettledCashOk

`func (o *MarginInitialBalance) GetUnsettledCashOk() (*float64, bool)`

GetUnsettledCashOk returns a tuple with the UnsettledCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsettledCash

`func (o *MarginInitialBalance) SetUnsettledCash(v float64)`

SetUnsettledCash sets UnsettledCash field to given value.

### HasUnsettledCash

`func (o *MarginInitialBalance) HasUnsettledCash() bool`

HasUnsettledCash returns a boolean if a field has been set.

### GetPendingDeposits

`func (o *MarginInitialBalance) GetPendingDeposits() float64`

GetPendingDeposits returns the PendingDeposits field if non-nil, zero value otherwise.

### GetPendingDepositsOk

`func (o *MarginInitialBalance) GetPendingDepositsOk() (*float64, bool)`

GetPendingDepositsOk returns a tuple with the PendingDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingDeposits

`func (o *MarginInitialBalance) SetPendingDeposits(v float64)`

SetPendingDeposits sets PendingDeposits field to given value.

### HasPendingDeposits

`func (o *MarginInitialBalance) HasPendingDeposits() bool`

HasPendingDeposits returns a boolean if a field has been set.

### GetMarginBalance

`func (o *MarginInitialBalance) GetMarginBalance() float64`

GetMarginBalance returns the MarginBalance field if non-nil, zero value otherwise.

### GetMarginBalanceOk

`func (o *MarginInitialBalance) GetMarginBalanceOk() (*float64, bool)`

GetMarginBalanceOk returns a tuple with the MarginBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBalance

`func (o *MarginInitialBalance) SetMarginBalance(v float64)`

SetMarginBalance sets MarginBalance field to given value.

### HasMarginBalance

`func (o *MarginInitialBalance) HasMarginBalance() bool`

HasMarginBalance returns a boolean if a field has been set.

### GetShortBalance

`func (o *MarginInitialBalance) GetShortBalance() float64`

GetShortBalance returns the ShortBalance field if non-nil, zero value otherwise.

### GetShortBalanceOk

`func (o *MarginInitialBalance) GetShortBalanceOk() (*float64, bool)`

GetShortBalanceOk returns a tuple with the ShortBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortBalance

`func (o *MarginInitialBalance) SetShortBalance(v float64)`

SetShortBalance sets ShortBalance field to given value.

### HasShortBalance

`func (o *MarginInitialBalance) HasShortBalance() bool`

HasShortBalance returns a boolean if a field has been set.

### GetAccountValue

`func (o *MarginInitialBalance) GetAccountValue() float64`

GetAccountValue returns the AccountValue field if non-nil, zero value otherwise.

### GetAccountValueOk

`func (o *MarginInitialBalance) GetAccountValueOk() (*float64, bool)`

GetAccountValueOk returns a tuple with the AccountValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountValue

`func (o *MarginInitialBalance) SetAccountValue(v float64)`

SetAccountValue sets AccountValue field to given value.

### HasAccountValue

`func (o *MarginInitialBalance) HasAccountValue() bool`

HasAccountValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


