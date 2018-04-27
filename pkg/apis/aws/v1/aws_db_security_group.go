
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbSecurityGroup describes a AwsDbSecurityGroup resource
type AwsDbSecurityGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDbSecurityGroupSpec	`json:"spec"`
}


// AwsDbSecurityGroupSpec is the spec for a AwsDbSecurityGroup Resource
type AwsDbSecurityGroupSpec struct {
	Tags	map[string]string	`json:"tags"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Ingress	AwsDbSecurityGroupIngress	`json:"ingress"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDbSecurityGroupList is a list of AwsDbSecurityGroup resources
type AwsDbSecurityGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDbSecurityGroup	`json:"items"`
}


// AwsDbSecurityGroupIngress is a AwsDbSecurityGroupIngress
type AwsDbSecurityGroupIngress struct {
	Cidr	string	`json:"cidr"`
}
