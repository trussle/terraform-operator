
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticsearchDomain describes a AwsElasticsearchDomain resource
type AwsElasticsearchDomain struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticsearchDomainSpec	`json:"spec"`
}


// AwsElasticsearchDomainSpec is the spec for a AwsElasticsearchDomain Resource
type AwsElasticsearchDomainSpec struct {
	ElasticsearchVersion	string	`json:"elasticsearch_version"`
	LogPublishingOptions	Generic	`json:"log_publishing_options"`
	Tags	map[string]Generic	`json:"tags"`
	DomainName	string	`json:"domain_name"`
	SnapshotOptions	[]Generic	`json:"snapshot_options"`
	VpcOptions	[]Generic	`json:"vpc_options"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticsearchDomainList is a list of AwsElasticsearchDomain resources
type AwsElasticsearchDomainList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticsearchDomain	`json:"items"`
}

