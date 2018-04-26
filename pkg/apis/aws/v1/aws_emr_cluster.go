
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrCluster describes a AwsEmrCluster resource
type AwsEmrCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEmrClusterSpec	`json:"spec"`
}


// AwsEmrClusterSpec is the spec for a AwsEmrCluster Resource
type AwsEmrClusterSpec struct {
	InstanceGroup	Generic	`json:"instance_group"`
	BootstrapAction	Generic	`json:"bootstrap_action"`
	Tags	map[string]Generic	`json:"tags"`
	ServiceRole	string	`json:"service_role"`
	SecurityConfiguration	string	`json:"security_configuration"`
	Applications	Generic	`json:"applications"`
	VisibleToAllUsers	bool	`json:"visible_to_all_users"`
	EbsRootVolumeSize	int	`json:"ebs_root_volume_size"`
	CustomAmiId	string	`json:"custom_ami_id"`
	Name	string	`json:"name"`
	CoreInstanceCount	int	`json:"core_instance_count"`
	KerberosAttributes	[]Generic	`json:"kerberos_attributes"`
	Ec2Attributes	[]Generic	`json:"ec2_attributes"`
	Configurations	string	`json:"configurations"`
	MasterInstanceType	string	`json:"master_instance_type"`
	LogUri	string	`json:"log_uri"`
	ReleaseLabel	string	`json:"release_label"`
	AutoscalingRole	string	`json:"autoscaling_role"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrClusterList is a list of AwsEmrCluster resources
type AwsEmrClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrCluster	`json:"items"`
}

