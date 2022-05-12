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

// DeviceAuthorizationResponse struct for DeviceAuthorizationResponse
type DeviceAuthorizationResponse struct {
	// The code which represents the result of the API call.
	ResultCode *string `json:"resultCode,omitempty"`
	// A short message which explains the result of the API call.
	ResultMessage *string `json:"resultMessage,omitempty"`
	// The next action that the authorization server implementation should take.
	Action *string `json:"action,omitempty"`
	// The content that the authorization server implementation is to return to the client application. Its format varies depending on the value of `action` parameter. 
	ResponseContent *string `json:"responseContent,omitempty"`
	// The client ID of the client application that has made the device authorization request. 
	ClientId *int64 `json:"clientId,omitempty"`
	// The client ID alias of the client application that has made the device authorization request. 
	ClientIdAlias *string `json:"clientIdAlias,omitempty"`
	// `true` if the value of the client_id request parameter included in the device authorization request is the client ID alias. `false` if the value is the original numeric client ID. 
	ClientIdAliasUsed *bool `json:"clientIdAliasUsed,omitempty"`
	// The name of the client application which has made the device authorization request. 
	ClientName *string `json:"clientName,omitempty"`
	// The client authentication method that should be performed at the device authorization endpoint. 
	ClientAuthMethod *string `json:"clientAuthMethod,omitempty"`
	// The scopes requested by the device authorization request.  Basically, this property holds the value of the scope request parameter in the device authorization request. However, because unregistered scopes are dropped on Authlete side, if the `scope` request parameter contains unknown scopes, the list returned by this property becomes different from the value of the `scope` request parameter.  Note that `description` property and `descriptions` property of each scope object in the array contained in this property is always `null` even if descriptions of the scopes are registered. 
	Scopes []string `json:"scopes,omitempty"`
	// The names of the claims which were requested indirectly via some special scopes. See [5.4. Requesting Claims using Scope Values](https://openid.net/specs/openid-connect-core-1_0.html#ScopeClaims) in OpenID Connect Core 1.0 for details. 
	ClaimNames []string `json:"claimNames,omitempty"`
	// The list of ACR values requested by the device authorization request.  Basically, this property holds the value of the `acr_values` request parameter in the device authorization request. However, because unsupported ACR values are dropped on Authlete side, if the `acr_values` request parameter contains unrecognized ACR values, the list returned by this property becomes different from the value of the `acr_values` request parameter. 
	Acrs []string `json:"acrs,omitempty"`
	// The device verification code. This corresponds to the `device_code` property in the response to the client. 
	DeviceCode *string `json:"deviceCode,omitempty"`
	// The end-user verification code. This corresponds to the `user_code` property in the response to the client. 
	UserCode *string `json:"userCode,omitempty"`
	// The end-user verification URI. This corresponds to the `verification_uri` property in the response to the client. 
	VerificationUri *string `json:"verificationUri,omitempty"`
	// The end-user verification URI that includes the end-user verification code. This corresponds to the `verification_uri_complete` property in the response to the client. 
	VerificationUriComplete *string `json:"verificationUriComplete,omitempty"`
	// The duration of the device verification code in seconds. This corresponds to the `expires_in` property in the response to the client. 
	ExpiresIn *int32 `json:"expiresIn,omitempty"`
	// The minimum amount of time in seconds that the client must wait for between polling requests to the token endpoint. This corresponds to the `interval` property in the response to the client. 
	Interval *int32 `json:"interval,omitempty"`
	// The warnings raised during processing the backchannel authentication request. 
	Warnings []string `json:"warnings,omitempty"`
	// The resources specified by the `resource` request parameters. See \"Resource Indicators for OAuth 2.0\" for details. 
	Resources []string `json:"resources,omitempty"`
	AuthorizationDetails *AuthorizationDetails `json:"authorizationDetails,omitempty"`
	// The attributes of this service that the client application belongs to. 
	ServiceAttributes []Pair `json:"serviceAttributes,omitempty"`
	// The attributes of the client. 
	ClientAttributes []Pair `json:"clientAttributes,omitempty"`
	// The dynamic scopes which the client application requested by the scope request parameter. 
	DynamicScopes []DynamicScope `json:"dynamicScopes,omitempty"`
}

// NewDeviceAuthorizationResponse instantiates a new DeviceAuthorizationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceAuthorizationResponse() *DeviceAuthorizationResponse {
	this := DeviceAuthorizationResponse{}
	return &this
}

// NewDeviceAuthorizationResponseWithDefaults instantiates a new DeviceAuthorizationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceAuthorizationResponseWithDefaults() *DeviceAuthorizationResponse {
	this := DeviceAuthorizationResponse{}
	return &this
}

// GetResultCode returns the ResultCode field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetResultCode() string {
	if o == nil || o.ResultCode == nil {
		var ret string
		return ret
	}
	return *o.ResultCode
}

// GetResultCodeOk returns a tuple with the ResultCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetResultCodeOk() (*string, bool) {
	if o == nil || o.ResultCode == nil {
		return nil, false
	}
	return o.ResultCode, true
}

// HasResultCode returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasResultCode() bool {
	if o != nil && o.ResultCode != nil {
		return true
	}

	return false
}

// SetResultCode gets a reference to the given string and assigns it to the ResultCode field.
func (o *DeviceAuthorizationResponse) SetResultCode(v string) {
	o.ResultCode = &v
}

// GetResultMessage returns the ResultMessage field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetResultMessage() string {
	if o == nil || o.ResultMessage == nil {
		var ret string
		return ret
	}
	return *o.ResultMessage
}

// GetResultMessageOk returns a tuple with the ResultMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetResultMessageOk() (*string, bool) {
	if o == nil || o.ResultMessage == nil {
		return nil, false
	}
	return o.ResultMessage, true
}

// HasResultMessage returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasResultMessage() bool {
	if o != nil && o.ResultMessage != nil {
		return true
	}

	return false
}

// SetResultMessage gets a reference to the given string and assigns it to the ResultMessage field.
func (o *DeviceAuthorizationResponse) SetResultMessage(v string) {
	o.ResultMessage = &v
}

// GetAction returns the Action field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetAction() string {
	if o == nil || o.Action == nil {
		var ret string
		return ret
	}
	return *o.Action
}

// GetActionOk returns a tuple with the Action field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetActionOk() (*string, bool) {
	if o == nil || o.Action == nil {
		return nil, false
	}
	return o.Action, true
}

// HasAction returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasAction() bool {
	if o != nil && o.Action != nil {
		return true
	}

	return false
}

// SetAction gets a reference to the given string and assigns it to the Action field.
func (o *DeviceAuthorizationResponse) SetAction(v string) {
	o.Action = &v
}

// GetResponseContent returns the ResponseContent field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetResponseContent() string {
	if o == nil || o.ResponseContent == nil {
		var ret string
		return ret
	}
	return *o.ResponseContent
}

// GetResponseContentOk returns a tuple with the ResponseContent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetResponseContentOk() (*string, bool) {
	if o == nil || o.ResponseContent == nil {
		return nil, false
	}
	return o.ResponseContent, true
}

// HasResponseContent returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasResponseContent() bool {
	if o != nil && o.ResponseContent != nil {
		return true
	}

	return false
}

// SetResponseContent gets a reference to the given string and assigns it to the ResponseContent field.
func (o *DeviceAuthorizationResponse) SetResponseContent(v string) {
	o.ResponseContent = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientId() int64 {
	if o == nil || o.ClientId == nil {
		var ret int64
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientIdOk() (*int64, bool) {
	if o == nil || o.ClientId == nil {
		return nil, false
	}
	return o.ClientId, true
}

// HasClientId returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientId() bool {
	if o != nil && o.ClientId != nil {
		return true
	}

	return false
}

// SetClientId gets a reference to the given int64 and assigns it to the ClientId field.
func (o *DeviceAuthorizationResponse) SetClientId(v int64) {
	o.ClientId = &v
}

// GetClientIdAlias returns the ClientIdAlias field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientIdAlias() string {
	if o == nil || o.ClientIdAlias == nil {
		var ret string
		return ret
	}
	return *o.ClientIdAlias
}

// GetClientIdAliasOk returns a tuple with the ClientIdAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientIdAliasOk() (*string, bool) {
	if o == nil || o.ClientIdAlias == nil {
		return nil, false
	}
	return o.ClientIdAlias, true
}

// HasClientIdAlias returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientIdAlias() bool {
	if o != nil && o.ClientIdAlias != nil {
		return true
	}

	return false
}

// SetClientIdAlias gets a reference to the given string and assigns it to the ClientIdAlias field.
func (o *DeviceAuthorizationResponse) SetClientIdAlias(v string) {
	o.ClientIdAlias = &v
}

// GetClientIdAliasUsed returns the ClientIdAliasUsed field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientIdAliasUsed() bool {
	if o == nil || o.ClientIdAliasUsed == nil {
		var ret bool
		return ret
	}
	return *o.ClientIdAliasUsed
}

// GetClientIdAliasUsedOk returns a tuple with the ClientIdAliasUsed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientIdAliasUsedOk() (*bool, bool) {
	if o == nil || o.ClientIdAliasUsed == nil {
		return nil, false
	}
	return o.ClientIdAliasUsed, true
}

// HasClientIdAliasUsed returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientIdAliasUsed() bool {
	if o != nil && o.ClientIdAliasUsed != nil {
		return true
	}

	return false
}

// SetClientIdAliasUsed gets a reference to the given bool and assigns it to the ClientIdAliasUsed field.
func (o *DeviceAuthorizationResponse) SetClientIdAliasUsed(v bool) {
	o.ClientIdAliasUsed = &v
}

// GetClientName returns the ClientName field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientName() string {
	if o == nil || o.ClientName == nil {
		var ret string
		return ret
	}
	return *o.ClientName
}

// GetClientNameOk returns a tuple with the ClientName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientNameOk() (*string, bool) {
	if o == nil || o.ClientName == nil {
		return nil, false
	}
	return o.ClientName, true
}

// HasClientName returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientName() bool {
	if o != nil && o.ClientName != nil {
		return true
	}

	return false
}

// SetClientName gets a reference to the given string and assigns it to the ClientName field.
func (o *DeviceAuthorizationResponse) SetClientName(v string) {
	o.ClientName = &v
}

// GetClientAuthMethod returns the ClientAuthMethod field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientAuthMethod() string {
	if o == nil || o.ClientAuthMethod == nil {
		var ret string
		return ret
	}
	return *o.ClientAuthMethod
}

// GetClientAuthMethodOk returns a tuple with the ClientAuthMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientAuthMethodOk() (*string, bool) {
	if o == nil || o.ClientAuthMethod == nil {
		return nil, false
	}
	return o.ClientAuthMethod, true
}

// HasClientAuthMethod returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientAuthMethod() bool {
	if o != nil && o.ClientAuthMethod != nil {
		return true
	}

	return false
}

// SetClientAuthMethod gets a reference to the given string and assigns it to the ClientAuthMethod field.
func (o *DeviceAuthorizationResponse) SetClientAuthMethod(v string) {
	o.ClientAuthMethod = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetScopes() []string {
	if o == nil || o.Scopes == nil {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetScopesOk() ([]string, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *DeviceAuthorizationResponse) SetScopes(v []string) {
	o.Scopes = v
}

// GetClaimNames returns the ClaimNames field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClaimNames() []string {
	if o == nil || o.ClaimNames == nil {
		var ret []string
		return ret
	}
	return o.ClaimNames
}

// GetClaimNamesOk returns a tuple with the ClaimNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClaimNamesOk() ([]string, bool) {
	if o == nil || o.ClaimNames == nil {
		return nil, false
	}
	return o.ClaimNames, true
}

// HasClaimNames returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClaimNames() bool {
	if o != nil && o.ClaimNames != nil {
		return true
	}

	return false
}

// SetClaimNames gets a reference to the given []string and assigns it to the ClaimNames field.
func (o *DeviceAuthorizationResponse) SetClaimNames(v []string) {
	o.ClaimNames = v
}

// GetAcrs returns the Acrs field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetAcrs() []string {
	if o == nil || o.Acrs == nil {
		var ret []string
		return ret
	}
	return o.Acrs
}

// GetAcrsOk returns a tuple with the Acrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetAcrsOk() ([]string, bool) {
	if o == nil || o.Acrs == nil {
		return nil, false
	}
	return o.Acrs, true
}

// HasAcrs returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasAcrs() bool {
	if o != nil && o.Acrs != nil {
		return true
	}

	return false
}

// SetAcrs gets a reference to the given []string and assigns it to the Acrs field.
func (o *DeviceAuthorizationResponse) SetAcrs(v []string) {
	o.Acrs = v
}

// GetDeviceCode returns the DeviceCode field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetDeviceCode() string {
	if o == nil || o.DeviceCode == nil {
		var ret string
		return ret
	}
	return *o.DeviceCode
}

// GetDeviceCodeOk returns a tuple with the DeviceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetDeviceCodeOk() (*string, bool) {
	if o == nil || o.DeviceCode == nil {
		return nil, false
	}
	return o.DeviceCode, true
}

// HasDeviceCode returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasDeviceCode() bool {
	if o != nil && o.DeviceCode != nil {
		return true
	}

	return false
}

// SetDeviceCode gets a reference to the given string and assigns it to the DeviceCode field.
func (o *DeviceAuthorizationResponse) SetDeviceCode(v string) {
	o.DeviceCode = &v
}

// GetUserCode returns the UserCode field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetUserCode() string {
	if o == nil || o.UserCode == nil {
		var ret string
		return ret
	}
	return *o.UserCode
}

// GetUserCodeOk returns a tuple with the UserCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetUserCodeOk() (*string, bool) {
	if o == nil || o.UserCode == nil {
		return nil, false
	}
	return o.UserCode, true
}

// HasUserCode returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasUserCode() bool {
	if o != nil && o.UserCode != nil {
		return true
	}

	return false
}

// SetUserCode gets a reference to the given string and assigns it to the UserCode field.
func (o *DeviceAuthorizationResponse) SetUserCode(v string) {
	o.UserCode = &v
}

// GetVerificationUri returns the VerificationUri field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetVerificationUri() string {
	if o == nil || o.VerificationUri == nil {
		var ret string
		return ret
	}
	return *o.VerificationUri
}

// GetVerificationUriOk returns a tuple with the VerificationUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetVerificationUriOk() (*string, bool) {
	if o == nil || o.VerificationUri == nil {
		return nil, false
	}
	return o.VerificationUri, true
}

// HasVerificationUri returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasVerificationUri() bool {
	if o != nil && o.VerificationUri != nil {
		return true
	}

	return false
}

// SetVerificationUri gets a reference to the given string and assigns it to the VerificationUri field.
func (o *DeviceAuthorizationResponse) SetVerificationUri(v string) {
	o.VerificationUri = &v
}

// GetVerificationUriComplete returns the VerificationUriComplete field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetVerificationUriComplete() string {
	if o == nil || o.VerificationUriComplete == nil {
		var ret string
		return ret
	}
	return *o.VerificationUriComplete
}

// GetVerificationUriCompleteOk returns a tuple with the VerificationUriComplete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetVerificationUriCompleteOk() (*string, bool) {
	if o == nil || o.VerificationUriComplete == nil {
		return nil, false
	}
	return o.VerificationUriComplete, true
}

// HasVerificationUriComplete returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasVerificationUriComplete() bool {
	if o != nil && o.VerificationUriComplete != nil {
		return true
	}

	return false
}

// SetVerificationUriComplete gets a reference to the given string and assigns it to the VerificationUriComplete field.
func (o *DeviceAuthorizationResponse) SetVerificationUriComplete(v string) {
	o.VerificationUriComplete = &v
}

// GetExpiresIn returns the ExpiresIn field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetExpiresIn() int32 {
	if o == nil || o.ExpiresIn == nil {
		var ret int32
		return ret
	}
	return *o.ExpiresIn
}

// GetExpiresInOk returns a tuple with the ExpiresIn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetExpiresInOk() (*int32, bool) {
	if o == nil || o.ExpiresIn == nil {
		return nil, false
	}
	return o.ExpiresIn, true
}

// HasExpiresIn returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasExpiresIn() bool {
	if o != nil && o.ExpiresIn != nil {
		return true
	}

	return false
}

// SetExpiresIn gets a reference to the given int32 and assigns it to the ExpiresIn field.
func (o *DeviceAuthorizationResponse) SetExpiresIn(v int32) {
	o.ExpiresIn = &v
}

// GetInterval returns the Interval field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetInterval() int32 {
	if o == nil || o.Interval == nil {
		var ret int32
		return ret
	}
	return *o.Interval
}

// GetIntervalOk returns a tuple with the Interval field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetIntervalOk() (*int32, bool) {
	if o == nil || o.Interval == nil {
		return nil, false
	}
	return o.Interval, true
}

// HasInterval returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasInterval() bool {
	if o != nil && o.Interval != nil {
		return true
	}

	return false
}

// SetInterval gets a reference to the given int32 and assigns it to the Interval field.
func (o *DeviceAuthorizationResponse) SetInterval(v int32) {
	o.Interval = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetWarnings() []string {
	if o == nil || o.Warnings == nil {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetWarningsOk() ([]string, bool) {
	if o == nil || o.Warnings == nil {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasWarnings() bool {
	if o != nil && o.Warnings != nil {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *DeviceAuthorizationResponse) SetWarnings(v []string) {
	o.Warnings = v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetResources() []string {
	if o == nil || o.Resources == nil {
		var ret []string
		return ret
	}
	return o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetResourcesOk() ([]string, bool) {
	if o == nil || o.Resources == nil {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasResources() bool {
	if o != nil && o.Resources != nil {
		return true
	}

	return false
}

// SetResources gets a reference to the given []string and assigns it to the Resources field.
func (o *DeviceAuthorizationResponse) SetResources(v []string) {
	o.Resources = v
}

// GetAuthorizationDetails returns the AuthorizationDetails field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetAuthorizationDetails() AuthorizationDetails {
	if o == nil || o.AuthorizationDetails == nil {
		var ret AuthorizationDetails
		return ret
	}
	return *o.AuthorizationDetails
}

// GetAuthorizationDetailsOk returns a tuple with the AuthorizationDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetAuthorizationDetailsOk() (*AuthorizationDetails, bool) {
	if o == nil || o.AuthorizationDetails == nil {
		return nil, false
	}
	return o.AuthorizationDetails, true
}

// HasAuthorizationDetails returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasAuthorizationDetails() bool {
	if o != nil && o.AuthorizationDetails != nil {
		return true
	}

	return false
}

// SetAuthorizationDetails gets a reference to the given AuthorizationDetails and assigns it to the AuthorizationDetails field.
func (o *DeviceAuthorizationResponse) SetAuthorizationDetails(v AuthorizationDetails) {
	o.AuthorizationDetails = &v
}

// GetServiceAttributes returns the ServiceAttributes field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetServiceAttributes() []Pair {
	if o == nil || o.ServiceAttributes == nil {
		var ret []Pair
		return ret
	}
	return o.ServiceAttributes
}

// GetServiceAttributesOk returns a tuple with the ServiceAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetServiceAttributesOk() ([]Pair, bool) {
	if o == nil || o.ServiceAttributes == nil {
		return nil, false
	}
	return o.ServiceAttributes, true
}

// HasServiceAttributes returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasServiceAttributes() bool {
	if o != nil && o.ServiceAttributes != nil {
		return true
	}

	return false
}

// SetServiceAttributes gets a reference to the given []Pair and assigns it to the ServiceAttributes field.
func (o *DeviceAuthorizationResponse) SetServiceAttributes(v []Pair) {
	o.ServiceAttributes = v
}

// GetClientAttributes returns the ClientAttributes field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetClientAttributes() []Pair {
	if o == nil || o.ClientAttributes == nil {
		var ret []Pair
		return ret
	}
	return o.ClientAttributes
}

// GetClientAttributesOk returns a tuple with the ClientAttributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetClientAttributesOk() ([]Pair, bool) {
	if o == nil || o.ClientAttributes == nil {
		return nil, false
	}
	return o.ClientAttributes, true
}

// HasClientAttributes returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasClientAttributes() bool {
	if o != nil && o.ClientAttributes != nil {
		return true
	}

	return false
}

// SetClientAttributes gets a reference to the given []Pair and assigns it to the ClientAttributes field.
func (o *DeviceAuthorizationResponse) SetClientAttributes(v []Pair) {
	o.ClientAttributes = v
}

// GetDynamicScopes returns the DynamicScopes field value if set, zero value otherwise.
func (o *DeviceAuthorizationResponse) GetDynamicScopes() []DynamicScope {
	if o == nil || o.DynamicScopes == nil {
		var ret []DynamicScope
		return ret
	}
	return o.DynamicScopes
}

// GetDynamicScopesOk returns a tuple with the DynamicScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceAuthorizationResponse) GetDynamicScopesOk() ([]DynamicScope, bool) {
	if o == nil || o.DynamicScopes == nil {
		return nil, false
	}
	return o.DynamicScopes, true
}

// HasDynamicScopes returns a boolean if a field has been set.
func (o *DeviceAuthorizationResponse) HasDynamicScopes() bool {
	if o != nil && o.DynamicScopes != nil {
		return true
	}

	return false
}

// SetDynamicScopes gets a reference to the given []DynamicScope and assigns it to the DynamicScopes field.
func (o *DeviceAuthorizationResponse) SetDynamicScopes(v []DynamicScope) {
	o.DynamicScopes = v
}

func (o DeviceAuthorizationResponse) MarshalJSON() ([]byte, error) {
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
	if o.ResponseContent != nil {
		toSerialize["responseContent"] = o.ResponseContent
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
	if o.ClientName != nil {
		toSerialize["clientName"] = o.ClientName
	}
	if o.ClientAuthMethod != nil {
		toSerialize["clientAuthMethod"] = o.ClientAuthMethod
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.ClaimNames != nil {
		toSerialize["claimNames"] = o.ClaimNames
	}
	if o.Acrs != nil {
		toSerialize["acrs"] = o.Acrs
	}
	if o.DeviceCode != nil {
		toSerialize["deviceCode"] = o.DeviceCode
	}
	if o.UserCode != nil {
		toSerialize["userCode"] = o.UserCode
	}
	if o.VerificationUri != nil {
		toSerialize["verificationUri"] = o.VerificationUri
	}
	if o.VerificationUriComplete != nil {
		toSerialize["verificationUriComplete"] = o.VerificationUriComplete
	}
	if o.ExpiresIn != nil {
		toSerialize["expiresIn"] = o.ExpiresIn
	}
	if o.Interval != nil {
		toSerialize["interval"] = o.Interval
	}
	if o.Warnings != nil {
		toSerialize["warnings"] = o.Warnings
	}
	if o.Resources != nil {
		toSerialize["resources"] = o.Resources
	}
	if o.AuthorizationDetails != nil {
		toSerialize["authorizationDetails"] = o.AuthorizationDetails
	}
	if o.ServiceAttributes != nil {
		toSerialize["serviceAttributes"] = o.ServiceAttributes
	}
	if o.ClientAttributes != nil {
		toSerialize["clientAttributes"] = o.ClientAttributes
	}
	if o.DynamicScopes != nil {
		toSerialize["dynamicScopes"] = o.DynamicScopes
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceAuthorizationResponse struct {
	value *DeviceAuthorizationResponse
	isSet bool
}

func (v NullableDeviceAuthorizationResponse) Get() *DeviceAuthorizationResponse {
	return v.value
}

func (v *NullableDeviceAuthorizationResponse) Set(val *DeviceAuthorizationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceAuthorizationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceAuthorizationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceAuthorizationResponse(val *DeviceAuthorizationResponse) *NullableDeviceAuthorizationResponse {
	return &NullableDeviceAuthorizationResponse{value: val, isSet: true}
}

func (v NullableDeviceAuthorizationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceAuthorizationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


