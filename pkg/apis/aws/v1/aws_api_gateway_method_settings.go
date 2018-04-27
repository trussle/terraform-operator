
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethodSettings describes a AwsApiGatewayMethodSettings resource
type AwsApiGatewayMethodSettings struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsApiGatewayMethodSettingsSpec	`json:"spec"`
}


// AwsApiGatewayMethodSettingsSpec is the spec for a AwsApiGatewayMethodSettings Resource
type AwsApiGatewayMethodSettingsSpec struct {
	RestApiId	string	`json:"rest_api_id"`
	StageName	string	`json:"stage_name"`
	MethodPath	string	`json:"method_path"`
	Settings	[]AwsApiGatewayMethodSettingsSettings	`json:"settings"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayMethodSettingsList is a list of AwsApiGatewayMethodSettings resources
type AwsApiGatewayMethodSettingsList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsApiGatewayMethodSettings	`json:"items"`
}


// AwsApiGatewayMethodSettingsSettings is a AwsApiGatewayMethodSettingsSettings
type AwsApiGatewayMethodSettingsSettings struct {
	ThrottlingRateLimit	float64	`json:"throttling_rate_limit"`
	CachingEnabled	bool	`json:"caching_enabled"`
	CacheDataEncrypted	bool	`json:"cache_data_encrypted"`
	RequireAuthorizationForCacheControl	bool	`json:"require_authorization_for_cache_control"`
	MetricsEnabled	bool	`json:"metrics_enabled"`
	LoggingLevel	string	`json:"logging_level"`
	DataTraceEnabled	bool	`json:"data_trace_enabled"`
	ThrottlingBurstLimit	int	`json:"throttling_burst_limit"`
	CacheTtlInSeconds	int	`json:"cache_ttl_in_seconds"`
	UnauthorizedCacheControlHeaderStrategy	string	`json:"unauthorized_cache_control_header_strategy"`
}
