
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Name	string	`json:"name"`
	AccessPolicy	string	`json:"access_policy"`
	Notification	[]interface{}	`json:"notification"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGlacierVaultList is a list of AwsGlacierVault resources
type AwsGlacierVaultList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGlacierVault	`json:"items"`
}

