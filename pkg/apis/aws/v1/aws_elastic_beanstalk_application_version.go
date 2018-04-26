
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkApplicationVersion describes a AwsElasticBeanstalkApplicationVersion resource
type AwsElasticBeanstalkApplicationVersion struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticBeanstalkApplicationVersionSpec	`json:"spec"`
}


// AwsElasticBeanstalkApplicationVersionSpec is the spec for a AwsElasticBeanstalkApplicationVersion Resource
type AwsElasticBeanstalkApplicationVersionSpec struct {
	Application	string	`json:"application"`
	Description	string	`json:"description"`
	Bucket	string	`json:"bucket"`
	Key	string	`json:"key"`
	Name	string	`json:"name"`
	ForceDelete	bool	`json:"force_delete"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkApplicationVersionList is a list of AwsElasticBeanstalkApplicationVersion resources
type AwsElasticBeanstalkApplicationVersionList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkApplicationVersion	`json:"items"`
}

