
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
	Tags	map[string]string	`json:"tags"`
	EmailConfiguration	[]AwsCognitoUserPoolEmailConfiguration	`json:"email_configuration"`
	MfaConfiguration	string	`json:"mfa_configuration"`
	Name	string	`json:"name"`
	DeviceConfiguration	[]AwsCognitoUserPoolDeviceConfiguration	`json:"device_configuration"`
	SmsAuthenticationMessage	string	`json:"sms_authentication_message"`
	Schema	AwsCognitoUserPoolSchema	`json:"schema"`
	SmsConfiguration	[]AwsCognitoUserPoolSmsConfiguration	`json:"sms_configuration"`
	SmsVerificationMessage	string	`json:"sms_verification_message"`
	UsernameAttributes	[]string	`json:"username_attributes"`
	AliasAttributes	string	`json:"alias_attributes"`
	AutoVerifiedAttributes	string	`json:"auto_verified_attributes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolList is a list of AwsCognitoUserPool resources
type AwsCognitoUserPoolList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPool	`json:"items"`
}


// AwsCognitoUserPoolSmsConfiguration is a AwsCognitoUserPoolSmsConfiguration
type AwsCognitoUserPoolSmsConfiguration struct {
	ExternalId	string	`json:"external_id"`
	SnsCallerArn	string	`json:"sns_caller_arn"`
}

// AwsCognitoUserPoolLambdaConfig is a AwsCognitoUserPoolLambdaConfig
type AwsCognitoUserPoolLambdaConfig struct {
	CreateAuthChallenge	string	`json:"create_auth_challenge"`
	DefineAuthChallenge	string	`json:"define_auth_challenge"`
	PostAuthentication	string	`json:"post_authentication"`
	PreSignUp	string	`json:"pre_sign_up"`
	UserMigration	string	`json:"user_migration"`
	CustomMessage	string	`json:"custom_message"`
	PostConfirmation	string	`json:"post_confirmation"`
	PreAuthentication	string	`json:"pre_authentication"`
	PreTokenGeneration	string	`json:"pre_token_generation"`
	VerifyAuthChallengeResponse	string	`json:"verify_auth_challenge_response"`
}

// AwsCognitoUserPoolDeviceConfiguration is a AwsCognitoUserPoolDeviceConfiguration
type AwsCognitoUserPoolDeviceConfiguration struct {
	DeviceOnlyRememberedOnUserPrompt	bool	`json:"device_only_remembered_on_user_prompt"`
	ChallengeRequiredOnNewDevice	bool	`json:"challenge_required_on_new_device"`
}

// AwsCognitoUserPoolNumberAttributeConstraints is a AwsCognitoUserPoolNumberAttributeConstraints
type AwsCognitoUserPoolNumberAttributeConstraints struct {
	MaxValue	string	`json:"max_value"`
	MinValue	string	`json:"min_value"`
}

// AwsCognitoUserPoolSchema is a AwsCognitoUserPoolSchema
type AwsCognitoUserPoolSchema struct {
	NumberAttributeConstraints	[]AwsCognitoUserPoolNumberAttributeConstraints	`json:"number_attribute_constraints"`
	Required	bool	`json:"required"`
	StringAttributeConstraints	[]AwsCognitoUserPoolStringAttributeConstraints	`json:"string_attribute_constraints"`
	AttributeDataType	string	`json:"attribute_data_type"`
	DeveloperOnlyAttribute	bool	`json:"developer_only_attribute"`
	Mutable	bool	`json:"mutable"`
	Name	string	`json:"name"`
}

// AwsCognitoUserPoolPasswordPolicy is a AwsCognitoUserPoolPasswordPolicy
type AwsCognitoUserPoolPasswordPolicy struct {
	RequireLowercase	bool	`json:"require_lowercase"`
	RequireNumbers	bool	`json:"require_numbers"`
	RequireSymbols	bool	`json:"require_symbols"`
	RequireUppercase	bool	`json:"require_uppercase"`
	MinimumLength	int	`json:"minimum_length"`
}

// AwsCognitoUserPoolStringAttributeConstraints is a AwsCognitoUserPoolStringAttributeConstraints
type AwsCognitoUserPoolStringAttributeConstraints struct {
	MinLength	string	`json:"min_length"`
	MaxLength	string	`json:"max_length"`
}

// AwsCognitoUserPoolEmailConfiguration is a AwsCognitoUserPoolEmailConfiguration
type AwsCognitoUserPoolEmailConfiguration struct {
	ReplyToEmailAddress	string	`json:"reply_to_email_address"`
	SourceArn	string	`json:"source_arn"`
}

// AwsCognitoUserPoolVerificationMessageTemplate is a AwsCognitoUserPoolVerificationMessageTemplate
type AwsCognitoUserPoolVerificationMessageTemplate struct {
	DefaultEmailOption	string	`json:"default_email_option"`
}

// AwsCognitoUserPoolInviteMessageTemplate is a AwsCognitoUserPoolInviteMessageTemplate
type AwsCognitoUserPoolInviteMessageTemplate struct {
	EmailSubject	string	`json:"email_subject"`
	SmsMessage	string	`json:"sms_message"`
	EmailMessage	string	`json:"email_message"`
}

// AwsCognitoUserPoolAdminCreateUserConfig is a AwsCognitoUserPoolAdminCreateUserConfig
type AwsCognitoUserPoolAdminCreateUserConfig struct {
	AllowAdminCreateUserOnly	bool	`json:"allow_admin_create_user_only"`
	InviteMessageTemplate	[]AwsCognitoUserPoolInviteMessageTemplate	`json:"invite_message_template"`
	UnusedAccountValidityDays	int	`json:"unused_account_validity_days"`
}
