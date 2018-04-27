
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecretsmanagerSecret describes a AwsSecretsmanagerSecret resource
type AwsSecretsmanagerSecret struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSecretsmanagerSecretSpec	`json:"spec"`
}


// AwsSecretsmanagerSecretSpec is the spec for a AwsSecretsmanagerSecret Resource
type AwsSecretsmanagerSecretSpec struct {
	Name	string	`json:"name"`
	RecoveryWindowInDays	int	`json:"recovery_window_in_days"`
	RotationLambdaArn	string	`json:"rotation_lambda_arn"`
	Description	string	`json:"description"`
	KmsKeyId	string	`json:"kms_key_id"`
	RotationRules	[]AwsSecretsmanagerSecretRotationRules	`json:"rotation_rules"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecretsmanagerSecretList is a list of AwsSecretsmanagerSecret resources
type AwsSecretsmanagerSecretList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecretsmanagerSecret	`json:"items"`
}


// AwsSecretsmanagerSecretRotationRules is a AwsSecretsmanagerSecretRotationRules
type AwsSecretsmanagerSecretRotationRules struct {
	AutomaticallyAfterDays	int	`json:"automatically_after_days"`
}
