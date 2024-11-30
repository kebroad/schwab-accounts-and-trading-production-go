# OrderLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AskPrice** | Pointer to **float64** |  | [optional] 
**BidPrice** | Pointer to **float64** |  | [optional] 
**LastPrice** | Pointer to **float64** |  | [optional] 
**MarkPrice** | Pointer to **float64** |  | [optional] 
**ProjectedCommission** | Pointer to **float64** |  | [optional] 
**Quantity** | Pointer to **float64** |  | [optional] 
**FinalSymbol** | Pointer to **string** |  | [optional] 
**LegId** | Pointer to **float32** |  | [optional] 
**AssetType** | Pointer to [**AssetType**](AssetType.md) |  | [optional] 
**Instruction** | Pointer to [**Instruction**](Instruction.md) |  | [optional] 

## Methods

### NewOrderLeg

`func NewOrderLeg() *OrderLeg`

NewOrderLeg instantiates a new OrderLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderLegWithDefaults

`func NewOrderLegWithDefaults() *OrderLeg`

NewOrderLegWithDefaults instantiates a new OrderLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAskPrice

`func (o *OrderLeg) GetAskPrice() float64`

GetAskPrice returns the AskPrice field if non-nil, zero value otherwise.

### GetAskPriceOk

`func (o *OrderLeg) GetAskPriceOk() (*float64, bool)`

GetAskPriceOk returns a tuple with the AskPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAskPrice

`func (o *OrderLeg) SetAskPrice(v float64)`

SetAskPrice sets AskPrice field to given value.

### HasAskPrice

`func (o *OrderLeg) HasAskPrice() bool`

HasAskPrice returns a boolean if a field has been set.

### GetBidPrice

`func (o *OrderLeg) GetBidPrice() float64`

GetBidPrice returns the BidPrice field if non-nil, zero value otherwise.

### GetBidPriceOk

`func (o *OrderLeg) GetBidPriceOk() (*float64, bool)`

GetBidPriceOk returns a tuple with the BidPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBidPrice

`func (o *OrderLeg) SetBidPrice(v float64)`

SetBidPrice sets BidPrice field to given value.

### HasBidPrice

`func (o *OrderLeg) HasBidPrice() bool`

HasBidPrice returns a boolean if a field has been set.

### GetLastPrice

`func (o *OrderLeg) GetLastPrice() float64`

GetLastPrice returns the LastPrice field if non-nil, zero value otherwise.

### GetLastPriceOk

`func (o *OrderLeg) GetLastPriceOk() (*float64, bool)`

GetLastPriceOk returns a tuple with the LastPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastPrice

`func (o *OrderLeg) SetLastPrice(v float64)`

SetLastPrice sets LastPrice field to given value.

### HasLastPrice

`func (o *OrderLeg) HasLastPrice() bool`

HasLastPrice returns a boolean if a field has been set.

### GetMarkPrice

`func (o *OrderLeg) GetMarkPrice() float64`

GetMarkPrice returns the MarkPrice field if non-nil, zero value otherwise.

### GetMarkPriceOk

`func (o *OrderLeg) GetMarkPriceOk() (*float64, bool)`

GetMarkPriceOk returns a tuple with the MarkPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkPrice

`func (o *OrderLeg) SetMarkPrice(v float64)`

SetMarkPrice sets MarkPrice field to given value.

### HasMarkPrice

`func (o *OrderLeg) HasMarkPrice() bool`

HasMarkPrice returns a boolean if a field has been set.

### GetProjectedCommission

`func (o *OrderLeg) GetProjectedCommission() float64`

GetProjectedCommission returns the ProjectedCommission field if non-nil, zero value otherwise.

### GetProjectedCommissionOk

`func (o *OrderLeg) GetProjectedCommissionOk() (*float64, bool)`

GetProjectedCommissionOk returns a tuple with the ProjectedCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedCommission

`func (o *OrderLeg) SetProjectedCommission(v float64)`

SetProjectedCommission sets ProjectedCommission field to given value.

### HasProjectedCommission

`func (o *OrderLeg) HasProjectedCommission() bool`

HasProjectedCommission returns a boolean if a field has been set.

### GetQuantity

`func (o *OrderLeg) GetQuantity() float64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *OrderLeg) GetQuantityOk() (*float64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *OrderLeg) SetQuantity(v float64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *OrderLeg) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetFinalSymbol

`func (o *OrderLeg) GetFinalSymbol() string`

GetFinalSymbol returns the FinalSymbol field if non-nil, zero value otherwise.

### GetFinalSymbolOk

`func (o *OrderLeg) GetFinalSymbolOk() (*string, bool)`

GetFinalSymbolOk returns a tuple with the FinalSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinalSymbol

`func (o *OrderLeg) SetFinalSymbol(v string)`

SetFinalSymbol sets FinalSymbol field to given value.

### HasFinalSymbol

`func (o *OrderLeg) HasFinalSymbol() bool`

HasFinalSymbol returns a boolean if a field has been set.

### GetLegId

`func (o *OrderLeg) GetLegId() float32`

GetLegId returns the LegId field if non-nil, zero value otherwise.

### GetLegIdOk

`func (o *OrderLeg) GetLegIdOk() (*float32, bool)`

GetLegIdOk returns a tuple with the LegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegId

`func (o *OrderLeg) SetLegId(v float32)`

SetLegId sets LegId field to given value.

### HasLegId

`func (o *OrderLeg) HasLegId() bool`

HasLegId returns a boolean if a field has been set.

### GetAssetType

`func (o *OrderLeg) GetAssetType() AssetType`

GetAssetType returns the AssetType field if non-nil, zero value otherwise.

### GetAssetTypeOk

`func (o *OrderLeg) GetAssetTypeOk() (*AssetType, bool)`

GetAssetTypeOk returns a tuple with the AssetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetType

`func (o *OrderLeg) SetAssetType(v AssetType)`

SetAssetType sets AssetType field to given value.

### HasAssetType

`func (o *OrderLeg) HasAssetType() bool`

HasAssetType returns a boolean if a field has been set.

### GetInstruction

`func (o *OrderLeg) GetInstruction() Instruction`

GetInstruction returns the Instruction field if non-nil, zero value otherwise.

### GetInstructionOk

`func (o *OrderLeg) GetInstructionOk() (*Instruction, bool)`

GetInstructionOk returns a tuple with the Instruction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstruction

`func (o *OrderLeg) SetInstruction(v Instruction)`

SetInstruction sets Instruction field to given value.

### HasInstruction

`func (o *OrderLeg) HasInstruction() bool`

HasInstruction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


