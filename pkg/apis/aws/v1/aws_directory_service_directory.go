
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDirectoryServiceDirectory struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDirectoryServiceDirectorySpec	`json:"spec"`
}

type AwsDirectoryServiceDirectoryList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDirectoryServiceDirectory	`json:"items"`
}

type AwsDirectoryServiceDirectorySpec struct {
	Type	string	`json:"type"`
	EnableSso	bool	`json:"enable_sso"`
	VpcSettings	[]interface{}	`json:"vpc_settings"`
	ConnectSettings	[]interface{}	`json:"connect_settings"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	Tags	map[string]interface{}	`json:"tags"`
	Password	string	`json:"password"`
}
