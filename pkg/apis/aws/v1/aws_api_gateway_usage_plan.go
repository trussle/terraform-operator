
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
	QuotaSettings	AwsApiGatewayUsagePlanQuotaSettings	`json:"quota_settings"`
	ThrottleSettings	AwsApiGatewayUsagePlanThrottleSettings	`json:"throttle_settings"`
	ProductCode	string	`json:"product_code"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	ApiStages	[]AwsApiGatewayUsagePlanApiStages	`json:"api_stages"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayUsagePlanList is a list of AwsApiGatewayUsagePlan resources
type AwsApiGatewayUsagePlanList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayUsagePlan	`json:"items"`
}


// AwsApiGatewayUsagePlanQuotaSettings is a AwsApiGatewayUsagePlanQuotaSettings
type AwsApiGatewayUsagePlanQuotaSettings struct {
	Limit	int	`json:"limit"`
	Offset	int	`json:"offset"`
	Period	string	`json:"period"`
}

// AwsApiGatewayUsagePlanThrottleSettings is a AwsApiGatewayUsagePlanThrottleSettings
type AwsApiGatewayUsagePlanThrottleSettings struct {
	RateLimit	float64	`json:"rate_limit"`
	BurstLimit	int	`json:"burst_limit"`
}

// AwsApiGatewayUsagePlanApiStages is a AwsApiGatewayUsagePlanApiStages
type AwsApiGatewayUsagePlanApiStages struct {
	ApiId	string	`json:"api_id"`
	Stage	string	`json:"stage"`
}
