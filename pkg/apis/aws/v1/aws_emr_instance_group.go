
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrInstanceGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEmrInstanceGroupSpec	`json:"spec"`
}

type AwsEmrInstanceGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrInstanceGroup	`json:"items"`
}

type AwsEmrInstanceGroupSpec struct {
	InstanceCount	int	`json:"instance_count"`
	Name	string	`json:"name"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	EbsConfig	interface{}	`json:"ebs_config"`
	ClusterId	string	`json:"cluster_id"`
	InstanceType	string	`json:"instance_type"`
}
