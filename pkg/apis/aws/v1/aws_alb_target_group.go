
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
	Protocol	string	`json:"protocol"`
	VpcId	string	`json:"vpc_id"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
	TargetType	string	`json:"target_type"`
	Tags	map[string]string	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	Port	int	`json:"port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroupList is a list of AwsAlbTargetGroup resources
type AwsAlbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbTargetGroup	`json:"items"`
}


// AwsAlbTargetGroupStickiness is a AwsAlbTargetGroupStickiness
type AwsAlbTargetGroupStickiness struct {
	Enabled	bool	`json:"enabled"`
	Type	string	`json:"type"`
	CookieDuration	int	`json:"cookie_duration"`
}

// AwsAlbTargetGroupHealthCheck is a AwsAlbTargetGroupHealthCheck
type AwsAlbTargetGroupHealthCheck struct {
	Port	string	`json:"port"`
	Protocol	string	`json:"protocol"`
	HealthyThreshold	int	`json:"healthy_threshold"`
	UnhealthyThreshold	int	`json:"unhealthy_threshold"`
	Interval	int	`json:"interval"`
}
