
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	WindowId	string	`json:"window_id"`
	ResourceType	string	`json:"resource_type"`
	Targets	[]Targets	`json:"targets"`
	OwnerInformation	string	`json:"owner_information"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTargetList is a list of AwsSsmMaintenanceWindowTarget resources
type AwsSsmMaintenanceWindowTargetList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindowTarget	`json:"items"`
}


// Targets is a Targets
type Targets struct {
	Key	string	`json:"key"`
	Values	[]string	`json:"values"`
}
