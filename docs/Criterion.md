# Criterion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldName** | Pointer to **string** |  | [optional] 
**Operator** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 

## Methods

### NewCriterion

`func NewCriterion() *Criterion`

NewCriterion instantiates a new Criterion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCriterionWithDefaults

`func NewCriterionWithDefaults() *Criterion`

NewCriterionWithDefaults instantiates a new Criterion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldName

`func (o *Criterion) GetFieldName() string`

GetFieldName returns the FieldName field if non-nil, zero value otherwise.

### GetFieldNameOk

`func (o *Criterion) GetFieldNameOk() (*string, bool)`

GetFieldNameOk returns a tuple with the FieldName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldName

`func (o *Criterion) SetFieldName(v string)`

SetFieldName sets FieldName field to given value.

### HasFieldName

`func (o *Criterion) HasFieldName() bool`

HasFieldName returns a boolean if a field has been set.

### GetOperator

`func (o *Criterion) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *Criterion) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *Criterion) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *Criterion) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetValue

`func (o *Criterion) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *Criterion) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *Criterion) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *Criterion) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


