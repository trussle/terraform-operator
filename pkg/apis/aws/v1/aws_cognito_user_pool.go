
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
	MfaConfiguration	string	`json:"mfa_configuration"`
	Schema	string	`json:"schema"`
	SmsVerificationMessage	string	`json:"sms_verification_message"`
	UsernameAttributes	[]string	`json:"username_attributes"`
	AutoVerifiedAttributes	string	`json:"auto_verified_attributes"`
	DeviceConfiguration	[]VVdHtAkn	`json:"device_configuration"`
	EmailConfiguration	[]jEQxCqkD	`json:"email_configuration"`
	Name	string	`json:"name"`
	SmsAuthenticationMessage	string	`json:"sms_authentication_message"`
	SmsConfiguration	[]sXTOdPOR	`json:"sms_configuration"`
	AliasAttributes	string	`json:"alias_attributes"`
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCognitoUserPoolList is a list of AwsCognitoUserPool resources
type AwsCognitoUserPoolList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCognitoUserPool	`json:"items"`
}


// dQrIPjyQ is a dQrIPjyQ
type dQrIPjyQ struct {
	EmailMessage	string	`json:"email_message"`
	EmailSubject	string	`json:"email_subject"`
	SmsMessage	string	`json:"sms_message"`
}

// PexOUAbU is a PexOUAbU
type PexOUAbU struct {
	AllowAdminCreateUserOnly	bool	`json:"allow_admin_create_user_only"`
	InviteMessageTemplate	[]dQrIPjyQ	`json:"invite_message_template"`
	UnusedAccountValidityDays	int	`json:"unused_account_validity_days"`
}

// yQFStFub is a yQFStFub
type yQFStFub struct {
	PreAuthentication	string	`json:"pre_authentication"`
	PreSignUp	string	`json:"pre_sign_up"`
	VerifyAuthChallengeResponse	string	`json:"verify_auth_challenge_response"`
	UserMigration	string	`json:"user_migration"`
	CreateAuthChallenge	string	`json:"create_auth_challenge"`
	CustomMessage	string	`json:"custom_message"`
	DefineAuthChallenge	string	`json:"define_auth_challenge"`
	PostAuthentication	string	`json:"post_authentication"`
	PostConfirmation	string	`json:"post_confirmation"`
	PreTokenGeneration	string	`json:"pre_token_generation"`
}

// VVdHtAkn is a VVdHtAkn
type VVdHtAkn struct {
	ChallengeRequiredOnNewDevice	bool	`json:"challenge_required_on_new_device"`
	DeviceOnlyRememberedOnUserPrompt	bool	`json:"device_only_remembered_on_user_prompt"`
}

// jEQxCqkD is a jEQxCqkD
type jEQxCqkD struct {
	ReplyToEmailAddress	string	`json:"reply_to_email_address"`
	SourceArn	string	`json:"source_arn"`
}

// IfTGXeJt is a IfTGXeJt
type IfTGXeJt struct {
	DefaultEmailOption	string	`json:"default_email_option"`
}

// uncbfqQU is a uncbfqQU
type uncbfqQU struct {
	RequireSymbols	bool	`json:"require_symbols"`
	RequireUppercase	bool	`json:"require_uppercase"`
	MinimumLength	int	`json:"minimum_length"`
	RequireLowercase	bool	`json:"require_lowercase"`
	RequireNumbers	bool	`json:"require_numbers"`
}

// sXTOdPOR is a sXTOdPOR
type sXTOdPOR struct {
	ExternalId	string	`json:"external_id"`
	SnsCallerArn	string	`json:"sns_caller_arn"`
}
