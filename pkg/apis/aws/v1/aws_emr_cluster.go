
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEmrCluster struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEmrClusterSpec	`json:"spec"`
}

type AwsEmrClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrCluster	`json:"items"`
}

type AwsEmrClusterSpec struct {
	VisibleToAllUsers	bool	`json:"visible_to_all_users"`
	EbsRootVolumeSize	int	`json:"ebs_root_volume_size"`
	CoreInstanceCount	int	`json:"core_instance_count"`
	LogUri	string	`json:"log_uri"`
	Ec2Attributes	[]interface{}	`json:"ec2_attributes"`
	ServiceRole	string	`json:"service_role"`
	SecurityConfiguration	string	`json:"security_configuration"`
	AutoscalingRole	string	`json:"autoscaling_role"`
	MasterInstanceType	string	`json:"master_instance_type"`
	Applications	interface{}	`json:"applications"`
	KerberosAttributes	[]interface{}	`json:"kerberos_attributes"`
	Tags	map[string]interface{}	`json:"tags"`
	Name	string	`json:"name"`
	ReleaseLabel	string	`json:"release_label"`
	InstanceGroup	interface{}	`json:"instance_group"`
	BootstrapAction	interface{}	`json:"bootstrap_action"`
	Configurations	string	`json:"configurations"`
	CustomAmiId	string	`json:"custom_ami_id"`
}
