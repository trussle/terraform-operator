
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCognitoUserPool struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoUserPoolSpec	`json:"spec"`
}

type AwsCognitoUserPoolList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPool	`json:"items"`
}

type AwsCognitoUserPoolSpec struct {
	AutoVerifiedAttributes	interface{}	`json:"auto_verified_attributes"`
	SmsAuthenticationMessage	string	`json:"sms_authentication_message"`
	UsernameAttributes	[]interface{}	`json:"username_attributes"`
	AliasAttributes	interface{}	`json:"alias_attributes"`
	MfaConfiguration	string	`json:"mfa_configuration"`
	Name	string	`json:"name"`
	Schema	interface{}	`json:"schema"`
	SmsConfiguration	[]interface{}	`json:"sms_configuration"`
	Tags	map[string]interface{}	`json:"tags"`
	EmailConfiguration	[]interface{}	`json:"email_configuration"`
	SmsVerificationMessage	string	`json:"sms_verification_message"`
	DeviceConfiguration	[]interface{}	`json:"device_configuration"`
}
