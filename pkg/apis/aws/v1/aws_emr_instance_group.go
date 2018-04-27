
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrInstanceGroup describes a AwsEmrInstanceGroup resource
type AwsEmrInstanceGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEmrInstanceGroupSpec	`json:"spec"`
}


// AwsEmrInstanceGroupSpec is the spec for a AwsEmrInstanceGroup Resource
type AwsEmrInstanceGroupSpec struct {
	EbsOptimized	bool	`json:"ebs_optimized"`
	EbsConfig	AwsEmrInstanceGroupEbsConfig	`json:"ebs_config"`
	ClusterId	string	`json:"cluster_id"`
	InstanceType	string	`json:"instance_type"`
	InstanceCount	int	`json:"instance_count"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrInstanceGroupList is a list of AwsEmrInstanceGroup resources
type AwsEmrInstanceGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrInstanceGroup	`json:"items"`
}


// AwsEmrInstanceGroupEbsConfig is a AwsEmrInstanceGroupEbsConfig
type AwsEmrInstanceGroupEbsConfig struct {
	Iops	int	`json:"iops"`
	Size	int	`json:"size"`
	Type	string	`json:"type"`
	VolumesPerInstance	int	`json:"volumes_per_instance"`
}
