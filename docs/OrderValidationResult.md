# OrderValidationResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alerts** | Pointer to [**[]OrderValidationDetail**](OrderValidationDetail.md) |  | [optional] 
**Accepts** | Pointer to [**[]OrderValidationDetail**](OrderValidationDetail.md) |  | [optional] 
**Rejects** | Pointer to [**[]OrderValidationDetail**](OrderValidationDetail.md) |  | [optional] 
**Reviews** | Pointer to [**[]OrderValidationDetail**](OrderValidationDetail.md) |  | [optional] 
**Warns** | Pointer to [**[]OrderValidationDetail**](OrderValidationDetail.md) |  | [optional] 

## Methods

### NewOrderValidationResult

`func NewOrderValidationResult() *OrderValidationResult`

NewOrderValidationResult instantiates a new OrderValidationResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderValidationResultWithDefaults

`func NewOrderValidationResultWithDefaults() *OrderValidationResult`

NewOrderValidationResultWithDefaults instantiates a new OrderValidationResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlerts

`func (o *OrderValidationResult) GetAlerts() []OrderValidationDetail`

GetAlerts returns the Alerts field if non-nil, zero value otherwise.

### GetAlertsOk

`func (o *OrderValidationResult) GetAlertsOk() (*[]OrderValidationDetail, bool)`

GetAlertsOk returns a tuple with the Alerts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlerts

`func (o *OrderValidationResult) SetAlerts(v []OrderValidationDetail)`

SetAlerts sets Alerts field to given value.

### HasAlerts

`func (o *OrderValidationResult) HasAlerts() bool`

HasAlerts returns a boolean if a field has been set.

### GetAccepts

`func (o *OrderValidationResult) GetAccepts() []OrderValidationDetail`

GetAccepts returns the Accepts field if non-nil, zero value otherwise.

### GetAcceptsOk

`func (o *OrderValidationResult) GetAcceptsOk() (*[]OrderValidationDetail, bool)`

GetAcceptsOk returns a tuple with the Accepts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepts

`func (o *OrderValidationResult) SetAccepts(v []OrderValidationDetail)`

SetAccepts sets Accepts field to given value.

### HasAccepts

`func (o *OrderValidationResult) HasAccepts() bool`

HasAccepts returns a boolean if a field has been set.

### GetRejects

`func (o *OrderValidationResult) GetRejects() []OrderValidationDetail`

GetRejects returns the Rejects field if non-nil, zero value otherwise.

### GetRejectsOk

`func (o *OrderValidationResult) GetRejectsOk() (*[]OrderValidationDetail, bool)`

GetRejectsOk returns a tuple with the Rejects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejects

`func (o *OrderValidationResult) SetRejects(v []OrderValidationDetail)`

SetRejects sets Rejects field to given value.

### HasRejects

`func (o *OrderValidationResult) HasRejects() bool`

HasRejects returns a boolean if a field has been set.

### GetReviews

`func (o *OrderValidationResult) GetReviews() []OrderValidationDetail`

GetReviews returns the Reviews field if non-nil, zero value otherwise.

### GetReviewsOk

`func (o *OrderValidationResult) GetReviewsOk() (*[]OrderValidationDetail, bool)`

GetReviewsOk returns a tuple with the Reviews field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviews

`func (o *OrderValidationResult) SetReviews(v []OrderValidationDetail)`

SetReviews sets Reviews field to given value.

### HasReviews

`func (o *OrderValidationResult) HasReviews() bool`

HasReviews returns a boolean if a field has been set.

### GetWarns

`func (o *OrderValidationResult) GetWarns() []OrderValidationDetail`

GetWarns returns the Warns field if non-nil, zero value otherwise.

### GetWarnsOk

`func (o *OrderValidationResult) GetWarnsOk() (*[]OrderValidationDetail, bool)`

GetWarnsOk returns a tuple with the Warns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarns

`func (o *OrderValidationResult) SetWarns(v []OrderValidationDetail)`

SetWarns sets Warns field to given value.

### HasWarns

`func (o *OrderValidationResult) HasWarns() bool`

HasWarns returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


