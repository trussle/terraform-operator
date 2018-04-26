
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
	Protocol	string	`json:"protocol"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
	TargetType	string	`json:"target_type"`
	NamePrefix	string	`json:"name_prefix"`
	Tags	map[string]???	`json:"tags"`
	Port	int	`json:"port"`
	VpcId	string	`json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbTargetGroupList is a list of AwsLbTargetGroup resources
type AwsLbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbTargetGroup	`json:"items"`
}


// hUhSWZdg is a hUhSWZdg
type hUhSWZdg struct {
	CookieDuration	int	`json:"cookie_duration"`
	Enabled	bool	`json:"enabled"`
	Type	string	`json:"type"`
}

// pyWsRCkc is a pyWsRCkc
type pyWsRCkc struct {
	UnhealthyThreshold	int	`json:"unhealthy_threshold"`
	Interval	int	`json:"interval"`
	Port	string	`json:"port"`
	Protocol	string	`json:"protocol"`
	HealthyThreshold	int	`json:"healthy_threshold"`
}
