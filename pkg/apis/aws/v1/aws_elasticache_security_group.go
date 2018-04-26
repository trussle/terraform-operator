
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheSecurityGroup describes a AwsElasticacheSecurityGroup resource
type AwsElasticacheSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticacheSecurityGroupSpec	`json:"spec"`
}


// AwsElasticacheSecurityGroupSpec is the spec for a AwsElasticacheSecurityGroup Resource
type AwsElasticacheSecurityGroupSpec struct {
	SecurityGroupNames	Generic	`json:"security_group_names"`
	Description	string	`json:"description"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheSecurityGroupList is a list of AwsElasticacheSecurityGroup resources
type AwsElasticacheSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticacheSecurityGroup	`json:"items"`
}

