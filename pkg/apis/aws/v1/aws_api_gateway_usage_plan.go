
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayUsagePlan describes a AwsApiGatewayUsagePlan resource
type AwsApiGatewayUsagePlan struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayUsagePlanSpec	`json:"spec"`
}


// AwsApiGatewayUsagePlanSpec is the spec for a AwsApiGatewayUsagePlan Resource
type AwsApiGatewayUsagePlanSpec struct {
	Description	string	`json:"description"`
	ApiStages	[]wQpOQgNc	`json:"api_stages"`
	QuotaSettings	string	`json:"quota_settings"`
	ThrottleSettings	string	`json:"throttle_settings"`
	ProductCode	string	`json:"product_code"`
	Name	string	`json:"name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayUsagePlanList is a list of AwsApiGatewayUsagePlan resources
type AwsApiGatewayUsagePlanList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayUsagePlan	`json:"items"`
}


// wQpOQgNc is a wQpOQgNc
type wQpOQgNc struct {
	ApiId	string	`json:"api_id"`
	Stage	string	`json:"stage"`
}
