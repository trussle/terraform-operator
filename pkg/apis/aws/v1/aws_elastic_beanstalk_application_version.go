
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticBeanstalkApplicationVersion struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticBeanstalkApplicationVersionSpec	`json:"spec"`
}

type AwsElasticBeanstalkApplicationVersionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkApplicationVersion	`json:"items"`
}

type AwsElasticBeanstalkApplicationVersionSpec struct {
	Description	string	`json:"description"`
	Bucket	string	`json:"bucket"`
	Key	string	`json:"key"`
	Name	string	`json:"name"`
	ForceDelete	bool	`json:"force_delete"`
	Application	string	`json:"application"`
}
