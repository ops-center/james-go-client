# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupAddr** | **string** | The address of the group | 
**MemberAddrs** | **[]string** | List of member addresses in the group | 

## Methods

### NewGroup

`func NewGroup(groupAddr string, memberAddrs []string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupAddr

`func (o *Group) GetGroupAddr() string`

GetGroupAddr returns the GroupAddr field if non-nil, zero value otherwise.

### GetGroupAddrOk

`func (o *Group) GetGroupAddrOk() (*string, bool)`

GetGroupAddrOk returns a tuple with the GroupAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupAddr

`func (o *Group) SetGroupAddr(v string)`

SetGroupAddr sets GroupAddr field to given value.


### GetMemberAddrs

`func (o *Group) GetMemberAddrs() []string`

GetMemberAddrs returns the MemberAddrs field if non-nil, zero value otherwise.

### GetMemberAddrsOk

`func (o *Group) GetMemberAddrsOk() (*[]string, bool)`

GetMemberAddrsOk returns a tuple with the MemberAddrs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberAddrs

`func (o *Group) SetMemberAddrs(v []string)`

SetMemberAddrs sets MemberAddrs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


