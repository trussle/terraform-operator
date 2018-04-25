
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
	VisibleToAllUsers	bool	`json:"visible_to_all_users"`
	EbsRootVolumeSize	int	`json:"ebs_root_volume_size"`
	LogUri	string	`json:"log_uri"`
	Applications	string	`json:"applications"`
	Tags	map[string]interface{}	`json:"tags"`
	ServiceRole	string	`json:"service_role"`
	CustomAmiId	string	`json:"custom_ami_id"`
	InstanceGroup	string	`json:"instance_group"`
	SecurityConfiguration	string	`json:"security_configuration"`
	KerberosAttributes	[]interface{}	`json:"kerberos_attributes"`
	AutoscalingRole	string	`json:"autoscaling_role"`
	Name	string	`json:"name"`
	ReleaseLabel	string	`json:"release_label"`
	CoreInstanceCount	int	`json:"core_instance_count"`
	MasterInstanceType	string	`json:"master_instance_type"`
	Ec2Attributes	[]interface{}	`json:"ec2_attributes"`
	BootstrapAction	string	`json:"bootstrap_action"`
	Configurations	string	`json:"configurations"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrClusterList is a list of AwsEmrCluster resources
type AwsEmrClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrCluster	`json:"items"`
}

