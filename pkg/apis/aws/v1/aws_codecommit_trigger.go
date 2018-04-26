
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodecommitTrigger describes a AwsCodecommitTrigger resource
type AwsCodecommitTrigger struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCodecommitTriggerSpec	`json:"spec"`
}


// AwsCodecommitTriggerSpec is the spec for a AwsCodecommitTrigger Resource
type AwsCodecommitTriggerSpec struct {
	Trigger	Trigger	`json:"trigger"`
	RepositoryName	string	`json:"repository_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodecommitTriggerList is a list of AwsCodecommitTrigger resources
type AwsCodecommitTriggerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodecommitTrigger	`json:"items"`
}


// Trigger is a Trigger
type Trigger struct {
	DestinationArn	string	`json:"destination_arn"`
	CustomData	string	`json:"custom_data"`
	Branches	[]string	`json:"branches"`
	Events	[]string	`json:"events"`
	Name	string	`json:"name"`
}
