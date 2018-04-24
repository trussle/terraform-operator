
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsInstanceSpec	`json:"spec"`
}

type AwsInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsInstance	`json:"items"`
}

type AwsInstanceSpec struct {
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	UserDataBase64	string	`json:"user_data_base64"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	GetPasswordData	bool	`json:"get_password_data"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	InstanceType	string	`json:"instance_type"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	Tags	map[string]interface{}	`json:"tags"`
	BlockDevice	map[string]interface{}	`json:"block_device"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	Monitoring	bool	`json:"monitoring"`
	UserData	string	`json:"user_data"`
	Ami	string	`json:"ami"`
}
