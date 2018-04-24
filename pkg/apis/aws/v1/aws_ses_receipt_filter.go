
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesReceiptFilter struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSesReceiptFilterSpec	`json:"spec"`
}

type AwsSesReceiptFilterList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSesReceiptFilter	`json:"items"`
}

type AwsSesReceiptFilterSpec struct {
	Name	string	`json:"name"`
	Cidr	string	`json:"cidr"`
	Policy	string	`json:"policy"`
}
