# SecuritiesAccountBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**RoundTrips** | Pointer to **int32** |  | [optional] 
**IsDayTrader** | Pointer to **bool** |  | [optional] [default to false]
**IsClosingOnlyRestricted** | Pointer to **bool** |  | [optional] [default to false]
**PfcbFlag** | Pointer to **bool** |  | [optional] [default to false]
**Positions** | Pointer to [**[]Position**](Position.md) |  | [optional] 

## Methods

### NewSecuritiesAccountBase

`func NewSecuritiesAccountBase() *SecuritiesAccountBase`

NewSecuritiesAccountBase instantiates a new SecuritiesAccountBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritiesAccountBaseWithDefaults

`func NewSecuritiesAccountBaseWithDefaults() *SecuritiesAccountBase`

NewSecuritiesAccountBaseWithDefaults instantiates a new SecuritiesAccountBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SecuritiesAccountBase) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecuritiesAccountBase) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecuritiesAccountBase) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecuritiesAccountBase) HasType() bool`

HasType returns a boolean if a field has been set.

### GetAccountNumber

`func (o *SecuritiesAccountBase) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *SecuritiesAccountBase) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *SecuritiesAccountBase) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *SecuritiesAccountBase) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetRoundTrips

`func (o *SecuritiesAccountBase) GetRoundTrips() int32`

GetRoundTrips returns the RoundTrips field if non-nil, zero value otherwise.

### GetRoundTripsOk

`func (o *SecuritiesAccountBase) GetRoundTripsOk() (*int32, bool)`

GetRoundTripsOk returns a tuple with the RoundTrips field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoundTrips

`func (o *SecuritiesAccountBase) SetRoundTrips(v int32)`

SetRoundTrips sets RoundTrips field to given value.

### HasRoundTrips

`func (o *SecuritiesAccountBase) HasRoundTrips() bool`

HasRoundTrips returns a boolean if a field has been set.

### GetIsDayTrader

`func (o *SecuritiesAccountBase) GetIsDayTrader() bool`

GetIsDayTrader returns the IsDayTrader field if non-nil, zero value otherwise.

### GetIsDayTraderOk

`func (o *SecuritiesAccountBase) GetIsDayTraderOk() (*bool, bool)`

GetIsDayTraderOk returns a tuple with the IsDayTrader field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDayTrader

`func (o *SecuritiesAccountBase) SetIsDayTrader(v bool)`

SetIsDayTrader sets IsDayTrader field to given value.

### HasIsDayTrader

`func (o *SecuritiesAccountBase) HasIsDayTrader() bool`

HasIsDayTrader returns a boolean if a field has been set.

### GetIsClosingOnlyRestricted

`func (o *SecuritiesAccountBase) GetIsClosingOnlyRestricted() bool`

GetIsClosingOnlyRestricted returns the IsClosingOnlyRestricted field if non-nil, zero value otherwise.

### GetIsClosingOnlyRestrictedOk

`func (o *SecuritiesAccountBase) GetIsClosingOnlyRestrictedOk() (*bool, bool)`

GetIsClosingOnlyRestrictedOk returns a tuple with the IsClosingOnlyRestricted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsClosingOnlyRestricted

`func (o *SecuritiesAccountBase) SetIsClosingOnlyRestricted(v bool)`

SetIsClosingOnlyRestricted sets IsClosingOnlyRestricted field to given value.

### HasIsClosingOnlyRestricted

`func (o *SecuritiesAccountBase) HasIsClosingOnlyRestricted() bool`

HasIsClosingOnlyRestricted returns a boolean if a field has been set.

### GetPfcbFlag

`func (o *SecuritiesAccountBase) GetPfcbFlag() bool`

GetPfcbFlag returns the PfcbFlag field if non-nil, zero value otherwise.

### GetPfcbFlagOk

`func (o *SecuritiesAccountBase) GetPfcbFlagOk() (*bool, bool)`

GetPfcbFlagOk returns a tuple with the PfcbFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPfcbFlag

`func (o *SecuritiesAccountBase) SetPfcbFlag(v bool)`

SetPfcbFlag sets PfcbFlag field to given value.

### HasPfcbFlag

`func (o *SecuritiesAccountBase) HasPfcbFlag() bool`

HasPfcbFlag returns a boolean if a field has been set.

### GetPositions

`func (o *SecuritiesAccountBase) GetPositions() []Position`

GetPositions returns the Positions field if non-nil, zero value otherwise.

### GetPositionsOk

`func (o *SecuritiesAccountBase) GetPositionsOk() (*[]Position, bool)`

GetPositionsOk returns a tuple with the Positions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositions

`func (o *SecuritiesAccountBase) SetPositions(v []Position)`

SetPositions sets Positions field to given value.

### HasPositions

`func (o *SecuritiesAccountBase) HasPositions() bool`

HasPositions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


