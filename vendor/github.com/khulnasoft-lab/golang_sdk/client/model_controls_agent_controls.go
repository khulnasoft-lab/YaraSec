/*
Khulnasoft Kengine

Khulnasoft Runtime API provides programmatic control over Khulnasoft microservice securing your container, kubernetes and cloud deployments. The API abstracts away underlying infrastructure details like cloud provider,  container distros, container orchestrator and type of deployment. This is one uniform API to manage and control security alerts, policies and response to alerts for microservices running anywhere i.e. managed pure greenfield container deployments or a mix of containers, VMs and serverless paradigms like AWS Fargate.

API version: v2.2.1
Contact: community@khulnasoft.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package client

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ControlsAgentControls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ControlsAgentControls{}

// ControlsAgentControls struct for ControlsAgentControls
type ControlsAgentControls struct {
	Beatrate int32 `json:"beatrate"`
	Commands []ControlsAction `json:"commands"`
}

type _ControlsAgentControls ControlsAgentControls

// NewControlsAgentControls instantiates a new ControlsAgentControls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewControlsAgentControls(beatrate int32, commands []ControlsAction) *ControlsAgentControls {
	this := ControlsAgentControls{}
	this.Beatrate = beatrate
	this.Commands = commands
	return &this
}

// NewControlsAgentControlsWithDefaults instantiates a new ControlsAgentControls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewControlsAgentControlsWithDefaults() *ControlsAgentControls {
	this := ControlsAgentControls{}
	return &this
}

// GetBeatrate returns the Beatrate field value
func (o *ControlsAgentControls) GetBeatrate() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Beatrate
}

// GetBeatrateOk returns a tuple with the Beatrate field value
// and a boolean to check if the value has been set.
func (o *ControlsAgentControls) GetBeatrateOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Beatrate, true
}

// SetBeatrate sets field value
func (o *ControlsAgentControls) SetBeatrate(v int32) {
	o.Beatrate = v
}

// GetCommands returns the Commands field value
// If the value is explicit nil, the zero value for []ControlsAction will be returned
func (o *ControlsAgentControls) GetCommands() []ControlsAction {
	if o == nil {
		var ret []ControlsAction
		return ret
	}

	return o.Commands
}

// GetCommandsOk returns a tuple with the Commands field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ControlsAgentControls) GetCommandsOk() ([]ControlsAction, bool) {
	if o == nil || IsNil(o.Commands) {
		return nil, false
	}
	return o.Commands, true
}

// SetCommands sets field value
func (o *ControlsAgentControls) SetCommands(v []ControlsAction) {
	o.Commands = v
}

func (o ControlsAgentControls) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ControlsAgentControls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["beatrate"] = o.Beatrate
	if o.Commands != nil {
		toSerialize["commands"] = o.Commands
	}
	return toSerialize, nil
}

func (o *ControlsAgentControls) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"beatrate",
		"commands",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varControlsAgentControls := _ControlsAgentControls{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varControlsAgentControls)

	if err != nil {
		return err
	}

	*o = ControlsAgentControls(varControlsAgentControls)

	return err
}

type NullableControlsAgentControls struct {
	value *ControlsAgentControls
	isSet bool
}

func (v NullableControlsAgentControls) Get() *ControlsAgentControls {
	return v.value
}

func (v *NullableControlsAgentControls) Set(val *ControlsAgentControls) {
	v.value = val
	v.isSet = true
}

func (v NullableControlsAgentControls) IsSet() bool {
	return v.isSet
}

func (v *NullableControlsAgentControls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableControlsAgentControls(val *ControlsAgentControls) *NullableControlsAgentControls {
	return &NullableControlsAgentControls{value: val, isSet: true}
}

func (v NullableControlsAgentControls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableControlsAgentControls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


