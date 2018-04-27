
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftFleet describes a AwsGameliftFleet resource
type AwsGameliftFleet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGameliftFleetSpec	`json:"spec"`
}


// AwsGameliftFleetSpec is the spec for a AwsGameliftFleet Resource
type AwsGameliftFleetSpec struct {
	Ec2InstanceType	string	`json:"ec2_instance_type"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	RuntimeConfiguration	[]AwsGameliftFleetRuntimeConfiguration	`json:"runtime_configuration"`
	BuildId	string	`json:"build_id"`
	NewGameSessionProtectionPolicy	string	`json:"new_game_session_protection_policy"`
	ResourceCreationLimitPolicy	[]AwsGameliftFleetResourceCreationLimitPolicy	`json:"resource_creation_limit_policy"`
	Ec2InboundPermission	[]AwsGameliftFleetEc2InboundPermission	`json:"ec2_inbound_permission"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftFleetList is a list of AwsGameliftFleet resources
type AwsGameliftFleetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftFleet	`json:"items"`
}


// AwsGameliftFleetResourceCreationLimitPolicy is a AwsGameliftFleetResourceCreationLimitPolicy
type AwsGameliftFleetResourceCreationLimitPolicy struct {
	NewGameSessionsPerCreator	int	`json:"new_game_sessions_per_creator"`
	PolicyPeriodInMinutes	int	`json:"policy_period_in_minutes"`
}

// AwsGameliftFleetEc2InboundPermission is a AwsGameliftFleetEc2InboundPermission
type AwsGameliftFleetEc2InboundPermission struct {
	FromPort	int	`json:"from_port"`
	IpRange	string	`json:"ip_range"`
	Protocol	string	`json:"protocol"`
	ToPort	int	`json:"to_port"`
}

// AwsGameliftFleetServerProcess is a AwsGameliftFleetServerProcess
type AwsGameliftFleetServerProcess struct {
	ConcurrentExecutions	int	`json:"concurrent_executions"`
	LaunchPath	string	`json:"launch_path"`
	Parameters	string	`json:"parameters"`
}

// AwsGameliftFleetRuntimeConfiguration is a AwsGameliftFleetRuntimeConfiguration
type AwsGameliftFleetRuntimeConfiguration struct {
	MaxConcurrentGameSessionActivations	int	`json:"max_concurrent_game_session_activations"`
	ServerProcess	[]AwsGameliftFleetServerProcess	`json:"server_process"`
	GameSessionActivationTimeoutSeconds	int	`json:"game_session_activation_timeout_seconds"`
}
