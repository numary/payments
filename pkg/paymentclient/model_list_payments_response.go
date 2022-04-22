/*
Sample API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package paymentclient

import (
	"encoding/json"
)

// ListPaymentsResponse struct for ListPaymentsResponse
type ListPaymentsResponse struct {
	Data []Payment `json:"data"`
}

// NewListPaymentsResponse instantiates a new ListPaymentsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListPaymentsResponse(data []Payment) *ListPaymentsResponse {
	this := ListPaymentsResponse{}
	this.Data = data
	return &this
}

// NewListPaymentsResponseWithDefaults instantiates a new ListPaymentsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListPaymentsResponseWithDefaults() *ListPaymentsResponse {
	this := ListPaymentsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ListPaymentsResponse) GetData() []Payment {
	if o == nil {
		var ret []Payment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ListPaymentsResponse) GetDataOk() ([]Payment, bool) {
	if o == nil  {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ListPaymentsResponse) SetData(v []Payment) {
	o.Data = v
}

func (o ListPaymentsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableListPaymentsResponse struct {
	value *ListPaymentsResponse
	isSet bool
}

func (v NullableListPaymentsResponse) Get() *ListPaymentsResponse {
	return v.value
}

func (v *NullableListPaymentsResponse) Set(val *ListPaymentsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableListPaymentsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableListPaymentsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableListPaymentsResponse(val *ListPaymentsResponse) *NullableListPaymentsResponse {
	return &NullableListPaymentsResponse{value: val, isSet: true}
}

func (v NullableListPaymentsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableListPaymentsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


