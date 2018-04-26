
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
//	"github.com/hashicorp/terraform/helper/schema"
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
	Tags	map[string]string	`json:"tags"`
	LogPublishingOptions	LogPublishingOptions	`json:"log_publishing_options"`
	SnapshotOptions	[]SnapshotOptions	`json:"snapshot_options"`
	ElasticsearchVersion	string	`json:"elasticsearch_version"`
	DomainName	string	`json:"domain_name"`
	VpcOptions	[]VpcOptions	`json:"vpc_options"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticsearchDomainList is a list of AwsElasticsearchDomain resources
type AwsElasticsearchDomainList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticsearchDomain	`json:"items"`
}


// ClusterConfig is a ClusterConfig
type ClusterConfig struct {
	DedicatedMasterType	string	`json:"dedicated_master_type"`
	InstanceCount	int	`json:"instance_count"`
	InstanceType	string	`json:"instance_type"`
	ZoneAwarenessEnabled	bool	`json:"zone_awareness_enabled"`
	DedicatedMasterCount	int	`json:"dedicated_master_count"`
	DedicatedMasterEnabled	bool	`json:"dedicated_master_enabled"`
}

// LogPublishingOptions is a LogPublishingOptions
type LogPublishingOptions struct {
	LogType	string	`json:"log_type"`
	CloudwatchLogGroupArn	string	`json:"cloudwatch_log_group_arn"`
	Enabled	bool	`json:"enabled"`
}

// EbsOptions is a EbsOptions
type EbsOptions struct {
	EbsEnabled	bool	`json:"ebs_enabled"`
	Iops	int	`json:"iops"`
	VolumeSize	int	`json:"volume_size"`
}

// EncryptAtRest is a EncryptAtRest
type EncryptAtRest struct {
	Enabled	bool	`json:"enabled"`
}

// SnapshotOptions is a SnapshotOptions
type SnapshotOptions struct {
	AutomatedSnapshotStartHour	int	`json:"automated_snapshot_start_hour"`
}

// VpcOptions is a VpcOptions
type VpcOptions struct {
	SecurityGroupIds	string	`json:"security_group_ids"`
	SubnetIds	string	`json:"subnet_ids"`
}
