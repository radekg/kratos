/*
 * Ory Kratos API
 *
 * Documentation for all public and administrative Ory Kratos APIs. Public and administrative APIs are exposed on different ports. Public APIs can face the public internet without any protection while administrative APIs should never be exposed without prior authorization. To protect the administative API port you should use something like Nginx, Ory Oathkeeper, or any other technology capable of authorizing incoming requests.
 *
 * API version: 1.0.0
 * Contact: hi@ory.sh
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kratos

import (
	"encoding/json"
)

// SettingsFlowMethod struct for SettingsFlowMethod
type SettingsFlowMethod struct {
	Config SettingsFlowMethodConfig `json:"config"`
	// Method is the name of this flow method.
	Method string `json:"method"`
}

// NewSettingsFlowMethod instantiates a new SettingsFlowMethod object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsFlowMethod(config SettingsFlowMethodConfig, method string) *SettingsFlowMethod {
	this := SettingsFlowMethod{}
	this.Config = config
	this.Method = method
	return &this
}

// NewSettingsFlowMethodWithDefaults instantiates a new SettingsFlowMethod object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsFlowMethodWithDefaults() *SettingsFlowMethod {
	this := SettingsFlowMethod{}
	return &this
}

// GetConfig returns the Config field value
func (o *SettingsFlowMethod) GetConfig() SettingsFlowMethodConfig {
	if o == nil {
		var ret SettingsFlowMethodConfig
		return ret
	}

	return o.Config
}

// GetConfigOk returns a tuple with the Config field value
// and a boolean to check if the value has been set.
func (o *SettingsFlowMethod) GetConfigOk() (*SettingsFlowMethodConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Config, true
}

// SetConfig sets field value
func (o *SettingsFlowMethod) SetConfig(v SettingsFlowMethodConfig) {
	o.Config = v
}

// GetMethod returns the Method field value
func (o *SettingsFlowMethod) GetMethod() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Method
}

// GetMethodOk returns a tuple with the Method field value
// and a boolean to check if the value has been set.
func (o *SettingsFlowMethod) GetMethodOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Method, true
}

// SetMethod sets field value
func (o *SettingsFlowMethod) SetMethod(v string) {
	o.Method = v
}

func (o SettingsFlowMethod) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["config"] = o.Config
	}
	if true {
		toSerialize["method"] = o.Method
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsFlowMethod struct {
	value *SettingsFlowMethod
	isSet bool
}

func (v NullableSettingsFlowMethod) Get() *SettingsFlowMethod {
	return v.value
}

func (v *NullableSettingsFlowMethod) Set(val *SettingsFlowMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsFlowMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsFlowMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsFlowMethod(val *SettingsFlowMethod) *NullableSettingsFlowMethod {
	return &NullableSettingsFlowMethod{value: val, isSet: true}
}

func (v NullableSettingsFlowMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsFlowMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
