# PreviewOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | Pointer to **int64** |  | [optional] 
**OrderStrategy** | Pointer to [**OrderStrategy**](OrderStrategy.md) |  | [optional] 
**OrderValidationResult** | Pointer to [**OrderValidationResult**](OrderValidationResult.md) |  | [optional] 
**CommissionAndFee** | Pointer to [**CommissionAndFee**](CommissionAndFee.md) |  | [optional] 

## Methods

### NewPreviewOrder

`func NewPreviewOrder() *PreviewOrder`

NewPreviewOrder instantiates a new PreviewOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPreviewOrderWithDefaults

`func NewPreviewOrderWithDefaults() *PreviewOrder`

NewPreviewOrderWithDefaults instantiates a new PreviewOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *PreviewOrder) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *PreviewOrder) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *PreviewOrder) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *PreviewOrder) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetOrderStrategy

`func (o *PreviewOrder) GetOrderStrategy() OrderStrategy`

GetOrderStrategy returns the OrderStrategy field if non-nil, zero value otherwise.

### GetOrderStrategyOk

`func (o *PreviewOrder) GetOrderStrategyOk() (*OrderStrategy, bool)`

GetOrderStrategyOk returns a tuple with the OrderStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderStrategy

`func (o *PreviewOrder) SetOrderStrategy(v OrderStrategy)`

SetOrderStrategy sets OrderStrategy field to given value.

### HasOrderStrategy

`func (o *PreviewOrder) HasOrderStrategy() bool`

HasOrderStrategy returns a boolean if a field has been set.

### GetOrderValidationResult

`func (o *PreviewOrder) GetOrderValidationResult() OrderValidationResult`

GetOrderValidationResult returns the OrderValidationResult field if non-nil, zero value otherwise.

### GetOrderValidationResultOk

`func (o *PreviewOrder) GetOrderValidationResultOk() (*OrderValidationResult, bool)`

GetOrderValidationResultOk returns a tuple with the OrderValidationResult field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderValidationResult

`func (o *PreviewOrder) SetOrderValidationResult(v OrderValidationResult)`

SetOrderValidationResult sets OrderValidationResult field to given value.

### HasOrderValidationResult

`func (o *PreviewOrder) HasOrderValidationResult() bool`

HasOrderValidationResult returns a boolean if a field has been set.

### GetCommissionAndFee

`func (o *PreviewOrder) GetCommissionAndFee() CommissionAndFee`

GetCommissionAndFee returns the CommissionAndFee field if non-nil, zero value otherwise.

### GetCommissionAndFeeOk

`func (o *PreviewOrder) GetCommissionAndFeeOk() (*CommissionAndFee, bool)`

GetCommissionAndFeeOk returns a tuple with the CommissionAndFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionAndFee

`func (o *PreviewOrder) SetCommissionAndFee(v CommissionAndFee)`

SetCommissionAndFee sets CommissionAndFee field to given value.

### HasCommissionAndFee

`func (o *PreviewOrder) HasCommissionAndFee() bool`

HasCommissionAndFee returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


