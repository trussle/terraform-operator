
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	DeploymentMode	string	`json:"deployment_mode"`
	EngineVersion	string	`json:"engine_version"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	ApplyImmediately	bool	`json:"apply_immediately"`
	BrokerName	string	`json:"broker_name"`
	EngineType	string	`json:"engine_type"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	HostInstanceType	string	`json:"host_instance_type"`
	SecurityGroups	string	`json:"security_groups"`
	User	string	`json:"user"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqBrokerList is a list of AwsMqBroker resources
type AwsMqBrokerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMqBroker	`json:"items"`
}


// VVxSAZWE is a VVxSAZWE
type VVxSAZWE struct {
	DayOfWeek	string	`json:"day_of_week"`
	TimeOfDay	string	`json:"time_of_day"`
	TimeZone	string	`json:"time_zone"`
}

// XejhAquX is a XejhAquX
type XejhAquX struct {
}

// daaaZlRH is a daaaZlRH
type daaaZlRH struct {
}
