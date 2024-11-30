# CommissionAndFee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Commission** | Pointer to [**Commission**](Commission.md) |  | [optional] 
**Fee** | Pointer to [**Fees**](Fees.md) |  | [optional] 
**TrueCommission** | Pointer to [**Commission**](Commission.md) |  | [optional] 

## Methods

### NewCommissionAndFee

`func NewCommissionAndFee() *CommissionAndFee`

NewCommissionAndFee instantiates a new CommissionAndFee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommissionAndFeeWithDefaults

`func NewCommissionAndFeeWithDefaults() *CommissionAndFee`

NewCommissionAndFeeWithDefaults instantiates a new CommissionAndFee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommission

`func (o *CommissionAndFee) GetCommission() Commission`

GetCommission returns the Commission field if non-nil, zero value otherwise.

### GetCommissionOk

`func (o *CommissionAndFee) GetCommissionOk() (*Commission, bool)`

GetCommissionOk returns a tuple with the Commission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommission

`func (o *CommissionAndFee) SetCommission(v Commission)`

SetCommission sets Commission field to given value.

### HasCommission

`func (o *CommissionAndFee) HasCommission() bool`

HasCommission returns a boolean if a field has been set.

### GetFee

`func (o *CommissionAndFee) GetFee() Fees`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *CommissionAndFee) GetFeeOk() (*Fees, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *CommissionAndFee) SetFee(v Fees)`

SetFee sets Fee field to given value.

### HasFee

`func (o *CommissionAndFee) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetTrueCommission

`func (o *CommissionAndFee) GetTrueCommission() Commission`

GetTrueCommission returns the TrueCommission field if non-nil, zero value otherwise.

### GetTrueCommissionOk

`func (o *CommissionAndFee) GetTrueCommissionOk() (*Commission, bool)`

GetTrueCommissionOk returns a tuple with the TrueCommission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrueCommission

`func (o *CommissionAndFee) SetTrueCommission(v Commission)`

SetTrueCommission sets TrueCommission field to given value.

### HasTrueCommission

`func (o *CommissionAndFee) HasTrueCommission() bool`

HasTrueCommission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


