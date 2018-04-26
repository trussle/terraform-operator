
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
	ApplyImmediately	bool	`json:"apply_immediately"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	DeploymentMode	string	`json:"deployment_mode"`
	HostInstanceType	string	`json:"host_instance_type"`
	BrokerName	string	`json:"broker_name"`
	EngineType	string	`json:"engine_type"`
	EngineVersion	string	`json:"engine_version"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	SecurityGroups	string	`json:"security_groups"`
	User	User	`json:"user"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsMqBrokerList is a list of AwsMqBroker resources
type AwsMqBrokerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsMqBroker	`json:"items"`
}


// MaintenanceWindowStartTime is a MaintenanceWindowStartTime
type MaintenanceWindowStartTime struct {
	TimeOfDay	string	`json:"time_of_day"`
	TimeZone	string	`json:"time_zone"`
	DayOfWeek	string	`json:"day_of_week"`
}

// Instances is a Instances
type Instances struct {
}

// User is a User
type User struct {
	Groups	string	`json:"groups"`
	Password	string	`json:"password"`
	Username	string	`json:"username"`
	ConsoleAccess	bool	`json:"console_access"`
}

// Configuration is a Configuration
type Configuration struct {
}
