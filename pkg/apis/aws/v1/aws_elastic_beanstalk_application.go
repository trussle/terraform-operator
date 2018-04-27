
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	AppversionLifecycle	[]AwsElasticBeanstalkApplicationAppversionLifecycle	`json:"appversion_lifecycle"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticBeanstalkApplicationList is a list of AwsElasticBeanstalkApplication resources
type AwsElasticBeanstalkApplicationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticBeanstalkApplication	`json:"items"`
}


// AwsElasticBeanstalkApplicationAppversionLifecycle is a AwsElasticBeanstalkApplicationAppversionLifecycle
type AwsElasticBeanstalkApplicationAppversionLifecycle struct {
	MaxCount	int	`json:"max_count"`
	DeleteSourceFromS3	bool	`json:"delete_source_from_s3"`
	ServiceRole	string	`json:"service_role"`
	MaxAgeInDays	int	`json:"max_age_in_days"`
}
