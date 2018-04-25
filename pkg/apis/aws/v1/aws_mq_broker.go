
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqBroker describes a AwsMqBroker resource
type AwsMqBroker struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsMqBrokerSpec	`json:"spec"`
}


// AwsMqBrokerSpec is the spec for a AwsMqBroker Resource
type AwsMqBrokerSpec struct {
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	ApplyImmediately	bool	`json:"apply_immediately"`
	EngineType	string	`json:"engine_type"`
	BrokerName	string	`json:"broker_name"`
	HostInstanceType	string	`json:"host_instance_type"`
	User	string	`json:"user"`
	DeploymentMode	string	`json:"deployment_mode"`
	EngineVersion	string	`json:"engine_version"`
	SecurityGroups	string	`json:"security_groups"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqBrokerList is a list of AwsMqBroker resources
type AwsMqBrokerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMqBroker	`json:"items"`
}

