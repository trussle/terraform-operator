
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkConfigurationTemplate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticBeanstalkConfigurationTemplateSpec	`json:"spec"`
}

type AwsElasticBeanstalkConfigurationTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkConfigurationTemplate	`json:"items"`
}

type AwsElasticBeanstalkConfigurationTemplateSpec struct {
	Name	string	`json:"name"`
	Application	string	`json:"application"`
	Description	string	`json:"description"`
	EnvironmentId	string	`json:"environment_id"`
	SolutionStackName	string	`json:"solution_stack_name"`
}
