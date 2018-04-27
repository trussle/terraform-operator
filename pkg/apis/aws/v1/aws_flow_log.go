
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsFlowLog describes a AwsFlowLog resource
type AwsFlowLog struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsFlowLogSpec	`json:"spec"`
}


// AwsFlowLogSpec is the spec for a AwsFlowLog Resource
type AwsFlowLogSpec struct {
	IamRoleArn	string	`json:"iam_role_arn"`
	LogGroupName	string	`json:"log_group_name"`
	VpcId	string	`json:"vpc_id"`
	SubnetId	string	`json:"subnet_id"`
	EniId	string	`json:"eni_id"`
	TrafficType	string	`json:"traffic_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsFlowLogList is a list of AwsFlowLog resources
type AwsFlowLogList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsFlowLog	`json:"items"`
}

