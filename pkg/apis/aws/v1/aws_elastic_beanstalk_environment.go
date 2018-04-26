
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
	Tier	string	`json:"tier"`
	Setting	string	`json:"setting"`
	WaitForReadyTimeout	string	`json:"wait_for_ready_timeout"`
	Tags	map[string]???	`json:"tags"`
	PollInterval	string	`json:"poll_interval"`
	Application	string	`json:"application"`
	Description	string	`json:"description"`
	TemplateName	string	`json:"template_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkEnvironmentList is a list of AwsElasticBeanstalkEnvironment resources
type AwsElasticBeanstalkEnvironmentList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkEnvironment	`json:"items"`
}

