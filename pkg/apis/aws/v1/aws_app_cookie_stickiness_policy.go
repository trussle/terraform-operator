
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppCookieStickinessPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsAppCookieStickinessPolicySpec	`json:"spec"`
}

type AwsAppCookieStickinessPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsAppCookieStickinessPolicy	`json:"items"`
}

type AwsAppCookieStickinessPolicySpec struct {
	LoadBalancer	string	`json:"load_balancer"`
	LbPort	int	`json:"lb_port"`
	CookieName	string	`json:"cookie_name"`
	Name	string	`json:"name"`
}
