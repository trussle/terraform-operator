
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
	Name	string	`json:"name"`
	Application	string	`json:"application"`
	Description	string	`json:"description"`
	Tier	string	`json:"tier"`
	PollInterval	string	`json:"poll_interval"`
	Tags	map[string]string	`json:"tags"`
	Setting	Setting	`json:"setting"`
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


// Setting is a Setting
type Setting struct {
	Namespace	string	`json:"namespace"`
	Name	string	`json:"name"`
	Value	string	`json:"value"`
	Resource	string	`json:"resource"`
}

// AllSettings is a AllSettings
type AllSettings struct {
	Namespace	string	`json:"namespace"`
	Name	string	`json:"name"`
	Value	string	`json:"value"`
	Resource	string	`json:"resource"`
}
