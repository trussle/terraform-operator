
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	ResourceCreationLimitPolicy	[]Generic	`json:"resource_creation_limit_policy"`
	RuntimeConfiguration	[]Generic	`json:"runtime_configuration"`
	BuildId	string	`json:"build_id"`
	Description	string	`json:"description"`
	Ec2InboundPermission	[]Generic	`json:"ec2_inbound_permission"`
	NewGameSessionProtectionPolicy	string	`json:"new_game_session_protection_policy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftFleetList is a list of AwsGameliftFleet resources
type AwsGameliftFleetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftFleet	`json:"items"`
}

