
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCloud9EnvironmentEc2 struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloud9EnvironmentEc2Spec	`json:"spec"`
}

type AwsCloud9EnvironmentEc2List struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloud9EnvironmentEc2	`json:"items"`
}

type AwsCloud9EnvironmentEc2Spec struct {
	Name	string	`json:"name"`
	InstanceType	string	`json:"instance_type"`
	AutomaticStopTimeMinutes	int	`json:"automatic_stop_time_minutes"`
	Description	string	`json:"description"`
	SubnetId	string	`json:"subnet_id"`
}
