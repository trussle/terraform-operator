
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
	DisableApiTermination	bool	`json:"disable_api_termination"`
	VolumeTags	map[string]Generic	`json:"volume_tags"`
	Tags	map[string]Generic	`json:"tags"`
	BlockDevice	map[string]Generic	`json:"block_device"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	SpotType	string	`json:"spot_type"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	Ami	string	`json:"ami"`
	UserData	string	`json:"user_data"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	SpotPrice	string	`json:"spot_price"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	GetPasswordData	bool	`json:"get_password_data"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	LaunchGroup	string	`json:"launch_group"`
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
	Monitoring	bool	`json:"monitoring"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	InstanceType	string	`json:"instance_type"`
	UserDataBase64	string	`json:"user_data_base64"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotInstanceRequestList is a list of AwsSpotInstanceRequest resources
type AwsSpotInstanceRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotInstanceRequest	`json:"items"`
}

