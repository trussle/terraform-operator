
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayUsagePlan struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayUsagePlanSpec	`json:"spec"`
}

type AwsApiGatewayUsagePlanList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayUsagePlan	`json:"items"`
}

type AwsApiGatewayUsagePlanSpec struct {
	QuotaSettings	interface{}	`json:"quota_settings"`
	ThrottleSettings	interface{}	`json:"throttle_settings"`
	ProductCode	string	`json:"product_code"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	ApiStages	[]interface{}	`json:"api_stages"`
}
