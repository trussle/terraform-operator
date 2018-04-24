
package v1

import (
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSqsQueue struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Spec	AwsSqsQueueSpec	`json:"spec"`
}

type AwsSqsQueueList struct {
	meta_v1.TypeMeta	`json:",inline"`
	meta_v1.ObjectMeta	`json:"metadata,omitempty"`

	Items	[]AwsSqsQueue	`json:"items"`
}

type AwsSqsQueueSpec struct {
	Tags	map[string]interface{}	`json:"tags"`
	FifoQueue	bool	`json:"fifo_queue"`
	NamePrefix	string	`json:"name_prefix"`
	DelaySeconds	int	`json:"delay_seconds"`
	ReceiveWaitTimeSeconds	int	`json:"receive_wait_time_seconds"`
	ContentBasedDeduplication	bool	`json:"content_based_deduplication"`
	KmsMasterKeyId	string	`json:"kms_master_key_id"`
	MaxMessageSize	int	`json:"max_message_size"`
	MessageRetentionSeconds	int	`json:"message_retention_seconds"`
	VisibilityTimeoutSeconds	int	`json:"visibility_timeout_seconds"`
	RedrivePolicy	string	`json:"redrive_policy"`
}
