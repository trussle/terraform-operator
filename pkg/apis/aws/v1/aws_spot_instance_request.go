
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSpotInstanceRequest struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSpotInstanceRequestSpec	`json:"spec"`
}

type AwsSpotInstanceRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotInstanceRequest	`json:"items"`
}

type AwsSpotInstanceRequestSpec struct {
	InstanceType	string	`json:"instance_type"`
	UserDataBase64	string	`json:"user_data_base64"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
	GetPasswordData	bool	`json:"get_password_data"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	Monitoring	bool	`json:"monitoring"`
	Tags	map[string]interface{}	`json:"tags"`
	SpotPrice	string	`json:"spot_price"`
	LaunchGroup	string	`json:"launch_group"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	BlockDevice	map[string]interface{}	`json:"block_device"`
	SpotType	string	`json:"spot_type"`
	UserData	string	`json:"user_data"`
	Ami	string	`json:"ami"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	VolumeTags	map[string]interface{}	`json:"volume_tags"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
}
