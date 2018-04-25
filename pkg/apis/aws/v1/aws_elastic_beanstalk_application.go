
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkApplication describes a AwsElasticBeanstalkApplication resource
type AwsElasticBeanstalkApplication struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticBeanstalkApplicationSpec	`json:"spec"`
}


// AwsElasticBeanstalkApplicationSpec is the spec for a AwsElasticBeanstalkApplication Resource
type AwsElasticBeanstalkApplicationSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkApplicationList is a list of AwsElasticBeanstalkApplication resources
type AwsElasticBeanstalkApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkApplication	`json:"items"`
}

