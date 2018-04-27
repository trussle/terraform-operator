
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
	RepositoryName	string	`json:"repository_name"`
	Trigger	AwsCodecommitTriggerTrigger	`json:"trigger"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCodecommitTriggerList is a list of AwsCodecommitTrigger resources
type AwsCodecommitTriggerList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCodecommitTrigger	`json:"items"`
}


// AwsCodecommitTriggerTrigger is a AwsCodecommitTriggerTrigger
type AwsCodecommitTriggerTrigger struct {
	Events	[]string	`json:"events"`
	Name	string	`json:"name"`
	DestinationArn	string	`json:"destination_arn"`
	CustomData	string	`json:"custom_data"`
	Branches	[]string	`json:"branches"`
}
