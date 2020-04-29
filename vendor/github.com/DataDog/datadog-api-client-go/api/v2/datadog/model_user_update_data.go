/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
)

// UserUpdateData Object to update a user.
type UserUpdateData struct {
	Attributes *UserUpdateAttributes `json:"attributes,omitempty"`
	// ID of the user.
	Id *string `json:"id,omitempty"`
	// Users resource type.
	Type *string `json:"type,omitempty"`
}

// NewUserUpdateData instantiates a new UserUpdateData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateData() *UserUpdateData {
	this := UserUpdateData{}
	var type_ string = "users"
	this.Type = &type_
	return &this
}

// NewUserUpdateDataWithDefaults instantiates a new UserUpdateData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateDataWithDefaults() *UserUpdateData {
	this := UserUpdateData{}
	var type_ string = "users"
	this.Type = &type_
	return &this
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *UserUpdateData) GetAttributes() UserUpdateAttributes {
	if o == nil || o.Attributes == nil {
		var ret UserUpdateAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateData) GetAttributesOk() (*UserUpdateAttributes, bool) {
	if o == nil || o.Attributes == nil {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *UserUpdateData) HasAttributes() bool {
	if o != nil && o.Attributes != nil {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given UserUpdateAttributes and assigns it to the Attributes field.
func (o *UserUpdateData) SetAttributes(v UserUpdateAttributes) {
	o.Attributes = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserUpdateData) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateData) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserUpdateData) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UserUpdateData) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *UserUpdateData) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateData) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *UserUpdateData) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *UserUpdateData) SetType(v string) {
	o.Type = &v
}

func (o UserUpdateData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Attributes != nil {
		toSerialize["attributes"] = o.Attributes
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableUserUpdateData struct {
	value *UserUpdateData
	isSet bool
}

func (v NullableUserUpdateData) Get() *UserUpdateData {
	return v.value
}

func (v *NullableUserUpdateData) Set(val *UserUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateData(val *UserUpdateData) *NullableUserUpdateData {
	return &NullableUserUpdateData{value: val, isSet: true}
}

func (v NullableUserUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
