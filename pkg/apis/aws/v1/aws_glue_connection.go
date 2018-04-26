
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueConnection describes a AwsGlueConnection resource
type AwsGlueConnection struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGlueConnectionSpec	`json:"spec"`
}


// AwsGlueConnectionSpec is the spec for a AwsGlueConnection Resource
type AwsGlueConnectionSpec struct {
	ConnectionProperties	map[string]Generic	`json:"connection_properties"`
	ConnectionType	string	`json:"connection_type"`
	Description	string	`json:"description"`
	MatchCriteria	[]Generic	`json:"match_criteria"`
	Name	string	`json:"name"`
	PhysicalConnectionRequirements	[]Generic	`json:"physical_connection_requirements"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlueConnectionList is a list of AwsGlueConnection resources
type AwsGlueConnectionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlueConnection	`json:"items"`
}

