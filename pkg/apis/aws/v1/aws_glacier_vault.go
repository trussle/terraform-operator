
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGlacierVault struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGlacierVaultSpec	`json:"spec"`
}

type AwsGlacierVaultList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlacierVault	`json:"items"`
}

type AwsGlacierVaultSpec struct {
	Name	string	`json:"name"`
	AccessPolicy	string	`json:"access_policy"`
	Notification	[]interface{}	`json:"notification"`
	Tags	map[string]interface{}	`json:"tags"`
}
