
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterInstance describes a AwsRdsClusterInstance resource
type AwsRdsClusterInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRdsClusterInstanceSpec	`json:"spec"`
}


// AwsRdsClusterInstanceSpec is the spec for a AwsRdsClusterInstance Resource
type AwsRdsClusterInstanceSpec struct {
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	Engine	string	`json:"engine"`
	ClusterIdentifier	string	`json:"cluster_identifier"`
	PromotionTier	int	`json:"promotion_tier"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	InstanceClass	string	`json:"instance_class"`
	Tags	map[string]Generic	`json:"tags"`
	MonitoringInterval	int	`json:"monitoring_interval"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterInstanceList is a list of AwsRdsClusterInstance resources
type AwsRdsClusterInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsClusterInstance	`json:"items"`
}

