
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
	ThrottleSettings	ThrottleSettings	`json:"throttle_settings"`
	ProductCode	string	`json:"product_code"`
	Name	string	`json:"name"`
	Description	string	`json:"description"`
	ApiStages	[]ApiStages	`json:"api_stages"`
	QuotaSettings	QuotaSettings	`json:"quota_settings"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayUsagePlanList is a list of AwsApiGatewayUsagePlan resources
type AwsApiGatewayUsagePlanList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayUsagePlan	`json:"items"`
}


// ThrottleSettings is a ThrottleSettings
type ThrottleSettings struct {
	BurstLimit	int	`json:"burst_limit"`
	RateLimit	float64	`json:"rate_limit"`
}

// ApiStages is a ApiStages
type ApiStages struct {
	ApiId	string	`json:"api_id"`
	Stage	string	`json:"stage"`
}

// QuotaSettings is a QuotaSettings
type QuotaSettings struct {
	Limit	int	`json:"limit"`
	Offset	int	`json:"offset"`
	Period	string	`json:"period"`
}
