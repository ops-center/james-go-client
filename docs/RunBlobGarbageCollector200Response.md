# RunBlobGarbageCollector200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BlobCount** | Pointer to **int32** | Count of blobs tried against the bloom filter | [optional] 
**BloomFilterAssociatedProbability** | Pointer to **float64** | Supplied associatedProbability query parameter | [optional] 
**BloomFilterExpectedBlobCount** | Pointer to **int32** | Supplied expectedBlobCount query parameter | [optional] 
**GcedBlobCount** | Pointer to **int32** | Count of blobs that were garbage collected | [optional] 
**ReferenceSourceCount** | Pointer to **int32** | Count of distinct blob references encountered while populating the bloom filter | [optional] 

## Methods

### NewRunBlobGarbageCollector200Response

`func NewRunBlobGarbageCollector200Response() *RunBlobGarbageCollector200Response`

NewRunBlobGarbageCollector200Response instantiates a new RunBlobGarbageCollector200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunBlobGarbageCollector200ResponseWithDefaults

`func NewRunBlobGarbageCollector200ResponseWithDefaults() *RunBlobGarbageCollector200Response`

NewRunBlobGarbageCollector200ResponseWithDefaults instantiates a new RunBlobGarbageCollector200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlobCount

`func (o *RunBlobGarbageCollector200Response) GetBlobCount() int32`

GetBlobCount returns the BlobCount field if non-nil, zero value otherwise.

### GetBlobCountOk

`func (o *RunBlobGarbageCollector200Response) GetBlobCountOk() (*int32, bool)`

GetBlobCountOk returns a tuple with the BlobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlobCount

`func (o *RunBlobGarbageCollector200Response) SetBlobCount(v int32)`

SetBlobCount sets BlobCount field to given value.

### HasBlobCount

`func (o *RunBlobGarbageCollector200Response) HasBlobCount() bool`

HasBlobCount returns a boolean if a field has been set.

### GetBloomFilterAssociatedProbability

`func (o *RunBlobGarbageCollector200Response) GetBloomFilterAssociatedProbability() float64`

GetBloomFilterAssociatedProbability returns the BloomFilterAssociatedProbability field if non-nil, zero value otherwise.

### GetBloomFilterAssociatedProbabilityOk

`func (o *RunBlobGarbageCollector200Response) GetBloomFilterAssociatedProbabilityOk() (*float64, bool)`

GetBloomFilterAssociatedProbabilityOk returns a tuple with the BloomFilterAssociatedProbability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloomFilterAssociatedProbability

`func (o *RunBlobGarbageCollector200Response) SetBloomFilterAssociatedProbability(v float64)`

SetBloomFilterAssociatedProbability sets BloomFilterAssociatedProbability field to given value.

### HasBloomFilterAssociatedProbability

`func (o *RunBlobGarbageCollector200Response) HasBloomFilterAssociatedProbability() bool`

HasBloomFilterAssociatedProbability returns a boolean if a field has been set.

### GetBloomFilterExpectedBlobCount

`func (o *RunBlobGarbageCollector200Response) GetBloomFilterExpectedBlobCount() int32`

GetBloomFilterExpectedBlobCount returns the BloomFilterExpectedBlobCount field if non-nil, zero value otherwise.

### GetBloomFilterExpectedBlobCountOk

`func (o *RunBlobGarbageCollector200Response) GetBloomFilterExpectedBlobCountOk() (*int32, bool)`

GetBloomFilterExpectedBlobCountOk returns a tuple with the BloomFilterExpectedBlobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBloomFilterExpectedBlobCount

`func (o *RunBlobGarbageCollector200Response) SetBloomFilterExpectedBlobCount(v int32)`

SetBloomFilterExpectedBlobCount sets BloomFilterExpectedBlobCount field to given value.

### HasBloomFilterExpectedBlobCount

`func (o *RunBlobGarbageCollector200Response) HasBloomFilterExpectedBlobCount() bool`

HasBloomFilterExpectedBlobCount returns a boolean if a field has been set.

### GetGcedBlobCount

`func (o *RunBlobGarbageCollector200Response) GetGcedBlobCount() int32`

GetGcedBlobCount returns the GcedBlobCount field if non-nil, zero value otherwise.

### GetGcedBlobCountOk

`func (o *RunBlobGarbageCollector200Response) GetGcedBlobCountOk() (*int32, bool)`

GetGcedBlobCountOk returns a tuple with the GcedBlobCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcedBlobCount

`func (o *RunBlobGarbageCollector200Response) SetGcedBlobCount(v int32)`

SetGcedBlobCount sets GcedBlobCount field to given value.

### HasGcedBlobCount

`func (o *RunBlobGarbageCollector200Response) HasGcedBlobCount() bool`

HasGcedBlobCount returns a boolean if a field has been set.

### GetReferenceSourceCount

`func (o *RunBlobGarbageCollector200Response) GetReferenceSourceCount() int32`

GetReferenceSourceCount returns the ReferenceSourceCount field if non-nil, zero value otherwise.

### GetReferenceSourceCountOk

`func (o *RunBlobGarbageCollector200Response) GetReferenceSourceCountOk() (*int32, bool)`

GetReferenceSourceCountOk returns a tuple with the ReferenceSourceCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceSourceCount

`func (o *RunBlobGarbageCollector200Response) SetReferenceSourceCount(v int32)`

SetReferenceSourceCount sets ReferenceSourceCount field to given value.

### HasReferenceSourceCount

`func (o *RunBlobGarbageCollector200Response) HasReferenceSourceCount() bool`

HasReferenceSourceCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


