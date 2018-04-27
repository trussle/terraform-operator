
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
	VpcSettings	[]AwsDirectoryServiceDirectoryVpcSettings	`json:"vpc_settings"`
	EnableSso	bool	`json:"enable_sso"`
	Type	string	`json:"type"`
	Password	string	`json:"password"`
	Description	string	`json:"description"`
	ConnectSettings	[]AwsDirectoryServiceDirectoryConnectSettings	`json:"connect_settings"`
	Tags	map[string]string	`json:"tags"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceDirectoryList is a list of AwsDirectoryServiceDirectory resources
type AwsDirectoryServiceDirectoryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDirectoryServiceDirectory	`json:"items"`
}


// AwsDirectoryServiceDirectoryVpcSettings is a AwsDirectoryServiceDirectoryVpcSettings
type AwsDirectoryServiceDirectoryVpcSettings struct {
	SubnetIds	string	`json:"subnet_ids"`
	VpcId	string	`json:"vpc_id"`
}

// AwsDirectoryServiceDirectoryConnectSettings is a AwsDirectoryServiceDirectoryConnectSettings
type AwsDirectoryServiceDirectoryConnectSettings struct {
	CustomerUsername	string	`json:"customer_username"`
	CustomerDnsIps	string	`json:"customer_dns_ips"`
	SubnetIds	string	`json:"subnet_ids"`
	VpcId	string	`json:"vpc_id"`
}
