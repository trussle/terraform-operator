
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDirectoryServiceConditionalForwarder struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDirectoryServiceConditionalForwarderSpec	`json:"spec"`
}

type AwsDirectoryServiceConditionalForwarderList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDirectoryServiceConditionalForwarder	`json:"items"`
}

type AwsDirectoryServiceConditionalForwarderSpec struct {
	DnsIps	[]interface{}	`json:"dns_ips"`
	RemoteDomainName	string	`json:"remote_domain_name"`
	DirectoryId	string	`json:"directory_id"`
}
