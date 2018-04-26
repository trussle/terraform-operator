
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroup describes a AwsAlbTargetGroup resource
type AwsAlbTargetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAlbTargetGroupSpec	`json:"spec"`
}


// AwsAlbTargetGroupSpec is the spec for a AwsAlbTargetGroup Resource
type AwsAlbTargetGroupSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	Protocol	string	`json:"protocol"`
	VpcId	string	`json:"vpc_id"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
	Tags	map[string]string	`json:"tags"`
	Port	int	`json:"port"`
	TargetType	string	`json:"target_type"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroupList is a list of AwsAlbTargetGroup resources
type AwsAlbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbTargetGroup	`json:"items"`
}


// Stickiness is a Stickiness
type Stickiness struct {
	Enabled	bool	`json:"enabled"`
	Type	string	`json:"type"`
	CookieDuration	int	`json:"cookie_duration"`
}

// HealthCheck is a HealthCheck
type HealthCheck struct {
	Protocol	string	`json:"protocol"`
	HealthyThreshold	int	`json:"healthy_threshold"`
	UnhealthyThreshold	int	`json:"unhealthy_threshold"`
	Interval	int	`json:"interval"`
	Port	string	`json:"port"`
}
