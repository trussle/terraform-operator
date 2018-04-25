
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchDashboard describes a AwsCloudwatchDashboard resource
type AwsCloudwatchDashboard struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsCloudwatchDashboardSpec	`json:"spec"`
}


// AwsCloudwatchDashboardSpec is the spec for a AwsCloudwatchDashboard Resource
type AwsCloudwatchDashboardSpec struct {
	DashboardBody	string	`json:"dashboard_body"`
	DashboardName	string	`json:"dashboard_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsCloudwatchDashboardList is a list of AwsCloudwatchDashboard resources
type AwsCloudwatchDashboardList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsCloudwatchDashboard	`json:"items"`
}

