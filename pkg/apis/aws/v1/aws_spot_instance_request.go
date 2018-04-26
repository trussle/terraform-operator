
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Tags	map[string]???	`json:"tags"`
	SourceDestCheck	bool	`json:"source_dest_check"`
	UserData	string	`json:"user_data"`
	GetPasswordData	bool	`json:"get_password_data"`
	DisableApiTermination	bool	`json:"disable_api_termination"`
	Ami	string	`json:"ami"`
	InstanceType	string	`json:"instance_type"`
	EbsOptimized	bool	`json:"ebs_optimized"`
	InstanceInitiatedShutdownBehavior	string	`json:"instance_initiated_shutdown_behavior"`
	IamInstanceProfile	string	`json:"iam_instance_profile"`
	LaunchGroup	string	`json:"launch_group"`
	BlockDurationMinutes	int	`json:"block_duration_minutes"`
	InstanceInterruptionBehaviour	string	`json:"instance_interruption_behaviour"`
	CreditSpecification	[]LVCxaSJl	`json:"credit_specification"`
	Monitoring	bool	`json:"monitoring"`
	SpotType	string	`json:"spot_type"`
	UserDataBase64	string	`json:"user_data_base64"`
	SpotPrice	string	`json:"spot_price"`
	WaitForFulfillment	bool	`json:"wait_for_fulfillment"`
	VolumeTags	map[string]???	`json:"volume_tags"`
	BlockDevice	map[string]???	`json:"block_device"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSpotInstanceRequestList is a list of AwsSpotInstanceRequest resources
type AwsSpotInstanceRequestList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSpotInstanceRequest	`json:"items"`
}


// zwVGqMZc is a zwVGqMZc
type zwVGqMZc struct {
	DeleteOnTermination	bool	`json:"delete_on_termination"`
}

// LVCxaSJl is a LVCxaSJl
type LVCxaSJl struct {
	CpuCredits	string	`json:"cpu_credits"`
}
