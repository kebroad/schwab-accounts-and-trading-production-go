# FeeLeg

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeValues** | Pointer to [**[]FeeValue**](FeeValue.md) |  | [optional] 

## Methods

### NewFeeLeg

`func NewFeeLeg() *FeeLeg`

NewFeeLeg instantiates a new FeeLeg object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeeLegWithDefaults

`func NewFeeLegWithDefaults() *FeeLeg`

NewFeeLegWithDefaults instantiates a new FeeLeg object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeValues

`func (o *FeeLeg) GetFeeValues() []FeeValue`

GetFeeValues returns the FeeValues field if non-nil, zero value otherwise.

### GetFeeValuesOk

`func (o *FeeLeg) GetFeeValuesOk() (*[]FeeValue, bool)`

GetFeeValuesOk returns a tuple with the FeeValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeValues

`func (o *FeeLeg) SetFeeValues(v []FeeValue)`

SetFeeValues sets FeeValues field to given value.

### HasFeeValues

`func (o *FeeLeg) HasFeeValues() bool`

HasFeeValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


