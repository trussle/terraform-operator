
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAmiCopy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAmiCopySpec	`json:"spec"`
}

type AwsAmiCopyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAmiCopy	`json:"items"`
}

type AwsAmiCopySpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	Name	string	`json:"name"`
	SourceAmiRegion	string	`json:"source_ami_region"`
	Description	string	`json:"description"`
	SourceAmiId	string	`json:"source_ami_id"`
	Encrypted	bool	`json:"encrypted"`
}
