# ApiError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Method** | **string** | The request method type e.g. GET, POST. | 
**RequestUri** | **string** | The request URI. | 
**Status** | **int32** | HTTP status code. | 
**Timestamp** | **string** | The date and time (dd-MM-yyyy hh:mm:ss) the error occured. | 
**Message** | **string** | High level error message. | 
**DebugMessage** | Pointer to **string** | Detailed error message. | [optional] 
**SubErrors** | Pointer to [**[]ApiValidationError**](ApiValidationError.md) | The list of invalid fields in the request. | [optional] 

## Methods

### NewApiError

`func NewApiError(method string, requestUri string, status int32, timestamp string, message string, ) *ApiError`

NewApiError instantiates a new ApiError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiErrorWithDefaults

`func NewApiErrorWithDefaults() *ApiError`

NewApiErrorWithDefaults instantiates a new ApiError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ApiError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiError) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiError) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMethod

`func (o *ApiError) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ApiError) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ApiError) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetRequestUri

`func (o *ApiError) GetRequestUri() string`

GetRequestUri returns the RequestUri field if non-nil, zero value otherwise.

### GetRequestUriOk

`func (o *ApiError) GetRequestUriOk() (*string, bool)`

GetRequestUriOk returns a tuple with the RequestUri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestUri

`func (o *ApiError) SetRequestUri(v string)`

SetRequestUri sets RequestUri field to given value.


### GetStatus

`func (o *ApiError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ApiError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ApiError) SetStatus(v int32)`

SetStatus sets Status field to given value.


### GetTimestamp

`func (o *ApiError) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *ApiError) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *ApiError) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetMessage

`func (o *ApiError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ApiError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ApiError) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetDebugMessage

`func (o *ApiError) GetDebugMessage() string`

GetDebugMessage returns the DebugMessage field if non-nil, zero value otherwise.

### GetDebugMessageOk

`func (o *ApiError) GetDebugMessageOk() (*string, bool)`

GetDebugMessageOk returns a tuple with the DebugMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugMessage

`func (o *ApiError) SetDebugMessage(v string)`

SetDebugMessage sets DebugMessage field to given value.

### HasDebugMessage

`func (o *ApiError) HasDebugMessage() bool`

HasDebugMessage returns a boolean if a field has been set.

### GetSubErrors

`func (o *ApiError) GetSubErrors() []ApiValidationError`

GetSubErrors returns the SubErrors field if non-nil, zero value otherwise.

### GetSubErrorsOk

`func (o *ApiError) GetSubErrorsOk() (*[]ApiValidationError, bool)`

GetSubErrorsOk returns a tuple with the SubErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubErrors

`func (o *ApiError) SetSubErrors(v []ApiValidationError)`

SetSubErrors sets SubErrors field to given value.

### HasSubErrors

`func (o *ApiError) HasSubErrors() bool`

HasSubErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


