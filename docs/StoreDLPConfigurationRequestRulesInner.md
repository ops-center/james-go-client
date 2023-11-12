# StoreDLPConfigurationRequestRulesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Explanation** | Pointer to **string** | Description of the configuration item | [optional] 
**Expression** | Pointer to **string** | Regular expression to match contents of targets | [optional] 
**Id** | Pointer to **string** | Unique identifier of the configuration item | [optional] 
**TargetsContent** | Pointer to **bool** | Indicates if the expression should be applied to the subject headers and textual bodies of the mail | [optional] 
**TargetsRecipients** | Pointer to **bool** | Indicates if the expression should be applied to the recipients of the mail | [optional] 
**TargetsSender** | Pointer to **bool** | Indicates if the expression should be applied to the sender and from headers of the mail | [optional] 

## Methods

### NewStoreDLPConfigurationRequestRulesInner

`func NewStoreDLPConfigurationRequestRulesInner() *StoreDLPConfigurationRequestRulesInner`

NewStoreDLPConfigurationRequestRulesInner instantiates a new StoreDLPConfigurationRequestRulesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreDLPConfigurationRequestRulesInnerWithDefaults

`func NewStoreDLPConfigurationRequestRulesInnerWithDefaults() *StoreDLPConfigurationRequestRulesInner`

NewStoreDLPConfigurationRequestRulesInnerWithDefaults instantiates a new StoreDLPConfigurationRequestRulesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExplanation

`func (o *StoreDLPConfigurationRequestRulesInner) GetExplanation() string`

GetExplanation returns the Explanation field if non-nil, zero value otherwise.

### GetExplanationOk

`func (o *StoreDLPConfigurationRequestRulesInner) GetExplanationOk() (*string, bool)`

GetExplanationOk returns a tuple with the Explanation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplanation

`func (o *StoreDLPConfigurationRequestRulesInner) SetExplanation(v string)`

SetExplanation sets Explanation field to given value.

### HasExplanation

`func (o *StoreDLPConfigurationRequestRulesInner) HasExplanation() bool`

HasExplanation returns a boolean if a field has been set.

### GetExpression

`func (o *StoreDLPConfigurationRequestRulesInner) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *StoreDLPConfigurationRequestRulesInner) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *StoreDLPConfigurationRequestRulesInner) SetExpression(v string)`

SetExpression sets Expression field to given value.

### HasExpression

`func (o *StoreDLPConfigurationRequestRulesInner) HasExpression() bool`

HasExpression returns a boolean if a field has been set.

### GetId

`func (o *StoreDLPConfigurationRequestRulesInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoreDLPConfigurationRequestRulesInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoreDLPConfigurationRequestRulesInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoreDLPConfigurationRequestRulesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetsContent

`func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsContent() bool`

GetTargetsContent returns the TargetsContent field if non-nil, zero value otherwise.

### GetTargetsContentOk

`func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsContentOk() (*bool, bool)`

GetTargetsContentOk returns a tuple with the TargetsContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsContent

`func (o *StoreDLPConfigurationRequestRulesInner) SetTargetsContent(v bool)`

SetTargetsContent sets TargetsContent field to given value.

### HasTargetsContent

`func (o *StoreDLPConfigurationRequestRulesInner) HasTargetsContent() bool`

HasTargetsContent returns a boolean if a field has been set.

### GetTargetsRecipients

`func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsRecipients() bool`

GetTargetsRecipients returns the TargetsRecipients field if non-nil, zero value otherwise.

### GetTargetsRecipientsOk

`func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsRecipientsOk() (*bool, bool)`

GetTargetsRecipientsOk returns a tuple with the TargetsRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsRecipients

`func (o *StoreDLPConfigurationRequestRulesInner) SetTargetsRecipients(v bool)`

SetTargetsRecipients sets TargetsRecipients field to given value.

### HasTargetsRecipients

`func (o *StoreDLPConfigurationRequestRulesInner) HasTargetsRecipients() bool`

HasTargetsRecipients returns a boolean if a field has been set.

### GetTargetsSender

`func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsSender() bool`

GetTargetsSender returns the TargetsSender field if non-nil, zero value otherwise.

### GetTargetsSenderOk

`func (o *StoreDLPConfigurationRequestRulesInner) GetTargetsSenderOk() (*bool, bool)`

GetTargetsSenderOk returns a tuple with the TargetsSender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetsSender

`func (o *StoreDLPConfigurationRequestRulesInner) SetTargetsSender(v bool)`

SetTargetsSender sets TargetsSender field to given value.

### HasTargetsSender

`func (o *StoreDLPConfigurationRequestRulesInner) HasTargetsSender() bool`

HasTargetsSender returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


