
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloudformationStack struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudformationStackSpec	`json:"spec"`
}

type AwsCloudformationStackList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudformationStack	`json:"items"`
}

type AwsCloudformationStackSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	TemplateUrl	string	`json:"template_url"`
	OnFailure	string	`json:"on_failure"`
	DisableRollback	bool	`json:"disable_rollback"`
	Name	string	`json:"name"`
	PolicyUrl	string	`json:"policy_url"`
	IamRoleArn	string	`json:"iam_role_arn"`
	Capabilities	interface{}	`json:"capabilities"`
	NotificationArns	interface{}	`json:"notification_arns"`
	TimeoutInMinutes	int	`json:"timeout_in_minutes"`
}
