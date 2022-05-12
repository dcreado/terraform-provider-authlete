/*
Authlete API

Authlete API Document. 

API version: 2.2.19
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authlete

import (
	"encoding/json"
	"fmt"
)

// JweEnc This is the encryption algorithm to be used when encrypting a JWT on client or server side. Depending upon the context, this refers to encryption done by the client or by the server. For instance:   - as `authorizationEncryptionEnc` value, it refers to the encryption algorithm used by server when creating a JARM response   - as `requestEncryptionEnc` value, it refers to the expected encryption algorithm used by the client when encrypting a Request Object   - as `idTokenEncryptionEnc` value, it refers to the algorithm used by the server to encrypt id_tokens 
type JweEnc string

// List of jwe_enc
const (
	JWEENC_A128_CBC_HS256 JweEnc = "A128CBC_HS256"
	JWEENC_A192_CBC_HS384 JweEnc = "A192CBC_HS384"
	JWEENC_A256_CBC_HS512 JweEnc = "A256CBC_HS512"
	JWEENC_A128_GCM JweEnc = "A128GCM"
	JWEENC_A192_GCM JweEnc = "A192GCM"
	JWEENC_A256_GCM JweEnc = "A256GCM"
)

// All allowed values of JweEnc enum
var AllowedJweEncEnumValues = []JweEnc{
	"A128CBC_HS256",
	"A192CBC_HS384",
	"A256CBC_HS512",
	"A128GCM",
	"A192GCM",
	"A256GCM",
}

func (v *JweEnc) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JweEnc(value)
	for _, existing := range AllowedJweEncEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JweEnc", value)
}

// NewJweEncFromValue returns a pointer to a valid JweEnc
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJweEncFromValue(v string) (*JweEnc, error) {
	ev := JweEnc(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JweEnc: valid values are %v", v, AllowedJweEncEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JweEnc) IsValid() bool {
	for _, existing := range AllowedJweEncEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to jwe_enc value
func (v JweEnc) Ptr() *JweEnc {
	return &v
}

type NullableJweEnc struct {
	value *JweEnc
	isSet bool
}

func (v NullableJweEnc) Get() *JweEnc {
	return v.value
}

func (v *NullableJweEnc) Set(val *JweEnc) {
	v.value = val
	v.isSet = true
}

func (v NullableJweEnc) IsSet() bool {
	return v.isSet
}

func (v *NullableJweEnc) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJweEnc(val *JweEnc) *NullableJweEnc {
	return &NullableJweEnc{value: val, isSet: true}
}

func (v NullableJweEnc) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJweEnc) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

