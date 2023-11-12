# CleanUploadRepositoryRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | Pointer to **string** | Specifies the scope of cleanup (e.g., \&quot;expired\&quot; uploads) | [optional] 
**Timestamp** | Pointer to **time.Time** | Timestamp of the task creation | [optional] 
**Type** | Pointer to **string** | Type of the task | [optional] 

## Methods

### NewCleanUploadRepositoryRequest

`func NewCleanUploadRepositoryRequest() *CleanUploadRepositoryRequest`

NewCleanUploadRepositoryRequest instantiates a new CleanUploadRepositoryRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCleanUploadRepositoryRequestWithDefaults

`func NewCleanUploadRepositoryRequestWithDefaults() *CleanUploadRepositoryRequest`

NewCleanUploadRepositoryRequestWithDefaults instantiates a new CleanUploadRepositoryRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *CleanUploadRepositoryRequest) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CleanUploadRepositoryRequest) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CleanUploadRepositoryRequest) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CleanUploadRepositoryRequest) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTimestamp

`func (o *CleanUploadRepositoryRequest) GetTimestamp() time.Time`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *CleanUploadRepositoryRequest) GetTimestampOk() (*time.Time, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *CleanUploadRepositoryRequest) SetTimestamp(v time.Time)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *CleanUploadRepositoryRequest) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetType

`func (o *CleanUploadRepositoryRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CleanUploadRepositoryRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CleanUploadRepositoryRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CleanUploadRepositoryRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


