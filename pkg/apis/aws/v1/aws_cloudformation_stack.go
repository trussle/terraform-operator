
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudformationStack describes a AwsCloudformationStack resource
type AwsCloudformationStack struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudformationStackSpec	`json:"spec"`
}


// AwsCloudformationStackSpec is the spec for a AwsCloudformationStack Resource
type AwsCloudformationStackSpec struct {
	Tags	map[string]Generic	`json:"tags"`
	OnFailure	string	`json:"on_failure"`
	Capabilities	Generic	`json:"capabilities"`
	DisableRollback	bool	`json:"disable_rollback"`
	TimeoutInMinutes	int	`json:"timeout_in_minutes"`
	Name	string	`json:"name"`
	TemplateUrl	string	`json:"template_url"`
	NotificationArns	Generic	`json:"notification_arns"`
	PolicyUrl	string	`json:"policy_url"`
	IamRoleArn	string	`json:"iam_role_arn"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudformationStackList is a list of AwsCloudformationStack resources
type AwsCloudformationStackList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudformationStack	`json:"items"`
}

