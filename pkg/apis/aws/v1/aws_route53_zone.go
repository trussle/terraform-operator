
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRoute53Zone struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRoute53ZoneSpec	`json:"spec"`
}

type AwsRoute53ZoneList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRoute53Zone	`json:"items"`
}

type AwsRoute53ZoneSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	ForceDestroy	bool	`json:"force_destroy"`
	Name	string	`json:"name"`
	Comment	string	`json:"comment"`
	VpcId	string	`json:"vpc_id"`
	DelegationSetId	string	`json:"delegation_set_id"`
}
