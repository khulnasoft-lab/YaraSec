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

// checks if the ModelScanResultsMaskRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ModelScanResultsMaskRequest{}

// ModelScanResultsMaskRequest struct for ModelScanResultsMaskRequest
type ModelScanResultsMaskRequest struct {
	MaskAction string `json:"mask_action"`
	ResultIds []string `json:"result_ids"`
	ScanId string `json:"scan_id"`
	ScanType string `json:"scan_type"`
}

type _ModelScanResultsMaskRequest ModelScanResultsMaskRequest

// NewModelScanResultsMaskRequest instantiates a new ModelScanResultsMaskRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModelScanResultsMaskRequest(maskAction string, resultIds []string, scanId string, scanType string) *ModelScanResultsMaskRequest {
	this := ModelScanResultsMaskRequest{}
	this.MaskAction = maskAction
	this.ResultIds = resultIds
	this.ScanId = scanId
	this.ScanType = scanType
	return &this
}

// NewModelScanResultsMaskRequestWithDefaults instantiates a new ModelScanResultsMaskRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModelScanResultsMaskRequestWithDefaults() *ModelScanResultsMaskRequest {
	this := ModelScanResultsMaskRequest{}
	return &this
}

// GetMaskAction returns the MaskAction field value
func (o *ModelScanResultsMaskRequest) GetMaskAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MaskAction
}

// GetMaskActionOk returns a tuple with the MaskAction field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsMaskRequest) GetMaskActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaskAction, true
}

// SetMaskAction sets field value
func (o *ModelScanResultsMaskRequest) SetMaskAction(v string) {
	o.MaskAction = v
}

// GetResultIds returns the ResultIds field value
// If the value is explicit nil, the zero value for []string will be returned
func (o *ModelScanResultsMaskRequest) GetResultIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ResultIds
}

// GetResultIdsOk returns a tuple with the ResultIds field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ModelScanResultsMaskRequest) GetResultIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ResultIds) {
		return nil, false
	}
	return o.ResultIds, true
}

// SetResultIds sets field value
func (o *ModelScanResultsMaskRequest) SetResultIds(v []string) {
	o.ResultIds = v
}

// GetScanId returns the ScanId field value
func (o *ModelScanResultsMaskRequest) GetScanId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanId
}

// GetScanIdOk returns a tuple with the ScanId field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsMaskRequest) GetScanIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanId, true
}

// SetScanId sets field value
func (o *ModelScanResultsMaskRequest) SetScanId(v string) {
	o.ScanId = v
}

// GetScanType returns the ScanType field value
func (o *ModelScanResultsMaskRequest) GetScanType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ScanType
}

// GetScanTypeOk returns a tuple with the ScanType field value
// and a boolean to check if the value has been set.
func (o *ModelScanResultsMaskRequest) GetScanTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScanType, true
}

// SetScanType sets field value
func (o *ModelScanResultsMaskRequest) SetScanType(v string) {
	o.ScanType = v
}

func (o ModelScanResultsMaskRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ModelScanResultsMaskRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mask_action"] = o.MaskAction
	if o.ResultIds != nil {
		toSerialize["result_ids"] = o.ResultIds
	}
	toSerialize["scan_id"] = o.ScanId
	toSerialize["scan_type"] = o.ScanType
	return toSerialize, nil
}

func (o *ModelScanResultsMaskRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"mask_action",
		"result_ids",
		"scan_id",
		"scan_type",
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

	varModelScanResultsMaskRequest := _ModelScanResultsMaskRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varModelScanResultsMaskRequest)

	if err != nil {
		return err
	}

	*o = ModelScanResultsMaskRequest(varModelScanResultsMaskRequest)

	return err
}

type NullableModelScanResultsMaskRequest struct {
	value *ModelScanResultsMaskRequest
	isSet bool
}

func (v NullableModelScanResultsMaskRequest) Get() *ModelScanResultsMaskRequest {
	return v.value
}

func (v *NullableModelScanResultsMaskRequest) Set(val *ModelScanResultsMaskRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableModelScanResultsMaskRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableModelScanResultsMaskRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModelScanResultsMaskRequest(val *ModelScanResultsMaskRequest) *NullableModelScanResultsMaskRequest {
	return &NullableModelScanResultsMaskRequest{value: val, isSet: true}
}

func (v NullableModelScanResultsMaskRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModelScanResultsMaskRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


