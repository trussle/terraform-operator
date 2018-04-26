
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	AutoVerifiedAttributes	string	`json:"auto_verified_attributes"`
	SmsVerificationMessage	string	`json:"sms_verification_message"`
	SmsAuthenticationMessage	string	`json:"sms_authentication_message"`
	EmailConfiguration	[]EmailConfiguration	`json:"email_configuration"`
	Name	string	`json:"name"`
	Schema	Schema	`json:"schema"`
	SmsConfiguration	[]SmsConfiguration	`json:"sms_configuration"`
	Tags	map[string]string	`json:"tags"`
	DeviceConfiguration	[]DeviceConfiguration	`json:"device_configuration"`
	MfaConfiguration	string	`json:"mfa_configuration"`
	UsernameAttributes	[]string	`json:"username_attributes"`
	AliasAttributes	string	`json:"alias_attributes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolList is a list of AwsCognitoUserPool resources
type AwsCognitoUserPoolList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPool	`json:"items"`
}


// LambdaConfig is a LambdaConfig
type LambdaConfig struct {
	CustomMessage	string	`json:"custom_message"`
	PreAuthentication	string	`json:"pre_authentication"`
	PreSignUp	string	`json:"pre_sign_up"`
	PreTokenGeneration	string	`json:"pre_token_generation"`
	VerifyAuthChallengeResponse	string	`json:"verify_auth_challenge_response"`
	CreateAuthChallenge	string	`json:"create_auth_challenge"`
	DefineAuthChallenge	string	`json:"define_auth_challenge"`
	PostAuthentication	string	`json:"post_authentication"`
	PostConfirmation	string	`json:"post_confirmation"`
	UserMigration	string	`json:"user_migration"`
}

// NumberAttributeConstraints is a NumberAttributeConstraints
type NumberAttributeConstraints struct {
	MaxValue	string	`json:"max_value"`
	MinValue	string	`json:"min_value"`
}

// Schema is a Schema
type Schema struct {
	StringAttributeConstraints	[]StringAttributeConstraints	`json:"string_attribute_constraints"`
	AttributeDataType	string	`json:"attribute_data_type"`
	DeveloperOnlyAttribute	bool	`json:"developer_only_attribute"`
	Mutable	bool	`json:"mutable"`
	Name	string	`json:"name"`
	NumberAttributeConstraints	[]NumberAttributeConstraints	`json:"number_attribute_constraints"`
	Required	bool	`json:"required"`
}

// SmsConfiguration is a SmsConfiguration
type SmsConfiguration struct {
	ExternalId	string	`json:"external_id"`
	SnsCallerArn	string	`json:"sns_caller_arn"`
}

// DeviceConfiguration is a DeviceConfiguration
type DeviceConfiguration struct {
	ChallengeRequiredOnNewDevice	bool	`json:"challenge_required_on_new_device"`
	DeviceOnlyRememberedOnUserPrompt	bool	`json:"device_only_remembered_on_user_prompt"`
}

// InviteMessageTemplate is a InviteMessageTemplate
type InviteMessageTemplate struct {
	EmailMessage	string	`json:"email_message"`
	EmailSubject	string	`json:"email_subject"`
	SmsMessage	string	`json:"sms_message"`
}

// EmailConfiguration is a EmailConfiguration
type EmailConfiguration struct {
	ReplyToEmailAddress	string	`json:"reply_to_email_address"`
	SourceArn	string	`json:"source_arn"`
}

// PasswordPolicy is a PasswordPolicy
type PasswordPolicy struct {
	RequireLowercase	bool	`json:"require_lowercase"`
	RequireNumbers	bool	`json:"require_numbers"`
	RequireSymbols	bool	`json:"require_symbols"`
	RequireUppercase	bool	`json:"require_uppercase"`
	MinimumLength	int	`json:"minimum_length"`
}

// StringAttributeConstraints is a StringAttributeConstraints
type StringAttributeConstraints struct {
	MinLength	string	`json:"min_length"`
	MaxLength	string	`json:"max_length"`
}

// VerificationMessageTemplate is a VerificationMessageTemplate
type VerificationMessageTemplate struct {
	DefaultEmailOption	string	`json:"default_email_option"`
}

// AdminCreateUserConfig is a AdminCreateUserConfig
type AdminCreateUserConfig struct {
	AllowAdminCreateUserOnly	bool	`json:"allow_admin_create_user_only"`
	InviteMessageTemplate	[]InviteMessageTemplate	`json:"invite_message_template"`
	UnusedAccountValidityDays	int	`json:"unused_account_validity_days"`
}
