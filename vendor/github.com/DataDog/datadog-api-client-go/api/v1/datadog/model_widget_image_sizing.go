/*
 * Unless explicitly stated otherwise all files in this repository are licensed under the Apache-2.0 License.
 * This product includes software developed at Datadog (https://www.datadoghq.com/).
 * Copyright 2019-Present Datadog, Inc.
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package datadog

import (
	"encoding/json"
	"fmt"
)

// WidgetImageSizing How to size the image on the widget.
type WidgetImageSizing string

// List of WidgetImageSizing
const (
	WIDGETIMAGESIZING_ZOOM   WidgetImageSizing = "zoom"
	WIDGETIMAGESIZING_FIT    WidgetImageSizing = "fit"
	WIDGETIMAGESIZING_CENTER WidgetImageSizing = "center"
)

func (v *WidgetImageSizing) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WidgetImageSizing(value)
	for _, existing := range []WidgetImageSizing{"zoom", "fit", "center"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WidgetImageSizing", value)
}

// Ptr returns reference to WidgetImageSizing value
func (v WidgetImageSizing) Ptr() *WidgetImageSizing {
	return &v
}

type NullableWidgetImageSizing struct {
	value *WidgetImageSizing
	isSet bool
}

func (v NullableWidgetImageSizing) Get() *WidgetImageSizing {
	return v.value
}

func (v *NullableWidgetImageSizing) Set(val *WidgetImageSizing) {
	v.value = val
	v.isSet = true
}

func (v NullableWidgetImageSizing) IsSet() bool {
	return v.isSet
}

func (v *NullableWidgetImageSizing) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWidgetImageSizing(val *WidgetImageSizing) *NullableWidgetImageSizing {
	return &NullableWidgetImageSizing{value: val, isSet: true}
}

func (v NullableWidgetImageSizing) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWidgetImageSizing) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
