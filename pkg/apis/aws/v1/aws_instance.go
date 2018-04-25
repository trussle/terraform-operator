
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
	EbsOptimized	bool	`json:"ebs_optimized"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	InstanceType	string	`json:"instance_type"`
	UserData	string	`json:"user_data"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	Monitoring	bool	`json:"monitoring"`
	Tags	map[string]interface{}	`json:"tags"`
	GetPasswordData	bool	`json:"get_password_data"`
	BlockDevice	map[string]interface{}	`json:"block_device"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	Ami	string	`json:"ami"`
	UserDataBase64	string	`json:"user_data_base64"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsInstanceList is a list of AwsInstance resources
type AwsInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInstance	`json:"items"`
}

