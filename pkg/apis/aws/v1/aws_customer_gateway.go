
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsCustomerGateway struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCustomerGatewaySpec	`json:"spec"`
}

type AwsCustomerGatewayList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCustomerGateway	`json:"items"`
}

type AwsCustomerGatewaySpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	BgpAsn	int	`json:"bgp_asn"`
	IpAddress	string	`json:"ip_address"`
	Type	string	`json:"type"`
}
