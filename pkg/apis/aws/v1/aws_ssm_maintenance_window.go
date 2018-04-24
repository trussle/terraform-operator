
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindow struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSsmMaintenanceWindowSpec	`json:"spec"`
}

type AwsSsmMaintenanceWindowList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSsmMaintenanceWindow	`json:"items"`
}

type AwsSsmMaintenanceWindowSpec struct {
	Name	string	`json:"name"`
	Schedule	string	`json:"schedule"`
	Duration	int	`json:"duration"`
	Cutoff	int	`json:"cutoff"`
	AllowUnassociatedTargets	bool	`json:"allow_unassociated_targets"`
	Enabled	bool	`json:"enabled"`
}
