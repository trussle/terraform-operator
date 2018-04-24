
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkEnvironment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticBeanstalkEnvironmentSpec	`json:"spec"`
}

type AwsElasticBeanstalkEnvironmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkEnvironment	`json:"items"`
}

type AwsElasticBeanstalkEnvironmentSpec struct {
	Setting	interface{}	`json:"setting"`
	Tags	map[string]interface{}	`json:"tags"`
	Application	string	`json:"application"`
	Tier	string	`json:"tier"`
	PollInterval	string	`json:"poll_interval"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	WaitForReadyTimeout	string	`json:"wait_for_ready_timeout"`
	TemplateName	string	`json:"template_name"`
}
