
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayUsagePlanKey struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayUsagePlanKeySpec	`json:"spec"`
}

type AwsApiGatewayUsagePlanKeyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayUsagePlanKey	`json:"items"`
}

type AwsApiGatewayUsagePlanKeySpec struct {
	KeyType	string	`json:"key_type"`
	UsagePlanId	string	`json:"usage_plan_id"`
	KeyId	string	`json:"key_id"`
}
