
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
	SecurityGroups	Generic	`json:"security_groups"`
	ApplyImmediately	bool	`json:"apply_immediately"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	EngineType	string	`json:"engine_type"`
	HostInstanceType	string	`json:"host_instance_type"`
	BrokerName	string	`json:"broker_name"`
	DeploymentMode	string	`json:"deployment_mode"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	EngineVersion	string	`json:"engine_version"`
	User	Generic	`json:"user"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqBrokerList is a list of AwsMqBroker resources
type AwsMqBrokerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMqBroker	`json:"items"`
}

