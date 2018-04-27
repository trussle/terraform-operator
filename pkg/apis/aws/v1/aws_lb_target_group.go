
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroup describes a AwsLbTargetGroup resource
type AwsLbTargetGroup struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbTargetGroupSpec	`json:"spec"`
}


// AwsLbTargetGroupSpec is the spec for a AwsLbTargetGroup Resource
type AwsLbTargetGroupSpec struct {
	NamePrefix	string	`json:"name_prefix"`
	Port	int	`json:"port"`
	Protocol	string	`json:"protocol"`
	VpcId	string	`json:"vpc_id"`
	TargetType	string	`json:"target_type"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroupList is a list of AwsLbTargetGroup resources
type AwsLbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbTargetGroup	`json:"items"`
}


// AwsLbTargetGroupHealthCheck is a AwsLbTargetGroupHealthCheck
type AwsLbTargetGroupHealthCheck struct {
	UnhealthyThreshold	int	`json:"unhealthy_threshold"`
	Interval	int	`json:"interval"`
	Port	string	`json:"port"`
	Protocol	string	`json:"protocol"`
	HealthyThreshold	int	`json:"healthy_threshold"`
}

// AwsLbTargetGroupStickiness is a AwsLbTargetGroupStickiness
type AwsLbTargetGroupStickiness struct {
	Enabled	bool	`json:"enabled"`
	Type	string	`json:"type"`
	CookieDuration	int	`json:"cookie_duration"`
}
