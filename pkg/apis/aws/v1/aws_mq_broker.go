
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
	HostInstanceType	string	`json:"host_instance_type"`
	BrokerName	string	`json:"broker_name"`
	EngineVersion	string	`json:"engine_version"`
	User	AwsMqBrokerUser	`json:"user"`
	ApplyImmediately	bool	`json:"apply_immediately"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	EngineType	string	`json:"engine_type"`
	SecurityGroups	string	`json:"security_groups"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqBrokerList is a list of AwsMqBroker resources
type AwsMqBrokerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMqBroker	`json:"items"`
}


// AwsMqBrokerConfiguration is a AwsMqBrokerConfiguration
type AwsMqBrokerConfiguration struct {
}

// AwsMqBrokerUser is a AwsMqBrokerUser
type AwsMqBrokerUser struct {
	ConsoleAccess	bool	`json:"console_access"`
	Groups	string	`json:"groups"`
	Password	string	`json:"password"`
	Username	string	`json:"username"`
}

// AwsMqBrokerMaintenanceWindowStartTime is a AwsMqBrokerMaintenanceWindowStartTime
type AwsMqBrokerMaintenanceWindowStartTime struct {
	DayOfWeek	string	`json:"day_of_week"`
	TimeOfDay	string	`json:"time_of_day"`
	TimeZone	string	`json:"time_zone"`
}

// AwsMqBrokerInstances is a AwsMqBrokerInstances
type AwsMqBrokerInstances struct {
}
