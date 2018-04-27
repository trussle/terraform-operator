
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	RoutingStrategy	[]AwsGameliftAliasRoutingStrategy	`json:"routing_strategy"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsGameliftAliasList is a list of AwsGameliftAlias resources
type AwsGameliftAliasList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsGameliftAlias	`json:"items"`
}


// AwsGameliftAliasRoutingStrategy is a AwsGameliftAliasRoutingStrategy
type AwsGameliftAliasRoutingStrategy struct {
	FleetId	string	`json:"fleet_id"`
	Message	string	`json:"message"`
	Type	string	`json:"type"`
}
