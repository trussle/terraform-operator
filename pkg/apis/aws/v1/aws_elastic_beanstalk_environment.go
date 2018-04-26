
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkEnvironment describes a AwsElasticBeanstalkEnvironment resource
type AwsElasticBeanstalkEnvironment struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticBeanstalkEnvironmentSpec	`json:"spec"`
}


// AwsElasticBeanstalkEnvironmentSpec is the spec for a AwsElasticBeanstalkEnvironment Resource
type AwsElasticBeanstalkEnvironmentSpec struct {
	Tags	map[string]Generic	`json:"tags"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	WaitForReadyTimeout	string	`json:"wait_for_ready_timeout"`
	Tier	string	`json:"tier"`
	PollInterval	string	`json:"poll_interval"`
	Application	string	`json:"application"`
	Setting	Generic	`json:"setting"`
	TemplateName	string	`json:"template_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkEnvironmentList is a list of AwsElasticBeanstalkEnvironment resources
type AwsElasticBeanstalkEnvironmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkEnvironment	`json:"items"`
}

