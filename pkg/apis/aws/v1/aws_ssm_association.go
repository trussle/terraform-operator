
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmAssociation describes a AwsSsmAssociation resource
type AwsSsmAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmAssociationSpec	`json:"spec"`
}


// AwsSsmAssociationSpec is the spec for a AwsSsmAssociation Resource
type AwsSsmAssociationSpec struct {
	InstanceId	string	`json:"instance_id"`
	ScheduleExpression	string	`json:"schedule_expression"`
	OutputLocation	[]interface{}	`json:"output_location"`
	AssociationName	string	`json:"association_name"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmAssociationList is a list of AwsSsmAssociation resources
type AwsSsmAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmAssociation	`json:"items"`
}

