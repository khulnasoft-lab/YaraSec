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

// checks if the ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance{}

// ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance struct for ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance
type ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance struct {
	New []ModelCloudCompliance `json:"new"`
}

type _ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance

// NewModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance instantiates a new ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance(new []ModelCloudCompliance) *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance {
	this := ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance{}
	this.New = new
	return &this
}

// NewModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudComplianceWithDefaults instantiates a new ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudComplianceWithDefaults() *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance {
	this := ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance{}
	return &this
}

// GetNew returns the New field value
// If the value is explicit nil, the zero value for []ModelCloudCompliance will be returned
func (o *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) GetNew() []ModelCloudCompliance {
	if o == nil {
		var ret []ModelCloudCompliance
		return ret
	}

	return o.New
}

// GetNewOk returns a tuple with the New field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) GetNewOk() ([]ModelCloudCompliance, bool) {
	if o == nil || IsNil(o.New) {
		return nil, false
	}
	return o.New, true
}

// SetNew sets field value
func (o *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) SetNew(v []ModelCloudCompliance) {
	o.New = v
}

func (o ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.New != nil {
		toSerialize["new"] = o.New
	}
	return toSerialize, nil
}

func (o *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"new",
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

	varModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance := _ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance)

	if err != nil {
		return err
	}

	*o = ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance(varModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance)

	return err
}

type NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance struct {
	value *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance
	isSet bool
}

func (v NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) Get() *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance {
	return v.value
}

func (v *NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) Set(val *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance(val *ModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) *NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance {
	return &NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance{value: val, isSet: true}
}

func (v NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanCompareResGithubComKhulnasoftKengineKhulnasoftServerModelCloudCompliance) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


