# OrderLegCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderLegType** | Pointer to **string** |  | [optional] 
**LegId** | Pointer to **int64** |  | [optional] 
**Instrument** | Pointer to **interface{}** |  | [optional] 
**Instruction** | Pointer to [**Instruction**](Instruction.md) |  | [optional] 
**PositionEffect** | Pointer to **string** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**QuantityType** | Pointer to **string** |  | [optional] 
**DivCapGains** | Pointer to **string** |  | [optional] 
**ToSymbol** | Pointer to **string** |  | [optional] 

## Methods

### NewOrderLegCollection

`func NewOrderLegCollection() *OrderLegCollection`

NewOrderLegCollection instantiates a new OrderLegCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLegCollectionWithDefaults

`func NewOrderLegCollectionWithDefaults() *OrderLegCollection`

NewOrderLegCollectionWithDefaults instantiates a new OrderLegCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderLegType

`func (o *OrderLegCollection) GetOrderLegType() string`

GetOrderLegType returns the OrderLegType field if non-nil, zero value otherwise.

### GetOrderLegTypeOk

`func (o *OrderLegCollection) GetOrderLegTypeOk() (*string, bool)`

GetOrderLegTypeOk returns a tuple with the OrderLegType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderLegType

`func (o *OrderLegCollection) SetOrderLegType(v string)`

SetOrderLegType sets OrderLegType field to given value.

### HasOrderLegType

`func (o *OrderLegCollection) HasOrderLegType() bool`

HasOrderLegType returns a boolean if a field has been set.

### GetLegId

`func (o *OrderLegCollection) GetLegId() int64`

GetLegId returns the LegId field if non-nil, zero value otherwise.

### GetLegIdOk

`func (o *OrderLegCollection) GetLegIdOk() (*int64, bool)`

GetLegIdOk returns a tuple with the LegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegId

`func (o *OrderLegCollection) SetLegId(v int64)`

SetLegId sets LegId field to given value.

### HasLegId

`func (o *OrderLegCollection) HasLegId() bool`

HasLegId returns a boolean if a field has been set.

### GetInstrument

`func (o *OrderLegCollection) GetInstrument() interface{}`

GetInstrument returns the Instrument field if non-nil, zero value otherwise.

### GetInstrumentOk

`func (o *OrderLegCollection) GetInstrumentOk() (*interface{}, bool)`

GetInstrumentOk returns a tuple with the Instrument field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstrument

`func (o *OrderLegCollection) SetInstrument(v interface{})`

SetInstrument sets Instrument field to given value.

### HasInstrument

`func (o *OrderLegCollection) HasInstrument() bool`

HasInstrument returns a boolean if a field has been set.

### SetInstrumentNil

`func (o *OrderLegCollection) SetInstrumentNil(b bool)`

 SetInstrumentNil sets the value for Instrument to be an explicit nil

### UnsetInstrument
`func (o *OrderLegCollection) UnsetInstrument()`

UnsetInstrument ensures that no value is present for Instrument, not even an explicit nil
### GetInstruction

`func (o *OrderLegCollection) GetInstruction() Instruction`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *OrderLegCollection) GetInstructionOk() (*Instruction, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *OrderLegCollection) SetInstruction(v Instruction)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *OrderLegCollection) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.

### GetPositionEffect

`func (o *OrderLegCollection) GetPositionEffect() string`

GetPositionEffect returns the PositionEffect field if non-nil, zero value otherwise.

### GetPositionEffectOk

`func (o *OrderLegCollection) GetPositionEffectOk() (*string, bool)`

GetPositionEffectOk returns a tuple with the PositionEffect field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionEffect

`func (o *OrderLegCollection) SetPositionEffect(v string)`

SetPositionEffect sets PositionEffect field to given value.

### HasPositionEffect

`func (o *OrderLegCollection) HasPositionEffect() bool`

HasPositionEffect returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderLegCollection) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderLegCollection) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderLegCollection) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderLegCollection) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetQuantityType

`func (o *OrderLegCollection) GetQuantityType() string`

GetQuantityType returns the QuantityType field if non-nil, zero value otherwise.

### GetQuantityTypeOk

`func (o *OrderLegCollection) GetQuantityTypeOk() (*string, bool)`

GetQuantityTypeOk returns a tuple with the QuantityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityType

`func (o *OrderLegCollection) SetQuantityType(v string)`

SetQuantityType sets QuantityType field to given value.

### HasQuantityType

`func (o *OrderLegCollection) HasQuantityType() bool`

HasQuantityType returns a boolean if a field has been set.

### GetDivCapGains

`func (o *OrderLegCollection) GetDivCapGains() string`

GetDivCapGains returns the DivCapGains field if non-nil, zero value otherwise.

### GetDivCapGainsOk

`func (o *OrderLegCollection) GetDivCapGainsOk() (*string, bool)`

GetDivCapGainsOk returns a tuple with the DivCapGains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivCapGains

`func (o *OrderLegCollection) SetDivCapGains(v string)`

SetDivCapGains sets DivCapGains field to given value.

### HasDivCapGains

`func (o *OrderLegCollection) HasDivCapGains() bool`

HasDivCapGains returns a boolean if a field has been set.

### GetToSymbol

`func (o *OrderLegCollection) GetToSymbol() string`

GetToSymbol returns the ToSymbol field if non-nil, zero value otherwise.

### GetToSymbolOk

`func (o *OrderLegCollection) GetToSymbolOk() (*string, bool)`

GetToSymbolOk returns a tuple with the ToSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToSymbol

`func (o *OrderLegCollection) SetToSymbol(v string)`

SetToSymbol sets ToSymbol field to given value.

### HasToSymbol

`func (o *OrderLegCollection) HasToSymbol() bool`

HasToSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


