# UserPreference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]UserPreferenceAccount**](UserPreferenceAccount.md) |  | [optional] 
**StreamerInfo** | Pointer to [**[]StreamerInfo**](StreamerInfo.md) |  | [optional] 
**Offers** | Pointer to [**[]Offer**](Offer.md) |  | [optional] 

## Methods

### NewUserPreference

`func NewUserPreference() *UserPreference`

NewUserPreference instantiates a new UserPreference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPreferenceWithDefaults

`func NewUserPreferenceWithDefaults() *UserPreference`

NewUserPreferenceWithDefaults instantiates a new UserPreference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *UserPreference) GetAccounts() []UserPreferenceAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UserPreference) GetAccountsOk() (*[]UserPreferenceAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UserPreference) SetAccounts(v []UserPreferenceAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UserPreference) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetStreamerInfo

`func (o *UserPreference) GetStreamerInfo() []StreamerInfo`

GetStreamerInfo returns the StreamerInfo field if non-nil, zero value otherwise.

### GetStreamerInfoOk

`func (o *UserPreference) GetStreamerInfoOk() (*[]StreamerInfo, bool)`

GetStreamerInfoOk returns a tuple with the StreamerInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamerInfo

`func (o *UserPreference) SetStreamerInfo(v []StreamerInfo)`

SetStreamerInfo sets StreamerInfo field to given value.

### HasStreamerInfo

`func (o *UserPreference) HasStreamerInfo() bool`

HasStreamerInfo returns a boolean if a field has been set.

### GetOffers

`func (o *UserPreference) GetOffers() []Offer`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *UserPreference) GetOffersOk() (*[]Offer, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *UserPreference) SetOffers(v []Offer)`

SetOffers sets Offers field to given value.

### HasOffers

`func (o *UserPreference) HasOffers() bool`

HasOffers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


