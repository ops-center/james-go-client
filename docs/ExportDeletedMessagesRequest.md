# ExportDeletedMessagesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Combinator** | Pointer to **string** |  | [optional] 
**Criteria** | Pointer to [**[]Criterion**](Criterion.md) |  | [optional] 

## Methods

### NewExportDeletedMessagesRequest

`func NewExportDeletedMessagesRequest() *ExportDeletedMessagesRequest`

NewExportDeletedMessagesRequest instantiates a new ExportDeletedMessagesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExportDeletedMessagesRequestWithDefaults

`func NewExportDeletedMessagesRequestWithDefaults() *ExportDeletedMessagesRequest`

NewExportDeletedMessagesRequestWithDefaults instantiates a new ExportDeletedMessagesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCombinator

`func (o *ExportDeletedMessagesRequest) GetCombinator() string`

GetCombinator returns the Combinator field if non-nil, zero value otherwise.

### GetCombinatorOk

`func (o *ExportDeletedMessagesRequest) GetCombinatorOk() (*string, bool)`

GetCombinatorOk returns a tuple with the Combinator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCombinator

`func (o *ExportDeletedMessagesRequest) SetCombinator(v string)`

SetCombinator sets Combinator field to given value.

### HasCombinator

`func (o *ExportDeletedMessagesRequest) HasCombinator() bool`

HasCombinator returns a boolean if a field has been set.

### GetCriteria

`func (o *ExportDeletedMessagesRequest) GetCriteria() []Criterion`

GetCriteria returns the Criteria field if non-nil, zero value otherwise.

### GetCriteriaOk

`func (o *ExportDeletedMessagesRequest) GetCriteriaOk() (*[]Criterion, bool)`

GetCriteriaOk returns a tuple with the Criteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCriteria

`func (o *ExportDeletedMessagesRequest) SetCriteria(v []Criterion)`

SetCriteria sets Criteria field to given value.

### HasCriteria

`func (o *ExportDeletedMessagesRequest) HasCriteria() bool`

HasCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


