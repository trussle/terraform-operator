
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationTask describes a AwsDmsReplicationTask resource
type AwsDmsReplicationTask struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsReplicationTaskSpec	`json:"spec"`
}


// AwsDmsReplicationTaskSpec is the spec for a AwsDmsReplicationTask Resource
type AwsDmsReplicationTaskSpec struct {
	ReplicationTaskSettings	string	`json:"replication_task_settings"`
	SourceEndpointArn	string	`json:"source_endpoint_arn"`
	TableMappings	string	`json:"table_mappings"`
	Tags	map[string]Generic	`json:"tags"`
	TargetEndpointArn	string	`json:"target_endpoint_arn"`
	CdcStartTime	string	`json:"cdc_start_time"`
	MigrationType	string	`json:"migration_type"`
	ReplicationInstanceArn	string	`json:"replication_instance_arn"`
	ReplicationTaskId	string	`json:"replication_task_id"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDmsReplicationTaskList is a list of AwsDmsReplicationTask resources
type AwsDmsReplicationTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationTask	`json:"items"`
}

