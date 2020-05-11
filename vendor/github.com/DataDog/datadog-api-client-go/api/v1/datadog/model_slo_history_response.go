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

// SLOHistoryResponse A service level objective history response.
type SLOHistoryResponse struct {
	Data *SLOHistoryResponseData `json:"data,omitempty"`
	// A list of errors while querying the history data for the service level objective.
	Errors *[]SLOHistoryResponseError `json:"errors,omitempty"`
}

// NewSLOHistoryResponse instantiates a new SLOHistoryResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSLOHistoryResponse() *SLOHistoryResponse {
	this := SLOHistoryResponse{}
	return &this
}

// NewSLOHistoryResponseWithDefaults instantiates a new SLOHistoryResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSLOHistoryResponseWithDefaults() *SLOHistoryResponse {
	this := SLOHistoryResponse{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *SLOHistoryResponse) GetData() SLOHistoryResponseData {
	if o == nil || o.Data == nil {
		var ret SLOHistoryResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistoryResponse) GetDataOk() (*SLOHistoryResponseData, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *SLOHistoryResponse) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given SLOHistoryResponseData and assigns it to the Data field.
func (o *SLOHistoryResponse) SetData(v SLOHistoryResponseData) {
	o.Data = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *SLOHistoryResponse) GetErrors() []SLOHistoryResponseError {
	if o == nil || o.Errors == nil {
		var ret []SLOHistoryResponseError
		return ret
	}
	return *o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SLOHistoryResponse) GetErrorsOk() (*[]SLOHistoryResponseError, bool) {
	if o == nil || o.Errors == nil {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *SLOHistoryResponse) HasErrors() bool {
	if o != nil && o.Errors != nil {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []SLOHistoryResponseError and assigns it to the Errors field.
func (o *SLOHistoryResponse) SetErrors(v []SLOHistoryResponseError) {
	o.Errors = &v
}

func (o SLOHistoryResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Errors != nil {
		toSerialize["errors"] = o.Errors
	}
	return json.Marshal(toSerialize)
}

type NullableSLOHistoryResponse struct {
	value *SLOHistoryResponse
	isSet bool
}

func (v NullableSLOHistoryResponse) Get() *SLOHistoryResponse {
	return v.value
}

func (v *NullableSLOHistoryResponse) Set(val *SLOHistoryResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSLOHistoryResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSLOHistoryResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSLOHistoryResponse(val *SLOHistoryResponse) *NullableSLOHistoryResponse {
	return &NullableSLOHistoryResponse{value: val, isSet: true}
}

func (v NullableSLOHistoryResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSLOHistoryResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}