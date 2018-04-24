
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsFlowLog struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsFlowLogSpec	`json:"spec"`
}

type AwsFlowLogList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsFlowLog	`json:"items"`
}

type AwsFlowLogSpec struct {
	LogGroupName	string	`json:"log_group_name"`
	VpcId	string	`json:"vpc_id"`
	SubnetId	string	`json:"subnet_id"`
	EniId	string	`json:"eni_id"`
	TrafficType	string	`json:"traffic_type"`
	IamRoleArn	string	`json:"iam_role_arn"`
}
