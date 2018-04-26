
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceDirectory describes a AwsDirectoryServiceDirectory resource
type AwsDirectoryServiceDirectory struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDirectoryServiceDirectorySpec	`json:"spec"`
}


// AwsDirectoryServiceDirectorySpec is the spec for a AwsDirectoryServiceDirectory Resource
type AwsDirectoryServiceDirectorySpec struct {
	Description	string	`json:"description"`
	VpcSettings	[]GZamCToZ	`json:"vpc_settings"`
	Tags	map[string]???	`json:"tags"`
	ConnectSettings	[]vPynaEph	`json:"connect_settings"`
	EnableSso	bool	`json:"enable_sso"`
	Name	string	`json:"name"`
	Password	string	`json:"password"`
	Type	string	`json:"type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceDirectoryList is a list of AwsDirectoryServiceDirectory resources
type AwsDirectoryServiceDirectoryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDirectoryServiceDirectory	`json:"items"`
}


// GZamCToZ is a GZamCToZ
type GZamCToZ struct {
	VpcId	string	`json:"vpc_id"`
	SubnetIds	string	`json:"subnet_ids"`
}

// vPynaEph is a vPynaEph
type vPynaEph struct {
	CustomerUsername	string	`json:"customer_username"`
	CustomerDnsIps	string	`json:"customer_dns_ips"`
	SubnetIds	string	`json:"subnet_ids"`
	VpcId	string	`json:"vpc_id"`
}
