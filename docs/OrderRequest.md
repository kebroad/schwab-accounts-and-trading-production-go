# OrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Session** | Pointer to [**Session**](Session.md) |  | [optional] 
**Duration** | Pointer to [**Duration**](Duration.md) |  | [optional] 
**OrderType** | Pointer to [**OrderTypeRequest**](OrderTypeRequest.md) |  | [optional] 
**CancelTime** | Pointer to **time.Time** |  | [optional] 
**ComplexOrderStrategyType** | Pointer to [**ComplexOrderStrategyType**](ComplexOrderStrategyType.md) |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**FilledQuantity** | Pointer to **float64** |  | [optional] 
**RemainingQuantity** | Pointer to **float64** |  | [optional] 
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
**AccountNumber** | Pointer to **int64** |  | [optional] 
**OrderActivityCollection** | Pointer to [**[]OrderActivity**](OrderActivity.md) |  | [optional] 
**ReplacingOrderCollection** | Pointer to [**[]OrderRequest**](OrderRequest.md) |  | [optional] 
**ChildOrderStrategies** | Pointer to [**[]OrderRequest**](OrderRequest.md) |  | [optional] 
**StatusDescription** | Pointer to **string** |  | [optional] 

## Methods

### NewOrderRequest

`func NewOrderRequest() *OrderRequest`

NewOrderRequest instantiates a new OrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderRequestWithDefaults

`func NewOrderRequestWithDefaults() *OrderRequest`

NewOrderRequestWithDefaults instantiates a new OrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSession

`func (o *OrderRequest) GetSession() Session`

GetSession returns the Session field if non-nil, zero value otherwise.

### GetSessionOk

`func (o *OrderRequest) GetSessionOk() (*Session, bool)`

GetSessionOk returns a tuple with the Session field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSession

`func (o *OrderRequest) SetSession(v Session)`

SetSession sets Session field to given value.

### HasSession

`func (o *OrderRequest) HasSession() bool`

HasSession returns a boolean if a field has been set.

### GetDuration

`func (o *OrderRequest) GetDuration() Duration`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *OrderRequest) GetDurationOk() (*Duration, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *OrderRequest) SetDuration(v Duration)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *OrderRequest) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetOrderType

`func (o *OrderRequest) GetOrderType() OrderTypeRequest`

GetOrderType returns the OrderType field if non-nil, zero value otherwise.

### GetOrderTypeOk

`func (o *OrderRequest) GetOrderTypeOk() (*OrderTypeRequest, bool)`

GetOrderTypeOk returns a tuple with the OrderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderType

`func (o *OrderRequest) SetOrderType(v OrderTypeRequest)`

SetOrderType sets OrderType field to given value.

### HasOrderType

`func (o *OrderRequest) HasOrderType() bool`

HasOrderType returns a boolean if a field has been set.

### GetCancelTime

`func (o *OrderRequest) GetCancelTime() time.Time`

GetCancelTime returns the CancelTime field if non-nil, zero value otherwise.

### GetCancelTimeOk

`func (o *OrderRequest) GetCancelTimeOk() (*time.Time, bool)`

GetCancelTimeOk returns a tuple with the CancelTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelTime

`func (o *OrderRequest) SetCancelTime(v time.Time)`

SetCancelTime sets CancelTime field to given value.

### HasCancelTime

`func (o *OrderRequest) HasCancelTime() bool`

HasCancelTime returns a boolean if a field has been set.

### GetComplexOrderStrategyType

`func (o *OrderRequest) GetComplexOrderStrategyType() ComplexOrderStrategyType`

GetComplexOrderStrategyType returns the ComplexOrderStrategyType field if non-nil, zero value otherwise.

### GetComplexOrderStrategyTypeOk

`func (o *OrderRequest) GetComplexOrderStrategyTypeOk() (*ComplexOrderStrategyType, bool)`

GetComplexOrderStrategyTypeOk returns a tuple with the ComplexOrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexOrderStrategyType

`func (o *OrderRequest) SetComplexOrderStrategyType(v ComplexOrderStrategyType)`

SetComplexOrderStrategyType sets ComplexOrderStrategyType field to given value.

### HasComplexOrderStrategyType

`func (o *OrderRequest) HasComplexOrderStrategyType() bool`

HasComplexOrderStrategyType returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderRequest) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderRequest) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderRequest) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderRequest) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetFilledQuantity

`func (o *OrderRequest) GetFilledQuantity() float64`

GetFilledQuantity returns the FilledQuantity field if non-nil, zero value otherwise.

### GetFilledQuantityOk

`func (o *OrderRequest) GetFilledQuantityOk() (*float64, bool)`

GetFilledQuantityOk returns a tuple with the FilledQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilledQuantity

`func (o *OrderRequest) SetFilledQuantity(v float64)`

SetFilledQuantity sets FilledQuantity field to given value.

### HasFilledQuantity

`func (o *OrderRequest) HasFilledQuantity() bool`

HasFilledQuantity returns a boolean if a field has been set.

### GetRemainingQuantity

`func (o *OrderRequest) GetRemainingQuantity() float64`

GetRemainingQuantity returns the RemainingQuantity field if non-nil, zero value otherwise.

### GetRemainingQuantityOk

`func (o *OrderRequest) GetRemainingQuantityOk() (*float64, bool)`

GetRemainingQuantityOk returns a tuple with the RemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingQuantity

`func (o *OrderRequest) SetRemainingQuantity(v float64)`

SetRemainingQuantity sets RemainingQuantity field to given value.

### HasRemainingQuantity

`func (o *OrderRequest) HasRemainingQuantity() bool`

HasRemainingQuantity returns a boolean if a field has been set.

### GetDestinationLinkName

`func (o *OrderRequest) GetDestinationLinkName() string`

GetDestinationLinkName returns the DestinationLinkName field if non-nil, zero value otherwise.

### GetDestinationLinkNameOk

`func (o *OrderRequest) GetDestinationLinkNameOk() (*string, bool)`

GetDestinationLinkNameOk returns a tuple with the DestinationLinkName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationLinkName

`func (o *OrderRequest) SetDestinationLinkName(v string)`

SetDestinationLinkName sets DestinationLinkName field to given value.

### HasDestinationLinkName

`func (o *OrderRequest) HasDestinationLinkName() bool`

HasDestinationLinkName returns a boolean if a field has been set.

### GetReleaseTime

`func (o *OrderRequest) GetReleaseTime() time.Time`

GetReleaseTime returns the ReleaseTime field if non-nil, zero value otherwise.

### GetReleaseTimeOk

`func (o *OrderRequest) GetReleaseTimeOk() (*time.Time, bool)`

GetReleaseTimeOk returns a tuple with the ReleaseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseTime

`func (o *OrderRequest) SetReleaseTime(v time.Time)`

SetReleaseTime sets ReleaseTime field to given value.

### HasReleaseTime

`func (o *OrderRequest) HasReleaseTime() bool`

HasReleaseTime returns a boolean if a field has been set.

### GetStopPrice

`func (o *OrderRequest) GetStopPrice() float64`

GetStopPrice returns the StopPrice field if non-nil, zero value otherwise.

### GetStopPriceOk

`func (o *OrderRequest) GetStopPriceOk() (*float64, bool)`

GetStopPriceOk returns a tuple with the StopPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPrice

`func (o *OrderRequest) SetStopPrice(v float64)`

SetStopPrice sets StopPrice field to given value.

### HasStopPrice

`func (o *OrderRequest) HasStopPrice() bool`

HasStopPrice returns a boolean if a field has been set.

### GetStopPriceLinkBasis

`func (o *OrderRequest) GetStopPriceLinkBasis() StopPriceLinkBasis`

GetStopPriceLinkBasis returns the StopPriceLinkBasis field if non-nil, zero value otherwise.

### GetStopPriceLinkBasisOk

`func (o *OrderRequest) GetStopPriceLinkBasisOk() (*StopPriceLinkBasis, bool)`

GetStopPriceLinkBasisOk returns a tuple with the StopPriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkBasis

`func (o *OrderRequest) SetStopPriceLinkBasis(v StopPriceLinkBasis)`

SetStopPriceLinkBasis sets StopPriceLinkBasis field to given value.

### HasStopPriceLinkBasis

`func (o *OrderRequest) HasStopPriceLinkBasis() bool`

HasStopPriceLinkBasis returns a boolean if a field has been set.

### GetStopPriceLinkType

`func (o *OrderRequest) GetStopPriceLinkType() StopPriceLinkType`

GetStopPriceLinkType returns the StopPriceLinkType field if non-nil, zero value otherwise.

### GetStopPriceLinkTypeOk

`func (o *OrderRequest) GetStopPriceLinkTypeOk() (*StopPriceLinkType, bool)`

GetStopPriceLinkTypeOk returns a tuple with the StopPriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceLinkType

`func (o *OrderRequest) SetStopPriceLinkType(v StopPriceLinkType)`

SetStopPriceLinkType sets StopPriceLinkType field to given value.

### HasStopPriceLinkType

`func (o *OrderRequest) HasStopPriceLinkType() bool`

HasStopPriceLinkType returns a boolean if a field has been set.

### GetStopPriceOffset

`func (o *OrderRequest) GetStopPriceOffset() float64`

GetStopPriceOffset returns the StopPriceOffset field if non-nil, zero value otherwise.

### GetStopPriceOffsetOk

`func (o *OrderRequest) GetStopPriceOffsetOk() (*float64, bool)`

GetStopPriceOffsetOk returns a tuple with the StopPriceOffset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopPriceOffset

`func (o *OrderRequest) SetStopPriceOffset(v float64)`

SetStopPriceOffset sets StopPriceOffset field to given value.

### HasStopPriceOffset

`func (o *OrderRequest) HasStopPriceOffset() bool`

HasStopPriceOffset returns a boolean if a field has been set.

### GetStopType

`func (o *OrderRequest) GetStopType() StopType`

GetStopType returns the StopType field if non-nil, zero value otherwise.

### GetStopTypeOk

`func (o *OrderRequest) GetStopTypeOk() (*StopType, bool)`

GetStopTypeOk returns a tuple with the StopType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopType

`func (o *OrderRequest) SetStopType(v StopType)`

SetStopType sets StopType field to given value.

### HasStopType

`func (o *OrderRequest) HasStopType() bool`

HasStopType returns a boolean if a field has been set.

### GetPriceLinkBasis

`func (o *OrderRequest) GetPriceLinkBasis() PriceLinkBasis`

GetPriceLinkBasis returns the PriceLinkBasis field if non-nil, zero value otherwise.

### GetPriceLinkBasisOk

`func (o *OrderRequest) GetPriceLinkBasisOk() (*PriceLinkBasis, bool)`

GetPriceLinkBasisOk returns a tuple with the PriceLinkBasis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkBasis

`func (o *OrderRequest) SetPriceLinkBasis(v PriceLinkBasis)`

SetPriceLinkBasis sets PriceLinkBasis field to given value.

### HasPriceLinkBasis

`func (o *OrderRequest) HasPriceLinkBasis() bool`

HasPriceLinkBasis returns a boolean if a field has been set.

### GetPriceLinkType

`func (o *OrderRequest) GetPriceLinkType() PriceLinkType`

GetPriceLinkType returns the PriceLinkType field if non-nil, zero value otherwise.

### GetPriceLinkTypeOk

`func (o *OrderRequest) GetPriceLinkTypeOk() (*PriceLinkType, bool)`

GetPriceLinkTypeOk returns a tuple with the PriceLinkType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLinkType

`func (o *OrderRequest) SetPriceLinkType(v PriceLinkType)`

SetPriceLinkType sets PriceLinkType field to given value.

### HasPriceLinkType

`func (o *OrderRequest) HasPriceLinkType() bool`

HasPriceLinkType returns a boolean if a field has been set.

### GetPrice

`func (o *OrderRequest) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *OrderRequest) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *OrderRequest) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *OrderRequest) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetTaxLotMethod

`func (o *OrderRequest) GetTaxLotMethod() TaxLotMethod`

GetTaxLotMethod returns the TaxLotMethod field if non-nil, zero value otherwise.

### GetTaxLotMethodOk

`func (o *OrderRequest) GetTaxLotMethodOk() (*TaxLotMethod, bool)`

GetTaxLotMethodOk returns a tuple with the TaxLotMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxLotMethod

`func (o *OrderRequest) SetTaxLotMethod(v TaxLotMethod)`

SetTaxLotMethod sets TaxLotMethod field to given value.

### HasTaxLotMethod

`func (o *OrderRequest) HasTaxLotMethod() bool`

HasTaxLotMethod returns a boolean if a field has been set.

### GetOrderLegCollection

`func (o *OrderRequest) GetOrderLegCollection() []OrderLegCollection`

GetOrderLegCollection returns the OrderLegCollection field if non-nil, zero value otherwise.

### GetOrderLegCollectionOk

`func (o *OrderRequest) GetOrderLegCollectionOk() (*[]OrderLegCollection, bool)`

GetOrderLegCollectionOk returns a tuple with the OrderLegCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegCollection

`func (o *OrderRequest) SetOrderLegCollection(v []OrderLegCollection)`

SetOrderLegCollection sets OrderLegCollection field to given value.

### HasOrderLegCollection

`func (o *OrderRequest) HasOrderLegCollection() bool`

HasOrderLegCollection returns a boolean if a field has been set.

### GetActivationPrice

`func (o *OrderRequest) GetActivationPrice() float64`

GetActivationPrice returns the ActivationPrice field if non-nil, zero value otherwise.

### GetActivationPriceOk

`func (o *OrderRequest) GetActivationPriceOk() (*float64, bool)`

GetActivationPriceOk returns a tuple with the ActivationPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationPrice

`func (o *OrderRequest) SetActivationPrice(v float64)`

SetActivationPrice sets ActivationPrice field to given value.

### HasActivationPrice

`func (o *OrderRequest) HasActivationPrice() bool`

HasActivationPrice returns a boolean if a field has been set.

### GetSpecialInstruction

`func (o *OrderRequest) GetSpecialInstruction() SpecialInstruction`

GetSpecialInstruction returns the SpecialInstruction field if non-nil, zero value otherwise.

### GetSpecialInstructionOk

`func (o *OrderRequest) GetSpecialInstructionOk() (*SpecialInstruction, bool)`

GetSpecialInstructionOk returns a tuple with the SpecialInstruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecialInstruction

`func (o *OrderRequest) SetSpecialInstruction(v SpecialInstruction)`

SetSpecialInstruction sets SpecialInstruction field to given value.

### HasSpecialInstruction

`func (o *OrderRequest) HasSpecialInstruction() bool`

HasSpecialInstruction returns a boolean if a field has been set.

### GetOrderStrategyType

`func (o *OrderRequest) GetOrderStrategyType() OrderStrategyType`

GetOrderStrategyType returns the OrderStrategyType field if non-nil, zero value otherwise.

### GetOrderStrategyTypeOk

`func (o *OrderRequest) GetOrderStrategyTypeOk() (*OrderStrategyType, bool)`

GetOrderStrategyTypeOk returns a tuple with the OrderStrategyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategyType

`func (o *OrderRequest) SetOrderStrategyType(v OrderStrategyType)`

SetOrderStrategyType sets OrderStrategyType field to given value.

### HasOrderStrategyType

`func (o *OrderRequest) HasOrderStrategyType() bool`

HasOrderStrategyType returns a boolean if a field has been set.

### GetOrderId

`func (o *OrderRequest) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *OrderRequest) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *OrderRequest) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *OrderRequest) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetCancelable

`func (o *OrderRequest) GetCancelable() bool`

GetCancelable returns the Cancelable field if non-nil, zero value otherwise.

### GetCancelableOk

`func (o *OrderRequest) GetCancelableOk() (*bool, bool)`

GetCancelableOk returns a tuple with the Cancelable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelable

`func (o *OrderRequest) SetCancelable(v bool)`

SetCancelable sets Cancelable field to given value.

### HasCancelable

`func (o *OrderRequest) HasCancelable() bool`

HasCancelable returns a boolean if a field has been set.

### GetEditable

`func (o *OrderRequest) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *OrderRequest) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *OrderRequest) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *OrderRequest) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetStatus

`func (o *OrderRequest) GetStatus() Status`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OrderRequest) GetStatusOk() (*Status, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OrderRequest) SetStatus(v Status)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OrderRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetEnteredTime

`func (o *OrderRequest) GetEnteredTime() time.Time`

GetEnteredTime returns the EnteredTime field if non-nil, zero value otherwise.

### GetEnteredTimeOk

`func (o *OrderRequest) GetEnteredTimeOk() (*time.Time, bool)`

GetEnteredTimeOk returns a tuple with the EnteredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnteredTime

`func (o *OrderRequest) SetEnteredTime(v time.Time)`

SetEnteredTime sets EnteredTime field to given value.

### HasEnteredTime

`func (o *OrderRequest) HasEnteredTime() bool`

HasEnteredTime returns a boolean if a field has been set.

### GetCloseTime

`func (o *OrderRequest) GetCloseTime() time.Time`

GetCloseTime returns the CloseTime field if non-nil, zero value otherwise.

### GetCloseTimeOk

`func (o *OrderRequest) GetCloseTimeOk() (*time.Time, bool)`

GetCloseTimeOk returns a tuple with the CloseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloseTime

`func (o *OrderRequest) SetCloseTime(v time.Time)`

SetCloseTime sets CloseTime field to given value.

### HasCloseTime

`func (o *OrderRequest) HasCloseTime() bool`

HasCloseTime returns a boolean if a field has been set.

### GetAccountNumber

`func (o *OrderRequest) GetAccountNumber() int64`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *OrderRequest) GetAccountNumberOk() (*int64, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *OrderRequest) SetAccountNumber(v int64)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *OrderRequest) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetOrderActivityCollection

`func (o *OrderRequest) GetOrderActivityCollection() []OrderActivity`

GetOrderActivityCollection returns the OrderActivityCollection field if non-nil, zero value otherwise.

### GetOrderActivityCollectionOk

`func (o *OrderRequest) GetOrderActivityCollectionOk() (*[]OrderActivity, bool)`

GetOrderActivityCollectionOk returns a tuple with the OrderActivityCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderActivityCollection

`func (o *OrderRequest) SetOrderActivityCollection(v []OrderActivity)`

SetOrderActivityCollection sets OrderActivityCollection field to given value.

### HasOrderActivityCollection

`func (o *OrderRequest) HasOrderActivityCollection() bool`

HasOrderActivityCollection returns a boolean if a field has been set.

### GetReplacingOrderCollection

`func (o *OrderRequest) GetReplacingOrderCollection() []OrderRequest`

GetReplacingOrderCollection returns the ReplacingOrderCollection field if non-nil, zero value otherwise.

### GetReplacingOrderCollectionOk

`func (o *OrderRequest) GetReplacingOrderCollectionOk() (*[]OrderRequest, bool)`

GetReplacingOrderCollectionOk returns a tuple with the ReplacingOrderCollection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplacingOrderCollection

`func (o *OrderRequest) SetReplacingOrderCollection(v []OrderRequest)`

SetReplacingOrderCollection sets ReplacingOrderCollection field to given value.

### HasReplacingOrderCollection

`func (o *OrderRequest) HasReplacingOrderCollection() bool`

HasReplacingOrderCollection returns a boolean if a field has been set.

### GetChildOrderStrategies

`func (o *OrderRequest) GetChildOrderStrategies() []OrderRequest`

GetChildOrderStrategies returns the ChildOrderStrategies field if non-nil, zero value otherwise.

### GetChildOrderStrategiesOk

`func (o *OrderRequest) GetChildOrderStrategiesOk() (*[]OrderRequest, bool)`

GetChildOrderStrategiesOk returns a tuple with the ChildOrderStrategies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOrderStrategies

`func (o *OrderRequest) SetChildOrderStrategies(v []OrderRequest)`

SetChildOrderStrategies sets ChildOrderStrategies field to given value.

### HasChildOrderStrategies

`func (o *OrderRequest) HasChildOrderStrategies() bool`

HasChildOrderStrategies returns a boolean if a field has been set.

### GetStatusDescription

`func (o *OrderRequest) GetStatusDescription() string`

GetStatusDescription returns the StatusDescription field if non-nil, zero value otherwise.

### GetStatusDescriptionOk

`func (o *OrderRequest) GetStatusDescriptionOk() (*string, bool)`

GetStatusDescriptionOk returns a tuple with the StatusDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDescription

`func (o *OrderRequest) SetStatusDescription(v string)`

SetStatusDescription sets StatusDescription field to given value.

### HasStatusDescription

`func (o *OrderRequest) HasStatusDescription() bool`

HasStatusDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


