
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticsearchDomain struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsElasticsearchDomainSpec	`json:"spec"`
}

type AwsElasticsearchDomainList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticsearchDomain	`json:"items"`
}

type AwsElasticsearchDomainSpec struct {
	SnapshotOptions	[]interface{}	`json:"snapshot_options"`
	Tags	map[string]interface{}	`json:"tags"`
	DomainName	string	`json:"domain_name"`
	VpcOptions	[]interface{}	`json:"vpc_options"`
	LogPublishingOptions	interface{}	`json:"log_publishing_options"`
	ElasticsearchVersion	string	`json:"elasticsearch_version"`
}
