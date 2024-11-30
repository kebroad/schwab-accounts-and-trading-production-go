# OrderActivity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityType** | Pointer to **string** |  | [optional] 
**ExecutionType** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**OrderRemainingQuantity** | Pointer to **float64** |  | [optional] 
**ExecutionLegs** | Pointer to [**[]ExecutionLeg**](ExecutionLeg.md) |  | [optional] 

## Methods

### NewOrderActivity

`func NewOrderActivity() *OrderActivity`

NewOrderActivity instantiates a new OrderActivity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderActivityWithDefaults

`func NewOrderActivityWithDefaults() *OrderActivity`

NewOrderActivityWithDefaults instantiates a new OrderActivity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityType

`func (o *OrderActivity) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *OrderActivity) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *OrderActivity) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *OrderActivity) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetExecutionType

`func (o *OrderActivity) GetExecutionType() string`

GetExecutionType returns the ExecutionType field if non-nil, zero value otherwise.

### GetExecutionTypeOk

`func (o *OrderActivity) GetExecutionTypeOk() (*string, bool)`

GetExecutionTypeOk returns a tuple with the ExecutionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionType

`func (o *OrderActivity) SetExecutionType(v string)`

SetExecutionType sets ExecutionType field to given value.

### HasExecutionType

`func (o *OrderActivity) HasExecutionType() bool`

HasExecutionType returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderActivity) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderActivity) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderActivity) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderActivity) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetOrderRemainingQuantity

`func (o *OrderActivity) GetOrderRemainingQuantity() float64`

GetOrderRemainingQuantity returns the OrderRemainingQuantity field if non-nil, zero value otherwise.

### GetOrderRemainingQuantityOk

`func (o *OrderActivity) GetOrderRemainingQuantityOk() (*float64, bool)`

GetOrderRemainingQuantityOk returns a tuple with the OrderRemainingQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderRemainingQuantity

`func (o *OrderActivity) SetOrderRemainingQuantity(v float64)`

SetOrderRemainingQuantity sets OrderRemainingQuantity field to given value.

### HasOrderRemainingQuantity

`func (o *OrderActivity) HasOrderRemainingQuantity() bool`

HasOrderRemainingQuantity returns a boolean if a field has been set.

### GetExecutionLegs

`func (o *OrderActivity) GetExecutionLegs() []ExecutionLeg`

GetExecutionLegs returns the ExecutionLegs field if non-nil, zero value otherwise.

### GetExecutionLegsOk

`func (o *OrderActivity) GetExecutionLegsOk() (*[]ExecutionLeg, bool)`

GetExecutionLegsOk returns a tuple with the ExecutionLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionLegs

`func (o *OrderActivity) SetExecutionLegs(v []ExecutionLeg)`

SetExecutionLegs sets ExecutionLegs field to given value.

### HasExecutionLegs

`func (o *OrderActivity) HasExecutionLegs() bool`

HasExecutionLegs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


