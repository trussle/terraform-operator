
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCustomerGateway describes a AwsCustomerGateway resource
type AwsCustomerGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCustomerGatewaySpec	`json:"spec"`
}


// AwsCustomerGatewaySpec is the spec for a AwsCustomerGateway Resource
type AwsCustomerGatewaySpec struct {
	BgpAsn	int	`json:"bgp_asn"`
	IpAddress	string	`json:"ip_address"`
	Type	string	`json:"type"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCustomerGatewayList is a list of AwsCustomerGateway resources
type AwsCustomerGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCustomerGateway	`json:"items"`
}

