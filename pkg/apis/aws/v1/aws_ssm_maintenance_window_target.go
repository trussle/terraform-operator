
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTarget describes a AwsSsmMaintenanceWindowTarget resource
type AwsSsmMaintenanceWindowTarget struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmMaintenanceWindowTargetSpec	`json:"spec"`
}


// AwsSsmMaintenanceWindowTargetSpec is the spec for a AwsSsmMaintenanceWindowTarget Resource
type AwsSsmMaintenanceWindowTargetSpec struct {
	OwnerInformation	string	`json:"owner_information"`
	WindowId	string	`json:"window_id"`
	ResourceType	string	`json:"resource_type"`
	Targets	[]Generic	`json:"targets"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTargetList is a list of AwsSsmMaintenanceWindowTarget resources
type AwsSsmMaintenanceWindowTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindowTarget	`json:"items"`
}

