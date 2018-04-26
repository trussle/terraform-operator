
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
	ConnectSettings	[]ConnectSettings	`json:"connect_settings"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Tags	map[string]string	`json:"tags"`
	VpcSettings	[]VpcSettings	`json:"vpc_settings"`
	Password	string	`json:"password"`
	Type	string	`json:"type"`
	EnableSso	bool	`json:"enable_sso"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceDirectoryList is a list of AwsDirectoryServiceDirectory resources
type AwsDirectoryServiceDirectoryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDirectoryServiceDirectory	`json:"items"`
}


// ConnectSettings is a ConnectSettings
type ConnectSettings struct {
	CustomerUsername	string	`json:"customer_username"`
	CustomerDnsIps	string	`json:"customer_dns_ips"`
	SubnetIds	string	`json:"subnet_ids"`
	VpcId	string	`json:"vpc_id"`
}

// VpcSettings is a VpcSettings
type VpcSettings struct {
	SubnetIds	string	`json:"subnet_ids"`
	VpcId	string	`json:"vpc_id"`
}
