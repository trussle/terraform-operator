
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftSecurityGroup describes a AwsRedshiftSecurityGroup resource
type AwsRedshiftSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRedshiftSecurityGroupSpec	`json:"spec"`
}


// AwsRedshiftSecurityGroupSpec is the spec for a AwsRedshiftSecurityGroup Resource
type AwsRedshiftSecurityGroupSpec struct {
	Ingress	string	`json:"ingress"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRedshiftSecurityGroupList is a list of AwsRedshiftSecurityGroup resources
type AwsRedshiftSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRedshiftSecurityGroup	`json:"items"`
}

