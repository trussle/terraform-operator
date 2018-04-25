
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPool describes a AwsCognitoUserPool resource
type AwsCognitoUserPool struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCognitoUserPoolSpec	`json:"spec"`
}


// AwsCognitoUserPoolSpec is the spec for a AwsCognitoUserPool Resource
type AwsCognitoUserPoolSpec struct {
	UsernameAttributes	[]interface{}	`json:"username_attributes"`
	AliasAttributes	string	`json:"alias_attributes"`
	Schema	string	`json:"schema"`
	SmsVerificationMessage	string	`json:"sms_verification_message"`
	AutoVerifiedAttributes	string	`json:"auto_verified_attributes"`
	DeviceConfiguration	[]interface{}	`json:"device_configuration"`
	SmsConfiguration	[]interface{}	`json:"sms_configuration"`
	Tags	map[string]interface{}	`json:"tags"`
	MfaConfiguration	string	`json:"mfa_configuration"`
	EmailConfiguration	[]interface{}	`json:"email_configuration"`
	Name	string	`json:"name"`
	SmsAuthenticationMessage	string	`json:"sms_authentication_message"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolList is a list of AwsCognitoUserPool resources
type AwsCognitoUserPoolList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPool	`json:"items"`
}

