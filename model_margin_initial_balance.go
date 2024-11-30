/*
Trader API - Account Access and User Preferences

Schwab Trader API access to Account, Order entry and User Preferences

API version: 1.0.0
Contact: TraderAPI@Schwab.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
)

// checks if the MarginInitialBalance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MarginInitialBalance{}

// MarginInitialBalance struct for MarginInitialBalance
type MarginInitialBalance struct {
	AccruedInterest *float64 `json:"accruedInterest,omitempty"`
	AvailableFundsNonMarginableTrade *float64 `json:"availableFundsNonMarginableTrade,omitempty"`
	BondValue *float64 `json:"bondValue,omitempty"`
	BuyingPower *float64 `json:"buyingPower,omitempty"`
	CashBalance *float64 `json:"cashBalance,omitempty"`
	CashAvailableForTrading *float64 `json:"cashAvailableForTrading,omitempty"`
	CashReceipts *float64 `json:"cashReceipts,omitempty"`
	DayTradingBuyingPower *float64 `json:"dayTradingBuyingPower,omitempty"`
	DayTradingBuyingPowerCall *float64 `json:"dayTradingBuyingPowerCall,omitempty"`
	DayTradingEquityCall *float64 `json:"dayTradingEquityCall,omitempty"`
	Equity *float64 `json:"equity,omitempty"`
	EquityPercentage *float64 `json:"equityPercentage,omitempty"`
	LiquidationValue *float64 `json:"liquidationValue,omitempty"`
	LongMarginValue *float64 `json:"longMarginValue,omitempty"`
	LongOptionMarketValue *float64 `json:"longOptionMarketValue,omitempty"`
	LongStockValue *float64 `json:"longStockValue,omitempty"`
	MaintenanceCall *float64 `json:"maintenanceCall,omitempty"`
	MaintenanceRequirement *float64 `json:"maintenanceRequirement,omitempty"`
	Margin *float64 `json:"margin,omitempty"`
	MarginEquity *float64 `json:"marginEquity,omitempty"`
	MoneyMarketFund *float64 `json:"moneyMarketFund,omitempty"`
	MutualFundValue *float64 `json:"mutualFundValue,omitempty"`
	RegTCall *float64 `json:"regTCall,omitempty"`
	ShortMarginValue *float64 `json:"shortMarginValue,omitempty"`
	ShortOptionMarketValue *float64 `json:"shortOptionMarketValue,omitempty"`
	ShortStockValue *float64 `json:"shortStockValue,omitempty"`
	TotalCash *float64 `json:"totalCash,omitempty"`
	IsInCall *float64 `json:"isInCall,omitempty"`
	UnsettledCash *float64 `json:"unsettledCash,omitempty"`
	PendingDeposits *float64 `json:"pendingDeposits,omitempty"`
	MarginBalance *float64 `json:"marginBalance,omitempty"`
	ShortBalance *float64 `json:"shortBalance,omitempty"`
	AccountValue *float64 `json:"accountValue,omitempty"`
}

// NewMarginInitialBalance instantiates a new MarginInitialBalance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarginInitialBalance() *MarginInitialBalance {
	this := MarginInitialBalance{}
	return &this
}

// NewMarginInitialBalanceWithDefaults instantiates a new MarginInitialBalance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarginInitialBalanceWithDefaults() *MarginInitialBalance {
	this := MarginInitialBalance{}
	return &this
}

// GetAccruedInterest returns the AccruedInterest field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetAccruedInterest() float64 {
	if o == nil || IsNil(o.AccruedInterest) {
		var ret float64
		return ret
	}
	return *o.AccruedInterest
}

// GetAccruedInterestOk returns a tuple with the AccruedInterest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetAccruedInterestOk() (*float64, bool) {
	if o == nil || IsNil(o.AccruedInterest) {
		return nil, false
	}
	return o.AccruedInterest, true
}

// HasAccruedInterest returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasAccruedInterest() bool {
	if o != nil && !IsNil(o.AccruedInterest) {
		return true
	}

	return false
}

// SetAccruedInterest gets a reference to the given float64 and assigns it to the AccruedInterest field.
func (o *MarginInitialBalance) SetAccruedInterest(v float64) {
	o.AccruedInterest = &v
}

// GetAvailableFundsNonMarginableTrade returns the AvailableFundsNonMarginableTrade field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetAvailableFundsNonMarginableTrade() float64 {
	if o == nil || IsNil(o.AvailableFundsNonMarginableTrade) {
		var ret float64
		return ret
	}
	return *o.AvailableFundsNonMarginableTrade
}

// GetAvailableFundsNonMarginableTradeOk returns a tuple with the AvailableFundsNonMarginableTrade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetAvailableFundsNonMarginableTradeOk() (*float64, bool) {
	if o == nil || IsNil(o.AvailableFundsNonMarginableTrade) {
		return nil, false
	}
	return o.AvailableFundsNonMarginableTrade, true
}

// HasAvailableFundsNonMarginableTrade returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasAvailableFundsNonMarginableTrade() bool {
	if o != nil && !IsNil(o.AvailableFundsNonMarginableTrade) {
		return true
	}

	return false
}

// SetAvailableFundsNonMarginableTrade gets a reference to the given float64 and assigns it to the AvailableFundsNonMarginableTrade field.
func (o *MarginInitialBalance) SetAvailableFundsNonMarginableTrade(v float64) {
	o.AvailableFundsNonMarginableTrade = &v
}

// GetBondValue returns the BondValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetBondValue() float64 {
	if o == nil || IsNil(o.BondValue) {
		var ret float64
		return ret
	}
	return *o.BondValue
}

// GetBondValueOk returns a tuple with the BondValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetBondValueOk() (*float64, bool) {
	if o == nil || IsNil(o.BondValue) {
		return nil, false
	}
	return o.BondValue, true
}

// HasBondValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasBondValue() bool {
	if o != nil && !IsNil(o.BondValue) {
		return true
	}

	return false
}

// SetBondValue gets a reference to the given float64 and assigns it to the BondValue field.
func (o *MarginInitialBalance) SetBondValue(v float64) {
	o.BondValue = &v
}

// GetBuyingPower returns the BuyingPower field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetBuyingPower() float64 {
	if o == nil || IsNil(o.BuyingPower) {
		var ret float64
		return ret
	}
	return *o.BuyingPower
}

// GetBuyingPowerOk returns a tuple with the BuyingPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetBuyingPowerOk() (*float64, bool) {
	if o == nil || IsNil(o.BuyingPower) {
		return nil, false
	}
	return o.BuyingPower, true
}

// HasBuyingPower returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasBuyingPower() bool {
	if o != nil && !IsNil(o.BuyingPower) {
		return true
	}

	return false
}

// SetBuyingPower gets a reference to the given float64 and assigns it to the BuyingPower field.
func (o *MarginInitialBalance) SetBuyingPower(v float64) {
	o.BuyingPower = &v
}

// GetCashBalance returns the CashBalance field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetCashBalance() float64 {
	if o == nil || IsNil(o.CashBalance) {
		var ret float64
		return ret
	}
	return *o.CashBalance
}

// GetCashBalanceOk returns a tuple with the CashBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetCashBalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.CashBalance) {
		return nil, false
	}
	return o.CashBalance, true
}

// HasCashBalance returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasCashBalance() bool {
	if o != nil && !IsNil(o.CashBalance) {
		return true
	}

	return false
}

// SetCashBalance gets a reference to the given float64 and assigns it to the CashBalance field.
func (o *MarginInitialBalance) SetCashBalance(v float64) {
	o.CashBalance = &v
}

// GetCashAvailableForTrading returns the CashAvailableForTrading field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetCashAvailableForTrading() float64 {
	if o == nil || IsNil(o.CashAvailableForTrading) {
		var ret float64
		return ret
	}
	return *o.CashAvailableForTrading
}

// GetCashAvailableForTradingOk returns a tuple with the CashAvailableForTrading field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetCashAvailableForTradingOk() (*float64, bool) {
	if o == nil || IsNil(o.CashAvailableForTrading) {
		return nil, false
	}
	return o.CashAvailableForTrading, true
}

// HasCashAvailableForTrading returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasCashAvailableForTrading() bool {
	if o != nil && !IsNil(o.CashAvailableForTrading) {
		return true
	}

	return false
}

// SetCashAvailableForTrading gets a reference to the given float64 and assigns it to the CashAvailableForTrading field.
func (o *MarginInitialBalance) SetCashAvailableForTrading(v float64) {
	o.CashAvailableForTrading = &v
}

// GetCashReceipts returns the CashReceipts field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetCashReceipts() float64 {
	if o == nil || IsNil(o.CashReceipts) {
		var ret float64
		return ret
	}
	return *o.CashReceipts
}

// GetCashReceiptsOk returns a tuple with the CashReceipts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetCashReceiptsOk() (*float64, bool) {
	if o == nil || IsNil(o.CashReceipts) {
		return nil, false
	}
	return o.CashReceipts, true
}

// HasCashReceipts returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasCashReceipts() bool {
	if o != nil && !IsNil(o.CashReceipts) {
		return true
	}

	return false
}

// SetCashReceipts gets a reference to the given float64 and assigns it to the CashReceipts field.
func (o *MarginInitialBalance) SetCashReceipts(v float64) {
	o.CashReceipts = &v
}

// GetDayTradingBuyingPower returns the DayTradingBuyingPower field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetDayTradingBuyingPower() float64 {
	if o == nil || IsNil(o.DayTradingBuyingPower) {
		var ret float64
		return ret
	}
	return *o.DayTradingBuyingPower
}

// GetDayTradingBuyingPowerOk returns a tuple with the DayTradingBuyingPower field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetDayTradingBuyingPowerOk() (*float64, bool) {
	if o == nil || IsNil(o.DayTradingBuyingPower) {
		return nil, false
	}
	return o.DayTradingBuyingPower, true
}

// HasDayTradingBuyingPower returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasDayTradingBuyingPower() bool {
	if o != nil && !IsNil(o.DayTradingBuyingPower) {
		return true
	}

	return false
}

// SetDayTradingBuyingPower gets a reference to the given float64 and assigns it to the DayTradingBuyingPower field.
func (o *MarginInitialBalance) SetDayTradingBuyingPower(v float64) {
	o.DayTradingBuyingPower = &v
}

// GetDayTradingBuyingPowerCall returns the DayTradingBuyingPowerCall field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetDayTradingBuyingPowerCall() float64 {
	if o == nil || IsNil(o.DayTradingBuyingPowerCall) {
		var ret float64
		return ret
	}
	return *o.DayTradingBuyingPowerCall
}

// GetDayTradingBuyingPowerCallOk returns a tuple with the DayTradingBuyingPowerCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetDayTradingBuyingPowerCallOk() (*float64, bool) {
	if o == nil || IsNil(o.DayTradingBuyingPowerCall) {
		return nil, false
	}
	return o.DayTradingBuyingPowerCall, true
}

// HasDayTradingBuyingPowerCall returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasDayTradingBuyingPowerCall() bool {
	if o != nil && !IsNil(o.DayTradingBuyingPowerCall) {
		return true
	}

	return false
}

// SetDayTradingBuyingPowerCall gets a reference to the given float64 and assigns it to the DayTradingBuyingPowerCall field.
func (o *MarginInitialBalance) SetDayTradingBuyingPowerCall(v float64) {
	o.DayTradingBuyingPowerCall = &v
}

// GetDayTradingEquityCall returns the DayTradingEquityCall field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetDayTradingEquityCall() float64 {
	if o == nil || IsNil(o.DayTradingEquityCall) {
		var ret float64
		return ret
	}
	return *o.DayTradingEquityCall
}

// GetDayTradingEquityCallOk returns a tuple with the DayTradingEquityCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetDayTradingEquityCallOk() (*float64, bool) {
	if o == nil || IsNil(o.DayTradingEquityCall) {
		return nil, false
	}
	return o.DayTradingEquityCall, true
}

// HasDayTradingEquityCall returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasDayTradingEquityCall() bool {
	if o != nil && !IsNil(o.DayTradingEquityCall) {
		return true
	}

	return false
}

// SetDayTradingEquityCall gets a reference to the given float64 and assigns it to the DayTradingEquityCall field.
func (o *MarginInitialBalance) SetDayTradingEquityCall(v float64) {
	o.DayTradingEquityCall = &v
}

// GetEquity returns the Equity field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetEquity() float64 {
	if o == nil || IsNil(o.Equity) {
		var ret float64
		return ret
	}
	return *o.Equity
}

// GetEquityOk returns a tuple with the Equity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.Equity) {
		return nil, false
	}
	return o.Equity, true
}

// HasEquity returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasEquity() bool {
	if o != nil && !IsNil(o.Equity) {
		return true
	}

	return false
}

// SetEquity gets a reference to the given float64 and assigns it to the Equity field.
func (o *MarginInitialBalance) SetEquity(v float64) {
	o.Equity = &v
}

// GetEquityPercentage returns the EquityPercentage field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetEquityPercentage() float64 {
	if o == nil || IsNil(o.EquityPercentage) {
		var ret float64
		return ret
	}
	return *o.EquityPercentage
}

// GetEquityPercentageOk returns a tuple with the EquityPercentage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetEquityPercentageOk() (*float64, bool) {
	if o == nil || IsNil(o.EquityPercentage) {
		return nil, false
	}
	return o.EquityPercentage, true
}

// HasEquityPercentage returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasEquityPercentage() bool {
	if o != nil && !IsNil(o.EquityPercentage) {
		return true
	}

	return false
}

// SetEquityPercentage gets a reference to the given float64 and assigns it to the EquityPercentage field.
func (o *MarginInitialBalance) SetEquityPercentage(v float64) {
	o.EquityPercentage = &v
}

// GetLiquidationValue returns the LiquidationValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetLiquidationValue() float64 {
	if o == nil || IsNil(o.LiquidationValue) {
		var ret float64
		return ret
	}
	return *o.LiquidationValue
}

// GetLiquidationValueOk returns a tuple with the LiquidationValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetLiquidationValueOk() (*float64, bool) {
	if o == nil || IsNil(o.LiquidationValue) {
		return nil, false
	}
	return o.LiquidationValue, true
}

// HasLiquidationValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasLiquidationValue() bool {
	if o != nil && !IsNil(o.LiquidationValue) {
		return true
	}

	return false
}

// SetLiquidationValue gets a reference to the given float64 and assigns it to the LiquidationValue field.
func (o *MarginInitialBalance) SetLiquidationValue(v float64) {
	o.LiquidationValue = &v
}

// GetLongMarginValue returns the LongMarginValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetLongMarginValue() float64 {
	if o == nil || IsNil(o.LongMarginValue) {
		var ret float64
		return ret
	}
	return *o.LongMarginValue
}

// GetLongMarginValueOk returns a tuple with the LongMarginValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetLongMarginValueOk() (*float64, bool) {
	if o == nil || IsNil(o.LongMarginValue) {
		return nil, false
	}
	return o.LongMarginValue, true
}

// HasLongMarginValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasLongMarginValue() bool {
	if o != nil && !IsNil(o.LongMarginValue) {
		return true
	}

	return false
}

// SetLongMarginValue gets a reference to the given float64 and assigns it to the LongMarginValue field.
func (o *MarginInitialBalance) SetLongMarginValue(v float64) {
	o.LongMarginValue = &v
}

// GetLongOptionMarketValue returns the LongOptionMarketValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetLongOptionMarketValue() float64 {
	if o == nil || IsNil(o.LongOptionMarketValue) {
		var ret float64
		return ret
	}
	return *o.LongOptionMarketValue
}

// GetLongOptionMarketValueOk returns a tuple with the LongOptionMarketValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetLongOptionMarketValueOk() (*float64, bool) {
	if o == nil || IsNil(o.LongOptionMarketValue) {
		return nil, false
	}
	return o.LongOptionMarketValue, true
}

// HasLongOptionMarketValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasLongOptionMarketValue() bool {
	if o != nil && !IsNil(o.LongOptionMarketValue) {
		return true
	}

	return false
}

// SetLongOptionMarketValue gets a reference to the given float64 and assigns it to the LongOptionMarketValue field.
func (o *MarginInitialBalance) SetLongOptionMarketValue(v float64) {
	o.LongOptionMarketValue = &v
}

// GetLongStockValue returns the LongStockValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetLongStockValue() float64 {
	if o == nil || IsNil(o.LongStockValue) {
		var ret float64
		return ret
	}
	return *o.LongStockValue
}

// GetLongStockValueOk returns a tuple with the LongStockValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetLongStockValueOk() (*float64, bool) {
	if o == nil || IsNil(o.LongStockValue) {
		return nil, false
	}
	return o.LongStockValue, true
}

// HasLongStockValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasLongStockValue() bool {
	if o != nil && !IsNil(o.LongStockValue) {
		return true
	}

	return false
}

// SetLongStockValue gets a reference to the given float64 and assigns it to the LongStockValue field.
func (o *MarginInitialBalance) SetLongStockValue(v float64) {
	o.LongStockValue = &v
}

// GetMaintenanceCall returns the MaintenanceCall field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetMaintenanceCall() float64 {
	if o == nil || IsNil(o.MaintenanceCall) {
		var ret float64
		return ret
	}
	return *o.MaintenanceCall
}

// GetMaintenanceCallOk returns a tuple with the MaintenanceCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetMaintenanceCallOk() (*float64, bool) {
	if o == nil || IsNil(o.MaintenanceCall) {
		return nil, false
	}
	return o.MaintenanceCall, true
}

// HasMaintenanceCall returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasMaintenanceCall() bool {
	if o != nil && !IsNil(o.MaintenanceCall) {
		return true
	}

	return false
}

// SetMaintenanceCall gets a reference to the given float64 and assigns it to the MaintenanceCall field.
func (o *MarginInitialBalance) SetMaintenanceCall(v float64) {
	o.MaintenanceCall = &v
}

// GetMaintenanceRequirement returns the MaintenanceRequirement field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetMaintenanceRequirement() float64 {
	if o == nil || IsNil(o.MaintenanceRequirement) {
		var ret float64
		return ret
	}
	return *o.MaintenanceRequirement
}

// GetMaintenanceRequirementOk returns a tuple with the MaintenanceRequirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetMaintenanceRequirementOk() (*float64, bool) {
	if o == nil || IsNil(o.MaintenanceRequirement) {
		return nil, false
	}
	return o.MaintenanceRequirement, true
}

// HasMaintenanceRequirement returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasMaintenanceRequirement() bool {
	if o != nil && !IsNil(o.MaintenanceRequirement) {
		return true
	}

	return false
}

// SetMaintenanceRequirement gets a reference to the given float64 and assigns it to the MaintenanceRequirement field.
func (o *MarginInitialBalance) SetMaintenanceRequirement(v float64) {
	o.MaintenanceRequirement = &v
}

// GetMargin returns the Margin field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetMargin() float64 {
	if o == nil || IsNil(o.Margin) {
		var ret float64
		return ret
	}
	return *o.Margin
}

// GetMarginOk returns a tuple with the Margin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetMarginOk() (*float64, bool) {
	if o == nil || IsNil(o.Margin) {
		return nil, false
	}
	return o.Margin, true
}

// HasMargin returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasMargin() bool {
	if o != nil && !IsNil(o.Margin) {
		return true
	}

	return false
}

// SetMargin gets a reference to the given float64 and assigns it to the Margin field.
func (o *MarginInitialBalance) SetMargin(v float64) {
	o.Margin = &v
}

// GetMarginEquity returns the MarginEquity field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetMarginEquity() float64 {
	if o == nil || IsNil(o.MarginEquity) {
		var ret float64
		return ret
	}
	return *o.MarginEquity
}

// GetMarginEquityOk returns a tuple with the MarginEquity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetMarginEquityOk() (*float64, bool) {
	if o == nil || IsNil(o.MarginEquity) {
		return nil, false
	}
	return o.MarginEquity, true
}

// HasMarginEquity returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasMarginEquity() bool {
	if o != nil && !IsNil(o.MarginEquity) {
		return true
	}

	return false
}

// SetMarginEquity gets a reference to the given float64 and assigns it to the MarginEquity field.
func (o *MarginInitialBalance) SetMarginEquity(v float64) {
	o.MarginEquity = &v
}

// GetMoneyMarketFund returns the MoneyMarketFund field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetMoneyMarketFund() float64 {
	if o == nil || IsNil(o.MoneyMarketFund) {
		var ret float64
		return ret
	}
	return *o.MoneyMarketFund
}

// GetMoneyMarketFundOk returns a tuple with the MoneyMarketFund field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetMoneyMarketFundOk() (*float64, bool) {
	if o == nil || IsNil(o.MoneyMarketFund) {
		return nil, false
	}
	return o.MoneyMarketFund, true
}

// HasMoneyMarketFund returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasMoneyMarketFund() bool {
	if o != nil && !IsNil(o.MoneyMarketFund) {
		return true
	}

	return false
}

// SetMoneyMarketFund gets a reference to the given float64 and assigns it to the MoneyMarketFund field.
func (o *MarginInitialBalance) SetMoneyMarketFund(v float64) {
	o.MoneyMarketFund = &v
}

// GetMutualFundValue returns the MutualFundValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetMutualFundValue() float64 {
	if o == nil || IsNil(o.MutualFundValue) {
		var ret float64
		return ret
	}
	return *o.MutualFundValue
}

// GetMutualFundValueOk returns a tuple with the MutualFundValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetMutualFundValueOk() (*float64, bool) {
	if o == nil || IsNil(o.MutualFundValue) {
		return nil, false
	}
	return o.MutualFundValue, true
}

// HasMutualFundValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasMutualFundValue() bool {
	if o != nil && !IsNil(o.MutualFundValue) {
		return true
	}

	return false
}

// SetMutualFundValue gets a reference to the given float64 and assigns it to the MutualFundValue field.
func (o *MarginInitialBalance) SetMutualFundValue(v float64) {
	o.MutualFundValue = &v
}

// GetRegTCall returns the RegTCall field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetRegTCall() float64 {
	if o == nil || IsNil(o.RegTCall) {
		var ret float64
		return ret
	}
	return *o.RegTCall
}

// GetRegTCallOk returns a tuple with the RegTCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetRegTCallOk() (*float64, bool) {
	if o == nil || IsNil(o.RegTCall) {
		return nil, false
	}
	return o.RegTCall, true
}

// HasRegTCall returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasRegTCall() bool {
	if o != nil && !IsNil(o.RegTCall) {
		return true
	}

	return false
}

// SetRegTCall gets a reference to the given float64 and assigns it to the RegTCall field.
func (o *MarginInitialBalance) SetRegTCall(v float64) {
	o.RegTCall = &v
}

// GetShortMarginValue returns the ShortMarginValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetShortMarginValue() float64 {
	if o == nil || IsNil(o.ShortMarginValue) {
		var ret float64
		return ret
	}
	return *o.ShortMarginValue
}

// GetShortMarginValueOk returns a tuple with the ShortMarginValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetShortMarginValueOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortMarginValue) {
		return nil, false
	}
	return o.ShortMarginValue, true
}

// HasShortMarginValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasShortMarginValue() bool {
	if o != nil && !IsNil(o.ShortMarginValue) {
		return true
	}

	return false
}

// SetShortMarginValue gets a reference to the given float64 and assigns it to the ShortMarginValue field.
func (o *MarginInitialBalance) SetShortMarginValue(v float64) {
	o.ShortMarginValue = &v
}

// GetShortOptionMarketValue returns the ShortOptionMarketValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetShortOptionMarketValue() float64 {
	if o == nil || IsNil(o.ShortOptionMarketValue) {
		var ret float64
		return ret
	}
	return *o.ShortOptionMarketValue
}

// GetShortOptionMarketValueOk returns a tuple with the ShortOptionMarketValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetShortOptionMarketValueOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortOptionMarketValue) {
		return nil, false
	}
	return o.ShortOptionMarketValue, true
}

// HasShortOptionMarketValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasShortOptionMarketValue() bool {
	if o != nil && !IsNil(o.ShortOptionMarketValue) {
		return true
	}

	return false
}

// SetShortOptionMarketValue gets a reference to the given float64 and assigns it to the ShortOptionMarketValue field.
func (o *MarginInitialBalance) SetShortOptionMarketValue(v float64) {
	o.ShortOptionMarketValue = &v
}

// GetShortStockValue returns the ShortStockValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetShortStockValue() float64 {
	if o == nil || IsNil(o.ShortStockValue) {
		var ret float64
		return ret
	}
	return *o.ShortStockValue
}

// GetShortStockValueOk returns a tuple with the ShortStockValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetShortStockValueOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortStockValue) {
		return nil, false
	}
	return o.ShortStockValue, true
}

// HasShortStockValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasShortStockValue() bool {
	if o != nil && !IsNil(o.ShortStockValue) {
		return true
	}

	return false
}

// SetShortStockValue gets a reference to the given float64 and assigns it to the ShortStockValue field.
func (o *MarginInitialBalance) SetShortStockValue(v float64) {
	o.ShortStockValue = &v
}

// GetTotalCash returns the TotalCash field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetTotalCash() float64 {
	if o == nil || IsNil(o.TotalCash) {
		var ret float64
		return ret
	}
	return *o.TotalCash
}

// GetTotalCashOk returns a tuple with the TotalCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetTotalCashOk() (*float64, bool) {
	if o == nil || IsNil(o.TotalCash) {
		return nil, false
	}
	return o.TotalCash, true
}

// HasTotalCash returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasTotalCash() bool {
	if o != nil && !IsNil(o.TotalCash) {
		return true
	}

	return false
}

// SetTotalCash gets a reference to the given float64 and assigns it to the TotalCash field.
func (o *MarginInitialBalance) SetTotalCash(v float64) {
	o.TotalCash = &v
}

// GetIsInCall returns the IsInCall field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetIsInCall() float64 {
	if o == nil || IsNil(o.IsInCall) {
		var ret float64
		return ret
	}
	return *o.IsInCall
}

// GetIsInCallOk returns a tuple with the IsInCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetIsInCallOk() (*float64, bool) {
	if o == nil || IsNil(o.IsInCall) {
		return nil, false
	}
	return o.IsInCall, true
}

// HasIsInCall returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasIsInCall() bool {
	if o != nil && !IsNil(o.IsInCall) {
		return true
	}

	return false
}

// SetIsInCall gets a reference to the given float64 and assigns it to the IsInCall field.
func (o *MarginInitialBalance) SetIsInCall(v float64) {
	o.IsInCall = &v
}

// GetUnsettledCash returns the UnsettledCash field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetUnsettledCash() float64 {
	if o == nil || IsNil(o.UnsettledCash) {
		var ret float64
		return ret
	}
	return *o.UnsettledCash
}

// GetUnsettledCashOk returns a tuple with the UnsettledCash field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetUnsettledCashOk() (*float64, bool) {
	if o == nil || IsNil(o.UnsettledCash) {
		return nil, false
	}
	return o.UnsettledCash, true
}

// HasUnsettledCash returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasUnsettledCash() bool {
	if o != nil && !IsNil(o.UnsettledCash) {
		return true
	}

	return false
}

// SetUnsettledCash gets a reference to the given float64 and assigns it to the UnsettledCash field.
func (o *MarginInitialBalance) SetUnsettledCash(v float64) {
	o.UnsettledCash = &v
}

// GetPendingDeposits returns the PendingDeposits field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetPendingDeposits() float64 {
	if o == nil || IsNil(o.PendingDeposits) {
		var ret float64
		return ret
	}
	return *o.PendingDeposits
}

// GetPendingDepositsOk returns a tuple with the PendingDeposits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetPendingDepositsOk() (*float64, bool) {
	if o == nil || IsNil(o.PendingDeposits) {
		return nil, false
	}
	return o.PendingDeposits, true
}

// HasPendingDeposits returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasPendingDeposits() bool {
	if o != nil && !IsNil(o.PendingDeposits) {
		return true
	}

	return false
}

// SetPendingDeposits gets a reference to the given float64 and assigns it to the PendingDeposits field.
func (o *MarginInitialBalance) SetPendingDeposits(v float64) {
	o.PendingDeposits = &v
}

// GetMarginBalance returns the MarginBalance field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetMarginBalance() float64 {
	if o == nil || IsNil(o.MarginBalance) {
		var ret float64
		return ret
	}
	return *o.MarginBalance
}

// GetMarginBalanceOk returns a tuple with the MarginBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetMarginBalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.MarginBalance) {
		return nil, false
	}
	return o.MarginBalance, true
}

// HasMarginBalance returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasMarginBalance() bool {
	if o != nil && !IsNil(o.MarginBalance) {
		return true
	}

	return false
}

// SetMarginBalance gets a reference to the given float64 and assigns it to the MarginBalance field.
func (o *MarginInitialBalance) SetMarginBalance(v float64) {
	o.MarginBalance = &v
}

// GetShortBalance returns the ShortBalance field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetShortBalance() float64 {
	if o == nil || IsNil(o.ShortBalance) {
		var ret float64
		return ret
	}
	return *o.ShortBalance
}

// GetShortBalanceOk returns a tuple with the ShortBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetShortBalanceOk() (*float64, bool) {
	if o == nil || IsNil(o.ShortBalance) {
		return nil, false
	}
	return o.ShortBalance, true
}

// HasShortBalance returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasShortBalance() bool {
	if o != nil && !IsNil(o.ShortBalance) {
		return true
	}

	return false
}

// SetShortBalance gets a reference to the given float64 and assigns it to the ShortBalance field.
func (o *MarginInitialBalance) SetShortBalance(v float64) {
	o.ShortBalance = &v
}

// GetAccountValue returns the AccountValue field value if set, zero value otherwise.
func (o *MarginInitialBalance) GetAccountValue() float64 {
	if o == nil || IsNil(o.AccountValue) {
		var ret float64
		return ret
	}
	return *o.AccountValue
}

// GetAccountValueOk returns a tuple with the AccountValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarginInitialBalance) GetAccountValueOk() (*float64, bool) {
	if o == nil || IsNil(o.AccountValue) {
		return nil, false
	}
	return o.AccountValue, true
}

// HasAccountValue returns a boolean if a field has been set.
func (o *MarginInitialBalance) HasAccountValue() bool {
	if o != nil && !IsNil(o.AccountValue) {
		return true
	}

	return false
}

// SetAccountValue gets a reference to the given float64 and assigns it to the AccountValue field.
func (o *MarginInitialBalance) SetAccountValue(v float64) {
	o.AccountValue = &v
}

func (o MarginInitialBalance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MarginInitialBalance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccruedInterest) {
		toSerialize["accruedInterest"] = o.AccruedInterest
	}
	if !IsNil(o.AvailableFundsNonMarginableTrade) {
		toSerialize["availableFundsNonMarginableTrade"] = o.AvailableFundsNonMarginableTrade
	}
	if !IsNil(o.BondValue) {
		toSerialize["bondValue"] = o.BondValue
	}
	if !IsNil(o.BuyingPower) {
		toSerialize["buyingPower"] = o.BuyingPower
	}
	if !IsNil(o.CashBalance) {
		toSerialize["cashBalance"] = o.CashBalance
	}
	if !IsNil(o.CashAvailableForTrading) {
		toSerialize["cashAvailableForTrading"] = o.CashAvailableForTrading
	}
	if !IsNil(o.CashReceipts) {
		toSerialize["cashReceipts"] = o.CashReceipts
	}
	if !IsNil(o.DayTradingBuyingPower) {
		toSerialize["dayTradingBuyingPower"] = o.DayTradingBuyingPower
	}
	if !IsNil(o.DayTradingBuyingPowerCall) {
		toSerialize["dayTradingBuyingPowerCall"] = o.DayTradingBuyingPowerCall
	}
	if !IsNil(o.DayTradingEquityCall) {
		toSerialize["dayTradingEquityCall"] = o.DayTradingEquityCall
	}
	if !IsNil(o.Equity) {
		toSerialize["equity"] = o.Equity
	}
	if !IsNil(o.EquityPercentage) {
		toSerialize["equityPercentage"] = o.EquityPercentage
	}
	if !IsNil(o.LiquidationValue) {
		toSerialize["liquidationValue"] = o.LiquidationValue
	}
	if !IsNil(o.LongMarginValue) {
		toSerialize["longMarginValue"] = o.LongMarginValue
	}
	if !IsNil(o.LongOptionMarketValue) {
		toSerialize["longOptionMarketValue"] = o.LongOptionMarketValue
	}
	if !IsNil(o.LongStockValue) {
		toSerialize["longStockValue"] = o.LongStockValue
	}
	if !IsNil(o.MaintenanceCall) {
		toSerialize["maintenanceCall"] = o.MaintenanceCall
	}
	if !IsNil(o.MaintenanceRequirement) {
		toSerialize["maintenanceRequirement"] = o.MaintenanceRequirement
	}
	if !IsNil(o.Margin) {
		toSerialize["margin"] = o.Margin
	}
	if !IsNil(o.MarginEquity) {
		toSerialize["marginEquity"] = o.MarginEquity
	}
	if !IsNil(o.MoneyMarketFund) {
		toSerialize["moneyMarketFund"] = o.MoneyMarketFund
	}
	if !IsNil(o.MutualFundValue) {
		toSerialize["mutualFundValue"] = o.MutualFundValue
	}
	if !IsNil(o.RegTCall) {
		toSerialize["regTCall"] = o.RegTCall
	}
	if !IsNil(o.ShortMarginValue) {
		toSerialize["shortMarginValue"] = o.ShortMarginValue
	}
	if !IsNil(o.ShortOptionMarketValue) {
		toSerialize["shortOptionMarketValue"] = o.ShortOptionMarketValue
	}
	if !IsNil(o.ShortStockValue) {
		toSerialize["shortStockValue"] = o.ShortStockValue
	}
	if !IsNil(o.TotalCash) {
		toSerialize["totalCash"] = o.TotalCash
	}
	if !IsNil(o.IsInCall) {
		toSerialize["isInCall"] = o.IsInCall
	}
	if !IsNil(o.UnsettledCash) {
		toSerialize["unsettledCash"] = o.UnsettledCash
	}
	if !IsNil(o.PendingDeposits) {
		toSerialize["pendingDeposits"] = o.PendingDeposits
	}
	if !IsNil(o.MarginBalance) {
		toSerialize["marginBalance"] = o.MarginBalance
	}
	if !IsNil(o.ShortBalance) {
		toSerialize["shortBalance"] = o.ShortBalance
	}
	if !IsNil(o.AccountValue) {
		toSerialize["accountValue"] = o.AccountValue
	}
	return toSerialize, nil
}

type NullableMarginInitialBalance struct {
	value *MarginInitialBalance
	isSet bool
}

func (v NullableMarginInitialBalance) Get() *MarginInitialBalance {
	return v.value
}

func (v *NullableMarginInitialBalance) Set(val *MarginInitialBalance) {
	v.value = val
	v.isSet = true
}

func (v NullableMarginInitialBalance) IsSet() bool {
	return v.isSet
}

func (v *NullableMarginInitialBalance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarginInitialBalance(val *MarginInitialBalance) *NullableMarginInitialBalance {
	return &NullableMarginInitialBalance{value: val, isSet: true}
}

func (v NullableMarginInitialBalance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarginInitialBalance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


