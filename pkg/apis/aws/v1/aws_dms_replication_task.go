
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDmsReplicationTask struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsDmsReplicationTaskSpec	`json:"spec"`
}

type AwsDmsReplicationTaskList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsDmsReplicationTask	`json:"items"`
}

type AwsDmsReplicationTaskSpec struct {
	MigrationType	string	`json:"migration_type"`
	ReplicationTaskSettings	string	`json:"replication_task_settings"`
	SourceEndpointArn	string	`json:"source_endpoint_arn"`
	TableMappings	string	`json:"table_mappings"`
	Tags	map[string]interface{}	`json:"tags"`
	TargetEndpointArn	string	`json:"target_endpoint_arn"`
	CdcStartTime	string	`json:"cdc_start_time"`
	ReplicationInstanceArn	string	`json:"replication_instance_arn"`
	ReplicationTaskId	string	`json:"replication_task_id"`
}
