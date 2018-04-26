
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayUsagePlanKey describes a AwsApiGatewayUsagePlanKey resource
type AwsApiGatewayUsagePlanKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayUsagePlanKeySpec	`json:"spec"`
}


// AwsApiGatewayUsagePlanKeySpec is the spec for a AwsApiGatewayUsagePlanKey Resource
type AwsApiGatewayUsagePlanKeySpec struct {
	KeyId	string	`json:"key_id"`
	KeyType	string	`json:"key_type"`
	UsagePlanId	string	`json:"usage_plan_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayUsagePlanKeyList is a list of AwsApiGatewayUsagePlanKey resources
type AwsApiGatewayUsagePlanKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayUsagePlanKey	`json:"items"`
}

