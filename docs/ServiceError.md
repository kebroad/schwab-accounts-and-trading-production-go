# ServiceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewServiceError

`func NewServiceError() *ServiceError`

NewServiceError instantiates a new ServiceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceErrorWithDefaults

`func NewServiceErrorWithDefaults() *ServiceError`

NewServiceErrorWithDefaults instantiates a new ServiceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ServiceError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ServiceError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ServiceError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *ServiceError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetErrors

`func (o *ServiceError) GetErrors() []string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ServiceError) GetErrorsOk() (*[]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ServiceError) SetErrors(v []string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ServiceError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


