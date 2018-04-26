
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	MonitoringInterval	int	`json:"monitoring_interval"`
	PubliclyAccessible	bool	`json:"publicly_accessible"`
	InstanceClass	string	`json:"instance_class"`
	AutoMinorVersionUpgrade	bool	`json:"auto_minor_version_upgrade"`
	ClusterIdentifier	string	`json:"cluster_identifier"`
	Tags	map[string]???	`json:"tags"`
	Engine	string	`json:"engine"`
	PromotionTier	int	`json:"promotion_tier"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsRdsClusterInstanceList is a list of AwsRdsClusterInstance resources
type AwsRdsClusterInstanceList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsRdsClusterInstance	`json:"items"`
}

