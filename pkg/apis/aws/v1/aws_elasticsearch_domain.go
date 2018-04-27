
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
	VpcOptions	[]AwsElasticsearchDomainVpcOptions	`json:"vpc_options"`
	LogPublishingOptions	AwsElasticsearchDomainLogPublishingOptions	`json:"log_publishing_options"`
	ElasticsearchVersion	string	`json:"elasticsearch_version"`
	DomainName	string	`json:"domain_name"`
	SnapshotOptions	[]AwsElasticsearchDomainSnapshotOptions	`json:"snapshot_options"`
	Tags	map[string]string	`json:"tags"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticsearchDomainList is a list of AwsElasticsearchDomain resources
type AwsElasticsearchDomainList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticsearchDomain	`json:"items"`
}


// AwsElasticsearchDomainSnapshotOptions is a AwsElasticsearchDomainSnapshotOptions
type AwsElasticsearchDomainSnapshotOptions struct {
	AutomatedSnapshotStartHour	int	`json:"automated_snapshot_start_hour"`
}

// AwsElasticsearchDomainVpcOptions is a AwsElasticsearchDomainVpcOptions
type AwsElasticsearchDomainVpcOptions struct {
	SecurityGroupIds	string	`json:"security_group_ids"`
	SubnetIds	string	`json:"subnet_ids"`
}

// AwsElasticsearchDomainLogPublishingOptions is a AwsElasticsearchDomainLogPublishingOptions
type AwsElasticsearchDomainLogPublishingOptions struct {
	CloudwatchLogGroupArn	string	`json:"cloudwatch_log_group_arn"`
	Enabled	bool	`json:"enabled"`
	LogType	string	`json:"log_type"`
}

// AwsElasticsearchDomainClusterConfig is a AwsElasticsearchDomainClusterConfig
type AwsElasticsearchDomainClusterConfig struct {
	InstanceType	string	`json:"instance_type"`
	ZoneAwarenessEnabled	bool	`json:"zone_awareness_enabled"`
	DedicatedMasterCount	int	`json:"dedicated_master_count"`
	DedicatedMasterEnabled	bool	`json:"dedicated_master_enabled"`
	DedicatedMasterType	string	`json:"dedicated_master_type"`
	InstanceCount	int	`json:"instance_count"`
}

// AwsElasticsearchDomainEbsOptions is a AwsElasticsearchDomainEbsOptions
type AwsElasticsearchDomainEbsOptions struct {
	EbsEnabled	bool	`json:"ebs_enabled"`
	Iops	int	`json:"iops"`
	VolumeSize	int	`json:"volume_size"`
}

// AwsElasticsearchDomainEncryptAtRest is a AwsElasticsearchDomainEncryptAtRest
type AwsElasticsearchDomainEncryptAtRest struct {
	Enabled	bool	`json:"enabled"`
}
