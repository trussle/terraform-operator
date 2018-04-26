
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Ec2Attributes	[]Ec2Attributes	`json:"ec2_attributes"`
	AutoscalingRole	string	`json:"autoscaling_role"`
	EbsRootVolumeSize	int	`json:"ebs_root_volume_size"`
	LogUri	string	`json:"log_uri"`
	InstanceGroup	InstanceGroup	`json:"instance_group"`
	ServiceRole	string	`json:"service_role"`
	SecurityConfiguration	string	`json:"security_configuration"`
	VisibleToAllUsers	bool	`json:"visible_to_all_users"`
	ReleaseLabel	string	`json:"release_label"`
	MasterInstanceType	string	`json:"master_instance_type"`
	Applications	string	`json:"applications"`
	Name	string	`json:"name"`
	KerberosAttributes	[]KerberosAttributes	`json:"kerberos_attributes"`
	BootstrapAction	BootstrapAction	`json:"bootstrap_action"`
	Tags	map[string]string	`json:"tags"`
	Configurations	string	`json:"configurations"`
	CustomAmiId	string	`json:"custom_ami_id"`
	CoreInstanceCount	int	`json:"core_instance_count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrClusterList is a list of AwsEmrCluster resources
type AwsEmrClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrCluster	`json:"items"`
}


// Ec2Attributes is a Ec2Attributes
type Ec2Attributes struct {
	AdditionalMasterSecurityGroups	string	`json:"additional_master_security_groups"`
	AdditionalSlaveSecurityGroups	string	`json:"additional_slave_security_groups"`
	EmrManagedMasterSecurityGroup	string	`json:"emr_managed_master_security_group"`
	EmrManagedSlaveSecurityGroup	string	`json:"emr_managed_slave_security_group"`
	InstanceProfile	string	`json:"instance_profile"`
	ServiceAccessSecurityGroup	string	`json:"service_access_security_group"`
	KeyName	string	`json:"key_name"`
	SubnetId	string	`json:"subnet_id"`
}

// EbsConfig is a EbsConfig
type EbsConfig struct {
	Type	string	`json:"type"`
	VolumesPerInstance	int	`json:"volumes_per_instance"`
	Iops	int	`json:"iops"`
	Size	int	`json:"size"`
}

// InstanceGroup is a InstanceGroup
type InstanceGroup struct {
	InstanceType	string	`json:"instance_type"`
	Name	string	`json:"name"`
	BidPrice	string	`json:"bid_price"`
	EbsConfig	EbsConfig	`json:"ebs_config"`
	InstanceCount	int	`json:"instance_count"`
	AutoscalingPolicy	string	`json:"autoscaling_policy"`
	InstanceRole	string	`json:"instance_role"`
}

// HadoopJarStep is a HadoopJarStep
type HadoopJarStep struct {
	Args	[]string	`json:"args"`
	Jar	string	`json:"jar"`
	MainClass	string	`json:"main_class"`
	Properties	map[string]string	`json:"properties"`
}

// Step is a Step
type Step struct {
	ActionOnFailure	string	`json:"action_on_failure"`
	HadoopJarStep	[]HadoopJarStep	`json:"hadoop_jar_step"`
	Name	string	`json:"name"`
}

// KerberosAttributes is a KerberosAttributes
type KerberosAttributes struct {
	KdcAdminPassword	string	`json:"kdc_admin_password"`
	Realm	string	`json:"realm"`
	AdDomainJoinPassword	string	`json:"ad_domain_join_password"`
	AdDomainJoinUser	string	`json:"ad_domain_join_user"`
	CrossRealmTrustPrincipalPassword	string	`json:"cross_realm_trust_principal_password"`
}

// BootstrapAction is a BootstrapAction
type BootstrapAction struct {
	Name	string	`json:"name"`
	Path	string	`json:"path"`
	Args	[]string	`json:"args"`
}
