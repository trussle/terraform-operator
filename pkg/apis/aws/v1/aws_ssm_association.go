
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmAssociation struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmAssociationSpec	`json:"spec"`
}

type AwsSsmAssociationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmAssociation	`json:"items"`
}

type AwsSsmAssociationSpec struct {
	InstanceId	string	`json:"instance_id"`
	Name	string	`json:"name"`
	ScheduleExpression	string	`json:"schedule_expression"`
	OutputLocation	[]interface{}	`json:"output_location"`
	AssociationName	string	`json:"association_name"`
}
