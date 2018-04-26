
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInstance describes a AwsInstance resource
type AwsInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInstanceSpec	`json:"spec"`
}


// AwsInstanceSpec is the spec for a AwsInstance Resource
type AwsInstanceSpec struct {
	UserData	string	`json:"user_data"`
	InstanceType	string	`json:"instance_type"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	UserDataBase64	string	`json:"user_data_base64"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	BlockDevice	map[string]???	`json:"block_device"`
	CreditSpecification	[]tPVGNpdG	`json:"credit_specification"`
	GetPasswordData	bool	`json:"get_password_data"`
	Monitoring	bool	`json:"monitoring"`
	Ami	string	`json:"ami"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	Tags	map[string]???	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInstanceList is a list of AwsInstance resources
type AwsInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInstance	`json:"items"`
}


// PqWARPXP is a PqWARPXP
type PqWARPXP struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}

// tPVGNpdG is a tPVGNpdG
type tPVGNpdG struct {
	CpuCredits	string	`json:"cpu_credits"`
}
