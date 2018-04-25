
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
	Name	string	`json:"name"`
	Password	string	`json:"password"`
	VpcSettings	[]interface{}	`json:"vpc_settings"`
	Description	string	`json:"description"`
	ConnectSettings	[]interface{}	`json:"connect_settings"`
	EnableSso	bool	`json:"enable_sso"`
	Type	string	`json:"type"`
	Tags	map[string]interface{}	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceDirectoryList is a list of AwsDirectoryServiceDirectory resources
type AwsDirectoryServiceDirectoryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDirectoryServiceDirectory	`json:"items"`
}

