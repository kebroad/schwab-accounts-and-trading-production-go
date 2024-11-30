# Fees

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeeLegs** | Pointer to [**[]FeeLeg**](FeeLeg.md) |  | [optional] 

## Methods

### NewFees

`func NewFees() *Fees`

NewFees instantiates a new Fees object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFeesWithDefaults

`func NewFeesWithDefaults() *Fees`

NewFeesWithDefaults instantiates a new Fees object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeeLegs

`func (o *Fees) GetFeeLegs() []FeeLeg`

GetFeeLegs returns the FeeLegs field if non-nil, zero value otherwise.

### GetFeeLegsOk

`func (o *Fees) GetFeeLegsOk() (*[]FeeLeg, bool)`

GetFeeLegsOk returns a tuple with the FeeLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeLegs

`func (o *Fees) SetFeeLegs(v []FeeLeg)`

SetFeeLegs sets FeeLegs field to given value.

### HasFeeLegs

`func (o *Fees) HasFeeLegs() bool`

HasFeeLegs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


