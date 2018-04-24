
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsRdsClusterInstance struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsRdsClusterInstanceSpec	`json:"spec"`
}

type AwsRdsClusterInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsClusterInstance	`json:"items"`
}

type AwsRdsClusterInstanceSpec struct {
	Engine	string	`json:"engine"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	PromotionTier	int	`json:"promotion_tier"`
	ClusterIdentifier	string	`json:"cluster_identifier"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	InstanceClass	string	`json:"instance_class"`
	MonitoringInterval	int	`json:"monitoring_interval"`
	Tags	map[string]interface{}	`json:"tags"`
}
