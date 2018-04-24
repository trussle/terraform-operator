
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEip struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsEipSpec	`json:"spec"`
}

type AwsEipList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsEip	`json:"items"`
}

type AwsEipSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	AssociateWithPrivateIp	string	`json:"associate_with_private_ip"`
}
