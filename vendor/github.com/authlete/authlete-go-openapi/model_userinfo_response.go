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

// UserinfoResponse struct for UserinfoResponse
type UserinfoResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The list of claims that the client application requests to be embedded in the ID token. 
	Claims []string `json:"claims,omitempty"`
	// The ID of the client application which is associated with the access token. 
	ClientId *string `json:"clientId,omitempty"`
	// The client ID alias when the authorization request for the access token was made. 
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// The flag which indicates whether the client ID alias was used when the authorization request for the access token was made. 
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The content that the authorization server implementation can use as the value of `WWW-Authenticate` header on errors. 
	ResponseContent *string `json:"responseContent,omitempty"`
	// The scopes covered by the access token. 
	Scopes []string `json:"scopes,omitempty"`
	// The subject (= resource owner's ID). 
	Subject *string `json:"subject,omitempty"`
	// The access token that came along with the userinfo request. 
	Token *string `json:"token,omitempty"`
	// The extra properties associated with the access token. 
	Properties []Property `json:"properties,omitempty"`
	// The value of the `userinfo` property in the `claims` request parameter or in the `claims` property in an authorization request object.  A client application may request certain claims be embedded in an ID token or in a response from the userInfo endpoint. There are several ways. Including the `claims` request parameter and including the `claims` property in a request object are such examples. In both cases, the value of the `claims` parameter/property is JSON. Its format is described in [5.5. Requesting Claims using the \"claims\" Request Parameter](https://openid.net/specs/openid-connect-core-1_0.html#ClaimsParameter).  The following is an excerpt from the specification. You can find `userinfo` and `id_token` are top-level properties.  ```json {   \"userinfo\":   {     \"given_name\": { \"essential\": true },     \"nickname\": null,     \"email\": { \"essential\": true },     \"email_verified\": { \"essential\": true },     \"picture\": null,     \"http://example.info/claims/groups\": null   },   \"id_token\":   {     \"auth_time\": { \"essential\": true },     \"acr\": { \"values\": [ \"urn:mace:incommon:iap:silver\" ] }   } } ````  The value of this property is the value of the `userinfo` property in JSON format. For example, if the JSON above is included in an authorization request, this property holds JSON equivalent to the following.  ```json {   \"given_name\": { \"essential\": true },   \"nickname\": null,   \"email\": { \"essential\": true },   \"email_verified\": { \"essential\": true },   \"picture\": null,   \"http://example.info/claims/groups\": null } ```  Note that if a request object is given and it contains the `claims` property and if the `claims` request parameter is also given, the value of this property holds the former value. 
	UserInfoClaims *string `json:"userInfoClaims,omitempty"`
	// The attributes of this service that the client application belongs to. 
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client. 
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
}

// NewUserinfoResponse instantiates a new UserinfoResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserinfoResponse() *UserinfoResponse {
	this := UserinfoResponse{}
	return &this
}

// NewUserinfoResponseWithDefaults instantiates a new UserinfoResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserinfoResponseWithDefaults() *UserinfoResponse {
	this := UserinfoResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *UserinfoResponse) GetResultCode() string {
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *UserinfoResponse) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *UserinfoResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *UserinfoResponse) GetResultMessage() string {
	if o == nil || o.ResultMessage == nil {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || o.ResultMessage == nil {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *UserinfoResponse) HasResultMessage() bool {
	if o != nil && o.ResultMessage != nil {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *UserinfoResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *UserinfoResponse) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *UserinfoResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *UserinfoResponse) SetAction(v string) {
	o.Action = &v
}

// GetClaims returns the Claims field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClaims() []string {
	if o == nil || o.Claims == nil {
		var ret []string
		return ret
	}
	return o.Claims
}

// GetClaimsOk returns a tuple with the Claims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClaimsOk() ([]string, bool) {
	if o == nil || o.Claims == nil {
		return nil, false
	}
	return o.Claims, true
}

// HasClaims returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClaims() bool {
	if o != nil && o.Claims != nil {
		return true
	}

	return false
}

// SetClaims gets a reference to the given []string and assigns it to the Claims field.
func (o *UserinfoResponse) SetClaims(v []string) {
	o.Claims = v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientId() string {
	if o == nil || o.ClientId == nil {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientIdOk() (*string, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *UserinfoResponse) SetClientId(v string) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientIdAlias() string {
	if o == nil || o.ClientIdAlias == nil {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || o.ClientIdAlias == nil {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientIdAlias() bool {
	if o != nil && o.ClientIdAlias != nil {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *UserinfoResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientIdAliasUsed() bool {
	if o == nil || o.ClientIdAliasUsed == nil {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || o.ClientIdAliasUsed == nil {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientIdAliasUsed() bool {
	if o != nil && o.ClientIdAliasUsed != nil {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *UserinfoResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *UserinfoResponse) GetResponseContent() string {
	if o == nil || o.ResponseContent == nil {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || o.ResponseContent == nil {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *UserinfoResponse) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *UserinfoResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *UserinfoResponse) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *UserinfoResponse) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *UserinfoResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *UserinfoResponse) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *UserinfoResponse) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *UserinfoResponse) SetSubject(v string) {
	o.Subject = &v
}

// GetToken returns the Token field value if set, zero value otherwise.
func (o *UserinfoResponse) GetToken() string {
	if o == nil || o.Token == nil {
		var ret string
		return ret
	}
	return *o.Token
}

// GetTokenOk returns a tuple with the Token field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetTokenOk() (*string, bool) {
	if o == nil || o.Token == nil {
		return nil, false
	}
	return o.Token, true
}

// HasToken returns a boolean if a field has been set.
func (o *UserinfoResponse) HasToken() bool {
	if o != nil && o.Token != nil {
		return true
	}

	return false
}

// SetToken gets a reference to the given string and assigns it to the Token field.
func (o *UserinfoResponse) SetToken(v string) {
	o.Token = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *UserinfoResponse) GetProperties() []Property {
	if o == nil || o.Properties == nil {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetPropertiesOk() ([]Property, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *UserinfoResponse) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *UserinfoResponse) SetProperties(v []Property) {
	o.Properties = v
}

// GetUserInfoClaims returns the UserInfoClaims field value if set, zero value otherwise.
func (o *UserinfoResponse) GetUserInfoClaims() string {
	if o == nil || o.UserInfoClaims == nil {
		var ret string
		return ret
	}
	return *o.UserInfoClaims
}

// GetUserInfoClaimsOk returns a tuple with the UserInfoClaims field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetUserInfoClaimsOk() (*string, bool) {
	if o == nil || o.UserInfoClaims == nil {
		return nil, false
	}
	return o.UserInfoClaims, true
}

// HasUserInfoClaims returns a boolean if a field has been set.
func (o *UserinfoResponse) HasUserInfoClaims() bool {
	if o != nil && o.UserInfoClaims != nil {
		return true
	}

	return false
}

// SetUserInfoClaims gets a reference to the given string and assigns it to the UserInfoClaims field.
func (o *UserinfoResponse) SetUserInfoClaims(v string) {
	o.UserInfoClaims = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *UserinfoResponse) GetServiceAttributes() []Pair {
	if o == nil || o.ServiceAttributes == nil {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || o.ServiceAttributes == nil {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *UserinfoResponse) HasServiceAttributes() bool {
	if o != nil && o.ServiceAttributes != nil {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *UserinfoResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *UserinfoResponse) GetClientAttributes() []Pair {
	if o == nil || o.ClientAttributes == nil {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserinfoResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || o.ClientAttributes == nil {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *UserinfoResponse) HasClientAttributes() bool {
	if o != nil && o.ClientAttributes != nil {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *UserinfoResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

func (o UserinfoResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ResultCode != nil {
		toSerialize["resultCode"] = o.ResultCode
	}
	if o.ResultMessage != nil {
		toSerialize["resultMessage"] = o.ResultMessage
	}
	if o.Action != nil {
		toSerialize["action"] = o.Action
	}
	if o.Claims != nil {
		toSerialize["claims"] = o.Claims
	}
	if o.ClientId != nil {
		toSerialize["clientId"] = o.ClientId
	}
	if o.ClientIdAlias != nil {
		toSerialize["clientIdAlias"] = o.ClientIdAlias
	}
	if o.ClientIdAliasUsed != nil {
		toSerialize["clientIdAliasUsed"] = o.ClientIdAliasUsed
	}
	if o.ResponseContent != nil {
		toSerialize["responseContent"] = o.ResponseContent
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Token != nil {
		toSerialize["token"] = o.Token
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.UserInfoClaims != nil {
		toSerialize["userInfoClaims"] = o.UserInfoClaims
	}
	if o.ServiceAttributes != nil {
		toSerialize["serviceAttributes"] = o.ServiceAttributes
	}
	if o.ClientAttributes != nil {
		toSerialize["clientAttributes"] = o.ClientAttributes
	}
	return json.Marshal(toSerialize)
}

type NullableUserinfoResponse struct {
	value *UserinfoResponse
	isSet bool
}

func (v NullableUserinfoResponse) Get() *UserinfoResponse {
	return v.value
}

func (v *NullableUserinfoResponse) Set(val *UserinfoResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserinfoResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserinfoResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserinfoResponse(val *UserinfoResponse) *NullableUserinfoResponse {
	return &NullableUserinfoResponse{value: val, isSet: true}
}

func (v NullableUserinfoResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserinfoResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


