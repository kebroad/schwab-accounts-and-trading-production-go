# OrderStrategy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountNumber** | Pointer to **string** |  | [optional] 
**AdvancedOrderType** | Pointer to **string** |  | [optional] 
**CloseTime** | Pointer to **time.Time** |  | [optional] 
**EnteredTime** | Pointer to **time.Time** |  | [optional] 
**OrderBalance** | Pointer to [**OrderBalance**](OrderBalance.md) |  | [optional] 
**OrderStrategyType** | Pointer to [**OrderStrategyType**](OrderStrategyType.md) |  | [optional] 
**OrderVersion** | Pointer to **float32** |  | [optional] 
**Session** | Pointer to [**Session**](Session.md) |  | [optional] 
**Status** | Pointer to [**ApiOrderStatus**](ApiOrderStatus.md) |  | [optional] 
**AllOrNone** | Pointer to **bool** |  | [optional] 
**Discretionary** | Pointer to **bool** |  | [optional] 
**Duration** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**FilledQuantity** | Pointer to **float64** |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**OrderValue** | Pointer to **float64** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**RemainingQuantity** | Pointer to **float64** |  | [optional] 
**SellNonMarginableFirst** | Pointer to **bool** |  | [optional] 
**SettlementInstruction** | Pointer to [**SettlementInstruction**](SettlementInstruction.md) |  | [optional] 
**Strategy** | Pointer to [**ComplexOrderStrategyType**](ComplexOrderStrategyType.md) |  | [optional] 
**AmountIndicator** | Pointer to [**AmountIndicator**](AmountIndicator.md) |  | [optional] 
**OrderLegs** | Pointer to [**[]OrderLeg**](OrderLeg.md) |  | [optional] 

## Methods

### NewOrderStrategy

`func NewOrderStrategy() *OrderStrategy`

NewOrderStrategy instantiates a new OrderStrategy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderStrategyWithDefaults

`func NewOrderStrategyWithDefaults() *OrderStrategy`

NewOrderStrategyWithDefaults instantiates a new OrderStrategy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountNumber

`func (o *OrderStrategy) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *OrderStrategy) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *OrderStrategy) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *OrderStrategy) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetAdvancedOrderType

`func (o *OrderStrategy) GetAdvancedOrderType() string`

GetAdvancedOrderType returns the AdvancedOrderType field if non-nil, zero value otherwise.

### GetAdvancedOrderTypeOk

`func (o *OrderStrategy) GetAdvancedOrderTypeOk() (*string, bool)`

GetAdvancedOrderTypeOk returns a tuple with the AdvancedOrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvancedOrderType

`func (o *OrderStrategy) SetAdvancedOrderType(v string)`

SetAdvancedOrderType sets AdvancedOrderType field to given value.

### HasAdvancedOrderType

`func (o *OrderStrategy) HasAdvancedOrderType() bool`

HasAdvancedOrderType returns a boolean if a field has been set.

### GetCloseTime

`func (o *OrderStrategy) GetCloseTime() time.Time`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *OrderStrategy) GetCloseTimeOk() (*time.Time, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *OrderStrategy) SetCloseTime(v time.Time)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *OrderStrategy) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetEnteredTime

`func (o *OrderStrategy) GetEnteredTime() time.Time`

GetEnteredTime returns the EnteredTime field if non-nil, zero value otherwise.

### GetEnteredTimeOk

`func (o *OrderStrategy) GetEnteredTimeOk() (*time.Time, bool)`

GetEnteredTimeOk returns a tuple with the EnteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredTime

`func (o *OrderStrategy) SetEnteredTime(v time.Time)`

SetEnteredTime sets EnteredTime field to given value.

### HasEnteredTime

`func (o *OrderStrategy) HasEnteredTime() bool`

HasEnteredTime returns a boolean if a field has been set.

### GetOrderBalance

`func (o *OrderStrategy) GetOrderBalance() OrderBalance`

GetOrderBalance returns the OrderBalance field if non-nil, zero value otherwise.

### GetOrderBalanceOk

`func (o *OrderStrategy) GetOrderBalanceOk() (*OrderBalance, bool)`

GetOrderBalanceOk returns a tuple with the OrderBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderBalance

`func (o *OrderStrategy) SetOrderBalance(v OrderBalance)`

SetOrderBalance sets OrderBalance field to given value.

### HasOrderBalance

`func (o *OrderStrategy) HasOrderBalance() bool`

HasOrderBalance returns a boolean if a field has been set.

### GetOrderStrategyType

`func (o *OrderStrategy) GetOrderStrategyType() OrderStrategyType`

GetOrderStrategyType returns the OrderStrategyType field if non-nil, zero value otherwise.

### GetOrderStrategyTypeOk

`func (o *OrderStrategy) GetOrderStrategyTypeOk() (*OrderStrategyType, bool)`

GetOrderStrategyTypeOk returns a tuple with the OrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategyType

`func (o *OrderStrategy) SetOrderStrategyType(v OrderStrategyType)`

SetOrderStrategyType sets OrderStrategyType field to given value.

### HasOrderStrategyType

`func (o *OrderStrategy) HasOrderStrategyType() bool`

HasOrderStrategyType returns a boolean if a field has been set.

### GetOrderVersion

`func (o *OrderStrategy) GetOrderVersion() float32`

GetOrderVersion returns the OrderVersion field if non-nil, zero value otherwise.

### GetOrderVersionOk

`func (o *OrderStrategy) GetOrderVersionOk() (*float32, bool)`

GetOrderVersionOk returns a tuple with the OrderVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderVersion

`func (o *OrderStrategy) SetOrderVersion(v float32)`

SetOrderVersion sets OrderVersion field to given value.

### HasOrderVersion

`func (o *OrderStrategy) HasOrderVersion() bool`

HasOrderVersion returns a boolean if a field has been set.

### GetSession

`func (o *OrderStrategy) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *OrderStrategy) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *OrderStrategy) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *OrderStrategy) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetStatus

`func (o *OrderStrategy) GetStatus() ApiOrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderStrategy) GetStatusOk() (*ApiOrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderStrategy) SetStatus(v ApiOrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderStrategy) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetAllOrNone

`func (o *OrderStrategy) GetAllOrNone() bool`

GetAllOrNone returns the AllOrNone field if non-nil, zero value otherwise.

### GetAllOrNoneOk

`func (o *OrderStrategy) GetAllOrNoneOk() (*bool, bool)`

GetAllOrNoneOk returns a tuple with the AllOrNone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllOrNone

`func (o *OrderStrategy) SetAllOrNone(v bool)`

SetAllOrNone sets AllOrNone field to given value.

### HasAllOrNone

`func (o *OrderStrategy) HasAllOrNone() bool`

HasAllOrNone returns a boolean if a field has been set.

### GetDiscretionary

`func (o *OrderStrategy) GetDiscretionary() bool`

GetDiscretionary returns the Discretionary field if non-nil, zero value otherwise.

### GetDiscretionaryOk

`func (o *OrderStrategy) GetDiscretionaryOk() (*bool, bool)`

GetDiscretionaryOk returns a tuple with the Discretionary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiscretionary

`func (o *OrderStrategy) SetDiscretionary(v bool)`

SetDiscretionary sets Discretionary field to given value.

### HasDiscretionary

`func (o *OrderStrategy) HasDiscretionary() bool`

HasDiscretionary returns a boolean if a field has been set.

### GetDuration

`func (o *OrderStrategy) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OrderStrategy) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OrderStrategy) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OrderStrategy) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *OrderStrategy) GetFilledQuantity() float64`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *OrderStrategy) GetFilledQuantityOk() (*float64, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *OrderStrategy) SetFilledQuantity(v float64)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *OrderStrategy) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderStrategy) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderStrategy) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderStrategy) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderStrategy) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetOrderValue

`func (o *OrderStrategy) GetOrderValue() float64`

GetOrderValue returns the OrderValue field if non-nil, zero value otherwise.

### GetOrderValueOk

`func (o *OrderStrategy) GetOrderValueOk() (*float64, bool)`

GetOrderValueOk returns a tuple with the OrderValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderValue

`func (o *OrderStrategy) SetOrderValue(v float64)`

SetOrderValue sets OrderValue field to given value.

### HasOrderValue

`func (o *OrderStrategy) HasOrderValue() bool`

HasOrderValue returns a boolean if a field has been set.

### GetPrice

`func (o *OrderStrategy) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderStrategy) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderStrategy) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderStrategy) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderStrategy) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderStrategy) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderStrategy) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderStrategy) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *OrderStrategy) GetRemainingQuantity() float64`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *OrderStrategy) GetRemainingQuantityOk() (*float64, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *OrderStrategy) SetRemainingQuantity(v float64)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *OrderStrategy) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetSellNonMarginableFirst

`func (o *OrderStrategy) GetSellNonMarginableFirst() bool`

GetSellNonMarginableFirst returns the SellNonMarginableFirst field if non-nil, zero value otherwise.

### GetSellNonMarginableFirstOk

`func (o *OrderStrategy) GetSellNonMarginableFirstOk() (*bool, bool)`

GetSellNonMarginableFirstOk returns a tuple with the SellNonMarginableFirst field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellNonMarginableFirst

`func (o *OrderStrategy) SetSellNonMarginableFirst(v bool)`

SetSellNonMarginableFirst sets SellNonMarginableFirst field to given value.

### HasSellNonMarginableFirst

`func (o *OrderStrategy) HasSellNonMarginableFirst() bool`

HasSellNonMarginableFirst returns a boolean if a field has been set.

### GetSettlementInstruction

`func (o *OrderStrategy) GetSettlementInstruction() SettlementInstruction`

GetSettlementInstruction returns the SettlementInstruction field if non-nil, zero value otherwise.

### GetSettlementInstructionOk

`func (o *OrderStrategy) GetSettlementInstructionOk() (*SettlementInstruction, bool)`

GetSettlementInstructionOk returns a tuple with the SettlementInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementInstruction

`func (o *OrderStrategy) SetSettlementInstruction(v SettlementInstruction)`

SetSettlementInstruction sets SettlementInstruction field to given value.

### HasSettlementInstruction

`func (o *OrderStrategy) HasSettlementInstruction() bool`

HasSettlementInstruction returns a boolean if a field has been set.

### GetStrategy

`func (o *OrderStrategy) GetStrategy() ComplexOrderStrategyType`

GetStrategy returns the Strategy field if non-nil, zero value otherwise.

### GetStrategyOk

`func (o *OrderStrategy) GetStrategyOk() (*ComplexOrderStrategyType, bool)`

GetStrategyOk returns a tuple with the Strategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrategy

`func (o *OrderStrategy) SetStrategy(v ComplexOrderStrategyType)`

SetStrategy sets Strategy field to given value.

### HasStrategy

`func (o *OrderStrategy) HasStrategy() bool`

HasStrategy returns a boolean if a field has been set.

### GetAmountIndicator

`func (o *OrderStrategy) GetAmountIndicator() AmountIndicator`

GetAmountIndicator returns the AmountIndicator field if non-nil, zero value otherwise.

### GetAmountIndicatorOk

`func (o *OrderStrategy) GetAmountIndicatorOk() (*AmountIndicator, bool)`

GetAmountIndicatorOk returns a tuple with the AmountIndicator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountIndicator

`func (o *OrderStrategy) SetAmountIndicator(v AmountIndicator)`

SetAmountIndicator sets AmountIndicator field to given value.

### HasAmountIndicator

`func (o *OrderStrategy) HasAmountIndicator() bool`

HasAmountIndicator returns a boolean if a field has been set.

### GetOrderLegs

`func (o *OrderStrategy) GetOrderLegs() []OrderLeg`

GetOrderLegs returns the OrderLegs field if non-nil, zero value otherwise.

### GetOrderLegsOk

`func (o *OrderStrategy) GetOrderLegsOk() (*[]OrderLeg, bool)`

GetOrderLegsOk returns a tuple with the OrderLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegs

`func (o *OrderStrategy) SetOrderLegs(v []OrderLeg)`

SetOrderLegs sets OrderLegs field to given value.

### HasOrderLegs

`func (o *OrderStrategy) HasOrderLegs() bool`

HasOrderLegs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


