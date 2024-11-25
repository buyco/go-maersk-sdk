# CreateAccessToken401Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**[]CreateAccessToken401ResponseErrorInner**](CreateAccessToken401ResponseErrorInner.md) |  | [optional] 

## Methods

### NewCreateAccessToken401Response

`func NewCreateAccessToken401Response() *CreateAccessToken401Response`

NewCreateAccessToken401Response instantiates a new CreateAccessToken401Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccessToken401ResponseWithDefaults

`func NewCreateAccessToken401ResponseWithDefaults() *CreateAccessToken401Response`

NewCreateAccessToken401ResponseWithDefaults instantiates a new CreateAccessToken401Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CreateAccessToken401Response) GetError() []CreateAccessToken401ResponseErrorInner`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateAccessToken401Response) GetErrorOk() (*[]CreateAccessToken401ResponseErrorInner, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateAccessToken401Response) SetError(v []CreateAccessToken401ResponseErrorInner)`

SetError sets Error field to given value.

### HasError

`func (o *CreateAccessToken401Response) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


