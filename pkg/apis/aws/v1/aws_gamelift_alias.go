
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftAlias describes a AwsGameliftAlias resource
type AwsGameliftAlias struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGameliftAliasSpec	`json:"spec"`
}


// AwsGameliftAliasSpec is the spec for a AwsGameliftAlias Resource
type AwsGameliftAliasSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	RoutingStrategy	[]Generic	`json:"routing_strategy"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftAliasList is a list of AwsGameliftAlias resources
type AwsGameliftAliasList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftAlias	`json:"items"`
}

