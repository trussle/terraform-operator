
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsNetworkInterface struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsNetworkInterfaceSpec	`json:"spec"`
}

type AwsNetworkInterfaceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsNetworkInterface	`json:"items"`
}

type AwsNetworkInterfaceSpec struct {
	SourceDestCheck	bool	`json:"source_dest_check"`
	Description	string	`json:"description"`
	Tags	map[string]interface{}	`json:"tags"`
	SubnetId	string	`json:"subnet_id"`
}
