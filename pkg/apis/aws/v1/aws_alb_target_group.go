
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
	Tags	map[string]???	`json:"tags"`
	NamePrefix	string	`json:"name_prefix"`
	Port	int	`json:"port"`
	DeregistrationDelay	int	`json:"deregistration_delay"`
	TargetType	string	`json:"target_type"`
	VpcId	string	`json:"vpc_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAlbTargetGroupList is a list of AwsAlbTargetGroup resources
type AwsAlbTargetGroupList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAlbTargetGroup	`json:"items"`
}


// zKaWlfVc is a zKaWlfVc
type zKaWlfVc struct {
	Protocol	string	`json:"protocol"`
	HealthyThreshold	int	`json:"healthy_threshold"`
	UnhealthyThreshold	int	`json:"unhealthy_threshold"`
	Interval	int	`json:"interval"`
	Port	string	`json:"port"`
}

// jYMcwVxD is a jYMcwVxD
type jYMcwVxD struct {
	Enabled	bool	`json:"enabled"`
	Type	string	`json:"type"`
	CookieDuration	int	`json:"cookie_duration"`
}
