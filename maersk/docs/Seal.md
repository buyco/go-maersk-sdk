# Seal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SealNumber** | **string** | Identifies a seal affixed to the container. | 
**SealSource** | Pointer to **string** | The source of the seal, namely who has affixed the seal. This attribute links to the Seal Source ID defined in the Seal Source reference data entity. * CAR (Carrier) * SHI (Shipper) * PHY (Phytosanitary) * VET (Veterinary) * CUS (Customs)  | [optional] 
**SealType** | **string** | The type of seal. This attribute links to the Seal Type ID defined in the Seal Type reference data entity. * KLP (Keyless padlock) * BLT (Bolt) * WIR (Wire)  | 

## Methods

### NewSeal

`func NewSeal(sealNumber string, sealType string, ) *Seal`

NewSeal instantiates a new Seal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSealWithDefaults

`func NewSealWithDefaults() *Seal`

NewSealWithDefaults instantiates a new Seal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSealNumber

`func (o *Seal) GetSealNumber() string`

GetSealNumber returns the SealNumber field if non-nil, zero value otherwise.

### GetSealNumberOk

`func (o *Seal) GetSealNumberOk() (*string, bool)`

GetSealNumberOk returns a tuple with the SealNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealNumber

`func (o *Seal) SetSealNumber(v string)`

SetSealNumber sets SealNumber field to given value.


### GetSealSource

`func (o *Seal) GetSealSource() string`

GetSealSource returns the SealSource field if non-nil, zero value otherwise.

### GetSealSourceOk

`func (o *Seal) GetSealSourceOk() (*string, bool)`

GetSealSourceOk returns a tuple with the SealSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealSource

`func (o *Seal) SetSealSource(v string)`

SetSealSource sets SealSource field to given value.

### HasSealSource

`func (o *Seal) HasSealSource() bool`

HasSealSource returns a boolean if a field has been set.

### GetSealType

`func (o *Seal) GetSealType() string`

GetSealType returns the SealType field if non-nil, zero value otherwise.

### GetSealTypeOk

`func (o *Seal) GetSealTypeOk() (*string, bool)`

GetSealTypeOk returns a tuple with the SealType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealType

`func (o *Seal) SetSealType(v string)`

SetSealType sets SealType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


