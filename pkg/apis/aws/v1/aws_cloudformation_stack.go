
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	TemplateUrl	string	`json:"template_url"`
	Capabilities	string	`json:"capabilities"`
	PolicyUrl	string	`json:"policy_url"`
	Name	string	`json:"name"`
	NotificationArns	string	`json:"notification_arns"`
	OnFailure	string	`json:"on_failure"`
	TimeoutInMinutes	int	`json:"timeout_in_minutes"`
	Tags	map[string]string	`json:"tags"`
	IamRoleArn	string	`json:"iam_role_arn"`
	DisableRollback	bool	`json:"disable_rollback"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudformationStackList is a list of AwsCloudformationStack resources
type AwsCloudformationStackList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudformationStack	`json:"items"`
}

