
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
	CoreInstanceCount	int	`json:"core_instance_count"`
	Applications	string	`json:"applications"`
	BootstrapAction	string	`json:"bootstrap_action"`
	Configurations	string	`json:"configurations"`
	Name	string	`json:"name"`
	MasterInstanceType	string	`json:"master_instance_type"`
	EbsRootVolumeSize	int	`json:"ebs_root_volume_size"`
	Ec2Attributes	[]fFSPqKXS	`json:"ec2_attributes"`
	AutoscalingRole	string	`json:"autoscaling_role"`
	KerberosAttributes	[]LEfZAPaJ	`json:"kerberos_attributes"`
	InstanceGroup	string	`json:"instance_group"`
	Tags	map[string]???	`json:"tags"`
	ServiceRole	string	`json:"service_role"`
	SecurityConfiguration	string	`json:"security_configuration"`
	ReleaseLabel	string	`json:"release_label"`
	LogUri	string	`json:"log_uri"`
	VisibleToAllUsers	bool	`json:"visible_to_all_users"`
	CustomAmiId	string	`json:"custom_ami_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEmrClusterList is a list of AwsEmrCluster resources
type AwsEmrClusterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEmrCluster	`json:"items"`
}


// fFSPqKXS is a fFSPqKXS
type fFSPqKXS struct {
	InstanceProfile	string	`json:"instance_profile"`
	ServiceAccessSecurityGroup	string	`json:"service_access_security_group"`
	KeyName	string	`json:"key_name"`
	SubnetId	string	`json:"subnet_id"`
	AdditionalMasterSecurityGroups	string	`json:"additional_master_security_groups"`
	AdditionalSlaveSecurityGroups	string	`json:"additional_slave_security_groups"`
	EmrManagedMasterSecurityGroup	string	`json:"emr_managed_master_security_group"`
	EmrManagedSlaveSecurityGroup	string	`json:"emr_managed_slave_security_group"`
}

// LEfZAPaJ is a LEfZAPaJ
type LEfZAPaJ struct {
	AdDomainJoinUser	string	`json:"ad_domain_join_user"`
	CrossRealmTrustPrincipalPassword	string	`json:"cross_realm_trust_principal_password"`
	KdcAdminPassword	string	`json:"kdc_admin_password"`
	Realm	string	`json:"realm"`
	AdDomainJoinPassword	string	`json:"ad_domain_join_password"`
}

// ITGHvBrQ is a ITGHvBrQ
type ITGHvBrQ struct {
	MainClass	string	`json:"main_class"`
	Properties	map[string]???	`json:"properties"`
	Args	[]string	`json:"args"`
	Jar	string	`json:"jar"`
}

// zldaUEdh is a zldaUEdh
type zldaUEdh struct {
	ActionOnFailure	string	`json:"action_on_failure"`
	HadoopJarStep	[]ITGHvBrQ	`json:"hadoop_jar_step"`
	Name	string	`json:"name"`
}
