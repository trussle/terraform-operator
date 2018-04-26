
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
	UsernameAttributes	[]Generic	`json:"username_attributes"`
	SmsAuthenticationMessage	string	`json:"sms_authentication_message"`
	Tags	map[string]Generic	`json:"tags"`
	MfaConfiguration	string	`json:"mfa_configuration"`
	EmailConfiguration	[]Generic	`json:"email_configuration"`
	Name	string	`json:"name"`
	SmsConfiguration	[]Generic	`json:"sms_configuration"`
	AutoVerifiedAttributes	Generic	`json:"auto_verified_attributes"`
	DeviceConfiguration	[]Generic	`json:"device_configuration"`
	Schema	Generic	`json:"schema"`
	SmsVerificationMessage	string	`json:"sms_verification_message"`
	AliasAttributes	Generic	`json:"alias_attributes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolList is a list of AwsCognitoUserPool resources
type AwsCognitoUserPoolList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPool	`json:"items"`
}

