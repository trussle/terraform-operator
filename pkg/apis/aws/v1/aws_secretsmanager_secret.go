
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
	KmsKeyId	string	`json:"kms_key_id"`
	RotationLambdaArn	string	`json:"rotation_lambda_arn"`
	RotationRules	[]RotationRules	`json:"rotation_rules"`
	Tags	map[string]string	`json:"tags"`
	Description	string	`json:"description"`
	Name	string	`json:"name"`
	RecoveryWindowInDays	int	`json:"recovery_window_in_days"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSecretsmanagerSecretList is a list of AwsSecretsmanagerSecret resources
type AwsSecretsmanagerSecretList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSecretsmanagerSecret	`json:"items"`
}


// RotationRules is a RotationRules
type RotationRules struct {
	AutomaticallyAfterDays	int	`json:"automatically_after_days"`
}
