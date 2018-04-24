
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsGameliftAlias struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsGameliftAliasSpec	`json:"spec"`
}

type AwsGameliftAliasList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftAlias	`json:"items"`
}

type AwsGameliftAliasSpec struct {
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	RoutingStrategy	[]interface{}	`json:"routing_strategy"`
}
