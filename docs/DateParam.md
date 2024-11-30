# DateParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to **string** | Valid ISO-8601 format is :&lt;br&gt; &lt;code&gt;yyyy-MM-dd&#39;T&#39;HH:mm:ss.SSSZ&lt;/code&gt; | [optional] 

## Methods

### NewDateParam

`func NewDateParam() *DateParam`

NewDateParam instantiates a new DateParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDateParamWithDefaults

`func NewDateParamWithDefaults() *DateParam`

NewDateParamWithDefaults instantiates a new DateParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *DateParam) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *DateParam) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *DateParam) SetDate(v string)`

SetDate sets Date field to given value.

### HasDate

`func (o *DateParam) HasDate() bool`

HasDate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


