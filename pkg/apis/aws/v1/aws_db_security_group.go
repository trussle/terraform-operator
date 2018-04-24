
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDbSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbSecurityGroupSpec	`json:"spec"`
}

type AwsDbSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbSecurityGroup	`json:"items"`
}

type AwsDbSecurityGroupSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Ingress	interface{}	`json:"ingress"`
	Tags	map[string]interface{}	`json:"tags"`
}
