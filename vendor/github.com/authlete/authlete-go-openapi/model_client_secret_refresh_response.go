/*
Authlete API

Authlete API Document. 

API version: 2.2.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
)

// ClientSecretRefreshResponse struct for ClientSecretRefreshResponse
type ClientSecretRefreshResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The new client secret. 
	NewClientSecret *string `json:"newClientSecret,omitempty"`
	// The old client secret. 
	OldClientSecret *string `json:"oldClientSecret,omitempty"`
}

// NewClientSecretRefreshResponse instantiates a new ClientSecretRefreshResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClientSecretRefreshResponse() *ClientSecretRefreshResponse {
	this := ClientSecretRefreshResponse{}
	return &this
}

// NewClientSecretRefreshResponseWithDefaults instantiates a new ClientSecretRefreshResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClientSecretRefreshResponseWithDefaults() *ClientSecretRefreshResponse {
	this := ClientSecretRefreshResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *ClientSecretRefreshResponse) GetResultCode() string {
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientSecretRefreshResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *ClientSecretRefreshResponse) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *ClientSecretRefreshResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *ClientSecretRefreshResponse) GetResultMessage() string {
	if o == nil || o.ResultMessage == nil {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientSecretRefreshResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || o.ResultMessage == nil {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *ClientSecretRefreshResponse) HasResultMessage() bool {
	if o != nil && o.ResultMessage != nil {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *ClientSecretRefreshResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetNewClientSecret returns the NewClientSecret field value if set, zero value otherwise.
func (o *ClientSecretRefreshResponse) GetNewClientSecret() string {
	if o == nil || o.NewClientSecret == nil {
		var ret string
		return ret
	}
	return *o.NewClientSecret
}

// GetNewClientSecretOk returns a tuple with the NewClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientSecretRefreshResponse) GetNewClientSecretOk() (*string, bool) {
	if o == nil || o.NewClientSecret == nil {
		return nil, false
	}
	return o.NewClientSecret, true
}

// HasNewClientSecret returns a boolean if a field has been set.
func (o *ClientSecretRefreshResponse) HasNewClientSecret() bool {
	if o != nil && o.NewClientSecret != nil {
		return true
	}

	return false
}

// SetNewClientSecret gets a reference to the given string and assigns it to the NewClientSecret field.
func (o *ClientSecretRefreshResponse) SetNewClientSecret(v string) {
	o.NewClientSecret = &v
}

// GetOldClientSecret returns the OldClientSecret field value if set, zero value otherwise.
func (o *ClientSecretRefreshResponse) GetOldClientSecret() string {
	if o == nil || o.OldClientSecret == nil {
		var ret string
		return ret
	}
	return *o.OldClientSecret
}

// GetOldClientSecretOk returns a tuple with the OldClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ClientSecretRefreshResponse) GetOldClientSecretOk() (*string, bool) {
	if o == nil || o.OldClientSecret == nil {
		return nil, false
	}
	return o.OldClientSecret, true
}

// HasOldClientSecret returns a boolean if a field has been set.
func (o *ClientSecretRefreshResponse) HasOldClientSecret() bool {
	if o != nil && o.OldClientSecret != nil {
		return true
	}

	return false
}

// SetOldClientSecret gets a reference to the given string and assigns it to the OldClientSecret field.
func (o *ClientSecretRefreshResponse) SetOldClientSecret(v string) {
	o.OldClientSecret = &v
}

func (o ClientSecretRefreshResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResultCode != nil {
		toSerialize["resultCode"] = o.ResultCode
	}
	if o.ResultMessage != nil {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if o.NewClientSecret != nil {
		toSerialize["newClientSecret"] = o.NewClientSecret
	}
	if o.OldClientSecret != nil {
		toSerialize["oldClientSecret"] = o.OldClientSecret
	}
	return json.Marshal(toSerialize)
}

type NullableClientSecretRefreshResponse struct {
	value *ClientSecretRefreshResponse
	isSet bool
}

func (v NullableClientSecretRefreshResponse) Get() *ClientSecretRefreshResponse {
	return v.value
}

func (v *NullableClientSecretRefreshResponse) Set(val *ClientSecretRefreshResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableClientSecretRefreshResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableClientSecretRefreshResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClientSecretRefreshResponse(val *ClientSecretRefreshResponse) *NullableClientSecretRefreshResponse {
	return &NullableClientSecretRefreshResponse{value: val, isSet: true}
}

func (v NullableClientSecretRefreshResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClientSecretRefreshResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


