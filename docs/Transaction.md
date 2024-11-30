# Transaction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityId** | Pointer to **int64** |  | [optional] 
**Time** | Pointer to **time.Time** |  | [optional] 
**User** | Pointer to [**UserDetails**](UserDetails.md) |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**AccountNumber** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**TransactionType**](TransactionType.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**SubAccount** | Pointer to **string** |  | [optional] 
**TradeDate** | Pointer to **time.Time** |  | [optional] 
**SettlementDate** | Pointer to **time.Time** |  | [optional] 
**PositionId** | Pointer to **int64** |  | [optional] 
**OrderId** | Pointer to **int64** |  | [optional] 
**NetAmount** | Pointer to **float64** |  | [optional] 
**ActivityType** | Pointer to **string** |  | [optional] 
**TransferItems** | Pointer to [**[]TransferItem**](TransferItem.md) |  | [optional] 

## Methods

### NewTransaction

`func NewTransaction() *Transaction`

NewTransaction instantiates a new Transaction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionWithDefaults

`func NewTransactionWithDefaults() *Transaction`

NewTransactionWithDefaults instantiates a new Transaction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityId

`func (o *Transaction) GetActivityId() int64`

GetActivityId returns the ActivityId field if non-nil, zero value otherwise.

### GetActivityIdOk

`func (o *Transaction) GetActivityIdOk() (*int64, bool)`

GetActivityIdOk returns a tuple with the ActivityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityId

`func (o *Transaction) SetActivityId(v int64)`

SetActivityId sets ActivityId field to given value.

### HasActivityId

`func (o *Transaction) HasActivityId() bool`

HasActivityId returns a boolean if a field has been set.

### GetTime

`func (o *Transaction) GetTime() time.Time`

GetTime returns the Time field if non-nil, zero value otherwise.

### GetTimeOk

`func (o *Transaction) GetTimeOk() (*time.Time, bool)`

GetTimeOk returns a tuple with the Time field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTime

`func (o *Transaction) SetTime(v time.Time)`

SetTime sets Time field to given value.

### HasTime

`func (o *Transaction) HasTime() bool`

HasTime returns a boolean if a field has been set.

### GetUser

`func (o *Transaction) GetUser() UserDetails`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Transaction) GetUserOk() (*UserDetails, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Transaction) SetUser(v UserDetails)`

SetUser sets User field to given value.

### HasUser

`func (o *Transaction) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetDescription

`func (o *Transaction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Transaction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Transaction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Transaction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetAccountNumber

`func (o *Transaction) GetAccountNumber() string`

GetAccountNumber returns the AccountNumber field if non-nil, zero value otherwise.

### GetAccountNumberOk

`func (o *Transaction) GetAccountNumberOk() (*string, bool)`

GetAccountNumberOk returns a tuple with the AccountNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountNumber

`func (o *Transaction) SetAccountNumber(v string)`

SetAccountNumber sets AccountNumber field to given value.

### HasAccountNumber

`func (o *Transaction) HasAccountNumber() bool`

HasAccountNumber returns a boolean if a field has been set.

### GetType

`func (o *Transaction) GetType() TransactionType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Transaction) GetTypeOk() (*TransactionType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Transaction) SetType(v TransactionType)`

SetType sets Type field to given value.

### HasType

`func (o *Transaction) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *Transaction) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transaction) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transaction) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transaction) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSubAccount

`func (o *Transaction) GetSubAccount() string`

GetSubAccount returns the SubAccount field if non-nil, zero value otherwise.

### GetSubAccountOk

`func (o *Transaction) GetSubAccountOk() (*string, bool)`

GetSubAccountOk returns a tuple with the SubAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubAccount

`func (o *Transaction) SetSubAccount(v string)`

SetSubAccount sets SubAccount field to given value.

### HasSubAccount

`func (o *Transaction) HasSubAccount() bool`

HasSubAccount returns a boolean if a field has been set.

### GetTradeDate

`func (o *Transaction) GetTradeDate() time.Time`

GetTradeDate returns the TradeDate field if non-nil, zero value otherwise.

### GetTradeDateOk

`func (o *Transaction) GetTradeDateOk() (*time.Time, bool)`

GetTradeDateOk returns a tuple with the TradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDate

`func (o *Transaction) SetTradeDate(v time.Time)`

SetTradeDate sets TradeDate field to given value.

### HasTradeDate

`func (o *Transaction) HasTradeDate() bool`

HasTradeDate returns a boolean if a field has been set.

### GetSettlementDate

`func (o *Transaction) GetSettlementDate() time.Time`

GetSettlementDate returns the SettlementDate field if non-nil, zero value otherwise.

### GetSettlementDateOk

`func (o *Transaction) GetSettlementDateOk() (*time.Time, bool)`

GetSettlementDateOk returns a tuple with the SettlementDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettlementDate

`func (o *Transaction) SetSettlementDate(v time.Time)`

SetSettlementDate sets SettlementDate field to given value.

### HasSettlementDate

`func (o *Transaction) HasSettlementDate() bool`

HasSettlementDate returns a boolean if a field has been set.

### GetPositionId

`func (o *Transaction) GetPositionId() int64`

GetPositionId returns the PositionId field if non-nil, zero value otherwise.

### GetPositionIdOk

`func (o *Transaction) GetPositionIdOk() (*int64, bool)`

GetPositionIdOk returns a tuple with the PositionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPositionId

`func (o *Transaction) SetPositionId(v int64)`

SetPositionId sets PositionId field to given value.

### HasPositionId

`func (o *Transaction) HasPositionId() bool`

HasPositionId returns a boolean if a field has been set.

### GetOrderId

`func (o *Transaction) GetOrderId() int64`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *Transaction) GetOrderIdOk() (*int64, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *Transaction) SetOrderId(v int64)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *Transaction) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetNetAmount

`func (o *Transaction) GetNetAmount() float64`

GetNetAmount returns the NetAmount field if non-nil, zero value otherwise.

### GetNetAmountOk

`func (o *Transaction) GetNetAmountOk() (*float64, bool)`

GetNetAmountOk returns a tuple with the NetAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetAmount

`func (o *Transaction) SetNetAmount(v float64)`

SetNetAmount sets NetAmount field to given value.

### HasNetAmount

`func (o *Transaction) HasNetAmount() bool`

HasNetAmount returns a boolean if a field has been set.

### GetActivityType

`func (o *Transaction) GetActivityType() string`

GetActivityType returns the ActivityType field if non-nil, zero value otherwise.

### GetActivityTypeOk

`func (o *Transaction) GetActivityTypeOk() (*string, bool)`

GetActivityTypeOk returns a tuple with the ActivityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityType

`func (o *Transaction) SetActivityType(v string)`

SetActivityType sets ActivityType field to given value.

### HasActivityType

`func (o *Transaction) HasActivityType() bool`

HasActivityType returns a boolean if a field has been set.

### GetTransferItems

`func (o *Transaction) GetTransferItems() []TransferItem`

GetTransferItems returns the TransferItems field if non-nil, zero value otherwise.

### GetTransferItemsOk

`func (o *Transaction) GetTransferItemsOk() (*[]TransferItem, bool)`

GetTransferItemsOk returns a tuple with the TransferItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferItems

`func (o *Transaction) SetTransferItems(v []TransferItem)`

SetTransferItems sets TransferItems field to given value.

### HasTransferItems

`func (o *Transaction) HasTransferItems() bool`

HasTransferItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


