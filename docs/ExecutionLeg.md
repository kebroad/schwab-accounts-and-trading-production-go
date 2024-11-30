# ExecutionLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LegId** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**MismarkedQuantity** | Pointer to **float64** |  | [optional] 
**InstrumentId** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewExecutionLeg

`func NewExecutionLeg() *ExecutionLeg`

NewExecutionLeg instantiates a new ExecutionLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecutionLegWithDefaults

`func NewExecutionLegWithDefaults() *ExecutionLeg`

NewExecutionLegWithDefaults instantiates a new ExecutionLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegId

`func (o *ExecutionLeg) GetLegId() int64`

GetLegId returns the LegId field if non-nil, zero value otherwise.

### GetLegIdOk

`func (o *ExecutionLeg) GetLegIdOk() (*int64, bool)`

GetLegIdOk returns a tuple with the LegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegId

`func (o *ExecutionLeg) SetLegId(v int64)`

SetLegId sets LegId field to given value.

### HasLegId

`func (o *ExecutionLeg) HasLegId() bool`

HasLegId returns a boolean if a field has been set.

### GetPrice

`func (o *ExecutionLeg) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ExecutionLeg) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ExecutionLeg) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ExecutionLeg) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetQuantity

`func (o *ExecutionLeg) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ExecutionLeg) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ExecutionLeg) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *ExecutionLeg) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetMismarkedQuantity

`func (o *ExecutionLeg) GetMismarkedQuantity() float64`

GetMismarkedQuantity returns the MismarkedQuantity field if non-nil, zero value otherwise.

### GetMismarkedQuantityOk

`func (o *ExecutionLeg) GetMismarkedQuantityOk() (*float64, bool)`

GetMismarkedQuantityOk returns a tuple with the MismarkedQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMismarkedQuantity

`func (o *ExecutionLeg) SetMismarkedQuantity(v float64)`

SetMismarkedQuantity sets MismarkedQuantity field to given value.

### HasMismarkedQuantity

`func (o *ExecutionLeg) HasMismarkedQuantity() bool`

HasMismarkedQuantity returns a boolean if a field has been set.

### GetInstrumentId

`func (o *ExecutionLeg) GetInstrumentId() int64`

GetInstrumentId returns the InstrumentId field if non-nil, zero value otherwise.

### GetInstrumentIdOk

`func (o *ExecutionLeg) GetInstrumentIdOk() (*int64, bool)`

GetInstrumentIdOk returns a tuple with the InstrumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrumentId

`func (o *ExecutionLeg) SetInstrumentId(v int64)`

SetInstrumentId sets InstrumentId field to given value.

### HasInstrumentId

`func (o *ExecutionLeg) HasInstrumentId() bool`

HasInstrumentId returns a boolean if a field has been set.

### GetTime

`func (o *ExecutionLeg) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *ExecutionLeg) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *ExecutionLeg) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *ExecutionLeg) HasTime() bool`

HasTime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


