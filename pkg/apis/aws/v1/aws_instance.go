
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	SourceDestCheck	bool	`json:"source_dest_check"`
	UserData	string	`json:"user_data"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	Ami	string	`json:"ami"`
	GetPasswordData	bool	`json:"get_password_data"`
	Monitoring	bool	`json:"monitoring"`
	InstanceType	string	`json:"instance_type"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	BlockDevice	map[string]Generic	`json:"block_device"`
	UserDataBase64	string	`json:"user_data_base64"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	Tags	map[string]Generic	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInstanceList is a list of AwsInstance resources
type AwsInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInstance	`json:"items"`
}

