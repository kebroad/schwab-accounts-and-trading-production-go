# Offer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Level2Permissions** | Pointer to **bool** |  | [optional] [default to false]
**MktDataPermission** | Pointer to **string** |  | [optional] 

## Methods

### NewOffer

`func NewOffer() *Offer`

NewOffer instantiates a new Offer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferWithDefaults

`func NewOfferWithDefaults() *Offer`

NewOfferWithDefaults instantiates a new Offer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLevel2Permissions

`func (o *Offer) GetLevel2Permissions() bool`

GetLevel2Permissions returns the Level2Permissions field if non-nil, zero value otherwise.

### GetLevel2PermissionsOk

`func (o *Offer) GetLevel2PermissionsOk() (*bool, bool)`

GetLevel2PermissionsOk returns a tuple with the Level2Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel2Permissions

`func (o *Offer) SetLevel2Permissions(v bool)`

SetLevel2Permissions sets Level2Permissions field to given value.

### HasLevel2Permissions

`func (o *Offer) HasLevel2Permissions() bool`

HasLevel2Permissions returns a boolean if a field has been set.

### GetMktDataPermission

`func (o *Offer) GetMktDataPermission() string`

GetMktDataPermission returns the MktDataPermission field if non-nil, zero value otherwise.

### GetMktDataPermissionOk

`func (o *Offer) GetMktDataPermissionOk() (*string, bool)`

GetMktDataPermissionOk returns a tuple with the MktDataPermission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMktDataPermission

`func (o *Offer) SetMktDataPermission(v string)`

SetMktDataPermission sets MktDataPermission field to given value.

### HasMktDataPermission

`func (o *Offer) HasMktDataPermission() bool`

HasMktDataPermission returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


