# OrderValidationDetail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ValidationRuleName** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**ActivityMessage** | Pointer to **string** |  | [optional] 
**OriginalSeverity** | Pointer to [**APIRuleAction**](APIRuleAction.md) |  | [optional] 
**OverrideName** | Pointer to **string** |  | [optional] 
**OverrideSeverity** | Pointer to [**APIRuleAction**](APIRuleAction.md) |  | [optional] 

## Methods

### NewOrderValidationDetail

`func NewOrderValidationDetail() *OrderValidationDetail`

NewOrderValidationDetail instantiates a new OrderValidationDetail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderValidationDetailWithDefaults

`func NewOrderValidationDetailWithDefaults() *OrderValidationDetail`

NewOrderValidationDetailWithDefaults instantiates a new OrderValidationDetail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValidationRuleName

`func (o *OrderValidationDetail) GetValidationRuleName() string`

GetValidationRuleName returns the ValidationRuleName field if non-nil, zero value otherwise.

### GetValidationRuleNameOk

`func (o *OrderValidationDetail) GetValidationRuleNameOk() (*string, bool)`

GetValidationRuleNameOk returns a tuple with the ValidationRuleName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidationRuleName

`func (o *OrderValidationDetail) SetValidationRuleName(v string)`

SetValidationRuleName sets ValidationRuleName field to given value.

### HasValidationRuleName

`func (o *OrderValidationDetail) HasValidationRuleName() bool`

HasValidationRuleName returns a boolean if a field has been set.

### GetMessage

`func (o *OrderValidationDetail) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OrderValidationDetail) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OrderValidationDetail) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *OrderValidationDetail) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetActivityMessage

`func (o *OrderValidationDetail) GetActivityMessage() string`

GetActivityMessage returns the ActivityMessage field if non-nil, zero value otherwise.

### GetActivityMessageOk

`func (o *OrderValidationDetail) GetActivityMessageOk() (*string, bool)`

GetActivityMessageOk returns a tuple with the ActivityMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityMessage

`func (o *OrderValidationDetail) SetActivityMessage(v string)`

SetActivityMessage sets ActivityMessage field to given value.

### HasActivityMessage

`func (o *OrderValidationDetail) HasActivityMessage() bool`

HasActivityMessage returns a boolean if a field has been set.

### GetOriginalSeverity

`func (o *OrderValidationDetail) GetOriginalSeverity() APIRuleAction`

GetOriginalSeverity returns the OriginalSeverity field if non-nil, zero value otherwise.

### GetOriginalSeverityOk

`func (o *OrderValidationDetail) GetOriginalSeverityOk() (*APIRuleAction, bool)`

GetOriginalSeverityOk returns a tuple with the OriginalSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalSeverity

`func (o *OrderValidationDetail) SetOriginalSeverity(v APIRuleAction)`

SetOriginalSeverity sets OriginalSeverity field to given value.

### HasOriginalSeverity

`func (o *OrderValidationDetail) HasOriginalSeverity() bool`

HasOriginalSeverity returns a boolean if a field has been set.

### GetOverrideName

`func (o *OrderValidationDetail) GetOverrideName() string`

GetOverrideName returns the OverrideName field if non-nil, zero value otherwise.

### GetOverrideNameOk

`func (o *OrderValidationDetail) GetOverrideNameOk() (*string, bool)`

GetOverrideNameOk returns a tuple with the OverrideName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideName

`func (o *OrderValidationDetail) SetOverrideName(v string)`

SetOverrideName sets OverrideName field to given value.

### HasOverrideName

`func (o *OrderValidationDetail) HasOverrideName() bool`

HasOverrideName returns a boolean if a field has been set.

### GetOverrideSeverity

`func (o *OrderValidationDetail) GetOverrideSeverity() APIRuleAction`

GetOverrideSeverity returns the OverrideSeverity field if non-nil, zero value otherwise.

### GetOverrideSeverityOk

`func (o *OrderValidationDetail) GetOverrideSeverityOk() (*APIRuleAction, bool)`

GetOverrideSeverityOk returns a tuple with the OverrideSeverity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrideSeverity

`func (o *OrderValidationDetail) SetOverrideSeverity(v APIRuleAction)`

SetOverrideSeverity sets OverrideSeverity field to given value.

### HasOverrideSeverity

`func (o *OrderValidationDetail) HasOverrideSeverity() bool`

HasOverrideSeverity returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


