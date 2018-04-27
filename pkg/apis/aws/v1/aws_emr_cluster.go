
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
	SecurityConfiguration	string	`json:"security_configuration"`
	MasterInstanceType	string	`json:"master_instance_type"`
	Tags	map[string]string	`json:"tags"`
	LogUri	string	`json:"log_uri"`
	ServiceRole	string	`json:"service_role"`
	Name	string	`json:"name"`
	ReleaseLabel	string	`json:"release_label"`
	CoreInstanceCount	int	`json:"core_instance_count"`
	InstanceGroup	AwsEmrClusterInstanceGroup	`json:"instance_group"`
	Configurations	string	`json:"configurations"`
	AutoscalingRole	string	`json:"autoscaling_role"`
	Applications	string	`json:"applications"`
	KerberosAttributes	[]AwsEmrClusterKerberosAttributes	`json:"kerberos_attributes"`
	BootstrapAction	AwsEmrClusterBootstrapAction	`json:"bootstrap_action"`
	VisibleToAllUsers	bool	`json:"visible_to_all_users"`
	EbsRootVolumeSize	int	`json:"ebs_root_volume_size"`
	CustomAmiId	string	`json:"custom_ami_id"`
	Ec2Attributes	[]AwsEmrClusterEc2Attributes	`json:"ec2_attributes"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrClusterList is a list of AwsEmrCluster resources
type AwsEmrClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrCluster	`json:"items"`
}


// AwsEmrClusterBootstrapAction is a AwsEmrClusterBootstrapAction
type AwsEmrClusterBootstrapAction struct {
	Name	string	`json:"name"`
	Path	string	`json:"path"`
	Args	[]string	`json:"args"`
}

// AwsEmrClusterEc2Attributes is a AwsEmrClusterEc2Attributes
type AwsEmrClusterEc2Attributes struct {
	KeyName	string	`json:"key_name"`
	SubnetId	string	`json:"subnet_id"`
	AdditionalMasterSecurityGroups	string	`json:"additional_master_security_groups"`
	AdditionalSlaveSecurityGroups	string	`json:"additional_slave_security_groups"`
	EmrManagedMasterSecurityGroup	string	`json:"emr_managed_master_security_group"`
	EmrManagedSlaveSecurityGroup	string	`json:"emr_managed_slave_security_group"`
	InstanceProfile	string	`json:"instance_profile"`
	ServiceAccessSecurityGroup	string	`json:"service_access_security_group"`
}

// AwsEmrClusterEbsConfig is a AwsEmrClusterEbsConfig
type AwsEmrClusterEbsConfig struct {
	Type	string	`json:"type"`
	VolumesPerInstance	int	`json:"volumes_per_instance"`
	Iops	int	`json:"iops"`
	Size	int	`json:"size"`
}

// AwsEmrClusterInstanceGroup is a AwsEmrClusterInstanceGroup
type AwsEmrClusterInstanceGroup struct {
	InstanceType	string	`json:"instance_type"`
	Name	string	`json:"name"`
	BidPrice	string	`json:"bid_price"`
	EbsConfig	AwsEmrClusterEbsConfig	`json:"ebs_config"`
	InstanceCount	int	`json:"instance_count"`
	AutoscalingPolicy	string	`json:"autoscaling_policy"`
	InstanceRole	string	`json:"instance_role"`
}

// AwsEmrClusterHadoopJarStep is a AwsEmrClusterHadoopJarStep
type AwsEmrClusterHadoopJarStep struct {
	MainClass	string	`json:"main_class"`
	Properties	map[string]string	`json:"properties"`
	Args	[]string	`json:"args"`
	Jar	string	`json:"jar"`
}

// AwsEmrClusterStep is a AwsEmrClusterStep
type AwsEmrClusterStep struct {
	ActionOnFailure	string	`json:"action_on_failure"`
	HadoopJarStep	[]AwsEmrClusterHadoopJarStep	`json:"hadoop_jar_step"`
	Name	string	`json:"name"`
}

// AwsEmrClusterKerberosAttributes is a AwsEmrClusterKerberosAttributes
type AwsEmrClusterKerberosAttributes struct {
	AdDomainJoinPassword	string	`json:"ad_domain_join_password"`
	AdDomainJoinUser	string	`json:"ad_domain_join_user"`
	CrossRealmTrustPrincipalPassword	string	`json:"cross_realm_trust_principal_password"`
	KdcAdminPassword	string	`json:"kdc_admin_password"`
	Realm	string	`json:"realm"`
}
