/*
Apache JAMES Web Admin API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi


// checks if the Group type satisfies the MappedNullable interface at compile time

// Group struct for Group
type Group struct {
	// The address of the group
	GroupAddr string `json:"groupAddr"`
	// List of member addresses in the group
	MemberAddrs []string `json:"memberAddrs"`
}

type _Group Group

// NewGroup instantiates a new Group object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroup(groupAddr string, memberAddrs []string) *Group {
	this := Group{}
	this.GroupAddr = groupAddr
	this.MemberAddrs = memberAddrs
	return &this
}