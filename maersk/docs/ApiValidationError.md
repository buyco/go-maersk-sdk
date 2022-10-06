# ApiValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | **string** | The field that has failed validation. | 
**RejectedValue** | **string** | The value that has failed validation. | 
**Message** | **string** | The reason and advice for failed validation. | 

## Methods

### NewApiValidationError

`func NewApiValidationError(field string, rejectedValue string, message string, ) *ApiValidationError`

NewApiValidationError instantiates a new ApiValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiValidationErrorWithDefaults

`func NewApiValidationErrorWithDefaults() *ApiValidationError`

NewApiValidationErrorWithDefaults instantiates a new ApiValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *ApiValidationError) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ApiValidationError) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ApiValidationError) SetField(v string)`

SetField sets Field field to given value.


### GetRejectedValue

`func (o *ApiValidationError) GetRejectedValue() string`

GetRejectedValue returns the RejectedValue field if non-nil, zero value otherwise.

### GetRejectedValueOk

`func (o *ApiValidationError) GetRejectedValueOk() (*string, bool)`

GetRejectedValueOk returns a tuple with the RejectedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedValue

`func (o *ApiValidationError) SetRejectedValue(v string)`

SetRejectedValue sets RejectedValue field to given value.


### GetMessage

`func (o *ApiValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


