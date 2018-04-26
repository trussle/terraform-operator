
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
	ConnectSettings	[]Generic	`json:"connect_settings"`
	Description	string	`json:"description"`
	VpcSettings	[]Generic	`json:"vpc_settings"`
	EnableSso	bool	`json:"enable_sso"`
	Tags	map[string]Generic	`json:"tags"`
	Type	string	`json:"type"`
	Name	string	`json:"name"`
	Password	string	`json:"password"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDirectoryServiceDirectoryList is a list of AwsDirectoryServiceDirectory resources
type AwsDirectoryServiceDirectoryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDirectoryServiceDirectory	`json:"items"`
}

