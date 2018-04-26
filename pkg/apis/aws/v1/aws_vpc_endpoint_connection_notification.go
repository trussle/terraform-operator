
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointConnectionNotification describes a AwsVpcEndpointConnectionNotification resource
type AwsVpcEndpointConnectionNotification struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointConnectionNotificationSpec	`json:"spec"`
}


// AwsVpcEndpointConnectionNotificationSpec is the spec for a AwsVpcEndpointConnectionNotification Resource
type AwsVpcEndpointConnectionNotificationSpec struct {
	VpcEndpointServiceId	string	`json:"vpc_endpoint_service_id"`
	VpcEndpointId	string	`json:"vpc_endpoint_id"`
	ConnectionNotificationArn	string	`json:"connection_notification_arn"`
	ConnectionEvents	string	`json:"connection_events"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsVpcEndpointConnectionNotificationList is a list of AwsVpcEndpointConnectionNotification resources
type AwsVpcEndpointConnectionNotificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointConnectionNotification	`json:"items"`
}

