
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloud9EnvironmentEc2 describes a AwsCloud9EnvironmentEc2 resource
type AwsCloud9EnvironmentEc2 struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloud9EnvironmentEc2Spec	`json:"spec"`
}


// AwsCloud9EnvironmentEc2Spec is the spec for a AwsCloud9EnvironmentEc2 Resource
type AwsCloud9EnvironmentEc2Spec struct {
	Name	string	`json:"name"`
	InstanceType	string	`json:"instance_type"`
	AutomaticStopTimeMinutes	int	`json:"automatic_stop_time_minutes"`
	Description	string	`json:"description"`
	SubnetId	string	`json:"subnet_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloud9EnvironmentEc2List is a list of AwsCloud9EnvironmentEc2 resources
type AwsCloud9EnvironmentEc2List struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloud9EnvironmentEc2	`json:"items"`
}

