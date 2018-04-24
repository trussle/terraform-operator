
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsMqBroker struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsMqBrokerSpec	`json:"spec"`
}

type AwsMqBrokerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMqBroker	`json:"items"`
}

type AwsMqBrokerSpec struct {
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	EngineType	string	`json:"engine_type"`
	HostInstanceType	string	`json:"host_instance_type"`
	ApplyImmediately	bool	`json:"apply_immediately"`
	BrokerName	string	`json:"broker_name"`
	DeploymentMode	string	`json:"deployment_mode"`
	EngineVersion	string	`json:"engine_version"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	SecurityGroups	interface{}	`json:"security_groups"`
	User	interface{}	`json:"user"`
}
