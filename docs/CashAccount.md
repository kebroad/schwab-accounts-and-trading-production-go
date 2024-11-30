# CashAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialBalances** | Pointer to [**CashInitialBalance**](CashInitialBalance.md) |  | [optional] 
**CurrentBalances** | Pointer to [**CashBalance**](CashBalance.md) |  | [optional] 
**ProjectedBalances** | Pointer to [**CashBalance**](CashBalance.md) |  | [optional] 

## Methods

### NewCashAccount

`func NewCashAccount() *CashAccount`

NewCashAccount instantiates a new CashAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCashAccountWithDefaults

`func NewCashAccountWithDefaults() *CashAccount`

NewCashAccountWithDefaults instantiates a new CashAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialBalances

`func (o *CashAccount) GetInitialBalances() CashInitialBalance`

GetInitialBalances returns the InitialBalances field if non-nil, zero value otherwise.

### GetInitialBalancesOk

`func (o *CashAccount) GetInitialBalancesOk() (*CashInitialBalance, bool)`

GetInitialBalancesOk returns a tuple with the InitialBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalances

`func (o *CashAccount) SetInitialBalances(v CashInitialBalance)`

SetInitialBalances sets InitialBalances field to given value.

### HasInitialBalances

`func (o *CashAccount) HasInitialBalances() bool`

HasInitialBalances returns a boolean if a field has been set.

### GetCurrentBalances

`func (o *CashAccount) GetCurrentBalances() CashBalance`

GetCurrentBalances returns the CurrentBalances field if non-nil, zero value otherwise.

### GetCurrentBalancesOk

`func (o *CashAccount) GetCurrentBalancesOk() (*CashBalance, bool)`

GetCurrentBalancesOk returns a tuple with the CurrentBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalances

`func (o *CashAccount) SetCurrentBalances(v CashBalance)`

SetCurrentBalances sets CurrentBalances field to given value.

### HasCurrentBalances

`func (o *CashAccount) HasCurrentBalances() bool`

HasCurrentBalances returns a boolean if a field has been set.

### GetProjectedBalances

`func (o *CashAccount) GetProjectedBalances() CashBalance`

GetProjectedBalances returns the ProjectedBalances field if non-nil, zero value otherwise.

### GetProjectedBalancesOk

`func (o *CashAccount) GetProjectedBalancesOk() (*CashBalance, bool)`

GetProjectedBalancesOk returns a tuple with the ProjectedBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedBalances

`func (o *CashAccount) SetProjectedBalances(v CashBalance)`

SetProjectedBalances sets ProjectedBalances field to given value.

### HasProjectedBalances

`func (o *CashAccount) HasProjectedBalances() bool`

HasProjectedBalances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


