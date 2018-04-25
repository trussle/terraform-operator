
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotInstanceRequest describes a AwsSpotInstanceRequest resource
type AwsSpotInstanceRequest struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSpotInstanceRequestSpec	`json:"spec"`
}


// AwsSpotInstanceRequestSpec is the spec for a AwsSpotInstanceRequest Resource
type AwsSpotInstanceRequestSpec struct {
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
	Monitoring	bool	`json:"monitoring"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	VolumeTags	map[string]interface{}	`json:"volume_tags"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	Ami	string	`json:"ami"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	BlockDevice	map[string]interface{}	`json:"block_device"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	SpotPrice	string	`json:"spot_price"`
	InstanceType	string	`json:"instance_type"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	Tags	map[string]interface{}	`json:"tags"`
	LaunchGroup	string	`json:"launch_group"`
	GetPasswordData	bool	`json:"get_password_data"`
	UserData	string	`json:"user_data"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	SpotType	string	`json:"spot_type"`
	UserDataBase64	string	`json:"user_data_base64"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotInstanceRequestList is a list of AwsSpotInstanceRequest resources
type AwsSpotInstanceRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotInstanceRequest	`json:"items"`
}

