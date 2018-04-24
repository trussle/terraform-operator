
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDefaultVpc struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDefaultVpcSpec	`json:"spec"`
}

type AwsDefaultVpcList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDefaultVpc	`json:"items"`
}

type AwsDefaultVpcSpec struct {
	EnableDnsSupport	bool	`json:"enable_dns_support"`
	Tags	map[string]interface{}	`json:"tags"`
}
