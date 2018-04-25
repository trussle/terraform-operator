
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	ForceDestroy	bool	`json:"force_destroy"`
	Name	string	`json:"name"`
	DelegationSetId	string	`json:"delegation_set_id"`
	Comment	string	`json:"comment"`
	VpcId	string	`json:"vpc_id"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRoute53ZoneList is a list of AwsRoute53Zone resources
type AwsRoute53ZoneList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Zone	`json:"items"`
}

