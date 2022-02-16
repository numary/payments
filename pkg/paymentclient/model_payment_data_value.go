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

// PaymentDataValue struct for PaymentDataValue
type PaymentDataValue struct {
	Amount int32 `json:"amount"`
	Asset string `json:"asset"`
}

// NewPaymentDataValue instantiates a new PaymentDataValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentDataValue(amount int32, asset string) *PaymentDataValue {
	this := PaymentDataValue{}
	this.Amount = amount
	this.Asset = asset
	return &this
}

// NewPaymentDataValueWithDefaults instantiates a new PaymentDataValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentDataValueWithDefaults() *PaymentDataValue {
	this := PaymentDataValue{}
	return &this
}

// GetAmount returns the Amount field value
func (o *PaymentDataValue) GetAmount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Amount
}

// GetAmountOk returns a tuple with the Amount field value
// and a boolean to check if the value has been set.
func (o *PaymentDataValue) GetAmountOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Amount, true
}

// SetAmount sets field value
func (o *PaymentDataValue) SetAmount(v int32) {
	o.Amount = v
}

// GetAsset returns the Asset field value
func (o *PaymentDataValue) GetAsset() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Asset
}

// GetAssetOk returns a tuple with the Asset field value
// and a boolean to check if the value has been set.
func (o *PaymentDataValue) GetAssetOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Asset, true
}

// SetAsset sets field value
func (o *PaymentDataValue) SetAsset(v string) {
	o.Asset = v
}

func (o PaymentDataValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["amount"] = o.Amount
	}
	if true {
		toSerialize["asset"] = o.Asset
	}
	return json.Marshal(toSerialize)
}

type NullablePaymentDataValue struct {
	value *PaymentDataValue
	isSet bool
}

func (v NullablePaymentDataValue) Get() *PaymentDataValue {
	return v.value
}

func (v *NullablePaymentDataValue) Set(val *PaymentDataValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentDataValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentDataValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentDataValue(val *PaymentDataValue) *NullablePaymentDataValue {
	return &NullablePaymentDataValue{value: val, isSet: true}
}

func (v NullablePaymentDataValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentDataValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


