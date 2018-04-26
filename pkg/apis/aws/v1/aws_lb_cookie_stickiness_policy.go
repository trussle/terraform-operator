
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbCookieStickinessPolicy describes a AwsLbCookieStickinessPolicy resource
type AwsLbCookieStickinessPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsLbCookieStickinessPolicySpec	`json:"spec"`
}


// AwsLbCookieStickinessPolicySpec is the spec for a AwsLbCookieStickinessPolicy Resource
type AwsLbCookieStickinessPolicySpec struct {
	CookieExpirationPeriod	int	`json:"cookie_expiration_period"`
	Name	string	`json:"name"`
	LoadBalancer	string	`json:"load_balancer"`
	LbPort	int	`json:"lb_port"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsLbCookieStickinessPolicyList is a list of AwsLbCookieStickinessPolicy resources
type AwsLbCookieStickinessPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsLbCookieStickinessPolicy	`json:"items"`
}

