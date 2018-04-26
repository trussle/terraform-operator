
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkConfigurationTemplate describes a AwsElasticBeanstalkConfigurationTemplate resource
type AwsElasticBeanstalkConfigurationTemplate struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticBeanstalkConfigurationTemplateSpec	`json:"spec"`
}


// AwsElasticBeanstalkConfigurationTemplateSpec is the spec for a AwsElasticBeanstalkConfigurationTemplate Resource
type AwsElasticBeanstalkConfigurationTemplateSpec struct {
	SolutionStackName	string	`json:"solution_stack_name"`
	Name	string	`json:"name"`
	Application	string	`json:"application"`
	Description	string	`json:"description"`
	EnvironmentId	string	`json:"environment_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkConfigurationTemplateList is a list of AwsElasticBeanstalkConfigurationTemplate resources
type AwsElasticBeanstalkConfigurationTemplateList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkConfigurationTemplate	`json:"items"`
}

