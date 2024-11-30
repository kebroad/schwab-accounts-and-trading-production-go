# FeeValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to **float64** |  | [optional] 
**Type** | Pointer to [**FeeType**](FeeType.md) |  | [optional] 

## Methods

### NewFeeValue

`func NewFeeValue() *FeeValue`

NewFeeValue instantiates a new FeeValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeValueWithDefaults

`func NewFeeValueWithDefaults() *FeeValue`

NewFeeValueWithDefaults instantiates a new FeeValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *FeeValue) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FeeValue) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FeeValue) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *FeeValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *FeeValue) GetType() FeeType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *FeeValue) GetTypeOk() (*FeeType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *FeeValue) SetType(v FeeType)`

SetType sets Type field to given value.

### HasType

`func (o *FeeValue) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


