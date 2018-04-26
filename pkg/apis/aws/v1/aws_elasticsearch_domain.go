
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
	Tags	map[string]???	`json:"tags"`
	DomainName	string	`json:"domain_name"`
	LogPublishingOptions	string	`json:"log_publishing_options"`
	SnapshotOptions	[]eoUqzMJG	`json:"snapshot_options"`
	VpcOptions	[]HOxYTsbf	`json:"vpc_options"`
	ElasticsearchVersion	string	`json:"elasticsearch_version"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticsearchDomainList is a list of AwsElasticsearchDomain resources
type AwsElasticsearchDomainList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsElasticsearchDomain	`json:"items"`
}


// tWdcUosn is a tWdcUosn
type tWdcUosn struct {
	EbsEnabled	bool	`json:"ebs_enabled"`
	Iops	int	`json:"iops"`
	VolumeSize	int	`json:"volume_size"`
}

// pMVHZdiy is a pMVHZdiy
type pMVHZdiy struct {
	Enabled	bool	`json:"enabled"`
}

// eoUqzMJG is a eoUqzMJG
type eoUqzMJG struct {
	AutomatedSnapshotStartHour	int	`json:"automated_snapshot_start_hour"`
}

// HOxYTsbf is a HOxYTsbf
type HOxYTsbf struct {
	SecurityGroupIds	string	`json:"security_group_ids"`
	SubnetIds	string	`json:"subnet_ids"`
}

// qAqtXjkV is a qAqtXjkV
type qAqtXjkV struct {
	InstanceType	string	`json:"instance_type"`
	ZoneAwarenessEnabled	bool	`json:"zone_awareness_enabled"`
	DedicatedMasterCount	int	`json:"dedicated_master_count"`
	DedicatedMasterEnabled	bool	`json:"dedicated_master_enabled"`
	DedicatedMasterType	string	`json:"dedicated_master_type"`
	InstanceCount	int	`json:"instance_count"`
}
