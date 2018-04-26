
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53Zone describes a AwsRoute53Zone resource
type AwsRoute53Zone struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53ZoneSpec	`json:"spec"`
}


// AwsRoute53ZoneSpec is the spec for a AwsRoute53Zone Resource
type AwsRoute53ZoneSpec struct {
	Name	string	`json:"name"`
	Comment	string	`json:"comment"`
	ForceDestroy	bool	`json:"force_destroy"`
	VpcId	string	`json:"vpc_id"`
	DelegationSetId	string	`json:"delegation_set_id"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53ZoneList is a list of AwsRoute53Zone resources
type AwsRoute53ZoneList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Zone	`json:"items"`
}

