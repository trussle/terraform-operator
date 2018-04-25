
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppCookieStickinessPolicy describes a AwsAppCookieStickinessPolicy resource
type AwsAppCookieStickinessPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppCookieStickinessPolicySpec	`json:"spec"`
}


// AwsAppCookieStickinessPolicySpec is the spec for a AwsAppCookieStickinessPolicy Resource
type AwsAppCookieStickinessPolicySpec struct {
	Name	string	`json:"name"`
	LoadBalancer	string	`json:"load_balancer"`
	LbPort	int	`json:"lb_port"`
	CookieName	string	`json:"cookie_name"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppCookieStickinessPolicyList is a list of AwsAppCookieStickinessPolicy resources
type AwsAppCookieStickinessPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppCookieStickinessPolicy	`json:"items"`
}

