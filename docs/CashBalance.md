# CashBalance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CashAvailableForTrading** | Pointer to **float64** |  | [optional] 
**CashAvailableForWithdrawal** | Pointer to **float64** |  | [optional] 
**CashCall** | Pointer to **float64** |  | [optional] 
**LongNonMarginableMarketValue** | Pointer to **float64** |  | [optional] 
**TotalCash** | Pointer to **float64** |  | [optional] 
**CashDebitCallValue** | Pointer to **float64** |  | [optional] 
**UnsettledCash** | Pointer to **float64** |  | [optional] 

## Methods

### NewCashBalance

`func NewCashBalance() *CashBalance`

NewCashBalance instantiates a new CashBalance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashBalanceWithDefaults

`func NewCashBalanceWithDefaults() *CashBalance`

NewCashBalanceWithDefaults instantiates a new CashBalance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCashAvailableForTrading

`func (o *CashBalance) GetCashAvailableForTrading() float64`

GetCashAvailableForTrading returns the CashAvailableForTrading field if non-nil, zero value otherwise.

### GetCashAvailableForTradingOk

`func (o *CashBalance) GetCashAvailableForTradingOk() (*float64, bool)`

GetCashAvailableForTradingOk returns a tuple with the CashAvailableForTrading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAvailableForTrading

`func (o *CashBalance) SetCashAvailableForTrading(v float64)`

SetCashAvailableForTrading sets CashAvailableForTrading field to given value.

### HasCashAvailableForTrading

`func (o *CashBalance) HasCashAvailableForTrading() bool`

HasCashAvailableForTrading returns a boolean if a field has been set.

### GetCashAvailableForWithdrawal

`func (o *CashBalance) GetCashAvailableForWithdrawal() float64`

GetCashAvailableForWithdrawal returns the CashAvailableForWithdrawal field if non-nil, zero value otherwise.

### GetCashAvailableForWithdrawalOk

`func (o *CashBalance) GetCashAvailableForWithdrawalOk() (*float64, bool)`

GetCashAvailableForWithdrawalOk returns a tuple with the CashAvailableForWithdrawal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashAvailableForWithdrawal

`func (o *CashBalance) SetCashAvailableForWithdrawal(v float64)`

SetCashAvailableForWithdrawal sets CashAvailableForWithdrawal field to given value.

### HasCashAvailableForWithdrawal

`func (o *CashBalance) HasCashAvailableForWithdrawal() bool`

HasCashAvailableForWithdrawal returns a boolean if a field has been set.

### GetCashCall

`func (o *CashBalance) GetCashCall() float64`

GetCashCall returns the CashCall field if non-nil, zero value otherwise.

### GetCashCallOk

`func (o *CashBalance) GetCashCallOk() (*float64, bool)`

GetCashCallOk returns a tuple with the CashCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashCall

`func (o *CashBalance) SetCashCall(v float64)`

SetCashCall sets CashCall field to given value.

### HasCashCall

`func (o *CashBalance) HasCashCall() bool`

HasCashCall returns a boolean if a field has been set.

### GetLongNonMarginableMarketValue

`func (o *CashBalance) GetLongNonMarginableMarketValue() float64`

GetLongNonMarginableMarketValue returns the LongNonMarginableMarketValue field if non-nil, zero value otherwise.

### GetLongNonMarginableMarketValueOk

`func (o *CashBalance) GetLongNonMarginableMarketValueOk() (*float64, bool)`

GetLongNonMarginableMarketValueOk returns a tuple with the LongNonMarginableMarketValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongNonMarginableMarketValue

`func (o *CashBalance) SetLongNonMarginableMarketValue(v float64)`

SetLongNonMarginableMarketValue sets LongNonMarginableMarketValue field to given value.

### HasLongNonMarginableMarketValue

`func (o *CashBalance) HasLongNonMarginableMarketValue() bool`

HasLongNonMarginableMarketValue returns a boolean if a field has been set.

### GetTotalCash

`func (o *CashBalance) GetTotalCash() float64`

GetTotalCash returns the TotalCash field if non-nil, zero value otherwise.

### GetTotalCashOk

`func (o *CashBalance) GetTotalCashOk() (*float64, bool)`

GetTotalCashOk returns a tuple with the TotalCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCash

`func (o *CashBalance) SetTotalCash(v float64)`

SetTotalCash sets TotalCash field to given value.

### HasTotalCash

`func (o *CashBalance) HasTotalCash() bool`

HasTotalCash returns a boolean if a field has been set.

### GetCashDebitCallValue

`func (o *CashBalance) GetCashDebitCallValue() float64`

GetCashDebitCallValue returns the CashDebitCallValue field if non-nil, zero value otherwise.

### GetCashDebitCallValueOk

`func (o *CashBalance) GetCashDebitCallValueOk() (*float64, bool)`

GetCashDebitCallValueOk returns a tuple with the CashDebitCallValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashDebitCallValue

`func (o *CashBalance) SetCashDebitCallValue(v float64)`

SetCashDebitCallValue sets CashDebitCallValue field to given value.

### HasCashDebitCallValue

`func (o *CashBalance) HasCashDebitCallValue() bool`

HasCashDebitCallValue returns a boolean if a field has been set.

### GetUnsettledCash

`func (o *CashBalance) GetUnsettledCash() float64`

GetUnsettledCash returns the UnsettledCash field if non-nil, zero value otherwise.

### GetUnsettledCashOk

`func (o *CashBalance) GetUnsettledCashOk() (*float64, bool)`

GetUnsettledCashOk returns a tuple with the UnsettledCash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnsettledCash

`func (o *CashBalance) SetUnsettledCash(v float64)`

SetUnsettledCash sets UnsettledCash field to given value.

### HasUnsettledCash

`func (o *CashBalance) HasUnsettledCash() bool`

HasUnsettledCash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


