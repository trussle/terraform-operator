
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
	ConnectionDrainingTimeout	int	`json:"connection_draining_timeout"`
	ConnectionDraining	bool	`json:"connection_draining"`
	AccessLogs	[]AccessLogs	`json:"access_logs"`
	Listener	Listener	`json:"listener"`
	NamePrefix	string	`json:"name_prefix"`
	CrossZoneLoadBalancing	bool	`json:"cross_zone_load_balancing"`
	IdleTimeout	int	`json:"idle_timeout"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElbList is a list of AwsElb resources
type AwsElbList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElb	`json:"items"`
}


// AccessLogs is a AccessLogs
type AccessLogs struct {
	Interval	int	`json:"interval"`
	Bucket	string	`json:"bucket"`
	BucketPrefix	string	`json:"bucket_prefix"`
	Enabled	bool	`json:"enabled"`
}

// Listener is a Listener
type Listener struct {
	InstancePort	int	`json:"instance_port"`
	InstanceProtocol	string	`json:"instance_protocol"`
	LbPort	int	`json:"lb_port"`
	LbProtocol	string	`json:"lb_protocol"`
	SslCertificateId	string	`json:"ssl_certificate_id"`
}

// HealthCheck is a HealthCheck
type HealthCheck struct {
	HealthyThreshold	int	`json:"healthy_threshold"`
	UnhealthyThreshold	int	`json:"unhealthy_threshold"`
	Target	string	`json:"target"`
	Interval	int	`json:"interval"`
	Timeout	int	`json:"timeout"`
}
