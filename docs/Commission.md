# Commission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommissionLegs** | Pointer to [**[]CommissionLeg**](CommissionLeg.md) |  | [optional] 

## Methods

### NewCommission

`func NewCommission() *Commission`

NewCommission instantiates a new Commission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionWithDefaults

`func NewCommissionWithDefaults() *Commission`

NewCommissionWithDefaults instantiates a new Commission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommissionLegs

`func (o *Commission) GetCommissionLegs() []CommissionLeg`

GetCommissionLegs returns the CommissionLegs field if non-nil, zero value otherwise.

### GetCommissionLegsOk

`func (o *Commission) GetCommissionLegsOk() (*[]CommissionLeg, bool)`

GetCommissionLegsOk returns a tuple with the CommissionLegs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommissionLegs

`func (o *Commission) SetCommissionLegs(v []CommissionLeg)`

SetCommissionLegs sets CommissionLegs field to given value.

### HasCommissionLegs

`func (o *Commission) HasCommissionLegs() bool`

HasCommissionLegs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


