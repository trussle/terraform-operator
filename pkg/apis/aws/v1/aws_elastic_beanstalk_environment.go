
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Setting	AwsElasticBeanstalkEnvironmentSetting	`json:"setting"`
	Name	string	`json:"name"`
	PollInterval	string	`json:"poll_interval"`
	Tags	map[string]string	`json:"tags"`
	Tier	string	`json:"tier"`
	Application	string	`json:"application"`
	Description	string	`json:"description"`
	TemplateName	string	`json:"template_name"`
	WaitForReadyTimeout	string	`json:"wait_for_ready_timeout"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkEnvironmentList is a list of AwsElasticBeanstalkEnvironment resources
type AwsElasticBeanstalkEnvironmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkEnvironment	`json:"items"`
}


// AwsElasticBeanstalkEnvironmentSetting is a AwsElasticBeanstalkEnvironmentSetting
type AwsElasticBeanstalkEnvironmentSetting struct {
	Value	string	`json:"value"`
	Resource	string	`json:"resource"`
	Namespace	string	`json:"namespace"`
	Name	string	`json:"name"`
}

// AwsElasticBeanstalkEnvironmentAllSettings is a AwsElasticBeanstalkEnvironmentAllSettings
type AwsElasticBeanstalkEnvironmentAllSettings struct {
	Resource	string	`json:"resource"`
	Namespace	string	`json:"namespace"`
	Name	string	`json:"name"`
	Value	string	`json:"value"`
}
