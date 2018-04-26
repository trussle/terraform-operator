
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
	BuildId	string	`json:"build_id"`
	Ec2InstanceType	string	`json:"ec2_instance_type"`
	Description	string	`json:"description"`
	Ec2InboundPermission	[]Ec2InboundPermission	`json:"ec2_inbound_permission"`
	NewGameSessionProtectionPolicy	string	`json:"new_game_session_protection_policy"`
	Name	string	`json:"name"`
	ResourceCreationLimitPolicy	[]ResourceCreationLimitPolicy	`json:"resource_creation_limit_policy"`
	RuntimeConfiguration	[]RuntimeConfiguration	`json:"runtime_configuration"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftFleetList is a list of AwsGameliftFleet resources
type AwsGameliftFleetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftFleet	`json:"items"`
}


// Ec2InboundPermission is a Ec2InboundPermission
type Ec2InboundPermission struct {
	Protocol	string	`json:"protocol"`
	ToPort	int	`json:"to_port"`
	FromPort	int	`json:"from_port"`
	IpRange	string	`json:"ip_range"`
}

// ResourceCreationLimitPolicy is a ResourceCreationLimitPolicy
type ResourceCreationLimitPolicy struct {
	NewGameSessionsPerCreator	int	`json:"new_game_sessions_per_creator"`
	PolicyPeriodInMinutes	int	`json:"policy_period_in_minutes"`
}

// ServerProcess is a ServerProcess
type ServerProcess struct {
	ConcurrentExecutions	int	`json:"concurrent_executions"`
	LaunchPath	string	`json:"launch_path"`
	Parameters	string	`json:"parameters"`
}

// RuntimeConfiguration is a RuntimeConfiguration
type RuntimeConfiguration struct {
	GameSessionActivationTimeoutSeconds	int	`json:"game_session_activation_timeout_seconds"`
	MaxConcurrentGameSessionActivations	int	`json:"max_concurrent_game_session_activations"`
	ServerProcess	[]ServerProcess	`json:"server_process"`
}
