
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticsearchDomainPolicy describes a AwsElasticsearchDomainPolicy resource
type AwsElasticsearchDomainPolicy struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticsearchDomainPolicySpec	`json:"spec"`
}


// AwsElasticsearchDomainPolicySpec is the spec for a AwsElasticsearchDomainPolicy Resource
type AwsElasticsearchDomainPolicySpec struct {
	DomainName	string	`json:"domain_name"`
	AccessPolicies	string	`json:"access_policies"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticsearchDomainPolicyList is a list of AwsElasticsearchDomainPolicy resources
type AwsElasticsearchDomainPolicyList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticsearchDomainPolicy	`json:"items"`
}

