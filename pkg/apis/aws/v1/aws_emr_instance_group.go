
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
	ClusterId	string	`json:"cluster_id"`
	InstanceType	string	`json:"instance_type"`
	InstanceCount	int	`json:"instance_count"`
	Name	string	`json:"name"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	EbsConfig	EbsConfig	`json:"ebs_config"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrInstanceGroupList is a list of AwsEmrInstanceGroup resources
type AwsEmrInstanceGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrInstanceGroup	`json:"items"`
}


// EbsConfig is a EbsConfig
type EbsConfig struct {
	Type	string	`json:"type"`
	VolumesPerInstance	int	`json:"volumes_per_instance"`
	Iops	int	`json:"iops"`
	Size	int	`json:"size"`
}
