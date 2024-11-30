# MarginAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InitialBalances** | Pointer to [**MarginInitialBalance**](MarginInitialBalance.md) |  | [optional] 
**CurrentBalances** | Pointer to [**MarginBalance**](MarginBalance.md) |  | [optional] 
**ProjectedBalances** | Pointer to [**MarginBalance**](MarginBalance.md) |  | [optional] 

## Methods

### NewMarginAccount

`func NewMarginAccount() *MarginAccount`

NewMarginAccount instantiates a new MarginAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarginAccountWithDefaults

`func NewMarginAccountWithDefaults() *MarginAccount`

NewMarginAccountWithDefaults instantiates a new MarginAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInitialBalances

`func (o *MarginAccount) GetInitialBalances() MarginInitialBalance`

GetInitialBalances returns the InitialBalances field if non-nil, zero value otherwise.

### GetInitialBalancesOk

`func (o *MarginAccount) GetInitialBalancesOk() (*MarginInitialBalance, bool)`

GetInitialBalancesOk returns a tuple with the InitialBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialBalances

`func (o *MarginAccount) SetInitialBalances(v MarginInitialBalance)`

SetInitialBalances sets InitialBalances field to given value.

### HasInitialBalances

`func (o *MarginAccount) HasInitialBalances() bool`

HasInitialBalances returns a boolean if a field has been set.

### GetCurrentBalances

`func (o *MarginAccount) GetCurrentBalances() MarginBalance`

GetCurrentBalances returns the CurrentBalances field if non-nil, zero value otherwise.

### GetCurrentBalancesOk

`func (o *MarginAccount) GetCurrentBalancesOk() (*MarginBalance, bool)`

GetCurrentBalancesOk returns a tuple with the CurrentBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentBalances

`func (o *MarginAccount) SetCurrentBalances(v MarginBalance)`

SetCurrentBalances sets CurrentBalances field to given value.

### HasCurrentBalances

`func (o *MarginAccount) HasCurrentBalances() bool`

HasCurrentBalances returns a boolean if a field has been set.

### GetProjectedBalances

`func (o *MarginAccount) GetProjectedBalances() MarginBalance`

GetProjectedBalances returns the ProjectedBalances field if non-nil, zero value otherwise.

### GetProjectedBalancesOk

`func (o *MarginAccount) GetProjectedBalancesOk() (*MarginBalance, bool)`

GetProjectedBalancesOk returns a tuple with the ProjectedBalances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectedBalances

`func (o *MarginAccount) SetProjectedBalances(v MarginBalance)`

SetProjectedBalances sets ProjectedBalances field to given value.

### HasProjectedBalances

`func (o *MarginAccount) HasProjectedBalances() bool`

HasProjectedBalances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


