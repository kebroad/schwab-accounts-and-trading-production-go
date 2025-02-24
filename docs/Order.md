# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to [**Session**](Session.md) |  | [optional] 
**Duration** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**OrderType** | Pointer to [**OrderType**](OrderType.md) |  | [optional] 
**CancelTime** | Pointer to **time.Time** |  | [optional] 
**ComplexOrderStrategyType** | Pointer to [**ComplexOrderStrategyType**](ComplexOrderStrategyType.md) |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**FilledQuantity** | Pointer to **float64** |  | [optional] 
**RemainingQuantity** | Pointer to **float64** |  | [optional] 
**RequestedDestination** | Pointer to [**RequestedDestination**](RequestedDestination.md) |  | [optional] 
**DestinationLinkName** | Pointer to **string** |  | [optional] 
**ReleaseTime** | Pointer to **time.Time** |  | [optional] 
**StopPrice** | Pointer to **float64** |  | [optional] 
**StopPriceLinkBasis** | Pointer to [**StopPriceLinkBasis**](StopPriceLinkBasis.md) |  | [optional] 
**StopPriceLinkType** | Pointer to [**StopPriceLinkType**](StopPriceLinkType.md) |  | [optional] 
**StopPriceOffset** | Pointer to **float64** |  | [optional] 
**StopType** | Pointer to [**StopType**](StopType.md) |  | [optional] 
**PriceLinkBasis** | Pointer to [**PriceLinkBasis**](PriceLinkBasis.md) |  | [optional] 
**PriceLinkType** | Pointer to [**PriceLinkType**](PriceLinkType.md) |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**TaxLotMethod** | Pointer to [**TaxLotMethod**](TaxLotMethod.md) |  | [optional] 
**OrderLegCollection** | Pointer to [**[]OrderLegCollection**](OrderLegCollection.md) |  | [optional] 
**ActivationPrice** | Pointer to **float64** |  | [optional] 
**SpecialInstruction** | Pointer to [**SpecialInstruction**](SpecialInstruction.md) |  | [optional] 
**OrderStrategyType** | Pointer to [**OrderStrategyType**](OrderStrategyType.md) |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**Cancelable** | Pointer to **bool** |  | [optional] [default to false]
**Editable** | Pointer to **bool** |  | [optional] [default to false]
**Status** | Pointer to [**Status**](Status.md) |  | [optional] 
**EnteredTime** | Pointer to **time.Time** |  | [optional] 
**CloseTime** | Pointer to **time.Time** |  | [optional] 
**Tag** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **int64** |  | [optional] 
**OrderActivityCollection** | Pointer to [**[]OrderActivity**](OrderActivity.md) |  | [optional] 
**ReplacingOrderCollection** | Pointer to [**[]Order**](Order.md) |  | [optional] 
**ChildOrderStrategies** | Pointer to [**[]Order**](Order.md) |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewOrder

`func NewOrder() *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *Order) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *Order) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *Order) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *Order) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetDuration

`func (o *Order) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *Order) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *Order) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *Order) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetOrderType

`func (o *Order) GetOrderType() OrderType`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *Order) GetOrderTypeOk() (*OrderType, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *Order) SetOrderType(v OrderType)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *Order) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetCancelTime

`func (o *Order) GetCancelTime() time.Time`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *Order) GetCancelTimeOk() (*time.Time, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *Order) SetCancelTime(v time.Time)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *Order) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetComplexOrderStrategyType

`func (o *Order) GetComplexOrderStrategyType() ComplexOrderStrategyType`

GetComplexOrderStrategyType returns the ComplexOrderStrategyType field if non-nil, zero value otherwise.

### GetComplexOrderStrategyTypeOk

`func (o *Order) GetComplexOrderStrategyTypeOk() (*ComplexOrderStrategyType, bool)`

GetComplexOrderStrategyTypeOk returns a tuple with the ComplexOrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexOrderStrategyType

`func (o *Order) SetComplexOrderStrategyType(v ComplexOrderStrategyType)`

SetComplexOrderStrategyType sets ComplexOrderStrategyType field to given value.

### HasComplexOrderStrategyType

`func (o *Order) HasComplexOrderStrategyType() bool`

HasComplexOrderStrategyType returns a boolean if a field has been set.

### GetQuantity

`func (o *Order) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *Order) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *Order) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *Order) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *Order) GetFilledQuantity() float64`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *Order) GetFilledQuantityOk() (*float64, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *Order) SetFilledQuantity(v float64)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *Order) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *Order) GetRemainingQuantity() float64`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *Order) GetRemainingQuantityOk() (*float64, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *Order) SetRemainingQuantity(v float64)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *Order) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetRequestedDestination

`func (o *Order) GetRequestedDestination() RequestedDestination`

GetRequestedDestination returns the RequestedDestination field if non-nil, zero value otherwise.

### GetRequestedDestinationOk

`func (o *Order) GetRequestedDestinationOk() (*RequestedDestination, bool)`

GetRequestedDestinationOk returns a tuple with the RequestedDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDestination

`func (o *Order) SetRequestedDestination(v RequestedDestination)`

SetRequestedDestination sets RequestedDestination field to given value.

### HasRequestedDestination

`func (o *Order) HasRequestedDestination() bool`

HasRequestedDestination returns a boolean if a field has been set.

### GetDestinationLinkName

`func (o *Order) GetDestinationLinkName() string`

GetDestinationLinkName returns the DestinationLinkName field if non-nil, zero value otherwise.

### GetDestinationLinkNameOk

`func (o *Order) GetDestinationLinkNameOk() (*string, bool)`

GetDestinationLinkNameOk returns a tuple with the DestinationLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLinkName

`func (o *Order) SetDestinationLinkName(v string)`

SetDestinationLinkName sets DestinationLinkName field to given value.

### HasDestinationLinkName

`func (o *Order) HasDestinationLinkName() bool`

HasDestinationLinkName returns a boolean if a field has been set.

### GetReleaseTime

`func (o *Order) GetReleaseTime() time.Time`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *Order) GetReleaseTimeOk() (*time.Time, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *Order) SetReleaseTime(v time.Time)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *Order) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetStopPrice

`func (o *Order) GetStopPrice() float64`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *Order) GetStopPriceOk() (*float64, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *Order) SetStopPrice(v float64)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *Order) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStopPriceLinkBasis

`func (o *Order) GetStopPriceLinkBasis() StopPriceLinkBasis`

GetStopPriceLinkBasis returns the StopPriceLinkBasis field if non-nil, zero value otherwise.

### GetStopPriceLinkBasisOk

`func (o *Order) GetStopPriceLinkBasisOk() (*StopPriceLinkBasis, bool)`

GetStopPriceLinkBasisOk returns a tuple with the StopPriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkBasis

`func (o *Order) SetStopPriceLinkBasis(v StopPriceLinkBasis)`

SetStopPriceLinkBasis sets StopPriceLinkBasis field to given value.

### HasStopPriceLinkBasis

`func (o *Order) HasStopPriceLinkBasis() bool`

HasStopPriceLinkBasis returns a boolean if a field has been set.

### GetStopPriceLinkType

`func (o *Order) GetStopPriceLinkType() StopPriceLinkType`

GetStopPriceLinkType returns the StopPriceLinkType field if non-nil, zero value otherwise.

### GetStopPriceLinkTypeOk

`func (o *Order) GetStopPriceLinkTypeOk() (*StopPriceLinkType, bool)`

GetStopPriceLinkTypeOk returns a tuple with the StopPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkType

`func (o *Order) SetStopPriceLinkType(v StopPriceLinkType)`

SetStopPriceLinkType sets StopPriceLinkType field to given value.

### HasStopPriceLinkType

`func (o *Order) HasStopPriceLinkType() bool`

HasStopPriceLinkType returns a boolean if a field has been set.

### GetStopPriceOffset

`func (o *Order) GetStopPriceOffset() float64`

GetStopPriceOffset returns the StopPriceOffset field if non-nil, zero value otherwise.

### GetStopPriceOffsetOk

`func (o *Order) GetStopPriceOffsetOk() (*float64, bool)`

GetStopPriceOffsetOk returns a tuple with the StopPriceOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceOffset

`func (o *Order) SetStopPriceOffset(v float64)`

SetStopPriceOffset sets StopPriceOffset field to given value.

### HasStopPriceOffset

`func (o *Order) HasStopPriceOffset() bool`

HasStopPriceOffset returns a boolean if a field has been set.

### GetStopType

`func (o *Order) GetStopType() StopType`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *Order) GetStopTypeOk() (*StopType, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *Order) SetStopType(v StopType)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *Order) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetPriceLinkBasis

`func (o *Order) GetPriceLinkBasis() PriceLinkBasis`

GetPriceLinkBasis returns the PriceLinkBasis field if non-nil, zero value otherwise.

### GetPriceLinkBasisOk

`func (o *Order) GetPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetPriceLinkBasisOk returns a tuple with the PriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkBasis

`func (o *Order) SetPriceLinkBasis(v PriceLinkBasis)`

SetPriceLinkBasis sets PriceLinkBasis field to given value.

### HasPriceLinkBasis

`func (o *Order) HasPriceLinkBasis() bool`

HasPriceLinkBasis returns a boolean if a field has been set.

### GetPriceLinkType

`func (o *Order) GetPriceLinkType() PriceLinkType`

GetPriceLinkType returns the PriceLinkType field if non-nil, zero value otherwise.

### GetPriceLinkTypeOk

`func (o *Order) GetPriceLinkTypeOk() (*PriceLinkType, bool)`

GetPriceLinkTypeOk returns a tuple with the PriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkType

`func (o *Order) SetPriceLinkType(v PriceLinkType)`

SetPriceLinkType sets PriceLinkType field to given value.

### HasPriceLinkType

`func (o *Order) HasPriceLinkType() bool`

HasPriceLinkType returns a boolean if a field has been set.

### GetPrice

`func (o *Order) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Order) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Order) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Order) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxLotMethod

`func (o *Order) GetTaxLotMethod() TaxLotMethod`

GetTaxLotMethod returns the TaxLotMethod field if non-nil, zero value otherwise.

### GetTaxLotMethodOk

`func (o *Order) GetTaxLotMethodOk() (*TaxLotMethod, bool)`

GetTaxLotMethodOk returns a tuple with the TaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotMethod

`func (o *Order) SetTaxLotMethod(v TaxLotMethod)`

SetTaxLotMethod sets TaxLotMethod field to given value.

### HasTaxLotMethod

`func (o *Order) HasTaxLotMethod() bool`

HasTaxLotMethod returns a boolean if a field has been set.

### GetOrderLegCollection

`func (o *Order) GetOrderLegCollection() []OrderLegCollection`

GetOrderLegCollection returns the OrderLegCollection field if non-nil, zero value otherwise.

### GetOrderLegCollectionOk

`func (o *Order) GetOrderLegCollectionOk() (*[]OrderLegCollection, bool)`

GetOrderLegCollectionOk returns a tuple with the OrderLegCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegCollection

`func (o *Order) SetOrderLegCollection(v []OrderLegCollection)`

SetOrderLegCollection sets OrderLegCollection field to given value.

### HasOrderLegCollection

`func (o *Order) HasOrderLegCollection() bool`

HasOrderLegCollection returns a boolean if a field has been set.

### GetActivationPrice

`func (o *Order) GetActivationPrice() float64`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *Order) GetActivationPriceOk() (*float64, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *Order) SetActivationPrice(v float64)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *Order) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetSpecialInstruction

`func (o *Order) GetSpecialInstruction() SpecialInstruction`

GetSpecialInstruction returns the SpecialInstruction field if non-nil, zero value otherwise.

### GetSpecialInstructionOk

`func (o *Order) GetSpecialInstructionOk() (*SpecialInstruction, bool)`

GetSpecialInstructionOk returns a tuple with the SpecialInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstruction

`func (o *Order) SetSpecialInstruction(v SpecialInstruction)`

SetSpecialInstruction sets SpecialInstruction field to given value.

### HasSpecialInstruction

`func (o *Order) HasSpecialInstruction() bool`

HasSpecialInstruction returns a boolean if a field has been set.

### GetOrderStrategyType

`func (o *Order) GetOrderStrategyType() OrderStrategyType`

GetOrderStrategyType returns the OrderStrategyType field if non-nil, zero value otherwise.

### GetOrderStrategyTypeOk

`func (o *Order) GetOrderStrategyTypeOk() (*OrderStrategyType, bool)`

GetOrderStrategyTypeOk returns a tuple with the OrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategyType

`func (o *Order) SetOrderStrategyType(v OrderStrategyType)`

SetOrderStrategyType sets OrderStrategyType field to given value.

### HasOrderStrategyType

`func (o *Order) HasOrderStrategyType() bool`

HasOrderStrategyType returns a boolean if a field has been set.

### GetOrderId

`func (o *Order) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Order) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Order) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Order) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCancelable

`func (o *Order) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *Order) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *Order) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *Order) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetEditable

`func (o *Order) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *Order) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *Order) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *Order) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetStatus

`func (o *Order) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Order) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Order) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Order) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnteredTime

`func (o *Order) GetEnteredTime() time.Time`

GetEnteredTime returns the EnteredTime field if non-nil, zero value otherwise.

### GetEnteredTimeOk

`func (o *Order) GetEnteredTimeOk() (*time.Time, bool)`

GetEnteredTimeOk returns a tuple with the EnteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredTime

`func (o *Order) SetEnteredTime(v time.Time)`

SetEnteredTime sets EnteredTime field to given value.

### HasEnteredTime

`func (o *Order) HasEnteredTime() bool`

HasEnteredTime returns a boolean if a field has been set.

### GetCloseTime

`func (o *Order) GetCloseTime() time.Time`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *Order) GetCloseTimeOk() (*time.Time, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *Order) SetCloseTime(v time.Time)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *Order) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetTag

`func (o *Order) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *Order) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *Order) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *Order) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Order) GetAccountNumber() int64`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Order) GetAccountNumberOk() (*int64, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Order) SetAccountNumber(v int64)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Order) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetOrderActivityCollection

`func (o *Order) GetOrderActivityCollection() []OrderActivity`

GetOrderActivityCollection returns the OrderActivityCollection field if non-nil, zero value otherwise.

### GetOrderActivityCollectionOk

`func (o *Order) GetOrderActivityCollectionOk() (*[]OrderActivity, bool)`

GetOrderActivityCollectionOk returns a tuple with the OrderActivityCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderActivityCollection

`func (o *Order) SetOrderActivityCollection(v []OrderActivity)`

SetOrderActivityCollection sets OrderActivityCollection field to given value.

### HasOrderActivityCollection

`func (o *Order) HasOrderActivityCollection() bool`

HasOrderActivityCollection returns a boolean if a field has been set.

### GetReplacingOrderCollection

`func (o *Order) GetReplacingOrderCollection() []Order`

GetReplacingOrderCollection returns the ReplacingOrderCollection field if non-nil, zero value otherwise.

### GetReplacingOrderCollectionOk

`func (o *Order) GetReplacingOrderCollectionOk() (*[]Order, bool)`

GetReplacingOrderCollectionOk returns a tuple with the ReplacingOrderCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacingOrderCollection

`func (o *Order) SetReplacingOrderCollection(v []Order)`

SetReplacingOrderCollection sets ReplacingOrderCollection field to given value.

### HasReplacingOrderCollection

`func (o *Order) HasReplacingOrderCollection() bool`

HasReplacingOrderCollection returns a boolean if a field has been set.

### GetChildOrderStrategies

`func (o *Order) GetChildOrderStrategies() []Order`

GetChildOrderStrategies returns the ChildOrderStrategies field if non-nil, zero value otherwise.

### GetChildOrderStrategiesOk

`func (o *Order) GetChildOrderStrategiesOk() (*[]Order, bool)`

GetChildOrderStrategiesOk returns a tuple with the ChildOrderStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOrderStrategies

`func (o *Order) SetChildOrderStrategies(v []Order)`

SetChildOrderStrategies sets ChildOrderStrategies field to given value.

### HasChildOrderStrategies

`func (o *Order) HasChildOrderStrategies() bool`

HasChildOrderStrategies returns a boolean if a field has been set.

### GetStatusDescription

`func (o *Order) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *Order) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *Order) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *Order) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


