
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftFleet struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGameliftFleetSpec	`json:"spec"`
}

type AwsGameliftFleetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftFleet	`json:"items"`
}

type AwsGameliftFleetSpec struct {
	Ec2InstanceType	string	`json:"ec2_instance_type"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	RuntimeConfiguration	[]interface{}	`json:"runtime_configuration"`
	BuildId	string	`json:"build_id"`
	Ec2InboundPermission	[]interface{}	`json:"ec2_inbound_permission"`
	NewGameSessionProtectionPolicy	string	`json:"new_game_session_protection_policy"`
	ResourceCreationLimitPolicy	[]interface{}	`json:"resource_creation_limit_policy"`
}
