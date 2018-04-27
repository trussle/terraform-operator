
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElb describes a AwsElb resource
type AwsElb struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElbSpec	`json:"spec"`
}


// AwsElbSpec is the spec for a AwsElb Resource
type AwsElbSpec struct {
	CrossZoneLoadBalancing	bool	`json:"cross_zone_load_balancing"`
	ConnectionDrainingTimeout	int	`json:"connection_draining_timeout"`
	IdleTimeout	int	`json:"idle_timeout"`
	ConnectionDraining	bool	`json:"connection_draining"`
	Listener	AwsElbListener	`json:"listener"`
	Tags	map[string]string	`json:"tags"`
	AccessLogs	[]AwsElbAccessLogs	`json:"access_logs"`
	NamePrefix	string	`json:"name_prefix"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElbList is a list of AwsElb resources
type AwsElbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElb	`json:"items"`
}


// AwsElbAccessLogs is a AwsElbAccessLogs
type AwsElbAccessLogs struct {
	Interval	int	`json:"interval"`
	Bucket	string	`json:"bucket"`
	BucketPrefix	string	`json:"bucket_prefix"`
	Enabled	bool	`json:"enabled"`
}

// AwsElbListener is a AwsElbListener
type AwsElbListener struct {
	InstancePort	int	`json:"instance_port"`
	InstanceProtocol	string	`json:"instance_protocol"`
	LbPort	int	`json:"lb_port"`
	LbProtocol	string	`json:"lb_protocol"`
	SslCertificateId	string	`json:"ssl_certificate_id"`
}

// AwsElbHealthCheck is a AwsElbHealthCheck
type AwsElbHealthCheck struct {
	HealthyThreshold	int	`json:"healthy_threshold"`
	UnhealthyThreshold	int	`json:"unhealthy_threshold"`
	Target	string	`json:"target"`
	Interval	int	`json:"interval"`
	Timeout	int	`json:"timeout"`
}
