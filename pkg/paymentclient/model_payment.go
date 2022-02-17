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

// Payment struct for Payment
type Payment struct {
	Provider string `json:"provider"`
	Reference *string `json:"reference,omitempty"`
	Scheme *string `json:"scheme,omitempty"`
	Status string `json:"status"`
	Value PaymentDataValue `json:"value"`
	Date string `json:"date"`
	Raw *map[string]interface{} `json:"raw,omitempty"`
	Id string `json:"id"`
}

// NewPayment instantiates a new Payment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPayment(provider string, status string, value PaymentDataValue, date string, id string) *Payment {
	this := Payment{}
	this.Provider = provider
	this.Status = status
	this.Value = value
	this.Date = date
	this.Id = id
	return &this
}

// NewPaymentWithDefaults instantiates a new Payment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentWithDefaults() *Payment {
	this := Payment{}
	return &this
}

// GetProvider returns the Provider field value
func (o *Payment) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *Payment) GetProviderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *Payment) SetProvider(v string) {
	o.Provider = v
}

// GetReference returns the Reference field value if set, zero value otherwise.
func (o *Payment) GetReference() string {
	if o == nil || o.Reference == nil {
		var ret string
		return ret
	}
	return *o.Reference
}

// GetReferenceOk returns a tuple with the Reference field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetReferenceOk() (*string, bool) {
	if o == nil || o.Reference == nil {
		return nil, false
	}
	return o.Reference, true
}

// HasReference returns a boolean if a field has been set.
func (o *Payment) HasReference() bool {
	if o != nil && o.Reference != nil {
		return true
	}

	return false
}

// SetReference gets a reference to the given string and assigns it to the Reference field.
func (o *Payment) SetReference(v string) {
	o.Reference = &v
}

// GetScheme returns the Scheme field value if set, zero value otherwise.
func (o *Payment) GetScheme() string {
	if o == nil || o.Scheme == nil {
		var ret string
		return ret
	}
	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetSchemeOk() (*string, bool) {
	if o == nil || o.Scheme == nil {
		return nil, false
	}
	return o.Scheme, true
}

// HasScheme returns a boolean if a field has been set.
func (o *Payment) HasScheme() bool {
	if o != nil && o.Scheme != nil {
		return true
	}

	return false
}

// SetScheme gets a reference to the given string and assigns it to the Scheme field.
func (o *Payment) SetScheme(v string) {
	o.Scheme = &v
}

// GetStatus returns the Status field value
func (o *Payment) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Payment) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Payment) SetStatus(v string) {
	o.Status = v
}

// GetValue returns the Value field value
func (o *Payment) GetValue() PaymentDataValue {
	if o == nil {
		var ret PaymentDataValue
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Payment) GetValueOk() (*PaymentDataValue, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Payment) SetValue(v PaymentDataValue) {
	o.Value = v
}

// GetDate returns the Date field value
func (o *Payment) GetDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Date
}

// GetDateOk returns a tuple with the Date field value
// and a boolean to check if the value has been set.
func (o *Payment) GetDateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Date, true
}

// SetDate sets field value
func (o *Payment) SetDate(v string) {
	o.Date = v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *Payment) GetRaw() map[string]interface{} {
	if o == nil || o.Raw == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Payment) GetRawOk() (*map[string]interface{}, bool) {
	if o == nil || o.Raw == nil {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *Payment) HasRaw() bool {
	if o != nil && o.Raw != nil {
		return true
	}

	return false
}

// SetRaw gets a reference to the given map[string]interface{} and assigns it to the Raw field.
func (o *Payment) SetRaw(v map[string]interface{}) {
	o.Raw = &v
}

// GetId returns the Id field value
func (o *Payment) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Payment) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Payment) SetId(v string) {
	o.Id = v
}

func (o Payment) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if o.Reference != nil {
		toSerialize["reference"] = o.Reference
	}
	if o.Scheme != nil {
		toSerialize["scheme"] = o.Scheme
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["date"] = o.Date
	}
	if o.Raw != nil {
		toSerialize["raw"] = o.Raw
	}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullablePayment struct {
	value *Payment
	isSet bool
}

func (v NullablePayment) Get() *Payment {
	return v.value
}

func (v *NullablePayment) Set(val *Payment) {
	v.value = val
	v.isSet = true
}

func (v NullablePayment) IsSet() bool {
	return v.isSet
}

func (v *NullablePayment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePayment(val *Payment) *NullablePayment {
	return &NullablePayment{value: val, isSet: true}
}

func (v NullablePayment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePayment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

