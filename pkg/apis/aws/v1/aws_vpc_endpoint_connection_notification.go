
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsVpcEndpointConnectionNotification struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsVpcEndpointConnectionNotificationSpec	`json:"spec"`
}

type AwsVpcEndpointConnectionNotificationList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsVpcEndpointConnectionNotification	`json:"items"`
}

type AwsVpcEndpointConnectionNotificationSpec struct {
	VpcEndpointServiceId	string	`json:"vpc_endpoint_service_id"`
	VpcEndpointId	string	`json:"vpc_endpoint_id"`
	ConnectionNotificationArn	string	`json:"connection_notification_arn"`
	ConnectionEvents	interface{}	`json:"connection_events"`
}
