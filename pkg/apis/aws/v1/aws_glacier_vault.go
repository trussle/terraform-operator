
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlacierVault describes a AwsGlacierVault resource
type AwsGlacierVault struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGlacierVaultSpec	`json:"spec"`
}


// AwsGlacierVaultSpec is the spec for a AwsGlacierVault Resource
type AwsGlacierVaultSpec struct {
	AccessPolicy	string	`json:"access_policy"`
	Notification	[]AwsGlacierVaultNotification	`json:"notification"`
	Tags	map[string]string	`json:"tags"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlacierVaultList is a list of AwsGlacierVault resources
type AwsGlacierVaultList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlacierVault	`json:"items"`
}


// AwsGlacierVaultNotification is a AwsGlacierVaultNotification
type AwsGlacierVaultNotification struct {
	Events	string	`json:"events"`
	SnsTopic	string	`json:"sns_topic"`
}
